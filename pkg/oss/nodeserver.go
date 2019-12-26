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

package oss

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"net"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
}

// Options contains options for target oss
type Options struct {
	Bucket           string `json:"bucket"`
	URL              string `json:"url"`
	OtherOpts        string `json:"otherOpts"`
	AkID             string `json:"akId"`
	AkSecret         string `json:"akSecret"`
	Path             string `json:"path"`
	Accelerate       string `json:"accelerate"`
	AccelerateVolume string `json:"accelerateVolume"`
}

const (
	//CredentialFile is the path of oss ak credential file
	CredentialFile = "/host/etc/passwd-ossfs"
	// NsenterCmd is nsenter mount command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt"
	// SocketPath is path of connector sock
	SocketPath = "/host/etc/csi-tool/connector.sock"
	// AkID is Ak ID
	AkID = "akId"
	// AkSecret is Ak Secret
	AkSecret = "akSecret"
	// AccelerateTypeLvm tag
	AccelerateTypeLvm = "lvm"
	// AccelerateTypeMemory tag
	AccelerateTypeMemory = "memory"
)

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// logout oss paras
	mountPath := req.GetTargetPath()
	opt := &Options{}
	for key, value := range req.VolumeContext {
		key = strings.ToLower(key)
		if key == "bucket" {
			opt.Bucket = strings.TrimSpace(value)
		} else if key == "url" {
			opt.URL = strings.TrimSpace(value)
		} else if key == "otheropts" {
			opt.OtherOpts = strings.TrimSpace(value)
		} else if key == "akid" {
			opt.AkID = strings.TrimSpace(value)
		} else if key == "aksecret" {
			opt.AkSecret = strings.TrimSpace(value)
		} else if key == "path" {
			opt.Path = strings.TrimSpace(value)
		} else if key == "accelerate" {
			opt.Accelerate = strings.TrimSpace(value)
		} else if key == "acceleratevolume" {
			opt.AccelerateVolume = strings.TrimSpace(value)
		}
	}
	if opt.Path == "" {
		opt.Path = "/"
	}

	// support set ak by secret
	if opt.AkID == "" || opt.AkSecret == "" {
		if value, ok := req.Secrets[AkID]; ok {
			opt.AkID = value
		}
		if value, ok := req.Secrets[AkSecret]; ok {
			opt.AkSecret = value
		}
	}

	// check parameters
	if err := checkOssOptions(opt); err != nil {
		log.Errorf("Check oss input error: %s", err.Error())
		return nil, errors.New("Check oss input error: " + err.Error())
	}
	if mountPath == "" {
		log.Errorf("Check oss input error: mountPath is empty")
		return nil, errors.New("mountPath is empty")
	}

	useCachePath := ""
	if opt.Accelerate == AccelerateTypeLvm {
		if opt.AccelerateVolume == "" {
			log.Errorf("Check Oss Accelerate input error: lvmVolume should not be empty")
			return nil, errors.New("Check Oss Accelerate input error: lvmVolume should not be empty ")
		}
		pathCmd := fmt.Sprintf("mount | grep %s | grep -v grep | awk 'NR==1{print $3}'", opt.AccelerateVolume)
		if out, err := utils.Run(pathCmd); err != nil {
			log.Errorf("Check lvm volume error: %s", err.Error())
			return nil, errors.New("Check lvm volume error: " + err.Error())
		} else {
			mntPath := strings.TrimSpace(out)
			if mntPath == "" {
				log.Warnf("lvm volume mntPath empty: %s, %s", opt.AccelerateVolume, mntPath)
				return nil, errors.New("lvm volume mntPath empty: " + opt.AccelerateVolume + mntPath)
			}
			useCachePath = filepath.Join(mntPath, "cache")
			if err := utils.CreateDest(useCachePath); err != nil {
				log.Errorf("lvm volume CreateDest error: %s", err.Error())
				return nil, errors.New("lvm volume CreateDest error: " + err.Error())
			}
		}
	} else if opt.Accelerate == AccelerateTypeMemory {

	}

	argStr := "Bucket: " + opt.Bucket + ", url: " + opt.URL + ", OtherOpts: " + opt.OtherOpts + ", Path: " + opt.Path
	log.Infof("NodePublishVolume:: Starting Oss Mount: %s", argStr)

	if utils.IsMounted(mountPath) {
		log.Infof("NodePublishVolume: The mountpoint is mounted: %s", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// Create Mount Path
	if err := utils.CreateDest(mountPath); err != nil {
		log.Errorf("Create Directory error: %s", err.Error())
		return nil, errors.New("Oss, Mount fail with create Path error: " + err.Error() + mountPath)
	}

	// Save ak file for ossfs
	if err := saveOssCredential(opt); err != nil {
		log.Errorf("Save oss ak error: %s", err.Error())
		return nil, errors.New("Oss, Save AK file fail: " + err.Error())
	}

	// default use allow_other
	mntCmd := fmt.Sprintf("systemd-run --scope -- /usr/local/bin/ossfs %s:%s %s -ourl=%s %s", opt.Bucket, opt.Path, mountPath, opt.URL, opt.OtherOpts)
	if opt.AccelerateVolume != "" {
		cacheOption := "-o use_cache=" + useCachePath
		log.Infof("OSS Volume: %s Set Accelerate with volume: %s", req.VolumeId, opt.AccelerateVolume)
		mntCmd = fmt.Sprintf("systemd-run --scope -- /usr/local/bin/ossfs %s:%s %s -ourl=%s %s %s", opt.Bucket, opt.Path, mountPath, opt.URL, opt.OtherOpts, cacheOption)
	}

	if out, err := connectorRun(mntCmd); err != nil {
		if err != nil {
			log.Errorf("Ossfs mount error: %s", err.Error())
			return nil, errors.New("Create OSS volume fail: " + err.Error() + ", out: " + out)
		}
	}

	log.Infof("NodePublishVolume:: Mount Oss successful: %s, with Command: %s", mountPath, mntCmd)
	return &csi.NodePublishVolumeResponse{}, nil
}

// save ak file: bucket:ak_id:ak_secret
func saveOssCredential(options *Options) error {

	oldContentByte := []byte{}
	if utils.IsFileExisting(CredentialFile) {
		tmpValue, err := ioutil.ReadFile(CredentialFile)
		if err != nil {
			return err
		}
		oldContentByte = tmpValue
	}

	oldContentStr := string(oldContentByte[:])
	newContentStr := ""
	for _, line := range strings.Split(oldContentStr, "\n") {
		lineList := strings.Split(line, ":")
		if len(lineList) != 3 || lineList[0] == options.Bucket {
			continue
		}
		newContentStr += line + "\n"
	}

	newContentStr = options.Bucket + ":" + options.AkID + ":" + options.AkSecret + "\n" + newContentStr
	if err := ioutil.WriteFile(CredentialFile, []byte(newContentStr), 0640); err != nil {
		log.Errorf("Save Credential File failed, %s, %s", newContentStr, err)
		return err
	}
	return nil
}

// Check oss options
func checkOssOptions(opt *Options) error {
	if opt.URL == "" || opt.Bucket == "" {
		return errors.New("Oss Parametes error: Url/Bucket empty ")
	}

	// if not input ak from user, use the default ak value
	if opt.AkID == "" || opt.AkSecret == "" {
		opt.AkID, opt.AkSecret = utils.GetLocalAK()
	}
	if opt.AkID == "" || opt.AkSecret == "" {
		return errors.New("Oss Parametes error: AK is empty ")
	}

	if opt.OtherOpts != "" {
		if !strings.HasPrefix(opt.OtherOpts, "-o ") {
			return errors.New("Oss OtherOpts error: start with -o ")
		}
	}

	if !strings.HasPrefix(opt.Path, "/") {
		return errors.New("Oss path error: start with " + opt.Path + ", should start with / ")
	}

	if opt.Accelerate != AccelerateTypeLvm && opt.Accelerate != AccelerateTypeMemory {
		return errors.New("Oss accelerate type error, only support memory/lvm ")
	}

	return nil
}

// Run shell command with host connector
// host connector is daemon running in host.
func connectorRun(cmd string) (string, error) {
	c, err := net.Dial("unix", SocketPath)
	if err != nil {
		log.Errorf("Oss connector Dial error: %s", err.Error())
		return err.Error(), err
	}
	defer c.Close()

	_, err = c.Write([]byte(cmd))
	if err != nil {
		log.Errorf("Oss connector write error: %s", err.Error())
		return err.Error(), err
	}

	buf := make([]byte, 2048)
	n, err := c.Read(buf[:])
	response := string(buf[0:n])
	if strings.HasPrefix(response, "Success") {
		respstr := response[8:]
		return respstr, nil
	}
	return response, errors.New("oss connector exec command error:" + response)
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

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {

	log.Infof("NodeUnpublishVolume:: Starting Umount OSS: %s", req.TargetPath)
	mountPoint := req.TargetPath
	if !utils.IsMounted(mountPoint) {
		log.Infof("Directory is not mounted: %s", mountPoint)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	umntCmd := fmt.Sprintf("umount -f %s", mountPoint)
	if _, err := utils.Run(umntCmd); err != nil {
		log.Errorf("Umount oss fail, with: %s", err.Error())
		return nil, errors.New("Oss, Umount oss Fail: " + err.Error())
	}

	log.Infof("NodeUnpublishVolume:: Umount OSS Successful: %s", mountPoint)
	return &csi.NodeUnpublishVolumeResponse{}, nil
}

func (ns *nodeServer) NodeStageVolume(
	ctx context.Context,
	req *csi.NodeStageVolumeRequest) (
	*csi.NodeStageVolumeResponse, error) {
	return &csi.NodeStageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeUnstageVolume(
	ctx context.Context,
	req *csi.NodeUnstageVolumeRequest) (
	*csi.NodeUnstageVolumeResponse, error) {
	return &csi.NodeUnstageVolumeResponse{}, nil
}

// IsHostMounted return status of host mounted or not
func IsHostMounted(mountPath string) bool {
	cmd := fmt.Sprintf("%s mount | grep \"%s type\" | grep -v grep", NsenterCmd, mountPath)
	out, err := utils.Run(cmd)
	if err != nil || out == "" {
		return false
	}
	return true
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
