// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaunchTemplateRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSystemDisk(v *CreateLaunchTemplateRequestSystemDisk) *CreateLaunchTemplateRequest
	GetSystemDisk() *CreateLaunchTemplateRequestSystemDisk
	SetAutoReleaseTime(v string) *CreateLaunchTemplateRequest
	GetAutoReleaseTime() *string
	SetAutoRenew(v bool) *CreateLaunchTemplateRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateLaunchTemplateRequest
	GetAutoRenewPeriod() *int32
	SetCreditSpecification(v string) *CreateLaunchTemplateRequest
	GetCreditSpecification() *string
	SetDataDisk(v []*CreateLaunchTemplateRequestDataDisk) *CreateLaunchTemplateRequest
	GetDataDisk() []*CreateLaunchTemplateRequestDataDisk
	SetDeletionProtection(v bool) *CreateLaunchTemplateRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *CreateLaunchTemplateRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *CreateLaunchTemplateRequest
	GetDescription() *string
	SetEnableVmOsConfig(v bool) *CreateLaunchTemplateRequest
	GetEnableVmOsConfig() *bool
	SetHostName(v string) *CreateLaunchTemplateRequest
	GetHostName() *string
	SetHttpEndpoint(v string) *CreateLaunchTemplateRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *CreateLaunchTemplateRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *CreateLaunchTemplateRequest
	GetHttpTokens() *string
	SetImageId(v string) *CreateLaunchTemplateRequest
	GetImageId() *string
	SetImageOptions(v *CreateLaunchTemplateRequestImageOptions) *CreateLaunchTemplateRequest
	GetImageOptions() *CreateLaunchTemplateRequestImageOptions
	SetImageOwnerAlias(v string) *CreateLaunchTemplateRequest
	GetImageOwnerAlias() *string
	SetInstanceChargeType(v string) *CreateLaunchTemplateRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *CreateLaunchTemplateRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateLaunchTemplateRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *CreateLaunchTemplateRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *CreateLaunchTemplateRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *CreateLaunchTemplateRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateLaunchTemplateRequest
	GetIoOptimized() *string
	SetIpv6AddressCount(v int32) *CreateLaunchTemplateRequest
	GetIpv6AddressCount() *int32
	SetKeyPairName(v string) *CreateLaunchTemplateRequest
	GetKeyPairName() *string
	SetLaunchTemplateName(v string) *CreateLaunchTemplateRequest
	GetLaunchTemplateName() *string
	SetNetworkInterface(v []*CreateLaunchTemplateRequestNetworkInterface) *CreateLaunchTemplateRequest
	GetNetworkInterface() []*CreateLaunchTemplateRequestNetworkInterface
	SetNetworkType(v string) *CreateLaunchTemplateRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateLaunchTemplateRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLaunchTemplateRequest
	GetOwnerId() *int64
	SetPasswordInherit(v bool) *CreateLaunchTemplateRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *CreateLaunchTemplateRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateLaunchTemplateRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *CreateLaunchTemplateRequest
	GetPrivateIpAddress() *string
	SetRamRoleName(v string) *CreateLaunchTemplateRequest
	GetRamRoleName() *string
	SetRegionId(v string) *CreateLaunchTemplateRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateLaunchTemplateRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateLaunchTemplateRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLaunchTemplateRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *CreateLaunchTemplateRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateLaunchTemplateRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *CreateLaunchTemplateRequest
	GetSecurityGroupIds() []*string
	SetSecurityOptions(v *CreateLaunchTemplateRequestSecurityOptions) *CreateLaunchTemplateRequest
	GetSecurityOptions() *CreateLaunchTemplateRequestSecurityOptions
	SetSpotDuration(v int32) *CreateLaunchTemplateRequest
	GetSpotDuration() *int32
	SetSpotPriceLimit(v float32) *CreateLaunchTemplateRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *CreateLaunchTemplateRequest
	GetSpotStrategy() *string
	SetTag(v []*CreateLaunchTemplateRequestTag) *CreateLaunchTemplateRequest
	GetTag() []*CreateLaunchTemplateRequestTag
	SetTemplateResourceGroupId(v string) *CreateLaunchTemplateRequest
	GetTemplateResourceGroupId() *string
	SetTemplateTag(v []*CreateLaunchTemplateRequestTemplateTag) *CreateLaunchTemplateRequest
	GetTemplateTag() []*CreateLaunchTemplateRequestTemplateTag
	SetUserData(v string) *CreateLaunchTemplateRequest
	GetUserData() *string
	SetVSwitchId(v string) *CreateLaunchTemplateRequest
	GetVSwitchId() *string
	SetVersionDescription(v string) *CreateLaunchTemplateRequest
	GetVersionDescription() *string
	SetVpcId(v string) *CreateLaunchTemplateRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateLaunchTemplateRequest
	GetZoneId() *string
}

