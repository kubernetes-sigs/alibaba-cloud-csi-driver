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

package jindofs

import (
	"context"
	"errors"
	"fmt"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	k8smount "k8s.io/utils/mount"
	"net"
	"strings"
)

type nodeServer struct {
	k8smounter k8smount.Interface
	*csicommon.DefaultNodeServer
}

// Options contains options for target oss
type Options struct {
	Bucket    string `json:"bucket"`
	URL       string `json:"url"`
	OtherOpts string `json:"otherOpts"`
	AkID      string `json:"akId"`
	AkSecret  string `json:"akSecret"`
	Path      string `json:"path"`
	AuthType  string `json:"authType"`
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
	// SharedPath is the shared mountpoint when UseSharedPath is "true"
	SharedPath = "/var/lib/kubelet/plugins/kubernetes.io/csi/pv/%s/globalmount"
	// JindoFsType is the oss filesystem type
	JindoFsType = "fuse.jindofs-fuse"
)

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// logout oss paras
	log.Infof("NodePublishVolume:: Starting Mount Jindofs volume: %s to path: %s", req.VolumeId, req.TargetPath)
	mountPath := req.GetTargetPath()

	opt := &Options{}
	for key, value := range req.VolumeContext {
		key = strings.ToLower(key)
		switch key {
		case "bucket":
			opt.Bucket = strings.TrimSpace(value)
		case "url":
			opt.URL = strings.TrimSpace(value)
		case "akid":
			opt.AkID = strings.TrimSpace(value)
		case "aksecret":
			opt.AkSecret = strings.TrimSpace(value)
		case "authtype":
			opt.AuthType = strings.ToLower(strings.TrimSpace(value))
		case "path":
			if v := strings.TrimSpace(value); v == "" {
				opt.Path = "/"
			} else {
				opt.Path = v
			}
		}
	}

	// Default oss path
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

	if mountPath == "" {
		log.Errorf("Check oss input error: mountPath is empty for volume %s ", req.VolumeId)
		return nil, errors.New("mountPath is empty")
	}
	if !utils.IsFileExisting(mountPath) {
		utils.CreateDest(mountPath)
	}

	argStr := fmt.Sprintf("Bucket: %s, url: %s, , OtherOpts: %s, Path: %s, autType: %s", opt.Bucket, opt.URL, opt.OtherOpts, opt.Path, opt.AuthType)
	log.Infof("NodePublishVolume:: Starting Oss Mount: %s", argStr)

	if IsJindofsMounted(mountPath) {
		log.Infof("NodePublishVolume: The mountpoint is mounted: %s", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// If you do not use sts authentication, save ak
	credentialProvider := ""
	if opt.AuthType == "sts" {
		credentialProvider = "-ocredential_provider=secrets:///etc/jindofs-credentials"
	} else if opt.AuthType == "ecs-role" {
		credentialProvider = "-ocredential_provider=ECS_ROLE"
	} else {
		// Save ak file for jindofs
		if err := saveOssCredential(opt); err != nil {
			log.Errorf("Save oss ak error: %s", err.Error())
			return nil, errors.New("Oss, Save AK file fail: " + err.Error())
		}
	}

	jindofsCmd := fmt.Sprintf("systemd-run --scope -- /etc/csi-tool/jindofs-fuse -obucket=%v -opath=%v -oendpoint=%v %v -oonly_sdk %s", opt.Bucket, opt.Path, opt.URL, credentialProvider, mountPath)
	out, err := connectorRun(jindofsCmd)
	if err != nil {
		log.Errorf("NodePublishVolume: Volume %s Jindofs mount error: %s, %v", req.VolumeId, err.Error(), out)
		return nil, errors.New("NodePublishVolume: Jindofs mount error: " + err.Error())
	}

	log.Infof("NodePublishVolume:: Mount JindoFS successful, volume %s, targetPath: %s, with Command: %s, with out: %s", req.VolumeId, mountPath, jindofsCmd, out)
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

	newContentStr = options.Bucket + "." + options.URL + ":" + options.AkID + ":" + options.AkSecret + "\n" + newContentStr
	if err := ioutil.WriteFile(CredentialFile, []byte(newContentStr), 0640); err != nil {
		log.Errorf("Save Credential File failed, %s, %s", newContentStr, err)
		return err
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

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Infof("NodeUnpublishVolume:: Starting Umount OSS: %s", req.TargetPath)
	mountPoint := req.TargetPath
	if !IsJindofsMounted(mountPoint) {
		log.Infof("Directory is not mounted: %s", mountPoint)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	pvName := req.GetVolumeId()
	var umntCmd string
	sharedMountPoint := fmt.Sprintf(SharedPath, pvName)
	if IsJindofsMounted(sharedMountPoint) {
		log.Infof("NodeUnpublishVolume:: Starting umount a shared path oss volume: %s", req.TargetPath)
		code, err := IsLastSharedVol(pvName)
		if err != nil {
			log.Errorf("Umount oss fail, with: %s", err.Error())
			return nil, errors.New("Oss, Umount oss Fail: " + err.Error())
		}
		if code == "1" {
			umntCmd = fmt.Sprintf("umount %s && umount -f %s", mountPoint, sharedMountPoint)
		} else {
			umntCmd = fmt.Sprintf("umount %s", mountPoint)
		}
	} else {
		umntCmd = fmt.Sprintf("umount -f %s", mountPoint)
	}
	log.Infof("Umount: %v", umntCmd)
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

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
