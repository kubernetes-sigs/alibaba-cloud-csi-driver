package types

import (
	"k8s.io/client-go/kubernetes"
)

// GlobalConfig var
type GlobalConfig struct {
	Region              string
	NodeID              string
	Scheduler           string
	PmemEnable          bool
	PmemType            string
	ControllerProvision bool
	KubeClient          *kubernetes.Clientset
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
)

const (
	// VolumeSpecLabel tag
	VolumeSpecLabel = "pv.csi.alibabacloud.com/volume.spec"
	// VolumeLifecycleLabel tag
	VolumeLifecycleLabel = "pv.csi.alibabacloud.com/volume.lifecycle"
	// VolumeLifecycleCreating tag
	VolumeLifecycleCreating = "creating"
	// VolumeLifecycleCreated tag
	VolumeLifecycleCreated = "created"
	// VolumeLifecycleDeleting tag
	VolumeLifecycleDeleting = "deleting"
	// VolumeLifecycleDeleted tag
	VolumeLifecycleDeleted = "deleted"
)
