package mounter

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	mountutils "k8s.io/mount-utils"
)

type Mounter interface {
	mountutils.Interface
	MountWithSecrets(source, target, fstype string, options []string, authCfg *utils.AuthConfig) error
}
