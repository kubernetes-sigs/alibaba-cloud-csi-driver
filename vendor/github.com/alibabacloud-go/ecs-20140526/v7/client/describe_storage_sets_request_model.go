// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DescribeStorageSetsRequest
	GetClientToken() *string
	SetOwnerAccount(v string) *DescribeStorageSetsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeStorageSetsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeStorageSetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStorageSetsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeStorageSetsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeStorageSetsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeStorageSetsRequest
	GetResourceOwnerId() *int64
	SetStorageSetIds(v string) *DescribeStorageSetsRequest
	GetStorageSetIds() *string
	SetStorageSetName(v string) *DescribeStorageSetsRequest
	GetStorageSetName() *string
	SetTag(v []*DescribeStorageSetsRequestTag) *DescribeStorageSetsRequest
	GetTag() []*DescribeStorageSetsRequestTag
	SetZoneId(v string) *DescribeStorageSetsRequest
	GetZoneId() *string
}

type DescribeStorageSetsRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the value, but you must make sure that it is unique among different requests. The value of **ClientToken*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// example:
	//
	// hide
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 111
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page. Maximum value: 100.
	//
	// Default value: 10
	//
	// example:
	//
	// 1
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
	// The IDs of storage sets. The value is a JSON array that consists of up to 100 storage set IDs. Separate the IDs with commas (,).
	//
	// example:
	//
	// ["ss-bp1d6tsvznfghy7y****", "ss-bp1ippxbaql9zet7****", … "ss-bp1ib7bcz07l****"]
	StorageSetIds *string `json:"StorageSetIds,omitempty" xml:"StorageSetIds,omitempty"`
	// The name of the storage set.
	//
	// example:
	//
	// storageSetTest
	StorageSetName *string                          `json:"StorageSetName,omitempty" xml:"StorageSetName,omitempty"`
	Tag            []*DescribeStorageSetsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID of the storage set. You can call the [DescribeZones](https://help.aliyun.com/document_detail/25610.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeStorageSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetsRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DescribeStorageSetsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeStorageSetsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStorageSetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStorageSetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStorageSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStorageSetsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeStorageSetsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeStorageSetsRequest) GetStorageSetIds() *string {
	return s.StorageSetIds
}

func (s *DescribeStorageSetsRequest) GetStorageSetName() *string {
	return s.StorageSetName
}

func (s *DescribeStorageSetsRequest) GetTag() []*DescribeStorageSetsRequestTag {
	return s.Tag
}

func (s *DescribeStorageSetsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeStorageSetsRequest) SetClientToken(v string) *DescribeStorageSetsRequest {
	s.ClientToken = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetOwnerAccount(v string) *DescribeStorageSetsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetOwnerId(v int64) *DescribeStorageSetsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetPageNumber(v int32) *DescribeStorageSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetPageSize(v int32) *DescribeStorageSetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetRegionId(v string) *DescribeStorageSetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetResourceOwnerAccount(v string) *DescribeStorageSetsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetResourceOwnerId(v int64) *DescribeStorageSetsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetStorageSetIds(v string) *DescribeStorageSetsRequest {
	s.StorageSetIds = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetStorageSetName(v string) *DescribeStorageSetsRequest {
	s.StorageSetName = &v
	return s
}

func (s *DescribeStorageSetsRequest) SetTag(v []*DescribeStorageSetsRequestTag) *DescribeStorageSetsRequest {
	s.Tag = v
	return s
}

func (s *DescribeStorageSetsRequest) SetZoneId(v string) *DescribeStorageSetsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeStorageSetsRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStorageSetsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeStorageSetsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageSetsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeStorageSetsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeStorageSetsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeStorageSetsRequestTag) SetKey(v string) *DescribeStorageSetsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeStorageSetsRequestTag) SetValue(v string) *DescribeStorageSetsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeStorageSetsRequestTag) Validate() error {
	return dara.Validate(s)
}
