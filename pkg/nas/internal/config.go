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
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
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

func mustGetKubeClients() (kubernetes.Interface, cnfsv1beta1.CNFSGetter) {
	cfg := options.MustGetRestConfig()
	crdCfg := options.GetRestConfigForCRD(*cfg)
	return kubernetes.NewForConfigOrDie(cfg), cnfsv1beta1.NewCNFSGetter(dynamic.NewForConfigOrDie(crdCfg))
}

func GetControllerConfig(meta *metadata.Metadata) (*ControllerConfig, error) {
	kubeClient, cnfsGetter := mustGetKubeClients()
	config := &ControllerConfig{
		Metadata:         meta,
		KubeClient:       kubeClient,
		CNFSGetter:       cnfsGetter,
		NasClientFactory: cloud.NewNasClientFactory(),
	}

	cm, err := kubeClient.CoreV1().ConfigMaps(configMapNamespace).Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return nil, err
		}
	} else {
		config.SkipSubpathCreation, _ = parseBool(cm.Data["nas-fake-provision"])
	}

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

func GetNodeConfig() (*NodeConfig, error) {
	kubeClient, cnfsGetter := mustGetKubeClients()
	config := &NodeConfig{
		// enable nfs port check by default
		EnablePortCheck: true,
		KubeClient:      kubeClient,
		CNFSGetter:      cnfsGetter,
	}

	// check if enable nfs port check
	if value := os.Getenv("NAS_PORT_CHECK"); value != "" {
		config.EnablePortCheck, _ = parseBool(value)
	}

	// get csi-plugin configmap
	cm, err := kubeClient.CoreV1().ConfigMaps(configMapNamespace).Get(context.Background(), configMapName, metav1.GetOptions{})
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return nil, err
		}
	} else {
		if value := cm.Data["nas-metric-enable"]; value != "" {
			config.EnableVolumeStats, _ = parseBool(value)
		}
		config.EnableEFCCache = cm.Data["cnfs-cache-properties"] != "" || cm.Data["nas-efc-cache"] != ""
	}
	if value := os.Getenv("NAS_METRIC_BY_PLUGIN"); value != "" {
		config.EnableVolumeStats, _ = parseBool(value)
	}
	if config.EnableVolumeStats {
		klog.Info("enabled nas volume stats")
	}

	// get node info
	nodeName := os.Getenv("KUBE_NODE_NAME")
	node, err := kubeClient.CoreV1().Nodes().Get(context.Background(), nodeName, metav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	config.NodeName = nodeName

	// check if losetup enabled
	if value := os.Getenv("NAS_LOSETUP_ENABLE"); value != "" {
		config.EnableLosetup, _ = parseBool(value)
	}
	if config.EnableLosetup {
		klog.Info("enabled nas losetup mode")
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
