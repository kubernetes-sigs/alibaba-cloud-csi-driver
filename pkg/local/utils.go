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

package local

import (
	"context"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	// MetadataURL is metadata server url
	MetadataURL = "http://100.100.100.200/latest/meta-data/"
	// InstanceID is the instance id tag
	InstanceID = "instance-id"
	// RegionIDTag is the region id tag
	RegionIDTag = "region-id"
)

// ErrParse is an error that is returned when parse operation fails
var ErrParse = errors.New("Cannot parse output of blkid")

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

func formatDevice(devicePath, fstype string) error {
	output, err := exec.Command("mkfs", "-t", fstype, devicePath).CombinedOutput()
	if err != nil {
		return errors.New("FormatDevice error: " + string(output))
	}
	return nil
}

func isVgExist(vgName string) (bool, error) {
	vgCmd := fmt.Sprintf("%s vgdisplay %s 2>&1 | grep 'VG Name' | grep %s | grep -v grep | wc -l", NsenterCmd, vgName, vgName)
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
	instanceID := GetMetaData(InstanceID)
	regionID := GetMetaData(RegionIDTag)
	localDeviceNum := 0
	akID, akSecret, token := utils.GetDefaultAK()
	client := utils.NewEcsClient(akID, akSecret, token)

	// Get Instance Type
	request := ecs.CreateDescribeInstancesRequest()
	request.RegionId = regionID
	request.InstanceIds = "[\"" + instanceID + "\"]"
	instanceResponse, err := client.DescribeInstances(request)
	if err != nil {
		log.Errorf("getLocalDeviceNum: Describe Instance: %s Error: %s", instanceID, err.Error())
		return -1, err
	}
	if instanceResponse == nil || len(instanceResponse.Instances.Instance) == 0 {
		log.Infof("getLocalDeviceNum: Describe Instance Error, with empty response: %s", instanceID)
		return -1, err
	}

	// Get Instance LocalDisk Number
	instanceTypeID := instanceResponse.Instances.Instance[0].InstanceType
	instanceTypeFamily := instanceResponse.Instances.Instance[0].InstanceTypeFamily
	instanceTypeRequest := ecs.CreateDescribeInstanceTypesRequest()
	instanceTypeRequest.InstanceTypeFamily = instanceTypeFamily
	response, err := client.DescribeInstanceTypes(instanceTypeRequest)
	if err != nil {
		log.Errorf("getLocalDeviceNum: Describe Instance: %s, Type: %s, Family: %s Error: %s", instanceID, instanceTypeID, instanceTypeFamily, err.Error())
		return -1, err
	}
	for _, instance := range response.InstanceTypes.InstanceType {
		if instance.InstanceTypeId == instanceTypeID {
			localDeviceNum = instance.LocalStorageAmount
			log.Infof("getLocalDeviceNum: Instance: %s, InstanceType: %s, InstanceLocalDiskNum: %d", instanceID, instanceTypeID, localDeviceNum)
			break
		}
	}
	return localDeviceNum, nil
}

