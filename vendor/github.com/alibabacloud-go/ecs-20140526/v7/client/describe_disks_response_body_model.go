// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDisksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDisks(v *DescribeDisksResponseBodyDisks) *DescribeDisksResponseBody
	GetDisks() *DescribeDisksResponseBodyDisks
	SetNextToken(v string) *DescribeDisksResponseBody
	GetNextToken() *string
	SetPageNumber(v int32) *DescribeDisksResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *DescribeDisksResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *DescribeDisksResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *DescribeDisksResponseBody
	GetTotalCount() *int32
}

type DescribeDisksResponseBody struct {
	// Details about the disks.
	Disks *DescribeDisksResponseBodyDisks `json:"Disks,omitempty" xml:"Disks,omitempty" type:"Struct"`
	// The returned pagination token which can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// AAAAAdDWBF2****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use `NextToken` and `MaxResults` for a paged query.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// >  This parameter will be removed in the future. We recommend that you use `NextToken` and `MaxResults` for a paged query.
	//
	// example:
	//
	// 1
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// > When using the `MaxResults` and `NextToken` parameters for a paginated query, the returned `TotalCount` parameter value is invalid.
	//
	// example:
	//
	// 15
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeDisksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBody) GetDisks() *DescribeDisksResponseBodyDisks {
	return s.Disks
}

func (s *DescribeDisksResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeDisksResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *DescribeDisksResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *DescribeDisksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDisksResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *DescribeDisksResponseBody) SetDisks(v *DescribeDisksResponseBodyDisks) *DescribeDisksResponseBody {
	s.Disks = v
	return s
}

