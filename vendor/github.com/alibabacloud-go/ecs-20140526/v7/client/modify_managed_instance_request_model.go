// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyManagedInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ModifyManagedInstanceRequest
	GetInstanceId() *string
	SetInstanceName(v string) *ModifyManagedInstanceRequest
	GetInstanceName() *string
	SetOwnerAccount(v string) *ModifyManagedInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyManagedInstanceRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyManagedInstanceRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyManagedInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyManagedInstanceRequest
	GetResourceOwnerId() *int64
}

type ModifyManagedInstanceRequest struct {
	// The new name of the managed instance. The name must be 1 to 128 characters in length. It must start with a letter and cannot start with a special character or a digit. It can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:) and cannot start with `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// mi-hz01nmcf****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The new name of the managed instance. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with a special character or a digit. The name can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:) and cannot start with `http://` or `https://`.
	//
	// This parameter is required.
	//
	// example:
	//
	// testInstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Supported regions: China (Qingdao), China (Beijing), China (Zhangjiakou), China (Hohhot), China (Ulanqab), China (Hangzhou), China (Shanghai), China (Shenzhen), China (Heyuan), China (Guangzhou), China (Chengdu), China (Hong Kong), Singapore, Japan (Tokyo), US (Silicon Valley), and US (Virginia).
	//
	// You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
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

func (s ModifyManagedInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyManagedInstanceRequest) GoString() string {
	return s.String()
}

func (s *ModifyManagedInstanceRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyManagedInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ModifyManagedInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyManagedInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyManagedInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyManagedInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyManagedInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyManagedInstanceRequest) SetInstanceId(v string) *ModifyManagedInstanceRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyManagedInstanceRequest) SetInstanceName(v string) *ModifyManagedInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *ModifyManagedInstanceRequest) SetOwnerAccount(v string) *ModifyManagedInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyManagedInstanceRequest) SetOwnerId(v int64) *ModifyManagedInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyManagedInstanceRequest) SetRegionId(v string) *ModifyManagedInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyManagedInstanceRequest) SetResourceOwnerAccount(v string) *ModifyManagedInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyManagedInstanceRequest) SetResourceOwnerId(v int64) *ModifyManagedInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyManagedInstanceRequest) Validate() error {
	return dara.Validate(s)
}