// create vg if not exist
func createVG(vgName string) (int, error) {
	pvNum := 0

	// step1: check vg is created or not
	vgCmd := fmt.Sprintf("%s vgdisplay %s 2>&1 | grep 'VG Name' | grep %s | grep -v grep | wc -l", NsenterCmd, vgName, vgName)
	vgline, err := utils.Run(vgCmd)
	if err != nil {
		return 0, err
	}
	if strings.TrimSpace(vgline) == "1" {
		pvNumCmd := fmt.Sprintf("%s vgdisplay %s 2>&1 | grep 'Cur PV' | grep -v grep | awk '{print $3}'", NsenterCmd, vgName)
		if pvNumStr, err := utils.Run(pvNumCmd); err != nil {
			return 0, err
		} else if pvNum, err = strconv.Atoi(strings.TrimSpace(pvNumStr)); err != nil {
			return 0, err
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
	//deviceStartChar := "b"
	deviceStartIndex := 0
	deviceNameLen := len(deviceStartWith)
	if deviceNameLen > 1 {
		deviceStartChar := deviceStartWith[deviceNameLen-1 : deviceNameLen]
		for index := 0; index < 15; index++ {
			if deviceStartChar == DeviceChars[index] {
				deviceStartIndex = index
			}
		}
		deviceNamePrefix = deviceStartWith[0 : deviceNameLen-1]
	}
	for i := deviceStartIndex; i < localDeviceNum; i++ {
		deviceName := deviceNamePrefix + DeviceChars[i]
		devicePath := filepath.Join("/dev", deviceName)
		localDeviceList = append(localDeviceList, devicePath)
	}
	log.Infof("doLocalVolumeMounts, Starting LocalDisk Mount: vgName: %s, LocalDisk Number: %d, LocalDisk: %v", vgName, localDeviceNum, localDeviceList)

	for _, devicePath := range localDeviceList {
		if !utils.IsFileExisting(devicePath) {
			log.Errorf("PV (%s) is not exist", devicePath)
			return 0, status.Error(codes.Internal, "PV is Not exit: "+devicePath)
		}
		pvCmd := fmt.Sprintf("%s pvdisplay %s 2>&1 | grep 'VG Name' | grep -v grep | awk '{print $3}'", NsenterCmd, devicePath)
		existVgName, err := utils.Run(pvCmd)
		if err != nil {
			log.Errorf("PV (%s) is Already in VG: %s", devicePath, strings.TrimSpace(existVgName))
			return 0, err
		}
	}
	localDeviceStr := strings.Join(localDeviceList, " ")
	vgAddCmd := fmt.Sprintf("%s vgcreate %s %s", NsenterCmd, vgName, localDeviceStr)
	_, err = utils.Run(vgAddCmd)
	if err != nil {
		log.Errorf("Add PV (%s) to VG: %s error: %s", localDeviceStr, strings.TrimSpace(vgName), err.Error())
		return 0, err
	}

	log.Infof("Successful add Local Disks to VG (%s): %v", vgName, localDeviceList)
	return localDeviceNum, nil
}

func getPVNumber(vgName string) int {
	var pvCount = 0
	vgList, err := server.ListVG()
	if err != nil {
		log.Errorf("Get pv for vg %s with error %s", vgName, err.Error())
		return 0
	}
	for _, vg := range vgList {
		if vg.Name == vgName {
			pvCount = int(vg.PvCount)
		}
	}
	return pvCount
}

func getPvObj(client kubernetes.Interface, volumeID string) (*v1.PersistentVolume, error) {
	return client.CoreV1().PersistentVolumes().Get(context.Background(), volumeID, metav1.GetOptions{})
}
func getPvSpec(client kubernetes.Interface, volumeID, driverName string) (string, string, *v1.PersistentVolume, error) {
	pv, err := getPvObj(client, volumeID)
	if err != nil {
		log.Errorf("Get Lvm Spec for volume %s, error with %v", volumeID, err)
		return "", "", nil, err
	}
	if pv.Spec.NodeAffinity == nil {
		log.Errorf("Get Lvm Spec for volume %s, with nil nodeAffinity", volumeID)
		return "", "", pv, errors.New("Get Lvm Spec for volume " + volumeID + ", with nil nodeAffinity")
	}
	if pv.Spec.NodeAffinity.Required == nil || len(pv.Spec.NodeAffinity.Required.NodeSelectorTerms) == 0 {
		log.Errorf("Get Lvm Spec for volume %s, with nil Required", volumeID)
		return "", "", pv, errors.New("Get Lvm Spec for volume " + volumeID + ", with nil Required")
	}
	if len(pv.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions) == 0 {
		log.Errorf("Get Lvm Spec for volume %s, with nil MatchExpressions", volumeID)
		return "", "", pv, errors.New("Get Lvm Spec for volume " + volumeID + ", with nil MatchExpressions")
	}
	key := pv.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions[0].Key
	if key != types.GlobalConfigVar.TopoKeyDefine && key != TopologyYodaNodeKey {
		log.Errorf("Get Lvm Spec for volume %s, with key %s", volumeID, key)
		return "", "", pv, errors.New("Get Lvm Spec for volume " + volumeID + ", with key" + key)
	}
	nodes := pv.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions[0].Values
	if len(nodes) == 0 {
		log.Errorf("Get Lvm Spec for volume %s, with empty nodes", volumeID)
		return "", "", pv, errors.New("Get Lvm Spec for volume " + volumeID + ", with empty nodes")
	}
	vgName := ""
	if value, ok := pv.Spec.CSI.VolumeAttributes["vgName"]; ok {
		vgName = value
	}

	log.Infof("Get Lvm Spec for volume %s, with VgName %s, Node %s", volumeID, pv.Spec.CSI.VolumeAttributes["vgName"], nodes[0])
	return nodes[0], vgName, pv, nil
}

// readIOPS: 1000
// writeIOPS: 10000
// readBPS: 100K
// writeBPS: 1M
func setVolumeIOLimit(devicePath string, req *csi.NodePublishVolumeRequest) error {
	readIOPS := req.VolumeContext["readIOPS"]
	writeIOPS := req.VolumeContext["writeIOPS"]
	readBPS := req.VolumeContext["readBPS"]
	writeBPS := req.VolumeContext["writeBPS"]

	// if no quota set, return;
	if readIOPS == "" && writeIOPS == "" && readBPS == "" && writeBPS == "" {
		return nil
	}

	// io limit parse
	readBPSInt, err := getBpsLimt(readBPS)
	if err != nil {
		log.Errorf("Volume(%s) Input Read BPS Limit format error: %s", req.VolumeId, err.Error())
		return err
	}
	writeBPSInt, err := getBpsLimt(writeBPS)
	if err != nil {
		log.Errorf("Volume(%s) Input Write BPS Limit format error: %s", req.VolumeId, err.Error())
		return err
	}
	readIOPSInt := 0
	if readIOPS != "" {
		readIOPSInt, err = strconv.Atoi(readIOPS)
		if err != nil {
			log.Errorf("Volume(%s) Input Read IOPS Limit format error: %s", req.VolumeId, err.Error())
			return err
		}
	}
	writeIOPSInt := 0
	if writeIOPS != "" {
		writeIOPSInt, err = strconv.Atoi(writeIOPS)
		if err != nil {
			log.Errorf("Volume(%s) Input Write IOPS Limit format error: %s", req.VolumeId, err.Error())
			return err
		}
	}

	// Get Device major/minor number
	majMinNum := getMajMinDevice(devicePath)
	if majMinNum == "" {
		log.Errorf("Volume(%s) Cannot get major/minor device number: %s", req.VolumeId, devicePath)
		return errors.New("Volume Cannot get major/minor device number: " + devicePath + req.VolumeId)
	}

	// Get pod uid
	podUid := req.VolumeContext["csi.storage.k8s.io/pod.uid"]
	if podUid == "" {
		log.Errorf("Volume(%s) Cannot get poduid and cannot set volume limit", req.VolumeId)
		return errors.New("Cannot get poduid and cannot set volume limit: " + req.VolumeId)
	}
	// /sys/fs/cgroup/blkio/kubepods.slice/kubepods-besteffort.slice/kubepods-besteffort-podaadcc749_6776_4933_990d_d50f260f5d46.slice/blkio.throttle.write_bps_device
	podUid = strings.ReplaceAll(podUid, "-", "_")
	podBlkIOPath := filepath.Join("/sys/fs/cgroup/blkio/kubepods.slice/kubepods-besteffort.slice", "kubepods-besteffort-pod"+podUid+".slice")
	if !utils.IsHostFileExist(podBlkIOPath) {
		podBlkIOPath = filepath.Join("/sys/fs/cgroup/blkio/kubepods.slice/kubepods-burstable.slice", "kubepods-besteffort-pod"+podUid+".slice")
	}
	if !utils.IsHostFileExist(podBlkIOPath) {
		log.Errorf("Volume(%s), Cannot get pod blkio/cgroup path: %s", req.VolumeId, podBlkIOPath)
		return errors.New("Cannot get pod blkio/cgroup path: " + podBlkIOPath)
	}

	// io limit set to blkio limit files
	if readIOPSInt != 0 {
		err := writeIoLimit(majMinNum, podBlkIOPath, "blkio.throttle.read_iops_device", readIOPSInt)
		if err != nil {
			return err
		}
	}
	if writeIOPSInt != 0 {
		err := writeIoLimit(majMinNum, podBlkIOPath, "blkio.throttle.write_iops_device", writeIOPSInt)
		if err != nil {
			return err
		}
	}
	if readBPSInt != 0 {
		err := writeIoLimit(majMinNum, podBlkIOPath, "blkio.throttle.read_bps_device", readBPSInt)
		if err != nil {
			return err
		}
	}
	if writeBPSInt != 0 {
		err := writeIoLimit(majMinNum, podBlkIOPath, "blkio.throttle.write_bps_device", writeBPSInt)
		if err != nil {
			return err
		}
	}
	log.Infof("Seccessful Set Volume(%s) IO Limit: readIOPS(%d), writeIOPS(%d), readBPS(%d), writeBPS(%d)", req.VolumeId, readIOPSInt, writeIOPSInt, readBPSInt, writeBPSInt)
	return nil
}

func writeIoLimit(majMinNum, podBlkIOPath, ioFile string, ioLimit int) error {
	targetPath := filepath.Join(podBlkIOPath, ioFile)
	content := majMinNum + " " + strconv.Itoa(ioLimit)
	cmd := fmt.Sprintf("%s sh -c 'echo %s > %s'", NsenterCmd, content, targetPath)
	_, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Errorf("WriteBps: Write file command(%s) error %v", cmd, err)
		return err
	}
	return nil
}

func getBpsLimt(bpsLimt string) (int, error) {
	if bpsLimt == "" {
		return 0, nil
	}

	bpsLimt = strings.ToLower(bpsLimt)
	convertNumber := 1
	intBpsStr := bpsLimt
	if strings.HasSuffix(bpsLimt, "k") && len(bpsLimt) > 1 {
		convertNumber = 1024 * 1024
		intBpsStr = strings.TrimSuffix(bpsLimt, "k")
	} else if strings.HasSuffix(bpsLimt, "m") && len(bpsLimt) > 1 {
		convertNumber = 1024 * 1024
		intBpsStr = strings.TrimSuffix(bpsLimt, "m")
	} else if strings.HasSuffix(bpsLimt, "g") && len(bpsLimt) > 1 {
		convertNumber = 1024 * 1024 * 1024
		intBpsStr = strings.TrimSuffix(bpsLimt, "g")
	}

	bpsIntValue, err := strconv.Atoi(intBpsStr)
	if err != nil {
		return 0, err
	}
	return bpsIntValue * convertNumber, nil
}

func getMajMinDevice(devicePath string) string {
	cmd := fmt.Sprintf("%s lsblk %s --noheadings --output MAJ:MIN", NsenterCmd, devicePath)
	out, err := exec.Command("sh", "-c", cmd).CombinedOutput()
	if err != nil {
		log.Errorf("getMajMinDevice with error: %s", err.Error())
		return ""
	}
	return strings.TrimSpace(string(out))
}
