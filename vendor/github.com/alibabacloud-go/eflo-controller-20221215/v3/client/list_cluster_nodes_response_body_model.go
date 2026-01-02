// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

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
	// The returned pagination token which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAXW/ZB9TBvH+0ZK0phtCibQgQmu1RbqplAI6Velo2OKR
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The nodes.
	Nodes []*ListClusterNodesResponseBodyNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The request ID.
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

type ListClusterNodesResponseBodyNodes struct {
	// The commodity code.
	//
	// example:
	//
	// bcccluster
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 1642472468000
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The time when the node expires.
	//
	// example:
	//
	// 1762185600000
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// Indicates whether file storage mounting is supported.
	//
	// example:
	//
	// False
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The hostname.
	//
	// example:
	//
	// 72432f80-273e-11ed-b57a-acde48001122
	Hostname *string `json:"Hostname,omitempty" xml:"Hostname,omitempty"`
	// The cluster number.
	//
	// example:
	//
	// A1
	HpnZone *string `json:"HpnZone,omitempty" xml:"HpnZone,omitempty"`
	// example:
	//
	// e01-cn-2r42tmj4z02
	HyperNodeId *string `json:"HyperNodeId,omitempty" xml:"HyperNodeId,omitempty"`
	// The system image ID.
	//
	// example:
	//
	// i190297201669099844192
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// Alinux3_x86_AMD_R_Host_D3_E3_24.13.00_UEFI_N_250121
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The node type.
	//
	// example:
	//
	// cn-wulanchabu-b11
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The network information.
	Networks []*ListClusterNodesResponseBodyNodesNetworks `json:"Networks,omitempty" xml:"Networks,omitempty" type:"Repeated"`
	// The node group ID.
	//
	// example:
	//
	// ng-e9b74f4d450cf18d
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The node group name.
	//
	// example:
	//
	// emr_master
	NodeGroupName *string `json:"NodeGroupName,omitempty" xml:"NodeGroupName,omitempty"`
	// The node ID.
	//
	// example:
	//
	// e01-cn-2r42tmj4z02
	NodeId   *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	NodeType *string `json:"NodeType,omitempty" xml:"NodeType,omitempty"`
	// The node status.
	//
	// Valid values:
	//
	// 	- Extending
	//
	// 	- UnusedNodeStopped
	//
	// 	- UnusedNodeStopping
	//
	// 	- Unused
	//
	// 	- Using
	//
	// 	- ReleaseLocking
	//
	// 	- Operating
	//
	// 	- Cutting
	//
	// 	- ClusterNodeStopped
	//
	// 	- UnusedNodeRecovering
	//
	// 	- ClusterNodeStopping
	//
	// 	- ClusterNodeRecovering
	//
	// 	- Replacing
	//
	// example:
	//
	// Extending
	OperatingState *string `json:"OperatingState,omitempty" xml:"OperatingState,omitempty"`
	// The serial number of the node.
	//
	// example:
	//
	// sn_tOuUk
	Sn *string `json:"Sn,omitempty" xml:"Sn,omitempty"`
	// The tags.
	Tags []*ListClusterNodesResponseBodyNodesTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
	// The job ID.
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
	// The VPC ID.
	//
	// example:
	//
	// vpc-0jltf9vinjz3if3lltdy7
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The zone ID.
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

func (s *ListClusterNodesResponseBodyNodes) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
}

func (s *ListClusterNodesResponseBodyNodes) GetHostname() *string {
	return s.Hostname
}

func (s *ListClusterNodesResponseBodyNodes) GetHpnZone() *string {
	return s.HpnZone
}

func (s *ListClusterNodesResponseBodyNodes) GetHyperNodeId() *string {
	return s.HyperNodeId
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

func (s *ListClusterNodesResponseBodyNodes) GetNodeType() *string {
	return s.NodeType
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

func (s *ListClusterNodesResponseBodyNodes) SetHyperNodeId(v string) *ListClusterNodesResponseBodyNodes {
	s.HyperNodeId = &v
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

func (s *ListClusterNodesResponseBodyNodes) SetNodeType(v string) *ListClusterNodesResponseBodyNodes {
	s.NodeType = &v
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
	if s.Networks != nil {
		for _, item := range s.Networks {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListClusterNodesResponseBodyNodesNetworks struct {
	// The name of the network port for the node.
	//
	// example:
	//
	// bond0
	BondName *string `json:"BondName,omitempty" xml:"BondName,omitempty"`
	// The IP address of the node in the virtual private cloud (VPC).
	//
	// example:
	//
	// 192.168.22.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// The subnet ID.
	//
	// example:
	//
	// subnet-fwekrvg9
	SubnetId *string `json:"SubnetId,omitempty" xml:"SubnetId,omitempty"`
	// The VPC ID.
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
	// The tag key.
	//
	// example:
	//
	// aa_key
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// example:
	//
	// aa_value
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
