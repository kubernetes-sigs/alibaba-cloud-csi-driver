package local

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	disk2 "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/manager"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// StorageCapacity define storage pool details
type StorageCapacity struct {
	Type       string `json:"type,omitempty"`
	Name       string `json:"name,omitempty"`
	Capacity   uint64 `json:"capacity,omitempty"`
	DeviceType string `json:"devicetype,omitempty"`
	MountPoint string `json:"mountPoint,omitempty"`
}

// LocalDeviceList the local device type list, only support alibaba cloud local disks;
var LocalDeviceList = []*StorageCapacity{}

// LocalDeviceUpdate alibaba cloud local disk is constant, no need reconcile
var LocalDeviceUpdate = false

// LocalDeviceLoopIndex local device check index
var LocalDeviceLoopIndex = 0

// updateNodeCapacity update node capacity annotations with the realtime storage capacities.
func updateNodeCapacity() {
	log.Infof("updateNodeCapacity: Starting to update volume capacity to node")
	CacheStorageCapacity := getCapacityFromNode()
	for {
		vgList := getVolumeGroup()
		qpList := getQuotaPath()
		ldList := getLoopDeviceCapacity()
		// local disk is not changed, no need reconcile
		if !LocalDeviceUpdate {
			updateLocalDiskList()
		}
		same := isTopologySame(vgList, qpList, ldList, CacheStorageCapacity)
		log.Infof("updateNodeCapacity: same: %v vgList: %v, ldList: %v, cacheCapacity: %v", same, vgList, ldList, CacheStorageCapacity)
		if !same {
			for _, item := range vgList {
				log.Infof("updateNodeCapacity: volumeGroup capacity %++v", item)
			}
			for _, item := range qpList {
				log.Infof("updateNodeCapacity: QuotaPath capacity %++v", item)
			}
			for _, item := range LocalDeviceList {
				log.Infof("updateNodeCapacity: Local Device capacity %++v", item)
			}
			for _, item := range CacheStorageCapacity {
				log.Infof("updateNodeCapacity: Cached capacity %++v", item)
			}

			err := updateCapacityToNode(vgList, qpList, ldList)
			if err == nil {
				CacheStorageCapacity = vgList
				CacheStorageCapacity = append(CacheStorageCapacity, qpList...)
				CacheStorageCapacity = append(CacheStorageCapacity, ldList...)
				CacheStorageCapacity = append(CacheStorageCapacity, LocalDeviceList...)
			}
		}

		// sleep with reconcile interval
		time.Sleep(time.Duration(60) * time.Second)
	}
}

// update node local disks
func updateLocalDiskList() {
	ac := utils.GetAccessControl()
	client := utils.NewEcsClient(ac)
	NodeRegionID, _ := utils.GetMetaData(RegionIDTag)
	NodeInstanceID, _ := utils.GetMetaData(InstanceID)
	if LocalDeviceLoopIndex > 5 {
		return
	}
	LocalDeviceLoopIndex = LocalDeviceLoopIndex + 1

	// Describe Disks with instanceID
	describeDisksRequest := ecs.CreateDescribeDisksRequest()
	describeDisksRequest.RegionId = NodeRegionID
	describeDisksRequest.InstanceId = NodeInstanceID
	diskResponse, err := client.DescribeDisks(describeDisksRequest)
	if err != nil {
		log.Errorf("updateLocalDiskList: Describe Disks for %s with error: %s", NodeInstanceID, err.Error())
		return
	}
	for _, disk := range diskResponse.Disks.Disk {
		if disk.Category == "local_ssd_pro" || disk.Category == "local_hdd_pro" {
			devicePaths, err := disk2.GetDeviceByVolumeID(disk.DiskId)
			rootDevice, subDevice, err := disk2.GetRootSubDevicePath(devicePaths)
			if err != nil {
				log.Errorf("updateLocalDiskList: get device by VolumeID(%s) with error: %s", disk.DiskId, err.Error())
				return
			}
			deviceName := disk2.ChooseDevice(rootDevice, subDevice)
			capacity := &StorageCapacity{
				Capacity:   uint64(utils.Gi2Bytes(int64(disk.Size))),
				Type:       "device",
				Name:       deviceName,
				DeviceType: disk.Category,
			}
			LocalDeviceList = append(LocalDeviceList, capacity)
		}
	}
	if len(LocalDeviceList) != 0 {
		for _, item := range LocalDeviceList {
			log.Infof("updateLocalDiskList: add local device with: %+v", item)
		}
	}
	log.Infof("updateLocalDiskList: Successful Update Local Device")
	LocalDeviceUpdate = true
	return
}

// get volumegroups from the node
func getVolumeGroup() []*StorageCapacity {
	resVGList := []*StorageCapacity{}
	vgList, err := server.ListVG()
	if err != nil {
		return nil
	}
	for _, item := range vgList {
		vg := &StorageCapacity{}
		vg.Name = item.Name
		vg.Capacity = item.Size
		vg.Type = "lvm"
		resVGList = append(resVGList, vg)
	}
	return resVGList
}

