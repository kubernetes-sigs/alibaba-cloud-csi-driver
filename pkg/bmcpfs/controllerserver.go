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

package bmcpfs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	efloclient "github.com/alibabacloud-go/eflo-controller-20221215/v2/client"
	nasclient "github.com/alibabacloud-go/nas-20170626/v4/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/bmcpfs/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/version"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	informerv1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/klog/v2"
)

type controllerServer struct {
	common.GenericControllerServer
	vscManager     *internal.PrimaryVscManagerWithCache
	attachDetacher internal.CPFSAttachDetacher
	filsetManager  internal.CPFSFileSetManager
	nasClient      *nasclient.Client
	nodeInformer   cache.SharedIndexInformer
	skipDetach     bool
}

type BMCPFSNodeVolume struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	LingjunID       string
	volumesAttached []v1.AttachedVolume
	volumesInUse    []v1.UniqueVolumeName
}

func mustGetKubeClient() kubernetes.Interface {
	cfg := options.MustGetRestConfig()
	return kubernetes.NewForConfigOrDie(cfg)
}

func newControllerServer(region string) (*controllerServer, error) {
	efloClient, err := newEfloClient(region)
	if err != nil {
		return nil, err
	}

	skipDetach := false
	if skipDetachVal, _ := strconv.ParseBool(os.Getenv("SKIP_BMCPFS_DETACH")); skipDetachVal {
		skipDetach = skipDetachVal
	}

	nasClient, err := cloud.NewNasClientV2(region)
	if err != nil {
		return nil, err
	}
	kubeconfig := mustGetKubeClient()
	nodeInformer := informerv1.NewFilteredNodeInformer(kubeconfig, 0, cache.Indexers{NODEID_INDEX: nodeIndexFunc}, func(options *metav1.ListOptions) {
		options.FieldSelector = fields.OneTermEqualSelector("alibabacloud.com/lingjun-worker", "true").String()
	})
	nodeInformer.SetTransform(nodeTransformFunc)
	return &controllerServer{
		vscManager:     internal.NewPrimaryVscManagerWithCache(efloClient),
		attachDetacher: internal.NewCPFSAttachDetacher(nasClient),
		filsetManager:  internal.NewCPFSFileSetManager(nasClient),
		nasClient:      nasClient,
		nodeInformer:   nodeInformer,
		skipDetach:     skipDetach,
	}, nil
}

