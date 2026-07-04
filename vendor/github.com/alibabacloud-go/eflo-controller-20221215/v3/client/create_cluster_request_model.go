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
	// The description of the cluster.
	//
	// example:
	//
	// Standard cluster test
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The name of the cluster.
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The type of the cluster.
	//
	// example:
	//
	// Lite
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The components (software instances).
	Components []*CreateClusterRequestComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Specifies whether to skip failed nodes. The default value is False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// The network information.
	Networks *CreateClusterRequestNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
	// The vSwitches for the node.
	NimizVSwitches []*string `json:"NimizVSwitches,omitempty" xml:"NimizVSwitches,omitempty" type:"Repeated"`
	// The list of node groups.
	NodeGroups []*CreateClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
	// Specifies whether the network interface supports jumbo frames.
	//
	// example:
	//
	// false
	OpenEniJumboFrame *bool `json:"OpenEniJumboFrame,omitempty" xml:"OpenEniJumboFrame,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The resource tags.
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
	// The component configuration.
	ComponentConfig *CreateClusterRequestComponentsComponentConfig `json:"ComponentConfig,omitempty" xml:"ComponentConfig,omitempty" type:"Struct"`
	// The component type.
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
	// The basic parameters of the component.
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
	// The node pool configuration. This is used to establish the mapping between node groups and node pools. This parameter is required when ComponentType is set to ACKEdge. Otherwise, leave it empty.
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
	// The IP address allocation policy.
	IpAllocationPolicy []*CreateClusterRequestNetworksIpAllocationPolicy `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty" type:"Repeated"`
	// The VPD configuration information.
	NewVpdInfo *CreateClusterRequestNetworksNewVpdInfo `json:"NewVpdInfo,omitempty" xml:"NewVpdInfo,omitempty" type:"Struct"`
	// The security group ID.
	//
	// example:
	//
	// sg-bp1d3dvbh9by7j5rujax
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IP version.
	//
	// example:
	//
	// IPv4
	TailIpVersion *string `json:"TailIpVersion,omitempty" xml:"TailIpVersion,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-asjdfklj
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The zone ID of the vSwitch.
	//
	// example:
	//
	// cn-shanghai-b
	VSwitchZoneId *string `json:"VSwitchZoneId,omitempty" xml:"VSwitchZoneId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-0jl36lqzmc06qogy0t5ll
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The information about the reused VPD.
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
	// The bond policy.
	BondPolicy *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy `json:"BondPolicy,omitempty" xml:"BondPolicy,omitempty" type:"Struct"`
	// The machine type allocation policy.
	MachineTypePolicy []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy `json:"MachineTypePolicy,omitempty" xml:"MachineTypePolicy,omitempty" type:"Repeated"`
	// The node allocation policy.
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
	// The default bond cluster subnet.
	//
	// example:
	//
	// 172.168.0.0/24
	BondDefaultSubnet *string `json:"BondDefaultSubnet,omitempty" xml:"BondDefaultSubnet,omitempty"`
	// The bond information.
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
	// The bond name.
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source cluster subnet for the IP address.
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
	// The bond information.
	Bonds []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	// The machine type.
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
	// The bond name.
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source cluster subnet for the IP address.
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
	// The bond information.
	Bonds []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	// The node ID.
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
	// The bond name.
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The source cluster subnet for the IP address.
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
	// The Cloud Enterprise Network (CEN) ID.
	//
	// example:
	//
	// cen-1gb1eftc5qp2ao75fo
	CenId *string `json:"CenId,omitempty" xml:"CenId,omitempty"`
	// The Cloud Link CIDR block.
	//
	// example:
	//
	// 172.16.0.0/24
	CloudLinkCidr *string `json:"CloudLinkCidr,omitempty" xml:"CloudLinkCidr,omitempty"`
	// The Cloud Link ID.
	//
	// example:
	//
	// vcc-cn-c4dtycm5i08
	CloudLinkId *string `json:"CloudLinkId,omitempty" xml:"CloudLinkId,omitempty"`
	// The VPC.
	//
	// example:
	//
	// vpc-0jl2x45apm6odc2c10h25
	MonitorVpcId *string `json:"MonitorVpcId,omitempty" xml:"MonitorVpcId,omitempty"`
	// The vSwitch.
	//
	// example:
	//
	// vsw-0jl2w3ffbghkss0x2foff
	MonitorVswitchId *string `json:"MonitorVswitchId,omitempty" xml:"MonitorVswitchId,omitempty"`
	// The CIDR block of the cluster.
	//
	// example:
	//
	// 192.168.0.0/16
	VpdCidr *string `json:"VpdCidr,omitempty" xml:"VpdCidr,omitempty"`
	// The subnets of the cluster.
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
	// The CIDR block of the subnet.
	//
	// example:
	//
	// 10.0.1.8/24
	SubnetCidr *string `json:"SubnetCidr,omitempty" xml:"SubnetCidr,omitempty"`
	// The subnet type.
	//
	// example:
	//
	// 10.0.2.8/24
	SubnetType *string `json:"SubnetType,omitempty" xml:"SubnetType,omitempty"`
	// The zone ID.
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
	// The VPD ID.
	//
	// example:
	//
	// vpd-vfuz6ejv
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
	// The list of cluster subnet IDs.
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
	// Specifies whether to enable file system mounting.
	//
	// example:
	//
	// false
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The list of supernodes.
	HyperNodes []*CreateClusterRequestNodeGroupsHyperNodes `json:"HyperNodes,omitempty" xml:"HyperNodes,omitempty" type:"Repeated"`
	// The OS image ID.
	//
	// example:
	//
	// i190297201634099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// test-keypair
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The logon password.
	//
	// example:
	//
	// Password
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The machine type.
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The description of the node group.
	//
	// example:
	//
	// Default node group
	NodeGroupDescription *string `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	// The name of the node group.
	//
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The list of nodes.
	Nodes []*CreateClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The name of the RAM role for the node. You can call the ListRoles operation of the RAM API to query the RAM roles that you have created. The trusted entity of the role must be Intelligent Computing LINGJUN.
	//
	// Note: You cannot clear an existing role.
	//
	// example:
	//
	// xianwen-test-ram-role
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The information about the system disk.
	SystemDisk *CreateClusterRequestNodeGroupsSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The instance user data. The data must be Base64-encoded. The raw data can be up to 16 KB in size.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Specifies whether to enable GPU virtualization.
	//
	// example:
	//
	// false
	VirtualGpuEnabled *bool `json:"VirtualGpuEnabled,omitempty" xml:"VirtualGpuEnabled,omitempty"`
	// The zone ID.
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

func (s *CreateClusterRequestNodeGroups) GetRamRoleName() *string {
	return s.RamRoleName
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

func (s *CreateClusterRequestNodeGroups) SetRamRoleName(v string) *CreateClusterRequestNodeGroups {
	s.RamRoleName = &v
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
	// The list of data disks.
	DataDisk []*CreateClusterRequestNodeGroupsHyperNodesDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The hostname.
	//
	// example:
	//
	// q25b01265.cloud.ng152
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The supernode ID.
	//
	// example:
	//
	// e01-dw72u2c23jk
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
	// The logon password.
	//
	// example:
	//
	// aaadddddfdsfdsfsdffd
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1xn9iq3s3p8218c4qu4
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-uf689skpx56nk7yfw0jhy
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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
	// Specifies whether to enable performance burst.
	//
	// example:
	//
	// true
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The disk type. Valid value:
	//
	// - cloud_essd: ESSD.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to delete the data disk when the node is released.
	//
	// example:
	//
	// true
	DeleteWithNode *bool `json:"DeleteWithNode,omitempty" xml:"DeleteWithNode,omitempty"`
	// The performance level of the ESSD that you create as a data disk. Valid values:
	//
	// - PL0: A single disk delivers up to 10,000 random read/write IOPS.
	//
	// - PL1: A single disk delivers up to 50,000 random read/write IOPS.
	//
	// - PL2: A single disk delivers up to 100,000 random read/write IOPS.
	//
	// - PL3: A single disk delivers up to 1,000,000 random read/write IOPS.
	//
	// Default value: PL1.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk.
	//
	// example:
	//
	// 10000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The disk size in GiB.
	//
	// example:
	//
	// 180
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
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
	// The specifications of the data disk.
	DataDisk []*CreateClusterRequestNodeGroupsNodesDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The hostname.
	//
	// example:
	//
	// 8d13b784-17a9-11ed-bc7b-acde48001122
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
	// e01poc-cn-i7m2wnivf0d
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp169pi5fj151rrms4sia
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The virtual private cloud (VPC) ID.
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
	// Specifies whether to enable performance burst.
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The type.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to delete the data disk when the node is released.
	//
	// example:
	//
	// true
	DeleteWithNode *bool `json:"DeleteWithNode,omitempty" xml:"DeleteWithNode,omitempty"`
	// The performance metric of the data disk.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned performance (IOPS). The value must be in the range of 0 to 50,000.
	//
	// example:
	//
	// 1000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The disk size.
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
	// The disk type. Valid value:
	//
	// - cloud_essd: enhanced SSD (ESSD).
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the ESSD that you create as a system disk. Valid values:
	//
	// - PL0: A single disk delivers up to 10,000 random read/write input/output operations per second (IOPS).
	//
	// - PL1: A single disk delivers up to 50,000 random read/write IOPS.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The unit is GB.
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
	// The key.
	//
	// example:
	//
	// env-name
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value.
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
