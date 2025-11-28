package throttle

import (
	"context"
	"errors"
	"fmt"
	"sync/atomic"
	"time"

	alierrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

type throttlingError struct {
	err error
}

func (e throttlingError) Error() string {
	return fmt.Sprintf("OpenAPI is throttling: %v", e.err)
}

func (t throttlingError) Unwrap() error {
	return t.err
}

type throttlingContext struct {
	probing chan struct{}

	lastThrottling time.Time
	retryDelay     time.Duration
}

type Throttler struct {
	initDelay time.Duration
	maxDelay  time.Duration

	clk clock.Clock

	current atomic.Pointer[throttlingContext]
}

func NewThrottler(clk clock.Clock, initDelay, maxDelay time.Duration) *Throttler {
	return &Throttler{
		initDelay: initDelay,
		maxDelay:  maxDelay,
		clk:       clk,
	}
}

// Throttle will call f and retry after delay when Throttling error code is returned.
// If one Throttling error is observed, all subsequent requests will be delayed.
// A random delayed request will be used to probe whether the throttling has ended.
// If a non-Throttling response is returned, all delayed requests will be sent out immediately.
//
// This works with Alibaba Cloud OpenAPI, which usually enforce a quota every minute.
// If we use up the quota in the first 10s, we will get Throttling error code in the following 50s.
func (t *Throttler) Throttle(ctx context.Context, f func() error) error {
	logger := klog.FromContext(ctx)
	start := t.clk.Now()
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		tCtx := t.current.Load()
		if tCtx != nil {
			logger.V(3).Info("OpenAPI throttling in effect")
			select {
			case <-ctx.Done():
				return throttlingError{ctx.Err()}
			case _, ok := <-tCtx.probing:
				if ok {
					// Currently throttling, we will probe if the throttling has ended.
					// Be sure to write or close tCtx.probing to notify others.
					select {
					case <-ctx.Done():
						tCtx.probing <- struct{}{}
						return throttlingError{ctx.Err()}
					case <-t.clk.After(tCtx.retryDelay - t.clk.Since(tCtx.lastThrottling)):
					}
					logger.V(2).Info("OpenAPI throttling, probing", "retry", tCtx.retryDelay)
				} else {
					logger.V(2).Info("OpenAPI no longer throttle", "delay", t.clk.Since(start))
					tCtx = nil
				}
			}
		}

		err := f()

		var alierr *alierrors.ServerError
		if err != nil && !errors.As(err, &alierr) {
			// We are not sure what's wrong, don't change state.
			if tCtx != nil {
				// Initiate the next probe immediately.
				tCtx.probing <- struct{}{}
			}
			return err
		}

		if err == nil || alierr.ErrorCode() != "Throttling" {
			if tCtx != nil {
				logger.V(1).Info("OpenAPI throttling ended")
				close(tCtx.probing)
				t.current.CompareAndSwap(tCtx, nil)
			}
			return err
		}

		// Throttling detected
		if tCtx == nil {
			tCtx = &throttlingContext{
				lastThrottling: t.clk.Now(),
				retryDelay:     t.initDelay,
				probing:        make(chan struct{}, 1),
			}
			tCtx.probing <- struct{}{}
			if t.current.CompareAndSwap(nil, tCtx) {
				logger.V(1).Info("OpenAPI throttling detected", "requestID", alierr.RequestId())
			} else {
				logger.V(3).Info("already throttling", "requestID", alierr.RequestId())
			}
		} else {
			tCtx.lastThrottling = t.clk.Now()
			tCtx.retryDelay *= 2
			if tCtx.retryDelay > t.maxDelay {
				tCtx.retryDelay = t.maxDelay
			}
			tCtx.probing <- struct{}{}
			logger.V(2).Info("still throttling", "retry", tCtx.retryDelay, "requestID", alierr.RequestId())
		}
	}
}

func Throttled[TReq any, TResp any](t *Throttler, f func(TReq) (TResp, error)) func(context.Context, TReq) (TResp, error) {
	return func(ctx context.Context, req TReq) (TResp, error) {
		var resp TResp
		err := t.Throttle(ctx, func() (err error) {
			resp, err = f(req)
			return err
		})
		return resp, err
	}
}
