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
	if s.MountTargets != nil {
		if err := s.MountTargets.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.MountTarget != nil {
		for _, item := range s.MountTarget {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMountTargetsResponseBodyMountTargetsMountTarget struct {
	AccessGroup                *string                                                                   `json:"AccessGroup,omitempty" xml:"AccessGroup,omitempty"`
	ClientMasterNodes          *DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodes `json:"ClientMasterNodes,omitempty" xml:"ClientMasterNodes,omitempty" type:"Struct"`
	DualStackMountTargetDomain *string                                                                   `json:"DualStackMountTargetDomain,omitempty" xml:"DualStackMountTargetDomain,omitempty"`
	IPVersion                  *string                                                                   `json:"IPVersion,omitempty" xml:"IPVersion,omitempty"`
	MountTargetDomain          *string                                                                   `json:"MountTargetDomain,omitempty" xml:"MountTargetDomain,omitempty"`
	NetworkType                *string                                                                   `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	Status                     *string                                                                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tags                       *DescribeMountTargetsResponseBodyMountTargetsMountTargetTags              `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	VpcId                      *string                                                                   `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VswId                      *string                                                                   `json:"VswId,omitempty" xml:"VswId,omitempty"`
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
	if s.ClientMasterNodes != nil {
		if err := s.ClientMasterNodes.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
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
	if s.ClientMasterNode != nil {
		for _, item := range s.ClientMasterNode {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeMountTargetsResponseBodyMountTargetsMountTargetClientMasterNodesClientMasterNode struct {
	DefaultPasswd *string `json:"DefaultPasswd,omitempty" xml:"DefaultPasswd,omitempty"`
	EcsId         *string `json:"EcsId,omitempty" xml:"EcsId,omitempty"`
	EcsIp         *string `json:"EcsIp,omitempty" xml:"EcsIp,omitempty"`
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

type DescribeMountTargetsResponseBodyMountTargetsMountTargetTagsTag struct {
	Key   *string `json:"Key,omitempty" xml:"Key,omitempty"`
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