// CreateVolume ...
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(2).Info("starting")

	// Handle static volume creation (existing fileset)
	csiVolume, err := staticFileSetCreate(req, cs.filsetManager)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "create static volume failed: %v", err)
	}
	if csiVolume != nil {
		klog.Infof("CreateVolume: static volume create successful, pvName: %s, VolumeId: %s, volumeContext: %v", req.Name, csiVolume.VolumeId, csiVolume.VolumeContext)
		return &csi.CreateVolumeResponse{Volume: csiVolume}, nil
	}

	// Validate parameters
	if err := validateFileSetParameters(req); err != nil {
		klog.Errorf("CreateVolume: error parameters from input: %v, with error: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, "Invalid parameters from input: %v, with error: %v", req.Name, err)
	}

	// Extract parameters
	params := req.GetParameters()
	bmcpfsID := params["bmcpfsId"]
	fullPath := req.Name
	if rootPath, ok := params["path"]; ok && rootPath != "" {
		fullPath = filepath.Join(rootPath, req.Name)
	}
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())

	// Create fileset
	fileSetID, err := cs.filsetManager.CreateFileSet(ctx, bmcpfsID, req.Name, fullPath, 1000000, volSizeBytes, false)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Prepare volume context
	volumeContext := req.GetParameters()
	if volumeContext == nil {
		volumeContext = make(map[string]string)
	}
	volumeContext = updateVolumeContext(volumeContext)

	klog.Infof("CreateVolume: Successfully created FileSet %s: id[%s], filesystem[%s], path[%s]", req.GetName(), fileSetID, bmcpfsID, fullPath)

	tmpVol := createVolumeResponse(fileSetID, bmcpfsID, volSizeBytes, volumeContext)

	return &csi.CreateVolumeResponse{Volume: tmpVol}, nil
}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	logger := klog.FromContext(ctx)
	logger.V(2).Info("starting")

	// Parse volume ID to extract filesystem ID and fileset ID
	fsID, fileSetID, err := parseVolumeID(req.VolumeId)
	if err != nil {
		klog.Errorf("DeleteVolume: failed to parse volume ID %s: %v", req.VolumeId, err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	klog.Infof("DeleteVolume: deleting fileset %s from filesystem %s", fileSetID, fsID)

	// Check if fileset exists before attempting to delete
	_, err = cs.filsetManager.DescribeFileset(ctx, fsID, fileSetID)
	if err != nil {
		// If fileset doesn't exist, consider it already deleted
		klog.V(2).Infof("DeleteVolume: fileset %s not found in filesystem %s, considering it already deleted", fileSetID, fsID)
		return &csi.DeleteVolumeResponse{}, nil
	}

	// Delete the fileset
	err = cs.filsetManager.DeleteFileSet(ctx, fsID, fileSetID)
	if err != nil {
		klog.Errorf("DeleteVolume: failed to delete fileset %s from filesystem %s: %v", fileSetID, fsID, err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	klog.Infof("DeleteVolume: successfully deleted fileset %s from filesystem %s", fileSetID, fsID)
	return &csi.DeleteVolumeResponse{}, nil
}

func nodeIndexFunc(obj interface{}) ([]string, error) {
	node := obj.(*BMCPFSNodeVolume)
	return []string{node.LingjunID}, nil
}

func nodeTransformFunc(obj interface{}) (interface{}, error) {
	if _, ok := obj.(cache.DeletedFinalStateUnknown); ok {
		return obj, nil
	}
	if _, ok := obj.(*BMCPFSNodeVolume); ok {
		return obj, nil // this may happen when resync
	}
	node := obj.(*v1.Node)
	nodeVolume := &BMCPFSNodeVolume{
		TypeMeta: metav1.TypeMeta{
			Kind:       "Node",
			APIVersion: "v1",
		},
		ObjectMeta:      node.ObjectMeta,
		LingjunID:       node.Spec.ProviderID,
		volumesAttached: node.Status.VolumesAttached,
		volumesInUse:    node.Status.VolumesInUse,
	}
	return nodeVolume, nil
}

func (cs *controllerServer) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	return &csi.ControllerGetCapabilitiesResponse{
		Capabilities: common.ControllerRPCCapabilities(
			csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
			csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
		),
	}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	if !strings.HasPrefix(req.NodeId, LingjunNodeIDPrefix) {
		if req.VolumeContext[_vpcMountTarget] == "" {
			return nil, status.Errorf(codes.InvalidArgument, "missing %q config in volume context", _vpcMountTarget)
		}
		// TODO: try to use existing vpc mount target
		klog.InfoS("ControllerPublishVolume: use VPC MountTarget", "nodeId", req.NodeId)
		return &csi.ControllerPublishVolumeResponse{
			PublishContext: map[string]string{
				_networkType:    networkTypeVPC,
				_vpcMountTarget: req.VolumeContext[_vpcMountTarget],
			},
		}, nil
	}

	cpfsID, _ := parseVolumeHandle(req.VolumeId)
	// Get VscMountTarget of filesystem
	mt := req.VolumeContext[_vscMountTarget]
	if mt == "" {
		var err error
		mt, err = getMountTarget(cs.nasClient, cpfsID, networkTypeVSC)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "failed to get VscMountTarget: %v", err)
		}
	}

	// Get Primary vsc of Lingjun node
	lingjunInstanceID := strings.TrimPrefix(req.NodeId, LingjunNodeIDPrefix)
	if LingjunNodeIDPrefix == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid node id")
	}
	vscID, err := cs.vscManager.EnsurePrimaryVsc(ctx, lingjunInstanceID, false)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	klog.Info("Use VSC MountTarget for lingjun node", "nodeId", req.NodeId, "vscId", vscID)

	// Attach CPFS to VSC
	err = cs.attachDetacher.Attach(ctx, cpfsID, vscID)
	if err != nil {
		if autoSwitch, _ := strconv.ParseBool(req.VolumeContext[_mpAutoSwitch]); autoSwitch && internal.IsAttachNotSupportedError(err) {
			if req.VolumeContext[_vpcMountTarget] == "" {
				return nil, status.Errorf(codes.InvalidArgument, "missing %q config in volume context as vsc mountpoint not supported", _vpcMountTarget)
			}
			return &csi.ControllerPublishVolumeResponse{
				PublishContext: map[string]string{
					_networkType:    networkTypeVPC,
					_vpcMountTarget: req.VolumeContext[_vpcMountTarget],
				},
			}, nil
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	// TODO: if the cached vscid is already deleted, try to recreate a new primary vsc for lingjun node

	klog.InfoS("ControllerPublishVolume: attached cpfs to vsc", "vscMountTarget", mt, "vscId", vscID, "node", req.NodeId)
	return &csi.ControllerPublishVolumeResponse{
		PublishContext: map[string]string{
			_networkType:    networkTypeVSC,
			_vscID:          vscID,
			_vscMountTarget: mt,
		},
	}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (
	*csi.ControllerUnpublishVolumeResponse, error,
) {
	if !strings.HasPrefix(req.NodeId, LingjunNodeIDPrefix) || cs.skipDetach {
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}
	// Create Primary vsc for Lingjun node
	lingjunInstanceID := strings.TrimPrefix(req.NodeId, LingjunNodeIDPrefix)
	if LingjunNodeIDPrefix == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid node id")
	}
	vsc, err := cs.vscManager.GetPrimaryVscOf(lingjunInstanceID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if vsc == nil {
		klog.InfoS("ControllerUnpublishVolume: skip detaching cpfs from vsc as vsc not found", "node", req.NodeId)
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}
	// If `req.VolumeId` is a combination of `cpfsID` and `fsetID`, Detach will trigger an error.
	err = cs.attachDetacher.Detach(ctx, req.VolumeId, vsc.VscID)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	klog.InfoS("ControllerUnpublishVolume: detached cpfs from vsc", "node", req.NodeId, "filesystem", req.VolumeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

// KubernetesAlicloudIdentity is the user agent string for Eflo client
var KubernetesAlicloudIdentity = fmt.Sprintf("Kubernetes.Alicloud/CsiProvision.Bmcpfs-%s", version.VERSION)

const efloConnTimeout = 10

func newEfloClient(region string) (*efloclient.Client, error) {
	// lingjun region could be different from ack region
	// use another environment variable EFLO_CONTROLLER_REGION for special lingjun pre regions
	if r := os.Getenv("EFLO_CONTROLLER_REGION"); r != "" {
		region = r
	}
	config := new(openapi.Config).
		SetUserAgent(KubernetesAlicloudIdentity).
		SetRegionId(region).
		SetConnectTimeout(efloConnTimeout).
		SetGlobalParameters(&openapi.GlobalParameters{
			Queries: map[string]*string{
				"RegionId": &region,
			},
		})
	// set credential
	cred, err := utils.GetCredentialV2()
	if err != nil {
		return nil, fmt.Errorf("init credential: %w", err)
	}
	config = config.SetCredential(cred)
	// set endpoint
	ep := os.Getenv("EFLO_CONTROLLER_ENDPOINT")
	if ep != "" {
		config = config.SetEndpoint(ep)
	} else {
		config = config.SetEndpoint(fmt.Sprintf("eflo-controller-vpc.%s.aliyuncs.com", region))
	}
	// set protocol
	scheme := "HTTPS"
	if strings.Contains(region, "test") {
		// must use HTTP in lingjun test regions
		scheme = "HTTP"
	}
	if e := os.Getenv("ALICLOUD_CLIENT_SCHEME"); e != "" {
		scheme = e
	}
	config = config.SetProtocol(scheme)
	// init client
	return efloclient.NewClient(config)
}

func parseVolumeHandle(volumeHandle string) (string, string) {
	parts := strings.Split(volumeHandle, volumeHandleDelimiter)
	if len(parts) == 2 {
		return parts[0], parts[1]
	}
	return parts[0], ""
}

func getMountTarget(client *nasclient.Client, fsID, networkType string) (string, error) {
	resp, err := client.DescribeFileSystems(&nasclient.DescribeFileSystemsRequest{
		FileSystemId: &fsID,
	})
	if err != nil {
		return "", fmt.Errorf("nas:DescribeFileSystems failed: %w", err)
	}
	klog.V(4).InfoS("nas:DescribeFileSystems succeeded", "response", resp.Body)
	filesystems := resp.Body.FileSystems
	if filesystems == nil || len(filesystems.FileSystem) == 0 || filesystems.FileSystem[0] == nil {
		return "", nil
	}
	mountTargets := filesystems.FileSystem[0].MountTargets
	if mountTargets == nil {
		return "", nil
	}
	for _, mt := range mountTargets.MountTarget {
		t := tea.StringValue(mt.NetworkType)
		// mount targets with empty network type is vsc type
		if t == "" {
			t = networkTypeVSC
		}
		if t == networkType {
			mountTarget := tea.StringValue(mt.MountTargetDomain)
			status := tea.StringValue(mt.Status)
			klog.V(2).InfoS("Found cpfs mount target", "filesystem", fsID, "networkType", networkType, "mountTarget", mountTarget, "status", status)
			if status == "Active" {
				return mountTarget, nil
			}
		}
	}
	return "", fmt.Errorf("no active %s mount target found", networkType)
}
