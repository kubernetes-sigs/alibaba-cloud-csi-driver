package pov

import (
	"context"
	"fmt"
	"strings"
	"time"

	csi "github.com/container-storage-interface/spec/lib/go/csi"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/log"
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
	// PVCNameKey contains name of the PVC for which is a volume provisioned.
	PVCNameKey = "csi.storage.k8s.io/pvc/name"
	// PVCNamespaceKey contains namespace of the PVC for which is a volume provisioned.
	PVCNamespaceKey = "csi.storage.k8s.io/pvc/namespace"
	// PVNameKey contains name of the final PV that will be used for the dynamically
	// provisioned volume
	PVNameKey   = "csi.storage.k8s.io/pv/name"
	TopologyKey = "topology." + DriverName + "/zone"

	// volumeContext starting with labelAppendPrefix will automatically added to pv lables
	labelAppendPrefix = "csi.alibabacloud.com/label-prefix/"
	annVSCIDPrefix    = "csi.alibabacloud.com/ann-vsc-id"
	annMPIDPrefix     = "csi.alibabacloud.com/ann-mp-id"
	labelFilesystemID = "csi.alibabacloud.com/label-pov-fsid"
)

// volumeCaps represents how the volume could be accessed.
// It is SINGLE_NODE_WRITER since EBS volume could only be
// attached to a single node at any given time.
var volumeCaps = []csi.VolumeCapability_AccessMode{
	{
		Mode: csi.VolumeCapability_AccessMode_MULTI_NODE_MULTI_WRITER,
	},
}

type controllerService struct {
	attachDescribeTimes int
	cloud               Cloud
	inFlight            *internal.InFlight
}

func newControllerService() controllerService {
	pov, err := newCloud()
	if err != nil {
		log.Log.Fatalf("newControllerService: init cloud err: %v", err)
	}
	return controllerService{
		cloud:    pov,
		inFlight: internal.NewInFlight(),
	}
}

func (d *controllerService) CreateVolume(ctx context.Context, req *csi.CreateVolumeRequest) (*csi.CreateVolumeResponse, error) {
	log.Log.Infof("CreateVolume: called args: %+v", *req)
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
		default:
			return nil, status.Errorf(codes.InvalidArgument, "Invalid parameter key %s for CreateVolume", key)
		}
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

	fsID, requestID, err := d.cloud.CreateVolume(ctx, volName, &opts)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Create pov volume failed, volName: %s, requestID: %s err: %+v", volName, requestID, err)
	}

	allParams := addAdditionalParams(req.GetParameters(), fsID)

	return newCreateVolumeResponse(fsID, zoneID, volSizeBytes, allParams), nil
}

func addAdditionalParams(params map[string]string, fsID string) map[string]string {
	params[labelAppendPrefix+labelFilesystemID] = fsID
	return params
}