func getLoopDeviceCapacity() []*StorageCapacity {
	lp := manager.NewLoopDevice(types.GlobalConfigVar.LocalSparseFileDir, types.GlobalConfigVar.LocalSparseFileTempSize)
	usedBytes, err := lp.GetUsedByteSize()
	log.Infof("getLoopDevice: get usedBytes: %v", usedBytes)
	if err != nil {
		log.Errorf("getLoopDevice: failed to get used bytes, err: %v", err)
		return nil
	}
	var totalAvailableBytes int64
	if types.GlobalConfigVar.LocalSparseTotalGi == "" {
		totalBytes, err := lp.GetTempDirTotalCapacity()
		if err != nil {
			log.Errorf("getLoopDevice: failed to get totalcapacity: %v. err: %v", totalBytes, err)
			return nil
		}
		percent := 0.9
		if types.GlobalConfigVar.LocalSparseTotalAvailablePercent != "" {
			p, _ := strconv.Atoi(types.GlobalConfigVar.LocalSparseTotalAvailablePercent)
			percent = float64(p)
		}
		totalAvailableBytes = int64(float64(totalBytes) * percent)
	} else {
		totalGi, err := strconv.Atoi(types.GlobalConfigVar.LocalSparseTotalGi)
		if err != nil {
			log.Errorf("getLoopDevice: failed to convert LocalSparseTotalGi: %s to int. err: %v", types.GlobalConfigVar.LocalSparseTotalGi, err)
			return nil
		}
		totalAvailableBytes = utils.Gi2Bytes(int64(totalGi))
	}
	log.Errorf("getLoopDevice: total available bytes: %v", totalAvailableBytes)
	availableBytes := totalAvailableBytes - usedBytes
	sc := StorageCapacity{
		Name:     types.GlobalConfigVar.LocalSparseFileDir,
		Type:     "loopdevice",
		Capacity: uint64(availableBytes),
	}
	scList := []*StorageCapacity{}
	scList = append(scList, &sc)
	return scList
}

// get quotapath from local node
// not support now
func getQuotaPath() []*StorageCapacity {
	return nil
}

// isTopologySame check the current volume capacity is same as cached.
func isTopologySame(vgList, qpList, ldList, cachedList []*StorageCapacity) bool {
	if len(vgList)+len(qpList)+len(ldList)+len(LocalDeviceList) != len(cachedList) {
		return false
	}
	for _, item := range cachedList {
		itemSearched := false
		for _, vgItem := range vgList {
			if isCapacityObjectSame(item, vgItem) {
				itemSearched = true
				break
			}
		}
		if itemSearched {
			continue
		}
		for _, qpItem := range qpList {
			if isCapacityObjectSame(item, qpItem) {
				itemSearched = true
				break
			}
		}

		if itemSearched {
			continue
		}

		for _, ldItem := range ldList {
			if isCapacityObjectSame(item, ldItem) {
				itemSearched = true
				break
			}
		}

		if itemSearched {
			continue
		}
		for _, deviceItem := range LocalDeviceList {
			if isCapacityObjectSame(item, deviceItem) {
				itemSearched = true
				break
			}
		}
		if !itemSearched {
			return false
		}
	}
	return true
}

func isCapacityObjectSame(item1, item2 *StorageCapacity) bool {
	if item1.Capacity != item2.Capacity {
		return false
	}
	if item1.DeviceType != item2.DeviceType {
		return false
	}
	if item1.Name != item2.Name {
		return false
	}
	if item1.MountPoint != item2.MountPoint {
		return false
	}
	if item1.Type != item2.Type {
		return false
	}
	return true
}

func updateCapacityToNode(vgList, qpList, ldList []*StorageCapacity) error {
	nodeInfo, err := types.GlobalConfigVar.KubeClient.CoreV1().Nodes().Get(context.Background(), types.GlobalConfigVar.NodeID, metav1.GetOptions{})
	if err != nil {
		log.Errorf("updateCapacityToNode:: get node info with error : %s", err.Error())
		return err
	}
	qpList = append(qpList, vgList...)
	qpList = append(qpList, ldList...)
	qpList = append(qpList, LocalDeviceList...)
	capacity, err := json.Marshal(qpList)
	if err != nil {
		log.Errorf("Update volumecapacity with json.Marshal error: %s", err.Error())
		return err
	}
	nodeInfo.Annotations["csi.alibabacloud.com/storage-topology"] = string(capacity)
	_, err = types.GlobalConfigVar.KubeClient.CoreV1().Nodes().Update(context.Background(), nodeInfo, metav1.UpdateOptions{})
	if err != nil {
		log.Errorf("Update volumecapacity to node with error: %s", err.Error())
		return err
	}
	log.Infof("Successful Update volumecapacity to node: %s", string(capacity))
	return nil
}

func getCapacityFromNode() []*StorageCapacity {
	capacityList := []*StorageCapacity{}
	nodeInfo, err := types.GlobalConfigVar.KubeClient.CoreV1().Nodes().Get(context.Background(), types.GlobalConfigVar.NodeID, metav1.GetOptions{})
	if err != nil {
		log.Errorf("getCapacityFromNode:: get node info with error : %s", err.Error())
		return nil
	}
	if value, ok := nodeInfo.Annotations["csi.alibabacloud.com/storage-topology"]; ok {
		err := json.Unmarshal([]byte(value), &capacityList)
		if err != nil {
			log.Errorf("getCapacityFromNode:: get node info with json.Unmarshal error : %s", err.Error())
			return nil
		}
	}
	for _, item := range capacityList {
		log.Infof("Successful Get storage capacity from node: %++v", item)
	}
	return capacityList
}
