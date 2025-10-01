// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *DeleteAccessPointRequest
	GetAccessPointId() *string
	SetFileSystemId(v string) *DeleteAccessPointRequest
	GetFileSystemId() *string
}

type DeleteAccessPointRequest struct {
	// The ID of the access point.
	//
	// This parameter is required.
	//
	// example:
	//
	// ap-ie15yd****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAccessPointRequest) GoString() string {
	return s.String()
}

func (s *DeleteAccessPointRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DeleteAccessPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteAccessPointRequest) SetAccessPointId(v string) *DeleteAccessPointRequest {
	s.AccessPointId = &v
	return s
}

func (s *DeleteAccessPointRequest) SetFileSystemId(v string) *DeleteAccessPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
