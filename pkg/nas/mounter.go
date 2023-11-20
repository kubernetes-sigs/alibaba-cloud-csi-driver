package nas

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/sirupsen/logrus"
	mountutils "k8s.io/mount-utils"
)

type NasMounter struct {
	mountutils.Interface
	fuseMounter mountutils.Interface
}

func (m *NasMounter) Mount(source string, target string, fstype string, options []string) (err error) {
	log := logrus.WithFields(logrus.Fields{
		"source":       source,
		"target":       target,
		"mountOptions": options,
		"fstype":       fstype,
	})
	switch fstype {
	case "alinas", "cpfs", "cpfs-nfs":
		err = m.fuseMounter.Mount(source, target, fstype, options)
	default:
		err = m.Interface.Mount(source, target, fstype, options)
	}
	if err != nil {
		log.Errorf("failed to mount: %v", err)
	} else {
		log.Info("mounted successfully")
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
