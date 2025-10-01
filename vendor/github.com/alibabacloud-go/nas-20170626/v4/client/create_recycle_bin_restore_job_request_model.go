// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecycleBinRestoreJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRecycleBinRestoreJobRequest
	GetClientToken() *string
	SetFileId(v string) *CreateRecycleBinRestoreJobRequest
	GetFileId() *string
	SetFileSystemId(v string) *CreateRecycleBinRestoreJobRequest
	GetFileSystemId() *string
	SetTargetFileId(v string) *CreateRecycleBinRestoreJobRequest
	GetTargetFileId() *string
}

type CreateRecycleBinRestoreJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the file or directory that you want to restore.
	//
	// You can call the [ListRecycleBinJobs](https://help.aliyun.com/document_detail/264192.html) operation to query the value of the FileId parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// 04***08
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the directory to which the file is restored.
	//
	// This parameter is required.
	//
	// example:
	//
	// 13***15
	TargetFileId *string `json:"TargetFileId,omitempty" xml:"TargetFileId,omitempty"`
}

func (s CreateRecycleBinRestoreJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecycleBinRestoreJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinRestoreJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRecycleBinRestoreJobRequest) GetFileId() *string {
	return s.FileId
}

func (s *CreateRecycleBinRestoreJobRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateRecycleBinRestoreJobRequest) GetTargetFileId() *string {
	return s.TargetFileId
}

func (s *CreateRecycleBinRestoreJobRequest) SetClientToken(v string) *CreateRecycleBinRestoreJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetFileId(v string) *CreateRecycleBinRestoreJobRequest {
	s.FileId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetFileSystemId(v string) *CreateRecycleBinRestoreJobRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) SetTargetFileId(v string) *CreateRecycleBinRestoreJobRequest {
	s.TargetFileId = &v
	return s
}

func (s *CreateRecycleBinRestoreJobRequest) Validate() error {
	return dara.Validate(s)
}
