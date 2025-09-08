// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

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
	SetNetworks(v *DescribeClusterResponseBodyNetworks) *DescribeClusterResponseBody
	GetNetworks() *DescribeClusterResponseBodyNetworks
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
	// The cluster description.
	//
	// example:
	//
	// Test cluster
	ClusterDescription *string `json:"ClusterDescription,omitempty" xml:"ClusterDescription,omitempty"`
	// The cluster ID.
	//
	// example:
	//
	// i116913051662373010974
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// Eflo-YJ-Test-Cluster
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The cluster type.
	//
	// example:
	//
	// AckEdgePro
	ClusterType *string `json:"ClusterType,omitempty" xml:"ClusterType,omitempty"`
	// The component information.
	Components []*DescribeClusterResponseBodyComponents `json:"Components,omitempty" xml:"Components,omitempty" type:"Repeated"`
	// The IP type of the computing network.
	//
	// example:
	//
	// IPv4
	ComputingIpVersion *string `json:"ComputingIpVersion,omitempty" xml:"ComputingIpVersion,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2022-06-08T07:05:11Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// A2
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// The network information.
	Networks *DescribeClusterResponseBodyNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Struct"`
	// The number of nodes.
	//
	// example:
	//
	// 2
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The number of node groups.
	//
	// example:
	//
	// 2
	NodeGroupCount *int64 `json:"NodeGroupCount,omitempty" xml:"NodeGroupCount,omitempty"`
	// The status of Jumbo Frames for the elastic network interface (ENI).
	//
	// example:
	//
	// unsupported
	OpenEniJumboFrame *string `json:"OpenEniJumboFrame,omitempty" xml:"OpenEniJumboFrame,omitempty"`
	// The cluster status.
	//
	// example:
	//
	// running
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 887FA855-89F4-5DB3-B305-C5879EC480E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The resource group ID.
	//
	// example:
	//
	// rg-aek2k3rqlvv6ytq
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The job ID.
	//
	// example:
	//
	// i152609221670466904596
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2022-08-23T06:36:17.000Z
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The ID of the virtual private cloud (VPC).
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

func (s *DescribeClusterResponseBody) GetNetworks() *DescribeClusterResponseBodyNetworks {
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

func (s *DescribeClusterResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeClusterResponseBodyComponents struct {
	// The component ID.
	//
	// example:
	//
	// i149549021660892626529
	ComponentId *string `json:"ComponentId,omitempty" xml:"ComponentId,omitempty"`
	// The component type.
	//
	// Valid values:
	//
	// 	- ARMS
	//
	// 	- ACKEdge
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
	// The ID of the CIDR block for the cluster.
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
