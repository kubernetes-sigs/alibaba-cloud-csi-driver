// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCapacityReservationInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCapacityReservationItem(v *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem) *DescribeCapacityReservationInstancesResponseBody
	GetCapacityReservationItem() *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem
	SetMaxResults(v int32) *DescribeCapacityReservationInstancesResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCapacityReservationInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeCapacityReservationInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeCapacityReservationInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeCapacityReservationInstancesResponseBody struct {
	// Details about the instances that match the capacity reservation.
	CapacityReservationItem *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem `json:"CapacityReservationItem,omitempty" xml:"CapacityReservationItem,omitempty" type:"Struct"`
	// The maximum number of entries per page.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 10
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeCapacityReservationInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationInstancesResponseBody) GetCapacityReservationItem() *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem {
	return s.CapacityReservationItem
}

func (s *DescribeCapacityReservationInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCapacityReservationInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCapacityReservationInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCapacityReservationInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeCapacityReservationInstancesResponseBody) SetCapacityReservationItem(v *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem) *DescribeCapacityReservationInstancesResponseBody {
	s.CapacityReservationItem = v
	return s
}

func (s *DescribeCapacityReservationInstancesResponseBody) SetMaxResults(v int32) *DescribeCapacityReservationInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeCapacityReservationInstancesResponseBody) SetNextToken(v string) *DescribeCapacityReservationInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeCapacityReservationInstancesResponseBody) SetRequestId(v string) *DescribeCapacityReservationInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCapacityReservationInstancesResponseBody) SetTotalCount(v int32) *DescribeCapacityReservationInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeCapacityReservationInstancesResponseBody) Validate() error {
	if s.CapacityReservationItem != nil {
		if err := s.CapacityReservationItem.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem struct {
	InstanceIdSet []*DescribeCapacityReservationInstancesResponseBodyCapacityReservationItemInstanceIdSet `json:"InstanceIdSet,omitempty" xml:"InstanceIdSet,omitempty" type:"Repeated"`
}

func (s DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem) GetInstanceIdSet() []*DescribeCapacityReservationInstancesResponseBodyCapacityReservationItemInstanceIdSet {
	return s.InstanceIdSet
}

func (s *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem) SetInstanceIdSet(v []*DescribeCapacityReservationInstancesResponseBodyCapacityReservationItemInstanceIdSet) *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem {
	s.InstanceIdSet = v
	return s
}

func (s *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItem) Validate() error {
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

type DescribeCapacityReservationInstancesResponseBodyCapacityReservationItemInstanceIdSet struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-bp67acfmxazb4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeCapacityReservationInstancesResponseBodyCapacityReservationItemInstanceIdSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeCapacityReservationInstancesResponseBodyCapacityReservationItemInstanceIdSet) GoString() string {
	return s.String()
}

func (s *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItemInstanceIdSet) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItemInstanceIdSet) SetInstanceId(v string) *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItemInstanceIdSet {
	s.InstanceId = &v
	return s
}

func (s *DescribeCapacityReservationInstancesResponseBodyCapacityReservationItemInstanceIdSet) Validate() error {
	return dara.Validate(s)
}
