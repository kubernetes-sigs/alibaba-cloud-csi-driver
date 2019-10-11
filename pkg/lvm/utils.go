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

package lvm

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	METADATA_URL = "http://100.100.100.200/latest/meta-data/"
	INSTANCE_ID  = "instance-id"
	REGIONID_TAG = "region-id"
)

// get host regionid, zoneid
func GetMetaData(resource string) string {
	resp, err := http.Get(METADATA_URL + resource)
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

func formatDevice(devicePath, fstype string) error {
	output, err := exec.Command("mkfs", "-t", fstype, devicePath).CombinedOutput()
	if err != nil {
		return errors.New("FormatDevice error: " + string(output))
	}
	return nil
}

func checkFSType(devicePath string) (string, error) {
	// We use `file -bsL` to determine whether any filesystem type is detected.
	// If a filesystem is detected (ie., the output is not "data", we use
	// `blkid` to determine what the filesystem is. We use `blkid` as `file`
	// has inconvenient output.
	// We do *not* use `lsblk` as that requires udev to be up-to-date which
	// is often not the case when a device is erased using `dd`.
	output, err := exec.Command("file", "-bsL", devicePath).CombinedOutput()
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(string(output)) == "data" {
		return "", nil
	}
	output, err = exec.Command("blkid", "-c", "/dev/null", "-o", "export", devicePath).CombinedOutput()
	if err != nil {
		return "", err
	}
	parseErr := errors.New("Cannot parse output of blkid.")
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		fields := strings.Split(strings.TrimSpace(line), "=")
		if len(fields) != 2 {
			return "", parseErr
		}
		if fields[0] == "TYPE" {
			return fields[1], nil
		}
	}
	return "", parseErr
}

func isVgExist(vgName string) (bool, error) {
	vgCmd := fmt.Sprintf("%s vgdisplay %s | grep 'VG Name' | grep %s | grep -v grep | wc -l", NSENTER_CMD, vgName, vgName)
	vgline, err := utils.Run(vgCmd)
	if err != nil {
		return false, err
	}
	if strings.TrimSpace(vgline) == "1" {
		return true, nil
	}
	return false, nil
}

// Get Local Disk Number from ecs API
// Requirements: The instance must have role which contains ecs::DescribeInstances, ecs::DescribeInstancesType.
func getLocalDeviceNum() (int, error) {
	instanceId := GetMetaData(INSTANCE_ID)
	regionId := GetMetaData(REGIONID_TAG)
	localDeviceNum := 0
	akId, akSecret, token := "", "", ""
	akId, akSecret = utils.GetLocalAK()
	if akId == "" || akSecret == "" {
		akId, akSecret, token = utils.GetSTSAK()
	}
	client := utils.NewEcsClient(akId, akSecret, token)

	// Get Instance Type
	request := ecs.CreateDescribeInstancesRequest()
	request.RegionId = regionId
	request.InstanceIds = "[\"" + instanceId + "\"]"
	instanceResponse, err := client.DescribeInstances(request)
	if err != nil {
		log.Errorf("getLocalDeviceNum: Describe Instance: %s Error: %s", instanceId, err.Error())
		return -1, err
	}
	if instanceResponse == nil || len(instanceResponse.Instances.Instance) == 0 {
		log.Infof("getLocalDeviceNum: Describe Instance Error, with empty response: %s", instanceId)
		return -1, err
	}

	// Get Instance LocalDisk Number
	instanceTypeId := instanceResponse.Instances.Instance[0].InstanceType
	instanceTypeFamily := instanceResponse.Instances.Instance[0].InstanceTypeFamily
	instanceTypeRequest := ecs.CreateDescribeInstanceTypesRequest()
	instanceTypeRequest.InstanceTypeFamily = instanceTypeFamily
	response, err := client.DescribeInstanceTypes(instanceTypeRequest)
	if err != nil {
		log.Errorf("getLocalDeviceNum: Describe Instance: %s, Type: %s, Family: %s Error: %s", instanceId, instanceTypeId, instanceTypeFamily, err.Error())
		return -1, err
	}
	for _, instance := range response.InstanceTypes.InstanceType {
		if instance.InstanceTypeId == instanceTypeId {
			localDeviceNum = instance.LocalStorageAmount
			log.Infof("getLocalDeviceNum: Instance: %s, InstanceType: %s, InstanceLocalDiskNum: %d", instanceId, instanceTypeId, localDeviceNum)
			break
		}
	}
	return localDeviceNum, nil
}

