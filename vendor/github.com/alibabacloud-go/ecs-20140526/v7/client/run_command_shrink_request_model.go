// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunCommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RunCommandShrinkRequest
	GetClientToken() *string
	SetCommandContent(v string) *RunCommandShrinkRequest
	GetCommandContent() *string
	SetContainerId(v string) *RunCommandShrinkRequest
	GetContainerId() *string
	SetContainerName(v string) *RunCommandShrinkRequest
	GetContainerName() *string
	SetContentEncoding(v string) *RunCommandShrinkRequest
	GetContentEncoding() *string
	SetDescription(v string) *RunCommandShrinkRequest
	GetDescription() *string
	SetEnableParameter(v bool) *RunCommandShrinkRequest
	GetEnableParameter() *bool
	SetFrequency(v string) *RunCommandShrinkRequest
	GetFrequency() *string
	SetInstanceId(v []*string) *RunCommandShrinkRequest
	GetInstanceId() []*string
	SetKeepCommand(v bool) *RunCommandShrinkRequest
	GetKeepCommand() *bool
	SetLauncher(v string) *RunCommandShrinkRequest
	GetLauncher() *string
	SetName(v string) *RunCommandShrinkRequest
	GetName() *string
	SetOssOutputDelivery(v string) *RunCommandShrinkRequest
	GetOssOutputDelivery() *string
	SetOwnerAccount(v string) *RunCommandShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RunCommandShrinkRequest
	GetOwnerId() *int64
	SetParametersShrink(v string) *RunCommandShrinkRequest
	GetParametersShrink() *string
	SetRegionId(v string) *RunCommandShrinkRequest
	GetRegionId() *string
	SetRepeatMode(v string) *RunCommandShrinkRequest
	GetRepeatMode() *string
	SetResourceGroupId(v string) *RunCommandShrinkRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *RunCommandShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RunCommandShrinkRequest
	GetResourceOwnerId() *int64
	SetResourceTag(v []*RunCommandShrinkRequestResourceTag) *RunCommandShrinkRequest
	GetResourceTag() []*RunCommandShrinkRequestResourceTag
	SetTag(v []*RunCommandShrinkRequestTag) *RunCommandShrinkRequest
	GetTag() []*RunCommandShrinkRequestTag
	SetTerminationMode(v string) *RunCommandShrinkRequest
	GetTerminationMode() *string
	SetTimed(v bool) *RunCommandShrinkRequest
	GetTimed() *bool
	SetTimeout(v int64) *RunCommandShrinkRequest
	GetTimeout() *int64
	SetType(v string) *RunCommandShrinkRequest
	GetType() *string
	SetUsername(v string) *RunCommandShrinkRequest
	GetUsername() *string
	SetWindowsPasswordName(v string) *RunCommandShrinkRequest
	GetWindowsPasswordName() *string
	SetWorkingDir(v string) *RunCommandShrinkRequest
	GetWorkingDir() *string
}

type RunCommandShrinkRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The content of the command. The command content can be plaintext or Base64-encoded. Take note of the following items:
	//
	// 	- If you want to retain the command, make sure that the size of the Base64-encoded command content does not exceed 18 KB. If you do not want to retain the command, make sure that the size of the Base64-encoded command content does not exceed 24 KB. You can set `KeepCommand` to specify whether to retain the command.
	//
	// 	- If the command content is Base64-encoded, set `ContentEncoding` to Base64.
	//
	// 	- If you specify `EnableParameter` to true, the custom parameter feature is enable. You can configure custom parameters based on the following rules:
	//
	//     	- Specify custom parameters in the `{{}}` format. The spaces and line feeds before and after the parameter names within `{{}}` are ignored.
	//
	//     	- You can specify up to 20 custom parameters.
	//
	//     	- A custom parameter name can contain letters, digits, underscores (_), and hyphens (-). The name is case-insensitive. The ACS:: prefix cannot be used to specify non-built-in environment parameters.
	//
	//     	- Each custom parameter name cannot exceed 64 bytes in length.
	//
	// 	- You can specify built-in environment parameters as custom parameters. When you run a command, the parameters are automatically specified by Cloud Assistant. You can specify the following built-in environment parameters:
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
	//     	- `{{ACS::InvokeId}}`: the task ID. If you want to specify `{{ACS::InvokeId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following versions:
	//
	//         	- Linux: 2.2.3.309
	//
	//         	- Windows: 2.1.3.309
	//
	//     	- `{{ACS::CommandId}}`: the command ID. If you want to specify `{{ACS::CommandId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following versions:
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
	// The container ID. Only 64-bit hexadecimal strings are supported. `docker://`, `containerd://`, or `cri-o://` can be used as the prefix of the container ID to specify the container runtime.
	//
	// Take note of the following items:
	//
	// 	- If you specify this parameter, Cloud Assistant runs the command in the specified container of the instances.
	//
	// 	- If you specify this parameter, make sure that the Cloud Assistant Agent version installed on Linux instances is 2.2.3.344 or later.
	//
	// 	- If you specify this parameter, `Username` and `WorkingDir` do not take effect. You can run the command in the default working directory of the container by using only the default user of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// >  Only shell scripts can run in Linux containers. You cannot add a command whose format is similar to `#!/usr/bin/python` at the beginning of a script to specify a script interpreter. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****
	ContainerId *string `json:"ContainerId,omitempty" xml:"ContainerId,omitempty"`
	// The container name.
	//
	// Take note of the following items:
	//
	// 	- If you specify this parameter, Cloud Assistant runs the command in the specified container of the instances.
	//
	// 	- If you specify this parameter, make sure that the Cloud Assistant Agent version installed on Linux instances is 2.2.3.344 or later.
	//
	// 	- If you specify this parameter, `Username` and `WorkingDir` do not take effect. You can run the command in the default working directory of the container by using only the default user of the container. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// >  Only shell scripts can run in Linux containers. You cannot add a command whose format is similar to `#!/usr/bin/python` at the beginning of a script to specify a script interpreter. For more information, see [Use Cloud Assistant to run commands in containers](https://help.aliyun.com/document_detail/456641.html).
	//
	// example:
	//
	// test-container
	ContainerName *string `json:"ContainerName,omitempty" xml:"ContainerName,omitempty"`
	// The encoding mode of command content (`CommandContent`). The valid values are case-insensitive. Valid values:
	//
	// 	- PlainText: The command content is not encoded.
	//
	// 	- Base64: The command content is encoded in Base64.
	//
	// Default value: PlainText. If the specified value of this parameter is invalid, PlainText is used by default.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The description of the command. The description supports all character sets and can be up to 512 characters in length.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to include custom parameters in the command.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The schedule on which to run the command. You can configure a command to run at a fixed interval based on a rate expression, run only once at a specified time, or run at designated times based on a cron expression.
	//
	// 	- To run a command at a fixed interval, use a rate expression to specify the interval. You can specify the interval in seconds, minutes, hours, or days. This option is suitable for scenarios in which tasks need to be executed at a fixed interval. Specify the interval in the following format: `rate(<Execution interval value> <Execution interval unit>)`. For example, specify `rate(5m)` to run the command every 5 minutes. When you specify an interval, take note of the following limits:
	//
	//     	- The interval can be anywhere from 60 seconds to 7 days, but must be longer than the timeout period of the scheduled task.
	//
	//     	- The interval is the amount of time that elapses between two consecutive executions. The interval is irrelevant to the amount of time that is required to run the command once. For example, assume that you set the interval to 5 minutes and that it takes 2 minutes to run the command each time. Each time the command is run, the system waits 3 minutes before the system runs the command again.
	//
	//     	- A task is not immediately executed after the task is created. For example, assume that you set the interval to 5 minutes for a task. The task begins to be executed 5 minutes after it is created.
	//
	// 	- To run a command only once at a specific time, specify a point in time and a time zone. Specify the point in time in the `at(yyyy-MM-dd HH:mm:ss <Time zone>)` format, which indicates `at(Year-Month-Day Hour:Minute:Second <Time zone>)`. If you do not specify a time zone, the UTC time zone is used by default. You can specify the time zone in the following forms:
	//
	//     	- The time zone name. Examples: `Asia/Shanghai` and `America/Los_Angeles`.
	//
	//     	- The time offset from GMT. Examples: `GMT+8:00` (UTC+8) and `GMT-7:00` (UTC-7). If you use the GMT format, you cannot add leading zeros to the hour value.
	//
	//     	- The time zone abbreviation. Only UTC is supported.
	//
	//     For example, to configure a command to run only once at 13:15:30 on June 6, 2022 (Shanghai time), set the time to `at(2022-06-06 13:15:30 Asia/Shanghai)`. To configure a command to run only once at 13:15:30 on June 6, 2022 (UTC-7), set the time to `at(2022-06-06 13:15:30 GMT-7:00)`.
	//
	// 	- To run a command at specific times, use a cron expression to define the schedule. Specify a schedule in the `<Cron expression> <Time zone>` format. The cron expression is in the `<seconds> <minutes> <hours> <day of the month> <month> <day of the week> <year (optional)>` format. The system calculates the execution times of the command based on the specified cron expression and time zone and runs the command as scheduled. If you do not specify a time zone, the system time zone of the instance on which you want to run the command is used by default. For more information about cron expressions, see [Cron expressions](https://help.aliyun.com/document_detail/64769.html). You can specify the time zone in the following forms:
	//
	//     	- The time zone name. Examples: `Asia/Shanghai` and `America/Los_Angeles`.
	//
	//     	- The time offset from GMT. Examples: `GMT+8:00` (UTC+8) and `GMT-7:00` (UTC-7). If you use the GMT format, you cannot add leading zeros to the hour value.
	//
	//     	- The time zone abbreviation. Only UTC is supported. For example, to configure a command to run at 10:15:00 every day in 2022 (Shanghai time), set the schedule to `0 15 10 ? 	- 	- 2022 Asia/Shanghai`. To configure a command to run every half an hour from 10:00:00 to 11:30:00 every day in 2022 (UTC+8), set the schedule to `0 0/30 10-11 	- 	- ? 2022 GMT+8:00`. To configure a command to run every 5 minutes from 14:00:00 to 14:55:00 every October every two years from 2022 in UTC, set the schedule to `0 0/5 14 	- 10 ? 2022/2 UTC`.
	//
	//     **
	//
	//     **Note*	- The minimum interval must be 10 seconds or more and cannot be shorter than the timeout period of scheduled executions.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The IDs of instances on which to create and run the command. Specify at least one instance ID. You can specify up to 100 instance IDs.
	//
	// If one of the specified instances does not meet the conditions for running the command, the call fails. To ensure that the call is successful, specify only the IDs of instances that meet the conditions.
	//
	// You can request a quota increase in the Quota Center console. The quota name is Maximum number of instances supported for command execution.
	//
	// example:
	//
	// i-bp185dy2o3o6neg****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// Specifies whether to retain the command after the command is run. Valid values:
	//
	// 	- true: retains the command. Then, you can call the InvokeCommand operation to rerun the command. The retained command counts against the quota of Cloud Assistant commands.
	//
	// 	- false: does not retain the command. The command is automatically deleted after it is run and does not count against the quota of Cloud Assistant commands.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	KeepCommand *bool `json:"KeepCommand,omitempty" xml:"KeepCommand,omitempty"`
	// The launcher for script execution. The value cannot exceed 1 KB in length.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The name of the command. The name supports all character sets and can be up to 128 characters in length.
	//
	// example:
	//
	// testName
	Name              *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OssOutputDelivery *string `json:"OssOutputDelivery,omitempty" xml:"OssOutputDelivery,omitempty"`
	OwnerAccount      *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId           *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key-value pairs of custom parameters to pass in when the command can include custom parameters. For example, the command content is `echo {{name}}`. You can use `Parameters` to pass in the `{"name":"Jack"}` key-value pair. The `name` key of the custom parameter is automatically replaced by the paired Jack value to generate a new command. As a result, the `echo Jack` command is run.
	//
	// You can specify 0 to 10 custom parameters. Take note of the following items:
	//
	// 	- The key of a custom parameter can be up to 64 characters in length and cannot be an empty string.
	//
	// 	- The value of a custom parameter can be an empty string.
	//
	// 	- If you want to retain a command, make sure that the command after Base64 encoding, including custom parameters and original command content, does not exceed 18 KB in size. If you do not want to retain the command, make sure that the command after Base64 encoding does not exceed 24 KB in size. You can set `KeepCommand` to specify whether to retain the command.
	//
	// 	- The custom parameter names that are specified by Parameters must be included in the custom parameter names that you specified when you created the command. You can use empty strings to represent the parameters that are not passed in.
	//
	// This parameter is left empty by default, which indicates that the custom parameter feature is disabled.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI*************"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID of the command. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies how to run the command. Valid values:
	//
	// 	- Once: immediately runs the command.
	//
	// 	- Period: runs the command based on a schedule. If you set this parameter to `Period`, you must specify `Frequency`.
	//
	// 	- NextRebootOnly: runs the command the next time the instances start.
	//
	// 	- EveryReboot: runs the command every time the instances start.
	//
	// 	- DryRun: performs only a dry run, without running the actual command. The system checks the request parameters, the execution environments on the instances, and the status of Cloud Assistant Agent.
	//
	// Default value:
	//
	// 	- If you do not specify `Frequency`, the default value is `Once`.
	//
	// 	- If you specify `Frequency`, RepeatMode is set to `Period` regardless of whether a value is already specified for RepeatMode.
	//
	// Take note of the following items:
	//
	// 	- You can call the [StopInvocation](https://help.aliyun.com/document_detail/64838.html) operation to stop the pending or scheduled executions of the command.
	//
	// 	- If you set this parameter to `Period` or `EveryReboot`, you can call the [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) operation with `IncludeHistory` set to true to query the historical scheduled executions.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The ID of the resource group to which to assign the command executions. When you set this parameter, take note of the following items:
	//
	// 	- The instances specified by InstanceId.N must belong to the specified resource group.
	//
	// 	- After the command is run, you can set this parameter to call the [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) or [DescribeInvocationResults](https://help.aliyun.com/document_detail/64845.html) operation to query the execution results in the specified resource group.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the instance. You can leave this parameter empty or specify up to 20 tags. If you do not specify InstanceId, the command is run on instances that have the specified tags.
	ResourceTag []*RunCommandShrinkRequestResourceTag `json:"ResourceTag,omitempty" xml:"ResourceTag,omitempty" type:"Repeated"`
	// The tags to add to the command task. You can leave this parameter empty or specify up to 20 tags.
	Tag []*RunCommandShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies how to stop the command task when a command execution is manually stopped or times out. Valid values:
	//
	// 	- Process: stops the process of the command.
	//
	// 	- ProcessTree: stops the process tree of the command. In this case, the process of the command and all subprocesses of the process are stopped.
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// >  This parameter is no longer used and does not take effect.
	//
	// example:
	//
	// true
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// The timeout period for the command execution. Unit: seconds.
	//
	// A timeout error occurs if the command cannot be run because the process slows down or because a specific module or Cloud Assistant Agent does not exist. When an execution times out, the command process is forcefully terminated.
	//
	// Default value: 60.
	//
	// example:
	//
	// 3600
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The type of the command. Valid values:
	//
	// 	- RunBatScript: batch command, applicable to Windows instances.
	//
	// 	- RunPowerShellScript: PowerShell command, applicable to Windows instances.
	//
	// 	- RunShellScript: shell command, applicable to Linux instances.
	//
	// This parameter is required.
	//
	// example:
	//
	// RunShellScript
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The username to use to run the command on the ECS instances. The username cannot exceed 255 characters in length.
	//
	// 	- For Linux instances, the root username is used by default.
	//
	// 	- For Windows instances, the System username is used by default.
	//
	// You can also specify other usernames that already exist in the instances to run the command. For security purposes, we recommend that you run Cloud Assistant commands as a regular user. For more information, see [Run Cloud Assistant commands as a regular user](https://help.aliyun.com/document_detail/203771.html).
	//
	// example:
	//
	// test
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The name of the password to use to run the command on a Windows instance. The name cannot exceed 255 characters in length.
	//
	// If you do not want to use the default System user to run the command on Windows instances, specify both WindowsPasswordName and `Username`. To mitigate the risk of password leaks, the password is stored in plaintext in CloudOps Orchestration Service (OOS) Parameter Store, and only the name of the password is passed in by using WindowsPasswordName. For more information, see [Manage encryption parameters](https://help.aliyun.com/document_detail/186828.html) and [Run Cloud Assistant commands as a regular user](https://help.aliyun.com/document_detail/203771.html).
	//
	// >  If you use the root username for Linux instances or the System username for Windows instances to run the command, you do not need to specify WindowsPasswordName.
	//
	// example:
	//
	// axtSecretPassword
	WindowsPasswordName *string `json:"WindowsPasswordName,omitempty" xml:"WindowsPasswordName,omitempty"`
	// The working directory of the command on the instance. The value can be up to 200 characters in length.
	//
	// Default values:
	//
	// 	- For Linux instances, the default value is `/root`, which is the home directory of the administrator (the root user).
	//
	// 	- For Windows instances, the default value is the directory where the Cloud Assistant Agent process resides, such as `C:\\Windows\\System32`.
	//
	// example:
	//
	// /home/user
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s RunCommandShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s RunCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunCommandShrinkRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunCommandShrinkRequest) GetContainerId() *string {
	return s.ContainerId
}

