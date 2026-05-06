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
	"syscall"
	"time"

	ecs20140526 "github.com/alibabacloud-go/ecs-20140526/v7/client"
	eflo_controller20221215 "github.com/alibabacloud-go/eflo-controller-20221215/v3/client"
	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	alicred_old "github.com/aliyun/credentials-go/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/bmcpfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/ens"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/labeler"
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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sversion "k8s.io/apimachinery/pkg/util/version"
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
	runLabeler           = flag.Bool("run-labeler", false, "Run the centralized node metadata labeler (patches max-disk annotation and disk type labels on nodes)")
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

	var kubeClient kubernetes.Interface
	cfg, err := options.GetRestConfig()
	var k8sVersion *k8sversion.Version
	if err != nil {
		klog.Warningf("newGlobalConfig: build kubeconfig failed: %v", err)
	} else {
		kubeClient, err = kubernetes.NewForConfig(cfg)
		if err != nil {
			klog.Warningf("Error building kubernetes clientset: %v", err)
		} else {
			meta.EnableKubernetes(kubeClient)

			// Detect Kubernetes version for feature compatibility
			k8sVersion, err = detectKubernetesVersion(kubeClient)
			if err != nil {
				klog.Warningf("Failed to detect kubernetes version: %v", err)
			}
		}
	}

	regionID, err := meta.Get(metadata.RegionID)
	if err != nil {
		klog.ErrorS(err, "failed to get regionID for metadata, use cn-hangzhou")
		regionID = "cn-hangzhou"
	} else {
		// OIDC recognize this env to enable vpc endpoint
		if _, ok := os.LookupEnv("ALIBABA_CLOUD_STS_REGION"); !ok {
			if err := os.Setenv("ALIBABA_CLOUD_STS_REGION", regionID); err != nil {
				klog.ErrorS(err, "failed to set ALIBABA_CLOUD_STS_REGION")
			}
		}
	}

	provider, err := credentials.NewProvider()
	if err != nil {
		klog.ErrorS(err, "failed to get credential for metadata, will not enable OpenAPI")
	}
	var ecsClient *ecs20140526.Client
	if provider != nil {
		cred := alicred_old.FromCredentialsProvider(provider.GetProviderName(), provider)

		efloClient, err := eflo_controller20221215.NewClient(utils.GetEfloControllerConfig(regionID).SetCredential(cred))
		if err != nil {
			klog.ErrorS(err, "failed to get efloClient for metadata")
		} else {
			meta.EnableEFLO(efloClient)
		}

		ecsClient, err = ecs20140526.NewClient(utils.GetEcsConfig(regionID).SetCredential(cred))
		if err != nil {
			klog.ErrorS(err, "failed to get ecsClient for metadata")
		} else {
			// Goes after EFLO, because if EFLO API confirms it's a lingjun instance, we can skip ECS API.
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

	ctx, cancelSignals := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancelSignals()

	if *runLabeler {
		if kubeClient == nil {
			klog.Fatal("labeler requires Kubernetes client; cluster config unavailable")
		}
		if ecsClient == nil {
			klog.Fatal("labeler requires Alibaba Cloud ECS v2 client; credential may be unavailable")
		}
		wg.Go(func() {
			klog.InfoS("starting metadata labeler")
			if err := labeler.Run(ctx, kubeClient, ecsClient, regionID, labeler.Options{}); err != nil {
				klog.ErrorS(err, "labeler exited")
			}
		})
	}

	if serviceType != 0 {
		for _, driverName := range driverNames {
			endpoint := replaceCsiEndpoint(driverName, *endpoint)
			klog.Infof("CSI endpoint for driver %s: %s", driverName, endpoint)

			var driver *common.Servers
			switch driverName {
			case TypePluginNAS:
				driver = nas.NewServers(meta, endpoint, serviceType, csiCfg)
			case TypePluginOSS:
				driver = oss.NewServers(endpoint, meta, serviceType, csiCfg, k8sVersion)
			case TypePluginDISK:
				driver = disk.NewServers(meta, ecsClient, endpoint, serviceType, csiCfg)
			case TypePluginCPFS:
				klog.Fatalf("%s is no longer supported, please switch to %s if you are using CPFS 2.0 protocol server", TypePluginCPFS, TypePluginNAS)
			case TypePluginENS:
				driver = ens.NewServers(*nodeID, endpoint, serviceType)
			case ExtenderAgent:
				klog.Fatalf("rund-csi protocol 1.0 is no longer supported.")
			case TypePluginPOV:
				driver = pov.NewServers(meta, endpoint, serviceType)
			case TypePluginBMCPFS:
				driver = bmcpfs.NewServers(meta, endpoint, serviceType)
			default:
				klog.Fatalf("CSI start failed, not support driver: %s", driverName)
			}

			server := common.NewCSIServer(driverName, driver)
			wg.Go(func() {
				common.Serve(server, endpoint)
			})
			wg.Go(func() {
				<-ctx.Done()
				server.Stop()
			})
		}
	}

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

	httpServer := &http.Server{Addr: fmt.Sprintf(":%s", servicePort), Handler: csiMux}
	wg.Go(func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			klog.Fatalf("Service port listen and serve err:%s", err.Error())
		}
	})
	wg.Go(func() {
		<-ctx.Done()
		klog.Infof("%v, exiting...", context.Cause(ctx))
		err := httpServer.Shutdown(context.Background())
		if err != nil {
			klog.ErrorS(err, "httpServer shutdown failed")
		}
	})

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

// detectKubernetesVersion detects the Kubernetes server version
func detectKubernetesVersion(clientset kubernetes.Interface) (*k8sversion.Version, error) {
	serverVersion, err := clientset.Discovery().ServerVersion()
	if err != nil {
		return nil, fmt.Errorf("failed to get server version: %w", err)
	}

	// Parse the version string (format: "v1.31.0" or "v1.31.0-eks-...")
	k8sVersion, err := k8sversion.ParseSemantic(serverVersion.GitVersion)
	if err != nil {
		return nil, fmt.Errorf("failed to parse version %q: %w", serverVersion.GitVersion, err)
	}

	return k8sVersion, nil
}
