// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageComponentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComponentType(v string) *DescribeImageComponentsRequest
	GetComponentType() *string
	SetComponentVersion(v string) *DescribeImageComponentsRequest
	GetComponentVersion() *string
	SetImageComponentId(v []*string) *DescribeImageComponentsRequest
	GetImageComponentId() []*string
	SetMaxResults(v int32) *DescribeImageComponentsRequest
	GetMaxResults() *int32
	SetName(v string) *DescribeImageComponentsRequest
	GetName() *string
	SetNextToken(v string) *DescribeImageComponentsRequest
	GetNextToken() *string
	SetOwner(v string) *DescribeImageComponentsRequest
	GetOwner() *string
	SetOwnerAccount(v string) *DescribeImageComponentsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeImageComponentsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeImageComponentsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeImageComponentsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeImageComponentsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeImageComponentsRequest
	GetResourceOwnerId() *int64
	SetSystemType(v string) *DescribeImageComponentsRequest
	GetSystemType() *string
	SetTag(v []*DescribeImageComponentsRequestTag) *DescribeImageComponentsRequest
	GetTag() []*DescribeImageComponentsRequestTag
}

type DescribeImageComponentsRequest struct {
	// The type of the image component.
	//
	// Valid values:
	//
	// 	- Build
	//
	// 	- Test
	//
	// example:
	//
	// null
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The version number of the image component in the \\<major>.\\<minor>.\\<patch> format. You can set \\<major>, \\<minor>, and \\<patch> to non-negative integers, or set one of \\<major>, \\<minor>, and \\<patch> to the wildcard (\\*) and the other two to non-negative integers.
	//
	// >  This parameter takes effect only if you specify Name.
	//
	// example:
	//
	// null
	ComponentVersion *string `json:"ComponentVersion,omitempty" xml:"ComponentVersion,omitempty"`
	// The IDs of image components. Valid values of N: 1 to 20.
	//
	// example:
	//
	// ic-bp67acfmxazb4p****
	ImageComponentId []*string `json:"ImageComponentId,omitempty" xml:"ImageComponentId,omitempty" type:"Repeated"`
	// The maximum number of entries per page. Valid values: 1 to 500.
	//
	// Default value: 50.
	//
	// example:
	//
	// 50
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The name of the image component. You must specify an exact name to search for the image component.
	//
	// example:
	//
	// testComponent
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of `NextToken`.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The type of the image component. Valid values:
	//
	// 	- SELF: the custom component that you created.
	//
	// 	- ALIYUN: the system component provided by Alibaba Cloud.
	//
	// example:
	//
	// SELF
	Owner        *string `json:"Owner,omitempty" xml:"Owner,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the image component. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If this parameter is specified to query resources, up to 1,000 resources that belong to the specified resource group can be displayed in the response.
	//
	// >  Resources in the default resource group are displayed in the response regardless of how this parameter is set.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the operating system supported by the image component.
	//
	// Valid values:
	//
	// 	- Linux
	//
	// 	- Windows
	//
	// example:
	//
	// null
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
	// The tags of the image component.
	Tag []*DescribeImageComponentsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeImageComponentsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageComponentsRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageComponentsRequest) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeImageComponentsRequest) GetComponentVersion() *string {
	return s.ComponentVersion
}

func (s *DescribeImageComponentsRequest) GetImageComponentId() []*string {
	return s.ImageComponentId
}

func (s *DescribeImageComponentsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeImageComponentsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeImageComponentsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeImageComponentsRequest) GetOwner() *string {
	return s.Owner
}

func (s *DescribeImageComponentsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeImageComponentsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeImageComponentsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageComponentsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeImageComponentsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeImageComponentsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeImageComponentsRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *DescribeImageComponentsRequest) GetTag() []*DescribeImageComponentsRequestTag {
	return s.Tag
}

func (s *DescribeImageComponentsRequest) SetComponentType(v string) *DescribeImageComponentsRequest {
	s.ComponentType = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetComponentVersion(v string) *DescribeImageComponentsRequest {
	s.ComponentVersion = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetImageComponentId(v []*string) *DescribeImageComponentsRequest {
	s.ImageComponentId = v
	return s
}

func (s *DescribeImageComponentsRequest) SetMaxResults(v int32) *DescribeImageComponentsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetName(v string) *DescribeImageComponentsRequest {
	s.Name = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetNextToken(v string) *DescribeImageComponentsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetOwner(v string) *DescribeImageComponentsRequest {
	s.Owner = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetOwnerAccount(v string) *DescribeImageComponentsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetOwnerId(v int64) *DescribeImageComponentsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetRegionId(v string) *DescribeImageComponentsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetResourceGroupId(v string) *DescribeImageComponentsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetResourceOwnerAccount(v string) *DescribeImageComponentsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetResourceOwnerId(v int64) *DescribeImageComponentsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetSystemType(v string) *DescribeImageComponentsRequest {
	s.SystemType = &v
	return s
}

func (s *DescribeImageComponentsRequest) SetTag(v []*DescribeImageComponentsRequestTag) *DescribeImageComponentsRequest {
	s.Tag = v
	return s
}

func (s *DescribeImageComponentsRequest) Validate() error {
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

type DescribeImageComponentsRequestTag struct {
	// The key of tag N. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageComponentsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageComponentsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeImageComponentsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeImageComponentsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeImageComponentsRequestTag) SetKey(v string) *DescribeImageComponentsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeImageComponentsRequestTag) SetValue(v string) *DescribeImageComponentsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeImageComponentsRequestTag) Validate() error {
	return dara.Validate(s)
}
