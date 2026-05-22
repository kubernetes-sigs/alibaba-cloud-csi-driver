// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVscsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeVscsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeVscsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeVscsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeVscsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeVscsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeVscsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeVscsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeVscsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeVscsRequest
	GetResourceOwnerId() *int64
	SetStatus(v string) *DescribeVscsRequest
	GetStatus() *string
	SetTag(v []*DescribeVscsRequestTag) *DescribeVscsRequest
	GetTag() []*DescribeVscsRequestTag
	SetVscIds(v []*string) *DescribeVscsRequest
	GetVscIds() []*string
}

type DescribeVscsRequest struct {
	// example:
	//
	// i-bp67acfmxazb4ph****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// AAAAAdDWBF2
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// example:
	//
	// Successful
	Status *string                   `json:"Status,omitempty" xml:"Status,omitempty"`
	Tag    []*DescribeVscsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	VscIds []*string                 `json:"VscIds,omitempty" xml:"VscIds,omitempty" type:"Repeated"`
}

func (s DescribeVscsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscsRequest) GoString() string {
	return s.String()
}

func (s *DescribeVscsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeVscsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeVscsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeVscsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeVscsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeVscsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeVscsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeVscsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeVscsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeVscsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeVscsRequest) GetTag() []*DescribeVscsRequestTag {
	return s.Tag
}

func (s *DescribeVscsRequest) GetVscIds() []*string {
	return s.VscIds
}

func (s *DescribeVscsRequest) SetInstanceId(v string) *DescribeVscsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeVscsRequest) SetMaxResults(v int32) *DescribeVscsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeVscsRequest) SetNextToken(v string) *DescribeVscsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeVscsRequest) SetOwnerAccount(v string) *DescribeVscsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeVscsRequest) SetOwnerId(v int64) *DescribeVscsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeVscsRequest) SetRegionId(v string) *DescribeVscsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeVscsRequest) SetResourceGroupId(v string) *DescribeVscsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeVscsRequest) SetResourceOwnerAccount(v string) *DescribeVscsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeVscsRequest) SetResourceOwnerId(v int64) *DescribeVscsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeVscsRequest) SetStatus(v string) *DescribeVscsRequest {
	s.Status = &v
	return s
}

func (s *DescribeVscsRequest) SetTag(v []*DescribeVscsRequestTag) *DescribeVscsRequest {
	s.Tag = v
	return s
}

func (s *DescribeVscsRequest) SetVscIds(v []*string) *DescribeVscsRequest {
	s.VscIds = v
	return s
}

func (s *DescribeVscsRequest) Validate() error {
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

type DescribeVscsRequestTag struct {
	// example:
	//
	// Environment
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeVscsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeVscsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeVscsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeVscsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeVscsRequestTag) SetKey(v string) *DescribeVscsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeVscsRequestTag) SetValue(v string) *DescribeVscsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeVscsRequestTag) Validate() error {
	return dara.Validate(s)
}
