// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRangeListAssociationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribePortRangeListAssociationsResponseBody
	GetNextToken() *string
	SetPortRangeListAssociations(v []*DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations) *DescribePortRangeListAssociationsResponseBody
	GetPortRangeListAssociations() []*DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations
	SetRequestId(v string) *DescribePortRangeListAssociationsResponseBody
	GetRequestId() *string
}

type DescribePortRangeListAssociationsResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results. If the return value is empty, no more data is returned.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The resources that are associated with the port list.
	PortRangeListAssociations []*DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations `json:"PortRangeListAssociations,omitempty" xml:"PortRangeListAssociations,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortRangeListAssociationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListAssociationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListAssociationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePortRangeListAssociationsResponseBody) GetPortRangeListAssociations() []*DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations {
	return s.PortRangeListAssociations
}

func (s *DescribePortRangeListAssociationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortRangeListAssociationsResponseBody) SetNextToken(v string) *DescribePortRangeListAssociationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePortRangeListAssociationsResponseBody) SetPortRangeListAssociations(v []*DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations) *DescribePortRangeListAssociationsResponseBody {
	s.PortRangeListAssociations = v
	return s
}

func (s *DescribePortRangeListAssociationsResponseBody) SetRequestId(v string) *DescribePortRangeListAssociationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortRangeListAssociationsResponseBody) Validate() error {
	if s.PortRangeListAssociations != nil {
		for _, item := range s.PortRangeListAssociations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations struct {
	// The ID of the resource.
	//
	// example:
	//
	// sg-2zefu72****
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// The type of the resource. Valid value: SecurityGroup.
	//
	// example:
	//
	// SecurityGroup
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations) SetResourceId(v string) *DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations {
	s.ResourceId = &v
	return s
}

func (s *DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations) SetResourceType(v string) *DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations {
	s.ResourceType = &v
	return s
}

func (s *DescribePortRangeListAssociationsResponseBodyPortRangeListAssociations) Validate() error {
	return dara.Validate(s)
}
