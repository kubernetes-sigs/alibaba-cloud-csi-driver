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
	// The automatic update interval. CPFS checks whether data is updated in the directory at the interval specified by this parameter. If data is updated, CPFS starts an automatic update task. Unit: minutes.
	//
	// Valid values: 5 to 526600. Default value: 10.
	//
	// example:
	//
	// 10
	AutoRefreshInterval *int64 `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	// The automatic update policy. The updated data in the source storage is imported into the CPFS file system based on the policy. Valid values:
	//
	// 	- None (default): Updated data in the source storage is not automatically imported into the CPFS file system. You can run a dataflow task to import the updated data from the source storage.
	//
	// 	- ImportChanged: Updated data in the source storage is automatically imported into the CPFS file system.
	//
	// example:
	//
	// None
	AutoRefreshPolicy *string `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	// The automatic update configurations.
	//
	// This parameter is required.
	AutoRefreshs []*ApplyDataFlowAutoRefreshRequestAutoRefreshs `json:"AutoRefreshs,omitempty" xml:"AutoRefreshs,omitempty" type:"Repeated"`
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
	// The automatic update directory. CPFS automatically checks whether the source data only in the directory is updated and imports the updated data.
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
