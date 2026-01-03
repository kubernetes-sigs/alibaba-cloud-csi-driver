// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEniMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeEniMonitorDataRequest
	GetEndTime() *string
	SetEniId(v string) *DescribeEniMonitorDataRequest
	GetEniId() *string
	SetInstanceId(v string) *DescribeEniMonitorDataRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeEniMonitorDataRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEniMonitorDataRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *DescribeEniMonitorDataRequest
	GetPeriod() *int32
	SetRegionId(v string) *DescribeEniMonitorDataRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeEniMonitorDataRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEniMonitorDataRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeEniMonitorDataRequest
	GetStartTime() *string
}

type DescribeEniMonitorDataRequest struct {
	// The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. If the value of seconds (ss) is not 00, the time is rounded up to the next minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-05-21T12:22:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The secondary ENI ID. By default, all secondary ENIs that are bound to the specified instance are queried.
	//
	// example:
	//
	// eni-bp19da36d6xdwey****
	EniId *string `json:"EniId,omitempty" xml:"EniId,omitempty"`
	// The ID of the instance to which the secondary ENI is bound.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1a5zr3u7nq9cx****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The interval at which to retrieve the monitoring data. Unit: seconds. Default value: Month. Valid values:
	//
	// 	- 60
	//
	// 	- 600
	//
	// 	- 3600
	//
	// Default: 60.
	//
	// example:
	//
	// 60
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. If the value of seconds (ss) is not 00, the time is rounded up to the next minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2018-05-21T12:19:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeEniMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEniMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEniMonitorDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEniMonitorDataRequest) GetEniId() *string {
	return s.EniId
}

func (s *DescribeEniMonitorDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeEniMonitorDataRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEniMonitorDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEniMonitorDataRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeEniMonitorDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEniMonitorDataRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEniMonitorDataRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEniMonitorDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEniMonitorDataRequest) SetEndTime(v string) *DescribeEniMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEniMonitorDataRequest) SetEniId(v string) *DescribeEniMonitorDataRequest {
	s.EniId = &v
	return s
}

func (s *DescribeEniMonitorDataRequest) SetInstanceId(v string) *DescribeEniMonitorDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeEniMonitorDataRequest) SetOwnerAccount(v string) *DescribeEniMonitorDataRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEniMonitorDataRequest) SetOwnerId(v int64) *DescribeEniMonitorDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEniMonitorDataRequest) SetPeriod(v int32) *DescribeEniMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeEniMonitorDataRequest) SetRegionId(v string) *DescribeEniMonitorDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEniMonitorDataRequest) SetResourceOwnerAccount(v string) *DescribeEniMonitorDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEniMonitorDataRequest) SetResourceOwnerId(v int64) *DescribeEniMonitorDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEniMonitorDataRequest) SetStartTime(v string) *DescribeEniMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEniMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
