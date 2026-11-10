package throttle

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	alierrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/ktesting"
)

var ErrThrottling error = alierrors.NewServerError(400, `{"Code": "Throttling"}`, "")

func TestV1Classifier(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		want      Class
		wantReqID string
	}{
		{name: "nil", err: nil, want: ClassOK},
		{name: "throttling", err: alierrors.NewServerError(400, `{"Code": "Throttling", "RequestId": "req-123"}`, ""), want: ClassThrottling, wantReqID: "req-123"},
		{name: "other server error", err: alierrors.NewServerError(400, `{"Code": "Test", "RequestId": "req-456"}`, ""), want: ClassOK, wantReqID: "req-456"},
		{name: "unknown error", err: errors.New("boom"), want: ClassUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			class, reqID := V1Classifier(tt.err)
			assert.Equal(t, tt.want, class)
			assert.Equal(t, tt.wantReqID, reqID)
		})
	}
}

func TestV2Classifier(t *testing.T) {
	tests := []struct {
		name      string
		err       error
		want      Class
		wantReqID string
	}{
		{name: "nil", err: nil, want: ClassOK},
		{name: "throttling", err: &tea.SDKError{Code: tea.String("Throttling"), Data: tea.String(`{"RequestId":"req-123"}`)}, want: ClassThrottling, wantReqID: "req-123"},
		{name: "throttling user prefix", err: &tea.SDKError{Code: tea.String("Throttling.User")}, want: ClassThrottling},
		{name: "throttling api prefix", err: &tea.SDKError{Code: tea.String("Throttling.Api")}, want: ClassThrottling},
		{name: "other sdk error", err: &tea.SDKError{Code: tea.String("InvalidParameter"), Data: tea.String(`{"RequestId":"req-789"}`)}, want: ClassOK, wantReqID: "req-789"},
		{name: "non-json data", err: &tea.SDKError{Code: tea.String("Throttling"), Data: tea.String("not json")}, want: ClassThrottling},
		{name: "unknown error", err: errors.New("boom"), want: ClassUnknown},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			class, reqID := V2Classifier(tt.err)
			assert.Equal(t, tt.want, class)
			assert.Equal(t, tt.wantReqID, reqID)
		})
	}
}

type controlledAPI struct {
	calls   chan chan error
	results chan error
}

func newControlledAPI(requests int) *controlledAPI {
	return &controlledAPI{
		calls:   make(chan chan error, requests),
		results: make(chan error, requests),
	}
}

func (a *controlledAPI) start(ctx context.Context, throttler *Throttler, requests int) {
	for range requests {
		go func() {
			a.results <- throttler.Throttle(ctx, func() error {
				response := make(chan error, 1)
				a.calls <- response
				select {
				case err := <-response:
					return err
				case <-ctx.Done():
					return ctx.Err()
				}
			})
		}()
	}
}

func (a *controlledAPI) finish(requests int, err error) {
	for range requests {
		(<-a.calls) <- err
	}
}

func startRecovery(t *testing.T, requests int) *controlledAPI {
	t.Helper()
	throttler := NewThrottler(time.Second, 8*time.Second, V1Classifier)
	api := newControlledAPI(requests)
	api.start(t.Context(), throttler, 1)
	synctest.Wait()
	api.finish(1, ErrThrottling)
	synctest.Wait()
	api.start(t.Context(), throttler, requests-1)
	synctest.Wait()
	require.Empty(t, api.calls)
	time.Sleep(time.Second)
	synctest.Wait()
	require.Len(t, api.calls, 1)
	api.finish(1, nil)
	synctest.Wait()
	require.Len(t, api.calls, 16)
	return api
}

func TestAdmissionWindowCap(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		throttler := NewThrottler(time.Second, 8*time.Second, V1Classifier)
		api := newControlledAPI(1024)
		api.start(t.Context(), throttler, 1024)
		synctest.Wait()
		require.Len(t, api.calls, 512)

		api.finish(1, nil)
		synctest.Wait()
		assert.Len(t, api.calls, 512)
		assert.NoError(t, <-api.results)
	})
}

func TestRecoveryWindowGrowth(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := startRecovery(t, 2048)
		for _, window := range []int{16, 32, 64, 128, 256} {
			api.finish(window, nil)
			synctest.Wait()
			require.Len(t, api.calls, window*2)
		}
		api.finish(1, nil)
		synctest.Wait()
		assert.Len(t, api.calls, 512)
	})
}

