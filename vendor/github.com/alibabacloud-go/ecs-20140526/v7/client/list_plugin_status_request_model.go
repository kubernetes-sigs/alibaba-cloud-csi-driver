// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v []*string) *ListPluginStatusRequest
	GetInstanceId() []*string
	SetMaxResults(v int32) *ListPluginStatusRequest
	GetMaxResults() *int32
	SetName(v string) *ListPluginStatusRequest
	GetName() *string
	SetNextToken(v string) *ListPluginStatusRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *ListPluginStatusRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ListPluginStatusRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *ListPluginStatusRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *ListPluginStatusRequest
	GetPageSize() *int64
	SetRegionId(v string) *ListPluginStatusRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ListPluginStatusRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ListPluginStatusRequest
	GetResourceOwnerId() *int64
}

type ListPluginStatusRequest struct {
	// The ID of the instance.
	InstanceId []*string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty" type:"Repeated"`
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
	// The name of the Cloud Assistant plug-in. The name supports all character sets and must be 1 to 255 characters in length.
	//
	// 	- If this parameter is not specified, the status of all Cloud Assistant plug-ins that are installed on the specified instances are queried.
	//
	//     **
	//
	//     **Note*	- If this parameter is not specified, only a single instance ID can be specified.
	//
	// 	- If this parameter is specified, the status of the specified Cloud Assistant plug-in is queried.
	//
	// example:
	//
	// testPluginName
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

func (s ListPluginStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ListPluginStatusRequest) GoString() string {
	return s.String()
}

func (s *ListPluginStatusRequest) GetInstanceId() []*string {
	return s.InstanceId
}

func (s *ListPluginStatusRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListPluginStatusRequest) GetName() *string {
	return s.Name
}

func (s *ListPluginStatusRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListPluginStatusRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ListPluginStatusRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ListPluginStatusRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListPluginStatusRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListPluginStatusRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ListPluginStatusRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ListPluginStatusRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ListPluginStatusRequest) SetInstanceId(v []*string) *ListPluginStatusRequest {
	s.InstanceId = v
	return s
}

func (s *ListPluginStatusRequest) SetMaxResults(v int32) *ListPluginStatusRequest {
	s.MaxResults = &v
	return s
}

func (s *ListPluginStatusRequest) SetName(v string) *ListPluginStatusRequest {
	s.Name = &v
	return s
}

func (s *ListPluginStatusRequest) SetNextToken(v string) *ListPluginStatusRequest {
	s.NextToken = &v
	return s
}

func (s *ListPluginStatusRequest) SetOwnerAccount(v string) *ListPluginStatusRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ListPluginStatusRequest) SetOwnerId(v int64) *ListPluginStatusRequest {
	s.OwnerId = &v
	return s
}

func (s *ListPluginStatusRequest) SetPageNumber(v int64) *ListPluginStatusRequest {
	s.PageNumber = &v
	return s
}

func (s *ListPluginStatusRequest) SetPageSize(v int64) *ListPluginStatusRequest {
	s.PageSize = &v
	return s
}

func (s *ListPluginStatusRequest) SetRegionId(v string) *ListPluginStatusRequest {
	s.RegionId = &v
	return s
}

func (s *ListPluginStatusRequest) SetResourceOwnerAccount(v string) *ListPluginStatusRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ListPluginStatusRequest) SetResourceOwnerId(v int64) *ListPluginStatusRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ListPluginStatusRequest) Validate() error {
	return dara.Validate(s)
}
