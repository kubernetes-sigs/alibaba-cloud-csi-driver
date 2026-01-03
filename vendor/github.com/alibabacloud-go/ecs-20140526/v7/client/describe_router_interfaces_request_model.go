// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouterInterfacesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilter(v []*DescribeRouterInterfacesRequestFilter) *DescribeRouterInterfacesRequest
	GetFilter() []*DescribeRouterInterfacesRequestFilter
	SetOwnerId(v int64) *DescribeRouterInterfacesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRouterInterfacesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouterInterfacesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRouterInterfacesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRouterInterfacesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRouterInterfacesRequest
	GetResourceOwnerId() *int64
}

type DescribeRouterInterfacesRequest struct {
	Filter     []*DescribeRouterInterfacesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerId    *int64                                   `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber *int32                                   `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                   `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeRouterInterfacesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesRequest) GetFilter() []*DescribeRouterInterfacesRequestFilter {
	return s.Filter
}

func (s *DescribeRouterInterfacesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRouterInterfacesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouterInterfacesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouterInterfacesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouterInterfacesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRouterInterfacesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRouterInterfacesRequest) SetFilter(v []*DescribeRouterInterfacesRequestFilter) *DescribeRouterInterfacesRequest {
	s.Filter = v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetOwnerId(v int64) *DescribeRouterInterfacesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetPageNumber(v int32) *DescribeRouterInterfacesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetPageSize(v int32) *DescribeRouterInterfacesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetRegionId(v string) *DescribeRouterInterfacesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetResourceOwnerAccount(v string) *DescribeRouterInterfacesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) SetResourceOwnerId(v int64) *DescribeRouterInterfacesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouterInterfacesRequest) Validate() error {
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

type DescribeRouterInterfacesRequestFilter struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribeRouterInterfacesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeRouterInterfacesRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribeRouterInterfacesRequestFilter) SetKey(v string) *DescribeRouterInterfacesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeRouterInterfacesRequestFilter) SetValue(v []*string) *DescribeRouterInterfacesRequestFilter {
	s.Value = v
	return s
}

func (s *DescribeRouterInterfacesRequestFilter) Validate() error {
	return dara.Validate(s)
}
