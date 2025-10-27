// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *DescribeAccessPointRequest
	GetAccessPointId() *string
	SetFileSystemId(v string) *DescribeAccessPointRequest
	GetFileSystemId() *string
}

type DescribeAccessPointRequest struct {
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
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DescribeAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeAccessPointRequest) GoString() string {
	return s.String()
}

func (s *DescribeAccessPointRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *DescribeAccessPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeAccessPointRequest) SetAccessPointId(v string) *DescribeAccessPointRequest {
	s.AccessPointId = &v
	return s
}

func (s *DescribeAccessPointRequest) SetFileSystemId(v string) *DescribeAccessPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
