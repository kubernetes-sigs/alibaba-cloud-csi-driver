// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualBorderRoutersForPhysicalConnectionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetFilter() []*DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter
	SetOwnerId(v int64) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetPageSize() *int32
	SetPhysicalConnectionId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetPhysicalConnectionId() *string
	SetRegionId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest
	GetResourceOwnerId() *int64
}

type DescribeVirtualBorderRoutersForPhysicalConnectionRequest struct {
	Filter     []*DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerId    *int64                                                            `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32                                                            `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                                            `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	PhysicalConnectionId *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetFilter() []*DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter {
	return s.Filter
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetFilter(v []*DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.Filter = v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetOwnerId(v int64) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetPageNumber(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetPageSize(v int32) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetPhysicalConnectionId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.PhysicalConnectionId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetRegionId(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetResourceOwnerAccount(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) SetResourceOwnerId(v int64) *DescribeVirtualBorderRoutersForPhysicalConnectionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequest) Validate() error {
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

type DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) SetKey(v string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) SetValue(v []*string) *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeVirtualBorderRoutersForPhysicalConnectionRequestFilter) Validate() error {
	return dara.Validate(s)
}
