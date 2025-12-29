// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticReportResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *CreateDiagnosticReportResponseBody
	GetReportId() *string
	SetRequestId(v string) *CreateDiagnosticReportResponseBody
	GetRequestId() *string
}

type CreateDiagnosticReportResponseBody struct {
	// The unique ID of the diagnostic report.
	//
	// example:
	//
	// dr-uf6i0tv2refv8wz*****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiagnosticReportResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticReportResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportResponseBody) GetReportId() *string {
	return s.ReportId
}

func (s *CreateDiagnosticReportResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiagnosticReportResponseBody) SetReportId(v string) *CreateDiagnosticReportResponseBody {
	s.ReportId = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) SetRequestId(v string) *CreateDiagnosticReportResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiagnosticReportResponseBody) Validate() error {
	return dara.Validate(s)
}
