// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyDataFlowRequest
	GetClientToken() *string
	SetDataFlowId(v string) *ModifyDataFlowRequest
	GetDataFlowId() *string
	SetDescription(v string) *ModifyDataFlowRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyDataFlowRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *ModifyDataFlowRequest
	GetFileSystemId() *string
	SetThroughput(v int64) *ModifyDataFlowRequest
	GetThroughput() *int64
}

type ModifyDataFlowRequest struct {
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
	// The description of the dataflow.
	//
	// Limits:
	//
	// 	- The description must be 2 to 128 characters in length.
	//
	// 	- The description must start with a letter but cannot start with http:// or https://.
	//
	// 	- The description can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no file system is created and no fee is incurred.
	//
	// Valid values:
	//
	// 	- true: performs a dry run. The system checks the required parameters, request syntax, limits, and available NAS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the FileSystemId parameter.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, a file system is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-125487\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for Lingjun file systems must start with `bmcpfs-`. Example: bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The maximum data flow throughput. Unit: MB/s. Valid values:
	//
	// 	- 600
	//
	// 	- 1200
	//
	// 	- 1500
	//
	// >  The data flow throughput must be less than the I/O throughput of the file system. This parameter is required for CPFS file systems.
	//
	// example:
	//
	// 600
	Throughput *int64 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
}

func (s ModifyDataFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataFlowRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDataFlowRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *ModifyDataFlowRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyDataFlowRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDataFlowRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyDataFlowRequest) GetThroughput() *int64 {
	return s.Throughput
}

func (s *ModifyDataFlowRequest) SetClientToken(v string) *ModifyDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDataFlowRequest) SetDataFlowId(v string) *ModifyDataFlowRequest {
	s.DataFlowId = &v
	return s
}

func (s *ModifyDataFlowRequest) SetDescription(v string) *ModifyDataFlowRequest {
	s.Description = &v
	return s
}

func (s *ModifyDataFlowRequest) SetDryRun(v bool) *ModifyDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDataFlowRequest) SetFileSystemId(v string) *ModifyDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyDataFlowRequest) SetThroughput(v int64) *ModifyDataFlowRequest {
	s.Throughput = &v
	return s
}

func (s *ModifyDataFlowRequest) Validate() error {
	return dara.Validate(s)
}
