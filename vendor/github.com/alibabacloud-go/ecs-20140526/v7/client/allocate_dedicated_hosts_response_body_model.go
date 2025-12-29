// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAllocateDedicatedHostsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostIdSets(v *AllocateDedicatedHostsResponseBodyDedicatedHostIdSets) *AllocateDedicatedHostsResponseBody
	GetDedicatedHostIdSets() *AllocateDedicatedHostsResponseBodyDedicatedHostIdSets
	SetOrderId(v string) *AllocateDedicatedHostsResponseBody
	GetOrderId() *string
	SetRequestId(v string) *AllocateDedicatedHostsResponseBody
	GetRequestId() *string
}

type AllocateDedicatedHostsResponseBody struct {
	// A list of dedicated host IDs.
	DedicatedHostIdSets *AllocateDedicatedHostsResponseBodyDedicatedHostIdSets `json:"DedicatedHostIdSets,omitempty" xml:"DedicatedHostIdSets,omitempty" type:"Struct"`
	// The ID of the order.
	//
	// >  This parameter has a return value only when the dedicated host is a subscription one (request parameter **ChargeType set to PrePaid**).
	//
	// example:
	//
	// 23841229****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E2A664A6-2933-4C64-88AE-5033D003****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AllocateDedicatedHostsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AllocateDedicatedHostsResponseBody) GoString() string {
	return s.String()
}

func (s *AllocateDedicatedHostsResponseBody) GetDedicatedHostIdSets() *AllocateDedicatedHostsResponseBodyDedicatedHostIdSets {
	return s.DedicatedHostIdSets
}

func (s *AllocateDedicatedHostsResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *AllocateDedicatedHostsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AllocateDedicatedHostsResponseBody) SetDedicatedHostIdSets(v *AllocateDedicatedHostsResponseBodyDedicatedHostIdSets) *AllocateDedicatedHostsResponseBody {
	s.DedicatedHostIdSets = v
	return s
}

func (s *AllocateDedicatedHostsResponseBody) SetOrderId(v string) *AllocateDedicatedHostsResponseBody {
	s.OrderId = &v
	return s
}

func (s *AllocateDedicatedHostsResponseBody) SetRequestId(v string) *AllocateDedicatedHostsResponseBody {
	s.RequestId = &v
	return s
}

func (s *AllocateDedicatedHostsResponseBody) Validate() error {
	if s.DedicatedHostIdSets != nil {
		if err := s.DedicatedHostIdSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type AllocateDedicatedHostsResponseBodyDedicatedHostIdSets struct {
	DedicatedHostId []*string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty" type:"Repeated"`
}

func (s AllocateDedicatedHostsResponseBodyDedicatedHostIdSets) String() string {
	return dara.Prettify(s)
}

func (s AllocateDedicatedHostsResponseBodyDedicatedHostIdSets) GoString() string {
	return s.String()
}

func (s *AllocateDedicatedHostsResponseBodyDedicatedHostIdSets) GetDedicatedHostId() []*string {
	return s.DedicatedHostId
}

func (s *AllocateDedicatedHostsResponseBodyDedicatedHostIdSets) SetDedicatedHostId(v []*string) *AllocateDedicatedHostsResponseBodyDedicatedHostIdSets {
	s.DedicatedHostId = v
	return s
}

func (s *AllocateDedicatedHostsResponseBodyDedicatedHostIdSets) Validate() error {
	return dara.Validate(s)
}
