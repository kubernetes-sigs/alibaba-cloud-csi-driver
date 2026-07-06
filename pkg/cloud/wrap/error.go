package wrap

import (
	"fmt"
)

type aliError interface {
	error
	Message() string
	ErrorCode() string
}

// briefAliError gives an SDK error a brief, stable message (no requestID) so
// that a repeating error renders identically, for efficient k8s event
// aggregation. It keeps the underlying error in the chain, so errors.As still
// reaches the SDK error (and thus cloud.ErrorCodeV2 works through it).
type briefAliError struct {
	err aliError
}

func (e briefAliError) Error() string {
	return fmt.Sprintf("OpenAPI returned error: %s (%s)", e.err.Message(), e.err.ErrorCode())
}

func (e briefAliError) Unwrap() error {
	return e.err
}
