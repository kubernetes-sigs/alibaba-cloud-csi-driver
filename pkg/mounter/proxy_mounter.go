package mounter

import (
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

func NewProxyMounter(socketPath string, inner mountutils.Interface) Mounter {
	return &ProxyMounter{
		socketPath: socketPath,
		Interface:  inner,
	}
}

func (m *ProxyMounter) ExtendedMount(source, target, fstype string, options []string, params ExtendedMountParams) error {
	dclient := client.NewClient(m.socketPath)
	resp, err := dclient.Mount(&proxy.MountRequest{
		Source:      source,
		Target:      target,
		Fstype:      fstype,
		Options:     options,
		Secrets:     params.Secrets,
		MetricsPath: params.MetricsPath,
	})
	if err != nil {
		return fmt.Errorf("call mounter daemon: %w", err)
	}
	err = resp.ToError()
	if err != nil {
		return fmt.Errorf("failed to mount: %w", err)
	}
	notMnt, err := m.IsLikelyNotMountPoint(target)
	if err != nil {
		return err
	}
	if notMnt {
		return errors.New("failed to mount")
	}
	return nil
}

func (m *ProxyMounter) Mount(source string, target string, fstype string, options []string) error {
	return m.ExtendedMount(source, target, fstype, options, ExtendedMountParams{})
}
