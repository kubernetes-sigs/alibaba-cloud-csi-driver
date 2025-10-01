// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDataFlowAutoRefreshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRefreshInterval(v int64) *ModifyDataFlowAutoRefreshRequest
	GetAutoRefreshInterval() *int64
	SetAutoRefreshPolicy(v string) *ModifyDataFlowAutoRefreshRequest
	GetAutoRefreshPolicy() *string
	SetClientToken(v string) *ModifyDataFlowAutoRefreshRequest
	GetClientToken() *string
	SetDataFlowId(v string) *ModifyDataFlowAutoRefreshRequest
	GetDataFlowId() *string
	SetDryRun(v bool) *ModifyDataFlowAutoRefreshRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *ModifyDataFlowAutoRefreshRequest
	GetFileSystemId() *string
}

type ModifyDataFlowAutoRefreshRequest struct {
	// The automatic update interval. CPFS checks whether data is updated in the directory at the interval. If data is updated, CPFS runs an AutoRefresh task. Unit: minutes.
	//
	// Valid values: 5 to 526600. Default value: 10.
	//
	// example:
	//
	// 10
	AutoRefreshInterval *int64 `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	// The automatic update policy. CPFS imports data updates in the Object Storage Service (OSS) bucket to the CPFS file system based on this policy. Valid values:
	//
	// 	- None: CPFS does not automatically import data updates in the OSS bucket to the CPFS file system. You can import the data updates by using a dataflow task.
	//
	// 	- ImportChanged: CPFS automatically imports data updates in the OSS bucket to the CPFS file system.
	//
	// example:
	//
	// None
	AutoRefreshPolicy *string `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
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
	// The dataflow ID.
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
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ModifyDataFlowAutoRefreshRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDataFlowAutoRefreshRequest) GoString() string {
	return s.String()
}

func (s *ModifyDataFlowAutoRefreshRequest) GetAutoRefreshInterval() *int64 {
	return s.AutoRefreshInterval
}

func (s *ModifyDataFlowAutoRefreshRequest) GetAutoRefreshPolicy() *string {
	return s.AutoRefreshPolicy
}

func (s *ModifyDataFlowAutoRefreshRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDataFlowAutoRefreshRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *ModifyDataFlowAutoRefreshRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyDataFlowAutoRefreshRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyDataFlowAutoRefreshRequest) SetAutoRefreshInterval(v int64) *ModifyDataFlowAutoRefreshRequest {
	s.AutoRefreshInterval = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetAutoRefreshPolicy(v string) *ModifyDataFlowAutoRefreshRequest {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetClientToken(v string) *ModifyDataFlowAutoRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetDataFlowId(v string) *ModifyDataFlowAutoRefreshRequest {
	s.DataFlowId = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetDryRun(v bool) *ModifyDataFlowAutoRefreshRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) SetFileSystemId(v string) *ModifyDataFlowAutoRefreshRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyDataFlowAutoRefreshRequest) Validate() error {
	return dara.Validate(s)
}
