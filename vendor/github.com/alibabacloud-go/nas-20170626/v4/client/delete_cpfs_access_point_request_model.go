// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCpfsAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *DeleteCpfsAccessPointRequest
	GetAccessPointId() *string
	SetFileSystemId(v string) *DeleteCpfsAccessPointRequest
	GetFileSystemId() *string
	SetRegionId(v string) *DeleteCpfsAccessPointRequest
	GetRegionId() *string
}

type DeleteCpfsAccessPointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ap-ie15yd****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DeleteCpfsAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCpfsAccessPointRequest) GoString() string {
	return s.String()
}

func (s *DeleteCpfsAccessPointRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DeleteCpfsAccessPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteCpfsAccessPointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCpfsAccessPointRequest) SetAccessPointId(v string) *DeleteCpfsAccessPointRequest {
	s.AccessPointId = &v
	return s
}

func (s *DeleteCpfsAccessPointRequest) SetFileSystemId(v string) *DeleteCpfsAccessPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteCpfsAccessPointRequest) SetRegionId(v string) *DeleteCpfsAccessPointRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCpfsAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
