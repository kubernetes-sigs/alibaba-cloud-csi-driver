// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupReferencesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSecurityGroupReferencesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSecurityGroupReferencesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSecurityGroupReferencesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSecurityGroupReferencesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSecurityGroupReferencesRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v []*string) *DescribeSecurityGroupReferencesRequest
	GetSecurityGroupId() []*string
}

type DescribeSecurityGroupReferencesRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of security groups. You can specify up to 10 IDs of security groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp14vtedjtobkvi****
	SecurityGroupId []*string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupReferencesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupReferencesRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupReferencesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSecurityGroupReferencesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSecurityGroupReferencesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityGroupReferencesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSecurityGroupReferencesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecurityGroupReferencesRequest) GetSecurityGroupId() []*string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupReferencesRequest) SetOwnerAccount(v string) *DescribeSecurityGroupReferencesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupReferencesRequest) SetOwnerId(v int64) *DescribeSecurityGroupReferencesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSecurityGroupReferencesRequest) SetRegionId(v string) *DescribeSecurityGroupReferencesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityGroupReferencesRequest) SetResourceOwnerAccount(v string) *DescribeSecurityGroupReferencesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupReferencesRequest) SetResourceOwnerId(v int64) *DescribeSecurityGroupReferencesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityGroupReferencesRequest) SetSecurityGroupId(v []*string) *DescribeSecurityGroupReferencesRequest {
	s.SecurityGroupId = v
	return s
}

func (s *DescribeSecurityGroupReferencesRequest) Validate() error {
	return dara.Validate(s)
}
