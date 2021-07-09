package utils

import (
	"os"
)

// KubeletRootDir kubelet root dir;
var KubeletRootDir = "/var/lib/kubelet"

func init() {
	rootDir := os.Getenv("KUBELET_ROOT_DIR")
	if rootDir != "" {
		KubeletRootDir = rootDir
	}
}
