// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticMetricSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribeDiagnosticMetricSetsRequest
	GetMaxResults() *int32
	SetMetricSetIds(v []*string) *DescribeDiagnosticMetricSetsRequest
	GetMetricSetIds() []*string
	SetNextToken(v string) *DescribeDiagnosticMetricSetsRequest
	GetNextToken() *string
	SetRegionId(v string) *DescribeDiagnosticMetricSetsRequest
	GetRegionId() *string
	SetResourceType(v string) *DescribeDiagnosticMetricSetsRequest
	GetResourceType() *string
	SetType(v string) *DescribeDiagnosticMetricSetsRequest
	GetType() *string
}

type DescribeDiagnosticMetricSetsRequest struct {
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value:
	//
	// 	- If this parameter is left empty, the default value is 10.
	//
	// 	- If you set this parameter to a value that is greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The IDs of diagnostic metric sets.
	MetricSetIds []*string `json:"MetricSetIds,omitempty" xml:"MetricSetIds,omitempty" type:"Repeated"`
	// The pagination token that is used in the request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The region ID of the diagnostic metric set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The resource type supported by the diagnostic metric set.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The type of the diagnostic metric set. Valid values:
	//
	// 	- User: custom diagnostic metric set
	//
	// 	- Common: public diagnostic metric set
	//
	// Default value: User.
	//
	// example:
	//
	// User
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDiagnosticMetricSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticMetricSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticMetricSetsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDiagnosticMetricSetsRequest) GetMetricSetIds() []*string {
	return s.MetricSetIds
}

func (s *DescribeDiagnosticMetricSetsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiagnosticMetricSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDiagnosticMetricSetsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDiagnosticMetricSetsRequest) GetType() *string {
	return s.Type
}

func (s *DescribeDiagnosticMetricSetsRequest) SetMaxResults(v int32) *DescribeDiagnosticMetricSetsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsRequest) SetMetricSetIds(v []*string) *DescribeDiagnosticMetricSetsRequest {
	s.MetricSetIds = v
	return s
}

func (s *DescribeDiagnosticMetricSetsRequest) SetNextToken(v string) *DescribeDiagnosticMetricSetsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsRequest) SetRegionId(v string) *DescribeDiagnosticMetricSetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsRequest) SetResourceType(v string) *DescribeDiagnosticMetricSetsRequest {
	s.ResourceType = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsRequest) SetType(v string) *DescribeDiagnosticMetricSetsRequest {
	s.Type = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsRequest) Validate() error {
	return dara.Validate(s)
}
