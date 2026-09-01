// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFileSystemShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ModifyFileSystemShrinkRequest
	GetDescription() *string
	SetFileSystemId(v string) *ModifyFileSystemShrinkRequest
	GetFileSystemId() *string
	SetOptionsShrink(v string) *ModifyFileSystemShrinkRequest
	GetOptionsShrink() *string
}

type ModifyFileSystemShrinkRequest struct {
	// The file system description.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter or Chinese character and cannot start with `http://` or `https://`.
	//
	// - The description can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// NAS-test-1
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file system ID.
	//
	// - General-purpose NAS: `31a8e4****`.
	//
	// - Extreme NAS: must start with `extreme-`, for example, `extreme-0015****`.
	//
	// - CPFS: must start with `cpfs-`, for example, `cpfs-125487****`.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The options.
	OptionsShrink *string `json:"Options,omitempty" xml:"Options,omitempty"`
}

func (s ModifyFileSystemShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFileSystemShrinkRequest) GoString() string {
	return s.String()
}

func (s *ModifyFileSystemShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFileSystemShrinkRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyFileSystemShrinkRequest) GetOptionsShrink() *string {
	return s.OptionsShrink
}

func (s *ModifyFileSystemShrinkRequest) SetDescription(v string) *ModifyFileSystemShrinkRequest {
	s.Description = &v
	return s
}

func (s *ModifyFileSystemShrinkRequest) SetFileSystemId(v string) *ModifyFileSystemShrinkRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFileSystemShrinkRequest) SetOptionsShrink(v string) *ModifyFileSystemShrinkRequest {
	s.OptionsShrink = &v
	return s
}

func (s *ModifyFileSystemShrinkRequest) Validate() error {
	return dara.Validate(s)
}
