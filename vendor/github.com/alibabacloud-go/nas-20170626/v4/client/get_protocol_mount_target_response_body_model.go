// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProtocolMountTargetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *GetProtocolMountTargetResponseBody
	GetNextToken() *string
	SetProtocolMountTarget(v *GetProtocolMountTargetResponseBodyProtocolMountTarget) *GetProtocolMountTargetResponseBody
	GetProtocolMountTarget() *GetProtocolMountTargetResponseBodyProtocolMountTarget
	SetRequestId(v string) *GetProtocolMountTargetResponseBody
	GetRequestId() *string
}

type GetProtocolMountTargetResponseBody struct {
	// example:
	//
	// M18xMA==
	NextToken           *string                                                `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ProtocolMountTarget *GetProtocolMountTargetResponseBodyProtocolMountTarget `json:"ProtocolMountTarget,omitempty" xml:"ProtocolMountTarget,omitempty" type:"Struct"`
	// example:
	//
	// 6299428C-3861-435D-AE54-9B330A00****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetProtocolMountTargetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetProtocolMountTargetResponseBody) GoString() string {
	return s.String()
}

func (s *GetProtocolMountTargetResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *GetProtocolMountTargetResponseBody) GetProtocolMountTarget() *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	return s.ProtocolMountTarget
}

func (s *GetProtocolMountTargetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetProtocolMountTargetResponseBody) SetNextToken(v string) *GetProtocolMountTargetResponseBody {
	s.NextToken = &v
	return s
}

func (s *GetProtocolMountTargetResponseBody) SetProtocolMountTarget(v *GetProtocolMountTargetResponseBodyProtocolMountTarget) *GetProtocolMountTargetResponseBody {
	s.ProtocolMountTarget = v
	return s
}

func (s *GetProtocolMountTargetResponseBody) SetRequestId(v string) *GetProtocolMountTargetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetProtocolMountTargetResponseBody) Validate() error {
	if s.ProtocolMountTarget != nil {
		if err := s.ProtocolMountTarget.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetProtocolMountTargetResponseBodyProtocolMountTarget struct {
	// example:
	//
	// DEFAULT_VPC_GROUP_NAME
	AccessGroupName *string `json:"AccessGroupName,omitempty" xml:"AccessGroupName,omitempty"`
	// example:
	//
	// 2025-12-22 17:49:25
	CreateTime  *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// exp-19abf5beab8d****
	ExportId *string `json:"ExportId,omitempty" xml:"ExportId,omitempty"`
	// Fileset ID。
	//
	// example:
	//
	// fset-299b4ca04de8****
	FsetId *string `json:"FsetId,omitempty" xml:"FsetId,omitempty"`
	// example:
	//
	// /
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// example:
	//
	// cpfs-0229cb80bcc0****-x******.cn-*****.cpfs.aliyuncs.com
	ProtocolMountTargetDomain *string `json:"ProtocolMountTargetDomain,omitempty" xml:"ProtocolMountTargetDomain,omitempty"`
	// example:
	//
	// NFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// example:
	//
	// AVAILABLE
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-8vb2qjnxs6hiobzve****
	VSwitchId  *string   `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VSwitchIds []*string `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Repeated"`
	// example:
	//
	// vpc-bp1h5mxoqfuo3xurf****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetProtocolMountTargetResponseBodyProtocolMountTarget) String() string {
	return dara.Prettify(s)
}

func (s GetProtocolMountTargetResponseBodyProtocolMountTarget) GoString() string {
	return s.String()
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetAccessGroupName() *string {
	return s.AccessGroupName
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetCreateTime() *string {
	return s.CreateTime
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetDescription() *string {
	return s.Description
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetExportId() *string {
	return s.ExportId
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetFsetId() *string {
	return s.FsetId
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetPath() *string {
	return s.Path
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetProtocolMountTargetDomain() *string {
	return s.ProtocolMountTargetDomain
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetStatus() *string {
	return s.Status
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetVSwitchIds() []*string {
	return s.VSwitchIds
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) GetVpcId() *string {
	return s.VpcId
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetAccessGroupName(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.AccessGroupName = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetCreateTime(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.CreateTime = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetDescription(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.Description = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetExportId(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.ExportId = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetFsetId(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.FsetId = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetPath(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.Path = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetProtocolMountTargetDomain(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.ProtocolMountTargetDomain = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetProtocolType(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.ProtocolType = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetStatus(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.Status = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetVSwitchId(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.VSwitchId = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetVSwitchIds(v []*string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.VSwitchIds = v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) SetVpcId(v string) *GetProtocolMountTargetResponseBodyProtocolMountTarget {
	s.VpcId = &v
	return s
}

func (s *GetProtocolMountTargetResponseBodyProtocolMountTarget) Validate() error {
	return dara.Validate(s)
}
