// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCommandRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommandId(v string) *DeleteCommandRequest
	GetCommandId() *string
	SetOwnerAccount(v string) *DeleteCommandRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *DeleteCommandRequest
	GetOwnerId() *int64
	SetRegionId(v string) *DeleteCommandRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *DeleteCommandRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *DeleteCommandRequest
	GetResourceOwnerId() *int64
}

type DeleteCommandRequest struct {
	// The ID of the command. You can call the [DescribeCommands](https://help.aliyun.com/document_detail/64843.html) operation to query all available command IDs.
	//
	// This parameter is required.
	//
	// example:
	//
	// c-4d34302d02424c5c8e10281e3a31****
	CommandId    *string `json:"CommandId,omitempty" xml:"CommandId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the command. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s DeleteCommandRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCommandRequest) GoString() string {
	return s.String()
}

func (s *DeleteCommandRequest) GetCommandId() *string {
	return s.CommandId
}

func (s *DeleteCommandRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *DeleteCommandRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *DeleteCommandRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *DeleteCommandRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *DeleteCommandRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *DeleteCommandRequest) SetCommandId(v string) *DeleteCommandRequest {
	s.CommandId = &v
	return s
}

func (s *DeleteCommandRequest) SetOwnerAccount(v string) *DeleteCommandRequest {
	s.OwnerAccount = &v
	return s
}

func (s *DeleteCommandRequest) SetOwnerId(v int64) *DeleteCommandRequest {
	s.OwnerId = &v
	return s
}

func (s *DeleteCommandRequest) SetRegionId(v string) *DeleteCommandRequest {
	s.RegionId = &v
	return s
}

func (s *DeleteCommandRequest) SetResourceOwnerAccount(v string) *DeleteCommandRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *DeleteCommandRequest) SetResourceOwnerId(v int64) *DeleteCommandRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *DeleteCommandRequest) Validate() error {
	return dara.Validate(s)
}
