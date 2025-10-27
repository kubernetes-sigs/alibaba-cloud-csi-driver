// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRecycleBinDeleteJobRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateRecycleBinDeleteJobRequest
	GetClientToken() *string
	SetFileId(v string) *CreateRecycleBinDeleteJobRequest
	GetFileId() *string
	SetFileSystemId(v string) *CreateRecycleBinDeleteJobRequest
	GetFileSystemId() *string
}

type CreateRecycleBinDeleteJobRequest struct {
	// The client token that is used to ensure the idempotence of the request.
	//
	// You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
	//
	// > If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the file or directory that you want to permanently delete.
	//
	// You can call the [ListRecycledDirectoriesAndFiles](https://help.aliyun.com/document_detail/264193.html) operation to query the value of the FileId parameter.
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
}

func (s CreateRecycleBinDeleteJobRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateRecycleBinDeleteJobRequest) GoString() string {
	return s.String()
}

func (s *CreateRecycleBinDeleteJobRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateRecycleBinDeleteJobRequest) GetFileId() *string {
	return s.FileId
}

func (s *CreateRecycleBinDeleteJobRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateRecycleBinDeleteJobRequest) SetClientToken(v string) *CreateRecycleBinDeleteJobRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateRecycleBinDeleteJobRequest) SetFileId(v string) *CreateRecycleBinDeleteJobRequest {
	s.FileId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobRequest) SetFileSystemId(v string) *CreateRecycleBinDeleteJobRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateRecycleBinDeleteJobRequest) Validate() error {
	return dara.Validate(s)
}
