package throttle

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"testing/synctest"
	"time"

	alierrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/stretchr/testify/assert"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/ktesting"
	"k8s.io/utils/clock"
	testclock "k8s.io/utils/clock/testing"
)

var ErrThrottling error = alierrors.NewServerError(400, `{"Code": "Throttling"}`, "")

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
	throttler := NewThrottler(clock.RealClock{}, 1*time.Second, 8*time.Second)

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

func TestThrottleErrorPassThrough(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)
	clk := testclock.NewFakeClock(time.Now())
	throttler := NewThrottler(clk, time.Second*1, time.Second*10)

	testErr := func(t *testing.T, expectedErr error) {
		f := func() error {
			return expectedErr
		}
		err := throttler.Throttle(ctx, f)
		assert.Same(t, expectedErr, err)
	}

	t.Run("unknown", func(t *testing.T) {
		testErr(t, errors.New("unknown error"))
	})
	t.Run("server error", func(t *testing.T) {
		testErr(t, alierrors.NewServerError(400, `{"Code": "Test"}`, ""))
	})
}

func TestCancelAtDelay(t *testing.T) { synctest.Test(t, testCancelAtDelaySync) }
func testCancelAtDelaySync(t *testing.T) {
	throttler := NewThrottler(clock.RealClock{}, time.Second*1, time.Second*10)

	_, ctx := ktesting.NewTestContext(t)
	ctxToCancel, cancel := context.WithCancel(ctx)
	for range 2 {
		// One should block at probing delay, another should block on reading tCtx.probing
		go func() {
			err := throttler.Throttle(ctxToCancel, func() error {
				return ErrThrottling
			})
			assert.ErrorIs(t, err, context.Canceled)
		}()
	}
	time.Sleep(time.Millisecond * 10)
	cancel()

	// Another request should still go through, canceling should not break throttler state.
	err := throttler.Throttle(ctx, func() error { return nil })
	assert.NoError(t, err)
}

func TestUnknownErrorOnProbing(t *testing.T) { synctest.Test(t, testUnknownErrorOnProbingSync) }
func testUnknownErrorOnProbingSync(t *testing.T) {
	_, ctx := ktesting.NewTestContext(t)
	throttler := NewThrottler(clock.RealClock{}, time.Second*1, time.Second*10)

	errUnknown := errors.New("unknown error")
	var errRet atomic.Pointer[error]
	errRet.Store(&ErrThrottling)
	go func() {
		err := throttler.Throttle(ctx, func() error { return *errRet.Load() })
		assert.Equal(t, err, errUnknown)
	}()

	time.Sleep(time.Millisecond * 10)
	errRet.Store(&errUnknown)
	time.Sleep(time.Second)

	// initiate next probe immediately
	err := throttler.Throttle(ctx, func() error { return nil })
	assert.NoError(t, err)
}
