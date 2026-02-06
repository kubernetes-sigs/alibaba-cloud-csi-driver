// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilesystemsAssociatedHpnZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilesystems(v []*DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) *DescribeFilesystemsAssociatedHpnZonesResponseBody
	GetFilesystems() []*DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems
	SetRequestId(v string) *DescribeFilesystemsAssociatedHpnZonesResponseBody
	GetRequestId() *string
}

type DescribeFilesystemsAssociatedHpnZonesResponseBody struct {
	Filesystems []*DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems `json:"Filesystems,omitempty" xml:"Filesystems,omitempty" type:"Repeated"`
	// example:
	//
	// 98696EF0-1607-4E9D-B01D-F20930B6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFilesystemsAssociatedHpnZonesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsAssociatedHpnZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBody) GetFilesystems() []*DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems {
	return s.Filesystems
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBody) SetFilesystems(v []*DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) *DescribeFilesystemsAssociatedHpnZonesResponseBody {
	s.Filesystems = v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBody) SetRequestId(v string) *DescribeFilesystemsAssociatedHpnZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBody) Validate() error {
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

type DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems struct {
	AssociatedHpnZones []*string `json:"AssociatedHpnZones,omitempty" xml:"AssociatedHpnZones,omitempty" type:"Repeated"`
	// example:
	//
	// bmcpfs-290t15yn4uo8lid****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// example:
	//
	// cn-hangzhou-k
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) GetAssociatedHpnZones() []*string {
	return s.AssociatedHpnZones
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) SetAssociatedHpnZones(v []*string) *DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems {
	s.AssociatedHpnZones = v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) SetFileSystemId(v string) *DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems {
	s.FileSystemId = &v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) SetZoneId(v string) *DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems {
	s.ZoneId = &v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesResponseBodyFilesystems) Validate() error {
	return dara.Validate(s)
}
