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

package main

import (
	"flag"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/metric"
	"io"
	"net/http"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/agent"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cpfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/dbfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/ens"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/lvm"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mem"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/om"
	_ "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/common/version"
	log "github.com/sirupsen/logrus"
)

func init() {
	flag.Set("logtostderr", "true")
}

func setPrometheusVersion() {
	version.Version = VERSION
	version.Revision = REVISION
	version.Branch = BRANCH
	version.BuildDate = BUILDTIME
}

const (
	// LogfilePrefix prefix of log file
	LogfilePrefix = "/var/log/alicloud/"
	// MBSIZE MB size
	MBSIZE = 1024 * 1024
	// TypePluginSuffix is the suffix of all storage plugins.
	TypePluginSuffix = "plugin.csi.alibabacloud.com"
	// TypePluginVar is the yaml variable that needs to be replaced.
	TypePluginVar = "driverplugin.csi.alibabacloud.com-replace"
	//PluginServicePort default port is 11260.
	PluginServicePort = "11260"
	//ProvisionerServicePort default port is 11270.
	ProvisionerServicePort = "11270"
	// TypePluginDBFS local type plugin
	TypePluginDBFS = "dbfsplugin.csi.alibabacloud.com"
	// TypePluginDISK DISK type plugin
	TypePluginDISK = "diskplugin.csi.alibabacloud.com"
	// TypePluginNAS NAS type plugin
	TypePluginNAS = "nasplugin.csi.alibabacloud.com"
	// TypePluginOSS OSS type plugin
	TypePluginOSS = "ossplugin.csi.alibabacloud.com"
	// TypePluginCPFS CPFS type plugin
	TypePluginCPFS = "cpfsplugin.csi.alibabacloud.com"
	// TypePluginLVM LVM type plugin
	TypePluginLVM = "lvmplugin.csi.alibabacloud.com"
	// TypePluginMEM memory type plugin
	TypePluginMEM = "memplugin.csi.alibabacloud.com"
	// TypePluginLOCAL local type plugin
	TypePluginLOCAL = "localplugin.csi.alibabacloud.com"
	// TypePluginYODA local type plugin
	TypePluginYODA = "yodaplugin.csi.alibabacloud.com"
	// TypePluginENS ENS type plugins
	TypePluginENS = "ensplugin.csi.alibabacloud.com"
	// ExtenderAgent agent component
	ExtenderAgent = "agent"
)

// BRANCH is CSI Driver Branch
var BRANCH = ""

// VERSION is CSI Driver Version
var VERSION = ""

// COMMITID is CSI Driver CommitID
var COMMITID = ""

// BUILDTIME is CSI Driver Buildtime
var BUILDTIME = ""

// REVISION is CSI Driver Revision
var REVISION = ""

var (
	endpoint        = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	nodeID          = flag.String("nodeid", "", "node id")
	runAsController = flag.Bool("run-as-controller", false, "Only run as controller service")
	driver          = flag.String("driver", TypePluginDISK, "CSI Driver")
	// Deprecated: rootDir is instead by KUBELET_ROOT_DIR env.
	rootDir = flag.String("rootdir", "/var/lib/kubelet/csi-plugins", "Kubernetes root directory")
)

type globalMetricConfig struct {
	enableMetric bool
	serviceType  string
}

