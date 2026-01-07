// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeploymentSetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetId(v string) *DeleteDeploymentSetRequest
	GetDeploymentSetId() *string
	SetOwnerAccount(v string) *DeleteDeploymentSetRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteDeploymentSetRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteDeploymentSetRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteDeploymentSetRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteDeploymentSetRequest
	GetResourceOwnerId() *int64
}

type DeleteDeploymentSetRequest struct {
	// The ID of the deployment set that you want to delete.
	//
	// This parameter is required.
	//
	// example:
	//
	// ds-bp1g5ahlkal88d7x****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s DeleteDeploymentSetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeploymentSetRequest) GoString() string {
	return s.String()
}

func (s *DeleteDeploymentSetRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *DeleteDeploymentSetRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteDeploymentSetRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteDeploymentSetRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteDeploymentSetRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteDeploymentSetRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteDeploymentSetRequest) SetDeploymentSetId(v string) *DeleteDeploymentSetRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *DeleteDeploymentSetRequest) SetOwnerAccount(v string) *DeleteDeploymentSetRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteDeploymentSetRequest) SetOwnerId(v int64) *DeleteDeploymentSetRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteDeploymentSetRequest) SetRegionId(v string) *DeleteDeploymentSetRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteDeploymentSetRequest) SetResourceOwnerAccount(v string) *DeleteDeploymentSetRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteDeploymentSetRequest) SetResourceOwnerId(v int64) *DeleteDeploymentSetRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteDeploymentSetRequest) Validate() error {
	return dara.Validate(s)
}
