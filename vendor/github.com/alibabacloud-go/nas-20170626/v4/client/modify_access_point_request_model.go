// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroup(v string) *ModifyAccessPointRequest
	GetAccessGroup() *string
	SetAccessPointId(v string) *ModifyAccessPointRequest
	GetAccessPointId() *string
	SetAccessPointName(v string) *ModifyAccessPointRequest
	GetAccessPointName() *string
	SetEnabledRam(v bool) *ModifyAccessPointRequest
	GetEnabledRam() *bool
	SetFileSystemId(v string) *ModifyAccessPointRequest
	GetFileSystemId() *string
}

type ModifyAccessPointRequest struct {
	// The name of the permission group.
	//
	// This parameter is required for a General-purpose File Storage NAS (NAS) file system.
	//
	// The default permission group for virtual private clouds (VPCs) is named DEFAULT_VPC_GROUP_NAME.
	//
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroup *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	// The ID of the access point.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-ie15yd****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The name of the access point.
	//
	// example:
	//
	// test
	AccessPointName *string `json:"AccessPointName,omitempty" xml:"AccessPointName,omitempty"`
	// Specifies whether to enable the Resource Access Management (RAM) policy. Valid values:
	//
	// 	- true: The RAM policy is enabled.
	//
	// 	- false (default): The RAM policy is disabled.
	//
	// >  After the RAM policy is enabled for access points, no RAM user is allowed to use access points to mount and access data by default. To use access points to mount and access data as a RAM user, you must grant the related access permissions to the RAM user. If the RAM policy is disabled, access points can be anonymously mounted.
	//
	// example:
	//
	// false
	EnabledRam *bool `json:"EnabledRam,omitempty" xml:"EnabledRam,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ModifyAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAccessPointRequest) GoString() string {
	return s.String()
}

func (s *ModifyAccessPointRequest) GetAccessGroup() *string {
	return s.AccessGroup
}

func (s *ModifyAccessPointRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *ModifyAccessPointRequest) GetAccessPointName() *string {
	return s.AccessPointName
}

func (s *ModifyAccessPointRequest) GetEnabledRam() *bool {
	return s.EnabledRam
}

func (s *ModifyAccessPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyAccessPointRequest) SetAccessGroup(v string) *ModifyAccessPointRequest {
	s.AccessGroup = &v
	return s
}

func (s *ModifyAccessPointRequest) SetAccessPointId(v string) *ModifyAccessPointRequest {
	s.AccessPointId = &v
	return s
}

func (s *ModifyAccessPointRequest) SetAccessPointName(v string) *ModifyAccessPointRequest {
	s.AccessPointName = &v
	return s
}

func (s *ModifyAccessPointRequest) SetEnabledRam(v bool) *ModifyAccessPointRequest {
	s.EnabledRam = &v
	return s
}

func (s *ModifyAccessPointRequest) SetFileSystemId(v string) *ModifyAccessPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
