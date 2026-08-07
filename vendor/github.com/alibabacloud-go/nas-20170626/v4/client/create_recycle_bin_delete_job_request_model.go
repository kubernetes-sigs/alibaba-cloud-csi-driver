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
	// Ensures the idempotence of the request.
	//
	// Generate a parameter value from your client to ensure that the value is unique among different requests. The value of ClientToken can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// >If you do not specify ClientToken, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The FileId of the file or directory that you want to permanently delete.
	//
	// You can call the [ListRecycledDirectoriesAndFiles](https://help.aliyun.com/document_detail/2412174.html) operation to query the FileId of deleted data.
	//
	// This parameter is required.
	//
	// example:
	//
	// 104
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The file system ID.
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
