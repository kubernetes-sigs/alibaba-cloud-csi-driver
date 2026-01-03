// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeInvocationsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DescribeInvocationsRequest
	GetCommandId() *string
	SetCommandName(v string) *DescribeInvocationsRequest
	GetCommandName() *string
	SetCommandType(v string) *DescribeInvocationsRequest
	GetCommandType() *string
	SetContentEncoding(v string) *DescribeInvocationsRequest
	GetContentEncoding() *string
	SetIncludeOutput(v bool) *DescribeInvocationsRequest
	GetIncludeOutput() *bool
	SetInstanceId(v string) *DescribeInvocationsRequest
	GetInstanceId() *string
	SetInvokeId(v string) *DescribeInvocationsRequest
	GetInvokeId() *string
	SetInvokeStatus(v string) *DescribeInvocationsRequest
	GetInvokeStatus() *string
	SetMaxResults(v int32) *DescribeInvocationsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeInvocationsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeInvocationsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeInvocationsRequest
	GetOwnerId() *int64
	SetPageNumber(v int64) *DescribeInvocationsRequest
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeInvocationsRequest
	GetPageSize() *int64
	SetRegionId(v string) *DescribeInvocationsRequest
	GetRegionId() *string
	SetRepeatMode(v string) *DescribeInvocationsRequest
	GetRepeatMode() *string
	SetResourceGroupId(v string) *DescribeInvocationsRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *DescribeInvocationsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeInvocationsRequest
	GetResourceOwnerId() *int64
	SetTag(v []*DescribeInvocationsRequestTag) *DescribeInvocationsRequest
	GetTag() []*DescribeInvocationsRequestTag
	SetTimed(v bool) *DescribeInvocationsRequest
	GetTimed() *bool
}

type DescribeInvocationsRequest struct {
	// $.parameters[15].schema.items.description
	//
	// example:
	//
	// c-hz0jdfwcsr****
	CommandId *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	// $.parameters[15].schema.items.example
	//
	// example:
	//
	// CommandTestName
	CommandName *string `json:"CommandName,omitempty" xml:"CommandName,omitempty"`
	// $.parameters[15].schema.items.enumValueTitles
	//
	// example:
	//
	// RunShellScript
	CommandType *string `json:"CommandType,omitempty" xml:"CommandType,omitempty"`
	// example:
	//
	// PlainText
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// example:
	//
	// false
	IncludeOutput *bool `json:"IncludeOutput,omitempty" xml:"IncludeOutput,omitempty"`
	// $.parameters[15].schema.enumValueTitles
	//
	// example:
	//
	// i-bp1i7gg30r52z2em****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// $.parameters[15].schema.items.properties.Value.enumValueTitles
	//
	// example:
	//
	// t-hz0jdfwd9f****
	InvokeId *string `json:"InvokeId,omitempty" xml:"InvokeId,omitempty"`
	// $.parameters[15].schema.example
	//
	// example:
	//
	// Finished
	InvokeStatus *string `json:"InvokeStatus,omitempty" xml:"InvokeStatus,omitempty"`
	// acs:ecs:{#regionId}:{#accountId}:instance/*
	//
	// example:
	//
	// 10
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// Instance
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
	// Command
	//
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// $.parameters[15].schema.items.properties.Value.description
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// FEATUREecsXZ3H4M
	//
	// example:
	//
	// Once
	RepeatMode *string `json:"RepeatMode,omitempty" xml:"RepeatMode,omitempty"`
	// $.parameters[15].schema.items.properties.Value.example
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// dubbo
	Tag []*DescribeInvocationsRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// $.parameters[15].schema.description
	//
	// example:
	//
	// true
	Timed *bool `json:"Timed,omitempty" xml:"Timed,omitempty"`
}

func (s DescribeInvocationsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsRequest) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DescribeInvocationsRequest) GetCommandName() *string {
	return s.CommandName
}

func (s *DescribeInvocationsRequest) GetCommandType() *string {
	return s.CommandType
}

