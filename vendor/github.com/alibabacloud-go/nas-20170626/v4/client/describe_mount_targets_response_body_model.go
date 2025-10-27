// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMountTargetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMountTargets(v *DescribeMountTargetsResponseBodyMountTargets) *DescribeMountTargetsResponseBody
	GetMountTargets() *DescribeMountTargetsResponseBodyMountTargets
	SetPageNumber(v int32) *DescribeMountTargetsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeMountTargetsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeMountTargetsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeMountTargetsResponseBody
	GetTotalCount() *int32
}

type DescribeMountTargetsResponseBody struct {
	// The queried mount targets.
	MountTargets *DescribeMountTargetsResponseBodyMountTargets `json:"MountTargets,omitempty" xml:"MountTargets,omitempty" type:"Struct"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 3BAB90FD-B4A0-48DA-9F09-2B963510****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of mount targets.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeMountTargetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBody) GetMountTargets() *DescribeMountTargetsResponseBodyMountTargets {
	return s.MountTargets
}

func (s *DescribeMountTargetsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeMountTargetsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeMountTargetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMountTargetsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeMountTargetsResponseBody) SetMountTargets(v *DescribeMountTargetsResponseBodyMountTargets) *DescribeMountTargetsResponseBody {
	s.MountTargets = v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetPageNumber(v int32) *DescribeMountTargetsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetPageSize(v int32) *DescribeMountTargetsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetRequestId(v string) *DescribeMountTargetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) SetTotalCount(v int32) *DescribeMountTargetsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeMountTargetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeMountTargetsResponseBodyMountTargets struct {
	MountTarget []*DescribeMountTargetsResponseBodyMountTargetsMountTarget `json:"MountTarget,omitempty" xml:"MountTarget,omitempty" type:"Repeated"`
}

func (s DescribeMountTargetsResponseBodyMountTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargets) GetMountTarget() []*DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	return s.MountTarget
}

func (s *DescribeMountTargetsResponseBodyMountTargets) SetMountTarget(v []*DescribeMountTargetsResponseBodyMountTargetsMountTarget) *DescribeMountTargetsResponseBodyMountTargets {
	s.MountTarget = v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargets) Validate() error {
	return dara.Validate(s)
}

type DescribeMountTargetsResponseBodyMountTargetsMountTarget struct {
	// The name of the permission group that is attached to the mount target.
	//
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroup *string `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	// The information about client management nodes.
	ClientMasterNodes *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes `json:"ClientMasterNodes,omitempty" xml:"ClientMasterNodes,omitempty" type:"Struct"`
	// The dual-stack (IPv4 and IPv6) domain name of the mount target.
	//
	// example:
	//
	// 1ca404****-x****.dualstack.cn-hangzhou.nas.aliyuncs.com
	DualStackMountTargetDomain *string `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	// The type of the mount target.
	//
	// 	- IPv4: an IPv4 mount target
	//
	// 	- DualStack: a dual-stack mount target
	//
	// example:
	//
	// IPv4
	IPVersion *string `json:"IPVersion,omitempty" xml:"IPVersion,omitempty"`
	// The IPv4 domain name of the mount target.
	//
	// example:
	//
	// 1ca404****-w****.cn-hangzhou.nas.aliyuncs.com
	MountTargetDomain *string `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	// The network type. Valid value: **Vpc**.
	//
	// example:
	//
	// Vpc
	NetworkType *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	// The status of the mount target.
	//
	// Valid values:
	//
	// 	- Active: The mount target is available.
	//
	// 	- Inactive: The mount target is unavailable.
	//
	// 	- Pending: The mount target is being created or modified.
	//
	// 	- Deleting: The mount target is being deleted.
	//
	// 	- Hibernating: The mount target is being hibernated.
	//
	// 	- Hibernated: The mount target is hibernated.
	//
	// > You can mount a file system only when the mount target of the file system is in the Active state.
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// An array of tags. The array may contain up to 20 tags. If the array contains multiple tags, each tag key is unique.
	Tags *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-2zesj9afh3y518k9o****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the vSwitch.
	//
	// example:
	//
	// vsw-2zevmwkwyztjuoffg****
	VswId *string `json:"VswId,omitempty" xml:"VswId,omitempty"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTarget) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTarget) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) GetAccessGroup() *string {
	return s.AccessGroup
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) GetClientMasterNodes() *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes {
	return s.ClientMasterNodes
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) GetDualStackMountTargetDomain() *string {
	return s.DualStackMountTargetDomain
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) GetIPVersion() *string {
	return s.IPVersion
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) GetMountTargetDomain() *string {
	return s.MountTargetDomain
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) GetStatus() *string {
	return s.Status
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) GetTags() *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags {
	return s.Tags
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) GetVswId() *string {
	return s.VswId
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetAccessGroup(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.AccessGroup = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetClientMasterNodes(v *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.ClientMasterNodes = v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetDualStackMountTargetDomain(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.DualStackMountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetIPVersion(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.IPVersion = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetMountTargetDomain(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.MountTargetDomain = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetNetworkType(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.NetworkType = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetStatus(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.Status = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetTags(v *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.Tags = v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetVpcId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.VpcId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) SetVswId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTarget {
	s.VswId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTarget) Validate() error {
	return dara.Validate(s)
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes struct {
	ClientMasterNode []*DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode `json:"ClientMasterNode,omitempty" xml:"ClientMasterNode,omitempty" type:"Repeated"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) GetClientMasterNode() []*DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	return s.ClientMasterNode
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) SetClientMasterNode(v []*DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes {
	s.ClientMasterNode = v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes) Validate() error {
	return dara.Validate(s)
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode struct {
	// The default logon password of the ECS instance.
	//
	// example:
	//
	// 12****
	DefaultPasswd *string `json:"DefaultPasswd,omitempty" xml:"DefaultPasswd,omitempty"`
	// The ID of the ECS instance on the client management node.
	//
	// example:
	//
	// i-hp3i3odi5ory1buo****
	EcsId *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	// The IP address of the ECS instance on the client management node.
	//
	// example:
	//
	// 192.168.1.0
	EcsIp *string `json:"EcsIp,omitempty" xml:"EcsIp,omitempty"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) GetDefaultPasswd() *string {
	return s.DefaultPasswd
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) GetEcsId() *string {
	return s.EcsId
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) GetEcsIp() *string {
	return s.EcsIp
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetDefaultPasswd(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.DefaultPasswd = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsId(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsId = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) SetEcsIp(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode {
	s.EcsIp = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode) Validate() error {
	return dara.Validate(s)
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetTags struct {
	Tag []*DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetTags) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags) GetTag() []*DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag {
	return s.Tag
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags) SetTag(v []*DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags {
	s.Tag = v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags) Validate() error {
	return dara.Validate(s)
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag struct {
	// The tag key. Limits:
	//
	// 	- The tag key cannot be null or an empty string.
	//
	// 	- The tag key can be up to 128 characters in length.
	//
	// 	- The key value cannot start with aliyun or acs:.
	//
	// 	- The key value cannot contain http:// or https://.
	//
	// example:
	//
	// nastest
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The tag value.
	//
	// Limits:
	//
	// 	- The tag value can be up to 128 characters in length.
	//
	// 	- The tag value cannot contain http:// or https://.
	//
	// example:
	//
	// mounttargettest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) GetKey() *string {
	return s.Key
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) GetValue() *string {
	return s.Value
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) SetKey(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag {
	s.Key = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) SetValue(v string) *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag {
	s.Value = &v
	return s
}

func (s *DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag) Validate() error {
	return dara.Validate(s)
}
