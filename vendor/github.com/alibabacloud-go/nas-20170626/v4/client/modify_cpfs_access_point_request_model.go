// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCpfsAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessPointId(v string) *ModifyCpfsAccessPointRequest
	GetAccessPointId() *string
	SetDescription(v string) *ModifyCpfsAccessPointRequest
	GetDescription() *string
	SetFileSystemId(v string) *ModifyCpfsAccessPointRequest
	GetFileSystemId() *string
	SetRegionId(v string) *ModifyCpfsAccessPointRequest
	GetRegionId() *string
}

type ModifyCpfsAccessPointRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ap-ie15yd****
	AccessPointId *string `json:"AccessPointId,omitempty" xml:"AccessPointId,omitempty"`
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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

func (s ModifyCpfsAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCpfsAccessPointRequest) GoString() string {
	return s.String()
}

func (s *ModifyCpfsAccessPointRequest) GetAccessPointId() *string {
	return s.AccessPointId
}

func (s *ModifyCpfsAccessPointRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCpfsAccessPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyCpfsAccessPointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCpfsAccessPointRequest) SetAccessPointId(v string) *ModifyCpfsAccessPointRequest {
	s.AccessPointId = &v
	return s
}

func (s *ModifyCpfsAccessPointRequest) SetDescription(v string) *ModifyCpfsAccessPointRequest {
	s.Description = &v
	return s
}

func (s *ModifyCpfsAccessPointRequest) SetFileSystemId(v string) *ModifyCpfsAccessPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyCpfsAccessPointRequest) SetRegionId(v string) *ModifyCpfsAccessPointRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCpfsAccessPointRequest) Validate() error {
	return dara.Validate(s)
}
