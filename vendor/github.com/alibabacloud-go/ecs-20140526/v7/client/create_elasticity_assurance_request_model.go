// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateElasticityAssuranceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *CreateElasticityAssuranceRequestPrivatePoolOptions) *CreateElasticityAssuranceRequest
	GetPrivatePoolOptions() *CreateElasticityAssuranceRequestPrivatePoolOptions
	SetAssuranceTimes(v string) *CreateElasticityAssuranceRequest
	GetAssuranceTimes() *string
	SetAutoRenew(v bool) *CreateElasticityAssuranceRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateElasticityAssuranceRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *CreateElasticityAssuranceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateElasticityAssuranceRequest
	GetDescription() *string
	SetInstanceAmount(v int32) *CreateElasticityAssuranceRequest
	GetInstanceAmount() *int32
	SetInstanceCpuCoreCount(v int32) *CreateElasticityAssuranceRequest
	GetInstanceCpuCoreCount() *int32
	SetInstanceType(v []*string) *CreateElasticityAssuranceRequest
	GetInstanceType() []*string
	SetOwnerAccount(v string) *CreateElasticityAssuranceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateElasticityAssuranceRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *CreateElasticityAssuranceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateElasticityAssuranceRequest
	GetPeriodUnit() *string
	SetRecurrenceRules(v []*CreateElasticityAssuranceRequestRecurrenceRules) *CreateElasticityAssuranceRequest
	GetRecurrenceRules() []*CreateElasticityAssuranceRequestRecurrenceRules
	SetRegionId(v string) *CreateElasticityAssuranceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateElasticityAssuranceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateElasticityAssuranceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateElasticityAssuranceRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *CreateElasticityAssuranceRequest
	GetStartTime() *string
	SetTag(v []*CreateElasticityAssuranceRequestTag) *CreateElasticityAssuranceRequest
	GetTag() []*CreateElasticityAssuranceRequestTag
	SetZoneId(v []*string) *CreateElasticityAssuranceRequest
	GetZoneId() []*string
}

type CreateElasticityAssuranceRequest struct {
	PrivatePoolOptions *CreateElasticityAssuranceRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The total number of times that the elasticity assurance can be used. Set the value to Unlimited. This value specifies that the elasticity assurance can be used for an unlimited number of times within its validity period.
	//
	// Default value: Unlimited.
	//
	// example:
	//
	// Unlimited
	AssuranceTimes *string `json:"AssuranceTimes,omitempty" xml:"AssuranceTimes,omitempty"`
	// Specifies whether to enable auto-renewal for the elasticity assurance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period. Unit: month. Valid values: 1, 2, 3, 6, 12, 24, and 36.
	//
	// 	- Default value when `PeriodUnit` is set to Month: 1.
	//
	// 	- Default value when `PeriodUnit` is set to Year: 12.
	//
	// >  If you set `AutoRenew` to `true`, you must specify this parameter.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that you want to use to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The `token` can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 0c593ea1-3bea-11e9-b96b-88e9fe637760
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the elasticity assurance. The description must be 2 to 256 characters in length. It cannot start with `http://` or `https://`.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The total number of instances of an instance type for which you want to reserve capacity.
	//
	// Valid values: 1 to 1000.
	//
	// >  You must specify this parameter.
	//
	// example:
	//
	// 2
	InstanceAmount *int32 `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	// > This parameter is no longer used.
	//
	// example:
	//
	// null
	InstanceCpuCoreCount *int32 `json:"InstanceCpuCoreCount,omitempty" xml:"InstanceCpuCoreCount,omitempty"`
	// The instance type. An elasticity assurance can be created to reserve the capacity of a single instance type.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.c6.xlarge
	InstanceType []*string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The validity period of the elasticity assurance. The unit of the validity period is determined by the value of `PeriodUnit`. Specifies whether to check the image used by the instance supports hot migration. Valid values:
	//
	// 	- When the value of `PeriodUnit` is `Month`, the valid values are 1, 2, 3, 4, 5, 6, 7, 8, and 9.
	//
	// 	- When the value of `PeriodUnit` is `Year`, the valid values are 1, 2, 3, 4, and 5.
	//
	// 	- When the value of `PeriodUnit` is `Day`, the valid values are 1 to 365.
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the validity period of the elasticity assurance. Valid values:
	//
	// 	- Month
	//
	// 	- Year
	//
	// 	- Day
	//
	//     **
	//
	//     **Note*	- If you set `PeriodUnit` to `Day`, you must specify RecurrenceRules to create a time-segmented elasticity assurance.
	//
	// Default value: Year.
	//
	// example:
	//
	// Year
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The assurance schedules based on which the capacity reservation takes effect.
	//
	// >  Time-segmented elasticity assurances are available only in specific regions and to specific users. To use time-segmented elasticity assurances, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl).
	RecurrenceRules []*CreateElasticityAssuranceRequestRecurrenceRules `json:"RecurrenceRules,omitempty" xml:"RecurrenceRules,omitempty" type:"Repeated"`
	// The ID of the region in which to create the elasticity assurance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the elasticity assurance.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The time when the elasticity assurance takes effect. The default value is the time when the CreateElasticityAssurance operation is called to create the elasticity assurance. Specify the time in the ISO 8601 standard in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC. For more information, see [ISO 8601](https://help.aliyun.com/document_detail/25696.html).
	//
	// example:
	//
	// 2020-10-30T06:32:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The tags to add to the elasticity assurance.
	Tag []*CreateElasticityAssuranceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the zone in which to create the elasticity assurance. An elasticity assurance can be used to reserve resources within a single zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId []*string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty" type:"Repeated"`
}

