// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemIds(v string) *CancelAutoSnapshotPolicyRequest
	GetFileSystemIds() *string
}

type CancelAutoSnapshotPolicyRequest struct {
	// The IDs of file systems.
	//
	// You can specify a maximum of 100 file system IDs. If you want to remove automatic snapshot policies from multiple file systems, separate the file system IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// extreme-233e6****,extreme-23vbp****,extreme-23vas****
	FileSystemIds *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
}

func (s CancelAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *CancelAutoSnapshotPolicyRequest) GetFileSystemIds() *string {
	return s.FileSystemIds
}

func (s *CancelAutoSnapshotPolicyRequest) SetFileSystemIds(v string) *CancelAutoSnapshotPolicyRequest {
	s.FileSystemIds = &v
	return s
}

func (s *CancelAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
