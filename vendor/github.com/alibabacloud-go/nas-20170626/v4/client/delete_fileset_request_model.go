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
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run for this deletion request.
	//
	// A dry run checks parameter validity and resource availability without actually deleting the instance.
	//
	// Valid values:
	//
	// - true: Sends a check request without deleting the instance. The check items include whether required parameters are specified, the request format, and business limitations. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is directly deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - CPFS: The ID must start with `cpfs-`, such as cpfs-099394bd928c****.
	//
	// - CPFS for Lingjun: The ID must start with `bmcpfs-`, such as bmcpfs-290w65p03ok64ya****.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
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
