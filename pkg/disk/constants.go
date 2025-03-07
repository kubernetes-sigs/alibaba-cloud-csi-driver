package disk

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
	//snapshotGetError means that the get snapshot error occurred
	snapshotGetError string = "SnapshotGetError"

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

	QuotaExceed_Snapshot = "QuotaExceed.Snapshot"

	// DiskHighAvail tag
	DiskHighAvail = "available"
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

	nodeDiskTypeLabelPrefix = "node.csi.alibabacloud.com/disktype."

	nodeDiskCountAnnotation = "node.csi.alibabacloud.com/allocatable-disk"
	// kubeNodeName ...
	kubeNodeName = "KUBE_NODE_NAME"
	// describeResourceType ...
	describeResourceType = "DataDisk"
	// NodeScheduleTag in annotations
	NodeScheduleTag = "volume.kubernetes.io/selected-node"
	// RetryMaxTimes ...
	RetryMaxTimes = 5

	labelAppendPrefix = "csi.alibabacloud.com/label-prefix/"
	annVolumeTopoKey  = "csi.alibabacloud.com/volume-topology"
	labelVolumeType   = "csi.alibabacloud.com/disktype"
	annAppendPrefix   = "csi.alibabacloud.com/annotation-prefix/"

	VolumeDeleteAutoSnapshotKey                    = "csi.alibabacloud.com/volume-delete-autosnapshot-retentiondays"
	VOLUME_DELETE_AUTO_SNAPSHOT_OP_RETENT_DAYS_KEY = "volumeDeleteSnapshotRetentionDays"

	PROVISIONED_IOPS_KEY = "provisionedIops"
	BURSTING_ENABLED_KEY = "burstingEnabled"

	CSI_DEFAULT_FS_TYPE = "csi.storage.k8s.io/fstype"
	FS_TYPE             = "fsType"
	EXT4_FSTYPE         = "ext4"
	EXT3_FSTYPE         = "ext3"
	XFS_FSTYPE          = "xfs"
	NTFS_FSTYPE         = "ntfs"

	SNAPSHOT_MAX_RETENTION_DAYS = 65536
	SNAPSHOT_MIN_RETENTION_DAYS = 1

	PUBLISH_CONTEXT_SERIAL = "serialNumber"
)
