// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StopDataFlowRequest
	GetClientToken() *string
	SetDataFlowId(v string) *StopDataFlowRequest
	GetDataFlowId() *string
	SetDryRun(v bool) *StopDataFlowRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *StopDataFlowRequest
	GetFileSystemId() *string
}

type StopDataFlowRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
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
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no fileset quota is canceled and no fee is incurred.
	//
	// Valid value:
	//
	// 	- true: performs a dry run. The system checks the required parameters, request syntax, service limits, and available Apsara File Storage NAS (NAS) resources. Otherwise, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the FileSystemId parameter.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, the fileset quota is canceled.
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

func (s StopDataFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s StopDataFlowRequest) GoString() string {
	return s.String()
}

func (s *StopDataFlowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StopDataFlowRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *StopDataFlowRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *StopDataFlowRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *StopDataFlowRequest) SetClientToken(v string) *StopDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *StopDataFlowRequest) SetDataFlowId(v string) *StopDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *StopDataFlowRequest) SetDryRun(v bool) *StopDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *StopDataFlowRequest) SetFileSystemId(v string) *StopDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *StopDataFlowRequest) Validate() error {
	return dara.Validate(s)
}
