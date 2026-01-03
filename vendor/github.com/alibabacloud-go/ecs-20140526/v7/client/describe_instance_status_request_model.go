// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeInstanceStatusRequest
	GetClusterId() *string
	SetInstanceId(v []*string) *DescribeInstanceStatusRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *DescribeInstanceStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceStatusRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeInstanceStatusRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceStatusRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeInstanceStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceStatusRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeInstanceStatusRequest
	GetZoneId() *string
}

type DescribeInstanceStatusRequest struct {
	// The ID of the cluster to which the ECS instances belong.
	//
	// >  This parameter is deprecated. We recommend that you do not use this parameter.
	//
	// example:
	//
	// cls-bp67acfmxazb4p****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The IDs of ECS instances. You can specify 1 to 100 instance IDs.
	//
	// Example: ["i-bp1j4i2jdf3owlhe\\*\\*\\*\\*", "i-bp1j4i2jdf3o1234\\*\\*\\*\\*"].
	//
	// example:
	//
	// i-bp1j4i2jdf3owlhe****
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
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
	// The zone ID of the instances. You can call the [DescribeZones](https://help.aliyun.com/document_detail/25610.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-d
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeInstanceStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceStatusRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeInstanceStatusRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeInstanceStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceStatusRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceStatusRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceStatusRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeInstanceStatusRequest) SetClusterId(v string) *DescribeInstanceStatusRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetInstanceId(v []*string) *DescribeInstanceStatusRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeInstanceStatusRequest) SetOwnerAccount(v string) *DescribeInstanceStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetOwnerId(v int64) *DescribeInstanceStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetPageNumber(v int32) *DescribeInstanceStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetPageSize(v int32) *DescribeInstanceStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetRegionId(v string) *DescribeInstanceStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetResourceOwnerAccount(v string) *DescribeInstanceStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetResourceOwnerId(v int64) *DescribeInstanceStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceStatusRequest) SetZoneId(v string) *DescribeInstanceStatusRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeInstanceStatusRequest) Validate() error {
	return dara.Validate(s)
}
