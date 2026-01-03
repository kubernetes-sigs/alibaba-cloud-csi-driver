// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *StopInstancesRequest
	GetBatchOptimization() *string
	SetDryRun(v bool) *StopInstancesRequest
	GetDryRun() *bool
	SetForceStop(v bool) *StopInstancesRequest
	GetForceStop() *bool
	SetInstanceId(v []*string) *StopInstancesRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *StopInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StopInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StopInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StopInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StopInstancesRequest
	GetResourceOwnerId() *int64
	SetStoppedMode(v string) *StopInstancesRequest
	GetStoppedMode() *string
}

type StopInstancesRequest struct {
	// Specifies the batch operation mode. Valid values:
	//
	// 	- AllTogether: The batch operation is successful only after all operations are successful. If any operation fails, the batch operation is considered failed, and all operations that have been performed are undone to restore the instances to the status before the batch operation.
	//
	// 	- SuccessFirst: allows each operation in a batch to be independently executed. If an operation fails, other operations can continue and confirm success. In this mode, successful operations are committed and failed operations are marked as failed, but the execution results of other operations are not affected.
	//
	// Default value: AllTogether.
	//
	// example:
	//
	// AllTogether
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	// Specifies whether to send a precheck request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, and instance status. If the check fails, the corresponding error message is returned. If the request passes the dry run, `DRYRUN.SUCCESS` is returned.
	//
	// >  If you set `BatchOptimization` to `SuccessFirst` and `DryRun` to true, only `DRYRUN.SUCCESS` is returned, regardless of whether the request passes the dry run.
	//
	// 	- false: performs a dry run and performs the actual request. If the request passes the dry run, instances are stopped.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully stop instances. Valid values:
	//
	// 	- true: forcefully stops the ECS instance.
	//
	//     **
	//
	//     **Alert*	- Force Stop: forcefully stops the instance. A force stop is equivalent to a physical shutdown and may cause data loss if instance data has not been written to disks.
	//
	// 	- false: normally stops the ECS instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The IDs of ECS instances. You can specify 1 to 100 instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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
	// The stop mode. Valid values:
	//
	// 	- StopCharging: economical mode. After the economical mode is enabled, billing for the following resources of the instance stops: computing resources (vCPUs and memory), image licenses, and public bandwidth of the static public IP address (if any) that uses the pay-by-bandwidth metering method. Billing for the following resources of the instance continues: system disk, data disks, and public bandwidth of the elastic IP address (EIP) (if any) that uses the pay-by-bandwidth metering method. For more information, see [Economical mode](https://help.aliyun.com/document_detail/63353.html).
	//
	//     **
	//
	//     **Note**
	//
	// 	- If the instance does not support the **economical*	- mode, the system stops the instance and does not report errors during the operation call. The economical mode cannot be enabled for instances of the classic network type, instances that use local disks, and instances that use persistent memory.
	//
	// 	- The instance may fail to restart due to the reclaimed computing resources or insufficient resources. Try again later or change the instance type of the instance.
	//
	// 	- If an EIP is associated with the instance before the instance is stopped, the EIP remains unchanged after the instance is restarted. If a static public IP address is associated with the instance, the static public IP address may change, but the private IP address does not change.
	//
	// 	- KeepCharging: standard mode. After the instance is stopped in standard mode, you continue to be charged for the instance. After the instance is stopped in standard mode, you continue to be charged for the instance.
	//
	// Default value: If the conditions for [enabling the economical mode for an instance in a VPC](~~63353#default~~) are met and you have enabled this mode in the ECS console, the default value is `StopCharging`. Otherwise, the default value is `KeepCharging`.
	//
	// example:
	//
	// KeepCharging
	StoppedMode *string `json:"StoppedMode,omitempty" xml:"StoppedMode,omitempty"`
}

func (s StopInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StopInstancesRequest) GoString() string {
	return s.String()
}

func (s *StopInstancesRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *StopInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StopInstancesRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *StopInstancesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *StopInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StopInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StopInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StopInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StopInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StopInstancesRequest) GetStoppedMode() *string {
	return s.StoppedMode
}

func (s *StopInstancesRequest) SetBatchOptimization(v string) *StopInstancesRequest {
	s.BatchOptimization = &v
	return s
}

func (s *StopInstancesRequest) SetDryRun(v bool) *StopInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *StopInstancesRequest) SetForceStop(v bool) *StopInstancesRequest {
	s.ForceStop = &v
	return s
}

func (s *StopInstancesRequest) SetInstanceId(v []*string) *StopInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *StopInstancesRequest) SetOwnerAccount(v string) *StopInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StopInstancesRequest) SetOwnerId(v int64) *StopInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *StopInstancesRequest) SetRegionId(v string) *StopInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *StopInstancesRequest) SetResourceOwnerAccount(v string) *StopInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StopInstancesRequest) SetResourceOwnerId(v int64) *StopInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StopInstancesRequest) SetStoppedMode(v string) *StopInstancesRequest {
	s.StoppedMode = &v
	return s
}

func (s *StopInstancesRequest) Validate() error {
	return dara.Validate(s)
}
