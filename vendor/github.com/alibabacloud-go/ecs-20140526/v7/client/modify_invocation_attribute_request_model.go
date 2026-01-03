// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInvocationAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandContent(v string) *ModifyInvocationAttributeRequest
	GetCommandContent() *string
	SetContentEncoding(v string) *ModifyInvocationAttributeRequest
	GetContentEncoding() *string
	SetEnableParameter(v bool) *ModifyInvocationAttributeRequest
	GetEnableParameter() *bool
	SetFrequency(v string) *ModifyInvocationAttributeRequest
	GetFrequency() *string
	SetInstanceId(v []*string) *ModifyInvocationAttributeRequest
	GetInstanceId() []*string
	SetInvokeId(v string) *ModifyInvocationAttributeRequest
	GetInvokeId() *string
	SetOwnerAccount(v string) *ModifyInvocationAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInvocationAttributeRequest
	GetOwnerId() *int64
	SetParameters(v map[string]interface{}) *ModifyInvocationAttributeRequest
	GetParameters() map[string]interface{}
	SetRegionId(v string) *ModifyInvocationAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInvocationAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInvocationAttributeRequest
	GetResourceOwnerId() *int64
}

type ModifyInvocationAttributeRequest struct {
	// The content of the command after modification. The command content can be plaintext or Base64-encoded. Take note of the following items:
	//
	// 	- You can specify whether to retain the command after the command is run when you created the command. If you specified to retain the command, the Base64-encoded command content cannot exceed 18 KB in size. If you specified not to retain the command, the Base64-encoded command content cannot exceed 24 KB in size.
	//
	// 	- If the command content is Base64-encoded, set `ContentEncoding` to Base64.
	//
	// 	- If you set `EnableParameter` to true, the custom parameter feature is enabled and you can configure custom parameters based on the following rules:
	//
	//     	- You can define custom parameters in the `{{}}` format. Within `{{}}`, the spaces and line feeds before and after the parameter names are ignored.
	//
	//     	- The number of custom parameters cannot exceed 20.
	//
	//     	- A custom parameter name can contain letters, digits, underscores (_), and hyphens (-). The name is case-insensitive. The ACS:: prefix cannot be used to specify non-built-in environment parameters.
	//
	//     	- Each custom parameter name cannot exceed 64 bytes in length.
	//
	// 	- You can specify built-in environment parameters as custom parameters. Then, when you run the command, these parameters are automatically specified by Cloud Assistant. You can specify the following built-in environment parameters:
	//
	//     	- `{{ACS::RegionId}}`: the region ID.
	//
	//     	- `{{ACS::AccountId}}`: the UID of the Alibaba Cloud account.
	//
	//     	- `{{ACS::InstanceId}}`: the instance ID. If you want to specify `{{ACS::InstanceId}}` as a built-in environment variable, make sure that the Cloud Assistant Agent version is not earlier than the following ones:
	//
	//         	- Linux: 2.2.3.309
	//
	//         	- Windows: 2.1.3.309
	//
	//     	- `{{ACS::InstanceName}}`: the instance name. When the command is run on multiple instances, if you want to specify `{{ACS::InstanceName}}` as a built-in environment variable, make sure that the Cloud Assistant Agent version is not earlier than the following ones:
	//
	//         	- Linux: 2.2.3.344
	//
	//         	- Windows: 2.1.3.344
	//
	//     	- `{{ACS::InvokeId}}`: the ID of the task. If you want to specify `{{ACS::InvokeId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following ones:
	//
	//         	- Linux: 2.2.3.309
	//
	//         	- Windows: 2.1.3.309
	//
	//     	- `{{ACS::CommandId}}`: the command ID. If you want to specify `{{ACS::CommandId}}` as a built-in environment parameter, make sure that the Cloud Assistant Agent version is not earlier than the following ones:
	//
	//         	- Linux: 2.2.3.309
	//
	//         	- Windows: 2.1.3.309
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The encoding mode of the command content that is specified by `CommandContent`. Valid values (case-insensitive):
	//
	// 	- PlainText: The command content is not encoded.
	//
	// 	- Base64: The command content is encoded in Base64.
	//
	// Default value: PlainText. If the value is invalid, the PlainText mode is used.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Specifies whether to include custom parameters in the command.
	//
	// 	- If you want to enable the custom parameter feature, or configure `Parameters` to modify the custom parameters in the command, set EnableParameter to `true`.
	//
	// 	- If you do not want to configure `Parameters` to modify the custom parameters in the command, leave EnableParameter empty or set it to `false`.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// The frequency at which to run the command. This parameter takes effect only when `RepeatMode` is set to `Period`. You can configure a command to run at a fixed interval based on a rate expression, run only once at a specific point in time, or run at designated points in time based on a cron expression.
	//
	// 	- To run the command at a fixed interval, use a rate expression to specify the interval. You can specify the interval in seconds, minutes, hours, or days. This option is suitable for scenarios in which tasks need to be executed at a fixed interval. Specify the interval in the following format: `rate(<Execution interval value> <Execution interval unit>)`. For example, specify `rate(5m)` to run the command every 5 minutes. Take note of the following limits when you specify an interval:
	//
	//     	- The specified interval must be in the range of 60 seconds to 7 days and must be longer than the timeout period specified when you created the scheduled task.
	//
	//     	- The interval is the amount of time that elapses between two consecutive executions. The interval is irrelevant to the amount of time that is required to run the command once. For example, you set the interval to 5 minutes and the command requires 2 minutes to run once. Each time the command running is complete, the system waits 3 minutes instead of 5 minutes before the system runs the command again.
	//
	//     	- The point in time at which the command is run the next time is calculated based on the creation time of the task (the [CreationTime](https://help.aliyun.com/document_detail/64840.html) value returned by the `DescribeInvocations` operation) and the modified execution interval.
	//
	// 	- To run a command only once at a specific time, specify a point in time and a time zone. Specify the point in time in the `at(yyyy-MM-dd HH:mm:ss <Time zone>)` format, which indicates `at(Year-Month-Day Hours:Minutes:Seconds <Time zone>)`. If you do not specify a time zone, the Coordinated Universal Time (UTC) time zone is used by default. You can specify a time zone in the following forms:
	//
	//     	- The time zone name. Examples: `Asia/Shanghai` and `America/Los_Angeles`.
	//
	//     	- The time offset from GMT. Examples: `GMT+8:00` (UTC+8) and `GMT-7:00` (UTC-7). If you use the GMT format, you cannot add leading zeros to the hour value.
	//
	//     	- The time zone abbreviation. Only UTC is supported.
	//
	//     For example, to configure a command to run only once at 13:15:30 on June 6, 2022 (Shanghai time), set the time to `at(2022-06-06 13:15:30 Asia/Shanghai)`. To configure a command to run only once at 13:15:30 on June 6, 2022 (UTC-7), set the time to `at(2022-06-06 13:15:30 GMT-7:00)`.
	//
	// 	- To run a command at designated points in time, use a cron expression to define the schedule. Specify a schedule in the `<Cron expression> <Time zone>` format. The cron expression is in the `<Seconds> <Minutes> <Hours> <Day of the month> <Month> <Day of the week> <Year (optional)>` format. The system calculates the execution times of the command based on the specified cron expression and time zone and runs the command as scheduled. If you do not specify a time zone, the system time zones of the instances on which you want to run the command are used by default. For information about cron expressions, see [Cron expressions](https://help.aliyun.com/document_detail/64769.html). You can specify the time zone in the following forms:
	//
	//     	- The time zone name. Examples: `Asia/Shanghai` and `America/Los_Angeles`.
	//
	//     	- The time offset from GMT. Examples: `GMT+8:00` (UTC+8) and `GMT-7:00` (UTC-7). If you use the GMT format, you cannot add leading zeros to the hour value.
	//
	//     	- The time zone abbreviation. Only UTC is supported. For example, to configure a command to run at 10:15:00 every day in 2022 (Shanghai time), set the schedule to `0 15 10 ? 	- 	- 2022 Asia/Shanghai`. To configure a command to run every half an hour from 10:00:00 to 11:30:00 every day in 2022 (UTC+8), set the schedule to `0 0/30 10-11 	- 	- ? 2022 GMT+8:00`. To configure a command to run every 5 minutes from 14:00:00 to 14:55:00 every October every two years from 2022 in UTC, set the schedule to `0 0/5 14 	- 10 ? 2022/2 UTC`.
	//
	//     **
	//
	//     **Note*	- The minimum interval must be 10 seconds or longer and cannot be shorter than the timeout period of scheduled executions.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The IDs of the ECS instances or managed instances that you want to add to the scheduled command task.
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The execution ID of the command.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId     *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The key-value pairs of the custom parameters that are passed in if custom parameters are included in the command.
	//
	// You can specify 0 to 10 custom parameters. Take note of the following items:
	//
	// 	- The key of a custom parameter can be up to 64 characters in length and cannot be an empty string.
	//
	// 	- The value of a custom parameter can be an empty string.
	//
	// 	- If you specified to retain the command when you create the command task, the total size of the custom parameters and original command content that are encoded in Base64 cannot exceed 18 KB. If you specified not to retain the command when you create the command task, the total size of the custom parameters and original command content that are encoded in Base64 cannot exceed 24 KB.
	//
	// 	- The custom parameter names that are specified by Parameters must be included in the custom parameter names that you specified when you created the command. You can use empty strings to represent the parameters that are not passed in.
	//
	// This parameter is empty by default, which indicates not to modify the key-value pairs of the custom parameters.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAI*************"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyInvocationAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInvocationAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInvocationAttributeRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *ModifyInvocationAttributeRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *ModifyInvocationAttributeRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *ModifyInvocationAttributeRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *ModifyInvocationAttributeRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *ModifyInvocationAttributeRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *ModifyInvocationAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInvocationAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInvocationAttributeRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *ModifyInvocationAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInvocationAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInvocationAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInvocationAttributeRequest) SetCommandContent(v string) *ModifyInvocationAttributeRequest {
	s.CommandContent = &v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetContentEncoding(v string) *ModifyInvocationAttributeRequest {
	s.ContentEncoding = &v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetEnableParameter(v bool) *ModifyInvocationAttributeRequest {
	s.EnableParameter = &v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetFrequency(v string) *ModifyInvocationAttributeRequest {
	s.Frequency = &v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetInstanceId(v []*string) *ModifyInvocationAttributeRequest {
	s.InstanceId = v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetInvokeId(v string) *ModifyInvocationAttributeRequest {
	s.InvokeId = &v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetOwnerAccount(v string) *ModifyInvocationAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetOwnerId(v int64) *ModifyInvocationAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetParameters(v map[string]interface{}) *ModifyInvocationAttributeRequest {
	s.Parameters = v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetRegionId(v string) *ModifyInvocationAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetResourceOwnerAccount(v string) *ModifyInvocationAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInvocationAttributeRequest) SetResourceOwnerId(v int64) *ModifyInvocationAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInvocationAttributeRequest) Validate() error {
	return dara.Validate(s)
}
