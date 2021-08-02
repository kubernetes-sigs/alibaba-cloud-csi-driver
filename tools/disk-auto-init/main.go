package main

import (
	"flag"
	"fmt"
	aliyunep "github.com/aliyun/alibaba-cloud-sdk-go/sdk/endpoints"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	utilexec "k8s.io/utils/exec"
	k8smount "k8s.io/utils/mount"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var (
	initModeInput    = flag.String("mode", "", "define disk init mode")
	deviceListInput  = flag.String("devices", "", "define which device to init")
	diskListInput    = flag.String("disks", "", "define which disk to init")
	mntPathListInput = flag.String("paths", "", "define the mountpoints of the disks")
)

var RegionID = "cn-hangzhou"
var InstanceID = ""
var K8sMounter = k8smount.New("")
var DiskMounter = &k8smount.SafeFormatAndMount{Interface: K8sMounter, Exec: utilexec.New()}

const (
	ExcludeFirstLastMode = "device-exclude-first-last"
	ExcludeFirstMode     = "device-exclude-first"
	DeviceDefinedMode    = "device-defined"
	DiskIdDefinedMode    = "device-disks"
)

func init() {
	RegionID, InstanceID, _ = utils.GetRegionAndInstanceID()
}

// disk-auto-init --mode device-exclude-first-last --paths /mnt/{devicename}
// disk-auto-init --mode device-exclude-first --paths /mnt/path{index}
// disk-auto-init --mode device-defined --devices /dev/vdb,/dev/vdc --paths /mnt/path1,/mnt/abc
// disk-auto-init --mode device-disks --disks d-2ze49dgr09ienk2exhma,d-2ze49dgr09ienk2exhmb --paths /home/kkk,/mnt/jjj

func main() {
	flag.Parse()

	diskInitMode := *initModeInput
	diskList := []string{}
	devicePathMap := map[string]string{}
	var err error
	log.Infof("Starting to init instance devices: %s", InstanceID)
	log.Infof("Starting to init instance devices, with Input: mode(%v), device(%v), disk(%v), path(%v)", *initModeInput, *deviceListInput, *diskListInput, *mntPathListInput)

	switch diskInitMode {
	case ExcludeFirstLastMode:
		deviceList := []string{}
		allDeviceList := getAllDevices()
		sort.Strings(allDeviceList)
		log.Infof("All Device List in this Instance: %v", allDeviceList)

		for i := 1; i < len(allDeviceList)-1; i++ {
			deviceList = append(deviceList, allDeviceList[i])
		}
		sort.Strings(deviceList)
		log.Infof("Target Device List in this Instance: %v", deviceList)

		devicePathMap, err = getDevicePathMap(deviceList, *mntPathListInput)
		if err != nil {
			log.Errorf("Get Device Path Map error: %v", err)
			os.Exit(1)
		}
		break
	case ExcludeFirstMode:
		deviceList := []string{}
		allDeviceList := getAllDevices()
		sort.Strings(allDeviceList)
		log.Infof("All Device List in this Instance: %v", allDeviceList)

		for i := 1; i < len(allDeviceList); i++ {
			deviceList = append(deviceList, allDeviceList[i])
		}
		sort.Strings(deviceList)
		log.Infof("Target Device List in this Instance: %v", deviceList)

		devicePathMap, err = getDevicePathMap(deviceList, *mntPathListInput)
		if err != nil {
			log.Errorf("Get Device Path Map error: %v", err)
			os.Exit(1)
		}
		break
	case DeviceDefinedMode:
		deviceList := []string{}
		if *deviceListInput != "" {
			deviceList = strings.Split(*deviceListInput, ",")
			if len(deviceList) != 0 {
				for _, deviceName := range deviceList {
					if !utils.IsFileExisting(deviceName) {
						log.Errorf("Device Defined mode, The input Device %s not exist", deviceName)
						os.Exit(1)
					}
				}
			} else {
				log.Errorf("Device Defined mode, device list empty: %s", *deviceListInput)
				os.Exit(1)
			}
		} else {
			log.Errorf("Device Defined mode, but not device input: %s", *deviceListInput)
			os.Exit(1)
		}
		sort.Strings(deviceList)
		log.Infof("Target Device List in this Instance: %v", deviceList)

		devicePathMap, err = getDevicePathMap(deviceList, *mntPathListInput)
		if err != nil {
			log.Errorf("Get Device Path Map error: %v", err)
			os.Exit(1)
		}
		break
	case DiskIdDefinedMode:
		if *diskListInput != "" {
			diskList = strings.Split(*diskListInput, ",")
			if len(diskList) != 0 {
				for _, disk := range diskList {
					disks := getDisk(disk)
					if len(disks) == 0 {
						log.Fatal("The input Disk %s not Found", disk)
					}
				}
			}
		}
	default:
		log.Errorf("Input mode is not supported: %s", *initModeInput)
		os.Exit(1)
	}
	log.Infof("Parse Device Path Map: %v", devicePathMap)

	for device, path := range devicePathMap {
		safeInitDevice(device, path)
	}

	log.Infof("Finished Init Devices...")
}

func safeInitDevice(device string, path string) error {
	log.Infof("Starting Mount Device %s to Path: %s", device, path)

	if !utils.IsFileExisting(device) {
		log.Fatal("The Input Device %s not Found", device)
	}
	mkfsOptions := make([]string, 0)
	mntOptions := make([]string, 0)
	mntOptions = append(mntOptions, "shared")

	if utils.IsFileExisting(path) {
		curDevice := getPathDevice(path)
		if curDevice != "" && curDevice == device {
			log.Infof("Device %s has Already mounted to Target Path %s", device, path)
			return nil
		}
		if curDevice != "" && curDevice != device {
			log.Fatalf("Path %s has Already mounted by other Device %s", path, curDevice)
			return nil
		}
	} else {
		utils.CreateDest(path)
		log.Infof("Create Path %s", path)
	}

	if err := disk.FormatAndMount(DiskMounter, device, path, "ext4", mkfsOptions, mntOptions); err != nil {
		log.Fatal("Device %s, Path %s, FormatAndMount got error %v", device, path, err)
	}
	log.Infof("Successful Mount Device %s to Path: %s", device, path)
	return nil
}

func getPathDevice(path string) string {
	cmd := fmt.Sprintf("findmnt %s -o SOURCE --noheadings | awk -F[ '{print $1}'", path)
	out, err := utils.Run(cmd)
	if err != nil {
		log.Fatal("Exec FindMnt error: %v", err)
	}
	return strings.TrimSpace(string(out))
}

func getDevicePathMap(deviceList []string, mntPathListStr string) (map[string]string, error) {
	devicePathMap := map[string]string{}
	if mntPathListStr != "" {
		if strings.Contains(mntPathListStr, "{devicename}") {
			pathPrexfix := strings.Replace(mntPathListStr, "{devicename}", "", 1)
			for _, deviceName := range deviceList {
				name := filepath.Base(deviceName)
				devicePathMap[deviceName] = filepath.Join(pathPrexfix, name)
			}
		} else if strings.Contains(mntPathListStr, "{index}") {
			pathPrexfix := strings.Replace(mntPathListStr, "{index}", "", 1)
			for index, deviceName := range deviceList {
				devicePathMap[deviceName] = pathPrexfix + strconv.Itoa(index+1)
			}
		} else {
			pathList := strings.Split(mntPathListStr, ",")
			if len(pathList) != len(deviceList) {
				return devicePathMap, fmt.Errorf("The input device list %v not match path list %v ", deviceList, pathList)
			}
			for i := 0; i < len(pathList); i++ {
				devicePathMap[deviceList[i]] = pathList[i]
			}
		}
	}
	return devicePathMap, nil
}

func getAllDevices() []string {
	devices := []string{}
	files, _ := ioutil.ReadDir("/dev")
	for _, file := range files {
		if !file.IsDir() && strings.HasPrefix(file.Name(), "vd") && !strings.HasSuffix(file.Name(), "1") {
			devices = append(devices, fmt.Sprintf("/dev/%s", file.Name()))
		}
	}
	return devices
}

func getEcsClient() (client *ecs.Client) {
	accessKeyID, accessSecret, accessToken := utils.GetDefaultAK()
	if accessToken != "" {
		client = newEcsClient(accessKeyID, accessSecret, accessToken)
	}
	return client
}

func newEcsClient(accessKeyID, accessKeySecret, accessToken string) (ecsClient *ecs.Client) {
	var err error
	if accessToken == "" {
		ecsClient, err = ecs.NewClientWithAccessKey(RegionID, accessKeyID, accessKeySecret)
		if err != nil {
			return nil
		}
	} else {
		ecsClient, err = ecs.NewClientWithStsToken(RegionID, accessKeyID, accessKeySecret, accessToken)
		if err != nil {
			return nil
		}
	}

	aliyunep.AddEndpointMapping(RegionID, "Ecs", "ecs-vpc."+RegionID+".aliyuncs.com")
	return
}

func getDisk(diskID string) []ecs.Disk {
	ecsClient := getEcsClient()
	// Step 1: Describe disk, if tag exist, return;
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = RegionID
	describeDisksRequest.DiskIds = "[\"" + diskID + "\"]"
	diskResponse, err := ecsClient.DescribeDisks(describeDisksRequest)
	if err != nil {
		log.Warnf("getDisk: error with DescribeDisks: %s, %s", diskID, err.Error())
		return []ecs.Disk{}
	}
	return diskResponse.Disks.Disk
}
