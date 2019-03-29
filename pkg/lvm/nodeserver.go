package lvm

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/AliyunContainerService/csi-plugin/pkg/utils"
	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	NSENTER_CMD = "/nsenter --mount=/proc/1/ns/mnt"
	VG_NAME_TAG = "vgName"
	LVM_SIZE    = "lvmSize"
	DEFAULTFS   = "ext4"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
	nodeID  string
	mounter utils.Mounter
}

func NewNodeServer(d *csicommon.CSIDriver, nodeID string) csi.NodeServer {
	return &nodeServer{
		DefaultNodeServer: csicommon.NewDefaultNodeServer(d),
		nodeID:            nodeID,
		mounter:           utils.NewMounter(),
	}
}

func (ns *nodeServer) GetNodeID() string {
	return ns.nodeID
}

func (ns *nodeServer) createVolume(ctx context.Context, volumeId string, vgName string, size string) (error) {
	ckCmd := fmt.Sprintf("%s vgck %s", NSENTER_CMD, vgName)
	_, err := utils.Run(ckCmd)
	if err != nil {
		return err
	}
	cmd := fmt.Sprintf("%s lvcreate -n %s -L %s %s", NSENTER_CMD, volumeId, size, vgName)
	_, err = utils.Run(cmd)
	if err != nil {
		return err
	}
	return nil
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	if targetPath == "" {
		return nil, status.Error(codes.Internal, "targetPath is empty")
	}

	vgName := ""
	if _, ok := req.VolumeContext[VG_NAME_TAG]; ok {
		vgName = req.VolumeContext[VG_NAME_TAG]
	}
	if vgName == "" {
		return nil, status.Error(codes.Internal, "error with input vgName is empty")
	}

	size := ""
	if _, ok := req.VolumeContext[LVM_SIZE]; ok {
		size = req.VolumeContext[LVM_SIZE]
	}
	if size == "" {
		return nil, status.Error(codes.Internal, "error with input lvmSize is empty")
	}

	volumeId := req.GetVolumeId()
	devicePath := filepath.Join("/dev/", vgName, volumeId)
	if _, err := os.Stat(devicePath); os.IsNotExist(err) {
		err := ns.createVolume(ctx, volumeId, vgName, size)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	isMnt, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			if err := os.MkdirAll(targetPath, 0750); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			isMnt = false
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	log.Printf("Check Filesystem type at %v", devicePath)
	exitFSType, err := checkFSType(devicePath)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "check fs type err: %v", err)
	}
	log.Printf("Device: %s fstype is '%v'", devicePath, exitFSType)
	if exitFSType == "" {
		log.Printf("The device %v has no filesystem, format %v", devicePath, exitFSType)
		if err := formatDevice(devicePath, DEFAULTFS); err != nil {
			return nil, status.Errorf(codes.Internal, "format fstype failed: err=%v", err)
		}
		exitFSType = DEFAULTFS
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

		err = ns.mounter.Mount(devicePath, targetPath, DEFAULTFS, options...)
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Infof("NodePublishVolume:: mount successful devicePath: %s, targetPath: %s, options: %v", devicePath, targetPath, options)
	}

	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	isMnt, err := ns.mounter.IsMounted(targetPath)
	if err != nil {
		if _, err := os.Stat(targetPath); os.IsNotExist(err) {
			return nil, status.Error(codes.NotFound, "TargetPath not found")
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !isMnt {
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	err = ns.mounter.Unmount(req.GetTargetPath())
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(ctx context.Context, req *csi.NodeUnstageVolumeRequest) (*csi.NodeUnstageVolumeResponse, error) {
	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(ctx context.Context, req *csi.NodeStageVolumeRequest) (*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, nil
}
