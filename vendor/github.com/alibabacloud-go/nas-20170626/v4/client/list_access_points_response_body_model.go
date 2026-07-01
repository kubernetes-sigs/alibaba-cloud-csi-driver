// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAccessPointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPoints(v []*ListAccessPointsResponseBodyAccessPoints) *ListAccessPointsResponseBody
	GetAccessPoints() []*ListAccessPointsResponseBodyAccessPoints
	SetNextToken(v string) *ListAccessPointsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListAccessPointsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAccessPointsResponseBody
	GetTotalCount() *int32
}

type ListAccessPointsResponseBody struct {
	AccessPoints []*ListAccessPointsResponseBodyAccessPoints `json:"AccessPoints,omitempty" xml:"AccessPoints,omitempty" type:"Repeated"`
	// example:
	//
	// 52frCAAAAABoZS90cm****==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 20
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAccessPointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResponseBody) GetAccessPoints() []*ListAccessPointsResponseBodyAccessPoints {
	return s.AccessPoints
}

func (s *ListAccessPointsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListAccessPointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAccessPointsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAccessPointsResponseBody) SetAccessPoints(v []*ListAccessPointsResponseBodyAccessPoints) *ListAccessPointsResponseBody {
	s.AccessPoints = v
	return s
}

