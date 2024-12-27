package features

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"
)

const (
	// Do disk attach/detach at controller, not node.
	// Historically, disks don't have serial number, so we need to attach/detach disks at node
	// to identify which device is the just attached disk.
	// If all your disks are created after 2020-06-10, you can enable this safely.
	// See: https://www.alibabacloud.com/help/en/ecs/user-guide/query-the-serial-number-of-a-disk
	//
	// Enable this at controller first, wait for the rollout to finish, then enable this at node.
	DiskADController featuregate.Feature = "DiskADController"

	// Attach multiple disks to the same node in parallel.
	// ECS don't allow parallel attach/detach to a node by default.
	// Enable this if you need faster attach, and only if your UID is whitelisted (by open a ticket),
	// or you have the supportConcurrencyAttach=true tag on your ECS instance.
	//
	// Only effective when DiskADController is also enabled.
	DiskParallelAttach featuregate.Feature = "DiskParallelAttach"

	// Detach multiple disks from the same node in parallel.
	// ECS does not allow parallel detach from a node currently. This feature gate is reserved for future use.
	//
	// Only effective when DiskADController is also enabled.
	DiskParallelDetach featuregate.Feature = "DiskParallelDetach"

	// Make volumeExpandAutoSnapshot parameter no-op.
	//
	// This feature is broken by new ECS intant available snapshot. And it is hard to fix it robustly.
	// It is very rare (if any) that a disk is corrupted by expand.
	DisableExpandAutoSnapshots featuregate.Feature = "DisableExpandAutoSnapshots"

	UpdatedOssfsVersion featuregate.Feature = "UpdatedOssfsVersion"

	RundCSIProtocol3 featuregate.Feature = "RundCSIProtocol3"

	// Expose metrics from kubelet stat/summary API
	MetricKubeletStatSummary featuregate.Feature = "MetricKubeletStatSummary"

	EnableVolumeGroupSnapshots featuregate.Feature = "EnableVolumeGroupSnapshots"
)

var (
	FunctionalMutableFeatureGate = featuregate.NewFeatureGate()
	defaultDiskFeatureGate       = map[featuregate.Feature]featuregate.FeatureSpec{
		DiskADController:           {Default: false, PreRelease: featuregate.Alpha},
		DiskParallelAttach:         {Default: false, PreRelease: featuregate.Alpha},
		DiskParallelDetach:         {Default: false, PreRelease: featuregate.Alpha},
		DisableExpandAutoSnapshots: {Default: true, PreRelease: featuregate.GA, LockToDefault: true},
		EnableVolumeGroupSnapshots: {Default: false, PreRelease: featuregate.Alpha},
	}

	defaultOSSFeatureGate = map[featuregate.Feature]featuregate.FeatureSpec{
		UpdatedOssfsVersion: {Default: true, PreRelease: featuregate.Beta},
	}
	defaultNASFeatureGate = map[featuregate.Feature]featuregate.FeatureSpec{
		RundCSIProtocol3: {Default: false, PreRelease: featuregate.Alpha},
	}

	defaultMetricFeatureGate = map[featuregate.Feature]featuregate.FeatureSpec{
		MetricKubeletStatSummary: {Default: true, PreRelease: featuregate.Beta},
	}
)

func init() {
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultDiskFeatureGate))
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultOSSFeatureGate))
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultNASFeatureGate))
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultMetricFeatureGate))
}
