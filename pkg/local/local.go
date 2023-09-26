/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package local

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// Local the Local struct
type Local struct {
	driver           *csicommon.CSIDriver
	endpoint         string
	idServer         *identityServer
	nodeServer       csi.NodeServer
	controllerServer *controllerServer

	cap   []*csi.VolumeCapability_AccessMode
	cscap []*csi.ControllerServiceCapability
}

const (
	defaultDriverName = "localplugin.csi.alibabacloud.com"
	localDriverName   = "localplugin.csi.alibabacloud.com"
	yodaDriverName    = "yodaplugin.csi.alibabacloud.com"
	csiVersion        = "1.0.0"
)

const (
	localSecretName      = "csi-local-plugin-cert"
	localSecretNamespace = "kube-system"
	caCertFileName       = "ca_cert.pem"
	serverCertFileName   = "server_cert.pem"
	serverKeyFileName    = "server_key.pem"
	clientCertFileName   = "client_cert.pem"
	clientKeyFileName    = "client_key.pem"
)

// PmemSupportType ...
var PmemSupportType = []string{types.PmemLVMType, types.PmemDirectType, types.PmemKmemType, types.PmemQuotaPathType}

// Init checks for the persistent volume file and loads all found volumes
// into a memory structure
func initDriver() {
}

func writeClientCert(caCert []byte, clientCert []byte, clientKey []byte) (string, string, string) {
	caCertPath := filepath.Join(utils.MountPathWithTLS, "/local/grpc")
	utils.CreateDest(caCertPath)
	caCertFile := caCertPath + "/" + caCertFileName
	clientCertFile := caCertPath + "/" + clientCertFileName
	clientKeyFile := caCertPath + "/" + clientKeyFileName
	err := utils.WriteAndSyncFile(caCertFile, caCert, 0644)
	if err != nil {
		log.Fatalf("WriteFile caCertFile is failed, caCertFile:%s, err:%s", caCertFile, err)
	}
	err = utils.WriteAndSyncFile(clientCertFile, clientCert, 0644)
	if err != nil {
		log.Fatalf("WriteFile clientCertFile is failed, clientCertFile:%s, err:%s", clientCertFile, err)
	}
	err = utils.WriteAndSyncFile(clientKeyFile, clientKey, 0644)
	if err != nil {
		log.Fatalf("WriteFile clientKeyFile is failed, clientKeyFile:%s, err:%s", clientKeyFile, err)
	}
	return caCertFile, clientCertFile, clientKeyFile
}

func createSecretWithTLS() (string, string, string) {
	secret, err := getSecret(localSecretNamespace, localSecretName)
	if err == nil {
		caCert, ok := secret.Data[caCertFileName]
		if !ok {
			log.Fatalf("caCert is not exist.")
		}
		clientCert, ok := secret.Data[clientCertFileName]
		if !ok {
			log.Fatalf("clientCert is not exist.")
		}
		clientKey, ok := secret.Data[clientKeyFileName]
		if !ok {
			log.Fatalf("clientKey is not exist.")
		}
		return writeClientCert(caCert, clientCert, clientKey)
	}

	begin := time.Now().Add(-time.Hour)
	end := time.Now().Add(time.Hour * 24 * 365 * 100)
	ca, err := utils.CreateCACert(utils.CertOption{
		CAName:          "csi-local-plugin-ca",
		CAOrganizations: []string{"csi-local-plugin"},
		DNSNames:        []string{"csi-local-plugin"},
		CommonName:      "csi-local-plugin",
	}, begin, end)
	if err != nil {
		log.Fatalf("CreateCACert is failed, err:%s", err)
	}
	// server cert
	serverCert, serverKey, err := utils.CreateCertPEM(utils.CertOption{
		DNSNames:   []string{"csi-local-plugin-server"},
		CommonName: "csi-local-plugin-server",
	}, ca, begin, end, false)
	if err != nil {
		log.Fatalf("CreatesServerCertPEM is failed, err:%s", err)
	}

	// client cert
	clientCert, clientKey, err := utils.CreateCertPEM(utils.CertOption{
		CommonName: "csi-local-plugin-client",
	}, ca, begin, end, true)
	if err != nil {
		log.Fatalf("CreateClientCertPEM is failed, err:%s", err)
	}
	secret = prepareSecret(localSecretNamespace, localSecretName, map[string]string{}, map[string]string{})
	addToSecretFromData(secret, caCertFileName, ca.CertPEM)
	addToSecretFromData(secret, serverCertFileName, serverCert)
	addToSecretFromData(secret, serverKeyFileName, serverKey)
	addToSecretFromData(secret, clientCertFileName, clientCert)
	addToSecretFromData(secret, clientKeyFileName, clientKey)
	err = createSecret(secret)
	if err != nil {
		log.Fatalf("CreateSecret is failed, err:+%v", err)
	}
	return writeClientCert(ca.CertPEM, clientCert, clientKey)
}

