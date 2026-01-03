// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForce(v bool) *DeleteSnapshotRequest
	GetForce() *bool
	SetOwnerAccount(v string) *DeleteSnapshotRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteSnapshotRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteSnapshotRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteSnapshotRequest
	GetResourceOwnerId() *int64
	SetSnapshotId(v string) *DeleteSnapshotRequest
	GetSnapshotId() *string
}

type DeleteSnapshotRequest struct {
	// Specifies whether to force delete the snapshot that has been used to create cloud disks. Valid values:
	//
	// 	- true: force deletes the snapshot. After the snapshot is force deleted, the cloud disks created from the snapshot cannot be re-initialized.
	//
	// 	- false: does not force delete the snapshot.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Force                *bool   `json:"Force,omitempty" xml:"Force,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp1c0doj0taqyzzl****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DeleteSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSnapshotRequest) GoString() string {
	return s.String()
}

func (s *DeleteSnapshotRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteSnapshotRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteSnapshotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteSnapshotRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteSnapshotRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DeleteSnapshotRequest) SetForce(v bool) *DeleteSnapshotRequest {
	s.Force = &v
	return s
}

func (s *DeleteSnapshotRequest) SetOwnerAccount(v string) *DeleteSnapshotRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteSnapshotRequest) SetOwnerId(v int64) *DeleteSnapshotRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetResourceOwnerAccount(v string) *DeleteSnapshotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteSnapshotRequest) SetResourceOwnerId(v int64) *DeleteSnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteSnapshotRequest) SetSnapshotId(v string) *DeleteSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *DeleteSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
