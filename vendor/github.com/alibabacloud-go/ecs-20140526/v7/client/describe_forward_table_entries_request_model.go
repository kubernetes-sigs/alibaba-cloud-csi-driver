// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeForwardTableEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetForwardEntryId(v string) *DescribeForwardTableEntriesRequest
	GetForwardEntryId() *string
	SetForwardTableId(v string) *DescribeForwardTableEntriesRequest
	GetForwardTableId() *string
	SetOwnerAccount(v string) *DescribeForwardTableEntriesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DescribeForwardTableEntriesRequest
	GetOwnerId() *int64
	SetPageNumber(v int32) *DescribeForwardTableEntriesRequest
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeForwardTableEntriesRequest
	GetPageSize() *int32
	SetRegionId(v string) *DescribeForwardTableEntriesRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DescribeForwardTableEntriesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DescribeForwardTableEntriesRequest
	GetResourceOwnerId() *int64
}

type DescribeForwardTableEntriesRequest struct {
	ForwardEntryId *string `json:"ForwardEntryId,omitempty" xml:"ForwardEntryId,omitempty"`
	// This parameter is required.
	ForwardTableId *string `json:"ForwardTableId,omitempty" xml:"ForwardTableId,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageNumber     *int32  `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	PageSize       *int32  `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// This parameter is required.
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s DescribeForwardTableEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeForwardTableEntriesRequest) GoString() string {
	return s.String()
}

func (s *DescribeForwardTableEntriesRequest) GetForwardEntryId() *string {
	return s.ForwardEntryId
}

func (s *DescribeForwardTableEntriesRequest) GetForwardTableId() *string {
	return s.ForwardTableId
}

func (s *DescribeForwardTableEntriesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DescribeForwardTableEntriesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DescribeForwardTableEntriesRequest) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeForwardTableEntriesRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeForwardTableEntriesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeForwardTableEntriesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DescribeForwardTableEntriesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DescribeForwardTableEntriesRequest) SetForwardEntryId(v string) *DescribeForwardTableEntriesRequest {
	s.ForwardEntryId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetForwardTableId(v string) *DescribeForwardTableEntriesRequest {
	s.ForwardTableId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetOwnerAccount(v string) *DescribeForwardTableEntriesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetOwnerId(v int64) *DescribeForwardTableEntriesRequest {
	s.OwnerId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetPageNumber(v int32) *DescribeForwardTableEntriesRequest {
	s.PageNumber = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetPageSize(v int32) *DescribeForwardTableEntriesRequest {
	s.PageSize = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetRegionId(v string) *DescribeForwardTableEntriesRequest {
	s.RegionId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetResourceOwnerAccount(v string) *DescribeForwardTableEntriesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) SetResourceOwnerId(v int64) *DescribeForwardTableEntriesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DescribeForwardTableEntriesRequest) Validate() error {
	return dara.Validate(s)
}
