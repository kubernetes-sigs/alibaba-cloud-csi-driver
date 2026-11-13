package throttle

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	alierrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"k8s.io/klog/v2"
)

// Class categorizes an error returned by an OpenAPI call for throttling purposes.
type Class int

const (
	// ClassUnknown means the error is not a recognized cloud API error, so we
	// cannot tell whether we are being throttled. Throttle state is left untouched.
	ClassUnknown Class = iota
	// ClassOK means the call was a recognized outcome that is not throttling
	// (success or a non-throttling cloud error). It allows admission to grow.
	ClassOK
	// ClassThrottling means the cloud responded with a throttling error, so we
	// should back off and retry.
	ClassThrottling
)

// Classifier maps an error (which may be nil) to a Class plus a best-effort
// request ID for logging (empty when unavailable). It lets the Throttler work
// with different Alibaba Cloud SDKs whose error types are incompatible.
type Classifier func(error) (Class, string)

// V1Classifier classifies errors from the old SDK
// (github.com/aliyun/alibaba-cloud-sdk-go), whose errors are *alierrors.ServerError.
func V1Classifier(err error) (Class, string) {
	if err == nil {
		return ClassOK, ""
	}
	alierr, ok := errors.AsType[*alierrors.ServerError](err)
	if !ok {
		return ClassUnknown, ""
	}
	if alierr.ErrorCode() == "Throttling" {
		return ClassThrottling, alierr.RequestId()
	}
	return ClassOK, alierr.RequestId()
}

// V2Classifier classifies errors from the new tea SDK
// (github.com/alibabacloud-go/*), whose errors are *tea.SDKError. Throttling
// codes take several forms (Throttling, Throttling.User, Throttling.Api), so we
// match on the "Throttling" prefix.
func V2Classifier(err error) (Class, string) {
	if err == nil {
		return ClassOK, ""
	}
	sdkErr, ok := errors.AsType[*tea.SDKError](err)
	if !ok {
		return ClassUnknown, ""
	}
	requestID := sdkErrorRequestID(sdkErr)
	if strings.HasPrefix(tea.StringValue(sdkErr.Code), "Throttling") {
		return ClassThrottling, requestID
	}
	return ClassOK, requestID
}

// sdkErrorRequestID extracts the request ID from a tea SDKError. The tea SDK
// stashes it inside the JSON-encoded Data field rather than a typed accessor
// (see pkg/cloud/wrap.transformV2ErrorForLog).
func sdkErrorRequestID(e *tea.SDKError) string {
	if e.Data == nil {
		return ""
	}
	data := *e.Data
	if data == "" || data[0] != '{' { // only attempt to parse a JSON object
		return ""
	}
	var parsed struct {
		RequestId string
	}
	if json.Unmarshal([]byte(data), &parsed) != nil {
		return ""
	}
	return parsed.RequestId
}

type throttlingError struct {
	err error
}

func (e throttlingError) Error() string {
	return fmt.Sprintf("OpenAPI is throttling: %v", e.err)
}

func (t throttlingError) Unwrap() error {
	return t.err
}

const (
	recoveryWindow = 16
	maxWindow      = 512
)

type admission struct {
	generation uint64
	probe      bool
}

type Throttler struct {
	initDelay time.Duration
	maxDelay  time.Duration
	classify  Classifier

	mu         sync.Mutex
	backoff    bool
	generation uint64
	window     int
	inFlight   int
	ready      chan struct{}

	retryDelay  time.Duration
	nextProbeAt time.Time
	probeOwned  bool
}

func NewThrottler(initDelay, maxDelay time.Duration, classify Classifier) *Throttler {
	return &Throttler{
		initDelay: initDelay,
		maxDelay:  maxDelay,
		classify:  classify,
		window:    maxWindow,
		ready:     make(chan struct{}, 1),
	}
}

// Throttle calls f, retrying throttling errors with exponential backoff and a
// single probe. An accepted probe resumes admission with 16 concurrent calls;
// each accepted completion grows the window by one, up to 512. This limits
// recovery bursts without imposing a fixed request rate on minute-based quotas.
func (t *Throttler) Throttle(ctx context.Context, f func() error) error {
	for {
		ticket, err := t.acquire(ctx)
		if err != nil {
			return err
		}
		if ticket.probe {
			klog.FromContext(ctx).V(2).Info("OpenAPI throttling, probing")
		}

		err = f()
		class, requestID := t.classify(err)
		t.complete(ctx, ticket, class, requestID)
		if class != ClassThrottling {
			return err
		}
	}
}

func (t *Throttler) acquire(ctx context.Context) (admission, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	// Pass the notification on after admission or cancellation, not failed checks.
	defer t.notifyLocked()
	for {
		if err := ctx.Err(); err != nil {
			if t.backoff {
				return admission{}, throttlingError{err}
			}
			return admission{}, err
		}
		if !t.backoff && t.inFlight < t.window {
			t.inFlight++
			return admission{generation: t.generation}, nil
		}
		if t.backoff && !t.probeOwned && t.inFlight < maxWindow {
			t.probeOwned = true
			delay := time.Until(t.nextProbeAt)
			if delay > 0 {
				t.mu.Unlock()
				select {
				case <-ctx.Done():
				case <-time.After(delay):
				}
				t.mu.Lock()
			}
			if err := ctx.Err(); err != nil {
				t.probeOwned = false
				return admission{}, throttlingError{err}
			}
			t.inFlight++
			return admission{generation: t.generation, probe: true}, nil
		}

		t.mu.Unlock()
		select {
		case <-ctx.Done():
		case <-t.ready:
		}
		t.mu.Lock()
	}
}

func (t *Throttler) complete(ctx context.Context, ticket admission, class Class, requestID string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	defer t.notifyLocked()
	t.inFlight--
	// Older calls still occupy slots, but cannot undo a newer backoff.
	if ticket.generation != t.generation {
		return
	}
	if ticket.probe {
		t.probeOwned = false
	}
	logger := klog.FromContext(ctx)
	switch class {
	case ClassOK:
		if ticket.probe {
			t.backoff = false
			t.window = recoveryWindow
			logger.V(1).Info("OpenAPI throttling ended, resuming admission", "window", t.window)
		} else if t.window < maxWindow {
			t.window++
		}
	case ClassThrottling:
		if !t.backoff {
			t.backoff = true
			t.generation++
			t.retryDelay = t.initDelay
			logger.V(1).Info("OpenAPI throttling detected", "requestID", requestID)
		} else {
			t.retryDelay = min(t.retryDelay*2, t.maxDelay)
			logger.V(2).Info("still throttling", "retry", t.retryDelay, "requestID", requestID)
		}
		t.nextProbeAt = time.Now().Add(t.retryDelay)
	}
}

func (t *Throttler) notifyLocked() {
	// Notifications are hints, not permits; admission is checked under mu.
	select {
	case t.ready <- struct{}{}:
	default:
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
