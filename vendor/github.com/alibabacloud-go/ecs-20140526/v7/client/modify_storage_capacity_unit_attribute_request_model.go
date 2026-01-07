// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyStorageCapacityUnitAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyStorageCapacityUnitAttributeRequest
	GetDescription() *string
	SetName(v string) *ModifyStorageCapacityUnitAttributeRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyStorageCapacityUnitAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyStorageCapacityUnitAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyStorageCapacityUnitAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyStorageCapacityUnitAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyStorageCapacityUnitAttributeRequest
	GetResourceOwnerId() *int64
	SetStorageCapacityUnitId(v string) *ModifyStorageCapacityUnitAttributeRequest
	GetStorageCapacityUnitId() *string
}

type ModifyStorageCapacityUnitAttributeRequest struct {
	// The new description of the SCU. The description must be 2 to 256 characters in length and cannot start with [http:// or https://.](http://https://。)
	//
	// example:
	//
	// testNewScuDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The new name of the SCU. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with [http:// or https://. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).](http://https://。、（:）、（_）（-）。)
	//
	// example:
	//
	// testNewScuName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// hide
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	// example:
	//
	// 111
	OwnerId *int64 `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the SCU. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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
	// The ID of the SCU.
	//
	// This parameter is required.
	//
	// example:
	//
	// scu-bp67acfmxazb4p****
	StorageCapacityUnitId *string `json:"StorageCapacityUnitId,omitempty" xml:"StorageCapacityUnitId,omitempty"`
}

func (s ModifyStorageCapacityUnitAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyStorageCapacityUnitAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyStorageCapacityUnitAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyStorageCapacityUnitAttributeRequest) GetName() *string {
	return s.Name
}

func (s *ModifyStorageCapacityUnitAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyStorageCapacityUnitAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyStorageCapacityUnitAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyStorageCapacityUnitAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyStorageCapacityUnitAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyStorageCapacityUnitAttributeRequest) GetStorageCapacityUnitId() *string {
	return s.StorageCapacityUnitId
}

func (s *ModifyStorageCapacityUnitAttributeRequest) SetDescription(v string) *ModifyStorageCapacityUnitAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeRequest) SetName(v string) *ModifyStorageCapacityUnitAttributeRequest {
	s.Name = &v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeRequest) SetOwnerAccount(v string) *ModifyStorageCapacityUnitAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeRequest) SetOwnerId(v int64) *ModifyStorageCapacityUnitAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeRequest) SetRegionId(v string) *ModifyStorageCapacityUnitAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeRequest) SetResourceOwnerAccount(v string) *ModifyStorageCapacityUnitAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeRequest) SetResourceOwnerId(v int64) *ModifyStorageCapacityUnitAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeRequest) SetStorageCapacityUnitId(v string) *ModifyStorageCapacityUnitAttributeRequest {
	s.StorageCapacityUnitId = &v
	return s
}

func (s *ModifyStorageCapacityUnitAttributeRequest) Validate() error {
	return dara.Validate(s)
}
