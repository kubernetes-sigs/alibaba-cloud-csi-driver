// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iApplyDataFlowAutoRefreshRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRefreshInterval(v int64) *ApplyDataFlowAutoRefreshRequest
	GetAutoRefreshInterval() *int64
	SetAutoRefreshPolicy(v string) *ApplyDataFlowAutoRefreshRequest
	GetAutoRefreshPolicy() *string
	SetAutoRefreshs(v []*ApplyDataFlowAutoRefreshRequestAutoRefreshs) *ApplyDataFlowAutoRefreshRequest
	GetAutoRefreshs() []*ApplyDataFlowAutoRefreshRequestAutoRefreshs
	SetClientToken(v string) *ApplyDataFlowAutoRefreshRequest
	GetClientToken() *string
	SetDataFlowId(v string) *ApplyDataFlowAutoRefreshRequest
	GetDataFlowId() *string
	SetDryRun(v bool) *ApplyDataFlowAutoRefreshRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *ApplyDataFlowAutoRefreshRequest
	GetFileSystemId() *string
}

type ApplyDataFlowAutoRefreshRequest struct {
	// The auto-refresh interval. CPFS checks the directory for data updates at this interval. If data updates exist, an auto-refresh task is started. Unit: minutes.
	//
	// Valid values: 10 to 525600. Default value: 10.
	//
	// example:
	//
	// 10
	AutoRefreshInterval *int64 `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	// The auto-refresh policy. This policy determines how data updates from the source are imported to CPFS. Valid values:
	//
	// - None: Data updates from the source are not automatically imported to CPFS. You can use a data flow task to import data updates from the source.
	//
	// - ImportChanged: Data updates from the source are automatically imported to CPFS.
	//
	// example:
	//
	// None
	AutoRefreshPolicy *string `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	// The collection of auto-refresh configurations.
	//
	// This parameter is required.
	AutoRefreshs []*ApplyDataFlowAutoRefreshRequestAutoRefreshs `json:"AutoRefreshs,omitempty" xml:"AutoRefreshs,omitempty" type:"Repeated"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the data flow.
	//
	// This parameter is required.
	//
	// example:
	//
	// df-194433a5be31****
	DataFlowId *string `json:"DataFlowId,omitempty" xml:"DataFlowId,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating an instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without creating an instance. The check items include required parameters, request format, business limits, and NAS inventory. If the check fails, the corresponding error is returned. If the check succeeds, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request. After the check succeeds, the instance is created.
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

func (s ApplyDataFlowAutoRefreshRequest) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshRequest) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshRequest) GetAutoRefreshInterval() *int64 {
	return s.AutoRefreshInterval
}

func (s *ApplyDataFlowAutoRefreshRequest) GetAutoRefreshPolicy() *string {
	return s.AutoRefreshPolicy
}

func (s *ApplyDataFlowAutoRefreshRequest) GetAutoRefreshs() []*ApplyDataFlowAutoRefreshRequestAutoRefreshs {
	return s.AutoRefreshs
}

func (s *ApplyDataFlowAutoRefreshRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ApplyDataFlowAutoRefreshRequest) GetDataFlowId() *string {
	return s.DataFlowId
}

func (s *ApplyDataFlowAutoRefreshRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ApplyDataFlowAutoRefreshRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ApplyDataFlowAutoRefreshRequest) SetAutoRefreshInterval(v int64) *ApplyDataFlowAutoRefreshRequest {
	s.AutoRefreshInterval = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetAutoRefreshPolicy(v string) *ApplyDataFlowAutoRefreshRequest {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetAutoRefreshs(v []*ApplyDataFlowAutoRefreshRequestAutoRefreshs) *ApplyDataFlowAutoRefreshRequest {
	s.AutoRefreshs = v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetClientToken(v string) *ApplyDataFlowAutoRefreshRequest {
	s.ClientToken = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetDataFlowId(v string) *ApplyDataFlowAutoRefreshRequest {
	s.DataFlowId = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetDryRun(v bool) *ApplyDataFlowAutoRefreshRequest {
	s.DryRun = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) SetFileSystemId(v string) *ApplyDataFlowAutoRefreshRequest {
	s.FileSystemId = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequest) Validate() error {
	if s.AutoRefreshs != nil {
		for _, item := range s.AutoRefreshs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ApplyDataFlowAutoRefreshRequestAutoRefreshs struct {
	// The auto-refresh directory. CPFS automatically checks whether data in this directory on the source has been updated and imports the updated data.
	//
	// Limits:
	//
	// - The path must be 2 to 1,024 characters in length.
	//
	// - The path must be encoded in UTF-8.
	//
	// - The path must start and end with a forward slash (/).
	//
	// > The directory must already exist in CPFS and must be in a fileset that has data flow enabled.
	//
	// This parameter is required.
	//
	// example:
	//
	// /prefix1/prefix2/
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s ApplyDataFlowAutoRefreshRequestAutoRefreshs) String() string {
	return dara.Prettify(s)
}

func (s ApplyDataFlowAutoRefreshRequestAutoRefreshs) GoString() string {
	return s.String()
}

func (s *ApplyDataFlowAutoRefreshRequestAutoRefreshs) GetRefreshPath() *string {
	return s.RefreshPath
}

func (s *ApplyDataFlowAutoRefreshRequestAutoRefreshs) SetRefreshPath(v string) *ApplyDataFlowAutoRefreshRequestAutoRefreshs {
	s.RefreshPath = &v
	return s
}

func (s *ApplyDataFlowAutoRefreshRequestAutoRefreshs) Validate() error {
	return dara.Validate(s)
}
