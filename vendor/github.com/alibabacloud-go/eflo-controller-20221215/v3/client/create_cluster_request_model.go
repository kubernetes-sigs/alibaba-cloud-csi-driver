// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterDescription(v string) *CreateClusterRequest
	GetClusterDescription() *string
	SetClusterName(v string) *CreateClusterRequest
	GetClusterName() *string
	SetClusterType(v string) *CreateClusterRequest
	GetClusterType() *string
	SetComponents(v []*CreateClusterRequestComponents) *CreateClusterRequest
	GetComponents() []*CreateClusterRequestComponents
	SetHpnZone(v string) *CreateClusterRequest
	GetHpnZone() *string
	SetIgnoreFailedNodeTasks(v bool) *CreateClusterRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNetworks(v *CreateClusterRequestNetworks) *CreateClusterRequest
	GetNetworks() *CreateClusterRequestNetworks
	SetNimizVSwitches(v []*string) *CreateClusterRequest
	GetNimizVSwitches() []*string
	SetNodeGroups(v []*CreateClusterRequestNodeGroups) *CreateClusterRequest
	GetNodeGroups() []*CreateClusterRequestNodeGroups
	SetOpenEniJumboFrame(v bool) *CreateClusterRequest
	GetOpenEniJumboFrame() *bool
	SetResourceGroupId(v string) *CreateClusterRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateClusterRequestTag) *CreateClusterRequest
	GetTag() []*CreateClusterRequestTag
}

