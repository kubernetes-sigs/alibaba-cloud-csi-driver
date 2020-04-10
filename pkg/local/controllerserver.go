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

package local

import (
	"errors"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/client"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"strings"
	"time"
)

type controllerServer struct {
	*csicommon.DefaultControllerServer
	client     kubernetes.Interface
	driverName string
}

const (
	// LvmVolumeType lvm volume type
	LvmVolumeType = "LVM"
	// MountPointType type
	MountPointType = "MountPoint"
	// DeviceVolumeType type
	DeviceVolumeType = "Device"
	// VolumeTypeKey volume type key words
	VolumeTypeKey = "volumeType"
	// connection timeout
	connectTimeout = 3 * time.Second
	// TopologyNodeKey define host name of node
	TopologyNodeKey = "kubernetes.io/hostname"
	// PvcNameTag in annotations
	PvcNameTag = "csi.storage.k8s.io/pvc/name"
	// PvcNsTag in annotations
	PvcNsTag = "csi.storage.k8s.io/pvc/namespace"
	// NodeSchTag in annotations
	NodeSchTag = "volume.kubernetes.io/selected-node"
	// StorageSchTag in annotations
	StorageSchTag = "volume.kubernetes.io/selected-storage"
	// LastAppliyAnnotationTag tag
	LastAppliyAnnotationTag = "kubectl.kubernetes.io/last-applied-configuration"
	// CsiProvisionerIdentity tag
	CsiProvisionerIdentity = "storage.kubernetes.io/csiProvisionerIdentity"
	// CsiProvisionerTag tag
	CsiProvisionerTag = "volume.beta.kubernetes.io/storage-provisioner"
)

// newControllerServer creates a controllerServer object
func newControllerServer(d *csicommon.CSIDriver) *controllerServer {
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}
	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	return &controllerServer{
		DefaultControllerServer: csicommon.NewDefaultControllerServer(d),
		client:                  kubeClient,
	}
}

