// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticReportRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *CreateDiagnosticReportRequest
	GetEndTime() *string
	SetMetricSetId(v string) *CreateDiagnosticReportRequest
	GetMetricSetId() *string
	SetRegionId(v string) *CreateDiagnosticReportRequest
	GetRegionId() *string
	SetResourceId(v string) *CreateDiagnosticReportRequest
	GetResourceId() *string
	SetStartTime(v string) *CreateDiagnosticReportRequest
	GetStartTime() *string
}

type CreateDiagnosticReportRequest struct {
	// The end time. This parameter takes effect only for diagnostic metrics that do not need to be assessed by running Cloud Assistant commands in guest operating systems.
	//
	// example:
	//
	// 2022-07-11T14:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The ID of the diagnostic metric set. If this parameter is left empty, the dms-instancedefault set is used, which is the default diagnostic metric set provided for Elastic Compute Service (ECS) instances.
	//
	// example:
	//
	// dms-uf6i0tv2refv8wz*****
	MetricSetId *string `json:"MetricSetId,omitempty" xml:"MetricSetId,omitempty"`
	// The region ID of the security group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of resource N.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-uf6i0tv2refv8wz*****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The start time. This parameter takes effect only for diagnostic metrics that do not need to be assessed by running Cloud Assistant commands in guest operating systems.
	//
	// example:
	//
	// 2022-07-11T12:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateDiagnosticReportRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticReportRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticReportRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateDiagnosticReportRequest) GetMetricSetId() *string {
	return s.MetricSetId
}

func (s *CreateDiagnosticReportRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDiagnosticReportRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateDiagnosticReportRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateDiagnosticReportRequest) SetEndTime(v string) *CreateDiagnosticReportRequest {
	s.EndTime = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetMetricSetId(v string) *CreateDiagnosticReportRequest {
	s.MetricSetId = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetRegionId(v string) *CreateDiagnosticReportRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetResourceId(v string) *CreateDiagnosticReportRequest {
	s.ResourceId = &v
	return s
}

func (s *CreateDiagnosticReportRequest) SetStartTime(v string) *CreateDiagnosticReportRequest {
	s.StartTime = &v
	return s
}

func (s *CreateDiagnosticReportRequest) Validate() error {
	return dara.Validate(s)
}