func TestRecoverySecondQuota(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := startRecovery(t, 1000)
		// The probe consumes one request of the 50-request quota.
		for range 49 {
			api.finish(1, nil)
			synctest.Wait()
		}
		require.Len(t, api.calls, 65)
		require.Len(t, api.results, 50)
		// Only 115 attempts (50 accepted + 65 outstanding) were issued,
		// rather than retrying the entire backlog in this quota period.
		api.finish(65, ErrThrottling)
		synctest.Wait()
		require.Empty(t, api.calls)
	})
}

func TestRecoveryCompletionClassification(t *testing.T) {
	tests := []struct {
		name   string
		err    error
		window int
	}{
		{name: "accepted cloud error", err: alierrors.NewServerError(400, `{"Code":"InvalidParameter"}`, ""), window: 17},
		{name: "unknown error", err: errors.New("network failure"), window: 16},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				api := startRecovery(t, 100)
				assert.NoError(t, <-api.results)
				api.finish(1, tt.err)
				synctest.Wait()
				assert.Len(t, api.calls, tt.window)
				assert.Same(t, tt.err, <-api.results)
			})
		})
	}
}

func TestStaleRecoveryCompletions(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		api := startRecovery(t, 100)
		var recovery [16]chan error
		for i := range recovery {
			recovery[i] = <-api.calls
		}

		recovery[0] <- ErrThrottling
		synctest.Wait()
		recovery[1] <- nil
		recovery[2] <- ErrThrottling
		synctest.Wait()
		require.Empty(t, api.calls)

		// Old failures do not extend backoff, and old successes cannot end it.
		time.Sleep(time.Second)
		synctest.Wait()
		require.Len(t, api.calls, 1)
		api.finish(1, nil)
		synctest.Wait()
		// Thirteen old calls still occupy slots in the new window.
		require.Len(t, api.calls, 3)

		recovery[3] <- nil
		synctest.Wait()
		require.Len(t, api.calls, 4)
		recovery[4] <- ErrThrottling
		synctest.Wait()
		assert.Len(t, api.calls, 5)
	})
}

func TestProbeBackoffDelay(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		throttler := NewThrottler(time.Second, 8*time.Second, V1Classifier)
		api := newControlledAPI(2)
		api.start(t.Context(), throttler, 2)
		synctest.Wait()
		require.Len(t, api.calls, 2)
		api.finish(2, ErrThrottling)
		synctest.Wait()

		for _, delay := range []time.Duration{time.Second, 2 * time.Second, 4 * time.Second, 8 * time.Second, 8 * time.Second} {
			time.Sleep(delay - time.Millisecond)
			synctest.Wait()
			require.Empty(t, api.calls)
			time.Sleep(time.Millisecond)
			synctest.Wait()
			require.Len(t, api.calls, 1)
			api.finish(1, ErrThrottling)
			synctest.Wait()
		}
	})
}

func TestUnknownErrorOnProbing(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		throttler := NewThrottler(time.Second, 8*time.Second, V1Classifier)
		api := newControlledAPI(1)
		api.start(t.Context(), throttler, 1)
		synctest.Wait()
		api.finish(1, ErrThrottling)
		synctest.Wait()
		time.Sleep(time.Second)
		synctest.Wait()
		require.Len(t, api.calls, 1)
		errUnknown := errors.New("unknown error")
		api.finish(1, errUnknown)
		synctest.Wait()
		assert.Same(t, errUnknown, <-api.results)

		// An unknown response releases ownership for another immediate probe.
		api.start(t.Context(), throttler, 1)
		synctest.Wait()
		require.Len(t, api.calls, 1)
		api.finish(1, nil)
		synctest.Wait()
		assert.NoError(t, <-api.results)
	})
}

func TestCanceledAdmissionHandoff(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		throttler := NewThrottler(time.Second, 8*time.Second, V1Classifier)
		throttler.window = 1
		ticket, err := throttler.acquire(t.Context())
		require.NoError(t, err)
		throttler.complete(t.Context(), ticket, ClassUnknown, "")
		require.Len(t, throttler.ready, 1)

		// A waiter receives the notification, then notices cancellation.
		<-throttler.ready
		ctx, cancel := context.WithCancel(t.Context())
		cancel()
		_, err = throttler.acquire(ctx)
		assert.ErrorIs(t, err, context.Canceled)
		assert.Len(t, throttler.ready, 1)
	})
}

