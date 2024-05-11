package nas

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

type NasMounter struct {
	mountutils.Interface
	fuseMounter mountutils.Interface
}

func (m *NasMounter) Mount(source string, target string, fstype string, options []string) (err error) {
	logger := klog.Background().WithValues(
		"source", source,
		"target", target,
		"options", options,
		"fstype", fstype,
	)
	switch fstype {
	case "alinas", "cpfs", "cpfs-nfs":
		err = m.fuseMounter.Mount(source, target, fstype, options)
	default:
		err = m.Interface.Mount(source, target, fstype, options)
	}
	if err != nil {
		logger.Error(err, "failed to mount")
	} else {
		logger.Info("mounted successfully")
	}
	return err
}

func NewNasMounter() mountutils.Interface {
	inner := mountutils.New("")
	return &NasMounter{
		Interface:   inner,
		fuseMounter: mounter.NewConnectorMounter(inner, ""),
	}
}
