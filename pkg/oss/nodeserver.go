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
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	csicommon "github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/client-go/dynamic"
	mountutils "k8s.io/mount-utils"
)

type nodeServer struct {
	*csicommon.DefaultNodeServer
	writeCredentialMutex sync.Mutex
	dynamicClient        dynamic.Interface
	sharedPathLock       *utils.VolumeLocks
	ossfsMounterFac      *mounter.ContainerizedFuseMounterFactory
}

// Options contains options for target oss
type Options struct {
	directAssigned bool

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
	ReadOnly      bool   `json:"readOnly"`
}

const (
	// OssfsCredentialFile is the path of oss ak credential file
	OssfsCredentialFile = "/host/etc/passwd-ossfs"
	// NsenterCmd is nsenter mount command
	NsenterCmd = "nsenter --mount=/proc/1/ns/mnt"
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
	// regionTag
	regionTag = "region-id"
)

func validateNodePublishVolumeRequest(req *csi.NodePublishVolumeRequest) error {
	valid, err := utils.CheckRequest(req.GetVolumeContext(), req.GetTargetPath())
	if !valid {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	return nil
}

func (ns *nodeServer) NodePublishVolume(ctx context.Context, req *csi.NodePublishVolumeRequest) (*csi.NodePublishVolumeResponse, error) {
	// logout oss paras
	log.Infof("NodePublishVolume:: Starting Mount volume: %s mount with req: %+v", req.VolumeId, req)
	mountPath := req.GetTargetPath()
	if err := validateNodePublishVolumeRequest(req); err != nil {
		return nil, err
	}
	var cnfsName string
	opt := &Options{}
	opt.UseSharedPath = true
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
		} else if key == "metricstop" {
			opt.MetricsTop = strings.ToLower(strings.TrimSpace(value))
		} else if key == "containernetworkfilesystem" {
			cnfsName = value
		} else if key == optDirectAssigned {
			opt.directAssigned, _ = strconv.ParseBool(strings.TrimSpace(value))
		}
	}

	if req.Readonly {
		opt.ReadOnly = true
	} else {
		switch req.VolumeCapability.AccessMode.GetMode() {
		case csi.VolumeCapability_AccessMode_SINGLE_NODE_WRITER, csi.VolumeCapability_AccessMode_MULTI_NODE_SINGLE_WRITER, csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER:
			opt.ReadOnly = false
		default:
			opt.ReadOnly = true
		}
	}

	if len(opt.Bucket) == 0 {
		cnfs, err := v1beta1.GetCnfsObject(ns.dynamicClient, cnfsName)
		if err != nil {
			return nil, err
		}
		if cnfs.Status.FsAttributes.EndPoint == nil {
			return nil, errors.New("Cnfs " + cnfsName + " is not ready, endpoint is empty.")
		}
		opt.Bucket = cnfs.Status.FsAttributes.BucketName
		opt.URL = cnfs.Status.FsAttributes.EndPoint.Internal
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

	argStr := fmt.Sprintf("Bucket: %s, url: %s, , OtherOpts: %s, Path: %s, UseSharedPath: %s, authType: %s", opt.Bucket, opt.URL, opt.OtherOpts, opt.Path, strconv.FormatBool(opt.UseSharedPath), opt.AuthType)
	log.Infof("NodePublishVolume:: Starting Oss Mount: %s", argStr)

	if opt.directAssigned {
		return ns.publishDirectVolume(ctx, req, opt)
	}
	mounter := ns.getMounterFor(opt.FuseType, req.VolumeId)
	notMnt, err := mountutils.IsNotMountPoint(mounter, mountPath)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(mountPath, os.ModePerm); err != nil {
				log.Errorf("NodePublishVolume: mkdir %s: %v", mountPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
			notMnt = true
		} else {
			log.Errorf("NodePublishVolume: check mountpoint %s: %v", mountPath, err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	if !notMnt {
		log.Infof("NodePublishVolume: %s already mounted", mountPath)
		return &csi.NodePublishVolumeResponse{}, nil
	}

	var mountOptions []string
	if req.VolumeCapability != nil && req.VolumeCapability.GetMount() != nil {
		mountOptions = req.VolumeCapability.GetMount().MountFlags
	}

	// If you do not use sts authentication, save ak
	if opt.AuthType != "sts" {
		if opt.FuseType == OssFsType {
			if err := ns.saveOssCredential(opt); err != nil {
				return nil, errors.New("Save ossfs ak is failed, err: " + err.Error())
			}
		} else if opt.FuseType == JindoFsType {
			mountOptions = append(mountOptions, fmt.Sprintf("fs.oss.accessKeyId=%s,fs.oss.accessKeySecret=%s", opt.AkID, opt.AkSecret))
		}
	} else {
		if opt.FuseType == OssFsType {
			mountOptions = append(mountOptions, GetRAMRoleOption())
		} else if opt.FuseType == JindoFsType {
			mountOptions = append(mountOptions, "fs.oss.provider.endpoint=ECS_ROLE")
		}
	}

	if opt.ReadOnly {
		mountOptions = append(mountOptions, "ro")
	}

	metaZoneID := GetMetaDataAsync(regionTag)
	if strings.Contains(opt.URL, metaZoneID) && !strings.Contains(opt.URL, "internal") && !utils.IsPrivateCloud() {
		originUrl := opt.URL
		opt.URL = strings.ReplaceAll(originUrl, metaZoneID, metaZoneID+"-internal")
	}
	if opt.UseSharedPath {
		// TODO: how to collect sharedPath metrics !!!
		sharedPath := GetGlobalMountPath(req.GetVolumeId())
		notMnt, err := mounter.IsLikelyNotMountPoint(sharedPath)
		if err != nil {
			if os.IsNotExist(err) {
				if err := os.MkdirAll(sharedPath, os.ModePerm); err != nil {
					log.Errorf("NodePublishVolume: mkdir %s: %v", sharedPath, err)
					return nil, status.Error(codes.Internal, err.Error())
				}
				notMnt = true
			} else {
				log.Errorf("NodePublishVolume: check mountpoint %s: %v", sharedPath, err)
				return nil, status.Error(codes.Internal, err.Error())
			}
		}
		if notMnt {
			// serialize node publish operations on the same volume when using sharedpath
			if lock := ns.sharedPathLock.TryAcquire(req.VolumeId); !lock {
				log.Errorf("NodePublishVolume: aborted because failed to acquire volume %s lock", req.VolumeId)
				return nil, status.Errorf(codes.Aborted, "NodePublishVolume operation on shared path of volume %s already exists", req.VolumeId)
			}
			defer ns.sharedPathLock.Release(req.VolumeId)
			if err := doMount(mounter, sharedPath, *opt, mountOptions); err != nil {
				log.Errorf("NodePublishVolume: failed to mount")
				return nil, err
			}
		} else {
			log.Infof("NodePublishVolume: %s already mounted", sharedPath)
			return &csi.NodePublishVolumeResponse{}, nil
		}
		log.Infof("NodePublishVolume:: Start mount operation from source [%s] to dest [%s]", sharedPath, mountPath)
		if err := mounter.Mount(sharedPath, mountPath, "", []string{"bind"}); err != nil {
			log.Errorf("Ossfs mount error: %v", err.Error())
			return nil, errors.New("Create oss volume fail: " + err.Error())
		}
	} else {
		if opt.FuseType == OssFsType {
			utils.WriteMetricsInfo(metricsPathPrefix, req, opt.MetricsTop, OssFsType, "oss", opt.Bucket)
		}
		if err := doMount(mounter, mountPath, *opt, mountOptions); err != nil {
			return nil, err
		}
	}

	log.Infof("NodePublishVolume: mounted oss successfully, volume %s, targetPath: %s", req.VolumeId, mountPath)
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

func validateNodeUnpublishVolumeRequest(req *csi.NodeUnpublishVolumeRequest) error {
	valid, err := utils.ValidatePath(req.GetTargetPath())
	if !valid {
		return status.Errorf(codes.InvalidArgument, err.Error())
	}
	return nil
}

func (ns *nodeServer) NodeUnpublishVolume(ctx context.Context, req *csi.NodeUnpublishVolumeRequest) (*csi.NodeUnpublishVolumeResponse, error) {
	log.Infof("NodeUnpublishVolume:: Starting Umount OSS: %s mount with req: %+v", req.TargetPath, req)
	mountPoint := req.TargetPath
	err := validateNodeUnpublishVolumeRequest(req)
	if err != nil {
		return nil, err
	}
	if isDirectVolumePath(mountPoint) {
		return ns.unPublishDirectVolume(ctx, req)
	}
	mounter := ns.getMounterFor("", req.VolumeId)
	forceUnmounter, ok := mounter.(mountutils.MounterForceUnmounter)
	if ok {
		err = mountutils.CleanupMountWithForce(mountPoint, forceUnmounter, true, time.Minute)
	} else {
		err = mountutils.CleanupMountPoint(mountPoint, mounter, true)
	}
	if err != nil {
		log.Errorf("NodeUnpublishVolume: failed to unmount %q: %v", mountPoint, err)
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", mountPoint, err)
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
	log.Infof("NodeUnstageVolume: starting to unmount volume, volumeId: %s, target: %v", req.VolumeId, req.StagingTargetPath)
	mounter := ns.getMounterFor("", req.VolumeId)
	mountPoint := GetGlobalMountPath(req.VolumeId)
	forceUnmounter, ok := mounter.(mountutils.MounterForceUnmounter)
	var err error
	if ok {
		err = mountutils.CleanupMountWithForce(req.StagingTargetPath, forceUnmounter, true, time.Minute)
	} else {
		err = mountutils.CleanupMountPoint(mountPoint, mounter, true)
	}
	if err != nil {
		log.Errorf("NodeUnpublishVolume: failed to unmount %q: %v", mountPoint, err)
		return nil, status.Errorf(codes.Internal, "failed to unmount target %q: %v", mountPoint, err)
	}
	log.Infof("NodeUnpublishVolume:: Umount OSS Successful: %s", mountPoint)
	return &csi.NodeUnstageVolumeResponse{}, nil
}

func (ns *nodeServer) NodeExpandVolume(ctx context.Context, req *csi.NodeExpandVolumeRequest) (
	*csi.NodeExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (ns *nodeServer) getMounterFor(fuseType, volumeId string) mountutils.Interface {
	if fuseType == JindoFsType {
		return mounter.NewConnectorMounter(mountutils.New(""), "/etc/jindofs-tool/jindo-fuse")
	}
	return ns.ossfsMounterFac.NewFuseMounter(volumeId)
}
