// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFaultReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *ListFaultReportsResponseBody
	GetAccessDeniedDetail() *string
	SetNextToken(v string) *ListFaultReportsResponseBody
	GetNextToken() *string
	SetReports(v []*ListFaultReportsResponseBodyReports) *ListFaultReportsResponseBody
	GetReports() []*ListFaultReportsResponseBodyReports
	SetRequestId(v string) *ListFaultReportsResponseBody
	GetRequestId() *string
}

type ListFaultReportsResponseBody struct {
	// The details of the access denial.
	//
	// example:
	//
	// None
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// The pagination token.
	//
	// example:
	//
	// DA4F08D
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The fault report details.
	Reports []*ListFaultReportsResponseBodyReports `json:"Reports,omitempty" xml:"Reports,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// DA4F08D4-DA54-5407-84B9-108C71D75B17
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFaultReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFaultReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFaultReportsResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *ListFaultReportsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFaultReportsResponseBody) GetReports() []*ListFaultReportsResponseBodyReports {
	return s.Reports
}

func (s *ListFaultReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFaultReportsResponseBody) SetAccessDeniedDetail(v string) *ListFaultReportsResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *ListFaultReportsResponseBody) SetNextToken(v string) *ListFaultReportsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFaultReportsResponseBody) SetReports(v []*ListFaultReportsResponseBodyReports) *ListFaultReportsResponseBody {
	s.Reports = v
	return s
}

func (s *ListFaultReportsResponseBody) SetRequestId(v string) *ListFaultReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFaultReportsResponseBody) Validate() error {
	if s.Reports != nil {
		for _, item := range s.Reports {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFaultReportsResponseBodyReports struct {
	// The time when the fault report was created. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2025-11-07 17:41:33
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
	// The time when the fault report was resolved. The time is in the YYYY-MM-DD HH:mm:ss format.
	//
	// example:
	//
	// 2025-11-07 17:41:33
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-3A244BE3
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The fault report ID.
	//
	// example:
	//
	// 100184932123
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The status of the fault report.
	//
	// example:
	//
	// Processing
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListFaultReportsResponseBodyReports) String() string {
	return dara.Prettify(s)
}

func (s ListFaultReportsResponseBodyReports) GoString() string {
	return s.String()
}

func (s *ListFaultReportsResponseBodyReports) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListFaultReportsResponseBodyReports) GetDescription() *string {
	return s.Description
}

func (s *ListFaultReportsResponseBodyReports) GetDiagnosisType() *string {
	return s.DiagnosisType
}

func (s *ListFaultReportsResponseBodyReports) GetFinishTime() *string {
	return s.FinishTime
}

func (s *ListFaultReportsResponseBodyReports) GetNodeId() *string {
	return s.NodeId
}

func (s *ListFaultReportsResponseBodyReports) GetReportId() *string {
	return s.ReportId
}

func (s *ListFaultReportsResponseBodyReports) GetStatus() *string {
	return s.Status
}

func (s *ListFaultReportsResponseBodyReports) SetCreateTime(v string) *ListFaultReportsResponseBodyReports {
	s.CreateTime = &v
	return s
}

func (s *ListFaultReportsResponseBodyReports) SetDescription(v string) *ListFaultReportsResponseBodyReports {
	s.Description = &v
	return s
}

func (s *ListFaultReportsResponseBodyReports) SetDiagnosisType(v string) *ListFaultReportsResponseBodyReports {
	s.DiagnosisType = &v
	return s
}

func (s *ListFaultReportsResponseBodyReports) SetFinishTime(v string) *ListFaultReportsResponseBodyReports {
	s.FinishTime = &v
	return s
}

func (s *ListFaultReportsResponseBodyReports) SetNodeId(v string) *ListFaultReportsResponseBodyReports {
	s.NodeId = &v
	return s
}

func (s *ListFaultReportsResponseBodyReports) SetReportId(v string) *ListFaultReportsResponseBodyReports {
	s.ReportId = &v
	return s
}

func (s *ListFaultReportsResponseBodyReports) SetStatus(v string) *ListFaultReportsResponseBodyReports {
	s.Status = &v
	return s
}

func (s *ListFaultReportsResponseBodyReports) Validate() error {
	return dara.Validate(s)
}
