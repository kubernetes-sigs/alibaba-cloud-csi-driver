// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceRamRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceRamRoleSets(v *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets) *DescribeInstanceRamRoleResponseBody
	GetInstanceRamRoleSets() *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets
	SetRegionId(v string) *DescribeInstanceRamRoleResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeInstanceRamRoleResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceRamRoleResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceRamRoleResponseBody struct {
	// The IDs of the ECS instances and the names of the corresponding instance RAM roles.
	InstanceRamRoleSets *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets `json:"InstanceRamRoleSets,omitempty" xml:"InstanceRamRoleSets,omitempty" type:"Struct"`
	// The region ID of the ECS instances.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The number of ECS instances returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceRamRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRamRoleResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRamRoleResponseBody) GetInstanceRamRoleSets() *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets {
	return s.InstanceRamRoleSets
}

func (s *DescribeInstanceRamRoleResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInstanceRamRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceRamRoleResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceRamRoleResponseBody) SetInstanceRamRoleSets(v *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets) *DescribeInstanceRamRoleResponseBody {
	s.InstanceRamRoleSets = v
	return s
}

func (s *DescribeInstanceRamRoleResponseBody) SetRegionId(v string) *DescribeInstanceRamRoleResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeInstanceRamRoleResponseBody) SetRequestId(v string) *DescribeInstanceRamRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceRamRoleResponseBody) SetTotalCount(v int32) *DescribeInstanceRamRoleResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceRamRoleResponseBody) Validate() error {
	if s.InstanceRamRoleSets != nil {
		if err := s.InstanceRamRoleSets.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets struct {
	InstanceRamRoleSet []*DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet `json:"InstanceRamRoleSet,omitempty" xml:"InstanceRamRoleSet,omitempty" type:"Repeated"`
}

func (s DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets) GetInstanceRamRoleSet() []*DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet {
	return s.InstanceRamRoleSet
}

func (s *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets) SetInstanceRamRoleSet(v []*DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet) *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets {
	s.InstanceRamRoleSet = v
	return s
}

func (s *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSets) Validate() error {
	if s.InstanceRamRoleSet != nil {
		for _, item := range s.InstanceRamRoleSet {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet struct {
	// The ID of the instance
	//
	// example:
	//
	// i-bp67acfmxazb4p****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The name of the instance RAM role.
	//
	// example:
	//
	// EcsServiceRole-EcsDocGuideTest
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
}

func (s DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet) GoString() string {
	return s.String()
}

func (s *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet) SetInstanceId(v string) *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet) SetRamRoleName(v string) *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet {
	s.RamRoleName = &v
	return s
}

func (s *DescribeInstanceRamRoleResponseBodyInstanceRamRoleSetsInstanceRamRoleSet) Validate() error {
	return dara.Validate(s)
}
