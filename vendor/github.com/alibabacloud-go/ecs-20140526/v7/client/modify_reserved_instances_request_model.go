// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyReservedInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetConfiguration(v []*ModifyReservedInstancesRequestConfiguration) *ModifyReservedInstancesRequest
	GetConfiguration() []*ModifyReservedInstancesRequestConfiguration
	SetOwnerAccount(v string) *ModifyReservedInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyReservedInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyReservedInstancesRequest
	GetRegionId() *string
	SetReservedInstanceId(v []*string) *ModifyReservedInstancesRequest
	GetReservedInstanceId() []*string
	SetResourceOwnerAccount(v string) *ModifyReservedInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyReservedInstancesRequest
	GetResourceOwnerId() *int64
}

type ModifyReservedInstancesRequest struct {
	// The configurations of the new reserved instances. You can specify up to 100 new reserved instances.
	Configuration []*ModifyReservedInstancesRequestConfiguration `json:"Configuration,omitempty" xml:"Configuration,omitempty" type:"Repeated"`
	OwnerAccount  *string                                        `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId       *int64                                         `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the reserved instance.
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IDs of reserved instances that you want to modify. You can specify up to 20 reserved instance IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// ReservedInstanceId.1="ecsri-bp1cx3****",ReservedInstanceId.2="ecsri-bp15xx2****"......
	ReservedInstanceId   []*string `json:"ReservedInstanceId,omitempty" xml:"ReservedInstanceId,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string   `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64    `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyReservedInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstancesRequest) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstancesRequest) GetConfiguration() []*ModifyReservedInstancesRequestConfiguration {
	return s.Configuration
}

func (s *ModifyReservedInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyReservedInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyReservedInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyReservedInstancesRequest) GetReservedInstanceId() []*string {
	return s.ReservedInstanceId
}

func (s *ModifyReservedInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyReservedInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyReservedInstancesRequest) SetConfiguration(v []*ModifyReservedInstancesRequestConfiguration) *ModifyReservedInstancesRequest {
	s.Configuration = v
	return s
}

func (s *ModifyReservedInstancesRequest) SetOwnerAccount(v string) *ModifyReservedInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyReservedInstancesRequest) SetOwnerId(v int64) *ModifyReservedInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyReservedInstancesRequest) SetRegionId(v string) *ModifyReservedInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyReservedInstancesRequest) SetReservedInstanceId(v []*string) *ModifyReservedInstancesRequest {
	s.ReservedInstanceId = v
	return s
}

func (s *ModifyReservedInstancesRequest) SetResourceOwnerAccount(v string) *ModifyReservedInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyReservedInstancesRequest) SetResourceOwnerId(v int64) *ModifyReservedInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyReservedInstancesRequest) Validate() error {
	if s.Configuration != nil {
		for _, item := range s.Configuration {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyReservedInstancesRequestConfiguration struct {
	// The number of pay-as-you-go instances of the specified instance type that the new reserved instance can match. The value of this parameter must be greater than or equal to 1.
	//
	// example:
	//
	// 1
	InstanceAmount *int32 `json:"InstanceAmount,omitempty" xml:"InstanceAmount,omitempty"`
	// The instance types that the new reserved instance can match.
	//
	// >  The supported instance types are continuously updated. For information about the instance types supported by reserved instances, see [Overview of reserved instances](~~100370#3c1b682051vt4~~).
	//
	// example:
	//
	// ecs.c5.4xlarge
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of the new reserved instance.
	//
	// The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with http:// or https://. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testReservedInstanceName
	ReservedInstanceName *string `json:"ReservedInstanceName,omitempty" xml:"ReservedInstanceName,omitempty"`
	// The scope level of the new reserved instance. Valid values:
	//
	// 	- Region
	//
	// 	- Zone
	//
	// Default value: Region.
	//
	// example:
	//
	// Zone
	Scope *string `json:"Scope,omitempty" xml:"Scope,omitempty"`
	// The zone ID of the new reserved instance.
	//
	// This parameter is required when you set `Scope` to `Zone`.
	//
	// You can call the [DescribeZones](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ModifyReservedInstancesRequestConfiguration) String() string {
	return dara.Prettify(s)
}

func (s ModifyReservedInstancesRequestConfiguration) GoString() string {
	return s.String()
}

func (s *ModifyReservedInstancesRequestConfiguration) GetInstanceAmount() *int32 {
	return s.InstanceAmount
}

func (s *ModifyReservedInstancesRequestConfiguration) GetInstanceType() *string {
	return s.InstanceType
}

func (s *ModifyReservedInstancesRequestConfiguration) GetReservedInstanceName() *string {
	return s.ReservedInstanceName
}

func (s *ModifyReservedInstancesRequestConfiguration) GetScope() *string {
	return s.Scope
}

func (s *ModifyReservedInstancesRequestConfiguration) GetZoneId() *string {
	return s.ZoneId
}

func (s *ModifyReservedInstancesRequestConfiguration) SetInstanceAmount(v int32) *ModifyReservedInstancesRequestConfiguration {
	s.InstanceAmount = &v
	return s
}

func (s *ModifyReservedInstancesRequestConfiguration) SetInstanceType(v string) *ModifyReservedInstancesRequestConfiguration {
	s.InstanceType = &v
	return s
}

func (s *ModifyReservedInstancesRequestConfiguration) SetReservedInstanceName(v string) *ModifyReservedInstancesRequestConfiguration {
	s.ReservedInstanceName = &v
	return s
}

func (s *ModifyReservedInstancesRequestConfiguration) SetScope(v string) *ModifyReservedInstancesRequestConfiguration {
	s.Scope = &v
	return s
}

func (s *ModifyReservedInstancesRequestConfiguration) SetZoneId(v string) *ModifyReservedInstancesRequestConfiguration {
	s.ZoneId = &v
	return s
}

func (s *ModifyReservedInstancesRequestConfiguration) Validate() error {
	return dara.Validate(s)
}
