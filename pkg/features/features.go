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
	// ECS don't allow parallel attach to a node by default.
	// Enable this if you need faster attach, and only if your UID is whitelisted (by open a ticket),
	// or you have the supportConcurrencyAttach=true tag on your ECS instance.
	//
	// Only effective when DiskADController is also enabled.
	DiskParallelAttach featuregate.Feature = "DiskParallelAttach"

	// Detach multiple disks from the same node in parallel.
	// ECS does not allow parallel detach from a node by default.
	// Enable this if you need faster detach, and only if your UID is whitelisted (by open a ticket),
	// or you have the supportConcurrencyDetach=true tag on your ECS instance.
	//
	// Only effective when DiskADController is also enabled.
	DiskParallelDetach featuregate.Feature = "DiskParallelDetach"

	// Make volumeExpandAutoSnapshot parameter no-op.
	//
	// This feature is broken by new ECS intant available snapshot. And it is hard to fix it robustly.
	// It is very rare (if any) that a disk is corrupted by expand.
	DisableExpandAutoSnapshots featuregate.Feature = "DisableExpandAutoSnapshots"

	// Take auto snapshots before delete disks.
	// This has the same functionality as setting the `VOLUME_DEL_AUTO_SNAP` environment variable.
	//
	// Unlike ExpandAutoSnapshots, new ECS available snapshots will not block disk deletion.
	EnableDeleteAutoSnapshots featuregate.Feature = "EnableDeleteAutoSnapshots"

	// Update ossfs version to v1.91 or later.
	//
	// This configuration only takes effect for newly mounted OSS volumes.
	UpdatedOssfsVersion featuregate.Feature = "UpdatedOssfsVersion"

	RundCSIProtocol3 featuregate.Feature = "RundCSIProtocol3"

	// Enable volume group snapshots.
	// This feature allows users to use the volume group snapshot functionality,
	// enabling snapshots of related disks under a workload through ECS's snapshot-group capability.
	EnableVolumeGroupSnapshots featuregate.Feature = "EnableVolumeGroupSnapshots"

	// Use cnfs-alinas-daemon instead of csiplugin-connector for alinas and efc mounting.
	AlinasMountProxy featuregate.Feature = "AlinasMountProxy"

	// ConstrainFusePodDeleteRV controls whether to constrain ResourceVersion to "0"
	// for fuse pod delete operations. When not explicitly set via --feature-gates,
	// the behavior is determined by the Kubernetes server version.
	ConstrainFusePodDeleteRV featuregate.Feature = "ConstrainFusePodDeleteRV"

	// Populate NodeGetInfo MaxVolumesPerNode from the ECS HighDensityDisk mode
	// (云盘高密模式) limit instead of the default DiskQuantity.
	//
	// Relies on non-public ECS OpenAPI fields (AdditionalInfo.EnableHighDensityMode and
	// the DensityDiskQuantity instance-type attribute). When enabled, the node/labeler
	// makes one extra DescribeInstances call per high-density-capable node; nodes whose
	// instance type does not support high density incur no extra call. When disabled,
	// the behavior is byte-identical to before (no extra API call).
	DiskHighDensityMode featuregate.Feature = "DiskHighDensityMode"

	// Enable FUSE fd passing for libfuse3-based clients.
	// When enabled, supported drivers (e.g., ossfs2) will use fd-passing protocol
	// where the client performs the kernel mount and passes the FUSE fd to the server.
	// This is an internal mechanism and not directly exposed to end users.
	//
	// Deployment note: this feature gate must be configured on BOTH the csi-provisioner
	// (controller) and the csi-plugin (node) components. The controller uses it to gate
	// capability validation in checkOssOptions; the node uses it to actually drive the
	// fd-passing code path. Mismatched configuration leads to inconsistent behavior.
	//
	// Mount leakage: when the fuse pod (RunC) or csi-agent container (RunD) exits
	// abnormally (e.g. OOM killed, crash, SIGKILL), the FUSE mount at attachPath will
	// become stale ("Transport endpoint is not connected") and remain on the node
	// until NodeUnstageVolume is called to clean it up. During this window, workload
	// pods using this volume will experience IO failures.
	EnableFUSEFdPassing featuregate.Feature = "EnableFUSEFdPassing"

	// Enable FUSE recovery for ossfs2.
	// When enabled, ossfs2 mounts will automatically recover (flush FUSE connection
	// and restart ossfs2) on non-SIGTERM process exit. This feature implies
	// fd-passing for ossfs2.
	//
	// Data integrity note: during recovery, in-flight FUSE requests are interrupted
	// with -EINTR via resend+flush. This means:
	//   - Reads: safe (idempotent). Applications retry and get correct data.
	//   - Writes: buffered writes in the dead daemon's memory are LOST. Applications
	//     receive -EINTR and know the write failed, but data already buffered by
	//     previous successful write() calls (not yet flushed to OSS) is unrecoverable.
	//   Applications requiring write durability MUST call fsync() and confirm success
	//   before considering data committed. Only data acknowledged by fsync() survives
	//   daemon crashes. This is consistent with POSIX semantics for any volatile cache.
	//
	// Deployment note: this feature gate must be configured on BOTH the csi-provisioner
	// (controller) and the csi-plugin (node) components, same as EnableFUSEFdPassing.
	//
	// Node OS/kernel requirement: before enabling this feature, the operator must
	// confirm that all nodes intended to host ossfs2 recovery mounts satisfy the
	// kernel/OS prerequisites enforced by the recovery kernel check (see
	// pkg/utils/os/kernel.go, ErrPrefixRecoveryKernel). Nodes that fail the check
	// will reject recovery-enabled mounts at runtime.
	EnableOssfs2Recovery featuregate.Feature = "EnableOssfs2Recovery"
)

var (
	FunctionalMutableFeatureGate = featuregate.NewFeatureGate()
	defaultDiskFeatureGate       = map[featuregate.Feature]featuregate.FeatureSpec{
		DiskADController:           {Default: false, PreRelease: featuregate.Alpha},
		DiskParallelAttach:         {Default: false, PreRelease: featuregate.Alpha},
		DiskParallelDetach:         {Default: false, PreRelease: featuregate.Alpha},
		DisableExpandAutoSnapshots: {Default: true, PreRelease: featuregate.GA, LockToDefault: true},
		EnableVolumeGroupSnapshots: {Default: false, PreRelease: featuregate.Alpha},
		EnableDeleteAutoSnapshots:  {Default: false, PreRelease: featuregate.Alpha},
		DiskHighDensityMode:        {Default: false, PreRelease: featuregate.Alpha},
	}

	defaultOSSFeatureGate = map[featuregate.Feature]featuregate.FeatureSpec{
		UpdatedOssfsVersion:      {Default: true, PreRelease: featuregate.Beta},
		ConstrainFusePodDeleteRV: {Default: true, PreRelease: featuregate.Beta},
		EnableOssfs2Recovery:     {Default: false, PreRelease: featuregate.Alpha},
	}

	defaultNasFeatureGate = map[featuregate.Feature]featuregate.FeatureSpec{
		AlinasMountProxy: {Default: false, PreRelease: featuregate.Alpha},
	}

	otherFeatureGate = map[featuregate.Feature]featuregate.FeatureSpec{
		RundCSIProtocol3:    {Default: false, PreRelease: featuregate.Alpha},
		EnableFUSEFdPassing: {Default: false, PreRelease: featuregate.Alpha},
	}
)

func init() {
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultDiskFeatureGate))
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultOSSFeatureGate))
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultNasFeatureGate))
	runtime.Must(FunctionalMutableFeatureGate.Add(otherFeatureGate))
}
