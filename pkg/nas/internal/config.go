package internal

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	cnfsv1beta1 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
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

	defaultAlinasMountProxySocket = "/run/cnfs/alinas-mounter.sock"
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

	// path of mount proxy socket
	MountProxySocket string
	AgentMode        bool

	// clients for kubernetes
	KubeClient kubernetes.Interface
	CNFSGetter cnfsv1beta1.CNFSGetter
}

func GetNodeConfig(csiCfg utils.Config) (*NodeConfig, error) {
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

	// check if enable nfs port check
	if value := os.Getenv("NAS_PORT_CHECK"); value != "" {
		config.EnablePortCheck, _ = parseBool(value)
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

	if features.FunctionalMutableFeatureGate.Enabled(features.AlinasMountProxy) {
		config.MountProxySocket = defaultAlinasMountProxySocket
	}

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
