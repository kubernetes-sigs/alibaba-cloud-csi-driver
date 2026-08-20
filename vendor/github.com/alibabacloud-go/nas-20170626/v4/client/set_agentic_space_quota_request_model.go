// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAgenticSpaceQuotaRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticSpaceId(v string) *SetAgenticSpaceQuotaRequest
	GetAgenticSpaceId() *string
	SetClientToken(v string) *SetAgenticSpaceQuotaRequest
	GetClientToken() *string
	SetDryRun(v bool) *SetAgenticSpaceQuotaRequest
	GetDryRun() *bool
	SetFileCountLimit(v int64) *SetAgenticSpaceQuotaRequest
	GetFileCountLimit() *int64
	SetFileSystemId(v string) *SetAgenticSpaceQuotaRequest
	GetFileSystemId() *string
	SetSizeLimit(v int64) *SetAgenticSpaceQuotaRequest
	GetSizeLimit() *int64
}

type SetAgenticSpaceQuotaRequest struct {
	// AgenticSpace Id。
	//
	// This parameter is required.
	//
	// example:
	//
	// agentic-229oypxjgpau2****
	AgenticSpaceId *string `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
	// Ensures the idempotency of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotency](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to perform a dry run for this request. A dry run checks parameter validity and dependencies without actually modifying the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without modifying the protocol service. The check items include required parameters, request format, and business dependency conditions. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned.
	//
	// - false (default): Sends a normal request. After the check passes, the protocol service is directly modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file count limit for the quota. Valid values:
	//
	// - Minimum value: 10,000.
	//
	// - Maximum value: 100,000,000.
	//
	// example:
	//
	// 10000
	FileCountLimit *int64 `json:"FileCountLimit,omitempty" xml:"FileCountLimit,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The total capacity limit for the quota. Unit: bytes.
	//
	// Valid values:
	//
	// - Minimum value: 10,737,418,240 (10 GiB).
	//
	// - Maximum value: 1,099,511,627,776,000 (1,024,000 GiB).
	//
	// - Step: 1,073,741,824 (1 GiB).
	//
	// example:
	//
	// 10737418240
	SizeLimit *int64 `json:"SizeLimit,omitempty" xml:"SizeLimit,omitempty"`
}

func (s SetAgenticSpaceQuotaRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAgenticSpaceQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetAgenticSpaceQuotaRequest) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *SetAgenticSpaceQuotaRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetAgenticSpaceQuotaRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *SetAgenticSpaceQuotaRequest) GetFileCountLimit() *int64 {
	return s.FileCountLimit
}

func (s *SetAgenticSpaceQuotaRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *SetAgenticSpaceQuotaRequest) GetSizeLimit() *int64 {
	return s.SizeLimit
}

func (s *SetAgenticSpaceQuotaRequest) SetAgenticSpaceId(v string) *SetAgenticSpaceQuotaRequest {
	s.AgenticSpaceId = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) SetClientToken(v string) *SetAgenticSpaceQuotaRequest {
	s.ClientToken = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) SetDryRun(v bool) *SetAgenticSpaceQuotaRequest {
	s.DryRun = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) SetFileCountLimit(v int64) *SetAgenticSpaceQuotaRequest {
	s.FileCountLimit = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) SetFileSystemId(v string) *SetAgenticSpaceQuotaRequest {
	s.FileSystemId = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) SetSizeLimit(v int64) *SetAgenticSpaceQuotaRequest {
	s.SizeLimit = &v
	return s
}

func (s *SetAgenticSpaceQuotaRequest) Validate() error {
	return dara.Validate(s)
}
