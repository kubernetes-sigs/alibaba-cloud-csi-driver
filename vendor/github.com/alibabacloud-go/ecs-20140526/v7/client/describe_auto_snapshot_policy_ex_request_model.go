// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPolicyExRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyExRequest
	GetAutoSnapshotPolicyId() *string
	SetAutoSnapshotPolicyName(v string) *DescribeAutoSnapshotPolicyExRequest
	GetAutoSnapshotPolicyName() *string
	SetOwnerAccount(v string) *DescribeAutoSnapshotPolicyExRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAutoSnapshotPolicyExRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAutoSnapshotPolicyExRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoSnapshotPolicyExRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAutoSnapshotPolicyExRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeAutoSnapshotPolicyExRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoSnapshotPolicyExRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoSnapshotPolicyExRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeAutoSnapshotPolicyExRequestTag) *DescribeAutoSnapshotPolicyExRequest
	GetTag() []*DescribeAutoSnapshotPolicyExRequestTag
}

type DescribeAutoSnapshotPolicyExRequest struct {
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-bp67acfmxazb4ph****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The name of the automatic snapshot policy.
	//
	// example:
	//
	// TestName
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	OwnerAccount           *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The region ID of the automatic snapshot policy. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If this parameter is specified to query resources, up to 1,000 resources that belong to the specified resource group can be displayed in the response.
	//
	// > Resources in the default resource group are displayed in the response regardless of how this parameter is set.
	//
	// example:
	//
	// rg-aek2kkmhmhs****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the automatic snapshot policy.
	Tag []*DescribeAutoSnapshotPolicyExRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotPolicyExRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetAutoSnapshotPolicyName() *string {
	return s.AutoSnapshotPolicyName
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoSnapshotPolicyExRequest) GetTag() []*DescribeAutoSnapshotPolicyExRequestTag {
	return s.Tag
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetAutoSnapshotPolicyName(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetOwnerAccount(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetOwnerId(v int64) *DescribeAutoSnapshotPolicyExRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetPageNumber(v int32) *DescribeAutoSnapshotPolicyExRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetPageSize(v int32) *DescribeAutoSnapshotPolicyExRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetRegionId(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetResourceGroupId(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetResourceOwnerAccount(v string) *DescribeAutoSnapshotPolicyExRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetResourceOwnerId(v int64) *DescribeAutoSnapshotPolicyExRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) SetTag(v []*DescribeAutoSnapshotPolicyExRequestTag) *DescribeAutoSnapshotPolicyExRequest {
	s.Tag = v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequest) Validate() error {
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

type DescribeAutoSnapshotPolicyExRequestTag struct {
	// The key of tag N of the automatic snapshot policy. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the automatic snapshot policy. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://. The tag value cannot start with acs:.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAutoSnapshotPolicyExRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPolicyExRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPolicyExRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeAutoSnapshotPolicyExRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeAutoSnapshotPolicyExRequestTag) SetKey(v string) *DescribeAutoSnapshotPolicyExRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequestTag) SetValue(v string) *DescribeAutoSnapshotPolicyExRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeAutoSnapshotPolicyExRequestTag) Validate() error {
	return dara.Validate(s)
}
