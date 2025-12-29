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
	// Details about account privileges in the specified region.
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
	// The type of the resource quota in the specified region. Valid values:
	//
	// 	- instance-network-type: the available network types.
	//
	// 	- max-security-groups: the maximum number of security groups.
	//
	// 	- max-elastic-network-interfaces: the maximum number of ENIs.
	//
	// 	- max-postpaid-instance-vcpu-count: the maximum number of vCPUs for pay-as-you-go instances.
	//
	// 	- max-spot-instance-vcpu-count: the maximum number of vCPUs for spot instances.
	//
	// 	- used-postpaid-instance-vcpu-count: the number of vCPUs that were allocated to pay-as-you-go instances.
	//
	// 	- used-spot-instance-vcpu-count: the number of vCPUs that were allocated to spot instances.
	//
	// 	- max-postpaid-yundisk-capacity: the maximum capacity of pay-as-you-go data disks. (The value is deprecated.)
	//
	// 	- used-postpaid-yundisk-capacity: the capacity of pay-as-you-go data disks that were created. (The value is deprecated.)
	//
	// 	- max-dedicated-hosts: the maximum number of dedicated hosts.
	//
	// 	- supported-postpaid-instance-types: the instance types of pay-as-you-go I/O optimized instances.
	//
	// 	- max-axt-command-count: the maximum number of Cloud Assistant commands.
	//
	// 	- max-axt-invocation-daily: the maximum number of Cloud Assistant command executions per day.
	//
	// 	- real-name-authentication: whether the account completed the real-name verification.
	//
	// 	- max-cloud-assistant-activation-count: the maximum number of activation codes that can be created to use to register managed instances.
	//
	// example:
	//
	// max-security-groups
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// The values of resource quotas.
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
	// The number of privilege attributes in the account.
	//
	// example:
	//
	// 3
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The data disk category. Valid values:
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: enhanced SSD (ESSD)
	//
	// example:
	//
	// cloud_ssd
	DiskCategory *string `json:"DiskCategory,omitempty" xml:"DiskCategory,omitempty"`
	// The expiration time of a privilege. This parameter is returned only when the account privilege has an expiration time. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-01-01T12:30:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The billing method of the instance.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The type of the instance.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The value of the resource quota in the specified region. Valid values:
	//
	// The values returned for the resource quotas to which the following AttributeName values correspond are 0 or positive integers:
	//
	// 	- max-security-groups
	//
	// 	- max-elastic-network-interfaces
	//
	// 	- max-postpaid-instance-vcpu-count
	//
	// 	- max-spot-instance-vcpu-count
	//
	// 	- used-postpaid-instance-vcpu-count
	//
	// 	- used-spot-instance-vcpu-count
	//
	// 	- max-postpaid-yundisk-capacity (the value is deprecated)
	//
	// 	- used-postpaid-yundisk-capacity (the value is deprecated)
	//
	// 	- max-dedicated-hosts
	//
	// 	- max-axt-command-count
	//
	// 	- max-axt-invocation-daily
	//
	// 	- max-cloud-assistant-activation-count
	//
	// When AttributeName is set to supported-postpay-instance-types, instance types are returned. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// When AttributeName is set to real-name-authentications, one of the following values is returned:
	//
	// 	- yes
	//
	// 	- none
	//
	// 	- unnecessary
	//
	// When AttributeName is set to instance-network-type, one of the following values is returned:
	//
	// 	- vpc
	//
	// 	- classic
	//
	// example:
	//
	// 800
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	// The ID of the zone in which the resource resides.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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
