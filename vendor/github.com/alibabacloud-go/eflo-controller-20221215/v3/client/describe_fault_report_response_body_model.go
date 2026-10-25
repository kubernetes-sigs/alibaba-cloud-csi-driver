// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *DescribeFaultReportResponseBody
	GetAccessDeniedDetail() *string
	SetCreateTime(v string) *DescribeFaultReportResponseBody
	GetCreateTime() *string
	SetDescription(v string) *DescribeFaultReportResponseBody
	GetDescription() *string
	SetDiagnosisType(v string) *DescribeFaultReportResponseBody
	GetDiagnosisType() *string
	SetErrorMessage(v string) *DescribeFaultReportResponseBody
	GetErrorMessage() *string
	SetFinishTime(v string) *DescribeFaultReportResponseBody
	GetFinishTime() *string
	SetReportId(v string) *DescribeFaultReportResponseBody
	GetReportId() *string
	SetRequestId(v string) *DescribeFaultReportResponseBody
	GetRequestId() *string
	SetStatus(v string) *DescribeFaultReportResponseBody
	GetStatus() *string
}

type DescribeFaultReportResponseBody struct {
	// The details about the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The time when the fault report was created. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2025-11-06 17:31:33
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the fault report.
	//
	// example:
	//
	// GPU 0 reporting XID 79 (fallen off bus)
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The diagnosis type.
	//
	// example:
	//
	// QUICK
	DiagnosisType *string `json:"DiagnosisType,omitempty" xml:"DiagnosisType,omitempty"`
	// The error message.
	//
	// example:
	//
	// Resource Not Found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The time when the fault report was processed. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2025-11-07 17:41:33
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The report ID.
	//
	// example:
	//
	// 100184932123
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// DA4F08D4-DA54-5407-84B9-108C71D75B17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the fault report.
	//
	// example:
	//
	// Processing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeFaultReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultReportResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFaultReportResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *DescribeFaultReportResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeFaultReportResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeFaultReportResponseBody) GetDiagnosisType() *string {
	return s.DiagnosisType
}

func (s *DescribeFaultReportResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeFaultReportResponseBody) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeFaultReportResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeFaultReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFaultReportResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeFaultReportResponseBody) SetAccessDeniedDetail(v string) *DescribeFaultReportResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *DescribeFaultReportResponseBody) SetCreateTime(v string) *DescribeFaultReportResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeFaultReportResponseBody) SetDescription(v string) *DescribeFaultReportResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeFaultReportResponseBody) SetDiagnosisType(v string) *DescribeFaultReportResponseBody {
	s.DiagnosisType = &v
	return s
}

func (s *DescribeFaultReportResponseBody) SetErrorMessage(v string) *DescribeFaultReportResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeFaultReportResponseBody) SetFinishTime(v string) *DescribeFaultReportResponseBody {
	s.FinishTime = &v
	return s
}

func (s *DescribeFaultReportResponseBody) SetReportId(v string) *DescribeFaultReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *DescribeFaultReportResponseBody) SetRequestId(v string) *DescribeFaultReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFaultReportResponseBody) SetStatus(v string) *DescribeFaultReportResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeFaultReportResponseBody) Validate() error {
	return dara.Validate(s)
}
