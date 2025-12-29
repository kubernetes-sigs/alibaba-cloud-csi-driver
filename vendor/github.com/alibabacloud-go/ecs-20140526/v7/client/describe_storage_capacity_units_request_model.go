// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStorageCapacityUnitsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocationType(v string) *DescribeStorageCapacityUnitsRequest
	GetAllocationType() *string
	SetCapacity(v int32) *DescribeStorageCapacityUnitsRequest
	GetCapacity() *int32
	SetName(v string) *DescribeStorageCapacityUnitsRequest
	GetName() *string
	SetOwnerAccount(v string) *DescribeStorageCapacityUnitsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeStorageCapacityUnitsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeStorageCapacityUnitsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeStorageCapacityUnitsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeStorageCapacityUnitsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeStorageCapacityUnitsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeStorageCapacityUnitsRequest
	GetResourceOwnerId() *int64
	SetStatus(v []*string) *DescribeStorageCapacityUnitsRequest
	GetStatus() []*string
	SetStorageCapacityUnitId(v []*string) *DescribeStorageCapacityUnitsRequest
	GetStorageCapacityUnitId() []*string
	SetTag(v []*DescribeStorageCapacityUnitsRequestTag) *DescribeStorageCapacityUnitsRequest
	GetTag() []*DescribeStorageCapacityUnitsRequestTag
}

type DescribeStorageCapacityUnitsRequest struct {
	// The allocation type. Valid values:
	//
	// 	- Normal: queries SCUs that belong to the current Alibaba Cloud account.
	//
	// 	- Shared: queries SCUs shared between the Alibaba Cloud account and RAM users.
	//
	// Default value: Normal.
	//
	// example:
	//
	// Normal
	AllocationType *string `json:"AllocationType,omitempty" xml:"AllocationType,omitempty"`
	// The capacity of the SCU. Unit: GiB. Valid values: 20, 40, 100, 200, 500, 1024, 2048, 5120, 10240, 20480, and 51200.
	//
	// example:
	//
	// 20
	Capacity *int32 `json:"Capacity,omitempty" xml:"Capacity,omitempty"`
	// The name of the SCU. The name must be 2 to 128 characters in length. It must start with a letter but cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testScuName
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of the page to return.
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
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the SCU. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The states of SCUs. The array is 1 to 4 in length.
	//
	// example:
	//
	// Active
	Status []*string `json:"Status,omitempty" xml:"Status,omitempty" type:"Repeated"`
	// The IDs of the SCUs. You can specify 1 to 100 SCU IDs.
	//
	// example:
	//
	// scu-bp67acfmxazb4p****
	StorageCapacityUnitId []*string `json:"StorageCapacityUnitId,omitempty" xml:"StorageCapacityUnitId,omitempty" type:"Repeated"`
	// The tags to add to the SCU. You can add up to 20 tags.
	Tag []*DescribeStorageCapacityUnitsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeStorageCapacityUnitsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageCapacityUnitsRequest) GoString() string {
	return s.String()
}

func (s *DescribeStorageCapacityUnitsRequest) GetAllocationType() *string {
	return s.AllocationType
}

func (s *DescribeStorageCapacityUnitsRequest) GetCapacity() *int32 {
	return s.Capacity
}

func (s *DescribeStorageCapacityUnitsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeStorageCapacityUnitsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeStorageCapacityUnitsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeStorageCapacityUnitsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeStorageCapacityUnitsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeStorageCapacityUnitsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeStorageCapacityUnitsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeStorageCapacityUnitsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeStorageCapacityUnitsRequest) GetStatus() []*string {
	return s.Status
}

func (s *DescribeStorageCapacityUnitsRequest) GetStorageCapacityUnitId() []*string {
	return s.StorageCapacityUnitId
}

func (s *DescribeStorageCapacityUnitsRequest) GetTag() []*DescribeStorageCapacityUnitsRequestTag {
	return s.Tag
}

func (s *DescribeStorageCapacityUnitsRequest) SetAllocationType(v string) *DescribeStorageCapacityUnitsRequest {
	s.AllocationType = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetCapacity(v int32) *DescribeStorageCapacityUnitsRequest {
	s.Capacity = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetName(v string) *DescribeStorageCapacityUnitsRequest {
	s.Name = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetOwnerAccount(v string) *DescribeStorageCapacityUnitsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetOwnerId(v int64) *DescribeStorageCapacityUnitsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetPageNumber(v int32) *DescribeStorageCapacityUnitsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetPageSize(v int32) *DescribeStorageCapacityUnitsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetRegionId(v string) *DescribeStorageCapacityUnitsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetResourceOwnerAccount(v string) *DescribeStorageCapacityUnitsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetResourceOwnerId(v int64) *DescribeStorageCapacityUnitsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetStatus(v []*string) *DescribeStorageCapacityUnitsRequest {
	s.Status = v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetStorageCapacityUnitId(v []*string) *DescribeStorageCapacityUnitsRequest {
	s.StorageCapacityUnitId = v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) SetTag(v []*DescribeStorageCapacityUnitsRequestTag) *DescribeStorageCapacityUnitsRequest {
	s.Tag = v
	return s
}

func (s *DescribeStorageCapacityUnitsRequest) Validate() error {
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

type DescribeStorageCapacityUnitsRequestTag struct {
	// The key of tag N to be added to the SCU.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to be added to the SCU.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeStorageCapacityUnitsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeStorageCapacityUnitsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeStorageCapacityUnitsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeStorageCapacityUnitsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeStorageCapacityUnitsRequestTag) SetKey(v string) *DescribeStorageCapacityUnitsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequestTag) SetValue(v string) *DescribeStorageCapacityUnitsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeStorageCapacityUnitsRequestTag) Validate() error {
	return dara.Validate(s)
}
