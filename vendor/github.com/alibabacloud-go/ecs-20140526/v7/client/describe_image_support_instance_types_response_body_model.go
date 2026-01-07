// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSupportInstanceTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageId(v string) *DescribeImageSupportInstanceTypesResponseBody
	GetImageId() *string
	SetInstanceTypes(v *DescribeImageSupportInstanceTypesResponseBodyInstanceTypes) *DescribeImageSupportInstanceTypesResponseBody
	GetInstanceTypes() *DescribeImageSupportInstanceTypesResponseBodyInstanceTypes
	SetRegionId(v string) *DescribeImageSupportInstanceTypesResponseBody
	GetRegionId() *string
	SetRequestId(v string) *DescribeImageSupportInstanceTypesResponseBody
	GetRequestId() *string
}

type DescribeImageSupportInstanceTypesResponseBody struct {
	// The key of filter N. Only the image ID can be used to filter instance types. Valid values:
	//
	// 	- imagId: image ID
	//
	// 	- filter: image ID
	//
	// example:
	//
	// m-o6w3gy99qf89rkga****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// {
	//
	//     "RequestId": "CF661E2D-4AFE-4BCD-959A-A65E14416B44",
	//
	//     "RegionId": "cn-hangzhou",
	//
	//     "ImageId": "ubuntu_16_0402_64_20G_alibase_20180409.vhd",
	//
	//     "InstanceTypes": {
	//
	//         "InstanceType": [{
	//
	//             "InstanceTypeId": "ecs.t1.xsmall",
	//
	//             "CpuCoreCount": 1,
	//
	//             "MemorySize": 0.5,
	//
	//             "InstanceTypeFamily": "ecs.t1"
	//
	//         },
	//
	//         {
	//
	//             "InstanceTypeId": "ecs.t1.small",
	//
	//             "CpuCoreCount": 1,
	//
	//             "MemorySize": 1,
	//
	//             "InstanceTypeFamily": "ecs.t1"
	//
	//         }]
	//
	//     }
	//
	// }
	InstanceTypes *DescribeImageSupportInstanceTypesResponseBodyInstanceTypes `json:"InstanceTypes,omitempty" xml:"InstanceTypes,omitempty" type:"Struct"`
	// {
	//
	//     "RequestId": "CF661E2D-4AFE-4BCD-959A-A65E14416B44",
	//
	//     "RegionId": "cn-hangzhou",
	//
	//     "ImageId": "ubuntu_16_0402_64_20G_alibase_20180409.vhd",
	//
	//     "InstanceTypes": {
	//
	//         "InstanceType": [{
	//
	//             "InstanceTypeId": "ecs.t1.xsmall",
	//
	//             "CpuCoreCount": 1,
	//
	//             "MemorySize": 0.5,
	//
	//             "InstanceTypeFamily": "ecs.t1"
	//
	//         },
	//
	//         {
	//
	//             "InstanceTypeId": "ecs.t1.small",
	//
	//             "CpuCoreCount": 1,
	//
	//             "MemorySize": 1,
	//
	//             "InstanceTypeFamily": "ecs.t1"
	//
	//         }]
	//
	//     }
	//
	// }
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
}

func (s DescribeImageSupportInstanceTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSupportInstanceTypesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeImageSupportInstanceTypesResponseBody) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageSupportInstanceTypesResponseBody) GetInstanceTypes() *DescribeImageSupportInstanceTypesResponseBodyInstanceTypes {
	return s.InstanceTypes
}

func (s *DescribeImageSupportInstanceTypesResponseBody) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageSupportInstanceTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeImageSupportInstanceTypesResponseBody) SetImageId(v string) *DescribeImageSupportInstanceTypesResponseBody {
	s.ImageId = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponseBody) SetInstanceTypes(v *DescribeImageSupportInstanceTypesResponseBodyInstanceTypes) *DescribeImageSupportInstanceTypesResponseBody {
	s.InstanceTypes = v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponseBody) SetRegionId(v string) *DescribeImageSupportInstanceTypesResponseBody {
	s.RegionId = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponseBody) SetRequestId(v string) *DescribeImageSupportInstanceTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponseBody) Validate() error {
	if s.InstanceTypes != nil {
		if err := s.InstanceTypes.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeImageSupportInstanceTypesResponseBodyInstanceTypes struct {
	InstanceType []*DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType `json:"InstanceType,omitempty" xml:"InstanceType,omitempty" type:"Repeated"`
}

func (s DescribeImageSupportInstanceTypesResponseBodyInstanceTypes) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSupportInstanceTypesResponseBodyInstanceTypes) GoString() string {
	return s.String()
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypes) GetInstanceType() []*DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType {
	return s.InstanceType
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypes) SetInstanceType(v []*DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) *DescribeImageSupportInstanceTypesResponseBodyInstanceTypes {
	s.InstanceType = v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypes) Validate() error {
	if s.InstanceType != nil {
		for _, item := range s.InstanceType {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType struct {
	// The number of vCPUs of the instance type.
	//
	// example:
	//
	// 1
	CpuCoreCount *int32 `json:"CpuCoreCount,omitempty" xml:"CpuCoreCount,omitempty"`
	// DescribeImageSupportInstanceTypes
	//
	// example:
	//
	// ecs.t1
	InstanceTypeFamily *string `json:"InstanceTypeFamily,omitempty" xml:"InstanceTypeFamily,omitempty"`
	// Queries the instance types supported by an image.
	//
	// example:
	//
	// ecs.t1.xsmall
	InstanceTypeId *string `json:"InstanceTypeId,omitempty" xml:"InstanceTypeId,omitempty"`
	// The memory size of the instance type. Unit: GiB.
	//
	// example:
	//
	// 1024
	MemorySize *float32 `json:"MemorySize,omitempty" xml:"MemorySize,omitempty"`
}

func (s DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) GoString() string {
	return s.String()
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) GetCpuCoreCount() *int32 {
	return s.CpuCoreCount
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceTypeFamily() *string {
	return s.InstanceTypeFamily
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) GetInstanceTypeId() *string {
	return s.InstanceTypeId
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) GetMemorySize() *float32 {
	return s.MemorySize
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) SetCpuCoreCount(v int32) *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.CpuCoreCount = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeFamily(v string) *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeFamily = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) SetInstanceTypeId(v string) *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.InstanceTypeId = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) SetMemorySize(v float32) *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType {
	s.MemorySize = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesResponseBodyInstanceTypesInstanceType) Validate() error {
	return dara.Validate(s)
}
