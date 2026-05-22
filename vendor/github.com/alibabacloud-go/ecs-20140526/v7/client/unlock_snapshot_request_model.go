// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnlockSnapshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *UnlockSnapshotRequest
	GetClientToken() *string
	SetDryRun(v bool) *UnlockSnapshotRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *UnlockSnapshotRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *UnlockSnapshotRequest
	GetOwnerId() *int64
	SetRegionId(v string) *UnlockSnapshotRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *UnlockSnapshotRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *UnlockSnapshotRequest
	GetResourceOwnerId() *int64
	SetSnapshotId(v string) *UnlockSnapshotRequest
	GetSnapshotId() *string
}

type UnlockSnapshotRequest struct {
	// A client token that is used to ensure the idempotence of the request. You can use the client to generate a client token. Make sure that a unique client token is used for each request. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/zh/ecs/developer-reference/how-to-ensure-idempotence?spm=a2c4g.11186623.0.0.2a29d467Bh2sO5).
	//
	// example:
	//
	// 5EC38E7D-389F-1925-ABE2-D7925A8F****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform the dry run. Valid values:
	//
	// 	- true: The request is checked and is not executed. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and service limits. If the check fails, an error is returned. If the check is passed, the error code DryRunOperation is returned.
	//
	// 	- false (default): sends the request. If the request passes the check, the request is directly executed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// example:
	//
	// 158704318252****
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 158704318252****
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/zh/ecs/developer-reference/api-ecs-2014-05-26-describeregions?spm=a2c4g.11186623.0.i2) operation to query the most recent region list.
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

func (s UnlockSnapshotRequest) String() string {
	return dara.Prettify(s)
}

func (s UnlockSnapshotRequest) GoString() string {
	return s.String()
}

func (s *UnlockSnapshotRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *UnlockSnapshotRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UnlockSnapshotRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *UnlockSnapshotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *UnlockSnapshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UnlockSnapshotRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *UnlockSnapshotRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *UnlockSnapshotRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *UnlockSnapshotRequest) SetClientToken(v string) *UnlockSnapshotRequest {
	s.ClientToken = &v
	return s
}

func (s *UnlockSnapshotRequest) SetDryRun(v bool) *UnlockSnapshotRequest {
	s.DryRun = &v
	return s
}

func (s *UnlockSnapshotRequest) SetOwnerAccount(v string) *UnlockSnapshotRequest {
	s.OwnerAccount = &v
	return s
}

func (s *UnlockSnapshotRequest) SetOwnerId(v int64) *UnlockSnapshotRequest {
	s.OwnerId = &v
	return s
}

func (s *UnlockSnapshotRequest) SetRegionId(v string) *UnlockSnapshotRequest {
	s.RegionId = &v
	return s
}

func (s *UnlockSnapshotRequest) SetResourceOwnerAccount(v string) *UnlockSnapshotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *UnlockSnapshotRequest) SetResourceOwnerId(v int64) *UnlockSnapshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *UnlockSnapshotRequest) SetSnapshotId(v string) *UnlockSnapshotRequest {
	s.SnapshotId = &v
	return s
}

func (s *UnlockSnapshotRequest) Validate() error {
	return dara.Validate(s)
}
