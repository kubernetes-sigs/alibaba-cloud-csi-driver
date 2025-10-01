// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFilesetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyFilesetRequest
	GetClientToken() *string
	SetDeletionProtection(v bool) *ModifyFilesetRequest
	GetDeletionProtection() *bool
	SetDescription(v string) *ModifyFilesetRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyFilesetRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *ModifyFilesetRequest
	GetFileSystemId() *string
	SetFsetId(v string) *ModifyFilesetRequest
	GetFsetId() *string
}

type ModifyFilesetRequest struct {
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
	// 	- false: disables release protection.
	//
	// >  This parameter can protect filesets only against manual releases, but not against automatic releases.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The fileset description.
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no fileset is modified and no fees incurred.
	//
	// Valid values:
	//
	// 	- true: performs only a dry run. The system checks the required parameters, request syntax, service limits, and Apsara File Storage NAS (NAS) inventory data. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, the specified fileset is modified.
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
	// The fileset ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
}

func (s ModifyFilesetRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFilesetRequest) GoString() string {
	return s.String()
}

func (s *ModifyFilesetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyFilesetRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyFilesetRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyFilesetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyFilesetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyFilesetRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *ModifyFilesetRequest) SetClientToken(v string) *ModifyFilesetRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyFilesetRequest) SetDeletionProtection(v bool) *ModifyFilesetRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyFilesetRequest) SetDescription(v string) *ModifyFilesetRequest {
	s.Description = &v
	return s
}

func (s *ModifyFilesetRequest) SetDryRun(v bool) *ModifyFilesetRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyFilesetRequest) SetFileSystemId(v string) *ModifyFilesetRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyFilesetRequest) SetFsetId(v string) *ModifyFilesetRequest {
	s.FsetId = &v
	return s
}

func (s *ModifyFilesetRequest) Validate() error {
	return dara.Validate(s)
}
