// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRecycleBinAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *GetRecycleBinAttributeRequest
	GetFileSystemId() *string
}

type GetRecycleBinAttributeRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s GetRecycleBinAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s GetRecycleBinAttributeRequest) GoString() string {
	return s.String()
}

func (s *GetRecycleBinAttributeRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetRecycleBinAttributeRequest) SetFileSystemId(v string) *GetRecycleBinAttributeRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetRecycleBinAttributeRequest) Validate() error {
	return dara.Validate(s)
}
