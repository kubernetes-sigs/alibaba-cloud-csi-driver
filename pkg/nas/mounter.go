package nas

import (
	"context"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

type NasMounter struct {
	mountutils.Interface
	alinasMounter mounter.Mounter
}

var _ mounter.Mounter = &NasMounter{}

func (m *NasMounter) ExtendedMount(ctx context.Context, op *mounter.MountOperation) (err error) {
	logger := klog.Background().WithValues(
		"source", op.Source,
		"target", op.Target,
		"options", op.Options,
		"fstype", op.FsType,
	)
	switch op.FsType {
	case "alinas", "cpfs", "cpfs-nfs":
		err = m.alinasMounter.ExtendedMount(ctx, op)
	default:
		err = m.Mount(op.Source, op.Target, op.FsType, op.Options)
	}
	if err != nil {
		logger.Error(err, "failed to mount")
	} else {
		logger.Info("mounted successfully")
	}
	return err
}

func newNasMounter(agentMode bool, socketPath string) mounter.Mounter {
	inner := mountutils.NewWithoutSystemd("")
	m := &NasMounter{
		Interface: inner,
	}
	switch {
	case socketPath != "":
		m.alinasMounter = mounter.NewProxyMounter(socketPath, inner)
	case !agentMode: // normal case, use connector mounter to ensure backward compatibility
		m.alinasMounter = mounter.NewConnectorMounter(inner, "")
	default:
		m.alinasMounter = mounter.NewForMounter(mounter.NewAdaptorMounter(inner), interceptors.AlinasSecretInterceptor)
	}
	return m
}
