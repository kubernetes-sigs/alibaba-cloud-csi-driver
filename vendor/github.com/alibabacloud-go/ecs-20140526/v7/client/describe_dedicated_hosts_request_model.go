// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDedicatedHostsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedHostClusterId(v string) *DescribeDedicatedHostsRequest
	GetDedicatedHostClusterId() *string
	SetDedicatedHostIds(v string) *DescribeDedicatedHostsRequest
	GetDedicatedHostIds() *string
	SetDedicatedHostName(v string) *DescribeDedicatedHostsRequest
	GetDedicatedHostName() *string
	SetDedicatedHostType(v string) *DescribeDedicatedHostsRequest
	GetDedicatedHostType() *string
	SetLockReason(v string) *DescribeDedicatedHostsRequest
	GetLockReason() *string
	SetMaxResults(v int32) *DescribeDedicatedHostsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeDedicatedHostsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeDedicatedHostsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeDedicatedHostsRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeDedicatedHostsRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDedicatedHostsRequest
	GetPageSize() *int32
	SetQueryInventory(v bool) *DescribeDedicatedHostsRequest
	GetQueryInventory() *bool
	SetRegionId(v string) *DescribeDedicatedHostsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeDedicatedHostsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeDedicatedHostsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeDedicatedHostsRequest
	GetResourceOwnerId() *int64
	SetSocketDetails(v string) *DescribeDedicatedHostsRequest
	GetSocketDetails() *string
	SetStatus(v string) *DescribeDedicatedHostsRequest
	GetStatus() *string
	SetTag(v []*DescribeDedicatedHostsRequestTag) *DescribeDedicatedHostsRequest
	GetTag() []*DescribeDedicatedHostsRequestTag
	SetZoneId(v string) *DescribeDedicatedHostsRequest
	GetZoneId() *string
}

type DescribeDedicatedHostsRequest struct {
	// The ID of the dedicated host cluster.
	//
	// example:
	//
	// dc-bp12wlf6am0vz9v2****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
	// The list of DDH IDs. You can specify up to 100 deployment set IDs in each request. Separate the deployment set IDs with commas (,).
	//
	// example:
	//
	// ["dh-bp165p6xk2tlw61e****", "dh-bp1f9vxmno7emy96****"]
	DedicatedHostIds *string `json:"DedicatedHostIds,omitempty" xml:"DedicatedHostIds,omitempty"`
	// The name of the dedicated host.
	//
	// example:
	//
	// MyDDHTestName
	DedicatedHostName *string `json:"DedicatedHostName,omitempty" xml:"DedicatedHostName,omitempty"`
	// The type of the DDH. You can call the [DescribeDedicatedHostTypes](https://help.aliyun.com/document_detail/134240.html) operation to query the most recent list of DDH types.
	//
	// example:
	//
	// ddh.g5
	DedicatedHostType *string `json:"DedicatedHostType,omitempty" xml:"DedicatedHostType,omitempty"`
	// The reason why the dedicated host is locked. Valid values:
	//
	// 	- financial: The dedicated host is locked due to overdue payments.
	//
	// 	- security: The dedicated host is locked due to security reasons.
	//
	// example:
	//
	// financial
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
	// The maximum number of entries per page. If you specify this parameter, both MaxResults and NextToken are used for a paged query.
	//
	// Valid values: 1 to 100.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of NextToken.
	//
	// example:
	//
	// e71d8a535bd9cc11
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// >  This parameter will be removed in the future. You can use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// >  This parameter will be removed in the future. You can use NextToken and MaxResults for a paged query.
	//
	// example:
	//
	// 10
	PageSize       *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	QueryInventory *bool  `json:"QueryInventory,omitempty" xml:"QueryInventory,omitempty"`
	// The region ID of the dedicated host. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the dedicated host belongs. When this parameter is specified to query resources, up to 1,000 resources that belong to the specified resource group can be displayed in the response.
	//
	// > Resources in the default resource group are displayed in the response regardless of how this parameter is set.
	//
	// example:
	//
	// rg-aek3b6jzp66****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to display socket information. You can view the remaining resources (vCPUs, memory usage, remaining resources, and total resources) based on the capacity information of the socket dimension. Then you can determine whether ECS instances of the corresponding specifications can be created. Valid values:
	//
	// 	- true Only some DDHs support the information about resources in the socket dimension. For more information, see [View and export information about DDHs](https://help.aliyun.com/document_detail/68989.html).
	//
	// 	- false
	//
	// >  Each DDH generally has two CPUs, and each CPU corresponds to Socket 0 and Socket 1. To maximize the performance of an ECS instance on a DDH, ECS instances are not created across sockets.
	//
	// 	- If one socket has available computing resources for creating the ECS instance, creation succeeds.
	//
	// 	- If not, creation fails even if the combined available resources of both sockets are sufficient. Although the remaining resources of the two sockets on the DDH are larger than the ECS instance type, the ECS instance cannot be created.
	//
	// example:
	//
	// true
	SocketDetails *string `json:"SocketDetails,omitempty" xml:"SocketDetails,omitempty"`
	// The service state of the dedicated host. Valid values:
	//
	// 	- Available: The dedicated host is running normally.
	//
	// 	- UnderAssessment: The dedicated host is available but has potential risks that may cause the ECS instances on the dedicated host to fail.
	//
	// 	- PermanentFailure: The dedicated host encounters permanent failures and is unavailable.
	//
	// 	- TempUnavailable: The dedicated host is temporarily unavailable.
	//
	// 	- Redeploying: The dedicated host is being restored.
	//
	// Default value: Available.
	//
	// example:
	//
	// Available
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The list of tags. The list length ranges from 0 to 20.
	Tag []*DescribeDedicatedHostsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The zone ID of the dedicated host. You can call the [DescribeZones](https://help.aliyun.com/document_detail/25610.html) operation to query the most recent zone list.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDedicatedHostsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsRequest) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsRequest) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *DescribeDedicatedHostsRequest) GetDedicatedHostIds() *string {
	return s.DedicatedHostIds
}

func (s *DescribeDedicatedHostsRequest) GetDedicatedHostName() *string {
	return s.DedicatedHostName
}

func (s *DescribeDedicatedHostsRequest) GetDedicatedHostType() *string {
	return s.DedicatedHostType
}

func (s *DescribeDedicatedHostsRequest) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDedicatedHostsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeDedicatedHostsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDedicatedHostsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeDedicatedHostsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeDedicatedHostsRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDedicatedHostsRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDedicatedHostsRequest) GetQueryInventory() *bool {
	return s.QueryInventory
}

