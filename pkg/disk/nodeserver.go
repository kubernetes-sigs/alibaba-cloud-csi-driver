/*
Copyright 2018 The Kubernetes Authors.

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

package disk

import (
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi/v0"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/kubernetes/pkg/util/mount"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
}

// csi disk driver: attach and mount
func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// check target mount path
	targetPath := req.GetTargetPath()
	log.Infof("NodePublishVolume: Starting with %s", targetPath)
	if !strings.HasSuffix(targetPath, "/mount") {
		return nil, status.Errorf(codes.InvalidArgument, "malformed the value of target path: %s", targetPath)
	}

	// get source path
	devicePath, err := DefaultAttachEntry.Get(req.GetVolumeId())
	if err != nil {
		log.Errorf("NodePublishVolume: failed to get device path%v", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	log.Infof("NodePublishVolume: Get source device: %s", devicePath)

	// check target mountpath is mounted
	notMnt, err := mount.New(devicePath).IsLikelyNotMountPoint(targetPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err = os.MkdirAll(targetPath, 0750); err != nil {
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMnt = true
		} else {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !notMnt {
		return &csi.NodePublishVolumeResponse{}, nil
	}

	fsType := req.GetVolumeCapability().GetMount().GetFsType()
	readOnly := req.GetReadonly()
	attrib := req.GetVolumeAttributes()
	mountFlags := req.GetVolumeCapability().GetMount().GetMountFlags()
	log.Infof("NodePublishVolume: Starting Format and Mount: target %v | fstype %v | device %v | readonly %v | attributes %v | mountflags %v",
		targetPath, fsType, devicePath, readOnly, attrib, mountFlags)

	// Start to format and Mount
	option := []string{}
	if readOnly {
		option = append(option, "ro")
	}
	diskMounter := &mount.SafeFormatAndMount{Interface: mount.New(""), Exec: mount.NewOsExec()}
	if err := diskMounter.FormatAndMount(devicePath, targetPath, fsType, option); err != nil {
		log.Errorf("NodePublishVolume: FormatAndMount error: %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	log.Infof("NodePublishVolume: Format and Mount Successful: target %v", targetPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	targetPath := req.GetTargetPath()
	log.Infof("NodeUnpublishVolume: Start to unpublish volume, target %v", targetPath)
	// Step 1: check mount point
	mounter := mount.New("")
	notMnt, err := mounter.IsLikelyNotMountPoint(targetPath)
	if err != nil {
		log.Errorf("NodeUnpublishVolume: ", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if notMnt {
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	// Step 2: umount target path
	cnt := GetDeviceMountNum(targetPath)
	err = mounter.Unmount(targetPath)
	if err != nil {
		log.Errorf("NodeUnpublishVolume: umount error:", err.Error())
		return nil, status.Error(codes.Internal, err.Error())
	}
	if cnt > 1 {
		log.Warnf("Only Unmount, with device mount by others: refer num ", cnt)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}
	log.Infof("NodeUnpublishVolume: Unpublish successful, target %v", targetPath)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(
	ctx context.Context,
	req *csi.NodeStageVolumeRequest) (
	*csi.NodeStageVolumeResponse, error) {

	return nil, status.Error(codes.Unimplemented, "")
}

func (ns *nodeServer) NodeUnstageVolume(
	ctx context.Context,
	req *csi.NodeUnstageVolumeRequest) (
	*csi.NodeUnstageVolumeResponse, error) {

	return nil, status.Error(codes.Unimplemented, "")
}
