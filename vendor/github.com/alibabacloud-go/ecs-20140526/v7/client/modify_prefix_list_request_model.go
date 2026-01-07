// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPrefixListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddEntry(v []*ModifyPrefixListRequestAddEntry) *ModifyPrefixListRequest
	GetAddEntry() []*ModifyPrefixListRequestAddEntry
	SetDescription(v string) *ModifyPrefixListRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyPrefixListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPrefixListRequest
	GetOwnerId() *int64
	SetPrefixListId(v string) *ModifyPrefixListRequest
	GetPrefixListId() *string
	SetPrefixListName(v string) *ModifyPrefixListRequest
	GetPrefixListName() *string
	SetRegionId(v string) *ModifyPrefixListRequest
	GetRegionId() *string
	SetRemoveEntry(v []*ModifyPrefixListRequestRemoveEntry) *ModifyPrefixListRequest
	GetRemoveEntry() []*ModifyPrefixListRequestRemoveEntry
	SetResourceOwnerAccount(v string) *ModifyPrefixListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPrefixListRequest
	GetResourceOwnerId() *int64
}

type ModifyPrefixListRequest struct {
	// The entries to be added to the prefix list.
	AddEntry []*ModifyPrefixListRequestAddEntry `json:"AddEntry,omitempty" xml:"AddEntry,omitempty" type:"Repeated"`
	// The description of the prefix list. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// This is description.
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the prefix list.
	//
	// This parameter is required.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The name of the prefix list. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://`, `https://`, `com.aliyun`, or `com.alibabacloud`. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// PrefixListNameSample
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
	// The region ID of the prefix list. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The entries to be deleted from the prefix list.
	RemoveEntry          []*ModifyPrefixListRequestRemoveEntry `json:"RemoveEntry,omitempty" xml:"RemoveEntry,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                               `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyPrefixListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrefixListRequest) GoString() string {
	return s.String()
}

func (s *ModifyPrefixListRequest) GetAddEntry() []*ModifyPrefixListRequestAddEntry {
	return s.AddEntry
}

func (s *ModifyPrefixListRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyPrefixListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPrefixListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPrefixListRequest) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *ModifyPrefixListRequest) GetPrefixListName() *string {
	return s.PrefixListName
}

func (s *ModifyPrefixListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPrefixListRequest) GetRemoveEntry() []*ModifyPrefixListRequestRemoveEntry {
	return s.RemoveEntry
}

func (s *ModifyPrefixListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPrefixListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPrefixListRequest) SetAddEntry(v []*ModifyPrefixListRequestAddEntry) *ModifyPrefixListRequest {
	s.AddEntry = v
	return s
}

func (s *ModifyPrefixListRequest) SetDescription(v string) *ModifyPrefixListRequest {
	s.Description = &v
	return s
}

func (s *ModifyPrefixListRequest) SetOwnerAccount(v string) *ModifyPrefixListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPrefixListRequest) SetOwnerId(v int64) *ModifyPrefixListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPrefixListRequest) SetPrefixListId(v string) *ModifyPrefixListRequest {
	s.PrefixListId = &v
	return s
}

func (s *ModifyPrefixListRequest) SetPrefixListName(v string) *ModifyPrefixListRequest {
	s.PrefixListName = &v
	return s
}

func (s *ModifyPrefixListRequest) SetRegionId(v string) *ModifyPrefixListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPrefixListRequest) SetRemoveEntry(v []*ModifyPrefixListRequestRemoveEntry) *ModifyPrefixListRequest {
	s.RemoveEntry = v
	return s
}

func (s *ModifyPrefixListRequest) SetResourceOwnerAccount(v string) *ModifyPrefixListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPrefixListRequest) SetResourceOwnerId(v int64) *ModifyPrefixListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPrefixListRequest) Validate() error {
	if s.AddEntry != nil {
		for _, item := range s.AddEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.RemoveEntry != nil {
		for _, item := range s.RemoveEntry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifyPrefixListRequestAddEntry struct {
	// The CIDR block in entry N to be added to the prefix list. Valid values of N: 0 to 200.
	//
	// Take note of the following items when you add the entries:
	//
	// 	- The total number of entries in the prefix list cannot exceed the maximum number of entries you specified for the prefix list. You can call the [DescribePrefixListAttributes](https://help.aliyun.com/document_detail/205872.html) operation to query the maximum number of entries that the prefix list can contain.
	//
	// 	- You cannot specify duplicate CIDR blocks.
	//
	// 	- The CIDR blocks cannot be the same as the `RemoveEntry.N.Cidr` values.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.2.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
	// The description in entry N. The description must be 2 to 32 characters in length and cannot start with `http://` or `https://`. Valid values of N: 0 to 200.
	//
	// example:
	//
	// Description Sample 01
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
}

func (s ModifyPrefixListRequestAddEntry) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrefixListRequestAddEntry) GoString() string {
	return s.String()
}

func (s *ModifyPrefixListRequestAddEntry) GetCidr() *string {
	return s.Cidr
}

func (s *ModifyPrefixListRequestAddEntry) GetDescription() *string {
	return s.Description
}

func (s *ModifyPrefixListRequestAddEntry) SetCidr(v string) *ModifyPrefixListRequestAddEntry {
	s.Cidr = &v
	return s
}

func (s *ModifyPrefixListRequestAddEntry) SetDescription(v string) *ModifyPrefixListRequestAddEntry {
	s.Description = &v
	return s
}

func (s *ModifyPrefixListRequestAddEntry) Validate() error {
	return dara.Validate(s)
}

type ModifyPrefixListRequestRemoveEntry struct {
	// The CIDR block in entry N to be deleted from the prefix list. Valid values of N: 0 to 200.
	//
	// Take note of the following items when you delete the entries:
	//
	// 	- You cannot specify duplicate CIDR blocks.
	//
	// 	- The CIDR blocks cannot be the same as the `AddEntry.N.Cidr` values.
	//
	// This parameter is required.
	//
	// example:
	//
	// 192.168.1.0/24
	Cidr *string `json:"Cidr,omitempty" xml:"Cidr,omitempty"`
}

func (s ModifyPrefixListRequestRemoveEntry) String() string {
	return dara.Prettify(s)
}

func (s ModifyPrefixListRequestRemoveEntry) GoString() string {
	return s.String()
}

func (s *ModifyPrefixListRequestRemoveEntry) GetCidr() *string {
	return s.Cidr
}

func (s *ModifyPrefixListRequestRemoveEntry) SetCidr(v string) *ModifyPrefixListRequestRemoveEntry {
	s.Cidr = &v
	return s
}

func (s *ModifyPrefixListRequestRemoveEntry) Validate() error {
	return dara.Validate(s)
}
