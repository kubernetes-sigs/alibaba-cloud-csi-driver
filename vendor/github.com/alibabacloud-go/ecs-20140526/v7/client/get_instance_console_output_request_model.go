// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceConsoleOutputRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetInstanceConsoleOutputRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *GetInstanceConsoleOutputRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *GetInstanceConsoleOutputRequest
	GetOwnerId() *int64
	SetRegionId(v string) *GetInstanceConsoleOutputRequest
	GetRegionId() *string
	SetRemoveSymbols(v bool) *GetInstanceConsoleOutputRequest
	GetRemoveSymbols() *bool
	SetResourceOwnerAccount(v string) *GetInstanceConsoleOutputRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetInstanceConsoleOutputRequest
	GetResourceOwnerId() *int64
}

type GetInstanceConsoleOutputRequest struct {
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1c1xhsrac2coiw****
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID of the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// Specifies whether to remove formatting symbols from the returned command output. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// example:
	//
	// false
	RemoveSymbols        *bool   `json:"RemoveSymbols,omitempty" xml:"RemoveSymbols,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s GetInstanceConsoleOutputRequest) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceConsoleOutputRequest) GoString() string {
	return s.String()
}

func (s *GetInstanceConsoleOutputRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetInstanceConsoleOutputRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *GetInstanceConsoleOutputRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetInstanceConsoleOutputRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetInstanceConsoleOutputRequest) GetRemoveSymbols() *bool {
	return s.RemoveSymbols
}

func (s *GetInstanceConsoleOutputRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetInstanceConsoleOutputRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetInstanceConsoleOutputRequest) SetInstanceId(v string) *GetInstanceConsoleOutputRequest {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceConsoleOutputRequest) SetOwnerAccount(v string) *GetInstanceConsoleOutputRequest {
	s.OwnerAccount = &v
	return s
}

func (s *GetInstanceConsoleOutputRequest) SetOwnerId(v int64) *GetInstanceConsoleOutputRequest {
	s.OwnerId = &v
	return s
}

func (s *GetInstanceConsoleOutputRequest) SetRegionId(v string) *GetInstanceConsoleOutputRequest {
	s.RegionId = &v
	return s
}

func (s *GetInstanceConsoleOutputRequest) SetRemoveSymbols(v bool) *GetInstanceConsoleOutputRequest {
	s.RemoveSymbols = &v
	return s
}

func (s *GetInstanceConsoleOutputRequest) SetResourceOwnerAccount(v string) *GetInstanceConsoleOutputRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetInstanceConsoleOutputRequest) SetResourceOwnerId(v int64) *GetInstanceConsoleOutputRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetInstanceConsoleOutputRequest) Validate() error {
	return dara.Validate(s)
}
