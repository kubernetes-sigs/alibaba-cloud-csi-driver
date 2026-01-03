// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIdSets(v *RunInstancesResponseBodyInstanceIdSets) *RunInstancesResponseBody
	GetInstanceIdSets() *RunInstancesResponseBodyInstanceIdSets
	SetOrderId(v string) *RunInstancesResponseBody
	GetOrderId() *string
	SetRequestId(v string) *RunInstancesResponseBody
	GetRequestId() *string
	SetTradePrice(v float32) *RunInstancesResponseBody
	GetTradePrice() *float32
}

type RunInstancesResponseBody struct {
	// The instance IDs.
	InstanceIdSets *RunInstancesResponseBodyInstanceIdSets `json:"InstanceIdSets,omitempty" xml:"InstanceIdSets,omitempty" type:"Struct"`
	// The ID of the order. This parameter is returned only when `InstanceChargeType` is set to PrePaid.
	//
	// example:
	//
	// 123456****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The transaction price.
	//
	// example:
	//
	// 0.165
	TradePrice *float32 `json:"TradePrice,omitempty" xml:"TradePrice,omitempty"`
}

func (s RunInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *RunInstancesResponseBody) GetInstanceIdSets() *RunInstancesResponseBodyInstanceIdSets {
	return s.InstanceIdSets
}

func (s *RunInstancesResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *RunInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunInstancesResponseBody) GetTradePrice() *float32 {
	return s.TradePrice
}

func (s *RunInstancesResponseBody) SetInstanceIdSets(v *RunInstancesResponseBodyInstanceIdSets) *RunInstancesResponseBody {
	s.InstanceIdSets = v
	return s
}

func (s *RunInstancesResponseBody) SetOrderId(v string) *RunInstancesResponseBody {
	s.OrderId = &v
	return s
}

func (s *RunInstancesResponseBody) SetRequestId(v string) *RunInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunInstancesResponseBody) SetTradePrice(v float32) *RunInstancesResponseBody {
	s.TradePrice = &v
	return s
}

func (s *RunInstancesResponseBody) Validate() error {
	if s.InstanceIdSets != nil {
		if err := s.InstanceIdSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunInstancesResponseBodyInstanceIdSets struct {
	InstanceIdSet []*string `json:"InstanceIdSet,omitempty" xml:"InstanceIdSet,omitempty" type:"Repeated"`
}

func (s RunInstancesResponseBodyInstanceIdSets) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesResponseBodyInstanceIdSets) GoString() string {
	return s.String()
}

func (s *RunInstancesResponseBodyInstanceIdSets) GetInstanceIdSet() []*string {
	return s.InstanceIdSet
}

func (s *RunInstancesResponseBodyInstanceIdSets) SetInstanceIdSet(v []*string) *RunInstancesResponseBodyInstanceIdSets {
	s.InstanceIdSet = v
	return s
}

func (s *RunInstancesResponseBodyInstanceIdSets) Validate() error {
	return dara.Validate(s)
}
