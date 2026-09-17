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
	SetTag(v []*CreateCpfsAccessPointRequestTag) *CreateCpfsAccessPointRequest
	GetTag() []*CreateCpfsAccessPointRequestTag
}

type CreateCpfsAccessPointRequest struct {
	// The description of the access point.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter.
	//
	// - The description cannot start with http:// or https://.
	//
	// - The description can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file system ID.
	//
	// - Cloud Parallel File Storage (CPFS): must start with `cpfs-`, such as cpfs-125487\\*\\*\\*\\*.
	//
	// - CPFS for Lingjun: must start with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The root directory of the access point. Default value: "/".
	RootDirectory *CreateCpfsAccessPointRequestRootDirectory `json:"RootDirectory,omitempty" xml:"RootDirectory,omitempty" type:"Struct"`
	// The list of tags for the CPFS access point.
	Tag []*CreateCpfsAccessPointRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
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

func (s *CreateCpfsAccessPointRequest) GetTag() []*CreateCpfsAccessPointRequestTag {
	return s.Tag
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

func (s *CreateCpfsAccessPointRequest) SetTag(v []*CreateCpfsAccessPointRequestTag) *CreateCpfsAccessPointRequest {
	s.Tag = v
	return s
}

func (s *CreateCpfsAccessPointRequest) Validate() error {
	if s.RootDirectory != nil {
		if err := s.RootDirectory.Validate(); err != nil {
			return err
		}
	}
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCpfsAccessPointRequestRootDirectory struct {
	// The root directory of the access point. The value must start and end with a forward slash (/).
	//
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

type CreateCpfsAccessPointRequestTag struct {
	// The tag key.
	//
	// example:
	//
	// testKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// testValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCpfsAccessPointRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCpfsAccessPointRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCpfsAccessPointRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCpfsAccessPointRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCpfsAccessPointRequestTag) SetKey(v string) *CreateCpfsAccessPointRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCpfsAccessPointRequestTag) SetValue(v string) *CreateCpfsAccessPointRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCpfsAccessPointRequestTag) Validate() error {
	return dara.Validate(s)
}
