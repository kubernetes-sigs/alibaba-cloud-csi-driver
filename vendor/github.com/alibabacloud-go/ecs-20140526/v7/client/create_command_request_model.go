// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandContent(v string) *CreateCommandRequest
	GetCommandContent() *string
	SetContentEncoding(v string) *CreateCommandRequest
	GetContentEncoding() *string
	SetDescription(v string) *CreateCommandRequest
	GetDescription() *string
	SetEnableParameter(v bool) *CreateCommandRequest
	GetEnableParameter() *bool
	SetLauncher(v string) *CreateCommandRequest
	GetLauncher() *string
	SetName(v string) *CreateCommandRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateCommandRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateCommandRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateCommandRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateCommandRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateCommandRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateCommandRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateCommandRequestTag) *CreateCommandRequest
	GetTag() []*CreateCommandRequestTag
	SetTimeout(v int64) *CreateCommandRequest
	GetTimeout() *int64
	SetType(v string) *CreateCommandRequest
	GetType() *string
	SetWorkingDir(v string) *CreateCommandRequest
	GetWorkingDir() *string
}

type CreateCommandRequest struct {
	// The Base64-encoded content of the command. Take note of the following items:
	//
	// 	- The value must be Base64-encoded and cannot exceed 18 KB in size.
	//
	// 	- You can use custom parameters in the command content. To enable the custom parameter feature, you must set `EnableParameter` to true.
	//
	//     	- Custom parameters are defined in the `{{}}` format. Within `{{}}`, the spaces and line feeds before and after the parameter names are ignored.
	//
	//     	- You can specify up to 20 custom parameters.
	//
	//     	- A custom parameter name can contain only letters, digits, underscores (_), and hyphens (-). The name is case-insensitive. The ACS:: prefix cannot be used to specify non-built-in environment parameters.
	//
	//     	- Each custom parameter name can be up to 64 bytes in length.
	//
	// 	- You can specify built-in environment parameters as custom parameters in a command. When you run the command, Cloud Assistant automatically uses the environment parameter values for the custom parameters. You can specify the following built-in environment variables:
	//
	//     	- `{{ACS::RegionId}}`: the region ID.
	//
	//     	- `{{ACS::AccountId}}`: the UID of the Alibaba Cloud account.
	//
	//     	- `{{ACS::InstanceId}}`: the instance ID. If you want to run the command on multiple instances and specify `{{ACS::InstanceId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following versions:
	//
	//         	- Linux: 2.2.3.309
	//
	//         	- Windows: 2.1.3.309
	//
	//     	- `{{ACS::InstanceName}}`: the instance name. If you want to run the command on multiple instances and specify `{{ACS::InstanceName}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following versions:
	//
	//         	- Linux: 2.2.3.344
	//
	//         	- Windows: 2.1.3.344
	//
	//     	- `{{ACS::InvokeId}}`: the ID of the task. If you want to specify `{{ACS::InvokeId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following versions:
	//
	//         	- Linux: 2.2.3.309
	//
	//         	- Windows: 2.1.3.309
	//
	//     	- `{{ACS::CommandId}}`: the command ID. If you want to call the [RunCommand](https://help.aliyun.com/document_detail/141751.html) operation to run the command and specify `{{ACS::CommandId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following versions:
	//
	//         	- Linux: 2.2.3.309
	//
	//         	- Windows: 2.1.3.309
	//
	// This parameter is required.
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The encoding mode of the command content (CommandContent). Valid values:
	//
	// 	- PlainText: The command content is not encoded.
	//
	// 	- Base64: The command content is Base64-encoded.
	//
	// Default value: Base64.
	//
	// > If the specified value of this parameter is invalid, Base64 is used by default.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The description of the command. The description supports all character sets and can be up to 512 characters in length.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to use custom parameters in the command.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The launcher for script execution. The value cannot exceed 1 KB in length.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The name of the command. The name supports all character sets and can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// testName
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The ID of the region in which to create the command. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the command.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags to add to the command.
	Tag []*CreateCommandRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// he maximum timeout period for the command execution on the instance. Unit: seconds. When a command that you created cannot be run, the command times out. When a command execution times out, Cloud Assistant Agent forcefully terminates the command process by canceling the PID.
	//
	// Default value: 60.
	//
	// example:
	//
	// 60
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The command type. Valid values:
	//
	// 	- RunBatScript: batch commands. These commands are applicable to Windows instances.
	//
	// 	- RunPowerShellScript: PowerShell commands. These commands are applicable to Windows instances.
	//
	// 	- RunShellScript: shell commands. These commands are applicable to Linux instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The execution path of the command on ECS instances. The value can be up to 200 characters in length.
	//
	// Default values:
	//
	// 	- For Linux instance, the default value is the home directory of the root user, which is the `/root` directory.
	//
	// 	- For Windows instances, the default value is the directory where the Cloud Assistant Agent process resides, such as `C:\\Windows\\System32\\`.
	//
	// >  If you set WorkingDir to a directory other than default ones, make sure that the directory exists on the instances.
	//
	// example:
	//
	// /root/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s CreateCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandRequest) GoString() string {
	return s.String()
}

