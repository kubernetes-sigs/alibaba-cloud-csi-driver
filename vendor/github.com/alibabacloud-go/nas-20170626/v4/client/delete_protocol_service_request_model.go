// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteProtocolServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteProtocolServiceRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteProtocolServiceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *DeleteProtocolServiceRequest
	GetFileSystemId() *string
	SetProtocolServiceId(v string) *DeleteProtocolServiceRequest
	GetProtocolServiceId() *string
}

type DeleteProtocolServiceRequest struct {
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
	// Specifies whether to perform only a dry run, without performing the actual request. The dry run checks parameter validity and prerequisites. The dry run does not delete the specified protocol service.
	//
	// Valid values:
	//
	// 	- true: performs only a dry run. The system checks the required parameters, request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, the specified protocol service is deleted.
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

func (s DeleteProtocolServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *DeleteProtocolServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteProtocolServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteProtocolServiceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteProtocolServiceRequest) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *DeleteProtocolServiceRequest) SetClientToken(v string) *DeleteProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetDryRun(v bool) *DeleteProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetFileSystemId(v string) *DeleteProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteProtocolServiceRequest) SetProtocolServiceId(v string) *DeleteProtocolServiceRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *DeleteProtocolServiceRequest) Validate() error {
	return dara.Validate(s)
}