func (s *RunCommandShrinkRequest) GetContainerName() *string {
	return s.ContainerName
}

func (s *RunCommandShrinkRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *RunCommandShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *RunCommandShrinkRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *RunCommandShrinkRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *RunCommandShrinkRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *RunCommandShrinkRequest) GetKeepCommand() *bool {
	return s.KeepCommand
}

func (s *RunCommandShrinkRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *RunCommandShrinkRequest) GetName() *string {
	return s.Name
}

func (s *RunCommandShrinkRequest) GetOssOutputDelivery() *string {
	return s.OssOutputDelivery
}

func (s *RunCommandShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RunCommandShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RunCommandShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *RunCommandShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunCommandShrinkRequest) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *RunCommandShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunCommandShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RunCommandShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RunCommandShrinkRequest) GetResourceTag() []*RunCommandShrinkRequestResourceTag {
	return s.ResourceTag
}

func (s *RunCommandShrinkRequest) GetTag() []*RunCommandShrinkRequestTag {
	return s.Tag
}

func (s *RunCommandShrinkRequest) GetTerminationMode() *string {
	return s.TerminationMode
}

func (s *RunCommandShrinkRequest) GetTimed() *bool {
	return s.Timed
}

func (s *RunCommandShrinkRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *RunCommandShrinkRequest) GetType() *string {
	return s.Type
}

