// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLimitationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimitation(v string) *DescribeLimitationRequest
	GetLimitation() *string
	SetOwnerAccount(v string) *DescribeLimitationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLimitationRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *DescribeLimitationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLimitationRequest
	GetResourceOwnerId() *int64
}

type DescribeLimitationRequest struct {
	// This parameter is required.
	Limitation           *string `json:"Limitation,omitempty" xml:"Limitation,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeLimitationRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLimitationRequest) GoString() string {
	return s.String()
}

func (s *DescribeLimitationRequest) GetLimitation() *string {
	return s.Limitation
}

func (s *DescribeLimitationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLimitationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLimitationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLimitationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLimitationRequest) SetLimitation(v string) *DescribeLimitationRequest {
	s.Limitation = &v
	return s
}

func (s *DescribeLimitationRequest) SetOwnerAccount(v string) *DescribeLimitationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLimitationRequest) SetOwnerId(v int64) *DescribeLimitationRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLimitationRequest) SetResourceOwnerAccount(v string) *DescribeLimitationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLimitationRequest) SetResourceOwnerId(v int64) *DescribeLimitationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLimitationRequest) Validate() error {
	return dara.Validate(s)
}
