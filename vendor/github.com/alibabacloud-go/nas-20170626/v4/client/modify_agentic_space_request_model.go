// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAgenticSpaceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAgenticSpaceId(v string) *ModifyAgenticSpaceRequest
	GetAgenticSpaceId() *string
	SetClientToken(v string) *ModifyAgenticSpaceRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyAgenticSpaceRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyAgenticSpaceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *ModifyAgenticSpaceRequest
	GetFileSystemId() *string
}

type ModifyAgenticSpaceRequest struct {
	// AgenticSpace Id。
	//
	// This parameter is required.
	//
	// example:
	//
	// agentic-229oypxjgpau2****
	AgenticSpaceId *string `json:"AgenticSpaceId,omitempty" xml:"AgenticSpaceId,omitempty"`
	// Ensures the idempotence of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may differ for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the AgenticSpace.
	//
	// example:
	//
	// Agentic Space Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and resource availability without actually modifying the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a check request without modifying the instance. The check items include whether required parameters are specified, request format, business limits, and NAS resource availability. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned, but FileSystemId is empty.
	//
	// - false (default): Sends a normal request and directly modifies the instance after the check passes.
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
	// 31a8e4****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
}

func (s ModifyAgenticSpaceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyAgenticSpaceRequest) GoString() string {
	return s.String()
}

func (s *ModifyAgenticSpaceRequest) GetAgenticSpaceId() *string {
	return s.AgenticSpaceId
}

func (s *ModifyAgenticSpaceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyAgenticSpaceRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyAgenticSpaceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyAgenticSpaceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyAgenticSpaceRequest) SetAgenticSpaceId(v string) *ModifyAgenticSpaceRequest {
	s.AgenticSpaceId = &v
	return s
}

func (s *ModifyAgenticSpaceRequest) SetClientToken(v string) *ModifyAgenticSpaceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyAgenticSpaceRequest) SetDescription(v string) *ModifyAgenticSpaceRequest {
	s.Description = &v
	return s
}

func (s *ModifyAgenticSpaceRequest) SetDryRun(v bool) *ModifyAgenticSpaceRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyAgenticSpaceRequest) SetFileSystemId(v string) *ModifyAgenticSpaceRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyAgenticSpaceRequest) Validate() error {
	return dara.Validate(s)
}
