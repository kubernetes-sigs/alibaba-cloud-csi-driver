// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotsUsageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerAccount(v string) *DescribeSnapshotsUsageRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSnapshotsUsageRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeSnapshotsUsageRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSnapshotsUsageRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSnapshotsUsageRequest
	GetResourceOwnerId() *int64
}

type DescribeSnapshotsUsageRequest struct {
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the snapshot. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DescribeSnapshotsUsageRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotsUsageRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotsUsageRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSnapshotsUsageRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSnapshotsUsageRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotsUsageRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSnapshotsUsageRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSnapshotsUsageRequest) SetOwnerAccount(v string) *DescribeSnapshotsUsageRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSnapshotsUsageRequest) SetOwnerId(v int64) *DescribeSnapshotsUsageRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSnapshotsUsageRequest) SetRegionId(v string) *DescribeSnapshotsUsageRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotsUsageRequest) SetResourceOwnerAccount(v string) *DescribeSnapshotsUsageRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSnapshotsUsageRequest) SetResourceOwnerId(v int64) *DescribeSnapshotsUsageRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSnapshotsUsageRequest) Validate() error {
	return dara.Validate(s)
}
