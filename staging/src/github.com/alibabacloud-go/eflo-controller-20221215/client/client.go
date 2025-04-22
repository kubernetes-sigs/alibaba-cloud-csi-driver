// This file is auto-generated, don't edit it. Thanks.
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/darabonba-openapi/v2/utils"
	"github.com/alibabacloud-go/tea/dara"
)

type iApproveOperationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *ApproveOperationRequest
	GetNodeId() *string
	SetOperationType(v string) *ApproveOperationRequest
	GetOperationType() *string
}

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
	return dara.Prettify(s)
}

func (s ApproveOperationRequest) GoString() string {
	return s.String()
}

func (s *ApproveOperationRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ApproveOperationRequest) GetOperationType() *string {
	return s.OperationType
}

func (s *ApproveOperationRequest) SetNodeId(v string) *ApproveOperationRequest {
	s.NodeId = &v
	return s
}

func (s *ApproveOperationRequest) SetOperationType(v string) *ApproveOperationRequest {
	s.OperationType = &v
	return s
}

func (s *ApproveOperationRequest) Validate() error {
	return dara.Validate(s)
}

type iApproveOperationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *ApproveOperationResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ApproveOperationResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ApproveOperationResponseBody) GoString() string {
	return s.String()
}

func (s *ApproveOperationResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ApproveOperationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ApproveOperationResponseBody) SetErrorMessage(v string) *ApproveOperationResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ApproveOperationResponseBody) SetRequestId(v string) *ApproveOperationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ApproveOperationResponseBody) Validate() error {
	return dara.Validate(s)
}

type iApproveOperationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ApproveOperationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ApproveOperationResponse
	GetStatusCode() *int32
	SetBody(v *ApproveOperationResponseBody) *ApproveOperationResponse
	GetBody() *ApproveOperationResponseBody
}

type ApproveOperationResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ApproveOperationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ApproveOperationResponse) String() string {
	return dara.Prettify(s)
}

func (s ApproveOperationResponse) GoString() string {
	return s.String()
}

func (s *ApproveOperationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ApproveOperationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ApproveOperationResponse) GetBody() *ApproveOperationResponseBody {
	return s.Body
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

func (s *ApproveOperationResponse) Validate() error {
	return dara.Validate(s)
}

type iChangeResourceGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetResourceGroupId(v string) *ChangeResourceGroupRequest
	GetResourceGroupId() *string
	SetResourceId(v string) *ChangeResourceGroupRequest
	GetResourceId() *string
	SetResourceRegionId(v string) *ChangeResourceGroupRequest
	GetResourceRegionId() *string
	SetResourceType(v string) *ChangeResourceGroupRequest
	GetResourceType() *string
	SetService(v string) *ChangeResourceGroupRequest
	GetService() *string
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
	Service      *string `json:"Service,omitempty" xml:"Service,omitempty"`
}

func (s ChangeResourceGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupRequest) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ChangeResourceGroupRequest) GetResourceId() *string {
	return s.ResourceId
}

func (s *ChangeResourceGroupRequest) GetResourceRegionId() *string {
	return s.ResourceRegionId
}

func (s *ChangeResourceGroupRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ChangeResourceGroupRequest) GetService() *string {
	return s.Service
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

func (s *ChangeResourceGroupRequest) SetService(v string) *ChangeResourceGroupRequest {
	s.Service = &v
	return s
}

func (s *ChangeResourceGroupRequest) Validate() error {
	return dara.Validate(s)
}

type iChangeResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ChangeResourceGroupResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ChangeResourceGroupResponseBody) SetRequestId(v string) *ChangeResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ChangeResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type iChangeResourceGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ChangeResourceGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ChangeResourceGroupResponse
	GetStatusCode() *int32
	SetBody(v *ChangeResourceGroupResponseBody) *ChangeResourceGroupResponse
	GetBody() *ChangeResourceGroupResponseBody
}

type ChangeResourceGroupResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ChangeResourceGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ChangeResourceGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s ChangeResourceGroupResponse) GoString() string {
	return s.String()
}

func (s *ChangeResourceGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ChangeResourceGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ChangeResourceGroupResponse) GetBody() *ChangeResourceGroupResponseBody {
	return s.Body
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

func (s *ChangeResourceGroupResponse) Validate() error {
	return dara.Validate(s)
}

type iCloseSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSessionId(v string) *CloseSessionRequest
	GetSessionId() *string
	SetSessionToken(v string) *CloseSessionRequest
	GetSessionToken() *string
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
	return dara.Prettify(s)
}

func (s CloseSessionRequest) GoString() string {
	return s.String()
}

func (s *CloseSessionRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *CloseSessionRequest) GetSessionToken() *string {
	return s.SessionToken
}

func (s *CloseSessionRequest) SetSessionId(v string) *CloseSessionRequest {
	s.SessionId = &v
	return s
}

func (s *CloseSessionRequest) SetSessionToken(v string) *CloseSessionRequest {
	s.SessionToken = &v
	return s
}

func (s *CloseSessionRequest) Validate() error {
	return dara.Validate(s)
}

type iCloseSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseSessionResponseBody
	GetRequestId() *string
	SetSessionId(v string) *CloseSessionResponseBody
	GetSessionId() *string
	SetState(v string) *CloseSessionResponseBody
	GetState() *string
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
	return dara.Prettify(s)
}

func (s CloseSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CloseSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *CloseSessionResponseBody) GetState() *string {
	return s.State
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

func (s *CloseSessionResponseBody) Validate() error {
	return dara.Validate(s)
}

type iCloseSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CloseSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CloseSessionResponse
	GetStatusCode() *int32
	SetBody(v *CloseSessionResponseBody) *CloseSessionResponse
	GetBody() *CloseSessionResponseBody
}

type CloseSessionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CloseSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CloseSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CloseSessionResponse) GoString() string {
	return s.String()
}

func (s *CloseSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CloseSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CloseSessionResponse) GetBody() *CloseSessionResponseBody {
	return s.Body
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

func (s *CloseSessionResponse) Validate() error {
	return dara.Validate(s)
}

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
	return dara.Validate(s)
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
	return dara.Validate(s)
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
	return dara.Validate(s)
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
	return dara.Validate(s)
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
	return dara.Validate(s)
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
	return dara.Validate(s)
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
	return dara.Validate(s)
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
	return dara.Validate(s)
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
	Disks *CreateClusterRequestNodeGroupsDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	// System image ID
	//
	// example:
	//
	// i190297201634099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
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
	return dara.Prettify(s)
}

func (s CreateClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroups) GetDisks() *CreateClusterRequestNodeGroupsDisks {
	return s.Disks
}

func (s *CreateClusterRequestNodeGroups) GetImageId() *string {
	return s.ImageId
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

func (s *CreateClusterRequestNodeGroups) GetUserData() *string {
	return s.UserData
}

func (s *CreateClusterRequestNodeGroups) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateClusterRequestNodeGroups) SetDisks(v *CreateClusterRequestNodeGroupsDisks) *CreateClusterRequestNodeGroups {
	s.Disks = v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetImageId(v string) *CreateClusterRequestNodeGroups {
	s.ImageId = &v
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

func (s *CreateClusterRequestNodeGroups) SetUserData(v string) *CreateClusterRequestNodeGroups {
	s.UserData = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) SetZoneId(v string) *CreateClusterRequestNodeGroups {
	s.ZoneId = &v
	return s
}

func (s *CreateClusterRequestNodeGroups) Validate() error {
	return dara.Validate(s)
}

type CreateClusterRequestNodeGroupsDisks struct {
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	SystemDiskSize             *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s CreateClusterRequestNodeGroupsDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsDisks) GoString() string {
	return s.String()
}

func (s *CreateClusterRequestNodeGroupsDisks) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *CreateClusterRequestNodeGroupsDisks) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateClusterRequestNodeGroupsDisks) SetSystemDiskPerformanceLevel(v string) *CreateClusterRequestNodeGroupsDisks {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsDisks) SetSystemDiskSize(v int32) *CreateClusterRequestNodeGroupsDisks {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateClusterRequestNodeGroupsDisks) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
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

type iCreateClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterDescription(v string) *CreateClusterShrinkRequest
	GetClusterDescription() *string
	SetClusterName(v string) *CreateClusterShrinkRequest
	GetClusterName() *string
	SetClusterType(v string) *CreateClusterShrinkRequest
	GetClusterType() *string
	SetComponentsShrink(v string) *CreateClusterShrinkRequest
	GetComponentsShrink() *string
	SetHpnZone(v string) *CreateClusterShrinkRequest
	GetHpnZone() *string
	SetIgnoreFailedNodeTasks(v bool) *CreateClusterShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNetworksShrink(v string) *CreateClusterShrinkRequest
	GetNetworksShrink() *string
	SetNimizVSwitchesShrink(v string) *CreateClusterShrinkRequest
	GetNimizVSwitchesShrink() *string
	SetNodeGroupsShrink(v string) *CreateClusterShrinkRequest
	GetNodeGroupsShrink() *string
	SetOpenEniJumboFrame(v bool) *CreateClusterShrinkRequest
	GetOpenEniJumboFrame() *bool
	SetResourceGroupId(v string) *CreateClusterShrinkRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateClusterShrinkRequestTag) *CreateClusterShrinkRequest
	GetTag() []*CreateClusterShrinkRequestTag
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
	return dara.Prettify(s)
}

func (s CreateClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequest) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *CreateClusterShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateClusterShrinkRequest) GetClusterType() *string {
	return s.ClusterType
}

func (s *CreateClusterShrinkRequest) GetComponentsShrink() *string {
	return s.ComponentsShrink
}

func (s *CreateClusterShrinkRequest) GetHpnZone() *string {
	return s.HpnZone
}

func (s *CreateClusterShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *CreateClusterShrinkRequest) GetNetworksShrink() *string {
	return s.NetworksShrink
}

func (s *CreateClusterShrinkRequest) GetNimizVSwitchesShrink() *string {
	return s.NimizVSwitchesShrink
}

func (s *CreateClusterShrinkRequest) GetNodeGroupsShrink() *string {
	return s.NodeGroupsShrink
}

func (s *CreateClusterShrinkRequest) GetOpenEniJumboFrame() *bool {
	return s.OpenEniJumboFrame
}

func (s *CreateClusterShrinkRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateClusterShrinkRequest) GetTag() []*CreateClusterShrinkRequestTag {
	return s.Tag
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

func (s *CreateClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateClusterShrinkRequestTag) GoString() string {
	return s.String()
}

func (s *CreateClusterShrinkRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateClusterShrinkRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateClusterShrinkRequestTag) SetKey(v string) *CreateClusterShrinkRequestTag {
	s.Key = &v
	return s
}

func (s *CreateClusterShrinkRequestTag) SetValue(v string) *CreateClusterShrinkRequestTag {
	s.Value = &v
	return s
}

func (s *CreateClusterShrinkRequestTag) Validate() error {
	return dara.Validate(s)
}

type iCreateClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateClusterResponseBody
	GetClusterId() *string
	SetRequestId(v string) *CreateClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateClusterResponseBody
	GetTaskId() *string
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
	return dara.Prettify(s)
}

func (s CreateClusterResponseBody) GoString() string {
	return s.String()
}

func (s *CreateClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateClusterResponseBody) GetTaskId() *string {
	return s.TaskId
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

func (s *CreateClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type iCreateClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateClusterResponse
	GetStatusCode() *int32
	SetBody(v *CreateClusterResponseBody) *CreateClusterResponse
	GetBody() *CreateClusterResponseBody
}

type CreateClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateClusterResponse) GoString() string {
	return s.String()
}

func (s *CreateClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateClusterResponse) GetBody() *CreateClusterResponseBody {
	return s.Body
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

func (s *CreateClusterResponse) Validate() error {
	return dara.Validate(s)
}

type iCreateDiagnosticTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiJobLogInfo(v *CreateDiagnosticTaskRequestAiJobLogInfo) *CreateDiagnosticTaskRequest
	GetAiJobLogInfo() *CreateDiagnosticTaskRequestAiJobLogInfo
	SetClusterId(v string) *CreateDiagnosticTaskRequest
	GetClusterId() *string
	SetDiagnosticType(v string) *CreateDiagnosticTaskRequest
	GetDiagnosticType() *string
	SetNodeIds(v []*string) *CreateDiagnosticTaskRequest
	GetNodeIds() []*string
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
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequest) GetAiJobLogInfo() *CreateDiagnosticTaskRequestAiJobLogInfo {
	return s.AiJobLogInfo
}

func (s *CreateDiagnosticTaskRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateDiagnosticTaskRequest) GetDiagnosticType() *string {
	return s.DiagnosticType
}

func (s *CreateDiagnosticTaskRequest) GetNodeIds() []*string {
	return s.NodeIds
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

func (s *CreateDiagnosticTaskRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfo) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) GetAiJobLogs() []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs {
	return s.AiJobLogs
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) GetEndTime() *string {
	return s.EndTime
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) GetStartTime() *string {
	return s.StartTime
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

func (s *CreateDiagnosticTaskRequestAiJobLogInfo) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) GetAiInstance() *string {
	return s.AiInstance
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) GetLogs() []*CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs {
	return s.Logs
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) GetNodeId() *string {
	return s.NodeId
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

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogs) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) GetDatetime() *string {
	return s.Datetime
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) GetLogContent() *string {
	return s.LogContent
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) SetDatetime(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs {
	s.Datetime = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) SetLogContent(v string) *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs {
	s.LogContent = &v
	return s
}

func (s *CreateDiagnosticTaskRequestAiJobLogInfoAiJobLogsLogs) Validate() error {
	return dara.Validate(s)
}

type iCreateDiagnosticTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAiJobLogInfoShrink(v string) *CreateDiagnosticTaskShrinkRequest
	GetAiJobLogInfoShrink() *string
	SetClusterId(v string) *CreateDiagnosticTaskShrinkRequest
	GetClusterId() *string
	SetDiagnosticType(v string) *CreateDiagnosticTaskShrinkRequest
	GetDiagnosticType() *string
	SetNodeIdsShrink(v string) *CreateDiagnosticTaskShrinkRequest
	GetNodeIdsShrink() *string
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
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskShrinkRequest) GetAiJobLogInfoShrink() *string {
	return s.AiJobLogInfoShrink
}

func (s *CreateDiagnosticTaskShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateDiagnosticTaskShrinkRequest) GetDiagnosticType() *string {
	return s.DiagnosticType
}

func (s *CreateDiagnosticTaskShrinkRequest) GetNodeIdsShrink() *string {
	return s.NodeIdsShrink
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

func (s *CreateDiagnosticTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iCreateDiagnosticTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosticId(v string) *CreateDiagnosticTaskResponseBody
	GetDiagnosticId() *string
	SetRequestId(v string) *CreateDiagnosticTaskResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskResponseBody) GetDiagnosticId() *string {
	return s.DiagnosticId
}

func (s *CreateDiagnosticTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDiagnosticTaskResponseBody) SetDiagnosticId(v string) *CreateDiagnosticTaskResponseBody {
	s.DiagnosticId = &v
	return s
}

func (s *CreateDiagnosticTaskResponseBody) SetRequestId(v string) *CreateDiagnosticTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDiagnosticTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type iCreateDiagnosticTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDiagnosticTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDiagnosticTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateDiagnosticTaskResponseBody) *CreateDiagnosticTaskResponse
	GetBody() *CreateDiagnosticTaskResponseBody
}

type CreateDiagnosticTaskResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDiagnosticTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDiagnosticTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDiagnosticTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateDiagnosticTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDiagnosticTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDiagnosticTaskResponse) GetBody() *CreateDiagnosticTaskResponseBody {
	return s.Body
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

func (s *CreateDiagnosticTaskResponse) Validate() error {
	return dara.Validate(s)
}

type iCreateNetTestTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNetTestTaskRequest
	GetClusterId() *string
	SetClusterName(v string) *CreateNetTestTaskRequest
	GetClusterName() *string
	SetCommTest(v *CreateNetTestTaskRequestCommTest) *CreateNetTestTaskRequest
	GetCommTest() *CreateNetTestTaskRequestCommTest
	SetDelayTest(v *CreateNetTestTaskRequestDelayTest) *CreateNetTestTaskRequest
	GetDelayTest() *CreateNetTestTaskRequestDelayTest
	SetNetTestType(v string) *CreateNetTestTaskRequest
	GetNetTestType() *string
	SetNetworkMode(v string) *CreateNetTestTaskRequest
	GetNetworkMode() *string
	SetPort(v string) *CreateNetTestTaskRequest
	GetPort() *string
	SetTrafficTest(v *CreateNetTestTaskRequestTrafficTest) *CreateNetTestTaskRequest
	GetTrafficTest() *CreateNetTestTaskRequestTrafficTest
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
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNetTestTaskRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateNetTestTaskRequest) GetCommTest() *CreateNetTestTaskRequestCommTest {
	return s.CommTest
}

func (s *CreateNetTestTaskRequest) GetDelayTest() *CreateNetTestTaskRequestDelayTest {
	return s.DelayTest
}

func (s *CreateNetTestTaskRequest) GetNetTestType() *string {
	return s.NetTestType
}

func (s *CreateNetTestTaskRequest) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *CreateNetTestTaskRequest) GetPort() *string {
	return s.Port
}

func (s *CreateNetTestTaskRequest) GetTrafficTest() *CreateNetTestTaskRequestTrafficTest {
	return s.TrafficTest
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

func (s *CreateNetTestTaskRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestCommTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestCommTest) GetGPUNum() *int64 {
	return s.GPUNum
}

func (s *CreateNetTestTaskRequestCommTest) GetHosts() []*CreateNetTestTaskRequestCommTestHosts {
	return s.Hosts
}

func (s *CreateNetTestTaskRequestCommTest) GetModel() *string {
	return s.Model
}

func (s *CreateNetTestTaskRequestCommTest) GetType() *string {
	return s.Type
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

func (s *CreateNetTestTaskRequestCommTest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestCommTestHosts) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestCommTestHosts) GetIP() *string {
	return s.IP
}

func (s *CreateNetTestTaskRequestCommTestHosts) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNetTestTaskRequestCommTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateNetTestTaskRequestCommTestHosts) GetServerName() *string {
	return s.ServerName
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

func (s *CreateNetTestTaskRequestCommTestHosts) Validate() error {
	return dara.Validate(s)
}

type CreateNetTestTaskRequestDelayTest struct {
	// 输入测试节点的hosts
	Hosts []*CreateNetTestTaskRequestDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s CreateNetTestTaskRequestDelayTest) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestDelayTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestDelayTest) GetHosts() []*CreateNetTestTaskRequestDelayTestHosts {
	return s.Hosts
}

func (s *CreateNetTestTaskRequestDelayTest) SetHosts(v []*CreateNetTestTaskRequestDelayTestHosts) *CreateNetTestTaskRequestDelayTest {
	s.Hosts = v
	return s
}

func (s *CreateNetTestTaskRequestDelayTest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestDelayTestHosts) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestDelayTestHosts) GetBond() *string {
	return s.Bond
}

func (s *CreateNetTestTaskRequestDelayTestHosts) GetIP() *string {
	return s.IP
}

func (s *CreateNetTestTaskRequestDelayTestHosts) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNetTestTaskRequestDelayTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateNetTestTaskRequestDelayTestHosts) GetServerName() *string {
	return s.ServerName
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

func (s *CreateNetTestTaskRequestDelayTestHosts) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTest) GetClients() []*CreateNetTestTaskRequestTrafficTestClients {
	return s.Clients
}

func (s *CreateNetTestTaskRequestTrafficTest) GetDuration() *int64 {
	return s.Duration
}

func (s *CreateNetTestTaskRequestTrafficTest) GetGDR() *bool {
	return s.GDR
}

func (s *CreateNetTestTaskRequestTrafficTest) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateNetTestTaskRequestTrafficTest) GetQP() *int64 {
	return s.QP
}

func (s *CreateNetTestTaskRequestTrafficTest) GetServers() []*CreateNetTestTaskRequestTrafficTestServers {
	return s.Servers
}

func (s *CreateNetTestTaskRequestTrafficTest) GetTrafficModel() *string {
	return s.TrafficModel
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

func (s *CreateNetTestTaskRequestTrafficTest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTestClients) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTestClients) GetBond() *string {
	return s.Bond
}

func (s *CreateNetTestTaskRequestTrafficTestClients) GetIP() *string {
	return s.IP
}

func (s *CreateNetTestTaskRequestTrafficTestClients) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNetTestTaskRequestTrafficTestClients) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateNetTestTaskRequestTrafficTestClients) GetServerName() *string {
	return s.ServerName
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

func (s *CreateNetTestTaskRequestTrafficTestClients) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s CreateNetTestTaskRequestTrafficTestServers) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskRequestTrafficTestServers) GetBond() *string {
	return s.Bond
}

func (s *CreateNetTestTaskRequestTrafficTestServers) GetIP() *string {
	return s.IP
}

func (s *CreateNetTestTaskRequestTrafficTestServers) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateNetTestTaskRequestTrafficTestServers) GetResourceId() *string {
	return s.ResourceId
}

func (s *CreateNetTestTaskRequestTrafficTestServers) GetServerName() *string {
	return s.ServerName
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

func (s *CreateNetTestTaskRequestTrafficTestServers) Validate() error {
	return dara.Validate(s)
}

type iCreateNetTestTaskShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNetTestTaskShrinkRequest
	GetClusterId() *string
	SetClusterName(v string) *CreateNetTestTaskShrinkRequest
	GetClusterName() *string
	SetCommTestShrink(v string) *CreateNetTestTaskShrinkRequest
	GetCommTestShrink() *string
	SetDelayTestShrink(v string) *CreateNetTestTaskShrinkRequest
	GetDelayTestShrink() *string
	SetNetTestType(v string) *CreateNetTestTaskShrinkRequest
	GetNetTestType() *string
	SetNetworkMode(v string) *CreateNetTestTaskShrinkRequest
	GetNetworkMode() *string
	SetPort(v string) *CreateNetTestTaskShrinkRequest
	GetPort() *string
	SetTrafficTestShrink(v string) *CreateNetTestTaskShrinkRequest
	GetTrafficTestShrink() *string
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
	return dara.Prettify(s)
}

func (s CreateNetTestTaskShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNetTestTaskShrinkRequest) GetClusterName() *string {
	return s.ClusterName
}

func (s *CreateNetTestTaskShrinkRequest) GetCommTestShrink() *string {
	return s.CommTestShrink
}

func (s *CreateNetTestTaskShrinkRequest) GetDelayTestShrink() *string {
	return s.DelayTestShrink
}

func (s *CreateNetTestTaskShrinkRequest) GetNetTestType() *string {
	return s.NetTestType
}

func (s *CreateNetTestTaskShrinkRequest) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *CreateNetTestTaskShrinkRequest) GetPort() *string {
	return s.Port
}

func (s *CreateNetTestTaskShrinkRequest) GetTrafficTestShrink() *string {
	return s.TrafficTestShrink
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

func (s *CreateNetTestTaskShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iCreateNetTestTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateNetTestTaskResponseBody
	GetRequestId() *string
	SetTestId(v string) *CreateNetTestTaskResponseBody
	GetTestId() *string
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
	return dara.Prettify(s)
}

func (s CreateNetTestTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNetTestTaskResponseBody) GetTestId() *string {
	return s.TestId
}

func (s *CreateNetTestTaskResponseBody) SetRequestId(v string) *CreateNetTestTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNetTestTaskResponseBody) SetTestId(v string) *CreateNetTestTaskResponseBody {
	s.TestId = &v
	return s
}

func (s *CreateNetTestTaskResponseBody) Validate() error {
	return dara.Validate(s)
}

type iCreateNetTestTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNetTestTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNetTestTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateNetTestTaskResponseBody) *CreateNetTestTaskResponse
	GetBody() *CreateNetTestTaskResponseBody
}

type CreateNetTestTaskResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNetTestTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNetTestTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNetTestTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateNetTestTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNetTestTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNetTestTaskResponse) GetBody() *CreateNetTestTaskResponseBody {
	return s.Body
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

func (s *CreateNetTestTaskResponse) Validate() error {
	return dara.Validate(s)
}

type iCreateNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNodeGroupRequest
	GetClusterId() *string
	SetNodeGroup(v *CreateNodeGroupRequestNodeGroup) *CreateNodeGroupRequest
	GetNodeGroup() *CreateNodeGroupRequestNodeGroup
	SetNodeUnit(v map[string]interface{}) *CreateNodeGroupRequest
	GetNodeUnit() map[string]interface{}
}

type CreateNodeGroupRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i118191731740041623425
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	NodeGroup *CreateNodeGroupRequestNodeGroup `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty" type:"Struct"`
	NodeUnit  map[string]interface{}           `json:"NodeUnit,omitempty" xml:"NodeUnit,omitempty"`
}

func (s CreateNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNodeGroupRequest) GetNodeGroup() *CreateNodeGroupRequestNodeGroup {
	return s.NodeGroup
}

