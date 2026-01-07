// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageSetDetailsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeStorageSetDetailsRequest
	GetClientToken() *string
	SetDiskIds(v string) *DescribeStorageSetDetailsRequest
	GetDiskIds() *string
	SetOwnerAccount(v string) *DescribeStorageSetDetailsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeStorageSetDetailsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeStorageSetDetailsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStorageSetDetailsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeStorageSetDetailsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeStorageSetDetailsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeStorageSetDetailsRequest
	GetResourceOwnerId() *int64
	SetStorageSetId(v string) *DescribeStorageSetDetailsRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *DescribeStorageSetDetailsRequest
	GetStorageSetPartitionNumber() *int32
}

type DescribeStorageSetDetailsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The IDs of disks or Shared Block Storage devices. The value can be a JSON array that consists of up to 100 disk or Shared Block Storage device IDs. Separate the IDs with commas (,).
	//
	// example:
	//
	// ["d-bp1d6tsvznfghy7y****", "d-bp1ippxbaql9zet7****", … "d-bp1ib7bcz07l****"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// example:
	//
	// hide
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 111
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Maximum value: 100
	//
	// Default value: 10
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the storage set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// hide
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	// example:
	//
	// 111
	ResourceOwnerId *int64 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the storage set.
	//
	// This parameter is required.
	//
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set.
	//
	// example:
	//
	// 3
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
}

func (s DescribeStorageSetDetailsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetDetailsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetDetailsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeStorageSetDetailsRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *DescribeStorageSetDetailsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeStorageSetDetailsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStorageSetDetailsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStorageSetDetailsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStorageSetDetailsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStorageSetDetailsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeStorageSetDetailsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeStorageSetDetailsRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *DescribeStorageSetDetailsRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *DescribeStorageSetDetailsRequest) SetClientToken(v string) *DescribeStorageSetDetailsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetDiskIds(v string) *DescribeStorageSetDetailsRequest {
	s.DiskIds = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetOwnerAccount(v string) *DescribeStorageSetDetailsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetOwnerId(v int64) *DescribeStorageSetDetailsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetPageNumber(v int32) *DescribeStorageSetDetailsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetPageSize(v int32) *DescribeStorageSetDetailsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetRegionId(v string) *DescribeStorageSetDetailsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetResourceOwnerAccount(v string) *DescribeStorageSetDetailsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetResourceOwnerId(v int64) *DescribeStorageSetDetailsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetStorageSetId(v string) *DescribeStorageSetDetailsRequest {
	s.StorageSetId = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) SetStorageSetPartitionNumber(v int32) *DescribeStorageSetDetailsRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *DescribeStorageSetDetailsRequest) Validate() error {
	return dara.Validate(s)
}
