package mounter

import (
	"context"

	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
	"k8s.io/mount-utils"
)

type ConnectorMounter struct {
	mounterPath string
	mount.Interface
}

var _ Mounter = &ConnectorMounter{}

func (m *ConnectorMounter) ExtendedMount(_ context.Context, req *mounterutils.MountRequest) error {
	if req == nil {
		return nil
	}
	args := mount.MakeMountArgs(req.Source, req.Target, req.Fstype, req.Options)
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
	return m.ExtendedMount(context.Background(), &mounterutils.MountRequest{
		Source:  source,
		Target:  target,
		Fstype:  fstype,
		Options: options,
	})
}

func NewConnectorMounter(inner mount.Interface, mounterPath string) Mounter {
	return &ConnectorMounter{
		mounterPath: mounterPath,
		Interface:   inner,
	}
}
