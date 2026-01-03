// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHibernationOptions(v *CreateInstanceRequestHibernationOptions) *CreateInstanceRequest
	GetHibernationOptions() *CreateInstanceRequestHibernationOptions
	SetPrivatePoolOptions(v *CreateInstanceRequestPrivatePoolOptions) *CreateInstanceRequest
	GetPrivatePoolOptions() *CreateInstanceRequestPrivatePoolOptions
	SetSystemDisk(v *CreateInstanceRequestSystemDisk) *CreateInstanceRequest
	GetSystemDisk() *CreateInstanceRequestSystemDisk
	SetAffinity(v string) *CreateInstanceRequest
	GetAffinity() *string
	SetArn(v []*CreateInstanceRequestArn) *CreateInstanceRequest
	GetArn() []*CreateInstanceRequestArn
	SetAutoRenew(v bool) *CreateInstanceRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateInstanceRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *CreateInstanceRequest
	GetClientToken() *string
	SetClusterId(v string) *CreateInstanceRequest
	GetClusterId() *string
	SetCreditSpecification(v string) *CreateInstanceRequest
	GetCreditSpecification() *string
	SetDataDisk(v []*CreateInstanceRequestDataDisk) *CreateInstanceRequest
	GetDataDisk() []*CreateInstanceRequestDataDisk
	SetDedicatedHostId(v string) *CreateInstanceRequest
	GetDedicatedHostId() *string
	SetDeletionProtection(v bool) *CreateInstanceRequest
	GetDeletionProtection() *bool
	SetDeploymentSetGroupNo(v int32) *CreateInstanceRequest
	GetDeploymentSetGroupNo() *int32
	SetDeploymentSetId(v string) *CreateInstanceRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *CreateInstanceRequest
	GetDescription() *string
	SetDryRun(v bool) *CreateInstanceRequest
	GetDryRun() *bool
	SetHostName(v string) *CreateInstanceRequest
	GetHostName() *string
	SetHpcClusterId(v string) *CreateInstanceRequest
	GetHpcClusterId() *string
	SetHttpEndpoint(v string) *CreateInstanceRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *CreateInstanceRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *CreateInstanceRequest
	GetHttpTokens() *string
	SetImageFamily(v string) *CreateInstanceRequest
	GetImageFamily() *string
	SetImageId(v string) *CreateInstanceRequest
	GetImageId() *string
	SetInnerIpAddress(v string) *CreateInstanceRequest
	GetInnerIpAddress() *string
	SetInstanceChargeType(v string) *CreateInstanceRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *CreateInstanceRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateInstanceRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *CreateInstanceRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *CreateInstanceRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *CreateInstanceRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateInstanceRequest
	GetIoOptimized() *string
	SetKeyPairName(v string) *CreateInstanceRequest
	GetKeyPairName() *string
	SetNodeControllerId(v string) *CreateInstanceRequest
	GetNodeControllerId() *string
	SetOwnerAccount(v string) *CreateInstanceRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateInstanceRequest
	GetOwnerId() *int64
	SetPassword(v string) *CreateInstanceRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *CreateInstanceRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *CreateInstanceRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateInstanceRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *CreateInstanceRequest
	GetPrivateIpAddress() *string
	SetRamRoleName(v string) *CreateInstanceRequest
	GetRamRoleName() *string
	SetRegionId(v string) *CreateInstanceRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateInstanceRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateInstanceRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateInstanceRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *CreateInstanceRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateInstanceRequest
	GetSecurityGroupId() *string
	SetSpotDuration(v int32) *CreateInstanceRequest
	GetSpotDuration() *int32
	SetSpotInterruptionBehavior(v string) *CreateInstanceRequest
	GetSpotInterruptionBehavior() *string
	SetSpotPriceLimit(v float32) *CreateInstanceRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *CreateInstanceRequest
	GetSpotStrategy() *string
	SetStorageSetId(v string) *CreateInstanceRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *CreateInstanceRequest
	GetStorageSetPartitionNumber() *int32
	SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest
	GetTag() []*CreateInstanceRequestTag
	SetTenancy(v string) *CreateInstanceRequest
	GetTenancy() *string
	SetUseAdditionalService(v bool) *CreateInstanceRequest
	GetUseAdditionalService() *bool
	SetUserData(v string) *CreateInstanceRequest
	GetUserData() *string
	SetVSwitchId(v string) *CreateInstanceRequest
	GetVSwitchId() *string
	SetVlanId(v string) *CreateInstanceRequest
	GetVlanId() *string
	SetZoneId(v string) *CreateInstanceRequest
	GetZoneId() *string
}

