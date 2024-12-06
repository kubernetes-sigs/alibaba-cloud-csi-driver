package throttle

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	alierrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/stretchr/testify/assert"
	"k8s.io/klog/v2"
	klogtest "k8s.io/klog/v2/test"
	testclock "k8s.io/utils/clock/testing"
)

var ErrThrottling error = alierrors.NewServerError(400, `{"Code": "Throttling"}`, "")

func TestThrottlingStress(t *testing.T) {
	klogtest.InitKlog(t).Set("logtostderr", "true")

	var throttling atomic.Bool
	var reqCount [16]atomic.Uint64
	var apiCount atomic.Uint64
	f := func() error {
		apiCount.Add(1)
		if throttling.Load() {
			return ErrThrottling
		} else {
			return nil
		}
	}

	clk := testclock.NewFakeClock(time.Now())
	throttler := NewThrottler(clk, 1*time.Second, 8*time.Second)
	ctx, cancel := context.WithCancel(context.Background())

	wg := sync.WaitGroup{}
	wg.Add(len(reqCount))
	for i := range reqCount {
		go func() {
			ctx := klog.NewContext(ctx, klog.Background().WithValues("worker", i))
			for {
				reqCount[i].Add(1)
				err := throttler.Throttle(ctx, f)
				if err != nil {
					// Throttling error will not be exposed to the caller
					assert.ErrorIs(t, err, context.Canceled)
					break
				}
			}
			wg.Done()
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

	for !allHasReq(16) {
		time.Sleep(time.Millisecond)
	}

	throttling.Store(true)
	time.Sleep(time.Millisecond * 10)

	c1 := totalCount()
	a1 := apiCount.Load()
	for range 16 {
		clk.Step(time.Second)
		time.Sleep(time.Millisecond * 10)
	}
	c2 := totalCount()
	a2 := apiCount.Load()
	// Expect 4 probe, but not return to caller
	assert.Equal(t, c1, c2)
	assert.Equal(t, 4, int(a2-a1))

	throttling.Store(false)
	clk.Step(time.Second * 20)

	for !allHasReq(c2 + 32) {
		time.Sleep(time.Millisecond)
	}

	cancel()
	wg.Wait()
}

func TestThrottleErrorPassThrough(t *testing.T) {
	clk := testclock.NewFakeClock(time.Now())
	throttler := NewThrottler(clk, time.Second*1, time.Second*10)

	testErr := func(t *testing.T, expectedErr error) {
		f := func() error {
			return expectedErr
		}
		err := throttler.Throttle(context.Background(), f)
		assert.Same(t, expectedErr, err)
	}

	t.Run("unknown", func(t *testing.T) {
		testErr(t, errors.New("unknown error"))
	})
	t.Run("server error", func(t *testing.T) {
		testErr(t, alierrors.NewServerError(400, `{"Code": "Test"}`, ""))
	})
}

func TestCancelAtDelay(t *testing.T) {
	clk := testclock.NewFakeClock(time.Now())
	throttler := NewThrottler(clk, time.Second*1, time.Second*10)

	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	wg.Add(2)
	for range 2 {
		// One should block at probing delay, another should block on reading tCtx.probing
		go func() {
			err := throttler.Throttle(ctx, func() error {
				return ErrThrottling
			})
			assert.ErrorIs(t, err, context.Canceled)
			wg.Done()
		}()
	}
	time.Sleep(time.Millisecond * 10)
	cancel()
	wg.Wait()

	go func() {
		time.Sleep(time.Millisecond * 10)
		clk.Step(time.Second)
	}()
	// Another request should still go through, canceling should not break throttler state.
	err := throttler.Throttle(context.Background(), func() error { return nil })
	assert.NoError(t, err)
}

func TestUnknownErrorOnProbing(t *testing.T) {
	clk := testclock.NewFakeClock(time.Now())
	throttler := NewThrottler(clk, time.Second*1, time.Second*10)

	errUnknown := errors.New("unknown error")
	var errRet atomic.Pointer[error]
	errRet.Store(&ErrThrottling)
	go func() {
		err := throttler.Throttle(context.Background(), func() error { return *errRet.Load() })
		assert.Equal(t, err, errUnknown)
	}()

	time.Sleep(time.Millisecond * 10)
	errRet.Store(&errUnknown)
	clk.Step(time.Second)

	// initiate next probe immediately
	go func() {
		time.Sleep(time.Millisecond * 10)
		clk.Step(time.Microsecond)
	}()
	err := throttler.Throttle(context.Background(), func() error { return nil })
	assert.NoError(t, err)
}
