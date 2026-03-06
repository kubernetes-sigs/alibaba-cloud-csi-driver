// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilesystemsAssociatedHpnZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilesystems(v []*DescribeFilesystemsAssociatedHpnZonesRequestFilesystems) *DescribeFilesystemsAssociatedHpnZonesRequest
	GetFilesystems() []*DescribeFilesystemsAssociatedHpnZonesRequestFilesystems
	SetRegionId(v string) *DescribeFilesystemsAssociatedHpnZonesRequest
	GetRegionId() *string
}

type DescribeFilesystemsAssociatedHpnZonesRequest struct {
	Filesystems []*DescribeFilesystemsAssociatedHpnZonesRequestFilesystems `json:"Filesystems,omitempty" xml:"Filesystems,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFilesystemsAssociatedHpnZonesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsAssociatedHpnZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsAssociatedHpnZonesRequest) GetFilesystems() []*DescribeFilesystemsAssociatedHpnZonesRequestFilesystems {
	return s.Filesystems
}

func (s *DescribeFilesystemsAssociatedHpnZonesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFilesystemsAssociatedHpnZonesRequest) SetFilesystems(v []*DescribeFilesystemsAssociatedHpnZonesRequestFilesystems) *DescribeFilesystemsAssociatedHpnZonesRequest {
	s.Filesystems = v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesRequest) SetRegionId(v string) *DescribeFilesystemsAssociatedHpnZonesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesRequest) Validate() error {
	if s.Filesystems != nil {
		for _, item := range s.Filesystems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeFilesystemsAssociatedHpnZonesRequestFilesystems struct {
	// example:
	//
	// bmcpfs-290t15yn4uo8lid****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DescribeFilesystemsAssociatedHpnZonesRequestFilesystems) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsAssociatedHpnZonesRequestFilesystems) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsAssociatedHpnZonesRequestFilesystems) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFilesystemsAssociatedHpnZonesRequestFilesystems) SetFileSystemId(v string) *DescribeFilesystemsAssociatedHpnZonesRequestFilesystems {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesRequestFilesystems) Validate() error {
	return dara.Validate(s)
}
