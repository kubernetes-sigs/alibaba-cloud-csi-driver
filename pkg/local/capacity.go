package local

import (
	"context"
	"encoding/json"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// StorageCapacity define storage pool details
// csi.alibabacloud.com/storage-capacity:
// [
//   {"type": "lvm", "name": "vg1", "capacity": 105151127552},
//   {"type": "quotapath", "name": "quotapath1", "capacity": 105151127552, "devicetype": "ssd", mountPoint": "/path1"}
// ]
// volumegorup 类型：name、capacity都是必选；
// quotapath 类型：mountPoint、capacity是必选；
type StorageCapacity struct {
	Type       string `json:"type,omitempty"`
	Name       string `json:"name,omitempty"`
	Capacity   uint64 `json:"capacity,omitempty"`
	DeviceType string `json:"devicetype,omitempty"`
	MountPoint string `json:"mountPoint,omitempty"`
}

func updateNodeCapacity() {
	log.Infof("updateNodeCapacity: Starting to update volume capacity to node")
	CacheStorageCapacity := getCapacityFromNode()
	for {
		vgList := getVolumeGroup()
		qpList := getQuotaPath()

		if !isCapacitySame(vgList, qpList, CacheStorageCapacity) {
			for _, item := range vgList {
				log.Infof("updateNodeCapacity: volumeGroup capacity %++v", item)
			}
			for _, item := range qpList {
				log.Infof("updateNodeCapacity: QuotaPath capacity %++v", item)
			}
			for _, item := range CacheStorageCapacity {
				log.Infof("updateNodeCapacity: Cached capacity %++v", item)
			}

			err := updateCapacityToNode(vgList, qpList)
			if err == nil {
				CacheStorageCapacity = vgList
				for _, item := range qpList {
					CacheStorageCapacity = append(CacheStorageCapacity, item)
				}
			}
		}

		// sleep with reconcile interval
		time.Sleep(time.Duration(60) * time.Second)
	}
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

// get quotapath from local node
// not support now
func getQuotaPath() []*StorageCapacity {
	return nil
}

// check the current volume capacity is same as cached.
func isCapacitySame(vgList, qpList, cachedList []*StorageCapacity) bool {
	if len(vgList)+len(qpList) != len(cachedList) {
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

func updateCapacityToNode(vgList, qpList []*StorageCapacity) error {
	nodeInfo, err := types.GlobalConfigVar.KubeClient.CoreV1().Nodes().Get(context.Background(), types.GlobalConfigVar.NodeID, metav1.GetOptions{})
	if err != nil {
		log.Errorf("updateCapacityToNode:: get node info with error : %s", err.Error())
		return err
	}
	for _, item := range vgList {
		qpList = append(qpList, item)
	}
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
		log.Infof("Successful Get volumecapacity from node: %++v", item)
	}
	return capacityList
}
