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
	"io"
	"os"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/agent"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cpfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/lvm"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mem"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/oss"
	log "github.com/sirupsen/logrus"
)

func init() {
	flag.Set("logtostderr", "true")
}

const (
	// LogfilePrefix prefix of log file
	LogfilePrefix = "/var/log/alicloud/"
	// MBSIZE MB size
	MBSIZE = 1024 * 1024
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

var (
	endpoint        = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	nodeID          = flag.String("nodeid", "", "node id")
	runAsController = flag.Bool("run-as-controller", false, "Only run as controller service")
	driver          = flag.String("driver", TypePluginDISK, "CSI Driver")
	rootDir         = flag.String("rootdir", "/var/lib/kubelet/csi-plugins", "Kubernetes root directory")
)

// Nas CSI Plugin
func main() {
	flag.Parse()

	// set log config
	setLogAttribute(*driver)

	if err := createPersistentStorage(path.Join(*rootDir, *driver, "controller")); err != nil {
		log.Errorf("failed to create persistent storage for controller: %v", err)
		os.Exit(1)
	}
	if err := createPersistentStorage(path.Join(*rootDir, *driver, "node")); err != nil {
		log.Errorf("failed to create persistent storage for node: %v", err)
		os.Exit(1)
	}

	driverName := *driver
	endPointName := *endpoint
	endPointNames := strings.Split(endPointName, ",")
	log.Infof("CSI Driver Name: %s, %s, %s", driverName, *nodeID, *endpoint)
	log.Infof("CSI Driver Branch: %s, Version: %s, Build time: %s\n", BRANCH, VERSION, BUILDTIME)
	driverNames := strings.Split(driverName, ",")
	var wg sync.WaitGroup
	for _, driverName := range driverNames {
		for _, endPointName := range endPointNames {
			if driverName == TypePluginNAS && strings.Contains(endPointName, TypePluginNAS) {
				go func() {
					defer wg.Done()
					driver := nas.NewDriver(*nodeID, endPointName)
					driver.Run()
				}()
			} else if driverName == TypePluginOSS && strings.Contains(endPointName, TypePluginOSS) {
				go func() {
					defer wg.Done()
					driver := oss.NewDriver(*nodeID, *endpoint)
					driver.Run()
				}()
			} else if driverName == TypePluginDISK && strings.Contains(endPointName, TypePluginDISK) {
				go func() {
					defer wg.Done()
					driver := disk.NewDriver(*nodeID, *endpoint, *runAsController)
					driver.Run()
				}()
			} else if driverName == TypePluginLVM && strings.Contains(endPointName, TypePluginLVM) {
				go func() {
					defer wg.Done()
					driver := lvm.NewDriver(*nodeID, *endpoint)
					driver.Run()
				}()
			} else if driverName == TypePluginCPFS && strings.Contains(endPointName, TypePluginCPFS) {
				go func() {
					defer wg.Done()
					driver := cpfs.NewDriver(*nodeID, *endpoint)
					driver.Run()
				}()
			} else if driverName == TypePluginMEM && strings.Contains(endPointName, TypePluginMEM) {
				go func() {
					defer wg.Done()
					driver := mem.NewDriver(*nodeID, *endpoint)
					driver.Run()
				}()
			} else if driverName == TypePluginLOCAL && strings.Contains(endPointName, TypePluginLOCAL) {
				go func() {
					defer wg.Done()
					driver := local.NewDriver(*nodeID, *endpoint)
					driver.Run()
				}()
			} else if driverName == ExtenderAgent && strings.Contains(endPointName, ExtenderAgent) {
				go func() {
					defer wg.Done()
					queryServer := agent.NewAgent()
					queryServer.RunAgent()
				}()
			} else {
				log.Errorf("CSI start failed, not support driver: %s", driverName)
			}
		}
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
