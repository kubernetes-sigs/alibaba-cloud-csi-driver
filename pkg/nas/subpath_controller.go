/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http:// www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package nas

import (
	"context"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	sdk "github.com/alibabacloud-go/nas-20170626/v3/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	acv1 "k8s.io/client-go/applyconfigurations/core/v1"
)

const (
	finalizerFieldManager    = "nas-subpath-deletion-controller"
	subpathArchiveFinalizer  = "csi.alibabacloud.com/nas-subpath-archive"
	subpathDeletionFinalizer = "csi.alibabacloud.com/nas-subpath-deletion"
)

type subpathController struct {
	config    *internal.ControllerConfig
	nasClient *cloud.NasClientV2
}

func newSubpathController(config *internal.ControllerConfig) (internal.Controller, error) {
	nasClient, err := config.NasClientFactory.V2(config.Region)
	if err != nil {
		return nil, err
	}
	return &subpathController{
		config:    config,
		nasClient: nasClient,
	}, nil
}

func (cs *subpathController) VolumeAs() string {
	return "subpath"
}

func (cs *subpathController) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	// parse parameters
	parameters := req.Parameters
	capacity := req.GetCapacityRange().GetRequiredBytes()
	var (
		path           string
		filesystemId   string
		filesystemType string
	)
	volumeContext := map[string]string{}
	// using cnfs or not
	if cnfsName := parameters["containerNetworkFileSystem"]; cnfsName != "" {
		cnfs, err := cs.config.CNFSGetter.GetCNFS(ctx, cnfsName)
		if err != nil {
			if apierrors.IsNotFound(err) {
				return nil, status.Errorf(codes.InvalidArgument, "CNFS not found: %s", cnfsName)
			}
			return nil, status.Errorf(codes.Internal, "failed to get CNFS %s: %v", cnfsName, err)
		}
		if path = parameters["path"]; path == "" {
			path = "/"
		} else {
			path = filepath.Clean(path)
		}
		filesystemId = cnfs.Status.FsAttributes.FilesystemID
		if filesystemId == "" {
			return nil, status.Error(codes.InvalidArgument, "missing filesystemId in CNFS status")
		}
		filesystemType = cnfs.Status.FsAttributes.FilesystemType
		// set volumeContext
		volumeContext["containerNetworkFileSystem"] = cnfs.Name
	} else {
		var server string
		server, path = muxServerSelector.SelectNfsServer(parameters["server"])
		filesystemId, _, _ := strings.Cut(server, "-")
		if server == "" || filesystemId == "" {
			return nil, status.Error(codes.InvalidArgument, "invalid nas server")
		}
		var err error
		filesystemType, err = cs.getFilesystemType(filesystemId)
		if err != nil {
			return nil, err
		}
		// set volumeContext
		if protocol := parameters["mountProtocol"]; protocol != "" {
			volumeContext["mountProtocol"] = protocol
		}
		volumeContext["server"] = server
	}
	// fill volumeContext
	path = filepath.Join(path, req.Name)
	volumeContext["volumeAs"] = "subpath"
	volumeContext["path"] = path
	if mountType := parameters["mountType"]; mountType != "" {
		if mountType == "losetup" {
			volumeContext["loopImageSize"] = strconv.FormatInt(capacity, 10)
		}
		volumeContext["mountType"] = mountType
	}

	resp := &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      req.Name,
			CapacityBytes: capacity,
			VolumeContext: volumeContext,
		},
	}
	// Only standard filesystems support "CreateDir" and "SetDirQuota" APIs.
	// Subpaths of other types filesystems will be truly created when NodePublishVolume.
	if filesystemType != cloud.FilesystemTypeStandard {
		return resp, nil
	}
	if cs.config.SkipSubpathCreation {
		logrus.Infof("skip creating subpath directory for %s", req.Name)
		return resp, nil
	}
	logrus.WithFields(logrus.Fields{
		"filesystemId": filesystemId,
		"path":         path,
	}).Info("start to create subpath directory for volume")
	// create dir
	if err := cs.nasClient.CreateDir(&sdk.CreateDirRequest{
		FileSystemId:  &filesystemId,
		OwnerGroupId:  tea.Int32(0),
		OwnerUserId:   tea.Int32(0),
		Permission:    tea.String("0777"),
		RootDirectory: &path,
	}); err != nil {
		return nil, status.Errorf(codes.Internal, "nas:CreateDir failed: %v", err)
	}
	// set dir quota
	if parameters["mountType"] != "losetup" && (parameters["volumeCapacity"] == "true" || parameters["allowVolumeExpansion"] == "true") {
		quota := (capacity + GiB - 1) >> 30
		resp.Volume.CapacityBytes = quota << 30
		resp.Volume.VolumeContext["volumeCapacity"] = "true"
		if err := cs.nasClient.SetDirQuota(&sdk.SetDirQuotaRequest{
			FileSystemId: &filesystemId,
			Path:         &path,
			SizeLimit:    &quota,
			QuotaType:    tea.String("Enforcement"),
			UserType:     tea.String("AllUsers"),
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "nas:SetDirQuota failed: %v", err)
		}
	}

	return resp, nil
}

