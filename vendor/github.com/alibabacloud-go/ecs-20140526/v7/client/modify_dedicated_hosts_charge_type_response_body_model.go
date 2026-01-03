// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedHostsChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeeOfInstances(v *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances) *ModifyDedicatedHostsChargeTypeResponseBody
	GetFeeOfInstances() *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances
	SetOrderId(v string) *ModifyDedicatedHostsChargeTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyDedicatedHostsChargeTypeResponseBody
	GetRequestId() *string
}

type ModifyDedicatedHostsChargeTypeResponseBody struct {
	// Details about the charges for the order.
	FeeOfInstances *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances `json:"FeeOfInstances,omitempty" xml:"FeeOfInstances,omitempty" type:"Struct"`
	// The ID of the order. This is returned only when the payment method is changed to subscription.
	//
	// example:
	//
	// 20413515388****
	OrderId *string `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B61C08E5-403A-46A2-96C1-F7B1216DB10C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDedicatedHostsChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostsChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostsChargeTypeResponseBody) GetFeeOfInstances() *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances {
	return s.FeeOfInstances
}

func (s *ModifyDedicatedHostsChargeTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyDedicatedHostsChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDedicatedHostsChargeTypeResponseBody) SetFeeOfInstances(v *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances) *ModifyDedicatedHostsChargeTypeResponseBody {
	s.FeeOfInstances = v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeResponseBody) SetOrderId(v string) *ModifyDedicatedHostsChargeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeResponseBody) SetRequestId(v string) *ModifyDedicatedHostsChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeResponseBody) Validate() error {
	if s.FeeOfInstances != nil {
		if err := s.FeeOfInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances struct {
	FeeOfInstance []*ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance `json:"FeeOfInstance,omitempty" xml:"FeeOfInstance,omitempty" type:"Repeated"`
}

func (s ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances) GetFeeOfInstance() []*ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance {
	return s.FeeOfInstance
}

func (s *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances) SetFeeOfInstance(v []*ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances {
	s.FeeOfInstance = v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstances) Validate() error {
	if s.FeeOfInstance != nil {
		for _, item := range s.FeeOfInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance struct {
	// The unit of currency for the bill.
	//
	// Alibaba Cloud China site (aliyun.com): CNY
	//
	// Alibaba Cloud International site (alibabacloud.com): USD
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The charged amount.
	//
	// example:
	//
	// 0
	Fee *string `json:"Fee,omitempty" xml:"Fee,omitempty"`
	// The IDs of the dedicated hosts.
	//
	// example:
	//
	// dh-bp181e5064b5sotrr****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) GetCurrency() *string {
	return s.Currency
}

func (s *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) GetFee() *string {
	return s.Fee
}

func (s *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) SetCurrency(v string) *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance {
	s.Currency = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) SetFee(v string) *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance {
	s.Fee = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) SetInstanceId(v string) *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance {
	s.InstanceId = &v
	return s
}

func (s *ModifyDedicatedHostsChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) Validate() error {
	return dara.Validate(s)
}
