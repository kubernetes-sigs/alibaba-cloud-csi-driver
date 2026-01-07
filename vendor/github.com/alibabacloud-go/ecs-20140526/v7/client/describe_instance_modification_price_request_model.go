// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceModificationPriceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSystemDisk(v *DescribeInstanceModificationPriceRequestSystemDisk) *DescribeInstanceModificationPriceRequest
	GetSystemDisk() *DescribeInstanceModificationPriceRequestSystemDisk
	SetDataDisk(v []*DescribeInstanceModificationPriceRequestDataDisk) *DescribeInstanceModificationPriceRequest
	GetDataDisk() []*DescribeInstanceModificationPriceRequestDataDisk
	SetInstanceId(v string) *DescribeInstanceModificationPriceRequest
	GetInstanceId() *string
	SetInstanceType(v string) *DescribeInstanceModificationPriceRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *DescribeInstanceModificationPriceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceModificationPriceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeInstanceModificationPriceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeInstanceModificationPriceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceModificationPriceRequest
	GetResourceOwnerId() *int64
}

type DescribeInstanceModificationPriceRequest struct {
	SystemDisk *DescribeInstanceModificationPriceRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The information about data disks.
	DataDisk []*DescribeInstanceModificationPriceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The ID of the instance for which you want to query pricing information for a configuration upgrade.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1f2o4ldh8l****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new instance type. We recommend that you call the [DescribeResourcesModification](https://help.aliyun.com/document_detail/66187.html) operation to query the instance types available for configuration upgrades in a specified zone.
	//
	// > When you call the DescribeInstanceModificationPrice operation, you must specify at least one of the following parameters: `InstanceType` and `DataDisk.N.*`.
	//
	// example:
	//
	// ecs.g6e.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeInstanceModificationPriceRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceRequest) GetSystemDisk() *DescribeInstanceModificationPriceRequestSystemDisk {
	return s.SystemDisk
}

func (s *DescribeInstanceModificationPriceRequest) GetDataDisk() []*DescribeInstanceModificationPriceRequestDataDisk {
	return s.DataDisk
}

func (s *DescribeInstanceModificationPriceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceModificationPriceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeInstanceModificationPriceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceModificationPriceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceModificationPriceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceModificationPriceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceModificationPriceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceModificationPriceRequest) SetSystemDisk(v *DescribeInstanceModificationPriceRequestSystemDisk) *DescribeInstanceModificationPriceRequest {
	s.SystemDisk = v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetDataDisk(v []*DescribeInstanceModificationPriceRequestDataDisk) *DescribeInstanceModificationPriceRequest {
	s.DataDisk = v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetInstanceId(v string) *DescribeInstanceModificationPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetInstanceType(v string) *DescribeInstanceModificationPriceRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetOwnerAccount(v string) *DescribeInstanceModificationPriceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetOwnerId(v int64) *DescribeInstanceModificationPriceRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetRegionId(v string) *DescribeInstanceModificationPriceRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetResourceOwnerAccount(v string) *DescribeInstanceModificationPriceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) SetResourceOwnerId(v int64) *DescribeInstanceModificationPriceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequest) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceModificationPriceRequestSystemDisk struct {
	// The category of the system disk. You must specify this parameter only when you upgrade a non-I/O optimized instance of a retired instance type to an I/O optimized instance of an available instance type. For more information about instance types, see [Instance families](https://help.aliyun.com/document_detail/25378.html) and [Retired instance types](https://help.aliyun.com/document_detail/55263.html).
	//
	// Valid values:
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
}

func (s DescribeInstanceModificationPriceRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeInstanceModificationPriceRequestSystemDisk) SetCategory(v string) *DescribeInstanceModificationPriceRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceModificationPriceRequestDataDisk struct {
	// The category of data disk N. You can specify this parameter if you want to query the pricing information about newly attached subscription data disks. Valid values of N: 1 to 16. Valid values:
	//
	// 	- cloud_efficiency: utra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud_essd: ESSD.
	//
	// 	- cloud: basic disk.
	//
	// This parameter is empty by default.
	//
	// >  When you call the DescribeInstanceModificationPrice operation, you must specify at least one of the following parameters: `InstanceType` and `DataDisk.N.*`.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of data disk N that is an enhanced SSD (ESSD). The value of N must be the same as that in `DataDisk.N.Category` when DataDisk.N.Category is set to cloud_essd. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
	//
	// For more information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The capacity of data disk N. Valid values of N: 1 to 16. Unit: GiB. Valid values:
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud_efficiency: 20 to 32768.
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud_ssd: 20 to 32768.
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud_essd: vary based on the `DataDisk.N.PerformanceLevel` value.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL0: 1 to 32768.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL1: 20 to 32768.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL2: 461 to 32768.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL3: 1261 to 32768.
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud: 5 to 2000.
	//
	// The default value is the minimum capacity allowed for the specified data disk category.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s DescribeInstanceModificationPriceRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceModificationPriceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) SetCategory(v string) *DescribeInstanceModificationPriceRequestDataDisk {
	s.Category = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) SetPerformanceLevel(v string) *DescribeInstanceModificationPriceRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) SetSize(v int32) *DescribeInstanceModificationPriceRequestDataDisk {
	s.Size = &v
	return s
}

func (s *DescribeInstanceModificationPriceRequestDataDisk) Validate() error {
	return dara.Validate(s)
}