func (s *CreateNodeGroupRequest) GetNodeUnit() map[string]interface{} {
	return s.NodeUnit
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

func (s *CreateNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}

type CreateNodeGroupRequestNodeGroup struct {
	// This parameter is required.
	//
	// example:
	//
	// cn-wulanchabu-b
	Az    *string                               `json:"Az,omitempty" xml:"Az,omitempty"`
	Disks *CreateNodeGroupRequestNodeGroupDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// i191887641687336652616
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// mock-machine-type3
	MachineType          *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	NodeGroupDescription *string `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// PAI-LINGJUN
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	UserData      *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s CreateNodeGroupRequestNodeGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupRequestNodeGroup) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequestNodeGroup) GetAz() *string {
	return s.Az
}

func (s *CreateNodeGroupRequestNodeGroup) GetDisks() *CreateNodeGroupRequestNodeGroupDisks {
	return s.Disks
}

func (s *CreateNodeGroupRequestNodeGroup) GetImageId() *string {
	return s.ImageId
}

func (s *CreateNodeGroupRequestNodeGroup) GetMachineType() *string {
	return s.MachineType
}

func (s *CreateNodeGroupRequestNodeGroup) GetNodeGroupDescription() *string {
	return s.NodeGroupDescription
}

func (s *CreateNodeGroupRequestNodeGroup) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *CreateNodeGroupRequestNodeGroup) GetUserData() *string {
	return s.UserData
}

func (s *CreateNodeGroupRequestNodeGroup) SetAz(v string) *CreateNodeGroupRequestNodeGroup {
	s.Az = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetDisks(v *CreateNodeGroupRequestNodeGroupDisks) *CreateNodeGroupRequestNodeGroup {
	s.Disks = v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) SetImageId(v string) *CreateNodeGroupRequestNodeGroup {
	s.ImageId = &v
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

func (s *CreateNodeGroupRequestNodeGroup) SetUserData(v string) *CreateNodeGroupRequestNodeGroup {
	s.UserData = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroup) Validate() error {
	return dara.Validate(s)
}

type CreateNodeGroupRequestNodeGroupDisks struct {
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	SystemDiskSize             *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s CreateNodeGroupRequestNodeGroupDisks) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupRequestNodeGroupDisks) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupRequestNodeGroupDisks) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *CreateNodeGroupRequestNodeGroupDisks) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *CreateNodeGroupRequestNodeGroupDisks) SetSystemDiskPerformanceLevel(v string) *CreateNodeGroupRequestNodeGroupDisks {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroupDisks) SetSystemDiskSize(v int32) *CreateNodeGroupRequestNodeGroupDisks {
	s.SystemDiskSize = &v
	return s
}

func (s *CreateNodeGroupRequestNodeGroupDisks) Validate() error {
	return dara.Validate(s)
}

type iCreateNodeGroupShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *CreateNodeGroupShrinkRequest
	GetClusterId() *string
	SetNodeGroupShrink(v string) *CreateNodeGroupShrinkRequest
	GetNodeGroupShrink() *string
	SetNodeUnitShrink(v string) *CreateNodeGroupShrinkRequest
	GetNodeUnitShrink() *string
}

type CreateNodeGroupShrinkRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// i118191731740041623425
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// This parameter is required.
	NodeGroupShrink *string `json:"NodeGroup,omitempty" xml:"NodeGroup,omitempty"`
	NodeUnitShrink  *string `json:"NodeUnit,omitempty" xml:"NodeUnit,omitempty"`
}

func (s CreateNodeGroupShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateNodeGroupShrinkRequest) GetNodeGroupShrink() *string {
	return s.NodeGroupShrink
}

func (s *CreateNodeGroupShrinkRequest) GetNodeUnitShrink() *string {
	return s.NodeUnitShrink
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

func (s *CreateNodeGroupShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iCreateNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *CreateNodeGroupResponseBody
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *CreateNodeGroupResponseBody
	GetNodeGroupName() *string
	SetRequestId(v string) *CreateNodeGroupResponseBody
	GetRequestId() *string
}

type CreateNodeGroupResponseBody struct {
	// example:
	//
	// ng-ec3c96ff0aa4c60d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// example:
	//
	// emr-default
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponseBody) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *CreateNodeGroupResponseBody) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *CreateNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *CreateNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type iCreateNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *CreateNodeGroupResponseBody) *CreateNodeGroupResponse
	GetBody() *CreateNodeGroupResponseBody
}

type CreateNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateNodeGroupResponse) GetBody() *CreateNodeGroupResponseBody {
	return s.Body
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

func (s *CreateNodeGroupResponse) Validate() error {
	return dara.Validate(s)
}

type iCreateSessionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *CreateSessionRequest
	GetNodeId() *string
	SetSessionType(v string) *CreateSessionRequest
	GetSessionType() *string
	SetStartTime(v string) *CreateSessionRequest
	GetStartTime() *string
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
	return dara.Prettify(s)
}

func (s CreateSessionRequest) GoString() string {
	return s.String()
}

func (s *CreateSessionRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateSessionRequest) GetSessionType() *string {
	return s.SessionType
}

func (s *CreateSessionRequest) GetStartTime() *string {
	return s.StartTime
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

func (s *CreateSessionRequest) Validate() error {
	return dara.Validate(s)
}

type iCreateSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateSessionResponseBody
	GetRequestId() *string
	SetServerSn(v string) *CreateSessionResponseBody
	GetServerSn() *string
	SetSessionId(v string) *CreateSessionResponseBody
	GetSessionId() *string
	SetSessionToken(v string) *CreateSessionResponseBody
	GetSessionToken() *string
	SetWssEndpoint(v string) *CreateSessionResponseBody
	GetWssEndpoint() *string
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
	return dara.Prettify(s)
}

func (s CreateSessionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSessionResponseBody) GetServerSn() *string {
	return s.ServerSn
}

func (s *CreateSessionResponseBody) GetSessionId() *string {
	return s.SessionId
}

func (s *CreateSessionResponseBody) GetSessionToken() *string {
	return s.SessionToken
}

func (s *CreateSessionResponseBody) GetWssEndpoint() *string {
	return s.WssEndpoint
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

func (s *CreateSessionResponseBody) Validate() error {
	return dara.Validate(s)
}

type iCreateSessionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateSessionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateSessionResponse
	GetStatusCode() *int32
	SetBody(v *CreateSessionResponseBody) *CreateSessionResponse
	GetBody() *CreateSessionResponseBody
}

type CreateSessionResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSessionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSessionResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateSessionResponse) GoString() string {
	return s.String()
}

func (s *CreateSessionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateSessionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateSessionResponse) GetBody() *CreateSessionResponseBody {
	return s.Body
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

func (s *CreateSessionResponse) Validate() error {
	return dara.Validate(s)
}

type iCreateVscRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateVscRequest
	GetClientToken() *string
	SetNodeId(v string) *CreateVscRequest
	GetNodeId() *string
	SetResourceGroupId(v string) *CreateVscRequest
	GetResourceGroupId() *string
	SetTag(v []*CreateVscRequestTag) *CreateVscRequest
	GetTag() []*CreateVscRequestTag
	SetVscName(v string) *CreateVscRequest
	GetVscName() *string
	SetVscType(v string) *CreateVscRequest
	GetVscType() *string
}

type CreateVscRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	NodeId          *string                `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ResourceGroupId *string                `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*CreateVscRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VscName         *string                `json:"VscName,omitempty" xml:"VscName,omitempty"`
	VscType         *string                `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s CreateVscRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateVscRequest) GoString() string {
	return s.String()
}

func (s *CreateVscRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateVscRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *CreateVscRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateVscRequest) GetTag() []*CreateVscRequestTag {
	return s.Tag
}

func (s *CreateVscRequest) GetVscName() *string {
	return s.VscName
}

func (s *CreateVscRequest) GetVscType() *string {
	return s.VscType
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

func (s *CreateVscRequest) Validate() error {
	return dara.Validate(s)
}

type CreateVscRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateVscRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateVscRequestTag) GoString() string {
	return s.String()
}

func (s *CreateVscRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateVscRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateVscRequestTag) SetKey(v string) *CreateVscRequestTag {
	s.Key = &v
	return s
}

func (s *CreateVscRequestTag) SetValue(v string) *CreateVscRequestTag {
	s.Value = &v
	return s
}

func (s *CreateVscRequestTag) Validate() error {
	return dara.Validate(s)
}

type iCreateVscResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateVscResponseBody
	GetRequestId() *string
	SetVscId(v string) *CreateVscResponseBody
	GetVscId() *string
}

type CreateVscResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VscId     *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s CreateVscResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateVscResponseBody) GoString() string {
	return s.String()
}

func (s *CreateVscResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateVscResponseBody) GetVscId() *string {
	return s.VscId
}

func (s *CreateVscResponseBody) SetRequestId(v string) *CreateVscResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateVscResponseBody) SetVscId(v string) *CreateVscResponseBody {
	s.VscId = &v
	return s
}

func (s *CreateVscResponseBody) Validate() error {
	return dara.Validate(s)
}

type iCreateVscResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateVscResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateVscResponse
	GetStatusCode() *int32
	SetBody(v *CreateVscResponseBody) *CreateVscResponse
	GetBody() *CreateVscResponseBody
}

type CreateVscResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateVscResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateVscResponse) GoString() string {
	return s.String()
}

func (s *CreateVscResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateVscResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateVscResponse) GetBody() *CreateVscResponseBody {
	return s.Body
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

func (s *CreateVscResponse) Validate() error {
	return dara.Validate(s)
}

type iDeleteClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteClusterRequest
	GetClusterId() *string
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
	return dara.Prettify(s)
}

func (s DeleteClusterRequest) GoString() string {
	return s.String()
}

func (s *DeleteClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteClusterRequest) SetClusterId(v string) *DeleteClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteClusterRequest) Validate() error {
	return dara.Validate(s)
}

type iDeleteClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteClusterResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s DeleteClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteClusterResponseBody) SetRequestId(v string) *DeleteClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type iDeleteClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteClusterResponse
	GetStatusCode() *int32
	SetBody(v *DeleteClusterResponseBody) *DeleteClusterResponse
	GetBody() *DeleteClusterResponseBody
}

type DeleteClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteClusterResponse) GoString() string {
	return s.String()
}

func (s *DeleteClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteClusterResponse) GetBody() *DeleteClusterResponseBody {
	return s.Body
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

func (s *DeleteClusterResponse) Validate() error {
	return dara.Validate(s)
}

type iDeleteNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DeleteNodeGroupRequest
	GetClusterId() *string
	SetNodeGroupId(v string) *DeleteNodeGroupRequest
	GetNodeGroupId() *string
}

type DeleteNodeGroupRequest struct {
	// example:
	//
	// i114444141733395242745
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// i121824791737080429819
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s DeleteNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteNodeGroupRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DeleteNodeGroupRequest) SetClusterId(v string) *DeleteNodeGroupRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteNodeGroupRequest) SetNodeGroupId(v string) *DeleteNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *DeleteNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}

type iDeleteNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNodeGroupResponseBody
	GetRequestId() *string
}

type DeleteNodeGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNodeGroupResponseBody) SetRequestId(v string) *DeleteNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type iDeleteNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *DeleteNodeGroupResponseBody) *DeleteNodeGroupResponse
	GetBody() *DeleteNodeGroupResponseBody
}

type DeleteNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteNodeGroupResponse) GetBody() *DeleteNodeGroupResponseBody {
	return s.Body
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

func (s *DeleteNodeGroupResponse) Validate() error {
	return dara.Validate(s)
}

type iDeleteVscRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *DeleteVscRequest
	GetClientToken() *string
	SetVscId(v string) *DeleteVscRequest
	GetVscId() *string
}

type DeleteVscRequest struct {
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// This parameter is required.
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DeleteVscRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscRequest) GoString() string {
	return s.String()
}

func (s *DeleteVscRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *DeleteVscRequest) GetVscId() *string {
	return s.VscId
}

func (s *DeleteVscRequest) SetClientToken(v string) *DeleteVscRequest {
	s.ClientToken = &v
	return s
}

func (s *DeleteVscRequest) SetVscId(v string) *DeleteVscRequest {
	s.VscId = &v
	return s
}

func (s *DeleteVscRequest) Validate() error {
	return dara.Validate(s)
}

type iDeleteVscResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVscResponseBody
	GetRequestId() *string
}

type DeleteVscResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVscResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVscResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVscResponseBody) SetRequestId(v string) *DeleteVscResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVscResponseBody) Validate() error {
	return dara.Validate(s)
}

type iDeleteVscResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteVscResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteVscResponse
	GetStatusCode() *int32
	SetBody(v *DeleteVscResponseBody) *DeleteVscResponse
	GetBody() *DeleteVscResponseBody
}

type DeleteVscResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteVscResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteVscResponse) GoString() string {
	return s.String()
}

func (s *DeleteVscResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteVscResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteVscResponse) GetBody() *DeleteVscResponseBody {
	return s.Body
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

func (s *DeleteVscResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeClusterRequest
	GetClusterId() *string
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
	return dara.Prettify(s)
}

func (s DescribeClusterRequest) GoString() string {
	return s.String()
}

func (s *DescribeClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterRequest) SetClusterId(v string) *DescribeClusterRequest {
	s.ClusterId = &v
	return s
}

func (s *DescribeClusterRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterDescription(v string) *DescribeClusterResponseBody
	GetClusterDescription() *string
	SetClusterId(v string) *DescribeClusterResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeClusterResponseBody
	GetClusterName() *string
	SetClusterType(v string) *DescribeClusterResponseBody
	GetClusterType() *string
	SetComponents(v []*DescribeClusterResponseBodyComponents) *DescribeClusterResponseBody
	GetComponents() []*DescribeClusterResponseBodyComponents
	SetComputingIpVersion(v string) *DescribeClusterResponseBody
	GetComputingIpVersion() *string
	SetCreateTime(v string) *DescribeClusterResponseBody
	GetCreateTime() *string
	SetHpnZone(v string) *DescribeClusterResponseBody
	GetHpnZone() *string
	SetNetworks(v []*DescribeClusterResponseBodyNetworks) *DescribeClusterResponseBody
	GetNetworks() []*DescribeClusterResponseBodyNetworks
	SetNodeCount(v int64) *DescribeClusterResponseBody
	GetNodeCount() *int64
	SetNodeGroupCount(v int64) *DescribeClusterResponseBody
	GetNodeGroupCount() *int64
	SetOpenEniJumboFrame(v string) *DescribeClusterResponseBody
	GetOpenEniJumboFrame() *string
	SetOperatingState(v string) *DescribeClusterResponseBody
	GetOperatingState() *string
	SetRequestId(v string) *DescribeClusterResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeClusterResponseBody
	GetResourceGroupId() *string
	SetTaskId(v string) *DescribeClusterResponseBody
	GetTaskId() *string
	SetUpdateTime(v string) *DescribeClusterResponseBody
	GetUpdateTime() *string
	SetVpcId(v string) *DescribeClusterResponseBody
	GetVpcId() *string
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
	Networks []*DescribeClusterResponseBodyNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
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
	return dara.Prettify(s)
}

func (s DescribeClusterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBody) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *DescribeClusterResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeClusterResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeClusterResponseBody) GetClusterType() *string {
	return s.ClusterType
}

func (s *DescribeClusterResponseBody) GetComponents() []*DescribeClusterResponseBodyComponents {
	return s.Components
}

func (s *DescribeClusterResponseBody) GetComputingIpVersion() *string {
	return s.ComputingIpVersion
}

func (s *DescribeClusterResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeClusterResponseBody) GetHpnZone() *string {
	return s.HpnZone
}

func (s *DescribeClusterResponseBody) GetNetworks() []*DescribeClusterResponseBodyNetworks {
	return s.Networks
}

func (s *DescribeClusterResponseBody) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *DescribeClusterResponseBody) GetNodeGroupCount() *int64 {
	return s.NodeGroupCount
}

func (s *DescribeClusterResponseBody) GetOpenEniJumboFrame() *string {
	return s.OpenEniJumboFrame
}

func (s *DescribeClusterResponseBody) GetOperatingState() *string {
	return s.OperatingState
}

func (s *DescribeClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeClusterResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeClusterResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeClusterResponseBody) GetVpcId() *string {
	return s.VpcId
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

func (s *DescribeClusterResponseBody) SetNetworks(v []*DescribeClusterResponseBodyNetworks) *DescribeClusterResponseBody {
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

func (s *DescribeClusterResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeClusterResponseBodyComponents) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyComponents) GetComponentId() *string {
	return s.ComponentId
}

func (s *DescribeClusterResponseBodyComponents) GetComponentType() *string {
	return s.ComponentType
}

func (s *DescribeClusterResponseBodyComponents) SetComponentId(v string) *DescribeClusterResponseBodyComponents {
	s.ComponentId = &v
	return s
}

func (s *DescribeClusterResponseBodyComponents) SetComponentType(v string) *DescribeClusterResponseBodyComponents {
	s.ComponentType = &v
	return s
}

func (s *DescribeClusterResponseBodyComponents) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterResponseBodyNetworks struct {
	// VPC Segment ID
	//
	// example:
	//
	// vpd-iqd7xunc
	VpdId *string `json:"VpdId,omitempty" xml:"VpdId,omitempty"`
}

func (s DescribeClusterResponseBodyNetworks) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResponseBodyNetworks) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponseBodyNetworks) GetVpdId() *string {
	return s.VpdId
}

func (s *DescribeClusterResponseBodyNetworks) SetVpdId(v string) *DescribeClusterResponseBodyNetworks {
	s.VpdId = &v
	return s
}

func (s *DescribeClusterResponseBodyNetworks) Validate() error {
	return dara.Validate(s)
}

type iDescribeClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeClusterResponse
	GetStatusCode() *int32
	SetBody(v *DescribeClusterResponseBody) *DescribeClusterResponse
	GetBody() *DescribeClusterResponseBody
}

type DescribeClusterResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeClusterResponse) GoString() string {
	return s.String()
}

