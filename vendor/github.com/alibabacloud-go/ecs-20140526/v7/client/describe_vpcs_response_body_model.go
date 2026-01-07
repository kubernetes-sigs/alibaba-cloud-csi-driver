// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVpcsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPageNumber(v int32) *DescribeVpcsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeVpcsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeVpcsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeVpcsResponseBody
	GetTotalCount() *int32
	SetVpcs(v *DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody
	GetVpcs() *DescribeVpcsResponseBodyVpcs
}

type DescribeVpcsResponseBody struct {
	PageNumber *int32                        `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize   *int32                        `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Vpcs       *DescribeVpcsResponseBodyVpcs `json:"Vpcs,omitempty" xml:"Vpcs,omitempty" type:"Struct"`
}

func (s DescribeVpcsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeVpcsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeVpcsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVpcsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeVpcsResponseBody) GetVpcs() *DescribeVpcsResponseBodyVpcs {
	return s.Vpcs
}

func (s *DescribeVpcsResponseBody) SetPageNumber(v int32) *DescribeVpcsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetPageSize(v int32) *DescribeVpcsResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetRequestId(v string) *DescribeVpcsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetTotalCount(v int32) *DescribeVpcsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeVpcsResponseBody) SetVpcs(v *DescribeVpcsResponseBodyVpcs) *DescribeVpcsResponseBody {
	s.Vpcs = v
	return s
}

func (s *DescribeVpcsResponseBody) Validate() error {
	if s.Vpcs != nil {
		if err := s.Vpcs.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpcsResponseBodyVpcs struct {
	Vpc []*DescribeVpcsResponseBodyVpcsVpc `json:"Vpc,omitempty" xml:"Vpc,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcs) GetVpc() []*DescribeVpcsResponseBodyVpcsVpc {
	return s.Vpc
}

func (s *DescribeVpcsResponseBodyVpcs) SetVpc(v []*DescribeVpcsResponseBodyVpcsVpc) *DescribeVpcsResponseBodyVpcs {
	s.Vpc = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcs) Validate() error {
	if s.Vpc != nil {
		for _, item := range s.Vpc {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeVpcsResponseBodyVpcsVpc struct {
	CidrBlock    *string                                    `json:"CidrBlock,omitempty" xml:"CidrBlock,omitempty"`
	CreationTime *string                                    `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	Description  *string                                    `json:"Description,omitempty" xml:"Description,omitempty"`
	IsDefault    *bool                                      `json:"IsDefault,omitempty" xml:"IsDefault,omitempty"`
	RegionId     *string                                    `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	Status       *string                                    `json:"Status,omitempty" xml:"Status,omitempty"`
	UserCidrs    *DescribeVpcsResponseBodyVpcsVpcUserCidrs  `json:"UserCidrs,omitempty" xml:"UserCidrs,omitempty" type:"Struct"`
	VRouterId    *string                                    `json:"VRouterId,omitempty" xml:"VRouterId,omitempty"`
	VSwitchIds   *DescribeVpcsResponseBodyVpcsVpcVSwitchIds `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty" type:"Struct"`
	VpcId        *string                                    `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	VpcName      *string                                    `json:"VpcName,omitempty" xml:"VpcName,omitempty"`
}

func (s DescribeVpcsResponseBodyVpcsVpc) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpc) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetCidrBlock() *string {
	return s.CidrBlock
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetDescription() *string {
	return s.Description
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetStatus() *string {
	return s.Status
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetUserCidrs() *DescribeVpcsResponseBodyVpcsVpcUserCidrs {
	return s.UserCidrs
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVRouterId() *string {
	return s.VRouterId
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVSwitchIds() *DescribeVpcsResponseBodyVpcsVpcVSwitchIds {
	return s.VSwitchIds
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeVpcsResponseBodyVpcsVpc) GetVpcName() *string {
	return s.VpcName
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetCidrBlock(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.CidrBlock = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetCreationTime(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.CreationTime = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetDescription(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.Description = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetIsDefault(v bool) *DescribeVpcsResponseBodyVpcsVpc {
	s.IsDefault = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetRegionId(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.RegionId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetStatus(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.Status = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetUserCidrs(v *DescribeVpcsResponseBodyVpcsVpcUserCidrs) *DescribeVpcsResponseBodyVpcsVpc {
	s.UserCidrs = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVRouterId(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.VRouterId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVSwitchIds(v *DescribeVpcsResponseBodyVpcsVpcVSwitchIds) *DescribeVpcsResponseBodyVpcsVpc {
	s.VSwitchIds = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVpcId(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.VpcId = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) SetVpcName(v string) *DescribeVpcsResponseBodyVpcsVpc {
	s.VpcName = &v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpc) Validate() error {
	if s.UserCidrs != nil {
		if err := s.UserCidrs.Validate(); err != nil {
			return err
		}
	}
	if s.VSwitchIds != nil {
		if err := s.VSwitchIds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeVpcsResponseBodyVpcsVpcUserCidrs struct {
	UserCidr []*string `json:"UserCidr,omitempty" xml:"UserCidr,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcsVpcUserCidrs) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcUserCidrs) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcUserCidrs) GetUserCidr() []*string {
	return s.UserCidr
}

func (s *DescribeVpcsResponseBodyVpcsVpcUserCidrs) SetUserCidr(v []*string) *DescribeVpcsResponseBodyVpcsVpcUserCidrs {
	s.UserCidr = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcUserCidrs) Validate() error {
	return dara.Validate(s)
}

type DescribeVpcsResponseBodyVpcsVpcVSwitchIds struct {
	VSwitchId []*string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty" type:"Repeated"`
}

func (s DescribeVpcsResponseBodyVpcsVpcVSwitchIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeVpcsResponseBodyVpcsVpcVSwitchIds) GoString() string {
	return s.String()
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchIds) GetVSwitchId() []*string {
	return s.VSwitchId
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchIds) SetVSwitchId(v []*string) *DescribeVpcsResponseBodyVpcsVpcVSwitchIds {
	s.VSwitchId = v
	return s
}

func (s *DescribeVpcsResponseBodyVpcsVpcVSwitchIds) Validate() error {
	return dara.Validate(s)
}
