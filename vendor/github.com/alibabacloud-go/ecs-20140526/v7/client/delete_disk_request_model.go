// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDiskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DeleteDiskRequest
	GetDiskId() *string
	SetOwnerAccount(v string) *DeleteDiskRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDiskRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteDiskRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDiskRequest
	GetResourceOwnerId() *int64
}

type DeleteDiskRequest struct {
	// The ID of the disk that you want to release.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp14k9cxvr5uzy5****
	DiskId               *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DeleteDiskRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDiskRequest) GoString() string {
	return s.String()
}

func (s *DeleteDiskRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DeleteDiskRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDiskRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDiskRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDiskRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDiskRequest) SetDiskId(v string) *DeleteDiskRequest {
	s.DiskId = &v
	return s
}

func (s *DeleteDiskRequest) SetOwnerAccount(v string) *DeleteDiskRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDiskRequest) SetOwnerId(v int64) *DeleteDiskRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDiskRequest) SetResourceOwnerAccount(v string) *DeleteDiskRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDiskRequest) SetResourceOwnerId(v int64) *DeleteDiskRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDiskRequest) Validate() error {
	return dara.Validate(s)
}
