// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeElasticityAssuranceAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetElasticityAssuranceRenewAttributes(v *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes) *DescribeElasticityAssuranceAutoRenewAttributeResponseBody
	GetElasticityAssuranceRenewAttributes() *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes
	SetRequestId(v string) *DescribeElasticityAssuranceAutoRenewAttributeResponseBody
	GetRequestId() *string
}

type DescribeElasticityAssuranceAutoRenewAttributeResponseBody struct {
	// The auto-renewal attribute of the elasticity assurances.
	ElasticityAssuranceRenewAttributes *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes `json:"ElasticityAssuranceRenewAttributes,omitempty" xml:"ElasticityAssuranceRenewAttributes,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeElasticityAssuranceAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBody) GetElasticityAssuranceRenewAttributes() *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes {
	return s.ElasticityAssuranceRenewAttributes
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBody) SetElasticityAssuranceRenewAttributes(v *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes) *DescribeElasticityAssuranceAutoRenewAttributeResponseBody {
	s.ElasticityAssuranceRenewAttributes = v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeElasticityAssuranceAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBody) Validate() error {
	if s.ElasticityAssuranceRenewAttributes != nil {
		if err := s.ElasticityAssuranceRenewAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes struct {
	ElasticityAssuranceRenewAttribute []*DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute `json:"ElasticityAssuranceRenewAttribute,omitempty" xml:"ElasticityAssuranceRenewAttribute,omitempty" type:"Repeated"`
}

func (s DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes) GetElasticityAssuranceRenewAttribute() []*DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute {
	return s.ElasticityAssuranceRenewAttribute
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes) SetElasticityAssuranceRenewAttribute(v []*DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes {
	s.ElasticityAssuranceRenewAttribute = v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributes) Validate() error {
	if s.ElasticityAssuranceRenewAttribute != nil {
		for _, item := range s.ElasticityAssuranceRenewAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute struct {
	// The auto-renewal period. Valid values: Valid values: 1, 2, 3, 6, 12, 24, and 36.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the auto-renewal period. Valid values:
	//
	// 	- Month (default)
	//
	// 	- Year
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The ID of the elasticity assurance.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	PrivatePoolOptionsId *string `json:"PrivatePoolOptionsId,omitempty" xml:"PrivatePoolOptionsId,omitempty"`
	// Indicates whether auto-renewal is enabled for the elasticity assurance. Valid values:
	//
	// 	- AutoRenewal: Auto-renewal is enabled for the elasticity assurance.
	//
	// 	- Normal: Auto-renewal is disabled for the elasticity assurance.
	//
	// 	- NotRenewal: The elasticity assurance is not renewed.
	//
	// example:
	//
	// Normal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) GetPrivatePoolOptionsId() *string {
	return s.PrivatePoolOptionsId
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) SetPeriod(v int32) *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute {
	s.Period = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) SetPeriodUnit(v string) *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) SetPrivatePoolOptionsId(v string) *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute {
	s.PrivatePoolOptionsId = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) SetRenewalStatus(v string) *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeElasticityAssuranceAutoRenewAttributeResponseBodyElasticityAssuranceRenewAttributesElasticityAssuranceRenewAttribute) Validate() error {
	return dara.Validate(s)
}
