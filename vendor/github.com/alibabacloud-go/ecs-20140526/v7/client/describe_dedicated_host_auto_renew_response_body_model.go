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
	AutoRenewEnabled *bool   `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	AutoRenewWithEcs *string `json:"AutoRenewWithEcs,omitempty" xml:"AutoRenewWithEcs,omitempty"`
	DedicatedHostId  *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	Duration         *int32  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	PeriodUnit       *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	RenewalStatus    *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
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