func (s *DescribeClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeClusterResponse) GetBody() *DescribeClusterResponseBody {
	return s.Body
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

func (s *DescribeClusterResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeDiagnosticResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosticId(v string) *DescribeDiagnosticResultRequest
	GetDiagnosticId() *string
}

type DescribeDiagnosticResultRequest struct {
	// example:
	//
	// diag-i151942361720577788844
	DiagnosticId *string `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
}

func (s DescribeDiagnosticResultRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultRequest) GetDiagnosticId() *string {
	return s.DiagnosticId
}

func (s *DescribeDiagnosticResultRequest) SetDiagnosticId(v string) *DescribeDiagnosticResultRequest {
	s.DiagnosticId = &v
	return s
}

func (s *DescribeDiagnosticResultRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeDiagnosticResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeDiagnosticResultResponseBody
	GetClusterId() *string
	SetCreatedTime(v string) *DescribeDiagnosticResultResponseBody
	GetCreatedTime() *string
	SetDiagnosticId(v string) *DescribeDiagnosticResultResponseBody
	GetDiagnosticId() *string
	SetDiagnosticResults(v []interface{}) *DescribeDiagnosticResultResponseBody
	GetDiagnosticResults() []interface{}
	SetDiagnosticState(v string) *DescribeDiagnosticResultResponseBody
	GetDiagnosticState() *string
	SetDiagnosticType(v string) *DescribeDiagnosticResultResponseBody
	GetDiagnosticType() *string
	SetEndTime(v string) *DescribeDiagnosticResultResponseBody
	GetEndTime() *string
	SetNodeIds(v []*string) *DescribeDiagnosticResultResponseBody
	GetNodeIds() []*string
	SetRequestId(v string) *DescribeDiagnosticResultResponseBody
	GetRequestId() *string
}

type DescribeDiagnosticResultResponseBody struct {
	// example:
	//
	// i118913031696573280136
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 2024-06-15T10:17:56
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// diag-i155363241720059671316
	DiagnosticId      *string       `json:"DiagnosticId,omitempty" xml:"DiagnosticId,omitempty"`
	DiagnosticResults []interface{} `json:"DiagnosticResults,omitempty" xml:"DiagnosticResults,omitempty" type:"Repeated"`
	// example:
	//
	// Fault
	DiagnosticState *string `json:"DiagnosticState,omitempty" xml:"DiagnosticState,omitempty"`
	// example:
	//
	// CheckByAiJobLogs
	DiagnosticType *string `json:"DiagnosticType,omitempty" xml:"DiagnosticType,omitempty"`
	// example:
	//
	// 2024-06-11T10:00:30
	EndTime *string   `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	NodeIds []*string `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDiagnosticResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeDiagnosticResultResponseBody) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *DescribeDiagnosticResultResponseBody) GetDiagnosticId() *string {
	return s.DiagnosticId
}

func (s *DescribeDiagnosticResultResponseBody) GetDiagnosticResults() []interface{} {
	return s.DiagnosticResults
}

func (s *DescribeDiagnosticResultResponseBody) GetDiagnosticState() *string {
	return s.DiagnosticState
}

func (s *DescribeDiagnosticResultResponseBody) GetDiagnosticType() *string {
	return s.DiagnosticType
}

func (s *DescribeDiagnosticResultResponseBody) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeDiagnosticResultResponseBody) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *DescribeDiagnosticResultResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *DescribeDiagnosticResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type iDescribeDiagnosticResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeDiagnosticResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeDiagnosticResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeDiagnosticResultResponseBody) *DescribeDiagnosticResultResponse
	GetBody() *DescribeDiagnosticResultResponseBody
}

type DescribeDiagnosticResultResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeDiagnosticResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeDiagnosticResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosticResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosticResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeDiagnosticResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeDiagnosticResultResponse) GetBody() *DescribeDiagnosticResultResponseBody {
	return s.Body
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

func (s *DescribeDiagnosticResultResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeInvocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContentEncoding(v string) *DescribeInvocationsRequest
	GetContentEncoding() *string
	SetIncludeOutput(v bool) *DescribeInvocationsRequest
	GetIncludeOutput() *bool
	SetInvokeId(v string) *DescribeInvocationsRequest
	GetInvokeId() *string
	SetNodeId(v string) *DescribeInvocationsRequest
	GetNodeId() *string
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
	return dara.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeInvocationsRequest) GetIncludeOutput() *bool {
	return s.IncludeOutput
}

func (s *DescribeInvocationsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationsRequest) GetNodeId() *string {
	return s.NodeId
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

func (s *DescribeInvocationsRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeInvocationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvocations(v *DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody
	GetInvocations() *DescribeInvocationsResponseBodyInvocations
	SetRequestId(v string) *DescribeInvocationsResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBody) GetInvocations() *DescribeInvocationsResponseBodyInvocations {
	return s.Invocations
}

func (s *DescribeInvocationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInvocationsResponseBody) SetInvocations(v *DescribeInvocationsResponseBodyInvocations) *DescribeInvocationsResponseBody {
	s.Invocations = v
	return s
}

func (s *DescribeInvocationsResponseBody) SetRequestId(v string) *DescribeInvocationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInvocationsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeInvocationsResponseBodyInvocations struct {
	// File delivery record.
	Invocation []*DescribeInvocationsResponseBodyInvocationsInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocations) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocations) GetInvocation() []*DescribeInvocationsResponseBodyInvocationsInvocation {
	return s.Invocation
}

func (s *DescribeInvocationsResponseBodyInvocations) SetInvocation(v []*DescribeInvocationsResponseBodyInvocationsInvocation) *DescribeInvocationsResponseBodyInvocations {
	s.Invocation = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocations) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocation) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCommandContent() *string {
	return s.CommandContent
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCommandDescription() *string {
	return s.CommandDescription
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCommandName() *string {
	return s.CommandName
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetFrequency() *string {
	return s.Frequency
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetInvokeNodes() *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes {
	return s.InvokeNodes
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetInvokeStatus() *string {
	return s.InvokeStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetParameters() *string {
	return s.Parameters
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetTimeout() *int32 {
	return s.Timeout
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetUsername() *string {
	return s.Username
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) GetWorkingDir() *string {
	return s.WorkingDir
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

func (s *DescribeInvocationsResponseBodyInvocationsInvocation) Validate() error {
	return dara.Validate(s)
}

type DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes struct {
	// Command execution records for nodes.
	InvokeNode []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode `json:"InvokeNode,omitempty" xml:"InvokeNode,omitempty" type:"Repeated"`
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) GetInvokeNode() []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	return s.InvokeNode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) SetInvokeNode(v []*DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes {
	s.InvokeNode = v
	return s
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodes) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetDropped() *int32 {
	return s.Dropped
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetExitCode() *int32 {
	return s.ExitCode
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetNodeInvokeStatus() *string {
	return s.NodeInvokeStatus
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetOutput() *string {
	return s.Output
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetRepeats() *int32 {
	return s.Repeats
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetStopTime() *string {
	return s.StopTime
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetTimed() *string {
	return s.Timed
}

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetUpdateTime() *string {
	return s.UpdateTime
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

func (s *DescribeInvocationsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) Validate() error {
	return dara.Validate(s)
}

type iDescribeInvocationsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeInvocationsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeInvocationsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeInvocationsResponseBody) *DescribeInvocationsResponse
	GetBody() *DescribeInvocationsResponseBody
}

type DescribeInvocationsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeInvocationsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeInvocationsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsResponse) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeInvocationsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeInvocationsResponse) GetBody() *DescribeInvocationsResponseBody {
	return s.Body
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

func (s *DescribeInvocationsResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeNetTestResultRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTestId(v string) *DescribeNetTestResultRequest
	GetTestId() *string
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
	return dara.Prettify(s)
}

func (s DescribeNetTestResultRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultRequest) GetTestId() *string {
	return s.TestId
}

func (s *DescribeNetTestResultRequest) SetTestId(v string) *DescribeNetTestResultRequest {
	s.TestId = &v
	return s
}

func (s *DescribeNetTestResultRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeNetTestResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeNetTestResultResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeNetTestResultResponseBody
	GetClusterName() *string
	SetCommTest(v *DescribeNetTestResultResponseBodyCommTest) *DescribeNetTestResultResponseBody
	GetCommTest() *DescribeNetTestResultResponseBodyCommTest
	SetCreationTime(v string) *DescribeNetTestResultResponseBody
	GetCreationTime() *string
	SetDelayTest(v *DescribeNetTestResultResponseBodyDelayTest) *DescribeNetTestResultResponseBody
	GetDelayTest() *DescribeNetTestResultResponseBodyDelayTest
	SetFinishedTime(v string) *DescribeNetTestResultResponseBody
	GetFinishedTime() *string
	SetNetTestType(v string) *DescribeNetTestResultResponseBody
	GetNetTestType() *string
	SetPort(v string) *DescribeNetTestResultResponseBody
	GetPort() *string
	SetRequestId(v string) *DescribeNetTestResultResponseBody
	GetRequestId() *string
	SetResultDetial(v string) *DescribeNetTestResultResponseBody
	GetResultDetial() *string
	SetStatus(v string) *DescribeNetTestResultResponseBody
	GetStatus() *string
	SetTestId(v string) *DescribeNetTestResultResponseBody
	GetTestId() *string
	SetTrafficTest(v *DescribeNetTestResultResponseBodyTrafficTest) *DescribeNetTestResultResponseBody
	GetTrafficTest() *DescribeNetTestResultResponseBodyTrafficTest
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
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeNetTestResultResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeNetTestResultResponseBody) GetCommTest() *DescribeNetTestResultResponseBodyCommTest {
	return s.CommTest
}

func (s *DescribeNetTestResultResponseBody) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeNetTestResultResponseBody) GetDelayTest() *DescribeNetTestResultResponseBodyDelayTest {
	return s.DelayTest
}

func (s *DescribeNetTestResultResponseBody) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *DescribeNetTestResultResponseBody) GetNetTestType() *string {
	return s.NetTestType
}

func (s *DescribeNetTestResultResponseBody) GetPort() *string {
	return s.Port
}

func (s *DescribeNetTestResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNetTestResultResponseBody) GetResultDetial() *string {
	return s.ResultDetial
}

func (s *DescribeNetTestResultResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeNetTestResultResponseBody) GetTestId() *string {
	return s.TestId
}

func (s *DescribeNetTestResultResponseBody) GetTrafficTest() *DescribeNetTestResultResponseBodyTrafficTest {
	return s.TrafficTest
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

func (s *DescribeNetTestResultResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyCommTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyCommTest) GetGPUNum() *string {
	return s.GPUNum
}

func (s *DescribeNetTestResultResponseBodyCommTest) GetHosts() []*DescribeNetTestResultResponseBodyCommTestHosts {
	return s.Hosts
}

func (s *DescribeNetTestResultResponseBodyCommTest) GetModel() *string {
	return s.Model
}

func (s *DescribeNetTestResultResponseBodyCommTest) GetType() *string {
	return s.Type
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

func (s *DescribeNetTestResultResponseBodyCommTest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyCommTestHosts) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) GetIP() *string {
	return s.IP
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetTestResultResponseBodyCommTestHosts) GetServerName() *string {
	return s.ServerName
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

func (s *DescribeNetTestResultResponseBodyCommTestHosts) Validate() error {
	return dara.Validate(s)
}

type DescribeNetTestResultResponseBodyDelayTest struct {
	// Input the hosts of the test nodes
	Hosts []*DescribeNetTestResultResponseBodyDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s DescribeNetTestResultResponseBodyDelayTest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyDelayTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyDelayTest) GetHosts() []*DescribeNetTestResultResponseBodyDelayTestHosts {
	return s.Hosts
}

func (s *DescribeNetTestResultResponseBodyDelayTest) SetHosts(v []*DescribeNetTestResultResponseBodyDelayTestHosts) *DescribeNetTestResultResponseBodyDelayTest {
	s.Hosts = v
	return s
}

func (s *DescribeNetTestResultResponseBodyDelayTest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyDelayTestHosts) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) GetBond() *string {
	return s.Bond
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) GetIP() *string {
	return s.IP
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) GetServerName() *string {
	return s.ServerName
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

func (s *DescribeNetTestResultResponseBodyDelayTestHosts) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTest) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetClients() []*DescribeNetTestResultResponseBodyTrafficTestClients {
	return s.Clients
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetGDR() *string {
	return s.GDR
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetProtocol() *string {
	return s.Protocol
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetQP() *int64 {
	return s.QP
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetServers() []*DescribeNetTestResultResponseBodyTrafficTestServers {
	return s.Servers
}

func (s *DescribeNetTestResultResponseBodyTrafficTest) GetTrafficModel() *string {
	return s.TrafficModel
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

func (s *DescribeNetTestResultResponseBodyTrafficTest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTestClients) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) GetBond() *string {
	return s.Bond
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) GetIP() *string {
	return s.IP
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) GetServerName() *string {
	return s.ServerName
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

func (s *DescribeNetTestResultResponseBodyTrafficTestClients) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponseBodyTrafficTestServers) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) GetBond() *string {
	return s.Bond
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) GetIP() *string {
	return s.IP
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) GetServerName() *string {
	return s.ServerName
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

func (s *DescribeNetTestResultResponseBodyTrafficTestServers) Validate() error {
	return dara.Validate(s)
}

type iDescribeNetTestResultResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNetTestResultResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNetTestResultResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNetTestResultResponseBody) *DescribeNetTestResultResponse
	GetBody() *DescribeNetTestResultResponseBody
}

type DescribeNetTestResultResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNetTestResultResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNetTestResultResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetTestResultResponse) GoString() string {
	return s.String()
}

func (s *DescribeNetTestResultResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNetTestResultResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNetTestResultResponse) GetBody() *DescribeNetTestResultResponseBody {
	return s.Body
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

func (s *DescribeNetTestResultResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeNodeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *DescribeNodeRequest
	GetNodeId() *string
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
	return dara.Prettify(s)
}

func (s DescribeNodeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeNodeRequest) SetNodeId(v string) *DescribeNodeRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeNodeRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeNodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeNodeResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeNodeResponseBody
	GetClusterName() *string
	SetCreateTime(v string) *DescribeNodeResponseBody
	GetCreateTime() *string
	SetExpiredTime(v string) *DescribeNodeResponseBody
	GetExpiredTime() *string
	SetHostname(v string) *DescribeNodeResponseBody
	GetHostname() *string
	SetHpnZone(v string) *DescribeNodeResponseBody
	GetHpnZone() *string
	SetImageId(v string) *DescribeNodeResponseBody
	GetImageId() *string
	SetImageName(v string) *DescribeNodeResponseBody
	GetImageName() *string
	SetMachineType(v string) *DescribeNodeResponseBody
	GetMachineType() *string
	SetNetworks(v []*DescribeNodeResponseBodyNetworks) *DescribeNodeResponseBody
	GetNetworks() []*DescribeNodeResponseBodyNetworks
	SetNodeGroupId(v string) *DescribeNodeResponseBody
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *DescribeNodeResponseBody
	GetNodeGroupName() *string
	SetNodeId(v string) *DescribeNodeResponseBody
	GetNodeId() *string
	SetOperatingState(v string) *DescribeNodeResponseBody
	GetOperatingState() *string
	SetRequestId(v string) *DescribeNodeResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeNodeResponseBody
	GetResourceGroupId() *string
	SetSn(v string) *DescribeNodeResponseBody
	GetSn() *string
	SetUserData(v string) *DescribeNodeResponseBody
	GetUserData() *string
	SetZoneId(v string) *DescribeNodeResponseBody
	GetZoneId() *string
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
	// Expiration time
	//
	// example:
	//
	// 2022-06-23T16:00:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
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
	// 资源组ID
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
	Sn       *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// Zone ID
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeNodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeNodeResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeNodeResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeNodeResponseBody) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeNodeResponseBody) GetHostname() *string {
	return s.Hostname
}

func (s *DescribeNodeResponseBody) GetHpnZone() *string {
	return s.HpnZone
}

func (s *DescribeNodeResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeNodeResponseBody) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeNodeResponseBody) GetMachineType() *string {
	return s.MachineType
}

func (s *DescribeNodeResponseBody) GetNetworks() []*DescribeNodeResponseBodyNetworks {
	return s.Networks
}

func (s *DescribeNodeResponseBody) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeNodeResponseBody) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *DescribeNodeResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeNodeResponseBody) GetOperatingState() *string {
	return s.OperatingState
}

func (s *DescribeNodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodeResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeNodeResponseBody) GetSn() *string {
	return s.Sn
}

func (s *DescribeNodeResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *DescribeNodeResponseBody) GetZoneId() *string {
	return s.ZoneId
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

func (s *DescribeNodeResponseBody) SetExpiredTime(v string) *DescribeNodeResponseBody {
	s.ExpiredTime = &v
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

func (s *DescribeNodeResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeNodeResponseBodyNetworks) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponseBodyNetworks) GetBondName() *string {
	return s.BondName
}

func (s *DescribeNodeResponseBodyNetworks) GetIp() *string {
	return s.Ip
}

func (s *DescribeNodeResponseBodyNetworks) GetSubnetId() *string {
	return s.SubnetId
}

func (s *DescribeNodeResponseBodyNetworks) GetVpdId() *string {
	return s.VpdId
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

func (s *DescribeNodeResponseBodyNetworks) Validate() error {
	return dara.Validate(s)
}

type iDescribeNodeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodeResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodeResponseBody) *DescribeNodeResponse
	GetBody() *DescribeNodeResponseBody
}

type DescribeNodeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodeResponse) GetBody() *DescribeNodeResponseBody {
	return s.Body
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

func (s *DescribeNodeResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *DescribeNodeGroupRequest
	GetNodeGroupId() *string
}

type DescribeNodeGroupRequest struct {
	// This parameter is required.
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
}

func (s DescribeNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeNodeGroupRequest) SetNodeGroupId(v string) *DescribeNodeGroupRequest {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAz(v string) *DescribeNodeGroupResponseBody
	GetAz() *string
	SetClusterId(v string) *DescribeNodeGroupResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeNodeGroupResponseBody
	GetClusterName() *string
	SetCreateTime(v string) *DescribeNodeGroupResponseBody
	GetCreateTime() *string
	SetDisks(v *DescribeNodeGroupResponseBodyDisks) *DescribeNodeGroupResponseBody
	GetDisks() *DescribeNodeGroupResponseBodyDisks
	SetImageId(v string) *DescribeNodeGroupResponseBody
	GetImageId() *string
	SetImageName(v string) *DescribeNodeGroupResponseBody
	GetImageName() *string
	SetMachineType(v string) *DescribeNodeGroupResponseBody
	GetMachineType() *string
	SetNodeCount(v string) *DescribeNodeGroupResponseBody
	GetNodeCount() *string
	SetNodeGroupDescription(v string) *DescribeNodeGroupResponseBody
	GetNodeGroupDescription() *string
	SetNodeGroupId(v string) *DescribeNodeGroupResponseBody
	GetNodeGroupId() *string
	SetNodeGroupName(v string) *DescribeNodeGroupResponseBody
	GetNodeGroupName() *string
	SetRequestId(v string) *DescribeNodeGroupResponseBody
	GetRequestId() *string
	SetUpdateTime(v string) *DescribeNodeGroupResponseBody
	GetUpdateTime() *string
	SetUserData(v string) *DescribeNodeGroupResponseBody
	GetUserData() *string
}

type DescribeNodeGroupResponseBody struct {
	Az                   *string                             `json:"Az,omitempty" xml:"Az,omitempty"`
	ClusterId            *string                             `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	ClusterName          *string                             `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	CreateTime           *string                             `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Disks                *DescribeNodeGroupResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	ImageId              *string                             `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	ImageName            *string                             `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	MachineType          *string                             `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	NodeCount            *string                             `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	NodeGroupDescription *string                             `json:"NodeGroupDescription,omitempty" xml:"NodeGroupDescription,omitempty"`
	NodeGroupId          *string                             `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeGroupName        *string                             `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	RequestId            *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	UpdateTime           *string                             `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	UserData             *string                             `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s DescribeNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupResponseBody) GetAz() *string {
	return s.Az
}

func (s *DescribeNodeGroupResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeNodeGroupResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeNodeGroupResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeNodeGroupResponseBody) GetDisks() *DescribeNodeGroupResponseBodyDisks {
	return s.Disks
}

func (s *DescribeNodeGroupResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeNodeGroupResponseBody) GetImageName() *string {
	return s.ImageName
}

func (s *DescribeNodeGroupResponseBody) GetMachineType() *string {
	return s.MachineType
}

func (s *DescribeNodeGroupResponseBody) GetNodeCount() *string {
	return s.NodeCount
}

func (s *DescribeNodeGroupResponseBody) GetNodeGroupDescription() *string {
	return s.NodeGroupDescription
}

func (s *DescribeNodeGroupResponseBody) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *DescribeNodeGroupResponseBody) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *DescribeNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeNodeGroupResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *DescribeNodeGroupResponseBody) GetUserData() *string {
	return s.UserData
}

func (s *DescribeNodeGroupResponseBody) SetAz(v string) *DescribeNodeGroupResponseBody {
	s.Az = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetClusterId(v string) *DescribeNodeGroupResponseBody {
	s.ClusterId = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetClusterName(v string) *DescribeNodeGroupResponseBody {
	s.ClusterName = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetCreateTime(v string) *DescribeNodeGroupResponseBody {
	s.CreateTime = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetDisks(v *DescribeNodeGroupResponseBodyDisks) *DescribeNodeGroupResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetImageId(v string) *DescribeNodeGroupResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetImageName(v string) *DescribeNodeGroupResponseBody {
	s.ImageName = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetMachineType(v string) *DescribeNodeGroupResponseBody {
	s.MachineType = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetNodeCount(v string) *DescribeNodeGroupResponseBody {
	s.NodeCount = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetNodeGroupDescription(v string) *DescribeNodeGroupResponseBody {
	s.NodeGroupDescription = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetNodeGroupId(v string) *DescribeNodeGroupResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetNodeGroupName(v string) *DescribeNodeGroupResponseBody {
	s.NodeGroupName = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetRequestId(v string) *DescribeNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetUpdateTime(v string) *DescribeNodeGroupResponseBody {
	s.UpdateTime = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) SetUserData(v string) *DescribeNodeGroupResponseBody {
	s.UserData = &v
	return s
}

func (s *DescribeNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeNodeGroupResponseBodyDisks struct {
	SystemDiskPerformanceLevel *string `json:"SystemDiskPerformanceLevel,omitempty" xml:"SystemDiskPerformanceLevel,omitempty"`
	SystemDiskSize             *int32  `json:"SystemDiskSize,omitempty" xml:"SystemDiskSize,omitempty"`
}

func (s DescribeNodeGroupResponseBodyDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupResponseBodyDisks) GetSystemDiskPerformanceLevel() *string {
	return s.SystemDiskPerformanceLevel
}

func (s *DescribeNodeGroupResponseBodyDisks) GetSystemDiskSize() *int32 {
	return s.SystemDiskSize
}

func (s *DescribeNodeGroupResponseBodyDisks) SetSystemDiskPerformanceLevel(v string) *DescribeNodeGroupResponseBodyDisks {
	s.SystemDiskPerformanceLevel = &v
	return s
}

func (s *DescribeNodeGroupResponseBodyDisks) SetSystemDiskSize(v int32) *DescribeNodeGroupResponseBodyDisks {
	s.SystemDiskSize = &v
	return s
}

func (s *DescribeNodeGroupResponseBodyDisks) Validate() error {
	return dara.Validate(s)
}

type iDescribeNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *DescribeNodeGroupResponseBody) *DescribeNodeGroupResponse
	GetBody() *DescribeNodeGroupResponseBody
}

type DescribeNodeGroupResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *DescribeNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeNodeGroupResponse) GetBody() *DescribeNodeGroupResponseBody {
	return s.Body
}

func (s *DescribeNodeGroupResponse) SetHeaders(v map[string]*string) *DescribeNodeGroupResponse {
	s.Headers = v
	return s
}

func (s *DescribeNodeGroupResponse) SetStatusCode(v int32) *DescribeNodeGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeNodeGroupResponse) SetBody(v *DescribeNodeGroupResponseBody) *DescribeNodeGroupResponse {
	s.Body = v
	return s
}

func (s *DescribeNodeGroupResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeRegionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeRegionsRequest
	GetAcceptLanguage() *string
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
	return dara.Prettify(s)
}

func (s DescribeRegionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeRegionsRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeRegionsRequest) SetAcceptLanguage(v string) *DescribeRegionsRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeRegionsRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeRegionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody
	GetRegions() []*DescribeRegionsResponseBodyRegions
	SetRequestId(v string) *DescribeRegionsResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBody) GetRegions() []*DescribeRegionsResponseBodyRegions {
	return s.Regions
}

func (s *DescribeRegionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRegionsResponseBody) SetRegions(v []*DescribeRegionsResponseBodyRegions) *DescribeRegionsResponseBody {
	s.Regions = v
	return s
}

func (s *DescribeRegionsResponseBody) SetRequestId(v string) *DescribeRegionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRegionsResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeRegionsResponseBodyRegions) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponseBodyRegions) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeRegionsResponseBodyRegions) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeRegionsResponseBodyRegions) SetLocalName(v string) *DescribeRegionsResponseBodyRegions {
	s.LocalName = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) SetRegionId(v string) *DescribeRegionsResponseBodyRegions {
	s.RegionId = &v
	return s
}

func (s *DescribeRegionsResponseBodyRegions) Validate() error {
	return dara.Validate(s)
}

type iDescribeRegionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeRegionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeRegionsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeRegionsResponseBody) *DescribeRegionsResponse
	GetBody() *DescribeRegionsResponseBody
}

type DescribeRegionsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeRegionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeRegionsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeRegionsResponse) GoString() string {
	return s.String()
}

func (s *DescribeRegionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeRegionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeRegionsResponse) GetBody() *DescribeRegionsResponseBody {
	return s.Body
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

func (s *DescribeRegionsResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeSendFileResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *DescribeSendFileResultsRequest
	GetInvokeId() *string
	SetNodeId(v string) *DescribeSendFileResultsRequest
	GetNodeId() *string
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
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeSendFileResultsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSendFileResultsRequest) SetInvokeId(v string) *DescribeSendFileResultsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetNodeId(v string) *DescribeSendFileResultsRequest {
	s.NodeId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeSendFileResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvocations(v *DescribeSendFileResultsResponseBodyInvocations) *DescribeSendFileResultsResponseBody
	GetInvocations() *DescribeSendFileResultsResponseBodyInvocations
	SetRequestId(v string) *DescribeSendFileResultsResponseBody
	GetRequestId() *string
	SetTotalCount(v string) *DescribeSendFileResultsResponseBody
	GetTotalCount() *string
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
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBody) GetInvocations() *DescribeSendFileResultsResponseBodyInvocations {
	return s.Invocations
}

func (s *DescribeSendFileResultsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSendFileResultsResponseBody) GetTotalCount() *string {
	return s.TotalCount
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

func (s *DescribeSendFileResultsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSendFileResultsResponseBodyInvocations struct {
	// Command execution ID.
	Invocation []*DescribeSendFileResultsResponseBodyInvocationsInvocation `json:"Invocation,omitempty" xml:"Invocation,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsResponseBodyInvocations) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocations) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocations) GetInvocation() []*DescribeSendFileResultsResponseBodyInvocationsInvocation {
	return s.Invocation
}

func (s *DescribeSendFileResultsResponseBodyInvocations) SetInvocation(v []*DescribeSendFileResultsResponseBodyInvocationsInvocation) *DescribeSendFileResultsResponseBodyInvocations {
	s.Invocation = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocations) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocation) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetContent() *string {
	return s.Content
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetContentType() *string {
	return s.ContentType
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetDescription() *string {
	return s.Description
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetFileGroup() *string {
	return s.FileGroup
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetFileMode() *string {
	return s.FileMode
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetFileOwner() *string {
	return s.FileOwner
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetInvokeNodes() *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes {
	return s.InvokeNodes
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetName() *string {
	return s.Name
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetNodeCount() *int32 {
	return s.NodeCount
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) GetTargetDir() *string {
	return s.TargetDir
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

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocation) Validate() error {
	return dara.Validate(s)
}

type DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes struct {
	// Record of file distribution for the node.
	InvokeNode []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode `json:"InvokeNode,omitempty" xml:"InvokeNode,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) GetInvokeNode() []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode {
	return s.InvokeNode
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) SetInvokeNode(v []*DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes {
	s.InvokeNode = v
	return s
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodes) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetErrorInfo() *string {
	return s.ErrorInfo
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetFinishTime() *string {
	return s.FinishTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) GetUpdateTime() *string {
	return s.UpdateTime
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

func (s *DescribeSendFileResultsResponseBodyInvocationsInvocationInvokeNodesInvokeNode) Validate() error {
	return dara.Validate(s)
}

type iDescribeSendFileResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSendFileResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSendFileResultsResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSendFileResultsResponseBody) *DescribeSendFileResultsResponse
	GetBody() *DescribeSendFileResultsResponseBody
}

type DescribeSendFileResultsResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSendFileResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSendFileResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsResponse) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSendFileResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSendFileResultsResponse) GetBody() *DescribeSendFileResultsResponseBody {
	return s.Body
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

func (s *DescribeSendFileResultsResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeTaskRequest
	GetTaskId() *string
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
	return dara.Prettify(s)
}

func (s DescribeTaskRequest) GoString() string {
	return s.String()
}

func (s *DescribeTaskRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskRequest) SetTaskId(v string) *DescribeTaskRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeTaskRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *DescribeTaskResponseBody
	GetClusterId() *string
	SetClusterName(v string) *DescribeTaskResponseBody
	GetClusterName() *string
	SetCreateTime(v string) *DescribeTaskResponseBody
	GetCreateTime() *string
	SetMessage(v string) *DescribeTaskResponseBody
	GetMessage() *string
	SetNodeIds(v []*string) *DescribeTaskResponseBody
	GetNodeIds() []*string
	SetRequestId(v string) *DescribeTaskResponseBody
	GetRequestId() *string
	SetSteps(v []*DescribeTaskResponseBodySteps) *DescribeTaskResponseBody
	GetSteps() []*DescribeTaskResponseBodySteps
	SetTaskState(v string) *DescribeTaskResponseBody
	GetTaskState() *string
	SetTaskType(v string) *DescribeTaskResponseBody
	GetTaskType() *string
	SetUpdateTime(v string) *DescribeTaskResponseBody
	GetUpdateTime() *string
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
	return dara.Prettify(s)
}

func (s DescribeTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBody) GetClusterId() *string {
	return s.ClusterId
}

func (s *DescribeTaskResponseBody) GetClusterName() *string {
	return s.ClusterName
}

func (s *DescribeTaskResponseBody) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeTaskResponseBody) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *DescribeTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTaskResponseBody) GetSteps() []*DescribeTaskResponseBodySteps {
	return s.Steps
}

func (s *DescribeTaskResponseBody) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeTaskResponseBody) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTaskResponseBody) GetUpdateTime() *string {
	return s.UpdateTime
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

func (s *DescribeTaskResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeTaskResponseBodySteps) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBodySteps) GetMessage() *string {
	return s.Message
}

func (s *DescribeTaskResponseBodySteps) GetStageTag() *string {
	return s.StageTag
}

func (s *DescribeTaskResponseBodySteps) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeTaskResponseBodySteps) GetStepName() *string {
	return s.StepName
}

func (s *DescribeTaskResponseBodySteps) GetStepState() *string {
	return s.StepState
}

func (s *DescribeTaskResponseBodySteps) GetStepType() *string {
	return s.StepType
}

func (s *DescribeTaskResponseBodySteps) GetSubTasks() []*DescribeTaskResponseBodyStepsSubTasks {
	return s.SubTasks
}

