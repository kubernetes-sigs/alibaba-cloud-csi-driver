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
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

const (
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// InstanceID is instance ID
	InstanceID = "instance-id"
	// RAMRoleResource is ram-role url subpath
	RAMRoleResource = "ram/security-credentials/"
	// RegionTag is region id
	RegionTag = "region-id"
	// Endpoint is OSS endpoint
	Endpoint = "oss-%s.aliyuncs.com"
	// InternalEndpoint is OSS internal endpoint
	InternalEndpoint = "oss-%s-internal.aliyuncs.com"
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

func getOssEndpoint(networkType, regionID string) (endpoint string) {
	if networkType == "vpc" {
		endpoint = fmt.Sprintf(InternalEndpoint, regionID)
	} else {
		endpoint = fmt.Sprintf(Endpoint, regionID)
	}
	return
}

func newOSSClient(customAccessKeyID, customAccessKeySecret, endpoint string) (ossClient *oss.Client, err error) {

	if customAccessKeyID != "" && customAccessKeySecret != "" {
		ossClient, err = oss.New(endpoint, customAccessKeyID, customAccessKeySecret)
	} else {
		accessKeyID, accessKeySecret, accessToken := utils.GetDefaultAK()
		ossClient, err = oss.New(endpoint, accessKeyID, accessKeySecret, oss.SecurityToken(accessToken))
	}

	if err != nil {
		return nil, err
	}
	return ossClient, nil
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
	checkMountCountCmd := fmt.Sprintf("%s mount | grep %s | grep %s | grep -v grep | wc -l", NsenterCmd, mountPath, OssFsType)
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
	checkMountCountCmd := fmt.Sprintf("%s mount | grep %s | grep %s | grep -v grep | wc -l", NsenterCmd, keyStr, OssFsType)
	out, err := utils.Run(checkMountCountCmd)
	if err != nil {
		return "0", err
	}
	return strings.TrimSpace(out), nil
}