type CreateClusterRequest struct {
	// Cluster description
	//
	// example:
	//
	// Cluster description
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// Cluster name
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Cluster type
	//
	// example:
	//
	// Lite
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Components (software instances)
	Components []*CreateClusterRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// Cluster number
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Whether to allow skipping failed nodes, the default value is False
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// Network information
	Networks *CreateClusterRequestNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
	// Node VSwitches
	NimizVSwitches []*string `json:"NimizVSwitches,omitempty" xml:"NimizVSwitches,omitempty" type:"Repeated"`
	// Node group list
	NodeGroups []*CreateClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// Whether the network interface supports jumbo frames
	//
	// example:
	//
	// false
	OpenEniJumboFrame *bool `json:"OpenEniJumboFrame,omitempty" xml:"OpenEniJumboFrame,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Resource tags
	Tag []*CreateClusterRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateClusterRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterRequest) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *CreateClusterRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreateClusterRequest) GetComponents() []*CreateClusterRequestComponents {
	return s.Components
}

func (s *CreateClusterRequest) GetHpnZone() *string {
	return s.HpnZone
}

func (s *CreateClusterRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *CreateClusterRequest) GetNetworks() *CreateClusterRequestNetworks {
	return s.Networks
}

func (s *CreateClusterRequest) GetNimizVSwitches() []*string {
	return s.NimizVSwitches
}

func (s *CreateClusterRequest) GetNodeGroups() []*CreateClusterRequestNodeGroups {
	return s.NodeGroups
}

func (s *CreateClusterRequest) GetOpenEniJumboFrame() *bool {
	return s.OpenEniJumboFrame
}

func (s *CreateClusterRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClusterRequest) GetTag() []*CreateClusterRequestTag {
	return s.Tag
}

func (s *CreateClusterRequest) SetClusterDescription(v string) *CreateClusterRequest {
	s.ClusterDescription = &v
	return s
}

func (s *CreateClusterRequest) SetClusterName(v string) *CreateClusterRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterRequest) SetClusterType(v string) *CreateClusterRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterRequest) SetComponents(v []*CreateClusterRequestComponents) *CreateClusterRequest {
	s.Components = v
	return s
}

func (s *CreateClusterRequest) SetHpnZone(v string) *CreateClusterRequest {
	s.HpnZone = &v
	return s
}

func (s *CreateClusterRequest) SetIgnoreFailedNodeTasks(v bool) *CreateClusterRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *CreateClusterRequest) SetNetworks(v *CreateClusterRequestNetworks) *CreateClusterRequest {
	s.Networks = v
	return s
}

func (s *CreateClusterRequest) SetNimizVSwitches(v []*string) *CreateClusterRequest {
	s.NimizVSwitches = v
	return s
}

func (s *CreateClusterRequest) SetNodeGroups(v []*CreateClusterRequestNodeGroups) *CreateClusterRequest {
	s.NodeGroups = v
	return s
}

func (s *CreateClusterRequest) SetOpenEniJumboFrame(v bool) *CreateClusterRequest {
	s.OpenEniJumboFrame = &v
	return s
}

func (s *CreateClusterRequest) SetResourceGroupId(v string) *CreateClusterRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterRequest) SetTag(v []*CreateClusterRequestTag) *CreateClusterRequest {
	s.Tag = v
	return s
}

func (s *CreateClusterRequest) Validate() error {
	if s.Components != nil {
		for _, item := range s.Components {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Networks != nil {
		if err := s.Networks.Validate(); err != nil {
			return err
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
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterRequestComponents struct {
	// Component configuration
	ComponentConfig *CreateClusterRequestComponentsComponentConfig `json:"ComponentConfig,omitempty" xml:"ComponentConfig,omitempty" type:"Struct"`
	// Component type
	//
	// example:
	//
	// ACKEdge
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
}

func (s CreateClusterRequestComponents) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestComponents) GetComponentConfig() *CreateClusterRequestComponentsComponentConfig {
	return s.ComponentConfig
}

func (s *CreateClusterRequestComponents) GetComponentType() *string {
	return s.ComponentType
}

func (s *CreateClusterRequestComponents) SetComponentConfig(v *CreateClusterRequestComponentsComponentConfig) *CreateClusterRequestComponents {
	s.ComponentConfig = v
	return s
}

func (s *CreateClusterRequestComponents) SetComponentType(v string) *CreateClusterRequestComponents {
	s.ComponentType = &v
	return s
}

func (s *CreateClusterRequestComponents) Validate() error {
	if s.ComponentConfig != nil {
		if err := s.ComponentConfig.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterRequestComponentsComponentConfig struct {
	// Basic component parameters
	//
	// example:
	//
	// {
	//
	//       "EndpointPublicAccess": false,
	//
	//       "ContainerCidr": "10.4.0.0/24",
	//
	//       "KeyPair": "test",
	//
	//       "NodeCidrMask": "25",
	//
	//       "ResourceGroupId": "rg-axsadm3sdzsdvdsndstdisd",
	//
	//       "WorkerSystemDiskCategory": "da",
	//
	//       "WorkerSystemDiskSize": 40,
	//
	//       "DeletionProtection": false,
	//
	//       "KubeProxy": "iptables",
	//
	//       "Name": "da",
	//
	//       "LoadBalancerSpec": "slb.s1.small",
	//
	//       "Runtime": {
	//
	//             "Version": "19.03.15",
	//
	//             "Name": "docker"
	//
	//       },
	//
	//       "IsEnterpriseSecurityGroup": true,
	//
	//       "Vpcid": "192.168.23.0/24",
	//
	//       "NumOfNodes": 1,
	//
	//       "VswitchIds": [
	//
	//             "dad"
	//
	//       ],
	//
	//       "ServiceCidr": "10.0.0.0/16",
	//
	//       "SnatEntry": false,
	//
	//       "kubernetesVersion": "1.20.11-aliyunedge.1",
	//
	//       "WorkerInstanceTypes": [
	//
	//             "da"
	//
	//       ]
	//
	// }
	BasicArgs interface{} `json:"BasicArgs,omitempty" xml:"BasicArgs,omitempty"`
	// Node pool configuration, used to establish the correspondence between node groups and node pools. Required when ComponentType is "ACKEdge", otherwise it can be empty.
	NodeUnits []interface{} `json:"NodeUnits,omitempty" xml:"NodeUnits,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestComponentsComponentConfig) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestComponentsComponentConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestComponentsComponentConfig) GetBasicArgs() interface{} {
	return s.BasicArgs
}

func (s *CreateClusterRequestComponentsComponentConfig) GetNodeUnits() []interface{} {
	return s.NodeUnits
}

