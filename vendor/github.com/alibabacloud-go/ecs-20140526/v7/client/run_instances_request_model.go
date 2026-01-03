// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunInstancesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCpuOptions(v *RunInstancesRequestCpuOptions) *RunInstancesRequest
	GetCpuOptions() *RunInstancesRequestCpuOptions
	SetHibernationOptions(v *RunInstancesRequestHibernationOptions) *RunInstancesRequest
	GetHibernationOptions() *RunInstancesRequestHibernationOptions
	SetPrivatePoolOptions(v *RunInstancesRequestPrivatePoolOptions) *RunInstancesRequest
	GetPrivatePoolOptions() *RunInstancesRequestPrivatePoolOptions
	SetSchedulerOptions(v *RunInstancesRequestSchedulerOptions) *RunInstancesRequest
	GetSchedulerOptions() *RunInstancesRequestSchedulerOptions
	SetSecurityOptions(v *RunInstancesRequestSecurityOptions) *RunInstancesRequest
	GetSecurityOptions() *RunInstancesRequestSecurityOptions
	SetSystemDisk(v *RunInstancesRequestSystemDisk) *RunInstancesRequest
	GetSystemDisk() *RunInstancesRequestSystemDisk
	SetAffinity(v string) *RunInstancesRequest
	GetAffinity() *string
	SetAmount(v int32) *RunInstancesRequest
	GetAmount() *int32
	SetArn(v []*RunInstancesRequestArn) *RunInstancesRequest
	GetArn() []*RunInstancesRequestArn
	SetAutoPay(v bool) *RunInstancesRequest
	GetAutoPay() *bool
	SetAutoReleaseTime(v string) *RunInstancesRequest
	GetAutoReleaseTime() *string
	SetAutoRenew(v bool) *RunInstancesRequest
	GetAutoRenew() *bool
	SetAutoRenewPeriod(v int32) *RunInstancesRequest
	GetAutoRenewPeriod() *int32
	SetClientToken(v string) *RunInstancesRequest
	GetClientToken() *string
	SetClockOptions(v *RunInstancesRequestClockOptions) *RunInstancesRequest
	GetClockOptions() *RunInstancesRequestClockOptions
	SetCreditSpecification(v string) *RunInstancesRequest
	GetCreditSpecification() *string
	SetDataDisk(v []*RunInstancesRequestDataDisk) *RunInstancesRequest
	GetDataDisk() []*RunInstancesRequestDataDisk
	SetDedicatedHostId(v string) *RunInstancesRequest
	GetDedicatedHostId() *string
	SetDeletionProtection(v bool) *RunInstancesRequest
	GetDeletionProtection() *bool
	SetDeploymentSetGroupNo(v int32) *RunInstancesRequest
	GetDeploymentSetGroupNo() *int32
	SetDeploymentSetId(v string) *RunInstancesRequest
	GetDeploymentSetId() *string
	SetDescription(v string) *RunInstancesRequest
	GetDescription() *string
	SetDryRun(v bool) *RunInstancesRequest
	GetDryRun() *bool
	SetHostName(v string) *RunInstancesRequest
	GetHostName() *string
	SetHostNames(v []*string) *RunInstancesRequest
	GetHostNames() []*string
	SetHpcClusterId(v string) *RunInstancesRequest
	GetHpcClusterId() *string
	SetHttpEndpoint(v string) *RunInstancesRequest
	GetHttpEndpoint() *string
	SetHttpPutResponseHopLimit(v int32) *RunInstancesRequest
	GetHttpPutResponseHopLimit() *int32
	SetHttpTokens(v string) *RunInstancesRequest
	GetHttpTokens() *string
	SetImageFamily(v string) *RunInstancesRequest
	GetImageFamily() *string
	SetImageId(v string) *RunInstancesRequest
	GetImageId() *string
	SetImageOptions(v *RunInstancesRequestImageOptions) *RunInstancesRequest
	GetImageOptions() *RunInstancesRequestImageOptions
	SetInstanceChargeType(v string) *RunInstancesRequest
	GetInstanceChargeType() *string
	SetInstanceName(v string) *RunInstancesRequest
	GetInstanceName() *string
	SetInstanceType(v string) *RunInstancesRequest
	GetInstanceType() *string
	SetInternetChargeType(v string) *RunInstancesRequest
	GetInternetChargeType() *string
	SetInternetMaxBandwidthIn(v int32) *RunInstancesRequest
	GetInternetMaxBandwidthIn() *int32
	SetInternetMaxBandwidthOut(v int32) *RunInstancesRequest
	GetInternetMaxBandwidthOut() *int32
	SetIoOptimized(v string) *RunInstancesRequest
	GetIoOptimized() *string
	SetIpv6Address(v []*string) *RunInstancesRequest
	GetIpv6Address() []*string
	SetIpv6AddressCount(v int32) *RunInstancesRequest
	GetIpv6AddressCount() *int32
	SetIsp(v string) *RunInstancesRequest
	GetIsp() *string
	SetKeyPairName(v string) *RunInstancesRequest
	GetKeyPairName() *string
	SetLaunchTemplateId(v string) *RunInstancesRequest
	GetLaunchTemplateId() *string
	SetLaunchTemplateName(v string) *RunInstancesRequest
	GetLaunchTemplateName() *string
	SetLaunchTemplateVersion(v int64) *RunInstancesRequest
	GetLaunchTemplateVersion() *int64
	SetMinAmount(v int32) *RunInstancesRequest
	GetMinAmount() *int32
	SetNetworkInterface(v []*RunInstancesRequestNetworkInterface) *RunInstancesRequest
	GetNetworkInterface() []*RunInstancesRequestNetworkInterface
	SetNetworkInterfaceQueueNumber(v int32) *RunInstancesRequest
	GetNetworkInterfaceQueueNumber() *int32
	SetNetworkOptions(v *RunInstancesRequestNetworkOptions) *RunInstancesRequest
	GetNetworkOptions() *RunInstancesRequestNetworkOptions
	SetOwnerAccount(v string) *RunInstancesRequest
	GetOwnerAccount() *string
	SetOwnerId(v int64) *RunInstancesRequest
	GetOwnerId() *int64
	SetPassword(v string) *RunInstancesRequest
	GetPassword() *string
	SetPasswordInherit(v bool) *RunInstancesRequest
	GetPasswordInherit() *bool
	SetPeriod(v int32) *RunInstancesRequest
	GetPeriod() *int32
	SetPeriodUnit(v string) *RunInstancesRequest
	GetPeriodUnit() *string
	SetPrivateDnsNameOptions(v *RunInstancesRequestPrivateDnsNameOptions) *RunInstancesRequest
	GetPrivateDnsNameOptions() *RunInstancesRequestPrivateDnsNameOptions
	SetPrivateIpAddress(v string) *RunInstancesRequest
	GetPrivateIpAddress() *string
	SetRamRoleName(v string) *RunInstancesRequest
	GetRamRoleName() *string
	SetRegionId(v string) *RunInstancesRequest
	GetRegionId() *string
	SetResourceGroupId(v string) *RunInstancesRequest
	GetResourceGroupId() *string
	SetResourceOwnerAccount(v string) *RunInstancesRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *RunInstancesRequest
	GetResourceOwnerId() *int64
	SetSecurityEnhancementStrategy(v string) *RunInstancesRequest
	GetSecurityEnhancementStrategy() *string
	SetSecurityGroupId(v string) *RunInstancesRequest
	GetSecurityGroupId() *string
	SetSecurityGroupIds(v []*string) *RunInstancesRequest
	GetSecurityGroupIds() []*string
	SetSpotDuration(v int32) *RunInstancesRequest
	GetSpotDuration() *int32
	SetSpotInterruptionBehavior(v string) *RunInstancesRequest
	GetSpotInterruptionBehavior() *string
	SetSpotPriceLimit(v float32) *RunInstancesRequest
	GetSpotPriceLimit() *float32
	SetSpotStrategy(v string) *RunInstancesRequest
	GetSpotStrategy() *string
	SetStorageSetId(v string) *RunInstancesRequest
	GetStorageSetId() *string
	SetStorageSetPartitionNumber(v int32) *RunInstancesRequest
	GetStorageSetPartitionNumber() *int32
	SetTag(v []*RunInstancesRequestTag) *RunInstancesRequest
	GetTag() []*RunInstancesRequestTag
	SetTenancy(v string) *RunInstancesRequest
	GetTenancy() *string
	SetUniqueSuffix(v bool) *RunInstancesRequest
	GetUniqueSuffix() *bool
	SetUserData(v string) *RunInstancesRequest
	GetUserData() *string
	SetVSwitchId(v string) *RunInstancesRequest
	GetVSwitchId() *string
	SetZoneId(v string) *RunInstancesRequest
	GetZoneId() *string
}

