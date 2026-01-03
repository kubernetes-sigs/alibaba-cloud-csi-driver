// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceClockOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyInstanceClockOptionsRequest
	GetClientToken() *string
	SetDryRun(v bool) *ModifyInstanceClockOptionsRequest
	GetDryRun() *bool
	SetInstanceId(v string) *ModifyInstanceClockOptionsRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceClockOptionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceClockOptionsRequest
	GetOwnerId() *int64
	SetPtpStatus(v string) *ModifyInstanceClockOptionsRequest
	GetPtpStatus() *string
	SetRegionId(v string) *ModifyInstanceClockOptionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceClockOptionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceClockOptionsRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceClockOptionsRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform only a dry run. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including invalid AccessKey pairs, unauthorized Resource Access Management (RAM) users, and missing parameter values. Otherwise, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- false: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// PTP status value. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// Default value: disabled.
	//
	// example:
	//
	// enabled
	PtpStatus *string `json:"PtpStatus,omitempty" xml:"PtpStatus,omitempty"`
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

func (s ModifyInstanceClockOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceClockOptionsRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceClockOptionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceClockOptionsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyInstanceClockOptionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceClockOptionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceClockOptionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceClockOptionsRequest) GetPtpStatus() *string {
	return s.PtpStatus
}

func (s *ModifyInstanceClockOptionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceClockOptionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceClockOptionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceClockOptionsRequest) SetClientToken(v string) *ModifyInstanceClockOptionsRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceClockOptionsRequest) SetDryRun(v bool) *ModifyInstanceClockOptionsRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyInstanceClockOptionsRequest) SetInstanceId(v string) *ModifyInstanceClockOptionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceClockOptionsRequest) SetOwnerAccount(v string) *ModifyInstanceClockOptionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceClockOptionsRequest) SetOwnerId(v int64) *ModifyInstanceClockOptionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceClockOptionsRequest) SetPtpStatus(v string) *ModifyInstanceClockOptionsRequest {
	s.PtpStatus = &v
	return s
}

func (s *ModifyInstanceClockOptionsRequest) SetRegionId(v string) *ModifyInstanceClockOptionsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceClockOptionsRequest) SetResourceOwnerAccount(v string) *ModifyInstanceClockOptionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceClockOptionsRequest) SetResourceOwnerId(v int64) *ModifyInstanceClockOptionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceClockOptionsRequest) Validate() error {
	return dara.Validate(s)
}
