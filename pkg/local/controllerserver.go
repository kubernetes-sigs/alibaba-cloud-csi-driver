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
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-csi/drivers/pkg/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/adapter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/generator"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/manager"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

type controllerServer struct {
	*csicommon.DefaultControllerServer
	client     kubernetes.Interface
	driverName string
}

const (
	// LvmVolumeType lvm volume type
	LvmVolumeType = "LVM"
	// PmemVolumeType lvm volume type
	PmemVolumeType = "PMEM"
	// QuotaPathVolumeType ...
	QuotaPathVolumeType = "QuotaPath"
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
	// TopologyYodaNodeKey define host name of node
	TopologyYodaNodeKey = "topology.yodaplugin.csi.alibabacloud.com/hostname"
	// PvcNameTag in annotations
	PvcNameTag = "csi.storage.k8s.io/pvc/name"
	// PvcNsTag in annotations
	PvcNsTag = "csi.storage.k8s.io/pvc/namespace"
	// NodeSchedueTag in annotations
	NodeSchedueTag = "volume.kubernetes.io/selected-node"
	// StorageSchedueTag in annotations
	StorageSchedueTag = "volume.kubernetes.io/selected-storage"
	// LastAppliyAnnotationTag tag
	LastAppliyAnnotationTag = "kubectl.kubernetes.io/last-applied-configuration"
	// CsiProvisionerIdentity tag
	CsiProvisionerIdentity = "storage.kubernetes.io/csiProvisionerIdentity"
	// CsiProvisionerTag tag
	CsiProvisionerTag = "volume.beta.kubernetes.io/storage-provisioner"
)

// the map of req.Name and csi.Volume
var createdVolumeMap = map[string]*csi.Volume{}

var supportVolumeTypes = []string{LvmVolumeType, PmemVolumeType, QuotaPathVolumeType, MountPointType, DeviceVolumeType}

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

