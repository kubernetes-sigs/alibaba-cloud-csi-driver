package utils

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

// KubeletRootDir kubelet root dir;
var KubeletRootDir = "/var/lib/kubelet"

// SubstrateRootDir is the host directory under which atelet creates
// per-actor volume mount points and staging directories. The bridge driver
// derives its actors directory from this (root + "/actors"), and the node
// server validator checks stage/publish paths against it.
var SubstrateRootDir = "/var/lib/ate"

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

func GetNetworkType() string {
	return os.Getenv("ALIBABA_CLOUD_NETWORK_TYPE")
}

const DefImageRegistry = "registry-cn-hangzhou.ack.aliyuncs.com"
const DefImageNamespace = "acs"

func GetRepositoryPrefix(region string) string {
	prefix := os.Getenv("IMAGE_REPOSITORY_PREFIX")
	if prefix != "" {
		return prefix
	}
	url := os.Getenv("DEFAULT_REGISTRY")
	if url != "" {
		return path.Join(url, DefImageNamespace)
	}
	if region != "" {
		url := fmt.Sprintf("registry-%s-vpc.ack.aliyuncs.com", region)
		klog.Warningf("DEFAULT_REGISTRY env not set, get current region: %v, fallback to default registry: %s", region, url)
		return path.Join(url, DefImageNamespace)
	}
	klog.Warningf("DEFAULT_REGISTRY env not set, failed to get current region, fallback to default registry: %s", DefImageRegistry)
	return path.Join(DefImageRegistry, DefImageNamespace)
}
