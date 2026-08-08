//go:build !windows

package disk

import (
	"context"
	"fmt"
	"net"
	"path/filepath"
	"strconv"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/containerd/ttrpc"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/proto"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	utilsio "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/io"
	perrors "github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2"
	k8smount "k8s.io/mount-utils"
)

func getUnmanagedDiskCount(getNode func() (*v1.Node, error), ecs cloud.ECSv2Interface, instanceID, regionID string, dev utilsio.DiskLister) (int, error) {
	// An attached disk is not managed by us if:
	// 1. it is not in node.Status.VolumesInUse or node.Status.VolumesAttached; and
	// 2. it does not have the xattr set.
	// 1 may fail because the info in node.Status is removed before ControllerUnpublishVolume.
	// 2 may fail because the disk may be just attached and not have the xattr set yet.
	// Combine 1 and 2 to get the accurate "not managed" disk list.
	// We should exclude these disks from available count.
	// e.g. static/dynamic PVs are managed, OS disk or manually attached disks are not managed.

	disks, err := listDiskXattrs(dev)
	if err != nil {
		return 0, fmt.Errorf("failed to list devices: %w", err)
	}
	managedDisks := sets.KeySet(disks)

	// To ensure all the managed attachedDisks also present in managedDisks,
	// ECS OpenAPI should goes after ListDisks because the just detached disk should
	// disappear from ListDisks after OpenAPI;
	// ECS OpenAPI should goes before getNode because the just attached disk should
	// appear in node before OpenAPI;
	attachedDisks, err := GetAttachedCloudDisks(ecs, instanceID, regionID)
	if err != nil {
		return 0, err
	}
	klog.Infof("getVolumeCount: found %d attached disks", len(attachedDisks))

	node, err := getNode()
	if err != nil {
		return 0, err
	}

	for id := range CollectManagedVolumes(node) {
		managedDisks.Insert(id)
	}

	unmanaged := 0
	for _, disk := range attachedDisks {
		if !managedDisks.Has(disk) {
			klog.Infof("getVolumeCount: disk %s is not managed by us", disk)
			unmanaged++
		}
	}

	return unmanaged, nil
}

// checkRundVolumeExpand
func checkRundVolumeExpand(req *csi.NodeExpandVolumeRequest) (bool, error) {
	klog.Infof("checkRundVolumeExpand: volumePath: %s", req.VolumePath)
	pvName := utils.GetPvNameFormPodMnt(req.VolumePath)
	if pvName == "" {
		klog.Errorf("checkRundVolumeExpand: cannot get pvname from volumePath %s", req.VolumePath)
		return false, perrors.Errorf("cannot get pvname from volumePath %s for volume %s", req.VolumePath, req.VolumeId)
	}
	var grpcVolume string
	socketFile := filepath.Join(RundSocketDir, pvName)
	if utils.IsFileExisting(socketFile) {
		grpcVolume = pvName
	} else {
		socketFile = filepath.Join(RundSocketDir, req.VolumeId)
		if !utils.IsFileExisting(socketFile) {
			klog.Infof("checkRundVolumeExpand: socketfile: %s not exists, trying runc expanding", socketFile)
			return false, nil
		}
		grpcVolume = req.VolumeId
	}
	klog.Infof("checkRundVolumeExpand: rund socket dir: %s exists", socketFile)

	// connect to rund server with timeout
	clientConn, err := net.DialTimeout("unix", socketFile, 1*time.Second)
	if err != nil {
		klog.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error: %s", req.VolumeId, req.VolumePath, err.Error())
		return true, perrors.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error: %s", req.VolumeId, req.VolumePath, err.Error())
	}
	defer func() {
		if err := clientConn.Close(); err != nil {
			klog.ErrorS(err, "close rund socket")
		}
	}()

	// send volume spec to rund to expand volume fs
	volumeSize := strconv.FormatInt(req.GetCapacityRange().GetRequiredBytes(), 10)
	client := proto.NewExtendedStatusClient(ttrpc.NewClient(clientConn))
	resp, err := client.ExpandVolume(context.Background(), &proto.ExpandVolumeRequest{
		Volume: grpcVolume,
	})
	if err != nil {
		klog.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error response: %s", req.VolumeId, req.VolumePath, err.Error())
		return true, perrors.Errorf("checkRundExpand: volume %s, volumepath %s, connect to rund server with error response: %s", req.VolumeId, req.VolumePath, err.Error())
	}

	klog.Infof("RundVolumeExpand: Expand VolumeFS(%s) to(%s) successful with response: %s", pvName, volumeSize, resp)
	return true, nil
}

func getMountedVolumeDevice(mnts []k8smount.MountInfo, targetPath string) string {
	for _, mnt := range mnts {
		if mnt.MountPoint == targetPath {
			return mnt.Root
		}
	}
	return ""
}

func isDeviceMountedAt(mnts []k8smount.MountInfo, device, targetPath string) bool {
	for _, mnt := range mnts {
		if mnt.MountPoint == targetPath && mnt.Source == device {
			return true
		}
	}
	return false
}

const mountInfoPath = "/proc/self/mountinfo"

func CheckDeviceAvailable(devicePath, volumeID, targetPath string) error {
	return checkDeviceAvailable(mountInfoPath, devicePath, volumeID, targetPath)
}

func checkDeviceAvailable(mountinfoPath, devicePath, volumeID, targetPath string) error {
	if devicePath == "" {
		return fmt.Errorf("devicePath is empty, cannot used for Volume")
	}

	mnts, err := k8smount.ParseMountInfo(mountinfoPath)
	if err != nil {
		return err
	}

	// TODO: remove in next version
	// Since devicePath has been change to symlink, we will never run into the following logics.
	// block volume
	if devicePath == "devtmpfs" {
		device := getMountedVolumeDevice(mnts, targetPath)
		newVolumeID, err := GetVolumeIDByDevice(device)
		if err != nil {
			return nil
		}
		if newVolumeID != volumeID {
			return fmt.Errorf("device [%s] associate with volumeID: [%s] rather than volumeID: [%s]", device, newVolumeID, volumeID)
		}

		return nil
	}

	if !utils.IsFileExisting(devicePath) {
		return fmt.Errorf("devicePath(%s) is empty, cannot used for Volume", devicePath)
	}

	// check the device is used for system
	if devicePath == "/dev/vda" || devicePath == "/dev/vda1" {
		klog.Warningf("checkDeviceAvailable: devicePath(%s) may be system device: %s", devicePath, volumeID)
	}

	if isDeviceMountedAt(mnts, devicePath, utils.KubeletRootDir) {
		return fmt.Errorf("devicePath(%s) is used as DataDisk for kubelet, cannot used for Volume", devicePath)
	}
	return nil
}
