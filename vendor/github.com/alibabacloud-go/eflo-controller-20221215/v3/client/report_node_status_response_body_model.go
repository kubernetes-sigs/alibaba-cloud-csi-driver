// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportNodeStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ReportNodeStatusResponseBody
	GetAccessDeniedDetail() *string
	SetErrorMessage(v string) *ReportNodeStatusResponseBody
	GetErrorMessage() *string
	SetReportId(v string) *ReportNodeStatusResponseBody
	GetReportId() *string
	SetRequestId(v string) *ReportNodeStatusResponseBody
	GetRequestId() *string
}

type ReportNodeStatusResponseBody struct {
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
	// The fault report ID.
	//
	// example:
	//
	// 156837191934
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DA4F08D4-DA54-5407-84B9-108C71D75B17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportNodeStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportNodeStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ReportNodeStatusResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ReportNodeStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ReportNodeStatusResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *ReportNodeStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportNodeStatusResponseBody) SetAccessDeniedDetail(v string) *ReportNodeStatusResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ReportNodeStatusResponseBody) SetErrorMessage(v string) *ReportNodeStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReportNodeStatusResponseBody) SetReportId(v string) *ReportNodeStatusResponseBody {
	s.ReportId = &v
	return s
}

func (s *ReportNodeStatusResponseBody) SetRequestId(v string) *ReportNodeStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportNodeStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
