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
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	alicred_old "github.com/aliyun/credentials-go/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/bmcpfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/ens"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	csilog "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/log"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/metric"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/om"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/pov"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	"github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"
	"golang.org/x/sys/unix"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/component-base/logs"
	logsapi "k8s.io/component-base/logs/api/v1"
	_ "k8s.io/component-base/logs/json/register"
	"k8s.io/klog/v2"
)

const (
	// TypePluginSuffix is the suffix of all storage plugins.
	TypePluginSuffix = "plugin.csi.alibabacloud.com"
	// TypePluginVar is the yaml variable that needs to be replaced.
	TypePluginVar = "driverplugin.csi.alibabacloud.com-replace"
	// PluginServicePort default port is 11260.
	PluginServicePort = "11260"
	// ProvisionerServicePort default port is 11270.
	ProvisionerServicePort = "11270"
	// TypePluginDISK DISK type plugin
	TypePluginDISK = "diskplugin.csi.alibabacloud.com"
	// TypePluginNAS NAS type plugin
	TypePluginNAS = "nasplugin.csi.alibabacloud.com"
	// TypePluginOSS OSS type plugin
	TypePluginOSS = "ossplugin.csi.alibabacloud.com"
	// TypePluginCPFS CPFS type plugin
	TypePluginCPFS = "cpfsplugin.csi.alibabacloud.com"
	// TypePluginENS ENS type plugins
	TypePluginENS = "ensplugin.csi.alibabacloud.com"
	// TypePluginPOV POV type plugins
	TypePluginPOV = "povplugin.csi.alibabacloud.com"
	// TypePluginBMCPFS BMCPFS type plugin
	TypePluginBMCPFS = "bmcpfsplugin.csi.alibabacloud.com"
	// ExtenderAgent agent component
	ExtenderAgent = "agent"
)

var (
	endpoint             = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	nodeID               = flag.String("nodeid", "", "node id")
	runAsController      = flag.Bool("run-as-controller", false, "Only run as controller service (deprecated)")
	runControllerService = flag.Bool("run-controller-service", true, "activate CSI controller service")
	runNodeService       = flag.Bool("run-node-service", true, "activate CSI node service")
	driver               = flag.String("driver", TypePluginDISK, "CSI Driver")
	// Deprecated: rootDir is instead by KUBELET_ROOT_DIR env.
	rootDir = flag.String("rootdir", "/var/lib/kubelet/csi-plugins", "Kubernetes root directory")
)

func setupFlags() {
	csilog.RedirectLogrusToLogr(logrus.StandardLogger(), klog.Background())

	utils.AddGoFlags(flag.CommandLine)

	fg := features.FunctionalMutableFeatureGate
	err := logsapi.AddFeatureGates(fg)
	if err != nil {
		panic(err) // just panic since the log infra is not initialized yet.
	}
	flag.Var(fg, "feature-gates", "A set of key=value pairs that describe feature gates for alpha/experimental features. "+
		"Options are:\n"+strings.Join(fg.KnownFeatures(), "\n"))

	c := logsapi.NewLoggingConfiguration()
	logsapi.AddFlags(c, flag.CommandLine)
	logs.InitLogs()
	flag.Parse()
	if err := logsapi.ValidateAndApply(c, fg); err != nil {
		klog.ErrorS(err, "LoggingConfiguration is invalid")
		klog.FlushAndExit(klog.ExitFlushTimeout, 1)
	}
}

