// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAdditionalAttributes(v []*string) *DescribeSnapshotGroupsRequest
	GetAdditionalAttributes() []*string
	SetInstanceId(v string) *DescribeSnapshotGroupsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeSnapshotGroupsRequest
	GetMaxResults() *int32
	SetName(v string) *DescribeSnapshotGroupsRequest
	GetName() *string
	SetNextToken(v string) *DescribeSnapshotGroupsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeSnapshotGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSnapshotGroupsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSnapshotGroupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSnapshotGroupsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSnapshotGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSnapshotGroupsRequest
	GetResourceOwnerId() *int64
	SetSnapshotGroupId(v []*string) *DescribeSnapshotGroupsRequest
	GetSnapshotGroupId() []*string
	SetStatus(v []*string) *DescribeSnapshotGroupsRequest
	GetStatus() []*string
	SetTag(v []*DescribeSnapshotGroupsRequestTag) *DescribeSnapshotGroupsRequest
	GetTag() []*DescribeSnapshotGroupsRequestTag
}

type DescribeSnapshotGroupsRequest struct {
	// This parameter is not publicly available.
	//
	// example:
	//
	// hide
	AdditionalAttributes []*string `json:"AdditionalAttributes,omitempty" xml:"AdditionalAttributes,omitempty" type:"Repeated"`
	// The ID of the instance.
	//
	// example:
	//
	// i-j6ca469urv8ei629****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the snapshot-consistent group.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The token that determines the start point of the next query. Set the value to the NextToken value that is returned from the last call.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a4883
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the snapshot-consistent group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the snapshot-consistent group belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of snapshot-consistent group N. Valid values of N: 1 to 10.
	//
	// example:
	//
	// ssg-j6ciyh3k52qp7ovm****
	SnapshotGroupId []*string `json:"SnapshotGroupId,omitempty" xml:"SnapshotGroupId,omitempty" type:"Repeated"`
	// The state of snapshot-consistent group N. Valid values of the second N: 1, 2, and 3. Valid values:
	//
	// 	- progressing: The snapshot-consistent group is being created.
	//
	// 	- accomplished: The snapshot-consistent group is created.
	//
	// 	- failed: The snapshot-consistent group fails to be created.
	//
	// example:
	//
	// accomplished
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The tags of the snapshot-consistent group.
	Tag []*DescribeSnapshotGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSnapshotGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsRequest) GetAdditionalAttributes() []*string {
	return s.AdditionalAttributes
}

func (s *DescribeSnapshotGroupsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSnapshotGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSnapshotGroupsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSnapshotGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSnapshotGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSnapshotGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSnapshotGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSnapshotGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSnapshotGroupsRequest) GetSnapshotGroupId() []*string {
	return s.SnapshotGroupId
}

func (s *DescribeSnapshotGroupsRequest) GetStatus() []*string {
	return s.Status
}

func (s *DescribeSnapshotGroupsRequest) GetTag() []*DescribeSnapshotGroupsRequestTag {
	return s.Tag
}

func (s *DescribeSnapshotGroupsRequest) SetAdditionalAttributes(v []*string) *DescribeSnapshotGroupsRequest {
	s.AdditionalAttributes = v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetInstanceId(v string) *DescribeSnapshotGroupsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetMaxResults(v int32) *DescribeSnapshotGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetName(v string) *DescribeSnapshotGroupsRequest {
	s.Name = &v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetNextToken(v string) *DescribeSnapshotGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetOwnerAccount(v string) *DescribeSnapshotGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetOwnerId(v int64) *DescribeSnapshotGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetRegionId(v string) *DescribeSnapshotGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetResourceGroupId(v string) *DescribeSnapshotGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetResourceOwnerAccount(v string) *DescribeSnapshotGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetResourceOwnerId(v int64) *DescribeSnapshotGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetSnapshotGroupId(v []*string) *DescribeSnapshotGroupsRequest {
	s.SnapshotGroupId = v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetStatus(v []*string) *DescribeSnapshotGroupsRequest {
	s.Status = v
	return s
}

func (s *DescribeSnapshotGroupsRequest) SetTag(v []*DescribeSnapshotGroupsRequestTag) *DescribeSnapshotGroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeSnapshotGroupsRequest) Validate() error {
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

type DescribeSnapshotGroupsRequestTag struct {
	// The key of tag N of the snapshot-consistent group. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the snapshot-consistent group. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSnapshotGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSnapshotGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSnapshotGroupsRequestTag) SetKey(v string) *DescribeSnapshotGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeSnapshotGroupsRequestTag) SetValue(v string) *DescribeSnapshotGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeSnapshotGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
