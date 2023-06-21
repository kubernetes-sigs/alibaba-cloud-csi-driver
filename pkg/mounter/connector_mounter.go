package mounter

import (
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	mountutils "k8s.io/mount-utils"
)

type ConnectorMounter struct {
	mountutils.Interface
}

func (m *ConnectorMounter) Mount(source string, target string, fstype string, options []string) error {
	args := mountutils.MakeMountArgs(source, target, fstype, options)
	mntCmd := "systemd-run --scope -- mount " + strings.Join(args, " ")
	out, err := utils.ConnectorRun(mntCmd)
	if len(out) > 0 {
		log.Infof("ConnectorRun: %q, output: %s", mntCmd, string(out))
	}
	return err
}

func NewConnectorMounter(inner mountutils.Interface) mountutils.Interface {
	return &ConnectorMounter{
		Interface: inner,
	}
}
