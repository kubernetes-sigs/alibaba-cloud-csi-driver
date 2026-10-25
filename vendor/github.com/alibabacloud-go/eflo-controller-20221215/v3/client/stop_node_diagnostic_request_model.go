// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopNodeDiagnosticRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *StopNodeDiagnosticRequest
	GetReportId() *string
}

type StopNodeDiagnosticRequest struct {
	// The report ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 15678327322
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s StopNodeDiagnosticRequest) String() string {
	return dara.Prettify(s)
}

func (s StopNodeDiagnosticRequest) GoString() string {
	return s.String()
}

func (s *StopNodeDiagnosticRequest) GetReportId() *string {
	return s.ReportId
}

func (s *StopNodeDiagnosticRequest) SetReportId(v string) *StopNodeDiagnosticRequest {
	s.ReportId = &v
	return s
}

func (s *StopNodeDiagnosticRequest) Validate() error {
	return dara.Validate(s)
}
