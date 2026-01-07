// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCapacityReservationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *ModifyCapacityReservationRequestPrivatePoolOptions) *ModifyCapacityReservationRequest
	GetPrivatePoolOptions() *ModifyCapacityReservationRequestPrivatePoolOptions
	SetDescription(v string) *ModifyCapacityReservationRequest
	GetDescription() *string
	SetEndTime(v string) *ModifyCapacityReservationRequest
	GetEndTime() *string
	SetEndTimeType(v string) *ModifyCapacityReservationRequest
	GetEndTimeType() *string
	SetInstanceAmount(v int32) *ModifyCapacityReservationRequest
	GetInstanceAmount() *int32
	SetOwnerAccount(v string) *ModifyCapacityReservationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCapacityReservationRequest
	GetOwnerId() *int64
	SetPlatform(v string) *ModifyCapacityReservationRequest
	GetPlatform() *string
	SetRegionId(v string) *ModifyCapacityReservationRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCapacityReservationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCapacityReservationRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *ModifyCapacityReservationRequest
	GetStartTime() *string
}

type ModifyCapacityReservationRequest struct {
	PrivatePoolOptions *ModifyCapacityReservationRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The description of the capacity reservation. The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The expiration time of the capacity reservation. This parameter takes effect only when `EndTimeType` is set to Limited. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2021-10-30T06:32:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The release mode of the capacity reservation. Valid values:
	//
	// 	- Limited: The capacity reservation is automatically released at the specified point in time. If you configure this parameter, you must also configure `EndTime`.
	//
	// 	- Unlimited: The capacity reservation must be manually released. You can release it anytime.
	//
	// example:
	//
	// Unlimited
	EndTimeType *string `json:"EndTimeType,omitempty" xml:"EndTimeType,omitempty"`
	// The total number of instances for which capacity is reserved. Valid values: the number of used instances to 1000.
	//
	// > When you increase the number of instances, the increase may fail due to insufficient resources.
	//
	// example:
	//
	// 100
	InstanceAmount *int32  `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The operating system of the image used by the instance. Valid values:
	//
	// 	- Windows
	//
	// 	- Linux
	//
	// > This parameter is unavailable.
	//
	// example:
	//
	// Linux
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// The region ID of the capacity reservation. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The mode in which the capacity reservation takes effect. Only immediate capacity reservations are supported. You do not need to specify a value for this parameter.
	//
	// > If you do not specify a value for this parameter, the capacity reservation immediately takes effect.
	//
	// example:
	//
	// Now
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyCapacityReservationRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCapacityReservationRequest) GoString() string {
	return s.String()
}

func (s *ModifyCapacityReservationRequest) GetPrivatePoolOptions() *ModifyCapacityReservationRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *ModifyCapacityReservationRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCapacityReservationRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyCapacityReservationRequest) GetEndTimeType() *string {
	return s.EndTimeType
}

func (s *ModifyCapacityReservationRequest) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *ModifyCapacityReservationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCapacityReservationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCapacityReservationRequest) GetPlatform() *string {
	return s.Platform
}

func (s *ModifyCapacityReservationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCapacityReservationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCapacityReservationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCapacityReservationRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyCapacityReservationRequest) SetPrivatePoolOptions(v *ModifyCapacityReservationRequestPrivatePoolOptions) *ModifyCapacityReservationRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyCapacityReservationRequest) SetDescription(v string) *ModifyCapacityReservationRequest {
	s.Description = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetEndTime(v string) *ModifyCapacityReservationRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetEndTimeType(v string) *ModifyCapacityReservationRequest {
	s.EndTimeType = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetInstanceAmount(v int32) *ModifyCapacityReservationRequest {
	s.InstanceAmount = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetOwnerAccount(v string) *ModifyCapacityReservationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetOwnerId(v int64) *ModifyCapacityReservationRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetPlatform(v string) *ModifyCapacityReservationRequest {
	s.Platform = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetRegionId(v string) *ModifyCapacityReservationRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetResourceOwnerAccount(v string) *ModifyCapacityReservationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetResourceOwnerId(v int64) *ModifyCapacityReservationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCapacityReservationRequest) SetStartTime(v string) *ModifyCapacityReservationRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyCapacityReservationRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCapacityReservationRequestPrivatePoolOptions struct {
	// The capacity reservation ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// crp-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The name of the capacity reservation. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// eapTestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ModifyCapacityReservationRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyCapacityReservationRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyCapacityReservationRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *ModifyCapacityReservationRequestPrivatePoolOptions) GetName() *string {
	return s.Name
}

func (s *ModifyCapacityReservationRequestPrivatePoolOptions) SetId(v string) *ModifyCapacityReservationRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ModifyCapacityReservationRequestPrivatePoolOptions) SetName(v string) *ModifyCapacityReservationRequestPrivatePoolOptions {
	s.Name = &v
	return s
}

func (s *ModifyCapacityReservationRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
