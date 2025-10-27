// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtocolMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteProtocolMountTargetRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteProtocolMountTargetRequest
	GetDryRun() *bool
	SetExportId(v string) *DeleteProtocolMountTargetRequest
	GetExportId() *string
	SetFileSystemId(v string) *DeleteProtocolMountTargetRequest
	GetFileSystemId() *string
	SetProtocolServiceId(v string) *DeleteProtocolMountTargetRequest
	GetProtocolServiceId() *string
}

type DeleteProtocolMountTargetRequest struct {
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
	// Specifies whether to perform only a dry run, without performing the actual request. The dry run checks parameter validity and prerequisites. The dry run does not delete the specified export directory or incur fees.
	//
	// Valid values:
	//
	// 	- true: performs only a dry run. The system checks the required parameters, request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, the specified export directory is deleted.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the export directory.
	//
	// This parameter is required.
	//
	// example:
	//
	// exp-123****
	ExportId *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-123****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the protocol service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ptc-123****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s DeleteProtocolMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtocolMountTargetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteProtocolMountTargetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteProtocolMountTargetRequest) GetExportId() *string {
	return s.ExportId
}

func (s *DeleteProtocolMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteProtocolMountTargetRequest) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *DeleteProtocolMountTargetRequest) SetClientToken(v string) *DeleteProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetDryRun(v bool) *DeleteProtocolMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetExportId(v string) *DeleteProtocolMountTargetRequest {
	s.ExportId = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetFileSystemId(v string) *DeleteProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) SetProtocolServiceId(v string) *DeleteProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *DeleteProtocolMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
