// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewaysRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNatGatewayId(v string) *DescribeNatGatewaysRequest
	GetNatGatewayId() *string
	SetOwnerAccount(v string) *DescribeNatGatewaysRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeNatGatewaysRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeNatGatewaysRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNatGatewaysRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeNatGatewaysRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeNatGatewaysRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNatGatewaysRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DescribeNatGatewaysRequest
	GetVpcId() *string
}

type DescribeNatGatewaysRequest struct {
	NatGatewayId *string `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber   *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	VpcId                *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNatGatewaysRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysRequest) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysRequest) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewaysRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeNatGatewaysRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNatGatewaysRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNatGatewaysRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNatGatewaysRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNatGatewaysRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNatGatewaysRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNatGatewaysRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNatGatewaysRequest) SetNatGatewayId(v string) *DescribeNatGatewaysRequest {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetOwnerAccount(v string) *DescribeNatGatewaysRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetOwnerId(v int64) *DescribeNatGatewaysRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetPageNumber(v int32) *DescribeNatGatewaysRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetPageSize(v int32) *DescribeNatGatewaysRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetRegionId(v string) *DescribeNatGatewaysRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetResourceOwnerAccount(v string) *DescribeNatGatewaysRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetResourceOwnerId(v int64) *DescribeNatGatewaysRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) SetVpcId(v string) *DescribeNatGatewaysRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeNatGatewaysRequest) Validate() error {
	return dara.Validate(s)
}
