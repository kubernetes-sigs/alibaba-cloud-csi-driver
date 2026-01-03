// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLaunchTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLaunchTemplateId(v []*string) *DescribeLaunchTemplatesRequest
	GetLaunchTemplateId() []*string
	SetLaunchTemplateName(v []*string) *DescribeLaunchTemplatesRequest
	GetLaunchTemplateName() []*string
	SetOwnerAccount(v string) *DescribeLaunchTemplatesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeLaunchTemplatesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeLaunchTemplatesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeLaunchTemplatesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeLaunchTemplatesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeLaunchTemplatesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeLaunchTemplatesRequest
	GetResourceOwnerId() *int64
	SetTemplateResourceGroupId(v string) *DescribeLaunchTemplatesRequest
	GetTemplateResourceGroupId() *string
	SetTemplateTag(v []*DescribeLaunchTemplatesRequestTemplateTag) *DescribeLaunchTemplatesRequest
	GetTemplateTag() []*DescribeLaunchTemplatesRequestTemplateTag
}

type DescribeLaunchTemplatesRequest struct {
	// The IDs of launch templates.
	//
	// 	- You can query up to 100 launch templates.
	//
	// 	- You must specify LaunchTemplateId or LaunchTemplateName to specify a launch template.
	//
	// example:
	//
	// lt-m5e3ofjr1zn1aw7q****
	LaunchTemplateId []*string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty" type:"Repeated"`
	// The names of launch templates.
	//
	// 	- You can query up to 100 launch templates.
	//
	// 	- You must specify LaunchTemplateId or LaunchTemplateName to specify a launch template.
	//
	// example:
	//
	// wd-152630748****
	LaunchTemplateName []*string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty" type:"Repeated"`
	OwnerAccount       *string   `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64    `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Page starts from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the launch template. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The ID of the resource group to which the launch template belongs. If you specify this parameter to query resources, up to 1,000 resources that belong to the specified resource group can be returned.
	//
	// >  The default resource group is not supported.
	//
	// example:
	//
	// rg-acfmxazb4p****
	TemplateResourceGroupId *string `json:"TemplateResourceGroupId,omitempty" xml:"TemplateResourceGroupId,omitempty"`
	// The tags of the launch template.
	//
	// >  You can only call API operations to add tags to and query the tags of a launch template. You cannot add tags to or view the tags of a launch template in the ECS console.
	TemplateTag []*DescribeLaunchTemplatesRequestTemplateTag `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty" type:"Repeated"`
}

func (s DescribeLaunchTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplatesRequest) GetLaunchTemplateId() []*string {
	return s.LaunchTemplateId
}

func (s *DescribeLaunchTemplatesRequest) GetLaunchTemplateName() []*string {
	return s.LaunchTemplateName
}

func (s *DescribeLaunchTemplatesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeLaunchTemplatesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeLaunchTemplatesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeLaunchTemplatesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeLaunchTemplatesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeLaunchTemplatesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeLaunchTemplatesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeLaunchTemplatesRequest) GetTemplateResourceGroupId() *string {
	return s.TemplateResourceGroupId
}

func (s *DescribeLaunchTemplatesRequest) GetTemplateTag() []*DescribeLaunchTemplatesRequestTemplateTag {
	return s.TemplateTag
}

func (s *DescribeLaunchTemplatesRequest) SetLaunchTemplateId(v []*string) *DescribeLaunchTemplatesRequest {
	s.LaunchTemplateId = v
	return s
}

func (s *DescribeLaunchTemplatesRequest) SetLaunchTemplateName(v []*string) *DescribeLaunchTemplatesRequest {
	s.LaunchTemplateName = v
	return s
}

func (s *DescribeLaunchTemplatesRequest) SetOwnerAccount(v string) *DescribeLaunchTemplatesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeLaunchTemplatesRequest) SetOwnerId(v int64) *DescribeLaunchTemplatesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeLaunchTemplatesRequest) SetPageNumber(v int32) *DescribeLaunchTemplatesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeLaunchTemplatesRequest) SetPageSize(v int32) *DescribeLaunchTemplatesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeLaunchTemplatesRequest) SetRegionId(v string) *DescribeLaunchTemplatesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeLaunchTemplatesRequest) SetResourceOwnerAccount(v string) *DescribeLaunchTemplatesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeLaunchTemplatesRequest) SetResourceOwnerId(v int64) *DescribeLaunchTemplatesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeLaunchTemplatesRequest) SetTemplateResourceGroupId(v string) *DescribeLaunchTemplatesRequest {
	s.TemplateResourceGroupId = &v
	return s
}

func (s *DescribeLaunchTemplatesRequest) SetTemplateTag(v []*DescribeLaunchTemplatesRequestTemplateTag) *DescribeLaunchTemplatesRequest {
	s.TemplateTag = v
	return s
}

func (s *DescribeLaunchTemplatesRequest) Validate() error {
	if s.TemplateTag != nil {
		for _, item := range s.TemplateTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLaunchTemplatesRequestTemplateTag struct {
	// The key of tag N of the launch template. Valid values of N: 1 to 20.
	//
	// If you specify a single tag to query resources, up to 1,000 resources to which the tag is added are returned. If you specify multiple tags to query resources, up to 1,000 resources to which all specified tags are added are returned. To query more than 1,000 resources that have specified tags added, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the launch template. Valid values of N: 1 to 20.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeLaunchTemplatesRequestTemplateTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeLaunchTemplatesRequestTemplateTag) GoString() string {
	return s.String()
}

func (s *DescribeLaunchTemplatesRequestTemplateTag) GetKey() *string {
	return s.Key
}

func (s *DescribeLaunchTemplatesRequestTemplateTag) GetValue() *string {
	return s.Value
}

func (s *DescribeLaunchTemplatesRequestTemplateTag) SetKey(v string) *DescribeLaunchTemplatesRequestTemplateTag {
	s.Key = &v
	return s
}

func (s *DescribeLaunchTemplatesRequestTemplateTag) SetValue(v string) *DescribeLaunchTemplatesRequestTemplateTag {
	s.Value = &v
	return s
}

func (s *DescribeLaunchTemplatesRequestTemplateTag) Validate() error {
	return dara.Validate(s)
}
