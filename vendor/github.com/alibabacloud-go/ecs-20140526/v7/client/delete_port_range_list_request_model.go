// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePortRangeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DeletePortRangeListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeletePortRangeListRequest
	GetOwnerId() *int64
	SetPortRangeListId(v string) *DeletePortRangeListRequest
	GetPortRangeListId() *string
	SetRegionId(v string) *DeletePortRangeListRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeletePortRangeListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeletePortRangeListRequest
	GetResourceOwnerId() *int64
}

type DeletePortRangeListRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the port list.
	//
	// >  If the port list is associated with other resources, you cannot delete the port list. You must disassociate the port list from the resources and then delete the port list.
	//
	// This parameter is required.
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The region ID of the port list. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DeletePortRangeListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePortRangeListRequest) GoString() string {
	return s.String()
}

func (s *DeletePortRangeListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeletePortRangeListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeletePortRangeListRequest) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *DeletePortRangeListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeletePortRangeListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeletePortRangeListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeletePortRangeListRequest) SetOwnerAccount(v string) *DeletePortRangeListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeletePortRangeListRequest) SetOwnerId(v int64) *DeletePortRangeListRequest {
	s.OwnerId = &v
	return s
}

func (s *DeletePortRangeListRequest) SetPortRangeListId(v string) *DeletePortRangeListRequest {
	s.PortRangeListId = &v
	return s
}

func (s *DeletePortRangeListRequest) SetRegionId(v string) *DeletePortRangeListRequest {
	s.RegionId = &v
	return s
}

func (s *DeletePortRangeListRequest) SetResourceOwnerAccount(v string) *DeletePortRangeListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeletePortRangeListRequest) SetResourceOwnerId(v int64) *DeletePortRangeListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeletePortRangeListRequest) Validate() error {
	return dara.Validate(s)
}
