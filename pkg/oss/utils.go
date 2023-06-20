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
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
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

func GetMetaDataAsync(resource string) string {
	c1 := make(chan string, 1)
	go func(r string) {
		ans := GetMetaData(r)
		c1 <- ans
	}(resource)
	select {
	case res := <-c1:
		return res
	case <-time.After(2 * time.Second):
		return ""
	}
}

func GetGlobalMountPath(volumeId string) string {

	result := sha256.Sum256([]byte(volumeId))
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
// TEST::
func IsOssfsMounted(mountPath string) (bool, []string, error) {
	checkMountCountCmd := fmt.Sprintf("%s mount", NsenterCmd)
	out, err := utils.RunWithFilter(checkMountCountCmd, mountPath, "fuse.ossfs")
	if err != nil {
		return false, out, err
	}
	if len(out) == 0 {
		return false, nil, nil
	}
	return true, out, nil
}

// IsLastSharedVol return code status to help check if this oss volume uses UseSharedPath and is the last one
func IsLastSharedVol(pvName string) (string, error) {
	keyStr := fmt.Sprintf("volumes/kubernetes.io~csi/%s/mount", pvName)
	checkMountCountCmd := fmt.Sprintf("%s mount", NsenterCmd)
	out, err := utils.RunWithFilter(checkMountCountCmd, keyStr, "fuse.ossfs")
	if err != nil {
		return "0", err
	}
	return string(rune(len(out))), nil
}