type CreateInstanceRequest struct {
	HibernationOptions *CreateInstanceRequestHibernationOptions `json:"HibernationOptions,omitempty" xml:"HibernationOptions,omitempty" type:"Struct"`
	PrivatePoolOptions *CreateInstanceRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SystemDisk         *CreateInstanceRequestSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Specifies whether to associate the instance on a dedicated host with the dedicated host. Valid values:
	//
	// 	- default: does not associate the instance with the dedicated host. When you start an instance that was stopped in economical mode, the instance is automatically deployed to another dedicated host in the automatic deployment resource pool if the available resources of the original dedicated host are insufficient.
	//
	// 	- host: associates the instance with the dedicated host. When you start an instance that was stopped in economical mode, the instance remains on the original dedicated host. If the available resources of the original dedicated host are insufficient, the instance cannot start.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	Arn []*CreateInstanceRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to enable auto-renewal for the instance. This parameter is valid only if `InstanceChargeType` is set to `PrePaid`. Valid values:
	//
	// 	- true: enables auto-renewal.
	//
	// 	- false: does not enable auto-renewal.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period of the instance. This parameter is required if AutoRenew is set to true.
	//
	// Valid values if PeriodUnit is set to Month: 1, 2, 3, 6, and 12.
	//
	// example:
	//
	// 2
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.***	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// The ID of the cluster in which to create the instance.
	//
	// >  This parameter will be removed in the future. We recommend that you use other parameters to ensure future compatibility.
	//
	// example:
	//
	// c-bp67acfmxazb4p****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The performance mode of the burstable instance. Valid values:
	//
	// 	- Standard: standard mode. For more information, see the "Standard mode" section in the [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html) topic.
	//
	// 	- Unlimited: unlimited mode. For more information, see the "Unlimited mode" section in the [Burstable instances](https://help.aliyun.com/document_detail/59977.html) topic.
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The data disks.
	DataDisk []*CreateInstanceRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The ID of the dedicated host on which to create the instance.
	//
	// You can call the [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/134242.html) operation to query the list of dedicated host IDs.
	//
	// > Spot instances (spot instances) cannot be created on dedicated hosts. If you specify DedicatedHostId, SpotStrategy and SpotPriceLimit are automatically ignored.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// Specifies whether to enable release protection for the instance. This parameter indicates whether you can use the ECS console or call the [DeleteInstance](https://help.aliyun.com/document_detail/25507.html) operation to release the instance.
	//
	// 	- true: enables release protection.
	//
	// 	- false (default): disables release protection.
	//
	// >  This parameter is applicable only to pay-as-you-go instances. It can protect instances against manual releases, but not against automatic releases.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The number of the deployment set group to which to deploy the instance. If the deployment set specified by the DeploymentSetId parameter uses the high availability group strategy (AvailabilityGroup), you can use the DeploymentSetGroupNo parameter to specify a deployment set group in the deployment set. Valid values: 1 to 7.
	//
	// example:
	//
	// 1
	DeploymentSetGroupNo *int32 `json:"DeploymentSetGroupNo,omitempty" xml:"DeploymentSetGroupNo,omitempty"`
	// The ID of the deployment set to which to deploy the instance.
	//
	// example:
	//
	// ds-bp1brhwhoqinyjd6****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The description of the instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// InstanceTest
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to perform only a dry run, without performing the actual request. Valid values:
	//
	// 	- true: performs only a dry run. The system checks the request for potential issues, including missing parameter values, incorrect request syntax, service limits, and unavailable ECS resources. If the request fails the dry run, an error message is returned. If the request passes the dry run, the `DryRunOperation` error code is returned.
	//
	// 	- false (default): performs a dry run and performs the actual request. If the request passes the dry run, the operation is performed.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The hostname of the instance.
	//
	// 	- The hostname cannot start or end with a period (.) or hyphen (-). It cannot contain consecutive periods (.) or hyphens (-).
	//
	// 	- For a Windows instance, the hostname must be 2 to 15 characters in length and cannot contain periods (.) or contain only digits. It can contain letters, digits, and hyphens (-).
	//
	// 	- For an instance that runs another type of operating system such as Linux, the hostname must be 2 to 64 characters in length. You can use periods (.) to separate the hostname into multiple segments. Each segment can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// LocalHostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The ID of the high performance computing (HPC) cluster to which to assign the instance.
	//
	// example:
	//
	// hpc-bp67acfmxazb4p****
	HpcClusterId *string `json:"HpcClusterId,omitempty" xml:"HpcClusterId,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// Default value: enabled.
	//
	// >  For more information about instance metadata, see [Overview of instance metadata](https://help.aliyun.com/document_detail/49122.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 0
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Specifies whether to forcibly use the security hardening mode (IMDSv2) to access instance metadata. Valid values:
	//
	// 	- optional: does not forcefully use the security hardening mode (IMDSv2).
	//
	// 	- required: forcefully uses the security hardening mode (IMDSv2). After you set this parameter to required, you cannot access instance metadata in normal mode.
	//
	// Default value: optional.
	//
	// >  For more information about the modes of accessing instance metadata, see [Access mode of instance metadata](https://help.aliyun.com/document_detail/150575.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. You can set this parameter to obtain the latest available custom image from the specified image family to create the instance.
	//
	// 	- ImageFamily must be empty if `ImageId` is specified.
	//
	// 	- ImageFamily can be specified if `ImageId` is not specified.
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image to use to create the instance. To use an Alibaba Cloud Marketplace image, you can view the `image ID` on the product page of the Alibaba Cloud Marketplace image. This parameter is required if you do not specify `ImageFamily` to obtain the latest available custom image from the specified image family.
	//
	// example:
	//
	// ubuntu_18_04_64_20G_alibase_20190624.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The internal IP address to assign to the instance.
	//
	// example:
	//
	// ``192.168.**.**``
	InnerIpAddress *string `json:"InnerIpAddress,omitempty" xml:"InnerIpAddress,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription. If you set this parameter to PrePaid, make sure that you have sufficient balance or credit in your account. Otherwise, an `InvalidPayMethod` error is returned.
	//
	// 	- PostPaid (default): pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the instance. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-). If you do not specify this parameter, the instance ID is used as the instance name by default.
	//
	// example:
	//
	// 2018-12-06T103200Z
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type.
	//
	// 	- Instance type selection: See [Instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the performance data of instance types, or see [Best practices for instance type selection](https://help.aliyun.com/document_detail/58291.html) to learn about how to select instance types.
	//
	// 	- Query of available resources: Call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/66186.html) operation to query resources available in a specific region or zone.
	//
	// This parameter is required.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for network usage. Default value: PayByTraffic. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth
	//
	// 	- PayByTraffic (default): pay-by-traffic
	//
	// >  When the **pay-by-traffic*	- billing method is used for network usage, the maximum inbound and outbound bandwidths are used as the upper limits of bandwidths instead of guaranteed performance specifications. In scenarios where demand outstrips resource supplies, these maximum bandwidth values may be limited. If you want guaranteed bandwidths for your instance, use the **pay-by-bandwidth*	- billing method for network usage.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// 	- When the purchased outbound public bandwidth is less than or equal to 10 Mbit/s, the valid values of this parameter are 1 to 10 and the default value is 10.
	//
	// 	- When the purchased outbound public bandwidth is greater than 10 Mbit/s, the valid values of this parameter are 1 to the `InternetMaxBandwidthOut` value and the default value is the `InternetMaxBandwidthOut` value.
	//
	// example:
	//
	// 50
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is I/O optimized. Valid values:
	//
	// 	- none: The instance is not I/O optimized.
	//
	// 	- optimized: The ECS instance is I/O optimized.
	//
	// For retired instance types, the default value is none. For more information, see [Retired instance types](https://help.aliyun.com/document_detail/55263.html).
	//
	// For other instance types, the default value is optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The name of the key pair.
	//
	// >  For Windows instances, this parameter is ignored. This parameter is empty by default. The `Password` parameter takes effect even if the KeyPairName parameter is specified.
	//
	// example:
	//
	// KeyPairTestName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// hide
	NodeControllerId *string `json:"NodeControllerId,omitempty" xml:"NodeControllerId,omitempty"`
	OwnerAccount     *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId          *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the instance. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. The following special characters are supported:
	//
	//     ( ) ` ~ ! @ # $ % ^ & 	- - _ + = | { } [ ] : ; \\" < > , . ? /
	//
	// Take note of the following items:
	//
	// 	- For security reasons, we recommend that you use HTTPS to send requests if the Password parameter is specified.
	//
	// 	- Passwords of Windows instances cannot start with a forward slash (/).
	//
	// 	- Passwords cannot be set for instances that run specific types of operating systems such as Others Linux and Fedora CoreOS. For these instances, only key pairs can be set.
	//
	// example:
	//
	// TestEcs123!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password preset in the image. When you use this parameter, leave the Password parameter empty and make sure that the selected image has a password preset.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription period of the instance. The unit is specified by `PeriodUnit`. This parameter is valid and required only when `InstanceChargeType` is set to `PrePaid`. If `DedicatedHostId` is specified, the value of Period must not exceed the subscription period of the specified dedicated host. Valid values:
	//
	// Valid values if PeriodUnit is set to Month: 1, 2, 3, 6, and 12.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// Month
	//
	// Default value: Month.
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private IP address to assign to the instance. The private IP address must be an available IP address in the CIDR block of the specified vSwitch.
	//
	// example:
	//
	// 172.16.236.*
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The name of the instance Resource Access Management (RAM) role. You can call the [ListRoles](https://help.aliyun.com/document_detail/28713.html) operation provided by RAM to query the instance RAM roles that you created.
	//
	// example:
	//
	// RAMTestName
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the region in which to create the instance. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the instance.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable security hardening. Valid values:
	//
	// 	- Active: enables security hardening. This value applies only to public images.
	//
	// 	- Deactive: disables security hardening. This value is applicable to all images.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which to assign the instance.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// 	- 1: After a spot instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After the 1-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// 	- 0: After a spot instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// >
	//
	// 	- You can set this parameter only to 0 or 1.
	//
	// 	- The spot instance is billed by second. Specify an appropriate protection period.
	//
	// 	- Alibaba Cloud sends an ECS system event to notify you 5 minutes before the instance is released.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The interruption mode of the spot instance. Valid values:
	//
	// 	- Terminate: The instance is released.
	//
	// 	- Stop: The instance is stopped in economical mode.
	//
	//     For information about the economical mode, see [Economical mode](https://help.aliyun.com/document_detail/63353.html).
	//
	// Default value: Terminate.
	//
	// example:
	//
	// Terminate
	SpotInterruptionBehavior *string `json:"SpotInterruptionBehavior,omitempty" xml:"SpotInterruptionBehavior,omitempty"`
	// The maximum hourly price of the instance. The value is accurate to three decimal places. This parameter is valid only when `SpotStrategy` is set to `SpotWithPriceLimit`.
	//
	// example:
	//
	// 0.98
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the pay-as-you-go instance. This parameter is valid only if you set `InstanceChargeType` to `PostPaid`. Valid values:
	//
	// 	- NoSpot (default): The instance is created as a regular pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is a spot instance for which you specify the maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is a spot instance for which the market price at the time of purchase is automatically used as the bid price. The market price can be up to the pay-as-you-go price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the storage set.
	//
	// example:
	//
	// ss-bp1j4i2jdf3owlhe****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set. Valid values: greater than or equal to 2.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The tags to add to the instance.
	Tag []*CreateInstanceRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to create the instance on a dedicated host. Valid values:
	//
	// 	- default: creates the instance on a non-dedicated host.
	//
	// 	- host: creates the instance on a dedicated host. If you do not specify `DedicatedHostId`, Alibaba Cloud selects a dedicated host for the instance.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// Specifies whether to use the system configurations for virtual machines. Alibaba Cloud provides the Network Time Protocol (NTP) and Key Management Service (KMS) system configurations for Windows and the NTP and Yellowdog Updater, Modified (YUM) system configurations for Linux.
	//
	// example:
	//
	// true
	UseAdditionalService *bool `json:"UseAdditionalService,omitempty" xml:"UseAdditionalService,omitempty"`
	// The user data of the instance. The user data must be encoded in Base64. The maximum size of raw data is 32 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the vSwitch to which to connect the instance. This parameter is required when you create an instance in a VPC. You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query available vSwitches.
	//
	// >  If `VSwitchId` is specified, the zone specified by `ZoneId` must be the zone where the specified vSwitch resides. You can also leave `ZoneId` empty. Then, the system selects the zone where the specified vSwitch resides.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the virtual LAN (VLAN).
	//
	// example:
	//
	// 10
	VlanId *string `json:"VlanId,omitempty" xml:"VlanId,omitempty"`
	// The ID of the zone in which to create the instance. You can call the [DescribeZones](https://help.aliyun.com/document_detail/25610.html) operation to query the zones in a specific region.
	//
	// >  If `VSwitchId` is specified, the zone specified by `ZoneId` must be the zone where the specified vSwitch resides. You can also leave `ZoneId` empty. Then, the system selects the zone where the specified vSwitch resides.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateInstanceRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequest) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequest) GetHibernationOptions() *CreateInstanceRequestHibernationOptions {
	return s.HibernationOptions
}

func (s *CreateInstanceRequest) GetPrivatePoolOptions() *CreateInstanceRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *CreateInstanceRequest) GetSystemDisk() *CreateInstanceRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateInstanceRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *CreateInstanceRequest) GetArn() []*CreateInstanceRequestArn {
	return s.Arn
}

