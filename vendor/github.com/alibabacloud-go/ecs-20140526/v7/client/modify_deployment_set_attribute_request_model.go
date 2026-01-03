// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeploymentSetAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetId(v string) *ModifyDeploymentSetAttributeRequest
	GetDeploymentSetId() *string
	SetDeploymentSetName(v string) *ModifyDeploymentSetAttributeRequest
	GetDeploymentSetName() *string
	SetDescription(v string) *ModifyDeploymentSetAttributeRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyDeploymentSetAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDeploymentSetAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDeploymentSetAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDeploymentSetAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDeploymentSetAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyDeploymentSetAttributeRequest struct {
	// The ID of the deployment set.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds-bp1frxuzdg87zh4p****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The new name of the deployment set. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// DeploymentSetTestName
	DeploymentSetName *string `json:"DeploymentSetName,omitempty" xml:"DeploymentSetName,omitempty"`
	// The new description of the deployment set. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// TestDescription
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the deployment set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s ModifyDeploymentSetAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeploymentSetAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDeploymentSetAttributeRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *ModifyDeploymentSetAttributeRequest) GetDeploymentSetName() *string {
	return s.DeploymentSetName
}

func (s *ModifyDeploymentSetAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDeploymentSetAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDeploymentSetAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDeploymentSetAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDeploymentSetAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDeploymentSetAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDeploymentSetAttributeRequest) SetDeploymentSetId(v string) *ModifyDeploymentSetAttributeRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *ModifyDeploymentSetAttributeRequest) SetDeploymentSetName(v string) *ModifyDeploymentSetAttributeRequest {
	s.DeploymentSetName = &v
	return s
}

func (s *ModifyDeploymentSetAttributeRequest) SetDescription(v string) *ModifyDeploymentSetAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyDeploymentSetAttributeRequest) SetOwnerAccount(v string) *ModifyDeploymentSetAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDeploymentSetAttributeRequest) SetOwnerId(v int64) *ModifyDeploymentSetAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDeploymentSetAttributeRequest) SetRegionId(v string) *ModifyDeploymentSetAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDeploymentSetAttributeRequest) SetResourceOwnerAccount(v string) *ModifyDeploymentSetAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDeploymentSetAttributeRequest) SetResourceOwnerId(v int64) *ModifyDeploymentSetAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDeploymentSetAttributeRequest) Validate() error {
	return dara.Validate(s)
}
