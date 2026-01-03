// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeEipMonitorDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationId(v string) *DescribeEipMonitorDataRequest
	GetAllocationId() *string
	SetEndTime(v string) *DescribeEipMonitorDataRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeEipMonitorDataRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeEipMonitorDataRequest
	GetOwnerId() *int64
	SetPeriod(v int32) *DescribeEipMonitorDataRequest
	GetPeriod() *int32
	SetRegionId(v string) *DescribeEipMonitorDataRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeEipMonitorDataRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeEipMonitorDataRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeEipMonitorDataRequest
	GetStartTime() *string
}

type DescribeEipMonitorDataRequest struct {
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

func (s DescribeEipMonitorDataRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeEipMonitorDataRequest) GoString() string {
	return s.String()
}

func (s *DescribeEipMonitorDataRequest) GetAllocationId() *string {
	return s.AllocationId
}

func (s *DescribeEipMonitorDataRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeEipMonitorDataRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeEipMonitorDataRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeEipMonitorDataRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *DescribeEipMonitorDataRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeEipMonitorDataRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeEipMonitorDataRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeEipMonitorDataRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeEipMonitorDataRequest) SetAllocationId(v string) *DescribeEipMonitorDataRequest {
	s.AllocationId = &v
	return s
}

func (s *DescribeEipMonitorDataRequest) SetEndTime(v string) *DescribeEipMonitorDataRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeEipMonitorDataRequest) SetOwnerAccount(v string) *DescribeEipMonitorDataRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeEipMonitorDataRequest) SetOwnerId(v int64) *DescribeEipMonitorDataRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeEipMonitorDataRequest) SetPeriod(v int32) *DescribeEipMonitorDataRequest {
	s.Period = &v
	return s
}

func (s *DescribeEipMonitorDataRequest) SetRegionId(v string) *DescribeEipMonitorDataRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeEipMonitorDataRequest) SetResourceOwnerAccount(v string) *DescribeEipMonitorDataRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeEipMonitorDataRequest) SetResourceOwnerId(v int64) *DescribeEipMonitorDataRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeEipMonitorDataRequest) SetStartTime(v string) *DescribeEipMonitorDataRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeEipMonitorDataRequest) Validate() error {
	return dara.Validate(s)
}
