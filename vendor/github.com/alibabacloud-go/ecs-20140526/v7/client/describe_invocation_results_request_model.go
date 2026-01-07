// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DescribeInvocationResultsRequest
	GetCommandId() *string
	SetContentEncoding(v string) *DescribeInvocationResultsRequest
	GetContentEncoding() *string
	SetIncludeHistory(v bool) *DescribeInvocationResultsRequest
	GetIncludeHistory() *bool
	SetInstanceId(v string) *DescribeInvocationResultsRequest
	GetInstanceId() *string
	SetInvokeId(v string) *DescribeInvocationResultsRequest
	GetInvokeId() *string
	SetInvokeRecordStatus(v string) *DescribeInvocationResultsRequest
	GetInvokeRecordStatus() *string
	SetMaxResults(v int32) *DescribeInvocationResultsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInvocationResultsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeInvocationResultsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInvocationResultsRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeInvocationResultsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInvocationResultsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeInvocationResultsRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *DescribeInvocationResultsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeInvocationResultsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInvocationResultsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeInvocationResultsRequestTag) *DescribeInvocationResultsRequest
	GetTag() []*DescribeInvocationResultsRequestTag
}

type DescribeInvocationResultsRequest struct {
	// $.parameters[11].schema.example
	//
	// example:
	//
	// c-hz0jdfwcsr****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// <DescribeInvocationResultsResponse>
	//
	//     <RequestId>473469C7-AA6F-4DC5-B3DB-A3DC0DE*****</RequestId>
	//
	//     <Invocation>
	//
	//         <InvocationResults>
	//
	//             <InvocationResult>
	//
	//                 <Dropped>0</Dropped>
	//
	//                 <InvocationStatus>Success</InvocationStatus>
	//
	//                 <InstanceId>i-bp1i7gg30r52z2em****</InstanceId>
	//
	//                 <ExitCode>0</ExitCode>
	//
	//                 <ErrorInfo>the specified instance does not exists</ErrorInfo>
	//
	//                 <StartTime>2019-12-20T06:15:55Z</StartTime>
	//
	//                 <Repeats>0</Repeats>
	//
	//                 <InvokeRecordStatus>Running</InvokeRecordStatus>
	//
	//                 <FinishedTime>2019-12-20T06:15:56Z</FinishedTime>
	//
	//                 <Output>MTU6MzA6MDEK</Output>
	//
	//                 <CommandId>c-hz0jdfwcsr****</CommandId>
	//
	//                 <ErrorCode>InstanceNotExists</ErrorCode>
	//
	//                 <InvokeId>t-hz0jdfwd9f****</InvokeId>
	//
	//                 <StopTime>2020-01-19T09:15:47Z</StopTime>
	//
	//                 <ContainerId>ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****</ContainerId>
	//
	//                 <ContainerName>test-container</ContainerName>
	//
	//                 <Tags>
	//
	//                     <TagKey>owner</TagKey>
	//
	//                     <TagValue>zhangsan</TagValue>
	//
	//                 </Tags>
	//
	//             </InvocationResult>
	//
	//         </InvocationResults>
	//
	//         <TotalCount>1</TotalCount>
	//
	//         <PageSize>1</PageSize>
	//
	//         <PageNumber>1</PageNumber>
	//
	//     </Invocation>
	//
	// </DescribeInvocationResultsResponse>
	//
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// {
	//
	//   "RequestId" : "473469C7-AA6F-4DC5-B3DB-A3DC0DE*****",
	//
	//   "Invocation" : {
	//
	//     "InvocationResults" : {
	//
	//       "InvocationResult" : [ {
	//
	//         "Dropped" : 0,
	//
	//         "InvocationStatus" : "Success",
	//
	//         "InstanceId" : "i-bp1i7gg30r52z2em****",
	//
	//         "ExitCode" : 0,
	//
	//         "ErrorInfo" : "the specified instance does not exists",
	//
	//         "StartTime" : "2019-12-20T06:15:55Z",
	//
	//         "Repeats" : 0,
	//
	//         "InvokeRecordStatus" : "Running",
	//
	//         "FinishedTime" : "2019-12-20T06:15:56Z",
	//
	//         "Output" : "MTU6MzA6MDEK",
	//
	//         "CommandId" : "c-hz0jdfwcsr****",
	//
	//         "ErrorCode" : "InstanceNotExists",
	//
	//         "InvokeId" : "t-hz0jdfwd9f****",
	//
	//         "StopTime" : "2020-01-19T09:15:47Z",
	//
	//         "ContainerId":"ab141ddfbacfe02d9dbc25966ed971536124527097398d419a6746873fea****",
	//
	//         "ContainerName":"test-container",
	//
	//         "Tags": [
	//
	//                     {
	//
	//                         "TagKey": "owner",
	//
	//                         "TagValue": "zhangsan"
	//
	//                     }
	//
	//                 ]
	//
	//       } ]
	//
	//     },
	//
	//     "TotalCount" : 1,
	//
	//     "PageSize" : 1,
	//
	//     "PageNumber" : 1
	//
	//   }
	//
	// }
	//
	// example:
	//
	// false
	IncludeHistory *bool `json:"IncludeHistory,omitempty" xml:"IncludeHistory,omitempty"`
	// $.parameters[11].schema.description
	//
	// example:
	//
	// i-bp1i7gg30r52z2em****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// $.parameters[11].schema.items.enumValueTitles
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// $.parameters[11].schema.enumValueTitles
	//
	// example:
	//
	// Running
	InvokeRecordStatus *string `json:"InvokeRecordStatus,omitempty" xml:"InvokeRecordStatus,omitempty"`
	// FEATUREecsXZ3H4M
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// dubbo
	//
	// example:
	//
	// AAAAAdDWBF2
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// acs:ecs:{#regionId}:{#accountId}:command/*
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// acs:ecs:{#regionId}:{#accountId}:instance/*
	//
	// example:
	//
	// 1
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// $.parameters[11].schema.items.description
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// $.parameters[11].schema.items.example
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The region ID of the command. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	Tag []*DescribeInvocationResultsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeInvocationResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeInvocationResultsRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeInvocationResultsRequest) GetIncludeHistory() *bool {
	return s.IncludeHistory
}

