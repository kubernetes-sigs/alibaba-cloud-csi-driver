// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePortRangeListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *DescribePortRangeListsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribePortRangeListsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribePortRangeListsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribePortRangeListsRequest
	GetOwnerId() *int64
	SetPortRangeListId(v []*string) *DescribePortRangeListsRequest
	GetPortRangeListId() []*string
	SetPortRangeListName(v string) *DescribePortRangeListsRequest
	GetPortRangeListName() *string
	SetRegionId(v string) *DescribePortRangeListsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribePortRangeListsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribePortRangeListsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribePortRangeListsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribePortRangeListsRequestTag) *DescribePortRangeListsRequest
	GetTag() []*DescribePortRangeListsRequestTag
}

type DescribePortRangeListsRequest struct {
	// The maximum number of entries per page.
	//
	// 	- Maximum value: 100
	//
	// 	- Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// 727d41872117f2816343eeb432fbc5bfd21dc824589d2a4be0b5e8707e68181f
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the port list. Valid values of N: 0 to 100.
	PortRangeListId []*string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty" type:"Repeated"`
	// The name of the port list. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with http://, https://, com.aliyun, or com.alibabacloud. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// PortRangeListNameSample
	PortRangeListName *string `json:"PortRangeListName,omitempty" xml:"PortRangeListName,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/2679950.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. If you specify this parameter to query resources, up to 1,000 resources that belong to the specified resource group can be returned in the response. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/2716558.html) operation to query the most recent resource group list.
	//
	// >  A default resource group is not supported.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags that are added to the port list.
	Tag []*DescribePortRangeListsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribePortRangeListsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListsRequest) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribePortRangeListsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribePortRangeListsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribePortRangeListsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribePortRangeListsRequest) GetPortRangeListId() []*string {
	return s.PortRangeListId
}

func (s *DescribePortRangeListsRequest) GetPortRangeListName() *string {
	return s.PortRangeListName
}

func (s *DescribePortRangeListsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribePortRangeListsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribePortRangeListsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribePortRangeListsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribePortRangeListsRequest) GetTag() []*DescribePortRangeListsRequestTag {
	return s.Tag
}

func (s *DescribePortRangeListsRequest) SetMaxResults(v int32) *DescribePortRangeListsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetNextToken(v string) *DescribePortRangeListsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetOwnerAccount(v string) *DescribePortRangeListsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetOwnerId(v int64) *DescribePortRangeListsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetPortRangeListId(v []*string) *DescribePortRangeListsRequest {
	s.PortRangeListId = v
	return s
}

func (s *DescribePortRangeListsRequest) SetPortRangeListName(v string) *DescribePortRangeListsRequest {
	s.PortRangeListName = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetRegionId(v string) *DescribePortRangeListsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetResourceGroupId(v string) *DescribePortRangeListsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetResourceOwnerAccount(v string) *DescribePortRangeListsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetResourceOwnerId(v int64) *DescribePortRangeListsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribePortRangeListsRequest) SetTag(v []*DescribePortRangeListsRequestTag) *DescribePortRangeListsRequest {
	s.Tag = v
	return s
}

func (s *DescribePortRangeListsRequest) Validate() error {
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

type DescribePortRangeListsRequestTag struct {
	// The key of tag N. Valid values: 1 to 20.
	//
	// If you specify a single tag to query resources, up to 1,000 resources to which the tag is added are returned. If you specify multiple tags to query resources, up to 1,000 resources to which all specified tags are added are returned. To query more than 1,000 resources that have specified tags added, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// example:
	//
	// key for PortRangeList
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N.
	//
	// example:
	//
	// value for PortRangeList
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribePortRangeListsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribePortRangeListsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribePortRangeListsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribePortRangeListsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribePortRangeListsRequestTag) SetKey(v string) *DescribePortRangeListsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribePortRangeListsRequestTag) SetValue(v string) *DescribePortRangeListsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribePortRangeListsRequestTag) Validate() error {
	return dara.Validate(s)
}
