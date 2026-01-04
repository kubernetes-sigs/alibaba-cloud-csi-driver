package mounter

import (
	"context"
	"errors"
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/client"
	mountutils "k8s.io/mount-utils"
)

type ProxyMounter struct {
	socketPath string
	mountutils.Interface
}

var _ Mounter = &ProxyMounter{}

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
	resp, err := dclient.Mount(&proxy.MountRequest{
		Source:      op.Source,
		Target:      op.Target,
		Fstype:      op.FsType,
		Options:     op.Options,
		Secrets:     op.Secrets,
		MetricsPath: op.MetricsPath,
		VolumeID:    op.VolumeID,
		AuthConfig:  op.AuthConfig,
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
