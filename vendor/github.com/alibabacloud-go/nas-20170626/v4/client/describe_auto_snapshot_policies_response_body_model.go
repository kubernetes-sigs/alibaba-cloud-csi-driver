// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoSnapshotPoliciesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicies(v *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) *DescribeAutoSnapshotPoliciesResponseBody
	GetAutoSnapshotPolicies() *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies
	SetPageNumber(v int32) *DescribeAutoSnapshotPoliciesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoSnapshotPoliciesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAutoSnapshotPoliciesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAutoSnapshotPoliciesResponseBody
	GetTotalCount() *int32
}

type DescribeAutoSnapshotPoliciesResponseBody struct {
	// The queried automatic snapshot policies.
	AutoSnapshotPolicies *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies `json:"AutoSnapshotPolicies,omitempty" xml:"AutoSnapshotPolicies,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of automatic snapshot policies.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoSnapshotPoliciesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) GetAutoSnapshotPolicies() *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies {
	return s.AutoSnapshotPolicies
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetAutoSnapshotPolicies(v *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) *DescribeAutoSnapshotPoliciesResponseBody {
	s.AutoSnapshotPolicies = v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetPageNumber(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetPageSize(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetRequestId(v string) *DescribeAutoSnapshotPoliciesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) SetTotalCount(v int32) *DescribeAutoSnapshotPoliciesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBody) Validate() error {
	if s.AutoSnapshotPolicies != nil {
		if err := s.AutoSnapshotPolicies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies struct {
	AutoSnapshotPolicy []*DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy `json:"AutoSnapshotPolicy,omitempty" xml:"AutoSnapshotPolicy,omitempty" type:"Repeated"`
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) GetAutoSnapshotPolicy() []*DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	return s.AutoSnapshotPolicy
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) SetAutoSnapshotPolicy(v []*DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies {
	s.AutoSnapshotPolicy = v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPolicies) Validate() error {
	if s.AutoSnapshotPolicy != nil {
		for _, item := range s.AutoSnapshotPolicy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy struct {
	// The ID of the automatic snapshot policy.
	//
	// example:
	//
	// sp-extreme-233e6****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The name of the automatic snapshot policy.
	//
	// example:
	//
	// FinanceJoshua
	AutoSnapshotPolicyName *string `json:"AutoSnapshotPolicyName,omitempty" xml:"AutoSnapshotPolicyName,omitempty"`
	// The time when the automatic snapshot policy was created.
	//
	// The time follows the [ISO8601](https://www.iso.org/iso-8601-date-and-time-format.html) standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2014-04-21T12:08:52Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The number of file systems to which the automatic snapshot policy applies.
	//
	// example:
	//
	// 2
	FileSystemNums *int32 `json:"FileSystemNums,omitempty" xml:"FileSystemNums,omitempty"`
	// The type of the file system.
	//
	// example:
	//
	// extreme
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The region ID of the automatic snapshot policy.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The days of a week on which auto snapshots are created.
	//
	// Auto snapshots are created on a weekly basis.
	//
	// Valid values: 1 to 7. The values from 1 to 7 indicate 7 days in a week from Monday to Sunday.
	//
	// example:
	//
	// 1,5
	RepeatWeekdays *string `json:"RepeatWeekdays,omitempty" xml:"RepeatWeekdays,omitempty"`
	// The retention period of auto snapshots.
	//
	// Unit: days.
	//
	// Valid values:
	//
	// 	- \\-1: Auto snapshots are permanently retained. After the number of auto snapshots exceeds the upper limit, the earliest auto snapshot is automatically deleted.
	//
	// 	- 1 to 65536: Auto snapshots are retained for the specified days. After the retention period of auto snapshots expires, the auto snapshots are automatically deleted.
	//
	// example:
	//
	// 30
	RetentionDays *int32 `json:"RetentionDays,omitempty" xml:"RetentionDays,omitempty"`
	// The status of the automatic snapshot policy.
	//
	// Valid values:
	//
	// 	- Creating: The automatic snapshot policy is being created.
	//
	// 	- Available: The automatic snapshot policy is available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The points in time at which auto snapshots are created.
	//
	// Unit: hours.
	//
	// Valid values: `0 to 23`. The values from 0 to 23 indicate a total of 24 hours from `00:00 to 23:00`. For example, 1 indicates 01:00. A maximum of 24 points in time can be returned. Multiple points in time are separated with commas (,).
	//
	// example:
	//
	// 4,19
	TimePoints *string `json:"TimePoints,omitempty" xml:"TimePoints,omitempty"`
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GoString() string {
	return s.String()
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetAutoSnapshotPolicyName() *string {
	return s.AutoSnapshotPolicyName
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetFileSystemNums() *int32 {
	return s.FileSystemNums
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetRepeatWeekdays() *string {
	return s.RepeatWeekdays
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetRetentionDays() *int32 {
	return s.RetentionDays
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetStatus() *string {
	return s.Status
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) GetTimePoints() *string {
	return s.TimePoints
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetAutoSnapshotPolicyId(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetAutoSnapshotPolicyName(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.AutoSnapshotPolicyName = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetCreateTime(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.CreateTime = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetFileSystemNums(v int32) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.FileSystemNums = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetFileSystemType(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRegionId(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRepeatWeekdays(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RepeatWeekdays = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetRetentionDays(v int32) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.RetentionDays = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetStatus(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.Status = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) SetTimePoints(v string) *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy {
	s.TimePoints = &v
	return s
}

func (s *DescribeAutoSnapshotPoliciesResponseBodyAutoSnapshotPoliciesAutoSnapshotPolicy) Validate() error {
	return dara.Validate(s)
}
