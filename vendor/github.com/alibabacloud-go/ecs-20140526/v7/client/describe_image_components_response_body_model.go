// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageComponentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageComponent(v *DescribeImageComponentsResponseBodyImageComponent) *DescribeImageComponentsResponseBody
	GetImageComponent() *DescribeImageComponentsResponseBodyImageComponent
	SetMaxResults(v int32) *DescribeImageComponentsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeImageComponentsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeImageComponentsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeImageComponentsResponseBody
	GetTotalCount() *int32
}

type DescribeImageComponentsResponseBody struct {
	// The information about the image components.
	ImageComponent *DescribeImageComponentsResponseBodyImageComponent `json:"ImageComponent,omitempty" xml:"ImageComponent,omitempty" type:"Struct"`
	// The number of entries per page.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. For information about how to use the returned value, see the "Usage notes" section of this topic.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of image components returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeImageComponentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageComponentsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageComponentsResponseBody) GetImageComponent() *DescribeImageComponentsResponseBodyImageComponent {
	return s.ImageComponent
}

func (s *DescribeImageComponentsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeImageComponentsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImageComponentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageComponentsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeImageComponentsResponseBody) SetImageComponent(v *DescribeImageComponentsResponseBodyImageComponent) *DescribeImageComponentsResponseBody {
	s.ImageComponent = v
	return s
}

func (s *DescribeImageComponentsResponseBody) SetMaxResults(v int32) *DescribeImageComponentsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *DescribeImageComponentsResponseBody) SetNextToken(v string) *DescribeImageComponentsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeImageComponentsResponseBody) SetRequestId(v string) *DescribeImageComponentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageComponentsResponseBody) SetTotalCount(v int32) *DescribeImageComponentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeImageComponentsResponseBody) Validate() error {
	if s.ImageComponent != nil {
		if err := s.ImageComponent.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageComponentsResponseBodyImageComponent struct {
	ImageComponentSet []*DescribeImageComponentsResponseBodyImageComponentImageComponentSet `json:"ImageComponentSet,omitempty" xml:"ImageComponentSet,omitempty" type:"Repeated"`
}

func (s DescribeImageComponentsResponseBodyImageComponent) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageComponentsResponseBodyImageComponent) GoString() string {
	return s.String()
}

func (s *DescribeImageComponentsResponseBodyImageComponent) GetImageComponentSet() []*DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	return s.ImageComponentSet
}

func (s *DescribeImageComponentsResponseBodyImageComponent) SetImageComponentSet(v []*DescribeImageComponentsResponseBodyImageComponentImageComponentSet) *DescribeImageComponentsResponseBodyImageComponent {
	s.ImageComponentSet = v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponent) Validate() error {
	if s.ImageComponentSet != nil {
		for _, item := range s.ImageComponentSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageComponentsResponseBodyImageComponentImageComponentSet struct {
	// The type of the image component.
	//
	// example:
	//
	// Build
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The version number of the image component.
	//
	// example:
	//
	// null
	ComponentVersion *string `json:"ComponentVersion,omitempty" xml:"ComponentVersion,omitempty"`
	// The content of the image component.
	//
	// example:
	//
	// RESTART
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The time when the image component was created.
	//
	// example:
	//
	// 2020-11-24T06:00:00Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The description of the image component.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the image component.
	//
	// example:
	//
	// ic-bp67acfmxazb4p****
	ImageComponentId *string `json:"ImageComponentId,omitempty" xml:"ImageComponentId,omitempty"`
	// The name of the image component.
	//
	// example:
	//
	// testComponent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the image component. Valid values:
	//
	// 	- SELF: the custom component that you created.
	//
	// 	- ALIYUN: the system component provided by Alibaba Cloud.
	//
	// example:
	//
	// SELF
	Owner *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	// The parameters contained in the image component.
	Parameters *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Struct"`
	// The ID of the resource group to which the image component belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The type of the operating system supported by the image component.
	//
	// example:
	//
	// Linux
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
	// The tags of the image component.
	Tags *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
}

func (s DescribeImageComponentsResponseBodyImageComponentImageComponentSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GoString() string {
	return s.String()
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetComponentVersion() *string {
	return s.ComponentVersion
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetContent() *string {
	return s.Content
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetDescription() *string {
	return s.Description
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetImageComponentId() *string {
	return s.ImageComponentId
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetName() *string {
	return s.Name
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetOwner() *string {
	return s.Owner
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetParameters() *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParameters {
	return s.Parameters
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetSystemType() *string {
	return s.SystemType
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) GetTags() *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTags {
	return s.Tags
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetComponentType(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.ComponentType = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetComponentVersion(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.ComponentVersion = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetContent(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.Content = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetCreationTime(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.CreationTime = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetDescription(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.Description = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetImageComponentId(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.ImageComponentId = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetName(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.Name = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetOwner(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.Owner = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetParameters(v *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParameters) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.Parameters = v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetResourceGroupId(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetSystemType(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.SystemType = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) SetTags(v *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTags) *DescribeImageComponentsResponseBodyImageComponentImageComponentSet {
	s.Tags = v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSet) Validate() error {
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageComponentsResponseBodyImageComponentImageComponentSetParameters struct {
	Parameter []*DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter `json:"Parameter,omitempty" xml:"Parameter,omitempty" type:"Repeated"`
}

func (s DescribeImageComponentsResponseBodyImageComponentImageComponentSetParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageComponentsResponseBodyImageComponentImageComponentSetParameters) GoString() string {
	return s.String()
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParameters) GetParameter() []*DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter {
	return s.Parameter
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParameters) SetParameter(v []*DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter) *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParameters {
	s.Parameter = v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParameters) Validate() error {
	if s.Parameter != nil {
		for _, item := range s.Parameter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter struct {
	// The default value of the parameter.
	//
	// example:
	//
	// null
	DefaultValue *string `json:"DefaultValue,omitempty" xml:"DefaultValue,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// null
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The type of the parameter.
	//
	// Valid values:
	//
	// 	- String
	//
	// 	- Number
	//
	// 	- Boolean
	//
	// example:
	//
	// null
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter) GoString() string {
	return s.String()
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter) GetDefaultValue() *string {
	return s.DefaultValue
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter) GetName() *string {
	return s.Name
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter) GetType() *string {
	return s.Type
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter) SetDefaultValue(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter {
	s.DefaultValue = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter) SetName(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter {
	s.Name = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter) SetType(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter {
	s.Type = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetParametersParameter) Validate() error {
	return dara.Validate(s)
}

type DescribeImageComponentsResponseBodyImageComponentImageComponentSetTags struct {
	Tag []*DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImageComponentsResponseBodyImageComponentImageComponentSetTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageComponentsResponseBodyImageComponentImageComponentSetTags) GoString() string {
	return s.String()
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTags) GetTag() []*DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag {
	return s.Tag
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTags) SetTag(v []*DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag) *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTags {
	s.Tag = v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTags) Validate() error {
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

type DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag struct {
	// The key of the tag.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The value of the tag.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag) SetTagKey(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag) SetTagValue(v string) *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeImageComponentsResponseBodyImageComponentImageComponentSetTagsTag) Validate() error {
	return dara.Validate(s)
}
