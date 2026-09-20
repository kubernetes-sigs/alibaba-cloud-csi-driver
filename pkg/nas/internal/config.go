package internal

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/interfaces"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const (
	configMapName      = "csi-plugin"
	configMapNamespace = "kube-system"

	// defaultRRSACAFile is where the chart projects the oidc-proxy CA secret.
	defaultRRSACAFile = "/etc/csi-plugin/oidc-proxy-ca/ca.crt"
)

type ControllerConfig struct {
	// cluster info
	Metadata metadata.MetadataProvider

	// subpath configs
	SkipSubpathCreation    bool
	EnableSubpathFinalizer bool
	// check whether recycle bin enabled before subpath deletion
	EnableRecycleBinCheck bool

	// clients for kubernetes
	KubeClient kubernetes.Interface
	CNFSGetter cnfsv1beta1.CNFSGetter

	// clients for alibaba cloud
	NasClientFactory interfaces.NasClientFactoryInterface
}

func getKubeClients() (kubernetes.Interface, cnfsv1beta1.CNFSGetter) {
	cfg, err := options.GetRestConfig()
	if err != nil {
		klog.ErrorS(err, "failed to get rest config")
		return nil, nil
	}
	crdCfg := options.GetRestConfigForCRD(*cfg)
	return kubernetes.NewForConfigOrDie(cfg), cnfsv1beta1.NewCNFSGetter(dynamic.NewForConfigOrDie(crdCfg))
}

func GetControllerConfig(meta *metadata.Metadata, csiCfg utils.Config) (*ControllerConfig, error) {
	kubeClient, cnfsGetter := getKubeClients()
	config := &ControllerConfig{
		Metadata:         meta,
		KubeClient:       kubeClient,
		CNFSGetter:       cnfsGetter,
		NasClientFactory: cloud.NewNasClientFactory(),
	}

	config.SkipSubpathCreation = csiCfg.GetBool("nas-fake-provision", "NAS_FAKE_PROVISION", false)
	config.EnableSubpathFinalizer, _ = parseBool(os.Getenv("ENABLE_NAS_SUBPATH_FINALIZER"))
	config.EnableRecycleBinCheck, _ = parseBool(os.Getenv("ENABLE_NAS_RECYCLEBIN_CHECK"))

	return config, nil
}

type NodeConfig struct {
	NodeName          string
	NodeIP            string
	EnableEFCCache    bool
	EnableMixRuntime  bool
	EnablePortCheck   bool
	EnableLosetup     bool
	EnableVolumeStats bool

	// MountProxySocket is the resolved socket path for mount-proxy-server.
	// Resolution is done in main.go:
	//   1. --mount-proxy-sock flag (sandbox agent scenario)
	//   2. AlinasMountProxy feature gate + default socket
	//   3. Empty string → NAS uses ConnectorMounter instead of ProxyMounter
	MountProxySocket string
	AgentMode        bool

	// RRSACAFile is the trust anchor for the oidc-proxy serving certificate used by
	// authType=rrsa volumes. Empty means the system pool, for a proxy behind a
	// publicly trusted certificate.
	RRSACAFile string

	// Region is used to build the default STS endpoint for authType=rrsa volumes
	// that name no rrsaEndpoint.
	Region string

	// AccountID and ClusterID complete an authType=rrsa volume that names only a
	// roleName, the way the OSS driver does.
	AccountID string
	ClusterID string

	// RRSADuration is the credential lifetime to ask for. Short means a small
	// window in which a credential outlives its Pod; long means fewer
	// AssumeRoleWithOIDC calls, which matters once a cluster has many volumes.
	// Zero leaves the choice to the endpoint, and the role's MaxSessionDuration
	// caps it either way.
	RRSADuration time.Duration

	// clients for kubernetes
	KubeClient kubernetes.Interface
	CNFSGetter cnfsv1beta1.CNFSGetter
}

func GetNodeConfig(meta metadata.MetadataProvider, csiCfg utils.Config, mountProxySock string) (*NodeConfig, error) {
	kubeClient, cnfsGetter := getKubeClients()
	config := &NodeConfig{
		// enable nfs port check by default
		EnablePortCheck:   true,
		KubeClient:        kubeClient,
		CNFSGetter:        cnfsGetter,
		EnableVolumeStats: csiCfg.GetBool("nas-metric-enable", "NAS_METRIC_BY_PLUGIN", false),
		EnableEFCCache: csiCfg.Get("cnfs-cache-properties", "CNFS_CACHE_PROPERTIES", "") != "" ||
			csiCfg.Get("nas-efc-cache", "NAS_EFC_CACHE", "") != "",
	}

	// Only used to build a default STS endpoint, so a provider that cannot supply
	// the region is not fatal: GetSTSEndpoint falls back to the public endpoint.
	if region, err := meta.Get(metadata.RegionID); err != nil {
		klog.V(4).InfoS("region unavailable, rrsa volumes without rrsaEndpoint will use the public STS endpoint", "error", err)
	} else {
		config.Region = region
	}
	// Only needed to complete a bare roleName, so a volume naming both ARNs works
	// without them.
	if accountID, err := meta.Get(metadata.AccountID); err != nil {
		klog.V(4).InfoS("account ID unavailable, rrsa volumes must name both ARNs", "error", err)
	} else {
		config.AccountID = accountID
	}
	if clusterID, err := meta.Get(metadata.ClusterID); err != nil {
		klog.V(4).InfoS("cluster ID unavailable, rrsa volumes must name both ARNs", "error", err)
	} else {
		config.ClusterID = clusterID
	}

	// check if enable nfs port check
	if value := os.Getenv("NAS_PORT_CHECK"); value != "" {
		config.EnablePortCheck, _ = parseBool(value)
	}

	// Mounted optionally, so the default path may not exist; mounts that need it
	// report that themselves.
	config.RRSACAFile = csiCfg.Get("nas-rrsa-ca-file", "NAS_RRSA_CA_FILE", defaultRRSACAFile)

	if v := csiCfg.Get("nas-rrsa-duration", "NAS_RRSA_DURATION", ""); v != "" {
		d, err := time.ParseDuration(v)
		if err != nil || d <= 0 {
			// Refuse to start rather than fall back to the endpoint's default: this
			// value is how long a credential may outlive its Pod, and silently
			// granting a different lifetime than asked for is not something an
			// operator would notice.
			return nil, fmt.Errorf("invalid nas-rrsa-duration %q: want a positive duration such as 15m", v)
		}
		config.RRSADuration = d
	}

	if config.EnableVolumeStats {
		klog.Info("enabled nas volume stats")
	}

	nodeName := os.Getenv("KUBE_NODE_NAME")
	config.NodeName = nodeName

	// check if losetup enabled
	config.EnableLosetup, _ = parseBool(os.Getenv("NAS_LOSETUP_ENABLE"))
	if config.EnableLosetup {
		klog.Info("enabled nas losetup mode")
		node, err := kubeClient.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
		if err != nil {
			return nil, err
		}

		for _, addr := range node.Status.Addresses {
			if addr.Type == corev1.NodeInternalIP {
				config.NodeIP = addr.Address
				break
			}
		}
		if config.NodeIP == "" {
			return nil, errors.New("enabled losetup mode but failed to get node IP")
		}
	}

	// mountProxySock is already resolved by main.go (flag > feature gate + default > empty).
	config.MountProxySocket = mountProxySock

	return config, nil
}

func parseBool(str string) (bool, error) {
	switch str {
	case "enable", "enabled", "yes":
		return true, nil
	case "no", "":
		return false, nil
	}
	return strconv.ParseBool(str)
}
