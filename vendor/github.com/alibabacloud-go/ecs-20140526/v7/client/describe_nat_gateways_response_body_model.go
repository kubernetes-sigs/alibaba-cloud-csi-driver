// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNatGatewaysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNatGateways(v *DescribeNatGatewaysResponseBodyNatGateways) *DescribeNatGatewaysResponseBody
	GetNatGateways() *DescribeNatGatewaysResponseBodyNatGateways
	SetPageNumber(v int32) *DescribeNatGatewaysResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeNatGatewaysResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeNatGatewaysResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeNatGatewaysResponseBody
	GetTotalCount() *int32
}

type DescribeNatGatewaysResponseBody struct {
	NatGateways *DescribeNatGatewaysResponseBodyNatGateways `json:"NatGateways,omitempty" xml:"NatGateways,omitempty" type:"Struct"`
	PageNumber  *int32                                      `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize    *int32                                      `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId   *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeNatGatewaysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBody) GetNatGateways() *DescribeNatGatewaysResponseBodyNatGateways {
	return s.NatGateways
}

func (s *DescribeNatGatewaysResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeNatGatewaysResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeNatGatewaysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNatGatewaysResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeNatGatewaysResponseBody) SetNatGateways(v *DescribeNatGatewaysResponseBodyNatGateways) *DescribeNatGatewaysResponseBody {
	s.NatGateways = v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetPageNumber(v int32) *DescribeNatGatewaysResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetPageSize(v int32) *DescribeNatGatewaysResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetRequestId(v string) *DescribeNatGatewaysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) SetTotalCount(v int32) *DescribeNatGatewaysResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeNatGatewaysResponseBody) Validate() error {
	if s.NatGateways != nil {
		if err := s.NatGateways.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNatGatewaysResponseBodyNatGateways struct {
	NatGateway []*DescribeNatGatewaysResponseBodyNatGatewaysNatGateway `json:"NatGateway,omitempty" xml:"NatGateway,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewaysResponseBodyNatGateways) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGateways) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) GetNatGateway() []*DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	return s.NatGateway
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) SetNatGateway(v []*DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) *DescribeNatGatewaysResponseBodyNatGateways {
	s.NatGateway = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGateways) Validate() error {
	if s.NatGateway != nil {
		for _, item := range s.NatGateway {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGateway struct {
	BandwidthPackageIds *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayBandwidthPackageIds `json:"BandwidthPackageIds,omitempty" xml:"BandwidthPackageIds,omitempty" type:"Struct"`
	BusinessStatus      *string                                                                  `json:"BusinessStatus,omitempty" xml:"BusinessStatus,omitempty"`
	CreationTime        *string                                                                  `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description         *string                                                                  `json:"Description,omitempty" xml:"Description,omitempty"`
	ForwardTableIds     *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds     `json:"ForwardTableIds,omitempty" xml:"ForwardTableIds,omitempty" type:"Struct"`
	InstanceChargeType  *string                                                                  `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	Name                *string                                                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	NatGatewayId        *string                                                                  `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	RegionId            *string                                                                  `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Spec                *string                                                                  `json:"Spec,omitempty" xml:"Spec,omitempty"`
	Status              *string                                                                  `json:"Status,omitempty" xml:"Status,omitempty"`
	VpcId               *string                                                                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetBandwidthPackageIds() *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayBandwidthPackageIds {
	return s.BandwidthPackageIds
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetBusinessStatus() *string {
	return s.BusinessStatus
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetDescription() *string {
	return s.Description
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetForwardTableIds() *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds {
	return s.ForwardTableIds
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetName() *string {
	return s.Name
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetSpec() *string {
	return s.Spec
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetStatus() *string {
	return s.Status
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetBandwidthPackageIds(v *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayBandwidthPackageIds) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.BandwidthPackageIds = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetBusinessStatus(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.BusinessStatus = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetCreationTime(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.CreationTime = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetDescription(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.Description = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetForwardTableIds(v *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.ForwardTableIds = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetInstanceChargeType(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetName(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.Name = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetNatGatewayId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.NatGatewayId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetRegionId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.RegionId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetSpec(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.Spec = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetStatus(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.Status = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) SetVpcId(v string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway {
	s.VpcId = &v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGateway) Validate() error {
	if s.BandwidthPackageIds != nil {
		if err := s.BandwidthPackageIds.Validate(); err != nil {
			return err
		}
	}
	if s.ForwardTableIds != nil {
		if err := s.ForwardTableIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayBandwidthPackageIds struct {
	BandwidthPackageId []*string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayBandwidthPackageIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayBandwidthPackageIds) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayBandwidthPackageIds) GetBandwidthPackageId() []*string {
	return s.BandwidthPackageId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayBandwidthPackageIds) SetBandwidthPackageId(v []*string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayBandwidthPackageIds {
	s.BandwidthPackageId = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayBandwidthPackageIds) Validate() error {
	return dara.Validate(s)
}

type DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds struct {
	ForwardTableId []*string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty" type:"Repeated"`
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) GoString() string {
	return s.String()
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) GetForwardTableId() []*string {
	return s.ForwardTableId
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) SetForwardTableId(v []*string) *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds {
	s.ForwardTableId = v
	return s
}

func (s *DescribeNatGatewaysResponseBodyNatGatewaysNatGatewayForwardTableIds) Validate() error {
	return dara.Validate(s)
}