type RunInstancesRequest struct {
	CpuOptions         *RunInstancesRequestCpuOptions         `json:"CpuOptions,omitempty" xml:"CpuOptions,omitempty" type:"Struct"`
	HibernationOptions *RunInstancesRequestHibernationOptions `json:"HibernationOptions,omitempty" xml:"HibernationOptions,omitempty" type:"Struct"`
	PrivatePoolOptions *RunInstancesRequestPrivatePoolOptions `json:"PrivatePoolOptions,omitempty" xml:"PrivatePoolOptions,omitempty" type:"Struct"`
	SchedulerOptions   *RunInstancesRequestSchedulerOptions   `json:"SchedulerOptions,omitempty" xml:"SchedulerOptions,omitempty" type:"Struct"`
	SecurityOptions    *RunInstancesRequestSecurityOptions    `json:"SecurityOptions,omitempty" xml:"SecurityOptions,omitempty" type:"Struct"`
	SystemDisk         *RunInstancesRequestSystemDisk         `json:"SystemDisk,omitempty" xml:"SystemDisk,omitempty" type:"Struct"`
	// Specifies whether to associate an instance on a dedicated host with the dedicated host. Valid values:
	//
	// 	- default: does not associate the instance with the dedicated host. When you start an instance that was stopped in economical mode, the instance is automatically deployed to another dedicated host in the automatic deployment resource pool if the available resources of the original dedicated host are insufficient.
	//
	// 	- host: associates the instance with the dedicated host. When you start an instance that was stopped in economical mode, the instance remains on the original dedicated host. If the available resources of the original dedicated host are insufficient, the instance cannot be started.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Affinity *string `json:"Affinity,omitempty" xml:"Affinity,omitempty"`
	// The desired number of ECS instances that you want to create. Valid values: 1 to 100.
	//
	// The number of ECS instances that can be created varies based on the Amount and MinAmount values.
	//
	// 	- If you do not specify MinAmount, the RunInstances operation creates ECS instances based on the Amount value. If the available resources are insufficient to create the desired number of ECS instances, the RunInstances operation returns an error response and no ECS instances are created.
	//
	// 	- If you specify MinAmount, take note of the following items:
	//
	//     	- If the available resources are insufficient to create the minimum number of ECS instances, no ECS instances are created and the RunInstances operation returns an error response.
	//
	//     	- If the available resources are insufficient to create the desired number of ECS instances but are sufficient to create the minimum number of ECS instances, the RunInstances operation uses the available resources to create ECS instances and returns a success response. In this case, the number of ECS instances that can be created is less than the desired number of ECS instances.
	//
	//     	- If the available resources are sufficient to create the desired number of ECS instances, the RunInstances operation uses the available resources to create the desired number of ECS instances and returns a success response.
	//
	// Default value: 1.
	//
	// example:
	//
	// 3
	Amount *int32 `json:"Amount,omitempty" xml:"Amount,omitempty"`
	// >  This parameter is not publicly available.
	Arn []*RunInstancesRequestArn `json:"Arn,omitempty" xml:"Arn,omitempty" type:"Repeated"`
	// Specifies whether to automatically complete the payment for instance creation. Valid values:
	//
	// 	- true: The payment is automatically completed.
	//
	//     **
	//
	//     **Note*	- Make sure that your account balance is sufficient. Otherwise, your order becomes invalid and is canceled. If your account balance is insufficient, you can set `AutoPay` to `false` to generate an unpaid order. Then, you can log on to the ECS console to pay for the order.
	//
	// 	- false: An order is generated but no payment is made.
	//
	//     **
	//
	//     **Note*	- When `InstanceChargeType` is set to `PostPaid`, `AutoPay` cannot be set to `false`.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	AutoPay *bool `json:"AutoPay,omitempty" xml:"AutoPay,omitempty"`
	// The time when to automatically release the pay-as-you-go instance. Specify the time in the [ISO 8601 standard](https://help.aliyun.com/document_detail/25696.html) in the `yyyy-MM-ddTHH:mm:ssZ` format. The time must be in UTC.
	//
	// 	- If the value of seconds (`ss`) is not `00`, the start time is automatically rounded to the nearest minute based on the value of minutes (`mm`).
	//
	// 	- The specified time must be at least 30 minutes later than the current time.
	//
	// 	- The specified time can be at most three years later than the current time.
	//
	// example:
	//
	// 2018-01-01T12:05:00Z
	AutoReleaseTime *string `json:"AutoReleaseTime,omitempty" xml:"AutoReleaseTime,omitempty"`
	// Specifies whether to enable auto-renewal for the instance. This parameter is valid only when the `InstanceChargeType` parameter is set to `PrePaid`. Valid values:
	//
	// 	- true: enables auto-renewal.
	//
	// 	- false: does not enable auto-renewal.
	//
	// Default value: false.
	//
	// example:
	//
	// true
	AutoRenew *bool `json:"AutoRenew,omitempty" xml:"AutoRenew,omitempty"`
	// The auto-renewal period of the instance. Valid values:
	//
	// 	- Valid values when PeriodUnit is set to Week: 1, 2, and 3.
	//
	// 	- Valid values when PeriodUnit is set to Month: 1, 2, 3, 6, 12, 24, 36, 48, and 60.
	//
	// Default value: 1.
	//
	// example:
	//
	// 1
	AutoRenewPeriod *int32 `json:"AutoRenewPeriod,omitempty" xml:"AutoRenewPeriod,omitempty"`
	// The client token that is used to ensure the idempotence of the request. You can use the client to generate the token, but you must make sure that the token is unique among different requests. The token can contain only ASCII characters and cannot exceed 64 characters in length.***	- For more information, see [How to ensure idempotence](https://help.aliyun.com/document_detail/25693.html).
	//
	// example:
	//
	// 123e4567-e89b-12d3-a456-426655440000
	ClientToken  *string                          `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	ClockOptions *RunInstancesRequestClockOptions `json:"ClockOptions,omitempty" xml:"ClockOptions,omitempty" type:"Struct"`
	// The performance mode of the burstable instance. Valid values:
	//
	// 	- Standard: the standard mode. For more information, see the "Standard mode" section in [Overview of burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// 	- Unlimited: the unlimited mode. For more information, see the "Unlimited mode" section in [Burstable instances](https://help.aliyun.com/document_detail/59977.html).
	//
	// example:
	//
	// Standard
	CreditSpecification *string `json:"CreditSpecification,omitempty" xml:"CreditSpecification,omitempty"`
	// The data disks.
	DataDisk []*RunInstancesRequestDataDisk `json:"DataDisk,omitempty" xml:"DataDisk,omitempty" type:"Repeated"`
	// The ID of the dedicated host.
	//
	// You can call the [DescribeDedicatedHosts](https://help.aliyun.com/document_detail/134242.html) operation to query the list of dedicated host IDs.
	//
	// > Spot instances cannot be created on dedicated hosts. If you specify DedicatedHostId, SpotStrategy and SpotPriceLimit are automatically ignored.
	//
	// example:
	//
	// dh-bp67acfmxazb4p****
	DedicatedHostId *string `json:"DedicatedHostId,omitempty" xml:"DedicatedHostId,omitempty"`
	// Specifies whether to enable release protection for the instance. This parameter determines whether you can use the ECS console or call the [DeleteInstance](https://help.aliyun.com/document_detail/25507.html) operation to release the instance. Valid values:
	//
	// 	- true: enables release protection for the instance.
	//
	// 	- false: disables release protection for the instance.
	//
	// Default value: false.
	//
	// > This parameter is applicable to only pay-as-you-go instances. It can protect instances against manual releases, but not against automatic releases.
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
	// The description of the instance. The description must be 2 to 256 characters in length, and cannot start with `http://` or `https://`.
	//
	// example:
	//
	// Instance_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// Specifies whether to check the validity of the request without actually making the request. Default value: false. Valid values:
	//
	// 	- true: The validity of the request is checked but the request is not made. Check items include whether required parameters are specified, the request format, service limits, and available ECS resources. If the check fails, the corresponding error code is returned. If the check succeeds, the `DryRunOperation` error code is returned.
	//
	// 	- false: The validity of the request is checked, and the request is made if the check succeeds.
	//
	// example:
	//
	// false
	DryRun *bool `json:"DryRun,omitempty" xml:"DryRun,omitempty"`
	// The hostname of the instance. Take note of the following items:
	//
	// 	- The hostname cannot start or end with a period (.) or hyphen (-). It cannot contain consecutive periods (.) or hyphens (-).
	//
	// 	- For Windows instances, the hostname must be 2 to 15 characters in length and cannot contain periods (.) or contain only digits. It can contain letters, digits, and hyphens (-).
	//
	// 	- For instances that run other operating systems such as Linux, take note of the following items:
	//
	//     	- The hostname must be 2 to 64 characters in length. You can use periods (.) to separate a hostname into multiple segments. Each segment can contain letters, digits, and hyphens (-).
	//
	//     	- You can use the `${instance_id}` placeholder to pass instance IDs into the hostname specified by `HostName`. For example, if you set `HostName` to k8s-${instance_id} and the instance is assigned an ID of `i-123abc****`, the hostname of the instance is `k8s-i-123abc****`.
	//
	// When you create multiple instances, you can perform the following operations:
	//
	// 	- Batch configure sequential hostnames for the instances. For more information, see [Batch configure sequential names or hostnames for multiple instances](https://help.aliyun.com/document_detail/196048.html).
	//
	// 	- Use the `HostNames.N` parameter to configure different hostnames for instances. You cannot specify both the `HostName` and `HostNames.N` parameters.
	//
	// example:
	//
	// k8s-node-[1,4]-ecshost
	HostName *string `json:"HostName,omitempty" xml:"HostName,omitempty"`
	// The hostname of instance N. You can use this parameter to specify different hostnames for multiple instances.
	//
	// example:
	//
	// ecs-host-01
	HostNames []*string `json:"HostNames,omitempty" xml:"HostNames,omitempty" type:"Repeated"`
	// The ID of the high performance computing (HPC) cluster to which the instance belongs.
	//
	// This parameter is required when you create instances of a Supper Computing Cluster (SCC) instance type. For information about how to create an HPC cluster, see [CreateHpcCluster](https://help.aliyun.com/document_detail/109138.html).
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
	// > For more information about instance metadata, see [Overview of ECS instance metadata](https://help.aliyun.com/document_detail/49122.html).
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
	// Specifies whether to forcefully use the security-enhanced mode (IMDSv2) to access instance metadata. Valid values:
	//
	// 	- optional: does not forcefully use the security-enhanced mode (IMDSv2).
	//
	// 	- required: forcefully uses the security-enhanced mode (IMDSv2). After you set this parameter to required, you cannot access instance metadata in normal mode.
	//
	// Default value: optional.
	//
	// > For more information about the modes of accessing instance metadata, see [Access mode of instance metadata](https://help.aliyun.com/document_detail/150575.html).
	//
	// example:
	//
	// optional
	HttpTokens *string `json:"HttpTokens,omitempty" xml:"HttpTokens,omitempty"`
	// The name of the image family. You can set this parameter to obtain the latest available custom image from the specified image family to create instances.
	//
	// The name must be 2 to 128 characters in length. The name cannot start with a digit, a special character, http://, or https://. The name can contain letters, digits, periods (.), underscores (_), hyphens (-), and colons (:).
	//
	// Take note of the following items:
	//
	// 	- If you specify `ImageId`, you cannot specify ImageFamily.
	//
	// 	- If you do not specify `ImageId` but use `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template that has `ImageId` specified, you cannot specify ImageFamily.
	//
	// 	- If you do not specify `ImageId` but use `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template that does not have `ImageId` specified, you can specify ImageFamily.
	//
	// 	- If you do not specify `ImageId`, `LaunchTemplateId`, or `LaunchTemplateName`, you can specify ImageFamily.
	//
	// >  For information about image families that are associated with Alibaba Cloud official images, see [Overview of public images](https://help.aliyun.com/document_detail/108393.html).
	//
	// example:
	//
	// hangzhou-daily-update
	ImageFamily *string `json:"ImageFamily,omitempty" xml:"ImageFamily,omitempty"`
	// The ID of the image. You can call the [DescribeImages](https://help.aliyun.com/document_detail/25534.html) operation to query available images. If you do not use `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template and do not set `ImageFamily` to obtain the latest available custom image from a specified image family, you must specify `ImageId`.
	//
	// example:
	//
	// aliyun_2_1903_x64_20G_alibase_20200324.vhd
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// Details about the image options.
	ImageOptions *RunInstancesRequestImageOptions `json:"ImageOptions,omitempty" xml:"ImageOptions,omitempty" type:"Struct"`
	// The billing method of the instance. Valid values:
	//
	// 	- PrePaid: subscription
	//
	// 	- PostPaid: pay-as-you-go
	//
	// Default value: PostPaid.
	//
	// If you set this parameter to PrePaid, make sure that your account has sufficient balance or credit. Otherwise, an `InvalidPayMethod` error is returned.
	//
	// example:
	//
	// PrePaid
	InstanceChargeType *string `json:"InstanceChargeType,omitempty" xml:"InstanceChargeType,omitempty"`
	// The name of the ECS instance. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-). The default value of this parameter is the `InstanceId` value.
	//
	// When you batch create instances, you can batch configure sequential names for the instances. The sequential names can contain brackets ([ ]) and commas (,). For more information, see [Batch configure sequential names or hostnames for multiple instances](https://help.aliyun.com/document_detail/196048.html).
	//
	// example:
	//
	// k8s-node-[1,4]-alibabacloud
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// The instance type. If you do not use `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template, you must set the `InstanceType` parameter.
	//
	// 	- Select an instance type. See [Instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the performance data of an instance type, or see [Best practices for instance type selection](https://help.aliyun.com/document_detail/58291.html) to learn about how to select instance types.
	//
	// 	- Query available resources. Call the [DescribeAvailableResource](https://help.aliyun.com/document_detail/66186.html) operation to query available resources in a specific region or zone.
	//
	// example:
	//
	// ecs.g6.large
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The billing method for network usage. Valid values:
	//
	// 	- PayByBandwidth: pay-by-bandwidth
	//
	// 	- PayByTraffic: pay-by-traffic
	//
	// Default value: PayByTraffic.
	//
	// > When the **pay-by-traffic*	- billing method for network usage is used, the maximum inbound and outbound bandwidths are used as the upper limits of bandwidths instead of guaranteed performance specifications. In scenarios where demand outstrips resource supplies, these maximum bandwidth values may not be reached. If you want guaranteed bandwidths for your instance, use the **pay-by-bandwidth*	- billing method for network usage.
	//
	// example:
	//
	// PayByTraffic
	InternetChargeType *string `json:"InternetChargeType,omitempty" xml:"InternetChargeType,omitempty"`
	// The maximum inbound public bandwidth. Unit: Mbit/s. Valid values:
	//
	// 	- When the purchased outbound public bandwidth is less than or equal to 10 Mbit/s, the valid values of InternetMaxBandwidthIn are 1 to 10, and the default value is 10.
	//
	// 	- When the purchased outbound public bandwidth is greater than 10 Mbit/s, the valid values of this parameter are 1 to the `InternetMaxBandwidthOut` value and the default value is the `InternetMaxBandwidthOut` value.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthIn *int32 `json:"InternetMaxBandwidthIn,omitempty" xml:"InternetMaxBandwidthIn,omitempty"`
	// The maximum outbound public bandwidth. Unit: Mbit/s. Valid values: 0 to 100.
	//
	// Default value: 0.
	//
	// example:
	//
	// 10
	InternetMaxBandwidthOut *int32 `json:"InternetMaxBandwidthOut,omitempty" xml:"InternetMaxBandwidthOut,omitempty"`
	// Specifies whether the instance is I/O optimized. For instances of [retired instance types](https://help.aliyun.com/document_detail/55263.html), the default value is none. For instances of other instance types, the default value is optimized. Valid values:
	//
	// 	- none: The instance is not I/O optimized.
	//
	// 	- optimized: The instance is I/O optimized.
	//
	// example:
	//
	// optimized
	IoOptimized *string `json:"IoOptimized,omitempty" xml:"IoOptimized,omitempty"`
	// IPv6 address N to be assigned to the primary ENI. Valid values of N: 1 to 10.
	//
	// Example: `Ipv6Address.1=2001:db8:1234:1a00::***`.
	//
	// Take note of the following items:
	//
	// 	- If the `Ipv6Address.N` parameter is specified, you must set the `Amount` parameter to 1 and leave the `Ipv6AddressCount` parameter empty.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot set `Ipv6Addresses.N` or `Ipv6AddressCount` and must set `NetworkInterface.N.Ipv6Addresses.N` or `NetworkInterface.N.Ipv6AddressCount`.
	//
	// example:
	//
	// Ipv6Address.1=2001:db8:1234:1a00::***
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10.
	//
	// Take note of the following items:
	//
	// 	- You cannot specify both the `Ipv6Addresses.N` and `Ipv6AddressCount` parameters.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot specify `Ipv6Address.N` or `Ipv6AddressCount` but can specify `NetworkInterface.N.Ipv6Address.N` or `NetworkInterface.N.Ipv6AddressCount`.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int32 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// > This parameter is in invitational preview and is unavailable.
	//
	// example:
	//
	// null
	Isp *string `json:"Isp,omitempty" xml:"Isp,omitempty"`
	// The name of the key pair.
	//
	// > For Windows instances, this parameter is ignored. This parameter is empty by default. The `Password` parameter takes effect even if the KeyPairName parameter is specified.
	//
	// example:
	//
	// KeyPair_Name
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The ID of the launch template. For more information, call the [DescribeLaunchTemplates](https://help.aliyun.com/document_detail/73759.html) operation.
	//
	// To use a launch template to create an instance, you must use the `LaunchTemplateId` or `LaunchTemplateName` parameter to specify the launch template.
	//
	// example:
	//
	// lt-bp1apo0bbbkuy0rj****
	LaunchTemplateId *string `json:"LaunchTemplateId,omitempty" xml:"LaunchTemplateId,omitempty"`
	// The name of the launch template.
	//
	// To use a launch template to create an instance, you must use the `LaunchTemplateId` or `LaunchTemplateName` parameter to specify the launch template.
	//
	// example:
	//
	// LaunchTemplate_Name
	LaunchTemplateName *string `json:"LaunchTemplateName,omitempty" xml:"LaunchTemplateName,omitempty"`
	// The version of the launch template. If you set the `LaunchTemplateId` or `LaunchTemplateName` parameter but do not set the version number of the launch template, the default template version is used.
	//
	// example:
	//
	// 3
	LaunchTemplateVersion *int64 `json:"LaunchTemplateVersion,omitempty" xml:"LaunchTemplateVersion,omitempty"`
	// The minimum number of ECS instances that you want to create. Valid values: 1 to 100.
	//
	// The number of ECS instances that can be created varies based on the Amount and MinAmount values.
	//
	// 	- If you do not specify MinAmount, the RunInstances operation creates ECS instances based on the Amount value. If the available resources are insufficient to create the desired number of ECS instances, the RunInstances operation returns an error response and no ECS instances are created.
	//
	// 	- If you specify MinAmount, take note of the following items:
	//
	//     	- If the available resources are insufficient to create the minimum number of ECS instances, no ECS instances are created and the RunInstances operation returns an error response.
	//
	//     	- If the available resources are insufficient to create the desired number of ECS instances but are sufficient to create the minimum number of ECS instances, the RunInstances operation uses the available resources to create ECS instances and returns a success response. In this case, the number of ECS instances that can be created is less than the desired number of ECS instances.
	//
	//     	- If the available resources are sufficient to create the desired number of ECS instances, the RunInstances operation uses the available resources to create the desired number of ECS instances and returns a success response.
	//
	// example:
	//
	// 2
	MinAmount *int32 `json:"MinAmount,omitempty" xml:"MinAmount,omitempty"`
	// The information of the elastic network interfaces (ENIs).
	NetworkInterface []*RunInstancesRequestNetworkInterface `json:"NetworkInterface,omitempty" xml:"NetworkInterface,omitempty" type:"Repeated"`
	// The number of queues supported by the primary ENI. Take note of the following items:
	//
	// 	- The value of this parameter cannot exceed the maximum number of queues per ENI allowed for the instance type.
	//
	// 	- The total number of queues for all ENIs on the instance cannot exceed the queue quota for the instance type. To query the maximum number of queues per ENI and the queue quota for an instance type, you can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation to query the `MaximumQueueNumberPerEni` and `TotalEniQueueQuantity` values.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot specify `NetworkInterfaceQueueNumber` but can specify `NetworkInterface.N.QueueNumber`.
	//
	// example:
	//
	// 8
	NetworkInterfaceQueueNumber *int32 `json:"NetworkInterfaceQueueNumber,omitempty" xml:"NetworkInterfaceQueueNumber,omitempty"`
	// Details about network options.
	NetworkOptions *RunInstancesRequestNetworkOptions `json:"NetworkOptions,omitempty" xml:"NetworkOptions,omitempty" type:"Struct"`
	OwnerAccount   *string                            `json:"OwnerAccount,omitempty" xml:"OwnerAccount,omitempty"`
	OwnerId        *int64                             `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The password of the instance. The password must be 8 to 30 characters in length and contain at least three of the following character types: uppercase letters, lowercase letters, digits, and special characters. Special characters include:
	//
	//     ()`~!@#$%^&*-_+=|{}[]:;\\"<>,.?/
	//
	// For Windows instances, the password cannot start with a forward slash (/).
	//
	// > If the `Password` parameter is specified, we recommend that you send requests over HTTPS to prevent password leaks.
	//
	// example:
	//
	// EcsV587!
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// Specifies whether to use the password preset in the image. Valid values:
	//
	// 	- true: uses the preset password.
	//
	// 	- false: does not use the preset password.
	//
	// Default value: false.
	//
	// > If you set this parameter to true, make sure that you leave the Password parameter empty and the selected image has a preset password.
	//
	// example:
	//
	// false
	PasswordInherit *bool `json:"PasswordInherit,omitempty" xml:"PasswordInherit,omitempty"`
	// The subscription period of the instance. The unit is specified by the `PeriodUnit` parameter. This parameter is valid and required only when `InstanceChargeType` is set to `PrePaid`. If the `DedicatedHostId` parameter is specified, the value of Period must not exceed the subscription period of the specified dedicated host. Valid values:
	//
	// 	- Valid values when PeriodUnit is set to Week: 1, 2, 3, and 4.
	//
	// 	- Valid values when PeriodUnit is set to Month: 1, 2, 3, 4, 5, 6, 7, 8, 9, 12, 24, 36, 48, and 60.
	//
	// example:
	//
	// 1
	Period *int32 `json:"Period,omitempty" xml:"Period,omitempty"`
	// The unit of the subscription period. Default value: Month. Valid values:
	//
	// 	- Week
	//
	// 	- Month
	//
	// example:
	//
	// Month
	PeriodUnit *string `json:"PeriodUnit,omitempty" xml:"PeriodUnit,omitempty"`
	// The private domain name options of the instance.
	//
	// For information about the resolution of ECS private domain names, see [ECS private DNS resolution](https://help.aliyun.com/document_detail/2844797.html).
	PrivateDnsNameOptions *RunInstancesRequestPrivateDnsNameOptions `json:"PrivateDnsNameOptions,omitempty" xml:"PrivateDnsNameOptions,omitempty" type:"Struct"`
	// The private IP address to assign to the instance. To assign a private IP address to an instance that resides in a VPC, make sure that the IP address is an idle IP address within the CIDR block of the vSwitch specified by `VSwitchId`.
	//
	// Take note of the following items:
	//
	// 	- If `PrivateIpAddress` is specified, take note of the following items:
	//
	//     	- If `Amount` is set to 1, a single instance is created and the specified private IP address is assigned to the instance.
	//
	//     	- If `Amount` is set to a numeric value greater than 1, the specified number of instances are created and consecutive private IP addresses starting from the specified one are assigned to the instances. In this case, you cannot specify parameters that start with `NetworkInterface.N` to attach secondary ENIs to the instances.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot specify `PrivateIpAddress` but can specify `NetworkInterface.N.PrimaryIpAddress`.
	//
	// >  The first IP address and last three IP addresses of each vSwitch CIDR block are reserved. You cannot specify the IP addresses. For example, if a vSwitch CIDR block is 192.168.1.0/24, the IP addresses 192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255 are reserved.
	//
	// example:
	//
	// ``10.1.**.**``
	PrivateIpAddress *string `json:"PrivateIpAddress,omitempty" xml:"PrivateIpAddress,omitempty"`
	// The name of the Resource Access Management (RAM) role. You can call the [ListRoles](https://help.aliyun.com/document_detail/28713.html) operation provided by RAM to query the instance RAM roles that you created.
	//
	// example:
	//
	// RAM_Name
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
	// 	- Active: enables security hardening. This value is applicable only to public images.
	//
	// 	- Deactive: does not enable security hardening. This value is applicable to all images.
	//
	// example:
	//
	// Active
	SecurityEnhancementStrategy *string `json:"SecurityEnhancementStrategy,omitempty" xml:"SecurityEnhancementStrategy,omitempty"`
	// The ID of the security group to which you want to assign the instance. Instances in the same security group can communicate with each other. The maximum number of instances allowed in a security group varies based on the type of the security group. For more information, see the "Security group limits" section in [Limits and quotas](~~25412#SecurityGroupQuota~~).
	//
	// >  The network type of the new instance is the same as the network type of the security group specified by `SecurityGroupId`. For example, if the specified security group is of the VPC type, the new instance is also of the VPC type and you must specify `VSwitchId`.
	//
	// If you do not use `LaunchTemplateId` or `LaunchTemplateName` to specify a launch template, you must specify a security group ID. When you specify this parameter, take note of the following items:
	//
	// 	- You can set `SecurityGroupId` to specify a single security group or set `SecurityGroupIds.N` to specify one or more security groups. However, you cannot specify both `SecurityGroupId` and `SecurityGroupIds.N` in the same request.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot specify `SecurityGroupId` or `SecurityGroupIds.N` but can specify `NetworkInterface.N.SecurityGroupId` or `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of security groups to which to assign the instance. The valid values of N vary based on the maximum number of security groups to which an instance can belong. For more information, see the [Security group limits](https://help.aliyun.com/document_detail/101348.html) section of the "Limits" topic.
	//
	// When you specify this parameter, take note of the following items:
	//
	// 	- You cannot specify both `SecurityGroupId` and `SecurityGroupIds.N` in the same request.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot specify `SecurityGroupId` or `SecurityGroupIds.N` but can specify `NetworkInterface.N.SecurityGroupId` or `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// The protection period of the spot instance. Unit: hours. Valid values:
	//
	// 	- 1: After a spot instance is created, Alibaba Cloud ensures that the instance is not automatically released within 1 hour. After the 1-hour protection period ends, the system compares the bid price with the market price and checks the resource inventory to determine whether to retain or release the instance.
	//
	// 	- 0: After a spot instance is created, Alibaba Cloud does not ensure that the instance can run for one hour. The system compares the biding price with the market prices and checks the resource inventory to determine whether to retain or release the instance.
	//
	// Default value: 1.
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
	// The maximum hourly price of the instance. The value is accurate to three decimal places. This parameter is valid only when the `SpotStrategy` parameter is set to `SpotWithPriceLimit`.
	//
	// example:
	//
	// 0.97
	SpotPriceLimit *float32 `json:"SpotPriceLimit,omitempty" xml:"SpotPriceLimit,omitempty"`
	// The bidding policy for the pay-as-you-go instance. This parameter is valid only when the `InstanceChargeType` parameter is set to `PostPaid`. Valid values:
	//
	// 	- NoSpot: The instance is created as a pay-as-you-go instance.
	//
	// 	- SpotWithPriceLimit: The instance is created as a spot instance with a user-defined maximum hourly price.
	//
	// 	- SpotAsPriceGo: The instance is created as a spot instance for which the market price at the time of purchase is automatically used as the bid price.
	//
	// Default value: NoSpot.
	//
	// example:
	//
	// NoSpot
	SpotStrategy *string `json:"SpotStrategy,omitempty" xml:"SpotStrategy,omitempty"`
	// The ID of the storage set.
	//
	// example:
	//
	// ss-bp67acfmxazb4p****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set. Valid values: integers greater than or equal to 1.
	//
	// example:
	//
	// 2
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The tags to add to the instance, disks, and primary ENI.
	Tag []*RunInstancesRequestTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
	// Specifies whether to create the instance on a dedicated host. Valid values:
	//
	// 	- default: creates the instance on a non-dedicated host.
	//
	// 	- host: creates the instance on a dedicated host. If you do not set the `DedicatedHostId` parameter, Alibaba Cloud selects a dedicated host for the instance.
	//
	// Default value: default.
	//
	// example:
	//
	// default
	Tenancy *string `json:"Tenancy,omitempty" xml:"Tenancy,omitempty"`
	// Specifies whether to automatically append incremental suffixes to the hostname specified by the `HostName` parameter and to the instance name specified by the `InstanceName` parameter when you batch create instances. The incremental suffixes can range from 001 to 999. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// Default value: false.
	//
	// When the `HostName` or `InstanceName` value is set in the `name_prefix[begin_number,bits]` format without `name_suffix`, the `UniqueSuffix` parameter does not take effect. The names are sorted in the specified sequence.
	//
	// For more information, see [Batch configure sequential names or hostnames for multiple instances](https://help.aliyun.com/document_detail/196048.html).
	//
	// example:
	//
	// true
	UniqueSuffix *bool `json:"UniqueSuffix,omitempty" xml:"UniqueSuffix,omitempty"`
	// The user data of the instance. You must specify Base64-encoded data. The instance user data cannot exceed 32 KB in size before Base64 encoding.
	//
	// For information about the limits, formats, and running frequencies of instance user data, see [Instance user data](https://help.aliyun.com/document_detail/49121.html).
	//
	// >  To ensure security, we recommend that you do not use plaintext to pass in confidential information, such as passwords or private keys, as user data. If you need to pass in confidential information, we recommend that you encrypt and encode the information in Base64 and then decode and decrypt the information in the same manner in the instance.
	//
	// example:
	//
	// ZWNobyBoZWxsbyBlY3Mh
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The ID of the vSwitch to which to connect to the instance. You must set this parameter when you create an instance of the VPC type. The specified vSwitch and security group must belong to the same VPC. You can call the [DescribeVSwitches](https://help.aliyun.com/document_detail/35748.html) operation to query available vSwitches.
	//
	// Take note of the following items:
	//
	// 	- If you specify the `VSwitchId` parameter, the zone specified by the `ZoneId` parameter must be the zone where the specified vSwitch is located. You can also leave the `ZoneId` parameter empty. Then, the system selects the zone where the specified vSwitch resides.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you cannot specify `VSwitchId` but can specify `NetworkInterface.N.VSwitchId`.
	//
	// example:
	//
	// vsw-bp1s5fnvk4gn2tws0****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The ID of the zone in which to create the instance. You can call the [DescribeZones](https://help.aliyun.com/document_detail/25610.html) operation to query the most recent zone list.
	//
	// > If you specify the `VSwitchId` parameter, the zone specified by the `ZoneId` parameter must be the zone where the vSwitch is located. You can also leave the `ZoneId` parameter empty. Then, the system selects the zone where the specified vSwitch is located.
	//
	// This parameter is empty by default.
	//
	// example:
	//
	// cn-hangzhou-g
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s RunInstancesRequest) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequest) GoString() string {
	return s.String()
}

func (s *RunInstancesRequest) GetCpuOptions() *RunInstancesRequestCpuOptions {
	return s.CpuOptions
}

func (s *RunInstancesRequest) GetHibernationOptions() *RunInstancesRequestHibernationOptions {
	return s.HibernationOptions
}

func (s *RunInstancesRequest) GetPrivatePoolOptions() *RunInstancesRequestPrivatePoolOptions {
	return s.PrivatePoolOptions
}

func (s *RunInstancesRequest) GetSchedulerOptions() *RunInstancesRequestSchedulerOptions {
	return s.SchedulerOptions
}

func (s *RunInstancesRequest) GetSecurityOptions() *RunInstancesRequestSecurityOptions {
	return s.SecurityOptions
}

func (s *RunInstancesRequest) GetSystemDisk() *RunInstancesRequestSystemDisk {
	return s.SystemDisk
}

func (s *RunInstancesRequest) GetAffinity() *string {
	return s.Affinity
}

func (s *RunInstancesRequest) GetAmount() *int32 {
	return s.Amount
}

func (s *RunInstancesRequest) GetArn() []*RunInstancesRequestArn {
	return s.Arn
}

func (s *RunInstancesRequest) GetAutoPay() *bool {
	return s.AutoPay
}

func (s *RunInstancesRequest) GetAutoReleaseTime() *string {
	return s.AutoReleaseTime
}

func (s *RunInstancesRequest) GetAutoRenew() *bool {
	return s.AutoRenew
}

func (s *RunInstancesRequest) GetAutoRenewPeriod() *int32 {
	return s.AutoRenewPeriod
}

func (s *RunInstancesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *RunInstancesRequest) GetClockOptions() *RunInstancesRequestClockOptions {
	return s.ClockOptions
}

func (s *RunInstancesRequest) GetCreditSpecification() *string {
	return s.CreditSpecification
}

func (s *RunInstancesRequest) GetDataDisk() []*RunInstancesRequestDataDisk {
	return s.DataDisk
}

func (s *RunInstancesRequest) GetDedicatedHostId() *string {
	return s.DedicatedHostId
}

func (s *RunInstancesRequest) GetDeletionProtection() *bool {
	return s.DeletionProtection
}

func (s *RunInstancesRequest) GetDeploymentSetGroupNo() *int32 {
	return s.DeploymentSetGroupNo
}

func (s *RunInstancesRequest) GetDeploymentSetId() *string {
	return s.DeploymentSetId
}

func (s *RunInstancesRequest) GetDescription() *string {
	return s.Description
}

func (s *RunInstancesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *RunInstancesRequest) GetHostName() *string {
	return s.HostName
}

func (s *RunInstancesRequest) GetHostNames() []*string {
	return s.HostNames
}

func (s *RunInstancesRequest) GetHpcClusterId() *string {
	return s.HpcClusterId
}

func (s *RunInstancesRequest) GetHttpEndpoint() *string {
	return s.HttpEndpoint
}

func (s *RunInstancesRequest) GetHttpPutResponseHopLimit() *int32 {
	return s.HttpPutResponseHopLimit
}

func (s *RunInstancesRequest) GetHttpTokens() *string {
	return s.HttpTokens
}

func (s *RunInstancesRequest) GetImageFamily() *string {
	return s.ImageFamily
}

func (s *RunInstancesRequest) GetImageId() *string {
	return s.ImageId
}

func (s *RunInstancesRequest) GetImageOptions() *RunInstancesRequestImageOptions {
	return s.ImageOptions
}

func (s *RunInstancesRequest) GetInstanceChargeType() *string {
	return s.InstanceChargeType
}

func (s *RunInstancesRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *RunInstancesRequest) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RunInstancesRequest) GetInternetChargeType() *string {
	return s.InternetChargeType
}

