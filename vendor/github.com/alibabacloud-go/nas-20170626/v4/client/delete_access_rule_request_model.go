// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *DeleteAccessRuleRequest
	GetAccessGroupName() *string
	SetAccessRuleId(v string) *DeleteAccessRuleRequest
	GetAccessRuleId() *string
	SetFileSystemType(v string) *DeleteAccessRuleRequest
	GetFileSystemType() *string
}

type DeleteAccessRuleRequest struct {
	// The name of the permission group.
	//
	// This parameter is required.
	//
	// example:
	//
	// vpc-test
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The rule ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	AccessRuleId *string `json:"AccessRuleId,omitempty" xml:"AccessRuleId,omitempty"`
	// The type of the file system.
	//
	// Valid values:
	//
	// 	- standard (default): General-purpose NAS file system
	//
	// 	- extreme: Extreme NAS file system
	//
	// example:
	//
	// standard
	FileSystemType *string `json:"FileSystemType,omitempty" xml:"FileSystemType,omitempty"`
}

func (s DeleteAccessRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessRuleRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessRuleRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *DeleteAccessRuleRequest) GetAccessRuleId() *string {
	return s.AccessRuleId
}

func (s *DeleteAccessRuleRequest) GetFileSystemType() *string {
	return s.FileSystemType
}

func (s *DeleteAccessRuleRequest) SetAccessGroupName(v string) *DeleteAccessRuleRequest {
	s.AccessGroupName = &v
	return s
}

func (s *DeleteAccessRuleRequest) SetAccessRuleId(v string) *DeleteAccessRuleRequest {
	s.AccessRuleId = &v
	return s
}

func (s *DeleteAccessRuleRequest) SetFileSystemType(v string) *DeleteAccessRuleRequest {
	s.FileSystemType = &v
	return s
}

func (s *DeleteAccessRuleRequest) Validate() error {
	return dara.Validate(s)
}
