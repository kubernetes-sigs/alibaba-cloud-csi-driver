// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConvertNatPublicIpToEipRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ConvertNatPublicIpToEipRequest
	GetInstanceId() *string
	SetOwnerId(v int64) *ConvertNatPublicIpToEipRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ConvertNatPublicIpToEipRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ConvertNatPublicIpToEipRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ConvertNatPublicIpToEipRequest
	GetResourceOwnerId() *int64
}

type ConvertNatPublicIpToEipRequest struct {
	// The ID of the instance whose public IP address you want to convert into an EIP.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp171jr36ge2ulvk****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerId    *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
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

func (s ConvertNatPublicIpToEipRequest) String() string {
	return dara.Prettify(s)
}

func (s ConvertNatPublicIpToEipRequest) GoString() string {
	return s.String()
}

func (s *ConvertNatPublicIpToEipRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ConvertNatPublicIpToEipRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ConvertNatPublicIpToEipRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ConvertNatPublicIpToEipRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ConvertNatPublicIpToEipRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ConvertNatPublicIpToEipRequest) SetInstanceId(v string) *ConvertNatPublicIpToEipRequest {
	s.InstanceId = &v
	return s
}

func (s *ConvertNatPublicIpToEipRequest) SetOwnerId(v int64) *ConvertNatPublicIpToEipRequest {
	s.OwnerId = &v
	return s
}

func (s *ConvertNatPublicIpToEipRequest) SetRegionId(v string) *ConvertNatPublicIpToEipRequest {
	s.RegionId = &v
	return s
}

func (s *ConvertNatPublicIpToEipRequest) SetResourceOwnerAccount(v string) *ConvertNatPublicIpToEipRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ConvertNatPublicIpToEipRequest) SetResourceOwnerId(v int64) *ConvertNatPublicIpToEipRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ConvertNatPublicIpToEipRequest) Validate() error {
	return dara.Validate(s)
}
