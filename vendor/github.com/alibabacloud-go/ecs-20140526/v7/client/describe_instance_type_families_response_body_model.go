// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceTypeFamiliesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceTypeFamilies(v *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) *DescribeInstanceTypeFamiliesResponseBody
	GetInstanceTypeFamilies() *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies
	SetRequestId(v string) *DescribeInstanceTypeFamiliesResponseBody
	GetRequestId() *string
}

type DescribeInstanceTypeFamiliesResponseBody struct {
	// The instance families.
	InstanceTypeFamilies *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies `json:"InstanceTypeFamilies,omitempty" xml:"InstanceTypeFamilies,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeInstanceTypeFamiliesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeFamiliesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeFamiliesResponseBody) GetInstanceTypeFamilies() *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies {
	return s.InstanceTypeFamilies
}

func (s *DescribeInstanceTypeFamiliesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceTypeFamiliesResponseBody) SetInstanceTypeFamilies(v *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) *DescribeInstanceTypeFamiliesResponseBody {
	s.InstanceTypeFamilies = v
	return s
}

func (s *DescribeInstanceTypeFamiliesResponseBody) SetRequestId(v string) *DescribeInstanceTypeFamiliesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceTypeFamiliesResponseBody) Validate() error {
	if s.InstanceTypeFamilies != nil {
		if err := s.InstanceTypeFamilies.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies struct {
	InstanceTypeFamily []*DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty" type:"Repeated"`
}

func (s DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) GetInstanceTypeFamily() []*DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily {
	return s.InstanceTypeFamily
}

func (s *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) SetInstanceTypeFamily(v []*DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies {
	s.InstanceTypeFamily = v
	return s
}

func (s *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamilies) Validate() error {
	if s.InstanceTypeFamily != nil {
		for _, item := range s.InstanceTypeFamily {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily struct {
	// The series of the instance family.
	//
	// example:
	//
	// ecs-5
	Generation *string `json:"Generation,omitempty" xml:"Generation,omitempty"`
	// The ID of the instance family.
	//
	// example:
	//
	// ecs.g6
	InstanceTypeFamilyId *string `json:"InstanceTypeFamilyId,omitempty" xml:"InstanceTypeFamilyId,omitempty"`
}

func (s DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) GoString() string {
	return s.String()
}

func (s *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) GetGeneration() *string {
	return s.Generation
}

func (s *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) GetInstanceTypeFamilyId() *string {
	return s.InstanceTypeFamilyId
}

func (s *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) SetGeneration(v string) *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily {
	s.Generation = &v
	return s
}

func (s *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) SetInstanceTypeFamilyId(v string) *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily {
	s.InstanceTypeFamilyId = &v
	return s
}

func (s *DescribeInstanceTypeFamiliesResponseBodyInstanceTypeFamiliesInstanceTypeFamily) Validate() error {
	return dara.Validate(s)
}
