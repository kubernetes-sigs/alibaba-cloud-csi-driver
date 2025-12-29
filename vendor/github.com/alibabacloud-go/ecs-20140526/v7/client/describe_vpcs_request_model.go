// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIsDefault(v bool) *DescribeVpcsRequest
	GetIsDefault() *bool
	SetOwnerAccount(v string) *DescribeVpcsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVpcsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVpcsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpcsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeVpcsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVpcsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVpcsRequest
	GetResourceOwnerId() *int64
	SetVpcId(v string) *DescribeVpcsRequest
	GetVpcId() *string
}

type DescribeVpcsRequest struct {
	IsDefault    *bool   `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
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

func (s DescribeVpcsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVpcsRequest) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVpcsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVpcsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpcsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVpcsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVpcsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcsRequest) SetIsDefault(v bool) *DescribeVpcsRequest {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsRequest) SetOwnerAccount(v string) *DescribeVpcsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVpcsRequest) SetOwnerId(v int64) *DescribeVpcsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVpcsRequest) SetPageNumber(v int32) *DescribeVpcsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcsRequest) SetPageSize(v int32) *DescribeVpcsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcsRequest) SetRegionId(v string) *DescribeVpcsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcsRequest) SetResourceOwnerAccount(v string) *DescribeVpcsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVpcsRequest) SetResourceOwnerId(v int64) *DescribeVpcsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVpcsRequest) SetVpcId(v string) *DescribeVpcsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcsRequest) Validate() error {
	return dara.Validate(s)
}
