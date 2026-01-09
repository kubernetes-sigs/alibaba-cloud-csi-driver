package nas

import (
	"context"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/interceptors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

type NasMounter struct {
	mountutils.Interface
	alinasMounter mounter.Mounter
}

var _ mounter.Mounter = &NasMounter{}

func (m *NasMounter) ExtendedMount(ctx context.Context, req *utils.MountRequest) (err error) {
	logger := klog.Background().WithValues(
		"source", req.Source,
		"target", req.Target,
		"options", req.Options,
		"fstype", req.Fstype,
	)
	switch req.Fstype {
	case "alinas", "cpfs", "cpfs-nfs":
		err = m.alinasMounter.ExtendedMount(ctx, req)
	default:
		err = m.Mount(req.Source, req.Target, req.Fstype, req.Options)
	}
	if err != nil {
		logger.Error(err, "failed to mount")
	} else {
		logger.Info("mounted successfully")
	}
	return err
}

func newNasMounter(agentMode bool, socketPath, regionID string) mounter.Mounter {
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
		interceptor, err := interceptors.NewAlinasSecretInterceptor(regionID)
		if err != nil {
			klog.ErrorS(err, "failed to create alinas secret interceptor")
			m.alinasMounter = mounter.NewForMounter(mounter.NewAdaptorMounter(inner))
		} else {
			m.alinasMounter = mounter.NewForMounter(mounter.NewAdaptorMounter(inner), interceptor.Intercept)
		}
	}
	return m
}
