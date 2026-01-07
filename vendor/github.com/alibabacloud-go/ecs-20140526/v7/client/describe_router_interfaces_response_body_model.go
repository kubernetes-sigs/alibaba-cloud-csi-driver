// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRouterInterfacesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeRouterInterfacesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeRouterInterfacesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeRouterInterfacesResponseBody
	GetRequestId() *string
	SetRouterInterfaceSet(v *DescribeRouterInterfacesResponseBodyRouterInterfaceSet) *DescribeRouterInterfacesResponseBody
	GetRouterInterfaceSet() *DescribeRouterInterfacesResponseBodyRouterInterfaceSet
	SetTotalCount(v int32) *DescribeRouterInterfacesResponseBody
	GetTotalCount() *int32
}

type DescribeRouterInterfacesResponseBody struct {
	PageNumber         *int32                                                  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize           *int32                                                  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId          *string                                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RouterInterfaceSet *DescribeRouterInterfacesResponseBodyRouterInterfaceSet `json:"RouterInterfaceSet,omitempty" xml:"RouterInterfaceSet,omitempty" type:"Struct"`
	TotalCount         *int32                                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeRouterInterfacesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeRouterInterfacesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeRouterInterfacesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRouterInterfacesResponseBody) GetRouterInterfaceSet() *DescribeRouterInterfacesResponseBodyRouterInterfaceSet {
	return s.RouterInterfaceSet
}

func (s *DescribeRouterInterfacesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeRouterInterfacesResponseBody) SetPageNumber(v int32) *DescribeRouterInterfacesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBody) SetPageSize(v int32) *DescribeRouterInterfacesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBody) SetRequestId(v string) *DescribeRouterInterfacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBody) SetRouterInterfaceSet(v *DescribeRouterInterfacesResponseBodyRouterInterfaceSet) *DescribeRouterInterfacesResponseBody {
	s.RouterInterfaceSet = v
	return s
}

