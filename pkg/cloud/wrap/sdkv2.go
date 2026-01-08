package wrap

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-logr/logr"
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

func transformV2ErrorForLog(errIn error) (code, requestID string, err error) {
	err = errIn
	// Intentionally not using errors.As, we should only deal with direct SDK errors.
	// Don't throw away other wrappers.
	if teaerr, ok := err.(*tea.SDKError); ok {
		e2 := &teaError{SDKError: teaerr}
		if teaerr.Code != nil {
			code = *teaerr.Code
		}
		if teaerr.Data != nil && (*teaerr.Data)[0] == '{' { // likely a json object
			var data struct {
				Message   string
				RequestId string
			}
			errJson := json.Unmarshal([]byte(*teaerr.Data), &data)
			if errJson == nil {
				requestID = data.RequestId
				e2.message = data.Message
			}
		}
		err = briefAliError{e2}
	}
	return
}

type v2Resp interface {
	comparable
	GetHeaders() map[string]*string
}

func V2[TReq any, TResp v2Resp](logger logr.Logger, f func(TReq) (TResp, error)) func(TReq) (TResp, error) {
	t := reflect.TypeFor[TReq]()
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	action := strings.TrimSuffix(t.Name(), "Request")
	logger = logger.WithValues("api", action)

	return func(req TReq) (TResp, error) {
		helper, logger := logger.WithCallStackHelper()
		helper()
		logger.V(5).Info("OpenAPI trace", "request", req)

		return v2Impl(logger, func() (TResp, error) { return f(req) })
	}
}

func v2Impl[TResp v2Resp](logger logr.Logger, f func() (TResp, error)) (TResp, error) {
	helper, logger := logger.WithCallStackHelper()
	helper()

	logger.V(3).Info("OpenAPI start")

	t := time.Now()
	resp, err := f()
	logger.V(5).Info("OpenAPI trace", "response", resp, "error", err)

	level := 2
	duration := time.Since(t)
	attrs := []any{"elapsed", duration}

	var respZero TResp
	var reqID string
	if resp != respZero {
		header := resp.GetHeaders()
		reqID = ptr.Deref(header["x-acs-request-id"], "")
	}
	if err != nil {
		level = 1
		code, r, e := transformV2ErrorForLog(err)
		if code != "" {
			attrs = append(attrs, "errorCode", code)
		} else {
			// No error code, just log entire error
			attrs = append(attrs, "error", err)
		}
		if r != "" {
			reqID = r
		}
		err = e
	}
	if reqID != "" {
		attrs = append(attrs, "requestID", reqID)
	}
	logger.V(level).Info("OpenAPI finish", attrs...)
	return resp, err

}
