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
	SetLifecyclePolicyId(v string) *ModifyLifecyclePolicyRequest
	GetLifecyclePolicyId() *string
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
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the lifecycle policy.
	//
	// example:
	//
	// lc-xxx
	LifecyclePolicyId *string `json:"LifecyclePolicyId,omitempty" xml:"LifecyclePolicyId,omitempty"`
	// The Policy Name of the lifecycle management policy.
	//
	// The name must be 3 to 64 characters in length, must start with an uppercase letter or lowercase letter, and can contain letters, digits, underscores (_), or hyphens (-).
	//
	// example:
	//
	// lifecyclepolicy_01
	LifecyclePolicyName *string `json:"LifecyclePolicyName,omitempty" xml:"LifecyclePolicyName,omitempty"`
	// The management rule associated with the lifecycle management policy.
	//
	// Valid values:
	//
	// - DEFAULT_ATIME_14: files that have not been accessed for 14 days.
	//
	// - DEFAULT_ATIME_30: files that have not been accessed for 30 days.
	//
	// - DEFAULT_ATIME_60: files that have not been accessed for 60 days.
	//
	// - DEFAULT_ATIME_90: files that have not been accessed for 90 days.
	//
	// - DEFAULT_ATIME_180: files that have not been accessed for 180 days. DEFAULT_ATIME_180 is supported only when StorageType is set to Archive.
	//
	// > If an IA storage class policy has already been configured for the directory, the time period specified for the archive policy must be longer than that of the IA storage class policy.
	//
	// example:
	//
	// DEFAULT_ATIME_14
	LifecycleRuleName *string `json:"LifecycleRuleName,omitempty" xml:"LifecycleRuleName,omitempty"`
	// The absolute path of a single directory configured in the lifecycle management policy.
	//
	// The path must start with a forward slash (/) and must be an existing path in the mount target.
	//
	// example:
	//
	// /pathway/to/folder
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The storage type.
	//
	// - InfrequentAccess: IA storage class.
	//
	// - Archive: Archive storage class.
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

func (s *ModifyLifecyclePolicyRequest) GetLifecyclePolicyId() *string {
	return s.LifecyclePolicyId
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

func (s *ModifyLifecyclePolicyRequest) SetLifecyclePolicyId(v string) *ModifyLifecyclePolicyRequest {
	s.LifecyclePolicyId = &v
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
