package disk

import (
	"fmt"
)

const (

	// ESSD_PERFORMANCE_LEVEL is storage class
	ESSD_PERFORMANCE_LEVEL = "performanceLevel"
	// DISKTAGKEY1 tag
	DISKTAGKEY1 = "k8s.aliyun.com"
	// DISKTAGVALUE1 value
	DISKTAGVALUE1 = "true"
	// DISKTAGKEY2 key
	DISKTAGKEY2 = "createdby"
	// DISKTAGVALUE2 value
	DISKTAGVALUE2 = "alibabacloud-csi-plugin"
	// DISKTAGKEY3 key
	DISKTAGKEY3 = "ack.aliyun.com"
	// ECS snapshot tag from old version, keep it for compatibility
	SNAPSHOTTAGKEY1 = "force.delete.snapshot.k8s.aliyun.com"
)

// keys used in CreateSnapshotRequest.Parameters
const (
	SNAPSHOTTYPE               = "snapshotType"
	INSTANTACCESS              = "InstantAccess"
	RETENTIONDAYS              = "retentionDays"
	INSTANTACCESSRETENTIONDAYS = "instantAccessRetentionDays"
	SNAPSHOTRESOURCEGROUPID    = "resourceGroupId"
	SNAPSHOT_TAG_PREFIX        = "snapshotTags/"
)

