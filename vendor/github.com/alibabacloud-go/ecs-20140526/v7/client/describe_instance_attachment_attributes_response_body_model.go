// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInstanceAttachmentAttributesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstances(v *DescribeInstanceAttachmentAttributesResponseBodyInstances) *DescribeInstanceAttachmentAttributesResponseBody
	GetInstances() *DescribeInstanceAttachmentAttributesResponseBodyInstances
	SetPageNumber(v int32) *DescribeInstanceAttachmentAttributesResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeInstanceAttachmentAttributesResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeInstanceAttachmentAttributesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeInstanceAttachmentAttributesResponseBody
	GetTotalCount() *int32
}

type DescribeInstanceAttachmentAttributesResponseBody struct {
	// The information about the association between private pools and instances.
	Instances *DescribeInstanceAttachmentAttributesResponseBodyInstances `json:"Instances,omitempty" xml:"Instances,omitempty" type:"Struct"`
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
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeInstanceAttachmentAttributesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttachmentAttributesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) GetInstances() *DescribeInstanceAttachmentAttributesResponseBodyInstances {
	return s.Instances
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) SetInstances(v *DescribeInstanceAttachmentAttributesResponseBodyInstances) *DescribeInstanceAttachmentAttributesResponseBody {
	s.Instances = v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) SetPageNumber(v int32) *DescribeInstanceAttachmentAttributesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) SetPageSize(v int32) *DescribeInstanceAttachmentAttributesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) SetRequestId(v string) *DescribeInstanceAttachmentAttributesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) SetTotalCount(v int32) *DescribeInstanceAttachmentAttributesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponseBody) Validate() error {
	if s.Instances != nil {
		if err := s.Instances.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeInstanceAttachmentAttributesResponseBodyInstances struct {
	Instance []*DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance `json:"Instance,omitempty" xml:"Instance,omitempty" type:"Repeated"`
}

func (s DescribeInstanceAttachmentAttributesResponseBodyInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttachmentAttributesResponseBodyInstances) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttachmentAttributesResponseBodyInstances) GetInstance() []*DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance {
	return s.Instance
}

func (s *DescribeInstanceAttachmentAttributesResponseBodyInstances) SetInstance(v []*DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance) *DescribeInstanceAttachmentAttributesResponseBodyInstances {
	s.Instance = v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponseBodyInstances) Validate() error {
	if s.Instance != nil {
		for _, item := range s.Instance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance struct {
	// The ID of the instance.
	//
	// example:
	//
	// i-bp67acfmxazb4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the private pool. If the value of `PrivatePoolOptionsMatchCriteria` is `Open`, the value of PrivatePoolOptionsId is the ID of the private pool that is automatically matched to the instance.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	PrivatePoolOptionsId *string `json:"PrivatePoolOptionsId,omitempty" xml:"PrivatePoolOptionsId,omitempty"`
	// The match mode of the private pool. Valid values:
	//
	// 	- Open: open private pool. Instances automatically match an open private pool.
	//
	// 	- Target: specified private pool. Instances match a specified private pool.
	//
	// 	- None: no private pool. Instances do not use private pools.
	//
	// example:
	//
	// Open
	PrivatePoolOptionsMatchCriteria *string `json:"PrivatePoolOptionsMatchCriteria,omitempty" xml:"PrivatePoolOptionsMatchCriteria,omitempty"`
}

func (s DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance) GoString() string {
	return s.String()
}

func (s *DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance) GetPrivatePoolOptionsId() *string {
	return s.PrivatePoolOptionsId
}

func (s *DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance) GetPrivatePoolOptionsMatchCriteria() *string {
	return s.PrivatePoolOptionsMatchCriteria
}

func (s *DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance) SetInstanceId(v string) *DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance) SetPrivatePoolOptionsId(v string) *DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance {
	s.PrivatePoolOptionsId = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance) SetPrivatePoolOptionsMatchCriteria(v string) *DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance {
	s.PrivatePoolOptionsMatchCriteria = &v
	return s
}

func (s *DescribeInstanceAttachmentAttributesResponseBodyInstancesInstance) Validate() error {
	return dara.Validate(s)
}
