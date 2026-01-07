// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOpenSnapshotServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *OpenSnapshotServiceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *OpenSnapshotServiceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *OpenSnapshotServiceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *OpenSnapshotServiceRequest
	GetResourceOwnerId() *int64
}

type OpenSnapshotServiceRequest struct {
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the port list. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-beijing
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s OpenSnapshotServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s OpenSnapshotServiceRequest) GoString() string {
	return s.String()
}

func (s *OpenSnapshotServiceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *OpenSnapshotServiceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *OpenSnapshotServiceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *OpenSnapshotServiceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *OpenSnapshotServiceRequest) SetOwnerId(v int64) *OpenSnapshotServiceRequest {
	s.OwnerId = &v
	return s
}

func (s *OpenSnapshotServiceRequest) SetRegionId(v string) *OpenSnapshotServiceRequest {
	s.RegionId = &v
	return s
}

func (s *OpenSnapshotServiceRequest) SetResourceOwnerAccount(v string) *OpenSnapshotServiceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *OpenSnapshotServiceRequest) SetResourceOwnerId(v int64) *OpenSnapshotServiceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *OpenSnapshotServiceRequest) Validate() error {
	return dara.Validate(s)
}
