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
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
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

// checkRRSAParams check parameters of RRSA
func checkRRSAParams(opt *Options) error {
	if opt.RoleArn != "" && opt.OidcProviderArn != "" {
		return nil
	}
	if opt.RoleArn != "" || opt.OidcProviderArn != "" {
		return errors.New("Oss parameters error: use RRSA but one of the ARNs is empty, roleArn:" + opt.RoleArn + ", oidcProviderArn:" + opt.OidcProviderArn)
	}
	if opt.RoleName == "" {
		return errors.New("Oss parameters error: use RRSA but roleName is empty")
	}
	return nil
}

// getRRSAConfig get oidcProviderArn and roleArn
func getRRSAConfig(opt *Options, m metadata.MetadataProvider) (rrsaCfg *mounter.RrsaConfig, err error) {
	saName := fuseServieAccountName
	if opt.ServiceAccountName != "" {
		saName = opt.ServiceAccountName
	}

	if opt.OidcProviderArn != "" && opt.RoleArn != "" {
		return &mounter.RrsaConfig{ServiceAccountName: saName, OidcProviderArn: opt.OidcProviderArn, RoleArn: opt.RoleArn}, nil
	}

	accountId, err := m.Get(metadata.AccountID)
	if err != nil {
		return nil, errors.New("Get accountId error: " + err.Error())
	}
	clusterId, err := m.Get(metadata.ClusterID)
	if err != nil {
		return nil, errors.New("Get clusterId error: " + err.Error())
	}
	provider := mounter.GetOIDCProvider(clusterId)
	oidcProviderArn, roleArn := mounter.GetArn(provider, accountId, opt.RoleName)
	return &mounter.RrsaConfig{ServiceAccountName: saName, OidcProviderArn: oidcProviderArn, RoleArn: roleArn}, nil
}

// parseOtherOpts extracts mount options from parameters.otherOpts string.
// example: "-o max_stat_cache_size=0 -o allow_other -ogid=1000,uid=1000" => {"max_stat_cache_size=0", "allow_other", "uid=1000", "gid=1000"}
func parseOtherOpts(otherOpts string) (mountOptions []string, err error) {
	elements := strings.Fields(otherOpts)
	accepting := false
	for _, ele := range elements {
		if accepting {
			eles := strings.Split(ele, ",")
			mountOptions = append(mountOptions, eles...)
			accepting = false
		} else {
			if ele == "-o" {
				accepting = true
			} else if strings.HasPrefix(ele, "-o") {
				eles := strings.Split(strings.TrimPrefix(ele, "-o"), ",")
				mountOptions = append(mountOptions, eles...)
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
	source = fmt.Sprintf("%s:%s", opts.Bucket, opts.Path)
	mountOptions = append(mountOptions, fmt.Sprintf("url=%s", opts.URL))
	otherOpts, err := parseOtherOpts(opts.OtherOpts)
	if err != nil {
		return err
	}
	mountOptions = append(mountOptions, otherOpts...)
	logger := log.WithFields(log.Fields{
		"source":  source,
		"target":  target,
		"options": mountOptions,
		"fstype":  opts.FuseType,
	})
	err = mounter.Mount(source, target, "", mountOptions)
	if err != nil {
		logger.Errorf("failed to mount: %v", err)
		return status.Errorf(codes.Internal, "failed to mount: %v", err)
	}
	logger.Info("mounted successfully")
	return nil
}

func setNetworkType(originURL, regionID string) (URL string, modified bool) {
	URL = originURL
	if utils.IsPrivateCloud() || !strings.HasSuffix(strings.TrimRight(URL, "/"), ".aliyuncs.com") {
		return
	}
	// compatible with the old OSS accelerator endpoint, remove after it is deprecated
	endpoint := strings.TrimPrefix(strings.TrimPrefix(originURL, "https://"), "http://")
	if strings.HasPrefix(endpoint, "oss-cache-") {
		return
	}
	if strings.Contains(originURL, regionID) && !strings.Contains(originURL, "internal") {
		URL = strings.ReplaceAll(originURL, regionID, regionID+"-internal")
		modified = true
	}
	return
}

func setTransmissionProtocol(originURL string) (URL string, modified bool) {
	URL = originURL
	if strings.HasPrefix(URL, "http://") || strings.HasPrefix(URL, "https://") {
		return
	}
	if strings.HasSuffix(strings.TrimRight(URL, "/"), "-internal.aliyuncs.com") {
		return "http://" + URL, true
	}
	return "https://" + URL, true
}