const (
	// DiskSnapshotID means snapshot id
	DiskSnapshotID = "csi.alibabacloud.com/disk-snapshot-id"
	// IAVolumeSnapshotKey tag
	IAVolumeSnapshotKey = "csi.alibabacloud.com/snapshot-ia"
	// SnapshotRequestTag interval limit
	SnapshotRequestTag = "SNAPSHOT_REQUEST_INTERVAL"
	// DefaultVolumeSnapshotClass ...
	DefaultVolumeSnapshotClass = "alibabacloud-disk-snapshot"
	// annDiskID tag
	annDiskID = "volume.alibabacloud.com/disk-id"
	// MultiAttach tag
	MultiAttach = "multiAttach"
	// MinimumDiskSizeInGB ...
	MinimumDiskSizeInGB = 20
	// MinimumDiskSizeInBytes ...
	MinimumDiskSizeInBytes = 21474836480

	// LastApplyKey key
	LastApplyKey = "kubectl.kubernetes.io/last-applied-configuration"
	// StorageProvisionerKey key
	StorageProvisionerKey = "volume.beta.kubernetes.io/storage-provisioner"

	//snapshotTooMany means that the previous Snapshot is greater than 1
	snapshotTooMany string = "SnapshotTooMany"
	//snapshotAlreadyExist means that the snapshot already exists
	snapshotAlreadyExist string = "SnapshotAlreadyExist"
	//snapshotCreateError means that the create snapshot error occurred
	snapshotCreateError string = "SnapshotCreateError"
	//snapshotCreatedSuccessfully means that the create snapshot success
	snapshotCreatedSuccessfully string = "SnapshotCreatedSuccessfully"
	//snapshotDeleteError means that the delete snapshot error occurred
	snapshotDeleteError string = "SnapshotDeleteError"
	//snapshotDeletedSuccessfully means that the delete snapshot success
	snapshotDeletedSuccessfully string = "SnapshotDeletedSuccessfully"

	// DocumentURL document URL
	DocumentURL = "http://100.100.100.200/latest/dynamic/instance-identity/document"
	// IncorrectDiskStatus incorrect disk status
	IncorrectDiskStatus = "IncorrectDiskStatus"
	// NeverAttached status belongs to IncorrectDiskStatus
	NeverAttached = "IncorrectDiskStatus.NeverAttached"
	// DiskCreatingSnapshot ...
	DiskCreatingSnapshot = "DiskCreatingSnapshot"
	// UserNotInTheWhiteList tag
	UserNotInTheWhiteList = "UserNotInTheWhiteList"
	// ZoneIDTag tag
	ZoneIDTag = "zone-id"

	// These are error codes of ECS OpenAPI

	DiskNotAvailable             = "InvalidDataDiskCategory" // InvalidDataDiskCategory.ValueNotSupported/NotSupported
	DiskNotAvailableVer2         = "'DataDisk.n.Category' is not valid in this region."
	DiskSizeNotAvailable1        = "InvalidDiskSize.NotSupported"             // for cloud/cloud_efficiency/cloud_ssd
	DiskSizeNotAvailable2        = "InvalidDataDiskSize.ValueNotSupported"    // for cloud_auto
	DiskPerformanceLevelNotMatch = "OperationDenied.PerformanceLevelNotMatch" // for cloud_essd
	DiskInvalidPL                = "InvalidPerformanceLevel.Malformed"
	DiskIopsLimitExceeded        = "InvalidProvisionedIops.LimitExceed"
	DiskLimitExceeded            = "InstanceDiskLimitExceeded"
	NotSupportDiskCategory       = "NotSupportDiskCategory"
	DiskNotPortable              = "DiskNotPortable"
	IdempotentParameterMismatch  = "IdempotentParameterMismatch"
	SnapshotNotFound             = "InvalidSnapshotId.NotFound"
	InstanceNotFound             = "InvalidInstanceId.NotFound"

	// DiskHighAvail tag
	DiskHighAvail = "available"
	// DiskCommon common disk type
	DiskCommon = "cloud"
	// DiskEfficiency efficiency disk type
	DiskEfficiency = "cloud_efficiency"
	// DiskSSD ssd disk type
	DiskSSD = "cloud_ssd"
	// DiskESSD essd disk type
	DiskESSD = "cloud_essd"
	// DiskESSDAuto  essd autopl disk type
	DiskESSDAuto = "cloud_auto"
	// DiskESSDEntry  essd entry disk type
	DiskESSDEntry = "cloud_essd_entry"
	// DiskHighPerformance
	DiskPPerf = "cloud_pperf"
	// DiskStandPerformance
	DiskSPerf = "cloud_sperf"
	// DiskSharedSSD shared sdd disk type
	DiskSharedSSD = "san_ssd"
	// DiskSharedEfficiency shared efficiency disk type
	DiskSharedEfficiency = "san_efficiency"
	// MBSIZE tag
	MBSIZE = 1024 * 1024
	// GBSIZE tag
	GBSIZE = 1024 * MBSIZE
	// ZoneID ...
	ZoneID = "zoneId"
	// instanceTypeLabel ...
	instanceTypeLabel = "beta.kubernetes.io/instance-type"
	// zoneIDLabel ...
	zoneIDLabel = "failure-domain.beta.kubernetes.io/zone"
	// sigmaLabel instance type ...
	sigmaInstanceTypeLabel = "sigma.ali/machine-model"
	// sigmaLabel zoneid ....
	sigmaLabelZoneId = "sigma.ali/ecs-zone-id"
	// nodeStorageLabel ...
	nodeStorageLabel = "node.csi.alibabacloud.com/disktype.%s"

	nodeDiskCountAnnotation = "node.csi.alibabacloud.com/allocatable-disk"
	// kubeNodeName ...
	kubeNodeName = "KUBE_NODE_NAME"
	// describeResourceType ...
	describeResourceType = "DataDisk"
	// NodeSchedueTag in annotations
	NodeSchedueTag = "volume.kubernetes.io/selected-node"
	// RetryMaxTimes ...
	RetryMaxTimes = 5

	labelAppendPrefix = "csi.alibabacloud.com/label-prefix/"
	annVolumeTopoKey  = "csi.alibabacloud.com/volume-topology"
	labelVolumeType   = "csi.alibabacloud.com/disktype"
	annAppendPrefix   = "csi.alibabacloud.com/annotation-prefix/"

	VolumeDeleteAutoSnapshotKey                    = "csi.alibabacloud.com/volume-delete-autosnapshot-retentiondays"
	VOLUME_EXPAND_AUTO_SNAPSHOT_OP_KEY             = "volumeExpandAutoSnapshot"
	VOLUME_DELETE_AUTO_SNAPSHOT_OP_RETENT_DAYS_KEY = "volumeDeleteSnapshotRetentionDays"

	PROVISIONED_IOPS_KEY = "provisionedIops"
	BURSTING_ENABLED_KEY = "burstingEnabled"

	CSI_DEFAULT_FS_TYPE = "csi.storage.k8s.io/fstype"
	FS_TYPE             = "fsType"
	EXT4_FSTYPE         = "ext4"
	EXT3_FSTYPE         = "ext3"
	XFS_FSTYPE          = "xfs"
	NTFS_FSTYPE         = "ntfs"

	DISK_PERFORMANCE_LEVEL0 = "PL0"
	DISK_PERFORMANCE_LEVEL1 = "PL1"
	DISK_PERFORMANCE_LEVEL2 = "PL2"
	DISK_PERFORMANCE_LEVEL3 = "PL3"

	SNAPSHOT_MAX_RETENTION_DAYS = 65536
	SNAPSHOT_MIN_RETENTION_DAYS = 1

	DISK_CLOUD_EFFICIENT_MIN_CAPACITY    = 20
	DISK_CLOUD_SSD_MIN_CAPACITY          = 20
	DISK_CLOUD_ESSD_PL0_MIN_CAPACITY     = 1
	DISK_CLOUD_ESSD_PL1_MIN_CAPACITY     = 20
	DISK_CLOUD_ESSD_PL2_MIN_CAPACITY     = 461
	DISK_CLOUD_ESSD_PL3_MIN_CAPACITY     = 1261
	DISK_CLOUD_ESSD_AUTO_PL_MIN_CAPACITY = 1
)

var DiskCapacityMapping = map[string]int{
	DiskEfficiency: DISK_CLOUD_EFFICIENT_MIN_CAPACITY,
	DiskSSD:        DISK_CLOUD_SSD_MIN_CAPACITY,
	fmt.Sprintf("%s.%s", DiskESSD, DISK_PERFORMANCE_LEVEL0): DISK_CLOUD_ESSD_PL0_MIN_CAPACITY,
	fmt.Sprintf("%s.%s", DiskESSD, DISK_PERFORMANCE_LEVEL1): DISK_CLOUD_ESSD_PL1_MIN_CAPACITY,
	fmt.Sprintf("%s.%s", DiskESSD, DISK_PERFORMANCE_LEVEL2): DISK_CLOUD_ESSD_PL2_MIN_CAPACITY,
	fmt.Sprintf("%s.%s", DiskESSD, DISK_PERFORMANCE_LEVEL3): DISK_CLOUD_ESSD_PL3_MIN_CAPACITY,
	DiskESSDAuto: DISK_CLOUD_ESSD_AUTO_PL_MIN_CAPACITY,
	DiskESSD:     DISK_CLOUD_ESSD_PL1_MIN_CAPACITY,
}
