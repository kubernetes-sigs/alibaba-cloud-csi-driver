// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostAutoReleaseTimeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoReleaseTime(v string) *ModifyDedicatedHostAutoReleaseTimeRequest
	GetAutoReleaseTime() *string
	SetDedicatedHostId(v string) *ModifyDedicatedHostAutoReleaseTimeRequest
	GetDedicatedHostId() *string
	SetOwnerAccount(v string) *ModifyDedicatedHostAutoReleaseTimeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDedicatedHostAutoReleaseTimeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDedicatedHostAutoReleaseTimeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDedicatedHostAutoReleaseTimeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDedicatedHostAutoReleaseTimeRequest
	GetResourceOwnerId() *int64
}

type ModifyDedicatedHostAutoReleaseTimeRequest struct {
	// The automatic release time of the dedicated host. Specify the time in the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC+0.
	//
	// 	- The automatic release time must be at least 30 minutes later than the current time.
	//
	// 	- The automatic release time can be up to 3 years earlier than the current time.
	//
	// 	- If the value of the seconds (ss) is not 00, it is automatically set to 00.
	//
	// 	- If `AutoReleaseTime` is not configured, the automatic release feature is disabled, and the dedicated host will not be automatically released.
	//
	// example:
	//
	// 2019-06-04T13:35:00Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// The ID of the dedicated host.
	//
	// This parameter is required.
	//
	// example:
	//
	// dh-bp165p6xk2tlw61e****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the dedicated host. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s ModifyDedicatedHostAutoReleaseTimeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostAutoReleaseTimeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) SetAutoReleaseTime(v string) *ModifyDedicatedHostAutoReleaseTimeRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) SetDedicatedHostId(v string) *ModifyDedicatedHostAutoReleaseTimeRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) SetOwnerAccount(v string) *ModifyDedicatedHostAutoReleaseTimeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) SetOwnerId(v int64) *ModifyDedicatedHostAutoReleaseTimeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) SetRegionId(v string) *ModifyDedicatedHostAutoReleaseTimeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) SetResourceOwnerAccount(v string) *ModifyDedicatedHostAutoReleaseTimeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) SetResourceOwnerId(v int64) *ModifyDedicatedHostAutoReleaseTimeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDedicatedHostAutoReleaseTimeRequest) Validate() error {
	return dara.Validate(s)
}