func (s *RunCommandShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *RunCommandShrinkRequest) GetWindowsPasswordName() *string {
	return s.WindowsPasswordName
}

func (s *RunCommandShrinkRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *RunCommandShrinkRequest) SetClientToken(v string) *RunCommandShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RunCommandShrinkRequest) SetCommandContent(v string) *RunCommandShrinkRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandShrinkRequest) SetContainerId(v string) *RunCommandShrinkRequest {
	s.ContainerId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetContainerName(v string) *RunCommandShrinkRequest {
	s.ContainerName = &v
	return s
}

func (s *RunCommandShrinkRequest) SetContentEncoding(v string) *RunCommandShrinkRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandShrinkRequest) SetDescription(v string) *RunCommandShrinkRequest {
	s.Description = &v
	return s
}

func (s *RunCommandShrinkRequest) SetEnableParameter(v bool) *RunCommandShrinkRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandShrinkRequest) SetFrequency(v string) *RunCommandShrinkRequest {
	s.Frequency = &v
	return s
}

func (s *RunCommandShrinkRequest) SetInstanceId(v []*string) *RunCommandShrinkRequest {
	s.InstanceId = v
	return s
}

func (s *RunCommandShrinkRequest) SetKeepCommand(v bool) *RunCommandShrinkRequest {
	s.KeepCommand = &v
	return s
}