func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Errorf("Invalid create volume req: %v", req)
		return nil, err
	}
	if req.Name == "" {
		log.Errorf("CreateVolume: volume Name is empty")
		return nil, status.Error(codes.InvalidArgument, "CreateVolume: Volume Name cannot be empty")
	}
	if req.VolumeCapabilities == nil {
		log.Errorf("Volume Capabilities cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Volume Capabilities cannot be empty")
	}
	pvcName, pvcNameSpace, volumeType, nodeSelected, storageSelected := "", "", "", "", ""

	volumeID := req.GetName()
	response := &csi.CreateVolumeResponse{}
	parameters := req.GetParameters()
	if value, ok := parameters[VolumeTypeKey]; ok {
		volumeType = value
	}
	if volumeType != LvmVolumeType && volumeType != MountPointType && volumeType != DeviceVolumeType {
		log.Errorf("CreateVolume: Create volume %s with error volumeType %v", volumeID, parameters)
		return nil, status.Error(codes.InvalidArgument, "Local driver only support LVM/MountPoint/Device volume type, no "+volumeType)
	}

	if value, ok := parameters[PvcNameTag]; ok {
		pvcName = value
	}
	if value, ok := parameters[PvcNsTag]; ok {
		pvcNameSpace = value
	}
	if value, ok := parameters[NodeSchTag]; ok {
		nodeSelected = value
	}
	if value, ok := parameters[StorageSchTag]; ok {
		storageSelected = value
	}
	// Log inputs
	log.Infof("Starting to Create %s volume %s with: pvcName(%s), pvcNameSpace(%s), nodeSelected(%s), storageSelected(%s)", volumeType, volumeID, pvcName, pvcNameSpace, nodeSelected, storageSelected)

	// Schedule lvm volume Info
	paraList := map[string]string{}
	if volumeType == LvmVolumeType {
		var err error
		// Node and Storage have been scheduled
		if storageSelected != "" && nodeSelected != "" {
			paraList, err = lvmScheduled(storageSelected, parameters)
			if err != nil {
				log.Errorf("CreateVolume: lvm all scheduled volume %s with error: %s", volumeID, err.Error())
				return nil, status.Error(codes.InvalidArgument, "Parse lvm all schedule info error "+err.Error())
			}
		} else if nodeSelected != "" {
			paraList, err = lvmPartScheduled(nodeSelected, pvcName, pvcNameSpace, parameters)
			if err != nil {
				log.Errorf("CreateVolume: lvm part scheduled volume %s with error: %s", volumeID, err.Error())
				return nil, status.Error(codes.InvalidArgument, "Parse lvm part schedule info error "+err.Error())
			}
		} else {
			nodeID := ""
			nodeID, paraList, err = lvmNoScheduled(parameters)
			if err != nil {
				log.Errorf("CreateVolume: lvm No scheduled volume %s with error: %s", volumeID, err.Error())
				return nil, status.Error(codes.InvalidArgument, "Parse lvm schedule info error "+err.Error())
			}
			nodeSelected = nodeID
		}
	} else if volumeType == MountPointType {
		var err error
		// Node and Storage have been scheduled
		if storageSelected != "" && nodeSelected != "" {
			paraList, err = mountpointScheduled(storageSelected, parameters)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, "Parse mountpoint all scheduled info error "+err.Error())
			}
		} else if nodeSelected != "" {
			paraList, err = mountpointPartScheduled(nodeSelected, pvcName, pvcNameSpace, parameters)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, "Parse lvm part schedule info error "+err.Error())
			}
		} else {
			nodeID := ""
			nodeID, paraList, err = mountpointNoScheduled(parameters)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, "Parse lvm schedule info error "+err.Error())
			}
			nodeSelected = nodeID
		}
	} else if volumeType == DeviceVolumeType {
		var err error
		// Node and Storage have been scheduled
		if storageSelected != "" && nodeSelected != "" {
			paraList, err = deviceScheduled(storageSelected, parameters)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, "Parse Device all scheduled info error "+err.Error())
			}
		} else if nodeSelected != "" {
			paraList, err = devicePartScheduled(nodeSelected, pvcName, pvcNameSpace, parameters)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, "Parse Device part schedule info error "+err.Error())
			}
		} else {
			nodeID := ""
			nodeID, paraList, err = deviceNoScheduled(parameters)
			if err != nil {
				return nil, status.Error(codes.InvalidArgument, "Parse Device schedule info error "+err.Error())
			}
			nodeSelected = nodeID
		}
	} else {
		log.Errorf("CreateVolume: Create with no support volume type %s", volumeType)
		return nil, status.Error(codes.InvalidArgument, "Create with no support type "+volumeType)
	}

	// Append necessary parameters
	for key, value := range paraList {
		parameters[key] = value
	}
	// remove not necessary labels
	for key := range parameters {
		if key == LastAppliyAnnotationTag {
			delete(parameters, key)
		} else if key == CsiProvisionerTag {
			delete(parameters, key)
		}
	}

	if nodeSelected == "" {
		response = &csi.CreateVolumeResponse{
			Volume: &csi.Volume{
				VolumeId:      volumeID,
				CapacityBytes: req.GetCapacityRange().GetRequiredBytes(),
				VolumeContext: parameters,
			},
		}
	} else {
		parameters[NodeSchTag] = nodeSelected
		response = &csi.CreateVolumeResponse{
			Volume: &csi.Volume{
				VolumeId:      volumeID,
				CapacityBytes: req.GetCapacityRange().GetRequiredBytes(),
				VolumeContext: parameters,
				AccessibleTopology: []*csi.Topology{
					{
						Segments: map[string]string{
							TopologyNodeKey: nodeSelected,
						},
					},
				},
			},
		}
	}

	log.Infof("Success create Volume: %s, Size: %d, Parameters: %v", volumeID, req.GetCapacityRange().GetRequiredBytes(), response.Volume)
	return response, nil
}

