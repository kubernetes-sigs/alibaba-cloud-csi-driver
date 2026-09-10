// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtocolServiceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateProtocolServiceRequest
	GetClientToken() *string
	SetDescription(v string) *CreateProtocolServiceRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateProtocolServiceRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CreateProtocolServiceRequest
	GetFileSystemId() *string
	SetProtocolSpec(v string) *CreateProtocolServiceRequest
	GetProtocolSpec() *string
	SetProtocolType(v string) *CreateProtocolServiceRequest
	GetProtocolType() *string
	SetThroughput(v int32) *CreateProtocolServiceRequest
	GetThroughput() *int32
	SetVSwitchId(v string) *CreateProtocolServiceRequest
	GetVSwitchId() *string
	SetVpcId(v string) *CreateProtocolServiceRequest
	GetVpcId() *string
}

type CreateProtocolServiceRequest struct {
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
	// The description of the protocol service. This value is displayed as "Protocol service name" in the console.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length.
	//
	// - The description must start with a letter or Chinese character and cannot start with `http://` or `https://`.
	//
	// - The description can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// 此协议服务的描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this request.
	//
	// A dry run checks parameter validity and dependencies without actually creating the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a check request without creating the protocol service. The check items include whether required parameters are specified, the request format, and business limit dependencies. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned, but ProtocolServiceId is empty.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is directly created.
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
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The specification of the protocol service.
	//
	// Valid values: General (default).
	//
	// This parameter is required.
	//
	// example:
	//
	// General
	ProtocolSpec *string `json:"ProtocolSpec,omitempty" xml:"ProtocolSpec,omitempty"`
	// The protocol type of the protocol service.
	//
	// Valid values: NFS (default). Only NFSv3 access is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The bandwidth of the protocol service.
	//
	// Unit: MB/s.
	//
	// example:
	//
	// 8000
	Throughput *int32 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	// The vSwitch ID of the protocol service.
	//
	// If the storage redundancy type of the file system is zone-redundant storage (ZRS), do not set this parameter. Otherwise, this parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vsw-2vc3c2lybvdllxyq4****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID of the protocol service. The VPC must be the same as the VPC of the file system.
	//
	// If the storage redundancy type of the file system is zone-redundant storage (ZRS), do not set this parameter. Otherwise, this parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vpc-2vct297b8157bth9z****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateProtocolServiceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProtocolServiceRequest) GoString() string {
	return s.String()
}

func (s *CreateProtocolServiceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProtocolServiceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProtocolServiceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateProtocolServiceRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateProtocolServiceRequest) GetProtocolSpec() *string {
	return s.ProtocolSpec
}

func (s *CreateProtocolServiceRequest) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *CreateProtocolServiceRequest) GetThroughput() *int32 {
	return s.Throughput
}

func (s *CreateProtocolServiceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateProtocolServiceRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateProtocolServiceRequest) SetClientToken(v string) *CreateProtocolServiceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetDescription(v string) *CreateProtocolServiceRequest {
	s.Description = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetDryRun(v bool) *CreateProtocolServiceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetFileSystemId(v string) *CreateProtocolServiceRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetProtocolSpec(v string) *CreateProtocolServiceRequest {
	s.ProtocolSpec = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetProtocolType(v string) *CreateProtocolServiceRequest {
	s.ProtocolType = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetThroughput(v int32) *CreateProtocolServiceRequest {
	s.Throughput = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetVSwitchId(v string) *CreateProtocolServiceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateProtocolServiceRequest) SetVpcId(v string) *CreateProtocolServiceRequest {
	s.VpcId = &v
	return s
}

func (s *CreateProtocolServiceRequest) Validate() error {
	return dara.Validate(s)
}
