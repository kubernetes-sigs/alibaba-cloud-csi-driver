// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairName(v string) *CreateKeyPairRequest
	GetKeyPairName() *string
	SetOwnerId(v int64) *CreateKeyPairRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateKeyPairRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateKeyPairRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateKeyPairRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateKeyPairRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateKeyPairRequestTag) *CreateKeyPairRequest
	GetTag() []*CreateKeyPairRequestTag
}

type CreateKeyPairRequest struct {
	// The name of the key pair. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which to create the key pair. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to add the key pair.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the key pair.
	Tag []*CreateKeyPairRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyPairRequest) GoString() string {
	return s.String()
}

func (s *CreateKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateKeyPairRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateKeyPairRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateKeyPairRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateKeyPairRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateKeyPairRequest) GetTag() []*CreateKeyPairRequestTag {
	return s.Tag
}

func (s *CreateKeyPairRequest) SetKeyPairName(v string) *CreateKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateKeyPairRequest) SetOwnerId(v int64) *CreateKeyPairRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateKeyPairRequest) SetRegionId(v string) *CreateKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *CreateKeyPairRequest) SetResourceGroupId(v string) *CreateKeyPairRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateKeyPairRequest) SetResourceOwnerAccount(v string) *CreateKeyPairRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateKeyPairRequest) SetResourceOwnerId(v int64) *CreateKeyPairRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateKeyPairRequest) SetTag(v []*CreateKeyPairRequestTag) *CreateKeyPairRequest {
	s.Tag = v
	return s
}

func (s *CreateKeyPairRequest) Validate() error {
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

type CreateKeyPairRequestTag struct {
	// The key of tag N to add to the key pair. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the key pair. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://. The tag value cannot start with acs:.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateKeyPairRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateKeyPairRequestTag) GoString() string {
	return s.String()
}

func (s *CreateKeyPairRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateKeyPairRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateKeyPairRequestTag) SetKey(v string) *CreateKeyPairRequestTag {
	s.Key = &v
	return s
}

func (s *CreateKeyPairRequestTag) SetValue(v string) *CreateKeyPairRequestTag {
	s.Value = &v
	return s
}

func (s *CreateKeyPairRequestTag) Validate() error {
	return dara.Validate(s)
}