func (s CreateElasticityAssuranceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticityAssuranceRequest) GoString() string {
	return s.String()
}

func (s *CreateElasticityAssuranceRequest) GetPrivatePoolOptions() *CreateElasticityAssuranceRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateElasticityAssuranceRequest) GetAssuranceTimes() *string {
	return s.AssuranceTimes
}

func (s *CreateElasticityAssuranceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateElasticityAssuranceRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateElasticityAssuranceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateElasticityAssuranceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateElasticityAssuranceRequest) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *CreateElasticityAssuranceRequest) GetInstanceCpuCoreCount() *int32 {
	return s.InstanceCpuCoreCount
}

func (s *CreateElasticityAssuranceRequest) GetInstanceType() []*string {
	return s.InstanceType
}

func (s *CreateElasticityAssuranceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateElasticityAssuranceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateElasticityAssuranceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateElasticityAssuranceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateElasticityAssuranceRequest) GetRecurrenceRules() []*CreateElasticityAssuranceRequestRecurrenceRules {
	return s.RecurrenceRules
}

func (s *CreateElasticityAssuranceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateElasticityAssuranceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateElasticityAssuranceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateElasticityAssuranceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateElasticityAssuranceRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *CreateElasticityAssuranceRequest) GetTag() []*CreateElasticityAssuranceRequestTag {
	return s.Tag
}

func (s *CreateElasticityAssuranceRequest) GetZoneId() []*string {
	return s.ZoneId
}

