//go:build !windows

package customfuse

import (
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	fpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager"
	customfusefpm "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/fuse_pod_manager/customfuse"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	k8sver "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

const (
	driverType = "customfuse"
	driverName = "customfuseplugin.csi.alibabacloud.com"
)

type Driver struct {
	endpoint string
	servers  common.Servers
}

func NewDriver(endpoint string, m metadata.MetadataProvider, serviceType utils.ServiceType, csiCfg utils.Config, k8sVersion *k8sver.Version) *Driver {
	klog.Infof("Driver: %v version: %v", driverName, version.VERSION)

	d := &Driver{endpoint: endpoint}

	nodeName := os.Getenv("KUBE_NODE_NAME")
	if serviceType&utils.Node > 0 {
		if nodeName == "" {
			klog.Fatal("env KUBE_NODE_NAME is empty")
		}
	} else {
		nodeName = "controller"
	}

	cfg, err := options.GetRestConfig()
	if err != nil {
		klog.ErrorS(err, "failed to get rest config")
	}

	var clientset kubernetes.Interface
	if cfg != nil {
		clientset = kubernetes.NewForConfigOrDie(cfg)
	}

	fuseManager := customfusefpm.NewCustomFuse(csiCfg, clientset)
	constrainRV := fpm.ShouldConstrainResourceVersion(k8sVersion)
	fusePodManager := fpm.NewFusePodManager(fuseManager, clientset, constrainRV)

	var servers common.Servers
	servers.IdentityServer = &identityServer{
		common.GenericIdentityServer{Name: driverName},
	}

	if serviceType&utils.Controller != 0 {
		servers.ControllerServer = &controllerServer{
			client:         clientset,
			metadata:       m,
			fusePodManager: fusePodManager,
		}
	}
	if serviceType&utils.Node != 0 {
		servers.NodeServer = &nodeServer{
			locks:      utils.NewVolumeLocks(),
			rawMounter: mountutils.NewWithoutSystemd(""),
			GenericNodeServer: common.GenericNodeServer{
				NodeID: nodeName,
			},
		}
	}
	d.servers = servers
	return d
}

func (d *Driver) Run() {
	klog.InfoS("Starting driver", "driver", driverType, "endpoint", d.endpoint)
	common.RunCSIServer(driverType, d.endpoint, d.servers)
}

type identityServer struct {
	common.GenericIdentityServer
}
