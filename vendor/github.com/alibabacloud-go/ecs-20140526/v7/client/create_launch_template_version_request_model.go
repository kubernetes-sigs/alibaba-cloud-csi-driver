// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLaunchTemplateVersionRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSystemDisk(v *CreateLaunchTemplateVersionRequestSystemDisk) *CreateLaunchTemplateVersionRequest
	GetSystemDisk() *CreateLaunchTemplateVersionRequestSystemDisk
	SetAutoReleaseTime(v string) *CreateLaunchTemplateVersionRequest
	GetAutoReleaseTime() *string
	SetAutoRenew(v bool) *CreateLaunchTemplateVersionRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *CreateLaunchTemplateVersionRequest
	GetAutoRenewPeriod() *int32
	SetCreditSpecification(v string) *CreateLaunchTemplateVersionRequest
	GetCreditSpecification() *string
	SetDataDisk(v []*CreateLaunchTemplateVersionRequestDataDisk) *CreateLaunchTemplateVersionRequest
	GetDataDisk() []*CreateLaunchTemplateVersionRequestDataDisk
	SetDeletionProtection(v bool) *CreateLaunchTemplateVersionRequest
	GetDeletionProtection() *bool
	SetDeploymentSetId(v string) *CreateLaunchTemplateVersionRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *CreateLaunchTemplateVersionRequest
	GetDescription() *string
	SetEnableVmOsConfig(v bool) *CreateLaunchTemplateVersionRequest
	GetEnableVmOsConfig() *bool
	SetHostName(v string) *CreateLaunchTemplateVersionRequest
	GetHostName() *string
	SetHttpEndpoint(v string) *CreateLaunchTemplateVersionRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *CreateLaunchTemplateVersionRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *CreateLaunchTemplateVersionRequest
	GetHttpTokens() *string
	SetImageId(v string) *CreateLaunchTemplateVersionRequest
	GetImageId() *string
	SetImageOptions(v *CreateLaunchTemplateVersionRequestImageOptions) *CreateLaunchTemplateVersionRequest
	GetImageOptions() *CreateLaunchTemplateVersionRequestImageOptions
	SetImageOwnerAlias(v string) *CreateLaunchTemplateVersionRequest
	GetImageOwnerAlias() *string
	SetInstanceChargeType(v string) *CreateLaunchTemplateVersionRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *CreateLaunchTemplateVersionRequest
	GetInstanceName() *string
	SetInstanceType(v string) *CreateLaunchTemplateVersionRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *CreateLaunchTemplateVersionRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *CreateLaunchTemplateVersionRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *CreateLaunchTemplateVersionRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *CreateLaunchTemplateVersionRequest
	GetIoOptimized() *string
	SetIpv6AddressCount(v int32) *CreateLaunchTemplateVersionRequest
	GetIpv6AddressCount() *int32
	SetKeyPairName(v string) *CreateLaunchTemplateVersionRequest
	GetKeyPairName() *string
	SetLaunchTemplateId(v string) *CreateLaunchTemplateVersionRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateName(v string) *CreateLaunchTemplateVersionRequest
	GetLaunchTemplateName() *string
	SetNetworkInterface(v []*CreateLaunchTemplateVersionRequestNetworkInterface) *CreateLaunchTemplateVersionRequest
	GetNetworkInterface() []*CreateLaunchTemplateVersionRequestNetworkInterface
	SetNetworkType(v string) *CreateLaunchTemplateVersionRequest
	GetNetworkType() *string
	SetOwnerAccount(v string) *CreateLaunchTemplateVersionRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *CreateLaunchTemplateVersionRequest
	GetOwnerId() *int64
	SetPasswordInherit(v bool) *CreateLaunchTemplateVersionRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *CreateLaunchTemplateVersionRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *CreateLaunchTemplateVersionRequest
	GetPeriodUnit() *string
	SetPrivateIpAddress(v string) *CreateLaunchTemplateVersionRequest
	GetPrivateIpAddress() *string
	SetRamRoleName(v string) *CreateLaunchTemplateVersionRequest
	GetRamRoleName() *string
	SetRegionId(v string) *CreateLaunchTemplateVersionRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *CreateLaunchTemplateVersionRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *CreateLaunchTemplateVersionRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *CreateLaunchTemplateVersionRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *CreateLaunchTemplateVersionRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *CreateLaunchTemplateVersionRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *CreateLaunchTemplateVersionRequest
	GetSecurityGroupIds() []*string
	SetSecurityOptions(v *CreateLaunchTemplateVersionRequestSecurityOptions) *CreateLaunchTemplateVersionRequest
	GetSecurityOptions() *CreateLaunchTemplateVersionRequestSecurityOptions
	SetSpotDuration(v int32) *CreateLaunchTemplateVersionRequest
	GetSpotDuration() *int32
	SetSpotPriceLimit(v float32) *CreateLaunchTemplateVersionRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *CreateLaunchTemplateVersionRequest
	GetSpotStrategy() *string
	SetTag(v []*CreateLaunchTemplateVersionRequestTag) *CreateLaunchTemplateVersionRequest
	GetTag() []*CreateLaunchTemplateVersionRequestTag
	SetUserData(v string) *CreateLaunchTemplateVersionRequest
	GetUserData() *string
	SetVSwitchId(v string) *CreateLaunchTemplateVersionRequest
	GetVSwitchId() *string
	SetVersionDescription(v string) *CreateLaunchTemplateVersionRequest
	GetVersionDescription() *string
	SetVpcId(v string) *CreateLaunchTemplateVersionRequest
	GetVpcId() *string
	SetZoneId(v string) *CreateLaunchTemplateVersionRequest
	GetZoneId() *string
}

