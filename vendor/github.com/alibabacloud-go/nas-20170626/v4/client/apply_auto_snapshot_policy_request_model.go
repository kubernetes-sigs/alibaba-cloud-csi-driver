// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyAutoSnapshotPolicyRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoSnapshotPolicyId(v string) *ApplyAutoSnapshotPolicyRequest
	GetAutoSnapshotPolicyId() *string
	SetFileSystemIds(v string) *ApplyAutoSnapshotPolicyRequest
	GetFileSystemIds() *string
}

type ApplyAutoSnapshotPolicyRequest struct {
	// The ID of the automatic snapshot policy.
	//
	// This parameter is required.
	//
	// example:
	//
	// sp-extreme-233e6****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The IDs of advanced Extreme NAS file systems.
	//
	// You can specify a maximum of 100 file system IDs at a time. If you want to apply an automatic snapshot policy to multiple file systems, separate the file system IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// extreme-233e6****,extreme -23vbp****,extreme -23vas****
	FileSystemIds *string `json:"FileSystemIds,omitempty" xml:"FileSystemIds,omitempty"`
}

func (s ApplyAutoSnapshotPolicyRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyAutoSnapshotPolicyRequest) GoString() string {
	return s.String()
}

func (s *ApplyAutoSnapshotPolicyRequest) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *ApplyAutoSnapshotPolicyRequest) GetFileSystemIds() *string {
	return s.FileSystemIds
}

func (s *ApplyAutoSnapshotPolicyRequest) SetAutoSnapshotPolicyId(v string) *ApplyAutoSnapshotPolicyRequest {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) SetFileSystemIds(v string) *ApplyAutoSnapshotPolicyRequest {
	s.FileSystemIds = &v
	return s
}

func (s *ApplyAutoSnapshotPolicyRequest) Validate() error {
	return dara.Validate(s)
}
