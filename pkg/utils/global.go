package utils

import (
	"flag"
	"os"

	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
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

func AddKlogFlags(fs *pflag.FlagSet) {
	local := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	klog.InitFlags(local)
	fs.AddGoFlagSet(local)
}

func AddGoFlags(fs *pflag.FlagSet) {
	fs.AddGoFlagSet(flag.CommandLine)
}
