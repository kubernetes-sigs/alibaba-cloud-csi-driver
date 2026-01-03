// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyPortRangeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddEntry(v []*ModifyPortRangeListRequestAddEntry) *ModifyPortRangeListRequest
	GetAddEntry() []*ModifyPortRangeListRequestAddEntry
	SetClientToken(v string) *ModifyPortRangeListRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyPortRangeListRequest
	GetDescription() *string
	SetOwnerAccount(v string) *ModifyPortRangeListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyPortRangeListRequest
	GetOwnerId() *int64
	SetPortRangeListId(v string) *ModifyPortRangeListRequest
	GetPortRangeListId() *string
	SetPortRangeListName(v string) *ModifyPortRangeListRequest
	GetPortRangeListName() *string
	SetRegionId(v string) *ModifyPortRangeListRequest
	GetRegionId() *string
	SetRemoveEntry(v []*ModifyPortRangeListRequestRemoveEntry) *ModifyPortRangeListRequest
	GetRemoveEntry() []*ModifyPortRangeListRequestRemoveEntry
	SetResourceOwnerAccount(v string) *ModifyPortRangeListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyPortRangeListRequest
	GetResourceOwnerId() *int64
}

type ModifyPortRangeListRequest struct {
	// The entries that you want to add or modify for the port list.
	AddEntry []*ModifyPortRangeListRequestAddEntry `json:"AddEntry,omitempty" xml:"AddEntry,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the port list. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// This is description.
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the port list.
	//
	// This parameter is required.
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The name of the port list. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http://, https://, com.aliyun, or com.alibabacloud. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// PortRangeListNameSample
	PortRangeListName *string `json:"PortRangeListName,omitempty" xml:"PortRangeListName,omitempty"`
	// The region ID of the port list. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The entries that you want to remove from the port list.
	RemoveEntry          []*ModifyPortRangeListRequestRemoveEntry `json:"RemoveEntry,omitempty" xml:"RemoveEntry,omitempty" type:"Repeated"`
	ResourceOwnerAccount *string                                  `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64                                   `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyPortRangeListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyPortRangeListRequest) GoString() string {
	return s.String()
}

func (s *ModifyPortRangeListRequest) GetAddEntry() []*ModifyPortRangeListRequestAddEntry {
	return s.AddEntry
}

func (s *ModifyPortRangeListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyPortRangeListRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyPortRangeListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyPortRangeListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyPortRangeListRequest) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *ModifyPortRangeListRequest) GetPortRangeListName() *string {
	return s.PortRangeListName
}

func (s *ModifyPortRangeListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyPortRangeListRequest) GetRemoveEntry() []*ModifyPortRangeListRequestRemoveEntry {
	return s.RemoveEntry
}

func (s *ModifyPortRangeListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyPortRangeListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyPortRangeListRequest) SetAddEntry(v []*ModifyPortRangeListRequestAddEntry) *ModifyPortRangeListRequest {
	s.AddEntry = v
	return s
}

func (s *ModifyPortRangeListRequest) SetClientToken(v string) *ModifyPortRangeListRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyPortRangeListRequest) SetDescription(v string) *ModifyPortRangeListRequest {
	s.Description = &v
	return s
}

func (s *ModifyPortRangeListRequest) SetOwnerAccount(v string) *ModifyPortRangeListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyPortRangeListRequest) SetOwnerId(v int64) *ModifyPortRangeListRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyPortRangeListRequest) SetPortRangeListId(v string) *ModifyPortRangeListRequest {
	s.PortRangeListId = &v
	return s
}

func (s *ModifyPortRangeListRequest) SetPortRangeListName(v string) *ModifyPortRangeListRequest {
	s.PortRangeListName = &v
	return s
}

func (s *ModifyPortRangeListRequest) SetRegionId(v string) *ModifyPortRangeListRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyPortRangeListRequest) SetRemoveEntry(v []*ModifyPortRangeListRequestRemoveEntry) *ModifyPortRangeListRequest {
	s.RemoveEntry = v
	return s
}

func (s *ModifyPortRangeListRequest) SetResourceOwnerAccount(v string) *ModifyPortRangeListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyPortRangeListRequest) SetResourceOwnerId(v int64) *ModifyPortRangeListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyPortRangeListRequest) Validate() error {
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

type ModifyPortRangeListRequestAddEntry struct {
	// The description of the port range in entry N. The description must be 2 to 32 characters in length and cannot start with http:// or https://. Valid values of N: 0 to 200.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port range in entry N. Valid values of N: 0 to 200. Take note of the following limits:
	//
	// 	- The total number of entries in the port list cannot exceed the `MaxEntries` value.
	//
	// 	- `PortRange` in different entries cannot be duplicated.
	//
	// 	- The value of this parameter cannot be the same as the value of `RemoveEntry.N.PortRange`.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
}

func (s ModifyPortRangeListRequestAddEntry) String() string {
	return dara.Prettify(s)
}

func (s ModifyPortRangeListRequestAddEntry) GoString() string {
	return s.String()
}

func (s *ModifyPortRangeListRequestAddEntry) GetDescription() *string {
	return s.Description
}

func (s *ModifyPortRangeListRequestAddEntry) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifyPortRangeListRequestAddEntry) SetDescription(v string) *ModifyPortRangeListRequestAddEntry {
	s.Description = &v
	return s
}

func (s *ModifyPortRangeListRequestAddEntry) SetPortRange(v string) *ModifyPortRangeListRequestAddEntry {
	s.PortRange = &v
	return s
}

func (s *ModifyPortRangeListRequestAddEntry) Validate() error {
	return dara.Validate(s)
}

type ModifyPortRangeListRequestRemoveEntry struct {
	// The port range in entry N. Valid values of N: 0 to 200. Take note of the following limits:
	//
	// 	- `PortRange` in different entries cannot be duplicated.
	//
	// 	- The value of this parameter cannot be the same as the value of `AddEntry.N.PortRange`.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
}

func (s ModifyPortRangeListRequestRemoveEntry) String() string {
	return dara.Prettify(s)
}

func (s ModifyPortRangeListRequestRemoveEntry) GoString() string {
	return s.String()
}

func (s *ModifyPortRangeListRequestRemoveEntry) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifyPortRangeListRequestRemoveEntry) SetPortRange(v string) *ModifyPortRangeListRequestRemoveEntry {
	s.PortRange = &v
	return s
}

func (s *ModifyPortRangeListRequestRemoveEntry) Validate() error {
	return dara.Validate(s)
}
