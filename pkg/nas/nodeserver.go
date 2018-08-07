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

package nas

import (
	"context"
	"errors"
	"fmt"
	"path"
	"strings"
	"sync"
	"time"

	"github.com/AliyunContainerService/csi-plugin/pkg/utils"
	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi/v0"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
}

type NasOptions struct {
	Server string `json:"host"`
	Path   string `json:"path"`
	Vers   string `json:"vers"`
	Mode   string `json:"mode"`
}

const (
	NAS_TEMP_MNTPath = "/mnt/acs_mnt/k8s_nas/temp" // used for create sub directory;
)

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {

	log.Infof("Nas Plugin Mount: %s", req.VolumeAttributes)

	// parse parameters
	mountPath := req.GetTargetPath()
	opt := &NasOptions{}
	for key, value := range req.VolumeAttributes {
		if key == "host" {
			opt.Server = value
		} else if key == "path" {
			opt.Path = value
		} else if key == "vers" {
			opt.Vers = value
		} else if key == "mode" {
			opt.Mode = value
		}
	}

	// check parameters
	if mountPath == "" {
		return nil, errors.New("mountPath is empty")
	}
	if opt.Server == "" {
		return nil, errors.New("host is empty, should input nas domain")
	}
	if opt.Path == "" {
		opt.Path = "/"
	}
	if !strings.HasPrefix(opt.Path, "/") {
		return nil, errors.New("the path format is illegal")
	}
	if opt.Vers == "" {
		opt.Vers = "4.0"
	}

	if utils.IsMounted(mountPath) {
		log.Infof("Nas, Mount Path Already Mount, options: %s", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// Create Mount Path
	if err := utils.CreateDest(mountPath); err != nil {
		return nil, errors.New("Nas, Mount error with create Path fail: " + mountPath)
	}

	// Do mount
	mntCmd := fmt.Sprintf("mount -t nfs -o vers=%s %s:%s %s", opt.Vers, opt.Server, opt.Path, mountPath)
	_, err := utils.Run(mntCmd)

	// Mount to nfs Sub-directory
	if err != nil && strings.Contains(err.Error(), "reason given by server: No such file or directory") && opt.Path != "/" {
		ns.createNasSubDir(opt)
		if _, err := utils.Run(mntCmd); err != nil {
			return nil, errors.New("Nas, Mount Nfs sub directory fail: " + err.Error())
		}
		// mount error
	} else if err != nil {
		return nil, errors.New("Nas, Mount nfs fail: " + err.Error())
	}

	// change the mode
	if opt.Mode != "" && opt.Path != "/" {
		var wg1 sync.WaitGroup
		wg1.Add(1)

		go func(*sync.WaitGroup) {
			cmd := fmt.Sprintf("chmod -R %s %s", opt.Mode, mountPath)
			if _, err := utils.Run(cmd); err != nil {
				log.Errorf("Nas chmod cmd fail: %s %s", cmd, err)
			} else {
				log.Infof("Nas chmod cmd success: %s", cmd)
			}
			wg1.Done()
		}(&wg1)

		if waitTimeout(&wg1, 1) {
			log.Infof("Chmod use more than 1s, running in Concurrency: %s", mountPath)
		}
	}

	// check mount
	if !utils.IsMounted(mountPath) {
		return nil, errors.New("Check mount fail after mount:" + mountPath)
	}
	log.Info("Mount success on: " + mountPath)

	return &csi.NodePublishVolumeResponse{}, nil
}

func waitTimeout(wg *sync.WaitGroup, timeout int) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false
	case <-time.After(time.Duration(timeout) * time.Second):
		return true
	}

}

func (ns *nodeServer) createNasSubDir(opt *NasOptions) {
	// step 1: create mount path
	if err := utils.CreateDest(NAS_TEMP_MNTPath); err != nil {
		log.Infof("Create Nas temp Directory err: " + err.Error())
		return
	}
	if utils.IsMounted(NAS_TEMP_MNTPath) {
		utils.Umount(NAS_TEMP_MNTPath)
	}

	// step 2: do mount
	mntCmd := fmt.Sprintf("mount -t nfs -o vers=%s %s:%s %s", opt.Vers, opt.Server, "/", NAS_TEMP_MNTPath)
	_, err := utils.Run(mntCmd)
	if err != nil {
		log.Infof("Nas, Mount to temp directory fail: " + err.Error())
		return
	}
	subPath := path.Join(NAS_TEMP_MNTPath, opt.Path)
	if err := utils.CreateDest(subPath); err != nil {
		log.Infof("Nas, Create Sub Directory err: " + err.Error())
		return
	}

	// step 3: umount after create
	utils.Umount(NAS_TEMP_MNTPath)
	log.Info("Create Sub Directory success: ", opt.Path)
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {

	mountPoint := req.TargetPath
	if !utils.IsMounted(mountPoint) {
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	umntCmd := fmt.Sprintf("umount -f %s", mountPoint)
	if _, err := utils.Run(umntCmd); err != nil {
		return nil, errors.New("Nas, Umount nfs Fail: " + err.Error())
	}

	log.Info("Umount nfs Successful: %s", mountPoint)
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
