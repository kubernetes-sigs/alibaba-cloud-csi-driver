// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAgenticSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticSpaceId(v string) *DeleteAgenticSpaceRequest
	GetAgenticSpaceId() *string
	SetClientToken(v string) *DeleteAgenticSpaceRequest
	GetClientToken() *string
	SetDryRun(v bool) *DeleteAgenticSpaceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *DeleteAgenticSpaceRequest
	GetFileSystemId() *string
}

type DeleteAgenticSpaceRequest struct {
	// AgenticSpace Id。
	//
	// This parameter is required.
	//
	// example:
	//
	// agentic-229oypxjgpau2****
	AgenticSpaceId *string `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
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
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually creating or deleting instances, and incurs no fees.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without deleting the instance. The check items include required parameters, request format, service limits, and NAS resource availability. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request and deletes the instance after the check is passed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The file system ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1ca404****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s DeleteAgenticSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAgenticSpaceRequest) GoString() string {
	return s.String()
}

func (s *DeleteAgenticSpaceRequest) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *DeleteAgenticSpaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteAgenticSpaceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *DeleteAgenticSpaceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *DeleteAgenticSpaceRequest) SetAgenticSpaceId(v string) *DeleteAgenticSpaceRequest {
	s.AgenticSpaceId = &v
	return s
}

func (s *DeleteAgenticSpaceRequest) SetClientToken(v string) *DeleteAgenticSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteAgenticSpaceRequest) SetDryRun(v bool) *DeleteAgenticSpaceRequest {
	s.DryRun = &v
	return s
}

func (s *DeleteAgenticSpaceRequest) SetFileSystemId(v string) *DeleteAgenticSpaceRequest {
	s.FileSystemId = &v
	return s
}

func (s *DeleteAgenticSpaceRequest) Validate() error {
	return dara.Validate(s)
}