func (s *ListAccessPointsResponseBody) SetNextToken(v string) *ListAccessPointsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListAccessPointsResponseBody) SetRequestId(v string) *ListAccessPointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAccessPointsResponseBody) SetTotalCount(v int32) *ListAccessPointsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAccessPointsResponseBody) Validate() error {
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

type ListAccessPointsResponseBodyAccessPoints struct {
	// example:
	//
	// acs:nas:cn-hangzhou:178321033379****:accesspoint/ap-ie15yd****
	ARN *string `json:"ARN,omitempty" xml:"ARN,omitempty"`
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroup *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	// example:
	//
	// ap-ie15y*****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// example:
	//
	// test
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// example:
	//
	// agentic-229oypxjgpau2****
	AgenticSpaceId *string `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
	// example:
	//
	// 2026-03-28T06:32:14Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// example:
	//
	// 2026-06-10T10:08:08Z
	CreateTimeUtc *string `json:"CreateTimeUtc,omitempty" xml:"CreateTimeUtc,omitempty"`
	// example:
	//
	// ap-ie15ydanoz.001014****-w****.cn-hangzhou.nas.aliyuncs.com
	DomainName *string `json:"DomainName,omitempty" xml:"DomainName,omitempty"`
	// example:
	//
	// false
	EnabledRam *bool `json:"EnabledRam,omitempty" xml:"EnabledRam,omitempty"`
	// example:
	//
	// 091yj49baxscll2****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// 2026-06-24T02:10:23Z
	ModifyTime *string `json:"ModifyTime,omitempty" xml:"ModifyTime,omitempty"`
	// example:
	//
	// 2026-06-10T10:08:08Z
	ModifyTimeUtc *string                                            `json:"ModifyTimeUtc,omitempty" xml:"ModifyTimeUtc,omitempty"`
	PosixUser     *ListAccessPointsResponseBodyAccessPointsPosixUser `json:"PosixUser,omitempty" xml:"PosixUser,omitempty" type:"Struct"`
	// example:
	//
	// /
	RootPath           *string                                                     `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
	RootPathPermission *ListAccessPointsResponseBodyAccessPointsRootPathPermission `json:"RootPathPermission,omitempty" xml:"RootPathPermission,omitempty" type:"Struct"`
	// example:
	//
	// 2
	RootPathStatus *string `json:"RootPathStatus,omitempty" xml:"RootPathStatus,omitempty"`
	// example:
	//
	// Active
	Status *string                                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListAccessPointsResponseBodyAccessPointsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// example:
	//
	// vsw-2zevmwkwyztjuoffg****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-2zesj9afh3y518k9o****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListAccessPointsResponseBodyAccessPoints) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResponseBodyAccessPoints) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetARN() *string {
	return s.ARN
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetAccessGroup() *string {
	return s.AccessGroup
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetCreateTimeUtc() *string {
	return s.CreateTimeUtc
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetDomainName() *string {
	return s.DomainName
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetEnabledRam() *bool {
	return s.EnabledRam
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetModifyTime() *string {
	return s.ModifyTime
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetModifyTimeUtc() *string {
	return s.ModifyTimeUtc
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetPosixUser() *ListAccessPointsResponseBodyAccessPointsPosixUser {
	return s.PosixUser
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetRootPath() *string {
	return s.RootPath
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetRootPathPermission() *ListAccessPointsResponseBodyAccessPointsRootPathPermission {
	return s.RootPathPermission
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetRootPathStatus() *string {
	return s.RootPathStatus
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetStatus() *string {
	return s.Status
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetTags() []*ListAccessPointsResponseBodyAccessPointsTags {
	return s.Tags
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListAccessPointsResponseBodyAccessPoints) GetVpcId() *string {
	return s.VpcId
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetARN(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.ARN = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetAccessGroup(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.AccessGroup = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetAccessPointId(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.AccessPointId = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetAccessPointName(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.AccessPointName = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetAgenticSpaceId(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.AgenticSpaceId = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetCreateTime(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.CreateTime = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetCreateTimeUtc(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.CreateTimeUtc = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetDomainName(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.DomainName = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetEnabledRam(v bool) *ListAccessPointsResponseBodyAccessPoints {
	s.EnabledRam = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetFileSystemId(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.FileSystemId = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetModifyTime(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.ModifyTime = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetModifyTimeUtc(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.ModifyTimeUtc = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetPosixUser(v *ListAccessPointsResponseBodyAccessPointsPosixUser) *ListAccessPointsResponseBodyAccessPoints {
	s.PosixUser = v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetRootPath(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.RootPath = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetRootPathPermission(v *ListAccessPointsResponseBodyAccessPointsRootPathPermission) *ListAccessPointsResponseBodyAccessPoints {
	s.RootPathPermission = v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetRootPathStatus(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.RootPathStatus = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetStatus(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.Status = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetTags(v []*ListAccessPointsResponseBodyAccessPointsTags) *ListAccessPointsResponseBodyAccessPoints {
	s.Tags = v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetVSwitchId(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.VSwitchId = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) SetVpcId(v string) *ListAccessPointsResponseBodyAccessPoints {
	s.VpcId = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPoints) Validate() error {
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

type ListAccessPointsResponseBodyAccessPointsPosixUser struct {
	// example:
	//
	// 10
	PosixGroupId           *int32   `json:"PosixGroupId,omitempty" xml:"PosixGroupId,omitempty"`
	PosixSecondaryGroupIds []*int32 `json:"PosixSecondaryGroupIds,omitempty" xml:"PosixSecondaryGroupIds,omitempty" type:"Repeated"`
	// example:
	//
	// 156
	PosixUserId *int32 `json:"PosixUserId,omitempty" xml:"PosixUserId,omitempty"`
}

func (s ListAccessPointsResponseBodyAccessPointsPosixUser) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResponseBodyAccessPointsPosixUser) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResponseBodyAccessPointsPosixUser) GetPosixGroupId() *int32 {
	return s.PosixGroupId
}

func (s *ListAccessPointsResponseBodyAccessPointsPosixUser) GetPosixSecondaryGroupIds() []*int32 {
	return s.PosixSecondaryGroupIds
}

func (s *ListAccessPointsResponseBodyAccessPointsPosixUser) GetPosixUserId() *int32 {
	return s.PosixUserId
}

func (s *ListAccessPointsResponseBodyAccessPointsPosixUser) SetPosixGroupId(v int32) *ListAccessPointsResponseBodyAccessPointsPosixUser {
	s.PosixGroupId = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPointsPosixUser) SetPosixSecondaryGroupIds(v []*int32) *ListAccessPointsResponseBodyAccessPointsPosixUser {
	s.PosixSecondaryGroupIds = v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPointsPosixUser) SetPosixUserId(v int32) *ListAccessPointsResponseBodyAccessPointsPosixUser {
	s.PosixUserId = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPointsPosixUser) Validate() error {
	return dara.Validate(s)
}

type ListAccessPointsResponseBodyAccessPointsRootPathPermission struct {
	// example:
	//
	// 12
	OwnerGroupId *int64 `json:"OwnerGroupId,omitempty" xml:"OwnerGroupId,omitempty"`
	// example:
	//
	// 1
	OwnerUserId *int64 `json:"OwnerUserId,omitempty" xml:"OwnerUserId,omitempty"`
	// example:
	//
	// 0755
	Permission *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
}

func (s ListAccessPointsResponseBodyAccessPointsRootPathPermission) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResponseBodyAccessPointsRootPathPermission) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResponseBodyAccessPointsRootPathPermission) GetOwnerGroupId() *int64 {
	return s.OwnerGroupId
}

func (s *ListAccessPointsResponseBodyAccessPointsRootPathPermission) GetOwnerUserId() *int64 {
	return s.OwnerUserId
}

func (s *ListAccessPointsResponseBodyAccessPointsRootPathPermission) GetPermission() *string {
	return s.Permission
}

func (s *ListAccessPointsResponseBodyAccessPointsRootPathPermission) SetOwnerGroupId(v int64) *ListAccessPointsResponseBodyAccessPointsRootPathPermission {
	s.OwnerGroupId = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPointsRootPathPermission) SetOwnerUserId(v int64) *ListAccessPointsResponseBodyAccessPointsRootPathPermission {
	s.OwnerUserId = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPointsRootPathPermission) SetPermission(v string) *ListAccessPointsResponseBodyAccessPointsRootPathPermission {
	s.Permission = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPointsRootPathPermission) Validate() error {
	return dara.Validate(s)
}

type ListAccessPointsResponseBodyAccessPointsTags struct {
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListAccessPointsResponseBodyAccessPointsTags) String() string {
	return dara.Prettify(s)
}

func (s ListAccessPointsResponseBodyAccessPointsTags) GoString() string {
	return s.String()
}

func (s *ListAccessPointsResponseBodyAccessPointsTags) GetKey() *string {
	return s.Key
}

func (s *ListAccessPointsResponseBodyAccessPointsTags) GetValue() *string {
	return s.Value
}

func (s *ListAccessPointsResponseBodyAccessPointsTags) SetKey(v string) *ListAccessPointsResponseBodyAccessPointsTags {
	s.Key = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPointsTags) SetValue(v string) *ListAccessPointsResponseBodyAccessPointsTags {
	s.Value = &v
	return s
}

func (s *ListAccessPointsResponseBodyAccessPointsTags) Validate() error {
	return dara.Validate(s)
}
