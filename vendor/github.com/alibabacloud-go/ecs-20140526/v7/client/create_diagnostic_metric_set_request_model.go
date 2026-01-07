// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDiagnosticMetricSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateDiagnosticMetricSetRequest
	GetDescription() *string
	SetMetricIds(v []*string) *CreateDiagnosticMetricSetRequest
	GetMetricIds() []*string
	SetMetricSetName(v string) *CreateDiagnosticMetricSetRequest
	GetMetricSetName() *string
	SetRegionId(v string) *CreateDiagnosticMetricSetRequest
	GetRegionId() *string
	SetResourceType(v string) *CreateDiagnosticMetricSetRequest
	GetResourceType() *string
}

type CreateDiagnosticMetricSetRequest struct {
	// The description of the diagnostic metric set.
	//
	// example:
	//
	// The ID of the request.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of diagnostic metrics. You can specify up to 100 diagnostic metric IDs.
	//
	// This parameter is required.
	MetricIds []*string `json:"MetricIds,omitempty" xml:"MetricIds,omitempty" type:"Repeated"`
	// The name of the diagnostic metric set.
	//
	// example:
	//
	// The IDs of diagnostic metrics. You can specify up to 100 diagnostic metric IDs.
	MetricSetName *string `json:"MetricSetName,omitempty" xml:"MetricSetName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The type of the resource.
	//
	// Default value: instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s CreateDiagnosticMetricSetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticMetricSetRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticMetricSetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDiagnosticMetricSetRequest) GetMetricIds() []*string {
	return s.MetricIds
}

func (s *CreateDiagnosticMetricSetRequest) GetMetricSetName() *string {
	return s.MetricSetName
}

func (s *CreateDiagnosticMetricSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateDiagnosticMetricSetRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *CreateDiagnosticMetricSetRequest) SetDescription(v string) *CreateDiagnosticMetricSetRequest {
	s.Description = &v
	return s
}

func (s *CreateDiagnosticMetricSetRequest) SetMetricIds(v []*string) *CreateDiagnosticMetricSetRequest {
	s.MetricIds = v
	return s
}

func (s *CreateDiagnosticMetricSetRequest) SetMetricSetName(v string) *CreateDiagnosticMetricSetRequest {
	s.MetricSetName = &v
	return s
}

func (s *CreateDiagnosticMetricSetRequest) SetRegionId(v string) *CreateDiagnosticMetricSetRequest {
	s.RegionId = &v
	return s
}

func (s *CreateDiagnosticMetricSetRequest) SetResourceType(v string) *CreateDiagnosticMetricSetRequest {
	s.ResourceType = &v
	return s
}

func (s *CreateDiagnosticMetricSetRequest) Validate() error {
	return dara.Validate(s)
}
