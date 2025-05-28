package mounter

import (
	"context"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server/ossfs"
	mountutils "k8s.io/mount-utils"
)

const timeout = time.Second * 10

type CmdMounter struct {
	execPath string
	mountutils.Interface
}

func NewCmdMounter(execPath string, inner mountutils.Interface) Mounter {
	return &CmdMounter{
		execPath:  execPath,
		Interface: inner,
	}
}

func (m *CmdMounter) MountWithSecrets(source, target, fstype string, options []string, secrets map[string]string) error {
	driver := ossfs.NewDriver()
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()
	return driver.Mount(ctx, &proxy.MountRequest{
		Source:  source,
		Target:  target,
		Fstype:  fstype,
		Options: options,
		Secrets: secrets,
	})
}
