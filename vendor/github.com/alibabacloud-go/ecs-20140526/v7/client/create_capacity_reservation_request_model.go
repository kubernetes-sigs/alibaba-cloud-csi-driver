// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCapacityReservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *CreateCapacityReservationRequestPrivatePoolOptions) *CreateCapacityReservationRequest
	GetPrivatePoolOptions() *CreateCapacityReservationRequestPrivatePoolOptions
	SetClientToken(v string) *CreateCapacityReservationRequest
	GetClientToken() *string
	SetDescription(v string) *CreateCapacityReservationRequest
	GetDescription() *string
	SetEndTime(v string) *CreateCapacityReservationRequest
	GetEndTime() *string
	SetEndTimeType(v string) *CreateCapacityReservationRequest
	GetEndTimeType() *string
	SetInstanceAmount(v int32) *CreateCapacityReservationRequest
	GetInstanceAmount() *int32
	SetInstanceType(v string) *CreateCapacityReservationRequest
	GetInstanceType() *string
	SetOwnerAccount(v string) *CreateCapacityReservationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCapacityReservationRequest
	GetOwnerId() *int64
	SetPlatform(v string) *CreateCapacityReservationRequest
	GetPlatform() *string
	SetRegionId(v string) *CreateCapacityReservationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateCapacityReservationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateCapacityReservationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCapacityReservationRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *CreateCapacityReservationRequest
	GetStartTime() *string
	SetTag(v []*CreateCapacityReservationRequestTag) *CreateCapacityReservationRequest
	GetTag() []*CreateCapacityReservationRequestTag
	SetZoneId(v []*string) *CreateCapacityReservationRequest
	GetZoneId() []*string
}

type CreateCapacityReservationRequest struct {
	PrivatePoolOptions *CreateCapacityReservationRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the capacity reservation. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the capacity reservation expires. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2021-10-30T06:32:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The release mode of the capacity reservation. Valid values:
	//
	// 	- Limited: The capacity reservation is automatically released at a specified time. If you specify this parameter, you must specify the `EndTime` parameter.
	//
	// 	- Unlimited: The capacity reservation is manually released. The capacity reservation can be released anytime.
	//
	// example:
	//
	// Unlimited
	EndTimeType *string `json:"EndTimeType,omitempty" xml:"EndTimeType,omitempty"`
	// The total number of instances for which the capacity of an instance type is reserved.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2
	InstanceAmount *int32 `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	// The instance type. You can create a capacity reservation to reserve the capacity of only one instance type. You can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the instance types provided by ECS.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g6.xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operating system of the image used by the instance. This parameter corresponds to the `Platform` parameter of regional reserved instances. If the operating system of a capacity reservation matches the operating system of a regional reserved instance, you can apply the regional reserved instance to offset fees of the unused capacity of the capacity reservation. Valid values:
	//
	// 	- Windows: Windows Server operating system
	//
	// 	- Linux: Linux and UNIX-like operating system
	//
	// Default value: Linux.
	//
	// > This parameter is unavailable.
	//
	// example:
	//
	// Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The ID of the region in which to create the capacity reservation. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which you want to assign the capacity reservation.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The mode in which the capacity reservation takes effect. You can call the CreateCapacityReservation operation to create only immediate capacity reservations.
	//
	// > If you do not specify this parameter, the capacity reservation immediately takes effect.
	//
	// example:
	//
	// 2021-10-30T05:32:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The tags to add to the capacity reservation.
	Tag []*CreateCapacityReservationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the zone in which you want to create the capacity reservation. A capacity reservation can reserve resources within only one zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId []*string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" type:"Repeated"`
}

func (s CreateCapacityReservationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCapacityReservationRequest) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationRequest) GetPrivatePoolOptions() *CreateCapacityReservationRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateCapacityReservationRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateCapacityReservationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCapacityReservationRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateCapacityReservationRequest) GetEndTimeType() *string {
	return s.EndTimeType
}

func (s *CreateCapacityReservationRequest) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *CreateCapacityReservationRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateCapacityReservationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCapacityReservationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCapacityReservationRequest) GetPlatform() *string {
	return s.Platform
}

func (s *CreateCapacityReservationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCapacityReservationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCapacityReservationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCapacityReservationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCapacityReservationRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateCapacityReservationRequest) GetTag() []*CreateCapacityReservationRequestTag {
	return s.Tag
}

func (s *CreateCapacityReservationRequest) GetZoneId() []*string {
	return s.ZoneId
}

func (s *CreateCapacityReservationRequest) SetPrivatePoolOptions(v *CreateCapacityReservationRequestPrivatePoolOptions) *CreateCapacityReservationRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateCapacityReservationRequest) SetClientToken(v string) *CreateCapacityReservationRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetDescription(v string) *CreateCapacityReservationRequest {
	s.Description = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetEndTime(v string) *CreateCapacityReservationRequest {
	s.EndTime = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetEndTimeType(v string) *CreateCapacityReservationRequest {
	s.EndTimeType = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetInstanceAmount(v int32) *CreateCapacityReservationRequest {
	s.InstanceAmount = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetInstanceType(v string) *CreateCapacityReservationRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetOwnerAccount(v string) *CreateCapacityReservationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetOwnerId(v int64) *CreateCapacityReservationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetPlatform(v string) *CreateCapacityReservationRequest {
	s.Platform = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetRegionId(v string) *CreateCapacityReservationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetResourceGroupId(v string) *CreateCapacityReservationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetResourceOwnerAccount(v string) *CreateCapacityReservationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetResourceOwnerId(v int64) *CreateCapacityReservationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetStartTime(v string) *CreateCapacityReservationRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCapacityReservationRequest) SetTag(v []*CreateCapacityReservationRequestTag) *CreateCapacityReservationRequest {
	s.Tag = v
	return s
}

func (s *CreateCapacityReservationRequest) SetZoneId(v []*string) *CreateCapacityReservationRequest {
	s.ZoneId = v
	return s
}

func (s *CreateCapacityReservationRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
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

type CreateCapacityReservationRequestPrivatePoolOptions struct {
	// The type of the private pool to generate after the capacity reservation takes effect. Valid values:
	//
	// 	- Open: open private pool
	//
	// 	- Target: targeted private pool
	//
	// Default value: Open.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
	// The capacity reservation name. The name must be 2 to 128 characters in length. It must start with a letter but cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// crpTestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateCapacityReservationRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateCapacityReservationRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *CreateCapacityReservationRequestPrivatePoolOptions) GetName() *string {
	return s.Name
}

func (s *CreateCapacityReservationRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateCapacityReservationRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateCapacityReservationRequestPrivatePoolOptions) SetName(v string) *CreateCapacityReservationRequestPrivatePoolOptions {
	s.Name = &v
	return s
}

func (s *CreateCapacityReservationRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateCapacityReservationRequestTag struct {
	// The key of tag N to add to the capacity reservation. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the capacity reservation. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCapacityReservationRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCapacityReservationRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCapacityReservationRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCapacityReservationRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCapacityReservationRequestTag) SetKey(v string) *CreateCapacityReservationRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCapacityReservationRequestTag) SetValue(v string) *CreateCapacityReservationRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCapacityReservationRequestTag) Validate() error {
	return dara.Validate(s)
}
