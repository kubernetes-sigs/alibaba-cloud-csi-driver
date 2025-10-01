// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNfsAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeNfsAclRequest
	GetFileSystemId() *string
}

type DescribeNfsAclRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 206614xxxx
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DescribeNfsAclRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNfsAclRequest) GoString() string {
	return s.String()
}

func (s *DescribeNfsAclRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeNfsAclRequest) SetFileSystemId(v string) *DescribeNfsAclRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeNfsAclRequest) Validate() error {
	return dara.Validate(s)
}
