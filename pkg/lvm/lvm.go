package lvm

import (
	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"os"
	"path"
)

type lvm struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         *identityServer
	nodeServer       csi.NodeServer
	controllerServer *controllerServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

const (
	PluginFolder = "/var/lib/kubelet/plugins/lvmplugin.csi.alibabacloud.com"
	driverName   = "lvmplugin.csi.alibabacloud.com"
	csiVersion   = "1.0.0"
)

// Init checks for the persistent volume file and loads all found volumes
// into a memory structure
func initDriver() {
	if _, err := os.Stat(path.Join(PluginFolder, "controller")); os.IsNotExist(err) {
		log.Infof("lvm: folder %s not found. Creating... ", path.Join(PluginFolder, "controller"))
		if err := os.Mkdir(path.Join(PluginFolder, "controller"), 0755); err != nil {
			log.Fatalf("Failed to create a controller's volumes folder with error: %v", err)
		}
		return
	}
}

func NewDriver(nodeID, endpoint string) *lvm {
	initDriver()
	tmplvm := &lvm{}
	tmplvm.endpoint = endpoint

	if nodeID == "" {
		nodeID = GetMetaData(INSTANCE_ID)
		log.Infof("Use node id : %s", nodeID)
	}
	csiDriver := csicommon.NewCSIDriver(driverName, csiVersion, nodeID)
	tmplvm.driver = csiDriver
	tmplvm.driver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
	})
	tmplvm.driver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})

	// Create GRPC servers
	tmplvm.idServer = NewIdentityServer(tmplvm.driver)
	tmplvm.nodeServer = NewNodeServer(tmplvm.driver, nodeID)
	tmplvm.controllerServer = NewControllerServer(tmplvm.driver)

	return tmplvm
}

func (lvm *lvm) Run() {
	log.Infof("Driver: %v ", driverName)

	server := csicommon.NewNonBlockingGRPCServer()
	server.Start(lvm.endpoint, lvm.idServer, lvm.controllerServer, lvm.nodeServer)
	server.Wait()
}
