//go:build unix

package nas

import (
	"context"
	"errors"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

type NasMounter struct {
	mountutils.Interface
	alinasMounter mounter.Mounter
}

var _ mounter.Mounter = &NasMounter{}

func (m *NasMounter) ExtendedMount(ctx context.Context, op *mounter.MountOperation) (err error) {
	logger := klog.Background().WithValues(
		"source", op.Source,
		"target", op.Target,
		"options", op.Options,
		"fstype", op.FsType,
	)
	switch op.FsType {
	case "alinas", "cpfs", "cpfs-nfs":
		err = m.alinasMounter.ExtendedMount(ctx, op)
	default:
		err = m.Mount(op.Source, op.Target, op.FsType, op.Options)
	}
	if err != nil {
		logger.Error(err, "failed to mount")
	} else {
		logger.Info("mounted successfully")
	}
	return err
}

// Unmount routes the unmount through the mount broker when the broker owns the
// target (i.e. it was mounted through the broker), and unmounts locally
// otherwise. Ownership is decided by the broker per target, NOT by the kernel
// fstype: alinas AccessPoint mounts are performed as fstype "alinas" but show up
// as "nfs" in the kernel mount table, so an fstype-based decision would route
// their unmounts locally and the umount.nfs RPC to tcp 12049 would be dropped by
// the csi_mount_proxy nftables rule (cgroup != 0), blocking ~3s.
func (m *NasMounter) Unmount(target string) error {
	return m.unmount(target, func() error { return m.Interface.Unmount(target) })
}

// UnmountWithForce is called by mount-utils CleanupMountWithForce. NasMounter
// satisfies mountutils.MounterForceUnmounter through its embedded Interface, so
// this override is required to keep broker-owned unmounts on the broker path;
// otherwise the forced local umount would still be dropped by the nftables rule.
func (m *NasMounter) UnmountWithForce(target string, umountTimeout time.Duration) error {
	return m.unmount(target, func() error {
		if fu, ok := m.Interface.(mountutils.MounterForceUnmounter); ok {
			return fu.UnmountWithForce(target, umountTimeout)
		}
		return m.Interface.Unmount(target)
	})
}

// unmount tries the mount broker first (when the alinas mounter supports it) and
// falls back to the provided local unmount when the broker does not own the
// target (or predates the unmount RPC). It never inspects the kernel fstype.
func (m *NasMounter) unmount(target string, local func() error) error {
	pu, ok := m.alinasMounter.(mounter.ProxyUnmounter)
	if !ok {
		// Not using the proxy mounter (agent/connector mode): unmount locally.
		return local()
	}
	logger := klog.Background().WithValues("target", target)
	err := pu.ExtendedUnmount(context.Background(), target)
	switch {
	case err == nil:
		logger.Info("unmounted successfully via mount broker")
		return nil
	case errors.Is(err, mounter.ErrTargetNotManagedByBroker):
		// Not a broker-owned mount (plain NFS, or broker too old): unmount locally.
		return local()
	default:
		// Broker is unreachable or returned an unexpected error. Do not fail the
		// unmount: fall back to a local unmount so unmounting never becomes
		// strictly dependent on broker availability. For broker-owned mounts this
		// local umount may hit the ~3s utab-fallback path, but that is strictly
		// better than failing NodeUnpublishVolume.
		logger.Error(err, "unmount via mount broker failed, falling back to local unmount")
		return local()
	}
}

func newNasMounter(agentMode bool, socketPath string) mounter.Mounter {
	inner := mountutils.NewWithoutSystemd("")
	m := &NasMounter{
		Interface: inner,
	}
	switch {
	case socketPath != "":
		m.alinasMounter = mounter.NewProxyMounter(socketPath, inner)
	case !agentMode: // normal case, use connector mounter to ensure backward compatibility
		m.alinasMounter = mounter.NewConnectorMounter(inner, "")
	default:
		m.alinasMounter = mounter.NewForMounter(mounter.NewAdaptorMounter(inner), interceptors.AlinasSecretInterceptor)
	}
	return m
}
