// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTerminalSessionShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandLine(v string) *StartTerminalSessionShrinkRequest
	GetCommandLine() *string
	SetConnectionType(v string) *StartTerminalSessionShrinkRequest
	GetConnectionType() *string
	SetEncryptionOptionsShrink(v string) *StartTerminalSessionShrinkRequest
	GetEncryptionOptionsShrink() *string
	SetInstanceId(v []*string) *StartTerminalSessionShrinkRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *StartTerminalSessionShrinkRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartTerminalSessionShrinkRequest
	GetOwnerId() *int64
	SetPasswordName(v string) *StartTerminalSessionShrinkRequest
	GetPasswordName() *string
	SetPortNumber(v int32) *StartTerminalSessionShrinkRequest
	GetPortNumber() *int32
	SetRegionId(v string) *StartTerminalSessionShrinkRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartTerminalSessionShrinkRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartTerminalSessionShrinkRequest
	GetResourceOwnerId() *int64
	SetTargetServer(v string) *StartTerminalSessionShrinkRequest
	GetTargetServer() *string
	SetUsername(v string) *StartTerminalSessionShrinkRequest
	GetUsername() *string
}

type StartTerminalSessionShrinkRequest struct {
	// The command to run after the session is initiated. The command length cannot exceed 512 characters.
	//
	// >  If you specify the `CommandLine` parameter, you cannot specify the `PortNumber` or `TargetServer` parameter.
	//
	// example:
	//
	// ssh root@192.168.0.246
	CommandLine *string `json:"CommandLine,omitempty" xml:"CommandLine,omitempty"`
	// The network type of the WebSocket URL required to connect to the instance. Valid values:
	//
	// 	- Internet (default)
	//
	// 	- Intranet
	//
	// example:
	//
	// Intranet
	ConnectionType          *string `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	EncryptionOptionsShrink *string `json:"EncryptionOptions,omitempty" xml:"EncryptionOptions,omitempty"`
	// The instance IDs.
	//
	// This parameter is required.
	InstanceId   []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	OwnerAccount *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PasswordName *string   `json:"PasswordName,omitempty" xml:"PasswordName,omitempty"`
	// The port number of the ECS instance. The port is used to forward data. After this parameter is configured, Cloud Assistant Agent forwards data to the specified port. For example, you can set this parameter to 22 for data forwarding over SSH.
	//
	// This parameter is empty by default, which indicates that no port is configured to forward data.
	//
	// example:
	//
	// 22
	PortNumber *int32 `json:"PortNumber,omitempty" xml:"PortNumber,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IP address of the instance. You can use the IP address to access the destination service in a virtual private cloud (VPC).
	//
	// >  If this parameter is not empty, `PortNumber` specifies the port number that is used by the managed instance to access the destination service in the VPC.
	//
	// example:
	//
	// 192.168.0.246
	TargetServer *string `json:"TargetServer,omitempty" xml:"TargetServer,omitempty"`
	// The username used for connection establishment.
	//
	// example:
	//
	// testUser
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s StartTerminalSessionShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTerminalSessionShrinkRequest) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionShrinkRequest) GetCommandLine() *string {
	return s.CommandLine
}

func (s *StartTerminalSessionShrinkRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *StartTerminalSessionShrinkRequest) GetEncryptionOptionsShrink() *string {
	return s.EncryptionOptionsShrink
}

func (s *StartTerminalSessionShrinkRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *StartTerminalSessionShrinkRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartTerminalSessionShrinkRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartTerminalSessionShrinkRequest) GetPasswordName() *string {
	return s.PasswordName
}

func (s *StartTerminalSessionShrinkRequest) GetPortNumber() *int32 {
	return s.PortNumber
}

func (s *StartTerminalSessionShrinkRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartTerminalSessionShrinkRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartTerminalSessionShrinkRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartTerminalSessionShrinkRequest) GetTargetServer() *string {
	return s.TargetServer
}

func (s *StartTerminalSessionShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *StartTerminalSessionShrinkRequest) SetCommandLine(v string) *StartTerminalSessionShrinkRequest {
	s.CommandLine = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetConnectionType(v string) *StartTerminalSessionShrinkRequest {
	s.ConnectionType = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetEncryptionOptionsShrink(v string) *StartTerminalSessionShrinkRequest {
	s.EncryptionOptionsShrink = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetInstanceId(v []*string) *StartTerminalSessionShrinkRequest {
	s.InstanceId = v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetOwnerAccount(v string) *StartTerminalSessionShrinkRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetOwnerId(v int64) *StartTerminalSessionShrinkRequest {
	s.OwnerId = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetPasswordName(v string) *StartTerminalSessionShrinkRequest {
	s.PasswordName = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetPortNumber(v int32) *StartTerminalSessionShrinkRequest {
	s.PortNumber = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetRegionId(v string) *StartTerminalSessionShrinkRequest {
	s.RegionId = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetResourceOwnerAccount(v string) *StartTerminalSessionShrinkRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetResourceOwnerId(v int64) *StartTerminalSessionShrinkRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetTargetServer(v string) *StartTerminalSessionShrinkRequest {
	s.TargetServer = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) SetUsername(v string) *StartTerminalSessionShrinkRequest {
	s.Username = &v
	return s
}

func (s *StartTerminalSessionShrinkRequest) Validate() error {
	return dara.Validate(s)
}
