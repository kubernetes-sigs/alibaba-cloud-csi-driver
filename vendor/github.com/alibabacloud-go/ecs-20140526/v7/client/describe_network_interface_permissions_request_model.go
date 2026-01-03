// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfacePermissionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkInterfaceId(v string) *DescribeNetworkInterfacePermissionsRequest
	GetNetworkInterfaceId() *string
	SetNetworkInterfacePermissionId(v []*string) *DescribeNetworkInterfacePermissionsRequest
	GetNetworkInterfacePermissionId() []*string
	SetOwnerAccount(v string) *DescribeNetworkInterfacePermissionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeNetworkInterfacePermissionsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeNetworkInterfacePermissionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNetworkInterfacePermissionsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeNetworkInterfacePermissionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeNetworkInterfacePermissionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNetworkInterfacePermissionsRequest
	GetResourceOwnerId() *int64
}

type DescribeNetworkInterfacePermissionsRequest struct {
	// The ID of ENI N. You must specify `NetworkInterfaceId` or `NetworkInterfacePermissionId.N` to determine the query range.
	//
	// example:
	//
	// eni-bp17pdijfczax****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The IDs of ENI permissions. You can specify up to 100 ENI permission IDs.
	//
	// example:
	//
	// eni-perm-bp1cs4lwn56lfb****
	NetworkInterfacePermissionId []*string `json:"NetworkInterfacePermissionId,omitempty" xml:"NetworkInterfacePermissionId,omitempty" type:"Repeated"`
	OwnerAccount                 *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the ENI permission. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeNetworkInterfacePermissionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfacePermissionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfacePermissionsRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeNetworkInterfacePermissionsRequest) GetNetworkInterfacePermissionId() []*string {
	return s.NetworkInterfacePermissionId
}

func (s *DescribeNetworkInterfacePermissionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeNetworkInterfacePermissionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNetworkInterfacePermissionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNetworkInterfacePermissionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNetworkInterfacePermissionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkInterfacePermissionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNetworkInterfacePermissionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNetworkInterfacePermissionsRequest) SetNetworkInterfaceId(v string) *DescribeNetworkInterfacePermissionsRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsRequest) SetNetworkInterfacePermissionId(v []*string) *DescribeNetworkInterfacePermissionsRequest {
	s.NetworkInterfacePermissionId = v
	return s
}

func (s *DescribeNetworkInterfacePermissionsRequest) SetOwnerAccount(v string) *DescribeNetworkInterfacePermissionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsRequest) SetOwnerId(v int64) *DescribeNetworkInterfacePermissionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsRequest) SetPageNumber(v int32) *DescribeNetworkInterfacePermissionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsRequest) SetPageSize(v int32) *DescribeNetworkInterfacePermissionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsRequest) SetRegionId(v string) *DescribeNetworkInterfacePermissionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsRequest) SetResourceOwnerAccount(v string) *DescribeNetworkInterfacePermissionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsRequest) SetResourceOwnerId(v int64) *DescribeNetworkInterfacePermissionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNetworkInterfacePermissionsRequest) Validate() error {
	return dara.Validate(s)
}
