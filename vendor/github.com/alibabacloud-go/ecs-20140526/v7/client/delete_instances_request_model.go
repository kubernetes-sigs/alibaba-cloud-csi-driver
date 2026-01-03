// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteInstancesRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteInstancesRequest
	GetDryRun() *bool
	SetForce(v bool) *DeleteInstancesRequest
	GetForce() *bool
	SetForceStop(v bool) *DeleteInstancesRequest
	GetForceStop() *bool
	SetInstanceId(v []*string) *DeleteInstancesRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *DeleteInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteInstancesRequest
	GetResourceOwnerId() *int64
	SetTerminateSubscription(v bool) *DeleteInstancesRequest
	GetTerminateSubscription() *bool
}

type DeleteInstancesRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. **The token can contain only ASCII characters and cannot exceed 64 characters in length.*	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized Resource Access Management (RAM) users, and missing parameter values. If the request fails the dry run, an error message is returned. If the request passes the dry run, the DRYRUN.SUCCESS error code is returned.
	//
	// 	- false: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// Default value: false.
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
	// The IDs of ECS instances. You can specify 1 to 100 ECS instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1g6zv0ce8oghu7****
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
	// Specifies whether to release the expired subscription instance.
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

func (s DeleteInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteInstancesRequest) GoString() string {
	return s.String()
}

func (s *DeleteInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteInstancesRequest) GetForce() *bool {
	return s.Force
}

func (s *DeleteInstancesRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *DeleteInstancesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DeleteInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteInstancesRequest) GetTerminateSubscription() *bool {
	return s.TerminateSubscription
}

func (s *DeleteInstancesRequest) SetClientToken(v string) *DeleteInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteInstancesRequest) SetDryRun(v bool) *DeleteInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteInstancesRequest) SetForce(v bool) *DeleteInstancesRequest {
	s.Force = &v
	return s
}

func (s *DeleteInstancesRequest) SetForceStop(v bool) *DeleteInstancesRequest {
	s.ForceStop = &v
	return s
}

func (s *DeleteInstancesRequest) SetInstanceId(v []*string) *DeleteInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *DeleteInstancesRequest) SetOwnerAccount(v string) *DeleteInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteInstancesRequest) SetOwnerId(v int64) *DeleteInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteInstancesRequest) SetRegionId(v string) *DeleteInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteInstancesRequest) SetResourceOwnerAccount(v string) *DeleteInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteInstancesRequest) SetResourceOwnerId(v int64) *DeleteInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteInstancesRequest) SetTerminateSubscription(v bool) *DeleteInstancesRequest {
	s.TerminateSubscription = &v
	return s
}

func (s *DeleteInstancesRequest) Validate() error {
	return dara.Validate(s)
}
