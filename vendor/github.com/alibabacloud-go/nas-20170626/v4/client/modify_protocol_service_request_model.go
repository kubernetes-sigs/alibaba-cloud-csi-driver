// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProtocolServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyProtocolServiceRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyProtocolServiceRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyProtocolServiceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *ModifyProtocolServiceRequest
	GetFileSystemId() *string
	SetProtocolServiceId(v string) *ModifyProtocolServiceRequest
	GetProtocolServiceId() *string
}

type ModifyProtocolServiceRequest struct {
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
	// The description of the protocol service.
	//
	// Limits:
	//
	// 	- The description must be 2 to 128 characters in length.
	//
	// 	- The description must start with a letter and cannot start with `http://` or `https://`.
	//
	// 	- The description can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// if can be null:
	// false
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. The dry run checks parameter validity and prerequisites. The dry run does not modify a file system or incur fees.
	//
	// Valid values:
	//
	// 	- true: performs only a dry run and does not modify the protocol service. The system checks the request format, service limits, prerequisites, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, a 200 HTTP status code is returned.
	//
	// 	- false (default): performs a dry run and performs the actual request. If the request passes the dry run, the service protocol is modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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

func (s ModifyProtocolServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtocolServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyProtocolServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyProtocolServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyProtocolServiceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyProtocolServiceRequest) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *ModifyProtocolServiceRequest) SetClientToken(v string) *ModifyProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetDescription(v string) *ModifyProtocolServiceRequest {
	s.Description = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetDryRun(v bool) *ModifyProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetFileSystemId(v string) *ModifyProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetProtocolServiceId(v string) *ModifyProtocolServiceRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *ModifyProtocolServiceRequest) Validate() error {
	return dara.Validate(s)
}
