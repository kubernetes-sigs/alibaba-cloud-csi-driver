package local

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/strategicpatch"
	"k8s.io/kubernetes/pkg/util/resizefs"
	utilexec "k8s.io/utils/exec"
	k8smount "k8s.io/utils/mount"
	"os"
	"path/filepath"
	"strings"
)

func (ns *nodeServer) mountLvm(ctx context.Context, req *csi.NodePublishVolumeRequest) error {
	targetPath := req.TargetPath
	// parse vgname, consider invalid if empty
	vgName := ""
	if _, ok := req.VolumeContext[VgNameTag]; ok {
		vgName = req.VolumeContext[VgNameTag]
	}
	if vgName == "" {
		log.Errorf("NodePublishVolume: request with empty vgName in volume: %s", req.VolumeId)
		return status.Error(codes.Internal, "error with input vgName is empty")
	}

	// special process on alibaba local disk type;
	pvType := CloudDisk
	if _, ok := req.VolumeContext[PvTypeTag]; ok {
		pvType = req.VolumeContext[PvTypeTag]
	}
	// parse lvm type and fstype
	lvmType := LinearType
	if _, ok := req.VolumeContext[LvmTypeTag]; ok {
		lvmType = req.VolumeContext[LvmTypeTag]
	}
	fsType := DefaultFs
	if _, ok := req.VolumeContext[FsTypeTag]; ok {
		fsType = req.VolumeContext[FsTypeTag]
	}
	nodeAffinity := DefaultNodeAffinity
	if _, ok := req.VolumeContext[NodeAffinity]; ok {
		nodeAffinity = req.VolumeContext[NodeAffinity]
	}
	log.Infof("NodePublishVolume: Starting to mount lvm at path: %s, with vg: %s, with volume: %s, PV Type: %s, LVM Type: %s, NodeAffinty: %s", targetPath, vgName, req.GetVolumeId(), pvType, lvmType, nodeAffinity)

	// Create LVM if not exist
	//volumeNewCreated := false
	volumeID := req.GetVolumeId()
	devicePath := filepath.Join("/dev/", vgName, volumeID)
	if _, err := os.Stat(devicePath); os.IsNotExist(err) {
		//volumeNewCreated = true
		err := createVolume(ctx, volumeID, vgName, pvType, lvmType)
		if err != nil {
			log.Errorf("NodePublishVolume: create volume %s with error: %s", volumeID, err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	}

	// Check target mounted
	isMnt, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			if err := os.MkdirAll(targetPath, 0750); err != nil {
				log.Errorf("NodePublishVolume: volume %s mkdir target path %s with error: %s", volumeID, targetPath, err.Error())
				return status.Error(codes.Internal, err.Error())
			}
			isMnt = false
		} else {
			log.Errorf("NodePublishVolume: check volume %s mounted with error: %s", volumeID, err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	}

	if !isMnt {
		var options []string
		if req.GetReadonly() {
			options = append(options, "ro")
		} else {
			options = append(options, "rw")
		}
		mountFlags := req.GetVolumeCapability().GetMount().GetMountFlags()
		options = append(options, mountFlags...)

		diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
		if err := diskMounter.FormatAndMount(devicePath, targetPath, fsType, options); err != nil {
			log.Errorf("NodeStageVolume: Volume: %s, Device: %s, FormatAndMount error: %s", req.VolumeId, devicePath, err.Error())
			return status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodePublishVolume:: mount successful devicePath: %s, targetPath: %s, options: %v", devicePath, targetPath, options)
	}
	return nil
}

func (ns *nodeServer) mountLocalVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) error {
	sourcePath := ""
	targetPath := req.TargetPath
	if value, ok := req.VolumeContext[MountPointType]; ok {
		sourcePath = value
	}
	if sourcePath == "" {
		log.Errorf("mountLocalVolume: volume: %s, sourcePath empty", req.VolumeId)
		return status.Error(codes.Internal, "Mount LocalVolume with empty source path "+req.VolumeId)
	}

	notmounted, err := ns.k8smounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		log.Errorf("mountLocalVolume: check volume: %s mounted with error %v", req.VolumeId, err)
		return status.Error(codes.Internal, err.Error())
	}
	if !notmounted {
		log.Infof("NodePublishVolume: VolumeId: %s, Local Path %s is already mounted", req.VolumeId, targetPath)
		return nil
	}

	// start to mount
	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "bind")
	if req.Readonly {
		options = append(options, "ro")
	}
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}
	log.Infof("NodePublishVolume: Starting mount local volume %s with flags %v and fsType %s", req.VolumeId, options, fsType)
	if err = ns.k8smounter.Mount(sourcePath, targetPath, fsType, options); err != nil {
		log.Errorf("mountLocalVolume: Mount volume: %s with error %v", req.VolumeId, err)
		return status.Error(codes.Internal, err.Error())
	}
	return nil
}

