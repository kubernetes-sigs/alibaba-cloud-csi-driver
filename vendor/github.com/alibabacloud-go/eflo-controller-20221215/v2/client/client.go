// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ApproveOperationRequest struct {
	// Node ID
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Operation Type
	//
	// example:
	//
	// RepairMachine
	OperationType *string `json:"OperationType,omitempty" xml:"OperationType,omitempty"`
}

func (s ApproveOperationRequest) String() string {
	return tea.Prettify(s)
}

func (s ApproveOperationRequest) GoString() string {
	return s.String()
}

func (s *ApproveOperationRequest) SetNodeId(v string) *ApproveOperationRequest {
	s.NodeId = &v
	return s
}

func (s *ApproveOperationRequest) SetOperationType(v string) *ApproveOperationRequest {
	s.OperationType = &v
	return s
}

type ApproveOperationResponseBody struct {
	// Error Message
	//
	// example:
	//
	// Resource not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Request ID
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ApproveOperationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ApproveOperationResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveOperationResponseBody) SetErrorMessage(v string) *ApproveOperationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ApproveOperationResponseBody) SetRequestId(v string) *ApproveOperationResponseBody {
	s.RequestId = &v
	return s
}

type ApproveOperationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveOperationResponse) String() string {
	return tea.Prettify(s)
}

func (s ApproveOperationResponse) GoString() string {
	return s.String()
}

func (s *ApproveOperationResponse) SetHeaders(v map[string]*string) *ApproveOperationResponse {
	s.Headers = v
	return s
}

func (s *ApproveOperationResponse) SetStatusCode(v int32) *ApproveOperationResponse {
	s.StatusCode = &v
	return s
}

func (s *ApproveOperationResponse) SetBody(v *ApproveOperationResponseBody) *ApproveOperationResponse {
	s.Body = v
	return s
}

type ChangeResourceGroupRequest struct {
	// System-defined parameter. Value: **ChangeResourceGroup**.
	//
	// This parameter is required.
	//
	// example:
	//
	// rg-aekzyqdwnfabx6q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// $.parameters[1].schema.example
	//
	// This parameter is required.
	//
	// example:
	//
	// i118099391667548921125
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Resource Group Change
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu
	ResourceRegionId *string `json:"ResourceRegionId,omitempty" xml:"ResourceRegionId,omitempty"`
	// $.parameters[3].schema.enumValueTitles
	//
	// example:
	//
	// Node
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) SetResourceGroupId(v string) *ChangeResourceGroupRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceId(v string) *ChangeResourceGroupRequest {
	s.ResourceId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceRegionId(v string) *ChangeResourceGroupRequest {
	s.ResourceRegionId = &v
	return s
}

func (s *ChangeResourceGroupRequest) SetResourceType(v string) *ChangeResourceGroupRequest {
	s.ResourceType = &v
	return s
}

type ChangeResourceGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ChangeResourceGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) SetHeaders(v map[string]*string) *ChangeResourceGroupResponse {
	s.Headers = v
	return s
}

func (s *ChangeResourceGroupResponse) SetStatusCode(v int32) *ChangeResourceGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ChangeResourceGroupResponse) SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse {
	s.Body = v
	return s
}

type CloseSessionRequest struct {
	// Session ID
	//
	// example:
	//
	// i207023871669364793713
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Session token
	//
	// example:
	//
	// 03f53c719015a9ad4f4f55d66cac2dac161b18e8065ca75a3220b89de389c980
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
}

func (s CloseSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseSessionRequest) GoString() string {
	return s.String()
}

func (s *CloseSessionRequest) SetSessionId(v string) *CloseSessionRequest {
	s.SessionId = &v
	return s
}

func (s *CloseSessionRequest) SetSessionToken(v string) *CloseSessionRequest {
	s.SessionToken = &v
	return s
}

type CloseSessionResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 07AA3A1F-321E-50D8-B834-88C411331C94
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Session ID.
	//
	// example:
	//
	// i206495551737511455528
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// ClosingActive
	//
	// example:
	//
	// Inactive
	State *string `json:"State,omitempty" xml:"State,omitempty"`
}

func (s CloseSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CloseSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseSessionResponseBody) SetRequestId(v string) *CloseSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseSessionResponseBody) SetSessionId(v string) *CloseSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *CloseSessionResponseBody) SetState(v string) *CloseSessionResponseBody {
	s.State = &v
	return s
}

type CloseSessionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseSessionResponse) GoString() string {
	return s.String()
}

func (s *CloseSessionResponse) SetHeaders(v map[string]*string) *CloseSessionResponse {
	s.Headers = v
	return s
}

func (s *CloseSessionResponse) SetStatusCode(v int32) *CloseSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CloseSessionResponse) SetBody(v *CloseSessionResponseBody) *CloseSessionResponse {
	s.Body = v
	return s
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
	// Whether to allow skipping failed nodes, default value is False
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
	// Open Eni Jumbo Frame
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
	return tea.Prettify(s)
}

func (s CreateClusterRequest) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s CreateClusterRequestComponents) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestComponents) SetComponentConfig(v *CreateClusterRequestComponentsComponentConfig) *CreateClusterRequestComponents {
	s.ComponentConfig = v
	return s
}

func (s *CreateClusterRequestComponents) SetComponentType(v string) *CreateClusterRequestComponents {
	s.ComponentType = &v
	return s
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
	return tea.Prettify(s)
}

func (s CreateClusterRequestComponentsComponentConfig) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestComponentsComponentConfig) SetBasicArgs(v interface{}) *CreateClusterRequestComponentsComponentConfig {
	s.BasicArgs = v
	return s
}

func (s *CreateClusterRequestComponentsComponentConfig) SetNodeUnits(v []interface{}) *CreateClusterRequestComponentsComponentConfig {
	s.NodeUnits = v
	return s
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
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworks) GoString() string {
	return s.String()
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

type CreateClusterRequestNetworksIpAllocationPolicy struct {
	// Bond policy
	BondPolicy *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy `json:"BondPolicy,omitempty" xml:"BondPolicy,omitempty" type:"Struct"`
	// Machine type allocation policy
	MachineTypePolicy []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy `json:"MachineTypePolicy,omitempty" xml:"MachineTypePolicy,omitempty" type:"Repeated"`
	// Node allocation policy
	NodePolicy []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicy `json:"NodePolicy,omitempty" xml:"NodePolicy,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicy) GoString() string {
	return s.String()
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

type CreateClusterRequestNetworksIpAllocationPolicyBondPolicy struct {
	// Default bond cluster subnet
	//
	// example:
	//
	// 172.168.0.0/24
	BondDefaultSubnet *string `json:"BondDefaultSubnet,omitempty" xml:"BondDefaultSubnet,omitempty"`
	// Bond information
	Bonds []*CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) SetBondDefaultSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy {
	s.BondDefaultSubnet = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy) SetBonds(v []*CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicy {
	s.Bonds = v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds struct {
	// Bond name
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// IP source cluster subnet
	//
	// example:
	//
	// 172.16.0.0/24
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) SetName(v string) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds) SetSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyBondPolicyBonds {
	s.Subnet = &v
	return s
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
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) SetBonds(v []*CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy {
	s.Bonds = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy) SetMachineType(v string) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicy {
	s.MachineType = &v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds struct {
	// Bond name
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// IP source cluster subnet
	//
	// example:
	//
	// 192.168.1.0/24
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) SetName(v string) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds) SetSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyMachineTypePolicyBonds {
	s.Subnet = &v
	return s
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
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) SetBonds(v []*CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy {
	s.Bonds = v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy) SetNodeId(v string) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicy {
	s.NodeId = &v
	return s
}

type CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds struct {
	// Bond name
	//
	// example:
	//
	// bond0
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// IP source cluster subnet
	//
	// example:
	//
	// 10.0.0.0/24
	Subnet *string `json:"Subnet,omitempty" xml:"Subnet,omitempty"`
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) SetName(v string) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds {
	s.Name = &v
	return s
}

func (s *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds) SetSubnet(v string) *CreateClusterRequestNetworksIpAllocationPolicyNodePolicyBonds {
	s.Subnet = &v
	return s
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
	// VPC
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
	// Cluster Network Segment
	//
	// example:
	//
	// 192.168.0.0/16
	VpdCidr *string `json:"VpdCidr,omitempty" xml:"VpdCidr,omitempty"`
	// Cluster subnets
	VpdSubnets []*CreateClusterRequestNetworksNewVpdInfoVpdSubnets `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty" type:"Repeated"`
}

func (s CreateClusterRequestNetworksNewVpdInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksNewVpdInfo) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksNewVpdInfoVpdSubnets) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s CreateClusterRequestNetworksVpdInfo) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNetworksVpdInfo) SetVpdId(v string) *CreateClusterRequestNetworksVpdInfo {
	s.VpdId = &v
	return s
}

func (s *CreateClusterRequestNetworksVpdInfo) SetVpdSubnets(v []*string) *CreateClusterRequestNetworksVpdInfo {
	s.VpdSubnets = v
	return s
}

type CreateClusterRequestNodeGroups struct {
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// System image ID
	//
	// example:
	//
	// i190297201634099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// sc-key
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
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
	// SystemDisk
	SystemDisk *CreateClusterRequestNodeGroupsSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Instance custom data. It needs to be Base64 encoded, and the original data should not exceed 16 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Zone ID
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateClusterRequestNodeGroups) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroups) SetFileSystemMountEnabled(v bool) *CreateClusterRequestNodeGroups {
	s.FileSystemMountEnabled = &v
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

func (s *CreateClusterRequestNodeGroups) SetZoneId(v string) *CreateClusterRequestNodeGroups {
	s.ZoneId = &v
	return s
}

type CreateClusterRequestNodeGroupsNodes struct {
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
	// Virtual switch ID
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
	return tea.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
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

type CreateClusterRequestNodeGroupsSystemDisk struct {
	// The disk type. Valid values:
	//
	// 	- cloud_ssd: standard SSD
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Valid values:
	//
	// 	- PL0: A single ESSD can provide up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can provide up to 50,000 random read/write IOPS.
	//
	//
	// Default value: PL1.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size. Unit: GB.
	//
	// example:
	//
	// 20
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateClusterRequestNodeGroupsSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsSystemDisk) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s CreateClusterRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestTag) SetKey(v string) *CreateClusterRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterRequestTag) SetValue(v string) *CreateClusterRequestTag {
	s.Value = &v
	return s
}

type CreateClusterShrinkRequest struct {
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
	ComponentsShrink *string `json:"Components,omitempty" xml:"Components,omitempty"`
	// Cluster number
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Whether to allow skipping failed nodes, default value is False
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// Network information
	NetworksShrink *string `json:"Networks,omitempty" xml:"Networks,omitempty"`
	// Node VSwitches
	NimizVSwitchesShrink *string `json:"NimizVSwitches,omitempty" xml:"NimizVSwitches,omitempty"`
	// Node group list
	NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
	// Open Eni Jumbo Frame
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
	Tag []*CreateClusterShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) SetClusterDescription(v string) *CreateClusterShrinkRequest {
	s.ClusterDescription = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterName(v string) *CreateClusterShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetClusterType(v string) *CreateClusterShrinkRequest {
	s.ClusterType = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetComponentsShrink(v string) *CreateClusterShrinkRequest {
	s.ComponentsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetHpnZone(v string) *CreateClusterShrinkRequest {
	s.HpnZone = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *CreateClusterShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetNetworksShrink(v string) *CreateClusterShrinkRequest {
	s.NetworksShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetNimizVSwitchesShrink(v string) *CreateClusterShrinkRequest {
	s.NimizVSwitchesShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetNodeGroupsShrink(v string) *CreateClusterShrinkRequest {
	s.NodeGroupsShrink = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetOpenEniJumboFrame(v bool) *CreateClusterShrinkRequest {
	s.OpenEniJumboFrame = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetResourceGroupId(v string) *CreateClusterShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateClusterShrinkRequest) SetTag(v []*CreateClusterShrinkRequestTag) *CreateClusterShrinkRequest {
	s.Tag = v
	return s
}

type CreateClusterShrinkRequestTag struct {
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

func (s CreateClusterShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequestTag) SetKey(v string) *CreateClusterShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterShrinkRequestTag) SetValue(v string) *CreateClusterShrinkRequestTag {
	s.Value = &v
	return s
}

type CreateClusterResponseBody struct {
	// Cluster ID
	//
	// example:
	//
	// i116913051663373010974
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Request ID
	//
	// example:
	//
	// 3C683243-7915-57FB-9570-A2932C1C0F78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task Id
	//
	// example:
	//
	// i159809891662373011015
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s CreateClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) SetClusterId(v string) *CreateClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *CreateClusterResponseBody) SetRequestId(v string) *CreateClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateClusterResponseBody) SetTaskId(v string) *CreateClusterResponseBody {
	s.TaskId = &v
	return s
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) SetHeaders(v map[string]*string) *CreateClusterResponse {
	s.Headers = v
	return s
}

func (s *CreateClusterResponse) SetStatusCode(v int32) *CreateClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateClusterResponse) SetBody(v *CreateClusterResponseBody) *CreateClusterResponse {
	s.Body = v
	return s
}

type CreateDiagnosticTaskRequest struct {
	// Log information
	AiJobLogInfo *CreateDiagnosticTaskRequestAiJobLogInfo `json:"AiJobLogInfo,omitempty" xml:"AiJobLogInfo,omitempty" type:"Struct"`
	// Cluster ID
	//
	// example:
	//
	// i118913031696573280136
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Diagnostic type.
	//
	// example:
	//
	// CheckByAiJobLogs
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// List of node IDs
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
}

func (s CreateDiagnosticTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequest) SetAiJobLogInfo(v *CreateDiagnosticTaskRequestAiJobLogInfo) *CreateDiagnosticTaskRequest {
	s.AiJobLogInfo = v
	return s
}

func (s *CreateDiagnosticTaskRequest) SetClusterId(v string) *CreateDiagnosticTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateDiagnosticTaskRequest) SetDiagnosticType(v string) *CreateDiagnosticTaskRequest {
	s.DiagnosticType = &v
	return s
}

func (s *CreateDiagnosticTaskRequest) SetNodeIds(v []*string) *CreateDiagnosticTaskRequest {
	s.NodeIds = v
	return s
}

type CreateDiagnosticTaskRequestAiJobLogInfo struct {
	// Task logs
	AiJobLogs []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs `json:"AiJobLogs,omitempty" xml:"AiJobLogs,omitempty" type:"Repeated"`
	// End time. In timestamp format, unit: seconds.
	//
	// > Must be on the hour or half-hour mark.
	//
	// example:
	//
	// 2024-08-05T10:10:01
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// Start time. In timestamp format, unit: seconds.
	//
	// > Must be on the hour or half-hour mark.
	//
	// example:
	//
	// 2024-10-11T00:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateDiagnosticTaskRequestAiJobLogInfo) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfo) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) SetAiJobLogs(v []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) *CreateDiagnosticTaskRequestAiJobLogInfo {
	s.AiJobLogs = v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) SetEndTime(v string) *CreateDiagnosticTaskRequestAiJobLogInfo {
	s.EndTime = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) SetStartTime(v string) *CreateDiagnosticTaskRequestAiJobLogInfo {
	s.StartTime = &v
	return s
}

type CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs struct {
	// Instance ID
	//
	// example:
	//
	// null
	AiInstance *string `json:"AiInstance,omitempty" xml:"AiInstance,omitempty"`
	// Log object
	Logs []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// Node ID
	//
	// example:
	//
	// e01-tw-p2p2al5u1hn
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) SetAiInstance(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	s.AiInstance = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) SetLogs(v []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	s.Logs = v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) SetNodeId(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	s.NodeId = &v
	return s
}

type CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs struct {
	// Sent date, in the format yyyymmdd.
	//
	// example:
	//
	// 2024-08-05T10:10:01
	Datetime *string `json:"Datetime,omitempty" xml:"Datetime,omitempty"`
	// Log content
	//
	// example:
	//
	// success
	LogContent *string `json:"LogContent,omitempty" xml:"LogContent,omitempty"`
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) SetDatetime(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs {
	s.Datetime = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) SetLogContent(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs {
	s.LogContent = &v
	return s
}

type CreateDiagnosticTaskShrinkRequest struct {
	// Log information
	AiJobLogInfoShrink *string `json:"AiJobLogInfo,omitempty" xml:"AiJobLogInfo,omitempty"`
	// Cluster ID
	//
	// example:
	//
	// i118913031696573280136
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Diagnostic type.
	//
	// example:
	//
	// CheckByAiJobLogs
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// List of node IDs
	NodeIdsShrink *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
}

func (s CreateDiagnosticTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskShrinkRequest) SetAiJobLogInfoShrink(v string) *CreateDiagnosticTaskShrinkRequest {
	s.AiJobLogInfoShrink = &v
	return s
}

func (s *CreateDiagnosticTaskShrinkRequest) SetClusterId(v string) *CreateDiagnosticTaskShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateDiagnosticTaskShrinkRequest) SetDiagnosticType(v string) *CreateDiagnosticTaskShrinkRequest {
	s.DiagnosticType = &v
	return s
}

func (s *CreateDiagnosticTaskShrinkRequest) SetNodeIdsShrink(v string) *CreateDiagnosticTaskShrinkRequest {
	s.NodeIdsShrink = &v
	return s
}

type CreateDiagnosticTaskResponseBody struct {
	// Diagnosis ID
	//
	// example:
	//
	// diag-i150553931717380274931
	DiagnosticId *string `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
	// Request ID
	//
	// example:
	//
	// A511C02A-0127-51AA-A9F9-966382C9A1B5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDiagnosticTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskResponseBody) SetDiagnosticId(v string) *CreateDiagnosticTaskResponseBody {
	s.DiagnosticId = &v
	return s
}

func (s *CreateDiagnosticTaskResponseBody) SetRequestId(v string) *CreateDiagnosticTaskResponseBody {
	s.RequestId = &v
	return s
}

type CreateDiagnosticTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiagnosticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiagnosticTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDiagnosticTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskResponse) SetHeaders(v map[string]*string) *CreateDiagnosticTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateDiagnosticTaskResponse) SetStatusCode(v int32) *CreateDiagnosticTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDiagnosticTaskResponse) SetBody(v *CreateDiagnosticTaskResponseBody) *CreateDiagnosticTaskResponse {
	s.Body = v
	return s
}

type CreateNetTestTaskRequest struct {
	// Cluster ID
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster name
	//
	// example:
	//
	// Eflo-YJ-Test-Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Required when the test type is communication library testing
	CommTest *CreateNetTestTaskRequestCommTest `json:"CommTest,omitempty" xml:"CommTest,omitempty" type:"Struct"`
	// Fill in this field when the network test type is delay testing.
	DelayTest *CreateNetTestTaskRequestDelayTest `json:"DelayTest,omitempty" xml:"DelayTest,omitempty" type:"Struct"`
	// Network test type.
	//
	// For example: DelayTest for latency testing, TrafficTest for traffic testing, CommTest for communication library testing.
	//
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// Network mode
	//
	// example:
	//
	// 2
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// Test port number.
	//
	// example:
	//
	// 23604
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// This field is empty if the TrafficModel is Fullmesh.
	TrafficTest *CreateNetTestTaskRequestTrafficTest `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty" type:"Struct"`
}

func (s CreateNetTestTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequest) SetClusterId(v string) *CreateNetTestTaskRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetClusterName(v string) *CreateNetTestTaskRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetCommTest(v *CreateNetTestTaskRequestCommTest) *CreateNetTestTaskRequest {
	s.CommTest = v
	return s
}

func (s *CreateNetTestTaskRequest) SetDelayTest(v *CreateNetTestTaskRequestDelayTest) *CreateNetTestTaskRequest {
	s.DelayTest = v
	return s
}

func (s *CreateNetTestTaskRequest) SetNetTestType(v string) *CreateNetTestTaskRequest {
	s.NetTestType = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetNetworkMode(v string) *CreateNetTestTaskRequest {
	s.NetworkMode = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetPort(v string) *CreateNetTestTaskRequest {
	s.Port = &v
	return s
}

func (s *CreateNetTestTaskRequest) SetTrafficTest(v *CreateNetTestTaskRequestTrafficTest) *CreateNetTestTaskRequest {
	s.TrafficTest = v
	return s
}

type CreateNetTestTaskRequestCommTest struct {
	// Number of GPUs
	//
	// example:
	//
	// 1
	GPUNum *int64 `json:"GPUNum,omitempty" xml:"GPUNum,omitempty"`
	// Resource ID
	Hosts []*CreateNetTestTaskRequestCommTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// Communication library model
	//
	// example:
	//
	// intention_v4
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Communication library test category: ACCL or NCCL
	//
	// example:
	//
	// ACCL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s CreateNetTestTaskRequestCommTest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestCommTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestCommTest) SetGPUNum(v int64) *CreateNetTestTaskRequestCommTest {
	s.GPUNum = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTest) SetHosts(v []*CreateNetTestTaskRequestCommTestHosts) *CreateNetTestTaskRequestCommTest {
	s.Hosts = v
	return s
}

func (s *CreateNetTestTaskRequestCommTest) SetModel(v string) *CreateNetTestTaskRequestCommTest {
	s.Model = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTest) SetType(v string) *CreateNetTestTaskRequestCommTest {
	s.Type = &v
	return s
}

type CreateNetTestTaskRequestCommTestHosts struct {
	// IP address.
	//
	// example:
	//
	// 169.253.253.15
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Node ID.
	//
	// example:
	//
	// e01-tw-bqisacl3z6l
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Resource ID
	//
	// example:
	//
	// i111670831721110797708
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// VBw
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestCommTestHosts) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestCommTestHosts) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetIP(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetNodeId(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetResourceId(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestCommTestHosts) SetServerName(v string) *CreateNetTestTaskRequestCommTestHosts {
	s.ServerName = &v
	return s
}

type CreateNetTestTaskRequestDelayTest struct {
	// 输入测试节点的hosts
	Hosts []*CreateNetTestTaskRequestDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s CreateNetTestTaskRequestDelayTest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestDelayTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestDelayTest) SetHosts(v []*CreateNetTestTaskRequestDelayTestHosts) *CreateNetTestTaskRequestDelayTest {
	s.Hosts = v
	return s
}

type CreateNetTestTaskRequestDelayTestHosts struct {
	// Network interface bond port
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// Node IP
	//
	// example:
	//
	// 125.210.225.48
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Node ID.
	//
	// example:
	//
	// e01-cn-fou43an0a05
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Resource ID
	//
	// example:
	//
	// e01-cn-bcd3u1aee06
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// NQU
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestDelayTestHosts) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestDelayTestHosts) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetBond(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.Bond = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetIP(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetNodeId(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetResourceId(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestDelayTestHosts) SetServerName(v string) *CreateNetTestTaskRequestDelayTestHosts {
	s.ServerName = &v
	return s
}

type CreateNetTestTaskRequestTrafficTest struct {
	// Resource ID.
	Clients []*CreateNetTestTaskRequestTrafficTestClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// The duration of the workflow task in seconds.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// Enter True/False when the protocol is RDMA,
	//
	// this field is empty when the protocol is TCP.
	//
	// example:
	//
	// False
	GDR *bool `json:"GDR,omitempty" xml:"GDR,omitempty"`
	// Network protocol, either RDMA or TCP.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// Enter the number of concurrent connections when the protocol is TCP, or enter the configured QP value when the protocol is RDMA.
	//
	// example:
	//
	// 1
	QP *int64 `json:"QP,omitempty" xml:"QP,omitempty"`
	// Service list
	Servers []*CreateNetTestTaskRequestTrafficTestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// Traffic model, either MTON or Fullmesh.
	//
	// example:
	//
	// Fullmesh
	TrafficModel *string `json:"TrafficModel,omitempty" xml:"TrafficModel,omitempty"`
}

func (s CreateNetTestTaskRequestTrafficTest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTest) SetClients(v []*CreateNetTestTaskRequestTrafficTestClients) *CreateNetTestTaskRequestTrafficTest {
	s.Clients = v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetDuration(v int64) *CreateNetTestTaskRequestTrafficTest {
	s.Duration = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetGDR(v bool) *CreateNetTestTaskRequestTrafficTest {
	s.GDR = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetProtocol(v string) *CreateNetTestTaskRequestTrafficTest {
	s.Protocol = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetQP(v int64) *CreateNetTestTaskRequestTrafficTest {
	s.QP = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetServers(v []*CreateNetTestTaskRequestTrafficTestServers) *CreateNetTestTaskRequestTrafficTest {
	s.Servers = v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTest) SetTrafficModel(v string) *CreateNetTestTaskRequestTrafficTest {
	s.TrafficModel = &v
	return s
}

type CreateNetTestTaskRequestTrafficTestClients struct {
	// Network card bond interface
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// Node IP
	//
	// example:
	//
	// 192.168.1.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Node ID
	//
	// example:
	//
	// e01-tw-w5elqg7pw18
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Resource ID
	//
	// example:
	//
	// e01-cn-20s41p6cx01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestTrafficTestClients) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTestClients) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetBond(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.Bond = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetIP(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetNodeId(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetResourceId(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestClients) SetServerName(v string) *CreateNetTestTaskRequestTrafficTestClients {
	s.ServerName = &v
	return s
}

type CreateNetTestTaskRequestTrafficTestServers struct {
	// Network card bond interface
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// Node IP
	//
	// example:
	//
	// 47.121.110.190
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Node ID
	//
	// example:
	//
	// e01-tw-bqisacl3z6l
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Resource ID
	//
	// example:
	//
	// e01-cn-wwo3etaqu0b
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s CreateNetTestTaskRequestTrafficTestServers) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTestServers) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetBond(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.Bond = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetIP(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.IP = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetNodeId(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.NodeId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetResourceId(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.ResourceId = &v
	return s
}

func (s *CreateNetTestTaskRequestTrafficTestServers) SetServerName(v string) *CreateNetTestTaskRequestTrafficTestServers {
	s.ServerName = &v
	return s
}

type CreateNetTestTaskShrinkRequest struct {
	// Cluster ID
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster name
	//
	// example:
	//
	// Eflo-YJ-Test-Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Required when the test type is communication library testing
	CommTestShrink *string `json:"CommTest,omitempty" xml:"CommTest,omitempty"`
	// Fill in this field when the network test type is delay testing.
	DelayTestShrink *string `json:"DelayTest,omitempty" xml:"DelayTest,omitempty"`
	// Network test type.
	//
	// For example: DelayTest for latency testing, TrafficTest for traffic testing, CommTest for communication library testing.
	//
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// Network mode
	//
	// example:
	//
	// 2
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// Test port number.
	//
	// example:
	//
	// 23604
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// This field is empty if the TrafficModel is Fullmesh.
	TrafficTestShrink *string `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty"`
}

func (s CreateNetTestTaskShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskShrinkRequest) SetClusterId(v string) *CreateNetTestTaskShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetClusterName(v string) *CreateNetTestTaskShrinkRequest {
	s.ClusterName = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetCommTestShrink(v string) *CreateNetTestTaskShrinkRequest {
	s.CommTestShrink = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetDelayTestShrink(v string) *CreateNetTestTaskShrinkRequest {
	s.DelayTestShrink = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetNetTestType(v string) *CreateNetTestTaskShrinkRequest {
	s.NetTestType = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetNetworkMode(v string) *CreateNetTestTaskShrinkRequest {
	s.NetworkMode = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetPort(v string) *CreateNetTestTaskShrinkRequest {
	s.Port = &v
	return s
}

func (s *CreateNetTestTaskShrinkRequest) SetTrafficTestShrink(v string) *CreateNetTestTaskShrinkRequest {
	s.TrafficTestShrink = &v
	return s
}

type CreateNetTestTaskResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 启动测试任务ID，网络测试任务的唯一标志。
	//
	// example:
	//
	// dr-uf6i0tv2refv8wz*****
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
}

func (s CreateNetTestTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskResponseBody) SetRequestId(v string) *CreateNetTestTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetTestTaskResponseBody) SetTestId(v string) *CreateNetTestTaskResponseBody {
	s.TestId = &v
	return s
}

type CreateNetTestTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetTestTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetTestTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNetTestTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskResponse) SetHeaders(v map[string]*string) *CreateNetTestTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateNetTestTaskResponse) SetStatusCode(v int32) *CreateNetTestTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNetTestTaskResponse) SetBody(v *CreateNetTestTaskResponseBody) *CreateNetTestTaskResponse {
	s.Body = v
	return s
}

type CreateNodeGroupRequest struct {
	// Cluster ID
	//
	// This parameter is required.
	//
	// example:
	//
	// i118191731740041623425
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Node ID.
	//
	// This parameter is required.
	NodeGroup *CreateNodeGroupRequestNodeGroup `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty" type:"Struct"`
	// Node information
	NodeUnit map[string]interface{} `json:"NodeUnit,omitempty" xml:"NodeUnit,omitempty"`
}

func (s CreateNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequest) SetClusterId(v string) *CreateNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodeGroupRequest) SetNodeGroup(v *CreateNodeGroupRequestNodeGroup) *CreateNodeGroupRequest {
	s.NodeGroup = v
	return s
}

func (s *CreateNodeGroupRequest) SetNodeUnit(v map[string]interface{}) *CreateNodeGroupRequest {
	s.NodeUnit = v
	return s
}

type CreateNodeGroupRequestNodeGroup struct {
	// Availability Zone
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	Az                     *string `json:"Az,omitempty" xml:"Az,omitempty"`
	FileSystemMountEnabled *bool   `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// Image ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i191887641687336652616
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// test-keypair
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// Machine type
	//
	// This parameter is required.
	//
	// example:
	//
	// mock-machine-type3
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Node group description
	//
	// example:
	//
	// describe for node group
	NodeGroupDescription *string `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	// Node group name
	//
	// This parameter is required.
	//
	// example:
	//
	// PAI-LINGJUN
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// SystemDisk
	SystemDisk *CreateNodeGroupRequestNodeGroupSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// user data
	//
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateNodeGroupRequestNodeGroup) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupRequestNodeGroup) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequestNodeGroup) SetAz(v string) *CreateNodeGroupRequestNodeGroup {
	s.Az = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetFileSystemMountEnabled(v bool) *CreateNodeGroupRequestNodeGroup {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetImageId(v string) *CreateNodeGroupRequestNodeGroup {
	s.ImageId = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetKeyPairName(v string) *CreateNodeGroupRequestNodeGroup {
	s.KeyPairName = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetMachineType(v string) *CreateNodeGroupRequestNodeGroup {
	s.MachineType = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetNodeGroupDescription(v string) *CreateNodeGroupRequestNodeGroup {
	s.NodeGroupDescription = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetNodeGroupName(v string) *CreateNodeGroupRequestNodeGroup {
	s.NodeGroupName = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetSystemDisk(v *CreateNodeGroupRequestNodeGroupSystemDisk) *CreateNodeGroupRequestNodeGroup {
	s.SystemDisk = v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetUserData(v string) *CreateNodeGroupRequestNodeGroup {
	s.UserData = &v
	return s
}

type CreateNodeGroupRequestNodeGroupSystemDisk struct {
	// Disk performance level
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The performance level of the disk if the disk is an ESSD. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	//
	// Default value: PL1.
	//
	// For information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL!
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// System disk size
	//
	// example:
	//
	// 250
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateNodeGroupRequestNodeGroupSystemDisk) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupRequestNodeGroupSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) SetCategory(v string) *CreateNodeGroupRequestNodeGroupSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) SetPerformanceLevel(v string) *CreateNodeGroupRequestNodeGroupSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroupSystemDisk) SetSize(v int32) *CreateNodeGroupRequestNodeGroupSystemDisk {
	s.Size = &v
	return s
}

type CreateNodeGroupShrinkRequest struct {
	// Cluster ID
	//
	// This parameter is required.
	//
	// example:
	//
	// i118191731740041623425
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Node ID.
	//
	// This parameter is required.
	NodeGroupShrink *string `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty"`
	// Node information
	NodeUnitShrink *string `json:"NodeUnit,omitempty" xml:"NodeUnit,omitempty"`
}

func (s CreateNodeGroupShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupShrinkRequest) SetClusterId(v string) *CreateNodeGroupShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateNodeGroupShrinkRequest) SetNodeGroupShrink(v string) *CreateNodeGroupShrinkRequest {
	s.NodeGroupShrink = &v
	return s
}

func (s *CreateNodeGroupShrinkRequest) SetNodeUnitShrink(v string) *CreateNodeGroupShrinkRequest {
	s.NodeUnitShrink = &v
	return s
}

type CreateNodeGroupResponseBody struct {
	// Node group ID
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// Node group name
	//
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponseBody) SetNodeGroupId(v string) *CreateNodeGroupResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *CreateNodeGroupResponseBody) SetNodeGroupName(v string) *CreateNodeGroupResponseBody {
	s.NodeGroupName = &v
	return s
}

func (s *CreateNodeGroupResponseBody) SetRequestId(v string) *CreateNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

type CreateNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponse) SetHeaders(v map[string]*string) *CreateNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateNodeGroupResponse) SetStatusCode(v int32) *CreateNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateNodeGroupResponse) SetBody(v *CreateNodeGroupResponseBody) *CreateNodeGroupResponse {
	s.Body = v
	return s
}

type CreateSessionRequest struct {
	// Instance ID.
	//
	// example:
	//
	// e01-cn-kvw44e6dn04
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Session type corresponding to the session package.
	//
	// example:
	//
	// N	两种：
	//
	// Sol：基于串口[默认]
	//
	// Assistant：基于云助手
	SessionType *string `json:"SessionType,omitempty" xml:"SessionType,omitempty"`
	// Initiation time, 13-digit timestamp.
	//
	// example:
	//
	// 1669340937156
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s CreateSessionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionRequest) SetNodeId(v string) *CreateSessionRequest {
	s.NodeId = &v
	return s
}

func (s *CreateSessionRequest) SetSessionType(v string) *CreateSessionRequest {
	s.SessionType = &v
	return s
}

func (s *CreateSessionRequest) SetStartTime(v string) *CreateSessionRequest {
	s.StartTime = &v
	return s
}

type CreateSessionResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// 节点  ID。
	//
	// example:
	//
	// 2A59143F1
	ServerSn *string `json:"ServerSn,omitempty" xml:"ServerSn,omitempty"`
	// Session ID.
	//
	// example:
	//
	// i207023871669364793713
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// Session token.
	//
	// example:
	//
	// 03f53c719015a9ad4f4f55d66cac2dac161b18e8065ca75a3220b89de389c980
	SessionToken *string `json:"SessionToken,omitempty" xml:"SessionToken,omitempty"`
	// WebSocket address
	//
	// example:
	//
	// ws://x.x.x.x:xx/calypso_web_console
	WssEndpoint *string `json:"WssEndpoint,omitempty" xml:"WssEndpoint,omitempty"`
}

func (s CreateSessionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionResponseBody) SetRequestId(v string) *CreateSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSessionResponseBody) SetServerSn(v string) *CreateSessionResponseBody {
	s.ServerSn = &v
	return s
}

func (s *CreateSessionResponseBody) SetSessionId(v string) *CreateSessionResponseBody {
	s.SessionId = &v
	return s
}

func (s *CreateSessionResponseBody) SetSessionToken(v string) *CreateSessionResponseBody {
	s.SessionToken = &v
	return s
}

func (s *CreateSessionResponseBody) SetWssEndpoint(v string) *CreateSessionResponseBody {
	s.WssEndpoint = &v
	return s
}

type CreateSessionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSessionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateSessionResponse) SetHeaders(v map[string]*string) *CreateSessionResponse {
	s.Headers = v
	return s
}

func (s *CreateSessionResponse) SetStatusCode(v int32) *CreateSessionResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSessionResponse) SetBody(v *CreateSessionResponseBody) *CreateSessionResponse {
	s.Body = v
	return s
}

type CreateVscRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*CreateVscRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// example:
	//
	// primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s CreateVscRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateVscRequest) GoString() string {
	return s.String()
}

func (s *CreateVscRequest) SetClientToken(v string) *CreateVscRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateVscRequest) SetNodeId(v string) *CreateVscRequest {
	s.NodeId = &v
	return s
}

func (s *CreateVscRequest) SetResourceGroupId(v string) *CreateVscRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateVscRequest) SetTag(v []*CreateVscRequestTag) *CreateVscRequest {
	s.Tag = v
	return s
}

func (s *CreateVscRequest) SetVscName(v string) *CreateVscRequest {
	s.VscName = &v
	return s
}

func (s *CreateVscRequest) SetVscType(v string) *CreateVscRequest {
	s.VscType = &v
	return s
}

type CreateVscRequestTag struct {
	// example:
	//
	// key001
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVscRequestTag) String() string {
	return tea.Prettify(s)
}

func (s CreateVscRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVscRequestTag) SetKey(v string) *CreateVscRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVscRequestTag) SetValue(v string) *CreateVscRequestTag {
	s.Value = &v
	return s
}

type CreateVscResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s CreateVscResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateVscResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVscResponseBody) SetRequestId(v string) *CreateVscResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVscResponseBody) SetVscId(v string) *CreateVscResponseBody {
	s.VscId = &v
	return s
}

type CreateVscResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVscResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateVscResponse) GoString() string {
	return s.String()
}

func (s *CreateVscResponse) SetHeaders(v map[string]*string) *CreateVscResponse {
	s.Headers = v
	return s
}

func (s *CreateVscResponse) SetStatusCode(v int32) *CreateVscResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateVscResponse) SetBody(v *CreateVscResponseBody) *CreateVscResponse {
	s.Body = v
	return s
}

type DeleteClusterRequest struct {
	// Cluster ID
	//
	// This parameter is required.
	//
	// example:
	//
	// i116913051662373010974
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DeleteClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

type DeleteClusterResponseBody struct {
	// Request Id
	//
	// example:
	//
	// 0FC4A1C7-421C-5EAB-9361-4C0338EFA287
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) SetHeaders(v map[string]*string) *DeleteClusterResponse {
	s.Headers = v
	return s
}

func (s *DeleteClusterResponse) SetStatusCode(v int32) *DeleteClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteClusterResponse) SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse {
	s.Body = v
	return s
}

type DeleteNodeGroupRequest struct {
	// Cluster ID
	//
	// example:
	//
	// i114444141733395242745
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Node Group ID
	//
	// example:
	//
	// i121824791737080429819
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s DeleteNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupRequest) SetClusterId(v string) *DeleteNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodeGroupRequest) SetNodeGroupId(v string) *DeleteNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

type DeleteNodeGroupResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponseBody) SetRequestId(v string) *DeleteNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

type DeleteNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponse) SetHeaders(v map[string]*string) *DeleteNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteNodeGroupResponse) SetStatusCode(v int32) *DeleteNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteNodeGroupResponse) SetBody(v *DeleteNodeGroupResponseBody) *DeleteNodeGroupResponse {
	s.Body = v
	return s
}

type DeleteVscRequest struct {
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DeleteVscRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteVscRequest) GoString() string {
	return s.String()
}

func (s *DeleteVscRequest) SetClientToken(v string) *DeleteVscRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVscRequest) SetVscId(v string) *DeleteVscRequest {
	s.VscId = &v
	return s
}

type DeleteVscResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVscResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteVscResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVscResponseBody) SetRequestId(v string) *DeleteVscResponseBody {
	s.RequestId = &v
	return s
}

type DeleteVscResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVscResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteVscResponse) GoString() string {
	return s.String()
}

func (s *DeleteVscResponse) SetHeaders(v map[string]*string) *DeleteVscResponse {
	s.Headers = v
	return s
}

func (s *DeleteVscResponse) SetStatusCode(v int32) *DeleteVscResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteVscResponse) SetBody(v *DeleteVscResponseBody) *DeleteVscResponse {
	s.Body = v
	return s
}

type DescribeClusterRequest struct {
	// Cluster ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s DescribeClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRequest) SetClusterId(v string) *DescribeClusterRequest {
	s.ClusterId = &v
	return s
}

type DescribeClusterResponseBody struct {
	// Cluster Description
	//
	// example:
	//
	// Default cluster
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// Cluster ID
	//
	// example:
	//
	// i116913051662373010974
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster Name
	//
	// example:
	//
	// Eflo-YJ-Test-Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Cluster Type
	//
	// example:
	//
	// AckEdgePro
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Component Information
	Components []*DescribeClusterResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// Type of IP in the compute network
	//
	// example:
	//
	// IPv4
	ComputingIpVersion *string `json:"ComputingIpVersion,omitempty" xml:"ComputingIpVersion,omitempty"`
	// Creation Time
	//
	// example:
	//
	// 2022-06-08T07:05:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Cluster Number
	//
	// example:
	//
	// A2
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Network Information
	Networks *DescribeClusterResponseBodyNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
	// Number of Nodes
	//
	// example:
	//
	// 2
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Number of Node Groups
	//
	// example:
	//
	// 2
	NodeGroupCount *int64 `json:"NodeGroupCount,omitempty" xml:"NodeGroupCount,omitempty"`
	// Open Eni Jumbo Frame
	//
	// example:
	//
	// close
	OpenEniJumboFrame *string `json:"OpenEniJumboFrame,omitempty" xml:"OpenEniJumboFrame,omitempty"`
	// Cluster State
	//
	// example:
	//
	// running
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Resource Group ID
	//
	// example:
	//
	// rg-aek2k3rqlvv6ytq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Task ID
	//
	// example:
	//
	// i152609221670466904596
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Update Time
	//
	// example:
	//
	// 2022-08-23T06:36:17.000Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-0jlkqysom5dmcviymep3f
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBody) SetClusterDescription(v string) *DescribeClusterResponseBody {
	s.ClusterDescription = &v
	return s
}

func (s *DescribeClusterResponseBody) SetClusterId(v string) *DescribeClusterResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetClusterName(v string) *DescribeClusterResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeClusterResponseBody) SetClusterType(v string) *DescribeClusterResponseBody {
	s.ClusterType = &v
	return s
}

func (s *DescribeClusterResponseBody) SetComponents(v []*DescribeClusterResponseBodyComponents) *DescribeClusterResponseBody {
	s.Components = v
	return s
}

func (s *DescribeClusterResponseBody) SetComputingIpVersion(v string) *DescribeClusterResponseBody {
	s.ComputingIpVersion = &v
	return s
}

func (s *DescribeClusterResponseBody) SetCreateTime(v string) *DescribeClusterResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeClusterResponseBody) SetHpnZone(v string) *DescribeClusterResponseBody {
	s.HpnZone = &v
	return s
}

func (s *DescribeClusterResponseBody) SetNetworks(v *DescribeClusterResponseBodyNetworks) *DescribeClusterResponseBody {
	s.Networks = v
	return s
}

func (s *DescribeClusterResponseBody) SetNodeCount(v int64) *DescribeClusterResponseBody {
	s.NodeCount = &v
	return s
}

func (s *DescribeClusterResponseBody) SetNodeGroupCount(v int64) *DescribeClusterResponseBody {
	s.NodeGroupCount = &v
	return s
}

func (s *DescribeClusterResponseBody) SetOpenEniJumboFrame(v string) *DescribeClusterResponseBody {
	s.OpenEniJumboFrame = &v
	return s
}

func (s *DescribeClusterResponseBody) SetOperatingState(v string) *DescribeClusterResponseBody {
	s.OperatingState = &v
	return s
}

func (s *DescribeClusterResponseBody) SetRequestId(v string) *DescribeClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetResourceGroupId(v string) *DescribeClusterResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetTaskId(v string) *DescribeClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *DescribeClusterResponseBody) SetUpdateTime(v string) *DescribeClusterResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeClusterResponseBody) SetVpcId(v string) *DescribeClusterResponseBody {
	s.VpcId = &v
	return s
}

type DescribeClusterResponseBodyComponents struct {
	// Component ID
	//
	// example:
	//
	// i149549021660892626529
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// Component Type
	//
	// example:
	//
	// ACKEdge
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
}

func (s DescribeClusterResponseBodyComponents) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyComponents) SetComponentId(v string) *DescribeClusterResponseBodyComponents {
	s.ComponentId = &v
	return s
}

func (s *DescribeClusterResponseBodyComponents) SetComponentType(v string) *DescribeClusterResponseBodyComponents {
	s.ComponentType = &v
	return s
}

type DescribeClusterResponseBodyNetworks struct {
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s DescribeClusterResponseBodyNetworks) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponseBodyNetworks) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyNetworks) SetVpdId(v string) *DescribeClusterResponseBodyNetworks {
	s.VpdId = &v
	return s
}

type DescribeClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponse) SetHeaders(v map[string]*string) *DescribeClusterResponse {
	s.Headers = v
	return s
}

func (s *DescribeClusterResponse) SetStatusCode(v int32) *DescribeClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeClusterResponse) SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse {
	s.Body = v
	return s
}

type DescribeDiagnosticResultRequest struct {
	// Diagnostic ID
	//
	// example:
	//
	// diag-i151942361720577788844
	DiagnosticId *string `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
}

func (s DescribeDiagnosticResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultRequest) SetDiagnosticId(v string) *DescribeDiagnosticResultRequest {
	s.DiagnosticId = &v
	return s
}

type DescribeDiagnosticResultResponseBody struct {
	// Cluster ID
	//
	// example:
	//
	// i118913031696573280136
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Device creation time.
	//
	// example:
	//
	// 2024-06-15T10:17:56
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// Diagnostic ID
	//
	// example:
	//
	// diag-i155363241720059671316
	DiagnosticId *string `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
	// Diagnostic Information
	DiagnosticResults []interface{} `json:"DiagnosticResults,omitempty" xml:"DiagnosticResults,omitempty" type:"Repeated"`
	// Diagnostic State
	//
	// example:
	//
	// Fault
	DiagnosticState *string `json:"DiagnosticState,omitempty" xml:"DiagnosticState,omitempty"`
	// Diagnostic Type
	//
	// example:
	//
	// CheckByAiJobLogs
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// End time of node anomaly issues. Represented according to the ISO8601 standard, in a timezone-aware format, formatted as yyyy-MM-ddTHH:mm:ss+0800
	//
	// example:
	//
	// 2024-06-11T10:00:30
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// List of Node IDs
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosticResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultResponseBody) SetClusterId(v string) *DescribeDiagnosticResultResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetCreatedTime(v string) *DescribeDiagnosticResultResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticId(v string) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticId = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticResults(v []interface{}) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticResults = v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticState(v string) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticState = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetDiagnosticType(v string) *DescribeDiagnosticResultResponseBody {
	s.DiagnosticType = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetEndTime(v string) *DescribeDiagnosticResultResponseBody {
	s.EndTime = &v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetNodeIds(v []*string) *DescribeDiagnosticResultResponseBody {
	s.NodeIds = v
	return s
}

func (s *DescribeDiagnosticResultResponseBody) SetRequestId(v string) *DescribeDiagnosticResultResponseBody {
	s.RequestId = &v
	return s
}

type DescribeDiagnosticResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosticResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosticResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeDiagnosticResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultResponse) SetHeaders(v map[string]*string) *DescribeDiagnosticResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeDiagnosticResultResponse) SetStatusCode(v int32) *DescribeDiagnosticResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeDiagnosticResultResponse) SetBody(v *DescribeDiagnosticResultResponseBody) *DescribeDiagnosticResultResponse {
	s.Body = v
	return s
}

