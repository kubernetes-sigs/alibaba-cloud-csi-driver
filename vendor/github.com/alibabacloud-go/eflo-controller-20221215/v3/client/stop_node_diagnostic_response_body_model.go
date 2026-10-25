// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopNodeDiagnosticResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *StopNodeDiagnosticResponseBody
	GetAccessDeniedDetail() *string
	SetErrorMessage(v string) *StopNodeDiagnosticResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *StopNodeDiagnosticResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopNodeDiagnosticResponseBody
	GetSuccess() *bool
}

type StopNodeDiagnosticResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The error message.
	//
	// example:
	//
	// Resource Not Found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DA4F08D4-DA54-5407-84B9-108C71D75B17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the stop operation is successful.
	//
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopNodeDiagnosticResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopNodeDiagnosticResponseBody) GoString() string {
	return s.String()
}

func (s *StopNodeDiagnosticResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *StopNodeDiagnosticResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *StopNodeDiagnosticResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopNodeDiagnosticResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopNodeDiagnosticResponseBody) SetAccessDeniedDetail(v string) *StopNodeDiagnosticResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *StopNodeDiagnosticResponseBody) SetErrorMessage(v string) *StopNodeDiagnosticResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *StopNodeDiagnosticResponseBody) SetRequestId(v string) *StopNodeDiagnosticResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopNodeDiagnosticResponseBody) SetSuccess(v bool) *StopNodeDiagnosticResponseBody {
	s.Success = &v
	return s
}

func (s *StopNodeDiagnosticResponseBody) Validate() error {
	return dara.Validate(s)
}