func (s *DescribeDedicatedHostsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDedicatedHostsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDedicatedHostsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeDedicatedHostsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeDedicatedHostsRequest) GetSocketDetails() *string {
	return s.SocketDetails
}

func (s *DescribeDedicatedHostsRequest) GetStatus() *string {
	return s.Status
}

func (s *DescribeDedicatedHostsRequest) GetTag() []*DescribeDedicatedHostsRequestTag {
	return s.Tag
}

func (s *DescribeDedicatedHostsRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostClusterId(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostIds(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostIds = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostName(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostName = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetDedicatedHostType(v string) *DescribeDedicatedHostsRequest {
	s.DedicatedHostType = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetLockReason(v string) *DescribeDedicatedHostsRequest {
	s.LockReason = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetMaxResults(v int32) *DescribeDedicatedHostsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetNextToken(v string) *DescribeDedicatedHostsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetOwnerAccount(v string) *DescribeDedicatedHostsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetOwnerId(v int64) *DescribeDedicatedHostsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetPageNumber(v int32) *DescribeDedicatedHostsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetPageSize(v int32) *DescribeDedicatedHostsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetQueryInventory(v bool) *DescribeDedicatedHostsRequest {
	s.QueryInventory = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetRegionId(v string) *DescribeDedicatedHostsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetResourceGroupId(v string) *DescribeDedicatedHostsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetResourceOwnerAccount(v string) *DescribeDedicatedHostsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetResourceOwnerId(v int64) *DescribeDedicatedHostsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetSocketDetails(v string) *DescribeDedicatedHostsRequest {
	s.SocketDetails = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetStatus(v string) *DescribeDedicatedHostsRequest {
	s.Status = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetTag(v []*DescribeDedicatedHostsRequestTag) *DescribeDedicatedHostsRequest {
	s.Tag = v
	return s
}

func (s *DescribeDedicatedHostsRequest) SetZoneId(v string) *DescribeDedicatedHostsRequest {
	s.ZoneId = &v
	return s
}

func (s *DescribeDedicatedHostsRequest) Validate() error {
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

type DescribeDedicatedHostsRequestTag struct {
	// The key of tag N of the DDH. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N of the DDH. You can specify empty strings as tag values. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeDedicatedHostsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDedicatedHostsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeDedicatedHostsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeDedicatedHostsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeDedicatedHostsRequestTag) SetKey(v string) *DescribeDedicatedHostsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeDedicatedHostsRequestTag) SetValue(v string) *DescribeDedicatedHostsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeDedicatedHostsRequestTag) Validate() error {
	return dara.Validate(s)
}
