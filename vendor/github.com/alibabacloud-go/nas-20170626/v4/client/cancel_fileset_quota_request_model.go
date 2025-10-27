// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelFilesetQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CancelFilesetQuotaRequest
	GetClientToken() *string
	SetDryRun(v bool) *CancelFilesetQuotaRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CancelFilesetQuotaRequest
	GetFileSystemId() *string
	SetFsetId(v string) *CancelFilesetQuotaRequest
	GetFsetId() *string
}

type CancelFilesetQuotaRequest struct {
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
	// Specifies whether to perform a dry run.
	//
	// During the dry run, the system checks whether the request parameters are valid and whether the requested resources are available. During the dry run, no fileset quota is canceled and no fee is incurred.
	//
	// Valid values:
	//
	// 	- true: performs a dry run. The system checks the required parameters, request syntax, and service limits. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the DataFlowld parameter.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, the fileset quota is canceled.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
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
}

func (s CancelFilesetQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelFilesetQuotaRequest) GoString() string {
	return s.String()
}

func (s *CancelFilesetQuotaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CancelFilesetQuotaRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CancelFilesetQuotaRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CancelFilesetQuotaRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *CancelFilesetQuotaRequest) SetClientToken(v string) *CancelFilesetQuotaRequest {
	s.ClientToken = &v
	return s
}

func (s *CancelFilesetQuotaRequest) SetDryRun(v bool) *CancelFilesetQuotaRequest {
	s.DryRun = &v
	return s
}

func (s *CancelFilesetQuotaRequest) SetFileSystemId(v string) *CancelFilesetQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *CancelFilesetQuotaRequest) SetFsetId(v string) *CancelFilesetQuotaRequest {
	s.FsetId = &v
	return s
}

func (s *CancelFilesetQuotaRequest) Validate() error {
	return dara.Validate(s)
}
