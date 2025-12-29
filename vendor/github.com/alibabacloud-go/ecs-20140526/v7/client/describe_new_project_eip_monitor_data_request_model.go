// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNewProjectEipMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *DescribeNewProjectEipMonitorDataRequest
	GetAllocationId() *string
	SetEndTime(v string) *DescribeNewProjectEipMonitorDataRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeNewProjectEipMonitorDataRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeNewProjectEipMonitorDataRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *DescribeNewProjectEipMonitorDataRequest
	GetPeriod() *int32
	SetRegionId(v string) *DescribeNewProjectEipMonitorDataRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeNewProjectEipMonitorDataRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNewProjectEipMonitorDataRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeNewProjectEipMonitorDataRequest
	GetStartTime() *string
}

type DescribeNewProjectEipMonitorDataRequest struct {
	// This parameter is required.
	AllocationId *string `json:"AllocationId,omitempty" xml:"AllocationId,omitempty"`
	// This parameter is required.
	EndTime              *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	Period               *int32  `json:"Period,omitempty" xml:"Period,omitempty"`
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// This parameter is required.
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeNewProjectEipMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNewProjectEipMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeNewProjectEipMonitorDataRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeNewProjectEipMonitorDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeNewProjectEipMonitorDataRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeNewProjectEipMonitorDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNewProjectEipMonitorDataRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeNewProjectEipMonitorDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNewProjectEipMonitorDataRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNewProjectEipMonitorDataRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNewProjectEipMonitorDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeNewProjectEipMonitorDataRequest) SetAllocationId(v string) *DescribeNewProjectEipMonitorDataRequest {
	s.AllocationId = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataRequest) SetEndTime(v string) *DescribeNewProjectEipMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataRequest) SetOwnerAccount(v string) *DescribeNewProjectEipMonitorDataRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataRequest) SetOwnerId(v int64) *DescribeNewProjectEipMonitorDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataRequest) SetPeriod(v int32) *DescribeNewProjectEipMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataRequest) SetRegionId(v string) *DescribeNewProjectEipMonitorDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataRequest) SetResourceOwnerAccount(v string) *DescribeNewProjectEipMonitorDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataRequest) SetResourceOwnerId(v int64) *DescribeNewProjectEipMonitorDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataRequest) SetStartTime(v string) *DescribeNewProjectEipMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeNewProjectEipMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
