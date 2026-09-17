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
	// Ensures the idempotence of the request. Generate a parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// > If you do not specify this parameter, the system automatically uses the RequestId of the API request as the ClientToken. The RequestId may vary for each API request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the protocol service export directory. This is displayed as **Export Directory Name*	- in the console.
	//
	// Limits:
	//
	// - The description must be 2 to 128 characters in length and can contain letters and Chinese characters.
	//
	// - The description must start with a letter or a Chinese character. It cannot start with `http://` or `https://`.
	//
	// - The description can contain digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// Description of this export directory
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform a dry run for this request. A dry run checks parameter validity and dependencies without actually creating the instance or incurring charges.
	//
	// Valid values:
	//
	// - true: Sends a dry run request without creating the export directory. The check items include whether required parameters are specified, request format, and business limit dependencies. If the check fails, the corresponding error is returned. If the check passes, HTTP status code 200 is returned, but ExportId is empty.
	//
	// - false (default): Sends a normal request. After the check passes, the instance is directly created.
	//
	// example:
	//
	// true
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The ID of the file system.
	//
	// > The target CPFS file system must be in the Running state before you call this operation. After creation, it takes approximately 20 to 40 minutes for the file system to reach the Running state. You can call [DescribeFileSystems](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-describefilesystems) to poll the Status field (Running indicates ready).
	//
	// This parameter is required.
	//
	// example:
	//
	// cpfs-099394bd928c****
	FileSystemId *string `json:"FileSystemId,omitempty" xml:"FileSystemId,omitempty"`
	// The ID of the fileset to export.
	//
	// Limits:
	//
	// - The fileset must already exist.
	//
	// - Only one export directory can be created for each fileset.
	//
	// - You must specify one and only one of FsetId and Path.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The path of the CPFS directory to export.
	//
	// Limits:
	//
	// - The directory must already exist on the CPFS file system.
	//
	// - Only one export can be created for the same directory.
	//
	// - You must specify one and only one of FsetId and Path.
	//
	// Format:
	//
	// - The path must be 1 to 1,024 characters in length.
	//
	// - The path must be encoded in UTF-8.
	//
	// - The path must start and end with a forward slash (/). The root directory is `/`.
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
	// ptc-197ed6a00f2b****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	// The ID of the vSwitch for the protocol service export.
	//
	// If the storage redundancy type of the file system is not zone-redundant storage (ZRS), this parameter is required when VpcId is specified.
	//
	// > The vSwitch must be in the same zone as the target CPFS file system (the ZoneId values must match). You can call [DescribeFileSystems](https://www.alibabacloud.com/help/en/nas/developer-reference/api-nas-2017-06-26-describefilesystems) to query the ZoneId field of the file system, and then select a vSwitch in the same zone.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vsw-2vc3c2lybvdllxyq4****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The list of vSwitch IDs for the protocol service export.
	//
	// if can be null:
	// true
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The ID of the VPC for the protocol service export.
	//
	// > The CIDR block of the selected VPC must not overlap with the VPC CIDR block of the target CPFS file system. You can call the DescribeVpcs operation of VPC to query the CidrBlock field of each VPC for comparison, and select a VPC that does not conflict and is in the same zone as the file system.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// vpc-2vct297b8157bth9z****
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
