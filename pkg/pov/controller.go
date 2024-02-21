package pov

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/pov/internal"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/template"
)

const (
	defaultVolumeSize            = 20 * utils.GiB
	ZONEID                       = "zoneid"
	DATAREDUNDANCYTYPE           = "dataredundancytype"
	PROTOCOLTYPE                 = "protocoltype"
	STORAGETYPE                  = "storagetype"
	FILESYSTEMNAME               = "filesystemname"
	SPACECAPACITY                = "spacecapacity"
	THROUGHPUTMODE               = "throughputmode"
	PROVISIONEDTHROUGHPUTINMIBPS = "provisionedthroughputinmibps"
	FILESYSTEMID                 = "filesystemid"
	// PVCNameKey contains name of the PVC for which is a volume provisioned.
	PVCNameKey = "csi.storage.k8s.io/pvc/name"
	// PVCNamespaceKey contains namespace of the PVC for which is a volume provisioned.
	PVCNamespaceKey = "csi.storage.k8s.io/pvc/namespace"
	// PVNameKey contains name of the final PV that will be used for the dynamically
	// provisioned volume
	PVNameKey   = "csi.storage.k8s.io/pv/name"
	TopologyKey = "topology.kubernetes.io/region"

	// volumeContext starting with labelAppendPrefix will automatically added to pv lables
	labelAppendPrefix = "csi.alibabacloud.com/label-prefix/"
	annVSCIDPrefix    = "csi.alibabacloud.com/ann-vsc-id"
	labelmpId         = "csi.alibabacloud.com/label-mp-id"
	labelFilesystemID = "csi.alibabacloud.com/label-pov-fsId"
)

// volumeCaps represents how the volume could be accessed.
// It is SINGLE_NODE_WRITER since EBS volume could only be
// attached to a single node at any given time.
var volumeCaps = []csi.VolumeCapability_AccessMode{
	{
		Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
	},
}

var controllerCaps = []csi.ControllerServiceCapability_RPC_Type{
	csi.ControllerServiceCapability_RPC_CREATE_DELETE_VOLUME,
	csi.ControllerServiceCapability_RPC_PUBLISH_UNPUBLISH_VOLUME,
}

type controllerService struct {
	attachDescribeTimes int
	cloud               Cloud
	inFlight            *internal.InFlight
}

func newControllerService() controllerService {
	pov, err := newCloud()
	if err != nil {
		log.Fatalf("newControllerService: init cloud err: %v", err)
	}
	return controllerService{
		cloud:               pov,
		inFlight:            internal.NewInFlight(),
		attachDescribeTimes: 12,
	}
}

func (d *controllerService) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Infof("CreateVolume: called args: %+v", *req)
	if err := validateCreateVolumeRequest(req); err != nil {
		return nil, err
	}

	volName := req.GetName()

	// check if a request is already in-flight
	if ok := d.inFlight.Insert(volName); !ok {
		msg := fmt.Sprintf("Create volume request for %s is already in progress", volName)
		return nil, status.Error(codes.Aborted, msg)
	}
	defer d.inFlight.Delete(volName)

	var (
		zoneID             string
		dataRedundancyType string
		protocolType       string
		storageType        string
		fileSystemName     string
		spaceCapacity      int64
		filesystemID       string

		// throughputMode must be binded with provisionedThroughputInMiBps
		throughputMode               string
		provisionedThroughputInMiBps string

		// volumeTags = map[string]string{}
	)
	pvProps := new(template.PVProps)

	for key, value := range req.GetParameters() {
		switch strings.ToLower(key) {
		case ZONEID:
			zoneID = value
		case DATAREDUNDANCYTYPE:
			dataRedundancyType = value
		case PROTOCOLTYPE:
			protocolType = value
		case STORAGETYPE:
			storageType = value
		case FILESYSTEMNAME:
			fileSystemName = value
		case THROUGHPUTMODE:
			throughputMode = value
		case PROVISIONEDTHROUGHPUTINMIBPS:
			provisionedThroughputInMiBps = value
		case PVCNameKey:
			pvProps.PVCName = value
		case PVCNamespaceKey:
			pvProps.PVCNamespace = value
		case PVNameKey:
			pvProps.PVName = value
		case FILESYSTEMID:
			filesystemID = value
			// default:
			// 	return nil, status.Errorf(codes.InvalidArgument, "Invalid parameter key %s for CreateVolume", key)
		}
	}

	if filesystemID != "" {
		return d.getMountPointWithFileSystemID(ctx, req, filesystemID)
	}

	volSizeBytes, err := getVolSizeBytes(req)
	if err != nil {
		return nil, err
	}
	spaceCapacity = utils.Bytes2GiB(volSizeBytes)

	opts := PovOptions{
		zoneID:                       zoneID,
		dataRedundancyType:           dataRedundancyType,
		protocolType:                 protocolType,
		storageType:                  storageType,
		capacity:                     spaceCapacity,
		fsName:                       fileSystemName,
		throughputMode:               throughputMode,
		provisionedThroughputInmibps: provisionedThroughputInMiBps,
	}

	fsId, requestID, err := d.cloud.CreateVolume(ctx, volName, &opts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Create pov volume failed, volName: %s, requestID: %s err: %+v", volName, requestID, err)
	}

	return d.getMountPointWithFileSystemID(ctx, req, fsId)
}