func main() {
	setupFlags()
	serviceType := utils.GetServiceType(*runAsController, *runControllerService, *runNodeService)

	// enable pprof analyse
	pprofPort := os.Getenv("PPROF_PORT")
	if pprofPort != "" {
		if _, err := strconv.Atoi(pprofPort); err == nil {
			klog.Infof("enable pprof & start port at %v", pprofPort)
			go func() {
				err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%v", pprofPort), nil)
				klog.Errorf("start server err: %v", err)
			}()
		}
	}

	klog.Infof("Multi CSI Driver Name: %s, nodeID: %s, endPoints: %s", *driver, *nodeID, *endpoint)
	klog.Infof("CSI Driver, Version: %s, Release time: %s", version.VERSION, version.GetTime())

	multiDriverNames := *driver
	driverNames := strings.Split(multiDriverNames, ",")
	var wg sync.WaitGroup

	// Storage devops
	go om.StorageOM()

	// initialize node metadata
	meta := metadata.NewMetadata()
	meta.EnableIMDS(http.DefaultTransport)

	cfg, err := options.GetRestConfig()
	if err != nil {
		klog.Warningf("newGlobalConfig: build kubeconfig failed: %v", err)
	} else {
		kubeClient, err := kubernetes.NewForConfig(cfg)
		if err != nil {
			klog.Warningf("Error building kubernetes clientset: %v", err)
		} else {
			meta.EnableKubernetes(kubeClient)
		}
	}

	provider, err := credentials.NewProvider()
	if err != nil {
		klog.ErrorS(err, "failed to get credential for metadata, will not enable OpenAPI")
	} else {
		regionID, err := meta.Get(metadata.RegionID)
		if err != nil {
			klog.ErrorS(err, "failed to get regionID for metadata, use cn-hangzhou")
			regionID = "cn-hangzhou"
		}
		cred := alicred_old.FromCredentialsProvider(provider.GetProviderName(), provider)

		ecsClient, err := ecs20140526.NewClient(utils.GetEcsConfig(regionID).SetCredential(cred))
		if err != nil {
			klog.ErrorS(err, "failed to get ecsClient for metadata")
		} else {
			meta.EnableOpenAPI(ecsClient)
		}

		stsClient, err := sts20150401.NewClient(utils.GetStsConfig(regionID).SetCredential(cred))
		if err != nil {
			klog.ErrorS(err, "failed to get stsClient for metadata")
		} else {
			meta.EnableSts(stsClient)
		}
	}

	for i, driverName := range driverNames {
		if !strings.Contains(driverName, TypePluginSuffix) && driverName != ExtenderAgent {
			driverNames[i] = joinCsiPluginSuffix(driverName)
		}
	}

	csiCfg := getCSIPluginConfig()

	for _, driverName := range driverNames {
		wg.Add(1)
		endPointName := replaceCsiEndpoint(driverName, *endpoint)
		klog.Infof("CSI endpoint for driver %s: %s", driverName, endPointName)

		switch driverName {
		case TypePluginNAS:
			go func(endPoint string) {
				defer wg.Done()
				driver := nas.NewDriver(meta, endPoint, serviceType, csiCfg)
				driver.Run()
			}(endPointName)
		case TypePluginOSS:
			go func(endPoint string) {
				defer wg.Done()
				driver := oss.NewDriver(endPoint, meta, serviceType, csiCfg)
				driver.Run()
			}(endPointName)
		case TypePluginDISK:
			go func(endPoint string) {
				defer wg.Done()
				driver := disk.NewDriver(meta, endPoint, serviceType, csiCfg)
				driver.Run()
			}(endPointName)

		case TypePluginCPFS:
			klog.Fatalf("%s is no longer supported, please switch to %s if you are using CPFS 2.0 protocol server", TypePluginCPFS, TypePluginNAS)
		case TypePluginENS:
			go func(endpoint string) {
				defer wg.Done()
				driver := ens.NewDriver(*nodeID, endpoint, serviceType)
				driver.Run()
			}(endPointName)
		case ExtenderAgent:
			klog.Fatalf("rund-csi protocol 1.0 is no longer supported.")
		case TypePluginPOV:
			go func(endPoint string) {
				defer wg.Done()
				driver := pov.NewDriver(meta, endPoint, serviceType)
				driver.Run()
			}(endPointName)
		case TypePluginBMCPFS:
			go func(endpoint string) {
				defer wg.Done()
				driver := bmcpfs.NewDriver(meta, endpoint, serviceType)
				driver.Run()
			}(endPointName)
		default:
			klog.Fatalf("CSI start failed, not support driver: %s", driverName)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, unix.SIGTERM)
	go func() {
		s := <-c
		klog.Infof("Got signal: %v, exiting...", s)
		os.Exit(0)
	}()

	servicePort := os.Getenv("SERVICE_PORT")

	if servicePort == "" {
		servicePort = PluginServicePort
		if serviceType&utils.Controller != 0 {
			servicePort = ProvisionerServicePort
		}
	}

	version.SetPrometheusVersion()
	enableMetric := true
	if os.Getenv("ENABLE_METRIC") == "false" {
		enableMetric = false
	}

	klog.Info("CSI is running status.")
	csiMux := http.NewServeMux()
	csiMux.HandleFunc("/healthz", healthHandler)
	klog.Infof("Healthz listening on address: /healthz")
	if enableMetric {
		metricHandler := metric.NewMetricHandler(driverNames, serviceType)
		csiMux.Handle("/metrics", metricHandler)
		klog.Infof("Metric listening on address: /metrics")
	}

	if err := http.ListenAndServe(fmt.Sprintf(":%s", servicePort), csiMux); err != nil {
		klog.Fatalf("Service port listen and serve err:%s", err.Error())
	}

	wg.Wait()
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
	_, _ = w.Write([]byte(message))
}

func getCSIPluginConfig() (config utils.Config) {
	cfg, err := options.GetRestConfig()
	if err != nil {
		klog.ErrorS(err, "error building kube config")
		return
	}

	client, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		if cfg != nil {
			klog.Fatal(err, "error building kube client")
		} else {
			klog.ErrorS(err, "error building kube client")
			return
		}
	}

	cm, err := client.CoreV1().ConfigMaps("kube-system").Get(context.Background(), "csi-plugin", metav1.GetOptions{})
	if err != nil {
		klog.ErrorS(err, "failed to get config map csi-plugin")
		return
	}

	config.ConfigMap = cm.Data
	return
}
