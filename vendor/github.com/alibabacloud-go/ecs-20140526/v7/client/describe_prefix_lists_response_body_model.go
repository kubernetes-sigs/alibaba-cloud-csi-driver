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
	NextToken   *string                                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
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
	AddressFamily    *string                                                   `json:"AddressFamily,omitempty" xml:"AddressFamily,omitempty"`
	AssociationCount *int32                                                    `json:"AssociationCount,omitempty" xml:"AssociationCount,omitempty"`
	CreationTime     *string                                                   `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description      *string                                                   `json:"Description,omitempty" xml:"Description,omitempty"`
	MaxEntries       *int32                                                    `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	PrefixListId     *string                                                   `json:"PrefixListId,omitempty" xml:"PrefixListId,omitempty"`
	PrefixListName   *string                                                   `json:"PrefixListName,omitempty" xml:"PrefixListName,omitempty"`
	ResourceGroupId  *string                                                   `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags             *DescribePrefixListsResponseBodyPrefixListsPrefixListTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
