// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPoints(v []*DescribeAccessPointsResponseBodyAccessPoints) *DescribeAccessPointsResponseBody
	GetAccessPoints() []*DescribeAccessPointsResponseBodyAccessPoints
	SetNextToken(v string) *DescribeAccessPointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeAccessPointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeAccessPointsResponseBody
	GetTotalCount() *int32
}

type DescribeAccessPointsResponseBody struct {
	// The information about the access point.
	AccessPoints []*DescribeAccessPointsResponseBodyAccessPoints `json:"AccessPoints,omitempty" xml:"AccessPoints,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If NextToken is empty, no next page exists.
	//
	// example:
	//
	// MTY4NzcxOTcwMjAzMDk2Nzc0MyM4MDM4****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6299428C-3861-435D-AE54-9B330A00****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of access points.
	//
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAccessPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBody) GetAccessPoints() []*DescribeAccessPointsResponseBodyAccessPoints {
	return s.AccessPoints
}

func (s *DescribeAccessPointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeAccessPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeAccessPointsResponseBody) SetAccessPoints(v []*DescribeAccessPointsResponseBodyAccessPoints) *DescribeAccessPointsResponseBody {
	s.AccessPoints = v
	return s
}

func (s *DescribeAccessPointsResponseBody) SetNextToken(v string) *DescribeAccessPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeAccessPointsResponseBody) SetRequestId(v string) *DescribeAccessPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessPointsResponseBody) SetTotalCount(v int32) *DescribeAccessPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAccessPointsResponseBody) Validate() error {
	if s.AccessPoints != nil {
		for _, item := range s.AccessPoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeAccessPointsResponseBodyAccessPoints struct {
	// The Alibaba Cloud Resource Name (ARN) of the access point.
	//
	// example:
	//
	// acs:nas:cn-hangzhou:178321033379****:accesspoint/ap-ie15yd****
	ARN *string `json:"ARN,omitempty" xml:"ARN,omitempty"`
	// The name of the permission group.
	//
	// example:
	//
	// test
	AccessGroup *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	// The ID of the access point.
	//
	// example:
	//
	// ap-ie15y*****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The name of the access point.
	//
	// example:
	//
	// test
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// The time when the access point was created.
	//
	// example:
	//
	// 1709619668276167
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The domain name of the access point.
	//
	// example:
	//
	// ap-ie15ydanoz.001014****-w****.cn-hangzhou.nas.aliyuncs.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Indicates whether the Resource Access Management (RAM) policy is enabled.
	//
	// example:
	//
	// false
	EnabledRam *bool `json:"EnabledRam,omitempty" xml:"EnabledRam,omitempty"`
	// The ID of the file system.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The time when the access point was modified.
	//
	// example:
	//
	// 1709619668276167
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The Portable Operating System Interface for UNIX (POSIX) user.
	PosixUser *DescribeAccessPointsResponseBodyAccessPointsPosixUser `json:"PosixUser,omitempty" xml:"PosixUser,omitempty" type:"Struct"`
	// The root directory.
	//
	// example:
	//
	// /
	RootPath *string `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
	// The permissions on the root directory.
	RootPathPermission *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission `json:"RootPathPermission,omitempty" xml:"RootPathPermission,omitempty" type:"Struct"`
	// The status of the root directory.
	//
	// Valid values:
	//
	// 	- 0: The rootpath status is unknown.
	//
	// 	- 1: The rootpath does not exist and may be deleted.
	//
	// 	- 2: The rootpath is normal.
	//
	// example:
	//
	// 2
	RootPathStatus *string `json:"RootPathStatus,omitempty" xml:"RootPathStatus,omitempty"`
	// The status of the access point.
	//
	// Valid values:
	//
	// 	- Active: The access point is available.
	//
	// 	- Inactive: The access point is unavailable.
	//
	// 	- Pending: The access point is being created.
	//
	// 	- Deleting: The access point is being deleted.
	//
	// >  You can mount a file system only if the access point is in the Active state.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The tags of the access point.
	Tags []*DescribeAccessPointsResponseBodyAccessPointsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Switch ID.
	//
	// example:
	//
	// vsw-2zevmwkwyztjuoffg****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-2zesj9afh3y518k9o****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeAccessPointsResponseBodyAccessPoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsResponseBodyAccessPoints) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetARN() *string {
	return s.ARN
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetAccessGroup() *string {
	return s.AccessGroup
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetEnabledRam() *bool {
	return s.EnabledRam
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetPosixUser() *DescribeAccessPointsResponseBodyAccessPointsPosixUser {
	return s.PosixUser
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetRootPath() *string {
	return s.RootPath
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetRootPathPermission() *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission {
	return s.RootPathPermission
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetRootPathStatus() *string {
	return s.RootPathStatus
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetStatus() *string {
	return s.Status
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetTags() []*DescribeAccessPointsResponseBodyAccessPointsTags {
	return s.Tags
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetARN(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.ARN = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetAccessGroup(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.AccessGroup = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetAccessPointId(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.AccessPointId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetAccessPointName(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.AccessPointName = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetCreateTime(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetDomainName(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.DomainName = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetEnabledRam(v bool) *DescribeAccessPointsResponseBodyAccessPoints {
	s.EnabledRam = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetFileSystemId(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.FileSystemId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetModifyTime(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.ModifyTime = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetPosixUser(v *DescribeAccessPointsResponseBodyAccessPointsPosixUser) *DescribeAccessPointsResponseBodyAccessPoints {
	s.PosixUser = v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetRootPath(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.RootPath = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetRootPathPermission(v *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) *DescribeAccessPointsResponseBodyAccessPoints {
	s.RootPathPermission = v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetRootPathStatus(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.RootPathStatus = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetStatus(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.Status = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetTags(v []*DescribeAccessPointsResponseBodyAccessPointsTags) *DescribeAccessPointsResponseBodyAccessPoints {
	s.Tags = v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetVSwitchId(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) SetVpcId(v string) *DescribeAccessPointsResponseBodyAccessPoints {
	s.VpcId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPoints) Validate() error {
	if s.PosixUser != nil {
		if err := s.PosixUser.Validate(); err != nil {
			return err
		}
	}
	if s.RootPathPermission != nil {
		if err := s.RootPathPermission.Validate(); err != nil {
			return err
		}
	}
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

type DescribeAccessPointsResponseBodyAccessPointsPosixUser struct {
	// The ID of the POSIX user group.
	//
	// example:
	//
	// 12
	PosixGroupId *int32 `json:"PosixGroupId,omitempty" xml:"PosixGroupId,omitempty"`
	// The IDs of the secondary user groups.
	PosixSecondaryGroupIds []*int32 `json:"PosixSecondaryGroupIds,omitempty" xml:"PosixSecondaryGroupIds,omitempty" type:"Repeated"`
	// The ID of the POSIX user.
	//
	// example:
	//
	// 123
	PosixUserId *int32 `json:"PosixUserId,omitempty" xml:"PosixUserId,omitempty"`
}

func (s DescribeAccessPointsResponseBodyAccessPointsPosixUser) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsResponseBodyAccessPointsPosixUser) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBodyAccessPointsPosixUser) GetPosixGroupId() *int32 {
	return s.PosixGroupId
}

func (s *DescribeAccessPointsResponseBodyAccessPointsPosixUser) GetPosixSecondaryGroupIds() []*int32 {
	return s.PosixSecondaryGroupIds
}

func (s *DescribeAccessPointsResponseBodyAccessPointsPosixUser) GetPosixUserId() *int32 {
	return s.PosixUserId
}

func (s *DescribeAccessPointsResponseBodyAccessPointsPosixUser) SetPosixGroupId(v int32) *DescribeAccessPointsResponseBodyAccessPointsPosixUser {
	s.PosixGroupId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsPosixUser) SetPosixSecondaryGroupIds(v []*int32) *DescribeAccessPointsResponseBodyAccessPointsPosixUser {
	s.PosixSecondaryGroupIds = v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsPosixUser) SetPosixUserId(v int32) *DescribeAccessPointsResponseBodyAccessPointsPosixUser {
	s.PosixUserId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsPosixUser) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessPointsResponseBodyAccessPointsRootPathPermission struct {
	// The ID of the owner group.
	//
	// example:
	//
	// 12
	OwnerGroupId *int64 `json:"OwnerGroupId,omitempty" xml:"OwnerGroupId,omitempty"`
	// The owner ID.
	//
	// example:
	//
	// 1
	OwnerUserId *int64 `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The POSIX permission.
	//
	// example:
	//
	// 0755
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
}

func (s DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) GetOwnerGroupId() *int64 {
	return s.OwnerGroupId
}

func (s *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) GetPermission() *string {
	return s.Permission
}

func (s *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) SetOwnerGroupId(v int64) *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission {
	s.OwnerGroupId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) SetOwnerUserId(v int64) *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) SetPermission(v string) *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission {
	s.Permission = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsRootPathPermission) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessPointsResponseBodyAccessPointsTags struct {
	// The key of the tag that is added to the resource.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAccessPointsResponseBodyAccessPointsTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointsResponseBodyAccessPointsTags) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointsResponseBodyAccessPointsTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAccessPointsResponseBodyAccessPointsTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAccessPointsResponseBodyAccessPointsTags) SetKey(v string) *DescribeAccessPointsResponseBodyAccessPointsTags {
	s.Key = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsTags) SetValue(v string) *DescribeAccessPointsResponseBodyAccessPointsTags {
	s.Value = &v
	return s
}

func (s *DescribeAccessPointsResponseBodyAccessPointsTags) Validate() error {
	return dara.Validate(s)
}
