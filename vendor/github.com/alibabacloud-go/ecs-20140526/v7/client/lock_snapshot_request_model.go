// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLockSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *LockSnapshotRequest
	GetClientToken() *string
	SetCoolOffPeriod(v int32) *LockSnapshotRequest
	GetCoolOffPeriod() *int32
	SetDryRun(v bool) *LockSnapshotRequest
	GetDryRun() *bool
	SetLockDuration(v int32) *LockSnapshotRequest
	GetLockDuration() *int32
	SetLockMode(v string) *LockSnapshotRequest
	GetLockMode() *string
	SetOwnerAccount(v string) *LockSnapshotRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *LockSnapshotRequest
	GetOwnerId() *int64
	SetRegionId(v string) *LockSnapshotRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *LockSnapshotRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *LockSnapshotRequest
	GetResourceOwnerId() *int64
	SetSnapshotId(v string) *LockSnapshotRequest
	GetSnapshotId() *string
}

type LockSnapshotRequest struct {
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate a client token. Make sure that a unique client token is used for each request. ClientToken only supports ASCII characters and cannot exceed 64 characters. For more information, see [How to ensure idempotence](https://help.aliyun.com/zh/ecs/developer-reference/how-to-ensure-idempotence?spm=a2c4g.11186623.0.0.2a29d467Bh2sO5).
	//
	// example:
	//
	// 5EC38E7D-389F-1925-ABE2-D7925A8F****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Cooling-off period. In compliance mode, you can set a cooling-off period or skip the cooling-off period to directly lock the snapshot.
	//
	// During the cooling-off period, users with corresponding RAM permissions can unlock snapshots, extend or shorten the cooling-off period, and extend or shorten the lock duration. Snapshots cannot be deleted during the cooling-off period.
	//
	// After the cooling-off period ends, only extending the lock duration is supported.
	//
	// Unit: hours.
	//
	// Valid values: 0 to 72. A value of 0 indicates skipping the cooling-off period and directly locking the snapshot.
	//
	// If the snapshot has entered the compliance mode lock period, set this parameter to 0 when extending the lock duration.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3
	CoolOffPeriod *int32 `json:"CoolOffPeriod,omitempty" xml:"CoolOffPeriod,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// 	- true: The request is checked and is not executed. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the check fails, the corresponding error is returned. If the check passes, the error code DryRunOperation is returned.
	//
	// 	- false (default): Sends a normal request, checks it, and executes the request directly if it passes the check.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Lock duration. After the lock duration ends, the snapshot lock will automatically expire.
	//
	// Unit: days.
	//
	// Valid values: 1 to 36500.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	LockDuration *int32 `json:"LockDuration,omitempty" xml:"LockDuration,omitempty"`
	// The lock mode. Valid values:
	//
	// 	- compliance: The snapshot is locked in compliance mode. A snapshot that is locked in compliance mode cannot be unlocked by any user. It can be deleted only after the lock duration expires. Users cannot shorten the lock duration, but users with the corresponding RAM permissions can extend the lock duration at any time. When locking a snapshot in compliance mode, you can optionally specify a cooling-off period.
	//
	// This parameter is required.
	//
	// example:
	//
	// compliance
	LockMode *string `json:"LockMode,omitempty" xml:"LockMode,omitempty"`
	// example:
	//
	// 158704318252****
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 158704318252****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID You can call the [DescribeRegions](https://help.aliyun.com/zh/ecs/developer-reference/api-ecs-2014-05-26-describeregions?spm=a2c4g.11186623.0.i2) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 158704318252****
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// example:
	//
	// 158704318252****
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-9dp2qojdpdfmgfmf****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s LockSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s LockSnapshotRequest) GoString() string {
	return s.String()
}

func (s *LockSnapshotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *LockSnapshotRequest) GetCoolOffPeriod() *int32 {
	return s.CoolOffPeriod
}

func (s *LockSnapshotRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *LockSnapshotRequest) GetLockDuration() *int32 {
	return s.LockDuration
}

func (s *LockSnapshotRequest) GetLockMode() *string {
	return s.LockMode
}

func (s *LockSnapshotRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *LockSnapshotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *LockSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *LockSnapshotRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *LockSnapshotRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *LockSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *LockSnapshotRequest) SetClientToken(v string) *LockSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *LockSnapshotRequest) SetCoolOffPeriod(v int32) *LockSnapshotRequest {
	s.CoolOffPeriod = &v
	return s
}

func (s *LockSnapshotRequest) SetDryRun(v bool) *LockSnapshotRequest {
	s.DryRun = &v
	return s
}

func (s *LockSnapshotRequest) SetLockDuration(v int32) *LockSnapshotRequest {
	s.LockDuration = &v
	return s
}

func (s *LockSnapshotRequest) SetLockMode(v string) *LockSnapshotRequest {
	s.LockMode = &v
	return s
}

func (s *LockSnapshotRequest) SetOwnerAccount(v string) *LockSnapshotRequest {
	s.OwnerAccount = &v
	return s
}

func (s *LockSnapshotRequest) SetOwnerId(v int64) *LockSnapshotRequest {
	s.OwnerId = &v
	return s
}

func (s *LockSnapshotRequest) SetRegionId(v string) *LockSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *LockSnapshotRequest) SetResourceOwnerAccount(v string) *LockSnapshotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *LockSnapshotRequest) SetResourceOwnerId(v int64) *LockSnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *LockSnapshotRequest) SetSnapshotId(v string) *LockSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *LockSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
