// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRouteEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationCidrBlock(v string) *DeleteRouteEntryRequest
	GetDestinationCidrBlock() *string
	SetNextHopId(v string) *DeleteRouteEntryRequest
	GetNextHopId() *string
	SetNextHopList(v []*DeleteRouteEntryRequestNextHopList) *DeleteRouteEntryRequest
	GetNextHopList() []*DeleteRouteEntryRequestNextHopList
	SetOwnerAccount(v string) *DeleteRouteEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteRouteEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteRouteEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteRouteEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteRouteEntryRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *DeleteRouteEntryRequest
	GetRouteTableId() *string
}

type DeleteRouteEntryRequest struct {
	// This parameter is required.
	DestinationCidrBlock *string                               `json:"DestinationCidrBlock,omitempty" xml:"DestinationCidrBlock,omitempty"`
	NextHopId            *string                               `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopList          []*DeleteRouteEntryRequestNextHopList `json:"NextHopList,omitempty" xml:"NextHopList,omitempty" type:"Repeated"`
	OwnerAccount         *string                               `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64                                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string                               `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	RouteTableId *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
}

func (s DeleteRouteEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntryRequest) GetDestinationCidrBlock() *string {
	return s.DestinationCidrBlock
}

func (s *DeleteRouteEntryRequest) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DeleteRouteEntryRequest) GetNextHopList() []*DeleteRouteEntryRequestNextHopList {
	return s.NextHopList
}

func (s *DeleteRouteEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteRouteEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteRouteEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteRouteEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteRouteEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteRouteEntryRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DeleteRouteEntryRequest) SetDestinationCidrBlock(v string) *DeleteRouteEntryRequest {
	s.DestinationCidrBlock = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetNextHopId(v string) *DeleteRouteEntryRequest {
	s.NextHopId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetNextHopList(v []*DeleteRouteEntryRequestNextHopList) *DeleteRouteEntryRequest {
	s.NextHopList = v
	return s
}

func (s *DeleteRouteEntryRequest) SetOwnerAccount(v string) *DeleteRouteEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetOwnerId(v int64) *DeleteRouteEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetRegionId(v string) *DeleteRouteEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetResourceOwnerAccount(v string) *DeleteRouteEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetResourceOwnerId(v int64) *DeleteRouteEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteRouteEntryRequest) SetRouteTableId(v string) *DeleteRouteEntryRequest {
	s.RouteTableId = &v
	return s
}

func (s *DeleteRouteEntryRequest) Validate() error {
	if s.NextHopList != nil {
		for _, item := range s.NextHopList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DeleteRouteEntryRequestNextHopList struct {
	NextHopId   *string `json:"NextHopId,omitempty" xml:"NextHopId,omitempty"`
	NextHopType *string `json:"NextHopType,omitempty" xml:"NextHopType,omitempty"`
}

func (s DeleteRouteEntryRequestNextHopList) String() string {
	return dara.Prettify(s)
}

func (s DeleteRouteEntryRequestNextHopList) GoString() string {
	return s.String()
}

func (s *DeleteRouteEntryRequestNextHopList) GetNextHopId() *string {
	return s.NextHopId
}

func (s *DeleteRouteEntryRequestNextHopList) GetNextHopType() *string {
	return s.NextHopType
}

func (s *DeleteRouteEntryRequestNextHopList) SetNextHopId(v string) *DeleteRouteEntryRequestNextHopList {
	s.NextHopId = &v
	return s
}

func (s *DeleteRouteEntryRequestNextHopList) SetNextHopType(v string) *DeleteRouteEntryRequestNextHopList {
	s.NextHopType = &v
	return s
}

func (s *DeleteRouteEntryRequestNextHopList) Validate() error {
	return dara.Validate(s)
}
