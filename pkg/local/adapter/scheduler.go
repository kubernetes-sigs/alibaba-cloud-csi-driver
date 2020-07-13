/*
Copyright 2020 The Kubernetes Authors.

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

package adapter

import (
	"encoding/json"
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/client"
	log "github.com/sirupsen/logrus"
	"os"
)

// BindingInfo represents the pvc and disk/lvm mapping
type BindingInfo struct {
	// node is the name of selected node
	Node string `json:"node"`
	// path for mount point
	Disk string `json:"disk"`
	// VgName is the name of selected volume group
	VgName string `json:"vgName"`
	// Device is the name for raw block device: /dev/vdb
	Device string `json:"device"`
	// [lvm] or [disk] or [device] or [quota]
	VolumeType string `json:"volumeType"`

	// PersistentVolumeClaim is the metakey for pvc: {namespace}/{name}
	PersistentVolumeClaim string `json:"persistentVolumeClaim"`
}

const (
	// SchedulerHostTag tag
	SchedulerHostTag = "SCHEDULER_HOST"
)

// URLHost default value
var UrlHost = "http://yoda-scheduler-extender-service:23000"

// ScheduleVolume make request and get expect schedule topology
func ScheduleVolume(volumeType, pvcName, pvcNamespace, vgName, nodeID string) (*BindingInfo, error) {
	bindingInfo := &BindingInfo{}
	hostEnv := os.Getenv(SchedulerHostTag)
	if hostEnv != "" {
		UrlHost = hostEnv
	}

	// make request url
	urlPath := fmt.Sprintf("/apis/scheduling/%s/persistentvolumeclaims/%s?nodeName=%s&volumeType=%s&vgName=%s", pvcNamespace, pvcName, nodeID, volumeType, vgName)
	if nodeID == "" && vgName != "" {
		urlPath = fmt.Sprintf("/apis/scheduling/%s/persistentvolumeclaims/%s?volumeType=%s&vgName=%s", pvcNamespace, pvcName, volumeType, vgName)
	} else if nodeID != "" && vgName == "" {
		urlPath = fmt.Sprintf("/apis/scheduling/%s/persistentvolumeclaims/%s?volumeType=%s&nodeName=%s", pvcNamespace, pvcName, volumeType, nodeID)
	} else if nodeID == "" && vgName == "" {
		urlPath = fmt.Sprintf("/apis/scheduling/%s/persistentvolumeclaims/%s?volumeType=%s", pvcNamespace, pvcName, volumeType)
	}
	url := UrlHost + urlPath

	// Request restful api
	respBody, err := client.DoRequest(url)
	if err != nil {
		log.Errorf("Schedule Volume with Url(%s) get error: %s", url, err.Error())
		return nil, err
	}
	// unmarshal json result.
	err = json.Unmarshal(respBody, bindingInfo)
	if err != nil {
		log.Errorf("Schedule Volume with Url(%s) get Unmarshal error: %s, and response: %s", url, err.Error(), string(respBody))
		return nil, err
	}

	log.Infof("Schedule Volume with Url(%s) Finished, get result: %v, %v", url, bindingInfo, string(respBody))
	return bindingInfo, nil
}

// ExpandVolume do volume capacity check
func ExpandVolume(pvcNameSpace, pvcName string, newSize int) error {
	urlPath := fmt.Sprintf("/apis/expand/%s/persistentvolumeclaims/%s?newSize=%d", pvcNameSpace, pvcName, newSize)
	url := UrlHost + urlPath

	// Request restful api
	respBody, err := client.DoRequest(url)
	if err != nil {
		log.Errorf("Volume Expand with Url(%s) get error: %s", url, err.Error())
		return err
	}

	log.Infof("Volume Expand with Url(%s) Finished, get result: %s", url, string(respBody))
	return nil
}
