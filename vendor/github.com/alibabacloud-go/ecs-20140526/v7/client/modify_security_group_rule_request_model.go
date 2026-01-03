// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifySecurityGroupRuleRequest
	GetClientToken() *string
	SetDescription(v string) *ModifySecurityGroupRuleRequest
	GetDescription() *string
	SetDestCidrIp(v string) *ModifySecurityGroupRuleRequest
	GetDestCidrIp() *string
	SetIpProtocol(v string) *ModifySecurityGroupRuleRequest
	GetIpProtocol() *string
	SetIpv6DestCidrIp(v string) *ModifySecurityGroupRuleRequest
	GetIpv6DestCidrIp() *string
	SetIpv6SourceCidrIp(v string) *ModifySecurityGroupRuleRequest
	GetIpv6SourceCidrIp() *string
	SetNicType(v string) *ModifySecurityGroupRuleRequest
	GetNicType() *string
	SetOwnerAccount(v string) *ModifySecurityGroupRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySecurityGroupRuleRequest
	GetOwnerId() *int64
	SetPolicy(v string) *ModifySecurityGroupRuleRequest
	GetPolicy() *string
	SetPortRange(v string) *ModifySecurityGroupRuleRequest
	GetPortRange() *string
	SetPortRangeListId(v string) *ModifySecurityGroupRuleRequest
	GetPortRangeListId() *string
	SetPriority(v string) *ModifySecurityGroupRuleRequest
	GetPriority() *string
	SetRegionId(v string) *ModifySecurityGroupRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySecurityGroupRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySecurityGroupRuleRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *ModifySecurityGroupRuleRequest
	GetSecurityGroupId() *string
	SetSecurityGroupRuleId(v string) *ModifySecurityGroupRuleRequest
	GetSecurityGroupRuleId() *string
	SetSourceCidrIp(v string) *ModifySecurityGroupRuleRequest
	GetSourceCidrIp() *string
	SetSourceGroupId(v string) *ModifySecurityGroupRuleRequest
	GetSourceGroupId() *string
	SetSourceGroupOwnerAccount(v string) *ModifySecurityGroupRuleRequest
	GetSourceGroupOwnerAccount() *string
	SetSourceGroupOwnerId(v int64) *ModifySecurityGroupRuleRequest
	GetSourceGroupOwnerId() *int64
	SetSourcePortRange(v string) *ModifySecurityGroupRuleRequest
	GetSourcePortRange() *string
	SetSourcePrefixListId(v string) *ModifySecurityGroupRuleRequest
	GetSourcePrefixListId() *string
}

type ModifySecurityGroupRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the security group rule. The description must be 1 to 512 characters in length.
	//
	// example:
	//
	// This is a new security group rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IPv4 CIDR block. IPv4 CIDR blocks and IPv4 addresses are supported.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// Network Layer /transport layer protocol. Two types of assignments are supported:
	//
	// 1.  The case-insensitive protocol name. Valid values:
	//
	// 	- ICMP
	//
	// 	- GRE
	//
	// 	- TCP
	//
	// 	- UDP
	//
	// 	- ALL: supports all protocols.
	//
	// 2.  The value of the IANA-compliant protocol number, which is an integer from 0 to 255. List of regions currently available:
	//
	// 	- Philippines (Manila)
	//
	// 	- UK (London)
	//
	// 	- Malaysia (Kuala Lumpur)
	//
	// 	- China (Hohhot)
	//
	// 	- China (Qingdao)
	//
	// 	- US (Silicon Valley)
	//
	// 	- Singapore
	//
	// example:
	//
	// all
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The destination IPv6 CIDR block. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// >  Only the IP addresses of instances in VPCs are supported. You cannot specify both Ipv6DestCidrIp and `DestCidrIp`.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// >  Only the IP addresses of instances in virtual private clouds (VPCs) are supported. You cannot specify both Ipv6SourceCidrIp and `SourceCidrIp`.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 2001:db8:1233:1a00::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The type of the network interface controller (NIC).
	//
	// >  You cannot modify this parameter when you modify a security group rule by specifying the ID of the rule. If you want to change the NIC type of a security group rule, you can create a security group rule of a desired NIC type and delete the existing rule.
	//
	// example:
	//
	// intranet
	NicType      *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The action of the security group rule. Valid values:
	//
	// 	- accept: allows access.
	//
	// 	- drop: denies access and returns no responses.
	//
	// Default value: accept.
	//
	// example:
	//
	// accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The range of destination ports that correspond to the transport layer protocol. Valid values:
	//
	// 	- If you set IpProtocol to TCP or UDP, the port number range is 1 to 65535. Separate the start port number and the end port number with a forward slash (/). Example: 1/200.
	//
	// 	- If you set IpProtocol to ICMP, the port number range is -1/-1.
	//
	// 	- If you set IpProtocol to GRE, the port number range is -1/-1.
	//
	// 	- If you set IpProtocol to ALL, the port number range is -1/-1.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the port list. You can call the `DescribePortRangeLists` operation to query the IDs of available port lists.
	//
	// 	- If you specify PortRange, this parameter is ignored.
	//
	// 	- If a security group is in the classic network, you cannot configure port lists in the rules of the security group. For information about the limits on security groups and port lists, see [Security group limits](~~25412#SecurityGroupQuota1~~).
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The priority of the security group rule. Valid values: 1 to 100.
	//
	// Default value: 1.
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
	// The security group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of the security group rule. You can call the [DescribeSecurityGroupAttribute](https://help.aliyun.com/document_detail/2679845.html) operation to query the IDs of security group rules in a security group.
	//
	// example:
	//
	// sgr-bp67acfmxa123b***
	SecurityGroupRuleId *string `json:"SecurityGroupRuleId,omitempty" xml:"SecurityGroupRuleId,omitempty"`
	// The source IPv4 CIDR block. IPv4 CIDR blocks and IPv4 addresses are supported.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The source security group ID. You must specify either `SourceGroupId` or `SourceCidrIp` or specify both of them.
	//
	// 	- If `SourceGroupId` is specified but `SourceCidrIp` is not specified, the value of `NicType` must be set to intranet.
	//
	// 	- If both `SourceGroupId` and `SourceCidrIp` are specified, the value of `SourceCidrIp` prevails by default.
	//
	// example:
	//
	// sg-bp67acfmxa123b****
	SourceGroupId *string `json:"SourceGroupId,omitempty" xml:"SourceGroupId,omitempty"`
	// The Alibaba Cloud account that manages the source security group when you configure a security group rule across accounts.
	//
	// 	- If both `SourceGroupOwnerId` and `SourceGroupOwnerAccount` are empty, access permissions are configured for another security group managed by your account.
	//
	// 	- If `SourceCidrIp` is specified, `SourceGroupOwnerAccount` is ignored.
	//
	// example:
	//
	// EcsforCloud@Alibaba.com
	SourceGroupOwnerAccount *string `json:"SourceGroupOwnerAccount,omitempty" xml:"SourceGroupOwnerAccount,omitempty"`
	// The ID of the Alibaba Cloud account that manages the source security group when you configure a security group rule across accounts.
	//
	// 	- If both `SourceGroupOwnerId` and `SourceGroupOwnerAccount` are empty, access permissions are configured for another security group managed by your account.
	//
	// 	- If `SourceCidrIp` is specified, `SourceGroupOwnerId` is ignored.
	//
	// example:
	//
	// 12345678910
	SourceGroupOwnerId *int64 `json:"SourceGroupOwnerId,omitempty" xml:"SourceGroupOwnerId,omitempty"`
	// The range of source ports that correspond to the transport layer protocol. Valid values:
	//
	// 	- If you set IpProtocol to TCP or UDP, the port number range is 1 to 65535. Separate the start port number and the end port number with a forward slash (/). Example: 1/200.
	//
	// 	- If you set IpProtocol to ICMP, the port number range is -1/-1.
	//
	// 	- If you set IpProtocol to GRE, the port number range is -1/-1.
	//
	// 	- If you set IpProtocol to ALL, the port number range is -1/-1.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The ID of the source prefix list to which you want to control access. You can call the [DescribePrefixLists](https://help.aliyun.com/document_detail/205046.html) operation to query the IDs of available prefix lists.
	//
	// If you specify `SourceCidrIp`, `Ipv6SourceCidrIp`, or `SourceGroupId`, this parameter is ignored.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	SourcePrefixListId *string `json:"SourcePrefixListId,omitempty" xml:"SourcePrefixListId,omitempty"`
}

