package wrap

import (
	"fmt"
)

type ErrorCode string

func (e ErrorCode) Error() string {
	return "Alibaba Cloud error: " + string(e)
}

type aliError interface {
	error
	Message() string
	ErrorCode() string
}

type briefAliError struct {
	err aliError
}

func (e briefAliError) Error() string {
	return fmt.Sprintf("OpenAPI returned error: %s (%s)", e.err.Message(), e.err.ErrorCode())
}

func (e briefAliError) Unwrap() []error {
	return []error{e.err, ErrorCode(e.err.ErrorCode())}
}
