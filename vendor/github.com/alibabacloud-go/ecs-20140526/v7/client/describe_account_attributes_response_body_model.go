// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccountAttributeItems(v *DescribeAccountAttributesResponseBodyAccountAttributeItems) *DescribeAccountAttributesResponseBody
	GetAccountAttributeItems() *DescribeAccountAttributesResponseBodyAccountAttributeItems
	SetRequestId(v string) *DescribeAccountAttributesResponseBody
	GetRequestId() *string
}

type DescribeAccountAttributesResponseBody struct {
	AccountAttributeItems *DescribeAccountAttributesResponseBodyAccountAttributeItems `json:"AccountAttributeItems,omitempty" xml:"AccountAttributeItems,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 8CE45CD5-31FB-47C2-959D-CA8144CE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccountAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccountAttributesResponseBody) GetAccountAttributeItems() *DescribeAccountAttributesResponseBodyAccountAttributeItems {
	return s.AccountAttributeItems
}

func (s *DescribeAccountAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccountAttributesResponseBody) SetAccountAttributeItems(v *DescribeAccountAttributesResponseBodyAccountAttributeItems) *DescribeAccountAttributesResponseBody {
	s.AccountAttributeItems = v
	return s
}

func (s *DescribeAccountAttributesResponseBody) SetRequestId(v string) *DescribeAccountAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccountAttributesResponseBody) Validate() error {
	if s.AccountAttributeItems != nil {
		if err := s.AccountAttributeItems.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountAttributesResponseBodyAccountAttributeItems struct {
	AccountAttributeItem []*DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem `json:"AccountAttributeItem,omitempty" xml:"AccountAttributeItem,omitempty" type:"Repeated"`
}

func (s DescribeAccountAttributesResponseBodyAccountAttributeItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAttributesResponseBodyAccountAttributeItems) GoString() string {
	return s.String()
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItems) GetAccountAttributeItem() []*DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem {
	return s.AccountAttributeItem
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItems) SetAccountAttributeItem(v []*DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem) *DescribeAccountAttributesResponseBodyAccountAttributeItems {
	s.AccountAttributeItem = v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItems) Validate() error {
	if s.AccountAttributeItem != nil {
		for _, item := range s.AccountAttributeItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem struct {
	AttributeName   *string                                                                                        `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	AttributeValues *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValues `json:"AttributeValues,omitempty" xml:"AttributeValues,omitempty" type:"Struct"`
}

func (s DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem) GoString() string {
	return s.String()
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem) GetAttributeName() *string {
	return s.AttributeName
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem) GetAttributeValues() *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValues {
	return s.AttributeValues
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem) SetAttributeName(v string) *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem {
	s.AttributeName = &v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem) SetAttributeValues(v *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValues) *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem {
	s.AttributeValues = v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItem) Validate() error {
	if s.AttributeValues != nil {
		if err := s.AttributeValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValues struct {
	ValueItem []*DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem `json:"ValueItem,omitempty" xml:"ValueItem,omitempty" type:"Repeated"`
}

func (s DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValues) GoString() string {
	return s.String()
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValues) GetValueItem() []*DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem {
	return s.ValueItem
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValues) SetValueItem(v []*DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValues {
	s.ValueItem = v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValues) Validate() error {
	if s.ValueItem != nil {
		for _, item := range s.ValueItem {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem struct {
	Count              *int32  `json:"Count,omitempty" xml:"Count,omitempty"`
	DiskCategory       *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	ExpiredTime        *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	InstanceType       *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	Value              *string `json:"Value,omitempty" xml:"Value,omitempty"`
	ZoneId             *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) GoString() string {
	return s.String()
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) GetCount() *int32 {
	return s.Count
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) GetDiskCategory() *string {
	return s.DiskCategory
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) GetValue() *string {
	return s.Value
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) SetCount(v int32) *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem {
	s.Count = &v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) SetDiskCategory(v string) *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem {
	s.DiskCategory = &v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) SetExpiredTime(v string) *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) SetInstanceChargeType(v string) *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem {
	s.InstanceChargeType = &v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) SetInstanceType(v string) *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem {
	s.InstanceType = &v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) SetValue(v string) *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem {
	s.Value = &v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) SetZoneId(v string) *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem {
	s.ZoneId = &v
	return s
}

func (s *DescribeAccountAttributesResponseBodyAccountAttributeItemsAccountAttributeItemAttributeValuesValueItem) Validate() error {
	return dara.Validate(s)
}