func (s ModifySecurityGroupRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifySecurityGroupRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySecurityGroupRuleRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *ModifySecurityGroupRuleRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifySecurityGroupRuleRequest) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *ModifySecurityGroupRuleRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *ModifySecurityGroupRuleRequest) GetNicType() *string {
	return s.NicType
}

func (s *ModifySecurityGroupRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySecurityGroupRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySecurityGroupRuleRequest) GetPolicy() *string {
	return s.Policy
}

func (s *ModifySecurityGroupRuleRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifySecurityGroupRuleRequest) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *ModifySecurityGroupRuleRequest) GetPriority() *string {
	return s.Priority
}

func (s *ModifySecurityGroupRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityGroupRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySecurityGroupRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityGroupRuleRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifySecurityGroupRuleRequest) GetSecurityGroupRuleId() *string {
	return s.SecurityGroupRuleId
}

func (s *ModifySecurityGroupRuleRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifySecurityGroupRuleRequest) GetSourceGroupId() *string {
	return s.SourceGroupId
}

func (s *ModifySecurityGroupRuleRequest) GetSourceGroupOwnerAccount() *string {
	return s.SourceGroupOwnerAccount
}

func (s *ModifySecurityGroupRuleRequest) GetSourceGroupOwnerId() *int64 {
	return s.SourceGroupOwnerId
}

func (s *ModifySecurityGroupRuleRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifySecurityGroupRuleRequest) GetSourcePrefixListId() *string {
	return s.SourcePrefixListId
}

func (s *ModifySecurityGroupRuleRequest) SetClientToken(v string) *ModifySecurityGroupRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetDescription(v string) *ModifySecurityGroupRuleRequest {
	s.Description = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetDestCidrIp(v string) *ModifySecurityGroupRuleRequest {
	s.DestCidrIp = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetIpProtocol(v string) *ModifySecurityGroupRuleRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetIpv6DestCidrIp(v string) *ModifySecurityGroupRuleRequest {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetIpv6SourceCidrIp(v string) *ModifySecurityGroupRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetNicType(v string) *ModifySecurityGroupRuleRequest {
	s.NicType = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetOwnerAccount(v string) *ModifySecurityGroupRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetOwnerId(v int64) *ModifySecurityGroupRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetPolicy(v string) *ModifySecurityGroupRuleRequest {
	s.Policy = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetPortRange(v string) *ModifySecurityGroupRuleRequest {
	s.PortRange = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetPortRangeListId(v string) *ModifySecurityGroupRuleRequest {
	s.PortRangeListId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetPriority(v string) *ModifySecurityGroupRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetRegionId(v string) *ModifySecurityGroupRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetResourceOwnerAccount(v string) *ModifySecurityGroupRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetResourceOwnerId(v int64) *ModifySecurityGroupRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSecurityGroupId(v string) *ModifySecurityGroupRuleRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSecurityGroupRuleId(v string) *ModifySecurityGroupRuleRequest {
	s.SecurityGroupRuleId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourceCidrIp(v string) *ModifySecurityGroupRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourceGroupId(v string) *ModifySecurityGroupRuleRequest {
	s.SourceGroupId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourceGroupOwnerAccount(v string) *ModifySecurityGroupRuleRequest {
	s.SourceGroupOwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourceGroupOwnerId(v int64) *ModifySecurityGroupRuleRequest {
	s.SourceGroupOwnerId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourcePortRange(v string) *ModifySecurityGroupRuleRequest {
	s.SourcePortRange = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) SetSourcePrefixListId(v string) *ModifySecurityGroupRuleRequest {
	s.SourcePrefixListId = &v
	return s
}

func (s *ModifySecurityGroupRuleRequest) Validate() error {
	return dara.Validate(s)
}
