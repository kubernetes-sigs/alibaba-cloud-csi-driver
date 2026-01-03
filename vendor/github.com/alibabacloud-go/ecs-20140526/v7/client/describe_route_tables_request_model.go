// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouteTablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeRouteTablesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeRouteTablesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeRouteTablesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouteTablesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeRouteTablesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeRouteTablesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeRouteTablesRequest
	GetResourceOwnerId() *int64
	SetRouteTableId(v string) *DescribeRouteTablesRequest
	GetRouteTableId() *string
	SetRouteTableName(v string) *DescribeRouteTablesRequest
	GetRouteTableName() *string
	SetRouterId(v string) *DescribeRouteTablesRequest
	GetRouterId() *string
	SetRouterType(v string) *DescribeRouteTablesRequest
	GetRouterType() *string
	SetVRouterId(v string) *DescribeRouteTablesRequest
	GetVRouterId() *string
}

type DescribeRouteTablesRequest struct {
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber           *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize             *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	RouteTableId         *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	RouteTableName       *string `json:"RouteTableName,omitempty" xml:"RouteTableName,omitempty"`
	RouterId             *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	RouterType           *string `json:"RouterType,omitempty" xml:"RouterType,omitempty"`
	VRouterId            *string `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
}

func (s DescribeRouteTablesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouteTablesRequest) GoString() string {
	return s.String()
}

func (s *DescribeRouteTablesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeRouteTablesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeRouteTablesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouteTablesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouteTablesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRouteTablesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeRouteTablesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeRouteTablesRequest) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeRouteTablesRequest) GetRouteTableName() *string {
	return s.RouteTableName
}

func (s *DescribeRouteTablesRequest) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeRouteTablesRequest) GetRouterType() *string {
	return s.RouterType
}

func (s *DescribeRouteTablesRequest) GetVRouterId() *string {
	return s.VRouterId
}

func (s *DescribeRouteTablesRequest) SetOwnerAccount(v string) *DescribeRouteTablesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetOwnerId(v int64) *DescribeRouteTablesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetPageNumber(v int32) *DescribeRouteTablesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetPageSize(v int32) *DescribeRouteTablesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetRegionId(v string) *DescribeRouteTablesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetResourceOwnerAccount(v string) *DescribeRouteTablesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetResourceOwnerId(v int64) *DescribeRouteTablesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetRouteTableId(v string) *DescribeRouteTablesRequest {
	s.RouteTableId = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetRouteTableName(v string) *DescribeRouteTablesRequest {
	s.RouteTableName = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetRouterId(v string) *DescribeRouteTablesRequest {
	s.RouterId = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetRouterType(v string) *DescribeRouteTablesRequest {
	s.RouterType = &v
	return s
}

func (s *DescribeRouteTablesRequest) SetVRouterId(v string) *DescribeRouteTablesRequest {
	s.VRouterId = &v
	return s
}

func (s *DescribeRouteTablesRequest) Validate() error {
	return dara.Validate(s)
}
