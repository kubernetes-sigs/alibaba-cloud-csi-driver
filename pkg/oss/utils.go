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
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"io/ioutil"
	"net/http"
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

// GetRAMRoleOption get command line's ram_role option
func GetRAMRoleOption(mntCmd string) string {
	ramRole := GetMetaData(RAMRoleResource)
	ramRoleOpt := MetadataURL + RAMRoleResource + ramRole
	mntCmd = fmt.Sprintf(mntCmd+" -oram_role=%s", ramRoleOpt)
	return mntCmd
}

// IsOssfsMounted return if oss mountPath is mounted
func IsOssfsMounted(mountPath string) bool {
	checkMountCountCmd := fmt.Sprintf("%s mount | grep %s | grep -E '%s|%s' | grep -v grep | wc -l", NsenterCmd, mountPath, OssFsType, JindoFsType)
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
	checkMountCountCmd := fmt.Sprintf("%s mount | grep %s | grep -E '%s|%s' | grep -v grep | wc -l", NsenterCmd, keyStr, OssFsType, JindoFsType)
	out, err := utils.Run(checkMountCountCmd)
	if err != nil {
		return "0", err
	}
	return strings.TrimSpace(out), nil
}