func (ns *nodeServer) mountDeviceVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) error {
	sourceDevice := ""
	targetPath := req.TargetPath
	if value, ok := req.VolumeContext[DeviceVolumeType]; ok {
		sourceDevice = value
	}
	if sourceDevice == "" {
		log.Errorf("mountDeviceVolume: device volume: %s, sourcePath empty", req.VolumeId)
		return status.Error(codes.Internal, "Mount Device with empty source path "+req.VolumeId)
	}

	// Step Start to format
	mnt := req.VolumeCapability.GetMount()
	options := append(mnt.MountFlags, "shared")
	fsType := "ext4"
	if mnt.FsType != "" {
		fsType = mnt.FsType
	}

	// do format-mount or mount
	diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
	if err := diskMounter.FormatAndMount(sourceDevice, targetPath, fsType, options); err != nil {
		log.Errorf("mountDeviceVolume: Volume: %s, Device: %s, FormatAndMount error: %s", req.VolumeId, sourceDevice, err.Error())
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

func (ns *nodeServer) mountPmemVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) error {
	targetPath := req.TargetPath
	// parse vgname, consider invalid if empty
	vgName := ""
	if _, ok := req.VolumeContext[VgNameTag]; ok {
		vgName = req.VolumeContext[VgNameTag]
	}
	if vgName == "" {
		vgName = lib.PmemVolumeGroupNameRegion0
	}
	pmemType := ""
	if _, ok := req.VolumeContext[PmemType]; ok {
		pmemType = req.VolumeContext[PmemType]
	}

	// parse lvm type and fstype
	lvmType := LinearType
	if _, ok := req.VolumeContext[LvmTypeTag]; ok {
		lvmType = req.VolumeContext[LvmTypeTag]
	}
	fsType := DefaultFs
	if _, ok := req.VolumeContext[FsTypeTag]; ok {
		fsType = req.VolumeContext[FsTypeTag]
	}
	nodeAffinity := DefaultNodeAffinity
	if _, ok := req.VolumeContext[NodeAffinity]; ok {
		nodeAffinity = req.VolumeContext[NodeAffinity]
	}
	pmemBlockDev := ""
	if value, ok := req.VolumeContext[PmemBlockDev]; ok {
		pmemBlockDev = value
	}
	log.Infof("NodePublishVolume: Starting to mount lvm at path: %s, with vg: %s, with volume: %s, LVM Type: %s, NodeAffinty: %s", targetPath, vgName, req.GetVolumeId(), lvmType, nodeAffinity)

	// Create LVM if not exist
	//volumeNewCreated := false
	volumeID := req.GetVolumeId()
	devicePath := filepath.Join("/dev/", vgName, volumeID)
	if pmemType != "kmem" && pmemType != "projquota" {
		if pmemBlockDev == "" {
			if _, err := os.Stat(devicePath); os.IsNotExist(err) {
				//volumeNewCreated = true
				err := createVolume(ctx, volumeID, vgName, "", lvmType)
				if err != nil {
					log.Errorf("NodePublishVolume: create volume %s with error: %s", volumeID, err.Error())
					return status.Error(codes.Internal, err.Error())
				}
			}
		} else {
			devicePath = filepath.Join("/dev/", pmemBlockDev)
			if _, err := os.Stat(devicePath); os.IsNotExist(err) {
				return status.Error(codes.Internal, err.Error())
			}
			ns.checkPmemNameSpaceResize(volumeID, targetPath)
		}
	}
	// Check target mounted
	isMnt, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			if err := os.MkdirAll(targetPath, 0750); err != nil {
				log.Errorf("NodePublishVolume: volume %s mkdir target path %s with error: %s", volumeID, targetPath, err.Error())
				return status.Error(codes.Internal, err.Error())
			}
			isMnt = false
		} else {
			log.Errorf("NodePublishVolume: check volume %s mounted with error: %s", volumeID, err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	}

	if !isMnt {
		switch pmemType {
		case "kmem":
			pvSize, pvSizeUnit, _ := getPvInfo(volumeID)
			devicePath = "tmpfs"
			kmemSize := fmt.Sprintf("%v%s", pvSize, pvSizeUnit)
			mpol, err := pickRegionForKMEM(kmemSize)
			if err != nil {
				log.Errorf("NodeStageVolume: Volume: %s, Device: %s, pick region for KMEM err: %s", req.VolumeId, devicePath, err.Error())
				return status.Error(codes.Internal, err.Error())
			}
			mountCmd := fmt.Sprintf("%s mount -t %s -o size=%s,mpol=%s %s %s", NsenterCmd, devicePath, kmemSize, mpol, devicePath, targetPath)
			_, err = utils.Run(mountCmd)
			if err != nil {
				log.Errorf("NodeStageVolume: Volume: %s, Device: %s, FormatAndMount error: %s", req.VolumeId, devicePath, err.Error())
				return status.Error(codes.Internal, err.Error())
			}
		case "projquota":
			projQuotaPath := ""
			if value, ok := req.VolumeContext[ProjQuotaFullPath]; ok {
				projQuotaPath = value
			}
			mountCmd := fmt.Sprintf("%s mount --bind %s %s", NsenterCmd, projQuotaPath, targetPath)
			_, err = utils.Run(mountCmd)
			if err != nil {
				log.Errorf("NodeStageVolume: Volume: %s, Device: %s, mount error: %s", req.VolumeId, devicePath, err.Error())
				return status.Error(codes.Internal, err.Error())
			}
		default:
			var options []string
			if req.GetReadonly() {
				options = append(options, "ro")
			} else {
				options = append(options, "rw")
			}
			mountFlags := req.GetVolumeCapability().GetMount().GetMountFlags()
			options = append(options, mountFlags...)

			diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()}
			if err := diskMounter.FormatAndMount(devicePath, targetPath, fsType, options); err != nil {
				log.Errorf("NodeStageVolume: Volume: %s, Device: %s, FormatAndMount error: %s", req.VolumeId, devicePath, err.Error())
				return status.Error(codes.Internal, err.Error())
			}
			log.Infof("NodePublishVolume:: mount successful devicePath: %s, targetPath: %s, options: %v", devicePath, targetPath, options)
		}
	}

	// upgrade PV with NodeAffinity
	if nodeAffinity == "true" {
		oldPv, err := ns.client.CoreV1().PersistentVolumes().Get(context.Background(), volumeID, metav1.GetOptions{})
		if err != nil {
			log.Errorf("NodePublishVolume: Get Persistent Volume(%s) Error: %s", volumeID, err.Error())
			return status.Error(codes.Internal, err.Error())
		}
		if oldPv.Spec.NodeAffinity == nil {
			oldData, err := json.Marshal(oldPv)
			if err != nil {
				log.Errorf("NodePublishVolume: Marshal Persistent Volume(%s) Error: %s", volumeID, err.Error())
				return status.Error(codes.Internal, err.Error())
			}
			pvClone := oldPv.DeepCopy()

			// construct new persistent volume data
			values := []string{ns.nodeID}
			nSR := v1.NodeSelectorRequirement{Key: "kubernetes.io/hostname", Operator: v1.NodeSelectorOpIn, Values: values}
			matchExpress := []v1.NodeSelectorRequirement{nSR}
			nodeSelectorTerm := v1.NodeSelectorTerm{MatchExpressions: matchExpress}
			nodeSelectorTerms := []v1.NodeSelectorTerm{nodeSelectorTerm}
			required := v1.NodeSelector{NodeSelectorTerms: nodeSelectorTerms}
			pvClone.Spec.NodeAffinity = &v1.VolumeNodeAffinity{Required: &required}
			newData, err := json.Marshal(pvClone)
			if err != nil {
				log.Errorf("NodePublishVolume: Marshal New Persistent Volume(%s) Error: %s", volumeID, err.Error())
				return status.Error(codes.Internal, err.Error())
			}
			patchBytes, err := strategicpatch.CreateTwoWayMergePatch(oldData, newData, pvClone)
			if err != nil {
				log.Errorf("NodePublishVolume: CreateTwoWayMergePatch Volume(%s) Error: %s", volumeID, err.Error())
				return status.Error(codes.Internal, err.Error())
			}

			// Upgrade PersistentVolume with NodeAffinity
			_, err = ns.client.CoreV1().PersistentVolumes().Patch(context.Background(), volumeID, types.StrategicMergePatchType, patchBytes, metav1.PatchOptions{})
			if err != nil {
				log.Errorf("NodePublishVolume: Patch Volume(%s) Error: %s", volumeID, err.Error())
				return status.Error(codes.Internal, err.Error())
			}
			log.Infof("NodePublishVolume: upgrade Persistent Volume(%s) with nodeAffinity: %s", volumeID, ns.nodeID)
		}
	}
	return nil
}

