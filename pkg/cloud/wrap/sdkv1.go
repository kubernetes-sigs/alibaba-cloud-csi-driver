package wrap

import (
	"net/http"
	"time"

	alierrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
	"github.com/go-logr/logr"
)

func transformErrorForLog(errIn error) (attrs []any, requestID string, err error) {
	err = errIn
	if alierr, ok := err.(alierrors.Error); ok {
		attrs = append(attrs, "errorCode", alierr.ErrorCode())
		err = &briefAliError{alierr}
		if svrerr, ok := alierr.(*alierrors.ServerError); ok {
			requestID = svrerr.RequestId()
		}
	} else {
		// No error code, just log entire error
		attrs = append(attrs, "error", err)
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
		logger.V(3).Info("OpenAPI start")

		t := time.Now()
		resp, err := f(req)

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
			a, r, e := transformErrorForLog(err)
			err = e
			attrs = append(attrs, a...)
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
