// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosticId(v string) *CreateDiagnosticTaskResponseBody
	GetDiagnosticId() *string
	SetRequestId(v string) *CreateDiagnosticTaskResponseBody
	GetRequestId() *string
}

type CreateDiagnosticTaskResponseBody struct {
	// The ID of the diagnostics task.
	//
	// example:
	//
	// diag-i150553931717380274931
	DiagnosticId *string `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A511C02A-0127-51AA-A9F9-966382C9A1B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiagnosticTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskResponseBody) GetDiagnosticId() *string {
	return s.DiagnosticId
}

func (s *CreateDiagnosticTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiagnosticTaskResponseBody) SetDiagnosticId(v string) *CreateDiagnosticTaskResponseBody {
	s.DiagnosticId = &v
	return s
}

func (s *CreateDiagnosticTaskResponseBody) SetRequestId(v string) *CreateDiagnosticTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiagnosticTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
