// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccountAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttributeName(v []*string) *DescribeAccountAttributesRequest
	GetAttributeName() []*string
	SetOwnerId(v int64) *DescribeAccountAttributesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeAccountAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAccountAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAccountAttributesRequest
	GetResourceOwnerId() *int64
	SetZoneId(v string) *DescribeAccountAttributesRequest
	GetZoneId() *string
}

type DescribeAccountAttributesRequest struct {
	// The type of resource quota N. Valid values of N: 1 to 8. Valid values:
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
	// 	- used-postpaid-instance-vcpu-count: the number of vCPUs that have been allocated to pay-as-you-go instances.
	//
	// 	- used-spot-instance-vcpu-count: the number of vCPUs that have been allocated to spot instances.
	//
	// 	- max-postpaid-yundisk-capacity: the maximum capacity of pay-as-you-go data disks. (The value is deprecated.)
	//
	// 	- used-postpaid-yundisk-capacity: the capacity of pay-as-you-go data disks that have been created. (The value is deprecated.)
	//
	// 	- max-dedicated-hosts: the maximum number of dedicated hosts.
	//
	// 	- supported-postpaid-instance-types: the instance types of pay-as-you-go I/O optimized instances.
	//
	// 	- max-axt-command-count: the maximum number of Cloud Assistant commands.
	//
	// 	- max-axt-invocation-daily: the maximum number of Cloud Assistant command executions per day.
	//
	// 	- real-name-authentication: whether the account has completed the real-name verification.
	//
	//     **
	//
	//     **Note*	- To create an ECS instance in a region in the Chinese mainland, you must complete the real-name verification.
	//
	// 	- max-cloud-assistant-activation-count: the maximum number of activation codes that can be created to use to register managed instances.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// max-security-groups
	AttributeName []*string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty" type:"Repeated"`
	OwnerId       *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the zone in which the resource resides.
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAccountAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccountAttributesRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccountAttributesRequest) GetAttributeName() []*string {
	return s.AttributeName
}

func (s *DescribeAccountAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAccountAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccountAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAccountAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAccountAttributesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAccountAttributesRequest) SetAttributeName(v []*string) *DescribeAccountAttributesRequest {
	s.AttributeName = v
	return s
}

func (s *DescribeAccountAttributesRequest) SetOwnerId(v int64) *DescribeAccountAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAccountAttributesRequest) SetRegionId(v string) *DescribeAccountAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAccountAttributesRequest) SetResourceOwnerAccount(v string) *DescribeAccountAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAccountAttributesRequest) SetResourceOwnerId(v int64) *DescribeAccountAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAccountAttributesRequest) SetZoneId(v string) *DescribeAccountAttributesRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeAccountAttributesRequest) Validate() error {
	return dara.Validate(s)
}
