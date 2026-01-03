// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiagnosticReportsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *DeleteDiagnosticReportsRequest
	GetRegionId() *string
	SetReportIds(v []*string) *DeleteDiagnosticReportsRequest
	GetReportIds() []*string
}

type DeleteDiagnosticReportsRequest struct {
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of the diagnostic reports. You can specify up to 100 resource IDs.
	//
	// This parameter is required.
	ReportIds []*string `json:"ReportIds,omitempty" xml:"ReportIds,omitempty" type:"Repeated"`
}

func (s DeleteDiagnosticReportsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiagnosticReportsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiagnosticReportsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDiagnosticReportsRequest) GetReportIds() []*string {
	return s.ReportIds
}

func (s *DeleteDiagnosticReportsRequest) SetRegionId(v string) *DeleteDiagnosticReportsRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDiagnosticReportsRequest) SetReportIds(v []*string) *DeleteDiagnosticReportsRequest {
	s.ReportIds = v
	return s
}

func (s *DeleteDiagnosticReportsRequest) Validate() error {
	return dara.Validate(s)
}