func (s *RunInstancesRequest) GetInternetMaxBandwidthIn() *int32 {
	return s.InternetMaxBandwidthIn
}

func (s *RunInstancesRequest) GetInternetMaxBandwidthOut() *int32 {
	return s.InternetMaxBandwidthOut
}

func (s *RunInstancesRequest) GetIoOptimized() *string {
	return s.IoOptimized
}

func (s *RunInstancesRequest) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *RunInstancesRequest) GetIpv6AddressCount() *int32 {
	return s.Ipv6AddressCount
}

func (s *RunInstancesRequest) GetIsp() *string {
	return s.Isp
}

func (s *RunInstancesRequest) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *RunInstancesRequest) GetLaunchTemplateId() *string {
	return s.LaunchTemplateId
}

func (s *RunInstancesRequest) GetLaunchTemplateName() *string {
	return s.LaunchTemplateName
}

func (s *RunInstancesRequest) GetLaunchTemplateVersion() *int64 {
	return s.LaunchTemplateVersion
}

func (s *RunInstancesRequest) GetMinAmount() *int32 {
	return s.MinAmount
}

func (s *RunInstancesRequest) GetNetworkInterface() []*RunInstancesRequestNetworkInterface {
	return s.NetworkInterface
}

func (s *RunInstancesRequest) GetNetworkInterfaceQueueNumber() *int32 {
	return s.NetworkInterfaceQueueNumber
}

