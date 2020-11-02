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

package dbfs

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dbfs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	// MetadataURL is metadata url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// RegionTag is region id
	RegionTag = "region-id"
)

var (
	// VERSION should be updated by hand at each release
	VERSION = "v1.14.8"
	// GITCOMMIT will be overwritten automatically by the build system
	GITCOMMIT = "HEAD"
	// KubernetesAlicloudIdentity is the system identity for ecs client request
	KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Nas-%s", VERSION)
)


func DoDBFSMount(opt *Options, mountPoint string, volumeID string) error {
	log.Infof("DoDBFSMount: do nothing")
	return nil
}

//CreateDest create the target
func CreateDest(dest string) error {
	fi, err := os.Lstat(dest)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dest, 0777); err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	if fi != nil && !fi.IsDir() {
		return fmt.Errorf("%v already exist but it's not a directory", dest)
	}
	return nil
}

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

func updateDbfsClient(client *dbfs.Client) *dbfs.Client {
	accessKeyID, accessSecret, accessToken := utils.GetDefaultAK()

	if accessToken != "" {
		client = newDbfsClient(accessKeyID, accessSecret, accessToken, GlobalConfigVar.Region)
	}
	return client
}

func newDbfsClient(accessKeyID, accessKeySecret, accessToken, regionID string) (dbfsClient *dbfs.Client) {
	var err error
	if regionID == "" {
		regionID = GetMetaData(RegionTag)
	}
	if accessToken == "" {
		dbfsClient, err = dbfs.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)
		if err != nil {
			return nil
		}
	} else {
		dbfsClient, err = dbfs.NewClientWithStsToken(regionID, accessKeyID, accessKeySecret, accessToken)
		if err != nil {
			return nil
		}
	}

	return
}
