// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudAssistantStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v []*string) *DescribeCloudAssistantStatusRequest
	GetInstanceId() []*string
	SetMaxResults(v int32) *DescribeCloudAssistantStatusRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeCloudAssistantStatusRequest
	GetNextToken() *string
	SetOSType(v string) *DescribeCloudAssistantStatusRequest
	GetOSType() *string
	SetOwnerAccount(v string) *DescribeCloudAssistantStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeCloudAssistantStatusRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeCloudAssistantStatusRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeCloudAssistantStatusRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeCloudAssistantStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeCloudAssistantStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeCloudAssistantStatusRequest
	GetResourceOwnerId() *int64
}

type DescribeCloudAssistantStatusRequest struct {
	// The instance ID.
	//
	// example:
	//
	// i-bp1iudwa5b1tqa****
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
	// The maximum number of entries per page. If you specify **InstanceId**, this parameter does not take effect.
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
	// AAAAAdDWBF2
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The operating system type of the instance. Valid values:
	//
	// 	- Windows
	//
	// 	- Linux
	//
	// 	- FreeBSD
	//
	// example:
	//
	// Windows
	OSType       *string `json:"OSType,omitempty" xml:"OSType,omitempty"`
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
	// The region ID of the instance. You can call [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) to query the most recent region list.
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

func (s DescribeCloudAssistantStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudAssistantStatusRequest) GoString() string {
	return s.String()
}

func (s *DescribeCloudAssistantStatusRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *DescribeCloudAssistantStatusRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeCloudAssistantStatusRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeCloudAssistantStatusRequest) GetOSType() *string {
	return s.OSType
}

func (s *DescribeCloudAssistantStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeCloudAssistantStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeCloudAssistantStatusRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeCloudAssistantStatusRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeCloudAssistantStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeCloudAssistantStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeCloudAssistantStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeCloudAssistantStatusRequest) SetInstanceId(v []*string) *DescribeCloudAssistantStatusRequest {
	s.InstanceId = v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetMaxResults(v int32) *DescribeCloudAssistantStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetNextToken(v string) *DescribeCloudAssistantStatusRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetOSType(v string) *DescribeCloudAssistantStatusRequest {
	s.OSType = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetOwnerAccount(v string) *DescribeCloudAssistantStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetOwnerId(v int64) *DescribeCloudAssistantStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetPageNumber(v int64) *DescribeCloudAssistantStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetPageSize(v int64) *DescribeCloudAssistantStatusRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetRegionId(v string) *DescribeCloudAssistantStatusRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetResourceOwnerAccount(v string) *DescribeCloudAssistantStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) SetResourceOwnerId(v int64) *DescribeCloudAssistantStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeCloudAssistantStatusRequest) Validate() error {
	return dara.Validate(s)
}
