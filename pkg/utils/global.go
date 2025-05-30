package utils

import (
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
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

func GetNetworkType() string {
	return os.Getenv("ALIBABA_CLOUD_NETWORK_TYPE")
}
const DefaultRegistry = "registry-cn-hangzhou.ack.aliyuncs.com"

func GetRepositoryPrefix(m metadata.MetadataProvider) string {
	prefix, err := m.Get(metadata.RepositoryPrefix)
	if err == nil && prefix != "" {
		return prefix
	}
	url, err := m.Get(metadata.RegistryURL)
	if err == nil && url != "" {
		return path.Join(url, "acs/")
	}
	region, err := m.Get(metadata.RegionID)
	if err == nil && region != "" {
		url := fmt.Sprintf("registry-%s-vpc.ack.aliyuncs.com", region)
		klog.Warningf("DEFAULT_REGISTRY env not set, get current region: %v, fallback to default registry: %s", region, url)
		return path.Join(url, "acs/")
	}
	klog.Warningf("DEFAULT_REGISTRY env not set, failed to get current region: %v, fallback to default registry: %s", err, DefaultRegistry)
	return path.Join(DefaultRegistry, "acs/")
}