func (s *DescribeInvocationsRequest) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *DescribeInvocationsRequest) GetIncludeOutput() *bool {
	return s.IncludeOutput
}

func (s *DescribeInvocationsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeInvocationsRequest) GetInvokeId() *string {
	return s.InvokeId
}

func (s *DescribeInvocationsRequest) GetInvokeStatus() *string {
	return s.InvokeStatus
}

func (s *DescribeInvocationsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeInvocationsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeInvocationsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeInvocationsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeInvocationsRequest) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeInvocationsRequest) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeInvocationsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeInvocationsRequest) GetRepeatMode() *string {
	return s.RepeatMode
}

func (s *DescribeInvocationsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeInvocationsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeInvocationsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeInvocationsRequest) GetTag() []*DescribeInvocationsRequestTag {
	return s.Tag
}

func (s *DescribeInvocationsRequest) GetTimed() *bool {
	return s.Timed
}

func (s *DescribeInvocationsRequest) SetCommandId(v string) *DescribeInvocationsRequest {
	s.CommandId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetCommandName(v string) *DescribeInvocationsRequest {
	s.CommandName = &v
	return s
}

func (s *DescribeInvocationsRequest) SetCommandType(v string) *DescribeInvocationsRequest {
	s.CommandType = &v
	return s
}

func (s *DescribeInvocationsRequest) SetContentEncoding(v string) *DescribeInvocationsRequest {
	s.ContentEncoding = &v
	return s
}

func (s *DescribeInvocationsRequest) SetIncludeOutput(v bool) *DescribeInvocationsRequest {
	s.IncludeOutput = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInstanceId(v string) *DescribeInvocationsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeId(v string) *DescribeInvocationsRequest {
	s.InvokeId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetInvokeStatus(v string) *DescribeInvocationsRequest {
	s.InvokeStatus = &v
	return s
}

func (s *DescribeInvocationsRequest) SetMaxResults(v int32) *DescribeInvocationsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeInvocationsRequest) SetNextToken(v string) *DescribeInvocationsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeInvocationsRequest) SetOwnerAccount(v string) *DescribeInvocationsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeInvocationsRequest) SetOwnerId(v int64) *DescribeInvocationsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetPageNumber(v int64) *DescribeInvocationsRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeInvocationsRequest) SetPageSize(v int64) *DescribeInvocationsRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeInvocationsRequest) SetRegionId(v string) *DescribeInvocationsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetRepeatMode(v string) *DescribeInvocationsRequest {
	s.RepeatMode = &v
	return s
}

func (s *DescribeInvocationsRequest) SetResourceGroupId(v string) *DescribeInvocationsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetResourceOwnerAccount(v string) *DescribeInvocationsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeInvocationsRequest) SetResourceOwnerId(v int64) *DescribeInvocationsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeInvocationsRequest) SetTag(v []*DescribeInvocationsRequestTag) *DescribeInvocationsRequest {
	s.Tag = v
	return s
}

func (s *DescribeInvocationsRequest) SetTimed(v bool) *DescribeInvocationsRequest {
	s.Timed = &v
	return s
}

func (s *DescribeInvocationsRequest) Validate() error {
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

type DescribeInvocationsRequestTag struct {
	// The command task ID.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The ID of the resource group. After you set this parameter, command execution results in the specified resource group are queried.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s DescribeInvocationsRequestTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeInvocationsRequestTag) GoString() string {
	return s.String()
}

func (s *DescribeInvocationsRequestTag) GetKey() *string {
	return s.Key
}

func (s *DescribeInvocationsRequestTag) GetValue() *string {
	return s.Value
}

func (s *DescribeInvocationsRequestTag) SetKey(v string) *DescribeInvocationsRequestTag {
	s.Key = &v
	return s
}

func (s *DescribeInvocationsRequestTag) SetValue(v string) *DescribeInvocationsRequestTag {
	s.Value = &v
	return s
}

func (s *DescribeInvocationsRequestTag) Validate() error {
	return dara.Validate(s)
}
