// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoProvisioningGroupInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeAutoProvisioningGroupInstancesResponseBodyInstances) *DescribeAutoProvisioningGroupInstancesResponseBody
	GetInstances() *DescribeAutoProvisioningGroupInstancesResponseBodyInstances
	SetPageNumber(v int32) *DescribeAutoProvisioningGroupInstancesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoProvisioningGroupInstancesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAutoProvisioningGroupInstancesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAutoProvisioningGroupInstancesResponseBody
	GetTotalCount() *int32
}

type DescribeAutoProvisioningGroupInstancesResponseBody struct {
	// The information about the instances in the auto provisioning group.
	Instances *DescribeAutoProvisioningGroupInstancesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
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
	// B48A12CD-1295-4A38-A8F0-0E92C937****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of queried instances in the auto provisioning group.
	//
	// example:
	//
	// 2
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAutoProvisioningGroupInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) GetInstances() *DescribeAutoProvisioningGroupInstancesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) SetInstances(v *DescribeAutoProvisioningGroupInstancesResponseBodyInstances) *DescribeAutoProvisioningGroupInstancesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) SetPageNumber(v int32) *DescribeAutoProvisioningGroupInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) SetPageSize(v int32) *DescribeAutoProvisioningGroupInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) SetRequestId(v string) *DescribeAutoProvisioningGroupInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) SetTotalCount(v int32) *DescribeAutoProvisioningGroupInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupInstancesResponseBodyInstances struct {
	Instance []*DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeAutoProvisioningGroupInstancesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupInstancesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstances) GetInstance() []*DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstances) SetInstance(v []*DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) *DescribeAutoProvisioningGroupInstancesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance struct {
	// The number of vCPU cores of the instance.
	//
	// example:
	//
	// 2
	CPU *int32 `json:"CPU,omitempty" xml:"CPU,omitempty"`
	// The time when the instance was created.
	//
	// example:
	//
	// 2017-12-10T04:04Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ECS instance type.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// Indicates whether the instance is an I/O optimized instance.
	//
	// example:
	//
	// true
	IoOptimized *bool `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// Indicates whether the instance is a spot instance.
	//
	// example:
	//
	// true
	IsSpot *bool `json:"IsSpot,omitempty" xml:"IsSpot,omitempty"`
	// The memory capacity of the instance. Unit: MiB.
	//
	// example:
	//
	// 1024
	Memory *int32 `json:"Memory,omitempty" xml:"Memory,omitempty"`
	// The network type of the instance. Valid values:
	//
	// 	- vpc: Virtual Private Cloud (VPC)
	//
	// 	- classic: classic network
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The operating system type of the instance. Valid values:
	//
	// 	- windows
	//
	// 	- linux
	//
	// example:
	//
	// linux
	OsType *string `json:"OsType,omitempty" xml:"OsType,omitempty"`
	// The region ID of the container group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The status of the instance.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the zone to which the instance belongs.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetCPU() *int32 {
	return s.CPU
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetIoOptimized() *bool {
	return s.IoOptimized
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetIsSpot() *bool {
	return s.IsSpot
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetMemory() *int32 {
	return s.Memory
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetOsType() *string {
	return s.OsType
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetStatus() *string {
	return s.Status
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetCPU(v int32) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.CPU = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetCreationTime(v string) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.CreationTime = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetInstanceType(v string) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.InstanceType = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetIoOptimized(v bool) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.IoOptimized = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetIsSpot(v bool) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.IsSpot = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetMemory(v int32) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.Memory = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetNetworkType(v string) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.NetworkType = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetOsType(v string) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.OsType = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetRegionId(v string) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetStatus(v string) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.Status = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) SetZoneId(v string) *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance {
	s.ZoneId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupInstancesResponseBodyInstancesInstance) Validate() error {
	return dara.Validate(s)
}
