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
	// The description of the protocol service. The name of the protocol service appears in the console.
	//
	// Limits:
	//
	// 	- The description must be 2 to 128 characters in length.
	//
	// 	- The description must start with a letter and cannot start with `http://` or `https://`.
	//
	// 	- The description can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run.
	//
	// The dry run checks parameter validity and prerequisites. The dry run does not create a protocol service or incur fees.
	//
	// Valid values:
	//
	// 	- true: performs only a dry run and does not create the protocol service. The system checks the request format, service limits, prerequisites, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the ProtocolServiceId parameter.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, a protocol service is created.
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
	// cpfs-123****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The specification of the protocol service.
	//
	// Set the value to General (default).
	//
	// Valid values:
	//
	// 	- CL2
	//
	// 	- General
	//
	// 	- CL1
	//
	// This parameter is required.
	//
	// example:
	//
	// General
	ProtocolSpec *string `json:"ProtocolSpec,omitempty" xml:"ProtocolSpec,omitempty"`
	// The protocol type of the protocol service.
	//
	// Valid value: NFS (default). Only NFSv3 is supported.
	//
	// This parameter is required.
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The throughput of the protocol service.
	//
	// Unit: MB/s.
	//
	// example:
	//
	// 8000
	Throughput *int32 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	// The vSwitch ID of the protocol service.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vsw-123****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID of the protocol service. The VPC ID of the protocol service must be the same as the VPC ID of the file system.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vpc-123****
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
