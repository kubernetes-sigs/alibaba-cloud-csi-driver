// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateImageComponentRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *CreateImageComponentRequest
	GetClientToken() *string
	SetComponentType(v string) *CreateImageComponentRequest
	GetComponentType() *string
	SetComponentVersion(v string) *CreateImageComponentRequest
	GetComponentVersion() *string
	SetContent(v string) *CreateImageComponentRequest
	GetContent() *string
	SetDescription(v string) *CreateImageComponentRequest
	GetDescription() *string
	SetName(v string) *CreateImageComponentRequest
	GetName() *string
	SetOwnerAccount(v string) *CreateImageComponentRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateImageComponentRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateImageComponentRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateImageComponentRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateImageComponentRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateImageComponentRequest
	GetResourceOwnerId() *int64
	SetSystemType(v string) *CreateImageComponentRequest
	GetSystemType() *string
	SetTag(v []*CreateImageComponentRequestTag) *CreateImageComponentRequest
	GetTag() []*CreateImageComponentRequestTag
}

type CreateImageComponentRequest struct {
	// The client token that is used to ensure the idempotency of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The type of the image component. Only image building components and image test components are supported.
	//
	// Valid values:
	//
	// 	- Build
	//
	// 	- Test
	//
	// Default value: Build.
	//
	// >  Image building components can be used only in image building templates. Image test components can be used only in image test templates.
	//
	// example:
	//
	// Build
	ComponentType *string `json:"ComponentType,omitempty" xml:"ComponentType,omitempty"`
	// The version number of the image component, which is used together with the name of the image component. The version number is in the \\<major>.\\<minor>.\\<patch> format. Set \\<major>, \\<minor>, and \\<patch> to non-negative integers.
	//
	// Default value: (x + 1).0.0, in which x is the maximum major version number of the image component.
	//
	// example:
	//
	// null
	ComponentVersion *string `json:"ComponentVersion,omitempty" xml:"ComponentVersion,omitempty"`
	// The content of the image component. The image component consists of multiple commands. The command content cannot exceed 16 KB in size. For information about the commands supported by Image Builder and the formats of the commands, see [Commands supported by Image Builder](https://help.aliyun.com/document_detail/200206.html).
	//
	// example:
	//
	// RUN yum update -y
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The description. The description must be 2 to 256 characters in length and cannot start with [http:// or https://](http://https://。).
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the image component. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with http:// or https://. The name can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// >  If you do not specify `Name`, the return value of `ImageComponentId` is used.
	//
	// example:
	//
	// testComponent
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The type of the operating system supported by the image component.
	//
	// Valid values:
	//
	// 	- Linux
	//
	// 	- Windows
	//
	// Default value: Linux.
	//
	// example:
	//
	// Linux
	SystemType *string `json:"SystemType,omitempty" xml:"SystemType,omitempty"`
	// The tags.
	Tag []*CreateImageComponentRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s CreateImageComponentRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateImageComponentRequest) GoString() string {
	return s.String()
}

func (s *CreateImageComponentRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateImageComponentRequest) GetComponentType() *string {
	return s.ComponentType
}

func (s *CreateImageComponentRequest) GetComponentVersion() *string {
	return s.ComponentVersion
}

func (s *CreateImageComponentRequest) GetContent() *string {
	return s.Content
}

func (s *CreateImageComponentRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateImageComponentRequest) GetName() *string {
	return s.Name
}

func (s *CreateImageComponentRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateImageComponentRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateImageComponentRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateImageComponentRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateImageComponentRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateImageComponentRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateImageComponentRequest) GetSystemType() *string {
	return s.SystemType
}

func (s *CreateImageComponentRequest) GetTag() []*CreateImageComponentRequestTag {
	return s.Tag
}

func (s *CreateImageComponentRequest) SetClientToken(v string) *CreateImageComponentRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateImageComponentRequest) SetComponentType(v string) *CreateImageComponentRequest {
	s.ComponentType = &v
	return s
}

func (s *CreateImageComponentRequest) SetComponentVersion(v string) *CreateImageComponentRequest {
	s.ComponentVersion = &v
	return s
}

func (s *CreateImageComponentRequest) SetContent(v string) *CreateImageComponentRequest {
	s.Content = &v
	return s
}

func (s *CreateImageComponentRequest) SetDescription(v string) *CreateImageComponentRequest {
	s.Description = &v
	return s
}

func (s *CreateImageComponentRequest) SetName(v string) *CreateImageComponentRequest {
	s.Name = &v
	return s
}

func (s *CreateImageComponentRequest) SetOwnerAccount(v string) *CreateImageComponentRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateImageComponentRequest) SetOwnerId(v int64) *CreateImageComponentRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateImageComponentRequest) SetRegionId(v string) *CreateImageComponentRequest {
	s.RegionId = &v
	return s
}

func (s *CreateImageComponentRequest) SetResourceGroupId(v string) *CreateImageComponentRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateImageComponentRequest) SetResourceOwnerAccount(v string) *CreateImageComponentRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateImageComponentRequest) SetResourceOwnerId(v int64) *CreateImageComponentRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateImageComponentRequest) SetSystemType(v string) *CreateImageComponentRequest {
	s.SystemType = &v
	return s
}

func (s *CreateImageComponentRequest) SetTag(v []*CreateImageComponentRequestTag) *CreateImageComponentRequest {
	s.Tag = v
	return s
}

func (s *CreateImageComponentRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateImageComponentRequestTag struct {
	// The key of tag N. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain [http:// or https://](http://https://。). The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain [http:// or https://](http://https://。). The tag value cannot start with acs:.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateImageComponentRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateImageComponentRequestTag) GoString() string {
	return s.String()
}

func (s *CreateImageComponentRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateImageComponentRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateImageComponentRequestTag) SetKey(v string) *CreateImageComponentRequestTag {
	s.Key = &v
	return s
}

func (s *CreateImageComponentRequestTag) SetValue(v string) *CreateImageComponentRequestTag {
	s.Value = &v
	return s
}

func (s *CreateImageComponentRequestTag) Validate() error {
	return dara.Validate(s)
}
