// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityGroupEgressRuleRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *ModifySecurityGroupEgressRuleRequest
	GetClientToken() *string
	SetDescription(v string) *ModifySecurityGroupEgressRuleRequest
	GetDescription() *string
	SetDestCidrIp(v string) *ModifySecurityGroupEgressRuleRequest
	GetDestCidrIp() *string
	SetDestGroupId(v string) *ModifySecurityGroupEgressRuleRequest
	GetDestGroupId() *string
	SetDestGroupOwnerAccount(v string) *ModifySecurityGroupEgressRuleRequest
	GetDestGroupOwnerAccount() *string
	SetDestGroupOwnerId(v int64) *ModifySecurityGroupEgressRuleRequest
	GetDestGroupOwnerId() *int64
	SetDestPrefixListId(v string) *ModifySecurityGroupEgressRuleRequest
	GetDestPrefixListId() *string
	SetIpProtocol(v string) *ModifySecurityGroupEgressRuleRequest
	GetIpProtocol() *string
	SetIpv6DestCidrIp(v string) *ModifySecurityGroupEgressRuleRequest
	GetIpv6DestCidrIp() *string
	SetIpv6SourceCidrIp(v string) *ModifySecurityGroupEgressRuleRequest
	GetIpv6SourceCidrIp() *string
	SetNicType(v string) *ModifySecurityGroupEgressRuleRequest
	GetNicType() *string
	SetOwnerAccount(v string) *ModifySecurityGroupEgressRuleRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifySecurityGroupEgressRuleRequest
	GetOwnerId() *int64
	SetPolicy(v string) *ModifySecurityGroupEgressRuleRequest
	GetPolicy() *string
	SetPortRange(v string) *ModifySecurityGroupEgressRuleRequest
	GetPortRange() *string
	SetPortRangeListId(v string) *ModifySecurityGroupEgressRuleRequest
	GetPortRangeListId() *string
	SetPriority(v string) *ModifySecurityGroupEgressRuleRequest
	GetPriority() *string
	SetRegionId(v string) *ModifySecurityGroupEgressRuleRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifySecurityGroupEgressRuleRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifySecurityGroupEgressRuleRequest
	GetResourceOwnerId() *int64
	SetSecurityGroupId(v string) *ModifySecurityGroupEgressRuleRequest
	GetSecurityGroupId() *string
	SetSecurityGroupRuleId(v string) *ModifySecurityGroupEgressRuleRequest
	GetSecurityGroupRuleId() *string
	SetSourceCidrIp(v string) *ModifySecurityGroupEgressRuleRequest
	GetSourceCidrIp() *string
	SetSourcePortRange(v string) *ModifySecurityGroupEgressRuleRequest
	GetSourcePortRange() *string
}

type ModifySecurityGroupEgressRuleRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.***	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The description of the security group rule. The description must be 1 to 512 characters in length.
	//
	// example:
	//
	// This is a new securitygroup rule.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination IPv4 CIDR block. IPv4 CIDR blocks and IPv4 addresses are supported.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 10.0.0.0/8
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The ID of the destination security group. You must specify at least one of `DestGroupId` and `DestCidrIp`.
	//
	// 	- At least one of DestGroupId, DestCidrIp, Ipv6DestCidrIp, and DestPrefixListId must be specified.
	//
	// 	- If DestGroupId is specified but DestCidrIp is not specified, the NicType parameter can be set only to intranet.
	//
	// 	- If both DestGroupId and DestCidrIp are specified, DestCidrIp takes precedence.
	//
	// example:
	//
	// sg-bp67acfmxa123b****
	DestGroupId *string `json:"DestGroupId,omitempty" xml:"DestGroupId,omitempty"`
	// The Alibaba Cloud account that manages the destination security group when you set security group rule N across accounts.
	//
	// example:
	//
	// EcsforCloud@Alibaba.com
	DestGroupOwnerAccount *string `json:"DestGroupOwnerAccount,omitempty" xml:"DestGroupOwnerAccount,omitempty"`
	// The ID of the Alibaba Cloud account that manages the destination security group when you set security group rule N across accounts.
	//
	// example:
	//
	// 1234567890
	DestGroupOwnerId *int64 `json:"DestGroupOwnerId,omitempty" xml:"DestGroupOwnerId,omitempty"`
	// The ID of the destination prefix list. You can call the [DescribePrefixLists](https://help.aliyun.com/document_detail/205046.html) operation to query the IDs of available prefix lists.
	//
	// If you specify `DestCidrIp`, `Ipv6DestCidrIp`, or `DestGroupId`, this parameter is ignored.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	DestPrefixListId *string `json:"DestPrefixListId,omitempty" xml:"DestPrefixListId,omitempty"`
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
	// tcp
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The destination IPv6 CIDR block. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// >  Only the IP addresses of instances in virtual private clouds (VPCs) are supported. You cannot specify both Ipv6DestCidrIp and `DestCidrIp`.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 2001:db8:1233:1a00::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block. IPv6 CIDR blocks and IPv6 addresses are supported.
	//
	// >  Only the IP addresses of instances in VPCs are supported. You cannot specify both Ipv6SourceCidrIp and `SourceCidrIp`.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The network interface controller (NIC) type.
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
	// The ID of the port list. You can call the `DescribePortRangeLists` operation to query the IDs of available prefix lists.
	//
	// 	- If you specify PortRange, the value of this parameter is ignored.
	//
	// 	- If the security group is of the classic network type, you cannot reference port lists in the security group rules. For information about the limits on security groups and port lists, see the [Security groups](~~25412#SecurityGroupQuota1~~) section of the "Limits and quotas" topic.
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
	// The ID of the security group.
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
	// sgr-bp67acfmxazb4q****
	SecurityGroupRuleId *string `json:"SecurityGroupRuleId,omitempty" xml:"SecurityGroupRuleId,omitempty"`
	// The source IPv4 CIDR block. IPv4 CIDR blocks and IPv4 addresses are supported.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 10.0.0.0/8
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
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
}

func (s ModifySecurityGroupEgressRuleRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityGroupEgressRuleRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityGroupEgressRuleRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifySecurityGroupEgressRuleRequest) GetDescription() *string {
	return s.Description
}

func (s *ModifySecurityGroupEgressRuleRequest) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *ModifySecurityGroupEgressRuleRequest) GetDestGroupId() *string {
	return s.DestGroupId
}

func (s *ModifySecurityGroupEgressRuleRequest) GetDestGroupOwnerAccount() *string {
	return s.DestGroupOwnerAccount
}

func (s *ModifySecurityGroupEgressRuleRequest) GetDestGroupOwnerId() *int64 {
	return s.DestGroupOwnerId
}

func (s *ModifySecurityGroupEgressRuleRequest) GetDestPrefixListId() *string {
	return s.DestPrefixListId
}

func (s *ModifySecurityGroupEgressRuleRequest) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *ModifySecurityGroupEgressRuleRequest) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *ModifySecurityGroupEgressRuleRequest) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *ModifySecurityGroupEgressRuleRequest) GetNicType() *string {
	return s.NicType
}

func (s *ModifySecurityGroupEgressRuleRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifySecurityGroupEgressRuleRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifySecurityGroupEgressRuleRequest) GetPolicy() *string {
	return s.Policy
}

func (s *ModifySecurityGroupEgressRuleRequest) GetPortRange() *string {
	return s.PortRange
}

func (s *ModifySecurityGroupEgressRuleRequest) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *ModifySecurityGroupEgressRuleRequest) GetPriority() *string {
	return s.Priority
}

func (s *ModifySecurityGroupEgressRuleRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityGroupEgressRuleRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifySecurityGroupEgressRuleRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifySecurityGroupEgressRuleRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ModifySecurityGroupEgressRuleRequest) GetSecurityGroupRuleId() *string {
	return s.SecurityGroupRuleId
}

func (s *ModifySecurityGroupEgressRuleRequest) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *ModifySecurityGroupEgressRuleRequest) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *ModifySecurityGroupEgressRuleRequest) SetClientToken(v string) *ModifySecurityGroupEgressRuleRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetDescription(v string) *ModifySecurityGroupEgressRuleRequest {
	s.Description = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetDestCidrIp(v string) *ModifySecurityGroupEgressRuleRequest {
	s.DestCidrIp = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetDestGroupId(v string) *ModifySecurityGroupEgressRuleRequest {
	s.DestGroupId = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetDestGroupOwnerAccount(v string) *ModifySecurityGroupEgressRuleRequest {
	s.DestGroupOwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetDestGroupOwnerId(v int64) *ModifySecurityGroupEgressRuleRequest {
	s.DestGroupOwnerId = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetDestPrefixListId(v string) *ModifySecurityGroupEgressRuleRequest {
	s.DestPrefixListId = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetIpProtocol(v string) *ModifySecurityGroupEgressRuleRequest {
	s.IpProtocol = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetIpv6DestCidrIp(v string) *ModifySecurityGroupEgressRuleRequest {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetIpv6SourceCidrIp(v string) *ModifySecurityGroupEgressRuleRequest {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetNicType(v string) *ModifySecurityGroupEgressRuleRequest {
	s.NicType = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetOwnerAccount(v string) *ModifySecurityGroupEgressRuleRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetOwnerId(v int64) *ModifySecurityGroupEgressRuleRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetPolicy(v string) *ModifySecurityGroupEgressRuleRequest {
	s.Policy = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetPortRange(v string) *ModifySecurityGroupEgressRuleRequest {
	s.PortRange = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetPortRangeListId(v string) *ModifySecurityGroupEgressRuleRequest {
	s.PortRangeListId = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetPriority(v string) *ModifySecurityGroupEgressRuleRequest {
	s.Priority = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetRegionId(v string) *ModifySecurityGroupEgressRuleRequest {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetResourceOwnerAccount(v string) *ModifySecurityGroupEgressRuleRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetResourceOwnerId(v int64) *ModifySecurityGroupEgressRuleRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetSecurityGroupId(v string) *ModifySecurityGroupEgressRuleRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetSecurityGroupRuleId(v string) *ModifySecurityGroupEgressRuleRequest {
	s.SecurityGroupRuleId = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetSourceCidrIp(v string) *ModifySecurityGroupEgressRuleRequest {
	s.SourceCidrIp = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) SetSourcePortRange(v string) *ModifySecurityGroupEgressRuleRequest {
	s.SourcePortRange = &v
	return s
}

func (s *ModifySecurityGroupEgressRuleRequest) Validate() error {
	return dara.Validate(s)
}