func (cs *subpathController) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest, pv *corev1.PersistentVolume) (*csi.DeleteVolumeResponse, error) {
	attributes := pv.Spec.CSI.VolumeAttributes
	var (
		filesystemId      string
		path              = attributes["path"]
		recycleBinEnabled bool
	)
	if cnfsName := attributes["containerNetworkFileSystem"]; cnfsName != "" {
		cnfs, err := cs.config.CNFSGetter.GetCNFS(ctx, cnfsName)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get CNFS %s: %v", cnfsName, err)
		}
		filesystemId = cnfs.Status.FsAttributes.FilesystemID
		recycleBinEnabled, _ = strconv.ParseBool(cnfs.Status.FsAttributes.EnableTrashCan)
	} else {
		server := attributes["server"]
		filesystemId, _, _ = strings.Cut(server, "-")
		var err error
		recycleBinEnabled, err = cs.isRecycleBinEnabled(filesystemId)
		if err != nil {
			return nil, err
		}
	}
	if filesystemId == "" {
		return nil, status.Error(codes.InvalidArgument, "empty filesystemId")
	}
	// cancel dir quota
	if attributes["volumeCapacity"] == "true" {
		if err := cs.nasClient.CancelDirQuota(&sdk.CancelDirQuotaRequest{
			FileSystemId: &filesystemId,
			Path:         &path,
			UserType:     tea.String("AllUsers"),
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "nas:CancelDirQuota failed: %v", err)
		}
	}
	if cs.config.DisableSubpathFinalizer {
		logrus.Warnf("deletion finalizer not enabled, skip subpath deletion for %s", req.VolumeId)
		return &csi.DeleteVolumeResponse{}, nil
	}

	// Patch finalizer on PV if need delete or archive subpath. The actual deletion/archiving will be executed
	// by storage-controller who is responsible to remove the finalzier.

	// Delete subpath only when parameters["archiveOnDelete"] exists and has a false value.
	// If StorageClass not found, always archive on delete.
	sc, err := cs.config.KubeClient.StorageV1().StorageClasses().Get(ctx, pv.Spec.StorageClassName, metav1.GetOptions{})
	if err != nil && !apierrors.IsNotFound(err) {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	archive := true
	if sc != nil {
		if value, exists := sc.Parameters["archiveOnDelete"]; exists {
			boolValue, err := strconv.ParseBool(value)
			if err == nil {
				archive = boolValue
			}
		}
	}

	var finalizer string
	if archive {
		finalizer = subpathArchiveFinalizer
	} else {
		if !recycleBinEnabled {
			logrus.WithFields(logrus.Fields{
				"filesystemId": filesystemId,
				"pv":           pv.Name,
			}).Warnf("skip subpath deletion because recycle bin of filesystem not enabled")
			return &csi.DeleteVolumeResponse{}, nil
		}
		finalizer = subpathDeletionFinalizer
	}
	if err := cs.patchFinalizerOnPV(ctx, pv, finalizer); err != nil {
		return nil, err
	}
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *subpathController) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest, pv *corev1.PersistentVolume) (*csi.ControllerExpandVolumeResponse, error) {
	attributes := pv.Spec.CSI.VolumeAttributes
	var (
		filesystemId string
		path         = attributes["path"]
	)
	// using cnfs or not
	if cnfsName := attributes["containerNetworkFileSystem"]; cnfsName != "" {
		cnfs, err := cs.config.CNFSGetter.GetCNFS(ctx, cnfsName)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get CNFS %s: %v", cnfsName, err)
		}
		filesystemId = cnfs.Status.FsAttributes.FilesystemID
	} else {
		server := attributes["server"]
		filesystemId, _, _ = strings.Cut(server, "-")
	}
	if filesystemId == "" {
		return nil, status.Error(codes.InvalidArgument, "empty filesystemId")
	}
	capacity := req.GetCapacityRange().GetRequiredBytes()
	if attributes["volumeCapacity"] == "true" {
		quota := (capacity + GiB - 1) >> 30
		if err := cs.nasClient.SetDirQuota(&sdk.SetDirQuotaRequest{
			FileSystemId: &filesystemId,
			Path:         &path,
			SizeLimit:    &quota,
			QuotaType:    tea.String("Enforcement"),
			UserType:     tea.String("AllUsers"),
		}); err != nil {
			return nil, status.Errorf(codes.Internal, "nas:SetDirQuota failed: %v", err)
		}
		return &csi.ControllerExpandVolumeResponse{CapacityBytes: quota << 30}, nil
	}
	if attributes["mountType"] == "losetup" {
		return &csi.ControllerExpandVolumeResponse{
			CapacityBytes:         capacity,
			NodeExpansionRequired: true,
		}, nil
	}
	logrus.Warn("volume capacity not enabled when provision, skip quota expandsion")
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: capacity}, nil
}