func validateCreateVolumeRequest(req *csi.CreateVolumeRequest) error {
	volName := req.GetName()
	if len(volName) == 0 {
		return status.Error(codes.InvalidArgument, "Volume name not provided")
	}

	volCaps := req.GetVolumeCapabilities()
	if len(volCaps) == 0 {
		return status.Error(codes.InvalidArgument, "Volume capabilities not provided")
	}

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

func newCreateVolumeResponse(fsID, zoneID string, blockSize int64, params map[string]string) *csi.CreateVolumeResponse {

	segments := map[string]string{TopologyKey: zoneID}

	return &csi.CreateVolumeResponse{
		Volume: &csi.Volume{
			VolumeId:      fsID,
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
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ControllerExpandVolume(ctx context.Context, req *csi.ControllerExpandVolumeRequest) (*csi.ControllerExpandVolumeResponse, error) {
	return nil, status.Error(codes.Unimplemented, "")
}

func (d *controllerService) ControllerPublishVolume(ctx context.Context, req *csi.ControllerPublishVolumeRequest) (*csi.ControllerPublishVolumeResponse, error) {
	if err := validateControllerPublishVolumeRequest(req); err != nil {
		return nil, err
	}

	vmp, err := d.cloud.DescribeVscMountPoints(ctx, req.GetVolumeId(), "")
	if err != nil {
		return nil, status.Error(codes.Internal, fmt.Sprintf("failed to describe fsid, err: %+v", err))
	}
	pvs, err := getPVByFsId(req.GetVolumeId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to get pv by fsid: %s err: %v", req.GetVolumeId(), err))
	}
	if len(pvs) == 0 {
		return nil, status.Error(codes.InvalidArgument, fmt.Sprintf("failed to get pv by fsid: %s", req.GetVolumeId()))
	}

	var mpID string
	if value, ok := pvs[0].Annotations[annMPIDPrefix]; ok {
		log.Log.Infof("ControllerPublishVolume: pv: %s already has mpid: %s", pvs[0].Name, value)
		mpID = value
	} else {
		if len(vmp.MountPoints) == 0 {
			mpID, err = d.cloud.CreateVolumeMountPoint(ctx, req.GetVolumeId())
			if err != nil {
				return nil, err
			}
			log.Log.Infof("ControllerPublishVolume: create vsc mountpoint success, volumeid: %s mountpointID: %s", req.GetVolumeId(), mpID)
		} else {
			// for test
			mpID = vmp.MountPoints[0].MountPointId
		}
	}

	requestID, err := d.cloud.AttachVscMountPoint(ctx, mpID, req.GetVolumeId(), req.GetNodeId())
	if err != nil {
		return nil, err
	}
	log.Log.Infof("ControllerPublishVolume: attach vsc volume: %s to node : %s success, requestid: %s", req.GetVolumeId(), req.GetNodeId(), requestID)

	var vscId string
	for i := 0; i < d.attachDescribeTimes; i++ {
		time.Sleep(1 * time.Second)
		vmp, err := d.cloud.DescribeVscMountPoints(ctx, req.GetVolumeId(), mpID)
		if err != nil {
			return nil, err
		}
		for _, mp := range vmp.MountPoints {
			for _, ins := range mp.Instances {
				if ins.InstanceId == req.GetNodeId() {
					if ins.Status == NORMAL.String() && ins.Vscs[0].VscStatus == NORMAL.String() {
						vscId = ins.Vscs[0].VscId
					}
				} else {
					continue
				}
			}
		}
	}
	if vscId == "" {
		return nil, status.Error(codes.Internal, fmt.Sprintf("can't get vscId by fsId: %s, mpId: %s, instanceId: %s", req.GetVolumeId(), mpID, req.GetNodeId()))
	}
	patchVscId2PV(req.GetVolumeId(), vscId, mpID)
	return &csi.ControllerPublishVolumeResponse{}, nil
}

func patchVscId2PV(fsID, vscId, mpID string) error {
	pvs, err := getPVByFsId(fsID)
	if err != nil {
		return err
	}
	if len(pvs) != 1 {
		log.Log.Errorf("patchVscId2PV: get invalid pv count by pov fs labels, pv count: %d", len(pvs))
	}
	npv := pvs[0].DeepCopy()
	npv.Annotations[annVSCIDPrefix] = vscId
	_, err = GlobalConfigVar.client.CoreV1().PersistentVolumes().Update(context.TODO(), npv, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	log.Log.Info("patchVscId2PV: Successfully updated vscid to pv's labels")
	return nil
}

func getPVByFsId(fsID string) ([]v1.PersistentVolume, error) {
	ctx := context.TODO()
	labelPvs := labels.SelectorFromSet(labels.Set(map[string]string{labelAppendPrefix + labelFilesystemID: fsID}))
	listPvOpts := metav1.ListOptions{
		LabelSelector: labelPvs.String(),
	}

	pvs, err := GlobalConfigVar.client.CoreV1().PersistentVolumes().List(ctx, listPvOpts)
	if err != nil {
		return []v1.PersistentVolume{}, err
	}
	return pvs.Items, nil
}

func validateControllerPublishVolumeRequest(req *csi.ControllerPublishVolumeRequest) error {
	if len(req.GetVolumeId()) == 0 {
		return status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	if len(req.GetNodeId()) == 0 {
		return status.Error(codes.InvalidArgument, "Node ID not provided")
	}

	volCap := req.GetVolumeCapability()
	if volCap == nil {
		return status.Error(codes.InvalidArgument, "Volume capability not provided")
	}

	return nil
}

func (d *controllerService) ControllerUnpublishVolume(ctx context.Context, req *csi.ControllerUnpublishVolumeRequest) (*csi.ControllerUnpublishVolumeResponse, error) {

	if len(req.GetVolumeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Volume ID not provided")
	}

	if len(req.GetNodeId()) == 0 {
		return nil, status.Error(codes.InvalidArgument, "Node ID not provided")
	}

	fsId := req.GetVolumeId()
	pvs, err := getPVByFsId(fsId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "get pv by fsid: %s failed", fsId)
	}
	if len(pvs) != 1 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid pv counts by fsid: %s", fsId)
	}

	var mpID string
	var exists bool
	if mpID, exists = pvs[0].Annotations[annMPIDPrefix]; !exists || mpID == "" {
		return nil, status.Errorf(codes.InvalidArgument, "invalid pv: %s, mpid annotation absent", pvs[0].Name)
	}

	requestID, err := d.cloud.DetachVscMountPoint(ctx, mpID, req.GetVolumeId(), req.GetNodeId())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "detach volume: %s failed err: %v", req.GetVolumeId(), err)
	}
	log.Log.Infof("ControllerUnpublishVolume: detach pov volume: %s success, reqid: %s", req.GetVolumeId(), requestID)

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
