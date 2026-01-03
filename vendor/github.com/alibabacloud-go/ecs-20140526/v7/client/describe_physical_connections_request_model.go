// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePhysicalConnectionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribePhysicalConnectionsRequest
	GetClientToken() *string
	SetFilter(v []*DescribePhysicalConnectionsRequestFilter) *DescribePhysicalConnectionsRequest
	GetFilter() []*DescribePhysicalConnectionsRequestFilter
	SetOwnerAccount(v string) *DescribePhysicalConnectionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePhysicalConnectionsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribePhysicalConnectionsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribePhysicalConnectionsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribePhysicalConnectionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribePhysicalConnectionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePhysicalConnectionsRequest
	GetResourceOwnerId() *int64
	SetUserCidr(v string) *DescribePhysicalConnectionsRequest
	GetUserCidr() *string
}

type DescribePhysicalConnectionsRequest struct {
	ClientToken  *string                                     `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Filter       []*DescribePhysicalConnectionsRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	OwnerAccount *string                                     `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64                                      `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber   *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize     *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	UserCidr             *string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty"`
}

func (s DescribePhysicalConnectionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsRequest) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribePhysicalConnectionsRequest) GetFilter() []*DescribePhysicalConnectionsRequestFilter {
	return s.Filter
}

func (s *DescribePhysicalConnectionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePhysicalConnectionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePhysicalConnectionsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribePhysicalConnectionsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribePhysicalConnectionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePhysicalConnectionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePhysicalConnectionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePhysicalConnectionsRequest) GetUserCidr() *string {
	return s.UserCidr
}

func (s *DescribePhysicalConnectionsRequest) SetClientToken(v string) *DescribePhysicalConnectionsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetFilter(v []*DescribePhysicalConnectionsRequestFilter) *DescribePhysicalConnectionsRequest {
	s.Filter = v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetOwnerAccount(v string) *DescribePhysicalConnectionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetOwnerId(v int64) *DescribePhysicalConnectionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetPageNumber(v int32) *DescribePhysicalConnectionsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetPageSize(v int32) *DescribePhysicalConnectionsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetRegionId(v string) *DescribePhysicalConnectionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetResourceOwnerAccount(v string) *DescribePhysicalConnectionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetResourceOwnerId(v int64) *DescribePhysicalConnectionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) SetUserCidr(v string) *DescribePhysicalConnectionsRequest {
	s.UserCidr = &v
	return s
}

func (s *DescribePhysicalConnectionsRequest) Validate() error {
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

type DescribePhysicalConnectionsRequestFilter struct {
	Key   *string   `json:"Key,omitempty" xml:"Key,omitempty"`
	Value []*string `json:"Value,omitempty" xml:"Value,omitempty" type:"Repeated"`
}

func (s DescribePhysicalConnectionsRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribePhysicalConnectionsRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribePhysicalConnectionsRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribePhysicalConnectionsRequestFilter) GetValue() []*string {
	return s.Value
}

func (s *DescribePhysicalConnectionsRequestFilter) SetKey(v string) *DescribePhysicalConnectionsRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribePhysicalConnectionsRequestFilter) SetValue(v []*string) *DescribePhysicalConnectionsRequestFilter {
	s.Value = v
	return s
}

func (s *DescribePhysicalConnectionsRequestFilter) Validate() error {
	return dara.Validate(s)
}
