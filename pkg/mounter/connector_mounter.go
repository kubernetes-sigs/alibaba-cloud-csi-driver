package mounter

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
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
		log.Infof("ConnectorRun: %q, output: %s", mntCmd, string(out))
	}
	return err
}

func NewConnectorMounter(inner mountutils.Interface, mounterPath string) mountutils.Interface {
	return &ConnectorMounter{
		mounterPath: mounterPath,
		Interface:   inner,
	}
}