type DescribeInvocationsRequest struct {
	// Sets the encoding method for the `CommandContent` and `Output` fields in the returned data. Possible values:
	//
	// - PlainText: Returns the original command content and output information.
	//
	// - Base64: Returns the Base64-encoded command content and output information.
	//
	// Default value: Base64.
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Indicates whether to return the output information of the command execution in the result.
	//
	// - true: Return. In this case, you must specify at least the `InvokeId` or `InstanceId` parameter.
	//
	// - false: Do not return.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	IncludeOutput *bool `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	// Command execution ID
	//
	// This parameter is required.
	//
	// example:
	//
	// t-cd03crwys0lrls0
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// Instance ID
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) SetContentEncoding(v string) *DescribeInvocationsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeOutput(v bool) *DescribeInvocationsRequest {
	s.IncludeOutput = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeId(v string) *DescribeInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetNodeId(v string) *DescribeInvocationsRequest {
	s.NodeId = &v
	return s
}

type DescribeInvocationsResponseBody struct {
	// Script execution record object.
	Invocations *DescribeInvocationsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Struct"`
	// Request ID
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInvocationsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v *DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeInvocationsResponseBodyInvocations struct {
	// File delivery record.
	Invocation []*DescribeInvocationsResponseBodyInvocationsInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvocation(v []*DescribeInvocationsResponseBodyInvocationsInvocation) *DescribeInvocationsResponseBodyInvocations {
	s.Invocation = v
	return s
}

type DescribeInvocationsResponseBodyInvocationsInvocation struct {
	// Command content.
	//
	// - If `ContentEncoding` is set to `PlainText`, the original script content is returned.
	//
	// - If `ContentEncoding` is set to `Base64`, the Base64-encoded script content is returned.
	//
	// example:
	//
	// cnBtIC1xYSB8IGdyZXAgdnNm****
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// Command description.
	//
	// example:
	//
	// testDescription
	CommandDescription *string `json:"CommandDescription,omitempty" xml:"CommandDescription,omitempty"`
	// Command name.
	//
	// example:
	//
	// CommandTestName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// The creation time of the task.
	//
	// example:
	//
	// 2020-01-19T09:15:46Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The execution time for scheduled commands.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// The overall execution status of the command, which depends on the common execution status of all instances involved in the call. Possible values:
	//
	// - Pending: The system is validating or sending the command. If at least one instance has a command execution status of Pending, the overall status is Pending.
	//
	// - Scheduled: The scheduled command has been sent and is waiting to run. If at least one instance has a command execution status of Scheduled, the overall status is Scheduled.
	//
	// - Running: The command is running on the instance. If at least one instance has a command execution status of Running, the overall status is Running.
	//
	// - Success: The command execution status of all instances is Stopped or Success, and at least one instance\\"s command execution status is Success. The overall status is Success.
	//
	//     - For immediately executed tasks: The command has completed with an exit code of 0.
	//
	//     - For periodically executed tasks: The most recent execution was successful with an exit code of 0, and the specified times have all been completed.
	//
	// - Failed: The command execution status of all instances is Stopped or Failed. The overall status is Failed if any of the following conditions apply to the instance\\"s command execution status:
	//
	//     - Command validation failed (Invalid).
	//
	//     - Command sending failed (Aborted).
	//
	//     - Command completed but the exit code is not 0 (Failed).
	//
	//     - Command execution timed out (Timeout).
	//
	//     - Command execution encountered an error (Error).
	//
	// - Stopping: The task is being stopped. If at least one instance has a command execution status of Stopping, the overall status is Stopping.
	//
	// - Stopped: The task has been stopped. If all instances\\" command execution statuses are Stopped, the overall status is Stopped. The overall status is Stopped if the instance\\"s command execution status is any of the following:
	//
	//     - The task was canceled (Cancelled).
	//
	//     - The task was terminated (Terminated).
	//
	// - PartialFailed: Some instances succeeded and some failed. If the command execution statuses of all instances are Success, Failed, or Stopped, the overall status is PartialFailed.
	//
	// > The `InvokeStatus` in the response parameters is similar in meaning to this parameter, but it is recommended that you check this return value.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// Command execution ID.
	//
	// example:
	//
	// t-ind3k9ytvvduoe8
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// Command execution records.
	InvokeNodes *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes `json:"InvokeNodes,omitempty" xml:"InvokeNodes,omitempty" type:"Struct"`
	// The overall execution status of the command. The overall execution status depends on the common execution status of one or more instances in the execution. Possible values:
	//
	// - Running:
	//
	//     - For scheduled execution: The execution status remains ongoing until the scheduled command is manually stopped.
	//
	//     - For single execution: If there is any command process running, the overall execution status is ongoing.
	//
	// - Finished:
	//
	//     - For scheduled execution: The command process cannot be completed.
	//
	//     - For single execution: All instances have completed execution, or some instances\\" command processes are manually stopped and the rest have completed.
	//
	// - Failed:
	//
	//     - For scheduled execution: The command process cannot fail.
	//
	//     - For single execution: All instances have failed to execute.
	//
	// - Stopped: The command has been stopped.
	//
	// - Stopping: The command is being stopped.
	//
	// - PartialFailed: Partial failure; if the `InstanceId` parameter is set, this does not apply.
	//
	// example:
	//
	// Running
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// Custom parameters in the command.
	//
	// example:
	//
	// {}
	Parameters *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// 命令执行的方式。可能值：
	//
	// Once：立即执行命令。
	//
	// Period：定时执行命令。
	//
	// NextRebootOnly：当实例下一次启动时，自动执行命令。
	//
	// EveryReboot：实例每一次启动都将自动执行命令。
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// Timeout for executing the command, in seconds.
	//
	// example:
	//
	// 60
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// Username for executing the command.
	//
	// example:
	//
	// root
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// The working directory of the command on the instance.
	//
	// example:
	//
	// /home
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocation) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocation) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandContent(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandContent = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandDescription(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandDescription = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCommandName(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CommandName = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetFrequency(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Frequency = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeId(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeNodes(v *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeNodes = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetParameters(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Parameters = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetRepeatMode(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.RepeatMode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetTimeout(v int32) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Timeout = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetUsername(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.Username = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) SetWorkingDir(v string) *DescribeInvocationsResponseBodyInvocationsInvocation {
	s.WorkingDir = &v
	return s
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes struct {
	// Command execution records for nodes.
	InvokeNode []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode `json:"InvokeNode,omitempty" xml:"InvokeNode,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) SetInvokeNode(v []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes {
	s.InvokeNode = v
	return s
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode struct {
	// The start time of the command execution.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The length of the text that is truncated and discarded when the length of the `Output` field exceeds 24 KB.
	//
	// example:
	//
	// 0
	Dropped *int32 `json:"Dropped,omitempty" xml:"Dropped,omitempty"`
	// Reason code for file delivery failure. Possible values:
	//
	// - Empty: File delivery is normal.
	//
	// - NodeNotExists: The specified instance does not exist or has been released.
	//
	// - NodeReleased: The instance was released during the file delivery process.
	//
	// - NodeNotRunning: The instance was not running when the file delivery task was created.
	//
	// - AccountNotExists: The specified account does not exist.
	//
	// - ClientNotRunning: The Cloud Assistant Agent is not running.
	//
	// - ClientNotResponse: The Cloud Assistant Agent is not responding.
	//
	// - ClientIsUpgrading: The Cloud Assistant Agent is currently upgrading.
	//
	// - ClientNeedUpgrade: The Cloud Assistant Agent needs to be upgraded.
	//
	// - DeliveryTimeout: File sending timed out.
	//
	// - FileCreateFail: File creation failed.
	//
	// - FileAlreadyExists: A file with the same name already exists at the specified path.
	//
	// - FileContentInvalid: The file content is invalid. - FileNameInvalid: The file name is invalid.
	//
	// - FilePathInvalid: The file path is invalid.
	//
	// - FileAuthorityInvalid: The file permissions are invalid.
	//
	// - UserGroupNotExists: The user group specified for file sending does not exist.
	//
	// example:
	//
	// NodeNotExists：
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Details of the reason for command delivery failure or execution failure, possible values:
	//
	// - Empty: The command executed normally.
	//
	// - the specified node does not exist: The specified instance does not exist or has been released.
	//
	// - the node was released when creating the task: The instance was released during command execution.
	//
	// - the node is not running when creating the task: The instance was not in a running state when the command was executed.
	//
	// - the command is not applicable: The command is not applicable to the specified instance.
	//
	// - the specified account does not exist: The specified account does not exist.
	//
	// - the specified directory does not exist: The specified directory does not exist.
	//
	// - the cron job expression is invalid: The specified execution time expression is invalid.
	//
	// - the aliyun service is not running on the instance: The Cloud Assistant Agent is not running.
	//
	// - the aliyun service in the instance does not respond: The Cloud Assistant Agent is not responding.
	//
	// - the aliyun service in the node is upgrading now: The Cloud Assistant Agent is currently being upgraded.
	//
	// - the aliyun service in the node needs upgrade: The Cloud Assistant Agent needs an upgrade.
	//
	// - the command delivery has timed out: Command delivery has timed out.
	//
	// - the command execution has timed out: Command execution has timed out.
	//
	// - the command execution got an exception: An exception occurred during command execution.
	//
	// - the command execution was interrupted: Command execution was interrupted.
	//
	//  - the command execution exit code is not zero: Command execution completed with a non-zero exit code.
	//
	// - the specified node has been released: The instance was released during file delivery.
	//
	// example:
	//
	// the specified node does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// The exit code of the command process. Possible values:
	//
	// - For Linux instances, it is the exit code of the Shell process. - For Windows instances, it is the exit code of the Bat or PowerShell process.
	//
	// example:
	//
	// 0
	ExitCode *int32 `json:"ExitCode,omitempty" xml:"ExitCode,omitempty"`
	// Completion time.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// The command progress status for a single instance. Possible values:
	//
	// -  Pending: The system is validating or sending the command.
	//
	// -  Invalid: The specified command type or parameters are incorrect.
	//
	// -  Aborted: Failed to send the command to the instance. The instance must be running, and the command should be sent within 1 minute.
	//
	// -  Running: The command is running on the instance.
	//
	// -  Success:
	//
	//     -  For a one-time execution command: The command has completed with an exit code of 0.
	//
	//     -  For a periodic execution command: The last run was successful with an exit code of 0, and the specified period has ended.
	//
	// -  Failed:
	//
	//     -  For a one-time execution command: The command has completed with a non-zero exit code.
	//
	//     -  For a periodic execution command: The last run was successful with a non-zero exit code, and the specified period will be terminated.
	//
	// -  Error: An exception occurred during command execution, and it cannot continue.
	//
	// -  Timeout: The command execution timed out.
	//
	// -  Cancelled: The command execution action has been canceled, and the command was never started.
	//
	// -  Stopping: The task is being stopped.
	//
	// -  Terminated: The command was terminated while running.
	//
	// -  Scheduled:
	//
	//     -  For a one-time execution command: Not applicable, will not appear.
	//
	//     -  For a periodic execution command: Waiting to run.
	//
	// example:
	//
	// Pending
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// Node ID
	//
	// example:
	//
	// e01-cn-lbj36wkp70b
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The command progress status of a single instance.
	//
	// example:
	//
	// Finished
	NodeInvokeStatus *string `json:"NodeInvokeStatus,omitempty" xml:"NodeInvokeStatus,omitempty"`
	// The output information of the command.
	//
	// - If `ContentEncoding` is set to `PlainText`, the original output information is returned.
	//
	// - If `ContentEncoding` is set to `Base64`, the Base64-encoded output information is returned.
	//
	// example:
	//
	// OutPutTestmsg
	Output *string `json:"Output,omitempty" xml:"Output,omitempty"`
	// The number of times the command has been executed on this instance.
	//
	// -  If the execution mode is one-time, the value is 0 or 1.
	//
	// -  If the execution mode is periodic, the value is the number of times it has been executed.
	//
	// example:
	//
	// 0
	Repeats *int32 `json:"Repeats,omitempty" xml:"Repeats,omitempty"`
	// Start Time
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The time when `StopInvocation` was called to stop the command execution.
	//
	// example:
	//
	// 2019-12-20T06:15:55Z
	StopTime *string `json:"StopTime,omitempty" xml:"StopTime,omitempty"`
	// Whether the queried command will be automatically executed in the future. The value range is as follows:
	//
	// - true: The command, when executed by calling `RunCommand` or `InvokeCommand`, has the `RepeatMode` parameter set to `Period`, `NextRebootOnly`, or `EveryReboot`.
	//
	// - false: Queries commands in the following two states.
	//
	// - When executing a command via `RunCommand` or `InvokeCommand`, the `RepeatMode` parameter is set to `Once`.
	//
	// - Commands that have been canceled, stopped, or completed.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	Timed *string `json:"Timed,omitempty" xml:"Timed,omitempty"`
	// Update Time
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetCreationTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.CreationTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetDropped(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Dropped = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorCode(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorInfo(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetExitCode(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ExitCode = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetFinishTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.FinishTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetInvocationStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetNodeId(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.NodeId = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetNodeInvokeStatus(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.NodeInvokeStatus = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetOutput(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Output = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetRepeats(v int32) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Repeats = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetStartTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.StartTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetStopTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.StopTime = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetTimed(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.Timed = &v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetUpdateTime(v string) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.UpdateTime = &v
	return s
}

type DescribeInvocationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvocationsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) SetHeaders(v map[string]*string) *DescribeInvocationsResponse {
	s.Headers = v
	return s
}

func (s *DescribeInvocationsResponse) SetStatusCode(v int32) *DescribeInvocationsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeInvocationsResponse) SetBody(v *DescribeInvocationsResponseBody) *DescribeInvocationsResponse {
	s.Body = v
	return s
}

type DescribeNetTestResultRequest struct {
	// Test task ID.
	//
	// example:
	//
	// dr-uf6i0tv2refv8wz*****
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
}

func (s DescribeNetTestResultRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultRequest) SetTestId(v string) *DescribeNetTestResultRequest {
	s.TestId = &v
	return s
}

type DescribeNetTestResultResponseBody struct {
	// Cluster ID.
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster name.
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Fill in when the traffic test type is communication library test
	CommTest *DescribeNetTestResultResponseBodyCommTest `json:"CommTest,omitempty" xml:"CommTest,omitempty" type:"Struct"`
	// Diagnosis task creation time.
	//
	// example:
	//
	// 2024-10-15T10:25:42+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Fill in when the network test type is latency test
	DelayTest *DescribeNetTestResultResponseBodyDelayTest `json:"DelayTest,omitempty" xml:"DelayTest,omitempty" type:"Struct"`
	// Diagnosis task completion time.
	//
	// example:
	//
	// 2024-10-16T02:04Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// Network test type.
	//
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// Test port number.
	//
	// example:
	//
	// 23604
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// Request ID
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Details of the diagnosis result. Returned as a JSON string.
	//
	// example:
	//
	// {}
	ResultDetial *string `json:"ResultDetial,omitempty" xml:"ResultDetial,omitempty"`
	// Diagnosis task status. Possible values:
	//
	// - InProgress: Diagnosis in progress.
	//
	// - Finished: Diagnosis completed.
	//
	// - Failed: Diagnosis failed.
	//
	// example:
	//
	// Failed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Initiated test task ID, which is the unique identifier for the network test task.
	//
	// example:
	//
	// af35ce53-7c35-4277-834a-fbf49c316a96
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
	// This field is empty if the traffic model (TrafficModel) is Fullmesh.
	TrafficTest *DescribeNetTestResultResponseBodyTrafficTest `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty" type:"Struct"`
}

func (s DescribeNetTestResultResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBody) SetClusterId(v string) *DescribeNetTestResultResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetClusterName(v string) *DescribeNetTestResultResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetCommTest(v *DescribeNetTestResultResponseBodyCommTest) *DescribeNetTestResultResponseBody {
	s.CommTest = v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetCreationTime(v string) *DescribeNetTestResultResponseBody {
	s.CreationTime = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetDelayTest(v *DescribeNetTestResultResponseBodyDelayTest) *DescribeNetTestResultResponseBody {
	s.DelayTest = v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetFinishedTime(v string) *DescribeNetTestResultResponseBody {
	s.FinishedTime = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetNetTestType(v string) *DescribeNetTestResultResponseBody {
	s.NetTestType = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetPort(v string) *DescribeNetTestResultResponseBody {
	s.Port = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetRequestId(v string) *DescribeNetTestResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetResultDetial(v string) *DescribeNetTestResultResponseBody {
	s.ResultDetial = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetStatus(v string) *DescribeNetTestResultResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetTestId(v string) *DescribeNetTestResultResponseBody {
	s.TestId = &v
	return s
}

func (s *DescribeNetTestResultResponseBody) SetTrafficTest(v *DescribeNetTestResultResponseBodyTrafficTest) *DescribeNetTestResultResponseBody {
	s.TrafficTest = v
	return s
}

type DescribeNetTestResultResponseBodyCommTest struct {
	// Number of GPUs
	//
	// example:
	//
	// 1
	GPUNum *string `json:"GPUNum,omitempty" xml:"GPUNum,omitempty"`
	// Resource ID
	Hosts []*DescribeNetTestResultResponseBodyCommTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// Communication library model
	//
	// example:
	//
	// intention_v4
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Communication library test category: ACCL or NCCL
	//
	// example:
	//
	// ACCL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNetTestResultResponseBodyCommTest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyCommTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetGPUNum(v string) *DescribeNetTestResultResponseBodyCommTest {
	s.GPUNum = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetHosts(v []*DescribeNetTestResultResponseBodyCommTestHosts) *DescribeNetTestResultResponseBodyCommTest {
	s.Hosts = v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetModel(v string) *DescribeNetTestResultResponseBodyCommTest {
	s.Model = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTest) SetType(v string) *DescribeNetTestResultResponseBodyCommTest {
	s.Type = &v
	return s
}

type DescribeNetTestResultResponseBodyCommTestHosts struct {
	// IP address
	//
	// example:
	//
	// 169.253.253.15
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Resource ID
	//
	// example:
	//
	// i111670831721110797708
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 服务名称。
	//
	// example:
	//
	// VBw
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyCommTestHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyCommTestHosts) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) SetIP(v string) *DescribeNetTestResultResponseBodyCommTestHosts {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) SetResourceId(v string) *DescribeNetTestResultResponseBodyCommTestHosts {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) SetServerName(v string) *DescribeNetTestResultResponseBodyCommTestHosts {
	s.ServerName = &v
	return s
}

type DescribeNetTestResultResponseBodyDelayTest struct {
	// Input the hosts of the test nodes
	Hosts []*DescribeNetTestResultResponseBodyDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s DescribeNetTestResultResponseBodyDelayTest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyDelayTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyDelayTest) SetHosts(v []*DescribeNetTestResultResponseBodyDelayTestHosts) *DescribeNetTestResultResponseBodyDelayTest {
	s.Hosts = v
	return s
}

type DescribeNetTestResultResponseBodyDelayTestHosts struct {
	// Network card bond interface
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// Node IP
	//
	// example:
	//
	// 125.210.225.48
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Resource ID
	//
	// example:
	//
	// e01-cn-bcd3u1aee06
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name
	//
	// example:
	//
	// NQU
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyDelayTestHosts) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyDelayTestHosts) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetBond(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.Bond = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetIP(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetResourceId(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) SetServerName(v string) *DescribeNetTestResultResponseBodyDelayTestHosts {
	s.ServerName = &v
	return s
}

type DescribeNetTestResultResponseBodyTrafficTest struct {
	// Resource ID.
	Clients []*DescribeNetTestResultResponseBodyTrafficTestClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// Duration of the workflow task in seconds.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// For RDMA, enter True/False; for TCP, this field is empty.
	//
	// example:
	//
	// False
	GDR *string `json:"GDR,omitempty" xml:"GDR,omitempty"`
	// Network protocol, either RDMA or TCP.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// For TCP, enter the number of concurrent test connections; for RDMA, enter the configured QP value.
	//
	// example:
	//
	// 1
	QP *int64 `json:"QP,omitempty" xml:"QP,omitempty"`
	// List of servers
	Servers []*DescribeNetTestResultResponseBodyTrafficTestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// Traffic model, either MTON or Fullmesh.
	//
	// example:
	//
	// Fullmesh
	TrafficModel *string `json:"TrafficModel,omitempty" xml:"TrafficModel,omitempty"`
}

func (s DescribeNetTestResultResponseBodyTrafficTest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetClients(v []*DescribeNetTestResultResponseBodyTrafficTestClients) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Clients = v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetDuration(v int64) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Duration = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetGDR(v string) *DescribeNetTestResultResponseBodyTrafficTest {
	s.GDR = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetProtocol(v string) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Protocol = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetQP(v int64) *DescribeNetTestResultResponseBodyTrafficTest {
	s.QP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetServers(v []*DescribeNetTestResultResponseBodyTrafficTestServers) *DescribeNetTestResultResponseBodyTrafficTest {
	s.Servers = v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) SetTrafficModel(v string) *DescribeNetTestResultResponseBodyTrafficTest {
	s.TrafficModel = &v
	return s
}

type DescribeNetTestResultResponseBodyTrafficTestClients struct {
	// Network card bond interface
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// Node IP
	//
	// example:
	//
	// 192.168.1.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Resource ID
	//
	// example:
	//
	// e01-cn-20s41p6cx01
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// 服务名称。
	//
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyTrafficTestClients) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTestClients) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetBond(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.Bond = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetIP(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetResourceId(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) SetServerName(v string) *DescribeNetTestResultResponseBodyTrafficTestClients {
	s.ServerName = &v
	return s
}

type DescribeNetTestResultResponseBodyTrafficTestServers struct {
	// Network card bond interface
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// Node IP
	//
	// example:
	//
	// 47.121.110.190
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Resource ID
	//
	// example:
	//
	// e01-cn-wwo3etaqu0b
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name
	//
	// example:
	//
	// xMv
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s DescribeNetTestResultResponseBodyTrafficTestServers) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTestServers) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetBond(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.Bond = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetIP(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.IP = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetResourceId(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.ResourceId = &v
	return s
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) SetServerName(v string) *DescribeNetTestResultResponseBodyTrafficTestServers {
	s.ServerName = &v
	return s
}

type DescribeNetTestResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetTestResultResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNetTestResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponse) SetHeaders(v map[string]*string) *DescribeNetTestResultResponse {
	s.Headers = v
	return s
}

func (s *DescribeNetTestResultResponse) SetStatusCode(v int32) *DescribeNetTestResultResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNetTestResultResponse) SetBody(v *DescribeNetTestResultResponseBody) *DescribeNetTestResultResponse {
	s.Body = v
	return s
}

type DescribeNodeRequest struct {
	// Node ID
	//
	// This parameter is required.
	//
	// example:
	//
	// mock-sn-2060
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeNodeRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeRequest) SetNodeId(v string) *DescribeNodeRequest {
	s.NodeId = &v
	return s
}

type DescribeNodeResponseBody struct {
	// Cluster ID
	//
	// example:
	//
	// i116913051662373010974
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster name
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Creation time
	//
	// example:
	//
	// 2022-09-30T03:35:53Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Disk infos
	Disks []*DescribeNodeResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Repeated"`
	// Expiration time
	//
	// example:
	//
	// 2022-06-23T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// 是否支持文件存储挂载
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// Hostname
	//
	// example:
	//
	// 31d38530-241e-11ed-bc63-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Cluster number
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Image ID
	//
	// example:
	//
	// i190297201634099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Image name
	//
	// example:
	//
	// Centos7.9_all_0811
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// Machine type
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Network information
	Networks []*DescribeNodeResponseBodyNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
	// Node group ID
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// Node group name
	//
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// Node ID
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Node status
	//
	// example:
	//
	// Using
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// Request ID
	//
	// example:
	//
	// AC4F0004-7BCE-52E0-891B-CAC7D64E3368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-acfmywpvugkh7kq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Unique machine identifier
	//
	// example:
	//
	// sag42ckf4jx
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The script by user defined
	//
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Zone ID
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeNodeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBody) SetClusterId(v string) *DescribeNodeResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetClusterName(v string) *DescribeNodeResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeNodeResponseBody) SetCreateTime(v string) *DescribeNodeResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeNodeResponseBody) SetDisks(v []*DescribeNodeResponseBodyDisks) *DescribeNodeResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeNodeResponseBody) SetExpiredTime(v string) *DescribeNodeResponseBody {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeNodeResponseBody) SetFileSystemMountEnabled(v bool) *DescribeNodeResponseBody {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *DescribeNodeResponseBody) SetHostname(v string) *DescribeNodeResponseBody {
	s.Hostname = &v
	return s
}

func (s *DescribeNodeResponseBody) SetHpnZone(v string) *DescribeNodeResponseBody {
	s.HpnZone = &v
	return s
}

func (s *DescribeNodeResponseBody) SetImageId(v string) *DescribeNodeResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetImageName(v string) *DescribeNodeResponseBody {
	s.ImageName = &v
	return s
}

func (s *DescribeNodeResponseBody) SetMachineType(v string) *DescribeNodeResponseBody {
	s.MachineType = &v
	return s
}

func (s *DescribeNodeResponseBody) SetNetworks(v []*DescribeNodeResponseBodyNetworks) *DescribeNodeResponseBody {
	s.Networks = v
	return s
}

func (s *DescribeNodeResponseBody) SetNodeGroupId(v string) *DescribeNodeResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetNodeGroupName(v string) *DescribeNodeResponseBody {
	s.NodeGroupName = &v
	return s
}

func (s *DescribeNodeResponseBody) SetNodeId(v string) *DescribeNodeResponseBody {
	s.NodeId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetOperatingState(v string) *DescribeNodeResponseBody {
	s.OperatingState = &v
	return s
}

func (s *DescribeNodeResponseBody) SetRequestId(v string) *DescribeNodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetResourceGroupId(v string) *DescribeNodeResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeNodeResponseBody) SetSn(v string) *DescribeNodeResponseBody {
	s.Sn = &v
	return s
}

func (s *DescribeNodeResponseBody) SetUserData(v string) *DescribeNodeResponseBody {
	s.UserData = &v
	return s
}

func (s *DescribeNodeResponseBody) SetZoneId(v string) *DescribeNodeResponseBody {
	s.ZoneId = &v
	return s
}

type DescribeNodeResponseBodyDisks struct {
	// The category of the disk.
	//
	// 	- cloud_ssd: all-flash disk.
	//
	// example:
	//
	// cloud_essd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-bp1fi88ryk4yah8a6yos
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The performance level of the ESSD. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the disk. Unit: GiB.
	//
	// example:
	//
	// 100
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The type of the disk. Valid values:
	//
	// 	- system: system disk
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeNodeResponseBodyDisks) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBodyDisks) SetCategory(v string) *DescribeNodeResponseBodyDisks {
	s.Category = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetDiskId(v string) *DescribeNodeResponseBodyDisks {
	s.DiskId = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetPerformanceLevel(v string) *DescribeNodeResponseBodyDisks {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetSize(v int32) *DescribeNodeResponseBodyDisks {
	s.Size = &v
	return s
}

func (s *DescribeNodeResponseBodyDisks) SetType(v string) *DescribeNodeResponseBodyDisks {
	s.Type = &v
	return s
}

type DescribeNodeResponseBodyNetworks struct {
	// Network interface port information
	//
	// example:
	//
	// Bond0
	BondName *string `json:"BondName,omitempty" xml:"BondName,omitempty"`
	// Machine IP
	//
	// example:
	//
	// 47.254.235.44
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// Cluster subnet ID
	//
	// example:
	//
	// vsw-uf68v51fldm5egmui5a6k
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// Cluster network ID
	//
	// example:
	//
	// vpd-xcuhjyrj
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s DescribeNodeResponseBodyNetworks) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeResponseBodyNetworks) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBodyNetworks) SetBondName(v string) *DescribeNodeResponseBodyNetworks {
	s.BondName = &v
	return s
}

func (s *DescribeNodeResponseBodyNetworks) SetIp(v string) *DescribeNodeResponseBodyNetworks {
	s.Ip = &v
	return s
}

func (s *DescribeNodeResponseBodyNetworks) SetSubnetId(v string) *DescribeNodeResponseBodyNetworks {
	s.SubnetId = &v
	return s
}

func (s *DescribeNodeResponseBodyNetworks) SetVpdId(v string) *DescribeNodeResponseBodyNetworks {
	s.VpdId = &v
	return s
}

type DescribeNodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeNodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponse) SetHeaders(v map[string]*string) *DescribeNodeResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeResponse) SetStatusCode(v int32) *DescribeNodeResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeResponse) SetBody(v *DescribeNodeResponseBody) *DescribeNodeResponse {
	s.Body = v
	return s
}

type DescribeRegionsRequest struct {
	// Filter the returned results based on Chinese, English, and Japanese. For more information, see RFC7231. Valid values:
	//
	// zh-CN
	//
	// en-US
	//
	// Default value: zh-CN
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeRegionsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeRegionsResponseBody struct {
	// List of region information.
	Regions []*DescribeRegionsResponseBodyRegions `json:"Regions,omitempty" xml:"Regions,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 1D2FBB36-C39B-5EBB-9928-FCC1A236D65D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRegionsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

type DescribeRegionsResponseBodyRegions struct {
	// Region name
	//
	// example:
	//
	// Hang Zhou
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// region id
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s DescribeRegionsResponseBodyRegions) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) SetHeaders(v map[string]*string) *DescribeRegionsResponse {
	s.Headers = v
	return s
}

func (s *DescribeRegionsResponse) SetStatusCode(v int32) *DescribeRegionsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeRegionsResponse) SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse {
	s.Body = v
	return s
}

type DescribeSendFileResultsRequest struct {
	// Command execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// t-bj038i0d6r8zoqo
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// Node ID
	//
	// example:
	//
	// e01-cn-zvp2tgykr08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s DescribeSendFileResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsRequest) SetInvokeId(v string) *DescribeSendFileResultsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetNodeId(v string) *DescribeSendFileResultsRequest {
	s.NodeId = &v
	return s
}

type DescribeSendFileResultsResponseBody struct {
	// Record of file distribution.
	Invocations *DescribeSendFileResultsResponseBodyInvocations `json:"Invocations,omitempty" xml:"Invocations,omitempty" type:"Struct"`
	// ID of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Total number of commands.
	//
	// example:
	//
	// 1
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeSendFileResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBody) SetInvocations(v *DescribeSendFileResultsResponseBodyInvocations) *DescribeSendFileResultsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetRequestId(v string) *DescribeSendFileResultsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSendFileResultsResponseBody) SetTotalCount(v string) *DescribeSendFileResultsResponseBody {
	s.TotalCount = &v
	return s
}

type DescribeSendFileResultsResponseBodyInvocations struct {
	// Command execution ID.
	Invocation []*DescribeSendFileResultsResponseBodyInvocationsInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsResponseBodyInvocations) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocations) SetInvocation(v []*DescribeSendFileResultsResponseBodyInvocationsInvocation) *DescribeSendFileResultsResponseBodyInvocations {
	s.Invocation = v
	return s
}

type DescribeSendFileResultsResponseBodyInvocationsInvocation struct {
	// Output information after command execution.
	//
	// If ContentEncoding is specified as PlainText, the original output information is returned.
	//
	// If ContentEncoding is specified as Base64, the Base64 encoded output information is returned.
	//
	// example:
	//
	// Base64
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// File content type.
	//
	// PlainText: Plain text.
	//
	// Base64: Base64 encoding.
	//
	// The default value is PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Creation time of the distribution.
	//
	// example:
	//
	// 2023-04-10T10:53:46.156+08:00
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Command description.
	//
	// example:
	//
	// 描述信息。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The user group of the file.
	//
	// example:
	//
	// root
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// File permissions.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file.
	//
	// example:
	//
	// root
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// Overall status of the file distribution. The overall status depends on the common execution status of all instances involved in this distribution, possible values are:
	//
	// - Pending: The system is verifying or distributing the file. If at least one instance has a file distribution status of Pending, the overall execution status will be Pending.
	//
	// - Running: The file is being distributed on the instance. If at least one instance has a file distribution status of Running, the overall execution status will be Running.
	//
	// - Success: All instances have a file distribution status of Success, then the overall execution status will be Success.
	//
	// - Failed: All instances have a file distribution status of Failed, then the overall execution status will be Failed. If any of the following conditions occur for the file distribution status on an instance, the return value will be Failed:
	//
	//     - The specified file parameters are incorrect, verification failed (Invalid).
	//
	//     - Failed to distribute the file to the instance (Aborted).
	//
	//     - The file creation failed within the instance (Failed).
	//
	//     - The file distribution timed out (Timeout).
	//
	//     - An exception occurred during file distribution and could not continue (Error).
	//
	// - PartialFailed: Some instances successfully received the file while others failed. If the file distribution status of all instances is either Success or Failed, the overall execution status will be PartialFailed.
	//
	// example:
	//
	// Pending
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// Record of file distribution.
	InvokeNodes *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes `json:"InvokeNodes,omitempty" xml:"InvokeNodes,omitempty" type:"Struct"`
	// Name of the file distribution.
	//
	// example:
	//
	// test
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Number of nodes
	//
	// example:
	//
	// 3
	NodeCount *int32 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Whether to overwrite the file if a file with the same name already exists in the target directory.
	//
	// - true: Overwrite.
	//
	// - false: Do not overwrite.
	//
	// The default value is false.
	//
	// example:
	//
	// true
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// Target path.
	//
	// example:
	//
	// /home/user
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocation) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocation) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetContent(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Content = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetContentType(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.ContentType = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetCreationTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.CreationTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetDescription(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Description = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetFileGroup(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.FileGroup = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetFileMode(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.FileMode = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetFileOwner(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.FileOwner = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetInvocationStatus(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetInvokeNodes(v *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.InvokeNodes = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetName(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Name = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetNodeCount(v int32) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.NodeCount = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetOverwrite(v bool) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.Overwrite = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) SetTargetDir(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocation {
	s.TargetDir = &v
	return s
}

type DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes struct {
	// Record of file distribution for the node.
	InvokeNode []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode `json:"InvokeNode,omitempty" xml:"InvokeNode,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) SetInvokeNode(v []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes {
	s.InvokeNode = v
	return s
}

type DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode struct {
	// The creation time of the file distribution task.
	//
	// example:
	//
	// 2023-02-06T07:12:50Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The failure reason code for file distribution. Possible values:
	//
	// - Empty: The file was distributed normally.
	//
	// - NodeNotExists: The specified instance does not exist or has been released.
	//
	// - NodeReleased: The instance was released during the file distribution process.
	//
	// - NodeNotRunning: The instance was not running when the file distribution task was created.
	//
	// - AccountNotExists: The specified account does not exist.
	//
	// - ClientNotRunning: The Cloud Assistant Agent is not running.
	//
	// - ClientNotResponse: The Cloud Assistant Agent is not responding.
	//
	// - ClientIsUpgrading: The Cloud Assistant Agent is currently being upgraded.
	//
	// - ClientNeedUpgrade: The Cloud Assistant Agent needs to be upgraded.
	//
	// - DeliveryTimeout: File delivery timed out.
	//
	// - FileCreateFail: Failed to create the file.
	//
	// - FileAlreadyExists: A file with the same name already exists at the specified path.
	//
	// - FileContentInvalid: The file content is invalid.
	//
	// - FileNameInvalid: The file name is invalid.
	//
	// - FilePathInvalid: The file path is invalid.
	//
	// - FileAuthorityInvalid: The file permissions are invalid.
	//
	// - UserGroupNotExists: The user group specified for file delivery does not exist.
	//
	// example:
	//
	// AccountNotExists
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// Details of the reason for command delivery failure or execution failure, possible values:
	//
	// - Empty: The command executed normally.
	//
	// - the specified instance does not exist: The specified instance does not exist or has been released.
	//
	// - the node has been released when creating task: The instance was released during the command execution.
	//
	// - the node is not running when creating task: The instance was not in a running state when the command was executed.
	//
	// - the command is not applicable: The command is not applicable to the specified instance.
	//
	// - the specified account does not exist: The specified account does not exist.
	//
	// - the specified directory does not exist: The specified directory does not exist.
	//
	// - the cron job expression is invalid: The specified execution time expression is invalid.
	//
	// - the aliyun service is not running on the instance: The Cloud Assistant Agent is not running.
	//
	// - the aliyun service in the instance does not respond: The Cloud Assistant Agent is not responding.
	//
	// - the aliyun service in the node is upgrading now: The Cloud Assistant Agent is currently being upgraded.
	//
	// - the aliyun service in the node needs upgrade: The Cloud Assistant Agent needs an upgrade.
	//
	// - the command delivery has timed out: Command delivery has timed out.
	//
	// - the command execution has timed out: Command execution has timed out.
	//
	// - the command execution got an exception: An exception occurred during command execution.
	//
	// - the command execution has been interrupted: The command execution was interrupted.
	//
	// - the command execution exit code is not zero: The command execution completed with a non-zero exit code.
	//
	// - the specified instance has been released: The instance was released during file delivery.
	//
	// example:
	//
	// the specified instance does not exists
	ErrorInfo *string `json:"ErrorInfo,omitempty" xml:"ErrorInfo,omitempty"`
	// Completion time, format: "2020-05-22T17:04:18".
	//
	// example:
	//
	// 2023-04-10T10:53:46.156+08:00
	FinishTime *string `json:"FinishTime,omitempty" xml:"FinishTime,omitempty"`
	// Status of the task on a single instance. Possible values:
	//
	// - Pending: The system is validating or distributing the file.
	//
	// - Invalid: The specified file parameters are incorrect, and validation failed.
	//
	// - Running: The file is being distributed to the instance.
	//
	// - Aborted: Failed to distribute the file to the instance.
	//
	// - Success: The file distribution is complete.
	//
	// - Failed: The file creation failed within the instance.
	//
	// - Error: An exception occurred during file distribution and could not continue.
	//
	// - Timeout: The file distribution timed out.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// Node ID.
	//
	// example:
	//
	// e01-cn-9lb3c15m81j
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Start Time
	//
	// example:
	//
	// 2023-03-30T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Update Time
	//
	// example:
	//
	// 2023-03-30T16:00:00Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetCreationTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.CreationTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorCode(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorCode = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetErrorInfo(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.ErrorInfo = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetFinishTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.FinishTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetInvocationStatus(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetNodeId(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.NodeId = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetStartTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.StartTime = &v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) SetUpdateTime(v string) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	s.UpdateTime = &v
	return s
}

type DescribeSendFileResultsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSendFileResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSendFileResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeSendFileResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponse) SetHeaders(v map[string]*string) *DescribeSendFileResultsResponse {
	s.Headers = v
	return s
}

func (s *DescribeSendFileResultsResponse) SetStatusCode(v int32) *DescribeSendFileResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSendFileResultsResponse) SetBody(v *DescribeSendFileResultsResponseBody) *DescribeSendFileResultsResponse {
	s.Body = v
	return s
}

type DescribeTaskRequest struct {
	// Task ID
	//
	// This parameter is required.
	//
	// example:
	//
	// i156331731670384438138
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskRequest) SetTaskId(v string) *DescribeTaskRequest {
	s.TaskId = &v
	return s
}

type DescribeTaskResponseBody struct {
	// Cluster ID
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster Name
	//
	// example:
	//
	// Standard_Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Start Time
	//
	// example:
	//
	// 2022-11-30T02:00:00.852Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Task Failure Message
	//
	// example:
	//
	// Releasing [prod_main_mid_26e234cf] in region [cn-beijing] with weight [0]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// List of node IDs
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// A7FD7411-9395-52E8-AF42-EB3A4A55446D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Execution Steps
	Steps []*DescribeTaskResponseBodySteps `json:"Steps,omitempty" xml:"Steps,omitempty" type:"Repeated"`
	// Task State
	//
	// example:
	//
	// running
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// Task Type
	//
	// example:
	//
	// cut_cluster
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Update Time
	//
	// example:
	//
	// 2022-11-30T03:40:14.852Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBody) SetClusterId(v string) *DescribeTaskResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetClusterName(v string) *DescribeTaskResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeTaskResponseBody) SetCreateTime(v string) *DescribeTaskResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeTaskResponseBody) SetMessage(v string) *DescribeTaskResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBody) SetNodeIds(v []*string) *DescribeTaskResponseBody {
	s.NodeIds = v
	return s
}

func (s *DescribeTaskResponseBody) SetRequestId(v string) *DescribeTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTaskResponseBody) SetSteps(v []*DescribeTaskResponseBodySteps) *DescribeTaskResponseBody {
	s.Steps = v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskState(v string) *DescribeTaskResponseBody {
	s.TaskState = &v
	return s
}

func (s *DescribeTaskResponseBody) SetTaskType(v string) *DescribeTaskResponseBody {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskResponseBody) SetUpdateTime(v string) *DescribeTaskResponseBody {
	s.UpdateTime = &v
	return s
}

type DescribeTaskResponseBodySteps struct {
	// Step Failure Message
	//
	// example:
	//
	// get taskinfo failed
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Stage Tag
	//
	// example:
	//
	// Node scaling
	StageTag *string `json:"StageTag,omitempty" xml:"StageTag,omitempty"`
	// Start Time
	//
	// example:
	//
	// 2022-11-30T2:00:00.852Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// Step Name
	//
	// example:
	//
	// create_vpd
	StepName *string `json:"StepName,omitempty" xml:"StepName,omitempty"`
	// Step Execution State
	//
	// example:
	//
	// execution_success
	StepState *string `json:"StepState,omitempty" xml:"StepState,omitempty"`
	// Step Type
	//
	// example:
	//
	// normal
	StepType *string `json:"StepType,omitempty" xml:"StepType,omitempty"`
	// Subtasks
	SubTasks []*DescribeTaskResponseBodyStepsSubTasks `json:"SubTasks,omitempty" xml:"SubTasks,omitempty" type:"Repeated"`
	// Update Time
	//
	// example:
	//
	// 2022-11-30T02:20:14.852Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTaskResponseBodySteps) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponseBodySteps) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBodySteps) SetMessage(v string) *DescribeTaskResponseBodySteps {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStageTag(v string) *DescribeTaskResponseBodySteps {
	s.StageTag = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStartTime(v string) *DescribeTaskResponseBodySteps {
	s.StartTime = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStepName(v string) *DescribeTaskResponseBodySteps {
	s.StepName = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStepState(v string) *DescribeTaskResponseBodySteps {
	s.StepState = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetStepType(v string) *DescribeTaskResponseBodySteps {
	s.StepType = &v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetSubTasks(v []*DescribeTaskResponseBodyStepsSubTasks) *DescribeTaskResponseBodySteps {
	s.SubTasks = v
	return s
}

func (s *DescribeTaskResponseBodySteps) SetUpdateTime(v string) *DescribeTaskResponseBodySteps {
	s.UpdateTime = &v
	return s
}

type DescribeTaskResponseBodyStepsSubTasks struct {
	// Creation Time
	//
	// example:
	//
	// 2022-11-30T2:00:00.852Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Subtask Failure Message
	//
	// example:
	//
	// Releasing [prod_main_mid_26e234cf] in region [cn-beijing] with weight [0]
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Task ID
	//
	// example:
	//
	// i158805051661047928377
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Task Execution State
	//
	// example:
	//
	// running
	TaskState *string `json:"TaskState,omitempty" xml:"TaskState,omitempty"`
	// Task Type
	//
	// example:
	//
	// cut_node_sub_task
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
	// Update Time
	//
	// example:
	//
	// 2022-11-30T02:20:14.852Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s DescribeTaskResponseBodyStepsSubTasks) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponseBodyStepsSubTasks) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetCreateTime(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.CreateTime = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetMessage(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.Message = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetTaskId(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetTaskState(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.TaskState = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetTaskType(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.TaskType = &v
	return s
}

func (s *DescribeTaskResponseBodyStepsSubTasks) SetUpdateTime(v string) *DescribeTaskResponseBodyStepsSubTasks {
	s.UpdateTime = &v
	return s
}

type DescribeTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponse) SetHeaders(v map[string]*string) *DescribeTaskResponse {
	s.Headers = v
	return s
}

func (s *DescribeTaskResponse) SetStatusCode(v int32) *DescribeTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeTaskResponse) SetBody(v *DescribeTaskResponseBody) *DescribeTaskResponse {
	s.Body = v
	return s
}

type DescribeVscRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DescribeVscRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscRequest) GoString() string {
	return s.String()
}

func (s *DescribeVscRequest) SetVscId(v string) *DescribeVscRequest {
	s.VscId = &v
	return s
}

type DescribeVscResponseBody struct {
	// example:
	//
	// e01-cn-kvw44e6dn04
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// rg-aek2k3rqlvv6ytq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// VscId
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// example:
	//
	// primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s DescribeVscResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVscResponseBody) SetNodeId(v string) *DescribeVscResponseBody {
	s.NodeId = &v
	return s
}

func (s *DescribeVscResponseBody) SetRequestId(v string) *DescribeVscResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVscResponseBody) SetResourceGroupId(v string) *DescribeVscResponseBody {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVscResponseBody) SetStatus(v string) *DescribeVscResponseBody {
	s.Status = &v
	return s
}

func (s *DescribeVscResponseBody) SetVscId(v string) *DescribeVscResponseBody {
	s.VscId = &v
	return s
}

func (s *DescribeVscResponseBody) SetVscName(v string) *DescribeVscResponseBody {
	s.VscName = &v
	return s
}

func (s *DescribeVscResponseBody) SetVscType(v string) *DescribeVscResponseBody {
	s.VscType = &v
	return s
}

type DescribeVscResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVscResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeVscResponse) GoString() string {
	return s.String()
}

func (s *DescribeVscResponse) SetHeaders(v map[string]*string) *DescribeVscResponse {
	s.Headers = v
	return s
}

func (s *DescribeVscResponse) SetStatusCode(v int32) *DescribeVscResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeVscResponse) SetBody(v *DescribeVscResponseBody) *DescribeVscResponse {
	s.Body = v
	return s
}

type DescribeZonesRequest struct {
	// Filter the returned results based on Chinese or English. For more information, see RFC7231. Valid values:
	//
	// zh-CN
	//
	// en-US
	//
	// Default value: zh-CN
	//
	// example:
	//
	// zh-CN
	AcceptLanguage *string `json:"AcceptLanguage,omitempty" xml:"AcceptLanguage,omitempty"`
}

func (s DescribeZonesRequest) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

type DescribeZonesResponseBody struct {
	// Request ID
	//
	// example:
	//
	// E9116F2D-82F8-501E-9ADB-2BE0C02B6A84
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// List of available zones
	Zones []*DescribeZonesResponseBodyZones `json:"Zones,omitempty" xml:"Zones,omitempty" type:"Repeated"`
}

func (s DescribeZonesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

type DescribeZonesResponseBodyZones struct {
	// Zone name
	//
	// example:
	//
	// Hang Zhou
	LocalName *string `json:"LocalName,omitempty" xml:"LocalName,omitempty"`
	// Zone ID
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeZonesResponseBodyZones) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) SetLocalName(v string) *DescribeZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZonesResponse) String() string {
	return tea.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) SetHeaders(v map[string]*string) *DescribeZonesResponse {
	s.Headers = v
	return s
}

func (s *DescribeZonesResponse) SetStatusCode(v int32) *DescribeZonesResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeZonesResponse) SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse {
	s.Body = v
	return s
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
	// IP allocation policy combination: Each policy can only choose one type, and multiple policies can be combined
	IpAllocationPolicy []*ExtendClusterRequestIpAllocationPolicy `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty" type:"Repeated"`
	// Node Group
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
	return tea.Prettify(s)
}

func (s ExtendClusterRequest) GoString() string {
	return s.String()
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

type ExtendClusterRequestIpAllocationPolicy struct {
	// Specify the cluster subnet ID based on the bond name
	BondPolicy *ExtendClusterRequestIpAllocationPolicyBondPolicy `json:"BondPolicy,omitempty" xml:"BondPolicy,omitempty" type:"Struct"`
	// Machine type allocation policy
	MachineTypePolicy []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicy `json:"MachineTypePolicy,omitempty" xml:"MachineTypePolicy,omitempty" type:"Repeated"`
	// Node allocation policy
	NodePolicy []*ExtendClusterRequestIpAllocationPolicyNodePolicy `json:"NodePolicy,omitempty" xml:"NodePolicy,omitempty" type:"Repeated"`
}

func (s ExtendClusterRequestIpAllocationPolicy) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicy) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyBondPolicy) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) SetBondDefaultSubnet(v string) *ExtendClusterRequestIpAllocationPolicyBondPolicy {
	s.BondDefaultSubnet = &v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) SetBonds(v []*ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) *ExtendClusterRequestIpAllocationPolicyBondPolicy {
	s.Bonds = v
	return s
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
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) SetName(v string) *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds {
	s.Name = &v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) SetSubnet(v string) *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds {
	s.Subnet = &v
	return s
}

type ExtendClusterRequestIpAllocationPolicyMachineTypePolicy struct {
	// Bond information
	Bonds []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	// Machine Type
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) SetBonds(v []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy {
	s.Bonds = v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) SetMachineType(v string) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy {
	s.MachineType = &v
	return s
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
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) SetName(v string) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds {
	s.Name = &v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) SetSubnet(v string) *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds {
	s.Subnet = &v
	return s
}

type ExtendClusterRequestIpAllocationPolicyNodePolicy struct {
	// Bond information
	Bonds []*ExtendClusterRequestIpAllocationPolicyNodePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	// Host name
	//
	// example:
	//
	// i22c11282.eu95sqa
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Node ID
	//
	// example:
	//
	// i-3fdodw2
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicy) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicy) GoString() string {
	return s.String()
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
	return tea.Prettify(s)
}

func (s ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) SetName(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds {
	s.Name = &v
	return s
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) SetSubnet(v string) *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds {
	s.Subnet = &v
	return s
}

type ExtendClusterRequestNodeGroups struct {
	// Number of nodes to purchase. Value range: 0–500.
	//
	// If the Amount parameter is set to 0, no nodes will be purchased. Existing nodes will be used for scaling.
	//
	// If the Amount parameter is set to 1–500, the specified number of nodes will be purchased and used for scaling.
	//
	// Default value: 0
	//
	// example:
	//
	// 4
	Amount *int64 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// Whether to enable auto-renewal for purchased nodes.
	//
	// Conditions: This parameter takes effect only when the Amount parameter is set to a non-zero value and the ChargeType is PrePaid.
	//
	// Valid values:
	//
	// True: Enable auto-renewal.
	//
	// False: Disable auto-renewal.
	//
	// Default value: False
	//
	// example:
	//
	// True
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// Billing method for nodes.
	//
	// This parameter takes effect only when the Amount parameter is set to a value other than 0.
	//
	// Valid values:
	//
	// PrePaid: Subscription (prepaid).
	//
	// PostPaid: Pay-as-you-go (postpaid).
	//
	// Default value: PrePaid
	//
	// example:
	//
	// PostPaid
	ChargeType *string `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// The hostnames of purchased nodes.
	//
	// This parameter takes effect only when the Amount parameter is set to a non-zero value.
	Hostnames []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	// The login password of node.
	//
	// example:
	//
	// Addk(*78
	LoginPassword *string `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// Node Group ID
	//
	// example:
	//
	// i16d4883a46cbadeb4bc9
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The tag of node
	NodeTag []*ExtendClusterRequestNodeGroupsNodeTag `json:"NodeTag,omitempty" xml:"NodeTag,omitempty" type:"Repeated"`
	// List of Nodes
	Nodes []*ExtendClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// Purchase duration for nodes (unit: month).
	//
	// Valid values: 1, 6, 12, 24, 36, 48.
	//
	// Conditions: This parameter takes effect only when the Amount parameter is set to a non-zero value and the ChargeType is PrePaid.
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
	// VSwitch Id
	//
	// example:
	//
	// vsw-0jly2d537ejphyq6h13ke
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// Vpc Id
	//
	// example:
	//
	// vpc-zq1econyv63tvyci5hefw
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Zone ID
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ExtendClusterRequestNodeGroups) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestNodeGroups) GoString() string {
	return s.String()
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

type ExtendClusterRequestNodeGroupsNodeTag struct {
	// The key of tag.
	//
	// example:
	//
	// my_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag.
	//
	// example:
	//
	// my_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsNodeTag) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsNodeTag) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) SetKey(v string) *ExtendClusterRequestNodeGroupsNodeTag {
	s.Key = &v
	return s
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) SetValue(v string) *ExtendClusterRequestNodeGroupsNodeTag {
	s.Value = &v
	return s
}

type ExtendClusterRequestNodeGroupsNodes struct {
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
	return tea.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
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

type ExtendClusterShrinkRequest struct {
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
	// IP allocation policy combination: Each policy can only choose one type, and multiple policies can be combined
	IpAllocationPolicyShrink *string `json:"IpAllocationPolicy,omitempty" xml:"IpAllocationPolicy,omitempty"`
	// Node Group
	NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
	// VSwitch availability zone ID
	//
	// example:
	//
	// cn-shanghai-b
	VSwitchZoneId *string `json:"VSwitchZoneId,omitempty" xml:"VSwitchZoneId,omitempty"`
	// List of cluster subnets
	VpdSubnetsShrink *string `json:"VpdSubnets,omitempty" xml:"VpdSubnets,omitempty"`
}

func (s ExtendClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExtendClusterShrinkRequest) SetClusterId(v string) *ExtendClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ExtendClusterShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *ExtendClusterShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ExtendClusterShrinkRequest) SetIpAllocationPolicyShrink(v string) *ExtendClusterShrinkRequest {
	s.IpAllocationPolicyShrink = &v
	return s
}

func (s *ExtendClusterShrinkRequest) SetNodeGroupsShrink(v string) *ExtendClusterShrinkRequest {
	s.NodeGroupsShrink = &v
	return s
}

func (s *ExtendClusterShrinkRequest) SetVSwitchZoneId(v string) *ExtendClusterShrinkRequest {
	s.VSwitchZoneId = &v
	return s
}

func (s *ExtendClusterShrinkRequest) SetVpdSubnetsShrink(v string) *ExtendClusterShrinkRequest {
	s.VpdSubnetsShrink = &v
	return s
}

type ExtendClusterResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 03668372-18FF-5959-98D9-6B36A4643C7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task ID
	//
	// example:
	//
	// i158475611663639202234
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ExtendClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ExtendClusterResponseBody) SetRequestId(v string) *ExtendClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExtendClusterResponseBody) SetTaskId(v string) *ExtendClusterResponseBody {
	s.TaskId = &v
	return s
}

type ExtendClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExtendClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExtendClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ExtendClusterResponse) GoString() string {
	return s.String()
}

func (s *ExtendClusterResponse) SetHeaders(v map[string]*string) *ExtendClusterResponse {
	s.Headers = v
	return s
}

func (s *ExtendClusterResponse) SetStatusCode(v int32) *ExtendClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ExtendClusterResponse) SetBody(v *ExtendClusterResponseBody) *ExtendClusterResponse {
	s.Body = v
	return s
}

type ListClusterNodesRequest struct {
	// Cluster ID
	//
	// This parameter is required.
	//
	// example:
	//
	// i15b480fbd2fcdbc2869cd80
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Number of items per page in a paginated query, with a default value of 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Query token (Token), which is the value of the NextToken parameter returned by the previous API call.
	//
	// example:
	//
	// AAAAAdQ3Z+oPlg49gsr2y8jb6wY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Node group ID
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// resource group id
	//
	// example:
	//
	// rg-xxkxkllss
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// tag information
	Tags []*ListClusterNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListClusterNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterNodesRequest) SetClusterId(v string) *ListClusterNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListClusterNodesRequest) SetMaxResults(v int64) *ListClusterNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClusterNodesRequest) SetNextToken(v string) *ListClusterNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListClusterNodesRequest) SetNodeGroupId(v string) *ListClusterNodesRequest {
	s.NodeGroupId = &v
	return s
}

func (s *ListClusterNodesRequest) SetResourceGroupId(v string) *ListClusterNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClusterNodesRequest) SetTags(v []*ListClusterNodesRequestTags) *ListClusterNodesRequest {
	s.Tags = v
	return s
}

type ListClusterNodesRequestTags struct {
	// The key of tag object
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag object
	//
	// example:
	//
	// aa_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClusterNodesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListClusterNodesRequestTags) SetKey(v string) *ListClusterNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListClusterNodesRequestTags) SetValue(v string) *ListClusterNodesRequestTags {
	s.Value = &v
	return s
}

type ListClusterNodesResponseBody struct {
	// The query token value returned by this call.
	//
	// example:
	//
	// AAAAAXW/ZB9TBvH+0ZK0phtCibQgQmu1RbqplAI6Velo2OKR
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// List of nodes
	Nodes []*ListClusterNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// 2BA76272-6608-5AEC-BBA8-B6F0D3D14CDB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClusterNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBody) SetNextToken(v string) *ListClusterNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListClusterNodesResponseBody) SetNodes(v []*ListClusterNodesResponseBodyNodes) *ListClusterNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListClusterNodesResponseBody) SetRequestId(v string) *ListClusterNodesResponseBody {
	s.RequestId = &v
	return s
}

type ListClusterNodesResponseBodyNodes struct {
	// product code
	//
	// example:
	//
	// bcccluster
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// Creation time
	//
	// example:
	//
	// 1642472468000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Machine expiration time
	//
	// example:
	//
	// 1762185600000
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// whether or not support file system mount
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// Hostname
	//
	// example:
	//
	// 72432f80-273e-11ed-b57a-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// Hpn Zone
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// System image ID
	//
	// example:
	//
	// i190297201669099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// image name
	//
	// example:
	//
	// Alinux3_x86_AMD_R_Host_D3_E3_24.13.00_UEFI_N_250121
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// Machine type
	//
	// example:
	//
	// cn-wulanchabu-b11
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Network information
	Networks []*ListClusterNodesResponseBodyNodesNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
	// Node group ID
	//
	// example:
	//
	// ng-e9b74f4d450cf18d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// Node group name
	//
	// example:
	//
	// emr_master
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// Node ID
	//
	// example:
	//
	// e01-cn-2r42tmj4z02
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Node status
	//
	// example:
	//
	// Extending
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// Machine SN
	//
	// example:
	//
	// sn_tOuUk
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// tag information
	Tags []*ListClusterNodesResponseBodyNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// task id
	//
	// example:
	//
	// i28ddkdkkdkdd
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1mxqhw8o20tgv3xk47h
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-0jltf9vinjz3if3lltdy7
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// Zone ID
	//
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListClusterNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBodyNodes) SetCommodityCode(v string) *ListClusterNodesResponseBodyNodes {
	s.CommodityCode = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetCreateTime(v string) *ListClusterNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetExpiredTime(v string) *ListClusterNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetFileSystemMountEnabled(v bool) *ListClusterNodesResponseBodyNodes {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetHostname(v string) *ListClusterNodesResponseBodyNodes {
	s.Hostname = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetHpnZone(v string) *ListClusterNodesResponseBodyNodes {
	s.HpnZone = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetImageId(v string) *ListClusterNodesResponseBodyNodes {
	s.ImageId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetImageName(v string) *ListClusterNodesResponseBodyNodes {
	s.ImageName = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetMachineType(v string) *ListClusterNodesResponseBodyNodes {
	s.MachineType = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetNetworks(v []*ListClusterNodesResponseBodyNodesNetworks) *ListClusterNodesResponseBodyNodes {
	s.Networks = v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetNodeGroupId(v string) *ListClusterNodesResponseBodyNodes {
	s.NodeGroupId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetNodeGroupName(v string) *ListClusterNodesResponseBodyNodes {
	s.NodeGroupName = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetNodeId(v string) *ListClusterNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetOperatingState(v string) *ListClusterNodesResponseBodyNodes {
	s.OperatingState = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetSn(v string) *ListClusterNodesResponseBodyNodes {
	s.Sn = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetTags(v []*ListClusterNodesResponseBodyNodesTags) *ListClusterNodesResponseBodyNodes {
	s.Tags = v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetTaskId(v string) *ListClusterNodesResponseBodyNodes {
	s.TaskId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetVSwitchId(v string) *ListClusterNodesResponseBodyNodes {
	s.VSwitchId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetVpcId(v string) *ListClusterNodesResponseBodyNodes {
	s.VpcId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodes) SetZoneId(v string) *ListClusterNodesResponseBodyNodes {
	s.ZoneId = &v
	return s
}

type ListClusterNodesResponseBodyNodesNetworks struct {
	// Machine network interface name
	//
	// example:
	//
	// bond0
	BondName *string `json:"BondName,omitempty" xml:"BondName,omitempty"`
	// IP address within the VPC
	//
	// example:
	//
	// 192.168.22.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// VPC subnet ID
	//
	// example:
	//
	// subnet-fwekrvg9
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpd-eoiy88ju
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s ListClusterNodesResponseBodyNodesNetworks) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponseBodyNodesNetworks) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBodyNodesNetworks) SetBondName(v string) *ListClusterNodesResponseBodyNodesNetworks {
	s.BondName = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodesNetworks) SetIp(v string) *ListClusterNodesResponseBodyNodesNetworks {
	s.Ip = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodesNetworks) SetSubnetId(v string) *ListClusterNodesResponseBodyNodesNetworks {
	s.SubnetId = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodesNetworks) SetVpdId(v string) *ListClusterNodesResponseBodyNodesNetworks {
	s.VpdId = &v
	return s
}

type ListClusterNodesResponseBodyNodesTags struct {
	// The key of tag object
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag object
	//
	// example:
	//
	// aa_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClusterNodesResponseBodyNodesTags) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponseBodyNodesTags) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBodyNodesTags) SetKey(v string) *ListClusterNodesResponseBodyNodesTags {
	s.Key = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodesTags) SetValue(v string) *ListClusterNodesResponseBodyNodesTags {
	s.Value = &v
	return s
}

type ListClusterNodesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponse) SetHeaders(v map[string]*string) *ListClusterNodesResponse {
	s.Headers = v
	return s
}

func (s *ListClusterNodesResponse) SetStatusCode(v int32) *ListClusterNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClusterNodesResponse) SetBody(v *ListClusterNodesResponseBody) *ListClusterNodesResponse {
	s.Body = v
	return s
}

type ListClustersRequest struct {
	// Number of items per page for paginated queries, with a default value of 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Query token, which is the value of the NextToken parameter returned by the previous API call.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-aek2bg6wyoox6jq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// tag info
	Tags []*ListClustersRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) SetMaxResults(v int64) *ListClustersRequest {
	s.MaxResults = &v
	return s
}

func (s *ListClustersRequest) SetNextToken(v string) *ListClustersRequest {
	s.NextToken = &v
	return s
}

func (s *ListClustersRequest) SetResourceGroupId(v string) *ListClustersRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersRequest) SetTags(v []*ListClustersRequestTags) *ListClustersRequest {
	s.Tags = v
	return s
}

type ListClustersRequestTags struct {
	// The key of tag object
	//
	// example:
	//
	// key_aa
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag object
	//
	// example:
	//
	// value_aa
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListClustersRequestTags) GoString() string {
	return s.String()
}

func (s *ListClustersRequestTags) SetKey(v string) *ListClustersRequestTags {
	s.Key = &v
	return s
}

func (s *ListClustersRequestTags) SetValue(v string) *ListClustersRequestTags {
	s.Value = &v
	return s
}

type ListClustersResponseBody struct {
	// Cluster information
	Clusters []*ListClustersResponseBodyClusters `json:"Clusters,omitempty" xml:"Clusters,omitempty" type:"Repeated"`
	// The query token value returned by this call.
	//
	// example:
	//
	// f4f9a292c17072a2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID
	//
	// example:
	//
	// 2FE2B22C-CF9D-59DE-BF63-DC9B9B33A9D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody {
	s.Clusters = v
	return s
}

func (s *ListClustersResponseBody) SetNextToken(v string) *ListClustersResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListClustersResponseBody) SetRequestId(v string) *ListClustersResponseBody {
	s.RequestId = &v
	return s
}

type ListClustersResponseBodyClusters struct {
	// Cluster description
	//
	// example:
	//
	// PPU-cluster2 bz
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// Cluster ID
	//
	// example:
	//
	// i137590131672134915401
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster name
	//
	// example:
	//
	// cnp_test_cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Cluster type
	//
	// example:
	//
	// AckEdgePro
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// Component information
	//
	// example:
	//
	// {}
	Components interface{} `json:"Components,omitempty" xml:"Components,omitempty"`
	// IP version of the computing network
	//
	// example:
	//
	// IPv4
	ComputingIpVersion *string `json:"ComputingIpVersion,omitempty" xml:"ComputingIpVersion,omitempty"`
	// Creation time
	//
	// example:
	//
	// 1672134938
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Cluster number
	//
	// example:
	//
	// B1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Number of nodes
	//
	// example:
	//
	// 12
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Number of node groups
	//
	// example:
	//
	// 2
	NodeGroupCount *int64 `json:"NodeGroupCount,omitempty" xml:"NodeGroupCount,omitempty"`
	// Cluster status
	//
	// example:
	//
	// initializing
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-aek2ajbjoloa23q
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// tag information
	Tags []*ListClustersResponseBodyClustersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Task ID
	//
	// example:
	//
	// i156365121663149566024
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// Update time
	//
	// example:
	//
	// 1672134968
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// VPC ID
	//
	// example:
	//
	// vpc-0jlx4hol2bjboafzmffvd
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s ListClustersResponseBodyClusters) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) SetClusterDescription(v string) *ListClustersResponseBodyClusters {
	s.ClusterDescription = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterId(v string) *ListClustersResponseBodyClusters {
	s.ClusterId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterName(v string) *ListClustersResponseBodyClusters {
	s.ClusterName = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetClusterType(v string) *ListClustersResponseBodyClusters {
	s.ClusterType = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetComponents(v interface{}) *ListClustersResponseBodyClusters {
	s.Components = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetComputingIpVersion(v string) *ListClustersResponseBodyClusters {
	s.ComputingIpVersion = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetCreateTime(v string) *ListClustersResponseBodyClusters {
	s.CreateTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetHpnZone(v string) *ListClustersResponseBodyClusters {
	s.HpnZone = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetNodeCount(v int64) *ListClustersResponseBodyClusters {
	s.NodeCount = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetNodeGroupCount(v int64) *ListClustersResponseBodyClusters {
	s.NodeGroupCount = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetOperatingState(v string) *ListClustersResponseBodyClusters {
	s.OperatingState = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetResourceGroupId(v string) *ListClustersResponseBodyClusters {
	s.ResourceGroupId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetTags(v []*ListClustersResponseBodyClustersTags) *ListClustersResponseBodyClusters {
	s.Tags = v
	return s
}

func (s *ListClustersResponseBodyClusters) SetTaskId(v string) *ListClustersResponseBodyClusters {
	s.TaskId = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetUpdateTime(v string) *ListClustersResponseBodyClusters {
	s.UpdateTime = &v
	return s
}

func (s *ListClustersResponseBodyClusters) SetVpcId(v string) *ListClustersResponseBodyClusters {
	s.VpcId = &v
	return s
}

type ListClustersResponseBodyClustersTags struct {
	// The key of tag object
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag object
	//
	// example:
	//
	// aa_value
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersResponseBodyClustersTags) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponseBodyClustersTags) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersTags) SetKey(v string) *ListClustersResponseBodyClustersTags {
	s.Key = &v
	return s
}

func (s *ListClustersResponseBodyClustersTags) SetValue(v string) *ListClustersResponseBodyClustersTags {
	s.Value = &v
	return s
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) SetHeaders(v map[string]*string) *ListClustersResponse {
	s.Headers = v
	return s
}

func (s *ListClustersResponse) SetStatusCode(v int32) *ListClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClustersResponse) SetBody(v *ListClustersResponseBody) *ListClustersResponse {
	s.Body = v
	return s
}

type ListDiagnosticResultsRequest struct {
	// Type of diagnosis
	//
	// example:
	//
	// NetDiag
	DiagType *string `json:"DiagType,omitempty" xml:"DiagType,omitempty"`
	// Number of items per page in a paginated query. The maximum value is 100.
	//
	// Default value:
	//
	// - If no value is set or the set value is less than 20, the default is 20.
	//
	// - If the set value is greater than 100, the default is 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Query token (Token), the value should be the NextToken parameter value returned from the previous API call.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-acfmywpvugkh7kq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListDiagnosticResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticResultsRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsRequest) SetDiagType(v string) *ListDiagnosticResultsRequest {
	s.DiagType = &v
	return s
}

func (s *ListDiagnosticResultsRequest) SetMaxResults(v int64) *ListDiagnosticResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnosticResultsRequest) SetNextToken(v string) *ListDiagnosticResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListDiagnosticResultsRequest) SetResourceGroupId(v string) *ListDiagnosticResultsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListDiagnosticResultsResponseBody struct {
	// Diagnostic information
	DiagnosticResults []*ListDiagnosticResultsResponseBodyDiagnosticResults `json:"DiagnosticResults,omitempty" xml:"DiagnosticResults,omitempty" type:"Repeated"`
	// 分页查询时每页行数。最大值为100。
	//
	// 默认值：
	//
	// •当不设置值或设置的值小于20时，默认值为20。
	//
	// •当设置的值大于100时，默认值为100。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// NextToken for the next page. Include this value when requesting the next page.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// ID of the request
	//
	// example:
	//
	// AC4F0004-7BCE-52E0-891B-CAC7D64E3368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDiagnosticResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponseBody) SetDiagnosticResults(v []*ListDiagnosticResultsResponseBodyDiagnosticResults) *ListDiagnosticResultsResponseBody {
	s.DiagnosticResults = v
	return s
}

func (s *ListDiagnosticResultsResponseBody) SetMaxResults(v int64) *ListDiagnosticResultsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDiagnosticResultsResponseBody) SetNextToken(v string) *ListDiagnosticResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDiagnosticResultsResponseBody) SetRequestId(v string) *ListDiagnosticResultsResponseBody {
	s.RequestId = &v
	return s
}

type ListDiagnosticResultsResponseBodyDiagnosticResults struct {
	// Cluster ID
	//
	// example:
	//
	// i118578141694745246055
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster name
	//
	// example:
	//
	// pjlab-lingjun
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Creation time of the diagnostic task.
	//
	// example:
	//
	// 2024-01-15T02:01:12Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Diagnosis content. For example, in network diagnosis, there are static configuration checks, dynamic operation checks, etc.
	//
	// example:
	//
	// diagcontent
	DiagContent *string `json:"DiagContent,omitempty" xml:"DiagContent,omitempty"`
	// Diagnosis ID
	//
	// example:
	//
	// 123
	DiagId *string `json:"DiagId,omitempty" xml:"DiagId,omitempty"`
	// Diagnosis result, success or failure.
	//
	// example:
	//
	// Success
	DiagResult *string `json:"DiagResult,omitempty" xml:"DiagResult,omitempty"`
	// Completion time of the diagnostic task.
	//
	// example:
	//
	// 2024-10-16T02:04Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// Resource ID
	//
	// example:
	//
	// e01-cn-bl03ofg6206
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// proxy-rps.mos.csvw.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
	// Governance status
	//
	// example:
	//
	// succeed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDiagnosticResultsResponseBodyDiagnosticResults) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticResultsResponseBodyDiagnosticResults) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetClusterId(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ClusterId = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetClusterName(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ClusterName = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetCreationTime(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.CreationTime = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetDiagContent(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.DiagContent = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetDiagId(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.DiagId = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetDiagResult(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.DiagResult = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetFinishedTime(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.FinishedTime = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetResourceId(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ResourceId = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetServerName(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.ServerName = &v
	return s
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) SetStatus(v string) *ListDiagnosticResultsResponseBodyDiagnosticResults {
	s.Status = &v
	return s
}

type ListDiagnosticResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnosticResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnosticResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDiagnosticResultsResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponse) SetHeaders(v map[string]*string) *ListDiagnosticResultsResponse {
	s.Headers = v
	return s
}

func (s *ListDiagnosticResultsResponse) SetStatusCode(v int32) *ListDiagnosticResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDiagnosticResultsResponse) SetBody(v *ListDiagnosticResultsResponseBody) *ListDiagnosticResultsResponse {
	s.Body = v
	return s
}

type ListFreeNodesRequest struct {
	// Cluster number
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Machine type
	//
	// example:
	//
	// mock-machine-type2
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Number of items per page for paginated queries, default is 20.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Query token (Token), the value should be the NextToken parameter value returned from the previous API call.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken       *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OperatingStates []*string `json:"OperatingStates,omitempty" xml:"OperatingStates,omitempty" type:"Repeated"`
	// Resource group ID
	//
	// example:
	//
	// rg-acfmxno4vh5muoq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Tag information
	Tags []*ListFreeNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListFreeNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesRequest) GoString() string {
	return s.String()
}

func (s *ListFreeNodesRequest) SetHpnZone(v string) *ListFreeNodesRequest {
	s.HpnZone = &v
	return s
}

func (s *ListFreeNodesRequest) SetMachineType(v string) *ListFreeNodesRequest {
	s.MachineType = &v
	return s
}

func (s *ListFreeNodesRequest) SetMaxResults(v int64) *ListFreeNodesRequest {
	s.MaxResults = &v
	return s
}

func (s *ListFreeNodesRequest) SetNextToken(v string) *ListFreeNodesRequest {
	s.NextToken = &v
	return s
}

func (s *ListFreeNodesRequest) SetOperatingStates(v []*string) *ListFreeNodesRequest {
	s.OperatingStates = v
	return s
}

func (s *ListFreeNodesRequest) SetResourceGroupId(v string) *ListFreeNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFreeNodesRequest) SetTags(v []*ListFreeNodesRequestTags) *ListFreeNodesRequest {
	s.Tags = v
	return s
}

type ListFreeNodesRequestTags struct {
	// The key of tag object
	//
	// example:
	//
	// key_aa
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag object
	//
	// example:
	//
	// value_aa
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFreeNodesRequestTags) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListFreeNodesRequestTags) SetKey(v string) *ListFreeNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListFreeNodesRequestTags) SetValue(v string) *ListFreeNodesRequestTags {
	s.Value = &v
	return s
}

type ListFreeNodesResponseBody struct {
	// The query token value returned by this call.
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// List of nodes
	Nodes []*ListFreeNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// Request ID
	//
	// example:
	//
	// AA14CB86-70C4-5CB7-9E7B-6CCA77F3512B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFreeNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBody) SetNextToken(v string) *ListFreeNodesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListFreeNodesResponseBody) SetNodes(v []*ListFreeNodesResponseBodyNodes) *ListFreeNodesResponseBody {
	s.Nodes = v
	return s
}

func (s *ListFreeNodesResponseBody) SetRequestId(v string) *ListFreeNodesResponseBody {
	s.RequestId = &v
	return s
}

type ListFreeNodesResponseBodyNodes struct {
	// Product Code
	//
	// example:
	//
	// bccluster_eflocomputing_public_cn
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// Creation time
	//
	// example:
	//
	// 1652321554
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Expiration time of the machine
	//
	// example:
	//
	// 1673107200
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// Cluster number
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Machine type
	//
	// example:
	//
	// efg1.nvga1
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Node ID
	//
	// example:
	//
	// e01-cn-7pp2x193801
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// Node status
	//
	// example:
	//
	// Unused
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-aekzkkbrpl4owgy
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// Machine SN
	//
	// example:
	//
	// sn_pozkHBgicd
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// Tags Info
	Tags []*ListFreeNodesResponseBodyNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Availability zone ID
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListFreeNodesResponseBodyNodes) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBodyNodes) SetCommodityCode(v string) *ListFreeNodesResponseBodyNodes {
	s.CommodityCode = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetCreateTime(v string) *ListFreeNodesResponseBodyNodes {
	s.CreateTime = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetExpiredTime(v string) *ListFreeNodesResponseBodyNodes {
	s.ExpiredTime = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetHpnZone(v string) *ListFreeNodesResponseBodyNodes {
	s.HpnZone = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetMachineType(v string) *ListFreeNodesResponseBodyNodes {
	s.MachineType = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetNodeId(v string) *ListFreeNodesResponseBodyNodes {
	s.NodeId = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetOperatingState(v string) *ListFreeNodesResponseBodyNodes {
	s.OperatingState = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetResourceGroupId(v string) *ListFreeNodesResponseBodyNodes {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetSn(v string) *ListFreeNodesResponseBodyNodes {
	s.Sn = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetTags(v []*ListFreeNodesResponseBodyNodesTags) *ListFreeNodesResponseBodyNodes {
	s.Tags = v
	return s
}

func (s *ListFreeNodesResponseBodyNodes) SetZoneId(v string) *ListFreeNodesResponseBodyNodes {
	s.ZoneId = &v
	return s
}

type ListFreeNodesResponseBodyNodesTags struct {
	// The key of tag object
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag object
	//
	// example:
	//
	// aa_vakye
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFreeNodesResponseBodyNodesTags) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesResponseBodyNodesTags) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBodyNodesTags) SetKey(v string) *ListFreeNodesResponseBodyNodesTags {
	s.Key = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodesTags) SetValue(v string) *ListFreeNodesResponseBodyNodesTags {
	s.Value = &v
	return s
}

type ListFreeNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFreeNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFreeNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFreeNodesResponse) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponse) SetHeaders(v map[string]*string) *ListFreeNodesResponse {
	s.Headers = v
	return s
}

func (s *ListFreeNodesResponse) SetStatusCode(v int32) *ListFreeNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFreeNodesResponse) SetBody(v *ListFreeNodesResponseBody) *ListFreeNodesResponse {
	s.Body = v
	return s
}

type ListImagesRequest struct {
	// Architecture
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// Image version
	//
	// example:
	//
	// 7.9
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// Platform
	//
	// example:
	//
	// ALinux3
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
}

func (s ListImagesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) SetArchitecture(v string) *ListImagesRequest {
	s.Architecture = &v
	return s
}

func (s *ListImagesRequest) SetImageVersion(v string) *ListImagesRequest {
	s.ImageVersion = &v
	return s
}

func (s *ListImagesRequest) SetPlatform(v string) *ListImagesRequest {
	s.Platform = &v
	return s
}

type ListImagesResponseBody struct {
	// Image details
	Images []*ListImagesResponseBodyImages `json:"Images,omitempty" xml:"Images,omitempty" type:"Repeated"`
	// NextToken for the next page, include this value when requesting the next page
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0FC4A1C7-421C-5EAB-9361-4C0338EFA287
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListImagesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody {
	s.Images = v
	return s
}

func (s *ListImagesResponseBody) SetNextToken(v string) *ListImagesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListImagesResponseBody) SetRequestId(v string) *ListImagesResponseBody {
	s.RequestId = &v
	return s
}

type ListImagesResponseBodyImages struct {
	// Architecture
	//
	// example:
	//
	// x86_64
	Architecture *string `json:"Architecture,omitempty" xml:"Architecture,omitempty"`
	// Description
	//
	// example:
	//
	// alibaba cloud linux 3 full for H800
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Image ID
	//
	// example:
	//
	// i190951671671438639388
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Image name
	//
	// example:
	//
	// CentOS_7.9_x86_64_FULL_20221110
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// Image version
	//
	// example:
	//
	// 7.9
	ImageVersion *string `json:"ImageVersion,omitempty" xml:"ImageVersion,omitempty"`
	// node count
	//
	// example:
	//
	// 20
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Platform
	//
	// example:
	//
	// ALinux3
	Platform *string `json:"Platform,omitempty" xml:"Platform,omitempty"`
	// File MD5
	//
	// example:
	//
	// 40741292480fc6d778138adcf8c
	ReleaseFileMd5 *string `json:"ReleaseFileMd5,omitempty" xml:"ReleaseFileMd5,omitempty"`
	// Image size
	//
	// example:
	//
	// 5.8G
	ReleaseFileSize *string `json:"ReleaseFileSize,omitempty" xml:"ReleaseFileSize,omitempty"`
	// image type
	//
	// example:
	//
	// Public
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) SetArchitecture(v string) *ListImagesResponseBodyImages {
	s.Architecture = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetDescription(v string) *ListImagesResponseBodyImages {
	s.Description = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageId(v string) *ListImagesResponseBodyImages {
	s.ImageId = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageName(v string) *ListImagesResponseBodyImages {
	s.ImageName = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetImageVersion(v string) *ListImagesResponseBodyImages {
	s.ImageVersion = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetNodeCount(v int64) *ListImagesResponseBodyImages {
	s.NodeCount = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetPlatform(v string) *ListImagesResponseBodyImages {
	s.Platform = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetReleaseFileMd5(v string) *ListImagesResponseBodyImages {
	s.ReleaseFileMd5 = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetReleaseFileSize(v string) *ListImagesResponseBodyImages {
	s.ReleaseFileSize = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetType(v string) *ListImagesResponseBodyImages {
	s.Type = &v
	return s
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImagesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) SetHeaders(v map[string]*string) *ListImagesResponse {
	s.Headers = v
	return s
}

func (s *ListImagesResponse) SetStatusCode(v int32) *ListImagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListImagesResponse) SetBody(v *ListImagesResponseBody) *ListImagesResponse {
	s.Body = v
	return s
}

type ListMachineNetworkInfoRequest struct {
	// Array
	MachineHpnInfo []*ListMachineNetworkInfoRequestMachineHpnInfo `json:"MachineHpnInfo,omitempty" xml:"MachineHpnInfo,omitempty" type:"Repeated"`
}

func (s ListMachineNetworkInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoRequest) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoRequest) SetMachineHpnInfo(v []*ListMachineNetworkInfoRequestMachineHpnInfo) *ListMachineNetworkInfoRequest {
	s.MachineHpnInfo = v
	return s
}

type ListMachineNetworkInfoRequestMachineHpnInfo struct {
	// Cluster ID
	//
	// example:
	//
	// C1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Machine type
	//
	// example:
	//
	// efg2.C48cNHmcn
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Region ID
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListMachineNetworkInfoRequestMachineHpnInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoRequestMachineHpnInfo) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) SetHpnZone(v string) *ListMachineNetworkInfoRequestMachineHpnInfo {
	s.HpnZone = &v
	return s
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) SetMachineType(v string) *ListMachineNetworkInfoRequestMachineHpnInfo {
	s.MachineType = &v
	return s
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) SetRegionId(v string) *ListMachineNetworkInfoRequestMachineHpnInfo {
	s.RegionId = &v
	return s
}

type ListMachineNetworkInfoShrinkRequest struct {
	// Array
	MachineHpnInfoShrink *string `json:"MachineHpnInfo,omitempty" xml:"MachineHpnInfo,omitempty"`
}

func (s ListMachineNetworkInfoShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoShrinkRequest) SetMachineHpnInfoShrink(v string) *ListMachineNetworkInfoShrinkRequest {
	s.MachineHpnInfoShrink = &v
	return s
}

type ListMachineNetworkInfoResponseBody struct {
	// Array
	MachineNetworkInfo []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo `json:"MachineNetworkInfo,omitempty" xml:"MachineNetworkInfo,omitempty" type:"Repeated"`
	// ID of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMachineNetworkInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponseBody) SetMachineNetworkInfo(v []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo) *ListMachineNetworkInfoResponseBody {
	s.MachineNetworkInfo = v
	return s
}

func (s *ListMachineNetworkInfoResponseBody) SetRequestId(v string) *ListMachineNetworkInfoResponseBody {
	s.RequestId = &v
	return s
}

type ListMachineNetworkInfoResponseBodyMachineNetworkInfo struct {
	// Cluster network
	//
	// example:
	//
	// vpc/acl
	ClusterNet *string `json:"ClusterNet,omitempty" xml:"ClusterNet,omitempty"`
	// Whether jumbo frame capability is enabled
	//
	// example:
	//
	// true
	EnableJumboFrame *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	// Cluster ID
	//
	// example:
	//
	// B1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// Whether it is in DPU mode
	//
	// example:
	//
	// true
	IsDpuMode *bool `json:"IsDpuMode,omitempty" xml:"IsDpuMode,omitempty"`
	// Machine type
	//
	// example:
	//
	// efg1.nvga8n
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Network architecture
	//
	// example:
	//
	// XX-7.0
	NetArch *string `json:"NetArch,omitempty" xml:"NetArch,omitempty"`
	// 地域ID。
	//
	// example:
	//
	// cn-heyuan
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s ListMachineNetworkInfoResponseBodyMachineNetworkInfo) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetClusterNet(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.ClusterNet = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetEnableJumboFrame(v bool) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.EnableJumboFrame = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetHpnZone(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.HpnZone = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetIsDpuMode(v bool) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.IsDpuMode = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetMachineType(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.MachineType = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetNetArch(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.NetArch = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) SetRegionId(v string) *ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	s.RegionId = &v
	return s
}

type ListMachineNetworkInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachineNetworkInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMachineNetworkInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMachineNetworkInfoResponse) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponse) SetHeaders(v map[string]*string) *ListMachineNetworkInfoResponse {
	s.Headers = v
	return s
}

func (s *ListMachineNetworkInfoResponse) SetStatusCode(v int32) *ListMachineNetworkInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachineNetworkInfoResponse) SetBody(v *ListMachineNetworkInfoResponseBody) *ListMachineNetworkInfoResponse {
	s.Body = v
	return s
}

type ListMachineTypesRequest struct {
	// Machine name
	//
	// example:
	//
	// efg1.nvga1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListMachineTypesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMachineTypesRequest) GoString() string {
	return s.String()
}

func (s *ListMachineTypesRequest) SetName(v string) *ListMachineTypesRequest {
	s.Name = &v
	return s
}

type ListMachineTypesResponseBody struct {
	// Details of the machine types
	MachineTypes []*ListMachineTypesResponseBodyMachineTypes `json:"MachineTypes,omitempty" xml:"MachineTypes,omitempty" type:"Repeated"`
	// NextToken for the next page, include this value when requesting the next page
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F16BA4D8-FF50-53B6-A026-F443FE31006C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListMachineTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMachineTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBody) SetMachineTypes(v []*ListMachineTypesResponseBodyMachineTypes) *ListMachineTypesResponseBody {
	s.MachineTypes = v
	return s
}

func (s *ListMachineTypesResponseBody) SetNextToken(v string) *ListMachineTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListMachineTypesResponseBody) SetRequestId(v string) *ListMachineTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListMachineTypesResponseBodyMachineTypes struct {
	// Number of bonds
	//
	// example:
	//
	// 2
	BondNum *int32 `json:"BondNum,omitempty" xml:"BondNum,omitempty"`
	// CPU information
	//
	// example:
	//
	// 2x Intel Icelake 8369B 32C CPU
	CpuInfo *string `json:"CpuInfo,omitempty" xml:"CpuInfo,omitempty"`
	// Disk information
	//
	// example:
	//
	// 2x 480GB SATA SSD
	DiskInfo *string `json:"DiskInfo,omitempty" xml:"DiskInfo,omitempty"`
	// GPU information
	//
	// example:
	//
	// 8x NVIDIA SXM4 80GB A100 GPU
	GpuInfo *string `json:"GpuInfo,omitempty" xml:"GpuInfo,omitempty"`
	// Memory information
	//
	// example:
	//
	// 32x 64GB DDR4 3200 Memory
	MemoryInfo *string `json:"MemoryInfo,omitempty" xml:"MemoryInfo,omitempty"`
	// Machine name
	//
	// example:
	//
	// efg1.nvga1
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Network information
	//
	// example:
	//
	// 2x 100Gbps DP NIC
	NetworkInfo *string `json:"NetworkInfo,omitempty" xml:"NetworkInfo,omitempty"`
	// Number of nodes
	//
	// example:
	//
	// 10
	NodeCount *string `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Number of CPU cores
	//
	// example:
	//
	// 48
	TotalCpuCore *int32 `json:"TotalCpuCore,omitempty" xml:"TotalCpuCore,omitempty"`
	// Type of machine
	//
	// example:
	//
	// Public
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListMachineTypesResponseBodyMachineTypes) String() string {
	return tea.Prettify(s)
}

func (s ListMachineTypesResponseBodyMachineTypes) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetBondNum(v int32) *ListMachineTypesResponseBodyMachineTypes {
	s.BondNum = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetCpuInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.CpuInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetDiskInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.DiskInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetGpuInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.GpuInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetMemoryInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.MemoryInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetName(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.Name = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetNetworkInfo(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.NetworkInfo = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetNodeCount(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.NodeCount = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetTotalCpuCore(v int32) *ListMachineTypesResponseBodyMachineTypes {
	s.TotalCpuCore = &v
	return s
}

func (s *ListMachineTypesResponseBodyMachineTypes) SetType(v string) *ListMachineTypesResponseBodyMachineTypes {
	s.Type = &v
	return s
}

type ListMachineTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachineTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMachineTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMachineTypesResponse) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponse) SetHeaders(v map[string]*string) *ListMachineTypesResponse {
	s.Headers = v
	return s
}

func (s *ListMachineTypesResponse) SetStatusCode(v int32) *ListMachineTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMachineTypesResponse) SetBody(v *ListMachineTypesResponseBody) *ListMachineTypesResponse {
	s.Body = v
	return s
}

type ListNetTestResultsRequest struct {
	// Number of items per page in a paginated query. The maximum value is 100.
	//
	// Default value:
	//
	// - If no value is set or the set value is less than 20, the default is 20.
	//
	// - If the set value is greater than 100, the default is 100.
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Type of network test.
	//
	// example:
	//
	// DelayTest
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// Query token (Token), which should be the value of the NextToken parameter returned from the previous API call.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-acfmxno4vh5muoq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
}

func (s ListNetTestResultsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsRequest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsRequest) SetMaxResults(v int64) *ListNetTestResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNetTestResultsRequest) SetNetTestType(v string) *ListNetTestResultsRequest {
	s.NetTestType = &v
	return s
}

func (s *ListNetTestResultsRequest) SetNextToken(v string) *ListNetTestResultsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNetTestResultsRequest) SetResourceGroupId(v string) *ListNetTestResultsRequest {
	s.ResourceGroupId = &v
	return s
}

type ListNetTestResultsResponseBody struct {
	// 分页查询时每页行数。最大值为100。
	//
	// 默认值：
	//
	// •当不设置值或设置的值小于20时，默认值为20。
	//
	// •当设置的值大于100时，默认值为100。
	//
	// example:
	//
	// 20
	MaxResults *int64 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// List of nodes
	NetTestResults []*ListNetTestResultsResponseBodyNetTestResults `json:"NetTestResults,omitempty" xml:"NetTestResults,omitempty" type:"Repeated"`
	// NextToken for the next page, to be included in the request for the next page
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 3C683243-7915-57FB-9570-A2932C1C0F78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNetTestResultsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBody) SetMaxResults(v int64) *ListNetTestResultsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListNetTestResultsResponseBody) SetNetTestResults(v []*ListNetTestResultsResponseBodyNetTestResults) *ListNetTestResultsResponseBody {
	s.NetTestResults = v
	return s
}

func (s *ListNetTestResultsResponseBody) SetNextToken(v string) *ListNetTestResultsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNetTestResultsResponseBody) SetRequestId(v string) *ListNetTestResultsResponseBody {
	s.RequestId = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResults struct {
	// Cluster ID
	//
	// example:
	//
	// i110667211718265012218
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster name
	//
	// example:
	//
	// test
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// To be filled when the network test type is communication library test
	CommTest *ListNetTestResultsResponseBodyNetTestResultsCommTest `json:"CommTest,omitempty" xml:"CommTest,omitempty" type:"Struct"`
	// Creation time.
	//
	// example:
	//
	// 2024-01-19T02:18:12Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Fill in when the network test type is latency test
	DelayTest *ListNetTestResultsResponseBodyNetTestResultsDelayTest `json:"DelayTest,omitempty" xml:"DelayTest,omitempty" type:"Struct"`
	// Completion time.
	//
	// example:
	//
	// 2024-10-30T02:07Z
	FinishedTime *string `json:"FinishedTime,omitempty" xml:"FinishedTime,omitempty"`
	// Type of network test.
	//
	// example:
	//
	// NetDiag
	NetTestType *string `json:"NetTestType,omitempty" xml:"NetTestType,omitempty"`
	// Network mode
	//
	// example:
	//
	// 01
	NetworkMode *string `json:"NetworkMode,omitempty" xml:"NetworkMode,omitempty"`
	// Test port number.
	//
	// example:
	//
	// 80
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
	// Status of the network test task. Possible values:</br>
	//
	// - InProgress: Testing in progress.</br>
	//
	// - Finished: Test completed.</br>
	//
	// - Failed: Test failed.
	//
	// example:
	//
	// InProgress
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Test ID. A unique identifier for the resource test task.
	//
	// example:
	//
	// String	i-uf6i0tv2refv8wz*****
	TestId *string `json:"TestId,omitempty" xml:"TestId,omitempty"`
	// Fill in when the network test type is traffic test.
	TrafficTest *ListNetTestResultsResponseBodyNetTestResultsTrafficTest `json:"TrafficTest,omitempty" xml:"TrafficTest,omitempty" type:"Struct"`
}

func (s ListNetTestResultsResponseBodyNetTestResults) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResults) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetClusterId(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.ClusterId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetClusterName(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.ClusterName = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetCommTest(v *ListNetTestResultsResponseBodyNetTestResultsCommTest) *ListNetTestResultsResponseBodyNetTestResults {
	s.CommTest = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetCreationTime(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.CreationTime = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetDelayTest(v *ListNetTestResultsResponseBodyNetTestResultsDelayTest) *ListNetTestResultsResponseBodyNetTestResults {
	s.DelayTest = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetFinishedTime(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.FinishedTime = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetNetTestType(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.NetTestType = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetNetworkMode(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.NetworkMode = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetPort(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.Port = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetStatus(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.Status = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetTestId(v string) *ListNetTestResultsResponseBodyNetTestResults {
	s.TestId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResults) SetTrafficTest(v *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) *ListNetTestResultsResponseBodyNetTestResults {
	s.TrafficTest = v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsCommTest struct {
	// Number of GPUs
	//
	// example:
	//
	// 4
	GPUNum *string `json:"GPUNum,omitempty" xml:"GPUNum,omitempty"`
	// Input hosts for the test nodes
	Hosts []*ListNetTestResultsResponseBodyNetTestResultsCommTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
	// Communication library model
	//
	// example:
	//
	// AllToAll
	Model *string `json:"Model,omitempty" xml:"Model,omitempty"`
	// Communication library test category: ACCL or NCCL
	//
	// example:
	//
	// ACCL
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTest) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetGPUNum(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.GPUNum = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetHosts(v []*ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.Hosts = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetModel(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.Model = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) SetType(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	s.Type = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsCommTestHosts struct {
	// Node IP
	//
	// example:
	//
	// 10.51.16.21
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Resource ID
	//
	// example:
	//
	// e01-cn-wwo3eteze19
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name
	//
	// example:
	//
	// www.xinjiaoyu.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	s.ServerName = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsDelayTest struct {
	// Resource list
	Hosts []*ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTest) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTest) SetHosts(v []*ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) *ListNetTestResultsResponseBodyNetTestResultsDelayTest {
	s.Hosts = v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts struct {
	// Bond interface of the network card
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// Node IP
	//
	// example:
	//
	// pgm-bp174z988a27wre71o.pg.rds.aliyuncs.com
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// 资源id
	//
	// example:
	//
	// e01-cn-wwo3eteze19
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name
	//
	// example:
	//
	// WrF
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetBond(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.Bond = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	s.ServerName = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsTrafficTest struct {
	// Clients
	Clients []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients `json:"Clients,omitempty" xml:"Clients,omitempty" type:"Repeated"`
	// Duration of the workflow task, in seconds.
	//
	// example:
	//
	// 1
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// 协议为RDMA时，填写True/False，
	//
	// 协议为TCP时，此字段为空。
	//
	// example:
	//
	// True
	GDR *string `json:"GDR,omitempty" xml:"GDR,omitempty"`
	// Network protocol, either RDMA or TCP.
	//
	// example:
	//
	// TCP
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// For TCP, enter the number of concurrent connections; for RDMA, enter the configured QP value.
	//
	// example:
	//
	// RDMA
	QP *int64 `json:"QP,omitempty" xml:"QP,omitempty"`
	// This field is empty when the traffic model (TrafficModel) is Fullmesh.
	Servers []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers `json:"Servers,omitempty" xml:"Servers,omitempty" type:"Repeated"`
	// Traffic model, either MTON or Fullmesh.
	//
	// example:
	//
	// Fullmesh
	TrafficModel *string `json:"TrafficModel,omitempty" xml:"TrafficModel,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTest) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetClients(v []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Clients = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetDuration(v int64) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Duration = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetGDR(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.GDR = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetProtocol(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Protocol = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetQP(v int64) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.QP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetServers(v []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.Servers = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) SetTrafficModel(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	s.TrafficModel = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients struct {
	// Network interface bond port
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// IP address.
	//
	// example:
	//
	// 74.73.100.1
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Resource ID.
	//
	// example:
	//
	// e01-cn-20p36bqet39
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// prod-gf-cn.juequling.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetBond(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.Bond = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	s.ServerName = &v
	return s
}

type ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers struct {
	// Network interface bond port
	//
	// example:
	//
	// bond1
	Bond *string `json:"Bond,omitempty" xml:"Bond,omitempty"`
	// Node IP
	//
	// example:
	//
	// 10.1.168.183
	IP *string `json:"IP,omitempty" xml:"IP,omitempty"`
	// Resource ID.
	//
	// example:
	//
	// e01-cn-wwo3eteze19
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Service name.
	//
	// example:
	//
	// prod-gf-cn.juequling.com
	ServerName *string `json:"ServerName,omitempty" xml:"ServerName,omitempty"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetBond(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.Bond = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetIP(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.IP = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetResourceId(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.ResourceId = &v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) SetServerName(v string) *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	s.ServerName = &v
	return s
}

type ListNetTestResultsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetTestResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetTestResultsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNetTestResultsResponse) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponse) SetHeaders(v map[string]*string) *ListNetTestResultsResponse {
	s.Headers = v
	return s
}

func (s *ListNetTestResultsResponse) SetStatusCode(v int32) *ListNetTestResultsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNetTestResultsResponse) SetBody(v *ListNetTestResultsResponseBody) *ListNetTestResultsResponse {
	s.Body = v
	return s
}

type ListNodeGroupsRequest struct {
	// Cluster ID
	//
	// example:
	//
	// i119982311660892626523
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Number of items per page in a paginated query. The maximum value is 100.
	//
	// Default value:
	//
	// - If no value is set or the set value is less than 20, the default value is 20.
	//
	// - If the set value is greater than 100, the default value is 100.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// NextToken for the next page, include this value when requesting the next page
	//
	// example:
	//
	// a3f2224a5ec7224116c4f5246120abe4
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Node group ID
	//
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s ListNodeGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsRequest) SetClusterId(v string) *ListNodeGroupsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListNodeGroupsRequest) SetMaxResults(v int32) *ListNodeGroupsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListNodeGroupsRequest) SetNextToken(v string) *ListNodeGroupsRequest {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupsRequest) SetNodeGroupId(v string) *ListNodeGroupsRequest {
	s.NodeGroupId = &v
	return s
}

type ListNodeGroupsResponseBody struct {
	// Cluster group information
	Groups []*ListNodeGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// NextToken for the next page, include this value when requesting the next page
	//
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodeGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBody) SetGroups(v []*ListNodeGroupsResponseBodyGroups) *ListNodeGroupsResponseBody {
	s.Groups = v
	return s
}

func (s *ListNodeGroupsResponseBody) SetNextToken(v string) *ListNodeGroupsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListNodeGroupsResponseBody) SetRequestId(v string) *ListNodeGroupsResponseBody {
	s.RequestId = &v
	return s
}

type ListNodeGroupsResponseBodyGroups struct {
	// Cluster ID
	//
	// example:
	//
	// i113952461729854708648
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Cluster name
	//
	// example:
	//
	// wzq-exclusivelite-71
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// Creation time
	//
	// example:
	//
	// 2024-02-27T13:16:31.599
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// Description
	//
	// example:
	//
	// created by ga2_prepare
	Description            *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileSystemMountEnabled *bool   `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// Group ID.
	//
	// example:
	//
	// 238276221
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// Group name.
	//
	// example:
	//
	// backend-group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// Image ID
	//
	// example:
	//
	// i194015071707321240258
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Image name
	//
	// example:
	//
	// CentOS_7.9_x86_64_FULL_20221110
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// Machine type
	//
	// example:
	//
	// efg1.nvga1n
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// Number of nodes
	//
	// example:
	//
	// 2
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// Update time
	//
	// example:
	//
	// 2023-09-22T00:03:05.114
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// 可用区id
	//
	// example:
	//
	// cn-shenzhen-c
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListNodeGroupsResponseBodyGroups) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBodyGroups) SetClusterId(v string) *ListNodeGroupsResponseBodyGroups {
	s.ClusterId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetClusterName(v string) *ListNodeGroupsResponseBodyGroups {
	s.ClusterName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetCreateTime(v string) *ListNodeGroupsResponseBodyGroups {
	s.CreateTime = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetDescription(v string) *ListNodeGroupsResponseBodyGroups {
	s.Description = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetFileSystemMountEnabled(v bool) *ListNodeGroupsResponseBodyGroups {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetGroupId(v string) *ListNodeGroupsResponseBodyGroups {
	s.GroupId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetGroupName(v string) *ListNodeGroupsResponseBodyGroups {
	s.GroupName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetImageId(v string) *ListNodeGroupsResponseBodyGroups {
	s.ImageId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetImageName(v string) *ListNodeGroupsResponseBodyGroups {
	s.ImageName = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetMachineType(v string) *ListNodeGroupsResponseBodyGroups {
	s.MachineType = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetNodeCount(v int64) *ListNodeGroupsResponseBodyGroups {
	s.NodeCount = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetUpdateTime(v string) *ListNodeGroupsResponseBodyGroups {
	s.UpdateTime = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetZoneId(v string) *ListNodeGroupsResponseBodyGroups {
	s.ZoneId = &v
	return s
}

type ListNodeGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListNodeGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponse) SetHeaders(v map[string]*string) *ListNodeGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListNodeGroupsResponse) SetStatusCode(v int32) *ListNodeGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListNodeGroupsResponse) SetBody(v *ListNodeGroupsResponseBody) *ListNodeGroupsResponse {
	s.Body = v
	return s
}

type ListTagResourcesRequest struct {
	// Query token (Token), the value should be the NextToken returned from the previous API call
	//
	// example:
	//
	// AAAAAdQ3Z+oPlg49gsr2y8jb6wY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// List of resource IDs
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// Resource type
	//
	// This parameter is required.
	//
	// example:
	//
	// Node
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// List of tags
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) SetNextToken(v string) *ListTagResourcesRequest {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesRequest) SetRegionId(v string) *ListTagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *ListTagResourcesRequest) SetResourceId(v []*string) *ListTagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *ListTagResourcesRequest) SetResourceType(v string) *ListTagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

type ListTagResourcesRequestTag struct {
	// Tag key
	//
	// example:
	//
	// PodName
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value
	//
	// example:
	//
	// WFT-OTC
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListTagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

type ListTagResourcesResponseBody struct {
	// NextToken for the next page, include this returned value when requesting the next page
	//
	// example:
	//
	// AAAAAdQ3Z+oPlg49gsr2y8jb6wY=
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Request ID
	//
	// example:
	//
	// 8F208B6D-4C42-5FD3-B6BE-E826E92A44DD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Tagged resources.
	TagResources *ListTagResourcesResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Struct"`
}

func (s ListTagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) SetNextToken(v string) *ListTagResourcesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetRequestId(v string) *ListTagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagResourcesResponseBody) SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody {
	s.TagResources = v
	return s
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

type ListTagResourcesResponseBodyTagResourcesTagResource struct {
	// Resource ID
	//
	// example:
	//
	// i15azeddnvf7uhw2oij57o0
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// Resource type
	//
	// example:
	//
	// Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Tag key
	//
	// example:
	//
	// env
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// Tag value
	//
	// example:
	//
	// dev
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceId(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceId = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetResourceType(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.ResourceType = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagKey(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagKey = &v
	return s
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) SetTagValue(v string) *ListTagResourcesResponseBodyTagResourcesTagResource {
	s.TagValue = &v
	return s
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) SetHeaders(v map[string]*string) *ListTagResourcesResponse {
	s.Headers = v
	return s
}

func (s *ListTagResourcesResponse) SetStatusCode(v int32) *ListTagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagResourcesResponse) SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse {
	s.Body = v
	return s
}

type ListUserClusterTypesResponseBody struct {
	// List of cluster types. The number of array elements N ranges from 1 to 100.
	ClusterTypes []*ListUserClusterTypesResponseBodyClusterTypes `json:"ClusterTypes,omitempty" xml:"ClusterTypes,omitempty" type:"Repeated"`
	// The NextToken for the next page. Include this value when requesting the next page.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListUserClusterTypesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserClusterTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponseBody) SetClusterTypes(v []*ListUserClusterTypesResponseBodyClusterTypes) *ListUserClusterTypesResponseBody {
	s.ClusterTypes = v
	return s
}

func (s *ListUserClusterTypesResponseBody) SetNextToken(v string) *ListUserClusterTypesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListUserClusterTypesResponseBody) SetRequestId(v string) *ListUserClusterTypesResponseBody {
	s.RequestId = &v
	return s
}

type ListUserClusterTypesResponseBodyClusterTypes struct {
	// 访问类型。
	//
	// example:
	//
	// Public
	AccessType *string `json:"AccessType,omitempty" xml:"AccessType,omitempty"`
	// Type name
	//
	// example:
	//
	// AckEdgePro
	TypeName *string `json:"TypeName,omitempty" xml:"TypeName,omitempty"`
}

func (s ListUserClusterTypesResponseBodyClusterTypes) String() string {
	return tea.Prettify(s)
}

func (s ListUserClusterTypesResponseBodyClusterTypes) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) SetAccessType(v string) *ListUserClusterTypesResponseBodyClusterTypes {
	s.AccessType = &v
	return s
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) SetTypeName(v string) *ListUserClusterTypesResponseBodyClusterTypes {
	s.TypeName = &v
	return s
}

type ListUserClusterTypesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserClusterTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserClusterTypesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserClusterTypesResponse) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponse) SetHeaders(v map[string]*string) *ListUserClusterTypesResponse {
	s.Headers = v
	return s
}

func (s *ListUserClusterTypesResponse) SetStatusCode(v int32) *ListUserClusterTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserClusterTypesResponse) SetBody(v *ListUserClusterTypesResponseBody) *ListUserClusterTypesResponse {
	s.Body = v
	return s
}

type ListVscsRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string   `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NodeIds   []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*ListVscsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
}

func (s ListVscsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVscsRequest) GoString() string {
	return s.String()
}

func (s *ListVscsRequest) SetMaxResults(v int32) *ListVscsRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVscsRequest) SetNextToken(v string) *ListVscsRequest {
	s.NextToken = &v
	return s
}

func (s *ListVscsRequest) SetNodeIds(v []*string) *ListVscsRequest {
	s.NodeIds = v
	return s
}

func (s *ListVscsRequest) SetResourceGroupId(v string) *ListVscsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVscsRequest) SetTag(v []*ListVscsRequestTag) *ListVscsRequest {
	s.Tag = v
	return s
}

func (s *ListVscsRequest) SetVscName(v string) *ListVscsRequest {
	s.VscName = &v
	return s
}

type ListVscsRequestTag struct {
	// example:
	//
	// key001
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVscsRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListVscsRequestTag) GoString() string {
	return s.String()
}

func (s *ListVscsRequestTag) SetKey(v string) *ListVscsRequestTag {
	s.Key = &v
	return s
}

func (s *ListVscsRequestTag) SetValue(v string) *ListVscsRequestTag {
	s.Value = &v
	return s
}

type ListVscsShrinkRequest struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken     *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NodeIdsShrink *string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty"`
	// example:
	//
	// rg-aek2xdkc6icwfha
	ResourceGroupId *string                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*ListVscsShrinkRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
}

func (s ListVscsShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListVscsShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListVscsShrinkRequest) SetMaxResults(v int32) *ListVscsShrinkRequest {
	s.MaxResults = &v
	return s
}

func (s *ListVscsShrinkRequest) SetNextToken(v string) *ListVscsShrinkRequest {
	s.NextToken = &v
	return s
}

func (s *ListVscsShrinkRequest) SetNodeIdsShrink(v string) *ListVscsShrinkRequest {
	s.NodeIdsShrink = &v
	return s
}

func (s *ListVscsShrinkRequest) SetResourceGroupId(v string) *ListVscsShrinkRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVscsShrinkRequest) SetTag(v []*ListVscsShrinkRequestTag) *ListVscsShrinkRequest {
	s.Tag = v
	return s
}

func (s *ListVscsShrinkRequest) SetVscName(v string) *ListVscsShrinkRequest {
	s.VscName = &v
	return s
}

type ListVscsShrinkRequestTag struct {
	// example:
	//
	// key001
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// value001
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVscsShrinkRequestTag) String() string {
	return tea.Prettify(s)
}

func (s ListVscsShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *ListVscsShrinkRequestTag) SetKey(v string) *ListVscsShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *ListVscsShrinkRequestTag) SetValue(v string) *ListVscsShrinkRequestTag {
	s.Value = &v
	return s
}

type ListVscsResponseBody struct {
	// example:
	//
	// 0
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 3a6b93229825ac667104463b56790c91
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 03668372-18FF-5959-98D9-6B36A4643C7A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vscs       []*ListVscsResponseBodyVscs `json:"Vscs,omitempty" xml:"Vscs,omitempty" type:"Repeated"`
}

func (s ListVscsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListVscsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBody) SetMaxResults(v int32) *ListVscsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListVscsResponseBody) SetNextToken(v string) *ListVscsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListVscsResponseBody) SetRequestId(v string) *ListVscsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListVscsResponseBody) SetTotalCount(v int32) *ListVscsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListVscsResponseBody) SetVscs(v []*ListVscsResponseBodyVscs) *ListVscsResponseBody {
	s.Vscs = v
	return s
}

type ListVscsResponseBodyVscs struct {
	// example:
	//
	// e01-cn-fzh47xd7u08
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// rg-acfm2zkwhkns57i
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// example:
	//
	// NORMAL
	Status *string                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags   []*ListVscsResponseBodyVscsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// VscId
	//
	// example:
	//
	// vsc-001
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	// example:
	//
	// test_name
	VscName *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	// example:
	//
	// primary
	VscType *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s ListVscsResponseBodyVscs) String() string {
	return tea.Prettify(s)
}

func (s ListVscsResponseBodyVscs) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBodyVscs) SetNodeId(v string) *ListVscsResponseBodyVscs {
	s.NodeId = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetResourceGroupId(v string) *ListVscsResponseBodyVscs {
	s.ResourceGroupId = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetStatus(v string) *ListVscsResponseBodyVscs {
	s.Status = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetTags(v []*ListVscsResponseBodyVscsTags) *ListVscsResponseBodyVscs {
	s.Tags = v
	return s
}

func (s *ListVscsResponseBodyVscs) SetVscId(v string) *ListVscsResponseBodyVscs {
	s.VscId = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetVscName(v string) *ListVscsResponseBodyVscs {
	s.VscName = &v
	return s
}

func (s *ListVscsResponseBodyVscs) SetVscType(v string) *ListVscsResponseBodyVscs {
	s.VscType = &v
	return s
}

type ListVscsResponseBodyVscsTags struct {
	// example:
	//
	// key001
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// example:
	//
	// value001
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListVscsResponseBodyVscsTags) String() string {
	return tea.Prettify(s)
}

func (s ListVscsResponseBodyVscsTags) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBodyVscsTags) SetTagKey(v string) *ListVscsResponseBodyVscsTags {
	s.TagKey = &v
	return s
}

func (s *ListVscsResponseBodyVscsTags) SetTagValue(v string) *ListVscsResponseBodyVscsTags {
	s.TagValue = &v
	return s
}

type ListVscsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVscsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVscsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListVscsResponse) GoString() string {
	return s.String()
}

func (s *ListVscsResponse) SetHeaders(v map[string]*string) *ListVscsResponse {
	s.Headers = v
	return s
}

func (s *ListVscsResponse) SetStatusCode(v int32) *ListVscsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListVscsResponse) SetBody(v *ListVscsResponseBody) *ListVscsResponse {
	s.Body = v
	return s
}

type RebootNodesRequest struct {
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
	// List of nodes
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s RebootNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootNodesRequest) GoString() string {
	return s.String()
}

func (s *RebootNodesRequest) SetClusterId(v string) *RebootNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *RebootNodesRequest) SetIgnoreFailedNodeTasks(v bool) *RebootNodesRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *RebootNodesRequest) SetNodes(v []*string) *RebootNodesRequest {
	s.Nodes = v
	return s
}

type RebootNodesShrinkRequest struct {
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
	// List of nodes
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s RebootNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RebootNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebootNodesShrinkRequest) SetClusterId(v string) *RebootNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *RebootNodesShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *RebootNodesShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *RebootNodesShrinkRequest) SetNodesShrink(v string) *RebootNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

type RebootNodesResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task Id
	//
	// example:
	//
	// i158475611663639202234
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RebootNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RebootNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootNodesResponseBody) SetRequestId(v string) *RebootNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootNodesResponseBody) SetTaskId(v string) *RebootNodesResponseBody {
	s.TaskId = &v
	return s
}

type RebootNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s RebootNodesResponse) GoString() string {
	return s.String()
}

func (s *RebootNodesResponse) SetHeaders(v map[string]*string) *RebootNodesResponse {
	s.Headers = v
	return s
}

func (s *RebootNodesResponse) SetStatusCode(v int32) *RebootNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *RebootNodesResponse) SetBody(v *RebootNodesResponseBody) *RebootNodesResponse {
	s.Body = v
	return s
}

type ReimageNodesRequest struct {
	// Cluster ID
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Whether to allow skipping failed node tasks, default value is False
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// Node list
	Nodes []*ReimageNodesRequestNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// Custom data
	//
	// example:
	//
	// #!/bin/sh
	//
	// echo "Hello World. The time is now $(date -R)!" | tee /root/userdata_test.txt
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ReimageNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s ReimageNodesRequest) GoString() string {
	return s.String()
}

func (s *ReimageNodesRequest) SetClusterId(v string) *ReimageNodesRequest {
	s.ClusterId = &v
	return s
}

func (s *ReimageNodesRequest) SetIgnoreFailedNodeTasks(v bool) *ReimageNodesRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ReimageNodesRequest) SetNodes(v []*ReimageNodesRequestNodes) *ReimageNodesRequest {
	s.Nodes = v
	return s
}

func (s *ReimageNodesRequest) SetUserData(v string) *ReimageNodesRequest {
	s.UserData = &v
	return s
}

type ReimageNodesRequestNodes struct {
	// Hostname
	//
	// example:
	//
	// 457db5ca-241d-11ed-9fd7-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// System image ID
	//
	// example:
	//
	// m-8vbf8rpv2nn14y7oybjy
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
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
	// e01-cn-zvp2tgykr0b
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ReimageNodesRequestNodes) String() string {
	return tea.Prettify(s)
}

func (s ReimageNodesRequestNodes) GoString() string {
	return s.String()
}

func (s *ReimageNodesRequestNodes) SetHostname(v string) *ReimageNodesRequestNodes {
	s.Hostname = &v
	return s
}

func (s *ReimageNodesRequestNodes) SetImageId(v string) *ReimageNodesRequestNodes {
	s.ImageId = &v
	return s
}

func (s *ReimageNodesRequestNodes) SetLoginPassword(v string) *ReimageNodesRequestNodes {
	s.LoginPassword = &v
	return s
}

func (s *ReimageNodesRequestNodes) SetNodeId(v string) *ReimageNodesRequestNodes {
	s.NodeId = &v
	return s
}

type ReimageNodesShrinkRequest struct {
	// Cluster ID
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Whether to allow skipping failed node tasks, default value is False
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// Node list
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// Custom data
	//
	// example:
	//
	// #!/bin/sh
	//
	// echo "Hello World. The time is now $(date -R)!" | tee /root/userdata_test.txt
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s ReimageNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ReimageNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReimageNodesShrinkRequest) SetClusterId(v string) *ReimageNodesShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ReimageNodesShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *ReimageNodesShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ReimageNodesShrinkRequest) SetNodesShrink(v string) *ReimageNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *ReimageNodesShrinkRequest) SetUserData(v string) *ReimageNodesShrinkRequest {
	s.UserData = &v
	return s
}

type ReimageNodesResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 15FBCD9B-C93F-54E8-A168-AADE7E66DAD2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task ID
	//
	// example:
	//
	// i158782151663841517926
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ReimageNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReimageNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ReimageNodesResponseBody) SetRequestId(v string) *ReimageNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReimageNodesResponseBody) SetTaskId(v string) *ReimageNodesResponseBody {
	s.TaskId = &v
	return s
}

type ReimageNodesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReimageNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReimageNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s ReimageNodesResponse) GoString() string {
	return s.String()
}

func (s *ReimageNodesResponse) SetHeaders(v map[string]*string) *ReimageNodesResponse {
	s.Headers = v
	return s
}

func (s *ReimageNodesResponse) SetStatusCode(v int32) *ReimageNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *ReimageNodesResponse) SetBody(v *ReimageNodesResponseBody) *ReimageNodesResponse {
	s.Body = v
	return s
}

type RunCommandRequest struct {
	// Ensures idempotence of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters. For more information, see How to Ensure Idempotence.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Command content. Please note the following:
	//
	// - Specify `EnableParameter=true` to enable custom parameters in the command content.
	//
	// - Define custom parameters using the {{}} format; spaces and newlines within `{{}}` will be ignored.
	//
	// - The number of custom parameters cannot exceed 20.
	//
	// - Custom parameter names can only contain a-zA-Z0-9-_, and are case-insensitive.
	//
	// - A single custom parameter name cannot exceed 64 bytes.
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// Command ID
	//
	// example:
	//
	// c-e996287206324975b5fbe1d***
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// Encoding method for the script content. Valid values:
	//
	// - PlainText: No encoding, transmitted in plain text.
	//
	// - Base64: Base64 encoding.
	//
	// Default value: PlainText. If an invalid value is provided, it will be treated as PlainText.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Command description.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether the command contains custom parameters.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// Execution time for scheduled commands. Currently, three types of scheduling methods are supported: fixed interval (based on Rate expression), one-time execution at a specific time, and clock-based scheduling (based on Cron expression).
	//
	// Fixed interval execution: Based on the Rate expression, the command is executed at the set interval. The interval can be set in seconds (s), minutes (m), hours (h), and days (d), suitable for scenarios where tasks need to be executed at fixed intervals. The format is rate(<interval value><interval unit>), such as rate(5m) for every 5 minutes. The following restrictions apply to fixed interval execution:
	//
	// - The interval should not exceed 7 days and should be no less than 60 seconds, and it must be greater than the timeout of the scheduled task.
	//
	// - The interval is based on a fixed frequency and is independent of the actual execution time of the task. For example, if the command is set to execute every 5 minutes and the task takes 2 minutes to complete, the next round will start 3 minutes after the completion of the task.
	//
	// - The task will not be executed immediately upon creation. For example, if the command is set to execute every 5 minutes, it will not be executed immediately upon creation but will start 5 minutes after the task is created.
	//
	// One-time execution at a specific time: Executes the command once at the specified time and timezone. The format is at(yyyy-MM-dd HH:mm:ss <timezone>), which is at(year-month-day hour:minute:second <timezone>). If no timezone is specified, UTC is used by default. Timezones support the following three formats:
	//
	// - Full timezone name: e.g., Asia/Shanghai (China/Shanghai time), America/Los_Angeles (America/Los Angeles time), etc.
	//
	// - Timezone offset from GMT: e.g., GMT+8:00 (UTC+8:00), GMT-7:00 (UTC-7:00). When using the GMT format, leading zeros are not allowed in the hour position.
	//
	// - Timezone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	// For example, to execute once at 13:15:30 on June 6, 2022, in Shanghai, China, the format would be: at(2022-06-06 13:15:30 Asia/Shanghai); to execute once at 13:15:30 on June 6, 2022, in the GMT-7:00 timezone, the format would be: at(2022-06-06 13:15:30 GMT-7:00).
	//
	// Clock-based scheduling (based on Cron expression): Based on the Cron expression, the command is executed according to the set schedule. The format is <second> <minute> <hour> <day> <month> <weekday> <year (optional)> <timezone>, i.e., <Cron expression> <timezone>. In the specified timezone, the scheduled task execution time is calculated based on the Cron expression and then executed. If no timezone is specified, the system\\"s internal timezone of the instance running the scheduled task is used by default. For more information about Cron expressions, see Cron Expressions. Timezones support the following three formats:
	//
	// - Full timezone name: e.g., Asia/Shanghai (China/Shanghai time), America/Los_Angeles (America/Los Angeles time), etc.
	//
	// - Timezone offset from GMT: e.g., GMT+8:00 (UTC+8:00), GMT-7:00 (UTC-7:00). When using the GMT format, leading zeros are not allowed in the hour position.
	//
	// - Timezone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	// For example, to execute the command at 10:15 AM every day in 2022 in Shanghai, China, the format would be 0 15 10 ? 	- 	- 2022 Asia/Shanghai; to execute the command every 30 minutes between 10:00 AM and 11:30 AM every day in 2022 in the GMT+8:00 timezone, the format would be 0 0/30 10-11 	- 	- ? 2022 GMT+8:00; to execute the command every 5 minutes between 2:00 PM and 2:55 PM every day in October every two years starting from 2022 in UTC, the format would be 0 0/5 14 	- 10 ? 2022/2 UTC.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// Bootstrap for script execution. The length must not exceed 1 KB.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// Command name.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of nodes.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// When the command contains custom parameters, you need to pass in key-value pairs of these custom parameters when executing the command. For example, if the command content is `echo {{name}}`, you can pass in the key-value pair `{"name":"Jack"}` through the `Parameter` parameter. The custom parameter will automatically replace the variable value `name`, resulting in a new command, which actually executes as `echo Jack`.
	//
	// The number of custom parameters ranges from 0 to 10, and you should note:
	//
	// - Keys cannot be empty strings and support up to 64 characters at most.
	//
	// - Values can be empty strings.
	//
	// - When combined with the original command content and Base64 encoded, if the command is saved, the size after Base64 encoding must not exceed 18 KB; if the command is not saved, the size after Base64 encoding must not exceed 24 KB. You can set whether to keep the command via `KeepCommand`.
	//
	// - The set of custom parameter names must be a subset of the parameter set defined during command creation. For parameters that are not passed in, you can use an empty string as a substitute.
	//
	// The default value is empty, indicating that the parameter is not set, thus disabling custom parameters.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAIdyvdIqaRY****"}
	Parameters map[string]interface{} `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Set the command execution mode. Valid values:
	//
	// - Once: Execute the command immediately.
	//
	// - Period: Execute the command at a scheduled time. When this parameter is set to `Period`, you must also specify the `Frequency` parameter.
	//
	// - NextRebootOnly: Automatically execute the command when the instance starts next time.
	//
	// - EveryReboot: Automatically execute the command every time the instance starts.
	//
	// Default value:
	//
	// - If the `Frequency` parameter is not specified, the default value is `Once`.
	//
	// - If the `Frequency` parameter is specified, regardless of whether this parameter is already set, it will be processed as `Period`.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The mode when stopping a task (manually or due to execution timeout). Possible values:
	//
	// Process: Stops the current script process. ProcessTree: Stops the current process tree (a collection of the script process and all its child processes).
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// Timeout for executing the command, in seconds. If the command cannot run due to process issues, missing modules, or the absence of the Cloud Assistant Agent, a timeout will occur. After a timeout, the command process will be forcibly terminated. Default value: 60.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The username to execute the command in the instance. The length must not exceed 255 characters.
	//
	//     For Linux systems, the command is executed by the root user by default.
	//
	// example:
	//
	// root
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// You can customize the execution path of the command. The default paths are as follows:
	//
	// - Linux instances: The default execution path is under the /home directory of the root user.
	//
	// example:
	//
	// /home/user
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s RunCommandRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) SetClientToken(v string) *RunCommandRequest {
	s.ClientToken = &v
	return s
}

func (s *RunCommandRequest) SetCommandContent(v string) *RunCommandRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandRequest) SetCommandId(v string) *RunCommandRequest {
	s.CommandId = &v
	return s
}

func (s *RunCommandRequest) SetContentEncoding(v string) *RunCommandRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandRequest) SetDescription(v string) *RunCommandRequest {
	s.Description = &v
	return s
}

func (s *RunCommandRequest) SetEnableParameter(v bool) *RunCommandRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandRequest) SetFrequency(v string) *RunCommandRequest {
	s.Frequency = &v
	return s
}

func (s *RunCommandRequest) SetLauncher(v string) *RunCommandRequest {
	s.Launcher = &v
	return s
}

func (s *RunCommandRequest) SetName(v string) *RunCommandRequest {
	s.Name = &v
	return s
}

func (s *RunCommandRequest) SetNodeIdList(v []*string) *RunCommandRequest {
	s.NodeIdList = v
	return s
}

func (s *RunCommandRequest) SetParameters(v map[string]interface{}) *RunCommandRequest {
	s.Parameters = v
	return s
}

func (s *RunCommandRequest) SetRepeatMode(v string) *RunCommandRequest {
	s.RepeatMode = &v
	return s
}

func (s *RunCommandRequest) SetTerminationMode(v string) *RunCommandRequest {
	s.TerminationMode = &v
	return s
}

func (s *RunCommandRequest) SetTimeout(v int32) *RunCommandRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandRequest) SetUsername(v string) *RunCommandRequest {
	s.Username = &v
	return s
}

func (s *RunCommandRequest) SetWorkingDir(v string) *RunCommandRequest {
	s.WorkingDir = &v
	return s
}

type RunCommandShrinkRequest struct {
	// Ensures idempotence of the request. Generate a unique parameter value from your client to ensure that the value is unique across different requests.
	//
	// ClientToken supports only ASCII characters and cannot exceed 64 characters. For more information, see How to Ensure Idempotence.
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Command content. Please note the following:
	//
	// - Specify `EnableParameter=true` to enable custom parameters in the command content.
	//
	// - Define custom parameters using the {{}} format; spaces and newlines within `{{}}` will be ignored.
	//
	// - The number of custom parameters cannot exceed 20.
	//
	// - Custom parameter names can only contain a-zA-Z0-9-_, and are case-insensitive.
	//
	// - A single custom parameter name cannot exceed 64 bytes.
	//
	// example:
	//
	// ZWNobyAxMjM=
	CommandContent *string `json:"CommandContent,omitempty" xml:"CommandContent,omitempty"`
	// Command ID
	//
	// example:
	//
	// c-e996287206324975b5fbe1d***
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// Encoding method for the script content. Valid values:
	//
	// - PlainText: No encoding, transmitted in plain text.
	//
	// - Base64: Base64 encoding.
	//
	// Default value: PlainText. If an invalid value is provided, it will be treated as PlainText.
	//
	// example:
	//
	// Base64
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// Command description.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Whether the command contains custom parameters.
	//
	// Default value: false.
	//
	// example:
	//
	// false
	EnableParameter *bool `json:"EnableParameter,omitempty" xml:"EnableParameter,omitempty"`
	// Execution time for scheduled commands. Currently, three types of scheduling methods are supported: fixed interval (based on Rate expression), one-time execution at a specific time, and clock-based scheduling (based on Cron expression).
	//
	// Fixed interval execution: Based on the Rate expression, the command is executed at the set interval. The interval can be set in seconds (s), minutes (m), hours (h), and days (d), suitable for scenarios where tasks need to be executed at fixed intervals. The format is rate(<interval value><interval unit>), such as rate(5m) for every 5 minutes. The following restrictions apply to fixed interval execution:
	//
	// - The interval should not exceed 7 days and should be no less than 60 seconds, and it must be greater than the timeout of the scheduled task.
	//
	// - The interval is based on a fixed frequency and is independent of the actual execution time of the task. For example, if the command is set to execute every 5 minutes and the task takes 2 minutes to complete, the next round will start 3 minutes after the completion of the task.
	//
	// - The task will not be executed immediately upon creation. For example, if the command is set to execute every 5 minutes, it will not be executed immediately upon creation but will start 5 minutes after the task is created.
	//
	// One-time execution at a specific time: Executes the command once at the specified time and timezone. The format is at(yyyy-MM-dd HH:mm:ss <timezone>), which is at(year-month-day hour:minute:second <timezone>). If no timezone is specified, UTC is used by default. Timezones support the following three formats:
	//
	// - Full timezone name: e.g., Asia/Shanghai (China/Shanghai time), America/Los_Angeles (America/Los Angeles time), etc.
	//
	// - Timezone offset from GMT: e.g., GMT+8:00 (UTC+8:00), GMT-7:00 (UTC-7:00). When using the GMT format, leading zeros are not allowed in the hour position.
	//
	// - Timezone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	// For example, to execute once at 13:15:30 on June 6, 2022, in Shanghai, China, the format would be: at(2022-06-06 13:15:30 Asia/Shanghai); to execute once at 13:15:30 on June 6, 2022, in the GMT-7:00 timezone, the format would be: at(2022-06-06 13:15:30 GMT-7:00).
	//
	// Clock-based scheduling (based on Cron expression): Based on the Cron expression, the command is executed according to the set schedule. The format is <second> <minute> <hour> <day> <month> <weekday> <year (optional)> <timezone>, i.e., <Cron expression> <timezone>. In the specified timezone, the scheduled task execution time is calculated based on the Cron expression and then executed. If no timezone is specified, the system\\"s internal timezone of the instance running the scheduled task is used by default. For more information about Cron expressions, see Cron Expressions. Timezones support the following three formats:
	//
	// - Full timezone name: e.g., Asia/Shanghai (China/Shanghai time), America/Los_Angeles (America/Los Angeles time), etc.
	//
	// - Timezone offset from GMT: e.g., GMT+8:00 (UTC+8:00), GMT-7:00 (UTC-7:00). When using the GMT format, leading zeros are not allowed in the hour position.
	//
	// - Timezone abbreviation: Only UTC (Coordinated Universal Time) is supported.
	//
	// For example, to execute the command at 10:15 AM every day in 2022 in Shanghai, China, the format would be 0 15 10 ? 	- 	- 2022 Asia/Shanghai; to execute the command every 30 minutes between 10:00 AM and 11:30 AM every day in 2022 in the GMT+8:00 timezone, the format would be 0 0/30 10-11 	- 	- ? 2022 GMT+8:00; to execute the command every 5 minutes between 2:00 PM and 2:55 PM every day in October every two years starting from 2022 in UTC, the format would be 0 0/5 14 	- 10 ? 2022/2 UTC.
	Frequency *string `json:"Frequency,omitempty" xml:"Frequency,omitempty"`
	// Bootstrap for script execution. The length must not exceed 1 KB.
	//
	// example:
	//
	// python3 -u {{ACS::ScriptFileName|Ext(".py")}}
	Launcher *string `json:"Launcher,omitempty" xml:"Launcher,omitempty"`
	// Command name.
	//
	// example:
	//
	// testName
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of nodes.
	NodeIdListShrink *string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty"`
	// When the command contains custom parameters, you need to pass in key-value pairs of these custom parameters when executing the command. For example, if the command content is `echo {{name}}`, you can pass in the key-value pair `{"name":"Jack"}` through the `Parameter` parameter. The custom parameter will automatically replace the variable value `name`, resulting in a new command, which actually executes as `echo Jack`.
	//
	// The number of custom parameters ranges from 0 to 10, and you should note:
	//
	// - Keys cannot be empty strings and support up to 64 characters at most.
	//
	// - Values can be empty strings.
	//
	// - When combined with the original command content and Base64 encoded, if the command is saved, the size after Base64 encoding must not exceed 18 KB; if the command is not saved, the size after Base64 encoding must not exceed 24 KB. You can set whether to keep the command via `KeepCommand`.
	//
	// - The set of custom parameter names must be a subset of the parameter set defined during command creation. For parameters that are not passed in, you can use an empty string as a substitute.
	//
	// The default value is empty, indicating that the parameter is not set, thus disabling custom parameters.
	//
	// example:
	//
	// {"name":"Jack", "accessKey":"LTAIdyvdIqaRY****"}
	ParametersShrink *string `json:"Parameters,omitempty" xml:"Parameters,omitempty"`
	// Set the command execution mode. Valid values:
	//
	// - Once: Execute the command immediately.
	//
	// - Period: Execute the command at a scheduled time. When this parameter is set to `Period`, you must also specify the `Frequency` parameter.
	//
	// - NextRebootOnly: Automatically execute the command when the instance starts next time.
	//
	// - EveryReboot: Automatically execute the command every time the instance starts.
	//
	// Default value:
	//
	// - If the `Frequency` parameter is not specified, the default value is `Once`.
	//
	// - If the `Frequency` parameter is specified, regardless of whether this parameter is already set, it will be processed as `Period`.
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// The mode when stopping a task (manually or due to execution timeout). Possible values:
	//
	// Process: Stops the current script process. ProcessTree: Stops the current process tree (a collection of the script process and all its child processes).
	//
	// example:
	//
	// ProcessTree
	TerminationMode *string `json:"TerminationMode,omitempty" xml:"TerminationMode,omitempty"`
	// Timeout for executing the command, in seconds. If the command cannot run due to process issues, missing modules, or the absence of the Cloud Assistant Agent, a timeout will occur. After a timeout, the command process will be forcibly terminated. Default value: 60.
	//
	// example:
	//
	// 3600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
	// The username to execute the command in the instance. The length must not exceed 255 characters.
	//
	//     For Linux systems, the command is executed by the root user by default.
	//
	// example:
	//
	// root
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
	// You can customize the execution path of the command. The default paths are as follows:
	//
	// - Linux instances: The default execution path is under the /home directory of the root user.
	//
	// example:
	//
	// /home/user
	WorkingDir *string `json:"WorkingDir,omitempty" xml:"WorkingDir,omitempty"`
}

func (s RunCommandShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s RunCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequest) SetClientToken(v string) *RunCommandShrinkRequest {
	s.ClientToken = &v
	return s
}

func (s *RunCommandShrinkRequest) SetCommandContent(v string) *RunCommandShrinkRequest {
	s.CommandContent = &v
	return s
}

func (s *RunCommandShrinkRequest) SetCommandId(v string) *RunCommandShrinkRequest {
	s.CommandId = &v
	return s
}

func (s *RunCommandShrinkRequest) SetContentEncoding(v string) *RunCommandShrinkRequest {
	s.ContentEncoding = &v
	return s
}

func (s *RunCommandShrinkRequest) SetDescription(v string) *RunCommandShrinkRequest {
	s.Description = &v
	return s
}

func (s *RunCommandShrinkRequest) SetEnableParameter(v bool) *RunCommandShrinkRequest {
	s.EnableParameter = &v
	return s
}

func (s *RunCommandShrinkRequest) SetFrequency(v string) *RunCommandShrinkRequest {
	s.Frequency = &v
	return s
}

func (s *RunCommandShrinkRequest) SetLauncher(v string) *RunCommandShrinkRequest {
	s.Launcher = &v
	return s
}

func (s *RunCommandShrinkRequest) SetName(v string) *RunCommandShrinkRequest {
	s.Name = &v
	return s
}

func (s *RunCommandShrinkRequest) SetNodeIdListShrink(v string) *RunCommandShrinkRequest {
	s.NodeIdListShrink = &v
	return s
}

func (s *RunCommandShrinkRequest) SetParametersShrink(v string) *RunCommandShrinkRequest {
	s.ParametersShrink = &v
	return s
}

func (s *RunCommandShrinkRequest) SetRepeatMode(v string) *RunCommandShrinkRequest {
	s.RepeatMode = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTerminationMode(v string) *RunCommandShrinkRequest {
	s.TerminationMode = &v
	return s
}

func (s *RunCommandShrinkRequest) SetTimeout(v int32) *RunCommandShrinkRequest {
	s.Timeout = &v
	return s
}

func (s *RunCommandShrinkRequest) SetUsername(v string) *RunCommandShrinkRequest {
	s.Username = &v
	return s
}

func (s *RunCommandShrinkRequest) SetWorkingDir(v string) *RunCommandShrinkRequest {
	s.WorkingDir = &v
	return s
}

type RunCommandResponseBody struct {
	// ID of the command execution.
	//
	// example:
	//
	// t-7d2a745b412b4601b2d47f6a768d*
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 2FE2B22C-CF9D-59DE-BF63-DC9B9B33A9D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunCommandResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

type RunCommandResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCommandResponse) String() string {
	return tea.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) SetHeaders(v map[string]*string) *RunCommandResponse {
	s.Headers = v
	return s
}

func (s *RunCommandResponse) SetStatusCode(v int32) *RunCommandResponse {
	s.StatusCode = &v
	return s
}

func (s *RunCommandResponse) SetBody(v *RunCommandResponseBody) *RunCommandResponse {
	s.Body = v
	return s
}

type SendFileRequest struct {
	// The content of the file. After Base64 encoding, the size cannot exceed 32 KB.
	//
	// - When the `ContentType` parameter is `PlainText`, this field is plain text.
	//
	// - When the `ContentType` parameter is `Base64`, this field is Base64 encoded text.
	//
	// This parameter is required.
	//
	// example:
	//
	// #!/bin/bash echo "Current User is :" echo $(ps | grep "$$" | awk \\"{print $2}\\") -------- oss://bucketName/objectName
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type of the file.
	//
	// PlainText: Plain text.
	//
	// Base64: Base64 encoded.
	//
	// The default value is PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Description information. Supports all character sets, and the length must not exceed 512 characters.
	//
	// example:
	//
	// This is a test file.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The group of the file. Applies only to Linux instances, and the default is root. The length must not exceed 64 characters.
	//
	// Note
	//
	// When using other groups, ensure that the group exists in the instance.
	//
	// example:
	//
	// test
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions of the file. Applies only to Linux instances, and the setting method is the same as the chmod command.
	//
	// The default value is 0644, which means the user has read and write permissions, while the group and other users have read-only permissions.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file. Applies only to Linux instances, and the default is root.
	//
	// example:
	//
	// root
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The name of the file. Supports all character sets, and the length must not exceed 255 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// file.txt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of nodes.
	//
	// This parameter is required.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
	// Whether to overwrite the file if a file with the same name already exists in the target directory.
	//
	// - true: Overwrite.
	//
	// - false: Do not overwrite.
	//
	// The default value is false.
	//
	// example:
	//
	// True
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The directory in the target Lingjun node where the file will be sent. If it does not exist, it will be automatically created.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
	// The timeout for sending the file. Unit: seconds.
	//
	// - A timeout may occur due to process reasons, missing modules, or missing Cloud Assistant Agent.
	//
	// - If the set timeout is less than 10 seconds, to ensure successful delivery, the system will automatically set the timeout to 10 seconds.
	//
	// The default value is 60.
	//
	// example:
	//
	// 600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s SendFileRequest) String() string {
	return tea.Prettify(s)
}

func (s SendFileRequest) GoString() string {
	return s.String()
}

func (s *SendFileRequest) SetContent(v string) *SendFileRequest {
	s.Content = &v
	return s
}

func (s *SendFileRequest) SetContentType(v string) *SendFileRequest {
	s.ContentType = &v
	return s
}

func (s *SendFileRequest) SetDescription(v string) *SendFileRequest {
	s.Description = &v
	return s
}

func (s *SendFileRequest) SetFileGroup(v string) *SendFileRequest {
	s.FileGroup = &v
	return s
}

func (s *SendFileRequest) SetFileMode(v string) *SendFileRequest {
	s.FileMode = &v
	return s
}

func (s *SendFileRequest) SetFileOwner(v string) *SendFileRequest {
	s.FileOwner = &v
	return s
}

func (s *SendFileRequest) SetName(v string) *SendFileRequest {
	s.Name = &v
	return s
}

func (s *SendFileRequest) SetNodeIdList(v []*string) *SendFileRequest {
	s.NodeIdList = v
	return s
}

func (s *SendFileRequest) SetOverwrite(v bool) *SendFileRequest {
	s.Overwrite = &v
	return s
}

func (s *SendFileRequest) SetTargetDir(v string) *SendFileRequest {
	s.TargetDir = &v
	return s
}

func (s *SendFileRequest) SetTimeout(v int32) *SendFileRequest {
	s.Timeout = &v
	return s
}

type SendFileShrinkRequest struct {
	// The content of the file. After Base64 encoding, the size cannot exceed 32 KB.
	//
	// - When the `ContentType` parameter is `PlainText`, this field is plain text.
	//
	// - When the `ContentType` parameter is `Base64`, this field is Base64 encoded text.
	//
	// This parameter is required.
	//
	// example:
	//
	// #!/bin/bash echo "Current User is :" echo $(ps | grep "$$" | awk \\"{print $2}\\") -------- oss://bucketName/objectName
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The content type of the file.
	//
	// PlainText: Plain text.
	//
	// Base64: Base64 encoded.
	//
	// The default value is PlainText.
	//
	// example:
	//
	// PlainText
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// Description information. Supports all character sets, and the length must not exceed 512 characters.
	//
	// example:
	//
	// This is a test file.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The group of the file. Applies only to Linux instances, and the default is root. The length must not exceed 64 characters.
	//
	// Note
	//
	// When using other groups, ensure that the group exists in the instance.
	//
	// example:
	//
	// test
	FileGroup *string `json:"FileGroup,omitempty" xml:"FileGroup,omitempty"`
	// The permissions of the file. Applies only to Linux instances, and the setting method is the same as the chmod command.
	//
	// The default value is 0644, which means the user has read and write permissions, while the group and other users have read-only permissions.
	//
	// example:
	//
	// 0644
	FileMode *string `json:"FileMode,omitempty" xml:"FileMode,omitempty"`
	// The owner of the file. Applies only to Linux instances, and the default is root.
	//
	// example:
	//
	// root
	FileOwner *string `json:"FileOwner,omitempty" xml:"FileOwner,omitempty"`
	// The name of the file. Supports all character sets, and the length must not exceed 255 characters.
	//
	// This parameter is required.
	//
	// example:
	//
	// file.txt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// List of nodes.
	//
	// This parameter is required.
	NodeIdListShrink *string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty"`
	// Whether to overwrite the file if a file with the same name already exists in the target directory.
	//
	// - true: Overwrite.
	//
	// - false: Do not overwrite.
	//
	// The default value is false.
	//
	// example:
	//
	// True
	Overwrite *bool `json:"Overwrite,omitempty" xml:"Overwrite,omitempty"`
	// The directory in the target Lingjun node where the file will be sent. If it does not exist, it will be automatically created.
	//
	// This parameter is required.
	//
	// example:
	//
	// /home
	TargetDir *string `json:"TargetDir,omitempty" xml:"TargetDir,omitempty"`
	// The timeout for sending the file. Unit: seconds.
	//
	// - A timeout may occur due to process reasons, missing modules, or missing Cloud Assistant Agent.
	//
	// - If the set timeout is less than 10 seconds, to ensure successful delivery, the system will automatically set the timeout to 10 seconds.
	//
	// The default value is 60.
	//
	// example:
	//
	// 600
	Timeout *int32 `json:"Timeout,omitempty" xml:"Timeout,omitempty"`
}

func (s SendFileShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SendFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendFileShrinkRequest) SetContent(v string) *SendFileShrinkRequest {
	s.Content = &v
	return s
}

func (s *SendFileShrinkRequest) SetContentType(v string) *SendFileShrinkRequest {
	s.ContentType = &v
	return s
}

func (s *SendFileShrinkRequest) SetDescription(v string) *SendFileShrinkRequest {
	s.Description = &v
	return s
}

func (s *SendFileShrinkRequest) SetFileGroup(v string) *SendFileShrinkRequest {
	s.FileGroup = &v
	return s
}

func (s *SendFileShrinkRequest) SetFileMode(v string) *SendFileShrinkRequest {
	s.FileMode = &v
	return s
}

func (s *SendFileShrinkRequest) SetFileOwner(v string) *SendFileShrinkRequest {
	s.FileOwner = &v
	return s
}

func (s *SendFileShrinkRequest) SetName(v string) *SendFileShrinkRequest {
	s.Name = &v
	return s
}

func (s *SendFileShrinkRequest) SetNodeIdListShrink(v string) *SendFileShrinkRequest {
	s.NodeIdListShrink = &v
	return s
}

func (s *SendFileShrinkRequest) SetOverwrite(v bool) *SendFileShrinkRequest {
	s.Overwrite = &v
	return s
}

func (s *SendFileShrinkRequest) SetTargetDir(v string) *SendFileShrinkRequest {
	s.TargetDir = &v
	return s
}

func (s *SendFileShrinkRequest) SetTimeout(v int32) *SendFileShrinkRequest {
	s.Timeout = &v
	return s
}

type SendFileResponseBody struct {
	// Command execution ID.
	//
	// example:
	//
	// t-hz03la52z1zkvls
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// ID of the request
	//
	// example:
	//
	// 3C683243-7915-57FB-9570-A2932C1C0F78
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SendFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SendFileResponseBody) GoString() string {
	return s.String()
}

func (s *SendFileResponseBody) SetInvokeId(v string) *SendFileResponseBody {
	s.InvokeId = &v
	return s
}

func (s *SendFileResponseBody) SetRequestId(v string) *SendFileResponseBody {
	s.RequestId = &v
	return s
}

type SendFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SendFileResponse) GoString() string {
	return s.String()
}

func (s *SendFileResponse) SetHeaders(v map[string]*string) *SendFileResponse {
	s.Headers = v
	return s
}

func (s *SendFileResponse) SetStatusCode(v int32) *SendFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SendFileResponse) SetBody(v *SendFileResponseBody) *SendFileResponse {
	s.Body = v
	return s
}

type ShrinkClusterRequest struct {
	// Cluster ID
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Whether to allow skipping failed node tasks, default value is False
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// Node group information
	NodeGroups []*ShrinkClusterRequestNodeGroups `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty" type:"Repeated"`
}

func (s ShrinkClusterRequest) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterRequest) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequest) SetClusterId(v string) *ShrinkClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *ShrinkClusterRequest) SetIgnoreFailedNodeTasks(v bool) *ShrinkClusterRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ShrinkClusterRequest) SetNodeGroups(v []*ShrinkClusterRequestNodeGroups) *ShrinkClusterRequest {
	s.NodeGroups = v
	return s
}