func (s *RunInstancesRequest) GetNetworkOptions() *RunInstancesRequestNetworkOptions {
	return s.NetworkOptions
}

func (s *RunInstancesRequest) GetOwnerAccount() *string {
	return s.OwnerAccount
}

func (s *RunInstancesRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *RunInstancesRequest) GetPassword() *string {
	return s.Password
}

func (s *RunInstancesRequest) GetPasswordInherit() *bool {
	return s.PasswordInherit
}

func (s *RunInstancesRequest) GetPeriod() *int32 {
	return s.Period
}

func (s *RunInstancesRequest) GetPeriodUnit() *string {
	return s.PeriodUnit
}

func (s *RunInstancesRequest) GetPrivateDnsNameOptions() *RunInstancesRequestPrivateDnsNameOptions {
	return s.PrivateDnsNameOptions
}

func (s *RunInstancesRequest) GetPrivateIpAddress() *string {
	return s.PrivateIpAddress
}

func (s *RunInstancesRequest) GetRamRoleName() *string {
	return s.RamRoleName
}

func (s *RunInstancesRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *RunInstancesRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *RunInstancesRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *RunInstancesRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *RunInstancesRequest) GetSecurityEnhancementStrategy() *string {
	return s.SecurityEnhancementStrategy
}

func (s *RunInstancesRequest) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RunInstancesRequest) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *RunInstancesRequest) GetSpotDuration() *int32 {
	return s.SpotDuration
}

func (s *RunInstancesRequest) GetSpotInterruptionBehavior() *string {
	return s.SpotInterruptionBehavior
}

func (s *RunInstancesRequest) GetSpotPriceLimit() *float32 {
	return s.SpotPriceLimit
}

func (s *RunInstancesRequest) GetSpotStrategy() *string {
	return s.SpotStrategy
}

func (s *RunInstancesRequest) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *RunInstancesRequest) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *RunInstancesRequest) GetTag() []*RunInstancesRequestTag {
	return s.Tag
}

func (s *RunInstancesRequest) GetTenancy() *string {
	return s.Tenancy
}

func (s *RunInstancesRequest) GetUniqueSuffix() *bool {
	return s.UniqueSuffix
}

func (s *RunInstancesRequest) GetUserData() *string {
	return s.UserData
}

func (s *RunInstancesRequest) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RunInstancesRequest) GetZoneId() *string {
	return s.ZoneId
}

func (s *RunInstancesRequest) SetCpuOptions(v *RunInstancesRequestCpuOptions) *RunInstancesRequest {
	s.CpuOptions = v
	return s
}

func (s *RunInstancesRequest) SetHibernationOptions(v *RunInstancesRequestHibernationOptions) *RunInstancesRequest {
	s.HibernationOptions = v
	return s
}

func (s *RunInstancesRequest) SetPrivatePoolOptions(v *RunInstancesRequestPrivatePoolOptions) *RunInstancesRequest {
	s.PrivatePoolOptions = v
	return s
}

func (s *RunInstancesRequest) SetSchedulerOptions(v *RunInstancesRequestSchedulerOptions) *RunInstancesRequest {
	s.SchedulerOptions = v
	return s
}

func (s *RunInstancesRequest) SetSecurityOptions(v *RunInstancesRequestSecurityOptions) *RunInstancesRequest {
	s.SecurityOptions = v
	return s
}

func (s *RunInstancesRequest) SetSystemDisk(v *RunInstancesRequestSystemDisk) *RunInstancesRequest {
	s.SystemDisk = v
	return s
}

func (s *RunInstancesRequest) SetAffinity(v string) *RunInstancesRequest {
	s.Affinity = &v
	return s
}

func (s *RunInstancesRequest) SetAmount(v int32) *RunInstancesRequest {
	s.Amount = &v
	return s
}

func (s *RunInstancesRequest) SetArn(v []*RunInstancesRequestArn) *RunInstancesRequest {
	s.Arn = v
	return s
}

func (s *RunInstancesRequest) SetAutoPay(v bool) *RunInstancesRequest {
	s.AutoPay = &v
	return s
}

func (s *RunInstancesRequest) SetAutoReleaseTime(v string) *RunInstancesRequest {
	s.AutoReleaseTime = &v
	return s
}

func (s *RunInstancesRequest) SetAutoRenew(v bool) *RunInstancesRequest {
	s.AutoRenew = &v
	return s
}

func (s *RunInstancesRequest) SetAutoRenewPeriod(v int32) *RunInstancesRequest {
	s.AutoRenewPeriod = &v
	return s
}

func (s *RunInstancesRequest) SetClientToken(v string) *RunInstancesRequest {
	s.ClientToken = &v
	return s
}

func (s *RunInstancesRequest) SetClockOptions(v *RunInstancesRequestClockOptions) *RunInstancesRequest {
	s.ClockOptions = v
	return s
}

func (s *RunInstancesRequest) SetCreditSpecification(v string) *RunInstancesRequest {
	s.CreditSpecification = &v
	return s
}

func (s *RunInstancesRequest) SetDataDisk(v []*RunInstancesRequestDataDisk) *RunInstancesRequest {
	s.DataDisk = v
	return s
}

func (s *RunInstancesRequest) SetDedicatedHostId(v string) *RunInstancesRequest {
	s.DedicatedHostId = &v
	return s
}

func (s *RunInstancesRequest) SetDeletionProtection(v bool) *RunInstancesRequest {
	s.DeletionProtection = &v
	return s
}

func (s *RunInstancesRequest) SetDeploymentSetGroupNo(v int32) *RunInstancesRequest {
	s.DeploymentSetGroupNo = &v
	return s
}

func (s *RunInstancesRequest) SetDeploymentSetId(v string) *RunInstancesRequest {
	s.DeploymentSetId = &v
	return s
}

func (s *RunInstancesRequest) SetDescription(v string) *RunInstancesRequest {
	s.Description = &v
	return s
}

func (s *RunInstancesRequest) SetDryRun(v bool) *RunInstancesRequest {
	s.DryRun = &v
	return s
}

func (s *RunInstancesRequest) SetHostName(v string) *RunInstancesRequest {
	s.HostName = &v
	return s
}

func (s *RunInstancesRequest) SetHostNames(v []*string) *RunInstancesRequest {
	s.HostNames = v
	return s
}

func (s *RunInstancesRequest) SetHpcClusterId(v string) *RunInstancesRequest {
	s.HpcClusterId = &v
	return s
}

func (s *RunInstancesRequest) SetHttpEndpoint(v string) *RunInstancesRequest {
	s.HttpEndpoint = &v
	return s
}

func (s *RunInstancesRequest) SetHttpPutResponseHopLimit(v int32) *RunInstancesRequest {
	s.HttpPutResponseHopLimit = &v
	return s
}

func (s *RunInstancesRequest) SetHttpTokens(v string) *RunInstancesRequest {
	s.HttpTokens = &v
	return s
}

func (s *RunInstancesRequest) SetImageFamily(v string) *RunInstancesRequest {
	s.ImageFamily = &v
	return s
}

func (s *RunInstancesRequest) SetImageId(v string) *RunInstancesRequest {
	s.ImageId = &v
	return s
}

func (s *RunInstancesRequest) SetImageOptions(v *RunInstancesRequestImageOptions) *RunInstancesRequest {
	s.ImageOptions = v
	return s
}

func (s *RunInstancesRequest) SetInstanceChargeType(v string) *RunInstancesRequest {
	s.InstanceChargeType = &v
	return s
}

func (s *RunInstancesRequest) SetInstanceName(v string) *RunInstancesRequest {
	s.InstanceName = &v
	return s
}

func (s *RunInstancesRequest) SetInstanceType(v string) *RunInstancesRequest {
	s.InstanceType = &v
	return s
}

func (s *RunInstancesRequest) SetInternetChargeType(v string) *RunInstancesRequest {
	s.InternetChargeType = &v
	return s
}

func (s *RunInstancesRequest) SetInternetMaxBandwidthIn(v int32) *RunInstancesRequest {
	s.InternetMaxBandwidthIn = &v
	return s
}

func (s *RunInstancesRequest) SetInternetMaxBandwidthOut(v int32) *RunInstancesRequest {
	s.InternetMaxBandwidthOut = &v
	return s
}

func (s *RunInstancesRequest) SetIoOptimized(v string) *RunInstancesRequest {
	s.IoOptimized = &v
	return s
}

func (s *RunInstancesRequest) SetIpv6Address(v []*string) *RunInstancesRequest {
	s.Ipv6Address = v
	return s
}

func (s *RunInstancesRequest) SetIpv6AddressCount(v int32) *RunInstancesRequest {
	s.Ipv6AddressCount = &v
	return s
}

func (s *RunInstancesRequest) SetIsp(v string) *RunInstancesRequest {
	s.Isp = &v
	return s
}

func (s *RunInstancesRequest) SetKeyPairName(v string) *RunInstancesRequest {
	s.KeyPairName = &v
	return s
}

func (s *RunInstancesRequest) SetLaunchTemplateId(v string) *RunInstancesRequest {
	s.LaunchTemplateId = &v
	return s
}

func (s *RunInstancesRequest) SetLaunchTemplateName(v string) *RunInstancesRequest {
	s.LaunchTemplateName = &v
	return s
}

func (s *RunInstancesRequest) SetLaunchTemplateVersion(v int64) *RunInstancesRequest {
	s.LaunchTemplateVersion = &v
	return s
}

