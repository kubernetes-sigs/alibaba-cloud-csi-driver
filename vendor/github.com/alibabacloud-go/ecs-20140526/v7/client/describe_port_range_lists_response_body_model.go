// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRangeListsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribePortRangeListsResponseBody
	GetNextToken() *string
	SetPortRangeLists(v []*DescribePortRangeListsResponseBodyPortRangeLists) *DescribePortRangeListsResponseBody
	GetPortRangeLists() []*DescribePortRangeListsResponseBodyPortRangeLists
	SetRequestId(v string) *DescribePortRangeListsResponseBody
	GetRequestId() *string
}

type DescribePortRangeListsResponseBody struct {
	// A pagination token. If the return value is empty, no more data is returned.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Details of the port lists.
	PortRangeLists []*DescribePortRangeListsResponseBodyPortRangeLists `json:"PortRangeLists,omitempty" xml:"PortRangeLists,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 6040AD98-11C3-5F78-9346-FCA8E9D8960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePortRangeListsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePortRangeListsResponseBody) GetPortRangeLists() []*DescribePortRangeListsResponseBodyPortRangeLists {
	return s.PortRangeLists
}

func (s *DescribePortRangeListsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePortRangeListsResponseBody) SetNextToken(v string) *DescribePortRangeListsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribePortRangeListsResponseBody) SetPortRangeLists(v []*DescribePortRangeListsResponseBodyPortRangeLists) *DescribePortRangeListsResponseBody {
	s.PortRangeLists = v
	return s
}

func (s *DescribePortRangeListsResponseBody) SetRequestId(v string) *DescribePortRangeListsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePortRangeListsResponseBody) Validate() error {
	if s.PortRangeLists != nil {
		for _, item := range s.PortRangeLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePortRangeListsResponseBodyPortRangeLists struct {
	// The number of associated resources.
	//
	// example:
	//
	// 1
	AssociationCount *int32 `json:"AssociationCount,omitempty" xml:"AssociationCount,omitempty"`
	// The time when the port list was created.
	//
	// example:
	//
	// 2024-12-04T07:11Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the port list.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of entries in the port list.
	//
	// example:
	//
	// 20
	MaxEntries *int32 `json:"MaxEntries,omitempty" xml:"MaxEntries,omitempty"`
	// The ID of the port list.
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The name of the port list.
	//
	// example:
	//
	// PortRangeListNameSample
	PortRangeListName *string `json:"PortRangeListName,omitempty" xml:"PortRangeListName,omitempty"`
	// The ID of the resource group to which to assign the port list.
	//
	// example:
	//
	// rg-2zeg82g****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The tags of the port list.
	Tags []*DescribePortRangeListsResponseBodyPortRangeListsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s DescribePortRangeListsResponseBodyPortRangeLists) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListsResponseBodyPortRangeLists) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) GetAssociationCount() *int32 {
	return s.AssociationCount
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) GetDescription() *string {
	return s.Description
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) GetMaxEntries() *int32 {
	return s.MaxEntries
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) GetPortRangeListName() *string {
	return s.PortRangeListName
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) GetTags() []*DescribePortRangeListsResponseBodyPortRangeListsTags {
	return s.Tags
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) SetAssociationCount(v int32) *DescribePortRangeListsResponseBodyPortRangeLists {
	s.AssociationCount = &v
	return s
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) SetCreationTime(v string) *DescribePortRangeListsResponseBodyPortRangeLists {
	s.CreationTime = &v
	return s
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) SetDescription(v string) *DescribePortRangeListsResponseBodyPortRangeLists {
	s.Description = &v
	return s
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) SetMaxEntries(v int32) *DescribePortRangeListsResponseBodyPortRangeLists {
	s.MaxEntries = &v
	return s
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) SetPortRangeListId(v string) *DescribePortRangeListsResponseBodyPortRangeLists {
	s.PortRangeListId = &v
	return s
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) SetPortRangeListName(v string) *DescribePortRangeListsResponseBodyPortRangeLists {
	s.PortRangeListName = &v
	return s
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) SetResourceGroupId(v string) *DescribePortRangeListsResponseBodyPortRangeLists {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) SetTags(v []*DescribePortRangeListsResponseBodyPortRangeListsTags) *DescribePortRangeListsResponseBodyPortRangeLists {
	s.Tags = v
	return s
}

func (s *DescribePortRangeListsResponseBodyPortRangeLists) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePortRangeListsResponseBodyPortRangeListsTags struct {
	// The key of tag N.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribePortRangeListsResponseBodyPortRangeListsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListsResponseBodyPortRangeListsTags) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListsResponseBodyPortRangeListsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribePortRangeListsResponseBodyPortRangeListsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribePortRangeListsResponseBodyPortRangeListsTags) SetTagKey(v string) *DescribePortRangeListsResponseBodyPortRangeListsTags {
	s.TagKey = &v
	return s
}

func (s *DescribePortRangeListsResponseBodyPortRangeListsTags) SetTagValue(v string) *DescribePortRangeListsResponseBodyPortRangeListsTags {
	s.TagValue = &v
	return s
}

func (s *DescribePortRangeListsResponseBodyPortRangeListsTags) Validate() error {
	return dara.Validate(s)
}
