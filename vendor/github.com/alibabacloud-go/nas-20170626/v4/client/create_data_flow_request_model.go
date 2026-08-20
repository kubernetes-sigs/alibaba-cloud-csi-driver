// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataFlowRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoRefreshInterval(v int64) *CreateDataFlowRequest
	GetAutoRefreshInterval() *int64
	SetAutoRefreshPolicy(v string) *CreateDataFlowRequest
	GetAutoRefreshPolicy() *string
	SetAutoRefreshs(v []*CreateDataFlowRequestAutoRefreshs) *CreateDataFlowRequest
	GetAutoRefreshs() []*CreateDataFlowRequestAutoRefreshs
	SetClientToken(v string) *CreateDataFlowRequest
	GetClientToken() *string
	SetDescription(v string) *CreateDataFlowRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateDataFlowRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CreateDataFlowRequest
	GetFileSystemId() *string
	SetFileSystemPath(v string) *CreateDataFlowRequest
	GetFileSystemPath() *string
	SetFsetId(v string) *CreateDataFlowRequest
	GetFsetId() *string
	SetSourceSecurityType(v string) *CreateDataFlowRequest
	GetSourceSecurityType() *string
	SetSourceStorage(v string) *CreateDataFlowRequest
	GetSourceStorage() *string
	SetSourceStoragePath(v string) *CreateDataFlowRequest
	GetSourceStoragePath() *string
	SetThroughput(v int64) *CreateDataFlowRequest
	GetThroughput() *int64
}

