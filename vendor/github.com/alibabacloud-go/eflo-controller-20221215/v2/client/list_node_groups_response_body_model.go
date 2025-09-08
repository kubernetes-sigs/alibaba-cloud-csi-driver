// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

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
	// The node groups.
	Groups []*ListNodeGroupsResponseBodyGroups `json:"Groups,omitempty" xml:"Groups,omitempty" type:"Repeated"`
	// The token that is used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// 563d42ae0b17572449ec8c97f7f66069
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
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
	// The cluster ID.
	//
	// example:
	//
	// i113952461729854708648
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The cluster name.
	//
	// example:
	//
	// wzq-exclusivelite-71
	ClusterName *string `json:"ClusterName,omitempty" xml:"ClusterName,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2024-02-27T13:16:31.599
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description.
	//
	// example:
	//
	// created by ga2_prepare
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Indicates whether file storage mounting is supported.
	//
	// example:
	//
	// False
	FileSystemMountEnabled *bool `json:"FileSystemMountEnabled,omitempty" xml:"FileSystemMountEnabled,omitempty"`
	// The group ID.
	//
	// example:
	//
	// 238276221
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The group name.
	//
	// example:
	//
	// backend-group
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The image ID.
	//
	// example:
	//
	// i194015071707321240258
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The image name.
	//
	// example:
	//
	// CentOS_7.9_x86_64_FULL_20221110
	ImageName *string `json:"ImageName,omitempty" xml:"ImageName,omitempty"`
	// The instance type.
	//
	// example:
	//
	// efg1.nvga1n
	MachineType *string `json:"MachineType,omitempty" xml:"MachineType,omitempty"`
	// The number of nodes.
	//
	// example:
	//
	// 2
	NodeCount *int64 `json:"NodeCount,omitempty" xml:"NodeCount,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2023-09-22T00:03:05.114
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// Whether to enable gpu virtualization or not
	//
	// example:
	//
	// false
	VirtualGpuEnabled *bool `json:"VirtualGpuEnabled,omitempty" xml:"VirtualGpuEnabled,omitempty"`
	// The zone ID.
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

func (s *ListNodeGroupsResponseBodyGroups) GetFileSystemMountEnabled() *bool {
	return s.FileSystemMountEnabled
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

func (s *ListNodeGroupsResponseBodyGroups) GetVirtualGpuEnabled() *bool {
	return s.VirtualGpuEnabled
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

func (s *ListNodeGroupsResponseBodyGroups) SetVirtualGpuEnabled(v bool) *ListNodeGroupsResponseBodyGroups {
	s.VirtualGpuEnabled = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) SetZoneId(v string) *ListNodeGroupsResponseBodyGroups {
	s.ZoneId = &v
	return s
}

func (s *ListNodeGroupsResponseBodyGroups) Validate() error {
	return dara.Validate(s)
}
