// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFilesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *GetFilesetRequest
	GetFileSystemId() *string
	SetFsetId(v string) *GetFilesetRequest
	GetFsetId() *string
}

type GetFilesetRequest struct {
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-125487\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for Lingjun file systems must start with `bmcpfs-`. Example: bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The fileset ID.
	//
	// >  This parameter is required for CPFS file systems.
	//
	// This parameter is required.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
}

func (s GetFilesetRequest) String() string {
	return dara.Prettify(s)
}

func (s GetFilesetRequest) GoString() string {
	return s.String()
}

func (s *GetFilesetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *GetFilesetRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *GetFilesetRequest) SetFileSystemId(v string) *GetFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *GetFilesetRequest) SetFsetId(v string) *GetFilesetRequest {
	s.FsetId = &v
	return s
}

func (s *GetFilesetRequest) Validate() error {
	return dara.Validate(s)
}
