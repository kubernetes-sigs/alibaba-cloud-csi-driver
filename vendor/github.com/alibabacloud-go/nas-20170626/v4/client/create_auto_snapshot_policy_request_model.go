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
	// - The name must be 2 to 128 characters in length.
	//
	// - The name must start with a letter or a Chinese character.
	//
	// - The name can contain digits, colons (:), underscores (_), or hyphens (-). It cannot start with `http://` or `https://`.
	//
	// - Default value: empty.
	//
	// example:
	//
	// FinanceJoshua
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	// The type of the file system.
	//
	// Valid values: extreme (Extreme NAS).
	//
	// This parameter is required.
	//
	// example:
	//
	// extreme
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The days of the week on which automatic snapshots are created.
	//
	// Cycle: week.
	//
	// Valid values: 1 to 7, which represent Monday through Sunday.
	//
	// To create automatic snapshots on multiple days in a week, specify multiple values separated by commas (,). You can specify a maximum of 7 values.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1,2,3
	RepeatWeekdays *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	// The retention period of automatic snapshots.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// - -1 (default): Automatic snapshots are permanently retained. When the snapshot quota is reached, the earliest automatic snapshots are automatically deleted.
	//
	// - 1 to 65536: Automatic snapshots are retained for the specified number of days. Snapshots are subject to automatic release after the retention period expires.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The time points at which automatic snapshots are created.
	//
	// Unit: hours.
	//
	// Valid values: 0 to 23, which represent the 24 time points from 00:00 to 23:00. For example, 1 indicates 01:00.
	//
	// To create multiple automatic snapshots within a day, specify multiple time points separated by commas (,). You can specify a maximum of 24 time points.
	//
	// This parameter is required.
	//
	// example:
	//
	// 0,1,…,23
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