func (d *controllerService) getMountPointWithFileSystemID(ctx context.Context, req *csi.CreateVolumeRequest, fsId string) (*csi.CreateVolumeResponse, error) {
	var err error
	// one filesystem only has one mountpoint by default, one filesystem can create 10 mountpoints at most
	mpId := d.getDefaultMountPoint(ctx, fsId)
	if mpId == "" {
		mpId, err = d.cloud.CreateVolumeMountPoint(ctx, fsId)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "create mountpoint failed, fsId: %s err: %v", fsId, err)
		}
	}

	allParams := addAdditionalParams(req.GetParameters(), fsId, mpId)
	return newCreateVolumeResponse(mpId, 0, allParams), nil
}

func (d *controllerService) getDefaultMountPoint(ctx context.Context, fsId string) string {

	vmp, err := d.cloud.DescribeVscMountPoints(ctx, fsId, "")
	if err != nil {
		log.Errorf("getDefaultMountPoint: describe vsc mountpoint failed, fsId: %s err: %v", fsId, err)
		return ""
	}
	if len(vmp.MountPoints) == 0 {
		log.Infof("getDefaultMountPoint: get empty mountpoint by fsId: %s", fsId)
		return ""
	}
	return vmp.MountPoints[0].MountPointId
}

func addAdditionalParams(params map[string]string, fsId string, mpId string) map[string]string {
	params[labelAppendPrefix+labelFilesystemID] = fsId
	params[labelAppendPrefix+labelmpId] = mpId
	return params
}

func validateCreateVolumeRequest(req *csi.CreateVolumeRequest) error {
	volCaps := req.GetVolumeCapabilities()
	if !isValidVolumeCapabilities(volCaps) {
		modes := utils.GetAccessModes(volCaps)
		stringModes := strings.Join(*modes, ", ")
		errString := "Volume capabilities " + stringModes + " not supported. Only AccessModes[ReadWriteMany] supported."
		return status.Error(codes.InvalidArgument, errString)
	}
	return nil
}

func isValidVolumeCapabilities(volCaps []*csi.VolumeCapability) bool {
	hasSupport := func(cap *csi.VolumeCapability) bool {
		for _, c := range volumeCaps {
			if c.GetMode() == cap.AccessMode.GetMode() {
				return true
			}
		}
		return false
	}

	foundAll := true
	for _, c := range volCaps {
		if !hasSupport(c) {
			foundAll = false
		}
	}
	return foundAll
}

func getVolSizeBytes(req *csi.CreateVolumeRequest) (int64, error) {
	var volSizeBytes int64
	capRange := req.GetCapacityRange()
	if capRange == nil {
		volSizeBytes = defaultVolumeSize
	} else {
		volSizeBytes = utils.RoundUpBytes(capRange.GetRequiredBytes())
		maxVolSize := capRange.GetLimitBytes()
		if maxVolSize > 0 && maxVolSize < volSizeBytes {
			return 0, status.Error(codes.InvalidArgument, "After round-up, volume size exceeds the limit specified")
		}
	}
	return volSizeBytes, nil
}

func (d *controllerService) DeleteVolume(ctx context.Context, req *csi.DeleteVolumeRequest) (*csi.DeleteVolumeResponse, error) {
	return &csi.DeleteVolumeResponse{}, nil
}

func newCreateVolumeResponse(mpId string, blockSize int64, params map[string]string) *csi.CreateVolumeResponse {

	segments := map[string]string{TopologyKey: GlobalConfigVar.regionID}

	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      mpId,
			CapacityBytes: blockSize,
			VolumeContext: params,
			AccessibleTopology: []*csi.Topology{
				{
					Segments: segments,
				},
			},
		},
	}
}

func (d *controllerService) ControllerGetCapabilities(ctx context.Context, req *csi.ControllerGetCapabilitiesRequest) (*csi.ControllerGetCapabilitiesResponse, error) {
	var caps []*csi.ControllerServiceCapability
	for _, cap := range controllerCaps {
		c := &csi.ControllerServiceCapability{
			Type: &csi.ControllerServiceCapability_Rpc{
				Rpc: &csi.ControllerServiceCapability_RPC{
					Type: cap,
				},
			},
		}
		caps = append(caps, c)
	}
	return &csi.ControllerGetCapabilitiesResponse{Capabilities: caps}, nil
}

func (d *controllerService) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	pvs, err := getPVBympId(req.GetVolumeId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get pv by mpId: %s err: %v", req.GetVolumeId(), err)
	}
	if len(pvs) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get pv by mpId: %s", req.GetVolumeId())
	}

	var fsId string
	if value, ok := pvs[0].Labels[labelFilesystemID]; ok {
		log.Infof("ControllerPublishVolume: pv:[%s]'s fsId: %s", pvs[0].Name, value)
		fsId = value
	} else {
		return nil, status.Errorf(codes.InvalidArgument, "failed to get fsId by mpId: %s, pvs: %+v", req.GetVolumeId(), pvs)
	}

	requestID, err := d.cloud.AttachVscMountPoint(ctx, req.GetVolumeId(), fsId, req.GetNodeId())
	if err != nil {
		return nil, err
	}
	log.Infof("ControllerPublishVolume: attach vsc volume: %s to node : %s success, requestid: %s", req.GetVolumeId(), req.GetNodeId(), requestID)

	var vscId string
