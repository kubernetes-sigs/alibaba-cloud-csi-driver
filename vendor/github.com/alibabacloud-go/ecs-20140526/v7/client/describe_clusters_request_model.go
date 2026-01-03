// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeClustersRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeClustersRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeClustersRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeClustersRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeClustersRequest
	GetResourceOwnerId() *int64
}

type DescribeClustersRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeClustersRequest) GoString() string {
	return s.String()
}

func (s *DescribeClustersRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeClustersRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeClustersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeClustersRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeClustersRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeClustersRequest) SetOwnerAccount(v string) *DescribeClustersRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeClustersRequest) SetOwnerId(v int64) *DescribeClustersRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeClustersRequest) SetRegionId(v string) *DescribeClustersRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeClustersRequest) SetResourceOwnerAccount(v string) *DescribeClustersRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeClustersRequest) SetResourceOwnerId(v int64) *DescribeClustersRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeClustersRequest) Validate() error {
	return dara.Validate(s)
}
