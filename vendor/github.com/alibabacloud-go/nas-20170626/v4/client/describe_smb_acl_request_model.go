// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSmbAclRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *DescribeSmbAclRequest
	GetFileSystemId() *string
}

type DescribeSmbAclRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DescribeSmbAclRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSmbAclRequest) GoString() string {
	return s.String()
}

func (s *DescribeSmbAclRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeSmbAclRequest) SetFileSystemId(v string) *DescribeSmbAclRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeSmbAclRequest) Validate() error {
	return dara.Validate(s)
}
