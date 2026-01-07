// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceScreenshotRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceScreenshotRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetInstanceScreenshotRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetInstanceScreenshotRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetInstanceScreenshotRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *GetInstanceScreenshotRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetInstanceScreenshotRequest
	GetResourceOwnerId() *int64
	SetWakeUp(v bool) *GetInstanceScreenshotRequest
	GetWakeUp() *bool
}

type GetInstanceScreenshotRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1gbz20g229bvu5****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-shenzhen
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to wake up the hibernated instance.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	WakeUp *bool `json:"WakeUp,omitempty" xml:"WakeUp,omitempty"`
}

func (s GetInstanceScreenshotRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceScreenshotRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceScreenshotRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceScreenshotRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetInstanceScreenshotRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetInstanceScreenshotRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceScreenshotRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetInstanceScreenshotRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetInstanceScreenshotRequest) GetWakeUp() *bool {
	return s.WakeUp
}

func (s *GetInstanceScreenshotRequest) SetInstanceId(v string) *GetInstanceScreenshotRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceScreenshotRequest) SetOwnerAccount(v string) *GetInstanceScreenshotRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetInstanceScreenshotRequest) SetOwnerId(v int64) *GetInstanceScreenshotRequest {
	s.OwnerId = &v
	return s
}

func (s *GetInstanceScreenshotRequest) SetRegionId(v string) *GetInstanceScreenshotRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceScreenshotRequest) SetResourceOwnerAccount(v string) *GetInstanceScreenshotRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetInstanceScreenshotRequest) SetResourceOwnerId(v int64) *GetInstanceScreenshotRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetInstanceScreenshotRequest) SetWakeUp(v bool) *GetInstanceScreenshotRequest {
	s.WakeUp = &v
	return s
}

func (s *GetInstanceScreenshotRequest) Validate() error {
	return dara.Validate(s)
}