func (s *CreateClusterRequestComponentsComponentConfig) SetBasicArgs(v interface{}) *CreateClusterRequestComponentsComponentConfig {
	s.BasicArgs = v
	return s
}

func (s *CreateClusterRequestComponentsComponentConfig) SetNodeUnits(v []interface{}) *CreateClusterRequestComponentsComponentConfig {
	s.NodeUnits = v
	return s
}

func (s *CreateClusterRequestComponentsComponentConfig) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestNetworks struct {
	// IP allocation policy
	IpAllocationPolicy []*CreateClusterRequestNetworksIpAllocationPolicy `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty" type:"Repeated"`
	// Vpd configuration information
	NewVpdInfo *CreateClusterRequestNetworksNewVpdInfo `json:"NewVpdInfo,omitempty" xml:"NewVpdInfo,omitempty" type:"Struct"`
	// Security group ID
	//
	// example:
	//
	// sg-bp1d3dvbh9by7j5rujax
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// IP version
	//
	// example:
	//
	// IPv4
	TailIpVersion *string `json:"TailIpVersion,omitempty" xml:"TailIpVersion,omitempty"`
	// VSwitch ID
	//
	// example:
	//
	// vsw-asjdfklj
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VSwitch Zone ID
	//
	// example:
	//
	// cn-shanghai-b
	VSwitchZoneId *string `json:"VSwitchZoneId,omitempty" xml:"VSwitchZoneId,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-0jl36lqzmc06qogy0t5ll
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Reuse VPD information
	VpdInfo *CreateClusterRequestNetworksVpdInfo `json:"VpdInfo,omitempty" xml:"VpdInfo,omitempty" type:"Struct"`
}

func (s CreateClusterRequestNetworks) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworks) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworks) GetIpAllocationPolicy() []*CreateClusterRequestNetworksIpAllocationPolicy {
	return s.IpAllocationPolicy
}

func (s *CreateClusterRequestNetworks) GetNewVpdInfo() *CreateClusterRequestNetworksNewVpdInfo {
	return s.NewVpdInfo
}

func (s *CreateClusterRequestNetworks) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateClusterRequestNetworks) GetTailIpVersion() *string {
	return s.TailIpVersion
}

func (s *CreateClusterRequestNetworks) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateClusterRequestNetworks) GetVSwitchZoneId() *string {
	return s.VSwitchZoneId
}

func (s *CreateClusterRequestNetworks) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterRequestNetworks) GetVpdInfo() *CreateClusterRequestNetworksVpdInfo {
	return s.VpdInfo
}

func (s *CreateClusterRequestNetworks) SetIpAllocationPolicy(v []*CreateClusterRequestNetworksIpAllocationPolicy) *CreateClusterRequestNetworks {
	s.IpAllocationPolicy = v
	return s
}

func (s *CreateClusterRequestNetworks) SetNewVpdInfo(v *CreateClusterRequestNetworksNewVpdInfo) *CreateClusterRequestNetworks {
	s.NewVpdInfo = v
	return s
}

func (s *CreateClusterRequestNetworks) SetSecurityGroupId(v string) *CreateClusterRequestNetworks {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateClusterRequestNetworks) SetTailIpVersion(v string) *CreateClusterRequestNetworks {
	s.TailIpVersion = &v
	return s
}

func (s *CreateClusterRequestNetworks) SetVSwitchId(v string) *CreateClusterRequestNetworks {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequestNetworks) SetVSwitchZoneId(v string) *CreateClusterRequestNetworks {
	s.VSwitchZoneId = &v
	return s
}

func (s *CreateClusterRequestNetworks) SetVpcId(v string) *CreateClusterRequestNetworks {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequestNetworks) SetVpdInfo(v *CreateClusterRequestNetworksVpdInfo) *CreateClusterRequestNetworks {
	s.VpdInfo = v
	return s
}

func (s *CreateClusterRequestNetworks) Validate() error {
	if s.IpAllocationPolicy != nil {
		for _, item := range s.IpAllocationPolicy {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.NewVpdInfo != nil {
		if err := s.NewVpdInfo.Validate(); err != nil {
			return err
		}
	}
	if s.VpdInfo != nil {
		if err := s.VpdInfo.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterRequestNetworksIpAllocationPolicy struct {
	// Bond policy
	BondPolicy *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy `json:"BondPolicy,omitempty" xml:"BondPolicy,omitempty" type:"Struct"`
	// Machine type allocation policy
	MachineTypePolicy []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy `json:"MachineTypePolicy,omitempty" xml:"MachineTypePolicy,omitempty" type:"Repeated"`
	// Node allocation policy
	NodePolicy []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicy `json:"NodePolicy,omitempty" xml:"NodePolicy,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicy) GetBondPolicy() *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy {
	return s.BondPolicy
}

