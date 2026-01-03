// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSendFileResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeSendFileResultsRequest
	GetInstanceId() *string
	SetInvocationStatus(v string) *DescribeSendFileResultsRequest
	GetInvocationStatus() *string
	SetInvokeId(v string) *DescribeSendFileResultsRequest
	GetInvokeId() *string
	SetMaxResults(v int32) *DescribeSendFileResultsRequest
	GetMaxResults() *int32
	SetName(v string) *DescribeSendFileResultsRequest
	GetName() *string
	SetNextToken(v string) *DescribeSendFileResultsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeSendFileResultsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSendFileResultsRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeSendFileResultsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSendFileResultsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeSendFileResultsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeSendFileResultsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeSendFileResultsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSendFileResultsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeSendFileResultsRequestTag) *DescribeSendFileResultsRequest
	GetTag() []*DescribeSendFileResultsRequestTag
}

type DescribeSendFileResultsRequest struct {
	// The ID of the instance for which you want to query file sending records.
	//
	// example:
	//
	// i-hz0jdfwd9f****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The overall sending status of the file. The overall sending status of the file varies based on the sending status of the file on all destination instances. Valid values:
	//
	// 	- Pending: The file is being verified or sent. If the sending state of the file on at least one instance is Pending, the overall sending state of the file is Pending.
	//
	// 	- Running: The file is being sent to the instances. If the sending state of the file on at least one instance is Running, the overall sending state of the file is Running.
	//
	// 	- Success: The file is sent. If the sending state of the file on all instances is Success, the overall sending state of the file is Success.
	//
	// 	- Failed: The file fails to be sent. If the sending state of the file on all instances is Failed, the overall sending state of the file is Failed.
	//
	// 	- PartialFailed: The file sending task succeeds on some instances and fails on other instances. If the sending state of the file is Success on some instances and is Failed on other instances, the overall sending state of the file is PartialFailed.
	//
	// example:
	//
	// Success
	InvocationStatus *string `json:"InvocationStatus,omitempty" xml:"InvocationStatus,omitempty"`
	// The ID of the file sending task.
	//
	// example:
	//
	// f-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
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
	// The name of the file whose sending records you want to query.
	//
	// example:
	//
	// test.txt
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the ECS instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group. After you set this parameter, file sending results in the specified resource group are queried.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags of the file sending task.
	Tag []*DescribeSendFileResultsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeSendFileResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSendFileResultsRequest) GetInvocationStatus() *string {
	return s.InvocationStatus
}

func (s *DescribeSendFileResultsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeSendFileResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSendFileResultsRequest) GetName() *string {
	return s.Name
}

func (s *DescribeSendFileResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSendFileResultsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSendFileResultsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSendFileResultsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSendFileResultsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSendFileResultsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSendFileResultsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeSendFileResultsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSendFileResultsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSendFileResultsRequest) GetTag() []*DescribeSendFileResultsRequestTag {
	return s.Tag
}

func (s *DescribeSendFileResultsRequest) SetInstanceId(v string) *DescribeSendFileResultsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetInvocationStatus(v string) *DescribeSendFileResultsRequest {
	s.InvocationStatus = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetInvokeId(v string) *DescribeSendFileResultsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetMaxResults(v int32) *DescribeSendFileResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetName(v string) *DescribeSendFileResultsRequest {
	s.Name = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetNextToken(v string) *DescribeSendFileResultsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetOwnerAccount(v string) *DescribeSendFileResultsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetOwnerId(v int64) *DescribeSendFileResultsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetPageNumber(v int64) *DescribeSendFileResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetPageSize(v int64) *DescribeSendFileResultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetRegionId(v string) *DescribeSendFileResultsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetResourceGroupId(v string) *DescribeSendFileResultsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetResourceOwnerAccount(v string) *DescribeSendFileResultsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetResourceOwnerId(v int64) *DescribeSendFileResultsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSendFileResultsRequest) SetTag(v []*DescribeSendFileResultsRequestTag) *DescribeSendFileResultsRequest {
	s.Tag = v
	return s
}

func (s *DescribeSendFileResultsRequest) Validate() error {
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

type DescribeSendFileResultsRequestTag struct {
	// The key of tag N of the file sending task. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that have this tag added can be displayed in the response. If multiple tags are specified to query resources, up to 1,000 resources that have all the tags added can be displayed in the response. To query more than 1,000 resources that have specified tags, call the [ListTagResources](https://help.aliyun.com/document_detail/110425.html) operation.
	//
	// The tag key can be up to 64 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the file sending task. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeSendFileResultsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeSendFileResultsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeSendFileResultsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeSendFileResultsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeSendFileResultsRequestTag) SetKey(v string) *DescribeSendFileResultsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeSendFileResultsRequestTag) SetValue(v string) *DescribeSendFileResultsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeSendFileResultsRequestTag) Validate() error {
	return dara.Validate(s)
}
