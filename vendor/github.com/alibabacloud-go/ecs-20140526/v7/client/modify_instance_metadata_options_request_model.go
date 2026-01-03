// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceMetadataOptionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHttpEndpoint(v string) *ModifyInstanceMetadataOptionsRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *ModifyInstanceMetadataOptionsRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *ModifyInstanceMetadataOptionsRequest
	GetHttpTokens() *string
	SetInstanceId(v string) *ModifyInstanceMetadataOptionsRequest
	GetInstanceId() *string
	SetInstanceMetadataTags(v string) *ModifyInstanceMetadataOptionsRequest
	GetInstanceMetadataTags() *string
	SetOwnerId(v int64) *ModifyInstanceMetadataOptionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyInstanceMetadataOptionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyInstanceMetadataOptionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceMetadataOptionsRequest
	GetResourceOwnerId() *int64
}

type ModifyInstanceMetadataOptionsRequest struct {
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// Default value: enabled.
	//
	// >  For information about instance metadata, see [Obtain instance metadata](https://help.aliyun.com/document_detail/49122.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 1
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Specifies whether to forcefully use the security hardening mode (IMDSv2) to access instance metadata. Valid values:
	//
	// 	- optional: does not forcefully use the security hardening mode (IMDSv2).
	//
	// 	- required: forcefully uses the security hardening mode (IMDSv2). After you set this parameter to required, you cannot access instance metadata in normal mode.
	//
	// Default value: optional.
	//
	// >  For more information about modes of accessing instance metadata, see [Obtain instance metadata](https://help.aliyun.com/document_detail/150575.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The ID of the instance.
	//
	// example:
	//
	// i-bp67acfmxaz****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// Default value: disabled.
	//
	// >  The tag key must be a combination of letters, digits, @, colons (:), underscores (_), hyphens (-), periods (.), equal signs (=), and commas (,). The tag key cannot be "." or "..". Otherwise, the tag key cannot be accessed in the metadata.
	//
	// example:
	//
	// null
	InstanceMetadataTags *string `json:"InstanceMetadataTags,omitempty" xml:"InstanceMetadataTags,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s ModifyInstanceMetadataOptionsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceMetadataOptionsRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceMetadataOptionsRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *ModifyInstanceMetadataOptionsRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *ModifyInstanceMetadataOptionsRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *ModifyInstanceMetadataOptionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceMetadataOptionsRequest) GetInstanceMetadataTags() *string {
	return s.InstanceMetadataTags
}

func (s *ModifyInstanceMetadataOptionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceMetadataOptionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyInstanceMetadataOptionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceMetadataOptionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceMetadataOptionsRequest) SetHttpEndpoint(v string) *ModifyInstanceMetadataOptionsRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsRequest) SetHttpPutResponseHopLimit(v int32) *ModifyInstanceMetadataOptionsRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsRequest) SetHttpTokens(v string) *ModifyInstanceMetadataOptionsRequest {
	s.HttpTokens = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsRequest) SetInstanceId(v string) *ModifyInstanceMetadataOptionsRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsRequest) SetInstanceMetadataTags(v string) *ModifyInstanceMetadataOptionsRequest {
	s.InstanceMetadataTags = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsRequest) SetOwnerId(v int64) *ModifyInstanceMetadataOptionsRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsRequest) SetRegionId(v string) *ModifyInstanceMetadataOptionsRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsRequest) SetResourceOwnerAccount(v string) *ModifyInstanceMetadataOptionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsRequest) SetResourceOwnerId(v int64) *ModifyInstanceMetadataOptionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceMetadataOptionsRequest) Validate() error {
	return dara.Validate(s)
}
