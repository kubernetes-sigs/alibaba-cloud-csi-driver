// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilesystemsAssociatedHpnZonesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilesystemsShrink(v string) *DescribeFilesystemsAssociatedHpnZonesShrinkRequest
	GetFilesystemsShrink() *string
	SetRegionId(v string) *DescribeFilesystemsAssociatedHpnZonesShrinkRequest
	GetRegionId() *string
}

type DescribeFilesystemsAssociatedHpnZonesShrinkRequest struct {
	FilesystemsShrink *string `json:"Filesystems,omitempty" xml:"Filesystems,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeFilesystemsAssociatedHpnZonesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilesystemsAssociatedHpnZonesShrinkRequest) GoString() string {
	return s.String()
}

func (s *DescribeFilesystemsAssociatedHpnZonesShrinkRequest) GetFilesystemsShrink() *string {
	return s.FilesystemsShrink
}

func (s *DescribeFilesystemsAssociatedHpnZonesShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeFilesystemsAssociatedHpnZonesShrinkRequest) SetFilesystemsShrink(v string) *DescribeFilesystemsAssociatedHpnZonesShrinkRequest {
	s.FilesystemsShrink = &v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesShrinkRequest) SetRegionId(v string) *DescribeFilesystemsAssociatedHpnZonesShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeFilesystemsAssociatedHpnZonesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