func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: deleting volume: %s", req.GetVolumeId())
	volumeID := req.GetVolumeId()
	pvObj, err := getPvObj(cs.client, volumeID)
	if err != nil {
		log.Errorf("DeleteVolume: get volume object %s error with: %s", volumeID, err.Error())
		return nil, err
	}
	if pvObj.Spec.CSI == nil {
		log.Errorf("DeleteVolume: Remove Lvm Failed, volume is not csi type %s", volumeID)
		return nil, errors.New("Remove Lvm Failed: volume is not csi type: " + volumeID)
	}
	volumeType := ""
	if value, ok := pvObj.Spec.CSI.VolumeAttributes[VolumeTypeKey]; ok {
		volumeType = value
	}

	if volumeType == LvmVolumeType {
		nodeName, vgName, err := getLvmSpec(cs.client, volumeID, cs.driverName)
		if err != nil {
			log.Errorf("DeleteVolume: get Lvm %s Spec with error %s", volumeID, err.Error())
			return nil, err
		}
		if nodeName != "" {
			addr, err := getLvmdAddr(cs.client, nodeName)
			if err != nil {
				log.Errorf("DeleteVolume: Get lvm volume %s address with error: %s", req.GetVolumeId(), err.Error())
				return nil, err
			}
			conn, err := client.NewLVMConnection(addr, connectTimeout)
			defer conn.Close()
			if err != nil {
				log.Errorf("DeleteVolume: New lvm %s Connection with error: %s", req.GetVolumeId(), err.Error())
				return nil, err
			}
			if _, err := conn.GetLvm(ctx, vgName, volumeID); err == nil {
				if err := conn.DeleteLvm(ctx, vgName, volumeID); err != nil {
					log.Errorf("DeleteVolume: Remove lvm %s/%s with error: %s", vgName, volumeID, err.Error())
					return nil, errors.New("Remove Lvm with error " + err.Error())
				}
			} else if strings.Contains(err.Error(), "Failed to find logical volume") {
				log.Infof("DeleteVolume: lvm volume not found, skip deleting %s", volumeID)
			} else if strings.Contains(err.Error(), "Volume group \""+vgName+"\" not found") {
				log.Infof("DeleteVolume: Volume group not found, skip deleting %s", volumeID)
			} else {
				log.Errorf("DeleteVolume: Get lvm for %s with error: %s", req.GetVolumeId(), err.Error())
				return nil, err
			}
		}
	} else if volumeType == MountPointType {
		if pvObj.Spec.PersistentVolumeReclaimPolicy == v1.PersistentVolumeReclaimDelete {
			nodes := pvObj.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions[0].Values
			if len(nodes) == 0 {
				log.Errorf("Get MountPoint Spec for volume %s, with empty nodes", volumeID)
				return nil, errors.New("MountPoint Pv is illegal, No node info")
			}
			nodeName := nodes[0]
			addr, err := getLvmdAddr(cs.client, nodeName)
			if err != nil {
				return nil, err
			}
			conn, err := client.NewLVMConnection(addr, connectTimeout)
			defer conn.Close()
			if err != nil {
				log.Errorf("DeleteVolume: New mountpoint %s Connection error: %s", req.GetVolumeId(), err.Error())
				return nil, err
			}
			path := ""
			if value, ok := pvObj.Spec.CSI.VolumeAttributes[MountPointType]; ok {
				path = value
			}
			if path == "" {
				log.Errorf("Get MountPoint Path for volume %s, with empty", volumeID)
				return nil, errors.New("MountPoint Path is empty")
			}
			if err := conn.CleanPath(ctx, path); err != nil {
				log.Errorf("DeleteVolume: Remove mountpoint for %s with error: %s", req.GetVolumeId(), err.Error())
				return nil, errors.New("Delete mountpoint Failed: " + err.Error())
			}
		}
		log.Infof("DeleteVolume: default to delete MountPoint volume(%s) type volume...", volumeID)
	} else if volumeType == DeviceVolumeType {
		log.Infof("DeleteVolume: default to delete Device volume type volume...")
	} else {
		log.Errorf("DeleteVolume: volumeType %s not supported %s", volumeType, volumeID)
		return nil, status.Error(codes.InvalidArgument, "Local driver only support LVM volume type, no "+volumeType)
	}
	log.Infof("DeleteVolume: successful delete volume %s", volumeID)
	return &csi.DeleteVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume is called, do nothing by now: %s", req.VolumeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof("ControllerPublishVolume is called, do nothing by now: %s", req.VolumeId)
	return &csi.ControllerPublishVolumeResponse{}, nil
}

func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	log.Infof("ControllerExpandVolume::: %v", req)
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
}