func (s *DescribeTaskResponseBodySteps) GetUpdateTime() *string {
	return s.UpdateTime
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

func (s *DescribeTaskResponseBodySteps) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeTaskResponseBodyStepsSubTasks) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetMessage() *string {
	return s.Message
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetTaskState() *string {
	return s.TaskState
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetTaskType() *string {
	return s.TaskType
}

func (s *DescribeTaskResponseBodyStepsSubTasks) GetUpdateTime() *string {
	return s.UpdateTime
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

func (s *DescribeTaskResponseBodyStepsSubTasks) Validate() error {
	return dara.Validate(s)
}

type iDescribeTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeTaskResponse
	GetStatusCode() *int32
	SetBody(v *DescribeTaskResponseBody) *DescribeTaskResponse
	GetBody() *DescribeTaskResponseBody
}

type DescribeTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeTaskResponse) GoString() string {
	return s.String()
}

func (s *DescribeTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeTaskResponse) GetBody() *DescribeTaskResponseBody {
	return s.Body
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

func (s *DescribeTaskResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeVscRequest interface {
	dara.Model
	String() string
	GoString() string
	SetVscId(v string) *DescribeVscRequest
	GetVscId() *string
}

type DescribeVscRequest struct {
	// This parameter is required.
	VscId *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
}

func (s DescribeVscRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscRequest) GoString() string {
	return s.String()
}

func (s *DescribeVscRequest) GetVscId() *string {
	return s.VscId
}

func (s *DescribeVscRequest) SetVscId(v string) *DescribeVscRequest {
	s.VscId = &v
	return s
}

func (s *DescribeVscRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeVscResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeId(v string) *DescribeVscResponseBody
	GetNodeId() *string
	SetRequestId(v string) *DescribeVscResponseBody
	GetRequestId() *string
	SetResourceGroupId(v string) *DescribeVscResponseBody
	GetResourceGroupId() *string
	SetStatus(v string) *DescribeVscResponseBody
	GetStatus() *string
	SetVscId(v string) *DescribeVscResponseBody
	GetVscId() *string
	SetVscName(v string) *DescribeVscResponseBody
	GetVscName() *string
	SetVscType(v string) *DescribeVscResponseBody
	GetVscType() *string
}

type DescribeVscResponseBody struct {
	NodeId          *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	RequestId       *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VscId           *string `json:"VscId,omitempty" xml:"VscId,omitempty"`
	VscName         *string `json:"VscName,omitempty" xml:"VscName,omitempty"`
	VscType         *string `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s DescribeVscResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVscResponseBody) GetNodeId() *string {
	return s.NodeId
}

func (s *DescribeVscResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVscResponseBody) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVscResponseBody) GetStatus() *string {
	return s.Status
}

func (s *DescribeVscResponseBody) GetVscId() *string {
	return s.VscId
}

func (s *DescribeVscResponseBody) GetVscName() *string {
	return s.VscName
}

func (s *DescribeVscResponseBody) GetVscType() *string {
	return s.VscType
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

func (s *DescribeVscResponseBody) Validate() error {
	return dara.Validate(s)
}

type iDescribeVscResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeVscResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeVscResponse
	GetStatusCode() *int32
	SetBody(v *DescribeVscResponseBody) *DescribeVscResponse
	GetBody() *DescribeVscResponseBody
}

type DescribeVscResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeVscResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeVscResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscResponse) GoString() string {
	return s.String()
}

func (s *DescribeVscResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeVscResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeVscResponse) GetBody() *DescribeVscResponseBody {
	return s.Body
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

func (s *DescribeVscResponse) Validate() error {
	return dara.Validate(s)
}

type iDescribeZonesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAcceptLanguage(v string) *DescribeZonesRequest
	GetAcceptLanguage() *string
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
	return dara.Prettify(s)
}

func (s DescribeZonesRequest) GoString() string {
	return s.String()
}

func (s *DescribeZonesRequest) GetAcceptLanguage() *string {
	return s.AcceptLanguage
}

func (s *DescribeZonesRequest) SetAcceptLanguage(v string) *DescribeZonesRequest {
	s.AcceptLanguage = &v
	return s
}

func (s *DescribeZonesRequest) Validate() error {
	return dara.Validate(s)
}

type iDescribeZonesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeZonesResponseBody
	GetRequestId() *string
	SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody
	GetZones() []*DescribeZonesResponseBodyZones
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
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeZonesResponseBody) GetZones() []*DescribeZonesResponseBodyZones {
	return s.Zones
}

func (s *DescribeZonesResponseBody) SetRequestId(v string) *DescribeZonesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeZonesResponseBody) SetZones(v []*DescribeZonesResponseBodyZones) *DescribeZonesResponseBody {
	s.Zones = v
	return s
}

func (s *DescribeZonesResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s DescribeZonesResponseBodyZones) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponseBodyZones) GetLocalName() *string {
	return s.LocalName
}

func (s *DescribeZonesResponseBodyZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeZonesResponseBodyZones) SetLocalName(v string) *DescribeZonesResponseBodyZones {
	s.LocalName = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) SetZoneId(v string) *DescribeZonesResponseBodyZones {
	s.ZoneId = &v
	return s
}

func (s *DescribeZonesResponseBodyZones) Validate() error {
	return dara.Validate(s)
}

type iDescribeZonesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeZonesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeZonesResponse
	GetStatusCode() *int32
	SetBody(v *DescribeZonesResponseBody) *DescribeZonesResponse
	GetBody() *DescribeZonesResponseBody
}

type DescribeZonesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeZonesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeZonesResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeZonesResponse) GoString() string {
	return s.String()
}

func (s *DescribeZonesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeZonesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeZonesResponse) GetBody() *DescribeZonesResponseBody {
	return s.Body
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

func (s *DescribeZonesResponse) Validate() error {
	return dara.Validate(s)
}

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
	return dara.Prettify(s)
}

func (s ExtendClusterRequest) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ExtendClusterRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ExtendClusterRequest) GetIpAllocationPolicy() []*ExtendClusterRequestIpAllocationPolicy {
	return s.IpAllocationPolicy
}

func (s *ExtendClusterRequest) GetNodeGroups() []*ExtendClusterRequestNodeGroups {
	return s.NodeGroups
}

func (s *ExtendClusterRequest) GetVSwitchZoneId() *string {
	return s.VSwitchZoneId
}

func (s *ExtendClusterRequest) GetVpdSubnets() []*string {
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

func (s *ExtendClusterRequestIpAllocationPolicy) GetBondPolicy() *ExtendClusterRequestIpAllocationPolicyBondPolicy {
	return s.BondPolicy
}

func (s *ExtendClusterRequestIpAllocationPolicy) GetMachineTypePolicy() []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicy {
	return s.MachineTypePolicy
}

func (s *ExtendClusterRequestIpAllocationPolicy) GetNodePolicy() []*ExtendClusterRequestIpAllocationPolicyNodePolicy {
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

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) GetBondDefaultSubnet() *string {
	return s.BondDefaultSubnet
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicy) GetBonds() []*ExtendClusterRequestIpAllocationPolicyBondPolicyBonds {
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

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) GetName() *string {
	return s.Name
}

func (s *ExtendClusterRequestIpAllocationPolicyBondPolicyBonds) GetSubnet() *string {
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
	// Machine Type
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

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) GetBonds() []*ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds {
	return s.Bonds
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicy) GetMachineType() *string {
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

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) GetName() *string {
	return s.Name
}

func (s *ExtendClusterRequestIpAllocationPolicyMachineTypePolicyBonds) GetSubnet() *string {
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
	Bonds    []*ExtendClusterRequestIpAllocationPolicyNodePolicyBonds `json:"Bonds,omitempty" xml:"Bonds,omitempty" type:"Repeated"`
	Hostname *string                                                  `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
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

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) GetBonds() []*ExtendClusterRequestIpAllocationPolicyNodePolicyBonds {
	return s.Bonds
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) GetHostname() *string {
	return s.Hostname
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicy) GetNodeId() *string {
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

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) GetName() *string {
	return s.Name
}

func (s *ExtendClusterRequestIpAllocationPolicyNodePolicyBonds) GetSubnet() *string {
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
	Amount        *int64    `json:"Amount,omitempty" xml:"Amount,omitempty"`
	AutoRenew     *bool     `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	ChargeType    *string   `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	Hostnames     []*string `json:"Hostnames,omitempty" xml:"Hostnames,omitempty" type:"Repeated"`
	LoginPassword *string   `json:"LoginPassword,omitempty" xml:"LoginPassword,omitempty"`
	// Node Group ID
	//
	// example:
	//
	// i16d4883a46cbadeb4bc9
	NodeGroupId *string                                  `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	NodeTag     []*ExtendClusterRequestNodeGroupsNodeTag `json:"NodeTag,omitempty" xml:"NodeTag,omitempty" type:"Repeated"`
	// List of Nodes
	Nodes  []*ExtendClusterRequestNodeGroupsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	Period *int64                                 `json:"Period,omitempty" xml:"Period,omitempty"`
	// Custom Data
	//
	// example:
	//
	// #!/bin/sh
	//
	// echo "Hello World. The time is now $(date -R)!" | tee /root/userdata_test.txt
	UserData  *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId     *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
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

func (s *ExtendClusterRequestNodeGroups) GetAmount() *int64 {
	return s.Amount
}

func (s *ExtendClusterRequestNodeGroups) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *ExtendClusterRequestNodeGroups) GetChargeType() *string {
	return s.ChargeType
}

func (s *ExtendClusterRequestNodeGroups) GetHostnames() []*string {
	return s.Hostnames
}

func (s *ExtendClusterRequestNodeGroups) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *ExtendClusterRequestNodeGroups) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ExtendClusterRequestNodeGroups) GetNodeTag() []*ExtendClusterRequestNodeGroupsNodeTag {
	return s.NodeTag
}

func (s *ExtendClusterRequestNodeGroups) GetNodes() []*ExtendClusterRequestNodeGroupsNodes {
	return s.Nodes
}

func (s *ExtendClusterRequestNodeGroups) GetPeriod() *int64 {
	return s.Period
}

func (s *ExtendClusterRequestNodeGroups) GetUserData() *string {
	return s.UserData
}

func (s *ExtendClusterRequestNodeGroups) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ExtendClusterRequestNodeGroups) GetVpcId() *string {
	return s.VpcId
}

func (s *ExtendClusterRequestNodeGroups) GetZoneId() *string {
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
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ExtendClusterRequestNodeGroupsNodeTag) String() string {
	return dara.Prettify(s)
}

func (s ExtendClusterRequestNodeGroupsNodeTag) GoString() string {
	return s.String()
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) GetKey() *string {
	return s.Key
}

func (s *ExtendClusterRequestNodeGroupsNodeTag) GetValue() *string {
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

func (s *ExtendClusterRequestNodeGroupsNodes) GetHostname() *string {
	return s.Hostname
}

func (s *ExtendClusterRequestNodeGroupsNodes) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *ExtendClusterRequestNodeGroupsNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ExtendClusterRequestNodeGroupsNodes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ExtendClusterRequestNodeGroupsNodes) GetVpcId() *string {
	return s.VpcId
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

type iExtendClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ExtendClusterShrinkRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *ExtendClusterShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetIpAllocationPolicyShrink(v string) *ExtendClusterShrinkRequest
	GetIpAllocationPolicyShrink() *string
	SetNodeGroupsShrink(v string) *ExtendClusterShrinkRequest
	GetNodeGroupsShrink() *string
	SetVSwitchZoneId(v string) *ExtendClusterShrinkRequest
	GetVSwitchZoneId() *string
	SetVpdSubnetsShrink(v string) *ExtendClusterShrinkRequest
	GetVpdSubnetsShrink() *string
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
	return dara.Prettify(s)
}

func (s ExtendClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ExtendClusterShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ExtendClusterShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ExtendClusterShrinkRequest) GetIpAllocationPolicyShrink() *string {
	return s.IpAllocationPolicyShrink
}

func (s *ExtendClusterShrinkRequest) GetNodeGroupsShrink() *string {
	return s.NodeGroupsShrink
}

func (s *ExtendClusterShrinkRequest) GetVSwitchZoneId() *string {
	return s.VSwitchZoneId
}

func (s *ExtendClusterShrinkRequest) GetVpdSubnetsShrink() *string {
	return s.VpdSubnetsShrink
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

func (s *ExtendClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iExtendClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ExtendClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ExtendClusterResponseBody
	GetTaskId() *string
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
	return dara.Prettify(s)
}

func (s ExtendClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ExtendClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ExtendClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ExtendClusterResponseBody) SetRequestId(v string) *ExtendClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ExtendClusterResponseBody) SetTaskId(v string) *ExtendClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ExtendClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type iExtendClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ExtendClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ExtendClusterResponse
	GetStatusCode() *int32
	SetBody(v *ExtendClusterResponseBody) *ExtendClusterResponse
	GetBody() *ExtendClusterResponseBody
}

type ExtendClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ExtendClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ExtendClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ExtendClusterResponse) GoString() string {
	return s.String()
}

func (s *ExtendClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ExtendClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ExtendClusterResponse) GetBody() *ExtendClusterResponseBody {
	return s.Body
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

func (s *ExtendClusterResponse) Validate() error {
	return dara.Validate(s)
}

type iListClusterNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListClusterNodesRequest
	GetClusterId() *string
	SetMaxResults(v int64) *ListClusterNodesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListClusterNodesRequest
	GetNextToken() *string
	SetNodeGroupId(v string) *ListClusterNodesRequest
	GetNodeGroupId() *string
	SetResourceGroupId(v string) *ListClusterNodesRequest
	GetResourceGroupId() *string
	SetTags(v []*ListClusterNodesRequestTags) *ListClusterNodesRequest
	GetTags() []*ListClusterNodesRequestTags
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
	NodeGroupId     *string                        `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	ResourceGroupId *string                        `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*ListClusterNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListClusterNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClusterNodesRequest) GoString() string {
	return s.String()
}

func (s *ListClusterNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClusterNodesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListClusterNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClusterNodesRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListClusterNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClusterNodesRequest) GetTags() []*ListClusterNodesRequestTags {
	return s.Tags
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

func (s *ListClusterNodesRequest) Validate() error {
	return dara.Validate(s)
}

type ListClusterNodesRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClusterNodesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListClusterNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListClusterNodesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListClusterNodesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListClusterNodesRequestTags) SetKey(v string) *ListClusterNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListClusterNodesRequestTags) SetValue(v string) *ListClusterNodesRequestTags {
	s.Value = &v
	return s
}

func (s *ListClusterNodesRequestTags) Validate() error {
	return dara.Validate(s)
}

type iListClusterNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListClusterNodesResponseBody
	GetNextToken() *string
	SetNodes(v []*ListClusterNodesResponseBodyNodes) *ListClusterNodesResponseBody
	GetNodes() []*ListClusterNodesResponseBodyNodes
	SetRequestId(v string) *ListClusterNodesResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListClusterNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClusterNodesResponseBody) GetNodes() []*ListClusterNodesResponseBodyNodes {
	return s.Nodes
}

func (s *ListClusterNodesResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *ListClusterNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListClusterNodesResponseBodyNodes struct {
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
	ImageId   *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
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
	Sn     *string                                  `json:"Sn,omitempty" xml:"Sn,omitempty"`
	Tags   []*ListClusterNodesResponseBodyNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	TaskId *string                                  `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// 专有网络交换机ID
	//
	// example:
	//
	// vsw-bp1mxqhw8o20tgv3xk47h
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// 专有网络ID
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
	return dara.Prettify(s)
}

func (s ListClusterNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBodyNodes) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListClusterNodesResponseBodyNodes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClusterNodesResponseBodyNodes) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListClusterNodesResponseBodyNodes) GetHostname() *string {
	return s.Hostname
}

func (s *ListClusterNodesResponseBodyNodes) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListClusterNodesResponseBodyNodes) GetImageId() *string {
	return s.ImageId
}

func (s *ListClusterNodesResponseBodyNodes) GetImageName() *string {
	return s.ImageName
}

func (s *ListClusterNodesResponseBodyNodes) GetMachineType() *string {
	return s.MachineType
}

func (s *ListClusterNodesResponseBodyNodes) GetNetworks() []*ListClusterNodesResponseBodyNodesNetworks {
	return s.Networks
}

func (s *ListClusterNodesResponseBodyNodes) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ListClusterNodesResponseBodyNodes) GetNodeGroupName() *string {
	return s.NodeGroupName
}

func (s *ListClusterNodesResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ListClusterNodesResponseBodyNodes) GetOperatingState() *string {
	return s.OperatingState
}

func (s *ListClusterNodesResponseBodyNodes) GetSn() *string {
	return s.Sn
}

func (s *ListClusterNodesResponseBodyNodes) GetTags() []*ListClusterNodesResponseBodyNodesTags {
	return s.Tags
}

func (s *ListClusterNodesResponseBodyNodes) GetTaskId() *string {
	return s.TaskId
}

func (s *ListClusterNodesResponseBodyNodes) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListClusterNodesResponseBodyNodes) GetVpcId() *string {
	return s.VpcId
}

func (s *ListClusterNodesResponseBodyNodes) GetZoneId() *string {
	return s.ZoneId
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

func (s *ListClusterNodesResponseBodyNodes) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListClusterNodesResponseBodyNodesNetworks) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBodyNodesNetworks) GetBondName() *string {
	return s.BondName
}

func (s *ListClusterNodesResponseBodyNodesNetworks) GetIp() *string {
	return s.Ip
}

func (s *ListClusterNodesResponseBodyNodesNetworks) GetSubnetId() *string {
	return s.SubnetId
}

func (s *ListClusterNodesResponseBodyNodesNetworks) GetVpdId() *string {
	return s.VpdId
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

func (s *ListClusterNodesResponseBodyNodesNetworks) Validate() error {
	return dara.Validate(s)
}

type ListClusterNodesResponseBodyNodesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClusterNodesResponseBodyNodesTags) String() string {
	return dara.Prettify(s)
}

func (s ListClusterNodesResponseBodyNodesTags) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponseBodyNodesTags) GetKey() *string {
	return s.Key
}

func (s *ListClusterNodesResponseBodyNodesTags) GetValue() *string {
	return s.Value
}

func (s *ListClusterNodesResponseBodyNodesTags) SetKey(v string) *ListClusterNodesResponseBodyNodesTags {
	s.Key = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodesTags) SetValue(v string) *ListClusterNodesResponseBodyNodesTags {
	s.Value = &v
	return s
}

func (s *ListClusterNodesResponseBodyNodesTags) Validate() error {
	return dara.Validate(s)
}

type iListClusterNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClusterNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClusterNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListClusterNodesResponseBody) *ListClusterNodesResponse
	GetBody() *ListClusterNodesResponseBody
}

type ListClusterNodesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClusterNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClusterNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClusterNodesResponse) GoString() string {
	return s.String()
}

func (s *ListClusterNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClusterNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClusterNodesResponse) GetBody() *ListClusterNodesResponseBody {
	return s.Body
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

func (s *ListClusterNodesResponse) Validate() error {
	return dara.Validate(s)
}

type iListClustersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListClustersRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListClustersRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListClustersRequest
	GetResourceGroupId() *string
	SetTags(v []*ListClustersRequestTags) *ListClustersRequest
	GetTags() []*ListClustersRequestTags
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
	ResourceGroupId *string                    `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*ListClustersRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListClustersRequest) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequest) GoString() string {
	return s.String()
}

func (s *ListClustersRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListClustersRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClustersRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClustersRequest) GetTags() []*ListClustersRequestTags {
	return s.Tags
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

func (s *ListClustersRequest) Validate() error {
	return dara.Validate(s)
}

type ListClustersRequestTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListClustersRequestTags) GoString() string {
	return s.String()
}

func (s *ListClustersRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListClustersRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListClustersRequestTags) SetKey(v string) *ListClustersRequestTags {
	s.Key = &v
	return s
}

func (s *ListClustersRequestTags) SetValue(v string) *ListClustersRequestTags {
	s.Value = &v
	return s
}

func (s *ListClustersRequestTags) Validate() error {
	return dara.Validate(s)
}

type iListClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusters(v []*ListClustersResponseBodyClusters) *ListClustersResponseBody
	GetClusters() []*ListClustersResponseBodyClusters
	SetNextToken(v string) *ListClustersResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListClustersResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBody) GetClusters() []*ListClustersResponseBodyClusters {
	return s.Clusters
}

func (s *ListClustersResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListClustersResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *ListClustersResponseBody) Validate() error {
	return dara.Validate(s)
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
	ResourceGroupId *string                                 `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*ListClustersResponseBodyClustersTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
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
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClusters) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClusters) GetClusterDescription() *string {
	return s.ClusterDescription
}

func (s *ListClustersResponseBodyClusters) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListClustersResponseBodyClusters) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListClustersResponseBodyClusters) GetClusterType() *string {
	return s.ClusterType
}

func (s *ListClustersResponseBodyClusters) GetComponents() interface{} {
	return s.Components
}

func (s *ListClustersResponseBodyClusters) GetComputingIpVersion() *string {
	return s.ComputingIpVersion
}

func (s *ListClustersResponseBodyClusters) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListClustersResponseBodyClusters) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListClustersResponseBodyClusters) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *ListClustersResponseBodyClusters) GetNodeGroupCount() *int64 {
	return s.NodeGroupCount
}

func (s *ListClustersResponseBodyClusters) GetOperatingState() *string {
	return s.OperatingState
}

func (s *ListClustersResponseBodyClusters) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListClustersResponseBodyClusters) GetTags() []*ListClustersResponseBodyClustersTags {
	return s.Tags
}

func (s *ListClustersResponseBodyClusters) GetTaskId() *string {
	return s.TaskId
}

func (s *ListClustersResponseBodyClusters) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListClustersResponseBodyClusters) GetVpcId() *string {
	return s.VpcId
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

func (s *ListClustersResponseBodyClusters) Validate() error {
	return dara.Validate(s)
}

type ListClustersResponseBodyClustersTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListClustersResponseBodyClustersTags) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponseBodyClustersTags) GoString() string {
	return s.String()
}

func (s *ListClustersResponseBodyClustersTags) GetKey() *string {
	return s.Key
}

func (s *ListClustersResponseBodyClustersTags) GetValue() *string {
	return s.Value
}

func (s *ListClustersResponseBodyClustersTags) SetKey(v string) *ListClustersResponseBodyClustersTags {
	s.Key = &v
	return s
}

func (s *ListClustersResponseBodyClustersTags) SetValue(v string) *ListClustersResponseBodyClustersTags {
	s.Value = &v
	return s
}

func (s *ListClustersResponseBodyClustersTags) Validate() error {
	return dara.Validate(s)
}

type iListClustersResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClustersResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClustersResponse
	GetStatusCode() *int32
	SetBody(v *ListClustersResponseBody) *ListClustersResponse
	GetBody() *ListClustersResponseBody
}

type ListClustersResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClustersResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClustersResponse) GoString() string {
	return s.String()
}

func (s *ListClustersResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClustersResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClustersResponse) GetBody() *ListClustersResponseBody {
	return s.Body
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

func (s *ListClustersResponse) Validate() error {
	return dara.Validate(s)
}

type iListDiagnosticResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiagType(v string) *ListDiagnosticResultsRequest
	GetDiagType() *string
	SetMaxResults(v int64) *ListDiagnosticResultsRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListDiagnosticResultsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListDiagnosticResultsRequest
	GetResourceGroupId() *string
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
	return dara.Prettify(s)
}

func (s ListDiagnosticResultsRequest) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsRequest) GetDiagType() *string {
	return s.DiagType
}

func (s *ListDiagnosticResultsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListDiagnosticResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDiagnosticResultsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *ListDiagnosticResultsRequest) Validate() error {
	return dara.Validate(s)
}

type iListDiagnosticResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDiagnosticResults(v []*ListDiagnosticResultsResponseBodyDiagnosticResults) *ListDiagnosticResultsResponseBody
	GetDiagnosticResults() []*ListDiagnosticResultsResponseBodyDiagnosticResults
	SetMaxResults(v int64) *ListDiagnosticResultsResponseBody
	GetMaxResults() *int64
	SetNextToken(v string) *ListDiagnosticResultsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDiagnosticResultsResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListDiagnosticResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponseBody) GetDiagnosticResults() []*ListDiagnosticResultsResponseBodyDiagnosticResults {
	return s.DiagnosticResults
}

func (s *ListDiagnosticResultsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListDiagnosticResultsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDiagnosticResultsResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *ListDiagnosticResultsResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListDiagnosticResultsResponseBodyDiagnosticResults) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetDiagContent() *string {
	return s.DiagContent
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetDiagId() *string {
	return s.DiagId
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetDiagResult() *string {
	return s.DiagResult
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetServerName() *string {
	return s.ServerName
}

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) GetStatus() *string {
	return s.Status
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

func (s *ListDiagnosticResultsResponseBodyDiagnosticResults) Validate() error {
	return dara.Validate(s)
}

type iListDiagnosticResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDiagnosticResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDiagnosticResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListDiagnosticResultsResponseBody) *ListDiagnosticResultsResponse
	GetBody() *ListDiagnosticResultsResponseBody
}

type ListDiagnosticResultsResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDiagnosticResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDiagnosticResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDiagnosticResultsResponse) GoString() string {
	return s.String()
}

func (s *ListDiagnosticResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDiagnosticResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDiagnosticResultsResponse) GetBody() *ListDiagnosticResultsResponseBody {
	return s.Body
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

func (s *ListDiagnosticResultsResponse) Validate() error {
	return dara.Validate(s)
}

type iListFreeNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHpnZone(v string) *ListFreeNodesRequest
	GetHpnZone() *string
	SetIncludeAbnormalNodes(v bool) *ListFreeNodesRequest
	GetIncludeAbnormalNodes() *bool
	SetMachineType(v string) *ListFreeNodesRequest
	GetMachineType() *string
	SetMaxResults(v int64) *ListFreeNodesRequest
	GetMaxResults() *int64
	SetNextToken(v string) *ListFreeNodesRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListFreeNodesRequest
	GetResourceGroupId() *string
	SetTags(v []*ListFreeNodesRequestTags) *ListFreeNodesRequest
	GetTags() []*ListFreeNodesRequestTags
}

