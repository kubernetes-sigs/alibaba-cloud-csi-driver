// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAutoReleaseTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReleaseTime(v string) *ModifyInstanceAutoReleaseTimeRequest
	GetAutoReleaseTime() *string
	SetInstanceId(v string) *ModifyInstanceAutoReleaseTimeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceAutoReleaseTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceAutoReleaseTimeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceAutoReleaseTimeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceAutoReleaseTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceAutoReleaseTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceAutoReleaseTimeRequest struct {
	// The automatic release time of the instance. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// 	- If the value of seconds (`ss`) is not `00`, the time is automatically rounded to the nearest minute based on the value of minutes (`mm`).
	//
	// 	- The release time must be at least 30 minutes later than the current time.
	//
	// 	- The release time must be at most three years later than the current time.
	//
	// If `AutoReleaseTime` is not configured, the automatic release feature is disabled, and the instance will not be automatically released.
	//
	// example:
	//
	// 2018-01-01T01:02:03Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1env7nl3mijm2t****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyInstanceAutoReleaseTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAutoReleaseTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAutoReleaseTimeRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *ModifyInstanceAutoReleaseTimeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAutoReleaseTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceAutoReleaseTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceAutoReleaseTimeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceAutoReleaseTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceAutoReleaseTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceAutoReleaseTimeRequest) SetAutoReleaseTime(v string) *ModifyInstanceAutoReleaseTimeRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeRequest) SetInstanceId(v string) *ModifyInstanceAutoReleaseTimeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeRequest) SetOwnerAccount(v string) *ModifyInstanceAutoReleaseTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeRequest) SetOwnerId(v int64) *ModifyInstanceAutoReleaseTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeRequest) SetRegionId(v string) *ModifyInstanceAutoReleaseTimeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAutoReleaseTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeRequest) SetResourceOwnerId(v int64) *ModifyInstanceAutoReleaseTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAutoReleaseTimeRequest) Validate() error {
	return dara.Validate(s)
}
