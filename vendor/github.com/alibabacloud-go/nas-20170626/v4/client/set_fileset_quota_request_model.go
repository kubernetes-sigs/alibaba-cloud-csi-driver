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
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may be different for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and dependencies without actually deleting the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a check request without deleting the export directory. The check items include whether required parameters are specified, the request format, and business limit dependencies. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is directly deleted.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file count limit of the quota. Valid values:
	//
	// - Minimum value: 10,000.
	//
	// - Maximum value: 10,000,000,000.
	//
	// > If this parameter is not specified, the file count is unlimited.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The ID of the CPFS for Lingjun file system. The ID must start with `bmcpfs-`, such as bmcpfs-290w65p03ok64ya****. You can call [DescribeFileSystems](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-describefilesystems) (FileSystemType=bmcpfs) to query existing file systems.
	//
	// This parameter is required.
	//
	// example:
	//
	// bmcpfs-290w65p03ok64ya****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// Fileset ID。
	//
	// This parameter is required.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The total capacity limit of the quota. Unit: bytes.
	//
	// Valid values:
	//
	// - Minimum value: 10,737,418,240 (10 GiB).
	//
	// - Increment: 1,073,741,824 (1 GiB).
	//
	// > If this parameter is not specified, the capacity is unlimited.
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