type ListFreeNodesRequest struct {
	// Cluster number
	//
	// example:
	//
	// A1
	HpnZone              *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	IncludeAbnormalNodes *bool   `json:"IncludeAbnormalNodes,omitempty" xml:"IncludeAbnormalNodes,omitempty"`
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
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Resource group ID
	//
	// example:
	//
	// rg-acfmxno4vh5muoq
	ResourceGroupId *string                     `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tags            []*ListFreeNodesRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListFreeNodesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesRequest) GoString() string {
	return s.String()
}

func (s *ListFreeNodesRequest) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListFreeNodesRequest) GetIncludeAbnormalNodes() *bool {
	return s.IncludeAbnormalNodes
}

func (s *ListFreeNodesRequest) GetMachineType() *string {
	return s.MachineType
}

func (s *ListFreeNodesRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListFreeNodesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFreeNodesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListFreeNodesRequest) GetTags() []*ListFreeNodesRequestTags {
	return s.Tags
}

func (s *ListFreeNodesRequest) SetHpnZone(v string) *ListFreeNodesRequest {
	s.HpnZone = &v
	return s
}

func (s *ListFreeNodesRequest) SetIncludeAbnormalNodes(v bool) *ListFreeNodesRequest {
	s.IncludeAbnormalNodes = &v
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

func (s *ListFreeNodesRequest) SetResourceGroupId(v string) *ListFreeNodesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ListFreeNodesRequest) SetTags(v []*ListFreeNodesRequestTags) *ListFreeNodesRequest {
	s.Tags = v
	return s
}

func (s *ListFreeNodesRequest) Validate() error {
	return dara.Validate(s)
}

type ListFreeNodesRequestTags struct {
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// 129
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFreeNodesRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesRequestTags) GoString() string {
	return s.String()
}

func (s *ListFreeNodesRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListFreeNodesRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListFreeNodesRequestTags) SetKey(v string) *ListFreeNodesRequestTags {
	s.Key = &v
	return s
}

func (s *ListFreeNodesRequestTags) SetValue(v string) *ListFreeNodesRequestTags {
	s.Value = &v
	return s
}

func (s *ListFreeNodesRequestTags) Validate() error {
	return dara.Validate(s)
}

type iListFreeNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListFreeNodesResponseBody
	GetNextToken() *string
	SetNodes(v []*ListFreeNodesResponseBodyNodes) *ListFreeNodesResponseBody
	GetNodes() []*ListFreeNodesResponseBodyNodes
	SetRequestId(v string) *ListFreeNodesResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListFreeNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListFreeNodesResponseBody) GetNodes() []*ListFreeNodesResponseBodyNodes {
	return s.Nodes
}

func (s *ListFreeNodesResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *ListFreeNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListFreeNodesResponseBodyNodes struct {
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
	NodeId         *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
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
	Sn   *string                               `json:"Sn,omitempty" xml:"Sn,omitempty"`
	Tags []*ListFreeNodesResponseBodyNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// Availability zone ID
	//
	// example:
	//
	// cn-hangzhou-j
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListFreeNodesResponseBodyNodes) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesResponseBodyNodes) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBodyNodes) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *ListFreeNodesResponseBodyNodes) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListFreeNodesResponseBodyNodes) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *ListFreeNodesResponseBodyNodes) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListFreeNodesResponseBodyNodes) GetMachineType() *string {
	return s.MachineType
}

func (s *ListFreeNodesResponseBodyNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ListFreeNodesResponseBodyNodes) GetOperatingState() *string {
	return s.OperatingState
}

func (s *ListFreeNodesResponseBodyNodes) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListFreeNodesResponseBodyNodes) GetSn() *string {
	return s.Sn
}

func (s *ListFreeNodesResponseBodyNodes) GetTags() []*ListFreeNodesResponseBodyNodesTags {
	return s.Tags
}

func (s *ListFreeNodesResponseBodyNodes) GetZoneId() *string {
	return s.ZoneId
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

func (s *ListFreeNodesResponseBodyNodes) Validate() error {
	return dara.Validate(s)
}

type ListFreeNodesResponseBodyNodesTags struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListFreeNodesResponseBodyNodesTags) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesResponseBodyNodesTags) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponseBodyNodesTags) GetKey() *string {
	return s.Key
}

func (s *ListFreeNodesResponseBodyNodesTags) GetValue() *string {
	return s.Value
}

func (s *ListFreeNodesResponseBodyNodesTags) SetKey(v string) *ListFreeNodesResponseBodyNodesTags {
	s.Key = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodesTags) SetValue(v string) *ListFreeNodesResponseBodyNodesTags {
	s.Value = &v
	return s
}

func (s *ListFreeNodesResponseBodyNodesTags) Validate() error {
	return dara.Validate(s)
}

type iListFreeNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListFreeNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListFreeNodesResponse
	GetStatusCode() *int32
	SetBody(v *ListFreeNodesResponseBody) *ListFreeNodesResponse
	GetBody() *ListFreeNodesResponseBody
}

type ListFreeNodesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFreeNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFreeNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListFreeNodesResponse) GoString() string {
	return s.String()
}

func (s *ListFreeNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListFreeNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListFreeNodesResponse) GetBody() *ListFreeNodesResponseBody {
	return s.Body
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

func (s *ListFreeNodesResponse) Validate() error {
	return dara.Validate(s)
}

type iListImagesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetArchitecture(v string) *ListImagesRequest
	GetArchitecture() *string
	SetImageVersion(v string) *ListImagesRequest
	GetImageVersion() *string
	SetPlatform(v string) *ListImagesRequest
	GetPlatform() *string
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
	return dara.Prettify(s)
}

func (s ListImagesRequest) GoString() string {
	return s.String()
}

func (s *ListImagesRequest) GetArchitecture() *string {
	return s.Architecture
}

func (s *ListImagesRequest) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *ListImagesRequest) GetPlatform() *string {
	return s.Platform
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

func (s *ListImagesRequest) Validate() error {
	return dara.Validate(s)
}

type iListImagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImages(v []*ListImagesResponseBodyImages) *ListImagesResponseBody
	GetImages() []*ListImagesResponseBodyImages
	SetNextToken(v string) *ListImagesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListImagesResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListImagesResponseBody) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBody) GetImages() []*ListImagesResponseBodyImages {
	return s.Images
}

func (s *ListImagesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListImagesResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *ListImagesResponseBody) Validate() error {
	return dara.Validate(s)
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
	ReleaseFileSize *int64 `json:"ReleaseFileSize,omitempty" xml:"ReleaseFileSize,omitempty"`
	// image type
	//
	// example:
	//
	// Public
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s ListImagesResponseBodyImages) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponseBodyImages) GoString() string {
	return s.String()
}

func (s *ListImagesResponseBodyImages) GetArchitecture() *string {
	return s.Architecture
}

func (s *ListImagesResponseBodyImages) GetDescription() *string {
	return s.Description
}

func (s *ListImagesResponseBodyImages) GetImageId() *string {
	return s.ImageId
}

func (s *ListImagesResponseBodyImages) GetImageName() *string {
	return s.ImageName
}

func (s *ListImagesResponseBodyImages) GetImageVersion() *string {
	return s.ImageVersion
}

func (s *ListImagesResponseBodyImages) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *ListImagesResponseBodyImages) GetPlatform() *string {
	return s.Platform
}

func (s *ListImagesResponseBodyImages) GetReleaseFileMd5() *string {
	return s.ReleaseFileMd5
}

func (s *ListImagesResponseBodyImages) GetReleaseFileSize() *int64 {
	return s.ReleaseFileSize
}

func (s *ListImagesResponseBodyImages) GetType() *string {
	return s.Type
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

func (s *ListImagesResponseBodyImages) SetReleaseFileSize(v int64) *ListImagesResponseBodyImages {
	s.ReleaseFileSize = &v
	return s
}

func (s *ListImagesResponseBodyImages) SetType(v string) *ListImagesResponseBodyImages {
	s.Type = &v
	return s
}

func (s *ListImagesResponseBodyImages) Validate() error {
	return dara.Validate(s)
}

type iListImagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListImagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListImagesResponse
	GetStatusCode() *int32
	SetBody(v *ListImagesResponseBody) *ListImagesResponse
	GetBody() *ListImagesResponseBody
}

type ListImagesResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListImagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListImagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListImagesResponse) GoString() string {
	return s.String()
}

func (s *ListImagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListImagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListImagesResponse) GetBody() *ListImagesResponseBody {
	return s.Body
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

func (s *ListImagesResponse) Validate() error {
	return dara.Validate(s)
}

type iListMachineNetworkInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMachineHpnInfo(v []*ListMachineNetworkInfoRequestMachineHpnInfo) *ListMachineNetworkInfoRequest
	GetMachineHpnInfo() []*ListMachineNetworkInfoRequestMachineHpnInfo
}

type ListMachineNetworkInfoRequest struct {
	// Array
	MachineHpnInfo []*ListMachineNetworkInfoRequestMachineHpnInfo `json:"MachineHpnInfo,omitempty" xml:"MachineHpnInfo,omitempty" type:"Repeated"`
}

func (s ListMachineNetworkInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoRequest) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoRequest) GetMachineHpnInfo() []*ListMachineNetworkInfoRequestMachineHpnInfo {
	return s.MachineHpnInfo
}

func (s *ListMachineNetworkInfoRequest) SetMachineHpnInfo(v []*ListMachineNetworkInfoRequestMachineHpnInfo) *ListMachineNetworkInfoRequest {
	s.MachineHpnInfo = v
	return s
}

func (s *ListMachineNetworkInfoRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoRequestMachineHpnInfo) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) GetMachineType() *string {
	return s.MachineType
}

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) GetRegionId() *string {
	return s.RegionId
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

func (s *ListMachineNetworkInfoRequestMachineHpnInfo) Validate() error {
	return dara.Validate(s)
}

type iListMachineNetworkInfoShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMachineHpnInfoShrink(v string) *ListMachineNetworkInfoShrinkRequest
	GetMachineHpnInfoShrink() *string
}

type ListMachineNetworkInfoShrinkRequest struct {
	// Array
	MachineHpnInfoShrink *string `json:"MachineHpnInfo,omitempty" xml:"MachineHpnInfo,omitempty"`
}

func (s ListMachineNetworkInfoShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoShrinkRequest) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoShrinkRequest) GetMachineHpnInfoShrink() *string {
	return s.MachineHpnInfoShrink
}

func (s *ListMachineNetworkInfoShrinkRequest) SetMachineHpnInfoShrink(v string) *ListMachineNetworkInfoShrinkRequest {
	s.MachineHpnInfoShrink = &v
	return s
}

func (s *ListMachineNetworkInfoShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iListMachineNetworkInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMachineNetworkInfo(v []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo) *ListMachineNetworkInfoResponseBody
	GetMachineNetworkInfo() []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo
	SetRequestId(v string) *ListMachineNetworkInfoResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponseBody) GetMachineNetworkInfo() []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo {
	return s.MachineNetworkInfo
}

func (s *ListMachineNetworkInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListMachineNetworkInfoResponseBody) SetMachineNetworkInfo(v []*ListMachineNetworkInfoResponseBodyMachineNetworkInfo) *ListMachineNetworkInfoResponseBody {
	s.MachineNetworkInfo = v
	return s
}

func (s *ListMachineNetworkInfoResponseBody) SetRequestId(v string) *ListMachineNetworkInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMachineNetworkInfoResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetClusterNet() *string {
	return s.ClusterNet
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetIsDpuMode() *bool {
	return s.IsDpuMode
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetMachineType() *string {
	return s.MachineType
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetNetArch() *string {
	return s.NetArch
}

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) GetRegionId() *string {
	return s.RegionId
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

func (s *ListMachineNetworkInfoResponseBodyMachineNetworkInfo) Validate() error {
	return dara.Validate(s)
}

type iListMachineNetworkInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMachineNetworkInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMachineNetworkInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListMachineNetworkInfoResponseBody) *ListMachineNetworkInfoResponse
	GetBody() *ListMachineNetworkInfoResponseBody
}

type ListMachineNetworkInfoResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachineNetworkInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMachineNetworkInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMachineNetworkInfoResponse) GoString() string {
	return s.String()
}

func (s *ListMachineNetworkInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMachineNetworkInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMachineNetworkInfoResponse) GetBody() *ListMachineNetworkInfoResponseBody {
	return s.Body
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

func (s *ListMachineNetworkInfoResponse) Validate() error {
	return dara.Validate(s)
}

type iListMachineTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *ListMachineTypesRequest
	GetName() *string
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
	return dara.Prettify(s)
}

func (s ListMachineTypesRequest) GoString() string {
	return s.String()
}

func (s *ListMachineTypesRequest) GetName() *string {
	return s.Name
}

func (s *ListMachineTypesRequest) SetName(v string) *ListMachineTypesRequest {
	s.Name = &v
	return s
}

func (s *ListMachineTypesRequest) Validate() error {
	return dara.Validate(s)
}

type iListMachineTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMachineTypes(v []*ListMachineTypesResponseBodyMachineTypes) *ListMachineTypesResponseBody
	GetMachineTypes() []*ListMachineTypesResponseBodyMachineTypes
	SetNextToken(v string) *ListMachineTypesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListMachineTypesResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListMachineTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBody) GetMachineTypes() []*ListMachineTypesResponseBodyMachineTypes {
	return s.MachineTypes
}

func (s *ListMachineTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListMachineTypesResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *ListMachineTypesResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListMachineTypesResponseBodyMachineTypes) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetBondNum() *int32 {
	return s.BondNum
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetCpuInfo() *string {
	return s.CpuInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetDiskInfo() *string {
	return s.DiskInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetGpuInfo() *string {
	return s.GpuInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetMemoryInfo() *string {
	return s.MemoryInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetName() *string {
	return s.Name
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetNetworkInfo() *string {
	return s.NetworkInfo
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetNodeCount() *string {
	return s.NodeCount
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetTotalCpuCore() *int32 {
	return s.TotalCpuCore
}

func (s *ListMachineTypesResponseBodyMachineTypes) GetType() *string {
	return s.Type
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

func (s *ListMachineTypesResponseBodyMachineTypes) Validate() error {
	return dara.Validate(s)
}

type iListMachineTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMachineTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMachineTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListMachineTypesResponseBody) *ListMachineTypesResponse
	GetBody() *ListMachineTypesResponseBody
}

type ListMachineTypesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMachineTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMachineTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMachineTypesResponse) GoString() string {
	return s.String()
}

func (s *ListMachineTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMachineTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMachineTypesResponse) GetBody() *ListMachineTypesResponseBody {
	return s.Body
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

func (s *ListMachineTypesResponse) Validate() error {
	return dara.Validate(s)
}

type iListNetTestResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListNetTestResultsRequest
	GetMaxResults() *int64
	SetNetTestType(v string) *ListNetTestResultsRequest
	GetNetTestType() *string
	SetNextToken(v string) *ListNetTestResultsRequest
	GetNextToken() *string
	SetResourceGroupId(v string) *ListNetTestResultsRequest
	GetResourceGroupId() *string
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
	return dara.Prettify(s)
}

func (s ListNetTestResultsRequest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsRequest) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListNetTestResultsRequest) GetNetTestType() *string {
	return s.NetTestType
}

func (s *ListNetTestResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNetTestResultsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
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

func (s *ListNetTestResultsRequest) Validate() error {
	return dara.Validate(s)
}

type iListNetTestResultsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int64) *ListNetTestResultsResponseBody
	GetMaxResults() *int64
	SetNetTestResults(v []*ListNetTestResultsResponseBodyNetTestResults) *ListNetTestResultsResponseBody
	GetNetTestResults() []*ListNetTestResultsResponseBodyNetTestResults
	SetNextToken(v string) *ListNetTestResultsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListNetTestResultsResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBody) GetMaxResults() *int64 {
	return s.MaxResults
}

func (s *ListNetTestResultsResponseBody) GetNetTestResults() []*ListNetTestResultsResponseBodyNetTestResults {
	return s.NetTestResults
}

func (s *ListNetTestResultsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNetTestResultsResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *ListNetTestResultsResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResults) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetCommTest() *ListNetTestResultsResponseBodyNetTestResultsCommTest {
	return s.CommTest
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetDelayTest() *ListNetTestResultsResponseBodyNetTestResultsDelayTest {
	return s.DelayTest
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetFinishedTime() *string {
	return s.FinishedTime
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetNetTestType() *string {
	return s.NetTestType
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetNetworkMode() *string {
	return s.NetworkMode
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetPort() *string {
	return s.Port
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetStatus() *string {
	return s.Status
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetTestId() *string {
	return s.TestId
}

func (s *ListNetTestResultsResponseBodyNetTestResults) GetTrafficTest() *ListNetTestResultsResponseBodyNetTestResultsTrafficTest {
	return s.TrafficTest
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

func (s *ListNetTestResultsResponseBodyNetTestResults) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) GetGPUNum() *string {
	return s.GPUNum
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) GetHosts() []*ListNetTestResultsResponseBodyNetTestResultsCommTestHosts {
	return s.Hosts
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) GetModel() *string {
	return s.Model
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) GetType() *string {
	return s.Type
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

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) GetIP() *string {
	return s.IP
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) GetServerName() *string {
	return s.ServerName
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

func (s *ListNetTestResultsResponseBodyNetTestResultsCommTestHosts) Validate() error {
	return dara.Validate(s)
}

type ListNetTestResultsResponseBodyNetTestResultsDelayTest struct {
	// Resource list
	Hosts []*ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts `json:"Hosts,omitempty" xml:"Hosts,omitempty" type:"Repeated"`
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTest) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTest) GetHosts() []*ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts {
	return s.Hosts
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTest) SetHosts(v []*ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) *ListNetTestResultsResponseBodyNetTestResultsDelayTest {
	s.Hosts = v
	return s
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GetBond() *string {
	return s.Bond
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GetIP() *string {
	return s.IP
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) GetServerName() *string {
	return s.ServerName
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

func (s *ListNetTestResultsResponseBodyNetTestResultsDelayTestHosts) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetClients() []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients {
	return s.Clients
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetDuration() *int64 {
	return s.Duration
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetGDR() *string {
	return s.GDR
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetProtocol() *string {
	return s.Protocol
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetQP() *int64 {
	return s.QP
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetServers() []*ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers {
	return s.Servers
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) GetTrafficModel() *string {
	return s.TrafficModel
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

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GetBond() *string {
	return s.Bond
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GetIP() *string {
	return s.IP
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) GetServerName() *string {
	return s.ServerName
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

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestClients) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GetBond() *string {
	return s.Bond
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GetIP() *string {
	return s.IP
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) GetServerName() *string {
	return s.ServerName
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

func (s *ListNetTestResultsResponseBodyNetTestResultsTrafficTestServers) Validate() error {
	return dara.Validate(s)
}

type iListNetTestResultsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNetTestResultsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNetTestResultsResponse
	GetStatusCode() *int32
	SetBody(v *ListNetTestResultsResponseBody) *ListNetTestResultsResponse
	GetBody() *ListNetTestResultsResponseBody
}

type ListNetTestResultsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNetTestResultsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNetTestResultsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNetTestResultsResponse) GoString() string {
	return s.String()
}

func (s *ListNetTestResultsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNetTestResultsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNetTestResultsResponse) GetBody() *ListNetTestResultsResponseBody {
	return s.Body
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

func (s *ListNetTestResultsResponse) Validate() error {
	return dara.Validate(s)
}

type iListNodeGroupsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ListNodeGroupsRequest
	GetClusterId() *string
	SetMaxResults(v int32) *ListNodeGroupsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListNodeGroupsRequest
	GetNextToken() *string
	SetNodeGroupId(v string) *ListNodeGroupsRequest
	GetNodeGroupId() *string
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
	return dara.Prettify(s)
}

func (s ListNodeGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodeGroupsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListNodeGroupsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodeGroupsRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
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

func (s *ListNodeGroupsRequest) Validate() error {
	return dara.Validate(s)
}

type iListNodeGroupsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroups(v []*ListNodeGroupsResponseBodyGroups) *ListNodeGroupsResponseBody
	GetGroups() []*ListNodeGroupsResponseBodyGroups
	SetNextToken(v string) *ListNodeGroupsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListNodeGroupsResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListNodeGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBody) GetGroups() []*ListNodeGroupsResponseBodyGroups {
	return s.Groups
}

func (s *ListNodeGroupsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListNodeGroupsResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *ListNodeGroupsResponseBody) Validate() error {
	return dara.Validate(s)
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
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
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
	return dara.Prettify(s)
}

func (s ListNodeGroupsResponseBodyGroups) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponseBodyGroups) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListNodeGroupsResponseBodyGroups) GetClusterName() *string {
	return s.ClusterName
}

func (s *ListNodeGroupsResponseBodyGroups) GetCreateTime() *string {
	return s.CreateTime
}

func (s *ListNodeGroupsResponseBodyGroups) GetDescription() *string {
	return s.Description
}

func (s *ListNodeGroupsResponseBodyGroups) GetGroupId() *string {
	return s.GroupId
}

func (s *ListNodeGroupsResponseBodyGroups) GetGroupName() *string {
	return s.GroupName
}

func (s *ListNodeGroupsResponseBodyGroups) GetImageId() *string {
	return s.ImageId
}

func (s *ListNodeGroupsResponseBodyGroups) GetImageName() *string {
	return s.ImageName
}

func (s *ListNodeGroupsResponseBodyGroups) GetMachineType() *string {
	return s.MachineType
}

func (s *ListNodeGroupsResponseBodyGroups) GetNodeCount() *int64 {
	return s.NodeCount
}

func (s *ListNodeGroupsResponseBodyGroups) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *ListNodeGroupsResponseBodyGroups) GetZoneId() *string {
	return s.ZoneId
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

func (s *ListNodeGroupsResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}

type iListNodeGroupsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListNodeGroupsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListNodeGroupsResponse
	GetStatusCode() *int32
	SetBody(v *ListNodeGroupsResponseBody) *ListNodeGroupsResponse
	GetBody() *ListNodeGroupsResponseBody
}

type ListNodeGroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListNodeGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListNodeGroupsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListNodeGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListNodeGroupsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListNodeGroupsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListNodeGroupsResponse) GetBody() *ListNodeGroupsResponseBody {
	return s.Body
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

func (s *ListNodeGroupsResponse) Validate() error {
	return dara.Validate(s)
}

type iListSyslogsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFromTime(v string) *ListSyslogsRequest
	GetFromTime() *string
	SetNextToken(v string) *ListSyslogsRequest
	GetNextToken() *string
	SetNodeId(v string) *ListSyslogsRequest
	GetNodeId() *string
	SetQuery(v string) *ListSyslogsRequest
	GetQuery() *string
	SetReverse(v bool) *ListSyslogsRequest
	GetReverse() *bool
	SetToTime(v string) *ListSyslogsRequest
	GetToTime() *string
}

type ListSyslogsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1659745800
	FromTime *string `json:"FromTime,omitempty" xml:"FromTime,omitempty"`
	// example:
	//
	// 392e8b4a03ed171433cc39f5b464ec9d
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// e01-cn-nwy37atbj44
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// *
	Query *string `json:"Query,omitempty" xml:"Query,omitempty"`
	// example:
	//
	// true
	Reverse *bool `json:"Reverse,omitempty" xml:"Reverse,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1665369329
	ToTime *string `json:"ToTime,omitempty" xml:"ToTime,omitempty"`
}

func (s ListSyslogsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListSyslogsRequest) GoString() string {
	return s.String()
}

func (s *ListSyslogsRequest) GetFromTime() *string {
	return s.FromTime
}

func (s *ListSyslogsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSyslogsRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *ListSyslogsRequest) GetQuery() *string {
	return s.Query
}

func (s *ListSyslogsRequest) GetReverse() *bool {
	return s.Reverse
}

func (s *ListSyslogsRequest) GetToTime() *string {
	return s.ToTime
}

func (s *ListSyslogsRequest) SetFromTime(v string) *ListSyslogsRequest {
	s.FromTime = &v
	return s
}

func (s *ListSyslogsRequest) SetNextToken(v string) *ListSyslogsRequest {
	s.NextToken = &v
	return s
}

func (s *ListSyslogsRequest) SetNodeId(v string) *ListSyslogsRequest {
	s.NodeId = &v
	return s
}

func (s *ListSyslogsRequest) SetQuery(v string) *ListSyslogsRequest {
	s.Query = &v
	return s
}

func (s *ListSyslogsRequest) SetReverse(v bool) *ListSyslogsRequest {
	s.Reverse = &v
	return s
}

func (s *ListSyslogsRequest) SetToTime(v string) *ListSyslogsRequest {
	s.ToTime = &v
	return s
}

func (s *ListSyslogsRequest) Validate() error {
	return dara.Validate(s)
}

type iListSyslogsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogs(v []*ListSyslogsResponseBodyLogs) *ListSyslogsResponseBody
	GetLogs() []*ListSyslogsResponseBodyLogs
	SetNextToken(v string) *ListSyslogsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListSyslogsResponseBody
	GetRequestId() *string
}

type ListSyslogsResponseBody struct {
	Logs []*ListSyslogsResponseBodyLogs `json:"Logs,omitempty" xml:"Logs,omitempty" type:"Repeated"`
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 2FE2B22C-CF9D-59DE-BF63-DC9B9B33A9D1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListSyslogsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListSyslogsResponseBody) GoString() string {
	return s.String()
}

func (s *ListSyslogsResponseBody) GetLogs() []*ListSyslogsResponseBodyLogs {
	return s.Logs
}

func (s *ListSyslogsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListSyslogsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListSyslogsResponseBody) SetLogs(v []*ListSyslogsResponseBodyLogs) *ListSyslogsResponseBody {
	s.Logs = v
	return s
}

func (s *ListSyslogsResponseBody) SetNextToken(v string) *ListSyslogsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListSyslogsResponseBody) SetRequestId(v string) *ListSyslogsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListSyslogsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListSyslogsResponseBodyLogs struct {
	// example:
	//
	// i119583961673208491760
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// ALIYUN_PUBLIC
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// kern
	Facility *string `json:"Facility,omitempty" xml:"Facility,omitempty"`
	// example:
	//
	// damo-m53kr8kd-0008
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// example:
	//
	// 114.55.254.44
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// example:
	//
	// bond4: failed to get link speed/duplex for eth8
	Msg *string `json:"Msg,omitempty" xml:"Msg,omitempty"`
	// example:
	//
	// e01-cn-9lb36u4s601
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// warning
	Severity *string `json:"Severity,omitempty" xml:"Severity,omitempty"`
	// example:
	//
	// 21A401332
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// example:
	//
	// damo-m53kr8kd-0008
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// example:
	//
	// kernel
	Syslogtag *string `json:"Syslogtag,omitempty" xml:"Syslogtag,omitempty"`
	// example:
	//
	// damo-m53kr8kd-0008
	TagHostname *string `json:"TagHostname,omitempty" xml:"TagHostname,omitempty"`
	// example:
	//
	// D990314D3C25D7E8-1080
	TagPackId *string `json:"TagPackId,omitempty" xml:"TagPackId,omitempty"`
	// example:
	//
	// /var/log/kern
	TagPath *string `json:"TagPath,omitempty" xml:"TagPath,omitempty"`
	// example:
	//
	// 1687363348
	TagReceiveTime *string `json:"TagReceiveTime,omitempty" xml:"TagReceiveTime,omitempty"`
	// example:
	//
	// application_b
	TagUserDefinedId *string `json:"TagUserDefinedId,omitempty" xml:"TagUserDefinedId,omitempty"`
	// example:
	//
	// 1687363346
	Time *string `json:"Time,omitempty" xml:"Time,omitempty"`
	// example:
	//
	// logserver
	Topic *string `json:"Topic,omitempty" xml:"Topic,omitempty"`
}

func (s ListSyslogsResponseBodyLogs) String() string {
	return dara.Prettify(s)
}

func (s ListSyslogsResponseBodyLogs) GoString() string {
	return s.String()
}

func (s *ListSyslogsResponseBodyLogs) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListSyslogsResponseBodyLogs) GetDomain() *string {
	return s.Domain
}

func (s *ListSyslogsResponseBodyLogs) GetFacility() *string {
	return s.Facility
}

func (s *ListSyslogsResponseBodyLogs) GetHostname() *string {
	return s.Hostname
}

func (s *ListSyslogsResponseBodyLogs) GetIp() *string {
	return s.Ip
}

func (s *ListSyslogsResponseBodyLogs) GetMsg() *string {
	return s.Msg
}

func (s *ListSyslogsResponseBodyLogs) GetNodeId() *string {
	return s.NodeId
}

func (s *ListSyslogsResponseBodyLogs) GetSeverity() *string {
	return s.Severity
}

func (s *ListSyslogsResponseBodyLogs) GetSn() *string {
	return s.Sn
}

func (s *ListSyslogsResponseBodyLogs) GetSource() *string {
	return s.Source
}

func (s *ListSyslogsResponseBodyLogs) GetSyslogtag() *string {
	return s.Syslogtag
}

func (s *ListSyslogsResponseBodyLogs) GetTagHostname() *string {
	return s.TagHostname
}

func (s *ListSyslogsResponseBodyLogs) GetTagPackId() *string {
	return s.TagPackId
}

func (s *ListSyslogsResponseBodyLogs) GetTagPath() *string {
	return s.TagPath
}

func (s *ListSyslogsResponseBodyLogs) GetTagReceiveTime() *string {
	return s.TagReceiveTime
}

func (s *ListSyslogsResponseBodyLogs) GetTagUserDefinedId() *string {
	return s.TagUserDefinedId
}

func (s *ListSyslogsResponseBodyLogs) GetTime() *string {
	return s.Time
}

func (s *ListSyslogsResponseBodyLogs) GetTopic() *string {
	return s.Topic
}