type ShrinkClusterRequestNodeGroups struct {
	// Node group ID
	//
	// example:
	//
	// ng-3b6fbd24b1b845a0
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// List of nodes
	Nodes []*ShrinkClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s ShrinkClusterRequestNodeGroups) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequestNodeGroups) SetNodeGroupId(v string) *ShrinkClusterRequestNodeGroups {
	s.NodeGroupId = &v
	return s
}

func (s *ShrinkClusterRequestNodeGroups) SetNodes(v []*ShrinkClusterRequestNodeGroupsNodes) *ShrinkClusterRequestNodeGroups {
	s.Nodes = v
	return s
}

type ShrinkClusterRequestNodeGroupsNodes struct {
	// Node ID
	//
	// example:
	//
	// e01poc-cn-zmb2ypjdc01
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
}

func (s ShrinkClusterRequestNodeGroupsNodes) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequestNodeGroupsNodes) SetNodeId(v string) *ShrinkClusterRequestNodeGroupsNodes {
	s.NodeId = &v
	return s
}

type ShrinkClusterShrinkRequest struct {
	// Cluster ID
	//
	// example:
	//
	// i15dfa12e8f27c44f4a006c2c8bb
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// Whether to allow skipping failed node tasks, default value is False
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// Node group information
	NodeGroupsShrink *string `json:"NodeGroups,omitempty" xml:"NodeGroups,omitempty"`
}

