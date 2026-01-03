// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityGroupAttributeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *DescribeSecurityGroupAttributeResponseBody
	GetDescription() *string
	SetInnerAccessPolicy(v string) *DescribeSecurityGroupAttributeResponseBody
	GetInnerAccessPolicy() *string
	SetNextToken(v string) *DescribeSecurityGroupAttributeResponseBody
	GetNextToken() *string
	SetPermissions(v *DescribeSecurityGroupAttributeResponseBodyPermissions) *DescribeSecurityGroupAttributeResponseBody
	GetPermissions() *DescribeSecurityGroupAttributeResponseBodyPermissions
	SetRegionId(v string) *DescribeSecurityGroupAttributeResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeSecurityGroupAttributeResponseBody
	GetRequestId() *string
	SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeResponseBody
	GetSecurityGroupId() *string
	SetSecurityGroupName(v string) *DescribeSecurityGroupAttributeResponseBody
	GetSecurityGroupName() *string
	SetSnapshotPolicyIds(v *DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds) *DescribeSecurityGroupAttributeResponseBody
	GetSnapshotPolicyIds() *DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds
	SetVpcId(v string) *DescribeSecurityGroupAttributeResponseBody
	GetVpcId() *string
}

type DescribeSecurityGroupAttributeResponseBody struct {
	// The description of the security group.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The access control policy of the security group. Valid values:
	//
	// 	- Accept: All instances in the security group can communicate with each other.
	//
	// 	- Drop: All instances in the security group are isolated from each other.
	//
	// example:
	//
	// Accept
	InnerAccessPolicy *string `json:"InnerAccessPolicy,omitempty" xml:"InnerAccessPolicy,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If the return value of this parameter is empty when you specify `MaxResults` and `NextToken` for a paged query, no more results are to be returned.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Details about the security group rules.
	Permissions *DescribeSecurityGroupAttributeResponseBodyPermissions `json:"Permissions,omitempty" xml:"Permissions,omitempty" type:"Struct"`
	// The ID of the region.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the security group.
	//
	// example:
	//
	// sg-bp1gxw6bznjjvhu3****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The name of the security group.
	//
	// example:
	//
	// SecurityGroupName Sample
	SecurityGroupName *string                                                      `json:"SecurityGroupName,omitempty" xml:"SecurityGroupName,omitempty"`
	SnapshotPolicyIds *DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds `json:"SnapshotPolicyIds,omitempty" xml:"SnapshotPolicyIds,omitempty" type:"Struct"`
	// The ID of the VPC. If a VPC ID is returned, the network type of the security group is VPC. If no VPC ID is returned, the network type of the security group is classic network.
	//
	// example:
	//
	// vpc-bp1opxu1zkhn00gzv****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeSecurityGroupAttributeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetInnerAccessPolicy() *string {
	return s.InnerAccessPolicy
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetPermissions() *DescribeSecurityGroupAttributeResponseBodyPermissions {
	return s.Permissions
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetSecurityGroupName() *string {
	return s.SecurityGroupName
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetSnapshotPolicyIds() *DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds {
	return s.SnapshotPolicyIds
}

func (s *DescribeSecurityGroupAttributeResponseBody) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetDescription(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.Description = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetInnerAccessPolicy(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.InnerAccessPolicy = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetNextToken(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetPermissions(v *DescribeSecurityGroupAttributeResponseBodyPermissions) *DescribeSecurityGroupAttributeResponseBody {
	s.Permissions = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetRegionId(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetRequestId(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetSecurityGroupId(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.SecurityGroupId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetSecurityGroupName(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.SecurityGroupName = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetSnapshotPolicyIds(v *DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds) *DescribeSecurityGroupAttributeResponseBody {
	s.SnapshotPolicyIds = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) SetVpcId(v string) *DescribeSecurityGroupAttributeResponseBody {
	s.VpcId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBody) Validate() error {
	if s.Permissions != nil {
		if err := s.Permissions.Validate(); err != nil {
			return err
		}
	}
	if s.SnapshotPolicyIds != nil {
		if err := s.SnapshotPolicyIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityGroupAttributeResponseBodyPermissions struct {
	Permission []*DescribeSecurityGroupAttributeResponseBodyPermissionsPermission `json:"Permission,omitempty" xml:"Permission,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissions) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissions) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissions) GetPermission() []*DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	return s.Permission
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissions) SetPermission(v []*DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) *DescribeSecurityGroupAttributeResponseBodyPermissions {
	s.Permission = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissions) Validate() error {
	if s.Permission != nil {
		for _, item := range s.Permission {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSecurityGroupAttributeResponseBodyPermissionsPermission struct {
	// The time when the security group rule was created. The time is displayed in UTC.
	//
	// example:
	//
	// 2018-12-12T07:28:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the security group.
	//
	// example:
	//
	// Description Sample 01
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The destination CIDR block for outbound access control.
	//
	// example:
	//
	// 0.0.0.0/0
	DestCidrIp *string `json:"DestCidrIp,omitempty" xml:"DestCidrIp,omitempty"`
	// The ID of the destination security group for outbound access control.
	//
	// example:
	//
	// sg-bp1czdx84jd88i7v****
	DestGroupId *string `json:"DestGroupId,omitempty" xml:"DestGroupId,omitempty"`
	// The name of the destination security group.
	//
	// example:
	//
	// testDestGroupName
	DestGroupName *string `json:"DestGroupName,omitempty" xml:"DestGroupName,omitempty"`
	// The ID of the Alibaba Cloud account to which the destination security group belongs.
	//
	// example:
	//
	// 1234567890
	DestGroupOwnerAccount *string `json:"DestGroupOwnerAccount,omitempty" xml:"DestGroupOwnerAccount,omitempty"`
	// The ID of the destination prefix list for outbound access control.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixabc****
	DestPrefixListId *string `json:"DestPrefixListId,omitempty" xml:"DestPrefixListId,omitempty"`
	// The name of the destination prefix list.
	//
	// example:
	//
	// DestPrefixListName Sample
	DestPrefixListName *string `json:"DestPrefixListName,omitempty" xml:"DestPrefixListName,omitempty"`
	// The direction in which the security group rule is applied.
	//
	// example:
	//
	// ingress
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	// The transport layer protocol.
	//
	// example:
	//
	// TCP
	IpProtocol *string `json:"IpProtocol,omitempty" xml:"IpProtocol,omitempty"`
	// The destination IPv6 CIDR block.
	//
	// example:
	//
	// 2001:db8:1233:1a00::***
	Ipv6DestCidrIp *string `json:"Ipv6DestCidrIp,omitempty" xml:"Ipv6DestCidrIp,omitempty"`
	// The source IPv6 CIDR block.
	//
	// example:
	//
	// 2001:db8:1234:1a00::***
	Ipv6SourceCidrIp *string `json:"Ipv6SourceCidrIp,omitempty" xml:"Ipv6SourceCidrIp,omitempty"`
	// The network type.
	//
	// example:
	//
	// intranet
	NicType *string `json:"NicType,omitempty" xml:"NicType,omitempty"`
	// The access control policy.
	//
	// example:
	//
	// Accept
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The port range.
	//
	// example:
	//
	// 80/80
	PortRange *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	// The ID of the port list.
	//
	// example:
	//
	// prl-2ze9743****
	PortRangeListId *string `json:"PortRangeListId,omitempty" xml:"PortRangeListId,omitempty"`
	// The name of the port list.
	//
	// example:
	//
	// PortRangeListNameSample
	PortRangeListName *string `json:"PortRangeListName,omitempty" xml:"PortRangeListName,omitempty"`
	// The priority of the rule.
	//
	// example:
	//
	// 1
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// The ID of the security group rule.
	//
	// example:
	//
	// sgr-bp12kewq32dfwrdi****
	SecurityGroupRuleId *string `json:"SecurityGroupRuleId,omitempty" xml:"SecurityGroupRuleId,omitempty"`
	// The source CIDR block for inbound access control.
	//
	// example:
	//
	// 0.0.0.0/0
	SourceCidrIp *string `json:"SourceCidrIp,omitempty" xml:"SourceCidrIp,omitempty"`
	// The source security group for inbound access control.
	//
	// example:
	//
	// sg-bp12kc4rqohaf2js****
	SourceGroupId *string `json:"SourceGroupId,omitempty" xml:"SourceGroupId,omitempty"`
	// The name of the source security group.
	//
	// example:
	//
	// testSourceGroupName1
	SourceGroupName *string `json:"SourceGroupName,omitempty" xml:"SourceGroupName,omitempty"`
	// The ID of the Alibaba Cloud account to which the source security group belongs.
	//
	// example:
	//
	// 1234567890
	SourceGroupOwnerAccount *string `json:"SourceGroupOwnerAccount,omitempty" xml:"SourceGroupOwnerAccount,omitempty"`
	// The source port range.
	//
	// example:
	//
	// 80/80
	SourcePortRange *string `json:"SourcePortRange,omitempty" xml:"SourcePortRange,omitempty"`
	// The ID of the source prefix list for inbound access control.
	//
	// example:
	//
	// pl-x1j1k5ykzqlixdcy****
	SourcePrefixListId *string `json:"SourcePrefixListId,omitempty" xml:"SourcePrefixListId,omitempty"`
	// The name of the source prefix list.
	//
	// example:
	//
	// SourcePrefixListName Sample
	SourcePrefixListName *string `json:"SourcePrefixListName,omitempty" xml:"SourcePrefixListName,omitempty"`
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDescription() *string {
	return s.Description
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDestCidrIp() *string {
	return s.DestCidrIp
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDestGroupId() *string {
	return s.DestGroupId
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDestGroupName() *string {
	return s.DestGroupName
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDestGroupOwnerAccount() *string {
	return s.DestGroupOwnerAccount
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDestPrefixListId() *string {
	return s.DestPrefixListId
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDestPrefixListName() *string {
	return s.DestPrefixListName
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetDirection() *string {
	return s.Direction
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetIpProtocol() *string {
	return s.IpProtocol
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetIpv6DestCidrIp() *string {
	return s.Ipv6DestCidrIp
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetIpv6SourceCidrIp() *string {
	return s.Ipv6SourceCidrIp
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetNicType() *string {
	return s.NicType
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetPortRange() *string {
	return s.PortRange
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetPortRangeListId() *string {
	return s.PortRangeListId
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetPortRangeListName() *string {
	return s.PortRangeListName
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetPriority() *string {
	return s.Priority
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetSecurityGroupRuleId() *string {
	return s.SecurityGroupRuleId
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetSourceCidrIp() *string {
	return s.SourceCidrIp
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetSourceGroupId() *string {
	return s.SourceGroupId
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetSourceGroupName() *string {
	return s.SourceGroupName
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetSourceGroupOwnerAccount() *string {
	return s.SourceGroupOwnerAccount
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetSourcePortRange() *string {
	return s.SourcePortRange
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetSourcePrefixListId() *string {
	return s.SourcePrefixListId
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) GetSourcePrefixListName() *string {
	return s.SourcePrefixListName
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetCreateTime(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.CreateTime = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDescription(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Description = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDestCidrIp(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.DestCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDestGroupId(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.DestGroupId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDestGroupName(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.DestGroupName = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDestGroupOwnerAccount(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.DestGroupOwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDestPrefixListId(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.DestPrefixListId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDestPrefixListName(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.DestPrefixListName = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetDirection(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Direction = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetIpProtocol(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.IpProtocol = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetIpv6DestCidrIp(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Ipv6DestCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetIpv6SourceCidrIp(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Ipv6SourceCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetNicType(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.NicType = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPolicy(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Policy = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPortRange(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.PortRange = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPortRangeListId(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.PortRangeListId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPortRangeListName(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.PortRangeListName = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetPriority(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.Priority = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSecurityGroupRuleId(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SecurityGroupRuleId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourceCidrIp(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourceCidrIp = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourceGroupId(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourceGroupId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourceGroupName(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourceGroupName = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourceGroupOwnerAccount(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourceGroupOwnerAccount = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourcePortRange(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourcePortRange = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourcePrefixListId(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourcePrefixListId = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) SetSourcePrefixListName(v string) *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission {
	s.SourcePrefixListName = &v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodyPermissionsPermission) Validate() error {
	return dara.Validate(s)
}

type DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds struct {
	SnapshotPolicyId []*string `json:"SnapshotPolicyId,omitempty" xml:"SnapshotPolicyId,omitempty" type:"Repeated"`
}

func (s DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds) GoString() string {
	return s.String()
}

func (s *DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds) GetSnapshotPolicyId() []*string {
	return s.SnapshotPolicyId
}

func (s *DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds) SetSnapshotPolicyId(v []*string) *DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds {
	s.SnapshotPolicyId = v
	return s
}

func (s *DescribeSecurityGroupAttributeResponseBodySnapshotPolicyIds) Validate() error {
	return dara.Validate(s)
}
