// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLifecycleRetrieveJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileSystemId(v string) *CreateLifecycleRetrieveJobRequest
	GetFileSystemId() *string
	SetPaths(v []*string) *CreateLifecycleRetrieveJobRequest
	GetPaths() []*string
	SetStorageType(v string) *CreateLifecycleRetrieveJobRequest
	GetStorageType() *string
}

type CreateLifecycleRetrieveJobRequest struct {
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The directories or files that you want to retrieve. You can specify a maximum of 10 paths.
	//
	// This parameter is required.
	//
	// example:
	//
	// Paths.1=/pathway/doc1,Paths.2=/pathway/doc2
	Paths []*string `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// The storage class.
	//
	// 	- InfrequentAccess (default): the Infrequent Access (IA) storage class.
	//
	// 	- Archive: the Archive storage class.
	//
	// example:
	//
	// InfrequentAccess
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s CreateLifecycleRetrieveJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLifecycleRetrieveJobRequest) GoString() string {
	return s.String()
}

func (s *CreateLifecycleRetrieveJobRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateLifecycleRetrieveJobRequest) GetPaths() []*string {
	return s.Paths
}

func (s *CreateLifecycleRetrieveJobRequest) GetStorageType() *string {
	return s.StorageType
}

func (s *CreateLifecycleRetrieveJobRequest) SetFileSystemId(v string) *CreateLifecycleRetrieveJobRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateLifecycleRetrieveJobRequest) SetPaths(v []*string) *CreateLifecycleRetrieveJobRequest {
	s.Paths = v
	return s
}

func (s *CreateLifecycleRetrieveJobRequest) SetStorageType(v string) *CreateLifecycleRetrieveJobRequest {
	s.StorageType = &v
	return s
}

func (s *CreateLifecycleRetrieveJobRequest) Validate() error {
	return dara.Validate(s)
}
