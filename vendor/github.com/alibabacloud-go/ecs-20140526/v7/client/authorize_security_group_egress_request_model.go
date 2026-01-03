// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSecurityGroupEgressRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AuthorizeSecurityGroupEgressRequest
	GetClientToken() *string
	SetDescription(v string) *AuthorizeSecurityGroupEgressRequest
	GetDescription() *string
	SetDestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest
	GetDestCidrIp() *string
	SetDestGroupId(v string) *AuthorizeSecurityGroupEgressRequest
	GetDestGroupId() *string
	SetDestGroupOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest
	GetDestGroupOwnerAccount() *string
	SetDestGroupOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest
	GetDestGroupOwnerId() *int64
	SetDestPrefixListId(v string) *AuthorizeSecurityGroupEgressRequest
	GetDestPrefixListId() *string
	SetIpProtocol(v string) *AuthorizeSecurityGroupEgressRequest
	GetIpProtocol() *string
	SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest
	GetIpv6DestCidrIp() *string
	SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequest
	GetIpv6SourceCidrIp() *string
	SetNicType(v string) *AuthorizeSecurityGroupEgressRequest
	GetNicType() *string
	SetOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest
	GetOwnerId() *int64
	SetPermissions(v []*AuthorizeSecurityGroupEgressRequestPermissions) *AuthorizeSecurityGroupEgressRequest
	GetPermissions() []*AuthorizeSecurityGroupEgressRequestPermissions
	SetPolicy(v string) *AuthorizeSecurityGroupEgressRequest
	GetPolicy() *string
	SetPortRange(v string) *AuthorizeSecurityGroupEgressRequest
	GetPortRange() *string
	SetPriority(v string) *AuthorizeSecurityGroupEgressRequest
	GetPriority() *string
	SetRegionId(v string) *AuthorizeSecurityGroupEgressRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *AuthorizeSecurityGroupEgressRequest
	GetSecurityGroupId() *string
	SetSourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequest
	GetSourceCidrIp() *string
	SetSourcePortRange(v string) *AuthorizeSecurityGroupEgressRequest
	GetSourcePortRange() *string
}

type AuthorizeSecurityGroupEgressRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.Description` to specify the description of security group rule N.
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
	// This parameter is deprecated. Use `Permissions.N.DestGroupId` to specify the ID of the destination security group.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	DestGroupId *string `json:"DestGroupId,omitempty" xml:"DestGroupId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.DestGroupOwnerAccount` to specify the Alibaba Cloud account that manages the destination security group.
	//
	// example:
	//
	// Test@aliyun.com
	DestGroupOwnerAccount *string `json:"DestGroupOwnerAccount,omitempty" xml:"DestGroupOwnerAccount,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.DestGroupOwnerId` to specify the ID of the Alibaba Cloud account that manages the destination security group.
	//
	// example:
	//
	// 12345678910
	DestGroupOwnerId *int64 `json:"DestGroupOwnerId,omitempty" xml:"DestGroupOwnerId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.DestPrefixListId` to specify the ID of the destination prefix list.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	DestPrefixListId *string `json:"DestPrefixListId,omitempty" xml:"DestPrefixListId,omitempty"`
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
	// This parameter is deprecated. Use `Permissions.N.NicType` to specify the NIC type.
	//
	// example:
	//
	// intranet
	NicType      *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The security group rules. You can specify 1 to 100 security group rules.
	Permissions []*AuthorizeSecurityGroupEgressRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.Policy` to specify whether to allow outbound access.
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
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.Priority` to specify the rule priority.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The region ID of the source security group. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
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
	// This parameter is deprecated. Use `Permissions.N.SourcePortRange` to specify the range of source ports.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s AuthorizeSecurityGroupEgressRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDescription() *string {
	return s.Description
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestGroupId() *string {
	return s.DestGroupId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestGroupOwnerAccount() *string {
	return s.DestGroupOwnerAccount
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestGroupOwnerId() *int64 {
	return s.DestGroupOwnerId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetDestPrefixListId() *string {
	return s.DestPrefixListId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeSecurityGroupEgressRequest) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequest) GetNicType() *string {
	return s.NicType
}

func (s *AuthorizeSecurityGroupEgressRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AuthorizeSecurityGroupEgressRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPermissions() []*AuthorizeSecurityGroupEgressRequestPermissions {
	return s.Permissions
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeSecurityGroupEgressRequest) GetPriority() *string {
	return s.Priority
}

func (s *AuthorizeSecurityGroupEgressRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AuthorizeSecurityGroupEgressRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AuthorizeSecurityGroupEgressRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeSecurityGroupEgressRequest) SetClientToken(v string) *AuthorizeSecurityGroupEgressRequest {
	s.ClientToken = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDescription(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Description = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestGroupId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestGroupOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestGroupOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestGroupOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest {
	s.DestGroupOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetDestPrefixListId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.DestPrefixListId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetIpProtocol(v string) *AuthorizeSecurityGroupEgressRequest {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetNicType(v string) *AuthorizeSecurityGroupEgressRequest {
	s.NicType = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest {
	s.OwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPermissions(v []*AuthorizeSecurityGroupEgressRequestPermissions) *AuthorizeSecurityGroupEgressRequest {
	s.Permissions = v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPolicy(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPortRange(v string) *AuthorizeSecurityGroupEgressRequest {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetPriority(v string) *AuthorizeSecurityGroupEgressRequest {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetRegionId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetResourceOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetResourceOwnerId(v int64) *AuthorizeSecurityGroupEgressRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSecurityGroupId(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) SetSourcePortRange(v string) *AuthorizeSecurityGroupEgressRequest {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequest) Validate() error {
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

type AuthorizeSecurityGroupEgressRequestPermissions struct {
	// The description of the security group rule. The description must be 1 to 512 characters in length.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IPv4 CIDR block of the security group rule. IPv4 CIDR blocks and IPv4 addresses are supported.
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The ID of the destination security group.
	//
	// 	- You must specify at least one of the following parameters: `DestGroupId`, `DestCidrIp`, `Ipv6DestCidrIp`, and `DestPrefixListId`.
	//
	// 	- If you specify `DestGroupId` but do not specify `DestCidrIp`, you must set `NicType` to intranet.
	//
	// 	- If you specify both `DestGroupId` and `DestCidrIp`, `DestCidrIp` takes precedence.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	DestGroupId *string `json:"DestGroupId,omitempty" xml:"DestGroupId,omitempty"`
	// The Alibaba Cloud account that manages the destination security group.
	//
	// 	- If both `DestGroupOwnerAccount` and `DestGroupOwnerId` are empty, the rule is created to control access to another security group in your Alibaba Cloud account.
	//
	// 	- If `DestCidrIp` is configured, `DestGroupOwnerAccount` is ignored.
	//
	// example:
	//
	// Test@aliyun.com
	DestGroupOwnerAccount *string `json:"DestGroupOwnerAccount,omitempty" xml:"DestGroupOwnerAccount,omitempty"`
	// The ID of the Alibaba Cloud account that manages the destination security group.
	//
	// 	- If both `DestGroupOwnerId` and `DestGroupOwnerAccount` are empty, the rule is created to control access to another security group in your Alibaba Cloud account.
	//
	// 	- If you specify `DestCidrIp`, `DestGroupOwnerId` is ignored.
	//
	// example:
	//
	// 12345678910
	DestGroupOwnerId *int64 `json:"DestGroupOwnerId,omitempty" xml:"DestGroupOwnerId,omitempty"`
	// The ID of the destination prefix list. You can call the [DescribePrefixLists](https://help.aliyun.com/document_detail/205046.html) operation to query the IDs of available prefix lists.
	//
	// When you specify this parameter, take note of the following items:
	//
	// 	- If a security group resides in the classic network, you cannot specify prefix lists in the rules of the security group. For information about the limits on security groups and prefix lists, see the [Security group limits](~~25412#SecurityGroupQuota1~~) section of the "Limits and quotas" topic.
	//
	// 	- If you specify `DestCidrIp`, `Ipv6DestCidrIp`, or `DestGroupId`, this parameter is ignored.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	DestPrefixListId *string `json:"DestPrefixListId,omitempty" xml:"DestPrefixListId,omitempty"`
	// The protocol. The values of this parameter are case-insensitive. Specifies whether to check that the CPU tag set of the source host is the subset of the CPU tag set of the destination host. Valid values:
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
	// ALL
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The destination IPv6 CIDR block of the security group rule. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// >  This parameter is valid only for ECS instances that reside in virtual private clouds (VPCs) and support IPv6 CIDR blocks. You cannot specify this parameter and `DestCidrIp` in the same request.
	//
	// example:
	//
	// 2001:db8:1233:1a00::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block or IPv6 address.
	//
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// >  This parameter is valid only for ECS instances that reside in VPCs and support IPv6 CIDR blocks. You cannot specify this parameter and `DestCidrIp` in the same request.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The network interface controller (NIC) type of the security group rule if the security group resides in the classic network. Specifies whether to check that the CPU tag set of the source host is the subset of the CPU tag set of the destination host. Valid values:
	//
	// 	- internet: public NIC.
	//
	// 	- intranet: internal NIC.
	//
	//     	- If the security group resides in a VPC, this parameter is set to intranet by default and cannot be changed.
	//
	//     	- If you specify only DestGroupId to create a rule that controls access to the specified security group, you must set this parameter to intranet.
	//
	// Default value: internet.
	//
	// example:
	//
	// intranet
	NicType *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	// The action of the security group rule. Specifies whether to check that the CPU tag set of the source host is the subset of the CPU tag set of the destination host. Valid values:
	//
	// 	- accept: allows outbound access.
	//
	// 	- drop: denies outbound access and returns no responses. In this case, the request times out or the connection cannot be established.
	//
	// Default value: accept.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination port numbers for the protocols specified in the security group rule. Specifies whether to check that the CPU tag set of the source host is the subset of the CPU tag set of the destination host. Valid values:
	//
	// 	- If you set IpProtocol to TCP or UDP, the port number range is 1 to 65535. Specify a port number range in the format of \\<Start port number>/\\<End port number>. Example: 1/200.
	//
	// 	- If you set IpProtocol to ICMP, the port number range is -1/-1.
	//
	// 	- If the IpProtocol parameter is set to GRE, the port number range is -1/-1, which indicates all ports.
	//
	// 	- If you set IpProtocol to ALL, the port number range is -1/-1.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
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
	// The source IPv4 CIDR blocks and IPv4 addresses are supported.
	//
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The range of source port numbers for the protocols specified in the security group rule. Specifies whether to check that the CPU tag set of the source host is the subset of the CPU tag set of the destination host. Valid values:
	//
	// 	- If you set IpProtocol to TCP or UDP, the port number range is 1 to 65535. Specify a port number range in the format of \\<Start port number>/\\<End port number>. Example: 1/200.
	//
	// 	- If you set IpProtocol to ICMP, the port number range is -1/-1.
	//
	// 	- If you set IpProtocol to GRE, the port number range is -1/-1.
	//
	// 	- If you set IpProtocol to ALL, the port number range is -1/-1.
	//
	// This property is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
}

func (s AuthorizeSecurityGroupEgressRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupEgressRequestPermissions) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDescription() *string {
	return s.Description
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDestGroupId() *string {
	return s.DestGroupId
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDestGroupOwnerAccount() *string {
	return s.DestGroupOwnerAccount
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDestGroupOwnerId() *int64 {
	return s.DestGroupOwnerId
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetDestPrefixListId() *string {
	return s.DestPrefixListId
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetNicType() *string {
	return s.NicType
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetPriority() *string {
	return s.Priority
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDescription(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.Description = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDestCidrIp(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDestGroupId(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.DestGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDestGroupOwnerAccount(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.DestGroupOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDestGroupOwnerId(v int64) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.DestGroupOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetDestPrefixListId(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.DestPrefixListId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetIpProtocol(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetNicType(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.NicType = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetPolicy(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetPortRange(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetPortRangeListId(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.PortRangeListId = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetPriority(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetSourceCidrIp(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) SetSourcePortRange(v string) *AuthorizeSecurityGroupEgressRequestPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupEgressRequestPermissions) Validate() error {
	return dara.Validate(s)
}
