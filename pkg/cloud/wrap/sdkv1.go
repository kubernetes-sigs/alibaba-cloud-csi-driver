package wrap

import (
	"net/http"
	"time"

	alierrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/go-logr/logr"
)

func transformErrorForLog(errIn error) (code, requestID string, err error) {
	err = errIn
	// Intentionally not using errors.As, we should only deal with direct SDK errors.
	// Don't throw away other wrappers.
	if alierr, ok := err.(alierrors.Error); ok {
		code = alierr.ErrorCode()
		err = &briefAliError{alierr}
		if svrerr, ok := alierr.(*alierrors.ServerError); ok {
			requestID = svrerr.RequestId()
		}
	}
	return
}

type tResp interface {
	responses.AcsResponse
	comparable
}

func V1[TReq requests.AcsRequest, TResp tResp](logger logr.Logger, f func(TReq) (TResp, error)) func(TReq) (TResp, error) {
	return func(req TReq) (TResp, error) {
		helper, logger := logger.WithCallStackHelper()
		logger = logger.WithValues("api", req.GetActionName())
		helper()
		logger.V(5).Info("OpenAPI trace", "request", req)
		logger.V(3).Info("OpenAPI start")

		t := time.Now()
		resp, err := f(req)
		logger.V(5).Info("OpenAPI trace", "response", resp, "error", err)

		level := 2
		duration := time.Since(t)
		attrs := []any{"elapsed", duration}
		var reqID string

		var respZero TResp
		if resp != respZero {
			var header http.Header = resp.GetHttpHeaders()
			reqID = header.Get("X-Acs-Request-Id")
		}
		if err != nil {
			level = 1
			code, r, e := transformErrorForLog(err)
			err = e
			if code != "" {
				attrs = append(attrs, "errorCode", code)
			} else {
				// No error code, just log entire error
				attrs = append(attrs, "error", err)
			}
			if r != "" {
				reqID = r
			}
		}
		if reqID != "" {
			attrs = append(attrs, "requestID", reqID)
		}
		logger.V(level).Info("OpenAPI finish", attrs...)
		return resp, err
	}
}
