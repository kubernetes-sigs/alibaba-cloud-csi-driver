// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyId(v string) *ModifyAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyId() *string
	SetAutoSnapshotPolicyName(v string) *ModifyAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyName() *string
	SetRepeatWeekdays(v string) *ModifyAutoSnapshotPolicyRequest
	GetRepeatWeekdays() *string
	SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyRequest
	GetRetentionDays() *int32
	SetTimePoints(v string) *ModifyAutoSnapshotPolicyRequest
	GetTimePoints() *string
}

type ModifyAutoSnapshotPolicyRequest struct {
	// The ID of the automatic snapshot policy.
	//
	// You can call the DescribeAutoSnapshotPolicies operation to view available automatic snapshot policies.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-extreme-233e6****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The name of the automatic snapshot policy. If you do not specify this parameter, the policy name is not changed.
	//
	// Limits:
	//
	// 	- The name must be 2 to 128 characters in length.
	//
	// 	- The name must start with a letter.
	//
	// 	- The name can contain digits, letters, colons (:), underscores (_), and hyphens (-). It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// FinanceJoshua
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	// The days of a week on which auto snapshots are created.
	//
	// Cycle: week.
	//
	// Valid values: 1 to 7. The value 1 indicates Monday. If you want to create multiple auto snapshots within a week, you can specify multiple days from Monday to Sunday and separate the days with commas (,). You can specify a maximum of seven days.
	//
	// example:
	//
	// 1,7
	RepeatWeekdays *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	// The retention period of auto snapshots.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// 	- \\-1 (default): Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	//
	// 	- 1 to 65536: Auto snapshots are retained for the specified number of days. After the retention period of auto snapshots expires, the auto snapshots are automatically deleted.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The points in time at which auto snapshots are created.
	//
	// Unit: hours.
	//
	// Valid values: 0 to 23. The values from 0 to 23 indicate a total of 24 hours from 00:00 to 23:00. For example, the value 1 indicates 01:00. If you want to create multiple auto snapshots within a day, you can specify multiple points in time and separate the points in time with commas (,). You can specify a maximum of 24 points in time.
	//
	// example:
	//
	// 0,1
	TimePoints *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s ModifyAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *ModifyAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyName() *string {
	return s.AutoSnapshotPolicyName
}

func (s *ModifyAutoSnapshotPolicyRequest) GetRepeatWeekdays() *string {
	return s.RepeatWeekdays
}

func (s *ModifyAutoSnapshotPolicyRequest) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *ModifyAutoSnapshotPolicyRequest) GetTimePoints() *string {
	return s.TimePoints
}

func (s *ModifyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *ModifyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyName(v string) *ModifyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRepeatWeekdays(v string) *ModifyAutoSnapshotPolicyRequest {
	s.RepeatWeekdays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetRetentionDays(v int32) *ModifyAutoSnapshotPolicyRequest {
	s.RetentionDays = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) SetTimePoints(v string) *ModifyAutoSnapshotPolicyRequest {
	s.TimePoints = &v
	return s
}

func (s *ModifyAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