func (s ShrinkClusterShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ShrinkClusterShrinkRequest) SetClusterId(v string) *ShrinkClusterShrinkRequest {
	s.ClusterId = &v
	return s
}

func (s *ShrinkClusterShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *ShrinkClusterShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *ShrinkClusterShrinkRequest) SetNodeGroupsShrink(v string) *ShrinkClusterShrinkRequest {
	s.NodeGroupsShrink = &v
	return s
}

type ShrinkClusterResponseBody struct {
	// Request ID
	//
	// example:
	//
	// CC9FEF89-9BE5-5E03-845E-238B48D7599B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// task id
	//
	// example:
	//
	// i159136551662516768776
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ShrinkClusterResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ShrinkClusterResponseBody) SetRequestId(v string) *ShrinkClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShrinkClusterResponseBody) SetTaskId(v string) *ShrinkClusterResponseBody {
	s.TaskId = &v
	return s
}

type ShrinkClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShrinkClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShrinkClusterResponse) String() string {
	return tea.Prettify(s)
}

func (s ShrinkClusterResponse) GoString() string {
	return s.String()
}

func (s *ShrinkClusterResponse) SetHeaders(v map[string]*string) *ShrinkClusterResponse {
	s.Headers = v
	return s
}

func (s *ShrinkClusterResponse) SetStatusCode(v int32) *ShrinkClusterResponse {
	s.StatusCode = &v
	return s
}

