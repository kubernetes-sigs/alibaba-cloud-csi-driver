// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroups(v *DescribeAccessGroupsResponseBodyAccessGroups) *DescribeAccessGroupsResponseBody
	GetAccessGroups() *DescribeAccessGroupsResponseBodyAccessGroups
	SetPageNumber(v int32) *DescribeAccessGroupsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeAccessGroupsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeAccessGroupsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAccessGroupsResponseBody
	GetTotalCount() *int32
}

type DescribeAccessGroupsResponseBody struct {
	// The queried permission groups.
	AccessGroups *DescribeAccessGroupsResponseBodyAccessGroups `json:"AccessGroups,omitempty" xml:"AccessGroups,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of permission groups returned per page.
	//
	// example:
	//
	// 2
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 2514F97E-FFF0-4A1F-BF04-729CEAC6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of permission groups.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessGroupsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBody) GetAccessGroups() *DescribeAccessGroupsResponseBodyAccessGroups {
	return s.AccessGroups
}

func (s *DescribeAccessGroupsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeAccessGroupsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeAccessGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessGroupsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAccessGroupsResponseBody) SetAccessGroups(v *DescribeAccessGroupsResponseBodyAccessGroups) *DescribeAccessGroupsResponseBody {
	s.AccessGroups = v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetPageNumber(v int32) *DescribeAccessGroupsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetPageSize(v int32) *DescribeAccessGroupsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetRequestId(v string) *DescribeAccessGroupsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) SetTotalCount(v int32) *DescribeAccessGroupsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessGroupsResponseBody) Validate() error {
	if s.AccessGroups != nil {
		if err := s.AccessGroups.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccessGroupsResponseBodyAccessGroups struct {
	AccessGroup []*DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty" type:"Repeated"`
}

func (s DescribeAccessGroupsResponseBodyAccessGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessGroupsResponseBodyAccessGroups) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBodyAccessGroups) GetAccessGroup() []*DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	return s.AccessGroup
}

func (s *DescribeAccessGroupsResponseBodyAccessGroups) SetAccessGroup(v []*DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) *DescribeAccessGroupsResponseBodyAccessGroups {
	s.AccessGroup = v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroups) Validate() error {
	if s.AccessGroup != nil {
		for _, item := range s.AccessGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup struct {
	// The name of the permission group.
	//
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The network type of the permission group. Valid value: **Vpc**.
	//
	// example:
	//
	// Vpc
	AccessGroupType *string `json:"AccessGroupType,omitempty" xml:"AccessGroupType,omitempty"`
	// The time when the permission group was created.
	//
	// example:
	//
	// 2020-01-05T16:00:00Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the permission group.
	//
	// example:
	//
	// This is a test access group.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard: General-purpose NAS
	//
	// 	- extreme: Extreme NAS
	//
	// 	- cpfs: CPFS
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
	// The number of mount targets to which the permission group is attached.
	//
	// example:
	//
	// 0
	MountTargetCount *int32 `json:"MountTargetCount,omitempty" xml:"MountTargetCount,omitempty"`
	// Region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The total number of rules in the permission group.
	//
	// example:
	//
	// 0
	RuleCount *int32 `json:"RuleCount,omitempty" xml:"RuleCount,omitempty"`
}

func (s DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GoString() string {
	return s.String()
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GetAccessGroupType() *string {
	return s.AccessGroupType
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GetDescription() *string {
	return s.Description
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GetMountTargetCount() *int32 {
	return s.MountTargetCount
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) GetRuleCount() *int32 {
	return s.RuleCount
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetAccessGroupName(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetAccessGroupType(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.AccessGroupType = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetCreateTime(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetDescription(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.Description = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetFileSystemType(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.FileSystemType = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetMountTargetCount(v int32) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.MountTargetCount = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetRegionId(v string) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) SetRuleCount(v int32) *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup {
	s.RuleCount = &v
	return s
}

func (s *DescribeAccessGroupsResponseBodyAccessGroupsAccessGroup) Validate() error {
	return dara.Validate(s)
}