func TestStaleReadyNotification(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		throttler := NewThrottler(time.Second, 8*time.Second, V1Classifier)
		throttler.window = 1
		ticket, err := throttler.acquire(t.Context())
		require.NoError(t, err)
		require.Len(t, throttler.ready, 1)

		api := newControlledAPI(1)
		api.start(t.Context(), throttler, 1)
		synctest.Wait()
		require.Empty(t, api.calls)
		require.Empty(t, throttler.ready)

		throttler.complete(t.Context(), ticket, ClassUnknown, "")
		synctest.Wait()
		require.Len(t, api.calls, 1)
		api.finish(1, nil)
		synctest.Wait()
		assert.NoError(t, <-api.results)
	})
}

func TestCanceledProbeHandoff(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		throttler := NewThrottler(time.Second, 8*time.Second, V1Classifier)
		ownerCtx, cancelOwner := context.WithCancel(t.Context())
		defer cancelOwner()
		ownerResult := make(chan error, 1)
		go func() {
			ownerResult <- throttler.Throttle(ownerCtx, func() error { return ErrThrottling })
		}()
		synctest.Wait()

		api := newControlledAPI(1)
		api.start(t.Context(), throttler, 1)
		synctest.Wait()
		require.Empty(t, api.calls)
		time.Sleep(500 * time.Millisecond)
		cancelOwner()
		synctest.Wait()
		assert.ErrorIs(t, <-ownerResult, context.Canceled)
		require.Empty(t, api.calls)

		// The surviving waiter inherits the original probe deadline.
		time.Sleep(500 * time.Millisecond)
		synctest.Wait()
		require.Len(t, api.calls, 1)
		api.finish(1, nil)
		synctest.Wait()
		assert.NoError(t, <-api.results)
	})
}

func TestCancelAtAdmissionWindow(t *testing.T) {
	synctest.Test(t, func(t *testing.T) {
		throttler := NewThrottler(time.Second, 8*time.Second, V1Classifier)
		api := newControlledAPI(512)
		api.start(t.Context(), throttler, 512)
		synctest.Wait()
		require.Len(t, api.calls, 512)

		waitingCtx, cancelWaiting := context.WithCancel(t.Context())
		result := make(chan error, 1)
		go func() {
			result <- throttler.Throttle(waitingCtx, func() error {
				t.Error("canceled waiter must not call the API")
				return nil
			})
		}()
		synctest.Wait()
		cancelWaiting()
		synctest.Wait()
		assert.ErrorIs(t, <-result, context.Canceled)

		api.finish(1, nil)
		synctest.Wait()
		assert.NoError(t, throttler.Throttle(t.Context(), func() error { return nil }))
	})
}

func TestThrottlingStress(t *testing.T) { synctest.Test(t, testThrottlingStressSync) }
func testThrottlingStressSync(t *testing.T) {
	logger, ctx := ktesting.NewTestContext(t)
	ctx, cancel := context.WithCancel(ctx)

	var throttling atomic.Bool
	var reqCount [16]atomic.Uint64
	var apiCount atomic.Uint64
	f := func() error {
		apiCount.Add(1)
		time.Sleep(10 * time.Millisecond)
		if throttling.Load() {
			return ErrThrottling
		} else {
			return nil
		}
	}
	throttler := NewThrottler(1*time.Second, 8*time.Second, V1Classifier)

	for i := range reqCount {
		go func() {
			ctx := klog.NewContext(ctx, logger.WithValues("worker", i))
			for {
				reqCount[i].Add(1)
				err := throttler.Throttle(ctx, f)
				if err != nil {
					// Throttling error will not be exposed to the caller
					assert.ErrorIs(t, err, context.Canceled)
					break
				}
			}
		}()
	}

	allHasReq := func(count uint64) bool {
		for i := range reqCount {
			if reqCount[i].Load() < count {
				return false
			}
		}
		return true
	}
	totalCount := func() uint64 {
		var c uint64
		for i := range reqCount {
			c += reqCount[i].Load()
		}
		return c
	}

	time.Sleep(11 * time.Second)
	assert.True(t, allHasReq(100))

	throttling.Store(true)
	time.Sleep(time.Millisecond * 20) // wait for all workers to finish current requests

	c1 := totalCount()
	a1 := apiCount.Load()
	time.Sleep(16 * time.Second)
	c2 := totalCount()
	a2 := apiCount.Load()
	// Expect 4 probe, but not return to caller
	assert.Equal(t, c1, c2)
	assert.Equal(t, 4, int(a2-a1))

	throttling.Store(false)
	time.Sleep(20 * time.Second)

	assert.True(t, allHasReq(120))
	cancel()
	time.Sleep(time.Millisecond * 20) // wait for all workers to finish current requests
}
