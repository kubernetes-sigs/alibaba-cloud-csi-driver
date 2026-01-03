// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *DescribeSnapshotMonitorDataRequest
	GetCategory() *string
	SetEndTime(v string) *DescribeSnapshotMonitorDataRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeSnapshotMonitorDataRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSnapshotMonitorDataRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *DescribeSnapshotMonitorDataRequest
	GetPeriod() *int32
	SetRegionId(v string) *DescribeSnapshotMonitorDataRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSnapshotMonitorDataRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSnapshotMonitorDataRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeSnapshotMonitorDataRequest
	GetStartTime() *string
}

type DescribeSnapshotMonitorDataRequest struct {
	// The type of the snapshot. Valid values:
	//
	// 	- Standard: standard snapshot
	//
	// 	- Flash: local snapshot
	//
	// 	- Archive: archive snapshot
	//
	// Default value: Standard.
	//
	// example:
	//
	// Standard
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The end of the time range to query. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC. If the value of seconds (ss) is not 00, the time is rounded up to the next minute.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2019-05-10T03:00:00Z
	EndTime      *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The interval at which to query the monitoring data of snapshot sizes. Unit: seconds. Valid values:
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
	// 2019-05-10T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeSnapshotMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotMonitorDataRequest) GetCategory() *string {
	return s.Category
}

func (s *DescribeSnapshotMonitorDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeSnapshotMonitorDataRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSnapshotMonitorDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSnapshotMonitorDataRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeSnapshotMonitorDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotMonitorDataRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSnapshotMonitorDataRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSnapshotMonitorDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSnapshotMonitorDataRequest) SetCategory(v string) *DescribeSnapshotMonitorDataRequest {
	s.Category = &v
	return s
}

func (s *DescribeSnapshotMonitorDataRequest) SetEndTime(v string) *DescribeSnapshotMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeSnapshotMonitorDataRequest) SetOwnerAccount(v string) *DescribeSnapshotMonitorDataRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSnapshotMonitorDataRequest) SetOwnerId(v int64) *DescribeSnapshotMonitorDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSnapshotMonitorDataRequest) SetPeriod(v int32) *DescribeSnapshotMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeSnapshotMonitorDataRequest) SetRegionId(v string) *DescribeSnapshotMonitorDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotMonitorDataRequest) SetResourceOwnerAccount(v string) *DescribeSnapshotMonitorDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSnapshotMonitorDataRequest) SetResourceOwnerId(v int64) *DescribeSnapshotMonitorDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSnapshotMonitorDataRequest) SetStartTime(v string) *DescribeSnapshotMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeSnapshotMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
