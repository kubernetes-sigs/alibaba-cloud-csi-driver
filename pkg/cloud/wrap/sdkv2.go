package wrap

import (
	"context"
	"encoding/json"
	"time"

	"github.com/alibabacloud-go/tea/dara"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/klog/v2"
	"k8s.io/utils/ptr"
)

type teaError struct {
	*tea.SDKError
	message string
}

func (e *teaError) Message() string {
	if e.message != "" {
		return e.message
	}
	// SDKError.Message contains too much info (includes requestID)
	// We don't like it, because we want the error message for a repeating error to be consistent,
	// for efficient K8s event aggregation.
	return ptr.Deref(e.SDKError.Message, "[NO MESSAGE]")
}

func (e *teaError) ErrorCode() string {
	return ptr.Deref(e.SDKError.Code, "[NO CODE]")
}

func (e *teaError) Unwrap() error {
	return e.SDKError
}

// transformV2ErrorForLog wraps a direct tea SDK error in a briefAliError (stable
// message for k8s event aggregation) and returns its code and request ID for logging.
func transformV2ErrorForLog(errIn error) (code, requestID string, err error) {
	err = errIn
	// Intentionally not using errors.As, we should only deal with direct SDK errors.
	// Don't throw away other wrappers.
	if teaerr, ok := err.(*tea.SDKError); ok {
		e2 := &teaError{SDKError: teaerr}
		code = ptr.Deref(teaerr.Code, "")
		if teaerr.Data != nil && len(*teaerr.Data) > 0 && (*teaerr.Data)[0] == '{' { // likely a json object
			var data struct {
				Message   string
				RequestId string
			}
			if json.Unmarshal([]byte(*teaerr.Data), &data) == nil {
				requestID = data.RequestId
				e2.message = data.Message
			}
		}
		err = briefAliError{e2}
	}
	return
}

// v2Resp is satisfied by every OpenAPI v2 response type; GetHeaders carries the request ID.
type v2Resp interface {
	comparable
	GetHeaders() map[string]*string
}

// logV2 runs one OpenAPI v2 call, logging it at the boundary (logger comes from ctx),
// recording metrics, and transforming the error for stable message.
func logV2[TReq any, TResp v2Resp](ctx context.Context, product, action string, req TReq, o *dara.RuntimeOptions, f func(context.Context, TReq, *dara.RuntimeOptions) (TResp, error)) (TResp, error) {
	logger := klog.FromContext(ctx).WithCallDepth(2).WithValues("action", action)
	if logger.V(5).Enabled() {
		logger.V(3).Info("OpenAPI start", "request", req)
	} else {
		logger.V(3).Info("OpenAPI start")
	}

	t := time.Now()
	resp, err := f(ctx, req, o)

	level := 2
	attrs := make([]any, 0, 8)
	elapsed := time.Since(t)
	attrs = append(attrs, "elapsed", elapsed)

	var respZero TResp
	var reqID string
	if resp != respZero {
		header := resp.GetHeaders()
		reqID = ptr.Deref(header["x-acs-request-id"], "")
	}
	var code = "OK"
	if err != nil {
		level = 1
		var r string
		code, r, err = transformV2ErrorForLog(err)
		if code != "" {
			attrs = append(attrs, "errorCode", code)
		} else {
			attrs = append(attrs, "error", err)
			code = "Unknown"
		}
		if r != "" {
			reqID = r
		}
	}
	if reqID != "" {
		attrs = append(attrs, "requestID", reqID)
	}
	if logger.V(5).Enabled() {
		attrs = append(attrs, "response", resp)
	}
	logger.V(level).Info("OpenAPI finish", attrs...)

	labels := prometheus.Labels{
		"product": product,
		"action":  action,
		"code":    code,
	}
	APIRequestsTotal.With(labels).Inc()
	APIDurationSecondsTotal.With(labels).Add(elapsed.Seconds())
	return resp, err
}