func (s *RunInstancesRequest) SetMinAmount(v int32) *RunInstancesRequest {
	s.MinAmount = &v
	return s
}

func (s *RunInstancesRequest) SetNetworkInterface(v []*RunInstancesRequestNetworkInterface) *RunInstancesRequest {
	s.NetworkInterface = v
	return s
}

func (s *RunInstancesRequest) SetNetworkInterfaceQueueNumber(v int32) *RunInstancesRequest {
	s.NetworkInterfaceQueueNumber = &v
	return s
}

func (s *RunInstancesRequest) SetNetworkOptions(v *RunInstancesRequestNetworkOptions) *RunInstancesRequest {
	s.NetworkOptions = v
	return s
}

func (s *RunInstancesRequest) SetOwnerAccount(v string) *RunInstancesRequest {
	s.OwnerAccount = &v
	return s
}

func (s *RunInstancesRequest) SetOwnerId(v int64) *RunInstancesRequest {
	s.OwnerId = &v
	return s
}

func (s *RunInstancesRequest) SetPassword(v string) *RunInstancesRequest {
	s.Password = &v
	return s
}

func (s *RunInstancesRequest) SetPasswordInherit(v bool) *RunInstancesRequest {
	s.PasswordInherit = &v
	return s
}

func (s *RunInstancesRequest) SetPeriod(v int32) *RunInstancesRequest {
	s.Period = &v
	return s
}

func (s *RunInstancesRequest) SetPeriodUnit(v string) *RunInstancesRequest {
	s.PeriodUnit = &v
	return s
}

func (s *RunInstancesRequest) SetPrivateDnsNameOptions(v *RunInstancesRequestPrivateDnsNameOptions) *RunInstancesRequest {
	s.PrivateDnsNameOptions = v
	return s
}

func (s *RunInstancesRequest) SetPrivateIpAddress(v string) *RunInstancesRequest {
	s.PrivateIpAddress = &v
	return s
}

func (s *RunInstancesRequest) SetRamRoleName(v string) *RunInstancesRequest {
	s.RamRoleName = &v
	return s
}

func (s *RunInstancesRequest) SetRegionId(v string) *RunInstancesRequest {
	s.RegionId = &v
	return s
}

func (s *RunInstancesRequest) SetResourceGroupId(v string) *RunInstancesRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *RunInstancesRequest) SetResourceOwnerAccount(v string) *RunInstancesRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *RunInstancesRequest) SetResourceOwnerId(v int64) *RunInstancesRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *RunInstancesRequest) SetSecurityEnhancementStrategy(v string) *RunInstancesRequest {
	s.SecurityEnhancementStrategy = &v
	return s
}

func (s *RunInstancesRequest) SetSecurityGroupId(v string) *RunInstancesRequest {
	s.SecurityGroupId = &v
	return s
}

func (s *RunInstancesRequest) SetSecurityGroupIds(v []*string) *RunInstancesRequest {
	s.SecurityGroupIds = v
	return s
}

func (s *RunInstancesRequest) SetSpotDuration(v int32) *RunInstancesRequest {
	s.SpotDuration = &v
	return s
}

func (s *RunInstancesRequest) SetSpotInterruptionBehavior(v string) *RunInstancesRequest {
	s.SpotInterruptionBehavior = &v
	return s
}

func (s *RunInstancesRequest) SetSpotPriceLimit(v float32) *RunInstancesRequest {
	s.SpotPriceLimit = &v
	return s
}

func (s *RunInstancesRequest) SetSpotStrategy(v string) *RunInstancesRequest {
	s.SpotStrategy = &v
	return s
}

func (s *RunInstancesRequest) SetStorageSetId(v string) *RunInstancesRequest {
	s.StorageSetId = &v
	return s
}

func (s *RunInstancesRequest) SetStorageSetPartitionNumber(v int32) *RunInstancesRequest {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *RunInstancesRequest) SetTag(v []*RunInstancesRequestTag) *RunInstancesRequest {
	s.Tag = v
	return s
}

func (s *RunInstancesRequest) SetTenancy(v string) *RunInstancesRequest {
	s.Tenancy = &v
	return s
}

func (s *RunInstancesRequest) SetUniqueSuffix(v bool) *RunInstancesRequest {
	s.UniqueSuffix = &v
	return s
}

func (s *RunInstancesRequest) SetUserData(v string) *RunInstancesRequest {
	s.UserData = &v
	return s
}

func (s *RunInstancesRequest) SetVSwitchId(v string) *RunInstancesRequest {
	s.VSwitchId = &v
	return s
}

func (s *RunInstancesRequest) SetZoneId(v string) *RunInstancesRequest {
	s.ZoneId = &v
	return s
}

