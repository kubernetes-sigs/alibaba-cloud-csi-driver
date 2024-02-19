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
	"path/filepath"
	"strings"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	mountutils "k8s.io/mount-utils"
)

const (
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// InstanceID is instance ID
	InstanceID = "instance-id"
	// RAMRoleResource is ram-role url subpath
	RAMRoleResource = "ram/security-credentials/"
)

func GetMetaDataAsync(resource string) string {
	c1 := make(chan string, 1)
	go func(r string) {
		ans, _ := utils.GetMetaData(r)
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
	ramRole, _ := utils.GetMetaData(RAMRoleResource)
	ramRoleOpt := MetadataURL + RAMRoleResource + ramRole
	mntCmdRamRole := fmt.Sprintf("ram_role=%s", ramRoleOpt)
	return mntCmdRamRole
}

// parseOtherOpts extracts mount options from parameters.otherOpts string.
// example: "-o max_stat_cache_size=0 -o allow_other" => {"max_stat_cache_size=0", "allow_other"}
func parseOtherOpts(otherOpts string) (mountOptions []string, err error) {
	elements := strings.Fields(otherOpts)
	accepting := false
	for _, ele := range elements {
		if accepting {
			mountOptions = append(mountOptions, ele)
			accepting = false
		} else {
			if ele == "-o" {
				accepting = true
			} else if strings.HasPrefix(ele, "-o") {
				mountOptions = append(mountOptions, strings.TrimPrefix(ele, "-o"))
			} else {
				// missing -o
				return nil, status.Errorf(codes.InvalidArgument, "invalid otherOpts: %q", otherOpts)
			}
		}
	}
	return mountOptions, nil
}

func doMount(mounter mountutils.Interface, target string, opts Options, mountOptions []string) error {
	var source string
	if opts.FuseType == JindoFsType {
		mountOptions = append(mountOptions, fmt.Sprintf("uri=oss://%s%s", opts.Bucket, opts.Path), fmt.Sprintf("fs.oss.endpoint=%s", opts.URL))
	} else {
		source = fmt.Sprintf("%s:%s", opts.Bucket, opts.Path)
		mountOptions = append(mountOptions, fmt.Sprintf("url=%s", opts.URL))
		otherOpts, err := parseOtherOpts(opts.OtherOpts)
		if err != nil {
			return err
		}
		mountOptions = append(mountOptions, otherOpts...)
	}
	logger := log.WithFields(log.Fields{
		"source":  source,
		"target":  target,
		"options": mountOptions,
		"fstype":  opts.FuseType,
	})
	err := mounter.Mount(source, target, "", mountOptions)
	if err != nil {
		logger.Errorf("failed to mount: %v", err)
		return status.Errorf(codes.Internal, "failed to mount: %v", err)
	}
	logger.Info("mounted successfully")
	return nil
}

func modifiedURL(originURL, regionID string) (URL string, modified bool) {
	URL = originURL
	if strings.Contains(originURL, regionID) && !strings.Contains(originURL, "internal") && !utils.IsPrivateCloud() {
		URL = strings.ReplaceAll(originURL, regionID, regionID+"-internal")
		modified = true
	}
	if strings.HasPrefix(URL, "http://") || strings.HasPrefix(URL, "https://") {
		return
	}
	if strings.HasSuffix(URL, "-internal.aliyuncs.com") {
		return "http://" + URL, true
	}
	return "https://" + URL, true
}
