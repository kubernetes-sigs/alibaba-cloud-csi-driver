// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLifecyclePolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *CreateLifecyclePolicyRequest
	GetFileSystemId() *string
	SetLifecyclePolicyName(v string) *CreateLifecyclePolicyRequest
	GetLifecyclePolicyName() *string
	SetLifecycleRuleName(v string) *CreateLifecyclePolicyRequest
	GetLifecycleRuleName() *string
	SetPath(v string) *CreateLifecyclePolicyRequest
	GetPath() *string
	SetPaths(v []*string) *CreateLifecyclePolicyRequest
	GetPaths() []*string
	SetStorageType(v string) *CreateLifecyclePolicyRequest
	GetStorageType() *string
}

type CreateLifecyclePolicyRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The name of the lifecycle policy. The name must be 3 to 64 characters in length and can contain letters, digits, underscores (_), and hyphens (-). The name must start with a letter.
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
	// The absolute path of the directory that is associated with the lifecycle policy.
	//
	// If you specify this parameter, you can associate the lifecycle policy with only one directory. The path must start with a forward slash (/) and must be a path that exists in the mount target.
	//
	// > We recommend that you specify the Paths.N parameter so that you can associate the lifecycle policy with multiple directories.
	//
	// example:
	//
	// /pathway/to/folder
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The absolute paths of the directories that are associated with the lifecycle policy.
	//
	// If you specify this parameter, you can associate the lifecycle policy with multiple directories. Each path must start with a forward slash (/) and must be a path that exists in the mount target. Valid values of N: 1 to 10.
	//
	// example:
	//
	// "/path1", "/path2"
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The storage type of the data that is dumped to the IA storage medium.
	//
	// Default value: InfrequentAccess (IA).
	//
	// This parameter is required.
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s CreateLifecyclePolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecyclePolicyRequest) GoString() string {
	return s.String()
}

func (s *CreateLifecyclePolicyRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateLifecyclePolicyRequest) GetLifecyclePolicyName() *string {
	return s.LifecyclePolicyName
}

func (s *CreateLifecyclePolicyRequest) GetLifecycleRuleName() *string {
	return s.LifecycleRuleName
}

func (s *CreateLifecyclePolicyRequest) GetPath() *string {
	return s.Path
}

func (s *CreateLifecyclePolicyRequest) GetPaths() []*string {
	return s.Paths
}

func (s *CreateLifecyclePolicyRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateLifecyclePolicyRequest) SetFileSystemId(v string) *CreateLifecyclePolicyRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetLifecyclePolicyName(v string) *CreateLifecyclePolicyRequest {
	s.LifecyclePolicyName = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetLifecycleRuleName(v string) *CreateLifecyclePolicyRequest {
	s.LifecycleRuleName = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetPath(v string) *CreateLifecyclePolicyRequest {
	s.Path = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetPaths(v []*string) *CreateLifecyclePolicyRequest {
	s.Paths = v
	return s
}

func (s *CreateLifecyclePolicyRequest) SetStorageType(v string) *CreateLifecyclePolicyRequest {
	s.StorageType = &v
	return s
}

func (s *CreateLifecyclePolicyRequest) Validate() error {
	return dara.Validate(s)
}
