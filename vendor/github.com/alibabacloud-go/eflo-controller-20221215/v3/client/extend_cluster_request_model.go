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
  // The cluster ID.
  // 
  // example:
  // 
  // i15b480fbd2fcdbc2869cd80
  ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
  // Specifies whether to skip failed nodes. Default value: False.
  // 
  // example:
  // 
  // False
  IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
  // The combined IP allocation policy. Each policy can use only one policy type, and multiple policies can be combined.
  IpAllocationPolicy []*ExtendClusterRequestIpAllocationPolicy `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty" type:"Repeated"`
  // The node groups.
  NodeGroups []*ExtendClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
  // The zone ID of the vSwitch.
  // 
  // example:
  // 
  // cn-shanghai-b
  VSwitchZoneId *string `json:"VSwitchZoneId,omitempty" xml:"VSwitchZoneId,omitempty"`
  // The list of cluster subnets.
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
  if s.IpAllocationPolicy != nil {
    for _, item := range s.IpAllocationPolicy {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.NodeGroups != nil {
    for _, item := range s.NodeGroups {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExtendClusterRequestIpAllocationPolicy struct {
  // Specifies the cluster subnet ID based on the bond name.
  BondPolicy *ExtendClusterRequestIpAllocationPolicyBondPolicy `json:"BondPolicy,omitempty" xml:"BondPolicy,omitempty" type:"Struct"`
  // The machine type allocation policy.
  MachineTypePolicy []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicy `json:"MachineTypePolicy,omitempty" xml:"MachineTypePolicy,omitempty" type:"Repeated"`
  // The node allocation policy.
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
  if s.BondPolicy != nil {
    if err := s.BondPolicy.Validate(); err != nil {
      return err
    }
  }
  if s.MachineTypePolicy != nil {
    for _, item := range s.MachineTypePolicy {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.NodePolicy != nil {
    for _, item := range s.NodePolicy {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExtendClusterRequestIpAllocationPolicyBondPolicy struct {
  // The default bond cluster subnet.
  // 
  // example:
  // 
  // subnet-3od2fe
  BondDefaultSubnet *string `json:"BondDefaultSubnet,omitempty" xml:"BondDefaultSubnet,omitempty"`
  // The bond information.
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
  if s.Bonds != nil {
    for _, item := range s.Bonds {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExtendClusterRequestIpAllocationPolicyBondPolicyBonds struct {
  // The bond name.
  // 
  // example:
  // 
  // Bond0
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The cluster subnet from which the IP address is allocated.
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
  // The bond information.
  Bonds []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
  // The machine type.
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
  if s.Bonds != nil {
    for _, item := range s.Bonds {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds struct {
  // The bond name.
  // 
  // example:
  // 
  // Bond0
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The cluster subnet from which the IP address is allocated.
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
  // The bond information.
  Bonds []*ExtendClusterRequestIpAllocationPolicyNodePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
  // The hostname.
  // 
  // example:
  // 
  // a100-xa5dza28-0085
  Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
  // The node ID.
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
  if s.Bonds != nil {
    for _, item := range s.Bonds {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExtendClusterRequestIpAllocationPolicyNodePolicyBonds struct {
  // The bond name.
  // 
  // example:
  // 
  // Bond0
  Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
  // The cluster subnet from which the IP address is allocated.
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
  // The number of nodes to purchase. Valid values: 0 to 500. If Amount is set to 0, no nodes are purchased and existing nodes are used for scale-out. If Amount is set to a value from 1 to 500, the specified number of nodes are purchased and used for scale-out. Default value: 0
  // 
  // example:
  // 
  // 4
  Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
  // Specifies whether to enable auto-renewal for the purchased nodes. This parameter takes effect when Amount is not 0 and ChargeType is set to PREPAY or POSTPAY. Valid values: True: Enable auto-renewal. False: Disable auto-renewal. Default value: False.
  // 
  // example:
  // 
  // True
  AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
  // The billing method of the nodes. This parameter does not take effect when Amount is set to 0. Valid values: PREPAY: subscription. POSTPAY: pay-as-you-go. Default value: PREPAY.
  // 
  // example:
  // 
  // PostPaid
  ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
  // The hostnames of the purchased nodes. This parameter does not take effect when Amount is set to 0.
  Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
  // The list of hyper nodes.
  HyperNodes []*ExtendClusterRequestNodeGroupsHyperNodes `json:"HyperNodes,omitempty" xml:"HyperNodes,omitempty" type:"Repeated"`
  // The logon password of the purchased nodes. This parameter does not take effect when Amount is set to 0.
  // 
  // example:
  // 
  // skkO(*89Y
  LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
  // The node group ID.
  // 
  // example:
  // 
  // i16d4883a46cbadeb4bc9
  NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
  // The node tags.
  NodeTag []*ExtendClusterRequestNodeGroupsNodeTag `json:"NodeTag,omitempty" xml:"NodeTag,omitempty" type:"Repeated"`
  // The node list.
  Nodes []*ExtendClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
  // The subscription duration of the purchased nodes. Unit: months. Valid values: 1, 6, 12, 24, 36, and 48. This parameter takes effect when Amount is not 0 and ChargeType is set to PREPAY.
  // 
  // example:
  // 
  // 6
  Period *int64 `json:"Period,omitempty" xml:"Period,omitempty"`
  // The savings plan ID.
  // 
  // example:
  // 
  // spn-25e985acAWbrwEBJ
  SavingsPlanId *string `json:"SavingsPlanId,omitempty" xml:"SavingsPlanId,omitempty"`
  // The custom executable shell script. The script must be Base64-encoded. The maximum size of the raw data is 16 KB.
  // 
  // example:
  // 
  // ZWNobyBoZWxsbyBlY3Mh
  UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
  // The vSwitch ID.
  // 
  // example:
  // 
  // vsw-uf65m8xqjgy55xj9jw92n
  VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
  // The VPC ID.
  // 
  // example:
  // 
  // vpc-0jl3b0c0ukydlfezr13n6
  VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
  // The zone ID.
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

func (s *ExtendClusterRequestNodeGroups) GetHyperNodes() []*ExtendClusterRequestNodeGroupsHyperNodes  {
  return s.HyperNodes
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

func (s *ExtendClusterRequestNodeGroups) GetSavingsPlanId() *string  {
  return s.SavingsPlanId
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

func (s *ExtendClusterRequestNodeGroups) SetHyperNodes(v []*ExtendClusterRequestNodeGroupsHyperNodes) *ExtendClusterRequestNodeGroups {
  s.HyperNodes = v
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

func (s *ExtendClusterRequestNodeGroups) SetSavingsPlanId(v string) *ExtendClusterRequestNodeGroups {
  s.SavingsPlanId = &v
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
  if s.HyperNodes != nil {
    for _, item := range s.HyperNodes {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.NodeTag != nil {
    for _, item := range s.NodeTag {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  if s.Nodes != nil {
    for _, item := range s.Nodes {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExtendClusterRequestNodeGroupsHyperNodes struct {
  // The list of cloud disk information.
  DataDisk []*ExtendClusterRequestNodeGroupsHyperNodesDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
  // The hostname.
  // 
  // example:
  // 
  // liliang-rmn7stf7-0000
  Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
  // The hyper node ID.
  // 
  // example:
  // 
  // e01-cn-2r42tmj4z02
  HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
  // The logon password.
  // 
  // example:
  // 
  // ***
  LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
  // The security group ID.
  // 
  // example:
  // 
  // sg-uf68xu2102avz7pl3t5d
  SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
  // The vSwitch ID.
  // 
  // example:
  // 
  // vsw-8vbobo4cvzsygw98f4j6b
  VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
  // The VPC ID.
  // 
  // example:
  // 
  // vpc-0jl8gs7qmx89739e210dn
  VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsHyperNodes) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsHyperNodes) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) GetDataDisk() []*ExtendClusterRequestNodeGroupsHyperNodesDataDisk  {
  return s.DataDisk
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) GetHostname() *string  {
  return s.Hostname
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) GetHyperNodeId() *string  {
  return s.HyperNodeId
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) GetLoginPassword() *string  {
  return s.LoginPassword
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) GetSecurityGroupId() *string  {
  return s.SecurityGroupId
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) GetVSwitchId() *string  {
  return s.VSwitchId
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) GetVpcId() *string  {
  return s.VpcId
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) SetDataDisk(v []*ExtendClusterRequestNodeGroupsHyperNodesDataDisk) *ExtendClusterRequestNodeGroupsHyperNodes {
  s.DataDisk = v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) SetHostname(v string) *ExtendClusterRequestNodeGroupsHyperNodes {
  s.Hostname = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) SetHyperNodeId(v string) *ExtendClusterRequestNodeGroupsHyperNodes {
  s.HyperNodeId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) SetLoginPassword(v string) *ExtendClusterRequestNodeGroupsHyperNodes {
  s.LoginPassword = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) SetSecurityGroupId(v string) *ExtendClusterRequestNodeGroupsHyperNodes {
  s.SecurityGroupId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) SetVSwitchId(v string) *ExtendClusterRequestNodeGroupsHyperNodes {
  s.VSwitchId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) SetVpcId(v string) *ExtendClusterRequestNodeGroupsHyperNodes {
  s.VpcId = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodes) Validate() error {
  if s.DataDisk != nil {
    for _, item := range s.DataDisk {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExtendClusterRequestNodeGroupsHyperNodesDataDisk struct {
  // Specifies whether to enable burst (I/O burst).
  // 
  // example:
  // 
  // false
  BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
  // The cloud disk type. Valid values:
  // 
  //  - cloud_essd: ESSD.
  // 
  // example:
  // 
  // cloud_essd
  Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
  // Specifies whether the data cloud disk is deleted when the node is unsubscribed.
  // 
  // example:
  // 
  // True
  DeleteWithNode *bool `json:"DeleteWithNode,omitempty" xml:"DeleteWithNode,omitempty"`
  // The performance level (PL) when an ESSD is used as a system cloud disk. Valid values:
  // 
  // - PL0: a maximum of 10,000 random read/write IOPS per disk.
  // 
  // - PL1: a maximum of 50,000 random read/write IOPS per disk.
  // 
  // example:
  // 
  // PL1
  PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
  // The provisioned performance (read/write IOPS) of a single ESSD AutoPL cloud disk.
  // 
  // example:
  // 
  // 9600
  ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
  // The cloud disk size. Unit: GiB.
  // 
  // example:
  // 
  // 10
  Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsHyperNodesDataDisk) String() string {
  return dara.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsHyperNodesDataDisk) GoString() string {
  return s.String()
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) GetBurstingEnabled() *bool  {
  return s.BurstingEnabled
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) GetCategory() *string  {
  return s.Category
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) GetDeleteWithNode() *bool  {
  return s.DeleteWithNode
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) GetPerformanceLevel() *string  {
  return s.PerformanceLevel
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) GetProvisionedIops() *int64  {
  return s.ProvisionedIops
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) GetSize() *int32  {
  return s.Size
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) SetBurstingEnabled(v bool) *ExtendClusterRequestNodeGroupsHyperNodesDataDisk {
  s.BurstingEnabled = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) SetCategory(v string) *ExtendClusterRequestNodeGroupsHyperNodesDataDisk {
  s.Category = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) SetDeleteWithNode(v bool) *ExtendClusterRequestNodeGroupsHyperNodesDataDisk {
  s.DeleteWithNode = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) SetPerformanceLevel(v string) *ExtendClusterRequestNodeGroupsHyperNodesDataDisk {
  s.PerformanceLevel = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) SetProvisionedIops(v int64) *ExtendClusterRequestNodeGroupsHyperNodesDataDisk {
  s.ProvisionedIops = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) SetSize(v int32) *ExtendClusterRequestNodeGroupsHyperNodesDataDisk {
  s.Size = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsHyperNodesDataDisk) Validate() error {
  return dara.Validate(s)
}

type ExtendClusterRequestNodeGroupsNodeTag struct {
  // The tag key of the node.
  // 
  // example:
  // 
  // key_my
  Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
  // The tag value of the node.
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
  // The data cloud disk specifications.
  DataDisk []*ExtendClusterRequestNodeGroupsNodesDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
  // The hostname.
  // 
  // example:
  // 
  // d044d220-33fd-11ed-86a6
  Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
  // The logon password.
  // 
  // example:
  // 
  // ***
  LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
  // The node ID.
  // 
  // example:
  // 
  // e01-cn-zvp2zdpy601
  NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
  // The security group ID.
  // 
  // example:
  // 
  // sg-uf68xu2102avz7pl3t5d
  SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
  // The vSwitch ID.
  // 
  // example:
  // 
  // vsw-bp169pi5fj151rrms4sia
  VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
  // The VPC ID.
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

func (s *ExtendClusterRequestNodeGroupsNodes) GetSecurityGroupId() *string  {
  return s.SecurityGroupId
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

func (s *ExtendClusterRequestNodeGroupsNodes) SetSecurityGroupId(v string) *ExtendClusterRequestNodeGroupsNodes {
  s.SecurityGroupId = &v
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
  if s.DataDisk != nil {
    for _, item := range s.DataDisk {
      if item != nil {
        if err := item.Validate(); err != nil {
          return err
        }
      }
    }
  }
  return nil
}

type ExtendClusterRequestNodeGroupsNodesDataDisk struct {
  // Specifies whether to enable burst (I/O burst).
  // 
  // example:
  // 
  // true
  BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
  // The type.
  // 
  // example:
  // 
  // cloud_essd
  Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
  // Specifies whether the data cloud disk is deleted when the node is unsubscribed.
  // 
  // example:
  // 
  // true
  DeleteWithNode *bool `json:"DeleteWithNode,omitempty" xml:"DeleteWithNode,omitempty"`
  // The performance metric of the data cloud disk.
  // 
  // example:
  // 
  // PL0
  PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
  // The provisioned performance (IOPS). Valid values: 0 to 50000.
  // 
  // example:
  // 
  // 1000
  ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
  // The cloud disk size.
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

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) GetBurstingEnabled() *bool  {
  return s.BurstingEnabled
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

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) GetProvisionedIops() *int64  {
  return s.ProvisionedIops
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) GetSize() *int32  {
  return s.Size
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetBurstingEnabled(v bool) *ExtendClusterRequestNodeGroupsNodesDataDisk {
  s.BurstingEnabled = &v
  return s
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

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetProvisionedIops(v int64) *ExtendClusterRequestNodeGroupsNodesDataDisk {
  s.ProvisionedIops = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) SetSize(v int32) *ExtendClusterRequestNodeGroupsNodesDataDisk {
  s.Size = &v
  return s
}

func (s *ExtendClusterRequestNodeGroupsNodesDataDisk) Validate() error {
  return dara.Validate(s)
}

