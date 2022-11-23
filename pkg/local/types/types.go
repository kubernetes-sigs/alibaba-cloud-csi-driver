package types

import (
	"k8s.io/client-go/kubernetes"
)

// GlobalConfig var
type GlobalConfig struct {
	Region                  string
	NodeID                  string
	Scheduler               string
	PmemEnable              bool
	PmemType                string
	CapacityToNode          bool
	GrpcProvision           bool
	HostNameAsTopo          bool
	TopoKeyDefine           string
	LocalSparseFileDir      string
	LocalSparseFileTempSize string
	KubeClient         *kubernetes.Clientset
}

var (
	// GlobalConfigVar var
	GlobalConfigVar GlobalConfig
)

const (
	// LocalDriverName tag
	LocalDriverName = "localplugin.csi.alibabacloud.com"
	// NodeSchedueTag in annotations
	NodeSchedueTag = "volume.kubernetes.io/selected-node"

	// PmemProjectQuotaBasePath ...
	PmemProjectQuotaBasePath = "/mnt/quotapath.%s"
	// PmemDeviceFilesystem ...
	PmemDeviceFilesystem = "ext4"
	// PmemLVMType ...
	PmemLVMType = "lvm"
	// PmemDirectType ...
	PmemDirectType = "direct"
	// PmemQuotaPathType ...
	PmemQuotaPathType = "quotapath"
	// PmemKmemType ...
	PmemKmemType = "kmem"
)

const (
	// VolumeSpecLabel tag
	VolumeSpecLabel = "pv.csi.alibabacloud.com/volume.spec"
	// VolumeLifecycleLabel tag
	VolumeLifecycleLabel = "pv.csi.alibabacloud.com/volume.lifecycle"
	// PmemNodeLable ...
	PmemNodeLable = "pmem.csi.alibabacloud.com/type"
	// VolumeLifecycleCreating tag
	VolumeLifecycleCreating = "creating"
	// VolumeLifecycleCreated tag
	VolumeLifecycleCreated = "created"
	// VolumeLifecycleDeleting tag
	VolumeLifecycleDeleting = "deleting"
	// VolumeLifecycleDeleted tag
	VolumeLifecycleDeleted = "deleted"
)
