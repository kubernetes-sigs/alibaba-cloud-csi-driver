// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateActivationRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDescription(v string) *CreateActivationRequest
	GetDescription() *string
	SetInstanceCount(v int32) *CreateActivationRequest
	GetInstanceCount() *int32
	SetInstanceName(v string) *CreateActivationRequest
	GetInstanceName() *string
	SetIpAddressRange(v string) *CreateActivationRequest
	GetIpAddressRange() *string
	SetOwnerAccount(v string) *CreateActivationRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateActivationRequest
	GetOwnerId() *int64
	SetRegionId(v string) *CreateActivationRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateActivationRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateActivationRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateActivationRequest
	GetResourceOwnerId() *int64
	SetTag(v []*CreateActivationRequestTag) *CreateActivationRequest
	GetTag() []*CreateActivationRequestTag
	SetTimeToLiveInHours(v int64) *CreateActivationRequest
	GetTimeToLiveInHours() *int64
}

type CreateActivationRequest struct {
	// The description of the activation code. The description must be 1 to 100 characters in length.
	//
	// example:
	//
	// This is description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The maximum number of times that you can use the activation code to register managed instances. Valid values: 1 to 1000.
	//
	// Default value: 10.
	//
	// example:
	//
	// 10
	InstanceCount *int32 `json:"InstanceCount,omitempty" xml:"InstanceCount,omitempty"`
	// The default instance name prefix. The prefix must be 2 to 50 characters in length and can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:). The prefix must start with a letter and cannot start with a digit, a special character, `http://`, or `https://`.
	//
	// If you use the activation code that is created by calling the CreateActivation operation to register managed instances, the instances are assigned sequential names that include the value of this parameter as a prefix. You can also specify a new instance name to replace the assigned sequential name when you register a managed instance.
	//
	// If you specify InstanceName when you register a managed instance, an instance name in the `<InstanceName>-<Number>` format is generated. The number of digits in the \\<Number> value varies based on the number of digits in the `InstanceCount` value. Example: `001`. If you do not specify InstanceName, the hostname (Hostname) is used as the instance name.
	//
	// example:
	//
	// test-InstanceName
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The IP addresses of hosts that can use the activation code. The value can be IPv4 addresses, IPv6 addresses, or CIDR blocks.
	//
	// example:
	//
	// 0.0.0.0/0
	IpAddressRange *string `json:"IpAddressRange,omitempty" xml:"IpAddressRange,omitempty"`
	OwnerAccount   *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The region ID. Supported regions: China (Qingdao), China (Beijing), China (Zhangjiakou), China (Hohhot), China (Ulanqab), China (Hangzhou), China (Shanghai), China (Shenzhen), China (Heyuan), China (Guangzhou), China (Chengdu), China (Hong Kong), Singapore, Japan (Tokyo), US (Silicon Valley), and US (Virginia). You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the activation code.
	//
	// example:
	//
	// rg-123******
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// The tags to add to the activation code.
	Tag []*CreateActivationRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The validity period of the activation code. After the validity period ends, you can no longer use the activation code to register managed instances. Unit: hours.
	//
	// Default value: 4.
	//
	// example:
	//
	// 4
	TimeToLiveInHours *int64 `json:"TimeToLiveInHours,omitempty" xml:"TimeToLiveInHours,omitempty"`
}

func (s CreateActivationRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateActivationRequest) GoString() string {
	return s.String()
}

func (s *CreateActivationRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateActivationRequest) GetInstanceCount() *int32 {
	return s.InstanceCount
}

func (s *CreateActivationRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateActivationRequest) GetIpAddressRange() *string {
	return s.IpAddressRange
}

func (s *CreateActivationRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateActivationRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateActivationRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateActivationRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateActivationRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateActivationRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateActivationRequest) GetTag() []*CreateActivationRequestTag {
	return s.Tag
}

func (s *CreateActivationRequest) GetTimeToLiveInHours() *int64 {
	return s.TimeToLiveInHours
}

func (s *CreateActivationRequest) SetDescription(v string) *CreateActivationRequest {
	s.Description = &v
	return s
}

func (s *CreateActivationRequest) SetInstanceCount(v int32) *CreateActivationRequest {
	s.InstanceCount = &v
	return s
}

func (s *CreateActivationRequest) SetInstanceName(v string) *CreateActivationRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateActivationRequest) SetIpAddressRange(v string) *CreateActivationRequest {
	s.IpAddressRange = &v
	return s
}

func (s *CreateActivationRequest) SetOwnerAccount(v string) *CreateActivationRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateActivationRequest) SetOwnerId(v int64) *CreateActivationRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateActivationRequest) SetRegionId(v string) *CreateActivationRequest {
	s.RegionId = &v
	return s
}

func (s *CreateActivationRequest) SetResourceGroupId(v string) *CreateActivationRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateActivationRequest) SetResourceOwnerAccount(v string) *CreateActivationRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateActivationRequest) SetResourceOwnerId(v int64) *CreateActivationRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateActivationRequest) SetTag(v []*CreateActivationRequestTag) *CreateActivationRequest {
	s.Tag = v
	return s
}

func (s *CreateActivationRequest) SetTimeToLiveInHours(v int64) *CreateActivationRequest {
	s.TimeToLiveInHours = &v
	return s
}

func (s *CreateActivationRequest) Validate() error {
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

type CreateActivationRequestTag struct {
	// The key of tag N to add to the activation code. Valid values of N: 1 to 20. The tag key cannot be an empty string.
	//
	// If a single tag is specified to query resources, up to 1,000 resources that have this tag added can be displayed in the response. If multiple tags are specified to query resources, up to 1,000 resources that have all these tags added can be displayed in the response. To query more than 1,000 resources that have specified tags, call [ListTagResources](https://help.aliyun.com/document_detail/110425.html).
	//
	// The tag key can be up to 64 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the activation code. Valid values of N: 1 to 20. The tag value can be an empty string.
	//
	// It can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateActivationRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateActivationRequestTag) GoString() string {
	return s.String()
}

func (s *CreateActivationRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateActivationRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateActivationRequestTag) SetKey(v string) *CreateActivationRequestTag {
	s.Key = &v
	return s
}

func (s *CreateActivationRequestTag) SetValue(v string) *CreateActivationRequestTag {
	s.Value = &v
	return s
}

func (s *CreateActivationRequestTag) Validate() error {
	return dara.Validate(s)
}
