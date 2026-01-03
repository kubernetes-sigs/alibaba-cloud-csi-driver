// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskDeploymentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskCategory(v string) *ModifyDiskDeploymentRequest
	GetDiskCategory() *string
	SetDiskId(v string) *ModifyDiskDeploymentRequest
	GetDiskId() *string
	SetDryRun(v bool) *ModifyDiskDeploymentRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ModifyDiskDeploymentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDiskDeploymentRequest
	GetOwnerId() *int64
	SetPerformanceLevel(v string) *ModifyDiskDeploymentRequest
	GetPerformanceLevel() *string
	SetResourceOwnerAccount(v string) *ModifyDiskDeploymentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDiskDeploymentRequest
	GetResourceOwnerId() *int64
	SetStorageClusterId(v string) *ModifyDiskDeploymentRequest
	GetStorageClusterId() *string
}

type ModifyDiskDeploymentRequest struct {
	// The new disk category. The parameter can be used only when you migrate a disk between dedicated block storage clusters. Only ESSDs can be created in dedicated block storage clusters. Set this parameter to cloud_essd.
	//
	// This parameter is empty by default, which indicates that the disk category is not changed.
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The ID of the disk.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp131n0q38u3a4zi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and unavailable ECS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DryRunOperation error code is returned.
	//
	// 	- false: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun       *bool   `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The new performance level of the ESSD. The parameter can be used only when you migrate data between dedicated block storage clusters. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50000 random read/write IOPS.
	//
	// This parameter is empty by default, which indicates that the performance level is not modified.
	//
	// example:
	//
	// PL1
	PerformanceLevel     *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the dedicated block storage cluster to which data disk N belongs.
	//
	// 	- When you migrate a disk to a dedicated block storage cluster, the`  StorageClusterId ` parameter must be specified.
	//
	// 	- When you migrate a disk to a public block storage cluster, the `StorageClusterId` parameter must be left empty.
	//
	// This parameter is empty by default, which indicates that the disk is migrated to a public block storage cluster.
	//
	// example:
	//
	// dbsc-cn-c4d2uea****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s ModifyDiskDeploymentRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskDeploymentRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskDeploymentRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *ModifyDiskDeploymentRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyDiskDeploymentRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDiskDeploymentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDiskDeploymentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDiskDeploymentRequest) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyDiskDeploymentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDiskDeploymentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDiskDeploymentRequest) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *ModifyDiskDeploymentRequest) SetDiskCategory(v string) *ModifyDiskDeploymentRequest {
	s.DiskCategory = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetDiskId(v string) *ModifyDiskDeploymentRequest {
	s.DiskId = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetDryRun(v bool) *ModifyDiskDeploymentRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetOwnerAccount(v string) *ModifyDiskDeploymentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetOwnerId(v int64) *ModifyDiskDeploymentRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetPerformanceLevel(v string) *ModifyDiskDeploymentRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetResourceOwnerAccount(v string) *ModifyDiskDeploymentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetResourceOwnerId(v int64) *ModifyDiskDeploymentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) SetStorageClusterId(v string) *ModifyDiskDeploymentRequest {
	s.StorageClusterId = &v
	return s
}

func (s *ModifyDiskDeploymentRequest) Validate() error {
	return dara.Validate(s)
}
