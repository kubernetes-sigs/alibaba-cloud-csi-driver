// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeReservedInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationType(v string) *DescribeReservedInstancesRequest
	GetAllocationType() *string
	SetInstanceType(v string) *DescribeReservedInstancesRequest
	GetInstanceType() *string
	SetInstanceTypeFamily(v string) *DescribeReservedInstancesRequest
	GetInstanceTypeFamily() *string
	SetLockReason(v string) *DescribeReservedInstancesRequest
	GetLockReason() *string
	SetOfferingType(v string) *DescribeReservedInstancesRequest
	GetOfferingType() *string
	SetOwnerAccount(v string) *DescribeReservedInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeReservedInstancesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeReservedInstancesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeReservedInstancesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeReservedInstancesRequest
	GetRegionId() *string
	SetReservedInstanceId(v []*string) *DescribeReservedInstancesRequest
	GetReservedInstanceId() []*string
	SetReservedInstanceName(v string) *DescribeReservedInstancesRequest
	GetReservedInstanceName() *string
	SetResourceOwnerAccount(v string) *DescribeReservedInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeReservedInstancesRequest
	GetResourceOwnerId() *int64
	SetScope(v string) *DescribeReservedInstancesRequest
	GetScope() *string
	SetStatus(v []*string) *DescribeReservedInstancesRequest
	GetStatus() []*string
	SetTag(v []*DescribeReservedInstancesRequestTag) *DescribeReservedInstancesRequest
	GetTag() []*DescribeReservedInstancesRequestTag
	SetZoneId(v string) *DescribeReservedInstancesRequest
	GetZoneId() *string
}

type DescribeReservedInstancesRequest struct {
	// The allocation type of the reserved instances. Valid values:
	//
	// 	- Normal: queries all reserved instances that belong to the current account.
	//
	// 	- Shared: queries the reserved instances that are shared between the current main account and linked accounts.
	//
	// Default value: Normal.
	//
	// example:
	//
	// Normal
	AllocationType *string `json:"AllocationType,omitempty" xml:"AllocationType,omitempty"`
	// The instance type of the reserved instance. For information about the valid values, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// >  Specify the instance type that you selected when you purchased the reserved instance. If the reserved instance is a regional reserved instance, it can be used to offset the bills of instance types that belong to the same instance family as the specified instance type, regardless of instance specifications.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The instance family of the reserved instance. For information about the valid values, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// ecs.g5
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// The reason why the reserved instance is locked. Valid values:
	//
	// 	- financial: The reserved instance is locked because the account has overdue payments or the service expires.
	//
	// 	- security: The reserved instance is locked due to security reasons.
	//
	// example:
	//
	// security
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The payment option of the reserved instance. Valid values:
	//
	// 	- No Upfront
	//
	// 	- Partial Upfront
	//
	// 	- All Upfront
	//
	// example:
	//
	// All Upfront
	OfferingType *string `json:"OfferingType,omitempty" xml:"OfferingType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the reserved instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of reserved instances. You can specify up to 100 IDs of reserved instances.
	//
	// example:
	//
	// ri-bpzhex2ulpzf53****
	ReservedInstanceId []*string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty" type:"Repeated"`
	// The name of the reserved instance.
	//
	// >  Only exact search is supported.
	//
	// example:
	//
	// testReservedInstanceName
	ReservedInstanceName *string `json:"ReservedInstanceName,omitempty" xml:"ReservedInstanceName,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The scope level of the reserved instance. Valid values:
	//
	// 	- Region: regional
	//
	// 	- Zone: zonal
	//
	// example:
	//
	// Region
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The status of the reserved instances.
	//
	// example:
	//
	// Active
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The tags of the reserved instance. You can specify up to 20 tags.
	Tag []*DescribeReservedInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID of the reserved instance. This parameter is valid and required if you set Scope to Zone. You can call the [DescribeZones](https://help.aliyun.com/document_detail/25610.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-z
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeReservedInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesRequest) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesRequest) GetAllocationType() *string {
	return s.AllocationType
}

func (s *DescribeReservedInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *DescribeReservedInstancesRequest) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeReservedInstancesRequest) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeReservedInstancesRequest) GetOfferingType() *string {
	return s.OfferingType
}

func (s *DescribeReservedInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeReservedInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeReservedInstancesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeReservedInstancesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeReservedInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeReservedInstancesRequest) GetReservedInstanceId() []*string {
	return s.ReservedInstanceId
}

func (s *DescribeReservedInstancesRequest) GetReservedInstanceName() *string {
	return s.ReservedInstanceName
}

func (s *DescribeReservedInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeReservedInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeReservedInstancesRequest) GetScope() *string {
	return s.Scope
}

func (s *DescribeReservedInstancesRequest) GetStatus() []*string {
	return s.Status
}

func (s *DescribeReservedInstancesRequest) GetTag() []*DescribeReservedInstancesRequestTag {
	return s.Tag
}

func (s *DescribeReservedInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeReservedInstancesRequest) SetAllocationType(v string) *DescribeReservedInstancesRequest {
	s.AllocationType = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetInstanceType(v string) *DescribeReservedInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetInstanceTypeFamily(v string) *DescribeReservedInstancesRequest {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetLockReason(v string) *DescribeReservedInstancesRequest {
	s.LockReason = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetOfferingType(v string) *DescribeReservedInstancesRequest {
	s.OfferingType = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetOwnerAccount(v string) *DescribeReservedInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetOwnerId(v int64) *DescribeReservedInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetPageNumber(v int32) *DescribeReservedInstancesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetPageSize(v int32) *DescribeReservedInstancesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetRegionId(v string) *DescribeReservedInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetReservedInstanceId(v []*string) *DescribeReservedInstancesRequest {
	s.ReservedInstanceId = v
	return s
}

func (s *DescribeReservedInstancesRequest) SetReservedInstanceName(v string) *DescribeReservedInstancesRequest {
	s.ReservedInstanceName = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetResourceOwnerAccount(v string) *DescribeReservedInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetResourceOwnerId(v int64) *DescribeReservedInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetScope(v string) *DescribeReservedInstancesRequest {
	s.Scope = &v
	return s
}

func (s *DescribeReservedInstancesRequest) SetStatus(v []*string) *DescribeReservedInstancesRequest {
	s.Status = v
	return s
}

func (s *DescribeReservedInstancesRequest) SetTag(v []*DescribeReservedInstancesRequestTag) *DescribeReservedInstancesRequest {
	s.Tag = v
	return s
}

func (s *DescribeReservedInstancesRequest) SetZoneId(v string) *DescribeReservedInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeReservedInstancesRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeReservedInstancesRequestTag struct {
	// The key of tag N of the reserved instance. The tag key cannot be empty and can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain http:// or https://.
	//
	// >  If you specify a single tag to query resources, up to 1,000 resources to which the tag is added are returned. If you specify multiple tags to query resources, up to 1,000 resources to which all specified tags are added are returned. To query more than 1,000 resources that have specified tags added, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the reserved instance. The tag value cannot be empty and can be up to 128 characters in length. It cannot start with `acs:` and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeReservedInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeReservedInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeReservedInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeReservedInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeReservedInstancesRequestTag) SetKey(v string) *DescribeReservedInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeReservedInstancesRequestTag) SetValue(v string) *DescribeReservedInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeReservedInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
