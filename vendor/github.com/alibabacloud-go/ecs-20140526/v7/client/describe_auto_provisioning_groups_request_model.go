// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoProvisioningGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoProvisioningGroupId(v []*string) *DescribeAutoProvisioningGroupsRequest
	GetAutoProvisioningGroupId() []*string
	SetAutoProvisioningGroupName(v string) *DescribeAutoProvisioningGroupsRequest
	GetAutoProvisioningGroupName() *string
	SetAutoProvisioningGroupStatus(v []*string) *DescribeAutoProvisioningGroupsRequest
	GetAutoProvisioningGroupStatus() []*string
	SetOwnerAccount(v string) *DescribeAutoProvisioningGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAutoProvisioningGroupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAutoProvisioningGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoProvisioningGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAutoProvisioningGroupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeAutoProvisioningGroupsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoProvisioningGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoProvisioningGroupsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeAutoProvisioningGroupsRequestTag) *DescribeAutoProvisioningGroupsRequest
	GetTag() []*DescribeAutoProvisioningGroupsRequestTag
}

type DescribeAutoProvisioningGroupsRequest struct {
	// The ID of the auto provisioning group. You can specify up to 20 IDs.
	//
	// example:
	//
	// apg-sn54avj8htgvtyh8****
	AutoProvisioningGroupId []*string `json:"AutoProvisioningGroupId,omitempty" xml:"AutoProvisioningGroupId,omitempty" type:"Repeated"`
	// The name of the auto provisioning group.
	//
	// example:
	//
	// testAutoProvisioningGroupName
	AutoProvisioningGroupName *string `json:"AutoProvisioningGroupName,omitempty" xml:"AutoProvisioningGroupName,omitempty"`
	// The status of the auto provisioning group.
	//
	// example:
	//
	// active
	AutoProvisioningGroupStatus []*string `json:"AutoProvisioningGroupStatus,omitempty" xml:"AutoProvisioningGroupStatus,omitempty" type:"Repeated"`
	OwnerAccount                *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the auto provisioning group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the auto provisioning group belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags that are added to the auto provisioning group.
	Tag []*DescribeAutoProvisioningGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsRequest) GetAutoProvisioningGroupId() []*string {
	return s.AutoProvisioningGroupId
}

func (s *DescribeAutoProvisioningGroupsRequest) GetAutoProvisioningGroupName() *string {
	return s.AutoProvisioningGroupName
}

func (s *DescribeAutoProvisioningGroupsRequest) GetAutoProvisioningGroupStatus() []*string {
	return s.AutoProvisioningGroupStatus
}

func (s *DescribeAutoProvisioningGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAutoProvisioningGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoProvisioningGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoProvisioningGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoProvisioningGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoProvisioningGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAutoProvisioningGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoProvisioningGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoProvisioningGroupsRequest) GetTag() []*DescribeAutoProvisioningGroupsRequestTag {
	return s.Tag
}

func (s *DescribeAutoProvisioningGroupsRequest) SetAutoProvisioningGroupId(v []*string) *DescribeAutoProvisioningGroupsRequest {
	s.AutoProvisioningGroupId = v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetAutoProvisioningGroupName(v string) *DescribeAutoProvisioningGroupsRequest {
	s.AutoProvisioningGroupName = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetAutoProvisioningGroupStatus(v []*string) *DescribeAutoProvisioningGroupsRequest {
	s.AutoProvisioningGroupStatus = v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetOwnerAccount(v string) *DescribeAutoProvisioningGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetOwnerId(v int64) *DescribeAutoProvisioningGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetPageNumber(v int32) *DescribeAutoProvisioningGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetPageSize(v int32) *DescribeAutoProvisioningGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetRegionId(v string) *DescribeAutoProvisioningGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetResourceGroupId(v string) *DescribeAutoProvisioningGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetResourceOwnerAccount(v string) *DescribeAutoProvisioningGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetResourceOwnerId(v int64) *DescribeAutoProvisioningGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) SetTag(v []*DescribeAutoProvisioningGroupsRequestTag) *DescribeAutoProvisioningGroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupsRequestTag struct {
	// The key of tag N that is added to the auto provisioning group.
	//
	// Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N that is added to the auto provisioning group.
	//
	// Valid values of N: 1 to 20. The tag value can be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAutoProvisioningGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAutoProvisioningGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAutoProvisioningGroupsRequestTag) SetKey(v string) *DescribeAutoProvisioningGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequestTag) SetValue(v string) *DescribeAutoProvisioningGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeAutoProvisioningGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