// Nas CSI Plugin
func main() {
	flag.Parse()
	serviceType := os.Getenv(utils.ServiceType)

	if len(serviceType) == 0 || serviceType == "" {
		serviceType = utils.PluginService
	}

	// When serviceType is neither plugin nor provisioner, the program will exits.
	if serviceType != utils.PluginService && serviceType != utils.ProvisionerService {
		log.Fatalf("Service type is unknown:%s", serviceType)
	}

	var logAttribute string
	switch serviceType {
	case utils.ProvisionerService:
		logAttribute = strings.Replace(TypePluginSuffix, utils.PluginService, utils.ProvisionerService, -1)
	case utils.PluginService:
		logAttribute = TypePluginSuffix
	default:
	}

	setLogAttribute(logAttribute)

	log.Infof("Multi CSI Driver Name: %s, nodeID: %s, endPoints: %s", *driver, *nodeID, *endpoint)
	log.Infof("CSI Driver Branch: %s, Version: %s, Build time: %s\n", BRANCH, VERSION, BUILDTIME)

	multiDriverNames := *driver
	endPointName := *endpoint
	driverNames := strings.Split(multiDriverNames, ",")
	var wg sync.WaitGroup

	// Storage devops
	go om.StorageOM()

	for _, driverName := range driverNames {
		wg.Add(1)
		if !strings.Contains(driverName, TypePluginSuffix) && driverName != ExtenderAgent {
			driverName = joinCsiPluginSuffix(driverName)
			if strings.Contains(*endpoint, TypePluginVar) {
				endPointName = replaceCsiEndpoint(driverName, *endpoint)
			} else {
				log.Fatalf("Csi endpoint:%s", *endpoint)
			}
		}
		if driverName == TypePluginYODA {
			driverName = TypePluginLOCAL
		}
		if err := createPersistentStorage(path.Join(utils.KubeletRootDir, "/csi-plugins", driverName, "controller")); err != nil {
			log.Errorf("failed to create persistent storage for controller: %v", err)
			os.Exit(1)
		}
		if err := createPersistentStorage(path.Join(utils.KubeletRootDir, "/csi-plugins", driverName, "node")); err != nil {
			log.Errorf("failed to create persistent storage for node: %v", err)
			os.Exit(1)
		}
		switch driverName {
		case TypePluginNAS:
			go func(endPoint string) {
				defer wg.Done()
				driver := nas.NewDriver(*nodeID, endPoint, serviceType)
				driver.Run()
			}(endPointName)
		case TypePluginOSS:
			go func(endPoint string) {
				defer wg.Done()
				driver := oss.NewDriver(*nodeID, endPoint)
				driver.Run()
			}(endPointName)
		case TypePluginDISK:
			go func(endPoint string) {
				defer wg.Done()
				driver := disk.NewDriver(*nodeID, endPoint, *runAsController)
				driver.Run()
			}(endPointName)

		case TypePluginLVM:
			go func(endPoint string) {
				defer wg.Done()
				driver := lvm.NewDriver(*nodeID, endPoint)
				driver.Run()
			}(endPointName)
		case TypePluginCPFS:
			go func(endPoint string) {
				defer wg.Done()
				driver := cpfs.NewDriver(*nodeID, endPoint)
				driver.Run()
			}(endPointName)
		case TypePluginMEM:
			go func(endPoint string) {
				defer wg.Done()
				driver := mem.NewDriver(*nodeID, endPoint)
				driver.Run()
			}(endPointName)
		case TypePluginLOCAL:
			go func(endPoint string) {
				defer wg.Done()
				driver := local.NewDriver(*nodeID, endPoint)
				driver.Run()
			}(endPointName)
		case TypePluginDBFS:
			go func(endPoint string) {
				defer wg.Done()
				driver := dbfs.NewDriver(*nodeID, endPoint)
				driver.Run()
			}(endPointName)
		case TypePluginENS:
			go func(endpoint string) {
				defer wg.Done()
				driver := ens.NewDriver(*nodeID, endpoint)
				driver.Run()
			}(endPointName)
		case ExtenderAgent:
			go func() {
				defer wg.Done()
				queryServer := agent.NewAgent()
				queryServer.RunAgent()
			}()
		default:
			log.Fatalf("CSI start failed, not support driver: %s", driverName)
		}
	}
	servicePort := os.Getenv("SERVICE_PORT")

	if len(servicePort) == 0 || servicePort == "" {
		switch serviceType {
		case utils.PluginService:
			servicePort = PluginServicePort
		case utils.ProvisionerService:
			servicePort = ProvisionerServicePort
		default:
		}
	}

	metricConfig := &globalMetricConfig{
		true,
		"plugin",
	}

	enableMetric := os.Getenv("ENABLE_METRIC")
	if enableMetric == "false" {
		setPrometheusVersion()
		metricConfig.enableMetric = false
	}
	metricConfig.serviceType = serviceType

	log.Info("CSI is running status.")
	server := &http.Server{Addr: ":" + servicePort}

	http.HandleFunc("/healthz", healthHandler)
	log.Infof("Metric listening on address: /healthz")
	if metricConfig.enableMetric {
		metricHandler := metric.NewMetricHandler(metricConfig.serviceType)
		http.Handle("/metrics", metricHandler)
		log.Infof("Metric listening on address: /metrics")
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Service port listen and serve err:%s", err.Error())
	}
	wg.Wait()
	os.Exit(0)
}
func createPersistentStorage(persistentStoragePath string) error {
	log.Infof("Create Stroage Path: %s", persistentStoragePath)
	return os.MkdirAll(persistentStoragePath, os.FileMode(0755))
}

// rotate log file by 2M bytes
// default print log to stdout and file both.
func setLogAttribute(driver string) {
	logType := os.Getenv("LOG_TYPE")
	logType = strings.ToLower(logType)
	if logType != "stdout" && logType != "host" {
		logType = "both"
	}
	if logType == "stdout" {
		return
	}

	os.MkdirAll(LogfilePrefix, os.FileMode(0755))
	logFile := LogfilePrefix + driver + ".log"
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(1)
	}

	// rotate the log file if too large
	if fi, err := f.Stat(); err == nil && fi.Size() > 2*MBSIZE {
		f.Close()
		timeStr := time.Now().Format("-2006-01-02-15:04:05")
		timedLogfile := LogfilePrefix + driver + timeStr + ".log"
		os.Rename(logFile, timedLogfile)
		f, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			os.Exit(1)
		}
	}
	if logType == "both" {
		mw := io.MultiWriter(os.Stdout, f)
		log.SetOutput(mw)
	} else {
		log.SetOutput(f)
	}
}

func joinCsiPluginSuffix(storageType string) string {
	return storageType + TypePluginSuffix
}

func replaceCsiEndpoint(pluginType string, endPointName string) string {
	return strings.Replace(endPointName, TypePluginVar, pluginType, -1)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	time := time.Now()
	message := "Liveness probe is OK, time:" + time.String()
	w.Write([]byte(message))
}