type CreateLaunchTemplateVersionRequest struct {
	SystemDisk *CreateLaunchTemplateVersionRequestSystemDisk `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
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
	// Specifies whether to enable auto-renewal for the instance. This parameter is valid only if `InstanceChargeType` is set to `PrePaid`. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
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
	// The information about the data disks.
	DataDisk []*CreateLaunchTemplateVersionRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
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
	// The description of the instance. The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to enable the operating system configuration of the instance.
	//
	// example:
	//
	// false
	EnableVmOsConfig *bool `json:"EnableVmOsConfig,omitempty" xml:"EnableVmOsConfig,omitempty"`
	// The hostname of the instance.
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
	// The ID of the image to use to create the Elastic Compute Service (ECS) instance. You can call the [DescribeImages](https://help.aliyun.com/document_detail/25534.html) operation to query available images.
	//
	// example:
	//
	// win2008r2_64_ent_sp1_en-us_40G_alibase_20170915.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Details about the image options.
	ImageOptions *CreateLaunchTemplateVersionRequestImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The source of the image.
	//
	// > This parameter will be removed in the future. We recommend that you use other parameters to ensure future compatibility.
	//
	// example:
	//
	// system
	ImageOwnerAlias *string `json:"ImageOwnerAlias,omitempty" xml:"ImageOwnerAlias,omitempty"`
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription. If you set this parameter to PrePaid, make sure that your account has sufficient credits. Otherwise, an `InvalidPayMethod` error is returned.
	//
	// 	- PostPaid: pay-as-you-go
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
	// The instance type. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html). You can also call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the most recent list of instance types.
	//
	// example:
	//
	// ecs.g5.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for network usage. Default value: PayByTraffic. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth
	//
	// 	- PayByTraffic: pay-by-traffic
	//
	// > When the **pay-by-traffic*	- billing method for network usage is used, the maximum inbound and outbound bandwidths are used as the upper limits of bandwidths instead of guaranteed performance specifications. In scenarios where demand outstrips resource supplies, these maximum bandwidth values may not be reached. If you want guaranteed bandwidths for your instance, use the **pay-by-bandwidth*	- billing method for network usage.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// 	- When the purchased outbound public bandwidth is less than or equal to 10 Mbit/s, the valid values of this parameter are 1 to 10 and the default value is 10.
	//
	// 	- If the purchased outbound public bandwidth is greater than 10 Mbit/s, the valid values of this parameter range from 1 to the `InternetMaxBandwidthOut` value and the default value is the `InternetMaxBandwidthOut` value.
	//
	// example:
	//
	// 50
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// example:
	//
	// 5
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether to create an I/O optimized instance. Valid values:
	//
	// 	- none: The instance is not I/O optimized.
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
	// The name of the key pair to bind to the instance.
	//
	// 	- For Windows instances, this parameter is ignored The `Password` parameter is valid even if the KeyPairName parameter is specified.
	//
	// 	- For Linux instances, the password-based logon method is disabled by default.
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The ID of the launch template. For more information, call the [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html) operation. You must specify `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template.
	//
	// example:
	//
	// lt-m5eiaupmvm2op9d****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The name of the launch template. The name must be 2 to 128 characters in length. The name must start with a letter and cannot start with `http://` or `https://`. The name can contain letters, digits, colons (:), underscores (_), and hyphens (-).
	//
	// example:
	//
	// testLaunchTemplateName
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// The information of the elastic network interfaces (ENIs).
	NetworkInterface []*CreateLaunchTemplateVersionRequestNetworkInterface `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty" type:"Repeated"`
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
	// Specifies whether to use the password that is preconfigured in the image. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// >  If you specify PasswordInherit, you must leave Password empty and make sure that a password is preconfigured for the image.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription period of the instance. Unit: months. This parameter is valid and required only when `InstanceChargeType` is set to `PrePaid`. Valid values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
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
	// To assign a private IP address to an instance of the VPC type, make sure that the IP address is an idle IP address within the CIDR block of the vSwitch specified by the `VSwitchId` parameter.
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
	// The region ID of the command. You can call the [DescribeRegions](https://help.aliyun.com/document_detail/25609.html) operation to query the most recent list of regions.
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
	// Specifies whether to enable security hardening for the operating system. Valid values:
	//
	// 	- Active: Security hardening is enabled. This value is applicable only to public images.
	//
	// 	- Deactive: Security hardening is disabled. This value is available to all types of images.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which to assign the ECS instance created based on the launch template version. Instances in the same security group can access each other.
	//
	// >  You cannot specify `SecurityGroupId` and `SecurityGroupIds.N` in the same request.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The ID of security group N to which to assign the instance. The valid values of N depend on the maximum number of security groups to which the instance can belong. For more information, see [Quantity limits](https://help.aliyun.com/document_detail/25412.html).
	//
	// > The `SecurityGroupId` parameter and the `SecurityGroupIds.N` parameter are mutually exclusive.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupIds []*string                                          `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	SecurityOptions  *CreateLaunchTemplateVersionRequestSecurityOptions `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
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
	// The maximum hourly price of the spot instance. A maximum of three decimal places are allowed.
	//
	// example:
	//
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The preemption policy for the pay-as-you-go instance. This parameter is valid only when the `InstanceChargeType` parameter is set to `PostPaid`. Default value: NoSpot. Valid values:
	//
	// 	- NoSpot: The instance is created as a pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instances of the compute node are spot instances. These types of instances have a specified maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is created as a spot instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The tags to add to the ECS instance, disks, and primary elastic network interface (ENI) created based on the launch template version.
	Tag []*CreateLaunchTemplateVersionRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// The user data of the instance. The user data must be encoded in Base64. The maximum size of raw data is 32 KB.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBl****
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
	// The ID of the virtual private cloud (VPC) in which to create the ECS instance.
	//
	// example:
	//
	// vpc-bp12433upq1y5scen****
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	// The ID of the zone to which the instance belongs.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s CreateLaunchTemplateVersionRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequest) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequest) GetSystemDisk() *CreateLaunchTemplateVersionRequestSystemDisk {
	return s.SystemDisk
}

func (s *CreateLaunchTemplateVersionRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *CreateLaunchTemplateVersionRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *CreateLaunchTemplateVersionRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *CreateLaunchTemplateVersionRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *CreateLaunchTemplateVersionRequest) GetDataDisk() []*CreateLaunchTemplateVersionRequestDataDisk {
	return s.DataDisk
}

func (s *CreateLaunchTemplateVersionRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *CreateLaunchTemplateVersionRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *CreateLaunchTemplateVersionRequest) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateVersionRequest) GetEnableVmOsConfig() *bool {
	return s.EnableVmOsConfig
}

func (s *CreateLaunchTemplateVersionRequest) GetHostName() *string {
	return s.HostName
}

func (s *CreateLaunchTemplateVersionRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *CreateLaunchTemplateVersionRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *CreateLaunchTemplateVersionRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *CreateLaunchTemplateVersionRequest) GetImageId() *string {
	return s.ImageId
}

func (s *CreateLaunchTemplateVersionRequest) GetImageOptions() *CreateLaunchTemplateVersionRequestImageOptions {
	return s.ImageOptions
}

func (s *CreateLaunchTemplateVersionRequest) GetImageOwnerAlias() *string {
	return s.ImageOwnerAlias
}

func (s *CreateLaunchTemplateVersionRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *CreateLaunchTemplateVersionRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *CreateLaunchTemplateVersionRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateLaunchTemplateVersionRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *CreateLaunchTemplateVersionRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *CreateLaunchTemplateVersionRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *CreateLaunchTemplateVersionRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *CreateLaunchTemplateVersionRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *CreateLaunchTemplateVersionRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *CreateLaunchTemplateVersionRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *CreateLaunchTemplateVersionRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *CreateLaunchTemplateVersionRequest) GetNetworkInterface() []*CreateLaunchTemplateVersionRequestNetworkInterface {
	return s.NetworkInterface
}

