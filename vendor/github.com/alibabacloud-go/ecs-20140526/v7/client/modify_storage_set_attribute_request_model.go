// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStorageSetAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyStorageSetAttributeRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyStorageSetAttributeRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyStorageSetAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyStorageSetAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyStorageSetAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyStorageSetAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyStorageSetAttributeRequest
	GetResourceOwnerId() *int64
	SetStorageSetId(v string) *ModifyStorageSetAttributeRequest
	GetStorageSetId() *string
	SetStorageSetName(v string) *ModifyStorageSetAttributeRequest
	GetStorageSetName() *string
}

type ModifyStorageSetAttributeRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the storage set.
	//
	// example:
	//
	// testStorageSetDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// hide
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 111
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the storage set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// hide
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// example:
	//
	// 111
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the storage set.
	//
	// This parameter is required.
	//
	// example:
	//
	// ss-bp67acfmxazb4ph****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The name of the storage set.
	//
	// example:
	//
	// testStorageSetName
	StorageSetName *string `json:"StorageSetName,omitempty" xml:"StorageSetName,omitempty"`
}

func (s ModifyStorageSetAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStorageSetAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyStorageSetAttributeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyStorageSetAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyStorageSetAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyStorageSetAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyStorageSetAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyStorageSetAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyStorageSetAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyStorageSetAttributeRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *ModifyStorageSetAttributeRequest) GetStorageSetName() *string {
	return s.StorageSetName
}

func (s *ModifyStorageSetAttributeRequest) SetClientToken(v string) *ModifyStorageSetAttributeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyStorageSetAttributeRequest) SetDescription(v string) *ModifyStorageSetAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyStorageSetAttributeRequest) SetOwnerAccount(v string) *ModifyStorageSetAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyStorageSetAttributeRequest) SetOwnerId(v int64) *ModifyStorageSetAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyStorageSetAttributeRequest) SetRegionId(v string) *ModifyStorageSetAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyStorageSetAttributeRequest) SetResourceOwnerAccount(v string) *ModifyStorageSetAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyStorageSetAttributeRequest) SetResourceOwnerId(v int64) *ModifyStorageSetAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyStorageSetAttributeRequest) SetStorageSetId(v string) *ModifyStorageSetAttributeRequest {
	s.StorageSetId = &v
	return s
}

func (s *ModifyStorageSetAttributeRequest) SetStorageSetName(v string) *ModifyStorageSetAttributeRequest {
	s.StorageSetName = &v
	return s
}

func (s *ModifyStorageSetAttributeRequest) Validate() error {
	return dara.Validate(s)
}
