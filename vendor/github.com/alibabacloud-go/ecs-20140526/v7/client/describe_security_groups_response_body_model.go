// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeSecurityGroupsResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeSecurityGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSecurityGroupsResponseBody
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSecurityGroupsResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeSecurityGroupsResponseBody
	GetRequestId() *string
	SetSecurityGroups(v *DescribeSecurityGroupsResponseBodySecurityGroups) *DescribeSecurityGroupsResponseBody
	GetSecurityGroups() *DescribeSecurityGroupsResponseBodySecurityGroups
	SetTotalCount(v int32) *DescribeSecurityGroupsResponseBody
	GetTotalCount() *int32
}

type DescribeSecurityGroupsResponseBody struct {
	// A pagination token. If the return value of this parameter is empty when MaxResults and NextToken are used for a paged query, no next page exists.
	//
	// example:
	//
	// e71d8a535bd9cc11
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The page number.
	//
	// > This parameter will be deprecated in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// > This parameter will be deprecated in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the security group.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId      *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SecurityGroups *DescribeSecurityGroupsResponseBodySecurityGroups `json:"SecurityGroups,omitempty" xml:"SecurityGroups,omitempty" type:"Struct"`
	// The total number of security groups returned. If `MaxResults` and `NextToken` are specified in the request, the value of this parameter is not returned.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSecurityGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSecurityGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSecurityGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSecurityGroupsResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityGroupsResponseBody) GetSecurityGroups() *DescribeSecurityGroupsResponseBodySecurityGroups {
	return s.SecurityGroups
}

func (s *DescribeSecurityGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeSecurityGroupsResponseBody) SetNextToken(v string) *DescribeSecurityGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetPageNumber(v int32) *DescribeSecurityGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetPageSize(v int32) *DescribeSecurityGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetRegionId(v string) *DescribeSecurityGroupsResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetRequestId(v string) *DescribeSecurityGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetSecurityGroups(v *DescribeSecurityGroupsResponseBodySecurityGroups) *DescribeSecurityGroupsResponseBody {
	s.SecurityGroups = v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) SetTotalCount(v int32) *DescribeSecurityGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBody) Validate() error {
	if s.SecurityGroups != nil {
		if err := s.SecurityGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityGroupsResponseBodySecurityGroups struct {
	SecurityGroup []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup `json:"SecurityGroup,omitempty" xml:"SecurityGroup,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroups) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroups) GetSecurityGroup() []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	return s.SecurityGroup
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroups) SetSecurityGroup(v []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) *DescribeSecurityGroupsResponseBodySecurityGroups {
	s.SecurityGroup = v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroups) Validate() error {
	if s.SecurityGroup != nil {
		for _, item := range s.SecurityGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup struct {
	AvailableInstanceAmount *int32                                                             `json:"AvailableInstanceAmount,omitempty" xml:"AvailableInstanceAmount,omitempty"`
	CreationTime            *string                                                            `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description             *string                                                            `json:"Description,omitempty" xml:"Description,omitempty"`
	EcsCount                *int32                                                             `json:"EcsCount,omitempty" xml:"EcsCount,omitempty"`
	GroupToGroupRuleCount   *int32                                                             `json:"GroupToGroupRuleCount,omitempty" xml:"GroupToGroupRuleCount,omitempty"`
	ResourceGroupId         *string                                                            `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	RuleCount               *int32                                                             `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
	SecurityGroupId         *string                                                            `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	SecurityGroupName       *string                                                            `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	SecurityGroupType       *string                                                            `json:"SecurityGroupType,omitempty" xml:"SecurityGroupType,omitempty"`
	ServiceID               *int64                                                             `json:"ServiceID,omitempty" xml:"ServiceID,omitempty"`
	ServiceManaged          *bool                                                              `json:"ServiceManaged,omitempty" xml:"ServiceManaged,omitempty"`
	Tags                    *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcId                   *string                                                            `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetAvailableInstanceAmount() *int32 {
	return s.AvailableInstanceAmount
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetEcsCount() *int32 {
	return s.EcsCount
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetGroupToGroupRuleCount() *int32 {
	return s.GroupToGroupRuleCount
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetSecurityGroupType() *string {
	return s.SecurityGroupType
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetServiceID() *int64 {
	return s.ServiceID
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetServiceManaged() *bool {
	return s.ServiceManaged
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetTags() *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTags {
	return s.Tags
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetAvailableInstanceAmount(v int32) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.AvailableInstanceAmount = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetCreationTime(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.CreationTime = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetDescription(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.Description = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetEcsCount(v int32) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.EcsCount = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetGroupToGroupRuleCount(v int32) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.GroupToGroupRuleCount = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetResourceGroupId(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetRuleCount(v int32) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.RuleCount = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetSecurityGroupId(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetSecurityGroupName(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetSecurityGroupType(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.SecurityGroupType = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetServiceID(v int64) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.ServiceID = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetServiceManaged(v bool) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.ServiceManaged = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetTags(v *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTags) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.Tags = v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) SetVpcId(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup {
	s.VpcId = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroup) Validate() error {
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTags struct {
	Tag []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTags) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTags) GetTag() []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag {
	return s.Tag
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTags) SetTag(v []*DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTags {
	s.Tag = v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTags) Validate() error {
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

type DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag) SetTagKey(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag) SetTagValue(v string) *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeSecurityGroupsResponseBodySecurityGroupsSecurityGroupTagsTag) Validate() error {
	return dara.Validate(s)
}
