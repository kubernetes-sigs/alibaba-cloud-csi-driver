// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeBlackListClientsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientIP(v string) *DescribeBlackListClientsRequest
	GetClientIP() *string
	SetFileSystemId(v string) *DescribeBlackListClientsRequest
	GetFileSystemId() *string
	SetRegionId(v string) *DescribeBlackListClientsRequest
	GetRegionId() *string
}

type DescribeBlackListClientsRequest struct {
	// The IP address of the client.
	//
	// example:
	//
	// 192.168.0.0
	ClientIP *string `json:"ClientIP,omitempty" xml:"ClientIP,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-123458****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the region where the file system resides.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeBlackListClientsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeBlackListClientsRequest) GoString() string {
	return s.String()
}

func (s *DescribeBlackListClientsRequest) GetClientIP() *string {
	return s.ClientIP
}

func (s *DescribeBlackListClientsRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeBlackListClientsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeBlackListClientsRequest) SetClientIP(v string) *DescribeBlackListClientsRequest {
	s.ClientIP = &v
	return s
}

func (s *DescribeBlackListClientsRequest) SetFileSystemId(v string) *DescribeBlackListClientsRequest {
	s.FileSystemId = &v
	return s
}

func (s *DescribeBlackListClientsRequest) SetRegionId(v string) *DescribeBlackListClientsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeBlackListClientsRequest) Validate() error {
	return dara.Validate(s)
}