// CreateVolume csi interface
func (cs *controllerServer) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	if err := cs.Driver.ValidateControllerServiceRequest(csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME); err != nil {
		log.Errorf("CreateVolume: Invalid create local volume req: %v", req)
		return nil, err
	}
	if req.Name == "" {
		log.Errorf("CreateVolume: local volume Name is empty")
		return nil, status.Error(codes.InvalidArgument, "CreateVolume: local Volume Name cannot be empty")
	}
	if req.VolumeCapabilities == nil {
		log.Errorf("CreateVolume: local Volume Capabilities cannot be empty")
		return nil, status.Error(codes.InvalidArgument, "Volume Capabilities cannot be empty")
	}
	pvcName, pvcNameSpace, volumeType, nodeSelected, storageSelected := "", "", "", "", ""

	volumeID := req.GetName()
	response := &csi.CreateVolumeResponse{}
	parameters := req.GetParameters()
	if value, ok := parameters[VolumeTypeKey]; ok {
		for _, supportVolType := range supportVolumeTypes {
			if supportVolType == value {
				volumeType = value
			}
		}
	}
	if volumeType == "" {
		log.Errorf("CreateVolume: Create volume %s with error volumeType %v", volumeID, parameters)
		return nil, status.Error(codes.InvalidArgument, "Local driver only support LVM/MountPoint/Device/PmemDirect/PmemQuotaPath volume type, no "+volumeType)
	}

	if value, ok := createdVolumeMap[req.Name]; ok {
		log.Infof("CreateVolume: local volume already be created, pvName: %s, VolumeId: %s, volumeContext: %v", req.Name, value.VolumeId, value.VolumeContext)
		return &csi.CreateVolumeResponse{Volume: value}, nil
	}

	if value, ok := parameters[PvcNameTag]; ok {
		pvcName = value
	}
	if value, ok := parameters[PvcNsTag]; ok {
		pvcNameSpace = value
	}
	if value, ok := parameters[NodeSchedueTag]; ok {
		nodeSelected = value
	}
	if value, ok := parameters[StorageSchedueTag]; ok {
		storageSelected = value
	}
	// Log inputs
	log.Infof("Starting to Create %s volume %s with: pvcName(%s), pvcNameSpace(%s), nodeSelected(%s), storageSelected(%s)", volumeType, volumeID, pvcName, pvcNameSpace, nodeSelected, storageSelected)

	// Schedule lvm volume Info
	paraList := map[string]string{}
	switch volumeType {
	case LvmVolumeType:
		var err error
		// Node and Storage have been scheduled (select volumeGroup)
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

		// if vgName configed in storageclass, use it first;
		if value, ok := paraList[VgNameTag]; ok && value != "" {
			storageSelected = value
		}

		// Volume Options
		options := &client.LVMOptions{}
		options.Name = req.Name
		options.VolumeGroup = storageSelected
		if value, ok := parameters[LvmTypeTag]; ok && value == StripingType {
			options.Striping = true
		}
		options.Size = uint64(req.GetCapacityRange().GetRequiredBytes())

		if types.GlobalConfigVar.GrpcProvision && nodeSelected != "" && storageSelected != "" {
			conn, err := cs.getNodeConn(nodeSelected)
			if err != nil {
				log.Errorf("CreateVolume: New lvm %s Connection to node %s with error: %s", req.Name, nodeSelected, err.Error())
				return nil, err
			}
			defer conn.Close()
			if lvmName, err := conn.GetLvm(ctx, storageSelected, volumeID); err == nil && lvmName == "" {
				outstr, err := conn.CreateLvm(ctx, options)
				if err != nil {
					log.Errorf("CreateVolume: Create lvm %s/%s, options: %v with error: %s", storageSelected, volumeID, options, err.Error())
					return nil, errors.New("Create Lvm with error " + err.Error())
				}
				log.Infof("CreateLvm: Successful Create lvm %s/%s in node %s with response %s", storageSelected, volumeID, nodeSelected, outstr)
			} else if err != nil {
				log.Errorf("CreateVolume: Get lvm %s from node %s with error: %s", req.Name, nodeSelected, err.Error())
				return nil, err
			} else {
				log.Infof("CreateVolume: lvm volume already created %s at node %s", req.Name, nodeSelected)
			}
		} else if !types.GlobalConfigVar.GrpcProvision && nodeSelected != "" && storageSelected != "" {
			createLabels := map[string]string{}
			optBytes, err := json.Marshal(options)
			if err != nil {
				log.Errorf("CreateVolume: Marshal lvm options error: %s, %s", req.Name, err.Error())
				return nil, err
			}
			createLabels[types.VolumeLifecycleLabel] = types.VolumeLifecycleCreating
			createLabels[types.VolumeSpecLabel] = string(optBytes)
			if err := generator.CreateVolumeWithAnnotations(pvcNameSpace, pvcName, createLabels); err != nil {
				log.Errorf("CreateVolume: create volume with label for volume %s %s at node %s error: %s", req.Name, storageSelected, nodeSelected, err.Error())
				return nil, err
			}
			log.Infof("CreateVolume: Successful create lvm volume without GRPC %s/%s at node %s", storageSelected, req.Name, nodeSelected)
		}
	case MountPointType:
		var err error
		// Node and Storage have been scheduled
		if storageSelected != "" && nodeSelected != "" {
			paraList, err = mountpointScheduled(storageSelected, parameters)
			if err != nil {
				log.Errorf("CreateVolume: create mountpoint volume %s/%s at node %s error: %s", storageSelected, req.Name, nodeSelected, err.Error())
				return nil, status.Error(codes.InvalidArgument, "CreateVolume: Parse mountpoint all scheduled info error "+err.Error())
			}
		} else if nodeSelected != "" {
			paraList, err = mountpointPartScheduled(nodeSelected, pvcName, pvcNameSpace, parameters)
			if err != nil {
				log.Errorf("CreateVolume: part schedule mountpoint volume %s at node %s error: %s", req.Name, nodeSelected, err.Error())
				return nil, status.Error(codes.InvalidArgument, "Parse mountpoint part schedule info error "+err.Error())
			}
		} else {
			nodeID := ""
			nodeID, paraList, err = mountpointNoScheduled(parameters)
			if err != nil {
				log.Errorf("CreateVolume: schedule mountpoint volume %s error: %s", req.Name, err.Error())
				return nil, status.Error(codes.InvalidArgument, "Parse mountpoint schedule info error "+err.Error())
			}
			nodeSelected = nodeID
		}
		log.Infof("CreateVolume: Successful create mountpoint volume %s/%s at node %s", storageSelected, req.Name, nodeSelected)
	case DeviceVolumeType:
		var err error
		// Node and Storage have been scheduled
		if storageSelected != "" && nodeSelected != "" {
			paraList, err = deviceScheduled(storageSelected, parameters)
			if err != nil {
				log.Errorf("CreateVolume: create device volume %s/%s at node %s error: %s", storageSelected, req.Name, nodeSelected, err.Error())
				return nil, status.Error(codes.InvalidArgument, "Parse Device all scheduled info error "+err.Error())
			}
		} else if nodeSelected != "" {
			paraList, err = devicePartScheduled(nodeSelected, pvcName, pvcNameSpace, parameters)
			if err != nil {
				log.Errorf("CreateVolume: part schedule device volume %s at node %s error: %s", req.Name, nodeSelected, err.Error())
				return nil, status.Error(codes.InvalidArgument, "Parse Device part schedule info error "+err.Error())
			}
		} else {
			nodeID := ""
			nodeID, paraList, err = deviceNoScheduled(parameters)
			if err != nil {
				log.Errorf("CreateVolume: schedule device volume %s error: %s", req.Name, err.Error())
				return nil, status.Error(codes.InvalidArgument, "Parse Device schedule info error "+err.Error())
			}
			nodeSelected = nodeID
		}
		log.Infof("CreateVolume: Successful create device volume %s/%s at node %s", storageSelected, req.Name, nodeSelected)
	case PmemVolumeType:
		if nodeSelected != "" {
			// only support pmem direct type
			conn, err := cs.getNodeConn(nodeSelected)
			if err != nil {
				log.Errorf("CreateVolume: create connect with node %s volume %s with error: %s", nodeSelected, req.Name, err.Error())
				return nil, err
			}
			defer conn.Close()
			options := &client.NameSpaceOptions{}
			options.Name = req.Name

			// pmem direct type need Region selection
			// First: scheduled aonnotation from pvc;
			// Second: pmemRegion setting in storageclass
			// Third: default as region0
			options.Region = manager.PmemRegionNameDefault
			if value, ok := parameters["pmemRegion"]; ok {
				options.Region = value
			}
			if storageSelected != "" {
				options.Region = storageSelected
			}

			options.Size = uint64(req.GetCapacityRange().GetRequiredBytes())
			if namespace, err := conn.GetNameSpace(ctx, options.Region, volumeID); err == nil && namespace == "" {
				newNameSpace, err := conn.CreateNameSpace(ctx, options)
				if err != nil {
					log.Errorf("CreateVolume: Create Pmem direct Namespace %s, options: %v at node with error: %s", volumeID, options, nodeSelected, err.Error())
					return nil, errors.New("Create Pmem direct with error " + err.Error())
				}
				parameters["pmemRegion"] = options.Region
				parameters["pmemNameSpace"] = newNameSpace.Dev
				parameters["pmemBlockDev"] = newNameSpace.BlockDev
				log.Infof("CreatePmem: Successful Create Pmem namespace %s with response %v", volumeID, newNameSpace)

			} else if err != nil {
				log.Errorf("CreateVolume: Get PMEM namespace %s at node %s with error: %s", req.Name, nodeSelected, err.Error())
				return nil, err
			} else {
				log.Infof("CreateVolume: PMEM namespace already created %s, %s at node %s", req.Name, namespace, nodeSelected)
			}
		} else {
			return nil, errors.New("CreateVolume: PMEMDirect type nodeselected must not be None")
		}
	case QuotaPathVolumeType:
		if nodeSelected != "" {
			conn, err := cs.getNodeConn(nodeSelected)
			if err != nil {
				log.Errorf("CreateVolume: New QuotaPath volume %s Connection node %s with error: %s", req.Name, nodeSelected, err.Error())
				return nil, err
			}
			defer conn.Close()
			size := fmt.Sprintf("%d", req.GetCapacityRange().GetRequiredBytes())
			kSize := fmt.Sprintf("%d", req.GetCapacityRange().GetRequiredBytes()/1024)
			log.Infof("CreateVolume: create quotaPath type volume %s with node(%s), rootPath(%s), size(%s)KB", req.Name, nodeSelected, storageSelected, kSize)
			_, projectQuotaSubpath, err := conn.CreateProjQuotaSubpath(ctx, req.Name, size, storageSelected)
			if err != nil {
				log.Errorf("CreateVolume: create quotaPath %s at node %s with error: %s", req.Name, nodeSelected, err.Error())
				return nil, err
			}
			_, err = conn.SetSubpathProjQuota(ctx, projectQuotaSubpath, kSize, kSize, "", "")
			if err != nil {
				log.Errorf("CreateVolume: set quotaPath %s at node %s with error: %s", req.Name, nodeSelected, err.Error())
				return nil, err
			}
			parameters[ProjQuotaFullPath] = projectQuotaSubpath
		} else {
			log.Errorf("CreateVolume: QuotaPath type nodeselected must not be None: %s", req.Name)
			return nil, errors.New("CreateVolume: QuotaPath type nodeselected must not be None")
		}
		log.Infof("CreateVolume: Successful Create QuotaPath volume %s at node %s", req.Name, nodeSelected)
	default:
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
		} else if key == CsiProvisionerIdentity {
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
		parameters[NodeSchedueTag] = nodeSelected
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

	createdVolumeMap[req.Name] = response.Volume
	log.Infof("Success create Volume: %s, Size: %d, Parameters: %v", volumeID, req.GetCapacityRange().GetRequiredBytes(), response.Volume)
	return response, nil
}

