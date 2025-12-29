// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVirtualBorderRoutersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVirtualBorderRoutersResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVirtualBorderRoutersResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVirtualBorderRoutersResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVirtualBorderRoutersResponseBody
	GetTotalCount() *int32
	SetVirtualBorderRouterSet(v *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) *DescribeVirtualBorderRoutersResponseBody
	GetVirtualBorderRouterSet() *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet
}

type DescribeVirtualBorderRoutersResponseBody struct {
	PageNumber             *int32                                                          `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize               *int32                                                          `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId              *string                                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount             *int32                                                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	VirtualBorderRouterSet *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet `json:"VirtualBorderRouterSet,omitempty" xml:"VirtualBorderRouterSet,omitempty" type:"Struct"`
}

func (s DescribeVirtualBorderRoutersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVirtualBorderRoutersResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVirtualBorderRoutersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVirtualBorderRoutersResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVirtualBorderRoutersResponseBody) GetVirtualBorderRouterSet() *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet {
	return s.VirtualBorderRouterSet
}

func (s *DescribeVirtualBorderRoutersResponseBody) SetPageNumber(v int32) *DescribeVirtualBorderRoutersResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBody) SetPageSize(v int32) *DescribeVirtualBorderRoutersResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBody) SetRequestId(v string) *DescribeVirtualBorderRoutersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBody) SetTotalCount(v int32) *DescribeVirtualBorderRoutersResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBody) SetVirtualBorderRouterSet(v *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) *DescribeVirtualBorderRoutersResponseBody {
	s.VirtualBorderRouterSet = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBody) Validate() error {
	if s.VirtualBorderRouterSet != nil {
		if err := s.VirtualBorderRouterSet.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet struct {
	VirtualBorderRouterType []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType `json:"VirtualBorderRouterType,omitempty" xml:"VirtualBorderRouterType,omitempty" type:"Repeated"`
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) GetVirtualBorderRouterType() []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	return s.VirtualBorderRouterType
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) SetVirtualBorderRouterType(v []*DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet {
	s.VirtualBorderRouterType = v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSet) Validate() error {
	if s.VirtualBorderRouterType != nil {
		for _, item := range s.VirtualBorderRouterType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType struct {
	AccessPointId                    *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	ActivationTime                   *string `json:"ActivationTime,omitempty" xml:"ActivationTime,omitempty"`
	CircuitCode                      *string `json:"CircuitCode,omitempty" xml:"CircuitCode,omitempty"`
	CreationTime                     *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description                      *string `json:"Description,omitempty" xml:"Description,omitempty"`
	LocalGatewayIp                   *string `json:"LocalGatewayIp,omitempty" xml:"LocalGatewayIp,omitempty"`
	Name                             *string `json:"Name,omitempty" xml:"Name,omitempty"`
	PeerGatewayIp                    *string `json:"PeerGatewayIp,omitempty" xml:"PeerGatewayIp,omitempty"`
	PeeringSubnetMask                *string `json:"PeeringSubnetMask,omitempty" xml:"PeeringSubnetMask,omitempty"`
	PhysicalConnectionBusinessStatus *string `json:"PhysicalConnectionBusinessStatus,omitempty" xml:"PhysicalConnectionBusinessStatus,omitempty"`
	PhysicalConnectionId             *string `json:"PhysicalConnectionId,omitempty" xml:"PhysicalConnectionId,omitempty"`
	PhysicalConnectionOwnerUid       *string `json:"PhysicalConnectionOwnerUid,omitempty" xml:"PhysicalConnectionOwnerUid,omitempty"`
	PhysicalConnectionStatus         *string `json:"PhysicalConnectionStatus,omitempty" xml:"PhysicalConnectionStatus,omitempty"`
	RecoveryTime                     *string `json:"RecoveryTime,omitempty" xml:"RecoveryTime,omitempty"`
	RouteTableId                     *string `json:"RouteTableId,omitempty" xml:"RouteTableId,omitempty"`
	Status                           *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TerminationTime                  *string `json:"TerminationTime,omitempty" xml:"TerminationTime,omitempty"`
	VbrId                            *string `json:"VbrId,omitempty" xml:"VbrId,omitempty"`
	VlanId                           *int32  `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	VlanInterfaceId                  *string `json:"VlanInterfaceId,omitempty" xml:"VlanInterfaceId,omitempty"`
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) String() string {
	return dara.Prettify(s)
}

func (s DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GoString() string {
	return s.String()
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetActivationTime() *string {
	return s.ActivationTime
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetCircuitCode() *string {
	return s.CircuitCode
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetDescription() *string {
	return s.Description
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetLocalGatewayIp() *string {
	return s.LocalGatewayIp
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetName() *string {
	return s.Name
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPeerGatewayIp() *string {
	return s.PeerGatewayIp
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPeeringSubnetMask() *string {
	return s.PeeringSubnetMask
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPhysicalConnectionBusinessStatus() *string {
	return s.PhysicalConnectionBusinessStatus
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPhysicalConnectionId() *string {
	return s.PhysicalConnectionId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPhysicalConnectionOwnerUid() *string {
	return s.PhysicalConnectionOwnerUid
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetPhysicalConnectionStatus() *string {
	return s.PhysicalConnectionStatus
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetRecoveryTime() *string {
	return s.RecoveryTime
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetRouteTableId() *string {
	return s.RouteTableId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetStatus() *string {
	return s.Status
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetTerminationTime() *string {
	return s.TerminationTime
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetVbrId() *string {
	return s.VbrId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetVlanId() *int32 {
	return s.VlanId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) GetVlanInterfaceId() *string {
	return s.VlanInterfaceId
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetAccessPointId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.AccessPointId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetActivationTime(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.ActivationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetCircuitCode(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.CircuitCode = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetCreationTime(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.CreationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetDescription(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.Description = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetLocalGatewayIp(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.LocalGatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetName(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.Name = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPeerGatewayIp(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PeerGatewayIp = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPeeringSubnetMask(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PeeringSubnetMask = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPhysicalConnectionBusinessStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PhysicalConnectionBusinessStatus = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPhysicalConnectionId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PhysicalConnectionId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPhysicalConnectionOwnerUid(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PhysicalConnectionOwnerUid = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetPhysicalConnectionStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.PhysicalConnectionStatus = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetRecoveryTime(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.RecoveryTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetRouteTableId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.RouteTableId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetStatus(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.Status = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetTerminationTime(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.TerminationTime = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetVbrId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.VbrId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetVlanId(v int32) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.VlanId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) SetVlanInterfaceId(v string) *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType {
	s.VlanInterfaceId = &v
	return s
}

func (s *DescribeVirtualBorderRoutersResponseBodyVirtualBorderRouterSetVirtualBorderRouterType) Validate() error {
	return dara.Validate(s)
}
