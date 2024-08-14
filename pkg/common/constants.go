package common

const (
	ECSInstanceIDTopologyKey = "alibabacloud.com/ecs-instance-id"
)

// constants of keys in volume parameters
const (
	// PVCNameKey contains name of the PVC for which is a volume provisioned.
	PVCNameKey = "csi.storage.k8s.io/pvc/name"

	// PVCNamespaceKey contains namespace of the PVC for which is a volume provisioned.
	PVCNamespaceKey = "csi.storage.k8s.io/pvc/namespace"

	// PVNameKey contains name of the final PV that will be used for the dynamically
	// provisioned volume
	PVNameKey = "csi.storage.k8s.io/pv/name"
)

// constants of keys in volume snapshot parameters
const (
	VolumeSnapshotNamespaceKey = "csi.storage.k8s.io/volumesnapshot/namespace"
	VolumeSnapshotNameKey      = "csi.storage.k8s.io/volumesnapshot/name"
)

const (
	// PVCNameTag is tag applied to provisioned alibaba cloud disk for compatibility
	// with in-tree volume plugin. Value of the tag is PVC name. It is applied only when
	// the external provisioner sidecar is started with --extra-create-metadata=true and
	// thus provides such metadata to the CSI driver.
	PVCNameTag = "kubernetes.io/created-for/pvc/name"

	// PVCNamespaceTag is tag applied to provisioned alibaba cloud disk for compatibility
	// with in-tree volume plugin. Value of the tag is PVC namespace. It is applied only when
	// the external provisioner sidecar is started with --extra-create-metadata=true and
	// thus provides such metadata to the CSI driver.
	PVCNamespaceTag = "kubernetes.io/created-for/pvc/namespace"

	// PVNameTag is tag applied to provisioned alibaba cloud disk for compatibility
	// with in-tree volume plugin. Value of the tag is PV name. It is applied only when
	// the external provisioner sidecar is started with --extra-create-metadata=true and
	// thus provides such metadata to the CSI driver.
	PVNameTag = "kubernetes.io/created-for/pv/name"

	// VolumeNameTag is tag applied to provisioned alibaba cloud disk
	// Disk name have many restrictions, so we use this tag to store the original name
	VolumeNameTag = "csi.alibabacloud.com/volume-name"
)

// Tags that will be added to ECS snapshots
const (
	VolumeSnapshotNameTag      = "csi.alibabacloud.com/snapshot/name"
	VolumeSnapshotNamespaceTag = "csi.alibabacloud.com/snapshot/namespace"
)
