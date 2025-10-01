// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyName(v string) *CreateAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyName() *string
	SetFileSystemType(v string) *CreateAutoSnapshotPolicyRequest
	GetFileSystemType() *string
	SetRepeatWeekdays(v string) *CreateAutoSnapshotPolicyRequest
	GetRepeatWeekdays() *string
	SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest
	GetRetentionDays() *int32
	SetTimePoints(v string) *CreateAutoSnapshotPolicyRequest
	GetTimePoints() *string
}

type CreateAutoSnapshotPolicyRequest struct {
	// The name of the automatic snapshot policy.
	//
	// Limits:
	//
	// 	- The name must be 2 to 128 characters in length.
	//
	// 	- The name must start with a letter.
	//
	// 	- The name can contain digits, colons (:), underscores (_), and hyphens (-). It cannot start with `http://` or `https://`.
	//
	// 	- This parameter is empty by default.
	//
	// example:
	//
	// FinanceJoshua
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	// The type of the file system.
	//
	// Valid value: extreme, which indicates Extreme NAS file systems.
	//
	// This parameter is required.
	//
	// example:
	//
	// extreme
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The days of a week on which to create automatic snapshots.
	//
	// Cycle: week.
	//
	// Valid values: 1 to 7. The values from 1 to 7 indicate the seven days in a week from Monday to Sunday.
	//
	// If you want to create multiple auto snapshots within a week, you can specify multiple days from Monday to Sunday and separate the days with commas (,). You can specify a maximum of seven days.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3
	RepeatWeekdays *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	// The retention period of auto snapshots.
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
	// The points in time at which auto snapshots were created.
	//
	// Unit: hours.
	//
	// Valid values: 0 to 23. The values from 0 to 23 indicate a total of 24 hours from 00:00 to 23:00. For example, the value 1 indicates 01:00.
	//
	// If you want to create multiple auto snapshots within a day, you can specify multiple points in time and separate the points in time with commas (,). You can specify a maximum of 24 points in time.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0,1,…23
	TimePoints *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s CreateAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyName() *string {
	return s.AutoSnapshotPolicyName
}

func (s *CreateAutoSnapshotPolicyRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *CreateAutoSnapshotPolicyRequest) GetRepeatWeekdays() *string {
	return s.RepeatWeekdays
}

func (s *CreateAutoSnapshotPolicyRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *CreateAutoSnapshotPolicyRequest) GetTimePoints() *string {
	return s.TimePoints
}

func (s *CreateAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyName(v string) *CreateAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetFileSystemType(v string) *CreateAutoSnapshotPolicyRequest {
	s.FileSystemType = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRepeatWeekdays(v string) *CreateAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *CreateAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) SetTimePoints(v string) *CreateAutoSnapshotPolicyRequest {
	s.TimePoints = &v
	return s
}

func (s *CreateAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
