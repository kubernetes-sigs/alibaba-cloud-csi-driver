// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelImagePipelineExecutionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetExecutionId(v string) *CancelImagePipelineExecutionRequest
	GetExecutionId() *string
	SetOwnerAccount(v string) *CancelImagePipelineExecutionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CancelImagePipelineExecutionRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CancelImagePipelineExecutionRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *CancelImagePipelineExecutionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CancelImagePipelineExecutionRequest
	GetResourceOwnerId() *int64
	SetTemplateTag(v []*CancelImagePipelineExecutionRequestTemplateTag) *CancelImagePipelineExecutionRequest
	GetTemplateTag() []*CancelImagePipelineExecutionRequestTemplateTag
}

type CancelImagePipelineExecutionRequest struct {
	// The ID of the image building task.
	//
	// This parameter is required.
	//
	// example:
	//
	// exec-5fb8facb8ed7427c****
	ExecutionId  *string `json:"ExecutionId,omitempty" xml:"ExecutionId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// >  This parameter is not publicly available.
	TemplateTag []*CancelImagePipelineExecutionRequestTemplateTag `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty" type:"Repeated"`
}

func (s CancelImagePipelineExecutionRequest) String() string {
	return dara.Prettify(s)
}

func (s CancelImagePipelineExecutionRequest) GoString() string {
	return s.String()
}

func (s *CancelImagePipelineExecutionRequest) GetExecutionId() *string {
	return s.ExecutionId
}

func (s *CancelImagePipelineExecutionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CancelImagePipelineExecutionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CancelImagePipelineExecutionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CancelImagePipelineExecutionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CancelImagePipelineExecutionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CancelImagePipelineExecutionRequest) GetTemplateTag() []*CancelImagePipelineExecutionRequestTemplateTag {
	return s.TemplateTag
}

func (s *CancelImagePipelineExecutionRequest) SetExecutionId(v string) *CancelImagePipelineExecutionRequest {
	s.ExecutionId = &v
	return s
}

func (s *CancelImagePipelineExecutionRequest) SetOwnerAccount(v string) *CancelImagePipelineExecutionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CancelImagePipelineExecutionRequest) SetOwnerId(v int64) *CancelImagePipelineExecutionRequest {
	s.OwnerId = &v
	return s
}

func (s *CancelImagePipelineExecutionRequest) SetRegionId(v string) *CancelImagePipelineExecutionRequest {
	s.RegionId = &v
	return s
}

func (s *CancelImagePipelineExecutionRequest) SetResourceOwnerAccount(v string) *CancelImagePipelineExecutionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CancelImagePipelineExecutionRequest) SetResourceOwnerId(v int64) *CancelImagePipelineExecutionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CancelImagePipelineExecutionRequest) SetTemplateTag(v []*CancelImagePipelineExecutionRequestTemplateTag) *CancelImagePipelineExecutionRequest {
	s.TemplateTag = v
	return s
}

func (s *CancelImagePipelineExecutionRequest) Validate() error {
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

type CancelImagePipelineExecutionRequestTemplateTag struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CancelImagePipelineExecutionRequestTemplateTag) String() string {
	return dara.Prettify(s)
}

func (s CancelImagePipelineExecutionRequestTemplateTag) GoString() string {
	return s.String()
}

func (s *CancelImagePipelineExecutionRequestTemplateTag) GetKey() *string {
	return s.Key
}

func (s *CancelImagePipelineExecutionRequestTemplateTag) GetValue() *string {
	return s.Value
}

func (s *CancelImagePipelineExecutionRequestTemplateTag) SetKey(v string) *CancelImagePipelineExecutionRequestTemplateTag {
	s.Key = &v
	return s
}

func (s *CancelImagePipelineExecutionRequestTemplateTag) SetValue(v string) *CancelImagePipelineExecutionRequestTemplateTag {
	s.Value = &v
	return s
}

func (s *CancelImagePipelineExecutionRequestTemplateTag) Validate() error {
	return dara.Validate(s)
}
