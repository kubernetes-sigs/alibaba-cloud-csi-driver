// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostAutoRenewResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostRenewAttributes(v *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes) *DescribeDedicatedHostAutoRenewResponseBody
	GetDedicatedHostRenewAttributes() *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes
	SetRequestId(v string) *DescribeDedicatedHostAutoRenewResponseBody
	GetRequestId() *string
}

type DescribeDedicatedHostAutoRenewResponseBody struct {
	// The array that consists of dedicated host auto-renewal attributes.
	DedicatedHostRenewAttributes *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes `json:"DedicatedHostRenewAttributes,omitempty" xml:"DedicatedHostRenewAttributes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDedicatedHostAutoRenewResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostAutoRenewResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostAutoRenewResponseBody) GetDedicatedHostRenewAttributes() *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes {
	return s.DedicatedHostRenewAttributes
}

func (s *DescribeDedicatedHostAutoRenewResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDedicatedHostAutoRenewResponseBody) SetDedicatedHostRenewAttributes(v *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes) *DescribeDedicatedHostAutoRenewResponseBody {
	s.DedicatedHostRenewAttributes = v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponseBody) SetRequestId(v string) *DescribeDedicatedHostAutoRenewResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponseBody) Validate() error {
	if s.DedicatedHostRenewAttributes != nil {
		if err := s.DedicatedHostRenewAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes struct {
	DedicatedHostRenewAttribute []*DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute `json:"DedicatedHostRenewAttribute,omitempty" xml:"DedicatedHostRenewAttribute,omitempty" type:"Repeated"`
}

func (s DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes) GetDedicatedHostRenewAttribute() []*DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute {
	return s.DedicatedHostRenewAttribute
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes) SetDedicatedHostRenewAttribute(v []*DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes {
	s.DedicatedHostRenewAttribute = v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributes) Validate() error {
	if s.DedicatedHostRenewAttribute != nil {
		for _, item := range s.DedicatedHostRenewAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute struct {
	// Indicates whether auto-renewal is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	// Indicates whether the dedicated host is automatically renewed if a subscription ECS instance it hosts, after being automatically renewed, has a new expiration time that is later than that of the dedicated host. Valid values:
	//
	// 	- AutoRenewWithEcs: The dedicated host is automatically renewed along with the ECS instance.
	//
	// 	- StopRenewWithEcs: The dedicated host is not automatically renewed along with the ECS instance.
	//
	// example:
	//
	// StopRenewWithEcs
	AutoRenewWithEcs *string `json:"AutoRenewWithEcs,omitempty" xml:"AutoRenewWithEcs,omitempty"`
	// The ID of the dedicated host.
	//
	// example:
	//
	// dh-bp165p6xk2tlw61e****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// The auto-renewal period.
	//
	// example:
	//
	// 0
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The unit of the auto-renewal duration. Valid values:
	//
	// 	- Week
	//
	// 	- Month
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// Indicates whether the subscription dedicated host is automatically renewed. Valid values:
	//
	// 	- AutoRenewal: The dedicated host is automatically renewed.
	//
	// 	- Normal: The dedicated host is not automatically renewed, but renewal notifications are sent.
	//
	// 	- NotRenewal: The dedicated host is not automatically renewed, and no expiration notification is sent. Alibaba Cloud sends only a non-renewal notice three days before the host expires. If the renewal status of a dedicated host is NotRenewal, you can change the value to Normal and then call [RenewDedicatedHosts](https://help.aliyun.com/document_detail/93287.html) to manually renew the dedicated host, or directly change the value to AutoRenewal.
	//
	// example:
	//
	// Normal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) GetAutoRenewEnabled() *bool {
	return s.AutoRenewEnabled
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) GetAutoRenewWithEcs() *string {
	return s.AutoRenewWithEcs
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) SetAutoRenewEnabled(v bool) *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute {
	s.AutoRenewEnabled = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) SetAutoRenewWithEcs(v string) *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute {
	s.AutoRenewWithEcs = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) SetDedicatedHostId(v string) *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute {
	s.DedicatedHostId = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) SetDuration(v int32) *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) SetPeriodUnit(v string) *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) SetRenewalStatus(v string) *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeDedicatedHostAutoRenewResponseBodyDedicatedHostRenewAttributesDedicatedHostRenewAttribute) Validate() error {
	return dara.Validate(s)
}
