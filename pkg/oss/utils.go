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

package oss

import (
	"crypto/sha256"
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

const (
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// InstanceID is instance ID
	InstanceID = "instance-id"
	// RAMRoleResource is ram-role url subpath
	RAMRoleResource = "ram/security-credentials/"
)

// GetMetaData get host regionid, zoneid
func GetMetaData(resource string) string {
	resp, err := http.Get(MetadataURL + resource)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

func GetGlobalMountPath(volumeId string) string {

	result := sha256.Sum256([]byte(fmt.Sprintf("%s", volumeId)))
	volSha := fmt.Sprintf("%x", result)

	globalFileVer1 := filepath.Join(utils.KubeletRootDir, "/plugins/kubernetes.io/csi/pv/", volumeId, "/globalmount")
	globalFileVer2 := filepath.Join(utils.KubeletRootDir, "/plugins/kubernetes.io/csi/", driverName, volSha, "/globalmount")

	if utils.IsFileExisting(globalFileVer1) {
		return globalFileVer1
	} else {
		return globalFileVer2
	}
}

// GetRAMRoleOption get command line's ram_role option
func GetRAMRoleOption() string {
	ramRole := GetMetaData(RAMRoleResource)
	ramRoleOpt := MetadataURL + RAMRoleResource + ramRole
	mntCmdRamRole := fmt.Sprintf("-oram_role=%s", ramRoleOpt)
	return mntCmdRamRole
}

// IsOssfsMounted return if oss mountPath is mounted
func IsOssfsMounted(mountPath string) bool {
	checkMountCountCmd := fmt.Sprintf("%s mount | grep %s | grep -E 'fuse.ossfs|fuse.jindo-fuse' | grep -v grep | wc -l", NsenterCmd, mountPath)
	out, err := utils.Run(checkMountCountCmd)
	if err != nil {
		return false
	}
	if strings.TrimSpace(out) == "0" {
		return false
	}
	return true
}

// IsLastSharedVol return code status to help check if this oss volume uses UseSharedPath and is the last one
func IsLastSharedVol(pvName string) (string, error) {
	keyStr := fmt.Sprintf("volumes/kubernetes.io~csi/%s/mount", pvName)
	checkMountCountCmd := fmt.Sprintf("%s mount | grep %s | grep -E 'fuse.ossfs|fuse.jindo-fuse' | grep -v grep | wc -l", NsenterCmd, keyStr)
	out, err := utils.Run(checkMountCountCmd)
	if err != nil {
		return "0", err
	}
	return strings.TrimSpace(out), nil
}

func WriteMetricsInfo(metricsPathPrefix string, req *csi.NodePublishVolumeRequest, opt Options) {
	podUIDPath := metricsPathPrefix + req.VolumeContext["csi.storage.k8s.io/pod.uid"] + "/"
	mountPointPath := podUIDPath + req.GetVolumeId() + "/"
	podInfoName := "pod_info"
	mountPointName := "mount_point_info"
	if !utils.IsFileExisting(mountPointPath) {
		_ = utils.MkdirAll(mountPointPath, os.FileMode(0755))
	}
	if !utils.IsFileExisting(podUIDPath + podInfoName) {
		info := req.VolumeContext["csi.storage.k8s.io/pod.namespace"] + " " +
			req.VolumeContext["csi.storage.k8s.io/pod.name"] + " " +
			req.VolumeContext["csi.storage.k8s.io/pod.uid"] + " " +
			opt.MetricsTop
		_ = utils.WriteAndSyncFile(podUIDPath+podInfoName, []byte(info), os.FileMode(0644))
	}

	if !utils.IsFileExisting(mountPointPath + mountPointName) {
		info := "ossfs" + " " +
			"oss" + " " +
			opt.Bucket + " " +
			req.GetVolumeId() + " " +
			req.TargetPath
		_ = utils.WriteAndSyncFile(mountPointPath+mountPointName, []byte(info), os.FileMode(0644))
	}
}
