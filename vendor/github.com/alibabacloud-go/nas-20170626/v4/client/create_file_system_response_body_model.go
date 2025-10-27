// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileSystemResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *CreateFileSystemResponseBody
	GetFileSystemId() *string
	SetRequestId(v string) *CreateFileSystemResponseBody
	GetRequestId() *string
}

type CreateFileSystemResponseBody struct {
	// The ID of the file system that is created.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateFileSystemResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateFileSystemResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileSystemResponseBody) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateFileSystemResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateFileSystemResponseBody) SetFileSystemId(v string) *CreateFileSystemResponseBody {
	s.FileSystemId = &v
	return s
}

func (s *CreateFileSystemResponseBody) SetRequestId(v string) *CreateFileSystemResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateFileSystemResponseBody) Validate() error {
	return dara.Validate(s)
}
