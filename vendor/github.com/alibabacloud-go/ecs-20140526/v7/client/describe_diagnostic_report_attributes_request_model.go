// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticReportAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DescribeDiagnosticReportAttributesRequest
	GetRegionId() *string
	SetReportId(v string) *DescribeDiagnosticReportAttributesRequest
	GetReportId() *string
}

type DescribeDiagnosticReportAttributesRequest struct {
	// The region ID of the diagnostic report. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the diagnostic report.
	//
	// This parameter is required.
	//
	// example:
	//
	// dr-i-uf6i0tv2refv8wz*****
	ReportId *string `json:"ReportId,omitempty" xml:"ReportId,omitempty"`
}

func (s DescribeDiagnosticReportAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticReportAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticReportAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnosticReportAttributesRequest) GetReportId() *string {
	return s.ReportId
}

func (s *DescribeDiagnosticReportAttributesRequest) SetRegionId(v string) *DescribeDiagnosticReportAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesRequest) SetReportId(v string) *DescribeDiagnosticReportAttributesRequest {
	s.ReportId = &v
	return s
}

func (s *DescribeDiagnosticReportAttributesRequest) Validate() error {
	return dara.Validate(s)
}