func (s *ListSyslogsResponseBodyLogs) SetClusterId(v string) *ListSyslogsResponseBodyLogs {
	s.ClusterId = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetDomain(v string) *ListSyslogsResponseBodyLogs {
	s.Domain = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetFacility(v string) *ListSyslogsResponseBodyLogs {
	s.Facility = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetHostname(v string) *ListSyslogsResponseBodyLogs {
	s.Hostname = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetIp(v string) *ListSyslogsResponseBodyLogs {
	s.Ip = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetMsg(v string) *ListSyslogsResponseBodyLogs {
	s.Msg = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetNodeId(v string) *ListSyslogsResponseBodyLogs {
	s.NodeId = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetSeverity(v string) *ListSyslogsResponseBodyLogs {
	s.Severity = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetSn(v string) *ListSyslogsResponseBodyLogs {
	s.Sn = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetSource(v string) *ListSyslogsResponseBodyLogs {
	s.Source = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetSyslogtag(v string) *ListSyslogsResponseBodyLogs {
	s.Syslogtag = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTagHostname(v string) *ListSyslogsResponseBodyLogs {
	s.TagHostname = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTagPackId(v string) *ListSyslogsResponseBodyLogs {
	s.TagPackId = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTagPath(v string) *ListSyslogsResponseBodyLogs {
	s.TagPath = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTagReceiveTime(v string) *ListSyslogsResponseBodyLogs {
	s.TagReceiveTime = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTagUserDefinedId(v string) *ListSyslogsResponseBodyLogs {
	s.TagUserDefinedId = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTime(v string) *ListSyslogsResponseBodyLogs {
	s.Time = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) SetTopic(v string) *ListSyslogsResponseBodyLogs {
	s.Topic = &v
	return s
}

func (s *ListSyslogsResponseBodyLogs) Validate() error {
	return dara.Validate(s)
}

type iListSyslogsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListSyslogsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListSyslogsResponse
	GetStatusCode() *int32
	SetBody(v *ListSyslogsResponseBody) *ListSyslogsResponse
	GetBody() *ListSyslogsResponseBody
}

type ListSyslogsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListSyslogsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListSyslogsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListSyslogsResponse) GoString() string {
	return s.String()
}

func (s *ListSyslogsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListSyslogsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListSyslogsResponse) GetBody() *ListSyslogsResponseBody {
	return s.Body
}

func (s *ListSyslogsResponse) SetHeaders(v map[string]*string) *ListSyslogsResponse {
	s.Headers = v
	return s
}

func (s *ListSyslogsResponse) SetStatusCode(v int32) *ListSyslogsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListSyslogsResponse) SetBody(v *ListSyslogsResponseBody) *ListSyslogsResponse {
	s.Body = v
	return s
}

func (s *ListSyslogsResponse) Validate() error {
	return dara.Validate(s)
}

type iListTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesRequest
	GetNextToken() *string
	SetRegionId(v string) *ListTagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *ListTagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *ListTagResourcesRequest
	GetResourceType() *string
	SetService(v string) *ListTagResourcesRequest
	GetService() *string
	SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest
	GetTag() []*ListTagResourcesRequestTag
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
	Service      *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// List of tags
	Tag []*ListTagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s ListTagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesRequest) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListTagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *ListTagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesRequest) GetService() *string {
	return s.Service
}

func (s *ListTagResourcesRequest) GetTag() []*ListTagResourcesRequestTag {
	return s.Tag
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

func (s *ListTagResourcesRequest) SetService(v string) *ListTagResourcesRequest {
	s.Service = &v
	return s
}

func (s *ListTagResourcesRequest) SetTag(v []*ListTagResourcesRequestTag) *ListTagResourcesRequest {
	s.Tag = v
	return s
}

func (s *ListTagResourcesRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListTagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *ListTagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListTagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListTagResourcesRequestTag) SetKey(v string) *ListTagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *ListTagResourcesRequestTag) SetValue(v string) *ListTagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *ListTagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}

type iListTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *ListTagResourcesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListTagResourcesResponseBody
	GetRequestId() *string
	SetTagResources(v *ListTagResourcesResponseBodyTagResources) *ListTagResourcesResponseBody
	GetTagResources() *ListTagResourcesResponseBodyTagResources
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
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListTagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTagResourcesResponseBody) GetTagResources() *ListTagResourcesResponseBodyTagResources {
	return s.TagResources
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

func (s *ListTagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListTagResourcesResponseBodyTagResources struct {
	TagResource []*ListTagResourcesResponseBodyTagResourcesTagResource `json:"TagResource,omitempty" xml:"TagResource,omitempty" type:"Repeated"`
}

func (s ListTagResourcesResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResources) GetTagResource() []*ListTagResourcesResponseBodyTagResourcesTagResource {
	return s.TagResource
}

func (s *ListTagResourcesResponseBodyTagResources) SetTagResource(v []*ListTagResourcesResponseBodyTagResourcesTagResource) *ListTagResourcesResponseBodyTagResources {
	s.TagResource = v
	return s
}

func (s *ListTagResourcesResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListTagResourcesResponseBodyTagResourcesTagResource) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetResourceId() *string {
	return s.ResourceId
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetResourceType() *string {
	return s.ResourceType
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetTagKey() *string {
	return s.TagKey
}

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) GetTagValue() *string {
	return s.TagValue
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

func (s *ListTagResourcesResponseBodyTagResourcesTagResource) Validate() error {
	return dara.Validate(s)
}

type iListTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListTagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListTagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *ListTagResourcesResponseBody) *ListTagResourcesResponse
	GetBody() *ListTagResourcesResponseBody
}

type ListTagResourcesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListTagResourcesResponse) GoString() string {
	return s.String()
}

func (s *ListTagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListTagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListTagResourcesResponse) GetBody() *ListTagResourcesResponseBody {
	return s.Body
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

func (s *ListTagResourcesResponse) Validate() error {
	return dara.Validate(s)
}

type iListUserClusterTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetClusterTypes(v []*ListUserClusterTypesResponseBodyClusterTypes) *ListUserClusterTypesResponseBody
	GetClusterTypes() []*ListUserClusterTypesResponseBodyClusterTypes
	SetNextToken(v string) *ListUserClusterTypesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListUserClusterTypesResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s ListUserClusterTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponseBody) GetClusterTypes() []*ListUserClusterTypesResponseBodyClusterTypes {
	return s.ClusterTypes
}

func (s *ListUserClusterTypesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListUserClusterTypesResponseBody) GetRequestId() *string {
	return s.RequestId
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

func (s *ListUserClusterTypesResponseBody) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ListUserClusterTypesResponseBodyClusterTypes) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) GetAccessType() *string {
	return s.AccessType
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) GetTypeName() *string {
	return s.TypeName
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) SetAccessType(v string) *ListUserClusterTypesResponseBodyClusterTypes {
	s.AccessType = &v
	return s
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) SetTypeName(v string) *ListUserClusterTypesResponseBodyClusterTypes {
	s.TypeName = &v
	return s
}

func (s *ListUserClusterTypesResponseBodyClusterTypes) Validate() error {
	return dara.Validate(s)
}

type iListUserClusterTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserClusterTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserClusterTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListUserClusterTypesResponseBody) *ListUserClusterTypesResponse
	GetBody() *ListUserClusterTypesResponseBody
}

type ListUserClusterTypesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserClusterTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserClusterTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserClusterTypesResponse) GoString() string {
	return s.String()
}

func (s *ListUserClusterTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserClusterTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserClusterTypesResponse) GetBody() *ListUserClusterTypesResponseBody {
	return s.Body
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

func (s *ListUserClusterTypesResponse) Validate() error {
	return dara.Validate(s)
}

type iListVscsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVscsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *ListVscsRequest
	GetNextToken() *string
	SetNodeIds(v []*string) *ListVscsRequest
	GetNodeIds() []*string
	SetResourceGroupId(v string) *ListVscsRequest
	GetResourceGroupId() *string
	SetTag(v []*ListVscsRequestTag) *ListVscsRequest
	GetTag() []*ListVscsRequestTag
	SetVscName(v string) *ListVscsRequest
	GetVscName() *string
}

type ListVscsRequest struct {
	MaxResults      *int32                `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	NextToken       *string               `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	NodeIds         []*string             `json:"NodeIds,omitempty" xml:"NodeIds,omitempty" type:"Repeated"`
	ResourceGroupId *string               `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Tag             []*ListVscsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VscName         *string               `json:"VscName,omitempty" xml:"VscName,omitempty"`
}

func (s ListVscsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListVscsRequest) GoString() string {
	return s.String()
}

func (s *ListVscsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVscsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVscsRequest) GetNodeIds() []*string {
	return s.NodeIds
}

func (s *ListVscsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVscsRequest) GetTag() []*ListVscsRequestTag {
	return s.Tag
}

func (s *ListVscsRequest) GetVscName() *string {
	return s.VscName
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

func (s *ListVscsRequest) Validate() error {
	return dara.Validate(s)
}

type ListVscsRequestTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListVscsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s ListVscsRequestTag) GoString() string {
	return s.String()
}

func (s *ListVscsRequestTag) GetKey() *string {
	return s.Key
}

func (s *ListVscsRequestTag) GetValue() *string {
	return s.Value
}

func (s *ListVscsRequestTag) SetKey(v string) *ListVscsRequestTag {
	s.Key = &v
	return s
}

func (s *ListVscsRequestTag) SetValue(v string) *ListVscsRequestTag {
	s.Value = &v
	return s
}

func (s *ListVscsRequestTag) Validate() error {
	return dara.Validate(s)
}

type iListVscsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListVscsResponseBody
	GetMaxResults() *int32
	SetNextToken(v string) *ListVscsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListVscsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListVscsResponseBody
	GetTotalCount() *int32
	SetVscs(v []*ListVscsResponseBodyVscs) *ListVscsResponseBody
	GetVscs() []*ListVscsResponseBodyVscs
}

type ListVscsResponseBody struct {
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// This parameter is required.
	NextToken  *string                     `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId  *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vscs       []*ListVscsResponseBodyVscs `json:"Vscs,omitempty" xml:"Vscs,omitempty" type:"Repeated"`
}

func (s ListVscsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListVscsResponseBody) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListVscsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListVscsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListVscsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListVscsResponseBody) GetVscs() []*ListVscsResponseBodyVscs {
	return s.Vscs
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

func (s *ListVscsResponseBody) Validate() error {
	return dara.Validate(s)
}

type ListVscsResponseBodyVscs struct {
	NodeId          *string                         `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	ResourceGroupId *string                         `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	Status          *string                         `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags            []*ListVscsResponseBodyVscsTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	VscId           *string                         `json:"VscId,omitempty" xml:"VscId,omitempty"`
	VscName         *string                         `json:"VscName,omitempty" xml:"VscName,omitempty"`
	VscType         *string                         `json:"VscType,omitempty" xml:"VscType,omitempty"`
}

func (s ListVscsResponseBodyVscs) String() string {
	return dara.Prettify(s)
}

func (s ListVscsResponseBodyVscs) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBodyVscs) GetNodeId() *string {
	return s.NodeId
}

func (s *ListVscsResponseBodyVscs) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ListVscsResponseBodyVscs) GetStatus() *string {
	return s.Status
}

func (s *ListVscsResponseBodyVscs) GetTags() []*ListVscsResponseBodyVscsTags {
	return s.Tags
}

func (s *ListVscsResponseBodyVscs) GetVscId() *string {
	return s.VscId
}

func (s *ListVscsResponseBodyVscs) GetVscName() *string {
	return s.VscName
}

func (s *ListVscsResponseBodyVscs) GetVscType() *string {
	return s.VscType
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

func (s *ListVscsResponseBodyVscs) Validate() error {
	return dara.Validate(s)
}

type ListVscsResponseBodyVscsTags struct {
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s ListVscsResponseBodyVscsTags) String() string {
	return dara.Prettify(s)
}

func (s ListVscsResponseBodyVscsTags) GoString() string {
	return s.String()
}

func (s *ListVscsResponseBodyVscsTags) GetTagKey() *string {
	return s.TagKey
}

func (s *ListVscsResponseBodyVscsTags) GetTagValue() *string {
	return s.TagValue
}

func (s *ListVscsResponseBodyVscsTags) SetTagKey(v string) *ListVscsResponseBodyVscsTags {
	s.TagKey = &v
	return s
}

func (s *ListVscsResponseBodyVscsTags) SetTagValue(v string) *ListVscsResponseBodyVscsTags {
	s.TagValue = &v
	return s
}

func (s *ListVscsResponseBodyVscsTags) Validate() error {
	return dara.Validate(s)
}

type iListVscsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListVscsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListVscsResponse
	GetStatusCode() *int32
	SetBody(v *ListVscsResponseBody) *ListVscsResponse
	GetBody() *ListVscsResponseBody
}

type ListVscsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListVscsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListVscsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListVscsResponse) GoString() string {
	return s.String()
}

func (s *ListVscsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListVscsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListVscsResponse) GetBody() *ListVscsResponseBody {
	return s.Body
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

func (s *ListVscsResponse) Validate() error {
	return dara.Validate(s)
}

type iRebootNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RebootNodesRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *RebootNodesRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodes(v []*string) *RebootNodesRequest
	GetNodes() []*string
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
	return dara.Prettify(s)
}

func (s RebootNodesRequest) GoString() string {
	return s.String()
}

func (s *RebootNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RebootNodesRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *RebootNodesRequest) GetNodes() []*string {
	return s.Nodes
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

func (s *RebootNodesRequest) Validate() error {
	return dara.Validate(s)
}

type iRebootNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *RebootNodesShrinkRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *RebootNodesShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodesShrink(v string) *RebootNodesShrinkRequest
	GetNodesShrink() *string
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
	return dara.Prettify(s)
}

func (s RebootNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *RebootNodesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *RebootNodesShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *RebootNodesShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
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

func (s *RebootNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iRebootNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RebootNodesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RebootNodesResponseBody
	GetTaskId() *string
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
	return dara.Prettify(s)
}

func (s RebootNodesResponseBody) GoString() string {
	return s.String()
}

func (s *RebootNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RebootNodesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RebootNodesResponseBody) SetRequestId(v string) *RebootNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *RebootNodesResponseBody) SetTaskId(v string) *RebootNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *RebootNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type iRebootNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RebootNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RebootNodesResponse
	GetStatusCode() *int32
	SetBody(v *RebootNodesResponseBody) *RebootNodesResponse
	GetBody() *RebootNodesResponseBody
}

type RebootNodesResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RebootNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RebootNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s RebootNodesResponse) GoString() string {
	return s.String()
}

func (s *RebootNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RebootNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RebootNodesResponse) GetBody() *RebootNodesResponseBody {
	return s.Body
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

func (s *RebootNodesResponse) Validate() error {
	return dara.Validate(s)
}

type iReimageNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ReimageNodesRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *ReimageNodesRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodes(v []*ReimageNodesRequestNodes) *ReimageNodesRequest
	GetNodes() []*ReimageNodesRequestNodes
	SetUserData(v string) *ReimageNodesRequest
	GetUserData() *string
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
	return dara.Prettify(s)
}

func (s ReimageNodesRequest) GoString() string {
	return s.String()
}

func (s *ReimageNodesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ReimageNodesRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ReimageNodesRequest) GetNodes() []*ReimageNodesRequestNodes {
	return s.Nodes
}

func (s *ReimageNodesRequest) GetUserData() *string {
	return s.UserData
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

func (s *ReimageNodesRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ReimageNodesRequestNodes) GoString() string {
	return s.String()
}

func (s *ReimageNodesRequestNodes) GetHostname() *string {
	return s.Hostname
}

func (s *ReimageNodesRequestNodes) GetImageId() *string {
	return s.ImageId
}

func (s *ReimageNodesRequestNodes) GetLoginPassword() *string {
	return s.LoginPassword
}

func (s *ReimageNodesRequestNodes) GetNodeId() *string {
	return s.NodeId
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

func (s *ReimageNodesRequestNodes) Validate() error {
	return dara.Validate(s)
}

type iReimageNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ReimageNodesShrinkRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *ReimageNodesShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodesShrink(v string) *ReimageNodesShrinkRequest
	GetNodesShrink() *string
	SetUserData(v string) *ReimageNodesShrinkRequest
	GetUserData() *string
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
	return dara.Prettify(s)
}

func (s ReimageNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReimageNodesShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ReimageNodesShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ReimageNodesShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *ReimageNodesShrinkRequest) GetUserData() *string {
	return s.UserData
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

func (s *ReimageNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iReimageNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReimageNodesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ReimageNodesResponseBody
	GetTaskId() *string
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
	return dara.Prettify(s)
}

func (s ReimageNodesResponseBody) GoString() string {
	return s.String()
}

func (s *ReimageNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReimageNodesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ReimageNodesResponseBody) SetRequestId(v string) *ReimageNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReimageNodesResponseBody) SetTaskId(v string) *ReimageNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *ReimageNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type iReimageNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReimageNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReimageNodesResponse
	GetStatusCode() *int32
	SetBody(v *ReimageNodesResponseBody) *ReimageNodesResponse
	GetBody() *ReimageNodesResponseBody
}

type ReimageNodesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReimageNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReimageNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s ReimageNodesResponse) GoString() string {
	return s.String()
}

func (s *ReimageNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReimageNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReimageNodesResponse) GetBody() *ReimageNodesResponseBody {
	return s.Body
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

func (s *ReimageNodesResponse) Validate() error {
	return dara.Validate(s)
}

type iReportNodesStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ReportNodesStatusRequest
	GetDescription() *string
	SetEndTime(v string) *ReportNodesStatusRequest
	GetEndTime() *string
	SetIssueCategory(v string) *ReportNodesStatusRequest
	GetIssueCategory() *string
	SetNodes(v []*string) *ReportNodesStatusRequest
	GetNodes() []*string
	SetReason(v string) *ReportNodesStatusRequest
	GetReason() *string
	SetStartTime(v string) *ReportNodesStatusRequest
	GetStartTime() *string
}

type ReportNodesStatusRequest struct {
	// example:
	//
	// dwd_mysql_lingwan_faxing_login_di
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2024-07-10T10:17:06
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// hardware-disk-error
	IssueCategory *string   `json:"IssueCategory,omitempty" xml:"IssueCategory,omitempty"`
	Nodes         []*string `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// example:
	//
	// SoftwareError
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 2024-09-22T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ReportNodesStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportNodesStatusRequest) GoString() string {
	return s.String()
}

func (s *ReportNodesStatusRequest) GetDescription() *string {
	return s.Description
}

func (s *ReportNodesStatusRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ReportNodesStatusRequest) GetIssueCategory() *string {
	return s.IssueCategory
}

func (s *ReportNodesStatusRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *ReportNodesStatusRequest) GetReason() *string {
	return s.Reason
}

func (s *ReportNodesStatusRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ReportNodesStatusRequest) SetDescription(v string) *ReportNodesStatusRequest {
	s.Description = &v
	return s
}

func (s *ReportNodesStatusRequest) SetEndTime(v string) *ReportNodesStatusRequest {
	s.EndTime = &v
	return s
}

func (s *ReportNodesStatusRequest) SetIssueCategory(v string) *ReportNodesStatusRequest {
	s.IssueCategory = &v
	return s
}

func (s *ReportNodesStatusRequest) SetNodes(v []*string) *ReportNodesStatusRequest {
	s.Nodes = v
	return s
}

func (s *ReportNodesStatusRequest) SetReason(v string) *ReportNodesStatusRequest {
	s.Reason = &v
	return s
}

func (s *ReportNodesStatusRequest) SetStartTime(v string) *ReportNodesStatusRequest {
	s.StartTime = &v
	return s
}

func (s *ReportNodesStatusRequest) Validate() error {
	return dara.Validate(s)
}

type iReportNodesStatusShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *ReportNodesStatusShrinkRequest
	GetDescription() *string
	SetEndTime(v string) *ReportNodesStatusShrinkRequest
	GetEndTime() *string
	SetIssueCategory(v string) *ReportNodesStatusShrinkRequest
	GetIssueCategory() *string
	SetNodesShrink(v string) *ReportNodesStatusShrinkRequest
	GetNodesShrink() *string
	SetReason(v string) *ReportNodesStatusShrinkRequest
	GetReason() *string
	SetStartTime(v string) *ReportNodesStatusShrinkRequest
	GetStartTime() *string
}

type ReportNodesStatusShrinkRequest struct {
	// example:
	//
	// dwd_mysql_lingwan_faxing_login_di
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// 2024-07-10T10:17:06
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// example:
	//
	// hardware-disk-error
	IssueCategory *string `json:"IssueCategory,omitempty" xml:"IssueCategory,omitempty"`
	NodesShrink   *string `json:"Nodes,omitempty" xml:"Nodes,omitempty"`
	// example:
	//
	// SoftwareError
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// example:
	//
	// 2024-09-22T16:00:00Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ReportNodesStatusShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s ReportNodesStatusShrinkRequest) GoString() string {
	return s.String()
}

func (s *ReportNodesStatusShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *ReportNodesStatusShrinkRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ReportNodesStatusShrinkRequest) GetIssueCategory() *string {
	return s.IssueCategory
}

func (s *ReportNodesStatusShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *ReportNodesStatusShrinkRequest) GetReason() *string {
	return s.Reason
}

func (s *ReportNodesStatusShrinkRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ReportNodesStatusShrinkRequest) SetDescription(v string) *ReportNodesStatusShrinkRequest {
	s.Description = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) SetEndTime(v string) *ReportNodesStatusShrinkRequest {
	s.EndTime = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) SetIssueCategory(v string) *ReportNodesStatusShrinkRequest {
	s.IssueCategory = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) SetNodesShrink(v string) *ReportNodesStatusShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) SetReason(v string) *ReportNodesStatusShrinkRequest {
	s.Reason = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) SetStartTime(v string) *ReportNodesStatusShrinkRequest {
	s.StartTime = &v
	return s
}

func (s *ReportNodesStatusShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iReportNodesStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorMessage(v string) *ReportNodesStatusResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *ReportNodesStatusResponseBody
	GetRequestId() *string
}

type ReportNodesStatusResponseBody struct {
	// Error Message
	//
	// example:
	//
	// Resource not found
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// B0699629-14FC-51E7-B49E-AAD83F6FEB60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReportNodesStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReportNodesStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ReportNodesStatusResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *ReportNodesStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReportNodesStatusResponseBody) SetErrorMessage(v string) *ReportNodesStatusResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *ReportNodesStatusResponseBody) SetRequestId(v string) *ReportNodesStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReportNodesStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type iReportNodesStatusResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ReportNodesStatusResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ReportNodesStatusResponse
	GetStatusCode() *int32
	SetBody(v *ReportNodesStatusResponseBody) *ReportNodesStatusResponse
	GetBody() *ReportNodesStatusResponseBody
}

type ReportNodesStatusResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ReportNodesStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ReportNodesStatusResponse) String() string {
	return dara.Prettify(s)
}

func (s ReportNodesStatusResponse) GoString() string {
	return s.String()
}

func (s *ReportNodesStatusResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ReportNodesStatusResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ReportNodesStatusResponse) GetBody() *ReportNodesStatusResponseBody {
	return s.Body
}

func (s *ReportNodesStatusResponse) SetHeaders(v map[string]*string) *ReportNodesStatusResponse {
	s.Headers = v
	return s
}

func (s *ReportNodesStatusResponse) SetStatusCode(v int32) *ReportNodesStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *ReportNodesStatusResponse) SetBody(v *ReportNodesStatusResponseBody) *ReportNodesStatusResponse {
	s.Body = v
	return s
}

func (s *ReportNodesStatusResponse) Validate() error {
	return dara.Validate(s)
}

type iRunCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RunCommandRequest
	GetClientToken() *string
	SetCommandContent(v string) *RunCommandRequest
	GetCommandContent() *string
	SetCommandId(v string) *RunCommandRequest
	GetCommandId() *string
	SetContentEncoding(v string) *RunCommandRequest
	GetContentEncoding() *string
	SetDescription(v string) *RunCommandRequest
	GetDescription() *string
	SetEnableParameter(v bool) *RunCommandRequest
	GetEnableParameter() *bool
	SetFrequency(v string) *RunCommandRequest
	GetFrequency() *string
	SetLauncher(v string) *RunCommandRequest
	GetLauncher() *string
	SetName(v string) *RunCommandRequest
	GetName() *string
	SetNodeIdList(v []*string) *RunCommandRequest
	GetNodeIdList() []*string
	SetParameters(v map[string]interface{}) *RunCommandRequest
	GetParameters() map[string]interface{}
	SetRepeatMode(v string) *RunCommandRequest
	GetRepeatMode() *string
	SetTerminationMode(v string) *RunCommandRequest
	GetTerminationMode() *string
	SetTimeout(v int32) *RunCommandRequest
	GetTimeout() *int32
	SetUsername(v string) *RunCommandRequest
	GetUsername() *string
	SetWorkingDir(v string) *RunCommandRequest
	GetWorkingDir() *string
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
	return dara.Prettify(s)
}

func (s RunCommandRequest) GoString() string {
	return s.String()
}

func (s *RunCommandRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunCommandRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunCommandRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *RunCommandRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *RunCommandRequest) GetDescription() *string {
	return s.Description
}

func (s *RunCommandRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *RunCommandRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *RunCommandRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *RunCommandRequest) GetName() *string {
	return s.Name
}

func (s *RunCommandRequest) GetNodeIdList() []*string {
	return s.NodeIdList
}

func (s *RunCommandRequest) GetParameters() map[string]interface{} {
	return s.Parameters
}

func (s *RunCommandRequest) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *RunCommandRequest) GetTerminationMode() *string {
	return s.TerminationMode
}

func (s *RunCommandRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *RunCommandRequest) GetUsername() *string {
	return s.Username
}

func (s *RunCommandRequest) GetWorkingDir() *string {
	return s.WorkingDir
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

func (s *RunCommandRequest) Validate() error {
	return dara.Validate(s)
}

type iRunCommandShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *RunCommandShrinkRequest
	GetClientToken() *string
	SetCommandContent(v string) *RunCommandShrinkRequest
	GetCommandContent() *string
	SetCommandId(v string) *RunCommandShrinkRequest
	GetCommandId() *string
	SetContentEncoding(v string) *RunCommandShrinkRequest
	GetContentEncoding() *string
	SetDescription(v string) *RunCommandShrinkRequest
	GetDescription() *string
	SetEnableParameter(v bool) *RunCommandShrinkRequest
	GetEnableParameter() *bool
	SetFrequency(v string) *RunCommandShrinkRequest
	GetFrequency() *string
	SetLauncher(v string) *RunCommandShrinkRequest
	GetLauncher() *string
	SetName(v string) *RunCommandShrinkRequest
	GetName() *string
	SetNodeIdListShrink(v string) *RunCommandShrinkRequest
	GetNodeIdListShrink() *string
	SetParametersShrink(v string) *RunCommandShrinkRequest
	GetParametersShrink() *string
	SetRepeatMode(v string) *RunCommandShrinkRequest
	GetRepeatMode() *string
	SetTerminationMode(v string) *RunCommandShrinkRequest
	GetTerminationMode() *string
	SetTimeout(v int32) *RunCommandShrinkRequest
	GetTimeout() *int32
	SetUsername(v string) *RunCommandShrinkRequest
	GetUsername() *string
	SetWorkingDir(v string) *RunCommandShrinkRequest
	GetWorkingDir() *string
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
	return dara.Prettify(s)
}

func (s RunCommandShrinkRequest) GoString() string {
	return s.String()
}

func (s *RunCommandShrinkRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunCommandShrinkRequest) GetCommandContent() *string {
	return s.CommandContent
}

func (s *RunCommandShrinkRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *RunCommandShrinkRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *RunCommandShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *RunCommandShrinkRequest) GetEnableParameter() *bool {
	return s.EnableParameter
}

func (s *RunCommandShrinkRequest) GetFrequency() *string {
	return s.Frequency
}

func (s *RunCommandShrinkRequest) GetLauncher() *string {
	return s.Launcher
}

func (s *RunCommandShrinkRequest) GetName() *string {
	return s.Name
}

func (s *RunCommandShrinkRequest) GetNodeIdListShrink() *string {
	return s.NodeIdListShrink
}

func (s *RunCommandShrinkRequest) GetParametersShrink() *string {
	return s.ParametersShrink
}

func (s *RunCommandShrinkRequest) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *RunCommandShrinkRequest) GetTerminationMode() *string {
	return s.TerminationMode
}

func (s *RunCommandShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
}

