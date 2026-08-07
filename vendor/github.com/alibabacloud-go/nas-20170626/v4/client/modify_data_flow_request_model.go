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
	// Ensures the idempotency of the request. Generate a parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The data flow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// The description of the data flow.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter and cannot start with http:// or https://.
	//
	// - The description can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Bucket01数据流动
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a check request without creating the instance. The check items include whether required parameters are specified, the request format, business limitations, and NAS inventory. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is directly created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - CPFS: The ID must start with `cpfs-`, such as cpfs-125487\\*\\*\\*\\*.
	//
	// - CPFS for Lingjun: The ID must start with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The maximum transmission bandwidth of the data flow. Unit: MB/s.
	//
	// Valid values:
	//
	// - 600
	//
	// - 1200
	//
	// - 1500
	//
	// > The transmission bandwidth of the data flow must be less than the I/O bandwidth of the file system. This parameter is required when the file system type is CPFS.
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
