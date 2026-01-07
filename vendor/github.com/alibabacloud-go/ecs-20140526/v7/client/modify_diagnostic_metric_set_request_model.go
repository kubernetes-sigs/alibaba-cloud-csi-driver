// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiagnosticMetricSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyDiagnosticMetricSetRequest
	GetDescription() *string
	SetMetricIds(v []*string) *ModifyDiagnosticMetricSetRequest
	GetMetricIds() []*string
	SetMetricSetId(v string) *ModifyDiagnosticMetricSetRequest
	GetMetricSetId() *string
	SetMetricSetName(v string) *ModifyDiagnosticMetricSetRequest
	GetMetricSetName() *string
	SetRegionId(v string) *ModifyDiagnosticMetricSetRequest
	GetRegionId() *string
	SetResourceType(v string) *ModifyDiagnosticMetricSetRequest
	GetResourceType() *string
}

type ModifyDiagnosticMetricSetRequest struct {
	// The description of the diagnostic metric set.
	//
	// example:
	//
	// connection diagnostics
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of diagnostic metrics.
	MetricIds []*string `json:"MetricIds,omitempty" xml:"MetricIds,omitempty" type:"Repeated"`
	// The IDs of the diagnostic metric sets.
	//
	// This parameter is required.
	//
	// example:
	//
	// dms-uf6i0tv2refv8wz*****
	MetricSetId *string `json:"MetricSetId,omitempty" xml:"MetricSetId,omitempty"`
	// The name of the diagnostic metric set.
	//
	// example:
	//
	// remoteConnectError
	MetricSetName *string `json:"MetricSetName,omitempty" xml:"MetricSetName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ModifyDiagnosticMetricSetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiagnosticMetricSetRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiagnosticMetricSetRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDiagnosticMetricSetRequest) GetMetricIds() []*string {
	return s.MetricIds
}

func (s *ModifyDiagnosticMetricSetRequest) GetMetricSetId() *string {
	return s.MetricSetId
}

func (s *ModifyDiagnosticMetricSetRequest) GetMetricSetName() *string {
	return s.MetricSetName
}

func (s *ModifyDiagnosticMetricSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDiagnosticMetricSetRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ModifyDiagnosticMetricSetRequest) SetDescription(v string) *ModifyDiagnosticMetricSetRequest {
	s.Description = &v
	return s
}

func (s *ModifyDiagnosticMetricSetRequest) SetMetricIds(v []*string) *ModifyDiagnosticMetricSetRequest {
	s.MetricIds = v
	return s
}

func (s *ModifyDiagnosticMetricSetRequest) SetMetricSetId(v string) *ModifyDiagnosticMetricSetRequest {
	s.MetricSetId = &v
	return s
}

func (s *ModifyDiagnosticMetricSetRequest) SetMetricSetName(v string) *ModifyDiagnosticMetricSetRequest {
	s.MetricSetName = &v
	return s
}

func (s *ModifyDiagnosticMetricSetRequest) SetRegionId(v string) *ModifyDiagnosticMetricSetRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiagnosticMetricSetRequest) SetResourceType(v string) *ModifyDiagnosticMetricSetRequest {
	s.ResourceType = &v
	return s
}

func (s *ModifyDiagnosticMetricSetRequest) Validate() error {
	return dara.Validate(s)
}