func (s *DescribeRouterInterfacesResponseBody) SetTotalCount(v int32) *DescribeRouterInterfacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBody) Validate() error {
	if s.RouterInterfaceSet != nil {
		if err := s.RouterInterfaceSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeRouterInterfacesResponseBodyRouterInterfaceSet struct {
	RouterInterfaceType []*DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType `json:"RouterInterfaceType,omitempty" xml:"RouterInterfaceType,omitempty" type:"Repeated"`
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSet) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSet) GetRouterInterfaceType() []*DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	return s.RouterInterfaceType
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSet) SetRouterInterfaceType(v []*DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) *DescribeRouterInterfacesResponseBodyRouterInterfaceSet {
	s.RouterInterfaceType = v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSet) Validate() error {
	if s.RouterInterfaceType != nil {
		for _, item := range s.RouterInterfaceType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType struct {
	AccessPointId                   *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	BusinessStatus                  *string `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	ChargeType                      *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	ConnectedTime                   *string `json:"ConnectedTime,omitempty" xml:"ConnectedTime,omitempty"`
	CreationTime                    *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description                     *string `json:"Description,omitempty" xml:"Description,omitempty"`
	EndTime                         *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	HealthCheckSourceIp             *string `json:"HealthCheckSourceIp,omitempty" xml:"HealthCheckSourceIp,omitempty"`
	HealthCheckTargetIp             *string `json:"HealthCheckTargetIp,omitempty" xml:"HealthCheckTargetIp,omitempty"`
	Name                            *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OppositeAccessPointId           *string `json:"OppositeAccessPointId,omitempty" xml:"OppositeAccessPointId,omitempty"`
	OppositeInterfaceBusinessStatus *string `json:"OppositeInterfaceBusinessStatus,omitempty" xml:"OppositeInterfaceBusinessStatus,omitempty"`
	OppositeInterfaceId             *string `json:"OppositeInterfaceId,omitempty" xml:"OppositeInterfaceId,omitempty"`
	OppositeInterfaceOwnerId        *string `json:"OppositeInterfaceOwnerId,omitempty" xml:"OppositeInterfaceOwnerId,omitempty"`
	OppositeInterfaceSpec           *string `json:"OppositeInterfaceSpec,omitempty" xml:"OppositeInterfaceSpec,omitempty"`
	OppositeInterfaceStatus         *string `json:"OppositeInterfaceStatus,omitempty" xml:"OppositeInterfaceStatus,omitempty"`
	OppositeRegionId                *string `json:"OppositeRegionId,omitempty" xml:"OppositeRegionId,omitempty"`
	OppositeRouterId                *string `json:"OppositeRouterId,omitempty" xml:"OppositeRouterId,omitempty"`
	OppositeRouterType              *string `json:"OppositeRouterType,omitempty" xml:"OppositeRouterType,omitempty"`
	Role                            *string `json:"Role,omitempty" xml:"Role,omitempty"`
	RouterId                        *string `json:"RouterId,omitempty" xml:"RouterId,omitempty"`
	RouterInterfaceId               *string `json:"RouterInterfaceId,omitempty" xml:"RouterInterfaceId,omitempty"`
	RouterType                      *string `json:"RouterType,omitempty" xml:"RouterType,omitempty"`
	Spec                            *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status                          *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GoString() string {
	return s.String()
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetChargeType() *string {
	return s.ChargeType
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetConnectedTime() *string {
	return s.ConnectedTime
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetDescription() *string {
	return s.Description
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetHealthCheckSourceIp() *string {
	return s.HealthCheckSourceIp
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetHealthCheckTargetIp() *string {
	return s.HealthCheckTargetIp
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetName() *string {
	return s.Name
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeAccessPointId() *string {
	return s.OppositeAccessPointId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeInterfaceBusinessStatus() *string {
	return s.OppositeInterfaceBusinessStatus
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeInterfaceId() *string {
	return s.OppositeInterfaceId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeInterfaceOwnerId() *string {
	return s.OppositeInterfaceOwnerId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeInterfaceSpec() *string {
	return s.OppositeInterfaceSpec
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeInterfaceStatus() *string {
	return s.OppositeInterfaceStatus
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeRegionId() *string {
	return s.OppositeRegionId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeRouterId() *string {
	return s.OppositeRouterId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetOppositeRouterType() *string {
	return s.OppositeRouterType
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetRole() *string {
	return s.Role
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetRouterId() *string {
	return s.RouterId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetRouterInterfaceId() *string {
	return s.RouterInterfaceId
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetRouterType() *string {
	return s.RouterType
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetSpec() *string {
	return s.Spec
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) GetStatus() *string {
	return s.Status
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetAccessPointId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.AccessPointId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetBusinessStatus(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetChargeType(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.ChargeType = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetConnectedTime(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.ConnectedTime = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetCreationTime(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.CreationTime = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetDescription(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Description = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetEndTime(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.EndTime = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetHealthCheckSourceIp(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.HealthCheckSourceIp = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetHealthCheckTargetIp(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.HealthCheckTargetIp = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetName(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Name = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeAccessPointId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeAccessPointId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeInterfaceBusinessStatus(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeInterfaceBusinessStatus = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeInterfaceId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeInterfaceId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeInterfaceOwnerId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeInterfaceOwnerId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeInterfaceSpec(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeInterfaceSpec = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeInterfaceStatus(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeInterfaceStatus = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeRegionId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeRegionId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeRouterId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeRouterId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetOppositeRouterType(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.OppositeRouterType = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetRole(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Role = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetRouterId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.RouterId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetRouterInterfaceId(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.RouterInterfaceId = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetRouterType(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.RouterType = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetSpec(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Spec = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) SetStatus(v string) *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType {
	s.Status = &v
	return s
}

func (s *DescribeRouterInterfacesResponseBodyRouterInterfaceSetRouterInterfaceType) Validate() error {
	return dara.Validate(s)
}
