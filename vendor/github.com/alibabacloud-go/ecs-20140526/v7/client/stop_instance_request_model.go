// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfirmStop(v bool) *StopInstanceRequest
	GetConfirmStop() *bool
	SetDryRun(v bool) *StopInstanceRequest
	GetDryRun() *bool
	SetForceStop(v bool) *StopInstanceRequest
	GetForceStop() *bool
	SetHibernate(v bool) *StopInstanceRequest
	GetHibernate() *bool
	SetInstanceId(v string) *StopInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *StopInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StopInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *StopInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopInstanceRequest
	GetResourceOwnerId() *int64
	SetStoppedMode(v string) *StopInstanceRequest
	GetStoppedMode() *string
}

type StopInstanceRequest struct {
	// This parameter will be removed in the future and is retained only to ensure compatibility. We recommend that you ignore this parameter.
	//
	// example:
	//
	// true
	ConfirmStop *bool `json:"ConfirmStop,omitempty" xml:"ConfirmStop,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and available ECS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- false: performs a dry run and performs the actual request.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully stop the ECS instance. Valid values:
	//
	// 	- true: forcefully stops the ECS instance. If you set ForceStop to true, this operation is equivalent to a power-off operation. Cache data that is not written to storage devices on the instance is lost.
	//
	// 	- false: normally stops the ECS instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// hide
	Hibernate *bool `json:"Hibernate,omitempty" xml:"Hibernate,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The stop mode of the pay-as-you-go instance. Valid values:
	//
	// 	- StopCharging: economical mode. After the economical mode is enabled, billing for the following resources of the instance stops: computing resources (vCPUs, memory, and GPUs), image licenses, and public bandwidth of the static public IP address (if any) that uses the pay-by-bandwidth metering method. Billing for the following resources of the instance continues: system disk, data disks, and public bandwidth of the elastic IP address (EIP) (if any) that uses the pay-by-bandwidth metering method. For more information, see [Economical mode](https://help.aliyun.com/document_detail/63353.html).
	//
	//     **
	//
	//     **Note**
	//
	// 	- If the instance does not support the **economical*	- mode, the system stops the instance and does not report errors during the operation call. The economical mode cannot be enabled for instances of the classic network type, instances that use local disks, and instances that use persistent memory.
	//
	// 	- The instance may fail to restart due to the reclaimed computing resources or insufficient resources. Try again later or change the instance type of the instance.
	//
	// 	- If an EIP is associated with the instance before the instance is stopped, the EIP remains unchanged after the instance is restarted. If a static public IP address is associated with the instance before the instance is stopped, the static public IP address may change, but the private IP address does not change.
	//
	// 	- KeepCharging: standard mode. After the instance is stopped in standard mode, you continue to be charged for the instance.
	//
	// Default value: If the conditions for [enabling the economical mode for an instance in a VPC](~~63353#default~~) are met and you have enabled this mode in the ECS console, the default value is `StopCharging`. Otherwise, the default value is `KeepCharging`.
	//
	// example:
	//
	// KeepCharging
	StoppedMode *string `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
}

func (s StopInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstanceRequest) GoString() string {
	return s.String()
}

func (s *StopInstanceRequest) GetConfirmStop() *bool {
	return s.ConfirmStop
}

func (s *StopInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StopInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *StopInstanceRequest) GetHibernate() *bool {
	return s.Hibernate
}

func (s *StopInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StopInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StopInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopInstanceRequest) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *StopInstanceRequest) SetConfirmStop(v bool) *StopInstanceRequest {
	s.ConfirmStop = &v
	return s
}

func (s *StopInstanceRequest) SetDryRun(v bool) *StopInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *StopInstanceRequest) SetForceStop(v bool) *StopInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *StopInstanceRequest) SetHibernate(v bool) *StopInstanceRequest {
	s.Hibernate = &v
	return s
}

func (s *StopInstanceRequest) SetInstanceId(v string) *StopInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *StopInstanceRequest) SetOwnerAccount(v string) *StopInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StopInstanceRequest) SetOwnerId(v int64) *StopInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *StopInstanceRequest) SetResourceOwnerAccount(v string) *StopInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopInstanceRequest) SetResourceOwnerId(v int64) *StopInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopInstanceRequest) SetStoppedMode(v string) *StopInstanceRequest {
	s.StoppedMode = &v
	return s
}

func (s *StopInstanceRequest) Validate() error {
	return dara.Validate(s)
}
