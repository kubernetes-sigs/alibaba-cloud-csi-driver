// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRangeListEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribePortRangeListEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePortRangeListEntriesRequest
	GetOwnerId() *int64
	SetPortRangeListId(v string) *DescribePortRangeListEntriesRequest
	GetPortRangeListId() *string
	SetRegionId(v string) *DescribePortRangeListEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePortRangeListEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePortRangeListEntriesRequest
	GetResourceOwnerId() *int64
}

type DescribePortRangeListEntriesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the port list.
	//
	// This parameter is required.
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The region ID of the port list. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2679950.html) operation to query the most recent region list.
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

func (s DescribePortRangeListEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePortRangeListEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePortRangeListEntriesRequest) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *DescribePortRangeListEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePortRangeListEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePortRangeListEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePortRangeListEntriesRequest) SetOwnerAccount(v string) *DescribePortRangeListEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePortRangeListEntriesRequest) SetOwnerId(v int64) *DescribePortRangeListEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePortRangeListEntriesRequest) SetPortRangeListId(v string) *DescribePortRangeListEntriesRequest {
	s.PortRangeListId = &v
	return s
}

func (s *DescribePortRangeListEntriesRequest) SetRegionId(v string) *DescribePortRangeListEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePortRangeListEntriesRequest) SetResourceOwnerAccount(v string) *DescribePortRangeListEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePortRangeListEntriesRequest) SetResourceOwnerId(v int64) *DescribePortRangeListEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePortRangeListEntriesRequest) Validate() error {
	return dara.Validate(s)
}