func (s *RunCommandShrinkRequest) SetLauncher(v string) *RunCommandShrinkRequest {
	s.Launcher = &v
	return s
}

func (s *RunCommandShrinkRequest) SetName(v string) *RunCommandShrinkRequest {
	s.Name = &v
	return s
}

func (s *RunCommandShrinkRequest) SetOssOutputDelivery(v string) *RunCommandShrinkRequest {
	s.OssOutputDelivery = &v
	return s
}

func (s *RunCommandShrinkRequest) SetOwnerAccount(v string) *RunCommandShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RunCommandShrinkRequest) SetOwnerId(v int64) *RunCommandShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetParametersShrink(v string) *RunCommandShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *RunCommandShrinkRequest) SetRegionId(v string) *RunCommandShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetRepeatMode(v string) *RunCommandShrinkRequest {
	s.RepeatMode = &v
	return s
}

func (s *RunCommandShrinkRequest) SetResourceGroupId(v string) *RunCommandShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetResourceOwnerAccount(v string) *RunCommandShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RunCommandShrinkRequest) SetResourceOwnerId(v int64) *RunCommandShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetResourceTag(v []*RunCommandShrinkRequestResourceTag) *RunCommandShrinkRequest {
	s.ResourceTag = v
	return s
}

