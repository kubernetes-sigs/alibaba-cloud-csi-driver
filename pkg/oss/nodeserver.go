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
	"strings"
	"sync"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	k8smount "k8s.io/kubernetes/pkg/util/mount"
	"strconv"
)

type nodeServer struct {
	k8smounter k8smount.Interface
	*csicommon.DefaultNodeServer
}

// Options contains options for target oss
type Options struct {
	Bucket        string `json:"bucket"`
	URL           string `json:"url"`
	OtherOpts     string `json:"otherOpts"`
	AkID          string `json:"akId"`
	AkSecret      string `json:"akSecret"`
	Path          string `json:"path"`
	UseSharedPath bool   `json:"useSharedPath"`
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
	// OssFsType is the oss filesystem type
	OssFsType = "fuse.ossfs"
)

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// logout oss paras
	mountPath := req.GetTargetPath()
	opt := &Options{}
	opt.UseSharedPath = false
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
			if v := strings.TrimSpace(value); v == "" {
				opt.Path = "/"
			} else {
				opt.Path = v
			}
		} else if key == "usesharedpath" {
			if strings.TrimSpace(value) == "true" || strings.TrimSpace(value) == "True" || strings.TrimSpace(value) == "1" {
				opt.UseSharedPath = true
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

	// check parameters
	if err := checkOssOptions(opt); err != nil {
		log.Errorf("Check oss input error: %s", err.Error())
		return nil, errors.New("Check oss input error: " + err.Error())
	}
	if mountPath == "" {
		log.Errorf("Check oss input error: mountPath is empty")
		return nil, errors.New("mountPath is empty")
	}

	argStr := fmt.Sprintf("Bucket: %s, url: %s, , OtherOpts: %s, Path: %s, UseSharedPath: %s", opt.Bucket, opt.URL, opt.OtherOpts, opt.Path, strconv.FormatBool(opt.UseSharedPath))
	log.Infof("NodePublishVolume:: Starting Oss Mount: %s", argStr)

	if IsOssfsMounted(mountPath) {
		log.Infof("NodePublishVolume: The mountpoint is mounted: %s", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// Save ak file for ossfs
	if err := saveOssCredential(opt); err != nil {
		log.Errorf("Save oss ak error: %s", err.Error())
		return nil, errors.New("Oss, Save AK file fail: " + err.Error())
	}

	// default use allow_other
	var mntCmd string
	if opt.UseSharedPath {
		sharedPath := fmt.Sprintf(SharedPath, req.GetVolumeId())
		if IsOssfsMounted(sharedPath) {
			log.Infof("NodePublishVolume: The shared path: %s is already mounted", sharedPath)
		} else {
			if err := utils.CreateDest(sharedPath); err != nil {
				log.Errorf("Ossfs mount error: %v", err.Error())
				return nil, errors.New("Create OSS volume fail: " + err.Error())
			}
			mntCmd = fmt.Sprintf("systemd-run --scope -- /usr/local/bin/ossfs %s:%s %s -ourl=%s %s", opt.Bucket, opt.Path, sharedPath, opt.URL, opt.OtherOpts)
			if out, err := connectorRun(mntCmd); err != nil {
				if err != nil {
					log.Errorf("Ossfs mount error: %s", err.Error())
					return nil, errors.New("Create OSS volume fail: " + err.Error() + ", out: " + out)
				}
			}
		}
		log.Infof("NodePublishVolume:: Start mount operation from source [%s] to dest [%s]", sharedPath, mountPath)
		options := []string{"bind"}
		if err := ns.k8smounter.Mount(sharedPath, mountPath, "", options); err != nil {
			log.Errorf("Ossfs mount error: %v", err.Error())
			return nil, errors.New("Create OSS volume fail: " + err.Error())
		}
	} else {
		// Create Mount Path
		if err := utils.CreateDest(mountPath); err != nil {
			log.Errorf("Create Directory error: %s", err.Error())
			return nil, errors.New("Oss, Mount fail with create Path error: " + err.Error() + mountPath)
		}
		mntCmd = fmt.Sprintf("systemd-run --scope -- /usr/local/bin/ossfs %s:%s %s -ourl=%s %s", opt.Bucket, opt.Path, mountPath, opt.URL, opt.OtherOpts)
		if out, err := connectorRun(mntCmd); err != nil {
			if err != nil {
				log.Errorf("Ossfs mount error: %s", err.Error())
				return nil, errors.New("Create OSS volume fail: " + err.Error() + ", out: " + out)
			}
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
	if !IsOssfsMounted(mountPoint) {
		log.Infof("Directory is not mounted: %s", mountPoint)
		return &csi.NodeUnpublishVolumeResponse{}, nil
	}

	pvName := req.GetVolumeId()
	var umntCmd string
	sharedMountPoint := fmt.Sprintf(SharedPath, pvName)
	if IsOssfsMounted(sharedMountPoint) {
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
