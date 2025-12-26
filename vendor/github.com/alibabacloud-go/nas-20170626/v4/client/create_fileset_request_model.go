// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFilesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateFilesetRequest
	GetClientToken() *string
	SetDeletionProtection(v bool) *CreateFilesetRequest
	GetDeletionProtection() *bool
	SetDescription(v string) *CreateFilesetRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateFilesetRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CreateFilesetRequest
	GetFileSystemId() *string
	SetFileSystemPath(v string) *CreateFilesetRequest
	GetFileSystemPath() *string
	SetQuota(v *CreateFilesetRequestQuota) *CreateFilesetRequest
	GetQuota() *CreateFilesetRequestQuota
}

type CreateFilesetRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable deletion protection to allow you to release the fileset by using the console or by calling the [DeleteFileset](https://help.aliyun.com/document_detail/2838077.html) operation.
	//
	// 	- true: enables release protection.
	//
	// 	- false (default): disables release protection.
	//
	// >  This parameter can protect filesets only against manual releases, but not against automatic releases.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of the fileset.
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
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no fileset is created and no fee is incurred.
	//
	// Valid values:
	//
	// 	- true: performs a dry run. The system checks the required parameters, request syntax, service limits, and available Apsara File Storage NAS (NAS) resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the FsetId parameter.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, a fileset is created.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the file system.
	//
	// 	- The IDs of CPFS file systems must start with `cpfs-`. Example: cpfs-099394bd928c\\*\\*\\*\\*.
	//
	// 	- The IDs of CPFS for LINGJUN file systems must start with `bmcpfs-`. Example: bmcpfs-290w65p03ok64ya\\*\\*\\*\\*.
	//
	// >  CPFS is not supported on the international site.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The absolute path of the fileset.
	//
	// 	- The path must be 2 to 1024 characters in length.
	//
	// 	- The path must start and end with a forward slash (/).
	//
	// 	- The fileset path must be a new path and cannot be an existing path. Fileset paths cannot be renamed and cannot be symbolic links.
	//
	// 	- The maximum depth supported by a fileset path is eight levels. The depth of the root directory / is 0 levels. For example, the fileset path /test/aaa/ccc/ has three levels.
	//
	// 	- If the fileset path is a multi-level path, the parent directory must be an existing directory.
	//
	// 	- Nested filesets are not supported. If a fileset is specified as a parent directory, its subdirectory cannot be a fileset. A fileset path supports only one quota.
	//
	// This parameter is required.
	//
	// example:
	//
	// /test/
	FileSystemPath *string `json:"FileSystemPath,omitempty" xml:"FileSystemPath,omitempty"`
	// The quota information.
	//
	// >  Only CPFS for LINGJUN V2.7.0 and later support this parameter.
	Quota *CreateFilesetRequestQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
}

func (s CreateFilesetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFilesetRequest) GoString() string {
	return s.String()
}

func (s *CreateFilesetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateFilesetRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateFilesetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateFilesetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateFilesetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateFilesetRequest) GetFileSystemPath() *string {
	return s.FileSystemPath
}

func (s *CreateFilesetRequest) GetQuota() *CreateFilesetRequestQuota {
	return s.Quota
}

func (s *CreateFilesetRequest) SetClientToken(v string) *CreateFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateFilesetRequest) SetDeletionProtection(v bool) *CreateFilesetRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateFilesetRequest) SetDescription(v string) *CreateFilesetRequest {
	s.Description = &v
	return s
}

func (s *CreateFilesetRequest) SetDryRun(v bool) *CreateFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateFilesetRequest) SetFileSystemId(v string) *CreateFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateFilesetRequest) SetFileSystemPath(v string) *CreateFilesetRequest {
	s.FileSystemPath = &v
	return s
}

func (s *CreateFilesetRequest) SetQuota(v *CreateFilesetRequestQuota) *CreateFilesetRequest {
	s.Quota = v
	return s
}

func (s *CreateFilesetRequest) Validate() error {
	if s.Quota != nil {
		if err := s.Quota.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateFilesetRequestQuota struct {
	// The number of files of the quota. Valid values:
	//
	// 	- Minimum value: 100000.
	//
	// 	- Maximum value: 10000000000.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The total capacity of the quota. Unit: bytes.
	//
	// Valid values:
	//
	// 	- Minimum value: 10737418240 (10 GiB).
	//
	// 	- Maximum value: 1073741824000 (1024000 GiB).
	//
	// 	- Step size: 1073741824 (1 GiB).
	//
	// example:
	//
	// 10737418240
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s CreateFilesetRequestQuota) String() string {
	return dara.Prettify(s)
}

func (s CreateFilesetRequestQuota) GoString() string {
	return s.String()
}

func (s *CreateFilesetRequestQuota) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *CreateFilesetRequestQuota) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *CreateFilesetRequestQuota) SetFileCountLimit(v int64) *CreateFilesetRequestQuota {
	s.FileCountLimit = &v
	return s
}

func (s *CreateFilesetRequestQuota) SetSizeLimit(v int64) *CreateFilesetRequestQuota {
	s.SizeLimit = &v
	return s
}

func (s *CreateFilesetRequestQuota) Validate() error {
	return dara.Validate(s)
}
