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

	// AgenticFS is distinguished by StorageType; its FileSystemType is still "standard". The
	// literal is a cross-repo contract with the status.fsAttributes.storageType that
	// managed-filesystem writes into a CNFS, so any change must be synced with that API type.
	StorageTypeAgentic = "Agentic"
)

var KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Nas-%s", version.VERSION)