VSCREADY:
	for i := 0; i < d.attachDescribeTimes && vscId == ""; i++ {
		time.Sleep(4 * time.Second)
		vmp, err := d.cloud.DescribeVscMountPoints(ctx, fsId, req.GetVolumeId())
		if err != nil {
			return nil, err
		}
		log.Infof("ControllerPublishVolume: describe vsc mountpoint success, volumeid: %s mountpointids: %v", req.GetVolumeId(), vmp.MountPoints)
		for _, mp := range vmp.MountPoints {
			for _, ins := range mp.Instances {
				if ins.InstanceId == req.GetNodeId() {
					if ins.Status == NORMAL.String() && ins.Vscs[0].VscStatus == NORMAL.String() {
						vscId = ins.Vscs[0].VscId
						break VSCREADY
					}
				} else {
					continue
				}
			}
		}
	}
	if vscId == "" {
		return nil, status.Errorf(codes.Internal, "can't get vscId by fsId: %s, mpId: %s, instanceId: %s", fsId, req.GetVolumeId(), req.GetNodeId())
	}
	log.Infof("ControllerPublishVolume: patch pv vsc vscid: %v mpId: %v", vscId, req.GetVolumeId())
	if err := patchVscId2PV(fsId, vscId, req.GetVolumeId(), req.GetNodeId()); err != nil {
		return nil, err
	}
	return &csi.ControllerPublishVolumeResponse{}, nil
}

func patchVscId2PV(fsId, vscId, mpId, nodeID string) error {
	pvs, err := getPVBympId(mpId)
	if err != nil {
		return err
	}
	if len(pvs) == 0 {
		log.Errorf("patchVscId2PV: get invalid pv count by pov mpid labels, pv count: %d", len(pvs))
	}
	for _, pv := range pvs {
		npv := pv.DeepCopy()
		// vscPrefix := fmt.Sprintf(annVSCIDPrefix, nodeID)
		if value, ok := npv.Annotations[annVSCIDPrefix]; ok && value == vscId {
			continue
		}
		npv.Annotations[annVSCIDPrefix] = vscId
		_, err = GlobalConfigVar.client.CoreV1().PersistentVolumes().Update(context.TODO(), npv, metav1.UpdateOptions{})
		if err != nil {
			return err
		}
	}
	log.Info("patchVscId2PV: Successfully updated vscid to pv's labels")
	return nil
}

func getPVBympId(mpId string) ([]v1.PersistentVolume, error) {
	ctx := context.TODO()
	labelPvs := labels.SelectorFromSet(labels.Set(map[string]string{labelmpId: mpId}))
	listPvOpts := metav1.ListOptions{
		LabelSelector: labelPvs.String(),
	}

	pvs, err := GlobalConfigVar.client.CoreV1().PersistentVolumes().List(ctx, listPvOpts)
	if err != nil {
		return []v1.PersistentVolume{}, err
	}
	return pvs.Items, nil
}

func (d *controllerService) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {

	if os.Getenv("ENABLE_SINGLE_TENANT") == "true" {
		return &csi.ControllerUnpublishVolumeResponse{}, nil
	}

	mpId := req.GetVolumeId()
	pvs, err := getPVBympId(mpId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "get pv by mpId: %s failed", mpId)
	}
	if len(pvs) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid pv counts by mpId: %s", mpId)
	}

	fsId, ok := pvs[0].Labels[labelFilesystemID]
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "invalid pv labels: %v by mpId: %s", pvs[0].Labels, mpId)
	}

	requestID, err := d.cloud.DetachVscMountPoint(ctx, mpId, fsId, req.GetNodeId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "detach volumepoint: %s failed err: %v", req.GetVolumeId(), err)
	}
	log.Infof("ControllerUnpublishVolume: detach pov mountpoint: %s, fsId: %s success, reqid: %s", req.GetVolumeId(), fsId, requestID)

	return &csi.ControllerUnpublishVolumeResponse{}, nil
}

func (d *controllerService) ValidateVolumeCapabilities(ctx context.Context, req *csi.ValidateVolumeCapabilitiesRequest) (*csi.ValidateVolumeCapabilitiesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) CreateSnapshot(ctx context.Context, req *csi.CreateSnapshotRequest) (*csi.CreateSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) GetCapacity(ctx context.Context, req *csi.GetCapacityRequest) (*csi.GetCapacityResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) DeleteSnapshot(ctx context.Context, req *csi.DeleteSnapshotRequest) (*csi.DeleteSnapshotResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ListSnapshots(ctx context.Context, req *csi.ListSnapshotsRequest) (*csi.ListSnapshotsResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ListVolumes(ctx context.Context, req *csi.ListVolumesRequest) (*csi.ListVolumesResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}
