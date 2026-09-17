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

	// AgenticFS is distinguished by StorageType; its FileSystemType is still "standard".
	// This canonical value is shared by CSI volumeAs, CNFS spec.type and
	// status.fsAttributes.storageType. Changes must be synced with the CNFS controller.
	StorageTypeAgentic = "Agentic"
)

var KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Nas-%s", version.VERSION)
