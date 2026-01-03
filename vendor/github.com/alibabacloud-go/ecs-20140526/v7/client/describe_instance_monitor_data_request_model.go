// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTime(v string) *DescribeInstanceMonitorDataRequest
	GetEndTime() *string
	SetInstanceId(v string) *DescribeInstanceMonitorDataRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *DescribeInstanceMonitorDataRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInstanceMonitorDataRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *DescribeInstanceMonitorDataRequest
	GetPeriod() *int32
	SetResourceOwnerAccount(v string) *DescribeInstanceMonitorDataRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInstanceMonitorDataRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeInstanceMonitorDataRequest
	GetStartTime() *string
}

type DescribeInstanceMonitorDataRequest struct {
	// The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. If the value of seconds (`ss`) is not `00`, the time is rounded up to the next minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2014-10-30T08:00:00Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1a36962lrhj4ab****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The interval at which to retrieve monitoring data. Unit: seconds. Valid values:
	//
	// 	- 60
	//
	// 	- 600
	//
	// 	- 3600
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. If the value of seconds (`ss`) is not `00`, the time is rounded up to the next minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2014-10-29T23:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeInstanceMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeInstanceMonitorDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeInstanceMonitorDataRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceMonitorDataRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInstanceMonitorDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInstanceMonitorDataRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeInstanceMonitorDataRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInstanceMonitorDataRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInstanceMonitorDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInstanceMonitorDataRequest) SetEndTime(v string) *DescribeInstanceMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetInstanceId(v string) *DescribeInstanceMonitorDataRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetOwnerAccount(v string) *DescribeInstanceMonitorDataRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetOwnerId(v int64) *DescribeInstanceMonitorDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetPeriod(v int32) *DescribeInstanceMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetResourceOwnerAccount(v string) *DescribeInstanceMonitorDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetResourceOwnerId(v int64) *DescribeInstanceMonitorDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) SetStartTime(v string) *DescribeInstanceMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeInstanceMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
