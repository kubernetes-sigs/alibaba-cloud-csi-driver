// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMaintenanceAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionOnMaintenance(v string) *ModifyInstanceMaintenanceAttributesRequest
	GetActionOnMaintenance() *string
	SetInstanceId(v []*string) *ModifyInstanceMaintenanceAttributesRequest
	GetInstanceId() []*string
	SetMaintenanceWindow(v []*ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow) *ModifyInstanceMaintenanceAttributesRequest
	GetMaintenanceWindow() []*ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow
	SetNotifyOnMaintenance(v bool) *ModifyInstanceMaintenanceAttributesRequest
	GetNotifyOnMaintenance() *bool
	SetOwnerAccount(v string) *ModifyInstanceMaintenanceAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceMaintenanceAttributesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceMaintenanceAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceMaintenanceAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceMaintenanceAttributesRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceMaintenanceAttributesRequest struct {
	// The maintenance action. Valid values:
	//
	// 	- Stop: stops the instance.
	//
	// 	- AutoRecover: automatically recovers the instance.
	//
	// 	- AutoRedeploy: redeploys the instance, which may damage the data disks attached to the instance.
	//
	// example:
	//
	// AutoRecover
	ActionOnMaintenance *string `json:"ActionOnMaintenance,omitempty" xml:"ActionOnMaintenance,omitempty"`
	// The ID of instance N. Valid values of N: 1 to 100.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The maintenance windows.
	MaintenanceWindow []*ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow `json:"MaintenanceWindow,omitempty" xml:"MaintenanceWindow,omitempty" type:"Repeated"`
	// Specifies whether to send an event notification before maintenance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	NotifyOnMaintenance *bool   `json:"NotifyOnMaintenance,omitempty" xml:"NotifyOnMaintenance,omitempty"`
	OwnerAccount        *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId             *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
}

func (s ModifyInstanceMaintenanceAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMaintenanceAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintenanceAttributesRequest) GetActionOnMaintenance() *string {
	return s.ActionOnMaintenance
}

func (s *ModifyInstanceMaintenanceAttributesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *ModifyInstanceMaintenanceAttributesRequest) GetMaintenanceWindow() []*ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow {
	return s.MaintenanceWindow
}

func (s *ModifyInstanceMaintenanceAttributesRequest) GetNotifyOnMaintenance() *bool {
	return s.NotifyOnMaintenance
}

func (s *ModifyInstanceMaintenanceAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceMaintenanceAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceMaintenanceAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceMaintenanceAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceMaintenanceAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceMaintenanceAttributesRequest) SetActionOnMaintenance(v string) *ModifyInstanceMaintenanceAttributesRequest {
	s.ActionOnMaintenance = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequest) SetInstanceId(v []*string) *ModifyInstanceMaintenanceAttributesRequest {
	s.InstanceId = v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequest) SetMaintenanceWindow(v []*ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow) *ModifyInstanceMaintenanceAttributesRequest {
	s.MaintenanceWindow = v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequest) SetNotifyOnMaintenance(v bool) *ModifyInstanceMaintenanceAttributesRequest {
	s.NotifyOnMaintenance = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequest) SetOwnerAccount(v string) *ModifyInstanceMaintenanceAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequest) SetOwnerId(v int64) *ModifyInstanceMaintenanceAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequest) SetRegionId(v string) *ModifyInstanceMaintenanceAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequest) SetResourceOwnerAccount(v string) *ModifyInstanceMaintenanceAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequest) SetResourceOwnerId(v int64) *ModifyInstanceMaintenanceAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequest) Validate() error {
	if s.MaintenanceWindow != nil {
		for _, item := range s.MaintenanceWindow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow struct {
	// The end time of the maintenance window. The time must be on the hour. You must configure both StartTime and EndTime. The value of EndTime must be 1 to 23 hours later than the value of StartTime. Specify the time in the `HH:mm:ss` format. The time must be in UTC+8. Set the value of N to 1.
	//
	// example:
	//
	// 18:00:00
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The start time of the maintenance window. The time must be on the hour. You must configure both StartTime and EndTime. The value of EndTime must be 1 to 23 hours later than the value of StartTime. Specify the time in the `HH:mm:ss` format. The time must be in UTC+8. Set the value of N to 1.
	//
	// example:
	//
	// 02:00:00
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow) SetEndTime(v string) *ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow {
	s.EndTime = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow) SetStartTime(v string) *ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow {
	s.StartTime = &v
	return s
}

func (s *ModifyInstanceMaintenanceAttributesRequestMaintenanceWindow) Validate() error {
	return dara.Validate(s)
}
