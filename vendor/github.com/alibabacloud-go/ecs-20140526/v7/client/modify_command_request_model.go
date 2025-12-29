// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandContent(v string) *ModifyCommandRequest
	GetCommandContent() *string
	SetCommandId(v string) *ModifyCommandRequest
	GetCommandId() *string
	SetDescription(v string) *ModifyCommandRequest
	GetDescription() *string
	SetLauncher(v string) *ModifyCommandRequest
	GetLauncher() *string
	SetName(v string) *ModifyCommandRequest
	GetName() *string
	SetOwnerAccount(v string) *ModifyCommandRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyCommandRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyCommandRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyCommandRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyCommandRequest
	GetResourceOwnerId() *int64
	SetTimeout(v int64) *ModifyCommandRequest
	GetTimeout() *int64
	SetWorkingDir(v string) *ModifyCommandRequest
	GetWorkingDir() *string
}

type ModifyCommandRequest struct {
	// >  This parameter is no longer used and does not take effect.
	//
	// example:
	//
	// echo
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// The command ID. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) operation to query all available command IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-hz01272yr52****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// The command description. The description supports all character sets and can be up to 512 characters in length.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The launcher for script execution. The value cannot exceed 1 KB in length.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// The command name. The name supports all character sets and can be up to 128 characters in length.
	//
	// example:
	//
	// test-CommandName
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the command. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The maximum timeout period for the command to be run on the instance. Unit: seconds. When a command cannot run within the specified time range, the command times out. Then, the command process is forcibly terminated by canceling the process ID (PID) of the command.
	//
	// example:
	//
	// 120
	Timeout *int64 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The working directory of the command. The value can be up to 200 characters in length.
	//
	// example:
	//
	// /home/
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s ModifyCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyCommandRequest) GoString() string {
	return s.String()
}

func (s *ModifyCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *ModifyCommandRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *ModifyCommandRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyCommandRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *ModifyCommandRequest) GetName() *string {
	return s.Name
}

func (s *ModifyCommandRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyCommandRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyCommandRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyCommandRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyCommandRequest) GetTimeout() *int64 {
	return s.Timeout
}

func (s *ModifyCommandRequest) GetWorkingDir() *string {
	return s.WorkingDir
}

func (s *ModifyCommandRequest) SetCommandContent(v string) *ModifyCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *ModifyCommandRequest) SetCommandId(v string) *ModifyCommandRequest {
	s.CommandId = &v
	return s
}

func (s *ModifyCommandRequest) SetDescription(v string) *ModifyCommandRequest {
	s.Description = &v
	return s
}

func (s *ModifyCommandRequest) SetLauncher(v string) *ModifyCommandRequest {
	s.Launcher = &v
	return s
}

func (s *ModifyCommandRequest) SetName(v string) *ModifyCommandRequest {
	s.Name = &v
	return s
}

func (s *ModifyCommandRequest) SetOwnerAccount(v string) *ModifyCommandRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyCommandRequest) SetOwnerId(v int64) *ModifyCommandRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyCommandRequest) SetRegionId(v string) *ModifyCommandRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyCommandRequest) SetResourceOwnerAccount(v string) *ModifyCommandRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyCommandRequest) SetResourceOwnerId(v int64) *ModifyCommandRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyCommandRequest) SetTimeout(v int64) *ModifyCommandRequest {
	s.Timeout = &v
	return s
}

func (s *ModifyCommandRequest) SetWorkingDir(v string) *ModifyCommandRequest {
	s.WorkingDir = &v
	return s
}

func (s *ModifyCommandRequest) Validate() error {
	return dara.Validate(s)
}
