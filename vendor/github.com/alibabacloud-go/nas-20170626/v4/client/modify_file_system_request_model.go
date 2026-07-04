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
	// - The description must be 2 to 128 characters.
	//
	// - It must start with an uppercase or lowercase letter or a Chinese character, and cannot start with `http://` or `https://`.
	//
	// - It can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// NAS-test-1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the file system.
	//
	// - General-purpose NAS: For example, `31a8e4****`.
	//
	// - Extreme NAS: The ID must start with `extreme-`. For example, `extreme-0015****`.
	//
	// - CPFS: The ID must start with `cpfs-`. For example, `cpfs-125487****`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Additional options for the file system.
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
	EnableABE *bool `json:"EnableABE,omitempty" xml:"EnableABE,omitempty"`
	// Specifies whether to enable OpLock. Valid values:
	//
	// - true: Enables OpLock.
	//
	// - false: Disables OpLock.
	//
	// > This feature is available only for file systems that use the SMB protocol.
	//
	// example:
	//
	// true
	EnableOplock             *bool `json:"EnableOplock,omitempty" xml:"EnableOplock,omitempty"`
	VscAccessPointAccessOnly *bool `json:"VscAccessPointAccessOnly,omitempty" xml:"VscAccessPointAccessOnly,omitempty"`
}

func (s ModifyFileSystemRequestOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemRequestOptions) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemRequestOptions) GetEnableABE() *bool {
	return s.EnableABE
}

func (s *ModifyFileSystemRequestOptions) GetEnableOplock() *bool {
	return s.EnableOplock
}

func (s *ModifyFileSystemRequestOptions) GetVscAccessPointAccessOnly() *bool {
	return s.VscAccessPointAccessOnly
}

func (s *ModifyFileSystemRequestOptions) SetEnableABE(v bool) *ModifyFileSystemRequestOptions {
	s.EnableABE = &v
	return s
}

func (s *ModifyFileSystemRequestOptions) SetEnableOplock(v bool) *ModifyFileSystemRequestOptions {
	s.EnableOplock = &v
	return s
}

func (s *ModifyFileSystemRequestOptions) SetVscAccessPointAccessOnly(v bool) *ModifyFileSystemRequestOptions {
	s.VscAccessPointAccessOnly = &v
	return s
}

func (s *ModifyFileSystemRequestOptions) Validate() error {
	return dara.Validate(s)
}
