// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserBusinessBehaviorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeUserBusinessBehaviorRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeUserBusinessBehaviorRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeUserBusinessBehaviorRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeUserBusinessBehaviorRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeUserBusinessBehaviorRequest
	GetResourceOwnerId() *int64
	SetStatusKey(v string) *DescribeUserBusinessBehaviorRequest
	GetStatusKey() *string
}

type DescribeUserBusinessBehaviorRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StatusKey *string `json:"statusKey,omitempty" xml:"statusKey,omitempty"`
}

func (s DescribeUserBusinessBehaviorRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserBusinessBehaviorRequest) GoString() string {
	return s.String()
}

func (s *DescribeUserBusinessBehaviorRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeUserBusinessBehaviorRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeUserBusinessBehaviorRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeUserBusinessBehaviorRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeUserBusinessBehaviorRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeUserBusinessBehaviorRequest) GetStatusKey() *string {
	return s.StatusKey
}

func (s *DescribeUserBusinessBehaviorRequest) SetOwnerAccount(v string) *DescribeUserBusinessBehaviorRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeUserBusinessBehaviorRequest) SetOwnerId(v int64) *DescribeUserBusinessBehaviorRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeUserBusinessBehaviorRequest) SetRegionId(v string) *DescribeUserBusinessBehaviorRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeUserBusinessBehaviorRequest) SetResourceOwnerAccount(v string) *DescribeUserBusinessBehaviorRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeUserBusinessBehaviorRequest) SetResourceOwnerId(v int64) *DescribeUserBusinessBehaviorRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeUserBusinessBehaviorRequest) SetStatusKey(v string) *DescribeUserBusinessBehaviorRequest {
	s.StatusKey = &v
	return s
}

func (s *DescribeUserBusinessBehaviorRequest) Validate() error {
	return dara.Validate(s)
}
