// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceChargeTypeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFeeOfInstances(v *ModifyInstanceChargeTypeResponseBodyFeeOfInstances) *ModifyInstanceChargeTypeResponseBody
	GetFeeOfInstances() *ModifyInstanceChargeTypeResponseBodyFeeOfInstances
	SetOrderId(v string) *ModifyInstanceChargeTypeResponseBody
	GetOrderId() *string
	SetRequestId(v string) *ModifyInstanceChargeTypeResponseBody
	GetRequestId() *string
}

type ModifyInstanceChargeTypeResponseBody struct {
	// Details about the charges for the order.
	FeeOfInstances *ModifyInstanceChargeTypeResponseBodyFeeOfInstances `json:"FeeOfInstances,omitempty" xml:"FeeOfInstances,omitempty" type:"Struct"`
	// The order ID.
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

func (s ModifyInstanceChargeTypeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceChargeTypeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeResponseBody) GetFeeOfInstances() *ModifyInstanceChargeTypeResponseBodyFeeOfInstances {
	return s.FeeOfInstances
}

func (s *ModifyInstanceChargeTypeResponseBody) GetOrderId() *string {
	return s.OrderId
}

func (s *ModifyInstanceChargeTypeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceChargeTypeResponseBody) SetFeeOfInstances(v *ModifyInstanceChargeTypeResponseBodyFeeOfInstances) *ModifyInstanceChargeTypeResponseBody {
	s.FeeOfInstances = v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBody) SetOrderId(v string) *ModifyInstanceChargeTypeResponseBody {
	s.OrderId = &v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBody) SetRequestId(v string) *ModifyInstanceChargeTypeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBody) Validate() error {
	if s.FeeOfInstances != nil {
		if err := s.FeeOfInstances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyInstanceChargeTypeResponseBodyFeeOfInstances struct {
	FeeOfInstance []*ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance `json:"FeeOfInstance,omitempty" xml:"FeeOfInstance,omitempty" type:"Repeated"`
}

func (s ModifyInstanceChargeTypeResponseBodyFeeOfInstances) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceChargeTypeResponseBodyFeeOfInstances) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeResponseBodyFeeOfInstances) GetFeeOfInstance() []*ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance {
	return s.FeeOfInstance
}

func (s *ModifyInstanceChargeTypeResponseBodyFeeOfInstances) SetFeeOfInstance(v []*ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) *ModifyInstanceChargeTypeResponseBodyFeeOfInstances {
	s.FeeOfInstance = v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBodyFeeOfInstances) Validate() error {
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

type ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance struct {
	// The unit of currency for the bill.
	//
	// Alibaba Cloud China site (aliyun.com): CNY.
	//
	// Alibaba Cloud International site (alibabacloud.com): USD.
	//
	// example:
	//
	// CNY
	Currency *string `json:"Currency,omitempty" xml:"Currency,omitempty"`
	// The cost value.
	//
	// example:
	//
	// 0
	Fee *string `json:"Fee,omitempty" xml:"Fee,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) GoString() string {
	return s.String()
}

func (s *ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) GetCurrency() *string {
	return s.Currency
}

func (s *ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) GetFee() *string {
	return s.Fee
}

func (s *ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) SetCurrency(v string) *ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance {
	s.Currency = &v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) SetFee(v string) *ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance {
	s.Fee = &v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) SetInstanceId(v string) *ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceChargeTypeResponseBodyFeeOfInstancesFeeOfInstance) Validate() error {
	return dara.Validate(s)
}
