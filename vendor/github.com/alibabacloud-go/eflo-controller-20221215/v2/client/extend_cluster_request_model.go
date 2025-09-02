// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iExtendClusterRequest interface {
  dara.Model
  String() string
  GoString() string
  SetClusterId(v string) *ExtendClusterRequest
  GetClusterId() *string 
  SetIgnoreFailedNodeTasks(v bool) *ExtendClusterRequest
  GetIgnoreFailedNodeTasks() *bool 
  SetIpAllocationPolicy(v []*ExtendClusterRequestIpAllocationPolicy) *ExtendClusterRequest
  GetIpAllocationPolicy() []*ExtendClusterRequestIpAllocationPolicy 
  SetNodeGroups(v []*ExtendClusterRequestNodeGroups) *ExtendClusterRequest
  GetNodeGroups() []*ExtendClusterRequestNodeGroups 
  SetVSwitchZoneId(v string) *ExtendClusterRequest
  GetVSwitchZoneId() *string 
  SetVpdSubnets(v []*string) *ExtendClusterRequest
  GetVpdSubnets() []*string 
}

type ExtendClusterRequest struct {
  // Cluster ID
  // 
  // example:
  // 
  // i15b480fbd2fcdbc2869cd80
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // Whether to allow skipping failed node tasks, default value is False
  // 
  // example:
  // 
  // False
  IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
  // IP allocation combination policy: Each policy can only choose one type, and multiple policies can be combined
  IpAllocationPolicy []*ExtendClusterRequestIpAllocationPolicy `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty" type:"Repeated"`
  // Node Groups
  NodeGroups []*ExtendClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
  // VSwitch availability zone ID
  // 
  // example:
  // 
  // cn-shanghai-b
  VSwitchZoneId *string `json:"VSwitchZoneId,omitempty" xml:"VSwitchZoneId,omitempty"`
  // List of cluster subnets
  VpdSubnets []*string `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty" type:"Repeated"`
}

func (s ExtendClusterRequest) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequest) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequest) GetClusterId() *string  {
  return s.ClusterId
}

func (s *ExtendClusterRequest) GetIgnoreFailedNodeTasks() *bool  {
  return s.IgnoreFailedNodeTasks
}

func (s *ExtendClusterRequest) GetIpAllocationPolicy() []*ExtendClusterRequestIpAllocationPolicy  {
  return s.IpAllocationPolicy
}

func (s *ExtendClusterRequest) GetNodeGroups() []*ExtendClusterRequestNodeGroups  {
  return s.NodeGroups
}

func (s *ExtendClusterRequest) GetVSwitchZoneId() *string  {
  return s.VSwitchZoneId
}

func (s *ExtendClusterRequest) GetVpdSubnets() []*string  {
  return s.VpdSubnets
}

func (s *ExtendClusterRequest) SetClusterId(v string) *ExtendClusterRequest {
  s.ClusterId = &v
  return s
}

func (s *ExtendClusterRequest) SetIgnoreFailedNodeTasks(v bool) *ExtendClusterRequest {
  s.IgnoreFailedNodeTasks = &v
  return s
}

func (s *ExtendClusterRequest) SetIpAllocationPolicy(v []*ExtendClusterRequestIpAllocationPolicy) *ExtendClusterRequest {
  s.IpAllocationPolicy = v
  return s
}

func (s *ExtendClusterRequest) SetNodeGroups(v []*ExtendClusterRequestNodeGroups) *ExtendClusterRequest {
  s.NodeGroups = v
  return s
}

func (s *ExtendClusterRequest) SetVSwitchZoneId(v string) *ExtendClusterRequest {
  s.VSwitchZoneId = &v
  return s
}

func (s *ExtendClusterRequest) SetVpdSubnets(v []*string) *ExtendClusterRequest {
  s.VpdSubnets = v
  return s
}

func (s *ExtendClusterRequest) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestIpAllocationPolicy struct {
  // Specify the cluster subnet ID based on the bond name
  BondPolicy *ExtendClusterRequestIpAllocationPolicyBondPolicy `json:"BondPolicy,omitempty" xml:"BondPolicy,omitempty" type:"Struct"`
  // Machine type allocation policy
  MachineTypePolicy []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicy `json:"MachineTypePolicy,omitempty" xml:"MachineTypePolicy,omitempty" type:"Repeated"`
  // Node allocation policy
  NodePolicy []*ExtendClusterRequestIpAllocationPolicyNodePolicy `json:"NodePolicy,omitempty" xml:"NodePolicy,omitempty" type:"Repeated"`
}

func (s ExtendClusterRequestIpAllocationPolicy) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicy) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicy) GetBondPolicy() *ExtendClusterRequestIpAllocationPolicyBondPolicy  {
  return s.BondPolicy
}