type CreateLaunchTemplateRequest struct {
	SystemDisk *CreateLaunchTemplateRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// The automatic release time of the instance. Specify the time in the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time must be in UTC.
	//
	// 	- If the value of `ss` is not `00`, the time is automatically rounded down to the nearest minute based on the value of `mm`.
	//
	// 	- The specified time must be at least 30 minutes later than the current time.
	//
	// 	- The specified time can be at most three years later than the current time.
	//
	// example:
	//
	// 2018-01-01T12:05:00Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// Specifies whether to enable auto-renewal. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  This parameter takes effect only if you set `InstanceChargeType` to `PrePaid`.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period of the instance. Valid values:
	//
	// Valid values when PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The performance mode of the burstable instance. Valid values:
	//
	// 	- Standard: the standard mode. For more information, see the "Standard mode" section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// 	- Unlimited: the unlimited mode. For more information, see the "Unlimited mode" section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The data disks.
	DataDisk []*CreateLaunchTemplateRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// Specifies whether to enable release protection for the instance. This parameter specifies whether you can use the ECS console or call the [DeleteInstance](https://help.aliyun.com/document_detail/25507.html) operation to release the instance. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  This parameter is applicable only to pay-as-you-go instances. The release protection feature can protect instances against manual releases, but not against automatic releases.
	//
	// example:
	//
	// false
	DeletionProtection *bool `json:"DeletionProtection,omitempty" xml:"DeletionProtection,omitempty"`
	// The ID of the deployment set to which to deploy the instance.
	//
	// example:
	//
	// ds-bp1brhwhoqinyjd6****
	DeploymentSetId *string `json:"DeploymentSetId,omitempty" xml:"DeploymentSetId,omitempty"`
	// The instance description. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testECSDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the operating system configuration of the instance.
	//
	// > This parameter will be removed in the future. To ensure future compatibility, we recommend that you use other parameters.
	//
	// example:
	//
	// false
	EnableVmOsConfig *bool `json:"EnableVmOsConfig,omitempty" xml:"EnableVmOsConfig,omitempty"`
	// The instance hostname.
	//
	// 	- The hostname cannot start or end with a period (.) or hyphen (-). It cannot contain consecutive periods (.) or hyphens (-).
	//
	// 	- For Windows instances, the hostname must be 2 to 15 characters in length and cannot contain periods (.) or contain only digits. It can contain letters, digits, and hyphens (-).
	//
	// 	- For instances that run other operating systems such as Linux, the hostname must be 2 to 64 characters in length. You can use periods (.) to separate the hostname into multiple segments. Each segment can contain letters, digits, and hyphens (-).
	//
	// example:
	//
	// testHostName
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// Specifies whether to enable the access channel for instance metadata. Valid values:
	//
	// 	- enabled
	//
	// 	- disabled
	//
	// Default value: enabled.
	//
	// >  For information about instance metadata, see [Obtain information about an ECS instance, such as instance attributes inside ECS instances from instance metadata service](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// enabled
	HttpEndpoint *string `json:"HttpEndpoint,omitempty" xml:"HttpEndpoint,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 3
	HttpPutResponseHopLimit *int32 `json:"HttpPutResponseHopLimit,omitempty" xml:"HttpPutResponseHopLimit,omitempty"`
	// Specifies whether to forcefully use the security hardening mode (IMDSv2) to access instance metadata. Valid values:
	//
	// 	- optional: does not forcefully use the security hardening mode (IMDSv2).
	//
	// 	- required: forcefully uses the security hardening mode (IMDSv2). After you set this parameter to required, you cannot access instance metadata in normal mode.
	//
	// Default value: optional.
	//
	// >  For information about the modes of accessing instance metadata, see [Obtain information about an ECS instance, such as instance attributes inside ECS instances from instance metadata service](https://help.aliyun.com/document_detail/108460.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The ID of the image to use to create the instance. You can call the [DescribeImages](https://help.aliyun.com/document_detail/25534.html) operation to query available images.
	//
	// example:
	//
	// win2008r2_64_ent_sp1_en-us_40G_alibase_20170915.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Details about the image options.
	ImageOptions *CreateLaunchTemplateRequestImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The source of the image. Valid values:
	//
	// 	- system: public image provided by Alibaba Cloud.
	//
	// 	- self: custom image that you created.
	//
	// 	- others: shared image from another Alibaba Cloud account.
	//
	// 	- marketplace:[Alibaba Cloud Marketplace](https://marketplace.alibabacloud.com/) image. If Alibaba Cloud Marketplace images are available, you can use the images without the need to subscribe to the images. Take note of the billing details of Alibaba Cloud Marketplace images.
	//
	// example:
	//
	// system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription. If you set this parameter to PrePaid, make sure that your account has sufficient credits.Otherwise, an `InvalidPayMethod` error is returned.
	//
	// 	- PostPaid: pay-as-you-go.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The instance name. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-). The default value of this parameter is the `InstanceId` value.
	//
	// When you create multiple ECS instances at a time, you can batch configure sequential names for the instances. The instance names can contain square brackets ([]) and commas (,). For more information, see [Batch configure sequential names or hostnames for multiple instances](https://help.aliyun.com/document_detail/196048.html).
	//
	// example:
	//
	// k8s-node-[1,4]-alibabacloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type of the instance. For more information, see [Instance families](https://help.aliyun.com/document_detail/25378.html). Alternatively, you can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the most recent instance type list.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth
	//
	// 	- PayByTraffic: pay-by-traffic
	//
	// > When the **pay-by-traffic*	- billing method for network usage is used, the maximum inbound and outbound bandwidths are used as the upper limits of bandwidths instead of guaranteed performance specifications. In scenarios where demand outstrips resource supplies, these maximum bandwidths may be limited. If you want guaranteed bandwidths for your instance, use the **pay-by-bandwidth*	- billing method for network usage.
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
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether to create an I/O optimized instance. Valid values:
	//
	// 	- none: creates a non-I/O optimized instance.
	//
	// 	- optimized: creates an I/O optimized instance.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// The number of IPv6 addresses to randomly generate for the primary elastic network interface (ENI). Valid values: 1 to 10.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The name of the key pair. This parameter is empty by default.
	//
	// 	- For Windows instances, this parameter is ignored The `Password` parameter takes effect even if the KeyPairName parameter is specified.
	//
	// 	- For Linux instances, the password-based logon method is disabled by default.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The name of the launch template. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// This parameter is required.
	//
	// example:
	//
	// testLaunchTemplateName
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// The information of the ENIs.
	NetworkInterface []*CreateLaunchTemplateRequestNetworkInterface `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty" type:"Repeated"`
	// The network type of the instance. Valid values:
	//
	// 	- classic: classic network
	//
	// 	- vpc: VPC
	//
	// example:
	//
	// vpc
	NetworkType  *string `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	OwnerAccount *string `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId      *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// Specifies whether to use the preset password of the image.
	//
	// > If you set the PasswordInherit parameter to true, make sure that you leave the Password parameter empty and the selected image has a preset password.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription duration of the instance. Unit: months. This parameter is valid and required only when `InstanceChargeType` is set to `PrePaid`. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Valid values:
	//
	// Month (default)
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private IP address to assign to the instance.
	//
	// To assign a private IP address to an instance that resides in a VPC, make sure that the IP address is an idle IP address within the CIDR block of the vSwitch specified by the `VSwitchId` parameter.
	//
	// example:
	//
	// ``10.1.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The name of the instance Resource Access Management (RAM) role. You can call the [ListRoles](https://help.aliyun.com/document_detail/28713.html) operation provided by RAM to query the instance RAM roles that you created.
	//
	// example:
	//
	// testRamRoleName
	RamRoleName *string `json:"RamRoleName,omitempty" xml:"RamRoleName,omitempty"`
	// The ID of the region in which to create the launch template. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent region list.
	//
	// This parameter is required.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which to assign the instance, Elastic Block Storage (EBS) device, and elastic network interface (ENI).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId      *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// Specifies whether to enable security hardening for the operating system. Valid values:
	//
	// 	- Active: enables security hardening. This value is applicable only to public images.
	//
	// 	- Deactive: does not enable security hardening. This value is applicable to all images.
	//
	// example:
	//
	// Deactive
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which to assign the instance. Instances in the same security group can communicate with each other. A security group can contain up to 1,000 instances.
	//
	// > You cannot specify both the `SecurityGroupId` and `SecurityGroupIds.N` parameters.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of the security groups to which to assign the instance. The valid values of N are based on the maximum number of security groups to which the instance can belong. For more information, see the "Security group limits" section in [Limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// > You cannot specify both the `SecurityGroupId` and `SecurityGroupIds.N` parameters.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupIds []*string                                   `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SecurityOptions  *CreateLaunchTemplateRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	// The protection period of the spot instance. Unit: hours. Default value: 1. Valid values:
	//
	// 	- 1: After a spot instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After the 1-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// 	- 0: After a spot instance is created, Alibaba Cloud does not ensure that the instance runs for 1 hour. The system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// Alibaba Cloud sends an ECS system event to notify you 5 minutes before the instance is released. Spot instances are billed by second. We recommend that you specify a protection period based on your business requirements.
	//
	// >  This parameter takes effect only if SpotStrategy is set to SpotWithPriceLimit or SpotAsPriceGo.
	//
	// example:
	//
	// 1
	SpotDuration *int32 `json:"SpotDuration,omitempty" xml:"SpotDuration,omitempty"`
	// The maximum hourly price of the instance. The value is accurate to three decimal places. This parameter is valid only when the `SpotStrategy` parameter is set to `SpotWithPriceLimit`.
	//
	// example:
	//
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the pay-as-you-go instance. This parameter is valid only when the `InstanceChargeType` parameter is set to `PostPaid`. Valid values:
	//
	// 	- NoSpot: The instance is a regular pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is created as a spot instance with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is created as a spot instance for which the market price at the time of purchase is automatically used as the bidding price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The tags to add to the instance, disks, and primary ENI that are created from the launch template.
	//
	// **Scenario**
	//
	// If you created a launch template by calling the CreateLaunchTemplate operation and use the default version that is automatically generated for the launch template to create instances, the specified tags are automatically added to the created instances, disks, and primary ENIs. For more information about the default versions of launch templates, see [xxxx]\\(url).
	Tag []*CreateLaunchTemplateRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The ID of the resource group to which the launch template belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	TemplateResourceGroupId *string `json:"TemplateResourceGroupId,omitempty" xml:"TemplateResourceGroupId,omitempty"`
	// The tags to add to the launch template.
	//
	// >  You can add tags to or query the tags of launch templates by calling API operations. You cannot add tags to or query the tags of launch templates in the ECS console.
	TemplateTag []*CreateLaunchTemplateRequestTemplateTag `json:"TemplateTag,omitempty" xml:"TemplateTag,omitempty" type:"Repeated"`
	// The user data of the instance. The user data must be encoded in Base64. The maximum size of raw data is 32 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the vSwitch to which to connect the instance. This parameter is required if you specify the VpcId parameter.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The description of the launch template version. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testVersionDescription
	VersionDescription *string `json:"VersionDescription,omitempty" xml:"VersionDescription,omitempty"`
	// The ID of the virtual private cloud (VPC).
	//
	// example:
	//
	// vpc-bp12433upq1y5scen****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone in which to create the instance.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLaunchTemplateRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequest) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequest) GetSystemDisk() *CreateLaunchTemplateRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateLaunchTemplateRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *CreateLaunchTemplateRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateLaunchTemplateRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateLaunchTemplateRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateLaunchTemplateRequest) GetDataDisk() []*CreateLaunchTemplateRequestDataDisk {
	return s.DataDisk
}

