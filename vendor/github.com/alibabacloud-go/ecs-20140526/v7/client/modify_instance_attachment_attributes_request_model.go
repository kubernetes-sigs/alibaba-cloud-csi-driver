// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceAttachmentAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPrivatePoolOptions(v *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) *ModifyInstanceAttachmentAttributesRequest
	GetPrivatePoolOptions() *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions
	SetInstanceId(v string) *ModifyInstanceAttachmentAttributesRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyInstanceAttachmentAttributesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceAttachmentAttributesRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceAttachmentAttributesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceAttachmentAttributesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceAttachmentAttributesRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceAttachmentAttributesRequest struct {
	PrivatePoolOptions *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	// The ID of the instance for which you want to modify the attributes of the private pool.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the private pool. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s ModifyInstanceAttachmentAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttachmentAttributesRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetPrivatePoolOptions() *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceAttachmentAttributesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetPrivatePoolOptions(v *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) *ModifyInstanceAttachmentAttributesRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetInstanceId(v string) *ModifyInstanceAttachmentAttributesRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetOwnerAccount(v string) *ModifyInstanceAttachmentAttributesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetOwnerId(v int64) *ModifyInstanceAttachmentAttributesRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetRegionId(v string) *ModifyInstanceAttachmentAttributesRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetResourceOwnerAccount(v string) *ModifyInstanceAttachmentAttributesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) SetResourceOwnerId(v int64) *ModifyInstanceAttachmentAttributesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequest) Validate() error {
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions struct {
	// The ID of the private pool. Set the value to the ID of the elasticity assurance or capacity reservation that generates the private pool.
	//
	// 	- This parameter is required when `PrivatePoolOptions.MatchCriteria` is set to `Target`.
	//
	// 	- This parameter must be empty when `PrivatePoolOptions.MatchCriteria` is set to `Open` or `None`.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The new type of private pool. Valid values:
	//
	// 	- Open: open private pool. The system matches the instance with an open private pool. If no matching open private pools exist, the system uses resources in the public pool to start the instance.
	//
	// 	- Target: specified private pool. The system uses the capacity in a specified private pool to start the instance. If the specified private pool is unavailable, the instance cannot be started. You must use `PrivatePoolOptions.Id` to specify the ID of a private pool.
	//
	// 	- None: no private pool. The capacity in private pools is not used to start the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) SetId(v string) *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) SetMatchCriteria(v string) *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *ModifyInstanceAttachmentAttributesRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}
