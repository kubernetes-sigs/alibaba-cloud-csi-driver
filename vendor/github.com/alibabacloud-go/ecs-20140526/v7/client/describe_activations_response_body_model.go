// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActivationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetActivationList(v []*DescribeActivationsResponseBodyActivationList) *DescribeActivationsResponseBody
	GetActivationList() []*DescribeActivationsResponseBodyActivationList
	SetNextToken(v string) *DescribeActivationsResponseBody
	GetNextToken() *string
	SetPageNumber(v int64) *DescribeActivationsResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeActivationsResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeActivationsResponseBody
	GetRequestId() *string
	SetTotalCount(v int64) *DescribeActivationsResponseBody
	GetTotalCount() *int64
}

type DescribeActivationsResponseBody struct {
	// The activation codes and their usage information.
	ActivationList []*DescribeActivationsResponseBodyActivationList `json:"ActivationList,omitempty" xml:"ActivationList,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F74625134
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeActivationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeActivationsResponseBody) GetActivationList() []*DescribeActivationsResponseBodyActivationList {
	return s.ActivationList
}

func (s *DescribeActivationsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeActivationsResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeActivationsResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeActivationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeActivationsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *DescribeActivationsResponseBody) SetActivationList(v []*DescribeActivationsResponseBodyActivationList) *DescribeActivationsResponseBody {
	s.ActivationList = v
	return s
}

func (s *DescribeActivationsResponseBody) SetNextToken(v string) *DescribeActivationsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeActivationsResponseBody) SetPageNumber(v int64) *DescribeActivationsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeActivationsResponseBody) SetPageSize(v int64) *DescribeActivationsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeActivationsResponseBody) SetRequestId(v string) *DescribeActivationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeActivationsResponseBody) SetTotalCount(v int64) *DescribeActivationsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeActivationsResponseBody) Validate() error {
	if s.ActivationList != nil {
		for _, item := range s.ActivationList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeActivationsResponseBodyActivationList struct {
	// The ID of the activation code.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F1234****
	ActivationId *string `json:"ActivationId,omitempty" xml:"ActivationId,omitempty"`
	// The time when the activation code was created.
	//
	// example:
	//
	// 2021-01-20T06:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The number of instances that were deregistered.
	//
	// example:
	//
	// 1
	DeregisteredCount *int32 `json:"DeregisteredCount,omitempty" xml:"DeregisteredCount,omitempty"`
	// The description of the activation code.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether the activation code is disabled.
	//
	// example:
	//
	// false
	Disabled *bool `json:"Disabled,omitempty" xml:"Disabled,omitempty"`
	// The maximum number of times that the activation code can be used to register managed instances.
	//
	// example:
	//
	// 1
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The default instance name prefix.
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP addresses of hosts that are allowed to use the activation code.
	//
	// example:
	//
	// 0.0.0.0/0
	IpAddressRange *string `json:"IpAddressRange,omitempty" xml:"IpAddressRange,omitempty"`
	// The number of instances that were registered.
	//
	// example:
	//
	// 1
	RegisteredCount *int32 `json:"RegisteredCount,omitempty" xml:"RegisteredCount,omitempty"`
	// The ID of the resource group to which the activation code belongs.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the activation code.
	Tags []*DescribeActivationsResponseBodyActivationListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The validity period of the activation code. Unit: hours.
	//
	// example:
	//
	// 4
	TimeToLiveInHours *int64 `json:"TimeToLiveInHours,omitempty" xml:"TimeToLiveInHours,omitempty"`
}

func (s DescribeActivationsResponseBodyActivationList) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationsResponseBodyActivationList) GoString() string {
	return s.String()
}

func (s *DescribeActivationsResponseBodyActivationList) GetActivationId() *string {
	return s.ActivationId
}

func (s *DescribeActivationsResponseBodyActivationList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeActivationsResponseBodyActivationList) GetDeregisteredCount() *int32 {
	return s.DeregisteredCount
}

func (s *DescribeActivationsResponseBodyActivationList) GetDescription() *string {
	return s.Description
}

func (s *DescribeActivationsResponseBodyActivationList) GetDisabled() *bool {
	return s.Disabled
}

func (s *DescribeActivationsResponseBodyActivationList) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *DescribeActivationsResponseBodyActivationList) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeActivationsResponseBodyActivationList) GetIpAddressRange() *string {
	return s.IpAddressRange
}

func (s *DescribeActivationsResponseBodyActivationList) GetRegisteredCount() *int32 {
	return s.RegisteredCount
}

func (s *DescribeActivationsResponseBodyActivationList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeActivationsResponseBodyActivationList) GetTags() []*DescribeActivationsResponseBodyActivationListTags {
	return s.Tags
}

func (s *DescribeActivationsResponseBodyActivationList) GetTimeToLiveInHours() *int64 {
	return s.TimeToLiveInHours
}

func (s *DescribeActivationsResponseBodyActivationList) SetActivationId(v string) *DescribeActivationsResponseBodyActivationList {
	s.ActivationId = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetCreationTime(v string) *DescribeActivationsResponseBodyActivationList {
	s.CreationTime = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetDeregisteredCount(v int32) *DescribeActivationsResponseBodyActivationList {
	s.DeregisteredCount = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetDescription(v string) *DescribeActivationsResponseBodyActivationList {
	s.Description = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetDisabled(v bool) *DescribeActivationsResponseBodyActivationList {
	s.Disabled = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetInstanceCount(v int32) *DescribeActivationsResponseBodyActivationList {
	s.InstanceCount = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetInstanceName(v string) *DescribeActivationsResponseBodyActivationList {
	s.InstanceName = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetIpAddressRange(v string) *DescribeActivationsResponseBodyActivationList {
	s.IpAddressRange = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetRegisteredCount(v int32) *DescribeActivationsResponseBodyActivationList {
	s.RegisteredCount = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetResourceGroupId(v string) *DescribeActivationsResponseBodyActivationList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetTags(v []*DescribeActivationsResponseBodyActivationListTags) *DescribeActivationsResponseBodyActivationList {
	s.Tags = v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) SetTimeToLiveInHours(v int64) *DescribeActivationsResponseBodyActivationList {
	s.TimeToLiveInHours = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationList) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeActivationsResponseBodyActivationListTags struct {
	// The tag key of the activation code.
	//
	// example:
	//
	// owner
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the activation code.
	//
	// example:
	//
	// zhangsan
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeActivationsResponseBodyActivationListTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationsResponseBodyActivationListTags) GoString() string {
	return s.String()
}

func (s *DescribeActivationsResponseBodyActivationListTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeActivationsResponseBodyActivationListTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeActivationsResponseBodyActivationListTags) SetTagKey(v string) *DescribeActivationsResponseBodyActivationListTags {
	s.TagKey = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationListTags) SetTagValue(v string) *DescribeActivationsResponseBodyActivationListTags {
	s.TagValue = &v
	return s
}

func (s *DescribeActivationsResponseBodyActivationListTags) Validate() error {
	return dara.Validate(s)
}
