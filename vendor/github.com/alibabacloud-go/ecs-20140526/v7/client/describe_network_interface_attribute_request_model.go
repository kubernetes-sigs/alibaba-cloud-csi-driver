// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNetworkInterfaceAttributeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAttribute(v string) *DescribeNetworkInterfaceAttributeRequest
	GetAttribute() *string
	SetNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributeRequest
	GetNetworkInterfaceId() *string
	SetOwnerAccount(v string) *DescribeNetworkInterfaceAttributeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeNetworkInterfaceAttributeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeNetworkInterfaceAttributeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeNetworkInterfaceAttributeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeNetworkInterfaceAttributeRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeNetworkInterfaceAttributeRequestTag) *DescribeNetworkInterfaceAttributeRequest
	GetTag() []*DescribeNetworkInterfaceAttributeRequestTag
}

type DescribeNetworkInterfaceAttributeRequest struct {
	// The attribute of the ENI. Valid values:
	//
	// attachment: member ENI attachment information of the trunk ENI. This value is in invitational preview and is not publicly available.
	//
	// connectionTrackingConfiguration: connection tracking configuration.
	//
	// Default value:
	//
	// example:
	//
	// attachment
	Attribute *string `json:"Attribute,omitempty" xml:"Attribute,omitempty"`
	// The ID of the ENI.
	//
	// This parameter is required.
	//
	// example:
	//
	// eni-bp67acfmxazb4p****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	OwnerAccount       *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId            *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the ENI. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// >  This parameter is unavailable.
	Tag []*DescribeNetworkInterfaceAttributeRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeNetworkInterfaceAttributeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeRequest) GetAttribute() *string {
	return s.Attribute
}

func (s *DescribeNetworkInterfaceAttributeRequest) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *DescribeNetworkInterfaceAttributeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeNetworkInterfaceAttributeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeNetworkInterfaceAttributeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeNetworkInterfaceAttributeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeNetworkInterfaceAttributeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeNetworkInterfaceAttributeRequest) GetTag() []*DescribeNetworkInterfaceAttributeRequestTag {
	return s.Tag
}

func (s *DescribeNetworkInterfaceAttributeRequest) SetAttribute(v string) *DescribeNetworkInterfaceAttributeRequest {
	s.Attribute = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeRequest) SetNetworkInterfaceId(v string) *DescribeNetworkInterfaceAttributeRequest {
	s.NetworkInterfaceId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeRequest) SetOwnerAccount(v string) *DescribeNetworkInterfaceAttributeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeRequest) SetOwnerId(v int64) *DescribeNetworkInterfaceAttributeRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeRequest) SetRegionId(v string) *DescribeNetworkInterfaceAttributeRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeRequest) SetResourceOwnerAccount(v string) *DescribeNetworkInterfaceAttributeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeRequest) SetResourceOwnerId(v int64) *DescribeNetworkInterfaceAttributeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeRequest) SetTag(v []*DescribeNetworkInterfaceAttributeRequestTag) *DescribeNetworkInterfaceAttributeRequest {
	s.Tag = v
	return s
}

func (s *DescribeNetworkInterfaceAttributeRequest) Validate() error {
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

type DescribeNetworkInterfaceAttributeRequestTag struct {
	// >  This parameter is unavailable.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// >  This parameter is unavailable.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeNetworkInterfaceAttributeRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeNetworkInterfaceAttributeRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeNetworkInterfaceAttributeRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeNetworkInterfaceAttributeRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeNetworkInterfaceAttributeRequestTag) SetKey(v string) *DescribeNetworkInterfaceAttributeRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeRequestTag) SetValue(v string) *DescribeNetworkInterfaceAttributeRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeNetworkInterfaceAttributeRequestTag) Validate() error {
	return dara.Validate(s)
}
