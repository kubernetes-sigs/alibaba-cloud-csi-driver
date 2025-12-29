// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDryRun(v bool) *DescribeSecurityGroupsRequest
	GetDryRun() *bool
	SetFuzzyQuery(v bool) *DescribeSecurityGroupsRequest
	GetFuzzyQuery() *bool
	SetIsQueryEcsCount(v bool) *DescribeSecurityGroupsRequest
	GetIsQueryEcsCount() *bool
	SetMaxResults(v int32) *DescribeSecurityGroupsRequest
	GetMaxResults() *int32
	SetNetworkType(v string) *DescribeSecurityGroupsRequest
	GetNetworkType() *string
	SetNextToken(v string) *DescribeSecurityGroupsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeSecurityGroupsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSecurityGroupsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSecurityGroupsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSecurityGroupsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSecurityGroupsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSecurityGroupsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSecurityGroupsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSecurityGroupsRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *DescribeSecurityGroupsRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v string) *DescribeSecurityGroupsRequest
	GetSecurityGroupIds() *string
	SetSecurityGroupName(v string) *DescribeSecurityGroupsRequest
	GetSecurityGroupName() *string
	SetSecurityGroupType(v string) *DescribeSecurityGroupsRequest
	GetSecurityGroupType() *string
	SetServiceManaged(v bool) *DescribeSecurityGroupsRequest
	GetServiceManaged() *bool
	SetTag(v []*DescribeSecurityGroupsRequestTag) *DescribeSecurityGroupsRequest
	GetTag() []*DescribeSecurityGroupsRequestTag
	SetVpcId(v string) *DescribeSecurityGroupsRequest
	GetVpcId() *string
}

type DescribeSecurityGroupsRequest struct {
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks your AccessKey pair, the permissions of the RAM user, and the required parameters. If the request passes the dry run, the DryRunOperation error code is returned. Otherwise, an error message is returned.
	//
	// 	- false: performs a dry run and performs the actual request. If the request passes the dry run, a 2xx HTTP status code is returned and the operation is performed.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// > This parameter is deprecated.
	//
	// example:
	//
	// null
	FuzzyQuery *bool `json:"FuzzyQuery,omitempty" xml:"FuzzyQuery,omitempty"`
	// Specifies whether to query the capacity of the security group. If you set this parameter to True, the `EcsCount` and `AvailableInstanceAmount` values in the response are valid.
	//
	// >  This parameter is deprecated.
	//
	// example:
	//
	// null
	IsQueryEcsCount *bool `json:"IsQueryEcsCount,omitempty" xml:"IsQueryEcsCount,omitempty"`
	// The maximum number of entries per page. If you specify this parameter, both `MaxResults` and `NextToken` are used for a paged query.
	//
	// Maximum value: 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The network type of the security group. Valid values:
	//
	// 	- vpc
	//
	// 	- classic
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// e71d8a535bd9cc11
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the security group belongs. If this parameter is specified to query resources, up to 1,000 resources that belong to the specified resource group can be displayed in the response. You can call the [ListResourceGroups](https://help.aliyun.com/document_detail/158855.html) operation to query the most recent resource group list.
	//
	// > Resources in the default resource group are displayed in the response regardless of how this parameter is configured.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The security group IDs. Set this parameter to a JSON array that consists of up to 100 security group IDs. Separate the security group IDs with commas (,).
	//
	// example:
	//
	// ["sg-bp67acfmxazb4p****", "sg-bp67acfmxazb4p****", "sg-bp67acfmxazb4p****",....]
	SecurityGroupIds *string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty"`
	// The name of the security group.
	//
	// example:
	//
	// SGTestName
	SecurityGroupName *string `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	// The type of the security group. Valid values:
	//
	// 	- normal: basic security group
	//
	// 	- enterprise: advanced security group
	//
	// > If you do not specify this parameter, both basic and advanced security groups are queried.
	//
	// example:
	//
	// normal
	SecurityGroupType *string `json:"SecurityGroupType,omitempty" xml:"SecurityGroupType,omitempty"`
	// Specifies whether to query managed security groups. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	ServiceManaged *bool `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	// The tags to add to the security groups.
	Tag []*DescribeSecurityGroupsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the virtual private cloud (VPC) to which the security group belongs.
	//
	// example:
	//
	// vpc-bp67acfmxazb4p****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeSecurityGroupsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DescribeSecurityGroupsRequest) GetFuzzyQuery() *bool {
	return s.FuzzyQuery
}

func (s *DescribeSecurityGroupsRequest) GetIsQueryEcsCount() *bool {
	return s.IsQueryEcsCount
}

func (s *DescribeSecurityGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSecurityGroupsRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeSecurityGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSecurityGroupsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSecurityGroupsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSecurityGroupsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSecurityGroupsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSecurityGroupsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityGroupsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSecurityGroupsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSecurityGroupsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSecurityGroupsRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupsRequest) GetSecurityGroupIds() *string {
	return s.SecurityGroupIds
}

func (s *DescribeSecurityGroupsRequest) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeSecurityGroupsRequest) GetSecurityGroupType() *string {
	return s.SecurityGroupType
}

