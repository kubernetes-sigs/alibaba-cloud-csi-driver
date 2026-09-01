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
	// The automatic update interval. CPFS checks the directory for data updates at each interval. If data updates exist, an automatic update task is started. Unit: minutes.
	//
	// Valid values: 5 to 526600. Default value: 10.
	//
	// example:
	//
	// 10
	AutoRefreshInterval *int64 `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	// The automatic update policy. This policy specifies how data updates from the source are imported to CPFS after the source data is updated. Valid values:
	//
	// - None: Data updates from the source are not automatically imported to CPFS. You can import source data updates by using a dataflow task.
	//
	// - ImportChanged: Data updates from the source are automatically imported to CPFS.
	//
	// example:
	//
	// None
	AutoRefreshPolicy *string `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
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
	// The dataflow ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating an instance or incurring fees.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without creating an instance. The check items include whether required parameters are specified, the request format, business limits, and NAS inventory. If the check fails, the corresponding error is returned. If the check succeeds, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request. After the check succeeds, the instance is directly created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
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