func (s *DescribeDisksResponseBody) SetNextToken(v string) *DescribeDisksResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeDisksResponseBody) SetPageNumber(v int32) *DescribeDisksResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeDisksResponseBody) SetPageSize(v int32) *DescribeDisksResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeDisksResponseBody) SetRequestId(v string) *DescribeDisksResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDisksResponseBody) SetTotalCount(v int32) *DescribeDisksResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeDisksResponseBody) Validate() error {
	if s.Disks != nil {
		if err := s.Disks.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDisksResponseBodyDisks struct {
	Disk []*DescribeDisksResponseBodyDisksDisk `json:"Disk,omitempty" xml:"Disk,omitempty" type:"Repeated"`
}

func (s DescribeDisksResponseBodyDisks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisks) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisks) GetDisk() []*DescribeDisksResponseBodyDisksDisk {
	return s.Disk
}

func (s *DescribeDisksResponseBodyDisks) SetDisk(v []*DescribeDisksResponseBodyDisksDisk) *DescribeDisksResponseBodyDisks {
	s.Disk = v
	return s
}

func (s *DescribeDisksResponseBodyDisks) Validate() error {
	if s.Disk != nil {
		for _, item := range s.Disk {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDisksResponseBodyDisksDisk struct {
	// The time when the disk was last attached. The time follows the ISO 8601 standard in the yyyy-MM-ddThh:mmZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-07T06:08:56Z
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// The attachment information of the disk. The value is an array that consists of the `Attachment` values. This value is not returned when you query Shared Block Storage devices.
	Attachments *DescribeDisksResponseBodyDisksDiskAttachments `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Struct"`
	// The ID of the automatic snapshot policy that is applied to the cloud disk.
	//
	// example:
	//
	// sp-bp67acfmxazb4p****
	AutoSnapshotPolicyId *string `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	// This parameter is in invitational preview and is not publicly available.
	//
	// example:
	//
	// null
	BdfId *string `json:"BdfId,omitempty" xml:"BdfId,omitempty"`
	// Indicates whether the performance burst feature is enabled. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// This parameter is available only if you set `Category` to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// false
	BurstingEnabled *bool `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	// The category of the disk. Valid values:
	//
	// 	- cloud: basic disk
	//
	// 	- cloud_efficiency: ultra disk
	//
	// 	- cloud_ssd: standard SSD
	//
	// 	- cloud_essd: ESSD
	//
	// 	- cloud_auto: ESSD AutoPL disk
	//
	// 	- local_ssd_pro: I/O-intensive local disk
	//
	// 	- local_hdd_pro: throughput-intensive local disk
	//
	// 	- cloud_essd_entry: ESSD Entry disk
	//
	// 	- elastic_ephemeral_disk_standard: standard elastic ephemeral disk
	//
	// 	- elastic_ephemeral_disk_premium: premium static ephemeral disk
	//
	// 	- ephemeral: retired local disk
	//
	// 	- ephemeral_ssd: retired local SSD
	//
	// example:
	//
	// cloud_ssd
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the disk was created.
	//
	// example:
	//
	// 2021-06-07T06:08:54Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// Indicates whether the automatic snapshots of the cloud disk are deleted when the cloud disk is released. Valid values:
	//
	// 	- true: The automatic snapshots of the cloud disk are deleted when the disk is released.
	//
	// 	- false: The automatic snapshots of the cloud disk are retained when the disk is released.
	//
	// Snapshots that were created in the ECS console or by calling the [CreateSnapshot](https://help.aliyun.com/document_detail/25524.html) operation are retained and not affected by this parameter.
	//
	// example:
	//
	// false
	DeleteAutoSnapshot *bool `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	// Indicates whether the disk is released when the instance to which the disk is attached is released. Valid values:
	//
	// 	- true: The disk is released when the associated instance is released.
	//
	// 	- false: The disk is retained when the associated instance is released.
	//
	// example:
	//
	// true
	DeleteWithInstance *bool `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	// The description of the disk.
	//
	// example:
	//
	// testDescription
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The time when the disk was last detached.
	//
	// example:
	//
	// 2021-06-07T21:01:22Z
	DetachedTime *string `json:"DetachedTime,omitempty" xml:"DetachedTime,omitempty"`
	// The device name of the disk on the instance to which the disk is attached. Example: /dev/xvdb. Take note of the following items:
	//
	// 	- This parameter has a value only when the `Status` value is `In_use` or `Detaching`.
	//
	// 	- This parameter is empty for cloud disks for which the multi-attach feature is enabled. You can query the attachment information of the cloud disk based on the returned list of `Attachment` objects.
	//
	// >  This parameter will be removed in the future. We recommend that you use other parameters to ensure future compatibility.
	//
	// example:
	//
	// /dev/xvdb
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The billing method of the disk. Valid values:
	//
	// 	- PrePaid: subscription
	//
	// 	- PostPaid: pay-as-you-go
	//
	// example:
	//
	// PrePaid
	DiskChargeType *string `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	// The ID of the disk.
	//
	// example:
	//
	// d-bp18um4r4f2fve24****
	DiskId *string `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	// The name of the disk.
	//
	// example:
	//
	// testDiskName
	DiskName *string `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	// Indicates whether the automatic snapshot policy feature is enabled for the cloud disk.
	//
	// >  This parameter is deprecated. By default, the automatic snapshot policy feature is enabled for cloud disks. You need to only apply an automatic snapshot policy to a cloud disk before you can use the automatic snapshot policy.
	//
	// example:
	//
	// false
	EnableAutoSnapshot *bool `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	// Indicates whether an automatic snapshot policy is applied to the cloud disk.
	//
	// example:
	//
	// false
	EnableAutomatedSnapshotPolicy *bool `json:"EnableAutomatedSnapshotPolicy,omitempty" xml:"EnableAutomatedSnapshotPolicy,omitempty"`
	// Indicates whether the cloud disk is encrypted.
	//
	// example:
	//
	// false
	Encrypted *bool `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	// The time when the subscription disk expires.
	//
	// example:
	//
	// 2021-07-07T16:00Z
	ExpiredTime *string `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	// The maximum number of read and write operations per second.
	//
	// example:
	//
	// 4000
	IOPS *int32 `json:"IOPS,omitempty" xml:"IOPS,omitempty"`
	// The maximum number of read operations per second.
	//
	// example:
	//
	// 2000
	IOPSRead *int32 `json:"IOPSRead,omitempty" xml:"IOPSRead,omitempty"`
	// The maximum number of write operations per second.
	//
	// example:
	//
	// 2000
	IOPSWrite *int32 `json:"IOPSWrite,omitempty" xml:"IOPSWrite,omitempty"`
	// The ID of the image that was used to create the instance. This parameter is empty unless the cloud disk was created from an image. The value of this parameter remains unchanged throughout the lifecycle of the cloud disk.
	//
	// example:
	//
	// m-bp13aqm171qynt3u***
	ImageId *string `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	// The ID of the instance to which the disk is attached. Take note of the following items:
	//
	// 	- This parameter has a value only when the `Status` value is `In_use` or `Detaching`.
	//
	// 	- This parameter is empty for cloud disks for which the multi-attach feature is enabled. You can query the attachment information of the cloud disk based on the returned `Attachment` objects.
	//
	// example:
	//
	// i-bp67acfmxazb4q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the KMS key that is used for the cloud disk.
	//
	// example:
	//
	// 0e478b7a-4262-4802-b8cb-00d3fb408***
	KMSKeyId *string `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	// The number of instances to which the Shared Block Storage device is attached.
	//
	// example:
	//
	// 1
	MountInstanceNum *int32 `json:"MountInstanceNum,omitempty" xml:"MountInstanceNum,omitempty"`
	// The attachment information of the Shared Block Storage device.
	MountInstances *DescribeDisksResponseBodyDisksDiskMountInstances `json:"MountInstances,omitempty" xml:"MountInstances,omitempty" type:"Struct"`
	// Indicates whether the multi-attach feature is enabled for the cloud disk.
	//
	// example:
	//
	// Disabled
	MultiAttach *string `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	// The reasons why the disk was locked.
	OperationLocks *DescribeDisksResponseBodyDisksDiskOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	// The performance level of the ESSD. Valid values:
	//
	// 	- PL0: A single ESSD can deliver up to 10,000 random read/write IOPS.
	//
	// 	- PL1: A single ESSD can deliver up to 50,000 random read/write IOPS.
	//
	// 	- PL2: A single ESSD can deliver up to 100,000 random read/write IOPS.
	//
	// 	- PL3: A single ESSD can deliver up to 1,000,000 random read/write IOPS.
	//
	// example:
	//
	// PL0
	PerformanceLevel *string `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	// The locations in which data is stored.
	//
	// This parameter is returned only if you specify `Placement` in the AdditionalAttributes.N request parameter.
	//
	// >  This parameter is valid only for Regional ESSDs (cloud_regional_disk_auto).
	Placement *DescribeDisksResponseBodyDisksDiskPlacement `json:"Placement,omitempty" xml:"Placement,omitempty" type:"Struct"`
	// Indicates whether the disk is removable.
	//
	// example:
	//
	// false
	Portable *bool `json:"Portable,omitempty" xml:"Portable,omitempty"`
	// The product code of the disk in Alibaba Cloud Marketplace.
	//
	// example:
	//
	// jxsc000204
	ProductCode *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	// The provisioned read/write IOPS of the ESSD AutoPL disk. Valid values: 0 to min{50,000, 1,000 × *Capacity - Baseline IOPS}. Baseline IOPS = min{1,800 + 50 × *Capacity, 50,000}
	//
	// This parameter is available only if you set `Category` to `cloud_auto`. For more information, see [ESSD AutoPL disks](https://help.aliyun.com/document_detail/368372.html).
	//
	// example:
	//
	// 40000
	ProvisionedIops *int64 `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	// The ID of the region to which the disk belongs.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The ID of the resource group to which the disk belongs.
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The serial number of the disk.
	//
	// example:
	//
	// bp18um4r4f2fve2****
	SerialNumber *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	// The size of the disk. Unit: GiB.
	//
	// example:
	//
	// 60
	Size *int32 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The ID of the snapshot that was used to create the cloud disk.
	//
	// This parameter is empty unless the cloud disk was created from a snapshot. The value of this parameter remains unchanged throughout the lifecycle of the cloud disk.
	//
	// example:
	//
	// s-bp67acfmxazb4p****
	SourceSnapshotId *string `json:"SourceSnapshotId,omitempty" xml:"SourceSnapshotId,omitempty"`
	// The status of the disk. Valid values:
	//
	// 	- In_use
	//
	// 	- Available
	//
	// 	- Attaching
	//
	// 	- Detaching
	//
	// 	- Creating
	//
	// 	- ReIniting
	//
	// example:
	//
	// In_use
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The ID of the dedicated block storage cluster to which the cloud disk belongs. If your cloud disk belongs to the public block storage cluster, an empty value is returned.
	//
	// example:
	//
	// dbsc-j5e1sf2vaf5he8m2****
	StorageClusterId *string `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
	// The ID of the storage set.
	//
	// example:
	//
	// ss-i-bp1j4i2jdf3owlhe****
	StorageSetId *string `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	// The maximum number of partitions in the storage set.
	//
	// example:
	//
	// 11
	StorageSetPartitionNumber *int32 `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	// The tags of the disk.
	Tags *DescribeDisksResponseBodyDisksDiskTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	// The amount of data that can be transferred per second. Unit: MB/s.
	//
	// example:
	//
	// 100
	Throughput *int32 `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	// The amount of data that can be read per second. Unit: MB/s.
	//
	// example:
	//
	// 100
	ThroughputRead *int32 `json:"ThroughputRead,omitempty" xml:"ThroughputRead,omitempty"`
	// The amount of data that can be written per second. Unit: MB/s.
	//
	// example:
	//
	// 100
	ThroughputWrite *int32 `json:"ThroughputWrite,omitempty" xml:"ThroughputWrite,omitempty"`
	// The type of the disk. Valid values:
	//
	// 	- system: system disk
	//
	// 	- data: data disk
	//
	// example:
	//
	// system
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The ID of the zone to which the disk belongs.
	//
	// example:
	//
	// cn-hangzhou-i
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeDisksResponseBodyDisksDisk) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDisk) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDisk) GetAttachedTime() *string {
	return s.AttachedTime
}

func (s *DescribeDisksResponseBodyDisksDisk) GetAttachments() *DescribeDisksResponseBodyDisksDiskAttachments {
	return s.Attachments
}

func (s *DescribeDisksResponseBodyDisksDisk) GetAutoSnapshotPolicyId() *string {
	return s.AutoSnapshotPolicyId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetBdfId() *string {
	return s.BdfId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetBurstingEnabled() *bool {
	return s.BurstingEnabled
}

func (s *DescribeDisksResponseBodyDisksDisk) GetCategory() *string {
	return s.Category
}

func (s *DescribeDisksResponseBodyDisksDisk) GetCreationTime() *string {
	return s.CreationTime
}

func (s *DescribeDisksResponseBodyDisksDisk) GetDeleteAutoSnapshot() *bool {
	return s.DeleteAutoSnapshot
}

func (s *DescribeDisksResponseBodyDisksDisk) GetDeleteWithInstance() *bool {
	return s.DeleteWithInstance
}

func (s *DescribeDisksResponseBodyDisksDisk) GetDescription() *string {
	return s.Description
}

func (s *DescribeDisksResponseBodyDisksDisk) GetDetachedTime() *string {
	return s.DetachedTime
}

func (s *DescribeDisksResponseBodyDisksDisk) GetDevice() *string {
	return s.Device
}

func (s *DescribeDisksResponseBodyDisksDisk) GetDiskChargeType() *string {
	return s.DiskChargeType
}

func (s *DescribeDisksResponseBodyDisksDisk) GetDiskId() *string {
	return s.DiskId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetDiskName() *string {
	return s.DiskName
}

func (s *DescribeDisksResponseBodyDisksDisk) GetEnableAutoSnapshot() *bool {
	return s.EnableAutoSnapshot
}

func (s *DescribeDisksResponseBodyDisksDisk) GetEnableAutomatedSnapshotPolicy() *bool {
	return s.EnableAutomatedSnapshotPolicy
}

func (s *DescribeDisksResponseBodyDisksDisk) GetEncrypted() *bool {
	return s.Encrypted
}

func (s *DescribeDisksResponseBodyDisksDisk) GetExpiredTime() *string {
	return s.ExpiredTime
}

func (s *DescribeDisksResponseBodyDisksDisk) GetIOPS() *int32 {
	return s.IOPS
}

func (s *DescribeDisksResponseBodyDisksDisk) GetIOPSRead() *int32 {
	return s.IOPSRead
}

func (s *DescribeDisksResponseBodyDisksDisk) GetIOPSWrite() *int32 {
	return s.IOPSWrite
}

func (s *DescribeDisksResponseBodyDisksDisk) GetImageId() *string {
	return s.ImageId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetKMSKeyId() *string {
	return s.KMSKeyId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetMountInstanceNum() *int32 {
	return s.MountInstanceNum
}

func (s *DescribeDisksResponseBodyDisksDisk) GetMountInstances() *DescribeDisksResponseBodyDisksDiskMountInstances {
	return s.MountInstances
}

func (s *DescribeDisksResponseBodyDisksDisk) GetMultiAttach() *string {
	return s.MultiAttach
}

func (s *DescribeDisksResponseBodyDisksDisk) GetOperationLocks() *DescribeDisksResponseBodyDisksDiskOperationLocks {
	return s.OperationLocks
}

func (s *DescribeDisksResponseBodyDisksDisk) GetPerformanceLevel() *string {
	return s.PerformanceLevel
}

func (s *DescribeDisksResponseBodyDisksDisk) GetPlacement() *DescribeDisksResponseBodyDisksDiskPlacement {
	return s.Placement
}

func (s *DescribeDisksResponseBodyDisksDisk) GetPortable() *bool {
	return s.Portable
}

func (s *DescribeDisksResponseBodyDisksDisk) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeDisksResponseBodyDisksDisk) GetProvisionedIops() *int64 {
	return s.ProvisionedIops
}

func (s *DescribeDisksResponseBodyDisksDisk) GetRegionId() *string {
	return s.RegionId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeDisksResponseBodyDisksDisk) GetSize() *int32 {
	return s.Size
}

func (s *DescribeDisksResponseBodyDisksDisk) GetSourceSnapshotId() *string {
	return s.SourceSnapshotId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetStatus() *string {
	return s.Status
}

func (s *DescribeDisksResponseBodyDisksDisk) GetStorageClusterId() *string {
	return s.StorageClusterId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetStorageSetId() *string {
	return s.StorageSetId
}

func (s *DescribeDisksResponseBodyDisksDisk) GetStorageSetPartitionNumber() *int32 {
	return s.StorageSetPartitionNumber
}

func (s *DescribeDisksResponseBodyDisksDisk) GetTags() *DescribeDisksResponseBodyDisksDiskTags {
	return s.Tags
}

func (s *DescribeDisksResponseBodyDisksDisk) GetThroughput() *int32 {
	return s.Throughput
}

func (s *DescribeDisksResponseBodyDisksDisk) GetThroughputRead() *int32 {
	return s.ThroughputRead
}

func (s *DescribeDisksResponseBodyDisksDisk) GetThroughputWrite() *int32 {
	return s.ThroughputWrite
}

func (s *DescribeDisksResponseBodyDisksDisk) GetType() *string {
	return s.Type
}

func (s *DescribeDisksResponseBodyDisksDisk) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeDisksResponseBodyDisksDisk) SetAttachedTime(v string) *DescribeDisksResponseBodyDisksDisk {
	s.AttachedTime = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetAttachments(v *DescribeDisksResponseBodyDisksDiskAttachments) *DescribeDisksResponseBodyDisksDisk {
	s.Attachments = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetAutoSnapshotPolicyId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.AutoSnapshotPolicyId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetBdfId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.BdfId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetBurstingEnabled(v bool) *DescribeDisksResponseBodyDisksDisk {
	s.BurstingEnabled = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetCategory(v string) *DescribeDisksResponseBodyDisksDisk {
	s.Category = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetCreationTime(v string) *DescribeDisksResponseBodyDisksDisk {
	s.CreationTime = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetDeleteAutoSnapshot(v bool) *DescribeDisksResponseBodyDisksDisk {
	s.DeleteAutoSnapshot = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetDeleteWithInstance(v bool) *DescribeDisksResponseBodyDisksDisk {
	s.DeleteWithInstance = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetDescription(v string) *DescribeDisksResponseBodyDisksDisk {
	s.Description = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetDetachedTime(v string) *DescribeDisksResponseBodyDisksDisk {
	s.DetachedTime = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetDevice(v string) *DescribeDisksResponseBodyDisksDisk {
	s.Device = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetDiskChargeType(v string) *DescribeDisksResponseBodyDisksDisk {
	s.DiskChargeType = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetDiskId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.DiskId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetDiskName(v string) *DescribeDisksResponseBodyDisksDisk {
	s.DiskName = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetEnableAutoSnapshot(v bool) *DescribeDisksResponseBodyDisksDisk {
	s.EnableAutoSnapshot = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetEnableAutomatedSnapshotPolicy(v bool) *DescribeDisksResponseBodyDisksDisk {
	s.EnableAutomatedSnapshotPolicy = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetEncrypted(v bool) *DescribeDisksResponseBodyDisksDisk {
	s.Encrypted = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetExpiredTime(v string) *DescribeDisksResponseBodyDisksDisk {
	s.ExpiredTime = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetIOPS(v int32) *DescribeDisksResponseBodyDisksDisk {
	s.IOPS = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetIOPSRead(v int32) *DescribeDisksResponseBodyDisksDisk {
	s.IOPSRead = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetIOPSWrite(v int32) *DescribeDisksResponseBodyDisksDisk {
	s.IOPSWrite = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetImageId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.ImageId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetInstanceId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.InstanceId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetKMSKeyId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.KMSKeyId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetMountInstanceNum(v int32) *DescribeDisksResponseBodyDisksDisk {
	s.MountInstanceNum = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetMountInstances(v *DescribeDisksResponseBodyDisksDiskMountInstances) *DescribeDisksResponseBodyDisksDisk {
	s.MountInstances = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetMultiAttach(v string) *DescribeDisksResponseBodyDisksDisk {
	s.MultiAttach = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetOperationLocks(v *DescribeDisksResponseBodyDisksDiskOperationLocks) *DescribeDisksResponseBodyDisksDisk {
	s.OperationLocks = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetPerformanceLevel(v string) *DescribeDisksResponseBodyDisksDisk {
	s.PerformanceLevel = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetPlacement(v *DescribeDisksResponseBodyDisksDiskPlacement) *DescribeDisksResponseBodyDisksDisk {
	s.Placement = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetPortable(v bool) *DescribeDisksResponseBodyDisksDisk {
	s.Portable = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetProductCode(v string) *DescribeDisksResponseBodyDisksDisk {
	s.ProductCode = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetProvisionedIops(v int64) *DescribeDisksResponseBodyDisksDisk {
	s.ProvisionedIops = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetRegionId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.RegionId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetResourceGroupId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.ResourceGroupId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetSerialNumber(v string) *DescribeDisksResponseBodyDisksDisk {
	s.SerialNumber = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetSize(v int32) *DescribeDisksResponseBodyDisksDisk {
	s.Size = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetSourceSnapshotId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.SourceSnapshotId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetStatus(v string) *DescribeDisksResponseBodyDisksDisk {
	s.Status = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetStorageClusterId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.StorageClusterId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetStorageSetId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.StorageSetId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetStorageSetPartitionNumber(v int32) *DescribeDisksResponseBodyDisksDisk {
	s.StorageSetPartitionNumber = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetTags(v *DescribeDisksResponseBodyDisksDiskTags) *DescribeDisksResponseBodyDisksDisk {
	s.Tags = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetThroughput(v int32) *DescribeDisksResponseBodyDisksDisk {
	s.Throughput = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetThroughputRead(v int32) *DescribeDisksResponseBodyDisksDisk {
	s.ThroughputRead = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetThroughputWrite(v int32) *DescribeDisksResponseBodyDisksDisk {
	s.ThroughputWrite = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetType(v string) *DescribeDisksResponseBodyDisksDisk {
	s.Type = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) SetZoneId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.ZoneId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDisk) Validate() error {
	if s.Attachments != nil {
		if err := s.Attachments.Validate(); err != nil {
			return err
		}
	}
	if s.MountInstances != nil {
		if err := s.MountInstances.Validate(); err != nil {
			return err
		}
	}
	if s.OperationLocks != nil {
		if err := s.OperationLocks.Validate(); err != nil {
			return err
		}
	}
	if s.Placement != nil {
		if err := s.Placement.Validate(); err != nil {
			return err
		}
	}
	if s.Tags != nil {
		if err := s.Tags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDisksResponseBodyDisksDiskAttachments struct {
	Attachment []*DescribeDisksResponseBodyDisksDiskAttachmentsAttachment `json:"Attachment,omitempty" xml:"Attachment,omitempty" type:"Repeated"`
}

func (s DescribeDisksResponseBodyDisksDiskAttachments) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDiskAttachments) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDiskAttachments) GetAttachment() []*DescribeDisksResponseBodyDisksDiskAttachmentsAttachment {
	return s.Attachment
}

func (s *DescribeDisksResponseBodyDisksDiskAttachments) SetAttachment(v []*DescribeDisksResponseBodyDisksDiskAttachmentsAttachment) *DescribeDisksResponseBodyDisksDiskAttachments {
	s.Attachment = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskAttachments) Validate() error {
	if s.Attachment != nil {
		for _, item := range s.Attachment {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDisksResponseBodyDisksDiskAttachmentsAttachment struct {
	// The time when the disk was attached. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-06-07T06:08:56Z
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// The device name of the disk.
	//
	// example:
	//
	// /dev/xvda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The ID of the instance to which the disk is attached.
	//
	// example:
	//
	// i-bp67acfmxazb4q****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDisksResponseBodyDisksDiskAttachmentsAttachment) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDiskAttachmentsAttachment) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDiskAttachmentsAttachment) GetAttachedTime() *string {
	return s.AttachedTime
}

func (s *DescribeDisksResponseBodyDisksDiskAttachmentsAttachment) GetDevice() *string {
	return s.Device
}

func (s *DescribeDisksResponseBodyDisksDiskAttachmentsAttachment) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDisksResponseBodyDisksDiskAttachmentsAttachment) SetAttachedTime(v string) *DescribeDisksResponseBodyDisksDiskAttachmentsAttachment {
	s.AttachedTime = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskAttachmentsAttachment) SetDevice(v string) *DescribeDisksResponseBodyDisksDiskAttachmentsAttachment {
	s.Device = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskAttachmentsAttachment) SetInstanceId(v string) *DescribeDisksResponseBodyDisksDiskAttachmentsAttachment {
	s.InstanceId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskAttachmentsAttachment) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksResponseBodyDisksDiskMountInstances struct {
	MountInstance []*DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance `json:"MountInstance,omitempty" xml:"MountInstance,omitempty" type:"Repeated"`
}

func (s DescribeDisksResponseBodyDisksDiskMountInstances) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDiskMountInstances) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDiskMountInstances) GetMountInstance() []*DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance {
	return s.MountInstance
}

func (s *DescribeDisksResponseBodyDisksDiskMountInstances) SetMountInstance(v []*DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance) *DescribeDisksResponseBodyDisksDiskMountInstances {
	s.MountInstance = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskMountInstances) Validate() error {
	if s.MountInstance != nil {
		for _, item := range s.MountInstance {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance struct {
	// The time when the disk was attached. The time follows the [ISO 8601](https://help.aliyun.com/document_detail/25696.html) standard in the yyyy-MM-ddTHH:mm:ssZ format. The time is displayed in UTC.
	//
	// example:
	//
	// 2017-12-05T2340:00Z
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	// The mount point of the disk.
	//
	// example:
	//
	// /dev/xvda
	Device *string `json:"Device,omitempty" xml:"Device,omitempty"`
	// The ID of the instance to which the disk is attached.
	//
	// example:
	//
	// i-bp1j4i2jdf3owlhe****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance) GetAttachedTime() *string {
	return s.AttachedTime
}

func (s *DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance) GetDevice() *string {
	return s.Device
}

func (s *DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance) SetAttachedTime(v string) *DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance {
	s.AttachedTime = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance) SetDevice(v string) *DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance {
	s.Device = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance) SetInstanceId(v string) *DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance {
	s.InstanceId = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskMountInstancesMountInstance) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksResponseBodyDisksDiskOperationLocks struct {
	OperationLock []*DescribeDisksResponseBodyDisksDiskOperationLocksOperationLock `json:"OperationLock,omitempty" xml:"OperationLock,omitempty" type:"Repeated"`
}

func (s DescribeDisksResponseBodyDisksDiskOperationLocks) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDiskOperationLocks) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDiskOperationLocks) GetOperationLock() []*DescribeDisksResponseBodyDisksDiskOperationLocksOperationLock {
	return s.OperationLock
}

func (s *DescribeDisksResponseBodyDisksDiskOperationLocks) SetOperationLock(v []*DescribeDisksResponseBodyDisksDiskOperationLocksOperationLock) *DescribeDisksResponseBodyDisksDiskOperationLocks {
	s.OperationLock = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskOperationLocks) Validate() error {
	if s.OperationLock != nil {
		for _, item := range s.OperationLock {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDisksResponseBodyDisksDiskOperationLocksOperationLock struct {
	// The reason why the disk was locked.
	//
	// example:
	//
	// security
	LockReason *string `json:"LockReason,omitempty" xml:"LockReason,omitempty"`
}

func (s DescribeDisksResponseBodyDisksDiskOperationLocksOperationLock) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDiskOperationLocksOperationLock) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDiskOperationLocksOperationLock) GetLockReason() *string {
	return s.LockReason
}

func (s *DescribeDisksResponseBodyDisksDiskOperationLocksOperationLock) SetLockReason(v string) *DescribeDisksResponseBodyDisksDiskOperationLocksOperationLock {
	s.LockReason = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskOperationLocksOperationLock) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksResponseBodyDisksDiskPlacement struct {
	// The IDs of the zones in which data is stored.
	//
	// example:
	//
	// cn-hangzhou-b
	//
	// cn-hangzhou-j
	ZoneIds *string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty"`
}

func (s DescribeDisksResponseBodyDisksDiskPlacement) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDiskPlacement) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDiskPlacement) GetZoneIds() *string {
	return s.ZoneIds
}

func (s *DescribeDisksResponseBodyDisksDiskPlacement) SetZoneIds(v string) *DescribeDisksResponseBodyDisksDiskPlacement {
	s.ZoneIds = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskPlacement) Validate() error {
	return dara.Validate(s)
}

type DescribeDisksResponseBodyDisksDiskTags struct {
	Tag []*DescribeDisksResponseBodyDisksDiskTagsTag `json:"Tag,omitempty" xml:"Tag,omitempty" type:"Repeated"`
}

func (s DescribeDisksResponseBodyDisksDiskTags) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDiskTags) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDiskTags) GetTag() []*DescribeDisksResponseBodyDisksDiskTagsTag {
	return s.Tag
}