func (s *DescribeSecurityGroupsRequest) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeSecurityGroupsRequest) GetTag() []*DescribeSecurityGroupsRequestTag {
	return s.Tag
}

func (s *DescribeSecurityGroupsRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeSecurityGroupsRequest) SetDryRun(v bool) *DescribeSecurityGroupsRequest {
	s.DryRun = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetFuzzyQuery(v bool) *DescribeSecurityGroupsRequest {
	s.FuzzyQuery = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetIsQueryEcsCount(v bool) *DescribeSecurityGroupsRequest {
	s.IsQueryEcsCount = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetMaxResults(v int32) *DescribeSecurityGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetNetworkType(v string) *DescribeSecurityGroupsRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetNextToken(v string) *DescribeSecurityGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetOwnerAccount(v string) *DescribeSecurityGroupsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetOwnerId(v int64) *DescribeSecurityGroupsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetPageNumber(v int32) *DescribeSecurityGroupsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetPageSize(v int32) *DescribeSecurityGroupsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetRegionId(v string) *DescribeSecurityGroupsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetResourceGroupId(v string) *DescribeSecurityGroupsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetResourceOwnerAccount(v string) *DescribeSecurityGroupsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetResourceOwnerId(v int64) *DescribeSecurityGroupsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetSecurityGroupId(v string) *DescribeSecurityGroupsRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetSecurityGroupIds(v string) *DescribeSecurityGroupsRequest {
	s.SecurityGroupIds = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetSecurityGroupName(v string) *DescribeSecurityGroupsRequest {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetSecurityGroupType(v string) *DescribeSecurityGroupsRequest {
	s.SecurityGroupType = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetServiceManaged(v bool) *DescribeSecurityGroupsRequest {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetTag(v []*DescribeSecurityGroupsRequestTag) *DescribeSecurityGroupsRequest {
	s.Tag = v
	return s
}

func (s *DescribeSecurityGroupsRequest) SetVpcId(v string) *DescribeSecurityGroupsRequest {
	s.VpcId = &v
	return s
}

func (s *DescribeSecurityGroupsRequest) Validate() error {
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

type DescribeSecurityGroupsRequestTag struct {
	// The key of tag N to add to the security group. Valid values of N: 1 to 20.
	//
	// Up to 1,000 resources that match the tags specified can be returned in the response. To query more than 1,000 resources that have specified tags added, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the security group. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSecurityGroupsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSecurityGroupsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSecurityGroupsRequestTag) SetKey(v string) *DescribeSecurityGroupsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeSecurityGroupsRequestTag) SetValue(v string) *DescribeSecurityGroupsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeSecurityGroupsRequestTag) Validate() error {
	return dara.Validate(s)
}