func (s *CreateCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *CreateCommandRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *CreateCommandRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateCommandRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *CreateCommandRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *CreateCommandRequest) GetName() *string {
	return s.Name
}

func (s *CreateCommandRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateCommandRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateCommandRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateCommandRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateCommandRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateCommandRequest) GetTag() []*CreateCommandRequestTag {
	return s.Tag
}

func (s *CreateCommandRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *CreateCommandRequest) GetType() *string {
	return s.Type
}

func (s *CreateCommandRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *CreateCommandRequest) SetCommandContent(v string) *CreateCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *CreateCommandRequest) SetContentEncoding(v string) *CreateCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *CreateCommandRequest) SetDescription(v string) *CreateCommandRequest {
	s.Description = &v
	return s
}

func (s *CreateCommandRequest) SetEnableParameter(v bool) *CreateCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *CreateCommandRequest) SetLauncher(v string) *CreateCommandRequest {
	s.Launcher = &v
	return s
}

func (s *CreateCommandRequest) SetName(v string) *CreateCommandRequest {
	s.Name = &v
	return s
}

func (s *CreateCommandRequest) SetOwnerAccount(v string) *CreateCommandRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateCommandRequest) SetOwnerId(v int64) *CreateCommandRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateCommandRequest) SetRegionId(v string) *CreateCommandRequest {
	s.RegionId = &v
	return s
}

func (s *CreateCommandRequest) SetResourceGroupId(v string) *CreateCommandRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateCommandRequest) SetResourceOwnerAccount(v string) *CreateCommandRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateCommandRequest) SetResourceOwnerId(v int64) *CreateCommandRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateCommandRequest) SetTag(v []*CreateCommandRequestTag) *CreateCommandRequest {
	s.Tag = v
	return s
}

func (s *CreateCommandRequest) SetTimeout(v int64) *CreateCommandRequest {
	s.Timeout = &v
	return s
}

func (s *CreateCommandRequest) SetType(v string) *CreateCommandRequest {
	s.Type = &v
	return s
}

func (s *CreateCommandRequest) SetWorkingDir(v string) *CreateCommandRequest {
	s.WorkingDir = &v
	return s
}

func (s *CreateCommandRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateCommandRequestTag struct {
	// The key of tag N. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that have this tag added can be displayed in the response. If multiple tags are specified to query resources, up to 1,000 resources that have all these tags added can be displayed in the response. To query more than 1,000 resources that have specified tags, call [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// It can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateCommandRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateCommandRequestTag) GoString() string {
	return s.String()
}

func (s *CreateCommandRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateCommandRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateCommandRequestTag) SetKey(v string) *CreateCommandRequestTag {
	s.Key = &v
	return s
}

func (s *CreateCommandRequestTag) SetValue(v string) *CreateCommandRequestTag {
	s.Value = &v
	return s
}

func (s *CreateCommandRequestTag) Validate() error {
	return dara.Validate(s)
}