func (s *ExtendClusterRequestIpAllocationPolicy) GetMachineTypePolicy() []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicy  {
  return s.MachineTypePolicy
}

func (s *ExtendClusterRequestIpAllocationPolicy) GetNodePolicy() []*ExtendClusterRequestIpAllocationPolicyNodePolicy  {
  return s.NodePolicy
}

func (s *ExtendClusterRequestIpAllocationPolicy) SetBondPolicy(v *ExtendClusterRequestIpAllocationPolicyBondPolicy) *ExtendClusterRequestIpAllocationPolicy {
  s.BondPolicy = v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicy) SetMachineTypePolicy(v []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) *ExtendClusterRequestIpAllocationPolicy {
  s.MachineTypePolicy = v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicy) SetNodePolicy(v []*ExtendClusterRequestIpAllocationPolicyNodePolicy) *ExtendClusterRequestIpAllocationPolicy {
  s.NodePolicy = v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicy) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestIpAllocationPolicyBondPolicy struct {
  // Default bond cluster subnet
  // 
  // example:
  // 
  // subnet-3od2fe
  BondDefaultSubnet *string `json:"BondDefaultSubnet,omitempty" xml:"BondDefaultSubnet,omitempty"`
  // Bond information
  Bonds []*ExtendClusterRequestIpAllocationPolicyBondPolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
}