func (s *CreateInstanceRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateInstanceRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateInstanceRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *CreateInstanceRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *CreateInstanceRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateInstanceRequest) GetDataDisk() []*CreateInstanceRequestDataDisk {
	return s.DataDisk
}

func (s *CreateInstanceRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *CreateInstanceRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateInstanceRequest) GetDeploymentSetGroupNo() *int32 {
	return s.DeploymentSetGroupNo
}

func (s *CreateInstanceRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateInstanceRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *CreateInstanceRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateInstanceRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *CreateInstanceRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *CreateInstanceRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *CreateInstanceRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *CreateInstanceRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *CreateInstanceRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateInstanceRequest) GetInnerIpAddress() *string {
	return s.InnerIpAddress
}

func (s *CreateInstanceRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateInstanceRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateInstanceRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateInstanceRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateInstanceRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateInstanceRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateInstanceRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateInstanceRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateInstanceRequest) GetNodeControllerId() *string {
	return s.NodeControllerId
}

func (s *CreateInstanceRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateInstanceRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateInstanceRequest) GetPassword() *string {
	return s.Password
}

func (s *CreateInstanceRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateInstanceRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateInstanceRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateInstanceRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateInstanceRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateInstanceRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateInstanceRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateInstanceRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateInstanceRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateInstanceRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateInstanceRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateInstanceRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateInstanceRequest) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *CreateInstanceRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *CreateInstanceRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateInstanceRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *CreateInstanceRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *CreateInstanceRequest) GetTag() []*CreateInstanceRequestTag {
	return s.Tag
}