func (s *RunCommandShrinkRequest) GetUsername() *string {
	return s.Username
}

func (s *RunCommandShrinkRequest) GetWorkingDir() *string {
	return s.WorkingDir
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

func (s *RunCommandShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iRunCommandResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *RunCommandResponseBody
	GetInvokeId() *string
	SetRequestId(v string) *RunCommandResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s RunCommandResponseBody) GoString() string {
	return s.String()
}

func (s *RunCommandResponseBody) GetInvokeId() *string {
	return s.InvokeId
}

func (s *RunCommandResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunCommandResponseBody) SetInvokeId(v string) *RunCommandResponseBody {
	s.InvokeId = &v
	return s
}

func (s *RunCommandResponseBody) SetRequestId(v string) *RunCommandResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunCommandResponseBody) Validate() error {
	return dara.Validate(s)
}

type iRunCommandResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RunCommandResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RunCommandResponse
	GetStatusCode() *int32
	SetBody(v *RunCommandResponseBody) *RunCommandResponse
	GetBody() *RunCommandResponseBody
}

type RunCommandResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RunCommandResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RunCommandResponse) String() string {
	return dara.Prettify(s)
}

func (s RunCommandResponse) GoString() string {
	return s.String()
}

func (s *RunCommandResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RunCommandResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RunCommandResponse) GetBody() *RunCommandResponseBody {
	return s.Body
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

func (s *RunCommandResponse) Validate() error {
	return dara.Validate(s)
}

type iSendFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SendFileRequest
	GetContent() *string
	SetContentType(v string) *SendFileRequest
	GetContentType() *string
	SetDescription(v string) *SendFileRequest
	GetDescription() *string
	SetFileGroup(v string) *SendFileRequest
	GetFileGroup() *string
	SetFileMode(v string) *SendFileRequest
	GetFileMode() *string
	SetFileOwner(v string) *SendFileRequest
	GetFileOwner() *string
	SetName(v string) *SendFileRequest
	GetName() *string
	SetNodeIdList(v []*string) *SendFileRequest
	GetNodeIdList() []*string
	SetOverwrite(v bool) *SendFileRequest
	GetOverwrite() *bool
	SetTargetDir(v string) *SendFileRequest
	GetTargetDir() *string
	SetTimeout(v int32) *SendFileRequest
	GetTimeout() *int32
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
	return dara.Prettify(s)
}

func (s SendFileRequest) GoString() string {
	return s.String()
}

func (s *SendFileRequest) GetContent() *string {
	return s.Content
}

func (s *SendFileRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SendFileRequest) GetDescription() *string {
	return s.Description
}

func (s *SendFileRequest) GetFileGroup() *string {
	return s.FileGroup
}

func (s *SendFileRequest) GetFileMode() *string {
	return s.FileMode
}

func (s *SendFileRequest) GetFileOwner() *string {
	return s.FileOwner
}

func (s *SendFileRequest) GetName() *string {
	return s.Name
}

func (s *SendFileRequest) GetNodeIdList() []*string {
	return s.NodeIdList
}

func (s *SendFileRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *SendFileRequest) GetTargetDir() *string {
	return s.TargetDir
}

func (s *SendFileRequest) GetTimeout() *int32 {
	return s.Timeout
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

func (s *SendFileRequest) Validate() error {
	return dara.Validate(s)
}

type iSendFileShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *SendFileShrinkRequest
	GetContent() *string
	SetContentType(v string) *SendFileShrinkRequest
	GetContentType() *string
	SetDescription(v string) *SendFileShrinkRequest
	GetDescription() *string
	SetFileGroup(v string) *SendFileShrinkRequest
	GetFileGroup() *string
	SetFileMode(v string) *SendFileShrinkRequest
	GetFileMode() *string
	SetFileOwner(v string) *SendFileShrinkRequest
	GetFileOwner() *string
	SetName(v string) *SendFileShrinkRequest
	GetName() *string
	SetNodeIdListShrink(v string) *SendFileShrinkRequest
	GetNodeIdListShrink() *string
	SetOverwrite(v bool) *SendFileShrinkRequest
	GetOverwrite() *bool
	SetTargetDir(v string) *SendFileShrinkRequest
	GetTargetDir() *string
	SetTimeout(v int32) *SendFileShrinkRequest
	GetTimeout() *int32
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
	return dara.Prettify(s)
}

func (s SendFileShrinkRequest) GoString() string {
	return s.String()
}

func (s *SendFileShrinkRequest) GetContent() *string {
	return s.Content
}

func (s *SendFileShrinkRequest) GetContentType() *string {
	return s.ContentType
}

func (s *SendFileShrinkRequest) GetDescription() *string {
	return s.Description
}

func (s *SendFileShrinkRequest) GetFileGroup() *string {
	return s.FileGroup
}

func (s *SendFileShrinkRequest) GetFileMode() *string {
	return s.FileMode
}

func (s *SendFileShrinkRequest) GetFileOwner() *string {
	return s.FileOwner
}

func (s *SendFileShrinkRequest) GetName() *string {
	return s.Name
}

func (s *SendFileShrinkRequest) GetNodeIdListShrink() *string {
	return s.NodeIdListShrink
}

func (s *SendFileShrinkRequest) GetOverwrite() *bool {
	return s.Overwrite
}

func (s *SendFileShrinkRequest) GetTargetDir() *string {
	return s.TargetDir
}

func (s *SendFileShrinkRequest) GetTimeout() *int32 {
	return s.Timeout
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

func (s *SendFileShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iSendFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *SendFileResponseBody
	GetInvokeId() *string
	SetRequestId(v string) *SendFileResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s SendFileResponseBody) GoString() string {
	return s.String()
}

func (s *SendFileResponseBody) GetInvokeId() *string {
	return s.InvokeId
}

func (s *SendFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SendFileResponseBody) SetInvokeId(v string) *SendFileResponseBody {
	s.InvokeId = &v
	return s
}

func (s *SendFileResponseBody) SetRequestId(v string) *SendFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *SendFileResponseBody) Validate() error {
	return dara.Validate(s)
}

type iSendFileResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SendFileResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SendFileResponse
	GetStatusCode() *int32
	SetBody(v *SendFileResponseBody) *SendFileResponse
	GetBody() *SendFileResponseBody
}

type SendFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SendFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SendFileResponse) String() string {
	return dara.Prettify(s)
}

func (s SendFileResponse) GoString() string {
	return s.String()
}

func (s *SendFileResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SendFileResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SendFileResponse) GetBody() *SendFileResponseBody {
	return s.Body
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

func (s *SendFileResponse) Validate() error {
	return dara.Validate(s)
}

type iShrinkClusterRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ShrinkClusterRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *ShrinkClusterRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodeGroups(v []*ShrinkClusterRequestNodeGroups) *ShrinkClusterRequest
	GetNodeGroups() []*ShrinkClusterRequestNodeGroups
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
	return dara.Prettify(s)
}

func (s ShrinkClusterRequest) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ShrinkClusterRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ShrinkClusterRequest) GetNodeGroups() []*ShrinkClusterRequestNodeGroups {
	return s.NodeGroups
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

func (s *ShrinkClusterRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ShrinkClusterRequestNodeGroups) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequestNodeGroups) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *ShrinkClusterRequestNodeGroups) GetNodes() []*ShrinkClusterRequestNodeGroupsNodes {
	return s.Nodes
}

func (s *ShrinkClusterRequestNodeGroups) SetNodeGroupId(v string) *ShrinkClusterRequestNodeGroups {
	s.NodeGroupId = &v
	return s
}

func (s *ShrinkClusterRequestNodeGroups) SetNodes(v []*ShrinkClusterRequestNodeGroupsNodes) *ShrinkClusterRequestNodeGroups {
	s.Nodes = v
	return s
}

func (s *ShrinkClusterRequestNodeGroups) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s ShrinkClusterRequestNodeGroupsNodes) GoString() string {
	return s.String()
}

func (s *ShrinkClusterRequestNodeGroupsNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *ShrinkClusterRequestNodeGroupsNodes) SetNodeId(v string) *ShrinkClusterRequestNodeGroupsNodes {
	s.NodeId = &v
	return s
}

func (s *ShrinkClusterRequestNodeGroupsNodes) Validate() error {
	return dara.Validate(s)
}

type iShrinkClusterShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClusterId(v string) *ShrinkClusterShrinkRequest
	GetClusterId() *string
	SetIgnoreFailedNodeTasks(v bool) *ShrinkClusterShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodeGroupsShrink(v string) *ShrinkClusterShrinkRequest
	GetNodeGroupsShrink() *string
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
	return dara.Prettify(s)
}

func (s ShrinkClusterShrinkRequest) GoString() string {
	return s.String()
}

func (s *ShrinkClusterShrinkRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ShrinkClusterShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *ShrinkClusterShrinkRequest) GetNodeGroupsShrink() *string {
	return s.NodeGroupsShrink
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

func (s *ShrinkClusterShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iShrinkClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ShrinkClusterResponseBody
	GetRequestId() *string
	SetTaskId(v string) *ShrinkClusterResponseBody
	GetTaskId() *string
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
	return dara.Prettify(s)
}

func (s ShrinkClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ShrinkClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ShrinkClusterResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *ShrinkClusterResponseBody) SetRequestId(v string) *ShrinkClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ShrinkClusterResponseBody) SetTaskId(v string) *ShrinkClusterResponseBody {
	s.TaskId = &v
	return s
}

func (s *ShrinkClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type iShrinkClusterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ShrinkClusterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ShrinkClusterResponse
	GetStatusCode() *int32
	SetBody(v *ShrinkClusterResponseBody) *ShrinkClusterResponse
	GetBody() *ShrinkClusterResponseBody
}

type ShrinkClusterResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShrinkClusterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShrinkClusterResponse) String() string {
	return dara.Prettify(s)
}

func (s ShrinkClusterResponse) GoString() string {
	return s.String()
}

func (s *ShrinkClusterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ShrinkClusterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ShrinkClusterResponse) GetBody() *ShrinkClusterResponseBody {
	return s.Body
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

func (s *ShrinkClusterResponse) Validate() error {
	return dara.Validate(s)
}

type iStopInvocationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *StopInvocationRequest
	GetInvokeId() *string
	SetNodeIdList(v []*string) *StopInvocationRequest
	GetNodeIdList() []*string
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
	return dara.Prettify(s)
}

func (s StopInvocationRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *StopInvocationRequest) GetNodeIdList() []*string {
	return s.NodeIdList
}

func (s *StopInvocationRequest) SetInvokeId(v string) *StopInvocationRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationRequest) SetNodeIdList(v []*string) *StopInvocationRequest {
	s.NodeIdList = v
	return s
}

func (s *StopInvocationRequest) Validate() error {
	return dara.Validate(s)
}

type iStopInvocationShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInvokeId(v string) *StopInvocationShrinkRequest
	GetInvokeId() *string
	SetNodeIdListShrink(v string) *StopInvocationShrinkRequest
	GetNodeIdListShrink() *string
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
	return dara.Prettify(s)
}

func (s StopInvocationShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopInvocationShrinkRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *StopInvocationShrinkRequest) GetNodeIdListShrink() *string {
	return s.NodeIdListShrink
}

func (s *StopInvocationShrinkRequest) SetInvokeId(v string) *StopInvocationShrinkRequest {
	s.InvokeId = &v
	return s
}

func (s *StopInvocationShrinkRequest) SetNodeIdListShrink(v string) *StopInvocationShrinkRequest {
	s.NodeIdListShrink = &v
	return s
}

func (s *StopInvocationShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iStopInvocationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopInvocationResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s StopInvocationResponseBody) GoString() string {
	return s.String()
}

func (s *StopInvocationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopInvocationResponseBody) SetRequestId(v string) *StopInvocationResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopInvocationResponseBody) Validate() error {
	return dara.Validate(s)
}

type iStopInvocationResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopInvocationResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopInvocationResponse
	GetStatusCode() *int32
	SetBody(v *StopInvocationResponseBody) *StopInvocationResponse
	GetBody() *StopInvocationResponseBody
}

type StopInvocationResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopInvocationResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopInvocationResponse) String() string {
	return dara.Prettify(s)
}

func (s StopInvocationResponse) GoString() string {
	return s.String()
}

func (s *StopInvocationResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopInvocationResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopInvocationResponse) GetBody() *StopInvocationResponseBody {
	return s.Body
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

func (s *StopInvocationResponse) Validate() error {
	return dara.Validate(s)
}

type iStopNodesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreFailedNodeTasks(v bool) *StopNodesRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodes(v []*string) *StopNodesRequest
	GetNodes() []*string
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
	return dara.Prettify(s)
}

func (s StopNodesRequest) GoString() string {
	return s.String()
}

func (s *StopNodesRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *StopNodesRequest) GetNodes() []*string {
	return s.Nodes
}

func (s *StopNodesRequest) SetIgnoreFailedNodeTasks(v bool) *StopNodesRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *StopNodesRequest) SetNodes(v []*string) *StopNodesRequest {
	s.Nodes = v
	return s
}

func (s *StopNodesRequest) Validate() error {
	return dara.Validate(s)
}

type iStopNodesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetIgnoreFailedNodeTasks(v bool) *StopNodesShrinkRequest
	GetIgnoreFailedNodeTasks() *bool
	SetNodesShrink(v string) *StopNodesShrinkRequest
	GetNodesShrink() *string
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
	return dara.Prettify(s)
}

func (s StopNodesShrinkRequest) GoString() string {
	return s.String()
}

func (s *StopNodesShrinkRequest) GetIgnoreFailedNodeTasks() *bool {
	return s.IgnoreFailedNodeTasks
}

func (s *StopNodesShrinkRequest) GetNodesShrink() *string {
	return s.NodesShrink
}

func (s *StopNodesShrinkRequest) SetIgnoreFailedNodeTasks(v bool) *StopNodesShrinkRequest {
	s.IgnoreFailedNodeTasks = &v
	return s
}

func (s *StopNodesShrinkRequest) SetNodesShrink(v string) *StopNodesShrinkRequest {
	s.NodesShrink = &v
	return s
}

func (s *StopNodesShrinkRequest) Validate() error {
	return dara.Validate(s)
}

type iStopNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopNodesResponseBody
	GetRequestId() *string
	SetTaskId(v string) *StopNodesResponseBody
	GetTaskId() *string
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
	return dara.Prettify(s)
}

func (s StopNodesResponseBody) GoString() string {
	return s.String()
}

func (s *StopNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopNodesResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *StopNodesResponseBody) SetRequestId(v string) *StopNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopNodesResponseBody) SetTaskId(v string) *StopNodesResponseBody {
	s.TaskId = &v
	return s
}

func (s *StopNodesResponseBody) Validate() error {
	return dara.Validate(s)
}

type iStopNodesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StopNodesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StopNodesResponse
	GetStatusCode() *int32
	SetBody(v *StopNodesResponseBody) *StopNodesResponse
	GetBody() *StopNodesResponseBody
}

type StopNodesResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StopNodesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StopNodesResponse) String() string {
	return dara.Prettify(s)
}

func (s StopNodesResponse) GoString() string {
	return s.String()
}

func (s *StopNodesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StopNodesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StopNodesResponse) GetBody() *StopNodesResponseBody {
	return s.Body
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

func (s *StopNodesResponse) Validate() error {
	return dara.Validate(s)
}

type iTagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRegionId(v string) *TagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *TagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *TagResourcesRequest
	GetResourceType() *string
	SetService(v string) *TagResourcesRequest
	GetService() *string
	SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest
	GetTag() []*TagResourcesRequestTag
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
	Service      *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// Tags
	//
	// This parameter is required.
	Tag []*TagResourcesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s TagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesRequest) GoString() string {
	return s.String()
}

func (s *TagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *TagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *TagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *TagResourcesRequest) GetService() *string {
	return s.Service
}

func (s *TagResourcesRequest) GetTag() []*TagResourcesRequestTag {
	return s.Tag
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

func (s *TagResourcesRequest) SetService(v string) *TagResourcesRequest {
	s.Service = &v
	return s
}

func (s *TagResourcesRequest) SetTag(v []*TagResourcesRequestTag) *TagResourcesRequest {
	s.Tag = v
	return s
}

func (s *TagResourcesRequest) Validate() error {
	return dara.Validate(s)
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
	return dara.Prettify(s)
}

func (s TagResourcesRequestTag) GoString() string {
	return s.String()
}

func (s *TagResourcesRequestTag) GetKey() *string {
	return s.Key
}

func (s *TagResourcesRequestTag) GetValue() *string {
	return s.Value
}

func (s *TagResourcesRequestTag) SetKey(v string) *TagResourcesRequestTag {
	s.Key = &v
	return s
}

func (s *TagResourcesRequestTag) SetValue(v string) *TagResourcesRequestTag {
	s.Value = &v
	return s
}

func (s *TagResourcesRequestTag) Validate() error {
	return dara.Validate(s)
}

type iTagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TagResourcesResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s TagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *TagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TagResourcesResponseBody) SetRequestId(v string) *TagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type iTagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *TagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *TagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *TagResourcesResponseBody) *TagResourcesResponse
	GetBody() *TagResourcesResponseBody
}

type TagResourcesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s TagResourcesResponse) GoString() string {
	return s.String()
}

func (s *TagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *TagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *TagResourcesResponse) GetBody() *TagResourcesResponseBody {
	return s.Body
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

func (s *TagResourcesResponse) Validate() error {
	return dara.Validate(s)
}

type iUntagResourcesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAll(v bool) *UntagResourcesRequest
	GetAll() *bool
	SetRegionId(v string) *UntagResourcesRequest
	GetRegionId() *string
	SetResourceId(v []*string) *UntagResourcesRequest
	GetResourceId() []*string
	SetResourceType(v string) *UntagResourcesRequest
	GetResourceType() *string
	SetService(v string) *UntagResourcesRequest
	GetService() *string
	SetTagKey(v []*string) *UntagResourcesRequest
	GetTagKey() []*string
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
	Service      *string `json:"Service,omitempty" xml:"Service,omitempty"`
	// List of tag keys
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s UntagResourcesRequest) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesRequest) GoString() string {
	return s.String()
}

func (s *UntagResourcesRequest) GetAll() *bool {
	return s.All
}

func (s *UntagResourcesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *UntagResourcesRequest) GetResourceId() []*string {
	return s.ResourceId
}

func (s *UntagResourcesRequest) GetResourceType() *string {
	return s.ResourceType
}

func (s *UntagResourcesRequest) GetService() *string {
	return s.Service
}

func (s *UntagResourcesRequest) GetTagKey() []*string {
	return s.TagKey
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

func (s *UntagResourcesRequest) SetService(v string) *UntagResourcesRequest {
	s.Service = &v
	return s
}

func (s *UntagResourcesRequest) SetTagKey(v []*string) *UntagResourcesRequest {
	s.TagKey = v
	return s
}

func (s *UntagResourcesRequest) Validate() error {
	return dara.Validate(s)
}

type iUntagResourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UntagResourcesResponseBody
	GetRequestId() *string
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
	return dara.Prettify(s)
}

func (s UntagResourcesResponseBody) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UntagResourcesResponseBody) SetRequestId(v string) *UntagResourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *UntagResourcesResponseBody) Validate() error {
	return dara.Validate(s)
}

type iUntagResourcesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UntagResourcesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UntagResourcesResponse
	GetStatusCode() *int32
	SetBody(v *UntagResourcesResponseBody) *UntagResourcesResponse
	GetBody() *UntagResourcesResponseBody
}

type UntagResourcesResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UntagResourcesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UntagResourcesResponse) String() string {
	return dara.Prettify(s)
}

func (s UntagResourcesResponse) GoString() string {
	return s.String()
}

func (s *UntagResourcesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UntagResourcesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UntagResourcesResponse) GetBody() *UntagResourcesResponseBody {
	return s.Body
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

func (s *UntagResourcesResponse) Validate() error {
	return dara.Validate(s)
}

type iUpdateNodeGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetNewNodeGroupName(v string) *UpdateNodeGroupRequest
	GetNewNodeGroupName() *string
	SetNodeGroupId(v string) *UpdateNodeGroupRequest
	GetNodeGroupId() *string
	SetUserData(v string) *UpdateNodeGroupRequest
	GetUserData() *string
}

type UpdateNodeGroupRequest struct {
	// example:
	//
	// test-update
	NewNodeGroupName *string `json:"NewNodeGroupName,omitempty" xml:"NewNodeGroupName,omitempty"`
	// example:
	//
	// i120021051733814190732
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	UserData    *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
}

func (s UpdateNodeGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupRequest) GetNewNodeGroupName() *string {
	return s.NewNodeGroupName
}

func (s *UpdateNodeGroupRequest) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *UpdateNodeGroupRequest) GetUserData() *string {
	return s.UserData
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

func (s *UpdateNodeGroupRequest) Validate() error {
	return dara.Validate(s)
}

type iUpdateNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNodeGroupResponseBody
	GetRequestId() *string
}

type UpdateNodeGroupResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8F065DDD-6996-5973-9691-9EC57BD0072E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNodeGroupResponseBody) SetRequestId(v string) *UpdateNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}

type iUpdateNodeGroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateNodeGroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateNodeGroupResponse
	GetStatusCode() *int32
	SetBody(v *UpdateNodeGroupResponseBody) *UpdateNodeGroupResponse
	GetBody() *UpdateNodeGroupResponseBody
}

type UpdateNodeGroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateNodeGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateNodeGroupResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateNodeGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateNodeGroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateNodeGroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateNodeGroupResponse) GetBody() *UpdateNodeGroupResponseBody {
	return s.Body
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

func (s *UpdateNodeGroupResponse) Validate() error {
	return dara.Validate(s)
}

type Client struct {
	openapi.Client
	DisableSDKError *bool
}

func NewClient(config *openapiutil.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapiutil.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	client.EndpointRule = dara.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(dara.String("eflo-controller"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !dara.IsNil(endpoint) {
		_result = endpoint
		return _result, _err
	}

	if !dara.IsNil(endpointMap) && !dara.IsNil(endpointMap[dara.StringValue(regionId)]) {
		_result = endpointMap[dara.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := openapiutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
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
func (client *Client) ApproveOperationWithOptions(request *ApproveOperationRequest, runtime *dara.RuntimeOptions) (_result *ApproveOperationResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.OperationType) {
		body["OperationType"] = request.OperationType
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ApproveOperation"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Approve Operation
//
// @param request - ApproveOperationRequest
//
// @return ApproveOperationResponse
func (client *Client) ApproveOperation(request *ApproveOperationRequest) (_result *ApproveOperationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ChangeResourceGroupWithOptions(request *ChangeResourceGroupRequest, runtime *dara.RuntimeOptions) (_result *ChangeResourceGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ResourceGroupId) {
		query["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceRegionId) {
		query["ResourceRegionId"] = request.ResourceRegionId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ChangeResourceGroup"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Target Resource Group
//
// @param request - ChangeResourceGroupRequest
//
// @return ChangeResourceGroupResponse
func (client *Client) ChangeResourceGroup(request *ChangeResourceGroupRequest) (_result *ChangeResourceGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) CloseSessionWithOptions(request *CloseSessionRequest, runtime *dara.RuntimeOptions) (_result *CloseSessionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.SessionId) {
		body["SessionId"] = request.SessionId
	}

	if !dara.IsNil(request.SessionToken) {
		body["SessionToken"] = request.SessionToken
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CloseSession"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

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
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) CreateClusterWithOptions(tmpReq *CreateClusterRequest, runtime *dara.RuntimeOptions) (_result *CreateClusterResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Components) {
		request.ComponentsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Components, dara.String("Components"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Networks) {
		request.NetworksShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Networks, dara.String("Networks"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NimizVSwitches) {
		request.NimizVSwitchesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NimizVSwitches, dara.String("NimizVSwitches"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeGroups) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, dara.String("NodeGroups"), dara.String("json"))
	}

	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterDescription) {
		body["ClusterDescription"] = request.ClusterDescription
	}

	if !dara.IsNil(request.ClusterName) {
		body["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.ClusterType) {
		body["ClusterType"] = request.ClusterType
	}

	if !dara.IsNil(request.ComponentsShrink) {
		body["Components"] = request.ComponentsShrink
	}

	if !dara.IsNil(request.HpnZone) {
		body["HpnZone"] = request.HpnZone
	}

	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.NetworksShrink) {
		body["Networks"] = request.NetworksShrink
	}

	if !dara.IsNil(request.NimizVSwitchesShrink) {
		body["NimizVSwitches"] = request.NimizVSwitchesShrink
	}

	if !dara.IsNil(request.NodeGroupsShrink) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	if !dara.IsNil(request.OpenEniJumboFrame) {
		body["OpenEniJumboFrame"] = request.OpenEniJumboFrame
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateCluster"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Create a large-scale computing cluster
//
// @param request - CreateClusterRequest
//
// @return CreateClusterResponse
func (client *Client) CreateCluster(request *CreateClusterRequest) (_result *CreateClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) CreateDiagnosticTaskWithOptions(tmpReq *CreateDiagnosticTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateDiagnosticTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateDiagnosticTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.AiJobLogInfo) {
		request.AiJobLogInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.AiJobLogInfo, dara.String("AiJobLogInfo"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeIds) {
		request.NodeIdsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIds, dara.String("NodeIds"), dara.String("simple"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.AiJobLogInfoShrink) {
		body["AiJobLogInfo"] = request.AiJobLogInfoShrink
	}

	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.DiagnosticType) {
		body["DiagnosticType"] = request.DiagnosticType
	}

	if !dara.IsNil(request.NodeIdsShrink) {
		body["NodeIds"] = request.NodeIdsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateDiagnosticTask"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Diagnostic Task Creation Interface
//
// @param request - CreateDiagnosticTaskRequest
//
// @return CreateDiagnosticTaskResponse
func (client *Client) CreateDiagnosticTask(request *CreateDiagnosticTaskRequest) (_result *CreateDiagnosticTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) CreateNetTestTaskWithOptions(tmpReq *CreateNetTestTaskRequest, runtime *dara.RuntimeOptions) (_result *CreateNetTestTaskResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateNetTestTaskShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.CommTest) {
		request.CommTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.CommTest, dara.String("CommTest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.DelayTest) {
		request.DelayTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.DelayTest, dara.String("DelayTest"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.TrafficTest) {
		request.TrafficTestShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.TrafficTest, dara.String("TrafficTest"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.ClusterName) {
		body["ClusterName"] = request.ClusterName
	}

	if !dara.IsNil(request.CommTestShrink) {
		body["CommTest"] = request.CommTestShrink
	}

	if !dara.IsNil(request.DelayTestShrink) {
		body["DelayTest"] = request.DelayTestShrink
	}

	if !dara.IsNil(request.NetTestType) {
		body["NetTestType"] = request.NetTestType
	}

	if !dara.IsNil(request.NetworkMode) {
		body["NetworkMode"] = request.NetworkMode
	}

	if !dara.IsNil(request.Port) {
		body["Port"] = request.Port
	}

	if !dara.IsNil(request.TrafficTestShrink) {
		body["TrafficTest"] = request.TrafficTestShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNetTestTask"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

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
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.CreateNetTestTaskWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建集群下的节点分组
//
// @param tmpReq - CreateNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateNodeGroupResponse
func (client *Client) CreateNodeGroupWithOptions(tmpReq *CreateNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *CreateNodeGroupResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &CreateNodeGroupShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeGroup) {
		request.NodeGroupShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroup, dara.String("NodeGroup"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeUnit) {
		request.NodeUnitShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeUnit, dara.String("NodeUnit"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeGroupShrink) {
		body["NodeGroup"] = request.NodeGroupShrink
	}

	if !dara.IsNil(request.NodeUnitShrink) {
		body["NodeUnit"] = request.NodeUnitShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateNodeGroup"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建集群下的节点分组
//
// @param request - CreateNodeGroupRequest
//
// @return CreateNodeGroupResponse
func (client *Client) CreateNodeGroup(request *CreateNodeGroupRequest) (_result *CreateNodeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) CreateSessionWithOptions(request *CreateSessionRequest, runtime *dara.RuntimeOptions) (_result *CreateSessionResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.SessionType) {
		body["SessionType"] = request.SessionType
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateSession"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

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
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) CreateVscWithOptions(request *CreateVscRequest, runtime *dara.RuntimeOptions) (_result *CreateVscResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.Tag) {
		body["Tag"] = request.Tag
	}

	if !dara.IsNil(request.VscName) {
		body["VscName"] = request.VscName
	}

	if !dara.IsNil(request.VscType) {
		body["VscType"] = request.VscType
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("CreateVsc"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 创建Vsc
//
// @param request - CreateVscRequest
//
// @return CreateVscResponse
func (client *Client) CreateVsc(request *CreateVscRequest) (_result *CreateVscResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DeleteClusterWithOptions(request *DeleteClusterRequest, runtime *dara.RuntimeOptions) (_result *DeleteClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteCluster"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Delete cluster instance
//
// @param request - DeleteClusterRequest
//
// @return DeleteClusterResponse
func (client *Client) DeleteCluster(request *DeleteClusterRequest) (_result *DeleteClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.DeleteClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除节点分组
//
// @param request - DeleteNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteNodeGroupResponse
func (client *Client) DeleteNodeGroupWithOptions(request *DeleteNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *DeleteNodeGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.NodeGroupId) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteNodeGroup"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除节点分组
//
// @param request - DeleteNodeGroupRequest
//
// @return DeleteNodeGroupResponse
func (client *Client) DeleteNodeGroup(request *DeleteNodeGroupRequest) (_result *DeleteNodeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DeleteVscWithOptions(request *DeleteVscRequest, runtime *dara.RuntimeOptions) (_result *DeleteVscResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		query["ClientToken"] = request.ClientToken
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.VscId) {
		body["VscId"] = request.VscId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DeleteVsc"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 删除Vsc
//
// @param request - DeleteVscRequest
//
// @return DeleteVscResponse
func (client *Client) DeleteVsc(request *DeleteVscRequest) (_result *DeleteVscResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DescribeClusterWithOptions(request *DescribeClusterRequest, runtime *dara.RuntimeOptions) (_result *DescribeClusterResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeCluster"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Cluster Details
//
// @param request - DescribeClusterRequest
//
// @return DescribeClusterResponse
func (client *Client) DescribeCluster(request *DescribeClusterRequest) (_result *DescribeClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.DescribeClusterWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 诊断任务查询接口
//
// @param request - DescribeDiagnosticResultRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeDiagnosticResultResponse
func (client *Client) DescribeDiagnosticResultWithOptions(request *DescribeDiagnosticResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeDiagnosticResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DiagnosticId) {
		body["DiagnosticId"] = request.DiagnosticId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeDiagnosticResult"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 诊断任务查询接口
//
// @param request - DescribeDiagnosticResultRequest
//
// @return DescribeDiagnosticResultResponse
func (client *Client) DescribeDiagnosticResult(request *DescribeDiagnosticResultRequest) (_result *DescribeDiagnosticResultResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DescribeInvocationsWithOptions(request *DescribeInvocationsRequest, runtime *dara.RuntimeOptions) (_result *DescribeInvocationsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ContentEncoding) {
		body["ContentEncoding"] = request.ContentEncoding
	}

	if !dara.IsNil(request.IncludeOutput) {
		body["IncludeOutput"] = request.IncludeOutput
	}

	if !dara.IsNil(request.InvokeId) {
		body["InvokeId"] = request.InvokeId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeInvocations"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Query the list and status of operations assistant command executions
//
// @param request - DescribeInvocationsRequest
//
// @return DescribeInvocationsResponse
func (client *Client) DescribeInvocations(request *DescribeInvocationsRequest) (_result *DescribeInvocationsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DescribeNetTestResultWithOptions(request *DescribeNetTestResultRequest, runtime *dara.RuntimeOptions) (_result *DescribeNetTestResultResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TestId) {
		body["TestId"] = request.TestId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNetTestResult"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

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
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DescribeNodeWithOptions(request *DescribeNodeRequest, runtime *dara.RuntimeOptions) (_result *DescribeNodeResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNode"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Query node list
//
// @param request - DescribeNodeRequest
//
// @return DescribeNodeResponse
func (client *Client) DescribeNode(request *DescribeNodeRequest) (_result *DescribeNodeResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.DescribeNodeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询节点分组
//
// @param request - DescribeNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DescribeNodeGroupResponse
func (client *Client) DescribeNodeGroupWithOptions(request *DescribeNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *DescribeNodeGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NodeGroupId) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeNodeGroup"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询节点分组
//
// @param request - DescribeNodeGroupRequest
//
// @return DescribeNodeGroupResponse
func (client *Client) DescribeNodeGroup(request *DescribeNodeGroupRequest) (_result *DescribeNodeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.DescribeNodeGroupWithOptions(request, runtime)
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
func (client *Client) DescribeRegionsWithOptions(request *DescribeRegionsRequest, runtime *dara.RuntimeOptions) (_result *DescribeRegionsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeRegions"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Region List
//
// @param request - DescribeRegionsRequest
//
// @return DescribeRegionsResponse
func (client *Client) DescribeRegions(request *DescribeRegionsRequest) (_result *DescribeRegionsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DescribeSendFileResultsWithOptions(request *DescribeSendFileResultsRequest, runtime *dara.RuntimeOptions) (_result *DescribeSendFileResultsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.InvokeId) {
		body["InvokeId"] = request.InvokeId
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeSendFileResults"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Query the list and status of files sent by the operation assistant
//
// @param request - DescribeSendFileResultsRequest
//
// @return DescribeSendFileResultsResponse
func (client *Client) DescribeSendFileResults(request *DescribeSendFileResultsRequest) (_result *DescribeSendFileResultsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DescribeTaskWithOptions(request *DescribeTaskRequest, runtime *dara.RuntimeOptions) (_result *DescribeTaskResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.TaskId) {
		body["TaskId"] = request.TaskId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeTask"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Query Task Details
//
// @param request - DescribeTaskRequest
//
// @return DescribeTaskResponse
func (client *Client) DescribeTask(request *DescribeTaskRequest) (_result *DescribeTaskResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DescribeVscWithOptions(request *DescribeVscRequest, runtime *dara.RuntimeOptions) (_result *DescribeVscResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.VscId) {
		body["VscId"] = request.VscId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeVsc"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 获取单个Vsc详情
//
// @param request - DescribeVscRequest
//
// @return DescribeVscResponse
func (client *Client) DescribeVsc(request *DescribeVscRequest) (_result *DescribeVscResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) DescribeZonesWithOptions(request *DescribeZonesRequest, runtime *dara.RuntimeOptions) (_result *DescribeZonesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.AcceptLanguage) {
		body["AcceptLanguage"] = request.AcceptLanguage
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("DescribeZones"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # List of Available Zones
//
// @param request - DescribeZonesRequest
//
// @return DescribeZonesResponse
func (client *Client) DescribeZones(request *DescribeZonesRequest) (_result *DescribeZonesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ExtendClusterWithOptions(tmpReq *ExtendClusterRequest, runtime *dara.RuntimeOptions) (_result *ExtendClusterResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ExtendClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.IpAllocationPolicy) {
		request.IpAllocationPolicyShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.IpAllocationPolicy, dara.String("IpAllocationPolicy"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.NodeGroups) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, dara.String("NodeGroups"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.VpdSubnets) {
		request.VpdSubnetsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.VpdSubnets, dara.String("VpdSubnets"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.IpAllocationPolicyShrink) {
		body["IpAllocationPolicy"] = request.IpAllocationPolicyShrink
	}

	if !dara.IsNil(request.NodeGroupsShrink) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	if !dara.IsNil(request.VSwitchZoneId) {
		body["VSwitchZoneId"] = request.VSwitchZoneId
	}

	if !dara.IsNil(request.VpdSubnetsShrink) {
		body["VpdSubnets"] = request.VpdSubnetsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ExtendCluster"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Cluster Scaling
//
// @param request - ExtendClusterRequest
//
// @return ExtendClusterResponse
func (client *Client) ExtendCluster(request *ExtendClusterRequest) (_result *ExtendClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ListClusterNodesWithOptions(request *ListClusterNodesRequest, runtime *dara.RuntimeOptions) (_result *ListClusterNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeGroupId) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusterNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # List of host groups under the cluster, and list of hosts under each group
//
// @param request - ListClusterNodesRequest
//
// @return ListClusterNodesResponse
func (client *Client) ListClusterNodes(request *ListClusterNodesRequest) (_result *ListClusterNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ListClustersWithOptions(request *ListClustersRequest, runtime *dara.RuntimeOptions) (_result *ListClustersResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListClusters"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Get the list of cluster instances
//
// @param request - ListClustersRequest
//
// @return ListClustersResponse
func (client *Client) ListClusters(request *ListClustersRequest) (_result *ListClustersResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ListDiagnosticResultsWithOptions(request *ListDiagnosticResultsRequest, runtime *dara.RuntimeOptions) (_result *ListDiagnosticResultsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.DiagType) {
		body["DiagType"] = request.DiagType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListDiagnosticResults"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

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
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ListFreeNodesWithOptions(request *ListFreeNodesRequest, runtime *dara.RuntimeOptions) (_result *ListFreeNodesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tags) {
		query["Tags"] = request.Tags
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.HpnZone) {
		body["HpnZone"] = request.HpnZone
	}

	if !dara.IsNil(request.IncludeAbnormalNodes) {
		body["IncludeAbnormalNodes"] = request.IncludeAbnormalNodes
	}

	if !dara.IsNil(request.MachineType) {
		body["MachineType"] = request.MachineType
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListFreeNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # List of Available Physical Machines
//
// @param request - ListFreeNodesRequest
//
// @return ListFreeNodesResponse
func (client *Client) ListFreeNodes(request *ListFreeNodesRequest) (_result *ListFreeNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ListImagesWithOptions(request *ListImagesRequest, runtime *dara.RuntimeOptions) (_result *ListImagesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Architecture) {
		body["Architecture"] = request.Architecture
	}

	if !dara.IsNil(request.ImageVersion) {
		body["ImageVersion"] = request.ImageVersion
	}

	if !dara.IsNil(request.Platform) {
		body["Platform"] = request.Platform
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListImages"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Query the list of images available to the user
//
// @param request - ListImagesRequest
//
// @return ListImagesResponse
func (client *Client) ListImages(request *ListImagesRequest) (_result *ListImagesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ListMachineNetworkInfoWithOptions(tmpReq *ListMachineNetworkInfoRequest, runtime *dara.RuntimeOptions) (_result *ListMachineNetworkInfoResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ListMachineNetworkInfoShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.MachineHpnInfo) {
		request.MachineHpnInfoShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.MachineHpnInfo, dara.String("MachineHpnInfo"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MachineHpnInfoShrink) {
		body["MachineHpnInfo"] = request.MachineHpnInfoShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMachineNetworkInfo"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

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
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ListMachineTypesWithOptions(request *ListMachineTypesRequest, runtime *dara.RuntimeOptions) (_result *ListMachineTypesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListMachineTypes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Query the list of machine types available to the user
//
// @param request - ListMachineTypesRequest
//
// @return ListMachineTypesResponse
func (client *Client) ListMachineTypes(request *ListMachineTypesRequest) (_result *ListMachineTypesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ListNetTestResultsWithOptions(request *ListNetTestResultsRequest, runtime *dara.RuntimeOptions) (_result *ListNetTestResultsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NetTestType) {
		body["NetTestType"] = request.NetTestType
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNetTestResults"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

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
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ListNodeGroupsWithOptions(request *ListNodeGroupsRequest, runtime *dara.RuntimeOptions) (_result *ListNodeGroupsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeGroupId) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListNodeGroups"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Query Node Group Information Under the Cluster
//
// @param request - ListNodeGroupsRequest
//
// @return ListNodeGroupsResponse
func (client *Client) ListNodeGroups(request *ListNodeGroupsRequest) (_result *ListNodeGroupsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.ListNodeGroupsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询系统日志
//
// @param request - ListSyslogsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListSyslogsResponse
func (client *Client) ListSyslogsWithOptions(request *ListSyslogsRequest, runtime *dara.RuntimeOptions) (_result *ListSyslogsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.FromTime) {
		body["FromTime"] = request.FromTime
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeId) {
		body["NodeId"] = request.NodeId
	}

	if !dara.IsNil(request.Query) {
		body["Query"] = request.Query
	}

	if !dara.IsNil(request.Reverse) {
		body["Reverse"] = request.Reverse
	}

	if !dara.IsNil(request.ToTime) {
		body["ToTime"] = request.ToTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListSyslogs"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询系统日志
//
// @param request - ListSyslogsRequest
//
// @return ListSyslogsResponse
func (client *Client) ListSyslogs(request *ListSyslogsRequest) (_result *ListSyslogsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.ListSyslogsWithOptions(request, runtime)
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
func (client *Client) ListTagResourcesWithOptions(request *ListTagResourcesRequest, runtime *dara.RuntimeOptions) (_result *ListTagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.NextToken) {
		query["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListTagResources"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Query Resource Tags
//
// @param request - ListTagResourcesRequest
//
// @return ListTagResourcesResponse
func (client *Client) ListTagResources(request *ListTagResourcesRequest) (_result *ListTagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ListUserClusterTypesWithOptions(runtime *dara.RuntimeOptions) (_result *ListUserClusterTypesResponse, _err error) {
	req := &openapiutil.OpenApiRequest{}
	params := &openapiutil.Params{
		Action:      dara.String("ListUserClusterTypes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

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
	runtime := &dara.RuntimeOptions{}
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
// @param request - ListVscsRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListVscsResponse
func (client *Client) ListVscsWithOptions(request *ListVscsRequest, runtime *dara.RuntimeOptions) (_result *ListVscsResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.MaxResults) {
		body["MaxResults"] = request.MaxResults
	}

	if !dara.IsNil(request.NextToken) {
		body["NextToken"] = request.NextToken
	}

	if !dara.IsNil(request.NodeIds) {
		body["NodeIds"] = request.NodeIds
	}

	if !dara.IsNil(request.ResourceGroupId) {
		body["ResourceGroupId"] = request.ResourceGroupId
	}

	if !dara.IsNil(request.VscName) {
		body["VscName"] = request.VscName
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
		Body:  openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ListVscs"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 查询Vsc列表
//
// @param request - ListVscsRequest
//
// @return ListVscsResponse
func (client *Client) ListVscs(request *ListVscsRequest) (_result *ListVscsResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) RebootNodesWithOptions(tmpReq *RebootNodesRequest, runtime *dara.RuntimeOptions) (_result *RebootNodesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RebootNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("Nodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.NodesShrink) {
		body["Nodes"] = request.NodesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RebootNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Reboot Machine
//
// @param request - RebootNodesRequest
//
// @return RebootNodesResponse
func (client *Client) RebootNodes(request *RebootNodesRequest) (_result *RebootNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ReimageNodesWithOptions(tmpReq *ReimageNodesRequest, runtime *dara.RuntimeOptions) (_result *ReimageNodesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ReimageNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("Nodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.NodesShrink) {
		body["Nodes"] = request.NodesShrink
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReimageNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Machine Reinstallation
//
// @param request - ReimageNodesRequest
//
// @return ReimageNodesResponse
func (client *Client) ReimageNodes(request *ReimageNodesRequest) (_result *ReimageNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.ReimageNodesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 节点异常问题上报
//
// @param tmpReq - ReportNodesStatusRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ReportNodesStatusResponse
func (client *Client) ReportNodesStatusWithOptions(tmpReq *ReportNodesStatusRequest, runtime *dara.RuntimeOptions) (_result *ReportNodesStatusResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ReportNodesStatusShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("Nodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EndTime) {
		body["EndTime"] = request.EndTime
	}

	if !dara.IsNil(request.IssueCategory) {
		body["IssueCategory"] = request.IssueCategory
	}

	if !dara.IsNil(request.NodesShrink) {
		body["Nodes"] = request.NodesShrink
	}

	if !dara.IsNil(request.Reason) {
		body["Reason"] = request.Reason
	}

	if !dara.IsNil(request.StartTime) {
		body["StartTime"] = request.StartTime
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ReportNodesStatus"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 节点异常问题上报
//
// @param request - ReportNodesStatusRequest
//
// @return ReportNodesStatusResponse
func (client *Client) ReportNodesStatus(request *ReportNodesStatusRequest) (_result *ReportNodesStatusResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.ReportNodesStatusWithOptions(request, runtime)
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
func (client *Client) RunCommandWithOptions(tmpReq *RunCommandRequest, runtime *dara.RuntimeOptions) (_result *RunCommandResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &RunCommandShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeIdList) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, dara.String("NodeIdList"), dara.String("json"))
	}

	if !dara.IsNil(tmpReq.Parameters) {
		request.ParametersShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Parameters, dara.String("Parameters"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClientToken) {
		body["ClientToken"] = request.ClientToken
	}

	if !dara.IsNil(request.CommandContent) {
		body["CommandContent"] = request.CommandContent
	}

	if !dara.IsNil(request.CommandId) {
		body["CommandId"] = request.CommandId
	}

	if !dara.IsNil(request.ContentEncoding) {
		body["ContentEncoding"] = request.ContentEncoding
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.EnableParameter) {
		body["EnableParameter"] = request.EnableParameter
	}

	if !dara.IsNil(request.Frequency) {
		body["Frequency"] = request.Frequency
	}

	if !dara.IsNil(request.Launcher) {
		body["Launcher"] = request.Launcher
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeIdListShrink) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	if !dara.IsNil(request.ParametersShrink) {
		body["Parameters"] = request.ParametersShrink
	}

	if !dara.IsNil(request.RepeatMode) {
		body["RepeatMode"] = request.RepeatMode
	}

	if !dara.IsNil(request.TerminationMode) {
		body["TerminationMode"] = request.TerminationMode
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	if !dara.IsNil(request.Username) {
		body["Username"] = request.Username
	}

	if !dara.IsNil(request.WorkingDir) {
		body["WorkingDir"] = request.WorkingDir
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("RunCommand"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Execute a Shell script on one or more Lingjun machines
//
// @param request - RunCommandRequest
//
// @return RunCommandResponse
func (client *Client) RunCommand(request *RunCommandRequest) (_result *RunCommandResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) SendFileWithOptions(tmpReq *SendFileRequest, runtime *dara.RuntimeOptions) (_result *SendFileResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &SendFileShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeIdList) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, dara.String("NodeIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.Content) {
		body["Content"] = request.Content
	}

	if !dara.IsNil(request.ContentType) {
		body["ContentType"] = request.ContentType
	}

	if !dara.IsNil(request.Description) {
		body["Description"] = request.Description
	}

	if !dara.IsNil(request.FileGroup) {
		body["FileGroup"] = request.FileGroup
	}

	if !dara.IsNil(request.FileMode) {
		body["FileMode"] = request.FileMode
	}

	if !dara.IsNil(request.FileOwner) {
		body["FileOwner"] = request.FileOwner
	}

	if !dara.IsNil(request.Name) {
		body["Name"] = request.Name
	}

	if !dara.IsNil(request.NodeIdListShrink) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	if !dara.IsNil(request.Overwrite) {
		body["Overwrite"] = request.Overwrite
	}

	if !dara.IsNil(request.TargetDir) {
		body["TargetDir"] = request.TargetDir
	}

	if !dara.IsNil(request.Timeout) {
		body["Timeout"] = request.Timeout
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("SendFile"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Send a remote file to one or more Lingjun machines
//
// @param request - SendFileRequest
//
// @return SendFileResponse
func (client *Client) SendFile(request *SendFileRequest) (_result *SendFileResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) ShrinkClusterWithOptions(tmpReq *ShrinkClusterRequest, runtime *dara.RuntimeOptions) (_result *ShrinkClusterResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &ShrinkClusterShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeGroups) {
		request.NodeGroupsShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeGroups, dara.String("NodeGroups"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.ClusterId) {
		body["ClusterId"] = request.ClusterId
	}

	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.NodeGroupsShrink) {
		body["NodeGroups"] = request.NodeGroupsShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("ShrinkCluster"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Shrink
//
// @param request - ShrinkClusterRequest
//
// @return ShrinkClusterResponse
func (client *Client) ShrinkCluster(request *ShrinkClusterRequest) (_result *ShrinkClusterResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) StopInvocationWithOptions(tmpReq *StopInvocationRequest, runtime *dara.RuntimeOptions) (_result *StopInvocationResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StopInvocationShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.NodeIdList) {
		request.NodeIdListShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.NodeIdList, dara.String("NodeIdList"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.InvokeId) {
		body["InvokeId"] = request.InvokeId
	}

	if !dara.IsNil(request.NodeIdListShrink) {
		body["NodeIdList"] = request.NodeIdListShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopInvocation"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Stop the operation assistant command process
//
// @param request - StopInvocationRequest
//
// @return StopInvocationResponse
func (client *Client) StopInvocation(request *StopInvocationRequest) (_result *StopInvocationResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) StopNodesWithOptions(tmpReq *StopNodesRequest, runtime *dara.RuntimeOptions) (_result *StopNodesResponse, _err error) {
	_err = tmpReq.Validate()
	if _err != nil {
		return _result, _err
	}
	request := &StopNodesShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !dara.IsNil(tmpReq.Nodes) {
		request.NodesShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Nodes, dara.String("Nodes"), dara.String("json"))
	}

	body := map[string]interface{}{}
	if !dara.IsNil(request.IgnoreFailedNodeTasks) {
		body["IgnoreFailedNodeTasks"] = request.IgnoreFailedNodeTasks
	}

	if !dara.IsNil(request.NodesShrink) {
		body["Nodes"] = request.NodesShrink
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("StopNodes"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Shut down the nodes
//
// @param request - StopNodesRequest
//
// @return StopNodesResponse
func (client *Client) StopNodes(request *StopNodesRequest) (_result *StopNodesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) TagResourcesWithOptions(request *TagResourcesRequest, runtime *dara.RuntimeOptions) (_result *TagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.Tag) {
		query["Tag"] = request.Tag
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("TagResources"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Tag User Resources
//
// @param request - TagResourcesRequest
//
// @return TagResourcesResponse
func (client *Client) TagResources(request *TagResourcesRequest) (_result *TagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
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
func (client *Client) UntagResourcesWithOptions(request *UntagResourcesRequest, runtime *dara.RuntimeOptions) (_result *UntagResourcesResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !dara.IsNil(request.All) {
		query["All"] = request.All
	}

	if !dara.IsNil(request.RegionId) {
		query["RegionId"] = request.RegionId
	}

	if !dara.IsNil(request.ResourceId) {
		query["ResourceId"] = request.ResourceId
	}

	if !dara.IsNil(request.ResourceType) {
		query["ResourceType"] = request.ResourceType
	}

	if !dara.IsNil(request.Service) {
		query["Service"] = request.Service
	}

	if !dara.IsNil(request.TagKey) {
		query["TagKey"] = request.TagKey
	}

	req := &openapiutil.OpenApiRequest{
		Query: openapiutil.Query(query),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UntagResources"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// # Remove user tags from resources
//
// @param request - UntagResourcesRequest
//
// @return UntagResourcesResponse
func (client *Client) UntagResources(request *UntagResourcesRequest) (_result *UntagResourcesResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.UntagResourcesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新节点分组
//
// @param request - UpdateNodeGroupRequest
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateNodeGroupResponse
func (client *Client) UpdateNodeGroupWithOptions(request *UpdateNodeGroupRequest, runtime *dara.RuntimeOptions) (_result *UpdateNodeGroupResponse, _err error) {
	_err = request.Validate()
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !dara.IsNil(request.NewNodeGroupName) {
		body["NewNodeGroupName"] = request.NewNodeGroupName
	}

	if !dara.IsNil(request.NodeGroupId) {
		body["NodeGroupId"] = request.NodeGroupId
	}

	if !dara.IsNil(request.UserData) {
		body["UserData"] = request.UserData
	}

	req := &openapiutil.OpenApiRequest{
		Body: openapiutil.ParseToMap(body),
	}
	params := &openapiutil.Params{
		Action:      dara.String("UpdateNodeGroup"),
		Version:     dara.String("2022-12-15"),
		Protocol:    dara.String("HTTPS"),
		Pathname:    dara.String("/"),
		Method:      dara.String("POST"),
		AuthType:    dara.String("AK"),
		Style:       dara.String("RPC"),
		ReqBodyType: dara.String("formData"),
		BodyType:    dara.String("json"),
	}
	if dara.IsNil(client.SignatureVersion) || dara.StringValue(client.SignatureVersion) != "v4" {
		_body, _err := client.CallApi(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	} else {
		_body, _err := client.Execute(params, req, runtime)
		if _err != nil {
			return _result, _err
		}
		_err = dara.Convert(_body, &_result)
		return _result, _err
	}

}

// Summary:
//
// 更新节点分组
//
// @param request - UpdateNodeGroupRequest
//
// @return UpdateNodeGroupResponse
func (client *Client) UpdateNodeGroup(request *UpdateNodeGroupRequest) (_result *UpdateNodeGroupResponse, _err error) {
	runtime := &dara.RuntimeOptions{}
	_body, _err := client.UpdateNodeGroupWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
