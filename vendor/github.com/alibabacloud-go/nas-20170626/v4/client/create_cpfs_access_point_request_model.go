// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCpfsAccessPointRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateCpfsAccessPointRequest
	GetDescription() *string
	SetFileSystemId(v string) *CreateCpfsAccessPointRequest
	GetFileSystemId() *string
	SetRegionId(v string) *CreateCpfsAccessPointRequest
	GetRegionId() *string
	SetRootDirectory(v *CreateCpfsAccessPointRequestRootDirectory) *CreateCpfsAccessPointRequest
	GetRootDirectory() *CreateCpfsAccessPointRequestRootDirectory
}

type CreateCpfsAccessPointRequest struct {
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
	RegionId      *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	RootDirectory *CreateCpfsAccessPointRequestRootDirectory `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty" type:"Struct"`
}

func (s CreateCpfsAccessPointRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCpfsAccessPointRequest) GoString() string {
	return s.String()
}

func (s *CreateCpfsAccessPointRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCpfsAccessPointRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateCpfsAccessPointRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCpfsAccessPointRequest) GetRootDirectory() *CreateCpfsAccessPointRequestRootDirectory {
	return s.RootDirectory
}

func (s *CreateCpfsAccessPointRequest) SetDescription(v string) *CreateCpfsAccessPointRequest {
	s.Description = &v
	return s
}

func (s *CreateCpfsAccessPointRequest) SetFileSystemId(v string) *CreateCpfsAccessPointRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateCpfsAccessPointRequest) SetRegionId(v string) *CreateCpfsAccessPointRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCpfsAccessPointRequest) SetRootDirectory(v *CreateCpfsAccessPointRequestRootDirectory) *CreateCpfsAccessPointRequest {
	s.RootDirectory = v
	return s
}

func (s *CreateCpfsAccessPointRequest) Validate() error {
	if s.RootDirectory != nil {
		if err := s.RootDirectory.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateCpfsAccessPointRequestRootDirectory struct {
	// example:
	//
	// /test/
	RootPath *string `json:"RootPath,omitempty" xml:"RootPath,omitempty"`
}

func (s CreateCpfsAccessPointRequestRootDirectory) String() string {
	return dara.Prettify(s)
}

func (s CreateCpfsAccessPointRequestRootDirectory) GoString() string {
	return s.String()
}

func (s *CreateCpfsAccessPointRequestRootDirectory) GetRootPath() *string {
	return s.RootPath
}

func (s *CreateCpfsAccessPointRequestRootDirectory) SetRootPath(v string) *CreateCpfsAccessPointRequestRootDirectory {
	s.RootPath = &v
	return s
}

func (s *CreateCpfsAccessPointRequestRootDirectory) Validate() error {
	return dara.Validate(s)
}