func (s *CreateInstanceRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *CreateInstanceRequest) GetUseAdditionalService() *bool {
	return s.UseAdditionalService
}

func (s *CreateInstanceRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateInstanceRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateInstanceRequest) GetVlanId() *string {
	return s.VlanId
}

func (s *CreateInstanceRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateInstanceRequest) SetHibernationOptions(v *CreateInstanceRequestHibernationOptions) *CreateInstanceRequest {
	s.HibernationOptions = v
	return s
}

func (s *CreateInstanceRequest) SetPrivatePoolOptions(v *CreateInstanceRequestPrivatePoolOptions) *CreateInstanceRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *CreateInstanceRequest) SetSystemDisk(v *CreateInstanceRequestSystemDisk) *CreateInstanceRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateInstanceRequest) SetAffinity(v string) *CreateInstanceRequest {
	s.Affinity = &v
	return s
}

func (s *CreateInstanceRequest) SetArn(v []*CreateInstanceRequestArn) *CreateInstanceRequest {
	s.Arn = v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenew(v bool) *CreateInstanceRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateInstanceRequest) SetAutoRenewPeriod(v int32) *CreateInstanceRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateInstanceRequest) SetClientToken(v string) *CreateInstanceRequest {
	s.ClientToken = &v
	return s
}

