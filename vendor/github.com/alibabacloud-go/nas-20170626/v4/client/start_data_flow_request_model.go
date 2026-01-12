// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDataFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartDataFlowRequest
	GetClientToken() *string
	SetDataFlowId(v string) *StartDataFlowRequest
	GetDataFlowId() *string
	SetDryRun(v bool) *StartDataFlowRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *StartDataFlowRequest
	GetFileSystemId() *string
}

type StartDataFlowRequest struct {
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
	// The ID of the dataflow.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. The dry run does not enable the specified dataflow or incur fees.
	//
	// Valid values:
	//
	// 	- true: performs only a dry run. The system checks the required parameters, request syntax, service limits, and available NAS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, the specified dataflow is enabled.
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
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s StartDataFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s StartDataFlowRequest) GoString() string {
	return s.String()
}

func (s *StartDataFlowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartDataFlowRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *StartDataFlowRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StartDataFlowRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *StartDataFlowRequest) SetClientToken(v string) *StartDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *StartDataFlowRequest) SetDataFlowId(v string) *StartDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *StartDataFlowRequest) SetDryRun(v bool) *StartDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *StartDataFlowRequest) SetFileSystemId(v string) *StartDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *StartDataFlowRequest) Validate() error {
	return dara.Validate(s)
}