func (s *CreateElasticityAssuranceRequest) SetPrivatePoolOptions(v *CreateElasticityAssuranceRequestPrivatePoolOptions) *CreateElasticityAssuranceRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetAssuranceTimes(v string) *CreateElasticityAssuranceRequest {
	s.AssuranceTimes = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetAutoRenew(v bool) *CreateElasticityAssuranceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetAutoRenewPeriod(v int32) *CreateElasticityAssuranceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetClientToken(v string) *CreateElasticityAssuranceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetDescription(v string) *CreateElasticityAssuranceRequest {
	s.Description = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetInstanceAmount(v int32) *CreateElasticityAssuranceRequest {
	s.InstanceAmount = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetInstanceCpuCoreCount(v int32) *CreateElasticityAssuranceRequest {
	s.InstanceCpuCoreCount = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetInstanceType(v []*string) *CreateElasticityAssuranceRequest {
	s.InstanceType = v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetOwnerAccount(v string) *CreateElasticityAssuranceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetOwnerId(v int64) *CreateElasticityAssuranceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetPeriod(v int32) *CreateElasticityAssuranceRequest {
	s.Period = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetPeriodUnit(v string) *CreateElasticityAssuranceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetRecurrenceRules(v []*CreateElasticityAssuranceRequestRecurrenceRules) *CreateElasticityAssuranceRequest {
	s.RecurrenceRules = v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetRegionId(v string) *CreateElasticityAssuranceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetResourceGroupId(v string) *CreateElasticityAssuranceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetResourceOwnerAccount(v string) *CreateElasticityAssuranceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetResourceOwnerId(v int64) *CreateElasticityAssuranceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetStartTime(v string) *CreateElasticityAssuranceRequest {
	s.StartTime = &v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetTag(v []*CreateElasticityAssuranceRequestTag) *CreateElasticityAssuranceRequest {
	s.Tag = v
	return s
}

func (s *CreateElasticityAssuranceRequest) SetZoneId(v []*string) *CreateElasticityAssuranceRequest {
	s.ZoneId = v
	return s
}

func (s *CreateElasticityAssuranceRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.RecurrenceRules != nil {
		for _, item := range s.RecurrenceRules {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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

type CreateElasticityAssuranceRequestPrivatePoolOptions struct {
	// The type of the private pool with which you want to associate the elasticity assurance. Valid values:
	//
	// 	- Open: open private pool. If you use the elasticity assurance to create ECS instances, the open private pool that is associated with the elasticity assurance is automatically matched. If no capacity is available in the open private pool, resources in the public pool are automatically used to create the ECS instances.
	//
	// 	- Target: targeted private pool. If you use the elasticity assurance to create ECS instances, the specified private pool that is associated with the elasticity assurance is automatically matched. If no capacity is available in the private pool, the ECS instances fail to be created.
	//
	// Default value: Open.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
	// The name of the elasticity assurance. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// eapTestName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s CreateElasticityAssuranceRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticityAssuranceRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateElasticityAssuranceRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *CreateElasticityAssuranceRequestPrivatePoolOptions) GetName() *string {
	return s.Name
}

func (s *CreateElasticityAssuranceRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateElasticityAssuranceRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateElasticityAssuranceRequestPrivatePoolOptions) SetName(v string) *CreateElasticityAssuranceRequestPrivatePoolOptions {
	s.Name = &v
	return s
}

func (s *CreateElasticityAssuranceRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateElasticityAssuranceRequestRecurrenceRules struct {
	// The end time of the assurance period for the capacity reservation. Specify an on-the-hour point in time.
	//
	// example:
	//
	// 10
	EndHour *int32 `json:"EndHour,omitempty" xml:"EndHour,omitempty"`
	// The type of the assurance schedule. Valid values:
	//
	// 	- Daily
	//
	// 	- Weekly
	//
	// 	- Monthly
	//
	// >  You must specify both `RecurrenceType` and `RecurrenceValue`.
	//
	// example:
	//
	// Daily
	RecurrenceType *string `json:"RecurrenceType,omitempty" xml:"RecurrenceType,omitempty"`
	// The days of the week or month on which the capacity reservation takes effect or the interval, in number of days, at which the capacity reservation takes effect.
	//
	// 	- If you set `RecurrenceType` to `Daily`, you can specify only one value for this parameter. Valid values: 1 to 31. The value specifies that the capacity reservation takes effect every few days.
	//
	// 	- If you set `RecurrenceType` to `Weekly`, you can specify multiple values for this parameter. Separate the values with commas (,). Valid values: 0, 1, 2, 3, 4, 5, and 6, which specify Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, and Saturday, respectively. Example: `1,2`, which specifies that the capacity reservation takes effect on Monday and Tuesday.
	//
	// 	- If you set `RecurrenceType` to `Monthly`, you can specify two values in the `A-B` format for this parameter. Valid values of A and B: 1 to 31. B must be greater than or equal to A. For example, `1-5` indicates that the execution is repeated from the 1st to 5th of each month.
	//
	// > You must specify both `RecurrenceType` and `RecurrenceValue`.
	//
	// example:
	//
	// 1
	RecurrenceValue *string `json:"RecurrenceValue,omitempty" xml:"RecurrenceValue,omitempty"`
	// The start time of the assurance period for the capacity reservation. Specify an on-the-hour point in time.
	//
	// >  You must specify both `StartHour` and `EndHour`. EndHour must be at least four hours later than StartHour.
	//
	// example:
	//
	// 4
	StartHour *int32 `json:"StartHour,omitempty" xml:"StartHour,omitempty"`
}

func (s CreateElasticityAssuranceRequestRecurrenceRules) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticityAssuranceRequestRecurrenceRules) GoString() string {
	return s.String()
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) GetEndHour() *int32 {
	return s.EndHour
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) GetRecurrenceType() *string {
	return s.RecurrenceType
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) GetRecurrenceValue() *string {
	return s.RecurrenceValue
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) GetStartHour() *int32 {
	return s.StartHour
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) SetEndHour(v int32) *CreateElasticityAssuranceRequestRecurrenceRules {
	s.EndHour = &v
	return s
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) SetRecurrenceType(v string) *CreateElasticityAssuranceRequestRecurrenceRules {
	s.RecurrenceType = &v
	return s
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) SetRecurrenceValue(v string) *CreateElasticityAssuranceRequestRecurrenceRules {
	s.RecurrenceValue = &v
	return s
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) SetStartHour(v int32) *CreateElasticityAssuranceRequestRecurrenceRules {
	s.StartHour = &v
	return s
}

func (s *CreateElasticityAssuranceRequestRecurrenceRules) Validate() error {
	return dara.Validate(s)
}

type CreateElasticityAssuranceRequestTag struct {
	// The key of tag N to add to the elasticity assurance. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the elasticity assurance. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot start with `acs:`. The tag value cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateElasticityAssuranceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateElasticityAssuranceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateElasticityAssuranceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateElasticityAssuranceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateElasticityAssuranceRequestTag) SetKey(v string) *CreateElasticityAssuranceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateElasticityAssuranceRequestTag) SetValue(v string) *CreateElasticityAssuranceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateElasticityAssuranceRequestTag) Validate() error {
	return dara.Validate(s)
}