func (s *RunInstancesRequest) Validate() error {
	if s.CpuOptions != nil {
		if err := s.CpuOptions.Validate(); err != nil {
			return err
		}
	}
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
	if s.SchedulerOptions != nil {
		if err := s.SchedulerOptions.Validate(); err != nil {
			return err
		}
	}
	if s.SecurityOptions != nil {
		if err := s.SecurityOptions.Validate(); err != nil {
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
	if s.ClockOptions != nil {
		if err := s.ClockOptions.Validate(); err != nil {
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
	if s.NetworkOptions != nil {
		if err := s.NetworkOptions.Validate(); err != nil {
			return err
		}
	}
	if s.PrivateDnsNameOptions != nil {
		if err := s.PrivateDnsNameOptions.Validate(); err != nil {
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

type RunInstancesRequestCpuOptions struct {
	// The number of CPU cores.
	//
	// example:
	//
	// 2
	Core *int32 `json:"Core,omitempty" xml:"Core,omitempty"`
	// This parameter is no longer used.
	//
	// example:
	//
	// 1
	Numa *string `json:"Numa,omitempty" xml:"Numa,omitempty"`
	// The number of threads per CPU core. The following formula is used to calculate the number of vCPUs of the instance: `CpuOptions.Core` value × `CpuOptions.ThreadsPerCore` value.
	//
	// 	- If `CpuOptionsThreadPerCore` is set to 1, Hyper-Threading (HT) is disabled.
	//
	// 	- This parameter is applicable only to specific instance types.
	//
	// example:
	//
	// 2
	ThreadsPerCore *int32 `json:"ThreadsPerCore,omitempty" xml:"ThreadsPerCore,omitempty"`
	// The CPU topology type of the instance. Valid values:
	//
	// 	- ContinuousCoreToHTMapping: The HT technology allows continuous threads to run on the same core in the CPU topology of the instance.``
	//
	// 	- DiscreteCoreToHTMapping: The HT technology allows discrete threads to run on the same core in the CPU topology of the instance.``
	//
	// This parameter is empty by default.
	//
	// >  This parameter is supported only for specific instance families. For more information about the supported instance families, see [View and modify the CPU topology](https://help.aliyun.com/document_detail/2636059.html).
	//
	// example:
	//
	// DiscreteCoreToHTMapping
	TopologyType *string `json:"TopologyType,omitempty" xml:"TopologyType,omitempty"`
}

func (s RunInstancesRequestCpuOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestCpuOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestCpuOptions) GetCore() *int32 {
	return s.Core
}

func (s *RunInstancesRequestCpuOptions) GetNuma() *string {
	return s.Numa
}

func (s *RunInstancesRequestCpuOptions) GetThreadsPerCore() *int32 {
	return s.ThreadsPerCore
}

func (s *RunInstancesRequestCpuOptions) GetTopologyType() *string {
	return s.TopologyType
}

func (s *RunInstancesRequestCpuOptions) SetCore(v int32) *RunInstancesRequestCpuOptions {
	s.Core = &v
	return s
}

func (s *RunInstancesRequestCpuOptions) SetNuma(v string) *RunInstancesRequestCpuOptions {
	s.Numa = &v
	return s
}

func (s *RunInstancesRequestCpuOptions) SetThreadsPerCore(v int32) *RunInstancesRequestCpuOptions {
	s.ThreadsPerCore = &v
	return s
}

func (s *RunInstancesRequestCpuOptions) SetTopologyType(v string) *RunInstancesRequestCpuOptions {
	s.TopologyType = &v
	return s
}

func (s *RunInstancesRequestCpuOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestHibernationOptions struct {
	// > This parameter is in invitational preview and is unavailable.
	//
	// example:
	//
	// false
	Configured *bool `json:"Configured,omitempty" xml:"Configured,omitempty"`
}

func (s RunInstancesRequestHibernationOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestHibernationOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestHibernationOptions) GetConfigured() *bool {
	return s.Configured
}

func (s *RunInstancesRequestHibernationOptions) SetConfigured(v bool) *RunInstancesRequestHibernationOptions {
	s.Configured = &v
	return s
}

func (s *RunInstancesRequestHibernationOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestPrivatePoolOptions struct {
	// The ID of the private pool. The ID of a private pool is the same as that of the elasticity assurance or capacity reservation for which the private pool is generated.
	//
	// example:
	//
	// eap-bp67acfmxazb4****
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The type of the private pool to use to create the instance. A private pool is generated after an elasticity assurance or a capacity reservation takes effect. You can select the private pool when you start an instance. Valid values:
	//
	// 	- Open: open private pool. The system selects a matching open private pool to create the instance. If no matching open private pools are found, resources in the public pool are used. When you set this parameter to Open, you can leave the `PrivatePoolOptions.Id` parameter empty.
	//
	// 	- Target: specified private pool. The system uses the capacity in a specified private pool to create the instance. If the specified private pool is unavailable, the instance cannot be created. If you set this parameter to Target, you must specify the `PrivatePoolOptions.Id` parameter.
	//
	// 	- None: no private pool. The capacity in private pools is not used.
	//
	// Default value: None.
	//
	// In the following scenarios, the PrivatePoolOptions.MatchCriteria parameter can be set only to `None` or left empty:
	//
	// 	- A spot instance is created.
	//
	// 	- The instance is created in the classic network.
	//
	// 	- The instance is created on a dedicated host.
	//
	// example:
	//
	// Open
	MatchCriteria *string `json:"MatchCriteria,omitempty" xml:"MatchCriteria,omitempty"`
}

func (s RunInstancesRequestPrivatePoolOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestPrivatePoolOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestPrivatePoolOptions) GetId() *string {
	return s.Id
}

func (s *RunInstancesRequestPrivatePoolOptions) GetMatchCriteria() *string {
	return s.MatchCriteria
}

func (s *RunInstancesRequestPrivatePoolOptions) SetId(v string) *RunInstancesRequestPrivatePoolOptions {
	s.Id = &v
	return s
}

func (s *RunInstancesRequestPrivatePoolOptions) SetMatchCriteria(v string) *RunInstancesRequestPrivatePoolOptions {
	s.MatchCriteria = &v
	return s
}

func (s *RunInstancesRequestPrivatePoolOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestSchedulerOptions struct {
	// The ID of the dedicated host cluster in which to create the instance. After this parameter is specified, the system selects one dedicated host from the specified cluster to create the instance.
	//
	// > This parameter is valid only when the `Tenancy` parameter is set to `host`.
	//
	// When you specify both the `DedicatedHostId` and `SchedulerOptions.DedicatedHostClusterId` parameters, take note of the following items:
	//
	// 	- If the specified dedicated host belongs to the specified dedicated host cluster, the instance is preferentially deployed on the specified dedicated host.
	//
	// 	- If the specified dedicated host does not belong to the specified dedicated host cluster, the instance cannot be created.
	//
	// You can call the [DescribeDedicatedHostClusters](https://help.aliyun.com/document_detail/184145.html) operation to query the list of dedicated host cluster IDs.
	//
	// example:
	//
	// dc-bp12wlf6am0vz9v2****
	DedicatedHostClusterId *string `json:"DedicatedHostClusterId,omitempty" xml:"DedicatedHostClusterId,omitempty"`
}

func (s RunInstancesRequestSchedulerOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestSchedulerOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestSchedulerOptions) GetDedicatedHostClusterId() *string {
	return s.DedicatedHostClusterId
}

func (s *RunInstancesRequestSchedulerOptions) SetDedicatedHostClusterId(v string) *RunInstancesRequestSchedulerOptions {
	s.DedicatedHostClusterId = &v
	return s
}

func (s *RunInstancesRequestSchedulerOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestSecurityOptions struct {
	// The confidential computing mode. Set the value to Enclave.
	//
	// A value of Enclave indicates that an enclave-based confidential computing environment is built on the instance. When you call the `RunInstances` operation, you can set this parameter only for c7, g7, or r7 instances to use enclave-based confidential computing. Take note of the following items:
	//
	// 	- The confidential computing feature is in invitational preview.
	//
	// 	- When you use the ECS API to create instances that support enclave-based confidential computing, you can call only the `RunInstances` operation. The `CreateInstance` operation does not support the `SecurityOptions.ConfidentialComputingMode` parameter.
	//
	// 	- Enclave-based confidential computing is implemented based on Alibaba Cloud Trusted System (vTPM). When you build a confidential computing environment on an instance by using Enclave, Alibaba Cloud Trusted System is enabled for the instance. If you set `SecurityOptions.ConfidentialComputingMode` to Enclave when you call this operation, the created instances use enclave-based confidential computing and Alibaba Cloud Trusted System regardless of whether `SecurityOptions.TrustedSystemMode` is set to vTPM.
	//
	// For more information about confidential computing, see [Build a confidential computing environment by using Enclave](https://help.aliyun.com/document_detail/203433.html).
	//
	// example:
	//
	// Enclave
	ConfidentialComputingMode *string `json:"ConfidentialComputingMode,omitempty" xml:"ConfidentialComputingMode,omitempty"`
	// The trusted system mode. Set the value to vTPM.
	//
	// The trusted system mode supports the following instance families:
	//
	// 	- g7, c7, and r7
	//
	// 	- Security-enhanced instance families: g7t, c7t, and r7t
	//
	// When you create instances of the preceding instance families, you must set this parameter. Take note of the following items:
	//
	// 	- To use the Alibaba Cloud trusted system, set this parameter to vTPM. Then, the Alibaba Cloud trusted system performs trust verifications when the instances start.
	//
	// 	- If you do not want to use the Alibaba Cloud trusted system, leave this parameter empty. Note that if your created instances use an enclave-based confidential computing environment (with `SecurityOptions.ConfidentialComputingMode` set to Enclave), the Alibaba Cloud trusted system is enabled for the instances.
	//
	// 	- When you use the ECS API to create instances that use the trusted system, you can call only the `RunInstances` operation. The `CreateInstance` operation does not support the `SecurityOptions.TrustedSystemMode` parameter.
	//
	// > If you have configured an instance as a trusted one when you created the instance, you can use only an image that support the trusted system to replace the system disk of the instance.
	//
	// For more information about the trusted system, see [Overview](https://help.aliyun.com/document_detail/201394.html).
	//
	// example:
	//
	// vTPM
	TrustedSystemMode *string `json:"TrustedSystemMode,omitempty" xml:"TrustedSystemMode,omitempty"`
}

func (s RunInstancesRequestSecurityOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestSecurityOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestSecurityOptions) GetConfidentialComputingMode() *string {
	return s.ConfidentialComputingMode
}

func (s *RunInstancesRequestSecurityOptions) GetTrustedSystemMode() *string {
	return s.TrustedSystemMode
}

func (s *RunInstancesRequestSecurityOptions) SetConfidentialComputingMode(v string) *RunInstancesRequestSecurityOptions {
	s.ConfidentialComputingMode = &v
	return s
}

func (s *RunInstancesRequestSecurityOptions) SetTrustedSystemMode(v string) *RunInstancesRequestSecurityOptions {
	s.TrustedSystemMode = &v
	return s
}

func (s *RunInstancesRequestSecurityOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestSystemDisk struct {
	// The ID of the automatic snapshot policy to apply to the system disk.
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// The category of the system disk. Valid values:
	//
	// 	- cloud_efficiency: utra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: enhanced SSD (ESSD)
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// 	- cloud_essd_entry: ESSD Entry disk
	//
	// >  The value of this parameter can be `cloud_essd_entry` only when `InstanceType` is set to `ecs.u1` or `ecs.e`. ecs.u1 indicates the u1 universal instance family and ecs.e indicates the e economy instance family. For information about the u1 and e instance families, see the [u1, universal instance family](https://help.aliyun.com/document_detail/457079.html) section in the "Universal instance families" topic and the [e, economy instance family](https://help.aliyun.com/document_detail/108489.html) section in the "Shared instance families" topic.
	//
	// For non-I/O optimized instances of retired instance types, the default value is cloud. For other types of instances, the default value is cloud_efficiency.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The description of the system disk. The description must be 2 to 256 characters in length. The description can contain letters but cannot start with `http://` or `https://`.
	//
	// example:
	//
	// SystemDisk_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The name of the system disk. The name must be 2 to 128 characters in length and support Unicode characters under the Decimal Number category and the categories whose names contain Letter. The name can contain colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// cloud_ssdSystem
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// The performance level of the ESSD to use as the system disk. Default value: PL1. Valid values:
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
	// The size of the system disk. Unit: GiB. Valid values:
	//
	// 	- Basic disk: 20 to 500.
	//
	// 	- ESSD: Valid values vary based on the performance level of the ESSD.
	//
	//     	- PL0 ESSD: 1 to 2048.
	//
	//     	- PL1 ESSD: 20 to 2048.
	//
	//     	- PL2 ESSD: 461 to 2048.
	//
	//     	- PL3 ESSD: 1261 to 2048.
	//
	// 	- ESSD AutoPL disk: 1 to 2048.
	//
	// 	- Other disk categories: 20 to 2048.
	//
	// The value of this parameter must be at least 1 and greater than or equal to the image size.
	//
	// Default value: 40 or the image size, whichever is greater.
	//
	// example:
	//
	// 40
	Size *string `json:"Size,omitempty" xml:"Size,omitempty"`
	// Specifies whether to enable the performance burst feature for the system disk. Valid values:
	//
	// 	- true: enables the performance burst feature for the system disk.
	//
	// 	- false: disables the performance burst feature for the system disk.
	//
	// >  This parameter is available only if you set `SystemDisk.Category` to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// ase-256
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt the system disk. Valid values:
	//
	// 	- true: encrypts the system disk.
	//
	// 	- false: does not encrypt the system disk.
	//
	// Default value: false.
	//
	// >  The system disks of instances cannot be encrypted during instance creation in Hong Kong Zone D or Singapore Zone A.
	//
	// >  When you use a shared encrypted image to create the disk based on an encrypted snapshot, you must set Encrypted to true to ensure that the disk uses an encryption key of your own.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key to use for the system disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk to use as the system disk. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}.
	//
	// Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// >  This parameter is available only if you set `SystemDisk.Category` to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The ID of the dedicated block storage cluster to which the system disk belongs. If you want to use disks in a dedicated block storage cluster as system disks when you create instances, specify this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s RunInstancesRequestSystemDisk) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestSystemDisk) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestSystemDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *RunInstancesRequestSystemDisk) GetCategory() *string {
	return s.Category
}

func (s *RunInstancesRequestSystemDisk) GetDescription() *string {
	return s.Description
}

func (s *RunInstancesRequestSystemDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *RunInstancesRequestSystemDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *RunInstancesRequestSystemDisk) GetSize() *string {
	return s.Size
}

func (s *RunInstancesRequestSystemDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *RunInstancesRequestSystemDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *RunInstancesRequestSystemDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *RunInstancesRequestSystemDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *RunInstancesRequestSystemDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *RunInstancesRequestSystemDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *RunInstancesRequestSystemDisk) SetAutoSnapshotPolicyId(v string) *RunInstancesRequestSystemDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetCategory(v string) *RunInstancesRequestSystemDisk {
	s.Category = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetDescription(v string) *RunInstancesRequestSystemDisk {
	s.Description = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetDiskName(v string) *RunInstancesRequestSystemDisk {
	s.DiskName = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetPerformanceLevel(v string) *RunInstancesRequestSystemDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetSize(v string) *RunInstancesRequestSystemDisk {
	s.Size = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetBurstingEnabled(v bool) *RunInstancesRequestSystemDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetEncryptAlgorithm(v string) *RunInstancesRequestSystemDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetEncrypted(v string) *RunInstancesRequestSystemDisk {
	s.Encrypted = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetKMSKeyId(v string) *RunInstancesRequestSystemDisk {
	s.KMSKeyId = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetProvisionedIops(v int64) *RunInstancesRequestSystemDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) SetStorageClusterId(v string) *RunInstancesRequestSystemDisk {
	s.StorageClusterId = &v
	return s
}

func (s *RunInstancesRequestSystemDisk) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestArn struct {
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// 0
	AssumeRoleFor *int64 `json:"AssumeRoleFor,omitempty" xml:"AssumeRoleFor,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	RoleType *string `json:"RoleType,omitempty" xml:"RoleType,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// null
	Rolearn *string `json:"Rolearn,omitempty" xml:"Rolearn,omitempty"`
}

func (s RunInstancesRequestArn) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestArn) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestArn) GetAssumeRoleFor() *int64 {
	return s.AssumeRoleFor
}

func (s *RunInstancesRequestArn) GetRoleType() *string {
	return s.RoleType
}

func (s *RunInstancesRequestArn) GetRolearn() *string {
	return s.Rolearn
}

func (s *RunInstancesRequestArn) SetAssumeRoleFor(v int64) *RunInstancesRequestArn {
	s.AssumeRoleFor = &v
	return s
}

func (s *RunInstancesRequestArn) SetRoleType(v string) *RunInstancesRequestArn {
	s.RoleType = &v
	return s
}

func (s *RunInstancesRequestArn) SetRolearn(v string) *RunInstancesRequestArn {
	s.Rolearn = &v
	return s
}

func (s *RunInstancesRequestArn) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestClockOptions struct {
	PtpStatus *string `json:"PtpStatus,omitempty" xml:"PtpStatus,omitempty"`
}

func (s RunInstancesRequestClockOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestClockOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestClockOptions) GetPtpStatus() *string {
	return s.PtpStatus
}

func (s *RunInstancesRequestClockOptions) SetPtpStatus(v string) *RunInstancesRequestClockOptions {
	s.PtpStatus = &v
	return s
}

func (s *RunInstancesRequestClockOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestDataDisk struct {
	// The ID of the automatic snapshot policy to apply to data disk N.
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// Specifies whether to enable the performance burst feature for data disk N. Valid values:
	//
	// 	- true: enables the performance burst feature for the system disk.
	//
	// 	- false: disables the performance burst feature for the data disk.
	//
	// >  This parameter is available only if you set DataDisk.N.Category to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
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
	// 	- cloud_regional_disk_auto: Regional ESSD.
	//
	// 	- cloud_essd_entry: ESSD Entry disk.
	//
	//     **
	//
	//     **Note*	- This parameter can be set to `cloud_essd_entry` only when `InstanceType` is set to `ecs.u1` or `ecs.e`.
	//
	// 	- elastic_ephemeral_disk_standard: standard elastic ephemeral disk.
	//
	// 	- elastic_ephemeral_disk_premium: premium elastic ephemeral disk
	//
	// For I/O optimized instances, the default value is cloud_efficiency. For non-I/O optimized instances, the default value is cloud.
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// Specifies whether to release data disk N when the associated instance is released. Valid values:
	//
	// 	- true: releases the data disk when the associated instance is released.
	//
	// 	- false: does not release the data disk when the associated instance is released.
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
	// DataDisk_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The mount point of data disk N. The mount points are named based on the number of data disks:
	//
	// 	- 1st to 25th data disks: /dev/xvd`[b-z]`.
	//
	// 	- From the 26th data disk on: /dev/xvd`[aa-zz]`. For example, the 26th data disk is named /dev/xvdaa, the 27th data disk is named /dev/xvdab, and so on.
	//
	// >
	//
	// 	- This parameter is applicable to scenarios in which a full image is used to create instances. A full image is an image that contains an operating system, application software, and business data. For these scenarios, you can set this parameter to the mount point of data disk N in the full image and modify `DataDisk.N.Size` and `DataDisk.N.Category` to change the category and size of data disk N created based on the image.
	//
	// 	- When you use a full image to create an ECS instance, the data disks in the image are created as the first N data disks of the instance.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The name of data disk N. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// example:
	//
	// cloud_ssdData
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// >  This parameter is not publicly available.
	//
	// example:
	//
	// aes-256
	EncryptAlgorithm *string `json:"EncryptAlgorithm,omitempty" xml:"EncryptAlgorithm,omitempty"`
	// Specifies whether to encrypt data disk N. Valid values:
	//
	// 	- true: encrypts the data disk.
	//
	// 	- false: does not encrypt the data disk.
	//
	// Default value: false.
	//
	// >  When you use a shared encrypted image to create the disk based on an encrypted snapshot, you must set Encrypted to true to ensure that the disk uses an encryption key of your own.
	//
	// example:
	//
	// false
	Encrypted *string `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The ID of the KMS key used for the data disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb40****
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
	// The provisioned read/write IOPS of the ESSD AutoPL disk to use as data disk N. Valid values: 0 to min{50,000, 1,000 × Capacity - Baseline IOPS}.
	//
	// Baseline IOPS = min{1,800 + 50 × Capacity, 50,000}.
	//
	// >  This parameter is available only if you set DataDisk.N.Category to cloud_auto. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The size of data disk N. Valid values of N: 1 to 16. Unit: GiB. Valid values:
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud_efficiency: 20 to 32768.
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud_ssd: 20 to 32768.
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud_essd: vary based on the value of `DataDisk.N.PerformanceLevel`.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL0: 1 to 65536.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL1: 20 to 65536.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL2: 461 to 65536.
	//
	//     	- Valid values when DataDisk.N.PerformanceLevel is set to PL3: 1261 to 65536.
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud: 5 to 2000.
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud_auto: 1 to 65536.
	//
	// 	- Valid values when DataDisk.N.Category is set to cloud_essd_entry: 10 to 32768.
	//
	// >  The value of this parameter must be greater than or equal to the size of the snapshot specified by `DataDisk.N.SnapshotId`.
	//
	// example:
	//
	// 2000
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot to use to create data disk N. Valid values of N: 1 to 16.
	//
	// When `DataDisk.N.SnapshotId` is specified, `DataDisk.N.Size` is ignored. The data disk is created with the size of the specified snapshot. Use snapshots created on or after July 15, 2013. Otherwise, an error is returned and your request is rejected.
	//
	// example:
	//
	// s-bp17441ohwka0yuh****
	SnapshotId *string `json:"SnapshotId,omitempty" xml:"SnapshotId,omitempty"`
	// The ID of the dedicated block storage cluster to which data disk N belongs. If you want to use a disk in a dedicated block storage cluster as data disk N when you create the instance, you must specify this parameter.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
}

func (s RunInstancesRequestDataDisk) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestDataDisk) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestDataDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *RunInstancesRequestDataDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *RunInstancesRequestDataDisk) GetCategory() *string {
	return s.Category
}

func (s *RunInstancesRequestDataDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *RunInstancesRequestDataDisk) GetDescription() *string {
	return s.Description
}

func (s *RunInstancesRequestDataDisk) GetDevice() *string {
	return s.Device
}

func (s *RunInstancesRequestDataDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *RunInstancesRequestDataDisk) GetEncryptAlgorithm() *string {
	return s.EncryptAlgorithm
}

func (s *RunInstancesRequestDataDisk) GetEncrypted() *string {
	return s.Encrypted
}

func (s *RunInstancesRequestDataDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *RunInstancesRequestDataDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *RunInstancesRequestDataDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *RunInstancesRequestDataDisk) GetSize() *int32 {
	return s.Size
}

func (s *RunInstancesRequestDataDisk) GetSnapshotId() *string {
	return s.SnapshotId
}

func (s *RunInstancesRequestDataDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *RunInstancesRequestDataDisk) SetAutoSnapshotPolicyId(v string) *RunInstancesRequestDataDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetBurstingEnabled(v bool) *RunInstancesRequestDataDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetCategory(v string) *RunInstancesRequestDataDisk {
	s.Category = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetDeleteWithInstance(v bool) *RunInstancesRequestDataDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetDescription(v string) *RunInstancesRequestDataDisk {
	s.Description = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetDevice(v string) *RunInstancesRequestDataDisk {
	s.Device = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetDiskName(v string) *RunInstancesRequestDataDisk {
	s.DiskName = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetEncryptAlgorithm(v string) *RunInstancesRequestDataDisk {
	s.EncryptAlgorithm = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetEncrypted(v string) *RunInstancesRequestDataDisk {
	s.Encrypted = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetKMSKeyId(v string) *RunInstancesRequestDataDisk {
	s.KMSKeyId = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetPerformanceLevel(v string) *RunInstancesRequestDataDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetProvisionedIops(v int64) *RunInstancesRequestDataDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetSize(v int32) *RunInstancesRequestDataDisk {
	s.Size = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetSnapshotId(v string) *RunInstancesRequestDataDisk {
	s.SnapshotId = &v
	return s
}

func (s *RunInstancesRequestDataDisk) SetStorageClusterId(v string) *RunInstancesRequestDataDisk {
	s.StorageClusterId = &v
	return s
}

func (s *RunInstancesRequestDataDisk) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestImageOptions struct {
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

func (s RunInstancesRequestImageOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestImageOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestImageOptions) GetLoginAsNonRoot() *bool {
	return s.LoginAsNonRoot
}

func (s *RunInstancesRequestImageOptions) SetLoginAsNonRoot(v bool) *RunInstancesRequestImageOptions {
	s.LoginAsNonRoot = &v
	return s
}

func (s *RunInstancesRequestImageOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestNetworkInterface struct {
	// Specifies whether to release ENI N when the associated instance is released. Valid values:
	//
	// 	- true: releases the ENI when the associated instance is released.
	//
	// 	- false: retains the ENI when the associated instance is released.
	//
	// Default value: true.
	//
	// >  This parameter takes effect only for secondary ENIs.
	//
	// example:
	//
	// true
	DeleteOnRelease *bool `json:"DeleteOnRelease,omitempty" xml:"DeleteOnRelease,omitempty"`
	// The description of ENI N.
	//
	// Take note of the following items:
	//
	// 	- The value of N cannot exceed the maximum number of ENIs per instance that the instance type supports. For the maximum number of ENIs per instance that an instance type supports, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) operation.
	//
	// 	- The description must be 2 to 256 characters in length and cannot start with `http://` or `https://`.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you do not need to specify this parameter.
	//
	// example:
	//
	// Network_Description
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The type of ENI N. The value of the first N in this parameter cannot exceed the maximum number of ENIs per instance that the instance type supports. For the maximum number of ENIs per instance that an instance type supports, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) operation.
	//
	// Valid values:
	//
	// 	- Primary: the primary ENI
	//
	// 	- Secondary
	//
	// Default value: Secondary.
	//
	// example:
	//
	// Secondary
	InstanceType *string `json:"InstanceType,omitempty" xml:"InstanceType,omitempty"`
	// The IPv6 addresses to assign to the primary ENI. You can assign up to 10 IPv6 addresses to the primary ENI. Valid values of the second N: 1 to 10.
	//
	// Example: `Ipv6Address.1=2001:db8:1234:1a00::***`.
	//
	// Take note of the following items:
	//
	// 	- This parameter takes effect only when `NetworkInterface.N.InstanceType` is set to `Primary`. If you set `NetworkInterface.N.InstanceType` to `Secondary` or leave NetworkInterface.N.InstanceType empty, you cannot specify this parameter.
	//
	// 	- If you specify this parameter, you must set `Amount` to 1 and cannot specify `Ipv6AddressCount`, `Ipv6Address.N`, or `NetworkInterface.N.Ipv6AddressCount`.
	Ipv6Address []*string `json:"Ipv6Address,omitempty" xml:"Ipv6Address,omitempty" type:"Repeated"`
	// The number of IPv6 addresses to randomly generate for the primary ENI. Valid values: 1 to 10.
	//
	// Take note of the following items:
	//
	// 	- This parameter takes effect only when `NetworkInterface.N.InstanceType` is set to `Primary`. If you set `NetworkInterface.N.InstanceType` to `Secondary` or leave NetworkInterface.N.InstanceType empty, you cannot specify this parameter.
	//
	// 	- If you specify this parameter, you cannot specify `Ipv6AddressCount`, `Ipv6Address.N`, or `NetworkInterface.N.Ipv6Address.N`.
	//
	// example:
	//
	// 1
	Ipv6AddressCount *int64 `json:"Ipv6AddressCount,omitempty" xml:"Ipv6AddressCount,omitempty"`
	// The index of the network card for ENI N.
	//
	// Take note of the following items:
	//
	// 	- You can specify NIC indexes only for instances of specific instance types.
	//
	// 	- If you set NetworkInterface.N.InstanceType to Primary, you can set NetworkInterface.N.NetworkCardIndex only to 0 for instance types that support network cards.
	//
	// 	- If you set NetworkInterface.N.InstanceType to Secondary or leave NetworkInterface.N.InstanceType empty, you can specify NetworkInterface.N.NetworkCardIndex based on instance types if the instance types support network cards. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// 0
	NetworkCardIndex *int32 `json:"NetworkCardIndex,omitempty" xml:"NetworkCardIndex,omitempty"`
	// The ID of the ENI to attach to the instance.
	//
	// If you specify this parameter, you must set `Amount` to 1.
	//
	// >  This parameter takes effect only for secondary ENIs. After you specify an existing secondary ENI, you cannot specify other ENI creation parameters.
	//
	// example:
	//
	// eni-bp1gn106np8jhxhj****
	NetworkInterfaceId *string `json:"NetworkInterfaceId,omitempty" xml:"NetworkInterfaceId,omitempty"`
	// The name of ENI N. The name must be 2 to 128 characters in length and can contain letters, digits, colons (:), underscores (_), periods (.), and hyphens (-).
	//
	// Take note of the following items:
	//
	// 	- The value of N cannot exceed the maximum number of ENIs per instance that the instance type supports. For the maximum number of ENIs per instance that an instance type supports, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) operation.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you do not need to specify this parameter.
	//
	// example:
	//
	// Network_Name
	NetworkInterfaceName *string `json:"NetworkInterfaceName,omitempty" xml:"NetworkInterfaceName,omitempty"`
	// The communication mode of ENI N. Valid values:
	//
	// 	- Standard: uses the TCP communication mode.
	//
	// 	- HighPerformance: uses the remote direct memory access (RDMA) communication mode with Elastic RDMA Interface (ERI) enabled.
	//
	// Default value: Standard.
	//
	// >  The number of ERIs on an instance cannot exceed the maximum number of ERIs that the instance type supports. For more information, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html).
	//
	// example:
	//
	// Standard
	NetworkInterfaceTrafficMode *string `json:"NetworkInterfaceTrafficMode,omitempty" xml:"NetworkInterfaceTrafficMode,omitempty"`
	// The primary IP address to assign to ENI N.
	//
	// Take note of the following items:
	//
	// 	- The value of N cannot exceed the maximum number of ENIs per instance that the instance type supports. For the maximum number of ENIs per instance that an instance type supports, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) operation.
	//
	//     	- If the value of N is 1, you can configure a primary or secondary ENI. If you specify this parameter, set `Amount` to a numeric value greater than 1, and set NetworkInterface.N.InstanceType to Primary, the specified number of instances are created and consecutive primary IP addresses starting from the specified IP address are assigned to the instances. In this case, you cannot attach secondary ENIs to the instances.
	//
	//     	- If you specify this parameter, set `Amount` to a numeric value greater than 1, and set NetworkInterface.N.InstanceType to Primary, you cannot set `NetworkInterface.2.InstanceType` to Secondary to attach a secondary ENI.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Primary`, this parameter is equivalent to `PrivateIpAddress`. You cannot specify both this parameter and `PrivateIpAddress` in the same request.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Secondary` or leave NetworkInterface.N.InstanceType empty, the specified primary IP address is assigned to the secondary ENI. The default value is an IP address that is randomly selected from within the CIDR block of the vSwitch to which to connect the secondary ENI.
	//
	// >
	//
	// 	- The first IP address and last three IP addresses of each vSwitch CIDR block are reserved. You cannot specify the IP addresses. For example, if a vSwitch CIDR block is 192.168.1.0/24, the following IP addresses are reserved: 192.168.1.0, 192.168.1.253, 192.168.1.254, and 192.168.1.255.
	//
	// example:
	//
	// ``172.16.**.**``
	PrimaryIpAddress *string `json:"PrimaryIpAddress,omitempty" xml:"PrimaryIpAddress,omitempty"`
	// The number of queues supported by ENI N.
	//
	// Take note of the following items:
	//
	// 	- The value of N cannot exceed the maximum number of ENIs per instance that the instance type supports. For the maximum number of ENIs per instance that an instance type supports, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) operation.
	//
	// 	- The value of this parameter cannot exceed the maximum number of queues allowed per ENI.
	//
	// 	- The total number of queues for all ENIs of an instance cannot exceed the queue quota for the instance type. To query the maximum number of queues per ENI and the queue quota for an instance type, you can call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/25620.html) operation and check the `MaximumQueueNumberPerEni` and `TotalEniQueueQuantity` values in the response.
	//
	// 	- If you specify this parameter and set `NetworkInterface.N.InstanceType` to `Primary`, you cannot specify `NetworkInterfaceQueueNumber`.
	//
	// example:
	//
	// 8
	QueueNumber *int32 `json:"QueueNumber,omitempty" xml:"QueueNumber,omitempty"`
	// The number of queue pairs (QPs) supported by the ERI.
	//
	// If you want to attach multiple ERIs to a created instance, we recommend that you specify QueuePairNumber for each ERI based on the value of `QueuePairNumber` supported by the instance type and the number of ERIs that you want to use. Make sure that the total number of QPs of all ERIs does not exceed the maximum number of QPs supported by the instance type. For information about the maximum number of QPs supported by an instance type, see [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html).
	//
	// >  If you do not specify QueuePairNumber for an ERI, the maximum number of QPs supported by the instance type is used as the number of QPs supported by the ERI. In this case, you cannot attach an additional ERI to the instance. However, you can attach other types of ENIs to the instance.
	//
	// example:
	//
	// 0
	QueuePairNumber *int64 `json:"QueuePairNumber,omitempty" xml:"QueuePairNumber,omitempty"`
	// The receive (Rx) queue depth of ENI N.
	//
	// >  This parameter is in invitational preview and is not publicly available. To use this parameter, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl).
	//
	// Take note of the following items:
	//
	// 	- This parameter is applicable only to 7th-generation or later ECS instance types.
	//
	// 	- This parameter is applicable to Linux images.
	//
	// 	- A larger Rx queue depth yields higher inbound throughput and reduces packet loss rates but consumes more memory.
	//
	// example:
	//
	// 8192
	RxQueueSize *int32 `json:"RxQueueSize,omitempty" xml:"RxQueueSize,omitempty"`
	// The ID of the security group to which to assign ENI N.
	//
	// Take note of the following items:
	//
	// 	- The value of N cannot exceed the maximum number of ENIs per instance that the instance type supports. For the maximum number of ENIs per instance that an instance type supports, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) operation.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you must specify this parameter. In this case, this parameter is equivalent to `SecurityGroupId` and you cannot specify `SecurityGroupId`, `SecurityGroupIds.N`, or `NetworkInterface.N.SecurityGroupIds.N`.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Secondary` or leave NetworkInterface.N.InstanceType empty, you do not need to specify this parameter. The default value is the ID of the security group to which to assign the instance.
	//
	// example:
	//
	// sg-bp67acfmxazb4p****
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// The IDs of security groups to which to assign ENI N.
	//
	// 	- The value of the first N in this parameter cannot exceed the maximum number of ENIs per instance that the instance type supports. For the maximum number of ENIs per instance that an instance type supports, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) operation.
	//
	// 	- The second N in this parameter indicates that one or more security group IDs can be specified. The valid values of the second N vary based on the maximum number of security groups to which an instance can belong. For more information, see [Security group limits](~~25412#SecurityGroupQuota1~~).
	//
	// Take note of the following items:
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Primary`, you must specify this parameter or `NetworkInterface.N.SecurityGroupId`. In this case, this parameter is equivalent to `SecurityGroupIds.N`, and you cannot specify `SecurityGroupId`, `SecurityGroupIds.N`, or `NetworkInterface.N.SecurityGroupId`.
	//
	// 	- If you set `NetworkInterface.N.InstanceType` to `Secondary` or leave NetworkInterface.N.InstanceType empty, you do not need to specify this parameter. The default value is the ID of the security group to which to assign the instance.
	//
	// example:
	//
	// sg-bp15ed6xe1yxeycg7****
	SecurityGroupIds []*string `json:"SecurityGroupIds,omitempty" xml:"SecurityGroupIds,omitempty" type:"Repeated"`
	// Specifies whether to enable the source and destination IP address check feature. We recommend that you enable the feature to improve network security. Valid value:
	//
	// 	- true: enables the performance burst feature for the system disk.
	//
	// 	- false: disables the performance burst feature for the data disk.
	//
	// Default value: false.
	//
	// >  This feature is available only in some regions. Before you use this method, read [Source and destination IP address check](https://help.aliyun.com/document_detail/2863210.html).
	//
	// example:
	//
	// false
	SourceDestCheck *bool `json:"SourceDestCheck,omitempty" xml:"SourceDestCheck,omitempty"`
	// The Tx queue depth of ENI N.
	//
	// >  This parameter is in invitational preview and is not publicly available. To use this parameter, [submit a ticket](https://smartservice.console.aliyun.com/service/create-ticket-intl).
	//
	// Take note of the following items:
	//
	// 	- This parameter is applicable only to 7th-generation or later ECS instance types.
	//
	// 	- This parameter is applicable to Linux images.
	//
	// 	- A larger Tx queue depth yields higher outbound throughput and reduces packet loss rates but consumes more memory.
	//
	// example:
	//
	// 8192
	TxQueueSize *int32 `json:"TxQueueSize,omitempty" xml:"TxQueueSize,omitempty"`
	// The ID of the vSwitch to which to connect ENI N.
	//
	// When you specify this parameter, take note of the following items:
	//
	// 	- The value of N cannot exceed the maximum number of ENIs per instance that the instance type supports. For the maximum number of ENIs per instance that an instance type supports, see [Overview of instance families](https://help.aliyun.com/document_detail/25378.html) or call the [DescribeInstanceTypes](https://help.aliyun.com/document_detail/2679699.html) operation.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Primary`, you must specify this parameter. In this case, this parameter is equivalent to `VSwitchId`. You cannot specify both NetworkInterface.N.VSwitchId and `VSwitchId` in the same request.
	//
	// 	- If `NetworkInterface.N.InstanceType` is set to `Secondary` or left empty, you do not need to specify this parameter. The default value is the VSwitchId value.
	//
	// example:
	//
	// vsw-bp67acfmxazb4p****
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
}

func (s RunInstancesRequestNetworkInterface) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestNetworkInterface) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestNetworkInterface) GetDeleteOnRelease() *bool {
	return s.DeleteOnRelease
}

func (s *RunInstancesRequestNetworkInterface) GetDescription() *string {
	return s.Description
}

func (s *RunInstancesRequestNetworkInterface) GetInstanceType() *string {
	return s.InstanceType
}

func (s *RunInstancesRequestNetworkInterface) GetIpv6Address() []*string {
	return s.Ipv6Address
}

func (s *RunInstancesRequestNetworkInterface) GetIpv6AddressCount() *int64 {
	return s.Ipv6AddressCount
}

func (s *RunInstancesRequestNetworkInterface) GetNetworkCardIndex() *int32 {
	return s.NetworkCardIndex
}

func (s *RunInstancesRequestNetworkInterface) GetNetworkInterfaceId() *string {
	return s.NetworkInterfaceId
}

func (s *RunInstancesRequestNetworkInterface) GetNetworkInterfaceName() *string {
	return s.NetworkInterfaceName
}

func (s *RunInstancesRequestNetworkInterface) GetNetworkInterfaceTrafficMode() *string {
	return s.NetworkInterfaceTrafficMode
}

func (s *RunInstancesRequestNetworkInterface) GetPrimaryIpAddress() *string {
	return s.PrimaryIpAddress
}

func (s *RunInstancesRequestNetworkInterface) GetQueueNumber() *int32 {
	return s.QueueNumber
}

func (s *RunInstancesRequestNetworkInterface) GetQueuePairNumber() *int64 {
	return s.QueuePairNumber
}

func (s *RunInstancesRequestNetworkInterface) GetRxQueueSize() *int32 {
	return s.RxQueueSize
}

func (s *RunInstancesRequestNetworkInterface) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *RunInstancesRequestNetworkInterface) GetSecurityGroupIds() []*string {
	return s.SecurityGroupIds
}

func (s *RunInstancesRequestNetworkInterface) GetSourceDestCheck() *bool {
	return s.SourceDestCheck
}

func (s *RunInstancesRequestNetworkInterface) GetTxQueueSize() *int32 {
	return s.TxQueueSize
}

func (s *RunInstancesRequestNetworkInterface) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *RunInstancesRequestNetworkInterface) SetDeleteOnRelease(v bool) *RunInstancesRequestNetworkInterface {
	s.DeleteOnRelease = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetDescription(v string) *RunInstancesRequestNetworkInterface {
	s.Description = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetInstanceType(v string) *RunInstancesRequestNetworkInterface {
	s.InstanceType = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetIpv6Address(v []*string) *RunInstancesRequestNetworkInterface {
	s.Ipv6Address = v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetIpv6AddressCount(v int64) *RunInstancesRequestNetworkInterface {
	s.Ipv6AddressCount = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetNetworkCardIndex(v int32) *RunInstancesRequestNetworkInterface {
	s.NetworkCardIndex = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetNetworkInterfaceId(v string) *RunInstancesRequestNetworkInterface {
	s.NetworkInterfaceId = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetNetworkInterfaceName(v string) *RunInstancesRequestNetworkInterface {
	s.NetworkInterfaceName = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetNetworkInterfaceTrafficMode(v string) *RunInstancesRequestNetworkInterface {
	s.NetworkInterfaceTrafficMode = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetPrimaryIpAddress(v string) *RunInstancesRequestNetworkInterface {
	s.PrimaryIpAddress = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetQueueNumber(v int32) *RunInstancesRequestNetworkInterface {
	s.QueueNumber = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetQueuePairNumber(v int64) *RunInstancesRequestNetworkInterface {
	s.QueuePairNumber = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetRxQueueSize(v int32) *RunInstancesRequestNetworkInterface {
	s.RxQueueSize = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetSecurityGroupId(v string) *RunInstancesRequestNetworkInterface {
	s.SecurityGroupId = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetSecurityGroupIds(v []*string) *RunInstancesRequestNetworkInterface {
	s.SecurityGroupIds = v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetSourceDestCheck(v bool) *RunInstancesRequestNetworkInterface {
	s.SourceDestCheck = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetTxQueueSize(v int32) *RunInstancesRequestNetworkInterface {
	s.TxQueueSize = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) SetVSwitchId(v string) *RunInstancesRequestNetworkInterface {
	s.VSwitchId = &v
	return s
}

func (s *RunInstancesRequestNetworkInterface) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestNetworkOptions struct {
	BandwidthWeighting *string `json:"BandwidthWeighting,omitempty" xml:"BandwidthWeighting,omitempty"`
	// Specifies whether to enable the Jumbo Frames feature for the instance. Valid values:
	//
	// 	- false: does not enable the Jumbo Frames feature for the instance. The maximum transmission unit (MTU) value of all ENIs on the instance is set to 1500.
	//
	// 	- true: enables the Jumbo Frames feature for the instance. The MTU value of all ENIs on the instance is set to 8500.
	//
	// Default value: true.
	//
	// >  The Jumbo Frames feature is supported by only 8th-generation or later instance types. For more information, see [Jumbo Frames](https://help.aliyun.com/document_detail/200512.html).
	//
	// example:
	//
	// false
	EnableJumboFrame        *bool `json:"EnableJumboFrame,omitempty" xml:"EnableJumboFrame,omitempty"`
	EnableNetworkEncryption *bool `json:"EnableNetworkEncryption,omitempty" xml:"EnableNetworkEncryption,omitempty"`
}

func (s RunInstancesRequestNetworkOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestNetworkOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestNetworkOptions) GetBandwidthWeighting() *string {
	return s.BandwidthWeighting
}

func (s *RunInstancesRequestNetworkOptions) GetEnableJumboFrame() *bool {
	return s.EnableJumboFrame
}

func (s *RunInstancesRequestNetworkOptions) GetEnableNetworkEncryption() *bool {
	return s.EnableNetworkEncryption
}

func (s *RunInstancesRequestNetworkOptions) SetBandwidthWeighting(v string) *RunInstancesRequestNetworkOptions {
	s.BandwidthWeighting = &v
	return s
}

func (s *RunInstancesRequestNetworkOptions) SetEnableJumboFrame(v bool) *RunInstancesRequestNetworkOptions {
	s.EnableJumboFrame = &v
	return s
}

func (s *RunInstancesRequestNetworkOptions) SetEnableNetworkEncryption(v bool) *RunInstancesRequestNetworkOptions {
	s.EnableNetworkEncryption = &v
	return s
}

func (s *RunInstancesRequestNetworkOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestPrivateDnsNameOptions struct {
	// Specifies whether DNS Resolution from the Instance ID-based Hostname to the Instance Primary Private IPv6 Address (AAAA Record) is enabled. Valid values:
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
	EnableInstanceIdDnsAAAARecord *bool `json:"EnableInstanceIdDnsAAAARecord,omitempty" xml:"EnableInstanceIdDnsAAAARecord,omitempty"`
	// Specifies whether DNS Resolution from the Instance ID-based Hostname to the Instance Primary Private IPv4 Address (A Record) is enabled. Valid values:
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
	EnableInstanceIdDnsARecord *bool `json:"EnableInstanceIdDnsARecord,omitempty" xml:"EnableInstanceIdDnsARecord,omitempty"`
	// Specifies whether DNS Resolution from the IP Address-based Hostname to the Instance Primary Private IPv4 Address (A Record) is enabled. Valid values:
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
	EnableIpDnsARecord *bool `json:"EnableIpDnsARecord,omitempty" xml:"EnableIpDnsARecord,omitempty"`
	// Specifies whether Reverse DNS Resolution from the Instance Primary Private IPv4 Address to the IP Address-based Hostname (PTR Record) is enabled. Valid values:
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
	EnableIpDnsPtrRecord *bool `json:"EnableIpDnsPtrRecord,omitempty" xml:"EnableIpDnsPtrRecord,omitempty"`
	// The type of hostname. Valid values:
	//
	// 	- Custom: custom hostname
	//
	// 	- IpBased: IP address-based hostname
	//
	// 	- InstanceIdBased: instance ID-based hostname
	//
	// Default value: Custom.
	//
	// example:
	//
	// Custom
	HostnameType *string `json:"HostnameType,omitempty" xml:"HostnameType,omitempty"`
}

func (s RunInstancesRequestPrivateDnsNameOptions) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestPrivateDnsNameOptions) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestPrivateDnsNameOptions) GetEnableInstanceIdDnsAAAARecord() *bool {
	return s.EnableInstanceIdDnsAAAARecord
}

func (s *RunInstancesRequestPrivateDnsNameOptions) GetEnableInstanceIdDnsARecord() *bool {
	return s.EnableInstanceIdDnsARecord
}

func (s *RunInstancesRequestPrivateDnsNameOptions) GetEnableIpDnsARecord() *bool {
	return s.EnableIpDnsARecord
}

func (s *RunInstancesRequestPrivateDnsNameOptions) GetEnableIpDnsPtrRecord() *bool {
	return s.EnableIpDnsPtrRecord
}

func (s *RunInstancesRequestPrivateDnsNameOptions) GetHostnameType() *string {
	return s.HostnameType
}

func (s *RunInstancesRequestPrivateDnsNameOptions) SetEnableInstanceIdDnsAAAARecord(v bool) *RunInstancesRequestPrivateDnsNameOptions {
	s.EnableInstanceIdDnsAAAARecord = &v
	return s
}

func (s *RunInstancesRequestPrivateDnsNameOptions) SetEnableInstanceIdDnsARecord(v bool) *RunInstancesRequestPrivateDnsNameOptions {
	s.EnableInstanceIdDnsARecord = &v
	return s
}

func (s *RunInstancesRequestPrivateDnsNameOptions) SetEnableIpDnsARecord(v bool) *RunInstancesRequestPrivateDnsNameOptions {
	s.EnableIpDnsARecord = &v
	return s
}

func (s *RunInstancesRequestPrivateDnsNameOptions) SetEnableIpDnsPtrRecord(v bool) *RunInstancesRequestPrivateDnsNameOptions {
	s.EnableIpDnsPtrRecord = &v
	return s
}

func (s *RunInstancesRequestPrivateDnsNameOptions) SetHostnameType(v string) *RunInstancesRequestPrivateDnsNameOptions {
	s.HostnameType = &v
	return s
}

func (s *RunInstancesRequestPrivateDnsNameOptions) Validate() error {
	return dara.Validate(s)
}

type RunInstancesRequestTag struct {
	// The key of tag N to add to the instance, disks, and primary ENI. Valid values of N: 1 to 20. The tag key cannot be an empty string. The tag key can be up to 128 characters in length and cannot contain http:// or https://. The tag key cannot start with acs: or aliyun.
	//
	// example:
	//
	// TestKey
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The value of tag N to add to the instance, disks, and primary ENI. Valid values of N: 1 to 20. The tag value can be an empty string. The tag value can be up to 128 characters in length and cannot contain http:// or https://.
	//
	// example:
	//
	// TestValue
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s RunInstancesRequestTag) String() string {
	return dara.Prettify(s)
}

func (s RunInstancesRequestTag) GoString() string {
	return s.String()
}

func (s *RunInstancesRequestTag) GetKey() *string {
	return s.Key
}

func (s *RunInstancesRequestTag) GetValue() *string {
	return s.Value
}

func (s *RunInstancesRequestTag) SetKey(v string) *RunInstancesRequestTag {
	s.Key = &v
	return s
}

func (s *RunInstancesRequestTag) SetValue(v string) *RunInstancesRequestTag {
	s.Value = &v
	return s
}

func (s *RunInstancesRequestTag) Validate() error {
	return dara.Validate(s)
}
