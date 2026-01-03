// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePrefixListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribePrefixListsResponseBody
	GetNextToken() *string
	SetPrefixLists(v *DescribePrefixListsResponseBodyPrefixLists) *DescribePrefixListsResponseBody
	GetPrefixLists() *DescribePrefixListsResponseBodyPrefixLists
	SetRequestId(v string) *DescribePrefixListsResponseBody
	GetRequestId() *string
}

type DescribePrefixListsResponseBody struct {
	// The query token that is returned in this call. If the return value is empty, no more data is returned.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Details about the prefix lists.
	PrefixLists *DescribePrefixListsResponseBodyPrefixLists `json:"PrefixLists,omitempty" xml:"PrefixLists,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 38793DB8-A4B2-4AEC-BFD3-111234E9188D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePrefixListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePrefixListsResponseBody) GetPrefixLists() *DescribePrefixListsResponseBodyPrefixLists {
	return s.PrefixLists
}

func (s *DescribePrefixListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePrefixListsResponseBody) SetNextToken(v string) *DescribePrefixListsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePrefixListsResponseBody) SetPrefixLists(v *DescribePrefixListsResponseBodyPrefixLists) *DescribePrefixListsResponseBody {
	s.PrefixLists = v
	return s
}

func (s *DescribePrefixListsResponseBody) SetRequestId(v string) *DescribePrefixListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePrefixListsResponseBody) Validate() error {
	if s.PrefixLists != nil {
		if err := s.PrefixLists.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePrefixListsResponseBodyPrefixLists struct {
	PrefixList []*DescribePrefixListsResponseBodyPrefixListsPrefixList `json:"PrefixList,omitempty" xml:"PrefixList,omitempty" type:"Repeated"`
}

func (s DescribePrefixListsResponseBodyPrefixLists) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsResponseBodyPrefixLists) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponseBodyPrefixLists) GetPrefixList() []*DescribePrefixListsResponseBodyPrefixListsPrefixList {
	return s.PrefixList
}

func (s *DescribePrefixListsResponseBodyPrefixLists) SetPrefixList(v []*DescribePrefixListsResponseBodyPrefixListsPrefixList) *DescribePrefixListsResponseBodyPrefixLists {
	s.PrefixList = v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixLists) Validate() error {
	if s.PrefixList != nil {
		for _, item := range s.PrefixList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePrefixListsResponseBodyPrefixListsPrefixList struct {
	// The IP address family of the prefix list. Valid values:
	//
	// 	- IPv4
	//
	// 	- IPv6
	//
	// example:
	//
	// IPv4
	AddressFamily *string `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	// The number of associated resources.
	//
	// example:
	//
	// 1
	AssociationCount *int32 `json:"AssociationCount,omitempty" xml:"AssociationCount,omitempty"`
	// The time when the prefix list was created.
	//
	// example:
	//
	// 2021-02-20T07:11Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the prefix list.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of entries that the prefix list can contain.
	//
	// example:
	//
	// 20
	MaxEntries *int32 `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	// The ID of the prefix list.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	PrefixListId *string `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	// The name of the prefix list.
	//
	// example:
	//
	// PrefixListNameSample
	PrefixListName *string `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
	// The ID of the resource group to which the prefix list belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the prefix list.
	Tags *DescribePrefixListsResponseBodyPrefixListsPrefixListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribePrefixListsResponseBodyPrefixListsPrefixList) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsResponseBodyPrefixListsPrefixList) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) GetAddressFamily() *string {
	return s.AddressFamily
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) GetAssociationCount() *int32 {
	return s.AssociationCount
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) GetDescription() *string {
	return s.Description
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) GetMaxEntries() *int32 {
	return s.MaxEntries
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) GetPrefixListId() *string {
	return s.PrefixListId
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) GetPrefixListName() *string {
	return s.PrefixListName
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) GetTags() *DescribePrefixListsResponseBodyPrefixListsPrefixListTags {
	return s.Tags
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) SetAddressFamily(v string) *DescribePrefixListsResponseBodyPrefixListsPrefixList {
	s.AddressFamily = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) SetAssociationCount(v int32) *DescribePrefixListsResponseBodyPrefixListsPrefixList {
	s.AssociationCount = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) SetCreationTime(v string) *DescribePrefixListsResponseBodyPrefixListsPrefixList {
	s.CreationTime = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) SetDescription(v string) *DescribePrefixListsResponseBodyPrefixListsPrefixList {
	s.Description = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) SetMaxEntries(v int32) *DescribePrefixListsResponseBodyPrefixListsPrefixList {
	s.MaxEntries = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) SetPrefixListId(v string) *DescribePrefixListsResponseBodyPrefixListsPrefixList {
	s.PrefixListId = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) SetPrefixListName(v string) *DescribePrefixListsResponseBodyPrefixListsPrefixList {
	s.PrefixListName = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) SetResourceGroupId(v string) *DescribePrefixListsResponseBodyPrefixListsPrefixList {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) SetTags(v *DescribePrefixListsResponseBodyPrefixListsPrefixListTags) *DescribePrefixListsResponseBodyPrefixListsPrefixList {
	s.Tags = v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixList) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribePrefixListsResponseBodyPrefixListsPrefixListTags struct {
	Tag []*DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribePrefixListsResponseBodyPrefixListsPrefixListTags) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsResponseBodyPrefixListsPrefixListTags) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixListTags) GetTag() []*DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag {
	return s.Tag
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixListTags) SetTag(v []*DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag) *DescribePrefixListsResponseBodyPrefixListsPrefixListTags {
	s.Tag = v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixListTags) Validate() error {
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

type DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag struct {
	// The tag value. A prefix list can have 1 to 20 tags. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http:// or https://`.
	//
	// example:
	//
	// TestValue
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag key. A prefix list can have 1 to 20 tags. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag) GoString() string {
	return s.String()
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag) SetTagKey(v string) *DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag) SetTagValue(v string) *DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribePrefixListsResponseBodyPrefixListsPrefixListTagsTag) Validate() error {
	return dara.Validate(s)
}
