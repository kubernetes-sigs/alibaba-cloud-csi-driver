// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportKeyPairRequest interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairName(v string) *ImportKeyPairRequest
	GetKeyPairName() *string
	SetOwnerId(v int64) *ImportKeyPairRequest
	GetOwnerId() *int64
	SetPublicKeyBody(v string) *ImportKeyPairRequest
	GetPublicKeyBody() *string
	SetRegionId(v string) *ImportKeyPairRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *ImportKeyPairRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *ImportKeyPairRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ImportKeyPairRequest
	GetResourceOwnerId() *int64
	SetTag(v []*ImportKeyPairRequestTag) *ImportKeyPairRequest
	GetTag() []*ImportKeyPairRequestTag
}

type ImportKeyPairRequest struct {
	// The name of the key pair. The name must be unique. It must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	OwnerId     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The public key of the key pair.
	//
	// This parameter is required.
	//
	// example:
	//
	// ABC1234567
	PublicKeyBody *string `json:"PublicKeyBody,omitempty" xml:"PublicKeyBody,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the enterprise resource group to which the SSH key pair belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags to add to the key pair.
	Tag []*ImportKeyPairRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ImportKeyPairRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyPairRequest) GoString() string {
	return s.String()
}

func (s *ImportKeyPairRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ImportKeyPairRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ImportKeyPairRequest) GetPublicKeyBody() *string {
	return s.PublicKeyBody
}

func (s *ImportKeyPairRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ImportKeyPairRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ImportKeyPairRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ImportKeyPairRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ImportKeyPairRequest) GetTag() []*ImportKeyPairRequestTag {
	return s.Tag
}

func (s *ImportKeyPairRequest) SetKeyPairName(v string) *ImportKeyPairRequest {
	s.KeyPairName = &v
	return s
}

func (s *ImportKeyPairRequest) SetOwnerId(v int64) *ImportKeyPairRequest {
	s.OwnerId = &v
	return s
}

func (s *ImportKeyPairRequest) SetPublicKeyBody(v string) *ImportKeyPairRequest {
	s.PublicKeyBody = &v
	return s
}

func (s *ImportKeyPairRequest) SetRegionId(v string) *ImportKeyPairRequest {
	s.RegionId = &v
	return s
}

func (s *ImportKeyPairRequest) SetResourceGroupId(v string) *ImportKeyPairRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ImportKeyPairRequest) SetResourceOwnerAccount(v string) *ImportKeyPairRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ImportKeyPairRequest) SetResourceOwnerId(v int64) *ImportKeyPairRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ImportKeyPairRequest) SetTag(v []*ImportKeyPairRequestTag) *ImportKeyPairRequest {
	s.Tag = v
	return s
}

func (s *ImportKeyPairRequest) Validate() error {
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

type ImportKeyPairRequestTag struct {
	// The key of tag N to add to the key pair. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain [http:// or https://](http://https://。). The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the key pair. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain [http:// or https://](http://https://。). The tag value cannot start with acs:.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ImportKeyPairRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ImportKeyPairRequestTag) GoString() string {
	return s.String()
}

func (s *ImportKeyPairRequestTag) GetKey() *string {
	return s.Key
}

func (s *ImportKeyPairRequestTag) GetValue() *string {
	return s.Value
}

func (s *ImportKeyPairRequestTag) SetKey(v string) *ImportKeyPairRequestTag {
	s.Key = &v
	return s
}

func (s *ImportKeyPairRequestTag) SetValue(v string) *ImportKeyPairRequestTag {
	s.Value = &v
	return s
}

func (s *ImportKeyPairRequestTag) Validate() error {
	return dara.Validate(s)
}