func (s *ShrinkClusterResponse) SetBody(v *ShrinkClusterResponseBody) *ShrinkClusterResponse {
	s.Body = v
	return s
}

type StopInvocationRequest struct {
	// Command execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f-hz044748dzepds0
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// List of nodes.
	NodeIdList []*string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty" type:"Repeated"`
}

func (s StopInvocationRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationRequest) SetInvokeId(v string) *StopInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationRequest) SetNodeIdList(v []*string) *StopInvocationRequest {
	s.NodeIdList = v
	return s
}

type StopInvocationShrinkRequest struct {
	// Command execution ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// f-hz044748dzepds0
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// List of nodes.
	NodeIdListShrink *string `json:"NodeIdList,omitempty" xml:"NodeIdList,omitempty"`
}

func (s StopInvocationShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationShrinkRequest) SetInvokeId(v string) *StopInvocationShrinkRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationShrinkRequest) SetNodeIdListShrink(v string) *StopInvocationShrinkRequest {
	s.NodeIdListShrink = &v
	return s
}

type StopInvocationResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// A7FD7411-9395-52E8-AF42-EB3A4A55446D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopInvocationResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *StopInvocationResponseBody) SetRequestId(v string) *StopInvocationResponseBody {
	s.RequestId = &v
	return s
}

type StopInvocationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInvocationResponse) String() string {
	return tea.Prettify(s)
}

func (s StopInvocationResponse) GoString() string {
	return s.String()
}

func (s *StopInvocationResponse) SetHeaders(v map[string]*string) *StopInvocationResponse {
	s.Headers = v
	return s
}

func (s *StopInvocationResponse) SetStatusCode(v int32) *StopInvocationResponse {
	s.StatusCode = &v
	return s
}

func (s *StopInvocationResponse) SetBody(v *StopInvocationResponseBody) *StopInvocationResponse {
	s.Body = v
	return s
}

type StopNodesRequest struct {
	// Whether to allow skipping failed node tasks, default value is False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// List of nodes.
	Nodes []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s StopNodesRequest) String() string {
	return tea.Prettify(s)
}

func (s StopNodesRequest) GoString() string {
	return s.String()
}

func (s *StopNodesRequest) SetIgnoreFailedNodeTasks(v bool) *StopNodesRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *StopNodesRequest) SetNodes(v []*string) *StopNodesRequest {
	s.Nodes = v
	return s
}

type StopNodesShrinkRequest struct {
	// Whether to allow skipping failed node tasks, default value is False.
	//
	// example:
	//
	// False
	IgnoreFailedNodeTasks *bool `json:"IgnoreFailedNodeTasks,omitempty" xml:"IgnoreFailedNodeTasks,omitempty"`
	// List of nodes.
	NodesShrink *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
}

func (s StopNodesShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s StopNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopNodesShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *StopNodesShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *StopNodesShrinkRequest) SetNodesShrink(v string) *StopNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

type StopNodesResponseBody struct {
	// Request ID
	//
	// example:
	//
	// 4FD06DF0-9167-5C6F-A145-F30CA4A15D54
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task ID
	//
	// example:
	//
	// i155847351716171893489
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s StopNodesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s StopNodesResponseBody) GoString() string {
	return s.String()
}

func (s *StopNodesResponseBody) SetRequestId(v string) *StopNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopNodesResponseBody) SetTaskId(v string) *StopNodesResponseBody {
	s.TaskId = &v
	return s
}

type StopNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopNodesResponse) String() string {
	return tea.Prettify(s)
}

func (s StopNodesResponse) GoString() string {
	return s.String()
}

func (s *StopNodesResponse) SetHeaders(v map[string]*string) *StopNodesResponse {
	s.Headers = v
	return s
}

func (s *StopNodesResponse) SetStatusCode(v int32) *StopNodesResponse {
	s.StatusCode = &v
	return s
}

func (s *StopNodesResponse) SetBody(v *StopNodesResponseBody) *StopNodesResponse {
	s.Body = v
	return s
}

type TagResourcesRequest struct {
	// Region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// List of resource IDs
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// Resource type
	//
	// This parameter is required.
	//
	// example:
	//
	// Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// Tags
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) SetRegionId(v string) *TagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *TagResourcesRequest) SetResourceId(v []*string) *TagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *TagResourcesRequest) SetResourceType(v string) *TagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

type TagResourcesRequestTag struct {
	// Tag key
	//
	// example:
	//
	// app
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// Tag value
	//
	// example:
	//
	// v3
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s TagResourcesRequestTag) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

type TagResourcesResponseBody struct {
	// ID of the request
	//
	// example:
	//
	// E7BB53E1-0B08-5C4E-BA66-9225548C3151
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s TagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) SetHeaders(v map[string]*string) *TagResourcesResponse {
	s.Headers = v
	return s
}

func (s *TagResourcesResponse) SetStatusCode(v int32) *TagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *TagResourcesResponse) SetBody(v *TagResourcesResponseBody) *TagResourcesResponse {
	s.Body = v
	return s
}

type UntagResourcesRequest struct {
	// Whether to remove all, only effective when TagKey.N is empty. Valid values:
	//
	// - True, remove all
	//
	// - False, do not remove all
	//
	// Default is False
	//
	// example:
	//
	// False
	All *bool `json:"All,omitempty" xml:"All,omitempty"`
	// Region ID
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// List of resource IDs
	//
	// This parameter is required.
	ResourceId []*string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty" type:"Repeated"`
	// Resource type
	//
	// This parameter is required.
	//
	// example:
	//
	// Cluster
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// List of tag keys
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) SetAll(v bool) *UntagResourcesRequest {
	s.All = &v
	return s
}

func (s *UntagResourcesRequest) SetRegionId(v string) *UntagResourcesRequest {
	s.RegionId = &v
	return s
}

func (s *UntagResourcesRequest) SetResourceId(v []*string) *UntagResourcesRequest {
	s.ResourceId = v
	return s
}

