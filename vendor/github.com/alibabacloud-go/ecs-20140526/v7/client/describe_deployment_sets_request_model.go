// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDeploymentSetsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeploymentSetIds(v string) *DescribeDeploymentSetsRequest
	GetDeploymentSetIds() *string
	SetDeploymentSetName(v string) *DescribeDeploymentSetsRequest
	GetDeploymentSetName() *string
	SetDomain(v string) *DescribeDeploymentSetsRequest
	GetDomain() *string
	SetGranularity(v string) *DescribeDeploymentSetsRequest
	GetGranularity() *string
	SetNetworkType(v string) *DescribeDeploymentSetsRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *DescribeDeploymentSetsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDeploymentSetsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDeploymentSetsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDeploymentSetsRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeDeploymentSetsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeDeploymentSetsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDeploymentSetsRequest
	GetResourceOwnerId() *int64
	SetStrategy(v string) *DescribeDeploymentSetsRequest
	GetStrategy() *string
}

type DescribeDeploymentSetsRequest struct {
	// The IDs of deployment sets. The value can be a JSON array that consists of deployment set IDs in the format of `["ds-xxxxxxxxx", "ds-yyyyyyyyy", ... "ds-zzzzzzzzz"]`. You can specify up to 100 deployment set IDs in each request. Separate the deployment set IDs with commas (,).
	//
	// example:
	//
	// ["ds-bp67acfmxazb4ph****", "ds-bp67acfmxazb4pi****", … "ds-bp67acfmxazb4pj****"]
	DeploymentSetIds *string `json:"DeploymentSetIds,omitempty" xml:"DeploymentSetIds,omitempty"`
	// The name of the deployment set. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain digits, letters, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testDeploymentSetName
	DeploymentSetName *string `json:"DeploymentSetName,omitempty" xml:"DeploymentSetName,omitempty"`
	// >  This parameter is deprecated.
	//
	// example:
	//
	// null
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// >  This parameter is deprecated.
	//
	// example:
	//
	// null
	Granularity *string `json:"Granularity,omitempty" xml:"Granularity,omitempty"`
	// >  This parameter is deprecated.
	//
	// example:
	//
	// null
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
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
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// Valid values: 1 to 50.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The region ID of the deployment set. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The deployment strategy. Valid values:
	//
	// 	- Availability: high availability strategy
	//
	// 	- AvailabilityGroup: high availability group strategy
	//
	// example:
	//
	// Availability
	Strategy *string `json:"Strategy,omitempty" xml:"Strategy,omitempty"`
}

func (s DescribeDeploymentSetsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDeploymentSetsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDeploymentSetsRequest) GetDeploymentSetIds() *string {
	return s.DeploymentSetIds
}

func (s *DescribeDeploymentSetsRequest) GetDeploymentSetName() *string {
	return s.DeploymentSetName
}

func (s *DescribeDeploymentSetsRequest) GetDomain() *string {
	return s.Domain
}

func (s *DescribeDeploymentSetsRequest) GetGranularity() *string {
	return s.Granularity
}

func (s *DescribeDeploymentSetsRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *DescribeDeploymentSetsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDeploymentSetsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDeploymentSetsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDeploymentSetsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDeploymentSetsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDeploymentSetsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDeploymentSetsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDeploymentSetsRequest) GetStrategy() *string {
	return s.Strategy
}

func (s *DescribeDeploymentSetsRequest) SetDeploymentSetIds(v string) *DescribeDeploymentSetsRequest {
	s.DeploymentSetIds = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetDeploymentSetName(v string) *DescribeDeploymentSetsRequest {
	s.DeploymentSetName = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetDomain(v string) *DescribeDeploymentSetsRequest {
	s.Domain = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetGranularity(v string) *DescribeDeploymentSetsRequest {
	s.Granularity = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetNetworkType(v string) *DescribeDeploymentSetsRequest {
	s.NetworkType = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetOwnerAccount(v string) *DescribeDeploymentSetsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetOwnerId(v int64) *DescribeDeploymentSetsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetPageNumber(v int32) *DescribeDeploymentSetsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetPageSize(v int32) *DescribeDeploymentSetsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetRegionId(v string) *DescribeDeploymentSetsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetResourceOwnerAccount(v string) *DescribeDeploymentSetsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetResourceOwnerId(v int64) *DescribeDeploymentSetsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) SetStrategy(v string) *DescribeDeploymentSetsRequest {
	s.Strategy = &v
	return s
}

func (s *DescribeDeploymentSetsRequest) Validate() error {
	return dara.Validate(s)
}
