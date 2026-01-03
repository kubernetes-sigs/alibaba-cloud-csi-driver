// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBatchOptimization(v string) *StartInstancesRequest
	GetBatchOptimization() *string
	SetDryRun(v bool) *StartInstancesRequest
	GetDryRun() *bool
	SetInstanceId(v []*string) *StartInstancesRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *StartInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartInstancesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartInstancesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartInstancesRequest
	GetResourceOwnerId() *int64
}

type StartInstancesRequest struct {
	// The batch operation mode. Valid values:
	//
	// 	- AllTogether: starts all ECS instances at the same time. If all ECS instances are started, a success message is returned. If an ECS instance fails to be started, all the specified instances fail to be started and an error message is returned.
	//
	// 	- SuccessFirst: separately starts each ECS instance. The response contains the operation results of each ECS instance.
	//
	// Default value: AllTogether.
	//
	// example:
	//
	// AllTogether
	BatchOptimization *string `json:"BatchOptimization,omitempty" xml:"BatchOptimization,omitempty"`
	// Specifies whether to perform a dry run. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including required parameters, request syntax, and instance status. If the request fails the dry run, an error message is returned. If the request passes the dry run, `DRYRUN.SUCCESS` is returned.
	//
	// > If you set `BatchOptimization` to `SuccessFirst` and `DryRun` to true, only `DRYRUN.SUCCESS` is returned regardless of whether the request passes the dry run.
	//
	// 	- false: performs a dry run and performs the actual request. If the request passes the dry run, the operation is performed.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The IDs of ECS instances. Valid values of N: 1 to 100.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the ECS instance. You can call the [DescribeRegions](~~DescribeRegions~~) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s StartInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s StartInstancesRequest) GoString() string {
	return s.String()
}

func (s *StartInstancesRequest) GetBatchOptimization() *string {
	return s.BatchOptimization
}

func (s *StartInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StartInstancesRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *StartInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartInstancesRequest) SetBatchOptimization(v string) *StartInstancesRequest {
	s.BatchOptimization = &v
	return s
}

func (s *StartInstancesRequest) SetDryRun(v bool) *StartInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *StartInstancesRequest) SetInstanceId(v []*string) *StartInstancesRequest {
	s.InstanceId = v
	return s
}

func (s *StartInstancesRequest) SetOwnerAccount(v string) *StartInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartInstancesRequest) SetOwnerId(v int64) *StartInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *StartInstancesRequest) SetRegionId(v string) *StartInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *StartInstancesRequest) SetResourceOwnerAccount(v string) *StartInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartInstancesRequest) SetResourceOwnerId(v int64) *StartInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartInstancesRequest) Validate() error {
	return dara.Validate(s)
}