func (s *CreateLaunchTemplateVersionRequest) GetNetworkType() *string {
	return s.NetworkType
}

func (s *CreateLaunchTemplateVersionRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *CreateLaunchTemplateVersionRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *CreateLaunchTemplateVersionRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *CreateLaunchTemplateVersionRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *CreateLaunchTemplateVersionRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *CreateLaunchTemplateVersionRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *CreateLaunchTemplateVersionRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *CreateLaunchTemplateVersionRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateLaunchTemplateVersionRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *CreateLaunchTemplateVersionRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *CreateLaunchTemplateVersionRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *CreateLaunchTemplateVersionRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *CreateLaunchTemplateVersionRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateLaunchTemplateVersionRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateLaunchTemplateVersionRequest) GetSecurityOptions() *CreateLaunchTemplateVersionRequestSecurityOptions {
	return s.SecurityOptions
}

func (s *CreateLaunchTemplateVersionRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *CreateLaunchTemplateVersionRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *CreateLaunchTemplateVersionRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *CreateLaunchTemplateVersionRequest) GetTag() []*CreateLaunchTemplateVersionRequestTag {
	return s.Tag
}

func (s *CreateLaunchTemplateVersionRequest) GetUserData() *string {
	return s.UserData
}

func (s *CreateLaunchTemplateVersionRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLaunchTemplateVersionRequest) GetVersionDescription() *string {
	return s.VersionDescription
}

func (s *CreateLaunchTemplateVersionRequest) GetVpcId() *string {
	return s.VpcId
}

func (s *CreateLaunchTemplateVersionRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *CreateLaunchTemplateVersionRequest) SetSystemDisk(v *CreateLaunchTemplateVersionRequestSystemDisk) *CreateLaunchTemplateVersionRequest {
	s.SystemDisk = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetAutoReleaseTime(v string) *CreateLaunchTemplateVersionRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetAutoRenew(v bool) *CreateLaunchTemplateVersionRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetAutoRenewPeriod(v int32) *CreateLaunchTemplateVersionRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetCreditSpecification(v string) *CreateLaunchTemplateVersionRequest {
	s.CreditSpecification = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetDataDisk(v []*CreateLaunchTemplateVersionRequestDataDisk) *CreateLaunchTemplateVersionRequest {
	s.DataDisk = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetDeletionProtection(v bool) *CreateLaunchTemplateVersionRequest {
	s.DeletionProtection = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetDeploymentSetId(v string) *CreateLaunchTemplateVersionRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetDescription(v string) *CreateLaunchTemplateVersionRequest {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetEnableVmOsConfig(v bool) *CreateLaunchTemplateVersionRequest {
	s.EnableVmOsConfig = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetHostName(v string) *CreateLaunchTemplateVersionRequest {
	s.HostName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetHttpEndpoint(v string) *CreateLaunchTemplateVersionRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetHttpPutResponseHopLimit(v int32) *CreateLaunchTemplateVersionRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetHttpTokens(v string) *CreateLaunchTemplateVersionRequest {
	s.HttpTokens = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetImageId(v string) *CreateLaunchTemplateVersionRequest {
	s.ImageId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetImageOptions(v *CreateLaunchTemplateVersionRequestImageOptions) *CreateLaunchTemplateVersionRequest {
	s.ImageOptions = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetImageOwnerAlias(v string) *CreateLaunchTemplateVersionRequest {
	s.ImageOwnerAlias = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInstanceChargeType(v string) *CreateLaunchTemplateVersionRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInstanceName(v string) *CreateLaunchTemplateVersionRequest {
	s.InstanceName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInstanceType(v string) *CreateLaunchTemplateVersionRequest {
	s.InstanceType = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInternetChargeType(v string) *CreateLaunchTemplateVersionRequest {
	s.InternetChargeType = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInternetMaxBandwidthIn(v int32) *CreateLaunchTemplateVersionRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetInternetMaxBandwidthOut(v int32) *CreateLaunchTemplateVersionRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetIoOptimized(v string) *CreateLaunchTemplateVersionRequest {
	s.IoOptimized = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetIpv6AddressCount(v int32) *CreateLaunchTemplateVersionRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetKeyPairName(v string) *CreateLaunchTemplateVersionRequest {
	s.KeyPairName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetLaunchTemplateId(v string) *CreateLaunchTemplateVersionRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetLaunchTemplateName(v string) *CreateLaunchTemplateVersionRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetNetworkInterface(v []*CreateLaunchTemplateVersionRequestNetworkInterface) *CreateLaunchTemplateVersionRequest {
	s.NetworkInterface = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetNetworkType(v string) *CreateLaunchTemplateVersionRequest {
	s.NetworkType = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetOwnerAccount(v string) *CreateLaunchTemplateVersionRequest {
	s.OwnerAccount = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetOwnerId(v int64) *CreateLaunchTemplateVersionRequest {
	s.OwnerId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetPasswordInherit(v bool) *CreateLaunchTemplateVersionRequest {
	s.PasswordInherit = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetPeriod(v int32) *CreateLaunchTemplateVersionRequest {
	s.Period = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetPeriodUnit(v string) *CreateLaunchTemplateVersionRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetPrivateIpAddress(v string) *CreateLaunchTemplateVersionRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetRamRoleName(v string) *CreateLaunchTemplateVersionRequest {
	s.RamRoleName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetRegionId(v string) *CreateLaunchTemplateVersionRequest {
	s.RegionId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetResourceGroupId(v string) *CreateLaunchTemplateVersionRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetResourceOwnerAccount(v string) *CreateLaunchTemplateVersionRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetResourceOwnerId(v int64) *CreateLaunchTemplateVersionRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSecurityEnhancementStrategy(v string) *CreateLaunchTemplateVersionRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSecurityGroupId(v string) *CreateLaunchTemplateVersionRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSecurityGroupIds(v []*string) *CreateLaunchTemplateVersionRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSecurityOptions(v *CreateLaunchTemplateVersionRequestSecurityOptions) *CreateLaunchTemplateVersionRequest {
	s.SecurityOptions = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSpotDuration(v int32) *CreateLaunchTemplateVersionRequest {
	s.SpotDuration = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSpotPriceLimit(v float32) *CreateLaunchTemplateVersionRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetSpotStrategy(v string) *CreateLaunchTemplateVersionRequest {
	s.SpotStrategy = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetTag(v []*CreateLaunchTemplateVersionRequestTag) *CreateLaunchTemplateVersionRequest {
	s.Tag = v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetUserData(v string) *CreateLaunchTemplateVersionRequest {
	s.UserData = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetVSwitchId(v string) *CreateLaunchTemplateVersionRequest {
	s.VSwitchId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetVersionDescription(v string) *CreateLaunchTemplateVersionRequest {
	s.VersionDescription = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetVpcId(v string) *CreateLaunchTemplateVersionRequest {
	s.VpcId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) SetZoneId(v string) *CreateLaunchTemplateVersionRequest {
	s.ZoneId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequest) Validate() error {
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
	return nil
}

type CreateLaunchTemplateVersionRequestSystemDisk struct {
	// The ID of the automatic snapshot policy to apply to the system disk.
	//
	// example:
	//
	// sp-bp1dgzpaxwc4load****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature. Valid values:
	//
	// 	- true: encrypts the disk.
	//
	// 	- false: does not enable the performance burst feature.
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
	// 	- cloud_auto: Enterprise SSD (ESSD) AutoPL disk.
	//
	// 	- cloud_essd: ESSD. You can use `SystemDisk.PerformanceLevel` to set the performance level of the ESSD to use as the system disk.
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
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false
	//
	// > You cannot encrypt system disks when you create instances in Hong Kong Zone D or Singapore Zone A.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 30000
	Iops *int32 `json:"Iops,omitempty" xml:"Iops,omitempty"`
	// The ID of the KMS key to use for the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The performance level of the ESSD to be used as the system disk. Default value: PL0. Valid values:
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
	// The provisioned read/write IOPS of the ESSD AutoPL disk to use as data disk N. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}
	//
	// Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}
	//
	// > This parameter is available only if you set the Category parameter to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html) and [Modify the performance configurations of an ESSD AutoPL disk](https://help.aliyun.com/document_detail/413275.html).
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

func (s CreateLaunchTemplateVersionRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetIops() *int32 {
	return s.Iops
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetBurstingEnabled(v bool) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetCategory(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetDeleteWithInstance(v bool) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetDescription(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetDiskName(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetEncrypted(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetIops(v int32) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.Iops = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetKMSKeyId(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetPerformanceLevel(v string) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetProvisionedIops(v int64) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) SetSize(v int32) *CreateLaunchTemplateVersionRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateVersionRequestDataDisk struct {
	// The ID of the automatic snapshot policy to apply to data disk N.
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
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
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// 	- cloud_essd: ESSD
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
	// 	- Valid values if DataDisk.N.Category is set to cloud: 5 to 2000.
	//
	// 	- Valid values if DataDisk.N.Category is set to cloud_efficiency: 20 to 32768.
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud_ssd: 20 to 32768.
	//
	// 	- Valid values if you set DataDisk.N.Category to cloud_essd: vary based on the `DataDisk.N.PerformanceLevel` value.
	//
	//     	- Valid values if you set DataDisk.N.PerformanceLevel to PL0: 1 to 32768.
	//
	//     	- Valid values if you set DataDisk.N.PerformanceLevel to PL1: 20 to 32768.
	//
	//     	- Valid values if you set DataDisk.N.PerformanceLevel to PL2: 461 to 32768.
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
	// The ID of the snapshot to use to create data disk N. Valid values of N: 1 to 16. When `DataDisk.N.SnapshotId` is specified, `DataDisk.N.Size` is ignored. The data disk is created with the size of the specified snapshot.
	//
	// Use snapshots created on or after July 15, 2013. Otherwise, an error is returned and your request is rejected.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
}

func (s CreateLaunchTemplateVersionRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestDataDisk) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetDevice() *string {
	return s.Device
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetAutoSnapshotPolicyId(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetBurstingEnabled(v bool) *CreateLaunchTemplateVersionRequestDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetCategory(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.Category = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetDeleteWithInstance(v bool) *CreateLaunchTemplateVersionRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetDescription(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetDevice(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.Device = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetDiskName(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetEncrypted(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetKMSKeyId(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetPerformanceLevel(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetProvisionedIops(v int64) *CreateLaunchTemplateVersionRequestDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetSize(v int32) *CreateLaunchTemplateVersionRequestDataDisk {
	s.Size = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) SetSnapshotId(v string) *CreateLaunchTemplateVersionRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateVersionRequestImageOptions struct {
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

func (s CreateLaunchTemplateVersionRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestImageOptions) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *CreateLaunchTemplateVersionRequestImageOptions) SetLoginAsNonRoot(v bool) *CreateLaunchTemplateVersionRequestImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestImageOptions) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateVersionRequestNetworkInterface struct {
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
	// testNetworkInterfaceDescription
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
	// The name of the secondary ENI. The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// example:
	//
	// testNetworkInterfaceName
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
	// The primary private IP address of the secondary ENI. The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// example:
	//
	// ``192.168.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// The ID of the security group to which to assign the secondary ENI. The security groups of the secondary ENI and of the instance must belong to the same VPC. The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// >  You cannot specify both `NetworkInterface.N.SecurityGroupId` and `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of the security groups to which to assign the secondary ENI. The security groups and the secondary ENI must reside in the same VPC. The valid values of N in `SecurityGroupIds.N` vary based on the maximum number of security groups to which a secondary ENI can belong. For more information, see the "Security group limits" section in [Limits](https://help.aliyun.com/document_detail/25412.html). The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// >  You cannot specify both `NetworkInterface.N.SecurityGroupId` and `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The ID of the vSwitch to which to connect the secondary ENI. The instance and the secondary ENI must reside in the same zone of the same VPC, but they can be connected to different vSwitches. The value of N in `NetworkInterface.N` cannot be greater than 1.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s CreateLaunchTemplateVersionRequestNetworkInterface) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestNetworkInterface) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetDescription() *string {
	return s.Description
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetInstanceType() *string {
	return s.InstanceType
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetDeleteOnRelease(v bool) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.DeleteOnRelease = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetDescription(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.Description = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetInstanceType(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.InstanceType = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetNetworkInterfaceName(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.NetworkInterfaceName = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetNetworkInterfaceTrafficMode(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetPrimaryIpAddress(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.PrimaryIpAddress = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetSecurityGroupId(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.SecurityGroupId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetSecurityGroupIds(v []*string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.SecurityGroupIds = v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) SetVSwitchId(v string) *CreateLaunchTemplateVersionRequestNetworkInterface {
	s.VSwitchId = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestNetworkInterface) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateVersionRequestSecurityOptions struct {
	TrustedSystemMode *string `json:"TrustedSystemMode,omitempty" xml:"TrustedSystemMode,omitempty"`
}

func (s CreateLaunchTemplateVersionRequestSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestSecurityOptions) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestSecurityOptions) GetTrustedSystemMode() *string {
	return s.TrustedSystemMode
}

func (s *CreateLaunchTemplateVersionRequestSecurityOptions) SetTrustedSystemMode(v string) *CreateLaunchTemplateVersionRequestSecurityOptions {
	s.TrustedSystemMode = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type CreateLaunchTemplateVersionRequestTag struct {
	// The key of tag N to add to the ECS instance, disks, and primary ENI created based on the launch template version. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the ECS instance, disks, and primary ENI created based on the launch template version. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s CreateLaunchTemplateVersionRequestTag) String() string {
	return dara.Prettify(s)
}

func (s CreateLaunchTemplateVersionRequestTag) GoString() string {
	return s.String()
}

func (s *CreateLaunchTemplateVersionRequestTag) GetKey() *string {
	return s.Key
}

func (s *CreateLaunchTemplateVersionRequestTag) GetValue() *string {
	return s.Value
}

func (s *CreateLaunchTemplateVersionRequestTag) SetKey(v string) *CreateLaunchTemplateVersionRequestTag {
	s.Key = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestTag) SetValue(v string) *CreateLaunchTemplateVersionRequestTag {
	s.Value = &v
	return s
}

func (s *CreateLaunchTemplateVersionRequestTag) Validate() error {
	return dara.Validate(s)
}
