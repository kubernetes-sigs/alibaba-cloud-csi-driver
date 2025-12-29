// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTerminalSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandLine(v string) *StartTerminalSessionRequest
	GetCommandLine() *string
	SetConnectionType(v string) *StartTerminalSessionRequest
	GetConnectionType() *string
	SetEncryptionOptions(v *StartTerminalSessionRequestEncryptionOptions) *StartTerminalSessionRequest
	GetEncryptionOptions() *StartTerminalSessionRequestEncryptionOptions
	SetInstanceId(v []*string) *StartTerminalSessionRequest
	GetInstanceId() []*string
	SetOwnerAccount(v string) *StartTerminalSessionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartTerminalSessionRequest
	GetOwnerId() *int64
	SetPasswordName(v string) *StartTerminalSessionRequest
	GetPasswordName() *string
	SetPortNumber(v int32) *StartTerminalSessionRequest
	GetPortNumber() *int32
	SetRegionId(v string) *StartTerminalSessionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartTerminalSessionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartTerminalSessionRequest
	GetResourceOwnerId() *int64
	SetTargetServer(v string) *StartTerminalSessionRequest
	GetTargetServer() *string
	SetUsername(v string) *StartTerminalSessionRequest
	GetUsername() *string
}

type StartTerminalSessionRequest struct {
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
	ConnectionType    *string                                       `json:"ConnectionType,omitempty" xml:"ConnectionType,omitempty"`
	EncryptionOptions *StartTerminalSessionRequestEncryptionOptions `json:"EncryptionOptions,omitempty" xml:"EncryptionOptions,omitempty" type:"Struct"`
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

func (s StartTerminalSessionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartTerminalSessionRequest) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionRequest) GetCommandLine() *string {
	return s.CommandLine
}

func (s *StartTerminalSessionRequest) GetConnectionType() *string {
	return s.ConnectionType
}

func (s *StartTerminalSessionRequest) GetEncryptionOptions() *StartTerminalSessionRequestEncryptionOptions {
	return s.EncryptionOptions
}

func (s *StartTerminalSessionRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *StartTerminalSessionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartTerminalSessionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartTerminalSessionRequest) GetPasswordName() *string {
	return s.PasswordName
}

func (s *StartTerminalSessionRequest) GetPortNumber() *int32 {
	return s.PortNumber
}

func (s *StartTerminalSessionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartTerminalSessionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartTerminalSessionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartTerminalSessionRequest) GetTargetServer() *string {
	return s.TargetServer
}

func (s *StartTerminalSessionRequest) GetUsername() *string {
	return s.Username
}

func (s *StartTerminalSessionRequest) SetCommandLine(v string) *StartTerminalSessionRequest {
	s.CommandLine = &v
	return s
}

func (s *StartTerminalSessionRequest) SetConnectionType(v string) *StartTerminalSessionRequest {
	s.ConnectionType = &v
	return s
}

func (s *StartTerminalSessionRequest) SetEncryptionOptions(v *StartTerminalSessionRequestEncryptionOptions) *StartTerminalSessionRequest {
	s.EncryptionOptions = v
	return s
}

func (s *StartTerminalSessionRequest) SetInstanceId(v []*string) *StartTerminalSessionRequest {
	s.InstanceId = v
	return s
}

func (s *StartTerminalSessionRequest) SetOwnerAccount(v string) *StartTerminalSessionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartTerminalSessionRequest) SetOwnerId(v int64) *StartTerminalSessionRequest {
	s.OwnerId = &v
	return s
}

func (s *StartTerminalSessionRequest) SetPasswordName(v string) *StartTerminalSessionRequest {
	s.PasswordName = &v
	return s
}

func (s *StartTerminalSessionRequest) SetPortNumber(v int32) *StartTerminalSessionRequest {
	s.PortNumber = &v
	return s
}

func (s *StartTerminalSessionRequest) SetRegionId(v string) *StartTerminalSessionRequest {
	s.RegionId = &v
	return s
}

func (s *StartTerminalSessionRequest) SetResourceOwnerAccount(v string) *StartTerminalSessionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartTerminalSessionRequest) SetResourceOwnerId(v int64) *StartTerminalSessionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartTerminalSessionRequest) SetTargetServer(v string) *StartTerminalSessionRequest {
	s.TargetServer = &v
	return s
}

func (s *StartTerminalSessionRequest) SetUsername(v string) *StartTerminalSessionRequest {
	s.Username = &v
	return s
}

func (s *StartTerminalSessionRequest) Validate() error {
	if s.EncryptionOptions != nil {
		if err := s.EncryptionOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartTerminalSessionRequestEncryptionOptions struct {
	Enabled  *bool   `json:"Enabled,omitempty" xml:"Enabled,omitempty"`
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	Mode     *string `json:"Mode,omitempty" xml:"Mode,omitempty"`
}

func (s StartTerminalSessionRequestEncryptionOptions) String() string {
	return dara.Prettify(s)
}

func (s StartTerminalSessionRequestEncryptionOptions) GoString() string {
	return s.String()
}

func (s *StartTerminalSessionRequestEncryptionOptions) GetEnabled() *bool {
	return s.Enabled
}

func (s *StartTerminalSessionRequestEncryptionOptions) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *StartTerminalSessionRequestEncryptionOptions) GetMode() *string {
	return s.Mode
}

func (s *StartTerminalSessionRequestEncryptionOptions) SetEnabled(v bool) *StartTerminalSessionRequestEncryptionOptions {
	s.Enabled = &v
	return s
}

func (s *StartTerminalSessionRequestEncryptionOptions) SetKMSKeyId(v string) *StartTerminalSessionRequestEncryptionOptions {
	s.KMSKeyId = &v
	return s
}

func (s *StartTerminalSessionRequestEncryptionOptions) SetMode(v string) *StartTerminalSessionRequestEncryptionOptions {
	s.Mode = &v
	return s
}

func (s *StartTerminalSessionRequestEncryptionOptions) Validate() error {
	return dara.Validate(s)
}
