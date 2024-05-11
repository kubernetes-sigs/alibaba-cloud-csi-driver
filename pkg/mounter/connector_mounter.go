package mounter

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

type ConnectorMounter struct {
	mounterPath string
	mountutils.Interface
}

func (m *ConnectorMounter) Mount(source string, target string, fstype string, options []string) error {
	args := mountutils.MakeMountArgs(source, target, fstype, options)
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

func NewConnectorMounter(inner mountutils.Interface, mounterPath string) mountutils.Interface {
	return &ConnectorMounter{
		mounterPath: mounterPath,
		Interface:   inner,
	}
}
