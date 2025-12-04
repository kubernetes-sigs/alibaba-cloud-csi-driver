// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyFileSystemRequest
	GetDescription() *string
	SetFileSystemId(v string) *ModifyFileSystemRequest
	GetFileSystemId() *string
	SetOptions(v *ModifyFileSystemRequestOptions) *ModifyFileSystemRequest
	GetOptions() *ModifyFileSystemRequestOptions
}

type ModifyFileSystemRequest struct {
	// The description of the file system.
	//
	// Limits:
	//
	// 	- The description must be 2 to 128 characters in length.
	//
	// 	- It must start with a letter but cannot start with `http://` or `https://`.
	//
	// 	- The description can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// NAS-test-1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the file system.
	//
	// 	- Sample ID of a General-purpose NAS file system: `31a8e4****`.
	//
	// 	- The IDs of Extreme NAS file systems must start with `extreme-`. Example: `extreme-0015****`.
	//
	// 	- The IDs of Cloud Paralleled File System (CPFS) file systems must start with `cpfs-`. Example: `cpfs-125487****`.
	//
	// >CPFS file systems are available only on the China site (aliyun.com).
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The options.
	Options *ModifyFileSystemRequestOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Struct"`
}

func (s ModifyFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFileSystemRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyFileSystemRequest) GetOptions() *ModifyFileSystemRequestOptions {
	return s.Options
}

func (s *ModifyFileSystemRequest) SetDescription(v string) *ModifyFileSystemRequest {
	s.Description = &v
	return s
}

func (s *ModifyFileSystemRequest) SetFileSystemId(v string) *ModifyFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFileSystemRequest) SetOptions(v *ModifyFileSystemRequestOptions) *ModifyFileSystemRequest {
	s.Options = v
	return s
}

func (s *ModifyFileSystemRequest) Validate() error {
	if s.Options != nil {
		if err := s.Options.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyFileSystemRequestOptions struct {
	// Specifies whether to enable the oplock feature. Valid values:
	//
	// 	- true: enables the feature.
	//
	// 	- false: disables the feature.
	//
	// >  Only Server Message Block (SMB) file systems support this feature.
	//
	// example:
	//
	// true
	EnableOplock *bool `json:"EnableOplock,omitempty" xml:"EnableOplock,omitempty"`
}

func (s ModifyFileSystemRequestOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemRequestOptions) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequestOptions) GetEnableOplock() *bool {
	return s.EnableOplock
}

func (s *ModifyFileSystemRequestOptions) SetEnableOplock(v bool) *ModifyFileSystemRequestOptions {
	s.EnableOplock = &v
	return s
}

func (s *ModifyFileSystemRequestOptions) Validate() error {
	return dara.Validate(s)
}
