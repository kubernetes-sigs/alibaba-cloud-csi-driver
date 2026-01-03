// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuOptions(v *ModifyInstanceAttributeRequestCpuOptions) *ModifyInstanceAttributeRequest
	GetCpuOptions() *ModifyInstanceAttributeRequestCpuOptions
	SetCreditSpecification(v string) *ModifyInstanceAttributeRequest
	GetCreditSpecification() *string
	SetDeletionProtection(v bool) *ModifyInstanceAttributeRequest
	GetDeletionProtection() *bool
	SetDescription(v string) *ModifyInstanceAttributeRequest
	GetDescription() *string
	SetEnableJumboFrame(v bool) *ModifyInstanceAttributeRequest
	GetEnableJumboFrame() *bool
	SetEnableNetworkEncryption(v bool) *ModifyInstanceAttributeRequest
	GetEnableNetworkEncryption() *bool
	SetHostName(v string) *ModifyInstanceAttributeRequest
	GetHostName() *string
	SetInstanceId(v string) *ModifyInstanceAttributeRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyInstanceAttributeRequest
	GetInstanceName() *string
	SetNetworkInterfaceQueueNumber(v int32) *ModifyInstanceAttributeRequest
	GetNetworkInterfaceQueueNumber() *int32
	SetOwnerAccount(v string) *ModifyInstanceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceAttributeRequest
	GetOwnerId() *int64
	SetPassword(v string) *ModifyInstanceAttributeRequest
	GetPassword() *string
	SetPrivateDnsNameOptions(v *ModifyInstanceAttributeRequestPrivateDnsNameOptions) *ModifyInstanceAttributeRequest
	GetPrivateDnsNameOptions() *ModifyInstanceAttributeRequestPrivateDnsNameOptions
	SetRecyclable(v bool) *ModifyInstanceAttributeRequest
	GetRecyclable() *bool
	SetRemoteConnectionOptions(v *ModifyInstanceAttributeRequestRemoteConnectionOptions) *ModifyInstanceAttributeRequest
	GetRemoteConnectionOptions() *ModifyInstanceAttributeRequestRemoteConnectionOptions
	SetResourceOwnerAccount(v string) *ModifyInstanceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceAttributeRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupIds(v []*string) *ModifyInstanceAttributeRequest
	GetSecurityGroupIds() []*string
	SetUserData(v string) *ModifyInstanceAttributeRequest
	GetUserData() *string
}

