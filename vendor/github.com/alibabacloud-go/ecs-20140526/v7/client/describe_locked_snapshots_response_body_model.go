// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLockedSnapshotsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLockedSnapshotsInfo(v []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) *DescribeLockedSnapshotsResponseBody
	GetLockedSnapshotsInfo() []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo
	SetNextToken(v string) *DescribeLockedSnapshotsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeLockedSnapshotsResponseBody
	GetRequestId() *string
}

type DescribeLockedSnapshotsResponseBody struct {
	// Details of the locked snapshots.
	LockedSnapshotsInfo []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo `json:"LockedSnapshotsInfo,omitempty" xml:"LockedSnapshotsInfo,omitempty" type:"Repeated"`
	// The returned pagination token which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLockedSnapshotsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLockedSnapshotsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLockedSnapshotsResponseBody) GetLockedSnapshotsInfo() []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	return s.LockedSnapshotsInfo
}

func (s *DescribeLockedSnapshotsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeLockedSnapshotsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLockedSnapshotsResponseBody) SetLockedSnapshotsInfo(v []*DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) *DescribeLockedSnapshotsResponseBody {
	s.LockedSnapshotsInfo = v
	return s
}

func (s *DescribeLockedSnapshotsResponseBody) SetNextToken(v string) *DescribeLockedSnapshotsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBody) SetRequestId(v string) *DescribeLockedSnapshotsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBody) Validate() error {
	if s.LockedSnapshotsInfo != nil {
		for _, item := range s.LockedSnapshotsInfo {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo struct {
	// The cooling-off period of the compliance mode. Unit: hours.
	//
	// example:
	//
	// 3
	CoolOffPeriod *int32 `json:"CoolOffPeriod,omitempty" xml:"CoolOffPeriod,omitempty"`
	// The end time of the cooling-off period in compliance mode. The time follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard in the yyyy-MM-ddTHH:mm:ssZ format (in UTC).
	//
	// example:
	//
	// 2025-10-15T13:00:00Z
	CoolOffPeriodExpiredTime *string `json:"CoolOffPeriodExpiredTime,omitempty" xml:"CoolOffPeriodExpiredTime,omitempty"`
	// The date and time at which the snapshot is locked. The time follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard in the yyyy-MM-ddTHH:mm:ssZ format (in UTC).
	//
	// example:
	//
	// 2025-10-15T10:00:00Z
	LockCreationTime *string `json:"LockCreationTime,omitempty" xml:"LockCreationTime,omitempty"`
	// The lock duration. The snapshot lock automatically expires after the lock duration ends. Unit: days.
	//
	// example:
	//
	// 1
	LockDuration *int32 `json:"LockDuration,omitempty" xml:"LockDuration,omitempty"`
	// The start time of the lock duration. The time follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard in the yyyy-MM-ddTHH:mm:ssZ format (in UTC). If you lock a snapshot that is in the Progressing state, the lock time is not calculated until the snapshot enters the Accomplished state.
	//
	// example:
	//
	// 2025-10-15T10:00:00Z
	LockDurationStartTime *string `json:"LockDurationStartTime,omitempty" xml:"LockDurationStartTime,omitempty"`
	// The time when the lock expires. The time follows the [ISO 8601](https://help.aliyun.com/zh/ecs/developer-reference/iso-8601-time-format?spm=a2c4g.11186623.0.0.277c6c92kl7kXM) standard in the yyyy-MM-ddTHH:mm:ssZ format (in UTC).
	//
	// example:
	//
	// 2025-10-16T10:00:00Z
	LockExpiredTime *string `json:"LockExpiredTime,omitempty" xml:"LockExpiredTime,omitempty"`
	// example:
	//
	// compliance
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// The lock status. Valid values:
	//
	// 	- compliance-cooloff: The snapshot is locked in compliance mode but is still in a cooling-off period. Snapshots cannot be deleted. However, users with the corresponding RAM permissions can unlock snapshots, extend or shorten the cooling-off period, and extend or shorten the lock duration.
	//
	// 	- compliance: The snapshot is locked in compliance mode and the cooling-off period has ended. Snapshots cannot be unlocked or deleted. However, users with the corresponding RAM permissions can extend the lock duration.
	//
	// 	- expired: The snapshot was once locked, but the lock duration has ended and the lock has expired. The snapshot is not locked and can be deleted.
	//
	// example:
	//
	// compliance-cooloff
	LockStatus *string `json:"LockStatus,omitempty" xml:"LockStatus,omitempty"`
	// The snapshot ID.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) String() string {
	return dara.Prettify(s)
}

func (s DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GoString() string {
	return s.String()
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetCoolOffPeriod() *int32 {
	return s.CoolOffPeriod
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetCoolOffPeriodExpiredTime() *string {
	return s.CoolOffPeriodExpiredTime
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockCreationTime() *string {
	return s.LockCreationTime
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockDuration() *int32 {
	return s.LockDuration
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockDurationStartTime() *string {
	return s.LockDurationStartTime
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockExpiredTime() *string {
	return s.LockExpiredTime
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockMode() *string {
	return s.LockMode
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetLockStatus() *string {
	return s.LockStatus
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetCoolOffPeriod(v int32) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.CoolOffPeriod = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetCoolOffPeriodExpiredTime(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.CoolOffPeriodExpiredTime = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockCreationTime(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockCreationTime = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockDuration(v int32) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockDuration = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockDurationStartTime(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockDurationStartTime = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockExpiredTime(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockExpiredTime = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockMode(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockMode = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetLockStatus(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.LockStatus = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) SetSnapshotId(v string) *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo {
	s.SnapshotId = &v
	return s
}

func (s *DescribeLockedSnapshotsResponseBodyLockedSnapshotsInfo) Validate() error {
	return dara.Validate(s)
}