func (ns *nodeServer) checkPmemNameSpaceResize(volumeID, targetPath string) error {
	pv, err := getPvObj(ns.client, volumeID)
	if err != nil {
		return err
	}

	if pv.Spec.CSI.FSType != "ext4" {
		return fmt.Errorf("pmem direct only support ext4 resize")
	}

	pvQuantity := pv.Spec.Capacity["storage"]
	expectedSize := pvQuantity.Value()
	pmemNameSpace := pv.Spec.CSI.VolumeAttributes["pmemNameSpace"]
	namespace, err := lib.GetNameSpace(pmemNameSpace)
	if err != nil {
		return err
	}
	pvSize := lib.GetNameSpaceCapacity(namespace)
	if expectedSize <= pvSize {
		return nil
	}
	pmemBlockDev := ""
	if value, ok := pv.Spec.CSI.VolumeAttributes["pmemBlockDev"]; ok {
		pmemBlockDev = value
	}

	// ndctl create-namespace -fe namespace0.0 -s 20G
	args := []string{NsenterCmd, "ndctl", "create-namespace", "-fe", pmemNameSpace, "-s", fmt.Sprintf("%d", expectedSize)}
	cmd := strings.Join(args, " ")
	_, err = utils.Run(cmd)
	if err != nil {
		log.Errorf("NodeExpandVolume: Pmem direct %s expand error with %v", pmemNameSpace, err)
		return err
	}

	devicePath := filepath.Join("/dev", pmemBlockDev)
	// use resizer to expand volume filesystem
	resizer := resizefs.NewResizeFs(&k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: utilexec.New()})
	ok, err := resizer.Resize(devicePath, targetPath)
	if err != nil {
		log.Errorf("NodeExpandVolume:: Lvm Resize Error, volumeId: %s, devicePath: %s, volumePath: %s, err: %s", volumeID, devicePath, targetPath, err.Error())
		return err
	}
	if !ok {
		log.Errorf("NodeExpandVolume:: Lvm Resize failed, volumeId: %s, devicePath: %s, volumePath: %s", volumeID, devicePath, targetPath)
		return status.Error(codes.Internal, "Fail to resize volume fs")
	}
	log.Infof("NodeExpandVolume:: lvm resizefs successful volumeId: %s, devicePath: %s, volumePath: %s", volumeID, devicePath, targetPath)
	return nil

}

