// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosticMetricSetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMetricSets(v []*DescribeDiagnosticMetricSetsResponseBodyMetricSets) *DescribeDiagnosticMetricSetsResponseBody
	GetMetricSets() []*DescribeDiagnosticMetricSetsResponseBodyMetricSets
	SetNextToken(v string) *DescribeDiagnosticMetricSetsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeDiagnosticMetricSetsResponseBody
	GetRequestId() *string
}

type DescribeDiagnosticMetricSetsResponseBody struct {
	// The diagnostic metric sets.
	MetricSets []*DescribeDiagnosticMetricSetsResponseBodyMetricSets `json:"MetricSets,omitempty" xml:"MetricSets,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosticMetricSetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticMetricSetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticMetricSetsResponseBody) GetMetricSets() []*DescribeDiagnosticMetricSetsResponseBodyMetricSets {
	return s.MetricSets
}

func (s *DescribeDiagnosticMetricSetsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDiagnosticMetricSetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosticMetricSetsResponseBody) SetMetricSets(v []*DescribeDiagnosticMetricSetsResponseBodyMetricSets) *DescribeDiagnosticMetricSetsResponseBody {
	s.MetricSets = v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponseBody) SetNextToken(v string) *DescribeDiagnosticMetricSetsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponseBody) SetRequestId(v string) *DescribeDiagnosticMetricSetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponseBody) Validate() error {
	if s.MetricSets != nil {
		for _, item := range s.MetricSets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDiagnosticMetricSetsResponseBodyMetricSets struct {
	// The description of the diagnostic metric set.
	//
	// example:
	//
	// connection issue diagnostics
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The IDs of the diagnostic metrics.
	MetricIds []*string `json:"MetricIds,omitempty" xml:"MetricIds,omitempty" type:"Repeated"`
	// The ID of the diagnostic metric set.
	//
	// example:
	//
	// dms-bp17p0qwtr72zmu*****
	MetricSetId *string `json:"MetricSetId,omitempty" xml:"MetricSetId,omitempty"`
	// The name of the diagnostic metric set.
	//
	// example:
	//
	// connection issue diagnostics
	MetricSetName *string `json:"MetricSetName,omitempty" xml:"MetricSetName,omitempty"`
	// The resource type supported by the diagnostic metric set.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The type of the diagnostic metric set. Valid values:
	//
	// 	- User: user-defined diagnostic metric set
	//
	// 	- Common: common diagnostic metric set
	//
	// example:
	//
	// User
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeDiagnosticMetricSetsResponseBodyMetricSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticMetricSetsResponseBodyMetricSets) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) GetDescription() *string {
	return s.Description
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) GetMetricIds() []*string {
	return s.MetricIds
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) GetMetricSetId() *string {
	return s.MetricSetId
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) GetMetricSetName() *string {
	return s.MetricSetName
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) GetType() *string {
	return s.Type
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) SetDescription(v string) *DescribeDiagnosticMetricSetsResponseBodyMetricSets {
	s.Description = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) SetMetricIds(v []*string) *DescribeDiagnosticMetricSetsResponseBodyMetricSets {
	s.MetricIds = v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) SetMetricSetId(v string) *DescribeDiagnosticMetricSetsResponseBodyMetricSets {
	s.MetricSetId = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) SetMetricSetName(v string) *DescribeDiagnosticMetricSetsResponseBodyMetricSets {
	s.MetricSetName = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) SetResourceType(v string) *DescribeDiagnosticMetricSetsResponseBodyMetricSets {
	s.ResourceType = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) SetType(v string) *DescribeDiagnosticMetricSetsResponseBodyMetricSets {
	s.Type = &v
	return s
}

func (s *DescribeDiagnosticMetricSetsResponseBodyMetricSets) Validate() error {
	return dara.Validate(s)
}