func (s *CreateInstanceRequest) SetClusterId(v string) *CreateInstanceRequest {
	s.ClusterId = &v
	return s
}

func (s *CreateInstanceRequest) SetCreditSpecification(v string) *CreateInstanceRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateInstanceRequest) SetDataDisk(v []*CreateInstanceRequestDataDisk) *CreateInstanceRequest {
	s.DataDisk = v
	return s
}

func (s *CreateInstanceRequest) SetDedicatedHostId(v string) *CreateInstanceRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *CreateInstanceRequest) SetDeletionProtection(v bool) *CreateInstanceRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateInstanceRequest) SetDeploymentSetGroupNo(v int32) *CreateInstanceRequest {
	s.DeploymentSetGroupNo = &v
	return s
}

func (s *CreateInstanceRequest) SetDeploymentSetId(v string) *CreateInstanceRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateInstanceRequest) SetDescription(v string) *CreateInstanceRequest {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequest) SetDryRun(v bool) *CreateInstanceRequest {
	s.DryRun = &v
	return s
}

func (s *CreateInstanceRequest) SetHostName(v string) *CreateInstanceRequest {
	s.HostName = &v
	return s
}

func (s *CreateInstanceRequest) SetHpcClusterId(v string) *CreateInstanceRequest {
	s.HpcClusterId = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpEndpoint(v string) *CreateInstanceRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpPutResponseHopLimit(v int32) *CreateInstanceRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *CreateInstanceRequest) SetHttpTokens(v string) *CreateInstanceRequest {
	s.HttpTokens = &v
	return s
}

