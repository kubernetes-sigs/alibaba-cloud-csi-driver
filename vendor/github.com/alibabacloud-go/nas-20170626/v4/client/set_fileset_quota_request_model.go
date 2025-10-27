// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetFilesetQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SetFilesetQuotaRequest
	GetClientToken() *string
	SetDryRun(v bool) *SetFilesetQuotaRequest
	GetDryRun() *bool
	SetFileCountLimit(v int64) *SetFilesetQuotaRequest
	GetFileCountLimit() *int64
	SetFileSystemId(v string) *SetFilesetQuotaRequest
	GetFileSystemId() *string
	SetFsetId(v string) *SetFilesetQuotaRequest
	GetFsetId() *string
	SetSizeLimit(v int64) *SetFilesetQuotaRequest
	GetSizeLimit() *int64
}

type SetFilesetQuotaRequest struct {
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
	// Specifies whether to perform a dry run. The dry run checks parameter validity and prerequisites. The dry run does not delete the specified quota or incur fees.
	//
	// Valid values:
	//
	// 	- true: performs only a dry run. The system checks the required parameters, request syntax, and service limits. If the request fails the dry run, an error code is returned. If the request passes the dry run, the HTTP status code 200 is returned.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, the quota is deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The number of files of the quota. Valid values:
	//
	// 	- Minimum value: 10,000.
	//
	// 	- Maximum value: 10,000,000,000.
	//
	// >  If you do not specify this parameter, the number of files is unlimited.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The ID of the CPFS for LINGJUN file system. The IDs of CPFS for LINGJUN file systems must start with `bmcpfs-`. Example: bmcpfs-290w65p03ok64ya\\*\\*\\*\\*.
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
	// The total capacity of the quota. Unit: bytes.
	//
	// Valid values:
	//
	// 	- Minimum value: 10,737,418,240 (10 GiB).
	//
	// 	- Step size: 1,073,741,824 (1 GiB).
	//
	// >  If you do not specify this parameter, the capacity is unlimited.
	//
	// example:
	//
	// 10737418240
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s SetFilesetQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s SetFilesetQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetFilesetQuotaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetFilesetQuotaRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *SetFilesetQuotaRequest) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *SetFilesetQuotaRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *SetFilesetQuotaRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *SetFilesetQuotaRequest) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *SetFilesetQuotaRequest) SetClientToken(v string) *SetFilesetQuotaRequest {
	s.ClientToken = &v
	return s
}

func (s *SetFilesetQuotaRequest) SetDryRun(v bool) *SetFilesetQuotaRequest {
	s.DryRun = &v
	return s
}

func (s *SetFilesetQuotaRequest) SetFileCountLimit(v int64) *SetFilesetQuotaRequest {
	s.FileCountLimit = &v
	return s
}

func (s *SetFilesetQuotaRequest) SetFileSystemId(v string) *SetFilesetQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *SetFilesetQuotaRequest) SetFsetId(v string) *SetFilesetQuotaRequest {
	s.FsetId = &v
	return s
}

func (s *SetFilesetQuotaRequest) SetSizeLimit(v int64) *SetFilesetQuotaRequest {
	s.SizeLimit = &v
	return s
}

func (s *SetFilesetQuotaRequest) Validate() error {
	return dara.Validate(s)
}
