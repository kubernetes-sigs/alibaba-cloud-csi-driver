// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAutoProvisioningGroupHistoryRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoProvisioningGroupId(v string) *DescribeAutoProvisioningGroupHistoryRequest
	GetAutoProvisioningGroupId() *string
	SetEndTime(v string) *DescribeAutoProvisioningGroupHistoryRequest
	GetEndTime() *string
	SetOwnerAccount(v string) *DescribeAutoProvisioningGroupHistoryRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeAutoProvisioningGroupHistoryRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeAutoProvisioningGroupHistoryRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAutoProvisioningGroupHistoryRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeAutoProvisioningGroupHistoryRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeAutoProvisioningGroupHistoryRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeAutoProvisioningGroupHistoryRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *DescribeAutoProvisioningGroupHistoryRequest
	GetStartTime() *string
}

type DescribeAutoProvisioningGroupHistoryRequest struct {
	// The ID of the auto provisioning group.
	//
	// This parameter is required.
	//
	// example:
	//
	// apg-bp67acfmxazb4p****
	AutoProvisioningGroupId *string `json:"AutoProvisioningGroupId,omitempty" xml:"AutoProvisioningGroupId,omitempty"`
	// The end of the time range of the queried data. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-06-20T15:10:20Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// 123456
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 123456
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number of the returned page. Pages start from page 1.
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// and the default value is 10.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the auto provisioning group.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// 123456
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// example:
	//
	// 123456
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The beginning of the time range of the queried data. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2019-04-01T15:10:20Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeAutoProvisioningGroupHistoryRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAutoProvisioningGroupHistoryRequest) GoString() string {
	return s.String()
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) GetAutoProvisioningGroupId() *string {
	return s.AutoProvisioningGroupId
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) SetAutoProvisioningGroupId(v string) *DescribeAutoProvisioningGroupHistoryRequest {
	s.AutoProvisioningGroupId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) SetEndTime(v string) *DescribeAutoProvisioningGroupHistoryRequest {
	s.EndTime = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) SetOwnerAccount(v string) *DescribeAutoProvisioningGroupHistoryRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) SetOwnerId(v int64) *DescribeAutoProvisioningGroupHistoryRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) SetPageNumber(v int32) *DescribeAutoProvisioningGroupHistoryRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) SetPageSize(v int32) *DescribeAutoProvisioningGroupHistoryRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) SetRegionId(v string) *DescribeAutoProvisioningGroupHistoryRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) SetResourceOwnerAccount(v string) *DescribeAutoProvisioningGroupHistoryRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) SetResourceOwnerId(v int64) *DescribeAutoProvisioningGroupHistoryRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) SetStartTime(v string) *DescribeAutoProvisioningGroupHistoryRequest {
	s.StartTime = &v
	return s
}

func (s *DescribeAutoProvisioningGroupHistoryRequest) Validate() error {
	return dara.Validate(s)
}
