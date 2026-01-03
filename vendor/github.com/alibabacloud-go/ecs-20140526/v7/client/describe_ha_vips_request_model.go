// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeHaVipsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeHaVipsRequestFilter) *DescribeHaVipsRequest
	GetFilter() []*DescribeHaVipsRequestFilter
	SetOwnerAccount(v string) *DescribeHaVipsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeHaVipsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeHaVipsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeHaVipsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeHaVipsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeHaVipsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeHaVipsRequest
	GetResourceOwnerId() *int64
}

type DescribeHaVipsRequest struct {
	// This parameter is required.
	Filter       []*DescribeHaVipsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerAccount *string                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber   *int32                         `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                         `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeHaVipsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsRequest) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsRequest) GetFilter() []*DescribeHaVipsRequestFilter {
	return s.Filter
}

func (s *DescribeHaVipsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeHaVipsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeHaVipsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeHaVipsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeHaVipsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeHaVipsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeHaVipsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeHaVipsRequest) SetFilter(v []*DescribeHaVipsRequestFilter) *DescribeHaVipsRequest {
	s.Filter = v
	return s
}

func (s *DescribeHaVipsRequest) SetOwnerAccount(v string) *DescribeHaVipsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeHaVipsRequest) SetOwnerId(v int64) *DescribeHaVipsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeHaVipsRequest) SetPageNumber(v int32) *DescribeHaVipsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeHaVipsRequest) SetPageSize(v int32) *DescribeHaVipsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeHaVipsRequest) SetRegionId(v string) *DescribeHaVipsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeHaVipsRequest) SetResourceOwnerAccount(v string) *DescribeHaVipsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeHaVipsRequest) SetResourceOwnerId(v int64) *DescribeHaVipsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeHaVipsRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeHaVipsRequestFilter struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeHaVipsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeHaVipsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeHaVipsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeHaVipsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeHaVipsRequestFilter) SetKey(v string) *DescribeHaVipsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeHaVipsRequestFilter) SetValue(v []*string) *DescribeHaVipsRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeHaVipsRequestFilter) Validate() error {
	return dara.Validate(s)
}