func (s *CreateLaunchTemplateRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateLaunchTemplateRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateLaunchTemplateRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateRequest) GetEnableVmOsConfig() *bool {
	return s.EnableVmOsConfig
}

func (s *CreateLaunchTemplateRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateLaunchTemplateRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *CreateLaunchTemplateRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *CreateLaunchTemplateRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *CreateLaunchTemplateRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateLaunchTemplateRequest) GetImageOptions() *CreateLaunchTemplateRequestImageOptions {
	return s.ImageOptions
}

func (s *CreateLaunchTemplateRequest) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *CreateLaunchTemplateRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateLaunchTemplateRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateLaunchTemplateRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateLaunchTemplateRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateLaunchTemplateRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateLaunchTemplateRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateLaunchTemplateRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateLaunchTemplateRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateLaunchTemplateRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateLaunchTemplateRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *CreateLaunchTemplateRequest) GetNetworkInterface() []*CreateLaunchTemplateRequestNetworkInterface {
	return s.NetworkInterface
}

func (s *CreateLaunchTemplateRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateLaunchTemplateRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLaunchTemplateRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLaunchTemplateRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateLaunchTemplateRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateLaunchTemplateRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateLaunchTemplateRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateLaunchTemplateRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateLaunchTemplateRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLaunchTemplateRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLaunchTemplateRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLaunchTemplateRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLaunchTemplateRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateLaunchTemplateRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateLaunchTemplateRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateLaunchTemplateRequest) GetSecurityOptions() *CreateLaunchTemplateRequestSecurityOptions {
	return s.SecurityOptions
}

