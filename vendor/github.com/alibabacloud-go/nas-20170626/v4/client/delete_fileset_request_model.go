// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteFilesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteFilesetRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteFilesetRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *DeleteFilesetRequest
	GetFileSystemId() *string
	SetFsetId(v string) *DeleteFilesetRequest
	GetFsetId() *string
}

type DeleteFilesetRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no fileset is deleted.
	//
	// Valid values:
	//
	// 	- true: performs only a dry run. The system checks the required parameters, request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, the fileset is deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-099394bd928c\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for LINGJUN file systems must start with `bmcpfs-`. Example: bmcpfs-290w65p03ok64ya\\*\\*\\*\\*.
	//
	// >  CPFS is not supported on the international site.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The fileset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
}

func (s DeleteFilesetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteFilesetRequest) GoString() string {
	return s.String()
}

func (s *DeleteFilesetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteFilesetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteFilesetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteFilesetRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *DeleteFilesetRequest) SetClientToken(v string) *DeleteFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteFilesetRequest) SetDryRun(v bool) *DeleteFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteFilesetRequest) SetFileSystemId(v string) *DeleteFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteFilesetRequest) SetFsetId(v string) *DeleteFilesetRequest {
	s.FsetId = &v
	return s
}

func (s *DeleteFilesetRequest) Validate() error {
	return dara.Validate(s)
}
