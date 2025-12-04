package mounter

import (
	"context"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

type ConnectorMounter struct {
	mounterPath string
	mountutils.Interface
}

var _ Mounter = &ConnectorMounter{}

func (m *ConnectorMounter) ExtendedMount(_ context.Context, op *MountOperation) error {
	args := mountutils.MakeMountArgs(op.Source, op.Target, op.FsType, op.Options)
	mntCmd := []string{"systemd-run", "--scope", "--"}
	if m.mounterPath == "" {
		mntCmd = append(mntCmd, "mount")
	} else {
		mntCmd = append(mntCmd, m.mounterPath)
	}
	mntCmd = append(mntCmd, args...)
	out, err := utils.ConnectorRun(mntCmd...)
	if len(out) > 0 {
		klog.Infof("ConnectorRun: %q, output: %s", mntCmd, string(out))
	}
	return err
}

func (m *ConnectorMounter) Mount(source string, target string, fstype string, options []string) error {
	return m.ExtendedMount(context.Background(), &MountOperation{
		Source:  source,
		Target:  target,
		FsType:  fstype,
		Options: options,
	})
}

func NewConnectorMounter(inner mountutils.Interface, mounterPath string) Mounter {
	return &ConnectorMounter{
		mounterPath: mounterPath,
		Interface:   inner,
	}
}
