// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployDedicatedHostRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostId(v string) *RedeployDedicatedHostRequest
	GetDedicatedHostId() *string
	SetMigrationType(v string) *RedeployDedicatedHostRequest
	GetMigrationType() *string
	SetOwnerAccount(v string) *RedeployDedicatedHostRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RedeployDedicatedHostRequest
	GetOwnerId() *int64
	SetRegionId(v string) *RedeployDedicatedHostRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RedeployDedicatedHostRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RedeployDedicatedHostRequest
	GetResourceOwnerId() *int64
}

type RedeployDedicatedHostRequest struct {
	// The ID of the dedicated host.
	//
	// This parameter is required.
	//
	// example:
	//
	// dh-bp165p6xk2tlw61e****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// Specifies whether to stop the instance before it is migrated to the destination dedicated host. Valid values:
	//
	// 	- reboot: stops the instance before migration.
	//
	// 	- LiveMigrationFirst: migrates the instance without stopping it. If you set MigrationType to LiveMigrationFirst, you must specify DedicatedHostId. In this case, you cannot change the instance type of the ECS instance when the instance is migrated. If the migration in LiveMigrationFirst mode fails, the system switches to the Reboot mode.
	//
	// Default value: reboot.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// Reboot
	MigrationType *string `json:"MigrationType,omitempty" xml:"MigrationType,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the dedicated host. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s RedeployDedicatedHostRequest) String() string {
	return dara.Prettify(s)
}

func (s RedeployDedicatedHostRequest) GoString() string {
	return s.String()
}

func (s *RedeployDedicatedHostRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *RedeployDedicatedHostRequest) GetMigrationType() *string {
	return s.MigrationType
}

func (s *RedeployDedicatedHostRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RedeployDedicatedHostRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RedeployDedicatedHostRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RedeployDedicatedHostRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RedeployDedicatedHostRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RedeployDedicatedHostRequest) SetDedicatedHostId(v string) *RedeployDedicatedHostRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *RedeployDedicatedHostRequest) SetMigrationType(v string) *RedeployDedicatedHostRequest {
	s.MigrationType = &v
	return s
}

func (s *RedeployDedicatedHostRequest) SetOwnerAccount(v string) *RedeployDedicatedHostRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RedeployDedicatedHostRequest) SetOwnerId(v int64) *RedeployDedicatedHostRequest {
	s.OwnerId = &v
	return s
}

func (s *RedeployDedicatedHostRequest) SetRegionId(v string) *RedeployDedicatedHostRequest {
	s.RegionId = &v
	return s
}

func (s *RedeployDedicatedHostRequest) SetResourceOwnerAccount(v string) *RedeployDedicatedHostRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RedeployDedicatedHostRequest) SetResourceOwnerId(v int64) *RedeployDedicatedHostRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RedeployDedicatedHostRequest) Validate() error {
	return dara.Validate(s)
}
