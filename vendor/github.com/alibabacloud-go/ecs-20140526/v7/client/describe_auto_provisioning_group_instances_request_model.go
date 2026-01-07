// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoProvisioningGroupInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoProvisioningGroupId(v string) *DescribeAutoProvisioningGroupInstancesRequest
	GetAutoProvisioningGroupId() *string
	SetOwnerAccount(v string) *DescribeAutoProvisioningGroupInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAutoProvisioningGroupInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAutoProvisioningGroupInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoProvisioningGroupInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAutoProvisioningGroupInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoProvisioningGroupInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoProvisioningGroupInstancesRequest
	GetResourceOwnerId() *int64
}

type DescribeAutoProvisioningGroupInstancesRequest struct {
	// The ID of the auto provisioning group.
	//
	// This parameter is required.
	//
	// example:
	//
	// apg-uf6jel2bbl62wh13****
	AutoProvisioningGroupId *string `json:"AutoProvisioningGroupId,omitempty" xml:"AutoProvisioningGroupId,omitempty"`
	// example:
	//
	// 123456
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 123456
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the auto provisioning group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 123456
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// example:
	//
	// 123456
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeAutoProvisioningGroupInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) GetAutoProvisioningGroupId() *string {
	return s.AutoProvisioningGroupId
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) SetAutoProvisioningGroupId(v string) *DescribeAutoProvisioningGroupInstancesRequest {
	s.AutoProvisioningGroupId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) SetOwnerAccount(v string) *DescribeAutoProvisioningGroupInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) SetOwnerId(v int64) *DescribeAutoProvisioningGroupInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) SetPageNumber(v int32) *DescribeAutoProvisioningGroupInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) SetPageSize(v int32) *DescribeAutoProvisioningGroupInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) SetRegionId(v string) *DescribeAutoProvisioningGroupInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) SetResourceOwnerAccount(v string) *DescribeAutoProvisioningGroupInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) SetResourceOwnerId(v int64) *DescribeAutoProvisioningGroupInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesRequest) Validate() error {
	return dara.Validate(s)
}
