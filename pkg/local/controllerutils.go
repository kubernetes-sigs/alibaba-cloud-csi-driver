/*
Copyright 2020 The Alibaba Cloud Authors.

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

package local

import (
	"encoding/json"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/adapter"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func lvmScheduled(storageSelected string, parameters map[string]string) (map[string]string, error) {
	vgName := ""
	paraList := map[string]string{}
	if storageSelected != "" {
		storageMap := map[string]string{}
		err := json.Unmarshal([]byte(storageSelected), &storageMap)
		if err != nil {
			paraList[VgNameTag] = storageSelected
			vgName = storageSelected
		} else if value, ok := storageMap["VolumeGroup"]; ok {
			paraList[VgNameTag] = value
			vgName = value
		}
	}
	if value, ok := parameters[VgNameTag]; ok && value != "" {
		if vgName != "" && value != vgName {
			return nil, status.Error(codes.InvalidArgument, "Storage Schedule is not expected "+value+vgName)
		}
	}
	if vgName == "" {
		return nil, status.Error(codes.InvalidArgument, "Node/Storage Schedule failed "+vgName)
	}
	return paraList, nil
}

func lvmPartScheduled(nodeSelected, pvcName, pvcNameSpace string, parameters map[string]string) (map[string]string, error) {
	vgName := ""
	paraList := map[string]string{}
	if value, ok := parameters[VgNameTag]; ok {
		vgName = value
	}
	if vgName == "" {
		volumeInfo, err := adapter.ScheduleVolume(LvmVolumeType, pvcName, pvcNameSpace, vgName, nodeSelected)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "lvm schedule with error "+err.Error())
		}
		if volumeInfo.VgName == "" || volumeInfo.Node == "" {
			log.Errorf("Lvm Schedule finished, but get empty: %v", volumeInfo)
			return nil, status.Error(codes.InvalidArgument, "lvm schedule finish but vgName/Node empty")
		}
		vgName = volumeInfo.VgName
	}
	paraList[VgNameTag] = vgName
	return paraList, nil
}

func lvmNoScheduled(parameters map[string]string) (string, map[string]string, error) {
	paraList := map[string]string{}
	return "", paraList, nil
}

func mountpointScheduled(storageSelected string, parameters map[string]string) (map[string]string, error) {
	mountpoint := ""
	paraList := map[string]string{}
	if storageSelected != "" {
		storageMap := map[string]string{}
		err := json.Unmarshal([]byte(storageSelected), &storageMap)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Scheduler provide error storage format: "+err.Error())
		}
		if value, ok := storageMap[MountPointType]; ok {
			paraList[MountPointType] = value
			mountpoint = value
		}
	}
	if mountpoint == "" {
		return nil, status.Error(codes.InvalidArgument, "Mountpoint Schedule failed "+mountpoint)
	}
	return paraList, nil
}

func mountpointPartScheduled(nodeSelected, pvcName, pvcNameSpace string, parameters map[string]string) (map[string]string, error) {
	paraList := map[string]string{}
	volumeInfo, err := adapter.ScheduleVolume(MountPointType, pvcName, pvcNameSpace, "", nodeSelected)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "lvm schedule with error "+err.Error())
	}
	if volumeInfo.Disk == "" {
		log.Errorf("mountpoint Schedule finished, but get empty Disk: %v", volumeInfo)
		return nil, status.Error(codes.InvalidArgument, "mountpoint schedule finish but Disk empty")
	}
	paraList[MountPointType] = volumeInfo.Disk
	return paraList, nil
}

func mountpointNoScheduled(parameters map[string]string) (string, map[string]string, error) {
	paraList := map[string]string{}
	return "", paraList, nil
}

func deviceScheduled(storageSelected string, parameters map[string]string) (map[string]string, error) {
	device := ""
	paraList := map[string]string{}
	if storageSelected != "" {
		storageMap := map[string]string{}
		err := json.Unmarshal([]byte(storageSelected), &storageMap)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Scheduler provide error storage format: "+err.Error())
		}
		if value, ok := storageMap[DeviceVolumeType]; ok {
			paraList[DeviceVolumeType] = value
			device = value
		}
	}
	if device == "" {
		return nil, status.Error(codes.InvalidArgument, "Device Schedule failed "+device)
	}
	return paraList, nil
}

func devicePartScheduled(nodeSelected, pvcName, pvcNameSpace string, parameters map[string]string) (map[string]string, error) {
	paraList := map[string]string{}
	volumeInfo, err := adapter.ScheduleVolume(DeviceVolumeType, pvcName, pvcNameSpace, "", nodeSelected)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "device schedule with error "+err.Error())
	}
	if volumeInfo.Disk == "" {
		log.Errorf("Device Schedule finished, but get empty Disk: %v", volumeInfo)
		return nil, status.Error(codes.InvalidArgument, "Device schedule finish but Disk empty")
	}
	paraList[DeviceVolumeType] = volumeInfo.Disk
	return paraList, nil
}

func deviceNoScheduled(parameters map[string]string) (string, map[string]string, error) {
	paraList := map[string]string{}
	return "", paraList, nil
}