// NewDriver create the identity/node/controller server and disk driver
func NewDriver(nodeID, endpoint string) *Local {
	initDriver()
	tmplvm := &Local{}
	tmplvm.endpoint = endpoint

	if nodeID == "" {
		nodeID = utils.RetryGetMetaData(InstanceID)
		log.Infof("Use node id : %s", nodeID)
	}

	// set driver name;
	driverName := defaultDriverName
	tmpValue := os.Getenv("DRIVER_VENDOR")
	if tmpValue != "" {
		driverName = tmpValue
	}

	// GlobalConfig Set
	GlobalConfigSet("", nodeID, driverName)

	csiDriver := csicommon.NewCSIDriver(driverName, csiVersion, nodeID)
	tmplvm.driver = csiDriver
	tmplvm.driver.AddControllerServiceCapabilities([]csi.ControllerServiceCapability_RPC_Type{
		csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
		csi.ControllerServiceCapability_RPC_EXPAND_VOLUME,
	})
	tmplvm.driver.AddVolumeCapabilityAccessModes([]csi.VolumeCapability_AccessMode_Mode{csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER})

	// Create GRPC servers
	tmplvm.idServer = newIdentityServer(tmplvm.driver)

	if os.Getenv(utils.ServiceType) == utils.ProvisionerService {
		//Create secret with TLS by Provisioner
		caCertFile, clientCertFile, clientKeyFile := createSecretWithTLS()
		tmplvm.controllerServer = newControllerServer(tmplvm.driver, caCertFile, clientCertFile, clientKeyFile)
	} else {
		tmplvm.nodeServer = NewNodeServer(tmplvm.driver, driverName, nodeID)
	}

	return tmplvm
}

// Run start a new server
func (local *Local) Run() {
	common.RunCSIServer(local.endpoint, local.idServer, local.controllerServer, local.nodeServer)
}

// GlobalConfigSet set Global Config
func GlobalConfigSet(region, nodeID, driverName string) {
	// Global Configs Set
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	if qps := os.Getenv("KUBE_CLI_API_QPS"); qps != "" {
		if qpsi, err := strconv.Atoi(qps); err == nil {
			cfg.QPS = float32(qpsi)
		}
	}
	if burst := os.Getenv("KUBE_CLI_API_BURST"); burst != "" {
		if qpsi, err := strconv.Atoi(burst); err == nil {
			cfg.Burst = qpsi
		}
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	nodeName := os.Getenv("KUBE_NODE_NAME")
	pmemEnable, pmeType := IsPmemSupported(nodeName, kubeClient)

	grpcProvision := true
	remoteConfig := os.Getenv(LOCAL_GRPC_PROVISION)
	if strings.ToLower(remoteConfig) == "false" {
		grpcProvision = false
	}

	hostNameAsTop := false
	hostNameEnv := os.Getenv(LOCAL_HOSTNAME_AS_TOPO)
	if strings.ToLower(hostNameEnv) == "true" {
		hostNameAsTop = true
	}

	topoKeyDefine := TopologyNodeKey
	topoKeyStr := os.Getenv(LOCAL_TOPO_KEY_DEFINED)
	if topoKeyStr != "" {
		log.Infof("Lcoal: use special topoloy key with LOCAL_TOPO_KEY_DEFINED: %s", topoKeyStr)
		topoKeyDefine = topoKeyStr
	}
	CapacityToNode := false
	capacityToNode := os.Getenv(CAPACITY_TO_NODE)
	if capacityToNode == "true" || capacityToNode == "yes" {
		log.Infof("Local: CAPACITY_TO_NODE is setting: %s", capacityToNode)
		CapacityToNode = true
	}
	localSparseTotalGi := ""
	if value, exists := os.LookupEnv(LOCAL_SPARSE_TOTAL_AVAILABLE_GI); exists {
		localSparseTotalGi = value
	}

	// Global Config Set
	types.GlobalConfigVar = types.GlobalConfig{
		Region:                           region,
		NodeID:                           nodeID,
		Scheduler:                        driverName,
		PmemEnable:                       pmemEnable,
		PmemType:                         pmeType,
		CapacityToNode:                   CapacityToNode,
		GrpcProvision:                    grpcProvision,
		KubeClient:                       kubeClient,
		HostNameAsTopo:                   hostNameAsTop,
		TopoKeyDefine:                    topoKeyDefine,
		LocalSparseFileDir:               os.Getenv(LOCAL_SPARSE_TEMPLATE_FILEDIR),
		LocalSparseFileTempSize:          os.Getenv(LOCAL_SPARSE_TEMPLATE_SIZE),
		LocalSparseTotalGi:               localSparseTotalGi,
		LocalSparseTotalAvailablePercent: os.Getenv(LOCAL_SPARSE_TOTAL_AVAILABLE_PERCENT),
	}
	log.Infof("Local Plugin Global Config is: %v", types.GlobalConfigVar)
}
