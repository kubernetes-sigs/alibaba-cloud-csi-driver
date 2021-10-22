package utils

import (
	"os"
)

// KubeletRootDir kubelet root dir;
var KubeletRootDir = "/var/lib/kubelet"

// MountPathWithTLS tls mount path;
var MountPathWithTLS = "/tls"

func init() {
	rootDir := os.Getenv("KUBELET_ROOT_DIR")
	if rootDir != "" {
		KubeletRootDir = rootDir
	}
	tlsDir := os.Getenv("TLS_MOUNT_DIR")
	if tlsDir != "" {
		MountPathWithTLS = tlsDir
	}
}