func (cs *controllerServer) getNodeConn(nodeSelected string) (client.Connection, error) {
	addr, err := getNodeAddr(cs.client, nodeSelected)
	if err != nil {
		log.Errorf("CreateVolume: Get node %s address with error: %s", nodeSelected, err.Error())
		return nil, err
	}
	conn, err := client.NewGrpcConnection(addr, connectTimeout)
	return conn, err
}

// DeleteVolume csi interface
func (cs *controllerServer) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	log.Infof("DeleteVolume: deleting local volume: %s", req.GetVolumeId())
	volumeID := req.GetVolumeId()
	nodeName, vgName, pvObj, err := getPvSpec(cs.client, volumeID, cs.driverName)
	if err != nil {
		log.Errorf("DeleteVolume: get pv spec %s with error: %s", volumeID, err.Error())
		return nil, err
	}
	volumeType := ""
	if value, ok := pvObj.Spec.CSI.VolumeAttributes[VolumeTypeKey]; ok {
		volumeType = value
	}

	switch volumeType {
	case LvmVolumeType:
		if types.GlobalConfigVar.GrpcProvision && nodeName != "" {
			conn, err := cs.getNodeConn(nodeName)
			if err != nil {
				log.Errorf("DeleteVolume: New lvm %s Connection at node %s with error: %s", req.GetVolumeId(), nodeName, err.Error())
				return nil, err
			}
			defer conn.Close()
			if lvmName, err := conn.GetLvm(ctx, vgName, volumeID); err == nil && lvmName != "" {
				if err := conn.DeleteLvm(ctx, vgName, volumeID); err != nil {
					log.Errorf("DeleteVolume: Remove lvm %s/%s at node %s with error: %s", vgName, volumeID, nodeName, err.Error())
					return nil, errors.New("DeleteVolume: Remove Lvm " + volumeID + " with error " + err.Error())
				}
				log.Infof("DeleteLvm: Successful Delete lvm %s/%s at node %s", vgName, volumeID, nodeName)
			} else if err == nil && lvmName == "" {
				log.Infof("DeleteVolume: get lvm empty, skip deleting %s", volumeID)
			} else if err != nil && strings.Contains(err.Error(), "Failed to find logical volume") {
				log.Infof("DeleteVolume: lvm volume not found, skip deleting %s", volumeID)
			} else if err != nil && strings.Contains(err.Error(), "Volume group \""+vgName+"\" not found") {
				log.Infof("DeleteVolume: Volume group not found, skip deleting %s", volumeID)
			} else {
				log.Errorf("DeleteVolume: Get lvm for %s with error: %s", req.GetVolumeId(), err.Error())
				return nil, err
			}
		} else if !types.GlobalConfigVar.GrpcProvision && nodeName != "" {
			createLabels := map[string]string{}
			createLabels[types.VolumeLifecycleLabel] = types.VolumeLifecycleDeleting
			createLabels[types.VolumeSpecLabel] = vgName + "/" + volumeID
			if err := generator.DeleteVolumeWithAnnotations(volumeID, createLabels); err != nil {
				log.Errorf("DeleteVolume: delete volume with label for volume %s error: %s", volumeID, err.Error())
				return nil, err
			}
			log.Infof("DeleteVolume: delete local volume %s with label at node %s", volumeID, nodeName)
		} else {
			log.Infof("DeleteVolume: delete local volume %s with node empty", volumeID)
		}

	case MountPointType:
		if pvObj.Spec.PersistentVolumeReclaimPolicy == v1.PersistentVolumeReclaimDelete {
			if pvObj.Spec.NodeAffinity == nil {
				log.Errorf("DeleteVolume: Get Lvm Spec for volume %s, with nil nodeAffinity", volumeID)
				return nil, errors.New("Get Lvm Spec for volume " + volumeID + ", with nil nodeAffinity")
			}
			if pvObj.Spec.NodeAffinity.Required == nil || len(pvObj.Spec.NodeAffinity.Required.NodeSelectorTerms) == 0 {
				log.Errorf("DeleteVolume: Get Lvm Spec for volume %s, with nil Required", volumeID)
				return nil, errors.New("Get Lvm Spec for volume " + volumeID + ", with nil Required")
			}
			if len(pvObj.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions) == 0 {
				log.Errorf("DeleteVolume: Get Lvm Spec for volume %s, with nil MatchExpressions", volumeID)
				return nil, errors.New("Get Lvm Spec for volume " + volumeID + ", with nil MatchExpressions")
			}
			key := pvObj.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions[0].Key
			if key != TopologyNodeKey && key != TopologyYodaNodeKey {
				log.Errorf("DeleteVolume: Get Lvm Spec for volume %s, with key %s", volumeID, key)
				return nil, errors.New("Get Lvm Spec for volume " + volumeID + ", with key" + key)
			}
			nodes := pvObj.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions[0].Values
			if len(nodes) == 0 {
				log.Errorf("DeleteVolume: Get MountPoint Spec for volume %s, with empty nodes", volumeID)
				return nil, errors.New("MountPoint Pv is illegal, No node info")
			}
			nodeName := nodes[0]
			conn, err := cs.getNodeConn(nodeName)
			if err != nil {
				log.Errorf("DeleteVolume: New mountpoint %s Connection error: %s", req.GetVolumeId(), err.Error())
				return nil, err
			}
			defer conn.Close()
			path := ""
			if value, ok := pvObj.Spec.CSI.VolumeAttributes[MountPointType]; ok {
				path = value
			}
			if path == "" {
				log.Errorf("DeleteVolume: Get MountPoint Path for volume %s, with empty", volumeID)
				return nil, errors.New("MountPoint Path is empty")
			}
			if err := conn.CleanPath(ctx, path); err != nil {
				log.Errorf("DeleteVolume: Remove mountpoint for %s with error: %s", req.GetVolumeId(), err.Error())
				return nil, errors.New("DeleteVolume: Delete mountpoint Failed: " + err.Error())
			}
		}
		log.Infof("DeleteVolume: successful delete MountPoint volume(%s)...", volumeID)
	case DeviceVolumeType:
		log.Infof("DeleteVolume: successful delete Device volume(%s)...", volumeID)
	case PmemVolumeType:
		if nodeName != "" {
			conn, err := cs.getNodeConn(nodeName)
			if err != nil {
				log.Errorf("DeleteVolume: New PMEM %s Connection at node %s with error: %s", req.GetVolumeId(), nodeName, err.Error())
				return nil, err
			}
			defer conn.Close()
			if _, ok := pvObj.Spec.CSI.VolumeAttributes["pmemNameSpace"]; !ok {
				log.Errorf("DeleteVolume: Direct PMEM volume can not found NameSpace: %s", volumeID)
				return nil, errors.New("DeleteVolume Direct PMEM volume can not found NameSpace " + volumeID)
			}
			namespace := pvObj.Spec.CSI.VolumeAttributes["pmemNameSpace"]
			if pmemName, err := conn.GetNameSpace(ctx, "", volumeID); err == nil && pmemName != "" {
				if err := conn.DeleteNameSpace(ctx, namespace); err != nil {
					log.Errorf("DeleteVolume: Remove PMEM direct volume %s at node %s with error: %s", volumeID, nodeName, err.Error())
					return nil, errors.New("DeleteVolume: Remove Pmem namespace with error " + err.Error())
				}
				log.Infof("DeleteLvm: Successful Delete PMEM volume %s at node %s", volumeID, nodeName)
			} else if err == nil && pmemName == "" {
				log.Infof("DeleteVolume: get PMEM empty at node %s, skip deleting %s", nodeName, volumeID)
			} else {
				log.Errorf("DeleteVolume: Get PMEM volume %s at node %s with error: %s", req.VolumeId, nodeName, err.Error())
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("DeleteVolume: failed to delete volume without nodeAffinity %s", volumeID)
		}
	case QuotaPathVolumeType:
		if nodeName != "" {
			conn, err := cs.getNodeConn(nodeName)
			if err != nil {
				log.Errorf("DeleteVolume: get QuotaPath volume %s Connection at node %s with error: %s", req.VolumeId, nodeName, err.Error())
				return nil, err
			}
			defer conn.Close()

			if _, ok := pvObj.Spec.CSI.VolumeAttributes[ProjQuotaFullPath]; !ok {
				log.Errorf("DeleteVolume: QuotaPath volume %s not have projQuotaFullPath parameter", req.VolumeId)
				return nil, fmt.Errorf("DeleteVolume: QuotaPath volume %s not have projQuotaFullPath parameter", req.VolumeId)
			}
			quotaPath := pvObj.Spec.CSI.VolumeAttributes[ProjQuotaFullPath]
			_, err = conn.RemoveProjQuotaSubpath(ctx, quotaPath)
			if err != nil {
				log.Errorf("DeleteVolume: Remove QuotaPath volume %s at node % with error %s", req.VolumeId, nodeName, err.Error())
				return nil, err
			}
		} else {
			log.Errorf("DeleteVolume: delete quotapath volume without nodeAffinity %s", volumeID)
			return nil, fmt.Errorf("DeleteVolume: delete quotapath volume without nodeAffinity %s", volumeID)
		}
	default:
		log.Errorf("DeleteVolume: volumeType %s not supported %s", volumeType, volumeID)
		return nil, status.Error(codes.InvalidArgument, "Local driver only support LVM volume type, no "+volumeType)
	}
	delete(createdVolumeMap, req.VolumeId)
	log.Infof("DeleteVolume: successful delete local volume %s", volumeID)
	return &csi.DeleteVolumeResponse{}, nil
}