func (s *CreateClusterRequestNetworksIpAllocationPolicy) GetMachineTypePolicy() []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy {
	return s.MachineTypePolicy
}

func (s *CreateClusterRequestNetworksIpAllocationPolicy) GetNodePolicy() []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicy {
	return s.NodePolicy
}

func (s *CreateClusterRequestNetworksIpAllocationPolicy) SetBondPolicy(v *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) *CreateClusterRequestNetworksIpAllocationPolicy {
	s.BondPolicy = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicy) SetMachineTypePolicy(v []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) *CreateClusterRequestNetworksIpAllocationPolicy {
	s.MachineTypePolicy = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicy) SetNodePolicy(v []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) *CreateClusterRequestNetworksIpAllocationPolicy {
	s.NodePolicy = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicy) Validate() error {
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

type CreateClusterRequestNetworksIpAllocationPolicyBondPolicy struct {
	// Default bond subnet for the cluster
	//
	// example:
	//
	// 172.168.0.0/24
	BondDefaultSubnet *string `json:"BondDefaultSubnet,omitempty" xml:"BondDefaultSubnet,omitempty"`
	// Bond information
	Bonds []*CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) GetBondDefaultSubnet() *string {
	return s.BondDefaultSubnet
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) GetBonds() []*CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds {
	return s.Bonds
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) SetBondDefaultSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy {
	s.BondDefaultSubnet = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) SetBonds(v []*CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy {
	s.Bonds = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) Validate() error {
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

type CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds struct {
	// Bond name
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// IP source subnet for the cluster
	//
	// example:
	//
	// 172.16.0.0/24
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) GetName() *string {
	return s.Name
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) GetSubnet() *string {
	return s.Subnet
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) SetName(v string) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) SetSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds {
	s.Subnet = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy struct {
	// Bond information
	Bonds []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	// Machine type
	//
	// example:
	//
	// efg1.nvga8n
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) GetBonds() []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds {
	return s.Bonds
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) GetMachineType() *string {
	return s.MachineType
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) SetBonds(v []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy {
	s.Bonds = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) SetMachineType(v string) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy {
	s.MachineType = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) Validate() error {
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

type CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds struct {
	// Bond name
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// IP source subnet for the cluster
	//
	// example:
	//
	// 192.168.1.0/24
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) GetName() *string {
	return s.Name
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) GetSubnet() *string {
	return s.Subnet
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) SetName(v string) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) SetSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds {
	s.Subnet = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestNetworksIpAllocationPolicyNodePolicy struct {
	// Bond information
	Bonds []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	// Node ID
	//
	// example:
	//
	// e01-cn-2r42vq62001
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) GetBonds() []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds {
	return s.Bonds
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) SetBonds(v []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy {
	s.Bonds = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) SetNodeId(v string) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy {
	s.NodeId = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) Validate() error {
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

type CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds struct {
	// Bond name
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// IP source subnet for the cluster
	//
	// example:
	//
	// 10.0.0.0/24
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) GetName() *string {
	return s.Name
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) GetSubnet() *string {
	return s.Subnet
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) SetName(v string) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) SetSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds {
	s.Subnet = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestNetworksNewVpdInfo struct {
	// Cloud Enterprise Network ID
	//
	// example:
	//
	// cen-1gb1eftc5qp2ao75fo
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// Cloud link CIDR
	//
	// example:
	//
	// 172.16.0.0/24
	CloudLinkCidr *string `json:"CloudLinkCidr,omitempty" xml:"CloudLinkCidr,omitempty"`
	// Cloud link ID
	//
	// example:
	//
	// vcc-cn-c4dtycm5i08
	CloudLinkId *string `json:"CloudLinkId,omitempty" xml:"CloudLinkId,omitempty"`
	// Virtual Private Cloud (VPC)
	//
	// example:
	//
	// vpc-0jl2x45apm6odc2c10h25
	MonitorVpcId *string `json:"MonitorVpcId,omitempty" xml:"MonitorVpcId,omitempty"`
	// VPC switch
	//
	// example:
	//
	// vsw-0jl2w3ffbghkss0x2foff
	MonitorVswitchId *string `json:"MonitorVswitchId,omitempty" xml:"MonitorVswitchId,omitempty"`
	// Cluster network segment
	//
	// example:
	//
	// 192.168.0.0/16
	VpdCidr *string `json:"VpdCidr,omitempty" xml:"VpdCidr,omitempty"`
	// Cluster subnets
	VpdSubnets []*CreateClusterRequestNetworksNewVpdInfoVpdSubnets `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksNewVpdInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworksNewVpdInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksNewVpdInfo) GetCenId() *string {
	return s.CenId
}

func (s *CreateClusterRequestNetworksNewVpdInfo) GetCloudLinkCidr() *string {
	return s.CloudLinkCidr
}

func (s *CreateClusterRequestNetworksNewVpdInfo) GetCloudLinkId() *string {
	return s.CloudLinkId
}

func (s *CreateClusterRequestNetworksNewVpdInfo) GetMonitorVpcId() *string {
	return s.MonitorVpcId
}

func (s *CreateClusterRequestNetworksNewVpdInfo) GetMonitorVswitchId() *string {
	return s.MonitorVswitchId
}

func (s *CreateClusterRequestNetworksNewVpdInfo) GetVpdCidr() *string {
	return s.VpdCidr
}

func (s *CreateClusterRequestNetworksNewVpdInfo) GetVpdSubnets() []*CreateClusterRequestNetworksNewVpdInfoVpdSubnets {
	return s.VpdSubnets
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetCenId(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.CenId = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetCloudLinkCidr(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.CloudLinkCidr = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetCloudLinkId(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.CloudLinkId = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetMonitorVpcId(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.MonitorVpcId = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetMonitorVswitchId(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.MonitorVswitchId = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetVpdCidr(v string) *CreateClusterRequestNetworksNewVpdInfo {
	s.VpdCidr = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) SetVpdSubnets(v []*CreateClusterRequestNetworksNewVpdInfoVpdSubnets) *CreateClusterRequestNetworksNewVpdInfo {
	s.VpdSubnets = v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfo) Validate() error {
	if s.VpdSubnets != nil {
		for _, item := range s.VpdSubnets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateClusterRequestNetworksNewVpdInfoVpdSubnets struct {
	// Subnet CIDR
	//
	// example:
	//
	// 10.0.1.8/24
	SubnetCidr *string `json:"SubnetCidr,omitempty" xml:"SubnetCidr,omitempty"`
	// Subnet type
	//
	// example:
	//
	// 10.0.2.8/24
	SubnetType *string `json:"SubnetType,omitempty" xml:"SubnetType,omitempty"`
	// Zone ID
	//
	// example:
	//
	// cn-wulanchabu-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequestNetworksNewVpdInfoVpdSubnets) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworksNewVpdInfoVpdSubnets) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksNewVpdInfoVpdSubnets) GetSubnetCidr() *string {
	return s.SubnetCidr
}

func (s *CreateClusterRequestNetworksNewVpdInfoVpdSubnets) GetSubnetType() *string {
	return s.SubnetType
}

func (s *CreateClusterRequestNetworksNewVpdInfoVpdSubnets) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateClusterRequestNetworksNewVpdInfoVpdSubnets) SetSubnetCidr(v string) *CreateClusterRequestNetworksNewVpdInfoVpdSubnets {
	s.SubnetCidr = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfoVpdSubnets) SetSubnetType(v string) *CreateClusterRequestNetworksNewVpdInfoVpdSubnets {
	s.SubnetType = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfoVpdSubnets) SetZoneId(v string) *CreateClusterRequestNetworksNewVpdInfoVpdSubnets {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequestNetworksNewVpdInfoVpdSubnets) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestNetworksVpdInfo struct {
	// VPC ID
	//
	// example:
	//
	// vpd-vfuz6ejv
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// List of cluster subnet IDs
	VpdSubnets []*string `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksVpdInfo) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNetworksVpdInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksVpdInfo) GetVpdId() *string {
	return s.VpdId
}

func (s *CreateClusterRequestNetworksVpdInfo) GetVpdSubnets() []*string {
	return s.VpdSubnets
}

func (s *CreateClusterRequestNetworksVpdInfo) SetVpdId(v string) *CreateClusterRequestNetworksVpdInfo {
	s.VpdId = &v
	return s
}

func (s *CreateClusterRequestNetworksVpdInfo) SetVpdSubnets(v []*string) *CreateClusterRequestNetworksVpdInfo {
	s.VpdSubnets = v
	return s
}

func (s *CreateClusterRequestNetworksVpdInfo) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestNodeGroups struct {
	// Whether to support file system mounting
	//
	// example:
	//
	// true
	FileSystemMountEnabled *bool                                       `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	HyperNodes             []*CreateClusterRequestNodeGroupsHyperNodes `json:"HyperNodes,omitempty" xml:"HyperNodes,omitempty" type:"Repeated"`
	// System image ID
	//
	// example:
	//
	// i190297201634099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Key pair name.
	//
	// example:
	//
	// test-keypair
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// Login password
	//
	// example:
	//
	// Password
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// Machine type
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Node group description
	//
	// example:
	//
	// Node group description
	NodeGroupDescription *string `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	// Node group name
	//
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// Node list
	Nodes []*CreateClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// System disk information
	SystemDisk *CreateClusterRequestNodeGroupsSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Instance custom data. It needs to be encoded in Base64, and the original data should not exceed 16 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Whether to enable gpu virtualization or not
	//
	// example:
	//
	// false
	VirtualGpuEnabled *bool `json:"VirtualGpuEnabled,omitempty" xml:"VirtualGpuEnabled,omitempty"`
	// Zone ID
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequestNodeGroups) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroups) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *CreateClusterRequestNodeGroups) GetHyperNodes() []*CreateClusterRequestNodeGroupsHyperNodes {
	return s.HyperNodes
}

func (s *CreateClusterRequestNodeGroups) GetImageId() *string {
	return s.ImageId
}

func (s *CreateClusterRequestNodeGroups) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateClusterRequestNodeGroups) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *CreateClusterRequestNodeGroups) GetMachineType() *string {
	return s.MachineType
}

func (s *CreateClusterRequestNodeGroups) GetNodeGroupDescription() *string {
	return s.NodeGroupDescription
}

func (s *CreateClusterRequestNodeGroups) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *CreateClusterRequestNodeGroups) GetNodes() []*CreateClusterRequestNodeGroupsNodes {
	return s.Nodes
}

func (s *CreateClusterRequestNodeGroups) GetSystemDisk() *CreateClusterRequestNodeGroupsSystemDisk {
	return s.SystemDisk
}

func (s *CreateClusterRequestNodeGroups) GetUserData() *string {
	return s.UserData
}

func (s *CreateClusterRequestNodeGroups) GetVirtualGpuEnabled() *bool {
	return s.VirtualGpuEnabled
}

func (s *CreateClusterRequestNodeGroups) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateClusterRequestNodeGroups) SetFileSystemMountEnabled(v bool) *CreateClusterRequestNodeGroups {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetHyperNodes(v []*CreateClusterRequestNodeGroupsHyperNodes) *CreateClusterRequestNodeGroups {
	s.HyperNodes = v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetImageId(v string) *CreateClusterRequestNodeGroups {
	s.ImageId = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetKeyPairName(v string) *CreateClusterRequestNodeGroups {
	s.KeyPairName = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetLoginPassword(v string) *CreateClusterRequestNodeGroups {
	s.LoginPassword = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetMachineType(v string) *CreateClusterRequestNodeGroups {
	s.MachineType = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetNodeGroupDescription(v string) *CreateClusterRequestNodeGroups {
	s.NodeGroupDescription = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetNodeGroupName(v string) *CreateClusterRequestNodeGroups {
	s.NodeGroupName = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetNodes(v []*CreateClusterRequestNodeGroupsNodes) *CreateClusterRequestNodeGroups {
	s.Nodes = v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetSystemDisk(v *CreateClusterRequestNodeGroupsSystemDisk) *CreateClusterRequestNodeGroups {
	s.SystemDisk = v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetUserData(v string) *CreateClusterRequestNodeGroups {
	s.UserData = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetVirtualGpuEnabled(v bool) *CreateClusterRequestNodeGroups {
	s.VirtualGpuEnabled = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetZoneId(v string) *CreateClusterRequestNodeGroups {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) Validate() error {
	if s.HyperNodes != nil {
		for _, item := range s.HyperNodes {
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
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateClusterRequestNodeGroupsHyperNodes struct {
	DataDisk      []*CreateClusterRequestNodeGroupsHyperNodesDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	Hostname      *string                                             `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	HyperNodeId   *string                                             `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
	LoginPassword *string                                             `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	VSwitchId     *string                                             `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId         *string                                             `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s CreateClusterRequestNodeGroupsHyperNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsHyperNodes) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) GetDataDisk() []*CreateClusterRequestNodeGroupsHyperNodesDataDisk {
	return s.DataDisk
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) GetHostname() *string {
	return s.Hostname
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) GetHyperNodeId() *string {
	return s.HyperNodeId
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) SetDataDisk(v []*CreateClusterRequestNodeGroupsHyperNodesDataDisk) *CreateClusterRequestNodeGroupsHyperNodes {
	s.DataDisk = v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) SetHostname(v string) *CreateClusterRequestNodeGroupsHyperNodes {
	s.Hostname = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) SetHyperNodeId(v string) *CreateClusterRequestNodeGroupsHyperNodes {
	s.HyperNodeId = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) SetLoginPassword(v string) *CreateClusterRequestNodeGroupsHyperNodes {
	s.LoginPassword = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) SetVSwitchId(v string) *CreateClusterRequestNodeGroupsHyperNodes {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) SetVpcId(v string) *CreateClusterRequestNodeGroupsHyperNodes {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodes) Validate() error {
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

type CreateClusterRequestNodeGroupsHyperNodesDataDisk struct {
	BurstingEnabled  *bool   `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category         *string `json:"Category,omitempty" xml:"Category,omitempty"`
	DeleteWithNode   *bool   `json:"DeleteWithNode,omitempty" xml:"DeleteWithNode,omitempty"`
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops  *int64  `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	Size             *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateClusterRequestNodeGroupsHyperNodesDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsHyperNodesDataDisk) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) GetDeleteWithNode() *bool {
	return s.DeleteWithNode
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) SetBurstingEnabled(v bool) *CreateClusterRequestNodeGroupsHyperNodesDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) SetCategory(v string) *CreateClusterRequestNodeGroupsHyperNodesDataDisk {
	s.Category = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) SetDeleteWithNode(v bool) *CreateClusterRequestNodeGroupsHyperNodesDataDisk {
	s.DeleteWithNode = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) SetPerformanceLevel(v string) *CreateClusterRequestNodeGroupsHyperNodesDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) SetProvisionedIops(v int64) *CreateClusterRequestNodeGroupsHyperNodesDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) SetSize(v int32) *CreateClusterRequestNodeGroupsHyperNodesDataDisk {
	s.Size = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsHyperNodesDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestNodeGroupsNodes struct {
	// Data disk specifications.
	DataDisk []*CreateClusterRequestNodeGroupsNodesDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// Hostname
	//
	// example:
	//
	// 8d13b784-17a9-11ed-bc7b-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Login password
	//
	// example:
	//
	// ***
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// Node ID
	//
	// example:
	//
	// e01poc-cn-i7m2wnivf0d
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

func (s CreateClusterRequestNodeGroupsNodes) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroupsNodes) GetDataDisk() []*CreateClusterRequestNodeGroupsNodesDataDisk {
	return s.DataDisk
}

func (s *CreateClusterRequestNodeGroupsNodes) GetHostname() *string {
	return s.Hostname
}

func (s *CreateClusterRequestNodeGroupsNodes) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *CreateClusterRequestNodeGroupsNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateClusterRequestNodeGroupsNodes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateClusterRequestNodeGroupsNodes) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateClusterRequestNodeGroupsNodes) SetDataDisk(v []*CreateClusterRequestNodeGroupsNodesDataDisk) *CreateClusterRequestNodeGroupsNodes {
	s.DataDisk = v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodes) SetHostname(v string) *CreateClusterRequestNodeGroupsNodes {
	s.Hostname = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodes) SetLoginPassword(v string) *CreateClusterRequestNodeGroupsNodes {
	s.LoginPassword = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodes) SetNodeId(v string) *CreateClusterRequestNodeGroupsNodes {
	s.NodeId = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodes) SetVSwitchId(v string) *CreateClusterRequestNodeGroupsNodes {
	s.VSwitchId = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodes) SetVpcId(v string) *CreateClusterRequestNodeGroupsNodes {
	s.VpcId = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodes) Validate() error {
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

type CreateClusterRequestNodeGroupsNodesDataDisk struct {
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// Type
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Whether the data disk is deleted with the node when it is unsubscribed
	//
	// example:
	//
	// true
	DeleteWithNode *bool `json:"DeleteWithNode,omitempty" xml:"DeleteWithNode,omitempty"`
	// Data disk performance level
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	ProvisionedIops  *int64  `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// Disk size
	//
	// example:
	//
	// 10
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateClusterRequestNodeGroupsNodesDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsNodesDataDisk) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) GetDeleteWithNode() *bool {
	return s.DeleteWithNode
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) SetBurstingEnabled(v bool) *CreateClusterRequestNodeGroupsNodesDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) SetCategory(v string) *CreateClusterRequestNodeGroupsNodesDataDisk {
	s.Category = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) SetDeleteWithNode(v bool) *CreateClusterRequestNodeGroupsNodesDataDisk {
	s.DeleteWithNode = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) SetPerformanceLevel(v string) *CreateClusterRequestNodeGroupsNodesDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) SetProvisionedIops(v int64) *CreateClusterRequestNodeGroupsNodesDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) SetSize(v int32) *CreateClusterRequestNodeGroupsNodesDataDisk {
	s.Size = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsNodesDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestNodeGroupsSystemDisk struct {
	// Disk type. The value range is:
	//
	// - cloud_essd: ESSD disk.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// When creating an ESSD disk as the system disk, set the performance level of the disk. The value range is:
	//
	// - PL0: Maximum random read/write IOPS for a single disk is 10,000.
	//
	// - PL1: Maximum random read/write IOPS for a single disk is 50,000.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// Unit: GB.
	//
	// example:
	//
	// 9999
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateClusterRequestNodeGroupsSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroupsSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateClusterRequestNodeGroupsSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateClusterRequestNodeGroupsSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateClusterRequestNodeGroupsSystemDisk) SetCategory(v string) *CreateClusterRequestNodeGroupsSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsSystemDisk) SetPerformanceLevel(v string) *CreateClusterRequestNodeGroupsSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsSystemDisk) SetSize(v int32) *CreateClusterRequestNodeGroupsSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestTag struct {
	// Key
	//
	// example:
	//
	// env-name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Value
	//
	// example:
	//
	// dev
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateClusterRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateClusterRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateClusterRequestTag) SetKey(v string) *CreateClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTag) SetValue(v string) *CreateClusterRequestTag {
	s.Value = &v
	return s
}

func (s *CreateClusterRequestTag) Validate() error {
	return dara.Validate(s)
}
