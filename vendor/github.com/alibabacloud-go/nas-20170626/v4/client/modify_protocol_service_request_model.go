// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyProtocolServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifyProtocolServiceRequest
	GetClientToken() *string
	SetDescription(v string) *ModifyProtocolServiceRequest
	GetDescription() *string
	SetDryRun(v bool) *ModifyProtocolServiceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *ModifyProtocolServiceRequest
	GetFileSystemId() *string
	SetProtocolServiceId(v string) *ModifyProtocolServiceRequest
	GetProtocolServiceId() *string
}

type ModifyProtocolServiceRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-42665544****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the protocol service.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter or a Chinese character and cannot start with `http://` or `https://`.
	//
	// - The description can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// - **Spaces and other special characters are not allowed**. Valid example: `My-Protocol-Service_01`.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// My-Protocol-Service_01
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this modification request.
	//
	// A dry run checks parameter validity and dependencies without actually modifying the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without modifying the protocol service. The dry run checks whether required parameters are specified, whether the request format is valid, and whether business constraints and dependencies are met. If the check fails, the corresponding error is returned. If the check succeeds, HTTP status code 200 is returned.
	//
	// - false (default): Sends a normal request. After the check succeeds, the protocol service is directly modified.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the protocol service. You can obtain the protocol service ID from the response of the [CreateProtocolService](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-createprotocolservice-cpfs) operation, or query it by calling the [DescribeProtocolService](https://www.alibabacloud.com/help/en/cpfs/cpfsonecs/developer-reference/api-nas-2017-06-26-describeprotocolservice-cpfs) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// ptc-197ed6a00f2b****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
}

func (s ModifyProtocolServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *ModifyProtocolServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyProtocolServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyProtocolServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *ModifyProtocolServiceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *ModifyProtocolServiceRequest) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *ModifyProtocolServiceRequest) SetClientToken(v string) *ModifyProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetDescription(v string) *ModifyProtocolServiceRequest {
	s.Description = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetDryRun(v bool) *ModifyProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetFileSystemId(v string) *ModifyProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *ModifyProtocolServiceRequest) SetProtocolServiceId(v string) *ModifyProtocolServiceRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *ModifyProtocolServiceRequest) Validate() error {
	return dara.Validate(s)
}
