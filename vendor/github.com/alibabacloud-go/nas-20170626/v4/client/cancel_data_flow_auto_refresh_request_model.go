// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelDataFlowAutoRefreshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelDataFlowAutoRefreshRequest
	GetClientToken() *string
	SetDataFlowId(v string) *CancelDataFlowAutoRefreshRequest
	GetDataFlowId() *string
	SetDryRun(v bool) *CancelDataFlowAutoRefreshRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CancelDataFlowAutoRefreshRequest
	GetFileSystemId() *string
	SetRefreshPath(v string) *CancelDataFlowAutoRefreshRequest
	GetRefreshPath() *string
}

type CancelDataFlowAutoRefreshRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The value of RequestId may be different for each API request.
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
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no file system is created and no fee is incurred.
	//
	// Valid values:
	//
	// 	- true: performs a dry run. The system checks the request format, service limits, prerequisites, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the DataFlowld parameter.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, a file system is created.
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
	// The directory for which you want to cancel AutoRefresh configurations.
	//
	// Limits:
	//
	// 	- The directory must be 2 to 1,024 characters in length.
	//
	// 	- The directory must be encoded in UTF-8.
	//
	// 	- The directory must start and end with a forward slash (/).
	//
	// >  The directory must be an existing directory in the CPFS file system and must be in a fileset where the dataflow is enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// /prefix1/prefix2/
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s CancelDataFlowAutoRefreshRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelDataFlowAutoRefreshRequest) GoString() string {
	return s.String()
}

func (s *CancelDataFlowAutoRefreshRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelDataFlowAutoRefreshRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *CancelDataFlowAutoRefreshRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CancelDataFlowAutoRefreshRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CancelDataFlowAutoRefreshRequest) GetRefreshPath() *string {
	return s.RefreshPath
}

func (s *CancelDataFlowAutoRefreshRequest) SetClientToken(v string) *CancelDataFlowAutoRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetDataFlowId(v string) *CancelDataFlowAutoRefreshRequest {
	s.DataFlowId = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetDryRun(v bool) *CancelDataFlowAutoRefreshRequest {
	s.DryRun = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetFileSystemId(v string) *CancelDataFlowAutoRefreshRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) SetRefreshPath(v string) *CancelDataFlowAutoRefreshRequest {
	s.RefreshPath = &v
	return s
}

func (s *CancelDataFlowAutoRefreshRequest) Validate() error {
	return dara.Validate(s)
}