// ControllerUnpublishVolume csi interface
func (cs *controllerServer) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {
	log.Infof("ControllerUnpublishVolume is called, do nothing by now: %s", req.VolumeId)
	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

// ControllerPublishVolume csi interface
func (cs *controllerServer) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	log.Infof("ControllerPublishVolume is called, do nothing by now: %s", req.VolumeId)
	return &csi.ControllerPublishVolumeResponse{}, nil
}

// ControllerExpandVolume csi interface
func (cs *controllerServer) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	log.Infof("ControllerExpandVolume::: %v", req)
	volSizeBytes := int64(req.GetCapacityRange().GetRequiredBytes())
	if types.GlobalConfigVar.Scheduler == yodaDriverName {
		volSizeGB := int((volSizeBytes + 1024*1024*1024 - 1) / (1024 * 1024 * 1024))
		volumeID := req.GetVolumeId()
		pvObj, err := getPvObj(cs.client, volumeID)
		if err != nil {
			log.Errorf("ControllerExpandVolume: get volume object %s error with: %s", volumeID, err.Error())
			return nil, err
		}
		if pvObj.Spec.CSI == nil {
			log.Errorf("ControllerExpandVolume: volume is not csi type %s", volumeID)
			return nil, errors.New("ControllerExpandVolume: volume is not csi type: " + volumeID)
		}
		attributes := pvObj.Spec.CSI.VolumeAttributes
		pvcName, pvcNameSpace := "", ""
		if value, ok := attributes[PvcNameTag]; ok {
			pvcName = value
		}
		if value, ok := attributes[PvcNsTag]; ok {
			pvcNameSpace = value
		}
		if err := adapter.ExpandVolume(pvcNameSpace, pvcName, volSizeGB); err != nil {
			log.Errorf("ControllerExpandVolume: expand volume %s to size %d meet error: %v", volumeID, volSizeGB, err)
			return nil, errors.New("ControllerExpandVolume: expand volume error " + err.Error())
		}

	}

	return &csi.ControllerExpandVolumeResponse{CapacityBytes: volSizeBytes, NodeExpansionRequired: true}, nil
}