func (s *RunCommandShrinkRequest) SetTag(v []*RunCommandShrinkRequestTag) *RunCommandShrinkRequest {
	s.Tag = v
	return s
}

func (s *RunCommandShrinkRequest) SetTerminationMode(v string) *RunCommandShrinkRequest {
	s.TerminationMode = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTimed(v bool) *RunCommandShrinkRequest {
	s.Timed = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTimeout(v int64) *RunCommandShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandShrinkRequest) SetType(v string) *RunCommandShrinkRequest {
	s.Type = &v
	return s
}

func (s *RunCommandShrinkRequest) SetUsername(v string) *RunCommandShrinkRequest {
	s.Username = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWindowsPasswordName(v string) *RunCommandShrinkRequest {
	s.WindowsPasswordName = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWorkingDir(v string) *RunCommandShrinkRequest {
	s.WorkingDir = &v
	return s
}

func (s *RunCommandShrinkRequest) Validate() error {
	if s.ResourceTag != nil {
		for _, item := range s.ResourceTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type RunCommandShrinkRequestResourceTag struct {
	// The key of tag N of the instance.
	//
	// Take note of the following items:
	//
	// 	- This parameter and InstanceId.N are mutually exclusive.
	//
	// 	- The tag key cannot be an empty string.
	//
	// 	- The number of instances that have the specified tags cannot exceed 100. Otherwise, we recommend that you use batch tags, such as batch: b1, to group the instances into batches of up to 100 instances.
	//
	// 	- The tag key can be up to 64 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the instance.
	//
	// Take note of the following items:
	//
	// 	- The tag value can be an empty string.
	//
	// 	- The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunCommandShrinkRequestResourceTag) String() string {
	return dara.Prettify(s)
}

func (s RunCommandShrinkRequestResourceTag) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequestResourceTag) GetKey() *string {
	return s.Key
}

func (s *RunCommandShrinkRequestResourceTag) GetValue() *string {
	return s.Value
}

func (s *RunCommandShrinkRequestResourceTag) SetKey(v string) *RunCommandShrinkRequestResourceTag {
	s.Key = &v
	return s
}

func (s *RunCommandShrinkRequestResourceTag) SetValue(v string) *RunCommandShrinkRequestResourceTag {
	s.Value = &v
	return s
}

func (s *RunCommandShrinkRequestResourceTag) Validate() error {
	return dara.Validate(s)
}

type RunCommandShrinkRequestTag struct {
	// The key of tag N to add to the command task. The tag key cannot be an empty string.
	//
	// If a tag is specified to query resources, up to 1,000 resources that have this tag can be displayed in the response. If multiple tags are specified to query resources, up to 1,000 resources that have all the tags can be displayed in the response. To query more than 1,000 resources that have the specified tags, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the command task. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunCommandShrinkRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RunCommandShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *RunCommandShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *RunCommandShrinkRequestTag) SetKey(v string) *RunCommandShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *RunCommandShrinkRequestTag) SetValue(v string) *RunCommandShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *RunCommandShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}
