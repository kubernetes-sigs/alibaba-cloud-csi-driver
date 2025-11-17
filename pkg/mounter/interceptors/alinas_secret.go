package interceptors

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
)

type AlinasSecretInterceptor struct{}

var _ mounter.MountInterceptor = AlinasSecretInterceptor{}

func NewAlinasSecretInterceptor() mounter.MountInterceptor {
	return AlinasSecretInterceptor{}
}

func (AlinasSecretInterceptor) BeforeMount(req *mounter.MountOperation) (*mounter.MountOperation, error) {
	return req, nil
}
