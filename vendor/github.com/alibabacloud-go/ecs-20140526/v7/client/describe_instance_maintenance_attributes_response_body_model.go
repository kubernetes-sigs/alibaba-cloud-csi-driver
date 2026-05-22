// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMaintenanceAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintenanceAttributes(v *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes) *DescribeInstanceMaintenanceAttributesResponseBody
	GetMaintenanceAttributes() *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes
	SetPageNumber(v int32) *DescribeInstanceMaintenanceAttributesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceMaintenanceAttributesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstanceMaintenanceAttributesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceMaintenanceAttributesResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceMaintenanceAttributesResponseBody struct {
	MaintenanceAttributes *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes `json:"MaintenanceAttributes,omitempty" xml:"MaintenanceAttributes,omitempty" type:"Struct"`
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
	// The total number of queried maintenance attributes.
	//
	// example:
	//
	// 100
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceMaintenanceAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMaintenanceAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) GetMaintenanceAttributes() *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes {
	return s.MaintenanceAttributes
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) SetMaintenanceAttributes(v *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes) *DescribeInstanceMaintenanceAttributesResponseBody {
	s.MaintenanceAttributes = v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) SetPageNumber(v int32) *DescribeInstanceMaintenanceAttributesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) SetPageSize(v int32) *DescribeInstanceMaintenanceAttributesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) SetRequestId(v string) *DescribeInstanceMaintenanceAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) SetTotalCount(v int32) *DescribeInstanceMaintenanceAttributesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBody) Validate() error {
	if s.MaintenanceAttributes != nil {
		if err := s.MaintenanceAttributes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes struct {
	MaintenanceAttribute []*DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute `json:"MaintenanceAttribute,omitempty" xml:"MaintenanceAttribute,omitempty" type:"Repeated"`
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes) GetMaintenanceAttribute() []*DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute {
	return s.MaintenanceAttribute
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes) SetMaintenanceAttribute(v []*DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes {
	s.MaintenanceAttribute = v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributes) Validate() error {
	if s.MaintenanceAttribute != nil {
		for _, item := range s.MaintenanceAttribute {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute struct {
	ActionOnMaintenance *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance `json:"ActionOnMaintenance,omitempty" xml:"ActionOnMaintenance,omitempty" type:"Struct"`
	InstanceId          *string                                                                                                        `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	MaintenanceWindows  *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindows  `json:"MaintenanceWindows,omitempty" xml:"MaintenanceWindows,omitempty" type:"Struct"`
	NotifyOnMaintenance *bool                                                                                                          `json:"NotifyOnMaintenance,omitempty" xml:"NotifyOnMaintenance,omitempty"`
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) GetActionOnMaintenance() *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance {
	return s.ActionOnMaintenance
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) GetMaintenanceWindows() *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindows {
	return s.MaintenanceWindows
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) GetNotifyOnMaintenance() *bool {
	return s.NotifyOnMaintenance
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) SetActionOnMaintenance(v *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute {
	s.ActionOnMaintenance = v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) SetInstanceId(v string) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) SetMaintenanceWindows(v *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindows) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute {
	s.MaintenanceWindows = v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) SetNotifyOnMaintenance(v bool) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute {
	s.NotifyOnMaintenance = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttribute) Validate() error {
	if s.ActionOnMaintenance != nil {
		if err := s.ActionOnMaintenance.Validate(); err != nil {
			return err
		}
	}
	if s.MaintenanceWindows != nil {
		if err := s.MaintenanceWindows.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance struct {
	DefaultValue    *string                                                                                                                       `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	SupportedValues *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenanceSupportedValues `json:"SupportedValues,omitempty" xml:"SupportedValues,omitempty" type:"Struct"`
	Value           *string                                                                                                                       `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance) GetSupportedValues() *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenanceSupportedValues {
	return s.SupportedValues
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance) GetValue() *string {
	return s.Value
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance) SetDefaultValue(v string) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance {
	s.DefaultValue = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance) SetSupportedValues(v *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenanceSupportedValues) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance {
	s.SupportedValues = v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance) SetValue(v string) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance {
	s.Value = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenance) Validate() error {
	if s.SupportedValues != nil {
		if err := s.SupportedValues.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenanceSupportedValues struct {
	SupportedValue []*string `json:"SupportedValue,omitempty" xml:"SupportedValue,omitempty" type:"Repeated"`
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenanceSupportedValues) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenanceSupportedValues) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenanceSupportedValues) GetSupportedValue() []*string {
	return s.SupportedValue
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenanceSupportedValues) SetSupportedValue(v []*string) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenanceSupportedValues {
	s.SupportedValue = v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeActionOnMaintenanceSupportedValues) Validate() error {
	return dara.Validate(s)
}

type DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindows struct {
	MaintenanceWindow []*DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow `json:"MaintenanceWindow,omitempty" xml:"MaintenanceWindow,omitempty" type:"Repeated"`
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindows) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindows) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindows) GetMaintenanceWindow() []*DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow {
	return s.MaintenanceWindow
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindows) SetMaintenanceWindow(v []*DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindows {
	s.MaintenanceWindow = v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindows) Validate() error {
	if s.MaintenanceWindow != nil {
		for _, item := range s.MaintenanceWindow {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow struct {
	EndTime   *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow) SetEndTime(v string) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow) SetStartTime(v string) *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceMaintenanceAttributesResponseBodyMaintenanceAttributesMaintenanceAttributeMaintenanceWindowsMaintenanceWindow) Validate() error {
	return dara.Validate(s)
}
