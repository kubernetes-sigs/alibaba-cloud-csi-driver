// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVRoutersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVRoutersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVRoutersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVRoutersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVRoutersResponseBody
	GetTotalCount() *int32
	SetVRouters(v *DescribeVRoutersResponseBodyVRouters) *DescribeVRoutersResponseBody
	GetVRouters() *DescribeVRoutersResponseBodyVRouters
}

type DescribeVRoutersResponseBody struct {
	PageNumber *int32                                `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                                `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VRouters   *DescribeVRoutersResponseBodyVRouters `json:"VRouters,omitempty" xml:"VRouters,omitempty" type:"Struct"`
}

func (s DescribeVRoutersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVRoutersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVRoutersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVRoutersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVRoutersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVRoutersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVRoutersResponseBody) GetVRouters() *DescribeVRoutersResponseBodyVRouters {
	return s.VRouters
}

func (s *DescribeVRoutersResponseBody) SetPageNumber(v int32) *DescribeVRoutersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVRoutersResponseBody) SetPageSize(v int32) *DescribeVRoutersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVRoutersResponseBody) SetRequestId(v string) *DescribeVRoutersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVRoutersResponseBody) SetTotalCount(v int32) *DescribeVRoutersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVRoutersResponseBody) SetVRouters(v *DescribeVRoutersResponseBodyVRouters) *DescribeVRoutersResponseBody {
	s.VRouters = v
	return s
}

func (s *DescribeVRoutersResponseBody) Validate() error {
	if s.VRouters != nil {
		if err := s.VRouters.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVRoutersResponseBodyVRouters struct {
	VRouter []*DescribeVRoutersResponseBodyVRoutersVRouter `json:"VRouter,omitempty" xml:"VRouter,omitempty" type:"Repeated"`
}

func (s DescribeVRoutersResponseBodyVRouters) String() string {
	return dara.Prettify(s)
}

func (s DescribeVRoutersResponseBodyVRouters) GoString() string {
	return s.String()
}

func (s *DescribeVRoutersResponseBodyVRouters) GetVRouter() []*DescribeVRoutersResponseBodyVRoutersVRouter {
	return s.VRouter
}

func (s *DescribeVRoutersResponseBodyVRouters) SetVRouter(v []*DescribeVRoutersResponseBodyVRoutersVRouter) *DescribeVRoutersResponseBodyVRouters {
	s.VRouter = v
	return s
}

func (s *DescribeVRoutersResponseBodyVRouters) Validate() error {
	if s.VRouter != nil {
		for _, item := range s.VRouter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVRoutersResponseBodyVRoutersVRouter struct {
	CreationTime  *string                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description   *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	RegionId      *string                                                   `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RouteTableIds *DescribeVRoutersResponseBodyVRoutersVRouterRouteTableIds `json:"RouteTableIds,omitempty" xml:"RouteTableIds,omitempty" type:"Struct"`
	VRouterId     *string                                                   `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	VRouterName   *string                                                   `json:"VRouterName,omitempty" xml:"VRouterName,omitempty"`
	VpcId         *string                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeVRoutersResponseBodyVRoutersVRouter) String() string {
	return dara.Prettify(s)
}

func (s DescribeVRoutersResponseBodyVRoutersVRouter) GoString() string {
	return s.String()
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) GetDescription() *string {
	return s.Description
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) GetRouteTableIds() *DescribeVRoutersResponseBodyVRoutersVRouterRouteTableIds {
	return s.RouteTableIds
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) GetVRouterId() *string {
	return s.VRouterId
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) GetVRouterName() *string {
	return s.VRouterName
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) SetCreationTime(v string) *DescribeVRoutersResponseBodyVRoutersVRouter {
	s.CreationTime = &v
	return s
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) SetDescription(v string) *DescribeVRoutersResponseBodyVRoutersVRouter {
	s.Description = &v
	return s
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) SetRegionId(v string) *DescribeVRoutersResponseBodyVRoutersVRouter {
	s.RegionId = &v
	return s
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) SetRouteTableIds(v *DescribeVRoutersResponseBodyVRoutersVRouterRouteTableIds) *DescribeVRoutersResponseBodyVRoutersVRouter {
	s.RouteTableIds = v
	return s
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) SetVRouterId(v string) *DescribeVRoutersResponseBodyVRoutersVRouter {
	s.VRouterId = &v
	return s
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) SetVRouterName(v string) *DescribeVRoutersResponseBodyVRoutersVRouter {
	s.VRouterName = &v
	return s
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) SetVpcId(v string) *DescribeVRoutersResponseBodyVRoutersVRouter {
	s.VpcId = &v
	return s
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouter) Validate() error {
	if s.RouteTableIds != nil {
		if err := s.RouteTableIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVRoutersResponseBodyVRoutersVRouterRouteTableIds struct {
	RouteTableId []*string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty" type:"Repeated"`
}

func (s DescribeVRoutersResponseBodyVRoutersVRouterRouteTableIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVRoutersResponseBodyVRoutersVRouterRouteTableIds) GoString() string {
	return s.String()
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouterRouteTableIds) GetRouteTableId() []*string {
	return s.RouteTableId
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouterRouteTableIds) SetRouteTableId(v []*string) *DescribeVRoutersResponseBodyVRoutersVRouterRouteTableIds {
	s.RouteTableId = v
	return s
}

func (s *DescribeVRoutersResponseBodyVRoutersVRouterRouteTableIds) Validate() error {
	return dara.Validate(s)
}
