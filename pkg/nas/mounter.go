package nas

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	mountutils "k8s.io/mount-utils"
)

type NasMounter struct {
	mountutils.Interface
	fuseMounter mountutils.Interface
}

func (m *NasMounter) Mount(source string, target string, fstype string, options []string) error {
	switch fstype {
	case "alinas", "cpfs", "cpfs-nfs":
		return m.fuseMounter.Mount(source, target, fstype, options)
	default:
		return m.Interface.Mount(source, target, fstype, options)
	}
}

func NewNasMounter() mountutils.Interface {
	inner := mountutils.New("")
	return &NasMounter{
		Interface:   inner,
		fuseMounter: mounter.NewConnectorMounter(inner, ""),
	}
}
