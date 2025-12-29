// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNatGatewayResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetBandwidthPackageIds(v *CreateNatGatewayResponseBodyBandwidthPackageIds) *CreateNatGatewayResponseBody
	GetBandwidthPackageIds() *CreateNatGatewayResponseBodyBandwidthPackageIds
	SetForwardTableIds(v *CreateNatGatewayResponseBodyForwardTableIds) *CreateNatGatewayResponseBody
	GetForwardTableIds() *CreateNatGatewayResponseBodyForwardTableIds
	SetNatGatewayId(v string) *CreateNatGatewayResponseBody
	GetNatGatewayId() *string
	SetRequestId(v string) *CreateNatGatewayResponseBody
	GetRequestId() *string
}

type CreateNatGatewayResponseBody struct {
	BandwidthPackageIds *CreateNatGatewayResponseBodyBandwidthPackageIds `json:"BandwidthPackageIds,omitempty" xml:"BandwidthPackageIds,omitempty" type:"Struct"`
	ForwardTableIds     *CreateNatGatewayResponseBodyForwardTableIds     `json:"ForwardTableIds,omitempty" xml:"ForwardTableIds,omitempty" type:"Struct"`
	NatGatewayId        *string                                          `json:"NatGatewayId,omitempty" xml:"NatGatewayId,omitempty"`
	RequestId           *string                                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNatGatewayResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayResponseBody) GetBandwidthPackageIds() *CreateNatGatewayResponseBodyBandwidthPackageIds {
	return s.BandwidthPackageIds
}

func (s *CreateNatGatewayResponseBody) GetForwardTableIds() *CreateNatGatewayResponseBodyForwardTableIds {
	return s.ForwardTableIds
}

func (s *CreateNatGatewayResponseBody) GetNatGatewayId() *string {
	return s.NatGatewayId
}

func (s *CreateNatGatewayResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNatGatewayResponseBody) SetBandwidthPackageIds(v *CreateNatGatewayResponseBodyBandwidthPackageIds) *CreateNatGatewayResponseBody {
	s.BandwidthPackageIds = v
	return s
}

func (s *CreateNatGatewayResponseBody) SetForwardTableIds(v *CreateNatGatewayResponseBodyForwardTableIds) *CreateNatGatewayResponseBody {
	s.ForwardTableIds = v
	return s
}

func (s *CreateNatGatewayResponseBody) SetNatGatewayId(v string) *CreateNatGatewayResponseBody {
	s.NatGatewayId = &v
	return s
}

func (s *CreateNatGatewayResponseBody) SetRequestId(v string) *CreateNatGatewayResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNatGatewayResponseBody) Validate() error {
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

type CreateNatGatewayResponseBodyBandwidthPackageIds struct {
	BandwidthPackageId []*string `json:"BandwidthPackageId,omitempty" xml:"BandwidthPackageId,omitempty" type:"Repeated"`
}

func (s CreateNatGatewayResponseBodyBandwidthPackageIds) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayResponseBodyBandwidthPackageIds) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayResponseBodyBandwidthPackageIds) GetBandwidthPackageId() []*string {
	return s.BandwidthPackageId
}

func (s *CreateNatGatewayResponseBodyBandwidthPackageIds) SetBandwidthPackageId(v []*string) *CreateNatGatewayResponseBodyBandwidthPackageIds {
	s.BandwidthPackageId = v
	return s
}

func (s *CreateNatGatewayResponseBodyBandwidthPackageIds) Validate() error {
	return dara.Validate(s)
}

type CreateNatGatewayResponseBodyForwardTableIds struct {
	ForwardTableId []*string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty" type:"Repeated"`
}

func (s CreateNatGatewayResponseBodyForwardTableIds) String() string {
	return dara.Prettify(s)
}

func (s CreateNatGatewayResponseBodyForwardTableIds) GoString() string {
	return s.String()
}

func (s *CreateNatGatewayResponseBodyForwardTableIds) GetForwardTableId() []*string {
	return s.ForwardTableId
}

func (s *CreateNatGatewayResponseBodyForwardTableIds) SetForwardTableId(v []*string) *CreateNatGatewayResponseBodyForwardTableIds {
	s.ForwardTableId = v
	return s
}

func (s *CreateNatGatewayResponseBodyForwardTableIds) Validate() error {
	return dara.Validate(s)
}
