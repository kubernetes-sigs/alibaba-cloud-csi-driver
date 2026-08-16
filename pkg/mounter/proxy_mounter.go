package mounter

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/client"
	mountutils "k8s.io/mount-utils"
)

type ProxyMounter struct {
	socketPath string
	mountutils.Interface
}

var (
	_ Mounter        = &ProxyMounter{}
	_ ProxyUnmounter = &ProxyMounter{}
)

func NewProxyMounter(socketPath string, inner mountutils.Interface) Mounter {
	return &ProxyMounter{
		socketPath: socketPath,
		Interface:  inner,
	}
}

func (m *ProxyMounter) ExtendedMount(ctx context.Context, op *MountOperation) error {
	if op == nil {
		return nil
	}
	dclient := client.NewClient(m.socketPath)
	resp, err := dclient.Mount(ctx, &proxy.MountRequest{
		Source:      op.Source,
		Target:      op.Target,
		Fstype:      op.FsType,
		Options:     op.Options,
		Secrets:     op.Secrets,
		MetricsPath: op.MetricsPath,
		VolumeID:    op.VolumeID,
		Overlay:     op.Overlay,
	})
	if err != nil {
		return fmt.Errorf("call mounter daemon: %w", err)
	}
	err = resp.ToError()
	if err != nil {
		return fmt.Errorf("failed to mount: %w", err)
	}
	notMnt, err := m.IsLikelyNotMountPoint(op.Target)
	if err != nil {
		return err
	}
	if notMnt {
		return errors.New("failed to mount")
	}
	return nil
}

func (m *ProxyMounter) Mount(source string, target string, fstype string, options []string) error {
	return m.ExtendedMount(context.Background(), &MountOperation{
		Source:  source,
		Target:  target,
		FsType:  fstype,
		Options: options,
	})
}

// ExtendedUnmount unmounts target through the mount broker (mount-proxy-server),
// so the umount runs in the daemon's cgroup 0. This is required for NAS
// AccessPoint mounts: the csi_mount_proxy nftables rule drops mount-broker
// traffic (tcp dport 12049) from any process in cgroup != 0, so a local
// umount.nfs issued from a container cgroup would have its RPC to the broker
// silently dropped (~3s timeout per attempt).
//
// The broker decides ownership by target (tracked at mount time), not by fstype.
// If the broker has no record of target (proxy.ErrTargetNotManaged) or predates
// the unmount RPC ("invalid method"), ExtendedUnmount returns
// ErrTargetNotManagedByBroker so the caller can fall back to a local unmount.
func (m *ProxyMounter) ExtendedUnmount(ctx context.Context, target string) error {
	dclient := client.NewClient(m.socketPath)
	resp, err := dclient.Unmount(ctx, &proxy.UnmountRequest{Target: target})
	if err != nil {
		return fmt.Errorf("call mounter daemon: %w", err)
	}
	if err := resp.ToError(); err != nil {
		// Old mount-proxy-server that predates the unmount RPC replies
		// "invalid method"; brokers that do not own the target reply
		// proxy.ErrTargetNotManaged. In both cases there is nothing for the
		// broker to unmount, so let the caller fall back to a local unmount.
		if strings.Contains(err.Error(), "invalid method") ||
			strings.Contains(err.Error(), proxy.ErrTargetNotManaged) {
			return ErrTargetNotManagedByBroker
		}
		return fmt.Errorf("failed to unmount via mount broker: %w", err)
	}
	return nil
}
