// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeImageSupportInstanceTypesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionType(v string) *DescribeImageSupportInstanceTypesRequest
	GetActionType() *string
	SetFilter(v []*DescribeImageSupportInstanceTypesRequestFilter) *DescribeImageSupportInstanceTypesRequest
	GetFilter() []*DescribeImageSupportInstanceTypesRequestFilter
	SetImageId(v string) *DescribeImageSupportInstanceTypesRequest
	GetImageId() *string
	SetOwnerId(v int64) *DescribeImageSupportInstanceTypesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeImageSupportInstanceTypesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeImageSupportInstanceTypesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeImageSupportInstanceTypesRequest
	GetResourceOwnerId() *int64
}

type DescribeImageSupportInstanceTypesRequest struct {
	// The scenario in which the image is used. Valid values:
	//
	// 	- CreateEcs (default): instance creation
	//
	// 	- ChangeOS: replacement of the system disk or operating system
	//
	// example:
	//
	// CreateEcs
	ActionType *string `json:"ActionType,omitempty" xml:"ActionType,omitempty"`
	// The number of vCPUs of the instance type.
	Filter []*DescribeImageSupportInstanceTypesRequestFilter `json:"Filter,omitempty" xml:"Filter,omitempty" type:"Repeated"`
	// The region ID of the image. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// example:
	//
	// m-o6w3gy99qf89rkga****
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	OwnerId *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Details about the instance types that are supported by the image.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeImageSupportInstanceTypesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSupportInstanceTypesRequest) GoString() string {
	return s.String()
}

func (s *DescribeImageSupportInstanceTypesRequest) GetActionType() *string {
	return s.ActionType
}

func (s *DescribeImageSupportInstanceTypesRequest) GetFilter() []*DescribeImageSupportInstanceTypesRequestFilter {
	return s.Filter
}

func (s *DescribeImageSupportInstanceTypesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeImageSupportInstanceTypesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeImageSupportInstanceTypesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeImageSupportInstanceTypesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeImageSupportInstanceTypesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeImageSupportInstanceTypesRequest) SetActionType(v string) *DescribeImageSupportInstanceTypesRequest {
	s.ActionType = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesRequest) SetFilter(v []*DescribeImageSupportInstanceTypesRequestFilter) *DescribeImageSupportInstanceTypesRequest {
	s.Filter = v
	return s
}

func (s *DescribeImageSupportInstanceTypesRequest) SetImageId(v string) *DescribeImageSupportInstanceTypesRequest {
	s.ImageId = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesRequest) SetOwnerId(v int64) *DescribeImageSupportInstanceTypesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesRequest) SetRegionId(v string) *DescribeImageSupportInstanceTypesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesRequest) SetResourceOwnerAccount(v string) *DescribeImageSupportInstanceTypesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesRequest) SetResourceOwnerId(v int64) *DescribeImageSupportInstanceTypesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesRequest) Validate() error {
	if s.Filter != nil {
		for _, item := range s.Filter {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeImageSupportInstanceTypesRequestFilter struct {
	// Filter N used to filter instance types.
	//
	// example:
	//
	// imageId
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the image.
	//
	// example:
	//
	// m-o6w3gy99qf89rkga****
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeImageSupportInstanceTypesRequestFilter) String() string {
	return dara.Prettify(s)
}

func (s DescribeImageSupportInstanceTypesRequestFilter) GoString() string {
	return s.String()
}

func (s *DescribeImageSupportInstanceTypesRequestFilter) GetKey() *string {
	return s.Key
}

func (s *DescribeImageSupportInstanceTypesRequestFilter) GetValue() *string {
	return s.Value
}

func (s *DescribeImageSupportInstanceTypesRequestFilter) SetKey(v string) *DescribeImageSupportInstanceTypesRequestFilter {
	s.Key = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesRequestFilter) SetValue(v string) *DescribeImageSupportInstanceTypesRequestFilter {
	s.Value = &v
	return s
}

func (s *DescribeImageSupportInstanceTypesRequestFilter) Validate() error {
	return dara.Validate(s)
}
