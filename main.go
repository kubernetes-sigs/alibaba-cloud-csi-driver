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
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cpfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/lvm"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/oss"
)

func init() {
	flag.Set("logtostderr", "true")
}

const (
	LOGFILE_PREFIX   = "/var/log/alicloud/"
	MB_SIZE          = 1024 * 1024
	TYPE_PLUGIN_DISK = "diskplugin.csi.alibabacloud.com"
	TYPE_PLUGIN_NAS  = "nasplugin.csi.alibabacloud.com"
	TYPE_PLUGIN_OSS  = "ossplugin.csi.alibabacloud.com"
	TYPE_PLUGIN_CPFS = "cpfsplugin.csi.alibabacloud.com"
	TYPE_PLUGIN_LVM  = "lvmplugin.csi.alibabacloud.com"
)

var _BRANCH_ = ""
var _VERSION_ = ""
var _COMMITID_ = ""
var _BUILDTIME_ = ""

var (
	endpoint        = flag.String("endpoint", "unix://tmp/csi.sock", "CSI endpoint")
	nodeId          = flag.String("nodeid", "", "node id")
	runAsController = flag.Bool("run-as-controller", false, "Only run as controller service")
	driver          = flag.String("driver", TYPE_PLUGIN_DISK, "CSI Driver")
	rootDir         = flag.String("rootdir", "/var/lib/kubelet", "Kubernetes root directory")
)

// Nas CSI Plugin
func main() {
	flag.Parse()

	// set log config
	setLogAttribute(*driver)

	if err := createPersistentStorage(path.Join(*rootDir, "plugins", *driver, "controller")); err != nil {
		log.Errorf("failed to create persistent storage for controller: %v", err)
		os.Exit(1)
	}
	if err := createPersistentStorage(path.Join(*rootDir, "plugins", *driver, "node")); err != nil {
		log.Errorf("failed to create persistent storage for node: %v", err)
		os.Exit(1)
	}

	drivername := *driver
	log.Infof("CSI Driver Name: %s, %s, %s", drivername, *nodeId, *endpoint)
	log.Infof("CSI Driver Branch: %s, Version: %s, Build time: %s\n", _BRANCH_, _VERSION_, _BUILDTIME_)
	if drivername == TYPE_PLUGIN_NAS {
		driver := nas.NewDriver(*nodeId, *endpoint)
		driver.Run()
	} else if drivername == TYPE_PLUGIN_OSS {
		driver := oss.NewDriver(*nodeId, *endpoint)
		driver.Run()
	} else if drivername == TYPE_PLUGIN_DISK {
		driver := disk.NewDriver(*nodeId, *endpoint, *runAsController)
		driver.Run()
	} else if drivername == TYPE_PLUGIN_LVM {
		driver := lvm.NewDriver(*nodeId, *endpoint)
		driver.Run()
	} else if drivername == TYPE_PLUGIN_CPFS {
		driver := cpfs.NewDriver(*nodeId, *endpoint)
		driver.Run()
	} else {
		log.Errorf("CSI start failed, not support driver: %s", drivername)
	}

	os.Exit(0)
}

func createPersistentStorage(persistentStoragePath string) error {
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

	os.MkdirAll(LOGFILE_PREFIX, os.FileMode(0755))
	logFile := LOGFILE_PREFIX + driver + ".log"
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(1)
	}

	// rotate the log file if too large
	if fi, err := f.Stat(); err == nil && fi.Size() > 2*MB_SIZE {
		f.Close()
		timeStr := time.Now().Format("-2006-01-02-15:04:05")
		timedLogfile := LOGFILE_PREFIX + driver + timeStr + ".log"
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
