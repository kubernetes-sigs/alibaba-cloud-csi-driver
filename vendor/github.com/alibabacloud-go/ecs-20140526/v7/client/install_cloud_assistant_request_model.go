// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInstallCloudAssistantRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v []*string) *InstallCloudAssistantRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *InstallCloudAssistantRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *InstallCloudAssistantRequest
	GetOwnerId() *int64
	SetRegionId(v string) *InstallCloudAssistantRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *InstallCloudAssistantRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *InstallCloudAssistantRequest
	GetResourceOwnerId() *int64
}

type InstallCloudAssistantRequest struct {
	// The ID of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1iudwa5b1tqa****
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of instances. You can specify up to 50 instance IDs in a single request.
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

func (s InstallCloudAssistantRequest) String() string {
	return dara.Prettify(s)
}

func (s InstallCloudAssistantRequest) GoString() string {
	return s.String()
}

func (s *InstallCloudAssistantRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *InstallCloudAssistantRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *InstallCloudAssistantRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *InstallCloudAssistantRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *InstallCloudAssistantRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *InstallCloudAssistantRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *InstallCloudAssistantRequest) SetInstanceId(v []*string) *InstallCloudAssistantRequest {
	s.InstanceId = v
	return s
}

func (s *InstallCloudAssistantRequest) SetOwnerAccount(v string) *InstallCloudAssistantRequest {
	s.OwnerAccount = &v
	return s
}

func (s *InstallCloudAssistantRequest) SetOwnerId(v int64) *InstallCloudAssistantRequest {
	s.OwnerId = &v
	return s
}

func (s *InstallCloudAssistantRequest) SetRegionId(v string) *InstallCloudAssistantRequest {
	s.RegionId = &v
	return s
}

func (s *InstallCloudAssistantRequest) SetResourceOwnerAccount(v string) *InstallCloudAssistantRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *InstallCloudAssistantRequest) SetResourceOwnerId(v int64) *InstallCloudAssistantRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *InstallCloudAssistantRequest) Validate() error {
	return dara.Validate(s)
}