func (s *DescribeInvocationResultsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInvocationResultsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationResultsRequest) GetInvokeRecordStatus() *string {
	return s.InvokeRecordStatus
}

func (s *DescribeInvocationResultsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInvocationResultsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInvocationResultsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInvocationResultsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInvocationResultsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInvocationResultsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInvocationResultsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInvocationResultsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInvocationResultsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInvocationResultsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInvocationResultsRequest) GetTag() []*DescribeInvocationResultsRequestTag {
	return s.Tag
}

func (s *DescribeInvocationResultsRequest) SetCommandId(v string) *DescribeInvocationResultsRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetContentEncoding(v string) *DescribeInvocationResultsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetIncludeHistory(v bool) *DescribeInvocationResultsRequest {
	s.IncludeHistory = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetInstanceId(v string) *DescribeInvocationResultsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetInvokeId(v string) *DescribeInvocationResultsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetInvokeRecordStatus(v string) *DescribeInvocationResultsRequest {
	s.InvokeRecordStatus = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetMaxResults(v int32) *DescribeInvocationResultsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetNextToken(v string) *DescribeInvocationResultsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetOwnerAccount(v string) *DescribeInvocationResultsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetOwnerId(v int64) *DescribeInvocationResultsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetPageNumber(v int64) *DescribeInvocationResultsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetPageSize(v int64) *DescribeInvocationResultsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetRegionId(v string) *DescribeInvocationResultsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetResourceGroupId(v string) *DescribeInvocationResultsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetResourceOwnerAccount(v string) *DescribeInvocationResultsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetResourceOwnerId(v int64) *DescribeInvocationResultsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInvocationResultsRequest) SetTag(v []*DescribeInvocationResultsRequestTag) *DescribeInvocationResultsRequest {
	s.Tag = v
	return s
}

func (s *DescribeInvocationResultsRequest) Validate() error {
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

type DescribeInvocationResultsRequestTag struct {
	// The ID of the instance.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the command task. You can call the [DescribeInvocations](https://help.aliyun.com/document_detail/64840.html) operation to query the IDs of all command tasks.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInvocationResultsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationResultsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInvocationResultsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInvocationResultsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInvocationResultsRequestTag) SetKey(v string) *DescribeInvocationResultsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInvocationResultsRequestTag) SetValue(v string) *DescribeInvocationResultsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInvocationResultsRequestTag) Validate() error {
	return dara.Validate(s)
}
