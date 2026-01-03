// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRenewReservedInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *RenewReservedInstancesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RenewReservedInstancesResponseBody
	GetRequestId() *string
	SetReservedInstanceIdSets(v *RenewReservedInstancesResponseBodyReservedInstanceIdSets) *RenewReservedInstancesResponseBody
	GetReservedInstanceIdSets() *RenewReservedInstancesResponseBodyReservedInstanceIdSets
}

type RenewReservedInstancesResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 2023912123****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8C314443-AF0D-4766-9562-C83B7F1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the reserved instances.
	ReservedInstanceIdSets *RenewReservedInstancesResponseBodyReservedInstanceIdSets `json:"ReservedInstanceIdSets,omitempty" xml:"ReservedInstanceIdSets,omitempty" type:"Struct"`
}

func (s RenewReservedInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RenewReservedInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RenewReservedInstancesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RenewReservedInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RenewReservedInstancesResponseBody) GetReservedInstanceIdSets() *RenewReservedInstancesResponseBodyReservedInstanceIdSets {
	return s.ReservedInstanceIdSets
}

func (s *RenewReservedInstancesResponseBody) SetOrderId(v string) *RenewReservedInstancesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RenewReservedInstancesResponseBody) SetRequestId(v string) *RenewReservedInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenewReservedInstancesResponseBody) SetReservedInstanceIdSets(v *RenewReservedInstancesResponseBodyReservedInstanceIdSets) *RenewReservedInstancesResponseBody {
	s.ReservedInstanceIdSets = v
	return s
}

func (s *RenewReservedInstancesResponseBody) Validate() error {
	if s.ReservedInstanceIdSets != nil {
		if err := s.ReservedInstanceIdSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RenewReservedInstancesResponseBodyReservedInstanceIdSets struct {
	ReservedInstanceId []*string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty" type:"Repeated"`
}

func (s RenewReservedInstancesResponseBodyReservedInstanceIdSets) String() string {
	return dara.Prettify(s)
}

func (s RenewReservedInstancesResponseBodyReservedInstanceIdSets) GoString() string {
	return s.String()
}

func (s *RenewReservedInstancesResponseBodyReservedInstanceIdSets) GetReservedInstanceId() []*string {
	return s.ReservedInstanceId
}

func (s *RenewReservedInstancesResponseBodyReservedInstanceIdSets) SetReservedInstanceId(v []*string) *RenewReservedInstancesResponseBodyReservedInstanceIdSets {
	s.ReservedInstanceId = v
	return s
}

func (s *RenewReservedInstancesResponseBodyReservedInstanceIdSets) Validate() error {
	return dara.Validate(s)
}
