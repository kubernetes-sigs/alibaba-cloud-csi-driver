// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotCategoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ModifySnapshotCategoryRequest
	GetCategory() *string
	SetOwnerAccount(v string) *ModifySnapshotCategoryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySnapshotCategoryRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySnapshotCategoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySnapshotCategoryRequest
	GetResourceOwnerId() *int64
	SetRetentionDays(v int32) *ModifySnapshotCategoryRequest
	GetRetentionDays() *int32
	SetSnapshotId(v string) *ModifySnapshotCategoryRequest
	GetSnapshotId() *string
}

type ModifySnapshotCategoryRequest struct {
	// The type of the snapshot. Valid value:
	//
	// 	- Archive: archive snapshot
	//
	// example:
	//
	// Archive
	Category             *string `json:"Category,omitempty" xml:"Category,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The retention period of the snapshot. Unit: days. The retention period started at the point in time when the snapshot was created. You can archive only standard snapshots that have been retained for at least 14 days.
	//
	// After the snapshot is archived, the minimum retention period (also called minimum archive period) is 60 days. When you calculate the retention period of archived snapshots, you must deduct the retention period of standard snapshots. If you delete the snapshot within 60 days after the snapshot is archived, you are charged archive tier storage fees for the snapshot for 60 days. For more information about the billing of snapshots, see [Snapshots](https://help.aliyun.com/document_detail/56159.html).
	//
	// Value range [74,65536]
	//
	// > If you do not specify this parameter, the snapshot is permanently retained.
	//
	// example:
	//
	// 60
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-123**sd
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ModifySnapshotCategoryRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotCategoryRequest) GoString() string {
	return s.String()
}

func (s *ModifySnapshotCategoryRequest) GetCategory() *string {
	return s.Category
}

func (s *ModifySnapshotCategoryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySnapshotCategoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySnapshotCategoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySnapshotCategoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySnapshotCategoryRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *ModifySnapshotCategoryRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ModifySnapshotCategoryRequest) SetCategory(v string) *ModifySnapshotCategoryRequest {
	s.Category = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetOwnerAccount(v string) *ModifySnapshotCategoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetOwnerId(v int64) *ModifySnapshotCategoryRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetResourceOwnerAccount(v string) *ModifySnapshotCategoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetResourceOwnerId(v int64) *ModifySnapshotCategoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetRetentionDays(v int32) *ModifySnapshotCategoryRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) SetSnapshotId(v string) *ModifySnapshotCategoryRequest {
	s.SnapshotId = &v
	return s
}

func (s *ModifySnapshotCategoryRequest) Validate() error {
	return dara.Validate(s)
}
