// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *AddTagsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *AddTagsRequest
	GetRegionId() *string
	SetResourceId(v string) *AddTagsRequest
	GetResourceId() *string
	SetResourceOwnerAccount(v string) *AddTagsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AddTagsRequest
	GetResourceOwnerId() *int64
	SetResourceType(v string) *AddTagsRequest
	GetResourceType() *string
	SetTag(v []*AddTagsRequestTag) *AddTagsRequest
	GetTag() []*AddTagsRequestTag
}

type AddTagsRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the resource. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource. When the resources are instances, this parameter can be interpreted as InstanceId.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1gtjxuuvwj17zr****
	ResourceId           *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the resource. Valid values:
	//
	// 	- instance: ECS instance
	//
	// 	- disk: disk
	//
	// 	- snapshot: snapshot
	//
	// 	- image: image
	//
	// 	- securitygroup: security group
	//
	// 	- volume: storage volume
	//
	// 	- eni: elastic network interface (ENI)
	//
	// 	- ddh: dedicated host
	//
	// 	- keypair: SSH key pair
	//
	// 	- launchtemplate: launch template
	//
	// 	- reservedinstance: reserved instance
	//
	// 	- snapshotpolicy: automatic snapshot policy
	//
	// All values must be lowercase.
	//
	// This parameter is required.
	//
	// example:
	//
	// instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// The tags.
	//
	// This parameter is required.
	Tag []*AddTagsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s AddTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTagsRequest) GoString() string {
	return s.String()
}

func (s *AddTagsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AddTagsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AddTagsRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *AddTagsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AddTagsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AddTagsRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *AddTagsRequest) GetTag() []*AddTagsRequestTag {
	return s.Tag
}

func (s *AddTagsRequest) SetOwnerId(v int64) *AddTagsRequest {
	s.OwnerId = &v
	return s
}

func (s *AddTagsRequest) SetRegionId(v string) *AddTagsRequest {
	s.RegionId = &v
	return s
}

func (s *AddTagsRequest) SetResourceId(v string) *AddTagsRequest {
	s.ResourceId = &v
	return s
}

func (s *AddTagsRequest) SetResourceOwnerAccount(v string) *AddTagsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AddTagsRequest) SetResourceOwnerId(v int64) *AddTagsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AddTagsRequest) SetResourceType(v string) *AddTagsRequest {
	s.ResourceType = &v
	return s
}

func (s *AddTagsRequest) SetTag(v []*AddTagsRequestTag) *AddTagsRequest {
	s.Tag = v
	return s
}

func (s *AddTagsRequest) Validate() error {
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

type AddTagsRequestTag struct {
	// The key of tag N of the resource. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain [http:// or https://](http://https://).
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the resource. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with acs: or aliyun. It cannot contain [http:// or https://](http://https://).
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s AddTagsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s AddTagsRequestTag) GoString() string {
	return s.String()
}

func (s *AddTagsRequestTag) GetKey() *string {
	return s.Key
}

func (s *AddTagsRequestTag) GetValue() *string {
	return s.Value
}

func (s *AddTagsRequestTag) SetKey(v string) *AddTagsRequestTag {
	s.Key = &v
	return s
}

func (s *AddTagsRequestTag) SetValue(v string) *AddTagsRequestTag {
	s.Value = &v
	return s
}

func (s *AddTagsRequestTag) Validate() error {
	return dara.Validate(s)
}
