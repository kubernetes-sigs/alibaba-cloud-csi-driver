// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddressFamily(v string) *DescribePrefixListsRequest
	GetAddressFamily() *string
	SetMaxResults(v int32) *DescribePrefixListsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePrefixListsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribePrefixListsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePrefixListsRequest
	GetOwnerId() *int64
	SetPrefixListId(v []*string) *DescribePrefixListsRequest
	GetPrefixListId() []*string
	SetPrefixListName(v string) *DescribePrefixListsRequest
	GetPrefixListName() *string
	SetRegionId(v string) *DescribePrefixListsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribePrefixListsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribePrefixListsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePrefixListsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribePrefixListsRequestTag) *DescribePrefixListsRequest
	GetTag() []*DescribePrefixListsRequestTag
}

type DescribePrefixListsRequest struct {
	// The IP address family. Valid values:
	//
	// 	- IPv4
	//
	// 	- IPv6
	//
	// This parameter is empty by default, which indicates that all prefix lists are queried.
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results. Set the value to the `NextToken` value returned in the last call to this operation. Leave this parameter empty the first time you call this operation.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The IDs of prefix lists. Valid values of N: 0 to 100.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	PrefixListId []*string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty" type:"Repeated"`
	// The name of the prefix list.
	//
	// example:
	//
	// PrefixListNameSample
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-chengdu
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the prefix list belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the prefix list.
	Tag []*DescribePrefixListsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribePrefixListsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsRequest) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsRequest) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *DescribePrefixListsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePrefixListsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePrefixListsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePrefixListsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePrefixListsRequest) GetPrefixListId() []*string {
	return s.PrefixListId
}

func (s *DescribePrefixListsRequest) GetPrefixListName() *string {
	return s.PrefixListName
}

func (s *DescribePrefixListsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePrefixListsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePrefixListsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePrefixListsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePrefixListsRequest) GetTag() []*DescribePrefixListsRequestTag {
	return s.Tag
}

func (s *DescribePrefixListsRequest) SetAddressFamily(v string) *DescribePrefixListsRequest {
	s.AddressFamily = &v
	return s
}

func (s *DescribePrefixListsRequest) SetMaxResults(v int32) *DescribePrefixListsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePrefixListsRequest) SetNextToken(v string) *DescribePrefixListsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePrefixListsRequest) SetOwnerAccount(v string) *DescribePrefixListsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePrefixListsRequest) SetOwnerId(v int64) *DescribePrefixListsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePrefixListsRequest) SetPrefixListId(v []*string) *DescribePrefixListsRequest {
	s.PrefixListId = v
	return s
}

func (s *DescribePrefixListsRequest) SetPrefixListName(v string) *DescribePrefixListsRequest {
	s.PrefixListName = &v
	return s
}

func (s *DescribePrefixListsRequest) SetRegionId(v string) *DescribePrefixListsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePrefixListsRequest) SetResourceGroupId(v string) *DescribePrefixListsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePrefixListsRequest) SetResourceOwnerAccount(v string) *DescribePrefixListsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePrefixListsRequest) SetResourceOwnerId(v int64) *DescribePrefixListsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePrefixListsRequest) SetTag(v []*DescribePrefixListsRequestTag) *DescribePrefixListsRequest {
	s.Tag = v
	return s
}

func (s *DescribePrefixListsRequest) Validate() error {
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

type DescribePrefixListsRequestTag struct {
	// The key of tag N of the prefix list. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length. It cannot start with aliyun or acs: and cannot contain `http:// or https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the prefix list. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http:// or https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePrefixListsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribePrefixListsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribePrefixListsRequestTag) SetKey(v string) *DescribePrefixListsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribePrefixListsRequestTag) SetValue(v string) *DescribePrefixListsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribePrefixListsRequestTag) Validate() error {
	return dara.Validate(s)
}
