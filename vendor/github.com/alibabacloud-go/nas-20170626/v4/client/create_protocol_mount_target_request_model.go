// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProtocolMountTargetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAccessGroupName(v string) *CreateProtocolMountTargetRequest
	GetAccessGroupName() *string
	SetClientToken(v string) *CreateProtocolMountTargetRequest
	GetClientToken() *string
	SetDescription(v string) *CreateProtocolMountTargetRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateProtocolMountTargetRequest
	GetDryRun() *bool
	SetFileSystemId(v string) *CreateProtocolMountTargetRequest
	GetFileSystemId() *string
	SetFsetId(v string) *CreateProtocolMountTargetRequest
	GetFsetId() *string
	SetPath(v string) *CreateProtocolMountTargetRequest
	GetPath() *string
	SetProtocolServiceId(v string) *CreateProtocolMountTargetRequest
	GetProtocolServiceId() *string
	SetVSwitchId(v string) *CreateProtocolMountTargetRequest
	GetVSwitchId() *string
	SetVSwitchIds(v []*string) *CreateProtocolMountTargetRequest
	GetVSwitchIds() []*string
	SetVpcId(v string) *CreateProtocolMountTargetRequest
	GetVpcId() *string
}

type CreateProtocolMountTargetRequest struct {
	// The name of the permission group.
	//
	// Default value: DEFAULT_VPC_GROUP_NAME.
	//
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests.
	//
	// The token can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How do I ensure the idempotence?](https://help.aliyun.com/document_detail/25693.html)
	//
	// >  If you do not specify this parameter, the system automatically uses the request ID as the client token. The request ID may be different for each request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the export directory for the protocol service. The **name of the export directory*	- appears in the console.
	//
	// Limits:
	//
	// 	- The description must be 2 to 128 characters in length.
	//
	// 	- The description must start with a letter but cannot start with `http://` or `https://`.
	//
	// 	- The description can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// test
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run. The dry run checks parameter validity and prerequisites. The dry run does not create an export directory or incur fees.
	//
	// Valid values:
	//
	// 	- true: performs a dry run. The system checks the request format, service limits, prerequisites, and whether the required parameters are specified. If the request fails the dry run, an error message is returned. If the request passes the dry run, the HTTP status code 200 is returned. No value is returned for the ExportId parameter.
	//
	// 	- false (default): performs a dry run and sends the request. If the request passes the dry run, an export directory is created.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the file system.
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-123****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the fileset that you want to export.
	//
	// Limits:
	//
	// 	- The fileset already exists.
	//
	// 	- You can create only one export directory for a fileset.
	//
	// 	- You can specify either a fileset or a path.
	//
	// example:
	//
	// fset-123****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The path of the CPFS directory that you want to export.
	//
	// Limits:
	//
	// 	- The directory already exists in the CPFS file system.
	//
	// 	- You can create only one export directory for a directory.
	//
	// 	- You can specify either a fileset or a path.
	//
	// Format:
	//
	// 	- The path must be 1 to 1,024 characters in length.
	//
	// 	- The path must be encoded in UTF-8.
	//
	// 	- The path must start and end with a forward slash (/). The root directory is `/`.
	//
	// example:
	//
	// /
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The ID of the protocol service.
	//
	// This parameter is required.
	//
	// example:
	//
	// ptc-123****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	// The vSwitch ID of the export directory.
	//
	// If the storage redundancy type of the file system is not zone-redundant (ZRS) and the VpcId is set, this field is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vsw-123****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The list of vSwitch IDs of the export directory.
	//
	// if can be null:
	// true
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID of the export directory.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vpc-123****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateProtocolMountTargetRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateProtocolMountTargetRequest) GoString() string {
	return s.String()
}

func (s *CreateProtocolMountTargetRequest) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *CreateProtocolMountTargetRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateProtocolMountTargetRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateProtocolMountTargetRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateProtocolMountTargetRequest) GetFileSystemId() *string {
	return s.FileSystemId
}

func (s *CreateProtocolMountTargetRequest) GetFsetId() *string {
	return s.FsetId
}

func (s *CreateProtocolMountTargetRequest) GetPath() *string {
	return s.Path
}

func (s *CreateProtocolMountTargetRequest) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *CreateProtocolMountTargetRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateProtocolMountTargetRequest) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *CreateProtocolMountTargetRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateProtocolMountTargetRequest) SetAccessGroupName(v string) *CreateProtocolMountTargetRequest {
	s.AccessGroupName = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetClientToken(v string) *CreateProtocolMountTargetRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetDescription(v string) *CreateProtocolMountTargetRequest {
	s.Description = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetDryRun(v bool) *CreateProtocolMountTargetRequest {
	s.DryRun = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetFileSystemId(v string) *CreateProtocolMountTargetRequest {
	s.FileSystemId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetFsetId(v string) *CreateProtocolMountTargetRequest {
	s.FsetId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetPath(v string) *CreateProtocolMountTargetRequest {
	s.Path = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetProtocolServiceId(v string) *CreateProtocolMountTargetRequest {
	s.ProtocolServiceId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetVSwitchId(v string) *CreateProtocolMountTargetRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetVSwitchIds(v []*string) *CreateProtocolMountTargetRequest {
	s.VSwitchIds = v
	return s
}

func (s *CreateProtocolMountTargetRequest) SetVpcId(v string) *CreateProtocolMountTargetRequest {
	s.VpcId = &v
	return s
}

func (s *CreateProtocolMountTargetRequest) Validate() error {
	return dara.Validate(s)
}
