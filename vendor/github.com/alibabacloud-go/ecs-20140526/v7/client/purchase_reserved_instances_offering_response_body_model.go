// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseReservedInstancesOfferingResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOrderId(v string) *PurchaseReservedInstancesOfferingResponseBody
	GetOrderId() *string
	SetRequestId(v string) *PurchaseReservedInstancesOfferingResponseBody
	GetRequestId() *string
	SetReservedInstanceIdSets(v *PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets) *PurchaseReservedInstancesOfferingResponseBody
	GetReservedInstanceIdSets() *PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets
}

type PurchaseReservedInstancesOfferingResponseBody struct {
	// The order ID.
	//
	// example:
	//
	// 23841229****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8C314443-AF0D-4766-9562-C83B7F1A3C8B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The IDs of the reserved instances.
	ReservedInstanceIdSets *PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets `json:"ReservedInstanceIdSets,omitempty" xml:"ReservedInstanceIdSets,omitempty" type:"Struct"`
}

func (s PurchaseReservedInstancesOfferingResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PurchaseReservedInstancesOfferingResponseBody) GoString() string {
	return s.String()
}

func (s *PurchaseReservedInstancesOfferingResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *PurchaseReservedInstancesOfferingResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PurchaseReservedInstancesOfferingResponseBody) GetReservedInstanceIdSets() *PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets {
	return s.ReservedInstanceIdSets
}

func (s *PurchaseReservedInstancesOfferingResponseBody) SetOrderId(v string) *PurchaseReservedInstancesOfferingResponseBody {
	s.OrderId = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingResponseBody) SetRequestId(v string) *PurchaseReservedInstancesOfferingResponseBody {
	s.RequestId = &v
	return s
}

func (s *PurchaseReservedInstancesOfferingResponseBody) SetReservedInstanceIdSets(v *PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets) *PurchaseReservedInstancesOfferingResponseBody {
	s.ReservedInstanceIdSets = v
	return s
}

func (s *PurchaseReservedInstancesOfferingResponseBody) Validate() error {
	if s.ReservedInstanceIdSets != nil {
		if err := s.ReservedInstanceIdSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets struct {
	ReservedInstanceId []*string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty" type:"Repeated"`
}

func (s PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets) String() string {
	return dara.Prettify(s)
}

func (s PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets) GoString() string {
	return s.String()
}

func (s *PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets) GetReservedInstanceId() []*string {
	return s.ReservedInstanceId
}

func (s *PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets) SetReservedInstanceId(v []*string) *PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets {
	s.ReservedInstanceId = v
	return s
}

func (s *PurchaseReservedInstancesOfferingResponseBodyReservedInstanceIdSets) Validate() error {
	return dara.Validate(s)
}
