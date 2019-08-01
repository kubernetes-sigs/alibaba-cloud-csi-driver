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

package oss

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"io/ioutil"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/AliyunContainerService/csi-plugin/pkg/utils"
	log "github.com/Sirupsen/logrus"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
}

type OssOptions struct {
	Bucket    string `json:"bucket"`
	Url       string `json:"url"`
	OtherOpts string `json:"otherOpts"`
	AkId      string `json:"akId"`
	AkSecret  string `json:"akSecret"`
}

const (
	CredentialFile = "/host/etc/passwd-ossfs"
	NSENTER_CMD    = "/nsenter --mount=/proc/1/ns/mnt"
	SOCKET_PATH    = "/host/etc/csi-tool/connector.sock"
	AK_ID          = "akId"
	AK_SECRET      = "akSecret"
)

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// logout oss paras
	mountPath := req.GetTargetPath()
	opt := &OssOptions{}
	for key, value := range req.VolumeContext {
		if key == "bucket" {
			opt.Bucket = value
		} else if key == "url" {
			opt.Url = value
		} else if key == "otherOpts" {
			opt.OtherOpts = value
		} else if key == "akId" {
			opt.AkId = value
		} else if key == "akSecret" {
			opt.AkSecret = value
		}
	}

	// support set ak by secret
	if opt.AkId == "" || opt.AkSecret == "" {
		if value, ok := req.Secrets[AK_ID]; ok {
			opt.AkId = value
		}
		if value, ok := req.Secrets[AK_SECRET]; ok {
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

	argStr := "AkId: " + opt.AkId + ", Bucket: " + opt.Bucket + ", url: " + opt.Url + ", OtherOpts: " + opt.OtherOpts
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
	mntCmd := fmt.Sprintf("systemd-run --scope -- /usr/local/bin/ossfs %s %s -ourl=%s -o allow_other %s", opt.Bucket, mountPath, opt.Url, opt.OtherOpts)
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
func saveOssCredential(options *OssOptions) error {

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

	newContentStr = options.Bucket + ":" + options.AkId + ":" + options.AkSecret + "\n" + newContentStr
	if err := ioutil.WriteFile(CredentialFile, []byte(newContentStr), 0640); err != nil {
		log.Errorf("Save Credential File failed, %s, %s", newContentStr, err)
		return err
	}
	return nil
}

// Check oss options
func checkOssOptions(opt *OssOptions) error {
	if opt.Url == "" || opt.Bucket == "" {
		return errors.New("Oss Parametes error: Url/Bucket empty ")
	}

	// if not input ak from user, use the default ak value
	if opt.AkId == "" || opt.AkSecret == "" {
		opt.AkId, opt.AkSecret = utils.GetLocalAK()
	}
	if opt.AkId == "" || opt.AkSecret == "" {
		return errors.New("Oss Parametes error: AK is empty ")
	}

	if opt.OtherOpts != "" {
		if !strings.HasPrefix(opt.OtherOpts, "-o ") {
			return errors.New("Oss OtherOpts error: start with -o ")
		}
	}
	return nil
}

// Run shell command with host connector
// host connector is daemon running in host.
func connectorRun(cmd string) (string, error) {
	c, err := net.Dial("unix", SOCKET_PATH)
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

func IsHostMounted(mountPath string) bool {
	cmd := fmt.Sprintf("%s mount | grep \"%s type\" | grep -v grep", NSENTER_CMD, mountPath)
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