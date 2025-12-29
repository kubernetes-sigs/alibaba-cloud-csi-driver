// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSnapshotLinksRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDiskIds(v string) *DescribeSnapshotLinksRequest
	GetDiskIds() *string
	SetInstanceId(v string) *DescribeSnapshotLinksRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeSnapshotLinksRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeSnapshotLinksRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeSnapshotLinksRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeSnapshotLinksRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeSnapshotLinksRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeSnapshotLinksRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeSnapshotLinksRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeSnapshotLinksRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeSnapshotLinksRequest
	GetResourceOwnerId() *int64
	SetSnapshotLinkIds(v string) *DescribeSnapshotLinksRequest
	GetSnapshotLinkIds() *string
}

type DescribeSnapshotLinksRequest struct {
	// The disk IDs. You can specify a JSON array that contains a maximum of 100 disk IDs. Separate the disk IDs with commas (,).
	//
	// example:
	//
	// ["d-bp1d6tsvznfghy7y****", "d-bp1ippxbaql9zet7****", … "d-bp1ib7bcz07lcxa9****"]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// i-bp1h6jmbefj2cyqs****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries per page. Maximum value: 100. Default value:
	//
	// 	- If you do not specify this parameter or if you set a value smaller than 10, the default value is 10.
	//
	// 	- If you set a value greater than 100, the default value is 100.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// caeba0bbb2be03f84eb48b699f0a****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The page number. Pages start from page 1.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page. Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 50
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the disk. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The snapshot chain IDs. You can specify a JSON array that contains a maximum of 100 snapshot chain IDs. Separate the snapshot chain IDs with commas (,).
	//
	// example:
	//
	// ["sl-bp1grgphbcc9brb5****", "sl-bp1c4izumvq0i5bs****", … "sl-bp1akk7isz866dds****"]
	SnapshotLinkIds *string `json:"SnapshotLinkIds,omitempty" xml:"SnapshotLinkIds,omitempty"`
}

func (s DescribeSnapshotLinksRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSnapshotLinksRequest) GoString() string {
	return s.String()
}

func (s *DescribeSnapshotLinksRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *DescribeSnapshotLinksRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeSnapshotLinksRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeSnapshotLinksRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeSnapshotLinksRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeSnapshotLinksRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeSnapshotLinksRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeSnapshotLinksRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeSnapshotLinksRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeSnapshotLinksRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeSnapshotLinksRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeSnapshotLinksRequest) GetSnapshotLinkIds() *string {
	return s.SnapshotLinkIds
}

func (s *DescribeSnapshotLinksRequest) SetDiskIds(v string) *DescribeSnapshotLinksRequest {
	s.DiskIds = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetInstanceId(v string) *DescribeSnapshotLinksRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetMaxResults(v int32) *DescribeSnapshotLinksRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetNextToken(v string) *DescribeSnapshotLinksRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetOwnerAccount(v string) *DescribeSnapshotLinksRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetOwnerId(v int64) *DescribeSnapshotLinksRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetPageNumber(v int32) *DescribeSnapshotLinksRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetPageSize(v int32) *DescribeSnapshotLinksRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetRegionId(v string) *DescribeSnapshotLinksRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetResourceOwnerAccount(v string) *DescribeSnapshotLinksRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetResourceOwnerId(v int64) *DescribeSnapshotLinksRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) SetSnapshotLinkIds(v string) *DescribeSnapshotLinksRequest {
	s.SnapshotLinkIds = &v
	return s
}

func (s *DescribeSnapshotLinksRequest) Validate() error {
	return dara.Validate(s)
}
