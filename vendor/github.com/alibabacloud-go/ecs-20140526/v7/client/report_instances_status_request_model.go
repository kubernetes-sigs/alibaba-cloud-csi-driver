// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReportInstancesStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ReportInstancesStatusRequest
	GetDescription() *string
	SetDevice(v []*string) *ReportInstancesStatusRequest
	GetDevice() []*string
	SetDiskId(v []*string) *ReportInstancesStatusRequest
	GetDiskId() []*string
	SetEndTime(v string) *ReportInstancesStatusRequest
	GetEndTime() *string
	SetInstanceId(v []*string) *ReportInstancesStatusRequest
	GetInstanceId() []*string
	SetIssueCategory(v string) *ReportInstancesStatusRequest
	GetIssueCategory() *string
	SetOwnerAccount(v string) *ReportInstancesStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ReportInstancesStatusRequest
	GetOwnerId() *int64
	SetReason(v string) *ReportInstancesStatusRequest
	GetReason() *string
	SetRegionId(v string) *ReportInstancesStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ReportInstancesStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ReportInstancesStatusRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *ReportInstancesStatusRequest
	GetStartTime() *string
}

type ReportInstancesStatusRequest struct {
	// The description of the exception.
	//
	// This parameter is required.
	//
	// example:
	//
	// The local disk is unavailable, the mount point is inaccessible, or files cannot be loaded.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The device names of disks on an instance that have the exception. You can specify to 100 device names in a single request.
	//
	// If you are using an ECS bare metal instance, enter the slot numbers of disks on the instance.
	//
	// > For ECS bare metal instances, this parameter is required when the value of the `Reason` parameter is `abnormal-local-disk` or `abnormal-cloud-disk` or when the value of the `IssueCategory` parameter is `hardware-disk-error`.
	//
	// example:
	//
	// /dev/xvdb
	Device []*string `json:"Device,omitempty" xml:"Device,omitempty" type:"Repeated"`
	// The IDs of disks on an instance that have the exception. You can specify up to 100 disk IDs in a single request. If you are using an ECS bare metal instance, enter the serial numbers of disks on the instance.
	//
	// > This parameter is required when the value of the `Reason` parameter is `abnormal-local-disk` or `abnormal-cloud-disk` or when the value of the `IssueCategory` parameter is `hardware-disk-error`.
	//
	// example:
	//
	// d-bp1aeljlfad7x6u1****
	DiskId []*string `json:"DiskId,omitempty" xml:"DiskId,omitempty" type:"Repeated"`
	// The end time of the instance exception. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-11-31T06:32:31Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The IDs of instances. You can specify up to 100 instance IDs in a single request.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp165p6xk2tmdhj0****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The category of the exception. This parameter is applicable only to ECS bare metal instances. Valid values:
	//
	// 	- hardware-cpu-error: CPU failure
	//
	// 	- hardware-motherboard-error: motherboard failure
	//
	// 	- hardware-mem-error: memory failure
	//
	// 	- hardware-power-error: power failure
	//
	// 	- hardware-disk-error: disk failure
	//
	// 	- hardware-networkcard-error: network interface controller (NIC) failure
	//
	// 	- hardware-raidcard-error: SAS/RAID card failure
	//
	// 	- hardware-fan-error: fan failure
	//
	// 	- others: other failures
	//
	// example:
	//
	// hardware-cpu-error
	IssueCategory *string `json:"IssueCategory,omitempty" xml:"IssueCategory,omitempty"`
	OwnerAccount  *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The impact of the exception on the instance. Valid values:
	//
	// 	- instance-hang: The instance is unavailable or cannot be connected.
	//
	// 	- instance-stuck-in-status: The instance is stuck in a state such as Starting or Stopping.
	//
	// 	- abnormal-network: The instance has a network exception.
	//
	// 	- abnormal-local-disk: A local disk attached to the instance has an exception.
	//
	// 	- abnormal-cloud-disk: A disk or a Shared Block Storage device attached to the instance has an exception.
	//
	// 	- others: other exception types. If the impact is not of the preceding types, you can set `Reason` to others and specify the `Description` parameter.
	//
	// example:
	//
	// abnormal-local-disk
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start time of the instance exception. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// example:
	//
	// 2017-11-30T06:32:31Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ReportInstancesStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportInstancesStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportInstancesStatusRequest) GetDescription() *string {
	return s.Description
}

func (s *ReportInstancesStatusRequest) GetDevice() []*string {
	return s.Device
}

func (s *ReportInstancesStatusRequest) GetDiskId() []*string {
	return s.DiskId
}

func (s *ReportInstancesStatusRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ReportInstancesStatusRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *ReportInstancesStatusRequest) GetIssueCategory() *string {
	return s.IssueCategory
}

func (s *ReportInstancesStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ReportInstancesStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ReportInstancesStatusRequest) GetReason() *string {
	return s.Reason
}

func (s *ReportInstancesStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ReportInstancesStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ReportInstancesStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ReportInstancesStatusRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ReportInstancesStatusRequest) SetDescription(v string) *ReportInstancesStatusRequest {
	s.Description = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetDevice(v []*string) *ReportInstancesStatusRequest {
	s.Device = v
	return s
}

func (s *ReportInstancesStatusRequest) SetDiskId(v []*string) *ReportInstancesStatusRequest {
	s.DiskId = v
	return s
}

func (s *ReportInstancesStatusRequest) SetEndTime(v string) *ReportInstancesStatusRequest {
	s.EndTime = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetInstanceId(v []*string) *ReportInstancesStatusRequest {
	s.InstanceId = v
	return s
}

func (s *ReportInstancesStatusRequest) SetIssueCategory(v string) *ReportInstancesStatusRequest {
	s.IssueCategory = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetOwnerAccount(v string) *ReportInstancesStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetOwnerId(v int64) *ReportInstancesStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetReason(v string) *ReportInstancesStatusRequest {
	s.Reason = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetRegionId(v string) *ReportInstancesStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetResourceOwnerAccount(v string) *ReportInstancesStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetResourceOwnerId(v int64) *ReportInstancesStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ReportInstancesStatusRequest) SetStartTime(v string) *ReportInstancesStatusRequest {
	s.StartTime = &v
	return s
}

func (s *ReportInstancesStatusRequest) Validate() error {
	return dara.Validate(s)
}
