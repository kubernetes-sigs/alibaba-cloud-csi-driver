// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClassicLinkInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeClassicLinkInstancesRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *DescribeClassicLinkInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v string) *DescribeClassicLinkInstancesRequest
	GetPageNumber() *string
	SetPageSize(v string) *DescribeClassicLinkInstancesRequest
	GetPageSize() *string
	SetRegionId(v string) *DescribeClassicLinkInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeClassicLinkInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClassicLinkInstancesRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DescribeClassicLinkInstancesRequest
	GetVpcId() *string
}

type DescribeClassicLinkInstancesRequest struct {
	// The instance ID. You can specify a maximum of 100 instance IDs in a single request. Separate the instance IDs with commas (,).
	//
	// example:
	//
	// i-bp1a5zr3u7nq9cxh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instances. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The VPC ID. The ClassicLink feature must be enabled for the specified VPC. For more information, see [Establish a ClassicLink connection](https://help.aliyun.com/document_detail/65413.html).
	//
	// example:
	//
	// vpc-bp1vwnn14rqpyiczj****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeClassicLinkInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClassicLinkInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeClassicLinkInstancesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeClassicLinkInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClassicLinkInstancesRequest) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeClassicLinkInstancesRequest) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeClassicLinkInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClassicLinkInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClassicLinkInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClassicLinkInstancesRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeClassicLinkInstancesRequest) SetInstanceId(v string) *DescribeClassicLinkInstancesRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeClassicLinkInstancesRequest) SetOwnerId(v int64) *DescribeClassicLinkInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClassicLinkInstancesRequest) SetPageNumber(v string) *DescribeClassicLinkInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeClassicLinkInstancesRequest) SetPageSize(v string) *DescribeClassicLinkInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeClassicLinkInstancesRequest) SetRegionId(v string) *DescribeClassicLinkInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClassicLinkInstancesRequest) SetResourceOwnerAccount(v string) *DescribeClassicLinkInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClassicLinkInstancesRequest) SetResourceOwnerId(v int64) *DescribeClassicLinkInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClassicLinkInstancesRequest) SetVpcId(v string) *DescribeClassicLinkInstancesRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeClassicLinkInstancesRequest) Validate() error {
	return dara.Validate(s)
}
