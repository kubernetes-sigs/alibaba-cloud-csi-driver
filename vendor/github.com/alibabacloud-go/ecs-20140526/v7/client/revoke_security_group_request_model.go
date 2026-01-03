// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevokeSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RevokeSecurityGroupRequest
	GetClientToken() *string
	SetDescription(v string) *RevokeSecurityGroupRequest
	GetDescription() *string
	SetDestCidrIp(v string) *RevokeSecurityGroupRequest
	GetDestCidrIp() *string
	SetIpProtocol(v string) *RevokeSecurityGroupRequest
	GetIpProtocol() *string
	SetIpv6DestCidrIp(v string) *RevokeSecurityGroupRequest
	GetIpv6DestCidrIp() *string
	SetIpv6SourceCidrIp(v string) *RevokeSecurityGroupRequest
	GetIpv6SourceCidrIp() *string
	SetNicType(v string) *RevokeSecurityGroupRequest
	GetNicType() *string
	SetOwnerAccount(v string) *RevokeSecurityGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RevokeSecurityGroupRequest
	GetOwnerId() *int64
	SetPermissions(v []*RevokeSecurityGroupRequestPermissions) *RevokeSecurityGroupRequest
	GetPermissions() []*RevokeSecurityGroupRequestPermissions
	SetPolicy(v string) *RevokeSecurityGroupRequest
	GetPolicy() *string
	SetPortRange(v string) *RevokeSecurityGroupRequest
	GetPortRange() *string
	SetPriority(v string) *RevokeSecurityGroupRequest
	GetPriority() *string
	SetRegionId(v string) *RevokeSecurityGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *RevokeSecurityGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RevokeSecurityGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *RevokeSecurityGroupRequest
	GetSecurityGroupId() *string
	SetSecurityGroupRuleId(v []*string) *RevokeSecurityGroupRequest
	GetSecurityGroupRuleId() []*string
	SetSourceCidrIp(v string) *RevokeSecurityGroupRequest
	GetSourceCidrIp() *string
	SetSourceGroupId(v string) *RevokeSecurityGroupRequest
	GetSourceGroupId() *string
	SetSourceGroupOwnerAccount(v string) *RevokeSecurityGroupRequest
	GetSourceGroupOwnerAccount() *string
	SetSourceGroupOwnerId(v int64) *RevokeSecurityGroupRequest
	GetSourceGroupOwnerId() *int64
	SetSourcePortRange(v string) *RevokeSecurityGroupRequest
	GetSourcePortRange() *string
	SetSourcePrefixListId(v string) *RevokeSecurityGroupRequest
	GetSourcePrefixListId() *string
}

type RevokeSecurityGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **ClientToken*	- value can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.Description` to specify the rule description.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.DestCidrIp` to specify the destination IPv4 CIDR block.
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.IpProtocol` to specify the protocol.
	//
	// example:
	//
	// ALL
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.Ipv6DestCidrIp` to specify the destination IPv6 CIDR block.
	//
	// example:
	//
	// 2001:db8:1233:1a00::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.Ipv6SourceCidrIp` to specify the source IPv6 CIDR block.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.NicType` to specify the network interface type.
	//
	// example:
	//
	// intranet
	NicType      *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The security group rules. You can specify up to 100 security group rules.
	Permissions []*RevokeSecurityGroupRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.Policy` to specify whether to allow inbound access.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.PortRange` to specify the range of destination ports.
	//
	// example:
	//
	// 1/200
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.Priority` to specify the rule priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the security group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the security group.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The security group rule IDs. You can specify 1 to 100 security group rule IDs in a request.
	SecurityGroupRuleId []*string `json:"SecurityGroupRuleId,omitempty" xml:"SecurityGroupRuleId,omitempty" type:"Repeated"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.SourceCidrIp` to specify the source IPv4 CIDR block.
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.SourceGroupId` to specify the ID of the source security group.
	//
	// example:
	//
	// sg-bp67acfmxa123b****
	SourceGroupId *string `json:"SourceGroupId,omitempty" xml:"SourceGroupId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.SourceGroupOwnerAccount` to specify the Alibaba Cloud account that manages the source security group.
	//
	// example:
	//
	// Test@aliyun.com
	SourceGroupOwnerAccount *string `json:"SourceGroupOwnerAccount,omitempty" xml:"SourceGroupOwnerAccount,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.SourceGroupOwnerId` to specify the ID of the Alibaba Cloud account that manages the source security group.
	//
	// example:
	//
	// 12345678910
	SourceGroupOwnerId *int64 `json:"SourceGroupOwnerId,omitempty" xml:"SourceGroupOwnerId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.SourcePortRange` to specify the range of source ports.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.SourcePrefixListId` to specify the ID of the source prefix list.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	SourcePrefixListId *string `json:"SourcePrefixListId,omitempty" xml:"SourcePrefixListId,omitempty"`
}

func (s RevokeSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RevokeSecurityGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *RevokeSecurityGroupRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *RevokeSecurityGroupRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *RevokeSecurityGroupRequest) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *RevokeSecurityGroupRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *RevokeSecurityGroupRequest) GetNicType() *string {
	return s.NicType
}

func (s *RevokeSecurityGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RevokeSecurityGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RevokeSecurityGroupRequest) GetPermissions() []*RevokeSecurityGroupRequestPermissions {
	return s.Permissions
}

func (s *RevokeSecurityGroupRequest) GetPolicy() *string {
	return s.Policy
}

func (s *RevokeSecurityGroupRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *RevokeSecurityGroupRequest) GetPriority() *string {
	return s.Priority
}

func (s *RevokeSecurityGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RevokeSecurityGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RevokeSecurityGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RevokeSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RevokeSecurityGroupRequest) GetSecurityGroupRuleId() []*string {
	return s.SecurityGroupRuleId
}

func (s *RevokeSecurityGroupRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *RevokeSecurityGroupRequest) GetSourceGroupId() *string {
	return s.SourceGroupId
}

func (s *RevokeSecurityGroupRequest) GetSourceGroupOwnerAccount() *string {
	return s.SourceGroupOwnerAccount
}

func (s *RevokeSecurityGroupRequest) GetSourceGroupOwnerId() *int64 {
	return s.SourceGroupOwnerId
}

func (s *RevokeSecurityGroupRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *RevokeSecurityGroupRequest) GetSourcePrefixListId() *string {
	return s.SourcePrefixListId
}

func (s *RevokeSecurityGroupRequest) SetClientToken(v string) *RevokeSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetDescription(v string) *RevokeSecurityGroupRequest {
	s.Description = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetDestCidrIp(v string) *RevokeSecurityGroupRequest {
	s.DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetIpProtocol(v string) *RevokeSecurityGroupRequest {
	s.IpProtocol = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetIpv6DestCidrIp(v string) *RevokeSecurityGroupRequest {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetIpv6SourceCidrIp(v string) *RevokeSecurityGroupRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetNicType(v string) *RevokeSecurityGroupRequest {
	s.NicType = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetOwnerAccount(v string) *RevokeSecurityGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetOwnerId(v int64) *RevokeSecurityGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPermissions(v []*RevokeSecurityGroupRequestPermissions) *RevokeSecurityGroupRequest {
	s.Permissions = v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPolicy(v string) *RevokeSecurityGroupRequest {
	s.Policy = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPortRange(v string) *RevokeSecurityGroupRequest {
	s.PortRange = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetPriority(v string) *RevokeSecurityGroupRequest {
	s.Priority = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetRegionId(v string) *RevokeSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetResourceOwnerAccount(v string) *RevokeSecurityGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetResourceOwnerId(v int64) *RevokeSecurityGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSecurityGroupId(v string) *RevokeSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSecurityGroupRuleId(v []*string) *RevokeSecurityGroupRequest {
	s.SecurityGroupRuleId = v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourceCidrIp(v string) *RevokeSecurityGroupRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourceGroupId(v string) *RevokeSecurityGroupRequest {
	s.SourceGroupId = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourceGroupOwnerAccount(v string) *RevokeSecurityGroupRequest {
	s.SourceGroupOwnerAccount = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourceGroupOwnerId(v int64) *RevokeSecurityGroupRequest {
	s.SourceGroupOwnerId = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourcePortRange(v string) *RevokeSecurityGroupRequest {
	s.SourcePortRange = &v
	return s
}

func (s *RevokeSecurityGroupRequest) SetSourcePrefixListId(v string) *RevokeSecurityGroupRequest {
	s.SourcePrefixListId = &v
	return s
}

func (s *RevokeSecurityGroupRequest) Validate() error {
	if s.Permissions != nil {
		for _, item := range s.Permissions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RevokeSecurityGroupRequestPermissions struct {
	// The description of the security group rule. The description must be 1 to 512 characters in length.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IPv4 CIDR block. IPv4 CIDR blocks and IPv4 addresses are supported.
	//
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The protocol. The values of this parameter are case-insensitive. Valid values:
	//
	// 	- TCP.
	//
	// 	- UDP.
	//
	// 	- ICMP.
	//
	// 	- ICMPv6.
	//
	// 	- GRE.
	//
	// 	- ALL: All protocols are supported.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The destination IPv6 CIDR block. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// >  This parameter is valid only for ECS instances that reside in VPCs and support IPv6 CIDR blocks. You cannot specify both this parameter and `DestCidrIp` in the same request.
	//
	// example:
	//
	// 2001:db8:1233:1a00::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block of the security group rule. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// >  This parameter is valid only for Elastic Compute Service (ECS) instances that reside in virtual private clouds (VPCs) and support IPv6 CIDR blocks. You cannot specify both this parameter and `SourceCidrIp` in the same request.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The network interface controller (NIC) type of the security group rule if the security group resides in the classic network. Valid values:
	//
	// 	- internet: public NIC.
	//
	// 	- intranet: internal NIC.
	//
	// If the security group resides in a VPC, this parameter is set to intranet by default and cannot be modified.
	//
	// If you specify `SourceGroupId` to delete inbound security group rules that reference the specified security group as an authorization object, you must set this parameter to intranet.
	//
	// Default value: internet.
	//
	// example:
	//
	// intranet
	NicType *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	// The action of the security group rule. Valid values:
	//
	// 	- accept: allows inbound access.
	//
	// 	- drop: denies inbound access and returns no responses. In this case, the request times out or the connection cannot be established.
	//
	// Default value: accept.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The destination port range of the security group rule. Valid values:
	//
	// 	- If you set IpProtocol to TCP or UDP, the valid values of this parameter are 1 to 65535. Specify a port range in the format of \\<Start port number>/\\<End port number>. Example: 1/200.
	//
	// 	- If you set IpProtocol to ICMP, the port range is -1/-1.
	//
	// 	- If you set IpProtocol to GRE, the port range is -1/-1.
	//
	// 	- If you set IpProtocol to ALL, the port range is -1/-1.
	//
	// example:
	//
	// 1/200
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the port list. You can call the `DescribePortRangeLists` operation to query the IDs of available port lists.
	//
	// 	- If you specify `Permissions.N.PortRange`, this parameter is ignored.
	//
	// 	- If a security group resides in the classic network, you cannot reference port lists in the rules of the security group. For information about the limits on security groups and port lists, see the [Security groups](~~25412#SecurityGroupQuota1~~) section of the "Limits and quotas on ECS" topic.
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The priority of the security group rule. A smaller value specifies a higher priority. Valid values: 1 to 100.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The source IPv4 CIDR block of the security group rule. IPv4 CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The ID of the source security group referenced in the security group rule.
	//
	// 	- You must specify at least one of the following parameters: `SourceGroupId`, `SourceCidrIp`, `Ipv6SourceCidrIp`, and `SourcePrefixListId`.
	//
	// 	- If you specify `SourceGroupId` but do not specify `SourceCidrIp` or `Ipv6SourceCidrIp`, you must set NicType to intranet.
	//
	// 	- If you specify both `SourceGroupId` and `SourceCidrIp`, `SourceCidrIp` takes precedence.
	//
	// Take note of the following items:
	//
	// 	- Advanced security groups do not support security group rules that reference security groups as authorization objects (sources or destinations of traffic).
	//
	// 	- Each basic security group can contain up to 20 security group rules that reference security groups as authorization objects.
	//
	// example:
	//
	// sg-bp67acfmxa123b****
	SourceGroupId *string `json:"SourceGroupId,omitempty" xml:"SourceGroupId,omitempty"`
	// The Alibaba Cloud account that manages the source security group referenced in the security group rule.
	//
	// 	- If both `SourceGroupOwnerAccount` and `SourceGroupOwnerId` are empty, access control on another security group in your Alibaba Cloud account is removed.
	//
	// 	- If you specify `SourceCidrIp`, `SourceGroupOwnerAccount` is ignored.
	//
	// example:
	//
	// Test@aliyun.com
	SourceGroupOwnerAccount *string `json:"SourceGroupOwnerAccount,omitempty" xml:"SourceGroupOwnerAccount,omitempty"`
	// The ID of the Alibaba Cloud account that manages the source security group referenced in the security group rule.
	//
	// 	- If both `SourceGroupOwnerId` and `SourceGroupOwnerAccount` are empty, access control on another security group in your Alibaba Cloud account is removed.
	//
	// 	- If you specify `SourceCidrIp`, `SourceGroupOwnerId` is ignored.
	//
	// example:
	//
	// 12345678910
	SourceGroupOwnerId *int64 `json:"SourceGroupOwnerId,omitempty" xml:"SourceGroupOwnerId,omitempty"`
	// The source port range of the security group rule. Valid values:
	//
	// 	- If you set IpProtocol to TCP or UDP, the valid values of this parameter are 1 to 65535. Specify a port range in the format of \\<Start port number>/\\<End port number>. Example: 1/200.
	//
	// 	- If you set IpProtocol to ICMP, the port range is -1/-1.
	//
	// 	- If you set IpProtocol to GRE, the port range is -1/-1.
	//
	// 	- If you set IpProtocol to ALL, the port range is -1/-1.
	//
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The ID of the source prefix list of the security group rule. You can call the [DescribePrefixLists](https://help.aliyun.com/document_detail/205046.html) operation to query the IDs of available prefix lists.
	//
	// Take note of the following items:
	//
	// 	- If a security group resides in the classic network, you cannot specify prefix lists in the rules of the security group. For information about the limits on security groups and prefix lists, see the [Security groups](~~25412#SecurityGroupQuota1~~) section of the "Limits and quotas on ECS" topic.
	//
	// 	- If you specify `SourceCidrIp`, `Ipv6SourceCidrIp`, or `SourceGroupId`, this parameter is ignored.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	SourcePrefixListId *string `json:"SourcePrefixListId,omitempty" xml:"SourcePrefixListId,omitempty"`
}

func (s RevokeSecurityGroupRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s RevokeSecurityGroupRequestPermissions) GoString() string {
	return s.String()
}

func (s *RevokeSecurityGroupRequestPermissions) GetDescription() *string {
	return s.Description
}

func (s *RevokeSecurityGroupRequestPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *RevokeSecurityGroupRequestPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *RevokeSecurityGroupRequestPermissions) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *RevokeSecurityGroupRequestPermissions) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *RevokeSecurityGroupRequestPermissions) GetNicType() *string {
	return s.NicType
}

func (s *RevokeSecurityGroupRequestPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *RevokeSecurityGroupRequestPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *RevokeSecurityGroupRequestPermissions) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *RevokeSecurityGroupRequestPermissions) GetPriority() *string {
	return s.Priority
}

func (s *RevokeSecurityGroupRequestPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *RevokeSecurityGroupRequestPermissions) GetSourceGroupId() *string {
	return s.SourceGroupId
}

func (s *RevokeSecurityGroupRequestPermissions) GetSourceGroupOwnerAccount() *string {
	return s.SourceGroupOwnerAccount
}

func (s *RevokeSecurityGroupRequestPermissions) GetSourceGroupOwnerId() *int64 {
	return s.SourceGroupOwnerId
}

func (s *RevokeSecurityGroupRequestPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *RevokeSecurityGroupRequestPermissions) GetSourcePrefixListId() *string {
	return s.SourcePrefixListId
}

func (s *RevokeSecurityGroupRequestPermissions) SetDescription(v string) *RevokeSecurityGroupRequestPermissions {
	s.Description = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetDestCidrIp(v string) *RevokeSecurityGroupRequestPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetIpProtocol(v string) *RevokeSecurityGroupRequestPermissions {
	s.IpProtocol = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetIpv6DestCidrIp(v string) *RevokeSecurityGroupRequestPermissions {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetIpv6SourceCidrIp(v string) *RevokeSecurityGroupRequestPermissions {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetNicType(v string) *RevokeSecurityGroupRequestPermissions {
	s.NicType = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetPolicy(v string) *RevokeSecurityGroupRequestPermissions {
	s.Policy = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetPortRange(v string) *RevokeSecurityGroupRequestPermissions {
	s.PortRange = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetPortRangeListId(v string) *RevokeSecurityGroupRequestPermissions {
	s.PortRangeListId = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetPriority(v string) *RevokeSecurityGroupRequestPermissions {
	s.Priority = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetSourceCidrIp(v string) *RevokeSecurityGroupRequestPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetSourceGroupId(v string) *RevokeSecurityGroupRequestPermissions {
	s.SourceGroupId = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetSourceGroupOwnerAccount(v string) *RevokeSecurityGroupRequestPermissions {
	s.SourceGroupOwnerAccount = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetSourceGroupOwnerId(v int64) *RevokeSecurityGroupRequestPermissions {
	s.SourceGroupOwnerId = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetSourcePortRange(v string) *RevokeSecurityGroupRequestPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) SetSourcePrefixListId(v string) *RevokeSecurityGroupRequestPermissions {
	s.SourcePrefixListId = &v
	return s
}

func (s *RevokeSecurityGroupRequestPermissions) Validate() error {
	return dara.Validate(s)
}
