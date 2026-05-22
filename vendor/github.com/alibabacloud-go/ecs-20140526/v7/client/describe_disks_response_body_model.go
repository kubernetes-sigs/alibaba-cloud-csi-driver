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
	AttachedTime                  *string                                           `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	Attachments                   *DescribeDisksResponseBodyDisksDiskAttachments    `json:"Attachments,omitempty" xml:"Attachments,omitempty" type:"Struct"`
	AutoSnapshotPolicyId          *string                                           `json:"AutoSnapshotPolicyId,omitempty" xml:"AutoSnapshotPolicyId,omitempty"`
	BdfId                         *string                                           `json:"BdfId,omitempty" xml:"BdfId,omitempty"`
	BurstingEnabled               *bool                                             `json:"BurstingEnabled,omitempty" xml:"BurstingEnabled,omitempty"`
	Category                      *string                                           `json:"Category,omitempty" xml:"Category,omitempty"`
	CreationTime                  *string                                           `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	DeleteAutoSnapshot            *bool                                             `json:"DeleteAutoSnapshot,omitempty" xml:"DeleteAutoSnapshot,omitempty"`
	DeleteWithInstance            *bool                                             `json:"DeleteWithInstance,omitempty" xml:"DeleteWithInstance,omitempty"`
	Description                   *string                                           `json:"Description,omitempty" xml:"Description,omitempty"`
	DetachedTime                  *string                                           `json:"DetachedTime,omitempty" xml:"DetachedTime,omitempty"`
	Device                        *string                                           `json:"Device,omitempty" xml:"Device,omitempty"`
	DiskChargeType                *string                                           `json:"DiskChargeType,omitempty" xml:"DiskChargeType,omitempty"`
	DiskId                        *string                                           `json:"DiskId,omitempty" xml:"DiskId,omitempty"`
	DiskName                      *string                                           `json:"DiskName,omitempty" xml:"DiskName,omitempty"`
	EnableAutoSnapshot            *bool                                             `json:"EnableAutoSnapshot,omitempty" xml:"EnableAutoSnapshot,omitempty"`
	EnableAutomatedSnapshotPolicy *bool                                             `json:"EnableAutomatedSnapshotPolicy,omitempty" xml:"EnableAutomatedSnapshotPolicy,omitempty"`
	Encrypted                     *bool                                             `json:"Encrypted,omitempty" xml:"Encrypted,omitempty"`
	ExpiredTime                   *string                                           `json:"ExpiredTime,omitempty" xml:"ExpiredTime,omitempty"`
	IOPS                          *int32                                            `json:"IOPS,omitempty" xml:"IOPS,omitempty"`
	IOPSRead                      *int32                                            `json:"IOPSRead,omitempty" xml:"IOPSRead,omitempty"`
	IOPSWrite                     *int32                                            `json:"IOPSWrite,omitempty" xml:"IOPSWrite,omitempty"`
	ImageId                       *string                                           `json:"ImageId,omitempty" xml:"ImageId,omitempty"`
	InstanceId                    *string                                           `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	KMSKeyId                      *string                                           `json:"KMSKeyId,omitempty" xml:"KMSKeyId,omitempty"`
	MountInstanceNum              *int32                                            `json:"MountInstanceNum,omitempty" xml:"MountInstanceNum,omitempty"`
	MountInstances                *DescribeDisksResponseBodyDisksDiskMountInstances `json:"MountInstances,omitempty" xml:"MountInstances,omitempty" type:"Struct"`
	MultiAttach                   *string                                           `json:"MultiAttach,omitempty" xml:"MultiAttach,omitempty"`
	OperationLocks                *DescribeDisksResponseBodyDisksDiskOperationLocks `json:"OperationLocks,omitempty" xml:"OperationLocks,omitempty" type:"Struct"`
	PerformanceLevel              *string                                           `json:"PerformanceLevel,omitempty" xml:"PerformanceLevel,omitempty"`
	Placement                     *DescribeDisksResponseBodyDisksDiskPlacement      `json:"Placement,omitempty" xml:"Placement,omitempty" type:"Struct"`
	Portable                      *bool                                             `json:"Portable,omitempty" xml:"Portable,omitempty"`
	ProductCode                   *string                                           `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProvisionedIops               *int64                                            `json:"ProvisionedIops,omitempty" xml:"ProvisionedIops,omitempty"`
	RegionId                      *string                                           `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	ResourceGroupId               *string                                           `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	SerialNumber                  *string                                           `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	Size                          *int32                                            `json:"Size,omitempty" xml:"Size,omitempty"`
	// example:
	//
	// d-123*********
	SourceDiskId              *string                                 `json:"SourceDiskId,omitempty" xml:"SourceDiskId,omitempty"`
	SourceSnapshotId          *string                                 `json:"SourceSnapshotId,omitempty" xml:"SourceSnapshotId,omitempty"`
	Status                    *string                                 `json:"Status,omitempty" xml:"Status,omitempty"`
	StorageClusterId          *string                                 `json:"StorageClusterId,omitempty" xml:"StorageClusterId,omitempty"`
	StorageSetId              *string                                 `json:"StorageSetId,omitempty" xml:"StorageSetId,omitempty"`
	StorageSetPartitionNumber *int32                                  `json:"StorageSetPartitionNumber,omitempty" xml:"StorageSetPartitionNumber,omitempty"`
	Tags                      *DescribeDisksResponseBodyDisksDiskTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Struct"`
	Throughput                *int32                                  `json:"Throughput,omitempty" xml:"Throughput,omitempty"`
	ThroughputRead            *int32                                  `json:"ThroughputRead,omitempty" xml:"ThroughputRead,omitempty"`
	ThroughputWrite           *int32                                  `json:"ThroughputWrite,omitempty" xml:"ThroughputWrite,omitempty"`
	Type                      *string                                 `json:"Type,omitempty" xml:"Type,omitempty"`
	ZoneId                    *string                                 `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
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

func (s *DescribeDisksResponseBodyDisksDisk) GetSourceDiskId() *string {
	return s.SourceDiskId
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

func (s *DescribeDisksResponseBodyDisksDisk) SetSourceDiskId(v string) *DescribeDisksResponseBodyDisksDisk {
	s.SourceDiskId = &v
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
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	Device       *string `json:"Device,omitempty" xml:"Device,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	AttachedTime *string `json:"AttachedTime,omitempty" xml:"AttachedTime,omitempty"`
	Device       *string `json:"Device,omitempty" xml:"Device,omitempty"`
	InstanceId   *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
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
	TagKey   *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
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
