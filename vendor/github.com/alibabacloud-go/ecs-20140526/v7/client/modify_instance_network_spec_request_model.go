// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceNetworkSpecRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllocatePublicIp(v bool) *ModifyInstanceNetworkSpecRequest
	GetAllocatePublicIp() *bool
	SetAutoPay(v bool) *ModifyInstanceNetworkSpecRequest
	GetAutoPay() *bool
	SetClientToken(v string) *ModifyInstanceNetworkSpecRequest
	GetClientToken() *string
	SetEndTime(v string) *ModifyInstanceNetworkSpecRequest
	GetEndTime() *string
	SetISP(v string) *ModifyInstanceNetworkSpecRequest
	GetISP() *string
	SetInstanceId(v string) *ModifyInstanceNetworkSpecRequest
	GetInstanceId() *string
	SetInternetMaxBandwidthIn(v int32) *ModifyInstanceNetworkSpecRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *ModifyInstanceNetworkSpecRequest
	GetInternetMaxBandwidthOut() *int32
	SetNetworkChargeType(v string) *ModifyInstanceNetworkSpecRequest
	GetNetworkChargeType() *string
	SetOwnerAccount(v string) *ModifyInstanceNetworkSpecRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *ModifyInstanceNetworkSpecRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *ModifyInstanceNetworkSpecRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *ModifyInstanceNetworkSpecRequest
	GetResourceOwnerId() *int64
	SetStartTime(v string) *ModifyInstanceNetworkSpecRequest
	GetStartTime() *string
}

type ModifyInstanceNetworkSpecRequest struct {
	// Specifies whether to assign a public IP address. Valid values:
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
	AllocatePublicIp *bool `json:"AllocatePublicIp,omitempty" xml:"AllocatePublicIp,omitempty"`
	// Specifies whether to automatically complete the payment. Valid values:
	//
	// 	- true: After you modify the bandwidth configurations, the payment is automatically completed. Make sure that your account balance is sufficient before you set AutoPay to true. If your account balance is insufficient, your order cannot be paid in the ECS console and becomes invalid. You must cancel the order.
	//
	// 	- false: After you modify the bandwidth configurations, an order is generated but the payment is not automatically completed. If your account balance is insufficient, you can set AutoPay to false to generate an unpaid order. Then, you can log on to the [ECS console](https://ecs.console.aliyun.com) to pay for the order.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. **The token can contain only ASCII characters and cannot exceed 64 characters in length.*	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The end time of the temporary bandwidth upgrade. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddThhZ format. The time must be in UTC and accurate to **hours*	- (hh).
	//
	// >  The interval between the end time and start time of temporary bandwidth upgrade must be greater than or equal to 3 hours.
	//
	// example:
	//
	// 2017-12-06T22Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// > This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	ISP *string `json:"ISP,omitempty" xml:"ISP,omitempty"`
	// The ID of the instance for which you want to modify network configurations.
	//
	// This parameter is required.
	//
	// example:
	//
	// i-bp67acfmxazb4****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The maximum inbound bandwidth from the Internet. Unit: Mbit/s. Valid values:
	//
	// 	- If the purchased outbound public bandwidth is less than or equal to 10 Mbit/s, the valid values of this parameter are 1 to 10 and the default value is 10.
	//
	// 	- If the purchased outbound public bandwidth is greater than 10 Mbit/s, the valid values of this parameter are 1 to the `InternetMaxBandwidthOut` value and the default value is the `InternetMaxBandwidthOut` value.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// 	- Valid values when the pay-by-traffic billing method for network usage is used: 0 to 100.
	//
	// 	- Valid values when the pay-by-bandwidth billing method for network usage is used:
	//
	//     	- Valid values for subscription instances: 0 to 200.
	//
	//     	- Valid values for pay-as-you-go instances: 0 to 100.
	//
	// >  The maximum outbound bandwidth of a single instance is also limited by the **network baseline bandwidth (Gbit/s) and network burst bandwidth (Gbit/s)*	- of the instance type. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// 	- PayByBandwidth
	//
	// 	- PayByTraffic
	//
	// > When the **pay-by-traffic*	- billing method for network usage is used, the maximum inbound and outbound bandwidth values are used as the upper limits of bandwidths instead of guaranteed values. In scenarios where demand outstrips resource supplies, these maximum bandwidths may be limited. If you want guaranteed bandwidths for your instance, use the **pay-by-bandwidth*	- billing method for network usage.
	//
	// example:
	//
	// PayByTraffic
	NetworkChargeType    *string `json:"NetworkChargeType,omitempty" xml:"NetworkChargeType,omitempty"`
	OwnerAccount         *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The start time of the temporary bandwidth upgrade. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddThh:mmZ format. The time must be in UTC and accurate to **minutes (mm)**.
	//
	// example:
	//
	// 2017-12-05T22:40Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s ModifyInstanceNetworkSpecRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceNetworkSpecRequest) GoString() string {
	return s.String()
}

func (s *ModifyInstanceNetworkSpecRequest) GetAllocatePublicIp() *bool {
	return s.AllocatePublicIp
}

func (s *ModifyInstanceNetworkSpecRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *ModifyInstanceNetworkSpecRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *ModifyInstanceNetworkSpecRequest) GetEndTime() *string {
	return s.EndTime
}

func (s *ModifyInstanceNetworkSpecRequest) GetISP() *string {
	return s.ISP
}

func (s *ModifyInstanceNetworkSpecRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ModifyInstanceNetworkSpecRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *ModifyInstanceNetworkSpecRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *ModifyInstanceNetworkSpecRequest) GetNetworkChargeType() *string {
	return s.NetworkChargeType
}

func (s *ModifyInstanceNetworkSpecRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *ModifyInstanceNetworkSpecRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *ModifyInstanceNetworkSpecRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *ModifyInstanceNetworkSpecRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *ModifyInstanceNetworkSpecRequest) GetStartTime() *string {
	return s.StartTime
}

func (s *ModifyInstanceNetworkSpecRequest) SetAllocatePublicIp(v bool) *ModifyInstanceNetworkSpecRequest {
	s.AllocatePublicIp = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetAutoPay(v bool) *ModifyInstanceNetworkSpecRequest {
	s.AutoPay = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetClientToken(v string) *ModifyInstanceNetworkSpecRequest {
	s.ClientToken = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetEndTime(v string) *ModifyInstanceNetworkSpecRequest {
	s.EndTime = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetISP(v string) *ModifyInstanceNetworkSpecRequest {
	s.ISP = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetInstanceId(v string) *ModifyInstanceNetworkSpecRequest {
	s.InstanceId = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetInternetMaxBandwidthIn(v int32) *ModifyInstanceNetworkSpecRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetInternetMaxBandwidthOut(v int32) *ModifyInstanceNetworkSpecRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetNetworkChargeType(v string) *ModifyInstanceNetworkSpecRequest {
	s.NetworkChargeType = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetOwnerAccount(v string) *ModifyInstanceNetworkSpecRequest {
	s.OwnerAccount = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetOwnerId(v int64) *ModifyInstanceNetworkSpecRequest {
	s.OwnerId = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetResourceOwnerAccount(v string) *ModifyInstanceNetworkSpecRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetResourceOwnerId(v int64) *ModifyInstanceNetworkSpecRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) SetStartTime(v string) *ModifyInstanceNetworkSpecRequest {
	s.StartTime = &v
	return s
}

func (s *ModifyInstanceNetworkSpecRequest) Validate() error {
	return dara.Validate(s)
}
