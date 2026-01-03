// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiskMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskId(v string) *DescribeDiskMonitorDataRequest
	GetDiskId() *string
	SetEndTime(v string) *DescribeDiskMonitorDataRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeDiskMonitorDataRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDiskMonitorDataRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *DescribeDiskMonitorDataRequest
	GetPeriod() *int32
	SetResourceOwnerAccount(v string) *DescribeDiskMonitorDataRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDiskMonitorDataRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeDiskMonitorDataRequest
	GetStartTime() *string
}

type DescribeDiskMonitorDataRequest struct {
	// The ID of the cloud disk that you want to query.
	//
	// This parameter is required.
	//
	// example:
	//
	// d-bp1bq5g3dxxo1x4o****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. If the value of seconds (ss) is not 00, the time is rounded up to the next minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2014-07-23T12:09:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The interval at which to retrieve the monitoring data. Unit: seconds. Valid values:
	//
	// 	- 60
	//
	// 	- 600
	//
	// 	- 3600
	//
	// Default value: 60.
	//
	// >  Up to 400 monitoring data entries can be returned at a time. Make sure that the TotalCount value does not exceed 400. The value is calculated by using the following formula: TotalCount = (EndTime - StartTime)/Period.
	//
	// example:
	//
	// 60
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. If the value of seconds (ss) is not 00, the time is rounded up to the next minute.
	//
	// >  You can query the monitoring data in the last 30 days. If the value of `StartTime` is more than 30 days earlier than the current time, an error is returned.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2014-07-23T12:07:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeDiskMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiskMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiskMonitorDataRequest) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDiskMonitorDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiskMonitorDataRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDiskMonitorDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDiskMonitorDataRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeDiskMonitorDataRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDiskMonitorDataRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDiskMonitorDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeDiskMonitorDataRequest) SetDiskId(v string) *DescribeDiskMonitorDataRequest {
	s.DiskId = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetEndTime(v string) *DescribeDiskMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetOwnerAccount(v string) *DescribeDiskMonitorDataRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetOwnerId(v int64) *DescribeDiskMonitorDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetPeriod(v int32) *DescribeDiskMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetResourceOwnerAccount(v string) *DescribeDiskMonitorDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetResourceOwnerId(v int64) *DescribeDiskMonitorDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) SetStartTime(v string) *DescribeDiskMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeDiskMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
