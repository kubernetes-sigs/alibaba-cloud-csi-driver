// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyLifecyclePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *ModifyLifecyclePolicyRequest
	GetFileSystemId() *string
	SetLifecyclePolicyName(v string) *ModifyLifecyclePolicyRequest
	GetLifecyclePolicyName() *string
	SetLifecycleRuleName(v string) *ModifyLifecyclePolicyRequest
	GetLifecycleRuleName() *string
	SetPath(v string) *ModifyLifecyclePolicyRequest
	GetPath() *string
	SetStorageType(v string) *ModifyLifecyclePolicyRequest
	GetStorageType() *string
}

type ModifyLifecyclePolicyRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the lifecycle policy.
	//
	// The name must be 3 to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
	//
	// This parameter is required.
	//
	// example:
	//
	// lifecyclepolicy_01
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The management rule that is associated with the lifecycle policy.
	//
	// Valid values:
	//
	// 	- DEFAULT_ATIME_14: Files that are not accessed in the last 14 days are dumped to the IA storage medium.
	//
	// 	- DEFAULT_ATIME_30: Files that are not accessed in the last 30 days are dumped to the IA storage medium.
	//
	// 	- DEFAULT_ATIME_60: Files that are not accessed in the last 60 days are dumped to the IA storage medium.
	//
	// 	- DEFAULT_ATIME_90: Files that are not accessed in the last 90 days are dumped to the IA storage medium.
	//
	// example:
	//
	// DEFAULT_ATIME_14
	LifecycleRuleName *string `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	// The absolute path of a directory with which the lifecycle policy is associated.
	//
	// The path must start with a forward slash (/) and must be a path that exists in the mount target.
	//
	// example:
	//
	// /pathway/to/folder
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The storage type of the data that is dumped to the IA storage medium.
	//
	// Default value: InfrequentAccess (IA).
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s ModifyLifecyclePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *ModifyLifecyclePolicyRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyLifecyclePolicyRequest) GetLifecyclePolicyName() *string {
	return s.LifecyclePolicyName
}

func (s *ModifyLifecyclePolicyRequest) GetLifecycleRuleName() *string {
	return s.LifecycleRuleName
}

func (s *ModifyLifecyclePolicyRequest) GetPath() *string {
	return s.Path
}

func (s *ModifyLifecyclePolicyRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *ModifyLifecyclePolicyRequest) SetFileSystemId(v string) *ModifyLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetLifecyclePolicyName(v string) *ModifyLifecyclePolicyRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetLifecycleRuleName(v string) *ModifyLifecyclePolicyRequest {
	s.LifecycleRuleName = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetPath(v string) *ModifyLifecyclePolicyRequest {
	s.Path = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) SetStorageType(v string) *ModifyLifecyclePolicyRequest {
	s.StorageType = &v
	return s
}

func (s *ModifyLifecyclePolicyRequest) Validate() error {
	return dara.Validate(s)
}