type CreateDataFlowRequest struct {
	// The auto-refresh interval. CPFS General-purpose checks the directory for data updates at this interval. If data updates exist, an auto-refresh task is started. Unit: minutes.
	//
	// Valid values: 10 to 525600. Default value: 10.
	//
	// > This parameter takes effect only when the file system type is CPFS General-purpose.
	//
	// example:
	//
	// 10
	AutoRefreshInterval *int64 `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	// The auto-refresh policy. Specifies the policy for importing data updates from the source storage to CPFS General-purpose after the source data is updated. Valid values:
	//
	// - None (default): Data updates in the source storage are not automatically imported to CPFS General-purpose. You can import data updates from the source storage by using a data flow task.
	//
	// - ImportChanged: Data updates in the source storage are automatically imported to CPFS General-purpose.
	//
	// > This parameter takes effect only when the file system type is CPFS General-purpose.
	//
	// example:
	//
	// None
	AutoRefreshPolicy *string `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	// The auto-refresh configuration collection.
	//
	// > This parameter takes effect only when the file system type is CPFS General-purpose.
	//
	// if can be null:
	// false
	AutoRefreshs []*CreateDataFlowRequestAutoRefreshs `json:"AutoRefreshs,omitempty" xml:"AutoRefreshs,omitempty" type:"Repeated"`
	// Ensures the idempotence of the request. Generate a parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the data flow.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter.
	//
	// - The description cannot start with `http://` or `https://`.
	//
	// - The description can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Bucket01 DataFlow
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this create request.
	//
	// A dry run checks parameter validity and resource availability without actually creating the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without creating the instance. The check items include required parameters, request format, business limits, and NAS inventory. If the check fails, the corresponding error is returned. If the check succeeds, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request and creates the instance after the check is passed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// - CPFS General-purpose: must start with `cpfs-`, such as cpfs-125487\\*\\*\\*\\*.
	//
	// - CPFS for Lingjun: must start with `bmcpfs-`, such as bmcpfs-0015\\*\\*\\*\\*.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The directory in the CPFS for Lingjun file system. Limits:
	//
	// - The path must start and end with a forward slash (/).
	//
	// - The directory must be an existing directory in the CPFS for Lingjun file system.
	//
	// - The path must be 1 to 1023 characters in length.
	//
	// - UTF-8 encoding is used.
	//
	// > This parameter is required when the file system type is CPFS for Lingjun.
	//
	// example:
	//
	// /path/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// The Fileset ID.
	//
	// > This parameter is required when the file system type is CPFS General-purpose.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The security protection type of the source storage. If the source storage must be accessed through security protection, specify the security protection type. Valid values:
	//
	// - None (default): The source storage does not require security protection for access.
	//
	// - SSL: Access is protected by an SSL certificate.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// SSL
	SourceSecurityType *string `json:"SourceSecurityType,omitempty" xml:"SourceSecurityType,omitempty"`
	// The access address of the source storage. Format: `<storage type>://[<account id>:]<path>`.
	//
	// Where:
	//
	// - storage type: only oss is supported.
	//
	// - account id: optional. The UID of the account that owns the source storage. This parameter is required when you use cross-account OSS.
	//
	// - path: the name of the OSS bucket. Limits:
	//
	//     - Only lowercase letters, digits, and hyphens (-) are supported. The name must start and end with a lowercase letter or digit.
	//
	//     - The maximum length is 128 characters.
	//
	//     - UTF-8 encoding is used.
	//
	// > - The OSS bucket must be an existing bucket in the region.
	//
	// > - The account id parameter is supported only by CPFS for Lingjun 2.6.0 and later.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://178321033379****:bucket-01
	SourceStorage *string `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	// The access path within the source storage bucket. Limits:
	//
	//    - The path must start and end with a forward slash (/).
	//
	// - The path is case-sensitive.
	//
	// - The path must be 1 to 1023 characters in length.
	//
	// - UTF-8 encoding is used.
	//
	// > This parameter is required when the file system type is CPFS for Lingjun.
	//
	// example:
	//
	// /prefix/
	SourceStoragePath *string `json:"SourceStoragePath,omitempty" xml:"SourceStoragePath,omitempty"`
	// The maximum transfer bandwidth of the data flow. Unit: MB/s. Valid values:
	//
	// - 600
	//
	// - 1200
	//
	// - 1500
	//
	// > The transfer bandwidth of the data flow must be less than the I/O bandwidth of the file system.
	//
	// > This parameter is required when the file system type is CPFS General-purpose.
	//
	// example:
	//
	// 600
	Throughput *int64 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
}

func (s CreateDataFlowRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowRequest) GoString() string {
	return s.String()
}

func (s *CreateDataFlowRequest) GetAutoRefreshInterval() *int64 {
	return s.AutoRefreshInterval
}

func (s *CreateDataFlowRequest) GetAutoRefreshPolicy() *string {
	return s.AutoRefreshPolicy
}

func (s *CreateDataFlowRequest) GetAutoRefreshs() []*CreateDataFlowRequestAutoRefreshs {
	return s.AutoRefreshs
}

func (s *CreateDataFlowRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateDataFlowRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateDataFlowRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateDataFlowRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateDataFlowRequest) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *CreateDataFlowRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *CreateDataFlowRequest) GetSourceSecurityType() *string {
	return s.SourceSecurityType
}

func (s *CreateDataFlowRequest) GetSourceStorage() *string {
	return s.SourceStorage
}

func (s *CreateDataFlowRequest) GetSourceStoragePath() *string {
	return s.SourceStoragePath
}

func (s *CreateDataFlowRequest) GetThroughput() *int64 {
	return s.Throughput
}

func (s *CreateDataFlowRequest) SetAutoRefreshInterval(v int64) *CreateDataFlowRequest {
	s.AutoRefreshInterval = &v
	return s
}

func (s *CreateDataFlowRequest) SetAutoRefreshPolicy(v string) *CreateDataFlowRequest {
	s.AutoRefreshPolicy = &v
	return s
}

func (s *CreateDataFlowRequest) SetAutoRefreshs(v []*CreateDataFlowRequestAutoRefreshs) *CreateDataFlowRequest {
	s.AutoRefreshs = v
	return s
}

func (s *CreateDataFlowRequest) SetClientToken(v string) *CreateDataFlowRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateDataFlowRequest) SetDescription(v string) *CreateDataFlowRequest {
	s.Description = &v
	return s
}

func (s *CreateDataFlowRequest) SetDryRun(v bool) *CreateDataFlowRequest {
	s.DryRun = &v
	return s
}

func (s *CreateDataFlowRequest) SetFileSystemId(v string) *CreateDataFlowRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateDataFlowRequest) SetFileSystemPath(v string) *CreateDataFlowRequest {
	s.FileSystemPath = &v
	return s
}

func (s *CreateDataFlowRequest) SetFsetId(v string) *CreateDataFlowRequest {
	s.FsetId = &v
	return s
}

func (s *CreateDataFlowRequest) SetSourceSecurityType(v string) *CreateDataFlowRequest {
	s.SourceSecurityType = &v
	return s
}

func (s *CreateDataFlowRequest) SetSourceStorage(v string) *CreateDataFlowRequest {
	s.SourceStorage = &v
	return s
}

func (s *CreateDataFlowRequest) SetSourceStoragePath(v string) *CreateDataFlowRequest {
	s.SourceStoragePath = &v
	return s
}

func (s *CreateDataFlowRequest) SetThroughput(v int64) *CreateDataFlowRequest {
	s.Throughput = &v
	return s
}

func (s *CreateDataFlowRequest) Validate() error {
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

type CreateDataFlowRequestAutoRefreshs struct {
	// The auto-refresh directory. CPFS General-purpose registers data modification events from the source storage and checks whether the source data in this directory has been updated, then automatically imports the updated data.
	//
	// The default value is empty, which means that data updates in the source storage are not automatically imported to CPFS General-purpose. You must manually create a task to import updates.
	//
	// Limits:
	//
	// - The path must be 2 to 1024 characters in length.
	//
	// - UTF-8 encoding is used.
	//
	// - The path must start and end with a forward slash (/).
	//
	// - The directory must be an existing directory in the CPFS General-purpose file system and must be located within the Fileset directory of the data flow.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// /prefix1/prefix2/
	RefreshPath *string `json:"RefreshPath,omitempty" xml:"RefreshPath,omitempty"`
}

func (s CreateDataFlowRequestAutoRefreshs) String() string {
	return dara.Prettify(s)
}

func (s CreateDataFlowRequestAutoRefreshs) GoString() string {
	return s.String()
}

func (s *CreateDataFlowRequestAutoRefreshs) GetRefreshPath() *string {
	return s.RefreshPath
}

func (s *CreateDataFlowRequestAutoRefreshs) SetRefreshPath(v string) *CreateDataFlowRequestAutoRefreshs {
	s.RefreshPath = &v
	return s
}

func (s *CreateDataFlowRequestAutoRefreshs) Validate() error {
	return dara.Validate(s)
}