func (s *DescribeDisksResponseBodyDisksDiskTags) SetTag(v []*DescribeDisksResponseBodyDisksDiskTagsTag) *DescribeDisksResponseBodyDisksDiskTags {
	s.Tag = v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskTags) Validate() error {
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

type DescribeDisksResponseBodyDisksDiskTagsTag struct {
	// The tag key of the disk.
	//
	// example:
	//
	// TestKey
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The tag value of the disk.
	//
	// example:
	//
	// TestValue
	TagValue *string `json:"TagValue,omitempty" xml:"TagValue,omitempty"`
}

func (s DescribeDisksResponseBodyDisksDiskTagsTag) String() string {
	return dara.Prettify(s)
}

func (s DescribeDisksResponseBodyDisksDiskTagsTag) GoString() string {
	return s.String()
}

func (s *DescribeDisksResponseBodyDisksDiskTagsTag) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeDisksResponseBodyDisksDiskTagsTag) GetTagValue() *string {
	return s.TagValue
}

func (s *DescribeDisksResponseBodyDisksDiskTagsTag) SetTagKey(v string) *DescribeDisksResponseBodyDisksDiskTagsTag {
	s.TagKey = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskTagsTag) SetTagValue(v string) *DescribeDisksResponseBodyDisksDiskTagsTag {
	s.TagValue = &v
	return s
}

func (s *DescribeDisksResponseBodyDisksDiskTagsTag) Validate() error {
	return dara.Validate(s)
}