// create vg if not exist
func createVG(vgName string) (int, error) {
	pvNum := 0

	// step1: check vg is created or not
	vgCmd := fmt.Sprintf("%s vgdisplay %s | grep 'VG Name' | grep %s | grep -v grep | wc -l", NSENTER_CMD, vgName, vgName)
	vgline, err := utils.Run(vgCmd)
	if err != nil {
		return 0, err
	}
	if strings.TrimSpace(vgline) == "1" {
		pvNumCmd := fmt.Sprintf("%s vgdisplay %s | grep 'Cur PV' | grep -v grep | awk '{print $3}'", NSENTER_CMD, vgName)
		if pvNumStr, err := utils.Run(pvNumCmd); err != nil {
			return 0, err
		} else {
			if pvNum, err = strconv.Atoi(strings.TrimSpace(pvNumStr)); err != nil {
				return 0, err
			}
		}
		return pvNum, nil
	}

	// Step 2: Get LocalDisk Number
	localDeviceList := []string{}
	localDeviceNum, err := getLocalDeviceNum()
	if err != nil {
		log.Errorf("LocalDiskMount:: Get Local Disk Number Error, Error: %s", err.Error())
		return 0, status.Error(codes.Internal, "Get Local Disk Number Error")
	}
	if localDeviceNum < 1 {
		log.Errorf("VG not exist and also not local disk exist, vgName: %s", vgName)
		return 0, status.Error(codes.Internal, "VG not exist, and also not local disk exist")
	}

	// Step 3: Get LocalDisk device
	deviceStartWith := "vdb"
	deviceNamePrefix := "vd"
	deviceStartChar := "b"
	deviceStartIndex := 0
	deviceNameLen := len(deviceStartWith)
	if deviceNameLen > 1 {
		deviceStartChar = deviceStartWith[deviceNameLen-1 : deviceNameLen]
		for index := 0; index < 15; index++ {
			if deviceStartChar == DEVICE_CHARS[index] {
				deviceStartIndex = index
			}
		}
		deviceNamePrefix = deviceStartWith[0 : deviceNameLen-1]
	}
	for i := deviceStartIndex; i < localDeviceNum; i++ {
		deviceName := deviceNamePrefix + DEVICE_CHARS[i]
		devicePath := filepath.Join("/dev", deviceName)
		localDeviceList = append(localDeviceList, devicePath)
	}
	log.Infof("doLocalVolumeMounts, Starting LocalDisk Mount: vgName: %s, LocalDisk Number: %d, LocalDisk: %v", vgName, localDeviceNum, localDeviceList)

	for _, devicePath := range localDeviceList {
		if !utils.IsFileExisting(devicePath) {
			log.Errorf("PV (%s) is not exist", devicePath)
			return 0, status.Error(codes.Internal, "PV is Not exit: "+devicePath)
		}
		pvCmd := fmt.Sprintf("%s pvdisplay %s | grep 'VG Name' | grep -v grep | awk '{print $3}'", NSENTER_CMD, devicePath)
		existVgName, err := utils.Run(pvCmd)
		if err != nil {
			log.Errorf("PV (%s) is Already in VG: %s", devicePath, strings.TrimSpace(existVgName))
			return 0, err
		}
	}
	localDeviceStr := strings.Join(localDeviceList, " ")
	vgAddCmd := fmt.Sprintf("%s vgcreate %s %s", NSENTER_CMD, vgName, localDeviceStr)
	_, err = utils.Run(vgAddCmd)
	if err != nil {
		log.Errorf("Add PV (%s) to VG: %s error: %s", localDeviceStr, strings.TrimSpace(vgName), err.Error())
		return 0, err
	}

	log.Infof("Successful add Local Disks to VG (%s): %v", vgName, localDeviceList)
	return localDeviceNum, nil
}
