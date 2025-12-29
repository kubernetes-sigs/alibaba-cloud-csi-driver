// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteStorageSetRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DeleteStorageSetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteStorageSetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteStorageSetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteStorageSetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteStorageSetRequest
	GetResourceOwnerId() *int64
	SetStorageSetId(v string) *DeleteStorageSetRequest
	GetStorageSetId() *string
}

type DeleteStorageSetRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
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
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
}

func (s DeleteStorageSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteStorageSetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteStorageSetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteStorageSetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteStorageSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteStorageSetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteStorageSetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteStorageSetRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *DeleteStorageSetRequest) SetClientToken(v string) *DeleteStorageSetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteStorageSetRequest) SetOwnerAccount(v string) *DeleteStorageSetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteStorageSetRequest) SetOwnerId(v int64) *DeleteStorageSetRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteStorageSetRequest) SetRegionId(v string) *DeleteStorageSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteStorageSetRequest) SetResourceOwnerAccount(v string) *DeleteStorageSetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteStorageSetRequest) SetResourceOwnerId(v int64) *DeleteStorageSetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteStorageSetRequest) SetStorageSetId(v string) *DeleteStorageSetRequest {
	s.StorageSetId = &v
	return s
}

func (s *DeleteStorageSetRequest) Validate() error {
	return dara.Validate(s)
}