func (s *CreateInstanceRequest) SetImageFamily(v string) *CreateInstanceRequest {
	s.ImageFamily = &v
	return s
}

func (s *CreateInstanceRequest) SetImageId(v string) *CreateInstanceRequest {
	s.ImageId = &v
	return s
}

func (s *CreateInstanceRequest) SetInnerIpAddress(v string) *CreateInstanceRequest {
	s.InnerIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceChargeType(v string) *CreateInstanceRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceName(v string) *CreateInstanceRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateInstanceRequest) SetInstanceType(v string) *CreateInstanceRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetChargeType(v string) *CreateInstanceRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetMaxBandwidthIn(v int32) *CreateInstanceRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateInstanceRequest) SetInternetMaxBandwidthOut(v int32) *CreateInstanceRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateInstanceRequest) SetIoOptimized(v string) *CreateInstanceRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateInstanceRequest) SetKeyPairName(v string) *CreateInstanceRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateInstanceRequest) SetNodeControllerId(v string) *CreateInstanceRequest {
	s.NodeControllerId = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerAccount(v string) *CreateInstanceRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetOwnerId(v int64) *CreateInstanceRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetPassword(v string) *CreateInstanceRequest {
	s.Password = &v
	return s
}

func (s *CreateInstanceRequest) SetPasswordInherit(v bool) *CreateInstanceRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriod(v int32) *CreateInstanceRequest {
	s.Period = &v
	return s
}

func (s *CreateInstanceRequest) SetPeriodUnit(v string) *CreateInstanceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateInstanceRequest) SetPrivateIpAddress(v string) *CreateInstanceRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateInstanceRequest) SetRamRoleName(v string) *CreateInstanceRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateInstanceRequest) SetRegionId(v string) *CreateInstanceRequest {
	s.RegionId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceGroupId(v string) *CreateInstanceRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerAccount(v string) *CreateInstanceRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateInstanceRequest) SetResourceOwnerId(v int64) *CreateInstanceRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateInstanceRequest) SetSecurityEnhancementStrategy(v string) *CreateInstanceRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateInstanceRequest) SetSecurityGroupId(v string) *CreateInstanceRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotDuration(v int32) *CreateInstanceRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotInterruptionBehavior(v string) *CreateInstanceRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotPriceLimit(v float32) *CreateInstanceRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateInstanceRequest) SetSpotStrategy(v string) *CreateInstanceRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSetId(v string) *CreateInstanceRequest {
	s.StorageSetId = &v
	return s
}

func (s *CreateInstanceRequest) SetStorageSetPartitionNumber(v int32) *CreateInstanceRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *CreateInstanceRequest) SetTag(v []*CreateInstanceRequestTag) *CreateInstanceRequest {
	s.Tag = v
	return s
}

func (s *CreateInstanceRequest) SetTenancy(v string) *CreateInstanceRequest {
	s.Tenancy = &v
	return s
}

func (s *CreateInstanceRequest) SetUseAdditionalService(v bool) *CreateInstanceRequest {
	s.UseAdditionalService = &v
	return s
}

func (s *CreateInstanceRequest) SetUserData(v string) *CreateInstanceRequest {
	s.UserData = &v
	return s
}

func (s *CreateInstanceRequest) SetVSwitchId(v string) *CreateInstanceRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateInstanceRequest) SetVlanId(v string) *CreateInstanceRequest {
	s.VlanId = &v
	return s
}

