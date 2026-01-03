// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySnapshotAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifySnapshotAttributeRequest
	GetDescription() *string
	SetDisableInstantAccess(v bool) *ModifySnapshotAttributeRequest
	GetDisableInstantAccess() *bool
	SetOwnerAccount(v string) *ModifySnapshotAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySnapshotAttributeRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifySnapshotAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySnapshotAttributeRequest
	GetResourceOwnerId() *int64
	SetRetentionDays(v int32) *ModifySnapshotAttributeRequest
	GetRetentionDays() *int32
	SetSnapshotId(v string) *ModifySnapshotAttributeRequest
	GetSnapshotId() *string
	SetSnapshotName(v string) *ModifySnapshotAttributeRequest
	GetSnapshotName() *string
}

type ModifySnapshotAttributeRequest struct {
	// The snapshot description. It can be empty or up to 256 characters in length. It cannot start with http:// or https://.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to disable the instant access feature. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  This parameter is no longer used. By default, new standard snapshots of Enterprise SSDs (ESSDs) are upgraded to instant access snapshots free of charge without the need for additional configurations. For more information, see [Use the instant access feature](https://help.aliyun.com/document_detail/193667.html).
	//
	// example:
	//
	// false
	DisableInstantAccess *bool   `json:"DisableInstantAccess,omitempty" xml:"DisableInstantAccess,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The retention period of the snapshot. After you specify this parameter, the end time of the new retention period is the specified number of days apart from the **creation time*	- of the snapshot, which follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format and is displayed in UTC. Valid values: 1 to 65536.
	//
	// >  You can extend the retention period of the snapshot and cannot shorten the retention period.
	//
	// example:
	//
	// 10
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The ID of the snapshot.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-bp199lyny9bb47pa****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The name of the snapshot. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// The name cannot start with auto because snapshots whose names start with auto are recognized as automatic snapshots.
	//
	// example:
	//
	// testSnapshotName
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s ModifySnapshotAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySnapshotAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifySnapshotAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySnapshotAttributeRequest) GetDisableInstantAccess() *bool {
	return s.DisableInstantAccess
}

func (s *ModifySnapshotAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySnapshotAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySnapshotAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySnapshotAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySnapshotAttributeRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *ModifySnapshotAttributeRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ModifySnapshotAttributeRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *ModifySnapshotAttributeRequest) SetDescription(v string) *ModifySnapshotAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetDisableInstantAccess(v bool) *ModifySnapshotAttributeRequest {
	s.DisableInstantAccess = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetOwnerAccount(v string) *ModifySnapshotAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetOwnerId(v int64) *ModifySnapshotAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetResourceOwnerAccount(v string) *ModifySnapshotAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetResourceOwnerId(v int64) *ModifySnapshotAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetRetentionDays(v int32) *ModifySnapshotAttributeRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetSnapshotId(v string) *ModifySnapshotAttributeRequest {
	s.SnapshotId = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) SetSnapshotName(v string) *ModifySnapshotAttributeRequest {
	s.SnapshotName = &v
	return s
}

func (s *ModifySnapshotAttributeRequest) Validate() error {
	return dara.Validate(s)
}