type ModifyInstanceAttributeRequest struct {
	CpuOptions *ModifyInstanceAttributeRequestCpuOptions `json:"CpuOptions,omitempty" xml:"CpuOptions,omitempty" type:"Struct"`
	// The performance mode of the burstable instance. Valid values:
	//
	// 	- Standard
	//
	// 	- Unlimited
	//
	// For more information about the performance modes of burstable instances, see [Overview](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The release protection attribute of the instance. This parameter specifies whether you can use the ECS console or call the [DeleteInstance](https://help.aliyun.com/document_detail/25507.html) operation to release the instance.
	//
	// >  This parameter is applicable only to pay-as-you-go instances. The release protection attribute can protect instances against manual releases, but not against automatic releases.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The description of the instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testInstanceDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the Jumbo Frames feature for the instance. Valid values:
	//
	// 	- true: The Jumbo Frame feature is enabled for the instance.
	//
	// 	- false: The Jumbo Frame feature is disabled for the instance.
	//
	// Take note of the following items:
	//
	// 	- The instance must be in the Running (`Running`) or Stopped (`Stopped`) state.
	//
	// 	- The instance must reside in a VPC.
	//
	// 	- After the Jumbo Frames feature is enabled, the MTU value of the instance is set to 8500. After the Jumbo Frames feature is disabled, the MTU value of the instance is set to 1500. You can enable the Jumbo Frames feature only for specific instance types. For more information, see [Jumbo Frames](https://help.aliyun.com/document_detail/200512.html).
	//
	// example:
	//
	// false
	EnableJumboFrame        *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	EnableNetworkEncryption *bool `json:"EnableNetworkEncryption,omitempty" xml:"EnableNetworkEncryption,omitempty"`
	// The hostname of the instance. Take note of the following items:
	//
	// 	- The instance cannot be in the Creating (`Pending`) or Starting (`Starting`) state. Otherwise, the new hostname and the configurations in the `/etc/hosts` file may not take effect. You can call the [DescribeInstances](https://help.aliyun.com/document_detail/25506.html) operation to query the status of the instance.
	//
	// 	- The parameter takes effect after the instance is restarted. You can restart an instance in the ECS console. For more information, see [Restart an instance](https://help.aliyun.com/document_detail/25440.html). You can also call the [RebootInstance](https://help.aliyun.com/document_detail/25502.html) operation to restart the instance. The parameter cannot take effect if you restart an instance within the operating system.
	//
	// The following limits apply to the hostnames of instances that run different operating systems:
	//
	// 	- For Windows Server, the hostname must be 2 to 15 characters in length and can contain letters, digits, and hyphens (-). The hostname cannot start or end with a hyphen (-), contain consecutive hyphens (-), or contain only digits.
	//
	// 	- For other operating systems such as Linux, the hostname must be 2 to 64 characters in length. You can use periods (.) to separate a hostname into multiple segments. Each segment can contain letters, digits, and hyphens (-). The hostname cannot contain consecutive periods (.) or hyphens (-). The hostname cannot start or end with a period (.) or a hyphen (-).
	//
	// example:
	//
	// testHostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The number of queues supported by the primary elastic network interface (ENI) of the instance. Take note of the following items:
	//
	// 	- The instance must be in the Stopped (`Stopped`) state.
	//
	// 	- The number of queues supported by an ENI cannot exceed the maximum number of queues that the instance type allows for each ENI. The total number of queues on all ENIs on the instance cannot exceed the queue quota that the instance type supports. To query the maximum number of queues that an instance type allows for each ENI and the queue quota for the instance type, call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation.
	//
	// 	- If you set this parameter to -1, the value is reset to the default value for the instance type. To query the default number of queues of an ENI of each instance type, call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation.
	//
	// example:
	//
	// 8
	NetworkInterfaceQueueNumber *int32  `json:"NetworkInterfaceQueueNumber,omitempty" xml:"NetworkInterfaceQueueNumber,omitempty"`
	OwnerAccount                *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId                     *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the instance. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include **( ) \\` ~ ! @ # $ % ^ & \\	- - _ + = | { } [ ] : ; \\" < > , . ? /*	- The password of a Windows instance cannot start with a forward slash (/). Take note of the following items:
	//
	// 	- The instance cannot be in the Starting (`Starting`) state.
	//
	// 	- The parameter takes effect after the instance is restarted. You can restart an instance in the ECS console. For more information, see [Restart an instance](https://help.aliyun.com/document_detail/25440.html). You can also call the [RebootInstance](https://help.aliyun.com/document_detail/25502.html) operation to restart the instance. The parameter cannot take effect if you restart an instance within the operating system.
	//
	// >  For security reasons, we recommend that you use HTTPS to send requests if `Password` is specified.
	//
	// example:
	//
	// Test123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The private domain name options of the ECS instance.
	//
	// For information about private domain name resolution, see [ECS private DNS resolution](https://help.aliyun.com/document_detail/2844797.html).
	PrivateDnsNameOptions *ModifyInstanceAttributeRequestPrivateDnsNameOptions `json:"PrivateDnsNameOptions,omitempty" xml:"PrivateDnsNameOptions,omitempty" type:"Struct"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// hide
	Recyclable *bool `json:"Recyclable,omitempty" xml:"Recyclable,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	RemoteConnectionOptions *ModifyInstanceAttributeRequestRemoteConnectionOptions `json:"RemoteConnectionOptions,omitempty" xml:"RemoteConnectionOptions,omitempty" type:"Struct"`
	ResourceOwnerAccount    *string                                                `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId         *int64                                                 `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The IDs of the new security groups to which to assign the instance. Take note of the following items:
	//
	// 	- The security group IDs in the array cannot be duplicate. The length of the array is related to the quota of security groups to which the instance can be assigned. For more information, see the [Security group limits](~~25412#SecurityGroupQuota1~~) section in the "Limits and quotas" topic.
	//
	// 	- The instance is moved from the current security groups to the replacement security groups. If you want the instance to remain in the current security groups, add the IDs of the current security groups to the array.
	//
	// 	- You can move the instance to security groups of a different type. However, the array cannot contain the IDs of both basic and advanced security groups.
	//
	// 	- The security groups and the instance must belong to the same VPC.
	//
	// 	- Security groups of instances in the classic network cannot be changed.
	//
	// >  New security groups become valid for the instance after a short delay.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7o****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The user data of the instance. We recommend that you encode the data in Base64. Take note of the following items:
	//
	// 	- The instance must meet the limits for user data. For more information, see [Initialize an instance by using instance user data](https://help.aliyun.com/document_detail/49121.html).
	//
	// 	- After you restart the instance, the new user data is displayed but not run as scripts.
	//
	// >  The maximum size of the raw data before encoding is 32 KB. We recommend that you do not pass in confidential information such as passwords and private keys in plaintext. If you must pass in confidential information, we recommend that you encrypt and Base64-encode the information before you pass it in. Then, you can decode and decrypt the information in the same way within the instance.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ModifyInstanceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequest) GetCpuOptions() *ModifyInstanceAttributeRequestCpuOptions {
	return s.CpuOptions
}

func (s *ModifyInstanceAttributeRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *ModifyInstanceAttributeRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *ModifyInstanceAttributeRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifyInstanceAttributeRequest) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *ModifyInstanceAttributeRequest) GetEnableNetworkEncryption() *bool {
	return s.EnableNetworkEncryption
}

func (s *ModifyInstanceAttributeRequest) GetHostName() *string {
	return s.HostName
}

func (s *ModifyInstanceAttributeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAttributeRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyInstanceAttributeRequest) GetNetworkInterfaceQueueNumber() *int32 {
	return s.NetworkInterfaceQueueNumber
}

func (s *ModifyInstanceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceAttributeRequest) GetPassword() *string {
	return s.Password
}

func (s *ModifyInstanceAttributeRequest) GetPrivateDnsNameOptions() *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	return s.PrivateDnsNameOptions
}

func (s *ModifyInstanceAttributeRequest) GetRecyclable() *bool {
	return s.Recyclable
}

func (s *ModifyInstanceAttributeRequest) GetRemoteConnectionOptions() *ModifyInstanceAttributeRequestRemoteConnectionOptions {
	return s.RemoteConnectionOptions
}

func (s *ModifyInstanceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceAttributeRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *ModifyInstanceAttributeRequest) GetUserData() *string {
	return s.UserData
}

func (s *ModifyInstanceAttributeRequest) SetCpuOptions(v *ModifyInstanceAttributeRequestCpuOptions) *ModifyInstanceAttributeRequest {
	s.CpuOptions = v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetCreditSpecification(v string) *ModifyInstanceAttributeRequest {
	s.CreditSpecification = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetDeletionProtection(v bool) *ModifyInstanceAttributeRequest {
	s.DeletionProtection = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetDescription(v string) *ModifyInstanceAttributeRequest {
	s.Description = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetEnableJumboFrame(v bool) *ModifyInstanceAttributeRequest {
	s.EnableJumboFrame = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetEnableNetworkEncryption(v bool) *ModifyInstanceAttributeRequest {
	s.EnableNetworkEncryption = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetHostName(v string) *ModifyInstanceAttributeRequest {
	s.HostName = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceId(v string) *ModifyInstanceAttributeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetInstanceName(v string) *ModifyInstanceAttributeRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetNetworkInterfaceQueueNumber(v int32) *ModifyInstanceAttributeRequest {
	s.NetworkInterfaceQueueNumber = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetOwnerAccount(v string) *ModifyInstanceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetOwnerId(v int64) *ModifyInstanceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetPassword(v string) *ModifyInstanceAttributeRequest {
	s.Password = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetPrivateDnsNameOptions(v *ModifyInstanceAttributeRequestPrivateDnsNameOptions) *ModifyInstanceAttributeRequest {
	s.PrivateDnsNameOptions = v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetRecyclable(v bool) *ModifyInstanceAttributeRequest {
	s.Recyclable = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetRemoteConnectionOptions(v *ModifyInstanceAttributeRequestRemoteConnectionOptions) *ModifyInstanceAttributeRequest {
	s.RemoteConnectionOptions = v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetResourceOwnerId(v int64) *ModifyInstanceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetSecurityGroupIds(v []*string) *ModifyInstanceAttributeRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *ModifyInstanceAttributeRequest) SetUserData(v string) *ModifyInstanceAttributeRequest {
	s.UserData = &v
	return s
}

func (s *ModifyInstanceAttributeRequest) Validate() error {
	if s.CpuOptions != nil {
		if err := s.CpuOptions.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateDnsNameOptions != nil {
		if err := s.PrivateDnsNameOptions.Validate(); err != nil {
			return err
		}
	}
	if s.RemoteConnectionOptions != nil {
		if err := s.RemoteConnectionOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyInstanceAttributeRequestCpuOptions struct {
	// The number of CPU cores. This parameter cannot be specified but only uses its default value.
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// The number of threads per CPU core. The following formula is used to calculate the number of vCPUs of the instance: `CpuOptions.Core` value × `CpuOptions.ThreadsPerCore` value.
	//
	// 	- If `CpuOptionsThreadPerCore` is set to 1, Hyper-Threading (HT) is disabled.
	//
	// 	- This parameter is applicable only to specific instance types.
	//
	// example:
	//
	// 2
	ThreadsPerCore *int32 `json:"ThreadsPerCore,omitempty" xml:"ThreadsPerCore,omitempty"`
	// The CPU topology type of the instance. Valid values:
	//
	// 	- ContinuousCoreToHTMapping: The Hyper-Threading (HT) technology allows continuous threads to run on the same core in the CPU topology of the instance.
	//
	// 	- DiscreteCoreToHTMapping: The HT technology allows discrete threads to run on the same core.
	//
	// This parameter is left empty by default.
	//
	// Take note of the following items:
	//
	// 	- The instance must be in the Stopped (`Stopped`) state.
	//
	// >  This parameter is supported only for specific instance families. For information about the supported instance families, see [View and modify CPU topologies](https://help.aliyun.com/document_detail/2636059.html).
	//
	// example:
	//
	// DiscreteCoreToHTMapping
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
}

func (s ModifyInstanceAttributeRequestCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequestCpuOptions) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequestCpuOptions) GetCore() *int32 {
	return s.Core
}

func (s *ModifyInstanceAttributeRequestCpuOptions) GetThreadsPerCore() *int32 {
	return s.ThreadsPerCore
}

func (s *ModifyInstanceAttributeRequestCpuOptions) GetTopologyType() *string {
	return s.TopologyType
}

func (s *ModifyInstanceAttributeRequestCpuOptions) SetCore(v int32) *ModifyInstanceAttributeRequestCpuOptions {
	s.Core = &v
	return s
}

func (s *ModifyInstanceAttributeRequestCpuOptions) SetThreadsPerCore(v int32) *ModifyInstanceAttributeRequestCpuOptions {
	s.ThreadsPerCore = &v
	return s
}

func (s *ModifyInstanceAttributeRequestCpuOptions) SetTopologyType(v string) *ModifyInstanceAttributeRequestCpuOptions {
	s.TopologyType = &v
	return s
}

func (s *ModifyInstanceAttributeRequestCpuOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyInstanceAttributeRequestPrivateDnsNameOptions struct {
	// Specifies whether DNS Resolution from the Instance ID-based Hostname to the Instance Primary Private IPv6 Address (AAAA Record) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableInstanceIdDnsAAAARecord *bool `json:"EnableInstanceIdDnsAAAARecord,omitempty" xml:"EnableInstanceIdDnsAAAARecord,omitempty"`
	// Specifies whether DNS Resolution from the Instance ID-based Hostname to the Instance Primary Private IPv4 Address (A Record) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableInstanceIdDnsARecord *bool `json:"EnableInstanceIdDnsARecord,omitempty" xml:"EnableInstanceIdDnsARecord,omitempty"`
	// Specifies whether DNS Resolution from the IP Address-based Hostname to the Instance Primary Private IPv4 Address (A Record) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableIpDnsARecord *bool `json:"EnableIpDnsARecord,omitempty" xml:"EnableIpDnsARecord,omitempty"`
	// Specifies whether Reverse DNS Resolution from the Instance Primary Private IPv4 Address to the IP Address-based Hostname (PTR Record) is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableIpDnsPtrRecord *bool `json:"EnableIpDnsPtrRecord,omitempty" xml:"EnableIpDnsPtrRecord,omitempty"`
	// The type of the hostname. Valid values:
	//
	// 	- Custom: custom hostname.
	//
	// 	- IpBased: IP address-based hostname.
	//
	// 	- InstanceIdBased: instance ID-based hostname.
	//
	// Default value: Custom.
	//
	// example:
	//
	// Custom
	HostnameType *string `json:"HostnameType,omitempty" xml:"HostnameType,omitempty"`
}

func (s ModifyInstanceAttributeRequestPrivateDnsNameOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequestPrivateDnsNameOptions) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) GetEnableInstanceIdDnsAAAARecord() *bool {
	return s.EnableInstanceIdDnsAAAARecord
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) GetEnableInstanceIdDnsARecord() *bool {
	return s.EnableInstanceIdDnsARecord
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) GetEnableIpDnsARecord() *bool {
	return s.EnableIpDnsARecord
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) GetEnableIpDnsPtrRecord() *bool {
	return s.EnableIpDnsPtrRecord
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) GetHostnameType() *string {
	return s.HostnameType
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) SetEnableInstanceIdDnsAAAARecord(v bool) *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	s.EnableInstanceIdDnsAAAARecord = &v
	return s
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) SetEnableInstanceIdDnsARecord(v bool) *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	s.EnableInstanceIdDnsARecord = &v
	return s
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) SetEnableIpDnsARecord(v bool) *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	s.EnableIpDnsARecord = &v
	return s
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) SetEnableIpDnsPtrRecord(v bool) *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	s.EnableIpDnsPtrRecord = &v
	return s
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) SetHostnameType(v string) *ModifyInstanceAttributeRequestPrivateDnsNameOptions {
	s.HostnameType = &v
	return s
}

func (s *ModifyInstanceAttributeRequestPrivateDnsNameOptions) Validate() error {
	return dara.Validate(s)
}

type ModifyInstanceAttributeRequestRemoteConnectionOptions struct {
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// hide
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// hide
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ModifyInstanceAttributeRequestRemoteConnectionOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttributeRequestRemoteConnectionOptions) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttributeRequestRemoteConnectionOptions) GetPassword() *string {
	return s.Password
}

func (s *ModifyInstanceAttributeRequestRemoteConnectionOptions) GetType() *string {
	return s.Type
}

func (s *ModifyInstanceAttributeRequestRemoteConnectionOptions) SetPassword(v string) *ModifyInstanceAttributeRequestRemoteConnectionOptions {
	s.Password = &v
	return s
}

func (s *ModifyInstanceAttributeRequestRemoteConnectionOptions) SetType(v string) *ModifyInstanceAttributeRequestRemoteConnectionOptions {
	s.Type = &v
	return s
}

func (s *ModifyInstanceAttributeRequestRemoteConnectionOptions) Validate() error {
	return dara.Validate(s)
}
