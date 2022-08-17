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
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io/ioutil"
	k8smount "k8s.io/utils/mount"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

type nodeServer struct {
	k8smounter k8smount.Interface
	*csicommon.DefaultNodeServer
	writeCredentialMutex sync.Mutex
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
	AuthType      string `json:"authType"`
	FuseType      string `json:"fuseType"`
	MetricsTop    string `json:"metricsTop"`
}

const (
	// OssfsCredentialFile is the path of oss ak credential file
	OssfsCredentialFile = "/host/etc/passwd-ossfs"
	// NsenterCmd is nsenter mount command
	NsenterCmd = "/nsenter --mount=/proc/1/ns/mnt"
	// AkID is Ak ID
	AkID = "akId"
	// AkSecret is Ak Secret
	AkSecret = "akSecret"
	// OssFsType is the oss filesystem type
	OssFsType = "ossfs"
	// JindoFsType tag
	JindoFsType = "jindofs"
	// metricsPathPrefix
	metricsPathPrefix = "/host/var/run/ossfs/"
	// metricsTop
	metricsTop = "10"
)

var (
	// SharedPath is the shared mountpoint when UseSharedPath is "true"
	SharedPath = filepath.Join(utils.KubeletRootDir, "/plugins/kubernetes.io/csi/pv/%s/globalmount")
)

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// logout oss paras
	log.Infof("NodePublishVolume:: Starting Mount volume: %s mount with req: %+v", req.VolumeId, req)
	mountPath := req.GetTargetPath()
	opt := &Options{}
	opt.UseSharedPath = false
	opt.FuseType = OssFsType
	opt.MetricsTop = "10"
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
		} else if key == "authtype" {
			opt.AuthType = strings.ToLower(strings.TrimSpace(value))
		} else if key == "fusetype" {
			opt.FuseType = strings.ToLower(strings.TrimSpace(value))
		} else if key == "metricstop"{
			opt.MetricsTop = strings.ToLower(strings.TrimSpace(value))
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

	argStr := fmt.Sprintf("Bucket: %s, url: %s, , OtherOpts: %s, Path: %s, UseSharedPath: %s, authType: %s", opt.Bucket, opt.URL, opt.OtherOpts, opt.Path, strconv.FormatBool(opt.UseSharedPath), opt.AuthType)
	log.Infof("NodePublishVolume:: Starting Oss Mount: %s", argStr)

	if IsOssfsMounted(mountPath) {
		log.Infof("NodePublishVolume: The mountpoint is mounted: %s", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	// If you do not use sts authentication, save ak
	credentialProvider := ""
	if opt.AuthType != "sts" {
		if opt.FuseType == OssFsType {
			if err := ns.saveOssCredential(opt); err != nil {
				return nil, errors.New("Save ossfs ak is failed, err: " + err.Error())
			}
		} else if opt.FuseType == JindoFsType {
			credentialProvider = fmt.Sprintf("-ofs.oss.accessKeyId=%s -ofs.oss.accessKeySecret=%s", opt.AkID, opt.AkSecret)
		}
	} else {
		if opt.FuseType == OssFsType {
			credentialProvider = GetRAMRoleOption()
		} else if opt.FuseType == JindoFsType {
			credentialProvider = "-ofs.oss.provider.endpoint=ECS_ROLE"
		}
	}

	// default use allow_other
	var mntCmd string
	if opt.UseSharedPath {
		sharedPath := fmt.Sprintf(SharedPath, req.GetVolumeId())
		if IsOssfsMounted(sharedPath) {
			log.Infof("NodePublishVolume: The shared path: %s is already mounted", sharedPath)
		} else {
			if err := utils.CreateDest(sharedPath); err != nil {
				log.Errorf("Ossfs mount is failed, err: %v", err.Error())
				return nil, errors.New("Create OSS volume fail: " + err.Error())
			}
			mntCmd = fmt.Sprintf("systemd-run --scope -- /usr/local/bin/ossfs %s:%s %s -ourl=%s %s %s", opt.Bucket, opt.Path, sharedPath, opt.URL, opt.OtherOpts, credentialProvider)
			if opt.FuseType == JindoFsType {
				mntCmd = fmt.Sprintf("systemd-run --scope -- /etc/jindofs-tool/jindo-fuse %s -ouri=oss://%s%s -ofs.oss.endpoint=%s %s", sharedPath, opt.Bucket, opt.Path, opt.URL, credentialProvider)
			}
			if err := utils.DoMountInHost(mntCmd); err != nil {
				return nil, err
			}
		}
		log.Infof("NodePublishVolume:: Start mount operation from source [%s] to dest [%s]", sharedPath, mountPath)
		options := []string{"bind"}
		if err := ns.k8smounter.Mount(sharedPath, mountPath, "", options); err != nil {
			log.Errorf("Ossfs mount error: %v", err.Error())
			return nil, errors.New("Create oss volume fail: " + err.Error())
		}
	} else {
		// Create Mount Path
		if err := utils.CreateDest(mountPath); err != nil {
			log.Errorf("Create directory is failed, err: %s", err.Error())
			return nil, errors.New("Mount is failed, with create path err: " + err.Error() + mountPath)
		}
		mntCmd = fmt.Sprintf("systemd-run --scope -- /usr/local/bin/ossfs %s:%s %s -ourl=%s %s %s", opt.Bucket, opt.Path, mountPath, opt.URL, opt.OtherOpts, credentialProvider)
		if opt.FuseType == JindoFsType {
			mntCmd = fmt.Sprintf("systemd-run --scope -- /etc/jindofs-tool/jindo-fuse %s -ouri=oss://%s%s -ofs.oss.endpoint=%s %s", mountPath, opt.Bucket, opt.Path, opt.URL, credentialProvider)
		}
		WriteMetricsInfo(metricsPathPrefix, req, *opt)
		if err := utils.DoMountInHost(mntCmd); err != nil {
			return nil, err
		}
	}

	log.Infof("NodePublishVolume:: Mount oss is successfully, volume %s, targetPath: %s, with Command: %s", req.VolumeId, mountPath, mntCmd)
	return &csi.NodePublishVolumeResponse{}, nil
}

// save ak file: bucket:ak_id:ak_secret
func saveOssfsCredential(options *Options) error {
	oldContentByte := []byte{}
	if utils.IsFileExisting(OssfsCredentialFile) {
		tmpValue, err := ioutil.ReadFile(OssfsCredentialFile)
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
	if err := utils.WriteAndSyncFile(OssfsCredentialFile, []byte(newContentStr), 0640); err != nil {
		log.Errorf("Save ossfs passwd-ossfs credential file is failed, err: %s", err)
		return err
	}
	return nil
}

func uniqOssfsCredential() {
	curOssInfoByte := []byte{}
	if utils.IsFileExisting(OssfsCredentialFile) {
		curOssInfoByte, _ = ioutil.ReadFile(OssfsCredentialFile)
	}
	curOssInfoStr := string(curOssInfoByte[:])
	curOssInfoStrArray := strings.Split(curOssInfoStr, "\n")
	uniqOssInfoStrArray := removeDuplicateElement(curOssInfoStrArray)
	uniqOssInfoStr := ""
	for _, line := range uniqOssInfoStrArray {
		uniqOssInfoStr = uniqOssInfoStr + line + "\n"
	}
	if err := utils.WriteAndSyncFile(OssfsCredentialFile, []byte(uniqOssInfoStr), 0640); err != nil {
		log.Errorf("Uniq credential file is failed, %s, %s", uniqOssInfoStr, err)
		return
	}
}

func removeDuplicateElement(languages []string) []string {
	result := make([]string, 0, len(languages))
	temp := map[string]struct{}{}
	for _, item := range languages {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// Check oss options
func checkOssOptions(opt *Options) error {
	if opt.URL == "" || opt.Bucket == "" {
		return errors.New("Oss Parametes error: Url/Bucket empty ")
	}

	if !strings.HasPrefix(opt.Path, "/") {
		return errors.New("Oss path error: start with " + opt.Path + ", should start with / ")
	}

	if opt.FuseType == JindoFsType {
		return nil
	}

	// if not input ak from user, use the default ak value
	if opt.AkID == "" || opt.AkSecret == "" {
		ac := utils.GetEnvAK()
		opt.AkID = ac.AccessKeyID
		opt.AkSecret = ac.AccessKeySecret
	}
	if opt.AkID == "" || opt.AkSecret == "" {
		if opt.AuthType == "" {
			return errors.New("Oss Parametes error: AK and authType are both empty ")
		}
	}

	if opt.OtherOpts != "" {
		if !strings.HasPrefix(opt.OtherOpts, "-o ") {
			return errors.New("Oss OtherOpts error: start with -o ")
		}
	}

	return nil
}

func (ns *nodeServer) saveOssCredential(opt *Options) error {
	// Save ak file for ossfs, exist same entry
	ns.writeCredentialMutex.Lock()
	defer ns.writeCredentialMutex.Unlock()
	if err := saveOssfsCredential(opt); err != nil {
		log.Errorf("Save ossfs ak is failed, err: %s", err.Error())
		return errors.New("Save ossfs ak is failed, err: " + err.Error())
	}
	//The same entry will exist concurrently, will to uniq same entry.
	uniqOssfsCredential()
	return nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Infof("NodeUnpublishVolume:: Starting Umount OSS: %s mount with req: %+v", req.TargetPath, req)
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
