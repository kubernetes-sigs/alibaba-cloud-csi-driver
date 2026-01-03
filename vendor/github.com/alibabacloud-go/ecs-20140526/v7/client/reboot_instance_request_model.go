// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRebootInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *RebootInstanceRequest
	GetDryRun() *bool
	SetForceStop(v bool) *RebootInstanceRequest
	GetForceStop() *bool
	SetInstanceId(v string) *RebootInstanceRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *RebootInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RebootInstanceRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *RebootInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RebootInstanceRequest
	GetResourceOwnerId() *int64
}

type RebootInstanceRequest struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and available ECS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- false: performs a dry run and sends the request. If the request passes the dry run, the ECS instance is restarted.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// Specifies whether to forcefully stop the ECS instance before the instance is restarted. Valid values:
	//
	// 	- true: forcefully stops the ECS instance. If you set this parameter to true, this operation is equivalent to a power-off operation. Cache data that is not written to storage devices on the instance is lost.
	//
	// 	- false: normally stops the ECS instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	ForceStop *bool `json:"ForceStop,omitempty" xml:"ForceStop,omitempty"`
	// The instance ID.
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
}

func (s RebootInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s RebootInstanceRequest) GoString() string {
	return s.String()
}

func (s *RebootInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RebootInstanceRequest) GetForceStop() *bool {
	return s.ForceStop
}

func (s *RebootInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *RebootInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RebootInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RebootInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RebootInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RebootInstanceRequest) SetDryRun(v bool) *RebootInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *RebootInstanceRequest) SetForceStop(v bool) *RebootInstanceRequest {
	s.ForceStop = &v
	return s
}

func (s *RebootInstanceRequest) SetInstanceId(v string) *RebootInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *RebootInstanceRequest) SetOwnerAccount(v string) *RebootInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RebootInstanceRequest) SetOwnerId(v int64) *RebootInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *RebootInstanceRequest) SetResourceOwnerAccount(v string) *RebootInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RebootInstanceRequest) SetResourceOwnerId(v int64) *RebootInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RebootInstanceRequest) Validate() error {
	return dara.Validate(s)
}