func (cs *subpathController) patchFinalizerOnPV(ctx context.Context, pv *corev1.PersistentVolume, finalizer string) error {
	ac, err := acv1.ExtractPersistentVolume(pv, finalizerFieldManager)
	if err != nil {
		return err
	}
	for _, f := range ac.Finalizers {
		if f == finalizer {
			return nil
		}
	}
	ac.WithFinalizers(finalizer)
	_, err = cs.config.KubeClient.CoreV1().PersistentVolumes().Apply(ctx, ac, metav1.ApplyOptions{FieldManager: finalizerFieldManager, Force: true})
	if err != nil {
		return status.Error(codes.Internal, err.Error())
	}
	logrus.Infof("patched finalizer %s on pv %s", finalizer, pv.Name)
	return nil
}

func (cs *subpathController) getFilesystemType(filesystemId string) (string, error) {
	resp, err := cs.nasClient.DescribeFileSystemBriefInfo(filesystemId)
	if err != nil {
		return "", status.Errorf(codes.Internal, "nas:DescribeFileSystemBriefInfos failed: %v", err)
	}
	if resp.Body == nil || resp.Body.FileSystems == nil || len(resp.Body.FileSystems.FileSystem) == 0 {
		return "", status.Errorf(codes.InvalidArgument, "filesystemId %q not found", filesystemId)
	}
	filesystemInfo := resp.Body.FileSystems.FileSystem[0]
	filesystemType := tea.StringValue(filesystemInfo.FileSystemType)
	logrus.Infof("nas:DescribeFileSystemBriefInfos succeeded, filesystemType of %s is %q", filesystemId, filesystemType)
	return filesystemType, nil
}

func (cs *subpathController) isRecycleBinEnabled(filesystemId string) (bool, error) {
	resp, err := cs.nasClient.GetRecycleBinAttribute(filesystemId)
	if err != nil {
		return false, status.Errorf(codes.Internal, "nas:GetRecycleBinAttribute failed: %v", err)
	}
	return resp.Body != nil && resp.Body.RecycleBinAttribute != nil && tea.StringValue(resp.Body.RecycleBinAttribute.Status) == "Enable", nil
}

var muxServerSelector = &rrMuxServerSelector{
	servers: make(map[string]int),
}

type rrMuxServerSelector struct {
	sync.Mutex
	servers map[string]int
}

func (s *rrMuxServerSelector) SelectNfsServer(muxServer string) (string, string) {
	s.Lock()
	defer s.Unlock()
	var servers, paths []string
	for _, str := range strings.Split(muxServer, ",") {
		server, path := s.parse(str)
		if server == "" {
			return "", ""
		}
		servers = append(servers, server)
		paths = append(paths, path)
	}
	if len(servers) < 2 {
		return servers[0], paths[0]
	}
	index := 0
	if i, ok := s.servers[muxServer]; ok {
		index = (i + 1) % len(servers)
	}
	s.servers[muxServer] = index
	return servers[index], paths[index]
}

func (s *rrMuxServerSelector) parse(serverWithPath string) (server, path string) {
	server, path, found := strings.Cut(serverWithPath, ":")
	if found && path != "" {
		path = filepath.Clean(path)
	} else {
		path = "/"
	}
	return
}
