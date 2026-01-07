// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReservedInstanceAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeReservedInstanceAutoRenewAttributeResponseBody
	GetRequestId() *string
	SetReservedInstanceRenewAttributes(v *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes) *DescribeReservedInstanceAutoRenewAttributeResponseBody
	GetReservedInstanceRenewAttributes() *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes
}

type DescribeReservedInstanceAutoRenewAttributeResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details about the auto-renewal settings of the reserved instances.
	ReservedInstanceRenewAttributes *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes `json:"ReservedInstanceRenewAttributes,omitempty" xml:"ReservedInstanceRenewAttributes,omitempty" type:"Struct"`
}

func (s DescribeReservedInstanceAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstanceAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBody) GetReservedInstanceRenewAttributes() *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes {
	return s.ReservedInstanceRenewAttributes
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeReservedInstanceAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBody) SetReservedInstanceRenewAttributes(v *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes) *DescribeReservedInstanceAutoRenewAttributeResponseBody {
	s.ReservedInstanceRenewAttributes = v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBody) Validate() error {
	if s.ReservedInstanceRenewAttributes != nil {
		if err := s.ReservedInstanceRenewAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes struct {
	ReservedInstanceRenewAttribute []*DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute `json:"ReservedInstanceRenewAttribute,omitempty" xml:"ReservedInstanceRenewAttribute,omitempty" type:"Repeated"`
}

func (s DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes) GetReservedInstanceRenewAttribute() []*DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute {
	return s.ReservedInstanceRenewAttribute
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes) SetReservedInstanceRenewAttribute(v []*DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes {
	s.ReservedInstanceRenewAttribute = v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributes) Validate() error {
	if s.ReservedInstanceRenewAttribute != nil {
		for _, item := range s.ReservedInstanceRenewAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute struct {
	// The auto-renewal duration.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The unit of the auto-renewal duration.
	//
	// Valid values: Year and Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The auto-renewal status of the reserved instance. Valid values:
	//
	// 	- AutoRenewal: automatically renews the reserved instance.
	//
	// 	- Normal: manually renews the reserved instances.
	//
	// example:
	//
	// AutoRenewal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
	// The ID of the reserved instance.
	//
	// example:
	//
	// ecsri-ajdfaj****
	ReservedInstanceId *string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty"`
}

func (s DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) GetReservedInstanceId() *string {
	return s.ReservedInstanceId
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) SetDuration(v int32) *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) SetPeriodUnit(v string) *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) SetRenewalStatus(v string) *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) SetReservedInstanceId(v string) *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute {
	s.ReservedInstanceId = &v
	return s
}

func (s *DescribeReservedInstanceAutoRenewAttributeResponseBodyReservedInstanceRenewAttributesReservedInstanceRenewAttribute) Validate() error {
	return dara.Validate(s)
}
