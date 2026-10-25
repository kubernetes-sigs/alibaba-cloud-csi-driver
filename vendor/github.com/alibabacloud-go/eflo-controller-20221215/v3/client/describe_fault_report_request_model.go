// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFaultReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetReportId(v string) *DescribeFaultReportRequest
	GetReportId() *string
}

type DescribeFaultReportRequest struct {
	// The report ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100184932123
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DescribeFaultReportRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFaultReportRequest) GoString() string {
	return s.String()
}

func (s *DescribeFaultReportRequest) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeFaultReportRequest) SetReportId(v string) *DescribeFaultReportRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeFaultReportRequest) Validate() error {
	return dara.Validate(s)
}
