// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProtocolMountTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeProtocolMountTargetResponseBody
	GetNextToken() *string
	SetProtocolMountTargets(v []*DescribeProtocolMountTargetResponseBodyProtocolMountTargets) *DescribeProtocolMountTargetResponseBody
	GetProtocolMountTargets() []*DescribeProtocolMountTargetResponseBodyProtocolMountTargets
	SetRequestId(v string) *DescribeProtocolMountTargetResponseBody
	GetRequestId() *string
}

type DescribeProtocolMountTargetResponseBody struct {
	// The marker used to retrieve the remaining export directories in subsequent queries.
	//
	// example:
	//
	// aBcdeg==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The collection of protocol service export directories.
	ProtocolMountTargets []*DescribeProtocolMountTargetResponseBodyProtocolMountTargets `json:"ProtocolMountTargets,omitempty" xml:"ProtocolMountTargets,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeProtocolMountTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProtocolMountTargetResponseBody) GetProtocolMountTargets() []*DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	return s.ProtocolMountTargets
}

func (s *DescribeProtocolMountTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProtocolMountTargetResponseBody) SetNextToken(v string) *DescribeProtocolMountTargetResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBody) SetProtocolMountTargets(v []*DescribeProtocolMountTargetResponseBodyProtocolMountTargets) *DescribeProtocolMountTargetResponseBody {
	s.ProtocolMountTargets = v
	return s
}

func (s *DescribeProtocolMountTargetResponseBody) SetRequestId(v string) *DescribeProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBody) Validate() error {
	if s.ProtocolMountTargets != nil {
		for _, item := range s.ProtocolMountTargets {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeProtocolMountTargetResponseBodyProtocolMountTargets struct {
	// The permission group associated with the protocol service export directory.
	//
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// The time when the protocol service export directory was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format.
	//
	// example:
	//
	// 2018-12-12T07:28:38Z
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The description of the protocol service export directory.
	//
	// example:
	//
	// 此协议服务导出目录的描述。
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the protocol service export directory.
	//
	// example:
	//
	// exp-19abf5beab8d****
	ExportId *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	// The fileset ID of the protocol service export directory.
	//
	// example:
	//
	// fset-1902718ea0ae****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// The exported directory of the protocol service.
	//
	// example:
	//
	// /path/
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// The domain name of the protocol service export directory.
	//
	// example:
	//
	// cpfs-123****.cn-hangzhou.cpfs.aliyuncs.com
	ProtocolMountTargetDomain *string `json:"ProtocolMountTargetDomain,omitempty" xml:"ProtocolMountTargetDomain,omitempty"`
	// The protocol service ID.
	//
	// example:
	//
	// ptc-123****
	ProtocolServiceId *string `json:"ProtocolServiceId,omitempty" xml:"ProtocolServiceId,omitempty"`
	// The protocol type supported by the protocol service.
	//
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// The mount target status.
	//
	// example:
	//
	// CREATING
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The vSwitch ID of the protocol service export directory.
	//
	// example:
	//
	// vsw-2vc3c2lybvdllxyq4****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The list of zone-redundant vSwitch IDs.
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// The VPC ID of the protocol service export directory.
	//
	// example:
	//
	// vpc-2vct297b8157bth9z****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeProtocolMountTargetResponseBodyProtocolMountTargets) String() string {
	return dara.Prettify(s)
}

func (s DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GoString() string {
	return s.String()
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetCreateTime() *string {
	return s.CreateTime
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetDescription() *string {
	return s.Description
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetExportId() *string {
	return s.ExportId
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetFsetId() *string {
	return s.FsetId
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetPath() *string {
	return s.Path
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetProtocolMountTargetDomain() *string {
	return s.ProtocolMountTargetDomain
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetProtocolServiceId() *string {
	return s.ProtocolServiceId
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetStatus() *string {
	return s.Status
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetAccessGroupName(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.AccessGroupName = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetCreateTime(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.CreateTime = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetDescription(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.Description = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetExportId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ExportId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetFsetId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.FsetId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetPath(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.Path = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetProtocolMountTargetDomain(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ProtocolMountTargetDomain = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetProtocolServiceId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ProtocolServiceId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetProtocolType(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.ProtocolType = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetStatus(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.Status = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetVSwitchId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.VSwitchId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetVSwitchIds(v []*string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.VSwitchIds = v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) SetVpcId(v string) *DescribeProtocolMountTargetResponseBodyProtocolMountTargets {
	s.VpcId = &v
	return s
}

func (s *DescribeProtocolMountTargetResponseBodyProtocolMountTargets) Validate() error {
	return dara.Validate(s)
}
