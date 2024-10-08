package cloud

import (
	"fmt"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
)

const (
	FilesystemTypeStandard = "standard"
	FilesystemTypeCpfs     = "cpfs"
	FilesystemTypeExtreme  = "extreme"
	DefaultAccessGroup     = "DEFAULT_VPC_GROUP_NAME"
)

var KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Nas-%s", version.VERSION)
