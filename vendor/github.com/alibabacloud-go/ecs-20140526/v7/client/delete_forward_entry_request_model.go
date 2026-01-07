// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteForwardEntryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardEntryId(v string) *DeleteForwardEntryRequest
	GetForwardEntryId() *string
	SetForwardTableId(v string) *DeleteForwardEntryRequest
	GetForwardTableId() *string
	SetOwnerAccount(v string) *DeleteForwardEntryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteForwardEntryRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteForwardEntryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteForwardEntryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteForwardEntryRequest
	GetResourceOwnerId() *int64
}

type DeleteForwardEntryRequest struct {
	// This parameter is required.
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// This parameter is required.
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteForwardEntryRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteForwardEntryRequest) GoString() string {
	return s.String()
}

func (s *DeleteForwardEntryRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DeleteForwardEntryRequest) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *DeleteForwardEntryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteForwardEntryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteForwardEntryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteForwardEntryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteForwardEntryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteForwardEntryRequest) SetForwardEntryId(v string) *DeleteForwardEntryRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *DeleteForwardEntryRequest) SetForwardTableId(v string) *DeleteForwardEntryRequest {
	s.ForwardTableId = &v
	return s
}

func (s *DeleteForwardEntryRequest) SetOwnerAccount(v string) *DeleteForwardEntryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteForwardEntryRequest) SetOwnerId(v int64) *DeleteForwardEntryRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteForwardEntryRequest) SetRegionId(v string) *DeleteForwardEntryRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteForwardEntryRequest) SetResourceOwnerAccount(v string) *DeleteForwardEntryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteForwardEntryRequest) SetResourceOwnerId(v int64) *DeleteForwardEntryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteForwardEntryRequest) Validate() error {
	return dara.Validate(s)
}
