// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeActivationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActivationId(v string) *DescribeActivationsRequest
	GetActivationId() *string
	SetInstanceName(v string) *DescribeActivationsRequest
	GetInstanceName() *string
	SetMaxResults(v int32) *DescribeActivationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeActivationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeActivationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeActivationsRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeActivationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeActivationsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeActivationsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeActivationsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeActivationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeActivationsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeActivationsRequestTag) *DescribeActivationsRequest
	GetTag() []*DescribeActivationsRequestTag
}

type DescribeActivationsRequest struct {
	// The ID of the activation code.
	//
	// example:
	//
	// 4ECEEE12-56F1-4FBC-9AB1-890F1234****
	ActivationId *string `json:"ActivationId,omitempty" xml:"ActivationId,omitempty"`
	// The default instance name prefix.
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The maximum number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number.
	//
	// Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID. Supported regions: China (Qingdao), China (Beijing), China (Zhangjiakou), China (Hohhot), China (Ulanqab), China (Hangzhou), China (Shanghai), China (Shenzhen), China (Heyuan), China (Guangzhou), China (Chengdu), China (Hong Kong), Singapore, Japan (Tokyo), US (Silicon Valley), and US (Virginia).
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the activation code belongs.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the activation code.
	Tag []*DescribeActivationsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeActivationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeActivationsRequest) GetActivationId() *string {
	return s.ActivationId
}

func (s *DescribeActivationsRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *DescribeActivationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeActivationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeActivationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeActivationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeActivationsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeActivationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeActivationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeActivationsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeActivationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeActivationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeActivationsRequest) GetTag() []*DescribeActivationsRequestTag {
	return s.Tag
}

func (s *DescribeActivationsRequest) SetActivationId(v string) *DescribeActivationsRequest {
	s.ActivationId = &v
	return s
}

func (s *DescribeActivationsRequest) SetInstanceName(v string) *DescribeActivationsRequest {
	s.InstanceName = &v
	return s
}

func (s *DescribeActivationsRequest) SetMaxResults(v int32) *DescribeActivationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeActivationsRequest) SetNextToken(v string) *DescribeActivationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeActivationsRequest) SetOwnerAccount(v string) *DescribeActivationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeActivationsRequest) SetOwnerId(v int64) *DescribeActivationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeActivationsRequest) SetPageNumber(v int64) *DescribeActivationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeActivationsRequest) SetPageSize(v int64) *DescribeActivationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeActivationsRequest) SetRegionId(v string) *DescribeActivationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeActivationsRequest) SetResourceGroupId(v string) *DescribeActivationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeActivationsRequest) SetResourceOwnerAccount(v string) *DescribeActivationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeActivationsRequest) SetResourceOwnerId(v int64) *DescribeActivationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeActivationsRequest) SetTag(v []*DescribeActivationsRequestTag) *DescribeActivationsRequest {
	s.Tag = v
	return s
}

func (s *DescribeActivationsRequest) Validate() error {
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

type DescribeActivationsRequestTag struct {
	// The key of tag N of the activation code. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that have this tag can be returned. If multiple tags are specified to query resources, up to 1,000 resources that have all these tags can be returned. To query more than 1,000 resources that have specified tags, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the activation code. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeActivationsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeActivationsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeActivationsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeActivationsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeActivationsRequestTag) SetKey(v string) *DescribeActivationsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeActivationsRequestTag) SetValue(v string) *DescribeActivationsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeActivationsRequestTag) Validate() error {
	return dara.Validate(s)
}
