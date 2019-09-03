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

package cpfs

import (
	"context"
	"errors"
	"fmt"
	"github.com/AliyunContainerService/csi-plugin/pkg/utils"
	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
}

type CpfsOptions struct {
	Server     string `json:"server"`
	FileSystem string `json:"fileSystem"`
	SubPath    string `json:"subPath"`
	Options    string `json:"options"`
}

const (
	CPFS_TEMP_MNTPath = "/mnt/acs_mnt/k8s_cpfs/temp" // used for create sub directory;
)

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {

	log.Infof("NodePublishVolume:: CPFS Mount with: %s", req.VolumeContext)

	// parse parameters
	mountPath := req.GetTargetPath()
	opt := &CpfsOptions{}
	for key, value := range req.VolumeContext {
		if key == "server" {
			opt.Server = value
		} else if key == "fileSystem" {
			opt.FileSystem = value
		} else if key == "subPath" {
			opt.SubPath = value
		} else if key == "options" {
			opt.Options = value
		}
	}

	// check parameters
	if mountPath == "" {
		return nil, errors.New("mountPath is empty")
	}
	if opt.Server == "" {
		return nil, errors.New("host is empty, should input Cpfs domain")
	}
	if opt.SubPath == "" {
		opt.SubPath = "/"
	}
	if !strings.HasPrefix(opt.SubPath, "/") {
		return nil, errors.New("the path format is illegal")
	}

	if utils.IsMounted(mountPath) {
		log.Infof("CPFS, Mount Path Already Mount, options: %s", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// if system not set cpfs, config it.
	//checkSystemCpfsConfig()

	// Create Mount Path
	if err := utils.CreateDest(mountPath); err != nil {
		return nil, errors.New("Cpfs, Mount error with create Path fail: " + mountPath)
	}

	// Do mount
	mntCmd := fmt.Sprintf("mount -t lustre %s:/%s%s %s", opt.Server, opt.FileSystem, opt.SubPath, mountPath)
	if opt.Options != "" {
		mntCmd = fmt.Sprintf("mount -t lustre -o %s %s:/%s%s %s", opt.Options, opt.Server, opt.FileSystem, opt.SubPath, mountPath)
	}
	_, err := utils.Run(mntCmd)

	// Mount to cpfs Sub-directory
	if err != nil && opt.SubPath != "/" {
		if strings.Contains(err.Error(), "No such file or directory") || strings.Contains(err.Error(), "access denied by server while mounting") {
			createCpfsSubDir(opt.Options, opt.Server, opt.FileSystem, opt.SubPath, req.VolumeId)
			if _, err := utils.Run(mntCmd); err != nil {
				log.Errorf("Cpfs, Mount Cpfs sub directory fail: %s", err.Error())
				return nil, errors.New("Cpfs, Mount Cpfs sub directory fail: %s" + err.Error())
			}
		} else {
			log.Errorf("Cpfs, Mount Cpfs fail with error: %s", err.Error())
			return nil, errors.New("Cpfs, Mount Cpfs fail with error: %s" + err.Error())
		}
		// mount error
	} else if err != nil {
		log.Errorf("Cpfs, Mount Cpfs fail: %s", err.Error())
		return nil, errors.New("Cpfs, Mount Cpfs fail: %s" + err.Error())
	}
	log.Infof("NodePublishVolume:: Exec mount command: %s", mntCmd)

	// check mount
	if !utils.IsMounted(mountPath) {
		return nil, errors.New("Check mount fail after mount:" + mountPath)
	}
	log.Infof("NodePublishVolume:: Mount success on mountpoint: %s", mountPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Infof("NodeUnpublishVolume:: Starting Umount Cpfs: %s", req.TargetPath)
	mountPoint := req.TargetPath
	if !utils.IsMounted(mountPoint) {
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	umntCmd := fmt.Sprintf("umount %s", mountPoint)
	if _, err := utils.Run(umntCmd); err != nil {
		return nil, errors.New("Cpfs, Umount cpfs Fail: " + err.Error())
	}

	log.Infof("Umount cpfs Successful on: %s", mountPoint)
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

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
