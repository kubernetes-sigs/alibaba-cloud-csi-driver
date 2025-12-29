// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAuthorizeSecurityGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *AuthorizeSecurityGroupRequest
	GetClientToken() *string
	SetDescription(v string) *AuthorizeSecurityGroupRequest
	GetDescription() *string
	SetDestCidrIp(v string) *AuthorizeSecurityGroupRequest
	GetDestCidrIp() *string
	SetIpProtocol(v string) *AuthorizeSecurityGroupRequest
	GetIpProtocol() *string
	SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupRequest
	GetIpv6DestCidrIp() *string
	SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupRequest
	GetIpv6SourceCidrIp() *string
	SetNicType(v string) *AuthorizeSecurityGroupRequest
	GetNicType() *string
	SetOwnerAccount(v string) *AuthorizeSecurityGroupRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *AuthorizeSecurityGroupRequest
	GetOwnerId() *int64
	SetPermissions(v []*AuthorizeSecurityGroupRequestPermissions) *AuthorizeSecurityGroupRequest
	GetPermissions() []*AuthorizeSecurityGroupRequestPermissions
	SetPolicy(v string) *AuthorizeSecurityGroupRequest
	GetPolicy() *string
	SetPortRange(v string) *AuthorizeSecurityGroupRequest
	GetPortRange() *string
	SetPriority(v string) *AuthorizeSecurityGroupRequest
	GetPriority() *string
	SetRegionId(v string) *AuthorizeSecurityGroupRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *AuthorizeSecurityGroupRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *AuthorizeSecurityGroupRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *AuthorizeSecurityGroupRequest
	GetSecurityGroupId() *string
	SetSourceCidrIp(v string) *AuthorizeSecurityGroupRequest
	GetSourceCidrIp() *string
	SetSourceGroupId(v string) *AuthorizeSecurityGroupRequest
	GetSourceGroupId() *string
	SetSourceGroupOwnerAccount(v string) *AuthorizeSecurityGroupRequest
	GetSourceGroupOwnerAccount() *string
	SetSourceGroupOwnerId(v int64) *AuthorizeSecurityGroupRequest
	GetSourceGroupOwnerId() *int64
	SetSourcePortRange(v string) *AuthorizeSecurityGroupRequest
	GetSourcePortRange() *string
	SetSourcePrefixListId(v string) *AuthorizeSecurityGroupRequest
	GetSourcePrefixListId() *string
}

type AuthorizeSecurityGroupRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
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
	// 2001:250:6000::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.Ipv6SourceCidrIp` to specify the source IPv6 CIDR block.
	//
	// example:
	//
	// 2001:250:6000::***
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
	// The security group rules. You can specify 1 to 100 security group rules in a request.
	Permissions []*AuthorizeSecurityGroupRequestPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Repeated"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.Policy` to specify whether to allow access.
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
	// 22/22
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
	// sg-bp67acfmxazb4p****
	SourceGroupId *string `json:"SourceGroupId,omitempty" xml:"SourceGroupId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.SourceGroupOwnerAccount` to specify the Alibaba Cloud account that manages the source security group.
	//
	// example:
	//
	// test@aliyun.com
	SourceGroupOwnerAccount *string `json:"SourceGroupOwnerAccount,omitempty" xml:"SourceGroupOwnerAccount,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.SourceGroupOwnerId` to specify the ID of the Alibaba Cloud account that manages the source security group.
	//
	// example:
	//
	// 1234567890
	SourceGroupOwnerId *int64 `json:"SourceGroupOwnerId,omitempty" xml:"SourceGroupOwnerId,omitempty"`
	// Deprecated
	//
	// This parameter is deprecated. Use `Permissions.N.SourcePortRange` to specify the range of source ports.
	//
	// example:
	//
	// 22/22
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

func (s AuthorizeSecurityGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *AuthorizeSecurityGroupRequest) GetDescription() *string {
	return s.Description
}

func (s *AuthorizeSecurityGroupRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *AuthorizeSecurityGroupRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeSecurityGroupRequest) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *AuthorizeSecurityGroupRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *AuthorizeSecurityGroupRequest) GetNicType() *string {
	return s.NicType
}

func (s *AuthorizeSecurityGroupRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *AuthorizeSecurityGroupRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *AuthorizeSecurityGroupRequest) GetPermissions() []*AuthorizeSecurityGroupRequestPermissions {
	return s.Permissions
}

func (s *AuthorizeSecurityGroupRequest) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeSecurityGroupRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeSecurityGroupRequest) GetPriority() *string {
	return s.Priority
}

func (s *AuthorizeSecurityGroupRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *AuthorizeSecurityGroupRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *AuthorizeSecurityGroupRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *AuthorizeSecurityGroupRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *AuthorizeSecurityGroupRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *AuthorizeSecurityGroupRequest) GetSourceGroupId() *string {
	return s.SourceGroupId
}

func (s *AuthorizeSecurityGroupRequest) GetSourceGroupOwnerAccount() *string {
	return s.SourceGroupOwnerAccount
}

func (s *AuthorizeSecurityGroupRequest) GetSourceGroupOwnerId() *int64 {
	return s.SourceGroupOwnerId
}

func (s *AuthorizeSecurityGroupRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeSecurityGroupRequest) GetSourcePrefixListId() *string {
	return s.SourcePrefixListId
}

func (s *AuthorizeSecurityGroupRequest) SetClientToken(v string) *AuthorizeSecurityGroupRequest {
	s.ClientToken = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetDescription(v string) *AuthorizeSecurityGroupRequest {
	s.Description = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetDestCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetIpProtocol(v string) *AuthorizeSecurityGroupRequest {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetNicType(v string) *AuthorizeSecurityGroupRequest {
	s.NicType = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetOwnerAccount(v string) *AuthorizeSecurityGroupRequest {
	s.OwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetOwnerId(v int64) *AuthorizeSecurityGroupRequest {
	s.OwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPermissions(v []*AuthorizeSecurityGroupRequestPermissions) *AuthorizeSecurityGroupRequest {
	s.Permissions = v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPolicy(v string) *AuthorizeSecurityGroupRequest {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPortRange(v string) *AuthorizeSecurityGroupRequest {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetPriority(v string) *AuthorizeSecurityGroupRequest {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetRegionId(v string) *AuthorizeSecurityGroupRequest {
	s.RegionId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetResourceOwnerAccount(v string) *AuthorizeSecurityGroupRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetResourceOwnerId(v int64) *AuthorizeSecurityGroupRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSecurityGroupId(v string) *AuthorizeSecurityGroupRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceCidrIp(v string) *AuthorizeSecurityGroupRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceGroupId(v string) *AuthorizeSecurityGroupRequest {
	s.SourceGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceGroupOwnerAccount(v string) *AuthorizeSecurityGroupRequest {
	s.SourceGroupOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourceGroupOwnerId(v int64) *AuthorizeSecurityGroupRequest {
	s.SourceGroupOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourcePortRange(v string) *AuthorizeSecurityGroupRequest {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) SetSourcePrefixListId(v string) *AuthorizeSecurityGroupRequest {
	s.SourcePrefixListId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequest) Validate() error {
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

type AuthorizeSecurityGroupRequestPermissions struct {
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
	// Valid values of N: 1 to 100.
	//
	// example:
	//
	// ALL
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The destination IPv6 CIDR block. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// This parameter is used to support quintuple rules. For more information, see [Security group quintuple rules](https://help.aliyun.com/document_detail/97439.html).
	//
	// >  This parameter is valid only for ECS instances that reside in VPCs and support IPv6 CIDR blocks. You cannot specify this parameter and `DestCidrIp` in the same request.
	//
	// example:
	//
	// 2001:250:6000::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block of the security group rule. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// >  This parameter is valid only for Elastic Compute Service (ECS) instances that reside in virtual private clouds (VPCs) and support IPv6 CIDR blocks. You cannot specify both this parameter and `SourceCidrIp` in the same request.
	//
	// example:
	//
	// 2001:250:6000::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The network interface controller (NIC) type of the security group rule if the security group resides in the classic network. Valid values:
	//
	// 	- internet: public NIC.
	//
	// 	- intranet: internal NIC.
	//
	// If the security group resides in a VPC, this parameter is set to intranet by default and cannot be modified.
	//
	// If you specify only DestGroupId when you configure access permissions between security groups, you must set this parameter to intranet.
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
	// For more information, see [Common ports](https://help.aliyun.com/document_detail/40724.html).
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the port list. You can call the `DescribePortRangeLists` operation to query the IDs of available port lists.
	//
	// 	- If you specify `Permissions.N.PortRange`, this parameter is ignored.
	//
	// 	- If a security group resides in the classic network, you cannot reference port lists in the security group rules. For information about the limits on security groups and port lists, see the [Security groups](~~25412#SecurityGroupQuota1~~) section of the "Limits and quotas on ECS" topic.
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
	// 	- If you specify `SourceGroupId` but do not specify `SourceCidrIp` or `Ipv6SourceCidrIp`, you must set `NicType` to `intranet`.
	//
	// 	- If you specify both `SourceGroupId` and `SourceCidrIp`, `SourceCidrIp` takes precedence.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SourceGroupId *string `json:"SourceGroupId,omitempty" xml:"SourceGroupId,omitempty"`
	// The Alibaba Cloud account that manages the source security group referenced in the security group rule.
	//
	// 	- If both `SourceGroupOwnerAccount` and `SourceGroupOwnerId` are empty, access permissions are configured for another security group in your Alibaba Cloud account.
	//
	// 	- If you specify `SourceCidrIp`, `SourceGroupOwnerAccount` is ignored.
	//
	// example:
	//
	// test@aliyun.com
	SourceGroupOwnerAccount *string `json:"SourceGroupOwnerAccount,omitempty" xml:"SourceGroupOwnerAccount,omitempty"`
	// The ID of the Alibaba Cloud account that manages the source security group referenced in the security group rule.
	//
	// 	- If both `SourceGroupOwnerAccount` and `SourceGroupOwnerId` are empty, access permissions are configured for another security group in your Alibaba Cloud account.
	//
	// 	- If you specify `SourceCidrIp`, `SourceGroupOwnerAccount` is ignored.
	//
	// example:
	//
	// 1234567890
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
	// 7000/8000
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

func (s AuthorizeSecurityGroupRequestPermissions) String() string {
	return dara.Prettify(s)
}

func (s AuthorizeSecurityGroupRequestPermissions) GoString() string {
	return s.String()
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetDescription() *string {
	return s.Description
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetNicType() *string {
	return s.NicType
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetPolicy() *string {
	return s.Policy
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetPortRange() *string {
	return s.PortRange
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetPriority() *string {
	return s.Priority
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourceGroupId() *string {
	return s.SourceGroupId
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourceGroupOwnerAccount() *string {
	return s.SourceGroupOwnerAccount
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourceGroupOwnerId() *int64 {
	return s.SourceGroupOwnerId
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *AuthorizeSecurityGroupRequestPermissions) GetSourcePrefixListId() *string {
	return s.SourcePrefixListId
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetDescription(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.Description = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetDestCidrIp(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetIpProtocol(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.IpProtocol = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetIpv6DestCidrIp(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetIpv6SourceCidrIp(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetNicType(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.NicType = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetPolicy(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.Policy = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetPortRange(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.PortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetPortRangeListId(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.PortRangeListId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetPriority(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.Priority = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourceCidrIp(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.SourceCidrIp = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourceGroupId(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.SourceGroupId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourceGroupOwnerAccount(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.SourceGroupOwnerAccount = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourceGroupOwnerId(v int64) *AuthorizeSecurityGroupRequestPermissions {
	s.SourceGroupOwnerId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourcePortRange(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.SourcePortRange = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) SetSourcePrefixListId(v string) *AuthorizeSecurityGroupRequestPermissions {
	s.SourcePrefixListId = &v
	return s
}

func (s *AuthorizeSecurityGroupRequestPermissions) Validate() error {
	return dara.Validate(s)
}
