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
	// The automatic update interval. CPFS checks whether data is updated in the directory at the interval specified by this parameter. If data is updated, CPFS starts an automatic update task. Unit: minutes.
	//
	// Valid values: 10 to 525600. Default value: 10.
	//
	// >  This parameter takes effect only for CPFS file systems.
	//
	// example:
	//
	// 10
	AutoRefreshInterval *int64 `json:"AutoRefreshInterval,omitempty" xml:"AutoRefreshInterval,omitempty"`
	// The automatic update policy. The updated data in the source storage is imported into the CPFS file system based on the policy.
	//
	// 	- None (default): Updated data in the source storage is not automatically imported into the CPFS file system. You can run a data flow task to import the updated data from the source storage.
	//
	// 	- ImportChanged: Updated data in the source storage is automatically imported into the CPFS file system.
	//
	// >  This parameter takes effect only for CPFS file systems.
	//
	// example:
	//
	// None
	AutoRefreshPolicy *string `json:"AutoRefreshPolicy,omitempty" xml:"AutoRefreshPolicy,omitempty"`
	// The automatic update configurations.
	//
	// >  This parameter takes effect only for CPFS file systems.
	//
	// if can be null:
	// false
	AutoRefreshs []*CreateDataFlowRequestAutoRefreshs `json:"AutoRefreshs,omitempty" xml:"AutoRefreshs,omitempty" type:"Repeated"`
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
	// The description of the dataflow.
	//
	// Limits:
	//
	// 	- The description must be 2 to 128 characters in length.
	//
	// 	- The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// 	- The description can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Bucket01 DataFlow
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
	// The directory in the CPFS for LINGJUN file system. Limits:
	//
	// 	- The directory must start and end with a forward slash (/).
	//
	// 	- The directory must be an existing directory in the CPFS for LINGJUN file system.
	//
	// 	- The directory must be 1 to 1023 characters in length.
	//
	// 	- The directory must be encoded in UTF-8.
	//
	// >  This parameter is required for CPFS for LINGJUN file systems.
	//
	// example:
	//
	// /path/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// The fileset ID.
	//
	// >  This parameter is required for CPFS file systems.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The type of security mechanism for the source storage. This parameter must be specified if the source storage is accessed with a security mechanism. Valid values:
	//
	// 	- None (default): The source storage can be accessed without a security mechanism.
	//
	// 	- SSL: The source storage must be accessed with an SSL certificate.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// SSL
	SourceSecurityType *string `json:"SourceSecurityType,omitempty" xml:"SourceSecurityType,omitempty"`
	// The access path of the source storage. Format: `<storage type>://[<account id>:]<path>`.
	//
	// Parameters:
	//
	// 	- storage type: Only OSS is supported.
	//
	// 	- account id (optional): the UID of the account of the source storage. This parameter is required when you use OSS buckets across accounts.
	//
	// 	- path: the name of the OSS bucket. Limits:
	//
	//     	- The name can contain only lowercase letters, digits, and hyphens (-). The name must start and end with a lowercase letter or digit.
	//
	//     	- The name can be up to 128 characters in length.
	//
	//     	- The name must be encoded in UTF-8.
	//
	// > 	- The OSS bucket must be an existing bucket in the region.
	//
	// > 	- Only CPFS for LINGJUN V2.6.0 and later support the account id parameter.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://bucket1
	SourceStorage *string `json:"SourceStorage,omitempty" xml:"SourceStorage,omitempty"`
	// The access path in the bucket of the source storage. Limits:
	//
	// 	- The path must start and end with a forward slash (/).
	//
	// 	- The path is case-sensitive.
	//
	// 	- The path must be 1 to 1023 characters in length.
	//
	// 	- The path must be encoded in UTF-8.
	//
	// >  This parameter is required for CPFS for LINGJUN file systems.
	//
	// example:
	//
	// /prefix/
	SourceStoragePath *string `json:"SourceStoragePath,omitempty" xml:"SourceStoragePath,omitempty"`
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
	// The automatic update directory. CPFS registers the data update event in the source storage, and automatically checks whether the source data in the directory is updated and imports the updated data.
	//
	// This parameter is empty by default. Updated data in the source storage is not automatically imported into the CPFS file system. You must import the updated data by running a manual task.
	//
	// Limits:
	//
	// 	- The directory must be 2 to 1,024 characters in length.
	//
	// 	- The directory must be encoded in UTF-8.
	//
	// 	- The directory must start and end with a forward slash (/).
	//
	// 	- The directory must be an existing directory in the CPFS file system and must be in a fileset where the data flow is enabled.
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
