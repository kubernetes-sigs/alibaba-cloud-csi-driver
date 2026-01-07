// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSimulatedSystemEventsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEventType(v string) *CreateSimulatedSystemEventsRequest
	GetEventType() *string
	SetInstanceId(v []*string) *CreateSimulatedSystemEventsRequest
	GetInstanceId() []*string
	SetNotBefore(v string) *CreateSimulatedSystemEventsRequest
	GetNotBefore() *string
	SetOwnerAccount(v string) *CreateSimulatedSystemEventsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateSimulatedSystemEventsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateSimulatedSystemEventsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CreateSimulatedSystemEventsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateSimulatedSystemEventsRequest
	GetResourceOwnerId() *int64
}

type CreateSimulatedSystemEventsRequest struct {
	// The type of the system event. Valid values:
	//
	// 	- SystemMaintenance.Reboot: The instance is restarted due to system maintenance.
	//
	// 	- SystemFailure.Reboot: The instance is restarted due to a system error.
	//
	// 	- InstanceFailure.Reboot: The instance is restarted due to an instance error.
	//
	// 	- SystemMaintenance.Stop: The instance is stopped due to system maintenance.
	//
	// 	- SystemMaintenance.Redeploy: The instance is redeployed due to system maintenance.
	//
	// 	- SystemFailure.Redeploy: The instance is redeployed due to a system error.
	//
	// 	- SystemFailure.Stop: The instance is stopped due to a system error.
	//
	// This parameter is required.
	//
	// example:
	//
	// SystemMaintenance.Reboot
	EventType *string `json:"EventType,omitempty" xml:"EventType,omitempty"`
	// The IDs of the instances. You can specify up to 100 instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1gtjxuuvwj17zr****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The scheduled start time of the event. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// > For events that occur due to system errors or instance errors, the simulated events of such events enter the `Executing` state when the simulated events are created. The value of `NotBefore` is the time when the simulated events enter the `Executed` state.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-12-01T06:32:31Z
	NotBefore    *string `json:"NotBefore,omitempty" xml:"NotBefore,omitempty"`
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

func (s CreateSimulatedSystemEventsRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSimulatedSystemEventsRequest) GoString() string {
	return s.String()
}

func (s *CreateSimulatedSystemEventsRequest) GetEventType() *string {
	return s.EventType
}

func (s *CreateSimulatedSystemEventsRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *CreateSimulatedSystemEventsRequest) GetNotBefore() *string {
	return s.NotBefore
}

func (s *CreateSimulatedSystemEventsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateSimulatedSystemEventsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateSimulatedSystemEventsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSimulatedSystemEventsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateSimulatedSystemEventsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateSimulatedSystemEventsRequest) SetEventType(v string) *CreateSimulatedSystemEventsRequest {
	s.EventType = &v
	return s
}

func (s *CreateSimulatedSystemEventsRequest) SetInstanceId(v []*string) *CreateSimulatedSystemEventsRequest {
	s.InstanceId = v
	return s
}

func (s *CreateSimulatedSystemEventsRequest) SetNotBefore(v string) *CreateSimulatedSystemEventsRequest {
	s.NotBefore = &v
	return s
}

func (s *CreateSimulatedSystemEventsRequest) SetOwnerAccount(v string) *CreateSimulatedSystemEventsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateSimulatedSystemEventsRequest) SetOwnerId(v int64) *CreateSimulatedSystemEventsRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateSimulatedSystemEventsRequest) SetRegionId(v string) *CreateSimulatedSystemEventsRequest {
	s.RegionId = &v
	return s
}

func (s *CreateSimulatedSystemEventsRequest) SetResourceOwnerAccount(v string) *CreateSimulatedSystemEventsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateSimulatedSystemEventsRequest) SetResourceOwnerId(v int64) *CreateSimulatedSystemEventsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateSimulatedSystemEventsRequest) Validate() error {
	return dara.Validate(s)
}
