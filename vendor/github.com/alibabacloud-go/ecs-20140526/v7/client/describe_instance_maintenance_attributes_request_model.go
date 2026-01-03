// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMaintenanceAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v []*string) *DescribeInstanceMaintenanceAttributesRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *DescribeInstanceMaintenanceAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceMaintenanceAttributesRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeInstanceMaintenanceAttributesRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInstanceMaintenanceAttributesRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeInstanceMaintenanceAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceMaintenanceAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceMaintenanceAttributesRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceMaintenanceAttributesRequest struct {
	// The instance IDs. You can specify up to 100 instance IDs.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
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
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeInstanceMaintenanceAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMaintenanceAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMaintenanceAttributesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeInstanceMaintenanceAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceMaintenanceAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceMaintenanceAttributesRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInstanceMaintenanceAttributesRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInstanceMaintenanceAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceMaintenanceAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceMaintenanceAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceMaintenanceAttributesRequest) SetInstanceId(v []*string) *DescribeInstanceMaintenanceAttributesRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesRequest) SetOwnerAccount(v string) *DescribeInstanceMaintenanceAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesRequest) SetOwnerId(v int64) *DescribeInstanceMaintenanceAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesRequest) SetPageNumber(v int64) *DescribeInstanceMaintenanceAttributesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesRequest) SetPageSize(v int64) *DescribeInstanceMaintenanceAttributesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesRequest) SetRegionId(v string) *DescribeInstanceMaintenanceAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesRequest) SetResourceOwnerAccount(v string) *DescribeInstanceMaintenanceAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesRequest) SetResourceOwnerId(v int64) *DescribeInstanceMaintenanceAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesRequest) Validate() error {
	return dara.Validate(s)
}
