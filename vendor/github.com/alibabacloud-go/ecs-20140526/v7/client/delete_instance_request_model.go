// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DeleteInstanceRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteInstanceRequest
	GetForce() *bool
	SetForceStop(v bool) *DeleteInstanceRequest
	GetForceStop() *bool
	SetInstanceId(v string) *DeleteInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DeleteInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DeleteInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteInstanceRequest
	GetResourceOwnerId() *int64
	SetTerminateSubscription(v bool) *DeleteInstanceRequest
	GetTerminateSubscription() *bool
}

type DeleteInstanceRequest struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and unavailable ECS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- false (default): performs a dry run and performs the actual request. If the request passes the dry run, the instance is released.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully release the ECS instance in the **Running*	- (`Running`) state. Valid values:
	//
	// 	- true: forcefully releases the ECS instance in the **Running*	- (`Running`) state.
	//
	// 	- false: normally releases the ECS instance. This value is valid only if the instance is in the **Stopped*	- (`Stopped`) state.
	//
	// Default value: false.
	//
	// **
	//
	// **Warning*	- When Force is set to true, this operation is equivalent to a power-off operation. Temporary data in the memory and storage of the instance is erased and cannot be restored.
	//
	// example:
	//
	// false
	Force *bool `json:"Force,omitempty" xml:"Force,omitempty"`
	// Specifies whether to forcefully stop the ECS instance in the **Running*	- (`Running`) state before the instance is released. This parameter takes effect only when `Force` is set to true. Valid values:
	//
	// 	- true: forcefully stops and releases the ECS instance. In this case, this operation is equivalent to a power-off operation. The instance directly enters the resource release process.
	//
	//     **
	//
	//     **Warning*	- A forceful stop and release is equivalent to a power-off operation. Temporary data in the memory and storage of the instance is erased and cannot be restored.
	//
	// 	- false: stops the ECS instance in the normal stop process and then releases the instance. In this case, the release process takes several minutes to complete. You can configure business drainage actions to reduce the noise of the business system on operating system shutdown.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1g6zv0ce8oghu7****
	InstanceId           *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to release the expired subscription instance. Valid values:
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
	TerminateSubscription *bool `json:"TerminateSubscription,omitempty" xml:"TerminateSubscription,omitempty"`
}

func (s DeleteInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstanceRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteInstanceRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *DeleteInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteInstanceRequest) GetTerminateSubscription() *bool {
	return s.TerminateSubscription
}

func (s *DeleteInstanceRequest) SetDryRun(v bool) *DeleteInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteInstanceRequest) SetForce(v bool) *DeleteInstanceRequest {
	s.Force = &v
	return s
}

func (s *DeleteInstanceRequest) SetForceStop(v bool) *DeleteInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *DeleteInstanceRequest) SetInstanceId(v string) *DeleteInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteInstanceRequest) SetOwnerAccount(v string) *DeleteInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteInstanceRequest) SetOwnerId(v int64) *DeleteInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerAccount(v string) *DeleteInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInstanceRequest) SetResourceOwnerId(v int64) *DeleteInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteInstanceRequest) SetTerminateSubscription(v bool) *DeleteInstanceRequest {
	s.TerminateSubscription = &v
	return s
}

func (s *DeleteInstanceRequest) Validate() error {
	return dara.Validate(s)
}
