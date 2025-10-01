// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProtocolMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyProtocolMountTargetRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyProtocolMountTargetRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyProtocolMountTargetRequest
	GetDryRun() *bool
	SetExportId(v string) *ModifyProtocolMountTargetRequest
	GetExportId() *string
	SetFileSystemId(v string) *ModifyProtocolMountTargetRequest
	GetFileSystemId() *string
	SetProtocolServiceId(v string) *ModifyProtocolMountTargetRequest
	GetProtocolServiceId() *string
}

type ModifyProtocolMountTargetRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the export directory for the protocol service.
	//
	// Limits:
	//
	// 	- The description must be 2 to 128 characters in length.
	//
	// 	- The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// 	- The description can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. The dry run checks parameter validity and prerequisites. The dry run does not modify the specified export directory or incur fees.
	//
	// Valid values:
	//
	// 	- true: performs only a dry run. The system checks the required parameters, request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned.
	//
	// 	- false (default): performs a dry run and sends the request.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the export directory for the protocol service.
	//
	// This parameter is required.
	//
	// example:
	//
	// exp-19abf5beab8d****
	ExportId *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the protocol service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ptc-197ed6a00f2b****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s ModifyProtocolMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtocolMountTargetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyProtocolMountTargetRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyProtocolMountTargetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyProtocolMountTargetRequest) GetExportId() *string {
	return s.ExportId
}

func (s *ModifyProtocolMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyProtocolMountTargetRequest) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *ModifyProtocolMountTargetRequest) SetClientToken(v string) *ModifyProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetDescription(v string) *ModifyProtocolMountTargetRequest {
	s.Description = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetDryRun(v bool) *ModifyProtocolMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetExportId(v string) *ModifyProtocolMountTargetRequest {
	s.ExportId = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetFileSystemId(v string) *ModifyProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) SetProtocolServiceId(v string) *ModifyProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *ModifyProtocolMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
