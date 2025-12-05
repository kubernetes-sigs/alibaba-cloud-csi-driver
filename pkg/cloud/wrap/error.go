package wrap

import (
	"fmt"

	alierrors "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
)

type ErrorCode string

func (e ErrorCode) Error() string {
	return "Alibaba Cloud error: " + string(e)
}

type briefAliError struct {
	err alierrors.Error
}

func (e *briefAliError) Error() string {
	return fmt.Sprintf("OpenAPI returned error: %s (%s)", e.err.Message(), e.err.ErrorCode())
}

func (e *briefAliError) Unwrap() []error {
	return []error{e.err, ErrorCode(e.err.ErrorCode())}
}