func (s *UntagResourcesRequest) SetResourceType(v string) *UntagResourcesRequest {
	s.ResourceType = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

type UntagResourcesResponseBody struct {
	// request id
	//
	// example:
	//
	// 81F648D0-5570-5351-AE98-6F501C7E957F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UntagResourcesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return tea.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) SetHeaders(v map[string]*string) *UntagResourcesResponse {
	s.Headers = v
	return s
}

func (s *UntagResourcesResponse) SetStatusCode(v int32) *UntagResourcesResponse {
	s.StatusCode = &v
	return s
}

func (s *UntagResourcesResponse) SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse {
	s.Body = v
	return s
}

type UpdateNodeGroupRequest struct {
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The name of the key pair.
	//
	// example:
	//
	// sc-key
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// Node group name
	//
	// example:
	//
	// test-update
	NewNodeGroupName *string `json:"NewNodeGroupName,omitempty" xml:"NewNodeGroupName,omitempty"`
	// Node group ID
	//
	// example:
	//
	// i120021051733814190732
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// user data
	//
	// example:
	//
	// #!/bin/bash
	//
	// uptime
	//
	// echo "aaaaaaa" >> /tmp/ttttt20250110141010.sh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateNodeGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupRequest) SetFileSystemMountEnabled(v bool) *UpdateNodeGroupRequest {
	s.FileSystemMountEnabled = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetKeyPairName(v string) *UpdateNodeGroupRequest {
	s.KeyPairName = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetNewNodeGroupName(v string) *UpdateNodeGroupRequest {
	s.NewNodeGroupName = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetNodeGroupId(v string) *UpdateNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *UpdateNodeGroupRequest) SetUserData(v string) *UpdateNodeGroupRequest {
	s.UserData = &v
	return s
}

type UpdateNodeGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Task ID
	//
	// example:
	//
	// i15374011238111706
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateNodeGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupResponseBody) SetRequestId(v string) *UpdateNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeGroupResponseBody) SetTaskId(v string) *UpdateNodeGroupResponseBody {
	s.TaskId = &v
	return s
}

type UpdateNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupResponse) SetHeaders(v map[string]*string) *UpdateNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateNodeGroupResponse) SetStatusCode(v int32) *UpdateNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateNodeGroupResponse) SetBody(v *UpdateNodeGroupResponseBody) *UpdateNodeGroupResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("eflo-controller"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Approve Operation
//
// @param request - ApproveOperationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ApproveOperationResponse
func (client *Client) ApproveOperationWithOptions(request *ApproveOperationRequest, runtime *util.RuntimeOptions) (_result *ApproveOperationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.OperationType)) {
		body["OperationType"] = request.OperationType
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ApproveOperation"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ApproveOperationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Approve Operation
//
// @param request - ApproveOperationRequest
//
// @return ApproveOperationResponse
func (client *Client) ApproveOperation(request *ApproveOperationRequest) (_result *ApproveOperationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ApproveOperationResponse{}
	_body, _err := client.ApproveOperationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Target Resource Group
//
// @param request - ChangeResourceGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *util.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceRegionId)) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ChangeResourceGroup"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Target Resource Group
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ChangeResourceGroupResponse{}
	_body, _err := client.ChangeResourceGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Disconnect Connection
//
// Description:
//
// # An interface for creating a session, returning the front-end EndPoint, and initiating a periodic task to track the session status
//
// @param request - CloseSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CloseSessionResponse
func (client *Client) CloseSessionWithOptions(request *CloseSessionRequest, runtime *util.RuntimeOptions) (_result *CloseSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.SessionId)) {
		body["SessionId"] = request.SessionId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionToken)) {
		body["SessionToken"] = request.SessionToken
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CloseSession"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CloseSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Disconnect Connection
//
// Description:
//
// # An interface for creating a session, returning the front-end EndPoint, and initiating a periodic task to track the session status
//
// @param request - CloseSessionRequest
//
// @return CloseSessionResponse
func (client *Client) CloseSession(request *CloseSessionRequest) (_result *CloseSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CloseSessionResponse{}
	_body, _err := client.CloseSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create a large-scale computing cluster
//
// @param tmpReq - CreateClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateClusterResponse
func (client *Client) CreateClusterWithOptions(tmpReq *CreateClusterRequest, runtime *util.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Components)) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, tea.String("Components"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Networks)) {
		request.NetworksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Networks, tea.String("Networks"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NimizVSwitches)) {
		request.NimizVSwitchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NimizVSwitches, tea.String("NimizVSwitches"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodeGroups)) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, tea.String("NodeGroups"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterDescription)) {
		body["ClusterDescription"] = request.ClusterDescription
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		body["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterType)) {
		body["ClusterType"] = request.ClusterType
	}

	if !tea.BoolValue(util.IsUnset(request.ComponentsShrink)) {
		body["Components"] = request.ComponentsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.HpnZone)) {
		body["HpnZone"] = request.HpnZone
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.NetworksShrink)) {
		body["Networks"] = request.NetworksShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NimizVSwitchesShrink)) {
		body["NimizVSwitches"] = request.NimizVSwitchesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupsShrink)) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.OpenEniJumboFrame)) {
		body["OpenEniJumboFrame"] = request.OpenEniJumboFrame
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCluster"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create a large-scale computing cluster
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateClusterResponse{}
	_body, _err := client.CreateClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Diagnostic Task Creation Interface
//
// @param tmpReq - CreateDiagnosticTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDiagnosticTaskResponse
func (client *Client) CreateDiagnosticTaskWithOptions(tmpReq *CreateDiagnosticTaskRequest, runtime *util.RuntimeOptions) (_result *CreateDiagnosticTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateDiagnosticTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.AiJobLogInfo)) {
		request.AiJobLogInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AiJobLogInfo, tea.String("AiJobLogInfo"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodeIds)) {
		request.NodeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIds, tea.String("NodeIds"), tea.String("simple"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AiJobLogInfoShrink)) {
		body["AiJobLogInfo"] = request.AiJobLogInfoShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.DiagnosticType)) {
		body["DiagnosticType"] = request.DiagnosticType
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdsShrink)) {
		body["NodeIds"] = request.NodeIdsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDiagnosticTask"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDiagnosticTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Diagnostic Task Creation Interface
//
// @param request - CreateDiagnosticTaskRequest
//
// @return CreateDiagnosticTaskResponse
func (client *Client) CreateDiagnosticTask(request *CreateDiagnosticTaskRequest) (_result *CreateDiagnosticTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateDiagnosticTaskResponse{}
	_body, _err := client.CreateDiagnosticTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Network Test Task
//
// Description:
//
// An interface for creating a session, which returns the frontend EndPoint and initiates a periodic task to track the session status.
//
// @param tmpReq - CreateNetTestTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNetTestTaskResponse
func (client *Client) CreateNetTestTaskWithOptions(tmpReq *CreateNetTestTaskRequest, runtime *util.RuntimeOptions) (_result *CreateNetTestTaskResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateNetTestTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.CommTest)) {
		request.CommTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CommTest, tea.String("CommTest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.DelayTest)) {
		request.DelayTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DelayTest, tea.String("DelayTest"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.TrafficTest)) {
		request.TrafficTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TrafficTest, tea.String("TrafficTest"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.ClusterName)) {
		body["ClusterName"] = request.ClusterName
	}

	if !tea.BoolValue(util.IsUnset(request.CommTestShrink)) {
		body["CommTest"] = request.CommTestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.DelayTestShrink)) {
		body["DelayTest"] = request.DelayTestShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NetTestType)) {
		body["NetTestType"] = request.NetTestType
	}

	if !tea.BoolValue(util.IsUnset(request.NetworkMode)) {
		body["NetworkMode"] = request.NetworkMode
	}

	if !tea.BoolValue(util.IsUnset(request.Port)) {
		body["Port"] = request.Port
	}

	if !tea.BoolValue(util.IsUnset(request.TrafficTestShrink)) {
		body["TrafficTest"] = request.TrafficTestShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNetTestTask"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNetTestTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Network Test Task
//
// Description:
//
// An interface for creating a session, which returns the frontend EndPoint and initiates a periodic task to track the session status.
//
// @param request - CreateNetTestTaskRequest
//
// @return CreateNetTestTaskResponse
func (client *Client) CreateNetTestTask(request *CreateNetTestTaskRequest) (_result *CreateNetTestTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNetTestTaskResponse{}
	_body, _err := client.CreateNetTestTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Node Group under Cluster
//
// Description:
//
// # An interface for creating a session, which returns the frontend EndPoint and initiates a periodic task to track the session status
//
// @param tmpReq - CreateNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeGroupResponse
func (client *Client) CreateNodeGroupWithOptions(tmpReq *CreateNodeGroupRequest, runtime *util.RuntimeOptions) (_result *CreateNodeGroupResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &CreateNodeGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeGroup)) {
		request.NodeGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroup, tea.String("NodeGroup"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodeUnit)) {
		request.NodeUnitShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeUnit, tea.String("NodeUnit"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupShrink)) {
		body["NodeGroup"] = request.NodeGroupShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeUnitShrink)) {
		body["NodeUnit"] = request.NodeUnitShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateNodeGroup"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Node Group under Cluster
//
// Description:
//
// # An interface for creating a session, which returns the frontend EndPoint and initiates a periodic task to track the session status
//
// @param request - CreateNodeGroupRequest
//
// @return CreateNodeGroupResponse
func (client *Client) CreateNodeGroup(request *CreateNodeGroupRequest) (_result *CreateNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateNodeGroupResponse{}
	_body, _err := client.CreateNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Create Web Terminal Session
//
// Description:
//
// An interface for creating a session, which returns the frontend EndPoint and initiates a periodic task to track the session status.
//
// @param request - CreateSessionRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSessionResponse
func (client *Client) CreateSessionWithOptions(request *CreateSessionRequest, runtime *util.RuntimeOptions) (_result *CreateSessionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.SessionType)) {
		body["SessionType"] = request.SessionType
	}

	if !tea.BoolValue(util.IsUnset(request.StartTime)) {
		body["StartTime"] = request.StartTime
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSession"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSessionResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Create Web Terminal Session
//
// Description:
//
// An interface for creating a session, which returns the frontend EndPoint and initiates a periodic task to track the session status.
//
// @param request - CreateSessionRequest
//
// @return CreateSessionResponse
func (client *Client) CreateSession(request *CreateSessionRequest) (_result *CreateSessionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateSessionResponse{}
	_body, _err := client.CreateSessionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建Vsc
//
// @param request - CreateVscRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateVscResponse
func (client *Client) CreateVscWithOptions(request *CreateVscRequest, runtime *util.RuntimeOptions) (_result *CreateVscResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		body["Tag"] = request.Tag
	}

	if !tea.BoolValue(util.IsUnset(request.VscName)) {
		body["VscName"] = request.VscName
	}

	if !tea.BoolValue(util.IsUnset(request.VscType)) {
		body["VscType"] = request.VscType
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateVsc"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateVscResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建Vsc
//
// @param request - CreateVscRequest
//
// @return CreateVscResponse
func (client *Client) CreateVsc(request *CreateVscRequest) (_result *CreateVscResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateVscResponse{}
	_body, _err := client.CreateVscWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete cluster instance
//
// @param request - DeleteClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteClusterResponse
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *util.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteCluster"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete cluster instance
//
// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteClusterResponse{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Delete Node Group
//
// Description:
//
// An interface for creating a session, which returns the front-end EndPoint and initiates a periodic task to track the session status.
//
// @param request - DeleteNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeGroupResponse
func (client *Client) DeleteNodeGroupWithOptions(request *DeleteNodeGroupRequest, runtime *util.RuntimeOptions) (_result *DeleteNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteNodeGroup"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Delete Node Group
//
// Description:
//
// An interface for creating a session, which returns the front-end EndPoint and initiates a periodic task to track the session status.
//
// @param request - DeleteNodeGroupRequest
//
// @return DeleteNodeGroupResponse
func (client *Client) DeleteNodeGroup(request *DeleteNodeGroupRequest) (_result *DeleteNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteNodeGroupResponse{}
	_body, _err := client.DeleteNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除Vsc
//
// @param request - DeleteVscRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteVscResponse
func (client *Client) DeleteVscWithOptions(request *DeleteVscRequest, runtime *util.RuntimeOptions) (_result *DeleteVscResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VscId)) {
		body["VscId"] = request.VscId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteVsc"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteVscResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除Vsc
//
// @param request - DeleteVscRequest
//
// @return DeleteVscResponse
func (client *Client) DeleteVsc(request *DeleteVscRequest) (_result *DeleteVscResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteVscResponse{}
	_body, _err := client.DeleteVscWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Cluster Details
//
// @param request - DescribeClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeClusterResponse
func (client *Client) DescribeClusterWithOptions(request *DescribeClusterRequest, runtime *util.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeCluster"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Cluster Details
//
// @param request - DescribeClusterRequest
//
// @return DescribeClusterResponse
func (client *Client) DescribeCluster(request *DescribeClusterRequest) (_result *DescribeClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeClusterResponse{}
	_body, _err := client.DescribeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Diagnostic Task Query Interface
//
// Description:
//
// An interface for creating a session, which returns the front-end EndPoint and initiates a periodic task to track the session status.
//
// @param request - DescribeDiagnosticResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosticResultResponse
func (client *Client) DescribeDiagnosticResultWithOptions(request *DescribeDiagnosticResultRequest, runtime *util.RuntimeOptions) (_result *DescribeDiagnosticResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiagnosticId)) {
		body["DiagnosticId"] = request.DiagnosticId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeDiagnosticResult"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeDiagnosticResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Diagnostic Task Query Interface
//
// Description:
//
// An interface for creating a session, which returns the front-end EndPoint and initiates a periodic task to track the session status.
//
// @param request - DescribeDiagnosticResultRequest
//
// @return DescribeDiagnosticResultResponse
func (client *Client) DescribeDiagnosticResult(request *DescribeDiagnosticResultRequest) (_result *DescribeDiagnosticResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeDiagnosticResultResponse{}
	_body, _err := client.DescribeDiagnosticResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list and status of operations assistant command executions
//
// @param request - DescribeInvocationsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *util.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ContentEncoding)) {
		body["ContentEncoding"] = request.ContentEncoding
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeOutput)) {
		body["IncludeOutput"] = request.IncludeOutput
	}

	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		body["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeInvocations"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list and status of operations assistant command executions
//
// @param request - DescribeInvocationsRequest
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeInvocationsResponse{}
	_body, _err := client.DescribeInvocationsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Network Test Result
//
// Description:
//
// # An interface for creating a session, returning the front-end EndPoint, and initiating a periodic task to track the session status
//
// @param request - DescribeNetTestResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNetTestResultResponse
func (client *Client) DescribeNetTestResultWithOptions(request *DescribeNetTestResultRequest, runtime *util.RuntimeOptions) (_result *DescribeNetTestResultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TestId)) {
		body["TestId"] = request.TestId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNetTestResult"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNetTestResultResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Network Test Result
//
// Description:
//
// # An interface for creating a session, returning the front-end EndPoint, and initiating a periodic task to track the session status
//
// @param request - DescribeNetTestResultRequest
//
// @return DescribeNetTestResultResponse
func (client *Client) DescribeNetTestResult(request *DescribeNetTestResultRequest) (_result *DescribeNetTestResultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNetTestResultResponse{}
	_body, _err := client.DescribeNetTestResultWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query node list
//
// @param request - DescribeNodeRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeResponse
func (client *Client) DescribeNodeWithOptions(request *DescribeNodeRequest, runtime *util.RuntimeOptions) (_result *DescribeNodeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeNode"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeNodeResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query node list
//
// @param request - DescribeNodeRequest
//
// @return DescribeNodeResponse
func (client *Client) DescribeNode(request *DescribeNodeRequest) (_result *DescribeNodeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeNodeResponse{}
	_body, _err := client.DescribeNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Region List
//
// @param request - DescribeRegionsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *util.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeRegions"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Region List
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeRegionsResponse{}
	_body, _err := client.DescribeRegionsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list and status of files sent by the operation assistant
//
// @param request - DescribeSendFileResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeSendFileResultsResponse
func (client *Client) DescribeSendFileResultsWithOptions(request *DescribeSendFileResultsRequest, runtime *util.RuntimeOptions) (_result *DescribeSendFileResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		body["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeId)) {
		body["NodeId"] = request.NodeId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeSendFileResults"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeSendFileResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list and status of files sent by the operation assistant
//
// @param request - DescribeSendFileResultsRequest
//
// @return DescribeSendFileResultsResponse
func (client *Client) DescribeSendFileResults(request *DescribeSendFileResultsRequest) (_result *DescribeSendFileResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeSendFileResultsResponse{}
	_body, _err := client.DescribeSendFileResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Task Details
//
// @param request - DescribeTaskRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeTaskResponse
func (client *Client) DescribeTaskWithOptions(request *DescribeTaskRequest, runtime *util.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["TaskId"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeTask"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeTaskResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Task Details
//
// @param request - DescribeTaskRequest
//
// @return DescribeTaskResponse
func (client *Client) DescribeTask(request *DescribeTaskRequest) (_result *DescribeTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeTaskResponse{}
	_body, _err := client.DescribeTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取单个Vsc详情
//
// @param request - DescribeVscRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeVscResponse
func (client *Client) DescribeVscWithOptions(request *DescribeVscRequest, runtime *util.RuntimeOptions) (_result *DescribeVscResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.VscId)) {
		body["VscId"] = request.VscId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeVsc"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeVscResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取单个Vsc详情
//
// @param request - DescribeVscRequest
//
// @return DescribeVscResponse
func (client *Client) DescribeVsc(request *DescribeVscRequest) (_result *DescribeVscResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeVscResponse{}
	_body, _err := client.DescribeVscWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of Available Zones
//
// @param request - DescribeZonesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeZonesResponse
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *util.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AcceptLanguage)) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DescribeZones"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &DescribeZonesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of Available Zones
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DescribeZonesResponse{}
	_body, _err := client.DescribeZonesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Cluster Scaling
//
// @param tmpReq - ExtendClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ExtendClusterResponse
func (client *Client) ExtendClusterWithOptions(tmpReq *ExtendClusterRequest, runtime *util.RuntimeOptions) (_result *ExtendClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ExtendClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.IpAllocationPolicy)) {
		request.IpAllocationPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpAllocationPolicy, tea.String("IpAllocationPolicy"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.NodeGroups)) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, tea.String("NodeGroups"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.VpdSubnets)) {
		request.VpdSubnetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VpdSubnets, tea.String("VpdSubnets"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.IpAllocationPolicyShrink)) {
		body["IpAllocationPolicy"] = request.IpAllocationPolicyShrink
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupsShrink)) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.VSwitchZoneId)) {
		body["VSwitchZoneId"] = request.VSwitchZoneId
	}

	if !tea.BoolValue(util.IsUnset(request.VpdSubnetsShrink)) {
		body["VpdSubnets"] = request.VpdSubnetsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ExtendCluster"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ExtendClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Cluster Scaling
//
// @param request - ExtendClusterRequest
//
// @return ExtendClusterResponse
func (client *Client) ExtendCluster(request *ExtendClusterRequest) (_result *ExtendClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ExtendClusterResponse{}
	_body, _err := client.ExtendClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of host groups under the cluster, and list of hosts under each group
//
// @param request - ListClusterNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClusterNodesResponse
func (client *Client) ListClusterNodesWithOptions(request *ListClusterNodesRequest, runtime *util.RuntimeOptions) (_result *ListClusterNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusterNodes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClusterNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of host groups under the cluster, and list of hosts under each group
//
// @param request - ListClusterNodesRequest
//
// @return ListClusterNodesResponse
func (client *Client) ListClusterNodes(request *ListClusterNodesRequest) (_result *ListClusterNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClusterNodesResponse{}
	_body, _err := client.ListClusterNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Get the list of cluster instances
//
// @param request - ListClustersRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListClustersResponse
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *util.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListClusters"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListClustersResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Get the list of cluster instances
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListClustersResponse{}
	_body, _err := client.ListClustersWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of Diagnostic Tasks
//
// Description:
//
// An interface for creating a session, which returns the frontend EndPoint and initiates a periodic task to track the session status.
//
// @param request - ListDiagnosticResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDiagnosticResultsResponse
func (client *Client) ListDiagnosticResultsWithOptions(request *ListDiagnosticResultsRequest, runtime *util.RuntimeOptions) (_result *ListDiagnosticResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DiagType)) {
		body["DiagType"] = request.DiagType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDiagnosticResults"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDiagnosticResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of Diagnostic Tasks
//
// Description:
//
// An interface for creating a session, which returns the frontend EndPoint and initiates a periodic task to track the session status.
//
// @param request - ListDiagnosticResultsRequest
//
// @return ListDiagnosticResultsResponse
func (client *Client) ListDiagnosticResults(request *ListDiagnosticResultsRequest) (_result *ListDiagnosticResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListDiagnosticResultsResponse{}
	_body, _err := client.ListDiagnosticResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # List of Available Physical Machines
//
// @param request - ListFreeNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFreeNodesResponse
func (client *Client) ListFreeNodesWithOptions(request *ListFreeNodesRequest, runtime *util.RuntimeOptions) (_result *ListFreeNodesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tags)) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.HpnZone)) {
		body["HpnZone"] = request.HpnZone
	}

	if !tea.BoolValue(util.IsUnset(request.MachineType)) {
		body["MachineType"] = request.MachineType
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.OperatingStates)) {
		body["OperatingStates"] = request.OperatingStates
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFreeNodes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFreeNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # List of Available Physical Machines
//
// @param request - ListFreeNodesRequest
//
// @return ListFreeNodesResponse
func (client *Client) ListFreeNodes(request *ListFreeNodesRequest) (_result *ListFreeNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFreeNodesResponse{}
	_body, _err := client.ListFreeNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of images available to the user
//
// @param request - ListImagesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListImagesResponse
func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *util.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Architecture)) {
		body["Architecture"] = request.Architecture
	}

	if !tea.BoolValue(util.IsUnset(request.ImageVersion)) {
		body["ImageVersion"] = request.ImageVersion
	}

	if !tea.BoolValue(util.IsUnset(request.Platform)) {
		body["Platform"] = request.Platform
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListImages"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListImagesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of images available to the user
//
// @param request - ListImagesRequest
//
// @return ListImagesResponse
func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListImagesResponse{}
	_body, _err := client.ListImagesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query machine network configuration using HPNZone and machine type
//
// Description:
//
// # An interface for creating a session, returning the frontend EndPoint, and initiating a periodic task to track the session status
//
// @param tmpReq - ListMachineNetworkInfoRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMachineNetworkInfoResponse
func (client *Client) ListMachineNetworkInfoWithOptions(tmpReq *ListMachineNetworkInfoRequest, runtime *util.RuntimeOptions) (_result *ListMachineNetworkInfoResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListMachineNetworkInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.MachineHpnInfo)) {
		request.MachineHpnInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MachineHpnInfo, tea.String("MachineHpnInfo"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MachineHpnInfoShrink)) {
		body["MachineHpnInfo"] = request.MachineHpnInfoShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMachineNetworkInfo"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMachineNetworkInfoResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query machine network configuration using HPNZone and machine type
//
// Description:
//
// # An interface for creating a session, returning the frontend EndPoint, and initiating a periodic task to track the session status
//
// @param request - ListMachineNetworkInfoRequest
//
// @return ListMachineNetworkInfoResponse
func (client *Client) ListMachineNetworkInfo(request *ListMachineNetworkInfoRequest) (_result *ListMachineNetworkInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMachineNetworkInfoResponse{}
	_body, _err := client.ListMachineNetworkInfoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the list of machine types available to the user
//
// @param request - ListMachineTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMachineTypesResponse
func (client *Client) ListMachineTypesWithOptions(request *ListMachineTypesRequest, runtime *util.RuntimeOptions) (_result *ListMachineTypesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMachineTypes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMachineTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the list of machine types available to the user
//
// @param request - ListMachineTypesRequest
//
// @return ListMachineTypesResponse
func (client *Client) ListMachineTypes(request *ListMachineTypesRequest) (_result *ListMachineTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMachineTypesResponse{}
	_body, _err := client.ListMachineTypesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Network Test List
//
// Description:
//
// An interface for creating a session, returning the frontend EndPoint, and initiating a periodic task to track the session status.
//
// @param request - ListNetTestResultsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNetTestResultsResponse
func (client *Client) ListNetTestResultsWithOptions(request *ListNetTestResultsRequest, runtime *util.RuntimeOptions) (_result *ListNetTestResultsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NetTestType)) {
		body["NetTestType"] = request.NetTestType
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNetTestResults"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNetTestResultsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Network Test List
//
// Description:
//
// An interface for creating a session, returning the frontend EndPoint, and initiating a periodic task to track the session status.
//
// @param request - ListNetTestResultsRequest
//
// @return ListNetTestResultsResponse
func (client *Client) ListNetTestResults(request *ListNetTestResultsRequest) (_result *ListNetTestResultsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNetTestResultsResponse{}
	_body, _err := client.ListNetTestResultsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Node Group Information Under the Cluster
//
// @param request - ListNodeGroupsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListNodeGroupsResponse
func (client *Client) ListNodeGroupsWithOptions(request *ListNodeGroupsRequest, runtime *util.RuntimeOptions) (_result *ListNodeGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListNodeGroups"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Node Group Information Under the Cluster
//
// @param request - ListNodeGroupsRequest
//
// @return ListNodeGroupsResponse
func (client *Client) ListNodeGroups(request *ListNodeGroupsRequest) (_result *ListNodeGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListNodeGroupsResponse{}
	_body, _err := client.ListNodeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query Resource Tags
//
// @param request - ListTagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *util.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		query["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTagResources"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query Resource Tags
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagResourcesResponse{}
	_body, _err := client.ListTagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Query the cluster types available to the user
//
// Description:
//
// # An interface for creating a session, which returns the front-end EndPoint and initiates a periodic task to track the session status
//
// @param request - ListUserClusterTypesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserClusterTypesResponse
func (client *Client) ListUserClusterTypesWithOptions(runtime *util.RuntimeOptions) (_result *ListUserClusterTypesResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	params := &openapi.Params{
		Action:      tea.String("ListUserClusterTypes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserClusterTypesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Query the cluster types available to the user
//
// Description:
//
// # An interface for creating a session, which returns the front-end EndPoint and initiates a periodic task to track the session status
//
// @return ListUserClusterTypesResponse
func (client *Client) ListUserClusterTypes() (_result *ListUserClusterTypesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListUserClusterTypesResponse{}
	_body, _err := client.ListUserClusterTypesWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询Vsc列表
//
// @param tmpReq - ListVscsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVscsResponse
func (client *Client) ListVscsWithOptions(tmpReq *ListVscsRequest, runtime *util.RuntimeOptions) (_result *ListVscsResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ListVscsShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeIds)) {
		request.NodeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIds, tea.String("NodeIds"), tea.String("json"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.MaxResults)) {
		body["MaxResults"] = request.MaxResults
	}

	if !tea.BoolValue(util.IsUnset(request.NextToken)) {
		body["NextToken"] = request.NextToken
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdsShrink)) {
		body["NodeIds"] = request.NodeIdsShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceGroupId)) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.VscName)) {
		body["VscName"] = request.VscName
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListVscs"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ListVscsResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询Vsc列表
//
// @param request - ListVscsRequest
//
// @return ListVscsResponse
func (client *Client) ListVscs(request *ListVscsRequest) (_result *ListVscsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListVscsResponse{}
	_body, _err := client.ListVscsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Reboot Machine
//
// @param tmpReq - RebootNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RebootNodesResponse
func (client *Client) RebootNodesWithOptions(tmpReq *RebootNodesRequest, runtime *util.RuntimeOptions) (_result *RebootNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RebootNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Nodes)) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, tea.String("Nodes"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.NodesShrink)) {
		body["Nodes"] = request.NodesShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RebootNodes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RebootNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Reboot Machine
//
// @param request - RebootNodesRequest
//
// @return RebootNodesResponse
func (client *Client) RebootNodes(request *RebootNodesRequest) (_result *RebootNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RebootNodesResponse{}
	_body, _err := client.RebootNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Machine Reinstallation
//
// @param tmpReq - ReimageNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReimageNodesResponse
func (client *Client) ReimageNodesWithOptions(tmpReq *ReimageNodesRequest, runtime *util.RuntimeOptions) (_result *ReimageNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ReimageNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Nodes)) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, tea.String("Nodes"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.NodesShrink)) {
		body["Nodes"] = request.NodesShrink
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ReimageNodes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ReimageNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Machine Reinstallation
//
// @param request - ReimageNodesRequest
//
// @return ReimageNodesResponse
func (client *Client) ReimageNodes(request *ReimageNodesRequest) (_result *ReimageNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReimageNodesResponse{}
	_body, _err := client.ReimageNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Execute a Shell script on one or more Lingjun machines
//
// @param tmpReq - RunCommandRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RunCommandResponse
func (client *Client) RunCommandWithOptions(tmpReq *RunCommandRequest, runtime *util.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &RunCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeIdList)) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, tea.String("NodeIdList"), tea.String("json"))
	}

	if !tea.BoolValue(util.IsUnset(tmpReq.Parameters)) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, tea.String("Parameters"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientToken)) {
		body["ClientToken"] = request.ClientToken
	}

	if !tea.BoolValue(util.IsUnset(request.CommandContent)) {
		body["CommandContent"] = request.CommandContent
	}

	if !tea.BoolValue(util.IsUnset(request.CommandId)) {
		body["CommandId"] = request.CommandId
	}

	if !tea.BoolValue(util.IsUnset(request.ContentEncoding)) {
		body["ContentEncoding"] = request.ContentEncoding
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.EnableParameter)) {
		body["EnableParameter"] = request.EnableParameter
	}

	if !tea.BoolValue(util.IsUnset(request.Frequency)) {
		body["Frequency"] = request.Frequency
	}

	if !tea.BoolValue(util.IsUnset(request.Launcher)) {
		body["Launcher"] = request.Launcher
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdListShrink)) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.ParametersShrink)) {
		body["Parameters"] = request.ParametersShrink
	}

	if !tea.BoolValue(util.IsUnset(request.RepeatMode)) {
		body["RepeatMode"] = request.RepeatMode
	}

	if !tea.BoolValue(util.IsUnset(request.TerminationMode)) {
		body["TerminationMode"] = request.TerminationMode
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["Timeout"] = request.Timeout
	}

	if !tea.BoolValue(util.IsUnset(request.Username)) {
		body["Username"] = request.Username
	}

	if !tea.BoolValue(util.IsUnset(request.WorkingDir)) {
		body["WorkingDir"] = request.WorkingDir
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RunCommand"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &RunCommandResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Execute a Shell script on one or more Lingjun machines
//
// @param request - RunCommandRequest
//
// @return RunCommandResponse
func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RunCommandResponse{}
	_body, _err := client.RunCommandWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Send a remote file to one or more Lingjun machines
//
// @param tmpReq - SendFileRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SendFileResponse
func (client *Client) SendFileWithOptions(tmpReq *SendFileRequest, runtime *util.RuntimeOptions) (_result *SendFileResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &SendFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeIdList)) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, tea.String("NodeIdList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Content)) {
		body["Content"] = request.Content
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["ContentType"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["Description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.FileGroup)) {
		body["FileGroup"] = request.FileGroup
	}

	if !tea.BoolValue(util.IsUnset(request.FileMode)) {
		body["FileMode"] = request.FileMode
	}

	if !tea.BoolValue(util.IsUnset(request.FileOwner)) {
		body["FileOwner"] = request.FileOwner
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["Name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdListShrink)) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	if !tea.BoolValue(util.IsUnset(request.Overwrite)) {
		body["Overwrite"] = request.Overwrite
	}

	if !tea.BoolValue(util.IsUnset(request.TargetDir)) {
		body["TargetDir"] = request.TargetDir
	}

	if !tea.BoolValue(util.IsUnset(request.Timeout)) {
		body["Timeout"] = request.Timeout
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SendFile"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &SendFileResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Send a remote file to one or more Lingjun machines
//
// @param request - SendFileRequest
//
// @return SendFileResponse
func (client *Client) SendFile(request *SendFileRequest) (_result *SendFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SendFileResponse{}
	_body, _err := client.SendFileWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Shrink
//
// @param tmpReq - ShrinkClusterRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ShrinkClusterResponse
func (client *Client) ShrinkClusterWithOptions(tmpReq *ShrinkClusterRequest, runtime *util.RuntimeOptions) (_result *ShrinkClusterResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &ShrinkClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeGroups)) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, tea.String("NodeGroups"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClusterId)) {
		body["ClusterId"] = request.ClusterId
	}

	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupsShrink)) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ShrinkCluster"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &ShrinkClusterResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Shrink
//
// @param request - ShrinkClusterRequest
//
// @return ShrinkClusterResponse
func (client *Client) ShrinkCluster(request *ShrinkClusterRequest) (_result *ShrinkClusterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ShrinkClusterResponse{}
	_body, _err := client.ShrinkClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Stop the operation assistant command process
//
// @param tmpReq - StopInvocationRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopInvocationResponse
func (client *Client) StopInvocationWithOptions(tmpReq *StopInvocationRequest, runtime *util.RuntimeOptions) (_result *StopInvocationResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StopInvocationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.NodeIdList)) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, tea.String("NodeIdList"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.InvokeId)) {
		body["InvokeId"] = request.InvokeId
	}

	if !tea.BoolValue(util.IsUnset(request.NodeIdListShrink)) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopInvocation"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopInvocationResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Stop the operation assistant command process
//
// @param request - StopInvocationRequest
//
// @return StopInvocationResponse
func (client *Client) StopInvocation(request *StopInvocationRequest) (_result *StopInvocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopInvocationResponse{}
	_body, _err := client.StopInvocationWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Shut down the nodes
//
// @param tmpReq - StopNodesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return StopNodesResponse
func (client *Client) StopNodesWithOptions(tmpReq *StopNodesRequest, runtime *util.RuntimeOptions) (_result *StopNodesResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &StopNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Nodes)) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, tea.String("Nodes"), tea.String("json"))
	}

	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IgnoreFailedNodeTasks)) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !tea.BoolValue(util.IsUnset(request.NodesShrink)) {
		body["Nodes"] = request.NodesShrink
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("StopNodes"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &StopNodesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Shut down the nodes
//
// @param request - StopNodesRequest
//
// @return StopNodesResponse
func (client *Client) StopNodes(request *StopNodesRequest) (_result *StopNodesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &StopNodesResponse{}
	_body, _err := client.StopNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Tag User Resources
//
// @param request - TagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TagResourcesResponse
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *util.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Tag)) {
		query["Tag"] = request.Tag
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("TagResources"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Tag User Resources
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagResourcesResponse{}
	_body, _err := client.TagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Remove user tags from resources
//
// @param request - UntagResourcesRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UntagResourcesResponse
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *util.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.All)) {
		query["All"] = request.All
	}

	if !tea.BoolValue(util.IsUnset(request.RegionId)) {
		query["RegionId"] = request.RegionId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceId)) {
		query["ResourceId"] = request.ResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ResourceType)) {
		query["ResourceType"] = request.ResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.TagKey)) {
		query["TagKey"] = request.TagKey
	}

	req := &openapi.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("UntagResources"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UntagResourcesResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Remove user tags from resources
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UntagResourcesResponse{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// # Update Node Group
//
// Description:
//
// An interface for creating a session, which returns the front-end EndPoint and initiates a periodic task to track the session status.
//
// @param request - UpdateNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeGroupResponse
func (client *Client) UpdateNodeGroupWithOptions(request *UpdateNodeGroupRequest, runtime *util.RuntimeOptions) (_result *UpdateNodeGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.FileSystemMountEnabled)) {
		body["FileSystemMountEnabled"] = request.FileSystemMountEnabled
	}

	if !tea.BoolValue(util.IsUnset(request.KeyPairName)) {
		body["KeyPairName"] = request.KeyPairName
	}

	if !tea.BoolValue(util.IsUnset(request.NewNodeGroupName)) {
		body["NewNodeGroupName"] = request.NewNodeGroupName
	}

	if !tea.BoolValue(util.IsUnset(request.NodeGroupId)) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["UserData"] = request.UserData
	}

	req := &openapi.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateNodeGroup"),
		Version:     tea.String("2022-12-15"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("RPC"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateNodeGroupResponse{}
	_body, _err := client.CallApi(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// # Update Node Group
//
// Description:
//
// An interface for creating a session, which returns the front-end EndPoint and initiates a periodic task to track the session status.
//
// @param request - UpdateNodeGroupRequest
//
// @return UpdateNodeGroupResponse
func (client *Client) UpdateNodeGroup(request *UpdateNodeGroupRequest) (_result *UpdateNodeGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &UpdateNodeGroupResponse{}
	_body, _err := client.UpdateNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
