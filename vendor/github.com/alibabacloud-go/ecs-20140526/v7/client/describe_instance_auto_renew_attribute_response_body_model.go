// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAutoRenewAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRenewAttributes(v *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) *DescribeInstanceAutoRenewAttributeResponseBody
	GetInstanceRenewAttributes() *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes
	SetPageNumber(v int32) *DescribeInstanceAutoRenewAttributeResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceAutoRenewAttributeResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstanceAutoRenewAttributeResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceAutoRenewAttributeResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceAutoRenewAttributeResponseBody struct {
	// The renewal attributes of instances.
	InstanceRenewAttributes *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes `json:"InstanceRenewAttributes,omitempty" xml:"InstanceRenewAttributes,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of queried instances.
	//
	// example:
	//
	// 6
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceAutoRenewAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) GetInstanceRenewAttributes() *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes {
	return s.InstanceRenewAttributes
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetInstanceRenewAttributes(v *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.InstanceRenewAttributes = v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetPageNumber(v int32) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetPageSize(v int32) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetRequestId(v string) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) SetTotalCount(v int32) *DescribeInstanceAutoRenewAttributeResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBody) Validate() error {
	if s.InstanceRenewAttributes != nil {
		if err := s.InstanceRenewAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes struct {
	InstanceRenewAttribute []*DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute `json:"InstanceRenewAttribute,omitempty" xml:"InstanceRenewAttribute,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) GetInstanceRenewAttribute() []*DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	return s.InstanceRenewAttribute
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) SetInstanceRenewAttribute(v []*DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes {
	s.InstanceRenewAttribute = v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributes) Validate() error {
	if s.InstanceRenewAttribute != nil {
		for _, item := range s.InstanceRenewAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute struct {
	// Indicates whether auto-renewal is enabled.
	//
	// example:
	//
	// false
	AutoRenewEnabled *bool `json:"AutoRenewEnabled,omitempty" xml:"AutoRenewEnabled,omitempty"`
	// The auto-renewal duration.
	//
	// example:
	//
	// 1
	Duration *int32 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp18x3z4hc7bixhx****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The unit of the auto-renewal duration.
	//
	// example:
	//
	// week
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The auto-renewal state of the instance. Valid values:
	//
	// 	- AutoRenewal: Auto-renewal is enabled for the instance.
	//
	// 	- Normal: Auto-renewal is disabled for the instance.
	//
	// 	- NotRenewal: The instance is not to be renewed. The system sends no more expiration reminders, but sends only a non-renewal reminder three days before the expiration date. For an instance that is not to be renewed, you can call the [ModifyInstanceAutoRenewAttribute](https://help.aliyun.com/document_detail/52843.html) operation to change its auto-renewal status to `Normal`. Then, you can manually renew the instance or enable auto-renewal for the instance.
	//
	// example:
	//
	// Normal
	RenewalStatus *string `json:"RenewalStatus,omitempty" xml:"RenewalStatus,omitempty"`
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GetAutoRenewEnabled() *bool {
	return s.AutoRenewEnabled
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GetDuration() *int32 {
	return s.Duration
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) GetRenewalStatus() *string {
	return s.RenewalStatus
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetAutoRenewEnabled(v bool) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.AutoRenewEnabled = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetDuration(v int32) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.Duration = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetInstanceId(v string) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetPeriodUnit(v string) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.PeriodUnit = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) SetRenewalStatus(v string) *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute {
	s.RenewalStatus = &v
	return s
}

func (s *DescribeInstanceAutoRenewAttributeResponseBodyInstanceRenewAttributesInstanceRenewAttribute) Validate() error {
	return dara.Validate(s)
}
