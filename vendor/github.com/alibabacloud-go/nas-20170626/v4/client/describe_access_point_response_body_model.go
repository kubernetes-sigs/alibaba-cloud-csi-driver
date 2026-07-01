// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessPointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPoint(v *DescribeAccessPointResponseBodyAccessPoint) *DescribeAccessPointResponseBody
	GetAccessPoint() *DescribeAccessPointResponseBodyAccessPoint
	SetRequestId(v string) *DescribeAccessPointResponseBody
	GetRequestId() *string
}

type DescribeAccessPointResponseBody struct {
	// The access point information.
	AccessPoint *DescribeAccessPointResponseBodyAccessPoint `json:"AccessPoint,omitempty" xml:"AccessPoint,omitempty" type:"Struct"`
	// The request ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeAccessPointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponseBody) GetAccessPoint() *DescribeAccessPointResponseBodyAccessPoint {
	return s.AccessPoint
}

func (s *DescribeAccessPointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAccessPointResponseBody) SetAccessPoint(v *DescribeAccessPointResponseBodyAccessPoint) *DescribeAccessPointResponseBody {
	s.AccessPoint = v
	return s
}

func (s *DescribeAccessPointResponseBody) SetRequestId(v string) *DescribeAccessPointResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAccessPointResponseBody) Validate() error {
	if s.AccessPoint != nil {
		if err := s.AccessPoint.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeAccessPointResponseBodyAccessPoint struct {
	// The access point ARN.
	//
	// example:
	//
	// acs:nas:cn-hangzhou:178321033379****:accesspoint/ap-ie15yd****
	ARN *string `json:"ARN,omitempty" xml:"ARN,omitempty"`
	// The permission group name.
	//
	// example:
	//
	// test
	AccessGroup *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	// The access point ID.
	//
	// example:
	//
	// ap-ie15yd****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The access point name.
	//
	// example:
	//
	// test
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// AgenticSpace Id。
	//
	// example:
	//
	// agentic-229oypxjgpau2****
	AgenticSpaceId *string `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
	// The time when the access point was created.
	//
	// example:
	//
	// 1709619668276167
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the AgenticSpace was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2026-06-10T10:08:08Z
	CreateTimeUtc *string `json:"CreateTimeUtc,omitempty" xml:"CreateTimeUtc,omitempty"`
	// The access point domain name.
	//
	// example:
	//
	// ap-ie15ydanoz.001014****-w****.cn-hangzhou.nas.aliyuncs.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// Indicates whether the RAM policy is enabled.
	//
	// example:
	//
	// false
	EnabledRam *bool `json:"EnabledRam,omitempty" xml:"EnabledRam,omitempty"`
	// The file system ID.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The time when the access point was last modified.
	//
	// example:
	//
	// 1709619668276167
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// The time when the AgenticSpace was last modified. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2026-06-10T10:08:08Z
	ModifyTimeUtc *string `json:"ModifyTimeUtc,omitempty" xml:"ModifyTimeUtc,omitempty"`
	// The POSIX user.
	PosixUser *DescribeAccessPointResponseBodyAccessPointPosixUser `json:"PosixUser,omitempty" xml:"PosixUser,omitempty" type:"Struct"`
	// The region ID.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The root directory.
	//
	// example:
	//
	// /
	RootPath *string `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
	// The permissions for creating the root directory.
	RootPathPermission *DescribeAccessPointResponseBodyAccessPointRootPathPermission `json:"RootPathPermission,omitempty" xml:"RootPathPermission,omitempty" type:"Struct"`
	// The current root directory status.
	//
	// Valid values:
	//
	// - 0: The root path status is unknown.
	//
	// - 1: The root path does not exist. It may have been deleted by the user.
	//
	// - 2: The root path status is normal.
	//
	// example:
	//
	// 2
	RootPathStatus *string `json:"RootPathStatus,omitempty" xml:"RootPathStatus,omitempty"`
	// The current access point status.
	//
	// Valid values:
	//
	// - Active: active
	//
	// - Inactive: inactive
	//
	// - Pending: being created
	//
	// - Deleting: being deleted
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of access point tags.
	Tags []*DescribeAccessPointResponseBodyAccessPointTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-2zevmwkwyztjuoffg****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// The VPC must be the same as the VPC of the Elastic Computing Service (ECS) server to which you want to mount the file system.
	//
	// example:
	//
	// vpc-2zesj9afh3y518k9o****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeAccessPointResponseBodyAccessPoint) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointResponseBodyAccessPoint) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetARN() *string {
	return s.ARN
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetAccessGroup() *string {
	return s.AccessGroup
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetCreateTimeUtc() *string {
	return s.CreateTimeUtc
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetDomainName() *string {
	return s.DomainName
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetEnabledRam() *bool {
	return s.EnabledRam
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetModifyTimeUtc() *string {
	return s.ModifyTimeUtc
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetPosixUser() *DescribeAccessPointResponseBodyAccessPointPosixUser {
	return s.PosixUser
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetRootPath() *string {
	return s.RootPath
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetRootPathPermission() *DescribeAccessPointResponseBodyAccessPointRootPathPermission {
	return s.RootPathPermission
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetRootPathStatus() *string {
	return s.RootPathStatus
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetStatus() *string {
	return s.Status
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetTags() []*DescribeAccessPointResponseBodyAccessPointTags {
	return s.Tags
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeAccessPointResponseBodyAccessPoint) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetARN(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.ARN = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetAccessGroup(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.AccessGroup = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetAccessPointId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.AccessPointId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetAccessPointName(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.AccessPointName = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetAgenticSpaceId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.AgenticSpaceId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetCreateTime(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.CreateTime = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetCreateTimeUtc(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.CreateTimeUtc = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetDomainName(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.DomainName = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetEnabledRam(v bool) *DescribeAccessPointResponseBodyAccessPoint {
	s.EnabledRam = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetFileSystemId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.FileSystemId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetModifyTime(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.ModifyTime = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetModifyTimeUtc(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.ModifyTimeUtc = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetPosixUser(v *DescribeAccessPointResponseBodyAccessPointPosixUser) *DescribeAccessPointResponseBodyAccessPoint {
	s.PosixUser = v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetRegionId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.RegionId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetRootPath(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.RootPath = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetRootPathPermission(v *DescribeAccessPointResponseBodyAccessPointRootPathPermission) *DescribeAccessPointResponseBodyAccessPoint {
	s.RootPathPermission = v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetRootPathStatus(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.RootPathStatus = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetStatus(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.Status = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetTags(v []*DescribeAccessPointResponseBodyAccessPointTags) *DescribeAccessPointResponseBodyAccessPoint {
	s.Tags = v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetVSwitchId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.VSwitchId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) SetVpcId(v string) *DescribeAccessPointResponseBodyAccessPoint {
	s.VpcId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPoint) Validate() error {
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

type DescribeAccessPointResponseBodyAccessPointPosixUser struct {
	// The POSIX user group ID.
	//
	// example:
	//
	// 12
	PosixGroupId *int32 `json:"PosixGroupId,omitempty" xml:"PosixGroupId,omitempty"`
	// The secondary user group ID.
	PosixSecondaryGroupIds []*int32 `json:"PosixSecondaryGroupIds,omitempty" xml:"PosixSecondaryGroupIds,omitempty" type:"Repeated"`
	// The POSIX user ID.
	//
	// example:
	//
	// 123
	PosixUserId *int32 `json:"PosixUserId,omitempty" xml:"PosixUserId,omitempty"`
}

func (s DescribeAccessPointResponseBodyAccessPointPosixUser) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointResponseBodyAccessPointPosixUser) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponseBodyAccessPointPosixUser) GetPosixGroupId() *int32 {
	return s.PosixGroupId
}

func (s *DescribeAccessPointResponseBodyAccessPointPosixUser) GetPosixSecondaryGroupIds() []*int32 {
	return s.PosixSecondaryGroupIds
}

func (s *DescribeAccessPointResponseBodyAccessPointPosixUser) GetPosixUserId() *int32 {
	return s.PosixUserId
}

func (s *DescribeAccessPointResponseBodyAccessPointPosixUser) SetPosixGroupId(v int32) *DescribeAccessPointResponseBodyAccessPointPosixUser {
	s.PosixGroupId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointPosixUser) SetPosixSecondaryGroupIds(v []*int32) *DescribeAccessPointResponseBodyAccessPointPosixUser {
	s.PosixSecondaryGroupIds = v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointPosixUser) SetPosixUserId(v int32) *DescribeAccessPointResponseBodyAccessPointPosixUser {
	s.PosixUserId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointPosixUser) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessPointResponseBodyAccessPointRootPathPermission struct {
	// The file group ID.
	//
	// example:
	//
	// 123
	OwnerGroupId *int32 `json:"OwnerGroupId,omitempty" xml:"OwnerGroupId,omitempty"`
	// The file owner ID.
	//
	// example:
	//
	// 1
	OwnerUserId *int32 `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// The POSIX permissions.
	//
	// example:
	//
	// 0755
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
}

func (s DescribeAccessPointResponseBodyAccessPointRootPathPermission) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointResponseBodyAccessPointRootPathPermission) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponseBodyAccessPointRootPathPermission) GetOwnerGroupId() *int32 {
	return s.OwnerGroupId
}

func (s *DescribeAccessPointResponseBodyAccessPointRootPathPermission) GetOwnerUserId() *int32 {
	return s.OwnerUserId
}

func (s *DescribeAccessPointResponseBodyAccessPointRootPathPermission) GetPermission() *string {
	return s.Permission
}

func (s *DescribeAccessPointResponseBodyAccessPointRootPathPermission) SetOwnerGroupId(v int32) *DescribeAccessPointResponseBodyAccessPointRootPathPermission {
	s.OwnerGroupId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointRootPathPermission) SetOwnerUserId(v int32) *DescribeAccessPointResponseBodyAccessPointRootPathPermission {
	s.OwnerUserId = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointRootPathPermission) SetPermission(v string) *DescribeAccessPointResponseBodyAccessPointRootPathPermission {
	s.Permission = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointRootPathPermission) Validate() error {
	return dara.Validate(s)
}

type DescribeAccessPointResponseBodyAccessPointTags struct {
	// The tag key.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeAccessPointResponseBodyAccessPointTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointResponseBodyAccessPointTags) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointResponseBodyAccessPointTags) GetKey() *string {
	return s.Key
}

func (s *DescribeAccessPointResponseBodyAccessPointTags) GetValue() *string {
	return s.Value
}

func (s *DescribeAccessPointResponseBodyAccessPointTags) SetKey(v string) *DescribeAccessPointResponseBodyAccessPointTags {
	s.Key = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointTags) SetValue(v string) *DescribeAccessPointResponseBodyAccessPointTags {
	s.Value = &v
	return s
}

func (s *DescribeAccessPointResponseBodyAccessPointTags) Validate() error {
	return dara.Validate(s)
}
