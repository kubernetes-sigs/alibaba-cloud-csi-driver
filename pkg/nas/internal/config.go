package internal

import (
	"context"
	"errors"
	"os"

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

const defaultAlinasMountProxySocket = "/run/cnfs/alinas-mounter.sock"

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
	cm := utils.DefaultConfig()

	config := &ControllerConfig{
		Metadata:         meta,
		KubeClient:       kubeClient,
		CNFSGetter:       cnfsGetter,
		NasClientFactory: cloud.NewNasClientFactory(),

		SkipSubpathCreation:    cm.GetBool("nas-fake-provision", "NAS_FAKE_PROVISION", false),
		EnableSubpathFinalizer: cm.GetBool("nas-subpath-finalizer", "ENABLE_NAS_SUBPATH_FINALIZER", true),
		EnableRecycleBinCheck:  cm.GetBool("nas-recyclebin-check", "ENABLE_NAS_RECYCLEBIN_CHECK", false),
	}
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

	// clients for kubernetes
	KubeClient kubernetes.Interface
	CNFSGetter cnfsv1beta1.CNFSGetter
}

func GetNodeConfig() (*NodeConfig, error) {
	kubeClient, cnfsGetter := mustGetKubeClients()
	cm := utils.DefaultConfig()

	config := &NodeConfig{
		KubeClient: kubeClient,
		CNFSGetter: cnfsGetter,

		EnablePortCheck:   cm.GetBool("nas-port-check", "NAS_PORT_CHECK", true),
		EnableVolumeStats: cm.GetBool("nas-metric-enable", "NAS_METRIC_BY_PLUGIN", false),
		EnableEFCCache: cm.Get("cnfs-cache-properties", "", "") != "" ||
			cm.Get("nas-efc-cache", "", "") != "",
		EnableLosetup: cm.GetBool("nas-losetup-enable", "NAS_LOSETUP_ENABLE", false),
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
