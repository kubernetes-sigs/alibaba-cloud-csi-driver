// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTerminalSessionsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DescribeTerminalSessionsRequest
	GetInstanceId() *string
	SetMaxResults(v int32) *DescribeTerminalSessionsRequest
	GetMaxResults() *int32
	SetNextToken(v string) *DescribeTerminalSessionsRequest
	GetNextToken() *string
	SetOwnerAccount(v string) *DescribeTerminalSessionsRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeTerminalSessionsRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DescribeTerminalSessionsRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeTerminalSessionsRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeTerminalSessionsRequest
	GetResourceOwnerId() *int64
	SetSessionId(v string) *DescribeTerminalSessionsRequest
	GetSessionId() *string
}

type DescribeTerminalSessionsRequest struct {
	// The instance ID.
	//
	// example:
	//
	// i-bp1i7gg30r52z2em****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum number of entries per page.
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
	// AAAAAdDWBF2****
	NextToken    *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The session ID.
	//
	// example:
	//
	// s-hz023od0x9****
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
}

func (s DescribeTerminalSessionsRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeTerminalSessionsRequest) GoString() string {
	return s.String()
}

func (s *DescribeTerminalSessionsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeTerminalSessionsRequest) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *DescribeTerminalSessionsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTerminalSessionsRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeTerminalSessionsRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeTerminalSessionsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeTerminalSessionsRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeTerminalSessionsRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeTerminalSessionsRequest) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribeTerminalSessionsRequest) SetInstanceId(v string) *DescribeTerminalSessionsRequest {
	s.InstanceId = &v
	return s
}

func (s *DescribeTerminalSessionsRequest) SetMaxResults(v int32) *DescribeTerminalSessionsRequest {
	s.MaxResults = &v
	return s
}

func (s *DescribeTerminalSessionsRequest) SetNextToken(v string) *DescribeTerminalSessionsRequest {
	s.NextToken = &v
	return s
}

func (s *DescribeTerminalSessionsRequest) SetOwnerAccount(v string) *DescribeTerminalSessionsRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeTerminalSessionsRequest) SetOwnerId(v int64) *DescribeTerminalSessionsRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeTerminalSessionsRequest) SetRegionId(v string) *DescribeTerminalSessionsRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeTerminalSessionsRequest) SetResourceOwnerAccount(v string) *DescribeTerminalSessionsRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeTerminalSessionsRequest) SetResourceOwnerId(v int64) *DescribeTerminalSessionsRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeTerminalSessionsRequest) SetSessionId(v string) *DescribeTerminalSessionsRequest {
	s.SessionId = &v
	return s
}

func (s *DescribeTerminalSessionsRequest) Validate() error {
	return dara.Validate(s)
}
