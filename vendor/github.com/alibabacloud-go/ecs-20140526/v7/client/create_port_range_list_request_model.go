// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePortRangeListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreatePortRangeListRequest
	GetClientToken() *string
	SetDescription(v string) *CreatePortRangeListRequest
	GetDescription() *string
	SetEntry(v []*CreatePortRangeListRequestEntry) *CreatePortRangeListRequest
	GetEntry() []*CreatePortRangeListRequestEntry
	SetMaxEntries(v int32) *CreatePortRangeListRequest
	GetMaxEntries() *int32
	SetOwnerAccount(v string) *CreatePortRangeListRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreatePortRangeListRequest
	GetOwnerId() *int64
	SetPortRangeListName(v string) *CreatePortRangeListRequest
	GetPortRangeListName() *string
	SetRegionId(v string) *CreatePortRangeListRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreatePortRangeListRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreatePortRangeListRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreatePortRangeListRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreatePortRangeListRequestTag) *CreatePortRangeListRequest
	GetTag() []*CreatePortRangeListRequestTag
}

type CreatePortRangeListRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the port list. The description must be 2 to 256 characters in length and cannot start with http:// or https://.
	//
	// example:
	//
	// Description information of PortRangeList
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The port list entries.
	Entry []*CreatePortRangeListRequestEntry `json:"Entry,omitempty" xml:"Entry,omitempty" type:"Repeated"`
	// The maximum number of entries in the port list. The value cannot be changed after you create the port list. Valid values: 1 to 2000.
	//
	// >  When you reference a port list in a resource, such as a security group, the maximum number of entries (instead of the actual number of entries) in the port list counts against the rule quota for the resource. Set a proper value for MaxEntries.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	MaxEntries   *int32  `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The name of the port list. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http://, https://, com.aliyun, or com.alibabacloud. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// This parameter is required.
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
	// The ID of the resource group to which the port list belongs.
	//
	// example:
	//
	// rg-aek3b6jzp66****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags to add to the port list. You can add 0 to 20 tags to the port list.
	Tag []*CreatePortRangeListRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreatePortRangeListRequest) String() string {
	return dara.Prettify(s)
}

func (s CreatePortRangeListRequest) GoString() string {
	return s.String()
}

func (s *CreatePortRangeListRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreatePortRangeListRequest) GetDescription() *string {
	return s.Description
}

func (s *CreatePortRangeListRequest) GetEntry() []*CreatePortRangeListRequestEntry {
	return s.Entry
}

func (s *CreatePortRangeListRequest) GetMaxEntries() *int32 {
	return s.MaxEntries
}

func (s *CreatePortRangeListRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreatePortRangeListRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreatePortRangeListRequest) GetPortRangeListName() *string {
	return s.PortRangeListName
}

func (s *CreatePortRangeListRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreatePortRangeListRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreatePortRangeListRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreatePortRangeListRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreatePortRangeListRequest) GetTag() []*CreatePortRangeListRequestTag {
	return s.Tag
}

func (s *CreatePortRangeListRequest) SetClientToken(v string) *CreatePortRangeListRequest {
	s.ClientToken = &v
	return s
}

func (s *CreatePortRangeListRequest) SetDescription(v string) *CreatePortRangeListRequest {
	s.Description = &v
	return s
}

func (s *CreatePortRangeListRequest) SetEntry(v []*CreatePortRangeListRequestEntry) *CreatePortRangeListRequest {
	s.Entry = v
	return s
}

func (s *CreatePortRangeListRequest) SetMaxEntries(v int32) *CreatePortRangeListRequest {
	s.MaxEntries = &v
	return s
}

func (s *CreatePortRangeListRequest) SetOwnerAccount(v string) *CreatePortRangeListRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreatePortRangeListRequest) SetOwnerId(v int64) *CreatePortRangeListRequest {
	s.OwnerId = &v
	return s
}

func (s *CreatePortRangeListRequest) SetPortRangeListName(v string) *CreatePortRangeListRequest {
	s.PortRangeListName = &v
	return s
}

func (s *CreatePortRangeListRequest) SetRegionId(v string) *CreatePortRangeListRequest {
	s.RegionId = &v
	return s
}

func (s *CreatePortRangeListRequest) SetResourceGroupId(v string) *CreatePortRangeListRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreatePortRangeListRequest) SetResourceOwnerAccount(v string) *CreatePortRangeListRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreatePortRangeListRequest) SetResourceOwnerId(v int64) *CreatePortRangeListRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreatePortRangeListRequest) SetTag(v []*CreatePortRangeListRequestTag) *CreatePortRangeListRequest {
	s.Tag = v
	return s
}

func (s *CreatePortRangeListRequest) Validate() error {
	if s.Entry != nil {
		for _, item := range s.Entry {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreatePortRangeListRequestEntry struct {
	// The description of port range N. The description must be 2 to 32 characters in length and cannot start with http:// or https://. Valid values of N: 0 to 200.
	//
	// example:
	//
	// Description information of Entry
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Port range N. Valid values of N: 0 to 200.
	//
	// 	- The total number of entries cannot exceed the `MaxEntries` value.
	//
	// 	- `PortRange` in multiple entries cannot be duplicated.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
}

func (s CreatePortRangeListRequestEntry) String() string {
	return dara.Prettify(s)
}

func (s CreatePortRangeListRequestEntry) GoString() string {
	return s.String()
}

func (s *CreatePortRangeListRequestEntry) GetDescription() *string {
	return s.Description
}

func (s *CreatePortRangeListRequestEntry) GetPortRange() *string {
	return s.PortRange
}

func (s *CreatePortRangeListRequestEntry) SetDescription(v string) *CreatePortRangeListRequestEntry {
	s.Description = &v
	return s
}

func (s *CreatePortRangeListRequestEntry) SetPortRange(v string) *CreatePortRangeListRequestEntry {
	s.PortRange = &v
	return s
}

func (s *CreatePortRangeListRequestEntry) Validate() error {
	return dara.Validate(s)
}

type CreatePortRangeListRequestTag struct {
	// The key of tag N to add to the port list.
	//
	// The tag key cannot be empty or an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// key for PortRangeList
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the port list.
	//
	// The tag value cannot be empty but can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// value for PortRangeList
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreatePortRangeListRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreatePortRangeListRequestTag) GoString() string {
	return s.String()
}

func (s *CreatePortRangeListRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreatePortRangeListRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreatePortRangeListRequestTag) SetKey(v string) *CreatePortRangeListRequestTag {
	s.Key = &v
	return s
}

func (s *CreatePortRangeListRequestTag) SetValue(v string) *CreatePortRangeListRequestTag {
	s.Value = &v
	return s
}

func (s *CreatePortRangeListRequestTag) Validate() error {
	return dara.Validate(s)
}