func (s ExtendClusterRequestIpAllocationPolicyBondPolicy) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyBondPolicy) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) GetBondDefaultSubnet() *string  {
  return s.BondDefaultSubnet
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) GetBonds() []*ExtendClusterRequestIpAllocationPolicyBondPolicyBonds  {
  return s.Bonds
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) SetBondDefaultSubnet(v string) *ExtendClusterRequestIpAllocationPolicyBondPolicy {
  s.BondDefaultSubnet = &v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) SetBonds(v []*ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) *ExtendClusterRequestIpAllocationPolicyBondPolicy {
  s.Bonds = v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestIpAllocationPolicyBondPolicyBonds struct {
  // Bond name
  // 
  // example:
  // 
  // Bond0
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // IP source cluster subnet
  // 
  // example:
  // 
  // subnet-3od2fe
  Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) GetName() *string  {
  return s.Name
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) GetSubnet() *string  {
  return s.Subnet
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) SetName(v string) *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds {
  s.Name = &v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) SetSubnet(v string) *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds {
  s.Subnet = &v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestIpAllocationPolicyMachineTypePolicy struct {
  // Bond information
  Bonds []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
  // Machine type
  // 
  // example:
  // 
  // efg1.nvga1
  MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) GetBonds() []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds  {
  return s.Bonds
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) GetMachineType() *string  {
  return s.MachineType
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) SetBonds(v []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy {
  s.Bonds = v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) SetMachineType(v string) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy {
  s.MachineType = &v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds struct {
  // Bond name
  // 
  // example:
  // 
  // Bond0
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // IP source cluster subnet
  // 
  // example:
  // 
  // subnet-fdo3dv
  Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) GetName() *string  {
  return s.Name
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) GetSubnet() *string  {
  return s.Subnet
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) SetName(v string) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds {
  s.Name = &v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) SetSubnet(v string) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds {
  s.Subnet = &v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestIpAllocationPolicyNodePolicy struct {
  // Bond information
  Bonds []*ExtendClusterRequestIpAllocationPolicyNodePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
  // Hostname
  // 
  // example:
  // 
  // a100-xa5dza28-0085
  Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
  // Node ID
  // 
  // example:
  // 
  // i-3fdodw2
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicy) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicy) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) GetBonds() []*ExtendClusterRequestIpAllocationPolicyNodePolicyBonds  {
  return s.Bonds
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) GetHostname() *string  {
  return s.Hostname
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) GetNodeId() *string  {
  return s.NodeId
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) SetBonds(v []*ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) *ExtendClusterRequestIpAllocationPolicyNodePolicy {
  s.Bonds = v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) SetHostname(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicy {
  s.Hostname = &v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) SetNodeId(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicy {
  s.NodeId = &v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestIpAllocationPolicyNodePolicyBonds struct {
  // Bond name
  // 
  // example:
  // 
  // Bond0
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // IP source cluster subnet
  // 
  // example:
  // 
  // subnet-fdo3dv
  Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) GetName() *string  {
  return s.Name
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) GetSubnet() *string  {
  return s.Subnet
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) SetName(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds {
  s.Name = &v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) SetSubnet(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds {
  s.Subnet = &v
  return s
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestNodeGroups struct {
  // Number of nodes to purchase. Range: 0~500. If the Amount parameter is set to 0, it means no new nodes will be purchased and existing nodes will be used for scaling. If the Amount parameter is set to 1~500, it means a certain number of nodes will be purchased and used for scaling. Default value: 0
  // 
  // example:
  // 
  // 4
  Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
  // Whether to automatically renew the purchased nodes. This parameter takes effect when the Amount parameter is not 0 and the ChargeType is set to PrePaid. Valid values: True (auto-renewal); False (no auto-renewal). Default value: False
  // 
  // example:
  // 
  // True
  AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
  // Payment method for the nodes. When the Amount parameter is set to 0, this parameter does not take effect. Valid values: PrePaid (Subscription); PostPaid (Pay-As-You-Go). Default value: PrePaid.
  // 
  // example:
  // 
  // PostPaid
  ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
  // Set the hostnames for the purchased nodes. This parameter does not take effect when the Amount parameter is set to 0.
  Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
  // Set the login password for the purchased nodes. This parameter is not effective when the Amount parameter is set to 0.
  // 
  // example:
  // 
  // skkO(*89Y
  LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
  // Node Group ID
  // 
  // example:
  // 
  // i16d4883a46cbadeb4bc9
  NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
  // Node tags
  NodeTag []*ExtendClusterRequestNodeGroupsNodeTag `json:"NodeTag,omitempty" xml:"NodeTag,omitempty" type:"Repeated"`
  // List of Nodes
  Nodes []*ExtendClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
  // Duration of the node purchase (in months). Valid values: 1, 6, 12, 24, 36, 48. This parameter takes effect when the Amount parameter is not 0 and the ChargeType is set to PrePaid.
  // 
  // example:
  // 
  // 6
  Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
  // Custom Data
  // 
  // example:
  // 
  // #!/bin/sh
  // 
  // echo "Hello World. The time is now $(date -R)!" | tee /root/userdata_test.txt
  UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
  // VSwitch ID
  // 
  // example:
  // 
  // vsw-uf65m8xqjgy55xj9jw92n
  VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
  // VPC ID
  // 
  // example:
  // 
  // vpc-0jl3b0c0ukydlfezr13n6
  VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
  // Zone ID
  // 
  // example:
  // 
  // cn-hangzhou-i
  ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ExtendClusterRequestNodeGroups) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestNodeGroups) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestNodeGroups) GetAmount() *int64  {
  return s.Amount
}

func (s *ExtendClusterRequestNodeGroups) GetAutoRenew() *bool  {
  return s.AutoRenew
}

func (s *ExtendClusterRequestNodeGroups) GetChargeType() *string  {
  return s.ChargeType
}

func (s *ExtendClusterRequestNodeGroups) GetHostnames() []*string  {
  return s.Hostnames
}

func (s *ExtendClusterRequestNodeGroups) GetLoginPassword() *string  {
  return s.LoginPassword
}

func (s *ExtendClusterRequestNodeGroups) GetNodeGroupId() *string  {
  return s.NodeGroupId
}

func (s *ExtendClusterRequestNodeGroups) GetNodeTag() []*ExtendClusterRequestNodeGroupsNodeTag  {
  return s.NodeTag
}

func (s *ExtendClusterRequestNodeGroups) GetNodes() []*ExtendClusterRequestNodeGroupsNodes  {
  return s.Nodes
}

func (s *ExtendClusterRequestNodeGroups) GetPeriod() *int64  {
  return s.Period
}

func (s *ExtendClusterRequestNodeGroups) GetUserData() *string  {
  return s.UserData
}

func (s *ExtendClusterRequestNodeGroups) GetVSwitchId() *string  {
  return s.VSwitchId
}

func (s *ExtendClusterRequestNodeGroups) GetVpcId() *string  {
  return s.VpcId
}

func (s *ExtendClusterRequestNodeGroups) GetZoneId() *string  {
  return s.ZoneId
}

func (s *ExtendClusterRequestNodeGroups) SetAmount(v int64) *ExtendClusterRequestNodeGroups {
  s.Amount = &v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetAutoRenew(v bool) *ExtendClusterRequestNodeGroups {
  s.AutoRenew = &v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetChargeType(v string) *ExtendClusterRequestNodeGroups {
  s.ChargeType = &v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetHostnames(v []*string) *ExtendClusterRequestNodeGroups {
  s.Hostnames = v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetLoginPassword(v string) *ExtendClusterRequestNodeGroups {
  s.LoginPassword = &v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetNodeGroupId(v string) *ExtendClusterRequestNodeGroups {
  s.NodeGroupId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetNodeTag(v []*ExtendClusterRequestNodeGroupsNodeTag) *ExtendClusterRequestNodeGroups {
  s.NodeTag = v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetNodes(v []*ExtendClusterRequestNodeGroupsNodes) *ExtendClusterRequestNodeGroups {
  s.Nodes = v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetPeriod(v int64) *ExtendClusterRequestNodeGroups {
  s.Period = &v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetUserData(v string) *ExtendClusterRequestNodeGroups {
  s.UserData = &v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetVSwitchId(v string) *ExtendClusterRequestNodeGroups {
  s.VSwitchId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetVpcId(v string) *ExtendClusterRequestNodeGroups {
  s.VpcId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroups) SetZoneId(v string) *ExtendClusterRequestNodeGroups {
  s.ZoneId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroups) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestNodeGroupsNodeTag struct {
  // Node tag key
  // 
  // example:
  // 
  // key_my
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // Node tag value
  // 
  // example:
  // 
  // value_my
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsNodeTag) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsNodeTag) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) GetKey() *string  {
  return s.Key
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) GetValue() *string  {
  return s.Value
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) SetKey(v string) *ExtendClusterRequestNodeGroupsNodeTag {
  s.Key = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) SetValue(v string) *ExtendClusterRequestNodeGroupsNodeTag {
  s.Value = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestNodeGroupsNodes struct {
  // Data Disk Specifications
  DataDisk []*ExtendClusterRequestNodeGroupsNodesDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
  // Hostname
  // 
  // example:
  // 
  // d044d220-33fd-11ed-86a6
  Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
  // Login Password
  // 
  // example:
  // 
  // ***
  LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
  // Node ID
  // 
  // example:
  // 
  // e01-cn-zvp2zdpy601
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
  // VSwitch ID
  // 
  // example:
  // 
  // vsw-bp169pi5fj151rrms4sia
  VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
  // VPC ID
  // 
  // example:
  // 
  // vpc-0jlasms92fdxqd3wlf8ny
  VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsNodes) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsNodes) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestNodeGroupsNodes) GetDataDisk() []*ExtendClusterRequestNodeGroupsNodesDataDisk  {
  return s.DataDisk
}

func (s *ExtendClusterRequestNodeGroupsNodes) GetHostname() *string  {
  return s.Hostname
}

func (s *ExtendClusterRequestNodeGroupsNodes) GetLoginPassword() *string  {
  return s.LoginPassword
}

func (s *ExtendClusterRequestNodeGroupsNodes) GetNodeId() *string  {
  return s.NodeId
}

func (s *ExtendClusterRequestNodeGroupsNodes) GetVSwitchId() *string  {
  return s.VSwitchId
}

func (s *ExtendClusterRequestNodeGroupsNodes) GetVpcId() *string  {
  return s.VpcId
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetDataDisk(v []*ExtendClusterRequestNodeGroupsNodesDataDisk) *ExtendClusterRequestNodeGroupsNodes {
  s.DataDisk = v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetHostname(v string) *ExtendClusterRequestNodeGroupsNodes {
  s.Hostname = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetLoginPassword(v string) *ExtendClusterRequestNodeGroupsNodes {
  s.LoginPassword = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetNodeId(v string) *ExtendClusterRequestNodeGroupsNodes {
  s.NodeId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetVSwitchId(v string) *ExtendClusterRequestNodeGroupsNodes {
  s.VSwitchId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodes) SetVpcId(v string) *ExtendClusterRequestNodeGroupsNodes {
  s.VpcId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodes) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestNodeGroupsNodesDataDisk struct {
  // Type
  // 
  // example:
  // 
  // cloud_essd
  Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
  // Whether the data disk is deleted with the node
  // 
  // example:
  // 
  // true
  DeleteWithNode *bool `json:"DeleteWithNode,omitempty" xml:"DeleteWithNode,omitempty"`
  // Data Disk Performance Level
  // 
  // example:
  // 
  // PL0
  PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
  // Disk Size
  // 
  // example:
  // 
  // 80
  Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsNodesDataDisk) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsNodesDataDisk) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) GetCategory() *string  {
  return s.Category
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) GetDeleteWithNode() *bool  {
  return s.DeleteWithNode
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) GetPerformanceLevel() *string  {
  return s.PerformanceLevel
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) GetSize() *int32  {
  return s.Size
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetCategory(v string) *ExtendClusterRequestNodeGroupsNodesDataDisk {
  s.Category = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetDeleteWithNode(v bool) *ExtendClusterRequestNodeGroupsNodesDataDisk {
  s.DeleteWithNode = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetPerformanceLevel(v string) *ExtendClusterRequestNodeGroupsNodesDataDisk {
  s.PerformanceLevel = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetSize(v int32) *ExtendClusterRequestNodeGroupsNodesDataDisk {
  s.Size = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) Validate() error {
  return dara.Validate(s)
}

