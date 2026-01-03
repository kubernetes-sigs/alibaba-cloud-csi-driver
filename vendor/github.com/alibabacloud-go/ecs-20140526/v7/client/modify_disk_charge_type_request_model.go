// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDiskChargeTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAutoPay(v bool) *ModifyDiskChargeTypeRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyDiskChargeTypeRequest
	GetClientToken() *string
	SetDiskChargeType(v string) *ModifyDiskChargeTypeRequest
	GetDiskChargeType() *string
	SetDiskIds(v string) *ModifyDiskChargeTypeRequest
	GetDiskIds() *string
	SetInstanceId(v string) *ModifyDiskChargeTypeRequest
	GetInstanceId() *string
	SetOwnerAccount(v string) *ModifyDiskChargeTypeRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyDiskChargeTypeRequest
	GetOwnerId() *int64
	SetRegionId(v string) *ModifyDiskChargeTypeRequest
	GetRegionId() *string
	SetResourceOwnerAccount(v string) *ModifyDiskChargeTypeRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyDiskChargeTypeRequest
	GetResourceOwnerId() *int64
}

type ModifyDiskChargeTypeRequest struct {
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 	- true (default): The payment is automatically completed. Maintain sufficient balance in your account. Otherwise, your order becomes invalid and must be canceled.
	//
	// 	- false: An order is generated but no payment is made. If your account balance is insufficient, you can set AutoPay to false to generate an unpaid order. Then, log on to the **Expenses and Costs console**, go to the [Orders page](https://usercenter2-intl.aliyun.com/order/list), and pay for the order.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The **token*	- can contain only ASCII characters and cannot exceed 64 characters in length. For more information, see [Ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The new billing method of the disk. Valid values:
	//
	// 	- PrePaid (default): changes the billing method from pay-as-you-go to subscription.
	//
	// 	- PostPaid: changes the billing method from subscription to pay-as-you-go.
	//
	// >  When you change the billing method of a pay-as-you-go disk to subscription, the billing cycle of the disk is automatically synchronized with that of the associated ECS instance.
	//
	// example:
	//
	// PostPaid
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The IDs of disks. The value is a JSON array that consists of up to 16 disk IDs. Separate the disk IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// [“d-bp67acfmxazb4ph****”, “d-bp67acfmxazb4pi****”, … “d-bp67acfmxazb4pj****”]
	DiskIds *string `json:"DiskIds,omitempty" xml:"DiskIds,omitempty"`
	// The ID of the instance to which disks are attached.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp1i778bq705cvx1****
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
	RegionId             *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
}

func (s ModifyDiskChargeTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyDiskChargeTypeRequest) GoString() string {
	return s.String()
}

func (s *ModifyDiskChargeTypeRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyDiskChargeTypeRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyDiskChargeTypeRequest) GetDiskChargeType() *string {
	return s.DiskChargeType
}

func (s *ModifyDiskChargeTypeRequest) GetDiskIds() *string {
	return s.DiskIds
}

func (s *ModifyDiskChargeTypeRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyDiskChargeTypeRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyDiskChargeTypeRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyDiskChargeTypeRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifyDiskChargeTypeRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyDiskChargeTypeRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyDiskChargeTypeRequest) SetAutoPay(v bool) *ModifyDiskChargeTypeRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyDiskChargeTypeRequest) SetClientToken(v string) *ModifyDiskChargeTypeRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyDiskChargeTypeRequest) SetDiskChargeType(v string) *ModifyDiskChargeTypeRequest {
	s.DiskChargeType = &v
	return s
}

func (s *ModifyDiskChargeTypeRequest) SetDiskIds(v string) *ModifyDiskChargeTypeRequest {
	s.DiskIds = &v
	return s
}

func (s *ModifyDiskChargeTypeRequest) SetInstanceId(v string) *ModifyDiskChargeTypeRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyDiskChargeTypeRequest) SetOwnerAccount(v string) *ModifyDiskChargeTypeRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyDiskChargeTypeRequest) SetOwnerId(v int64) *ModifyDiskChargeTypeRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyDiskChargeTypeRequest) SetRegionId(v string) *ModifyDiskChargeTypeRequest {
	s.RegionId = &v
	return s
}

func (s *ModifyDiskChargeTypeRequest) SetResourceOwnerAccount(v string) *ModifyDiskChargeTypeRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyDiskChargeTypeRequest) SetResourceOwnerId(v int64) *ModifyDiskChargeTypeRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyDiskChargeTypeRequest) Validate() error {
	return dara.Validate(s)
}
