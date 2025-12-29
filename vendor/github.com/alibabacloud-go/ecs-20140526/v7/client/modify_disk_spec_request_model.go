// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDestinationZoneId(v string) *ModifyDiskSpecRequest
	GetDestinationZoneId() *string
	SetDiskCategory(v string) *ModifyDiskSpecRequest
	GetDiskCategory() *string
	SetDiskId(v string) *ModifyDiskSpecRequest
	GetDiskId() *string
	SetDryRun(v bool) *ModifyDiskSpecRequest
	GetDryRun() *bool
	SetOwnerAccount(v string) *ModifyDiskSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDiskSpecRequest
	GetOwnerId() *int64
	SetPerformanceControlOptions(v *ModifyDiskSpecRequestPerformanceControlOptions) *ModifyDiskSpecRequest
	GetPerformanceControlOptions() *ModifyDiskSpecRequestPerformanceControlOptions
	SetPerformanceLevel(v string) *ModifyDiskSpecRequest
	GetPerformanceLevel() *string
	SetProvisionedIops(v int64) *ModifyDiskSpecRequest
	GetProvisionedIops() *int64
	SetResourceOwnerAccount(v string) *ModifyDiskSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDiskSpecRequest
	GetResourceOwnerId() *int64
}

type ModifyDiskSpecRequest struct {
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// cn-hangzhou-g
	DestinationZoneId *string `json:"DestinationZoneId,omitempty" xml:"DestinationZoneId,omitempty"`
	// The new disk category of the cloud disk. Valid values:
	//
	// 	- cloud_essd: ESSD
	//
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_efficiency: utra disk
	//
	// This parameter is empty by default, which indicates that the disk category is not changed.
	//
	// >
	//
	// 	- The preceding values are listed in descending order of disk performance. Subscription disks cannot be downgraded.
	//
	// example:
	//
	// cloud_essd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The disk ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp131n0q38u3a4zi****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and insufficient ECS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
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
	// The disk performance specifications.
	PerformanceControlOptions *ModifyDiskSpecRequestPerformanceControlOptions `json:"PerformanceControlOptions,omitempty" xml:"PerformanceControlOptions,omitempty" type:"Struct"`
	// The new performance level of the ESSD. Valid values:
	//
	// 	- PL0: An ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: An ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: An ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: An ESSD delivers up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
	//
	// example:
	//
	// PL2
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk.
	//
	// Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}.
	//
	// Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// >  This parameter is available only if you set `DiskCategory` to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html) and [Modify the performance configurations of an ESSD AutoPL disk](https://help.aliyun.com/document_detail/413275.html).
	//
	// example:
	//
	// 50000
	ProvisionedIops      *int64  `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDiskSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskSpecRequest) GetDestinationZoneId() *string {
	return s.DestinationZoneId
}

func (s *ModifyDiskSpecRequest) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *ModifyDiskSpecRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *ModifyDiskSpecRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDiskSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDiskSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDiskSpecRequest) GetPerformanceControlOptions() *ModifyDiskSpecRequestPerformanceControlOptions {
	return s.PerformanceControlOptions
}

func (s *ModifyDiskSpecRequest) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *ModifyDiskSpecRequest) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *ModifyDiskSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDiskSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDiskSpecRequest) SetDestinationZoneId(v string) *ModifyDiskSpecRequest {
	s.DestinationZoneId = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetDiskCategory(v string) *ModifyDiskSpecRequest {
	s.DiskCategory = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetDiskId(v string) *ModifyDiskSpecRequest {
	s.DiskId = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetDryRun(v bool) *ModifyDiskSpecRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetOwnerAccount(v string) *ModifyDiskSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetOwnerId(v int64) *ModifyDiskSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetPerformanceControlOptions(v *ModifyDiskSpecRequestPerformanceControlOptions) *ModifyDiskSpecRequest {
	s.PerformanceControlOptions = v
	return s
}

func (s *ModifyDiskSpecRequest) SetPerformanceLevel(v string) *ModifyDiskSpecRequest {
	s.PerformanceLevel = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetProvisionedIops(v int64) *ModifyDiskSpecRequest {
	s.ProvisionedIops = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetResourceOwnerAccount(v string) *ModifyDiskSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDiskSpecRequest) SetResourceOwnerId(v int64) *ModifyDiskSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDiskSpecRequest) Validate() error {
	if s.PerformanceControlOptions != nil {
		if err := s.PerformanceControlOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDiskSpecRequestPerformanceControlOptions struct {
	// The new IOPS rate of the cloud disk. You can modify the IOPS rate of only cloud disks in dedicated block storage clusters.
	//
	// Valid values: 900 to maximum IOPS per disk (with an increment of 100).
	//
	// For more information, see [Block storage performance](https://help.aliyun.com/document_detail/25382.html).
	//
	// example:
	//
	// 2000
	IOPS *int32 `json:"IOPS,omitempty" xml:"IOPS,omitempty"`
	// Specifies whether to reset the IOPS rate and throughput of the cloud disk. This parameter takes effect only when the cloud disk belongs to a dedicated block storage cluster.
	//
	// After you specify this parameter, PerformanceControlOptions.IOPS and PerformanceControlOptions.Throughput do not take effect.
	//
	// Set the value to All, which indicates that the IOPS rate and throughput of the cloud disk are reset to the initial values.
	//
	// example:
	//
	// All
	Recover *string `json:"Recover,omitempty" xml:"Recover,omitempty"`
	// The new throughput of the cloud disk. You can change the throughput of only cloud disks in dedicated block storage clusters. Unit: MB/s.
	//
	// Valid values: 60 to maximum throughput per disk.
	//
	// For more information, see [Block storage performance](https://help.aliyun.com/document_detail/25382.html).
	//
	// example:
	//
	// 200
	Throughput *int32 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
}

func (s ModifyDiskSpecRequestPerformanceControlOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskSpecRequestPerformanceControlOptions) GoString() string {
	return s.String()
}

func (s *ModifyDiskSpecRequestPerformanceControlOptions) GetIOPS() *int32 {
	return s.IOPS
}

func (s *ModifyDiskSpecRequestPerformanceControlOptions) GetRecover() *string {
	return s.Recover
}

func (s *ModifyDiskSpecRequestPerformanceControlOptions) GetThroughput() *int32 {
	return s.Throughput
}

func (s *ModifyDiskSpecRequestPerformanceControlOptions) SetIOPS(v int32) *ModifyDiskSpecRequestPerformanceControlOptions {
	s.IOPS = &v
	return s
}

func (s *ModifyDiskSpecRequestPerformanceControlOptions) SetRecover(v string) *ModifyDiskSpecRequestPerformanceControlOptions {
	s.Recover = &v
	return s
}

func (s *ModifyDiskSpecRequestPerformanceControlOptions) SetThroughput(v int32) *ModifyDiskSpecRequestPerformanceControlOptions {
	s.Throughput = &v
	return s
}

func (s *ModifyDiskSpecRequestPerformanceControlOptions) Validate() error {
	return dara.Validate(s)
}