// create lvm volume
func createVolume(ctx context.Context, volumeID, vgName, pvType, lvmType string) error {
	pvSize, unit, _ := getPvInfo(volumeID)
	if pvSize == 0 {
		log.Errorf("createVolume: Volume: %s, VG: %s, parse pv Size zero", volumeID, vgName)
		return status.Error(codes.Internal, "parse pv Size zero")
	}
	var err error
	// Create VG if vg not exist,
	if pvType == LocalDisk {
		if _, err = createVG(vgName); err != nil {
			log.Errorf("createVolume: Volume: %s, VG: %s, error: %s", volumeID, vgName, err.Error())
			return err
		}
	}

	// check vg exist
	ckCmd := fmt.Sprintf("%s vgck %s", NsenterCmd, vgName)
	_, err = utils.Run(ckCmd)
	if err != nil {
		log.Errorf("createVolume:: VG is not exist: %s", vgName)
		return err
	}

	// Create lvm volume
	if err := createLvm(vgName, volumeID, lvmType, unit, pvSize); err != nil {
		return err
	}

	return nil
}

func createLvm(vgName, volumeID, lvmType, unit string, pvSize int64) error {
	// Create lvm volume
	if lvmType == StripingType {
		pvNumber := getPVNumber(vgName)
		if pvNumber == 0 {
			log.Errorf("createVolume:: VG is exist: %s, bug get pv number as 0", vgName)
			return errors.New("")
		}
		cmd := fmt.Sprintf("%s lvcreate -i %d -n %s -L %d%s %s", NsenterCmd, pvNumber, volumeID, pvSize, unit, vgName)
		_, err := utils.Run(cmd)
		if err != nil {
			log.Errorf("createVolume:: lvcreate command %s error: %v", cmd, err)
			return err
		}
		log.Infof("Successful Create Striping LVM volume: %s, with command: %s", volumeID, cmd)
	} else if lvmType == LinearType {
		cmd := fmt.Sprintf("%s lvcreate -n %s -L %d%s %s", NsenterCmd, volumeID, pvSize, unit, vgName)
		_, err := utils.Run(cmd)
		if err != nil {
			log.Errorf("createVolume:: lvcreate linear command %s error: %v", cmd, err)
			return err
		}
		log.Infof("Successful Create Linear LVM volume: %s, with command: %s", volumeID, cmd)
	}
	return nil
}

func pickRegionForKMEM(size string) (string, error) {
	return "bind:3", nil
}
