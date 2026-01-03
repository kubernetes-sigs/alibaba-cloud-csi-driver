// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClassicLinkInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLinks(v *DescribeClassicLinkInstancesResponseBodyLinks) *DescribeClassicLinkInstancesResponseBody
	GetLinks() *DescribeClassicLinkInstancesResponseBodyLinks
	SetPageNumber(v int32) *DescribeClassicLinkInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeClassicLinkInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeClassicLinkInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeClassicLinkInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeClassicLinkInstancesResponseBody struct {
	// The details of the ClassicLink connections between the instances reside in the classic network and VPCs.
	Links *DescribeClassicLinkInstancesResponseBodyLinks `json:"Links,omitempty" xml:"Links,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of ClassicLink connections.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeClassicLinkInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassicLinkInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClassicLinkInstancesResponseBody) GetLinks() *DescribeClassicLinkInstancesResponseBodyLinks {
	return s.Links
}

func (s *DescribeClassicLinkInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeClassicLinkInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeClassicLinkInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClassicLinkInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeClassicLinkInstancesResponseBody) SetLinks(v *DescribeClassicLinkInstancesResponseBodyLinks) *DescribeClassicLinkInstancesResponseBody {
	s.Links = v
	return s
}

func (s *DescribeClassicLinkInstancesResponseBody) SetPageNumber(v int32) *DescribeClassicLinkInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeClassicLinkInstancesResponseBody) SetPageSize(v int32) *DescribeClassicLinkInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeClassicLinkInstancesResponseBody) SetRequestId(v string) *DescribeClassicLinkInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClassicLinkInstancesResponseBody) SetTotalCount(v int32) *DescribeClassicLinkInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeClassicLinkInstancesResponseBody) Validate() error {
	if s.Links != nil {
		if err := s.Links.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeClassicLinkInstancesResponseBodyLinks struct {
	Link []*DescribeClassicLinkInstancesResponseBodyLinksLink `json:"Link,omitempty" xml:"Link,omitempty" type:"Repeated"`
}

func (s DescribeClassicLinkInstancesResponseBodyLinks) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassicLinkInstancesResponseBodyLinks) GoString() string {
	return s.String()
}

func (s *DescribeClassicLinkInstancesResponseBodyLinks) GetLink() []*DescribeClassicLinkInstancesResponseBodyLinksLink {
	return s.Link
}

func (s *DescribeClassicLinkInstancesResponseBodyLinks) SetLink(v []*DescribeClassicLinkInstancesResponseBodyLinksLink) *DescribeClassicLinkInstancesResponseBodyLinks {
	s.Link = v
	return s
}

func (s *DescribeClassicLinkInstancesResponseBodyLinks) Validate() error {
	if s.Link != nil {
		for _, item := range s.Link {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeClassicLinkInstancesResponseBodyLinksLink struct {
	// The instance ID.
	//
	// example:
	//
	// i-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-test
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeClassicLinkInstancesResponseBodyLinksLink) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassicLinkInstancesResponseBodyLinksLink) GoString() string {
	return s.String()
}

func (s *DescribeClassicLinkInstancesResponseBodyLinksLink) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeClassicLinkInstancesResponseBodyLinksLink) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClassicLinkInstancesResponseBodyLinksLink) SetInstanceId(v string) *DescribeClassicLinkInstancesResponseBodyLinksLink {
	s.InstanceId = &v
	return s
}

func (s *DescribeClassicLinkInstancesResponseBodyLinksLink) SetVpcId(v string) *DescribeClassicLinkInstancesResponseBodyLinksLink {
	s.VpcId = &v
	return s
}

func (s *DescribeClassicLinkInstancesResponseBodyLinksLink) Validate() error {
	return dara.Validate(s)
}