func (s *CreateInstanceRequest) SetZoneId(v string) *CreateInstanceRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateInstanceRequest) Validate() error {
	if s.HibernationOptions != nil {
		if err := s.HibernationOptions.Validate(); err != nil {
			return err
		}
	}
	if s.PrivatePoolOptions != nil {
		if err := s.PrivatePoolOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
		}
	}
	if s.Arn != nil {
		for _, item := range s.Arn {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.DataDisk != nil {
		for _, item := range s.DataDisk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
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

type CreateInstanceRequestHibernationOptions struct {
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// false
	Configured *bool `json:"Configured,omitempty" xml:"Configured,omitempty"`
}

func (s CreateInstanceRequestHibernationOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestHibernationOptions) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestHibernationOptions) GetConfigured() *bool {
	return s.Configured
}

func (s *CreateInstanceRequestHibernationOptions) SetConfigured(v bool) *CreateInstanceRequestHibernationOptions {
	s.Configured = &v
	return s
}

func (s *CreateInstanceRequestHibernationOptions) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestPrivatePoolOptions struct {
	// The ID of the private pool. The ID of a private pool is the same as that of the elasticity assurance or capacity reservation for which the private pool is generated.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the private pool to use to start the instance. A private pool is generated when an elasticity assurance or a capacity reservation takes effect. You can select a private pool to start instances. Valid values:
	//
	// 	- Open: open private pool. The system selects a matching open private pool to start the instance. If no matching open private pools are found, resources in the public pool are used. When you set this parameter to Open, you can leave the `PrivatePoolOptions.Id` parameter empty.
	//
	// 	- Target: specified private pool. The system uses the capacity in a specified private pool to start the instance. If the specified private pool is unavailable, the instance cannot be started. If you set this parameter to Target, you must specify the `PrivatePoolOptions.Id` parameter.
	//
	// 	- None: no private pool. The capacity in private pools is not used.
	//
	// Default value: none.
	//
	// In the following scenarios, the PrivatePoolOptions.MatchCriteria parameter can be set only to `None` or left empty:
	//
	// 	- Create a spot instance.
	//
	// 	- Create an instance in the classic network.
	//
	// 	- Create an instance on a dedicated host.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s CreateInstanceRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *CreateInstanceRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *CreateInstanceRequestPrivatePoolOptions) SetId(v string) *CreateInstanceRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *CreateInstanceRequestPrivatePoolOptions) SetMatchCriteria(v string) *CreateInstanceRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *CreateInstanceRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestSystemDisk struct {
	// The category of the system disk. Valid values:
	//
	// 	- cloud_essd: ESSD. If SystemDisk.Category is set to this value, you can use `SystemDisk.PerformanceLevel` to specify the performance level of the disk.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud: basic disk.
	//
	// For non-I/O optimized instances of retired instance types, the default value is cloud. For other types of instances, the default value is cloud_efficiency.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length. It must start with a letter and cannot start with `http://` or `https://`. It can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// SystemDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The performance level of the ESSD that is used as the system disk. Default value: PL1. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1 (default): A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// For information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values:
	//
	// 	- Basic disks: 20 to 500.
	//
	// 	- Other disks: 20 to 2048.
	//
	// The value of this parameter must be at least 20 and greater than or equal to the size of the image.
	//
	// Default value: 40 or the size of the image, whichever is greater.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the dedicated block storage cluster. If you want to use disks in a dedicated block storage cluster as system disks when you create instances, you need to specify this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s CreateInstanceRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateInstanceRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateInstanceRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateInstanceRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateInstanceRequestSystemDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *CreateInstanceRequestSystemDisk) SetCategory(v string) *CreateInstanceRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetDescription(v string) *CreateInstanceRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetDiskName(v string) *CreateInstanceRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetPerformanceLevel(v string) *CreateInstanceRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetSize(v int32) *CreateInstanceRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) SetStorageClusterId(v string) *CreateInstanceRequestSystemDisk {
	s.StorageClusterId = &v
	return s
}

func (s *CreateInstanceRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestArn struct {
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// 1234567890
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// Primary
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// >  This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// acs:ram::123456789012****:role/adminrole
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s CreateInstanceRequestArn) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestArn) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *CreateInstanceRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *CreateInstanceRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *CreateInstanceRequestArn) SetAssumeRoleFor(v int64) *CreateInstanceRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *CreateInstanceRequestArn) SetRoleType(v string) *CreateInstanceRequestArn {
	s.RoleType = &v
	return s
}

func (s *CreateInstanceRequestArn) SetRolearn(v string) *CreateInstanceRequestArn {
	s.Rolearn = &v
	return s
}

