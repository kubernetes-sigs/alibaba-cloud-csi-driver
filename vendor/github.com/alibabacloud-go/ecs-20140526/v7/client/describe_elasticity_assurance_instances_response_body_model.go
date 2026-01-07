// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticityAssuranceInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticityAssuranceItem(v *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem) *DescribeElasticityAssuranceInstancesResponseBody
	GetElasticityAssuranceItem() *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem
	SetMaxResults(v int32) *DescribeElasticityAssuranceInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeElasticityAssuranceInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeElasticityAssuranceInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeElasticityAssuranceInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeElasticityAssuranceInstancesResponseBody struct {
	// Details about the instances that match and use the elasticity assurance.
	ElasticityAssuranceItem *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem `json:"ElasticityAssuranceItem,omitempty" xml:"ElasticityAssuranceItem,omitempty" type:"Struct"`
	// The number of entries returned per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The token used to start the next query.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeElasticityAssuranceInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) GetElasticityAssuranceItem() *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem {
	return s.ElasticityAssuranceItem
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) SetElasticityAssuranceItem(v *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem) *DescribeElasticityAssuranceInstancesResponseBody {
	s.ElasticityAssuranceItem = v
	return s
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) SetMaxResults(v int32) *DescribeElasticityAssuranceInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) SetNextToken(v string) *DescribeElasticityAssuranceInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) SetRequestId(v string) *DescribeElasticityAssuranceInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) SetTotalCount(v int32) *DescribeElasticityAssuranceInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesResponseBody) Validate() error {
	if s.ElasticityAssuranceItem != nil {
		if err := s.ElasticityAssuranceItem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem struct {
	InstanceIdSet []*DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItemInstanceIdSet `json:"InstanceIdSet,omitempty" xml:"InstanceIdSet,omitempty" type:"Repeated"`
}

func (s DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem) GetInstanceIdSet() []*DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItemInstanceIdSet {
	return s.InstanceIdSet
}

func (s *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem) SetInstanceIdSet(v []*DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItemInstanceIdSet) *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem {
	s.InstanceIdSet = v
	return s
}

func (s *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItem) Validate() error {
	if s.InstanceIdSet != nil {
		for _, item := range s.InstanceIdSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItemInstanceIdSet struct {
	// The instance ID
	//
	// example:
	//
	// i-bp67acfmxazb4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItemInstanceIdSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItemInstanceIdSet) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItemInstanceIdSet) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItemInstanceIdSet) SetInstanceId(v string) *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItemInstanceIdSet {
	s.InstanceId = &v
	return s
}

func (s *DescribeElasticityAssuranceInstancesResponseBodyElasticityAssuranceItemInstanceIdSet) Validate() error {
	return dara.Validate(s)
}