func (s *CreateLaunchTemplateRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateLaunchTemplateRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *CreateLaunchTemplateRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateLaunchTemplateRequest) GetTag() []*CreateLaunchTemplateRequestTag {
	return s.Tag
}

func (s *CreateLaunchTemplateRequest) GetTemplateResourceGroupId() *string {
	return s.TemplateResourceGroupId
}

func (s *CreateLaunchTemplateRequest) GetTemplateTag() []*CreateLaunchTemplateRequestTemplateTag {
	return s.TemplateTag
}

func (s *CreateLaunchTemplateRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateLaunchTemplateRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLaunchTemplateRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *CreateLaunchTemplateRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateLaunchTemplateRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateLaunchTemplateRequest) SetSystemDisk(v *CreateLaunchTemplateRequestSystemDisk) *CreateLaunchTemplateRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetAutoReleaseTime(v string) *CreateLaunchTemplateRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetAutoRenew(v bool) *CreateLaunchTemplateRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetAutoRenewPeriod(v int32) *CreateLaunchTemplateRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetCreditSpecification(v string) *CreateLaunchTemplateRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetDataDisk(v []*CreateLaunchTemplateRequestDataDisk) *CreateLaunchTemplateRequest {
	s.DataDisk = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetDeletionProtection(v bool) *CreateLaunchTemplateRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetDeploymentSetId(v string) *CreateLaunchTemplateRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetDescription(v string) *CreateLaunchTemplateRequest {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetEnableVmOsConfig(v bool) *CreateLaunchTemplateRequest {
	s.EnableVmOsConfig = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetHostName(v string) *CreateLaunchTemplateRequest {
	s.HostName = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetHttpEndpoint(v string) *CreateLaunchTemplateRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetHttpPutResponseHopLimit(v int32) *CreateLaunchTemplateRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetHttpTokens(v string) *CreateLaunchTemplateRequest {
	s.HttpTokens = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetImageId(v string) *CreateLaunchTemplateRequest {
	s.ImageId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetImageOptions(v *CreateLaunchTemplateRequestImageOptions) *CreateLaunchTemplateRequest {
	s.ImageOptions = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetImageOwnerAlias(v string) *CreateLaunchTemplateRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInstanceChargeType(v string) *CreateLaunchTemplateRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInstanceName(v string) *CreateLaunchTemplateRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInstanceType(v string) *CreateLaunchTemplateRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInternetChargeType(v string) *CreateLaunchTemplateRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInternetMaxBandwidthIn(v int32) *CreateLaunchTemplateRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetInternetMaxBandwidthOut(v int32) *CreateLaunchTemplateRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetIoOptimized(v string) *CreateLaunchTemplateRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetIpv6AddressCount(v int32) *CreateLaunchTemplateRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetKeyPairName(v string) *CreateLaunchTemplateRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetLaunchTemplateName(v string) *CreateLaunchTemplateRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetNetworkInterface(v []*CreateLaunchTemplateRequestNetworkInterface) *CreateLaunchTemplateRequest {
	s.NetworkInterface = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetNetworkType(v string) *CreateLaunchTemplateRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetOwnerAccount(v string) *CreateLaunchTemplateRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetOwnerId(v int64) *CreateLaunchTemplateRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetPasswordInherit(v bool) *CreateLaunchTemplateRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetPeriod(v int32) *CreateLaunchTemplateRequest {
	s.Period = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetPeriodUnit(v string) *CreateLaunchTemplateRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetPrivateIpAddress(v string) *CreateLaunchTemplateRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetRamRoleName(v string) *CreateLaunchTemplateRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetRegionId(v string) *CreateLaunchTemplateRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetResourceGroupId(v string) *CreateLaunchTemplateRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetResourceOwnerAccount(v string) *CreateLaunchTemplateRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetResourceOwnerId(v int64) *CreateLaunchTemplateRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSecurityEnhancementStrategy(v string) *CreateLaunchTemplateRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSecurityGroupId(v string) *CreateLaunchTemplateRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSecurityGroupIds(v []*string) *CreateLaunchTemplateRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSecurityOptions(v *CreateLaunchTemplateRequestSecurityOptions) *CreateLaunchTemplateRequest {
	s.SecurityOptions = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSpotDuration(v int32) *CreateLaunchTemplateRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSpotPriceLimit(v float32) *CreateLaunchTemplateRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetSpotStrategy(v string) *CreateLaunchTemplateRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetTag(v []*CreateLaunchTemplateRequestTag) *CreateLaunchTemplateRequest {
	s.Tag = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetTemplateResourceGroupId(v string) *CreateLaunchTemplateRequest {
	s.TemplateResourceGroupId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetTemplateTag(v []*CreateLaunchTemplateRequestTemplateTag) *CreateLaunchTemplateRequest {
	s.TemplateTag = v
	return s
}

func (s *CreateLaunchTemplateRequest) SetUserData(v string) *CreateLaunchTemplateRequest {
	s.UserData = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetVSwitchId(v string) *CreateLaunchTemplateRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetVersionDescription(v string) *CreateLaunchTemplateRequest {
	s.VersionDescription = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetVpcId(v string) *CreateLaunchTemplateRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) SetZoneId(v string) *CreateLaunchTemplateRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateLaunchTemplateRequest) Validate() error {
	if s.SystemDisk != nil {
		if err := s.SystemDisk.Validate(); err != nil {
			return err
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
	if s.ImageOptions != nil {
		if err := s.ImageOptions.Validate(); err != nil {
			return err
		}
	}
	if s.NetworkInterface != nil {
		for _, item := range s.NetworkInterface {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SecurityOptions != nil {
		if err := s.SecurityOptions.Validate(); err != nil {
			return err
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
	if s.TemplateTag != nil {
		for _, item := range s.TemplateTag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateLaunchTemplateRequestSystemDisk struct {
	// The ID of the automatic snapshot policy to apply to the system disk.
	//
	// example:
	//
	// sp-gc7c37d4ylw7mtnk****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature for the system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of the system disk. Valid values:
	//
	// 	- cloud: basic disk.
	//
	// 	- cloud_efficiency: ultra disk.
	//
	// 	- cloud_ssd: standard SSD.
	//
	// 	- cloud_essd: Enterprise SSD (ESSD). You can use `SystemDisk.PerformanceLevel` to set the performance level of the ESSD to use as the system disk.
	//
	// 	- cloud_auto: ESSD AutoPL disk.
	//
	// 	- cloud_essd_entry: ESSD Entry disk.
	//
	// For non-I/O optimized instances of retired instance types, the default value is cloud. For other types of instances, the default value is cloud_efficiency.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release the system disk when the instance is released. Valid values:
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
	// The description of the system disk. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testSystemDiskDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testSystemDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  If you create an instance in Hong Kong Zone D or Singapore Zone A, you cannot encrypt the system disk.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// > This parameter is in invitational preview and is unavailable for general users.
	//
	// example:
	//
	// null
	Iops *int32 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	// The ID of the KMS key to use for the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the ESSD to use as the system disk. Default value: PL0. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// For more information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk to use as the system disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}
	//
	// Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}
	//
	// > This parameter is available only if you set the SystemDisk.Category parameter to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html) and [Modify the performance configurations of an ESSD AutoPL disk](https://help.aliyun.com/document_detail/413275.html).
	//
	// example:
	//
	// 50000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of the system disk. Unit: GiB. Valid values:
	//
	// 	- Valid values if you set SystemDisk.Category to cloud: 20 to 500.
	//
	// 	- Valid values if you set SystemDisk.Category to other disk categories: 20 to 2048.
	//
	// The value of this parameter must be at least 20 and greater than or equal to the size of the image.
	//
	// example:
	//
	// 40
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
}

func (s CreateLaunchTemplateRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetIops() *int32 {
	return s.Iops
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateLaunchTemplateRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetBurstingEnabled(v bool) *CreateLaunchTemplateRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetCategory(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetDeleteWithInstance(v bool) *CreateLaunchTemplateRequestSystemDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetDescription(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetDiskName(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetEncrypted(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetIops(v int32) *CreateLaunchTemplateRequestSystemDisk {
	s.Iops = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetKMSKeyId(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetPerformanceLevel(v string) *CreateLaunchTemplateRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetProvisionedIops(v int64) *CreateLaunchTemplateRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) SetSize(v int32) *CreateLaunchTemplateRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateLaunchTemplateRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestDataDisk struct {
	// The ID of the automatic snapshot policy to apply to data disk N.
	//
	// example:
	//
	// sp-m5e7fa9ute44ssa****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature for the system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of data disk N. Valid values:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_efficiency: utra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: ESSD
	//
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// 	- cloud_essd_entry: ESSD Entry disk
	//
	// For I/O optimized instances, the default value is cloud_efficiency. For non-I/O optimized instances, the default value is cloud.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release data disk N when the associated instance is released. Valid values:
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
	// testDataDiskDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of data disk N. The mount points are named based on the number of data disks:
	//
	// 	- 1st to 25th data disks: /dev/xvd`[b-z]`.
	//
	// 	- From the 26th data disk on: /dev/xvd`[aa-zz]`. For example, the 26th data disk is named /dev/xvdaa, the 27th data disk is named /dev/xvdab, and so on.
	//
	// >  This parameter is applicable to scenarios in which a full image is used to create instances. A full image is an image that contains an operating system, application software, and business data. For these scenarios, you can set the parameter to the mount point of data disk N contained in the full image and modify `DataDisk.N.Size` and `DataDisk.N.Category` to change the category and size of data disk N created based on the image.
	//
	// example:
	//
	// null
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of data disk N. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testDataDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt data disk N.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key used for the data disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the ESSD to use as data disk N. The value of N must be the same as that in `DataDisk.N.Category` when DataDisk.N.Category is set to cloud_essd. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10000 random read/write IOPS.
	//
	// 	- PL1 (default): A single ESSD can deliver up to 50000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1000000 random read/write IOPS.
	//
	// For information about ESSD performance levels, see [ESSDs](https://help.aliyun.com/document_detail/122389.html).
	//
	// example:
	//
	// PL1
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}.
	//
	// Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// >  This parameter is available only if you set DiskCategory to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html) and [Modify the performance configurations of an ESSD AutoPL disk](https://help.aliyun.com/document_detail/413275.html).
	//
	// example:
	//
	// 50000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of data disk N. Valid values of N: 1 to 16. Unit: GiB. Valid values:
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud: 5 to 2000.
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud_efficiency: 20 to 32768.
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud_ssd: 20 to 32768.
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud_essd: vary based on the value of `DataDisk.N.PerformanceLevel`.
	//
	//     	- Valid values if DataDisk.N.PerformanceLevel is set to PL0: 1 to 32768.
	//
	//     	- Valid values if DataDisk.N.PerformanceLevel is set to PL1: 20 to 32768.
	//
	//     	- Valid values if DataDisk.N.PerformanceLevel is set to PL2: 461 to 32768.
	//
	//     	- Valid values if you set DataDisk.N.PerformanceLevel to PL3: 1261 to 32768.
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud_auto: 1 to 32768.
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud_essd_entry: 10 to 32768.
	//
	// The value of this parameter must be greater than or equal to the size of the snapshot specified by `SnapshotId`.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot to use to create data disk N. Valid values of N: 1 to 16. If you specify `DataDisk.N.SnapshotId`, `DataDisk.N.Size` is ignored. The data disk is created with the size of the specified snapshot.
	//
	// >  Use snapshots created on or after July 15, 2013. Otherwise, an error is returned and your request is rejected.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateLaunchTemplateRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestDataDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateLaunchTemplateRequestDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateLaunchTemplateRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateLaunchTemplateRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateLaunchTemplateRequestDataDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateRequestDataDisk) GetDevice() *string {
	return s.Device
}

func (s *CreateLaunchTemplateRequestDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateLaunchTemplateRequestDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateLaunchTemplateRequestDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateLaunchTemplateRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateLaunchTemplateRequestDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateLaunchTemplateRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateLaunchTemplateRequestDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateLaunchTemplateRequestDataDisk) SetAutoSnapshotPolicyId(v string) *CreateLaunchTemplateRequestDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetBurstingEnabled(v bool) *CreateLaunchTemplateRequestDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetCategory(v string) *CreateLaunchTemplateRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetDeleteWithInstance(v bool) *CreateLaunchTemplateRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetDescription(v string) *CreateLaunchTemplateRequestDataDisk {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetDevice(v string) *CreateLaunchTemplateRequestDataDisk {
	s.Device = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetDiskName(v string) *CreateLaunchTemplateRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetEncrypted(v string) *CreateLaunchTemplateRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetKMSKeyId(v string) *CreateLaunchTemplateRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetPerformanceLevel(v string) *CreateLaunchTemplateRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetProvisionedIops(v int64) *CreateLaunchTemplateRequestDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetSize(v int32) *CreateLaunchTemplateRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) SetSnapshotId(v string) *CreateLaunchTemplateRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateLaunchTemplateRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestImageOptions struct {
	// Specifies whether the instance that uses the image supports logons of the ecs-user user. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	LoginAsNonRoot *bool `json:"LoginAsNonRoot,omitempty" xml:"LoginAsNonRoot,omitempty"`
}

func (s CreateLaunchTemplateRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestImageOptions) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *CreateLaunchTemplateRequestImageOptions) SetLoginAsNonRoot(v bool) *CreateLaunchTemplateRequestImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *CreateLaunchTemplateRequestImageOptions) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestNetworkInterface struct {
	// Specifies whether to release ENI N when the instance is released. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: true.
	//
	// >  This parameter takes effect only for secondary ENIs.
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of the secondary ENI. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`. The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// example:
	//
	// testEniDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of ENI N. Valid values of N: 1 and 2. If the value of N is 1, you can configure a primary or secondary ENI. If the value of N is 2, you must configure a primary ENI and a secondary ENI.
	//
	// Valid values:
	//
	// 	- Primary
	//
	// 	- Secondary
	//
	// Default value: Secondary.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The name of ENI N.
	//
	// Take note of the following items:
	//
	// 	- Valid values of N: 1 and 2. If the value of N is 1, you can configure a primary or secondary ENI. If the value of N is 2, you must configure a primary ENI and a secondary ENI.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Primary`, you do not need to specify this parameter.
	//
	// example:
	//
	// testEniName
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication mode of the primary ENI. Valid values:
	//
	// 	- Standard: uses the TCP communication mode.
	//
	// 	- HighPerformance: uses the remote direct memory access (RDMA) communication mode with Elastic RDMA Interface (ERI) enabled.
	//
	// example:
	//
	// Standard
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The primary IP address to assign to ENI N.
	//
	// Take note of the following items:
	//
	// 	- Valid values of N: 1 and 2.
	//
	//     	- If the value of N is 1, you can configure a primary or secondary ENI. If you specify this parameter, set `Amount` to a numeric value greater than 1, and set NetworkInterface.N.InstanceType to Primary, the specified number of instances are created and consecutive primary IP addresses starting from the specified IP address are assigned to the instances. In this case, you cannot attach secondary ENIs to the instances.
	//
	//     	- If the value of N is 2, you must configure a primary ENI and a secondary ENI. If you specify this parameter, set `Amount` to a numeric value greater than 1, and set NetworkInterface.N.InstanceType to Primary, you cannot set `NetworkInterface.2.InstanceType` to Secondary to attach a secondary ENI.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Primary`, this parameter is equivalent to `PrivateIpAddress`. You cannot specify both this parameter and `PrivateIpAddress` in the same request.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Secondary` or leave NetworkInterface.N.InstanceType empty, the specified primary IP address is assigned to the secondary ENI. The default value is an IP address that is randomly selected from within the CIDR block of the vSwitch to which to connect the secondary ENI.
	//
	// >  You can attach only a single secondary ENI when you create an instance. After the instance is created, you can call the [CreateNetworkInterface](https://help.aliyun.com/document_detail/58504.html) and [AttachNetworkInterface](https://help.aliyun.com/document_detail/58515.html) operations to attach more secondary ENIs.
	//
	// example:
	//
	// ``192.168.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// The ID of the security group to which to assign ENI N.
	//
	// Take note of the following items:
	//
	// 	- Valid values of N: 1 and 2. If the value of N is 1, you can configure a primary or secondary ENI. If the value of N is 2, you must configure a primary ENI and a secondary ENI.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Primary`, you must specify this parameter. In this case, this parameter is equivalent to `SecurityGroupId`, and you cannot specify `SecurityGroupId`, `SecurityGroupIds.N`, or `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Secondary` or leave NetworkInterface.N.InstanceType empty, you do not need to specify this parameter. The default value is the ID of the security group to which to assign the instance.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of security groups to which to assign ENI N.
	//
	// 	- Valid values of the first N: 1 and 2. If the value of N is 1, you can configure a primary or secondary ENI. If the value of N is 2, you must configure a primary ENI and a secondary ENI.
	//
	// 	- The second N in this parameter indicates that one or more security group IDs can be specified. The valid values of N vary based on the maximum number of security groups to which an instance can belong. For more information, see the [Security group limits](~~25412#SecurityGroupQuota1~~) section of the "Limits" topic.
	//
	// Take note of the following items:
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Primary`, you must specify this parameter or `NetworkInterface.N.SecurityGroupId`. In this case, this parameter is equivalent to `SecurityGroupIds.N`, and you cannot specify `SecurityGroupId`, `SecurityGroupIds.N`, or `NetworkInterface.N.SecurityGroupId`.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Secondary` or leave NetworkInterface.N.InstanceType empty, you do not need to specify this parameter. The default value is the ID of the security group to which to assign the instance.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which to connect ENI N.
	//
	// Take note of the following items:
	//
	// 	- Valid values of N: 1 and 2. If the value of N is 1, you can configure a primary or secondary ENI. If the value of N is 2, you must configure a primary ENI and a secondary ENI.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Primary`, you must specify this parameter. In this case, this parameter is equivalent to `VSwitchId`. You cannot specify both NetworkInterface.N.VSwitchId and `VSwitchId` in the same request.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Secondary` or leave NetworkInterface.N.InstanceType empty, you do not need to specify this parameter. The default value is the VSwitchId value.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateLaunchTemplateRequestNetworkInterface) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestNetworkInterface) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateLaunchTemplateRequestNetworkInterface) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetDeleteOnRelease(v bool) *CreateLaunchTemplateRequestNetworkInterface {
	s.DeleteOnRelease = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetDescription(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetInstanceType(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.InstanceType = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetNetworkInterfaceName(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.NetworkInterfaceName = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetNetworkInterfaceTrafficMode(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetPrimaryIpAddress(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.PrimaryIpAddress = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetSecurityGroupId(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetSecurityGroupIds(v []*string) *CreateLaunchTemplateRequestNetworkInterface {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) SetVSwitchId(v string) *CreateLaunchTemplateRequestNetworkInterface {
	s.VSwitchId = &v
	return s
}

func (s *CreateLaunchTemplateRequestNetworkInterface) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestSecurityOptions struct {
	TrustedSystemMode *string `json:"TrustedSystemMode,omitempty" xml:"TrustedSystemMode,omitempty"`
}

func (s CreateLaunchTemplateRequestSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestSecurityOptions) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestSecurityOptions) GetTrustedSystemMode() *string {
	return s.TrustedSystemMode
}

func (s *CreateLaunchTemplateRequestSecurityOptions) SetTrustedSystemMode(v string) *CreateLaunchTemplateRequestSecurityOptions {
	s.TrustedSystemMode = &v
	return s
}

func (s *CreateLaunchTemplateRequestSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestTag struct {
	// The key of tag N to add to the instance, disks, and primary ENI that are created from the launch template. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the instance, disks, and primary ENI that are created from the launch template. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLaunchTemplateRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLaunchTemplateRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLaunchTemplateRequestTag) SetKey(v string) *CreateLaunchTemplateRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLaunchTemplateRequestTag) SetValue(v string) *CreateLaunchTemplateRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLaunchTemplateRequestTag) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateRequestTemplateTag struct {
	// The key of tag N to add to the launch template. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag key cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the launch template. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain `http://` or `https://`. The tag value cannot start with `acs:` or `aliyun`.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLaunchTemplateRequestTemplateTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateRequestTemplateTag) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateRequestTemplateTag) GetKey() *string {
	return s.Key
}

func (s *CreateLaunchTemplateRequestTemplateTag) GetValue() *string {
	return s.Value
}

func (s *CreateLaunchTemplateRequestTemplateTag) SetKey(v string) *CreateLaunchTemplateRequestTemplateTag {
	s.Key = &v
	return s
}

func (s *CreateLaunchTemplateRequestTemplateTag) SetValue(v string) *CreateLaunchTemplateRequestTemplateTag {
	s.Value = &v
	return s
}

func (s *CreateLaunchTemplateRequestTemplateTag) Validate() error {
	return dara.Validate(s)
}