func (s *CreateInstanceRequestArn) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestDataDisk struct {
	// The category of data disk N. Valid values:
	//
	// 	- cloud_efficiency: utra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud_essd: ESSD.
	//
	// 	- cloud: basic disk.
	//
	// 	- cloud_auto: ESSD AutoPL disk.
	//
	// 	- cloud_essd_entry: ESSD Entry disk.
	//
	//     **
	//
	//     **Note*	- This parameter can be set to `cloud_essd_entry` only when `InstanceType` is set to `ecs.u1` or `ecs.e`.
	//
	// 	- elastic_ephemeral_disk_standard: standard elastic ephemeral disk.
	//
	// 	- elastic_ephemeral_disk_premium: premium elastic ephemeral disk.
	//
	// For I/O optimized instances, the default value is cloud_efficiency. For non-I/O optimized instances, the default value is cloud.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release data disk N when the instance is released. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of data disk N. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// TestDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of data disk N.
	//
	// >  This parameter is applicable to scenarios in which a full image is used to create instances. A full image is an image that contains an operating system, application software, and business data. For these scenarios, you can set this parameter to the mount point of data disk N contained in the full image and modify the `DataDisk.N.Size` and `DataDisk.N.Category` parameters to change the category and size of data disk N created based on the image.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of data disk N. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// DataDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// hide
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt data disk N. Valid values:
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
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key to use for data disk N.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the ESSD to use as data disk N. The value of N must be the same as that in `DataDisk.N.Category` when DataDisk.N.Category is set to cloud_essd. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1 (default): A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// For more information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL2
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The size of data disk N. Valid values of N: 1 to 16. Unit: GiB. Valid values:
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud_efficiency: 20 to 32768.
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud_ssd: 20 to 32768.
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud_essd: vary based on the `DataDisk.N.PerformanceLevel` value.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL0: 1 to 65536.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL1: 20 to 65536.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL2: 461 to 65536.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL3: 1261 to 65536.
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud: 5 to 2000.
	//
	// >  The value of this parameter must be greater than or equal to the size of the snapshot specified by `SnapshotId`.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot to use to create data disk N. Valid values of N: 1 to 16.
	//
	// 	- If `DataDisk.N.SnapshotId` is specified, `DataDisk.N.Size` is ignored. The data disk is created based on the size of the specified snapshot.
	//
	// 	- Use snapshots created on or after July 15, 2013. Otherwise, an error is returned and your request is rejected.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The ID of the dedicated block storage cluster to which data disk N belongs. If you want to use a disk in a dedicated block storage cluster as data disk N when you create the instance, specify this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s CreateInstanceRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateInstanceRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateInstanceRequestDataDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateInstanceRequestDataDisk) GetDevice() *string {
	return s.Device
}

func (s *CreateInstanceRequestDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateInstanceRequestDataDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *CreateInstanceRequestDataDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *CreateInstanceRequestDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateInstanceRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateInstanceRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateInstanceRequestDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateInstanceRequestDataDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *CreateInstanceRequestDataDisk) SetCategory(v string) *CreateInstanceRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDeleteWithInstance(v bool) *CreateInstanceRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDescription(v string) *CreateInstanceRequestDataDisk {
	s.Description = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDevice(v string) *CreateInstanceRequestDataDisk {
	s.Device = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetDiskName(v string) *CreateInstanceRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetEncryptAlgorithm(v string) *CreateInstanceRequestDataDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetEncrypted(v bool) *CreateInstanceRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetKMSKeyId(v string) *CreateInstanceRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetPerformanceLevel(v string) *CreateInstanceRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetSize(v int32) *CreateInstanceRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetSnapshotId(v string) *CreateInstanceRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) SetStorageClusterId(v string) *CreateInstanceRequestDataDisk {
	s.StorageClusterId = &v
	return s
}

func (s *CreateInstanceRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceRequestTag struct {
	// The key of tag N to add to the instance, disks, and primary ENI. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot start with `acs:` or `aliyun`. It cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the instance, disks, and primary ENI. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateInstanceRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceRequestTag) GoString() string {
	return s.String()
}

func (s *CreateInstanceRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateInstanceRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateInstanceRequestTag) SetKey(v string) *CreateInstanceRequestTag {
	s.Key = &v
	return s
}

func (s *CreateInstanceRequestTag) SetValue(v string) *CreateInstanceRequestTag {
	s.Value = &v
	return s
}

func (s *CreateInstanceRequestTag) Validate() error {
	return dara.Validate(s)
}
