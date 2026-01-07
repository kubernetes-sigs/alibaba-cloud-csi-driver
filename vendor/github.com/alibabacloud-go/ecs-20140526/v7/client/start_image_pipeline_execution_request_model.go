// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartImagePipelineExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *StartImagePipelineExecutionRequest
	GetClientToken() *string
	SetImagePipelineId(v string) *StartImagePipelineExecutionRequest
	GetImagePipelineId() *string
	SetOwnerAccount(v string) *StartImagePipelineExecutionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *StartImagePipelineExecutionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *StartImagePipelineExecutionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *StartImagePipelineExecutionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *StartImagePipelineExecutionRequest
	GetResourceOwnerId() *int64
	SetTag(v []*StartImagePipelineExecutionRequestTag) *StartImagePipelineExecutionRequest
	GetTag() []*StartImagePipelineExecutionRequestTag
	SetTemplateTag(v []*StartImagePipelineExecutionRequestTemplateTag) *StartImagePipelineExecutionRequest
	GetTemplateTag() []*StartImagePipelineExecutionRequestTemplateTag
}

type StartImagePipelineExecutionRequest struct {
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but make sure that the token is unique among requests. **The token can contain only ASCII characters and cannot exceed 64 characters in length.*	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the image template.
	//
	// This parameter is required.
	//
	// example:
	//
	// ip-2ze5tsl5bp6nf2b3****
	ImagePipelineId *string `json:"ImagePipelineId,omitempty" xml:"ImagePipelineId,omitempty"`
	OwnerAccount    *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId         *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags.
	Tag []*StartImagePipelineExecutionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Deprecated
	//
	// >  This parameter is deprecated.
	TemplateTag []*StartImagePipelineExecutionRequestTemplateTag `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty" type:"Repeated"`
}

func (s StartImagePipelineExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s StartImagePipelineExecutionRequest) GoString() string {
	return s.String()
}

func (s *StartImagePipelineExecutionRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *StartImagePipelineExecutionRequest) GetImagePipelineId() *string {
	return s.ImagePipelineId
}

func (s *StartImagePipelineExecutionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *StartImagePipelineExecutionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *StartImagePipelineExecutionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *StartImagePipelineExecutionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *StartImagePipelineExecutionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *StartImagePipelineExecutionRequest) GetTag() []*StartImagePipelineExecutionRequestTag {
	return s.Tag
}

func (s *StartImagePipelineExecutionRequest) GetTemplateTag() []*StartImagePipelineExecutionRequestTemplateTag {
	return s.TemplateTag
}

func (s *StartImagePipelineExecutionRequest) SetClientToken(v string) *StartImagePipelineExecutionRequest {
	s.ClientToken = &v
	return s
}

func (s *StartImagePipelineExecutionRequest) SetImagePipelineId(v string) *StartImagePipelineExecutionRequest {
	s.ImagePipelineId = &v
	return s
}

func (s *StartImagePipelineExecutionRequest) SetOwnerAccount(v string) *StartImagePipelineExecutionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *StartImagePipelineExecutionRequest) SetOwnerId(v int64) *StartImagePipelineExecutionRequest {
	s.OwnerId = &v
	return s
}

func (s *StartImagePipelineExecutionRequest) SetRegionId(v string) *StartImagePipelineExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *StartImagePipelineExecutionRequest) SetResourceOwnerAccount(v string) *StartImagePipelineExecutionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *StartImagePipelineExecutionRequest) SetResourceOwnerId(v int64) *StartImagePipelineExecutionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *StartImagePipelineExecutionRequest) SetTag(v []*StartImagePipelineExecutionRequestTag) *StartImagePipelineExecutionRequest {
	s.Tag = v
	return s
}

func (s *StartImagePipelineExecutionRequest) SetTemplateTag(v []*StartImagePipelineExecutionRequestTemplateTag) *StartImagePipelineExecutionRequest {
	s.TemplateTag = v
	return s
}

func (s *StartImagePipelineExecutionRequest) Validate() error {
	if s.Tag != nil {
		for _, item := range s.Tag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type StartImagePipelineExecutionRequestTag struct {
	// The key of tag N. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. The tag key cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length. The tag value cannot start with `acs:` or contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s StartImagePipelineExecutionRequestTag) String() string {
	return dara.Prettify(s)
}

func (s StartImagePipelineExecutionRequestTag) GoString() string {
	return s.String()
}

func (s *StartImagePipelineExecutionRequestTag) GetKey() *string {
	return s.Key
}

func (s *StartImagePipelineExecutionRequestTag) GetValue() *string {
	return s.Value
}

func (s *StartImagePipelineExecutionRequestTag) SetKey(v string) *StartImagePipelineExecutionRequestTag {
	s.Key = &v
	return s
}

func (s *StartImagePipelineExecutionRequestTag) SetValue(v string) *StartImagePipelineExecutionRequestTag {
	s.Value = &v
	return s
}

func (s *StartImagePipelineExecutionRequestTag) Validate() error {
	return dara.Validate(s)
}

type StartImagePipelineExecutionRequestTemplateTag struct {
	// >  This parameter is deprecated.
	//
	// example:
	//
	// null
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// >  This parameter is deprecated.
	//
	// example:
	//
	// null
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s StartImagePipelineExecutionRequestTemplateTag) String() string {
	return dara.Prettify(s)
}

func (s StartImagePipelineExecutionRequestTemplateTag) GoString() string {
	return s.String()
}

func (s *StartImagePipelineExecutionRequestTemplateTag) GetKey() *string {
	return s.Key
}

func (s *StartImagePipelineExecutionRequestTemplateTag) GetValue() *string {
	return s.Value
}

func (s *StartImagePipelineExecutionRequestTemplateTag) SetKey(v string) *StartImagePipelineExecutionRequestTemplateTag {
	s.Key = &v
	return s
}

func (s *StartImagePipelineExecutionRequestTemplateTag) SetValue(v string) *StartImagePipelineExecutionRequestTemplateTag {
	s.Value = &v
	return s
}

func (s *StartImagePipelineExecutionRequestTemplateTag) Validate() error {
	return dara.Validate(s)
}
