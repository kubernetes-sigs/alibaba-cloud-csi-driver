// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateSnapshotRequest
	GetDescription() *string
	SetFileSystemId(v string) *CreateSnapshotRequest
	GetFileSystemId() *string
	SetRetentionDays(v int32) *CreateSnapshotRequest
	GetRetentionDays() *int32
	SetSnapshotName(v string) *CreateSnapshotRequest
	GetSnapshotName() *string
}

type CreateSnapshotRequest struct {
	// The description of the snapshot.
	//
	// Limits:
	//
	// 	- The description must be 2 to 256 characters in length.
	//
	// 	- The description cannot start with `http://` or `https://`.
	//
	// 	- This parameter is empty by default.
	//
	// example:
	//
	// FinanceDepet
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the advanced Extreme NAS file system. The value must start with `extreme-`, for example, `extreme-01dd****`.
	//
	// This parameter is required.
	//
	// example:
	//
	// extreme-01dd****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The retention period of the snapshot.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// 	- \\-1 (default). Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	//
	// 	- 1 to 65536: Auto snapshots are retained for the specified days. After the retention period of auto snapshots expires, the auto snapshots are automatically deleted.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The snapshot name.
	//
	// Limits:
	//
	// 	- The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`.
	//
	// 	- The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// 	- The name cannot start with auto because snapshots whose names start with auto are recognized as auto snapshots.
	//
	// example:
	//
	// FinanceJoshua
	SnapshotName *string `json:"SnapshotName,omitempty" xml:"SnapshotName,omitempty"`
}

func (s CreateSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSnapshotRequest) GoString() string {
	return s.String()
}

func (s *CreateSnapshotRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateSnapshotRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateSnapshotRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *CreateSnapshotRequest) GetSnapshotName() *string {
	return s.SnapshotName
}

func (s *CreateSnapshotRequest) SetDescription(v string) *CreateSnapshotRequest {
	s.Description = &v
	return s
}

func (s *CreateSnapshotRequest) SetFileSystemId(v string) *CreateSnapshotRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateSnapshotRequest) SetRetentionDays(v int32) *CreateSnapshotRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateSnapshotRequest) SetSnapshotName(v string) *CreateSnapshotRequest {
	s.SnapshotName = &v
	return s
}

func (s *CreateSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
