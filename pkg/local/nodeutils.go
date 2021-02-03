package local

import (
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"

	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	k8smount "k8s.io/kubernetes/pkg/util/mount"
	"os"
	"path/filepath"
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
		err := ns.createVolume(ctx, volumeID, vgName, pvType, lvmType)
		if err != nil {
			log.Errorf("NodePublishVolume: create volume %s with error: %s", volumeID, err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	}

	// Check targetPath
	if _, err := os.Stat(targetPath); os.IsNotExist(err) {
		if err := os.MkdirAll(targetPath, 0750); err != nil {
			log.Errorf("NodePublishVolume: volume %s mkdir target path %s with error: %s", volumeID, targetPath, err.Error())
			return status.Error(codes.Internal, err.Error())
		}
	}

	isMnt := utils.IsMounted(targetPath)
	if !isMnt {
		var options []string
		if req.GetReadonly() {
			options = append(options, "ro")
		} else {
			options = append(options, "rw")
		}
		mountFlags := req.GetVolumeCapability().GetMount().GetMountFlags()
		options = append(options, mountFlags...)

		diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: k8smount.NewOsExec()}
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
	diskMounter := &k8smount.SafeFormatAndMount{Interface: ns.k8smounter, Exec: k8smount.NewOsExec()}
	if err := diskMounter.FormatAndMount(sourceDevice, targetPath, fsType, options); err != nil {
		log.Errorf("mountDeviceVolume: Volume: %s, Device: %s, FormatAndMount error: %s", req.VolumeId, sourceDevice, err.Error())
		return status.Error(codes.Internal, err.Error())
	}

	return nil
}

// create lvm volume
func (ns *nodeServer) createVolume(ctx context.Context, volumeID, vgName, pvType, lvmType string) error {
	pvSize, unit, _ := ns.getPvInfo(volumeID)
	if pvSize == 0 {
		log.Errorf("createVolume: Volume: %s, VG: %s, parse pv Size zero", volumeID, vgName)
		return status.Error(codes.Internal, "parse pv Size zero")
	}
	pvNumber := 0
	var err error
	// Create VG if vg not exist,
	if pvType == LocalDisk {
		if pvNumber, err = createVG(vgName); err != nil {
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
	if lvmType == StripingType {
		pvNumber = getPVNumber(vgName)
		if pvNumber == 0 {
			log.Errorf("createVolume:: VG is exist: %s, bug get pv number as 0", vgName)
			return err
		}
		cmd := fmt.Sprintf("%s lvcreate -i %d -n %s -L %d%s %s", NsenterCmd, pvNumber, volumeID, pvSize, unit, vgName)
		_, err = utils.Run(cmd)
		if err != nil {
			log.Errorf("createVolume:: lvcreate command %s error: %v", cmd, err)
			return err
		}
		log.Infof("Successful Create Striping LVM volume: %s, with command: %s", volumeID, cmd)
	} else if lvmType == LinearType {
		cmd := fmt.Sprintf("%s lvcreate -n %s -L %d%s %s", NsenterCmd, volumeID, pvSize, unit, vgName)
		_, err = utils.Run(cmd)
		if err != nil {
			log.Errorf("createVolume:: lvcreate linear command %s error: %v", cmd, err)
			return err
		}
		log.Infof("Successful Create Linear LVM volume: %s, with command: %s", volumeID, cmd)
	}
	return nil
}
