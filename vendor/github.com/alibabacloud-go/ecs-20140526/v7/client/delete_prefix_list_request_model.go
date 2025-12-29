// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePrefixListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeletePrefixListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeletePrefixListRequest
	GetOwnerId() *int64
	SetPrefixListId(v string) *DeletePrefixListRequest
	GetPrefixListId() *string
	SetRegionId(v string) *DeletePrefixListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeletePrefixListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletePrefixListRequest
	GetResourceOwnerId() *int64
}

type DeletePrefixListRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the prefix list.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The region ID of the prefix list. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeletePrefixListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePrefixListRequest) GoString() string {
	return s.String()
}

func (s *DeletePrefixListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeletePrefixListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePrefixListRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *DeletePrefixListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePrefixListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePrefixListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletePrefixListRequest) SetOwnerAccount(v string) *DeletePrefixListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeletePrefixListRequest) SetOwnerId(v int64) *DeletePrefixListRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePrefixListRequest) SetPrefixListId(v string) *DeletePrefixListRequest {
	s.PrefixListId = &v
	return s
}

func (s *DeletePrefixListRequest) SetRegionId(v string) *DeletePrefixListRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePrefixListRequest) SetResourceOwnerAccount(v string) *DeletePrefixListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePrefixListRequest) SetResourceOwnerId(v int64) *DeletePrefixListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletePrefixListRequest) Validate() error {
	return dara.Validate(s)
}
