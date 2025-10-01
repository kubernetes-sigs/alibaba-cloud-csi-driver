// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableAndCleanRecycleBinRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DisableAndCleanRecycleBinRequest
	GetFileSystemId() *string
}

type DisableAndCleanRecycleBinRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DisableAndCleanRecycleBinRequest) String() string {
	return dara.Prettify(s)
}

func (s DisableAndCleanRecycleBinRequest) GoString() string {
	return s.String()
}

func (s *DisableAndCleanRecycleBinRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DisableAndCleanRecycleBinRequest) SetFileSystemId(v string) *DisableAndCleanRecycleBinRequest {
	s.FileSystemId = &v
	return s
}

func (s *DisableAndCleanRecycleBinRequest) Validate() error {
	return dara.Validate(s)
}
