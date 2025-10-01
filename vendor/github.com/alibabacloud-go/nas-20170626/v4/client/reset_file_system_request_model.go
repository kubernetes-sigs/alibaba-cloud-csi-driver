// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResetFileSystemRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *ResetFileSystemRequest
	GetFileSystemId() *string
	SetSnapshotId(v string) *ResetFileSystemRequest
	GetSnapshotId() *string
}

type ResetFileSystemRequest struct {
	// The ID of the advanced Extreme NAS file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// extreme-012dd****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The snapshot ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// s-extreme-snapsho****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s ResetFileSystemRequest) String() string {
	return dara.Prettify(s)
}

func (s ResetFileSystemRequest) GoString() string {
	return s.String()
}

func (s *ResetFileSystemRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ResetFileSystemRequest) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *ResetFileSystemRequest) SetFileSystemId(v string) *ResetFileSystemRequest {
	s.FileSystemId = &v
	return s
}

func (s *ResetFileSystemRequest) SetSnapshotId(v string) *ResetFileSystemRequest {
	s.SnapshotId = &v
	return s
}

func (s *ResetFileSystemRequest) Validate() error {
	return dara.Validate(s)
}
