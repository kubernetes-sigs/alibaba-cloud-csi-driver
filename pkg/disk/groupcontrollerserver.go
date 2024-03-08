package disk

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes/timestamp"
	groupVolumeSnasphotV1alpha1 "github.com/kubernetes-csi/external-snapshotter/client/v7/apis/volumegroupsnapshot/v1alpha1"
	snapClientset "github.com/kubernetes-csi/external-snapshotter/client/v7/clientset/versioned"
	csicommon "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/agent/csi-common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/crds"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/api/core/v1"
	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	crd "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/tools/record"
)

const (
	MaxCapacityInGiB = 32 << 10
)

var groupControllerCaps = []csi.GroupControllerServiceCapability_RPC_Type{
	csi.GroupControllerServiceCapability_RPC_CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT,
}

// groupcontroller server try to create/delete group snapshots
type groupControllerServer struct {
	*csicommon.DefaultGroupControllerServer
	recorder record.EventRecorder
}

// NewGroupControllerServer is to create controller server
func NewGroupControllerServer(d *csicommon.CSIDriver, client *crd.Clientset) csi.GroupControllerServer {
	installCRD := true
	installCRDStr := os.Getenv(utils.InstallSnapshotCRD)
	if installCRDStr == "false" {
		installCRD = false
	}

	// parse input snapshot request interval
	intervalStr := os.Getenv(SnapshotRequestTag)
	if intervalStr != "" {
		interval, err := strconv.ParseInt(intervalStr, 10, 64)
		if err != nil {
			log.Fatalf("Input SnapshotRequestTag is illegal: %s", intervalStr)
		}
		SnapshotRequestInterval = interval
	}

	serviceType := os.Getenv(utils.ServiceType)
	if serviceType == utils.ProvisionerService && installCRD {
		checkInstallVolumeGroupSnapshotCRD(client)
		checkInstallDefaultVolumeGroupSnapshotClass(GlobalConfigVar.SnapClient)
	}
	c := &groupControllerServer{
		DefaultGroupControllerServer: csicommon.NewDefaultGroupControllerServer(d),
		recorder:                     utils.NewEventRecorder(),
	}
	return c
}

// the map of req.Name and csi.VolumeGroupSnapshot
var createdGroupSnapshotMap = map[string]*csi.VolumeGroupSnapshot{}

// GroupSnapshotRequestMap snapshot request limit
var GroupSnapshotRequestMap = map[string]int64{}

// GroupSnapshotRequestInterval snapshot request limit
var GroupSnapshotRequestInterval = int64(10)

func (cs *groupControllerServer) GroupControllerGetCapabilities(ctx context.Context, req *csi.GroupControllerGetCapabilitiesRequest) (*csi.GroupControllerGetCapabilitiesResponse, error) {
	var caps []*csi.GroupControllerServiceCapability
	for _, cap := range groupControllerCaps {
		c := &csi.GroupControllerServiceCapability{
			Type: &csi.GroupControllerServiceCapability_Rpc{
				Rpc: &csi.GroupControllerServiceCapability_RPC{
					Type: cap,
				},
			},
		}
		caps = append(caps, c)
	}
	return &csi.GroupControllerGetCapabilitiesResponse{Capabilities: caps}, nil
}

func (cs *groupControllerServer) CreateVolumeGroupSnapshot(ctx context.Context, req *csi.CreateVolumeGroupSnapshotRequest) (*csi.CreateVolumeGroupSnapshotResponse, error) {
	// request limit
	cur := time.Now().Unix()
	if initTime, ok := SnapshotRequestMap[req.Name]; ok {
		if cur-initTime < SnapshotRequestInterval {
			err := fmt.Errorf("CreateVolumeGroupSnapshot: volume group snapshot create request limit %s", req.Name)
			return nil, err
		}
	}
	SnapshotRequestMap[req.Name] = cur

	// used for snapshot events
	groupSnapshotName := req.Parameters[common.VolumeGroupSnapshotNameKey]
	groupSnapshotNamespace := req.Parameters[common.VolumeGroupSnapshotNamespaceKey]

	ref := &v1.ObjectReference{
		Kind:      "VolumeGroupSnapshot",
		Name:      groupSnapshotName,
		UID:       "",
		Namespace: groupSnapshotNamespace,
	}

	params, err := getVolumeGroupSnapshotConfig(req)
	if err != nil {
		return nil, err
	}

	log.Infof("CreateVolumeGroupSnapshot:: Starting to create volumegroupsnapshot: %+v", req)
	if err := cs.Driver.ValidateGroupControllerServiceRequest(csi.GroupControllerServiceCapability_RPC_CREATE_DELETE_GET_VOLUME_GROUP_SNAPSHOT); err != nil {
		log.Errorf("CreateVolumeGroupSnapshot: driver not support Create volume group snapshot: %v", err)
		return nil, err
	}

	sourceVolumeIds := req.GetSourceVolumeIds()
	for index, id := range sourceVolumeIds {
		sourceVolumeIds[index] = strings.TrimSpace(id)
	}
	// Need to check for already existing snapshot name
	GlobalConfigVar.EcsClient = updateEcsClient(GlobalConfigVar.EcsClient)
	// check if groupvolumesnapshot has already exists
	groupSnapshots, snapNum, err := findGroupSnapshot(req.GetName(), "")
	if snapNum == 0 {
		groupSnapshots, snapNum, err = findGroupSnapshot("", req.GetName())
	}
	switch {
	case snapNum == 1:
		// Since err is nil, it means the snapshot with the same name already exists need
		// to check if the sourceVolumeId of existing snapshot is the same as in new request.
		existsGroupSnapshot := groupSnapshots.SnapshotGroups.SnapshotGroup[0]
		match := ifExistsGroupSnapshotMatch(&existsGroupSnapshot, sourceVolumeIds)
		if !match {
			log.Errorf("CreateVolumeGroupSnapshot:: GroupSnapshot already exist with same name: name[%s], volumeIDs[%v]", req.Name, sourceVolumeIds)
			err := status.Errorf(codes.AlreadyExists, "groupSnapshot with the same name: %s but with different SourceVolumeIds already exist", req.GetName())
			utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotAlreadyExist, err.Error())
			return nil, err
		}
		groupSnapshot, err := formatGroupSnapshot(&existsGroupSnapshot)
		if err != nil {
			return nil, err
		}
		log.Infof("CreateVolumeGroupSnapshot:: GroupSnapshot already created: name[%s], sourceIds[%v], status[%v]", req.Name, sourceVolumeIds, groupSnapshot.ReadyToUse)
		if groupSnapshot.ReadyToUse {
			utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully, fmt.Sprintf("VolumeGroupSnapshot: name: %s, id: %s is ready to use.", req.GetName(), groupSnapshot.GroupSnapshotId))
			delete(GroupSnapshotRequestMap, req.GetName())
		}
		return &csi.CreateVolumeGroupSnapshotResponse{
			GroupSnapshot: groupSnapshot,
		}, nil
	case snapNum > 1:
		log.Errorf("CreateVolumeGroupSnapshot:: Find groupSnapshot name[%s], but get more than 1 instance", req.GetName())
		err := status.Error(codes.Internal, "CreateVolumeGroupSnapshot: get snapshot more than 1 instance")
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotTooMany, err.Error())
		return nil, err
	case err != nil:
		log.Errorf("CreateVolumeGroupSnapshot:: Expect to find groupSnapshot name[%s], but get error: %v", req.GetName(), err)
		e := status.Errorf(codes.Internal, "CreateVolumeGroupSnapshot: get groupSnapshot with error: %s", err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
		return nil, e
	}

	// check groupsnapshot again, if ram has no auth to describe groupsnapshot, there will always 0 response.
	if value, ok := createdGroupSnapshotMap[req.Name]; ok {
		str := fmt.Sprintf("CreateVolumeGroupSnapshot:: groupSnapshot already created, Name: %s, Info: %v", req.Name, value)
		log.Info(str)
		return &csi.CreateVolumeGroupSnapshotResponse{
			GroupSnapshot: value,
		}, nil
	}

	// Check the status, AZ, type, whether it is multi-mounted of each disk, and the total capacity
	err = checkSourceVolumes(sourceVolumeIds)
	if err != nil {
		log.Errorf("CreateVolumeGroupSnapshot:: source volumes do not meet the requirements for creating group snapshot, err: %v", err)
		e := status.Errorf(codes.Internal, "CreateVolumeGroupSnapshot:: source volumes do not meet the requirements for creating group snapshot: %s", err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, e.Error())
	}

	// init createSnapshotGroupRequest and parameters
	params.SourceVolumeIDs = sourceVolumeIds
	params.SnapshotName = req.GetName()
	snapshotResponse, err := requestAndCreateSnapshotGroup(GlobalConfigVar.EcsClient, params)

	if err != nil {
		log.Errorf("CreateVolumeGroupSnapshot:: create groupSnapshot failed: snapshotName[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, err
	}

	ecsGroupSnapshots, snapNum, err := findGroupSnapshot(req.GetName(), snapshotResponse.SnapshotGroupId)
	if err != nil || snapNum == 0 {
		log.Errorf("CreateVolumeGroupSnapshot:: find created groupSnapshot failed or not found: snapshotName[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, err
	}
	groupSnapshot, err := formatGroupSnapshot(&ecsGroupSnapshots.SnapshotGroups.SnapshotGroup[0])
	if err != nil {
		log.Errorf("CreateVolumeGroupSnapshot:: format created groupSnapshot failed: snapshotName[%s], sourceIds[%v], error[%s]", req.GetName(), sourceVolumeIds, err.Error())
		utils.CreateEvent(cs.recorder, ref, v1.EventTypeWarning, snapshotCreateError, err.Error())
		return nil, err
	}

	groupSnapshot.ReadyToUse = true

	// always IA snapshot, so set readyToUse true immediately
	delete(GroupSnapshotRequestMap, req.GetName())
	str := fmt.Sprintf("CreateVolumeGroupSnapshot:: GroupSnapshot create successful: snapshotName[%s], sourceIds[%v], snapshotId[%s]", req.GetName(), sourceVolumeIds, snapshotResponse.SnapshotGroupId)
	log.Infof(str)
	utils.CreateEvent(cs.recorder, ref, v1.EventTypeNormal, snapshotCreatedSuccessfully, str)
	return &csi.CreateVolumeGroupSnapshotResponse{
		GroupSnapshot: groupSnapshot,
	}, nil
}

func (cs *groupControllerServer) DeleteVolumeGroupSnapshot(ctx context.Context, req *csi.DeleteVolumeGroupSnapshotRequest) (*csi.DeleteVolumeGroupSnapshotResponse, error) {
	return nil, nil
}
func (cs *groupControllerServer) GetVolumeGroupSnapshot(ctx context.Context, req *csi.GetVolumeGroupSnapshotRequest) (*csi.GetVolumeGroupSnapshotResponse, error) {
	return nil, nil
}

func getVolumeGroupSnapshotConfig(req *csi.CreateVolumeGroupSnapshotRequest) (*createGroupSnapshotParams, error) {
	var ecsParams createGroupSnapshotParams
	if req.Parameters != nil {
		err := parseGroupSnapshotParameters(req.Parameters, &ecsParams)
		if err != nil {
			log.Errorf("CreateSnapshot:: Snapshot name[%s], parse config failed: %v", req.Name, err)
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
	}

	vsName := req.Parameters[common.VolumeGroupSnapshotNameKey]
	vsNameSpace := req.Parameters[common.VolumeGroupSnapshotNamespaceKey]
	// volumesnapshot not in parameters, just retrun
	if vsName == "" || vsNameSpace == "" {
		return &ecsParams, nil
	}

	volumeGroupSnapshot, err := GlobalConfigVar.SnapClient.GroupsnapshotV1alpha1().VolumeGroupSnapshots(vsNameSpace).Get(context.Background(), vsName, metav1.GetOptions{})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get VolumeGroupSnapshot: %s/%s: %v", vsNameSpace, vsName, err)
	}
	err = parseGroupSnapshotAnnotations(volumeGroupSnapshot.Annotations, &ecsParams)
	if err != nil {
		log.Errorf("CreateVolumeGroupSnapshot:: Snapshot name[%s], parse annotation failed: %v", req.Name, err)
		return nil, status.Errorf(codes.InvalidArgument, err.Error())
	}
	return &ecsParams, nil
}

func parseGroupSnapshotParameters(params map[string]string, ecsParams *createGroupSnapshotParams) (err error) {
	for k, v := range params {
		switch k {
		case SNAPSHOTTYPE:
			// use instanctAccess by default
		case RETENTIONDAYS:
			//ecsParams.RetentionDays, err = strconv.Atoi(v)
			//if err != nil {
			//	return fmt.Errorf("failed to parse retentionDays: %w", err)
			//}
			return fmt.Errorf("groupSnapshot do not support retentionDays: %w", err)
		case INSTANTACCESSRETENTIONDAYS:
			// instantAccess RetentionDays always equals to RetentionDays
		case SNAPSHOTRESOURCEGROUPID:
			ecsParams.ResourceGroupID = v
		case common.VolumeGroupSnapshotNameKey:
			ecsParams.SnapshotTags = append(ecsParams.SnapshotTags, ecs.CreateSnapshotGroupTag{
				Key:   common.VolumeGroupSnapshotNameTag,
				Value: v,
			})
		case common.VolumeGroupSnapshotNamespaceKey:
			ecsParams.SnapshotTags = append(ecsParams.SnapshotTags, ecs.CreateSnapshotGroupTag{
				Key:   common.VolumeGroupSnapshotNamespaceTag,
				Value: v,
			})
		default:
			if strings.HasPrefix(k, SNAPSHOT_TAG_PREFIX) {
				ecsParams.SnapshotTags = append(ecsParams.SnapshotTags, ecs.CreateSnapshotGroupTag{
					Key:   k[len(SNAPSHOT_TAG_PREFIX):],
					Value: v,
				})
			}
		}
	}
	return nil
}

// if volumesnapshot have Annotations, use it first.
// storage.alibabacloud.com/snapshot-ttl
func parseGroupSnapshotAnnotations(anno map[string]string, ecsParams *createGroupSnapshotParams) error {
	snapshotTTL := anno["storage.alibabacloud.com/snapshot-ttl"]

	if snapshotTTL != "" {
		//ecsParams.RetentionDays, err = strconv.Atoi(snapshotTTL)
		//if err != nil {
		//	return fmt.Errorf("failed to parse annotation snapshot-ttl: %w", err)
		//}
		return fmt.Errorf("groupSnapshot do not support retentionDays")
	}

	return nil
}

func formatGroupSnapshot(groupSnapshot *ecs.SnapshotGroup) (*csi.VolumeGroupSnapshot, error) {
	t, err := time.Parse(time.RFC3339, groupSnapshot.CreationTime)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to parse groupSnapshot creation time: %s", groupSnapshot.CreationTime)
	}

	readyToUse := true
	snapshots := []*csi.Snapshot{}
	for _, ecsSnapshot := range groupSnapshot.Snapshots.Snapshot {
		snapshot, err := formatCSISnapshot(&ecsSnapshot)
		if err != nil {
			return nil, err
		}
		if !snapshot.ReadyToUse && !ecsSnapshot.InstantAccess {
			// set readyToUse for groupsnapshots according to each snapshot status
			readyToUse = false
		}
		snapshots = append(snapshots, snapshot)
	}

	return &csi.VolumeGroupSnapshot{
		GroupSnapshotId: groupSnapshot.SnapshotGroupId,
		Snapshots:       snapshots,
		CreationTime:    &timestamp.Timestamp{Seconds: t.Unix()},
		ReadyToUse:      readyToUse,
	}, nil
}

func ifExistsGroupSnapshotMatch(existsGroupSnapshot *ecs.SnapshotGroup, sourceVolumeIds []string) bool {
	existsSnapshotSourceIdsMap := make(map[string]struct{})
	for _, snapshot := range existsGroupSnapshot.Snapshots.Snapshot {
		existsSnapshotSourceIdsMap[snapshot.SourceDiskId] = struct{}{}
	}
	matach := true
	for _, id := range sourceVolumeIds {
		if _, ok := existsSnapshotSourceIdsMap[id]; !ok {
			matach = false
			break
		}
	}
	return matach
}

func checkSourceVolumes(sourceVolumeIds []string) error {
	var zoneId string
	var capacityInGiB int
	for _, sourceVolumeId := range sourceVolumeIds {
		disks := getDisk(sourceVolumeId, GlobalConfigVar.EcsClient)
		if len(disks) == 0 {
			return fmt.Errorf("no disk found: %s", sourceVolumeId)
		} else if len(disks) != 1 {
			return fmt.Errorf("multi disk found: %s", sourceVolumeId)
		}

		// TODO: skip status check as CreateSnapshot

		switch disks[0].Category {
		case DiskESSD, DiskESSDAuto, DiskESSDEntry:
			// do nothing
		default:
			return fmt.Errorf("groupSnapshot only support ESSD disks, but disk %s is %s: %v", sourceVolumeId, disks[0].Category)
		}

		if zoneId == "" {
			zoneId = disks[0].ZoneId
		} else if zoneId != disks[0].ZoneId {
			return fmt.Errorf("groupSnapshot only support disks from one zone, but get disks from %s and %s", zoneId, disks[0].ZoneId)
		}

		if disks[0].MultiAttach != "Disabled" {
			return fmt.Errorf("groupSnapshot do not support multiAttached disk, but disk %s enables multiAttached", sourceVolumeId)
		}

		capacityInGiB += disks[0].Size
		if capacityInGiB > MaxCapacityInGiB {
			return fmt.Errorf("total capacity is larger than %d", MaxCapacityInGiB)
		}
	}
	return nil
}

func checkInstallVolumeGroupSnapshotCRD(crdClient *crd.Clientset) {

	snapshotCRDNames := map[string]string{
		"volumegroupsnapshotclasses.groupsnapshot.storage.k8s.io":  "GetVolumeGroupSnapshotClassesCRDv1",
		"volumegroupsnapshotcontents.groupsnapshot.storage.k8s.io": "GetVolumeGroupSnapshotContentsCRDv1",
		"volumegroupsnapshots.groupsnapshot.storage.k8s.io":        "GetVolumeGroupSnapshotsCRDv1",
	}

	ctx := context.Background()
	listOpts := metav1.ListOptions{}
	crdList, err := crdClient.ApiextensionsV1().CustomResourceDefinitions().List(ctx, listOpts)
	if err != nil {
		log.Errorf("checkInstallVolumeGroupSnapshotCRD:: list CustomResourceDefinitions error: %v", err)
		return
	}
	for _, crd := range crdList.Items {
		delete(snapshotCRDNames, crd.Name)
		if len(snapshotCRDNames) == 0 {
			return
		}
	}
	temp := &crds.Template{}
	info, err := GlobalConfigVar.ClientSet.ServerVersion()
	kVersion := ""
	if err != nil || info == nil {
		log.Errorf("checkInstallVolumeGroupSnapshotCRD: get server version error : %v", err)
		kVersion = "v1.28.3-aliyun.1"
	} else {
		kVersion = info.GitVersion
	}
	log.Infof("checkInstallVolumeGroupSnapshotCRD: need to create crd counts: %v", len(snapshotCRDNames))
	for _, value := range snapshotCRDNames {
		crdStrings := reflect.ValueOf(temp).MethodByName(value).Call([]reflect.Value{reflect.ValueOf(kVersion)})
		crdToBeCreated := crdv1.CustomResourceDefinition{}
		yamlString := crdStrings[0].Interface().(string)
		crdDecoder := yaml.NewYAMLOrJSONDecoder(bytes.NewReader([]byte(yamlString)), 4096)
		err := crdDecoder.Decode(&crdToBeCreated)
		if err != nil {
			log.Errorf("checkInstallVolumeGroupSnapshotCRD: yaml unmarshal error: %v", err)
			return
		}
		force := true
		yamlBytes := []byte(yamlString)
		_, err = crdClient.ApiextensionsV1().CustomResourceDefinitions().Patch(ctx, crdToBeCreated.Name, types.ApplyPatchType, yamlBytes, metav1.PatchOptions{
			Force:        &force,
			FieldManager: "alibaba-cloud-csi-driver",
		})
		if err != nil {
			log.Infof("checkInstallVolumeGroupSnapshotCRD: crd apply error: %v", err)
			return
		}
	}
}

func checkInstallDefaultVolumeGroupSnapshotClass(snapClient *snapClientset.Clientset) {
	_, err := snapClient.GroupsnapshotV1alpha1().VolumeGroupSnapshotClasses().Get(context.TODO(), DefaultVolumeGroupSnapshotClass, metav1.GetOptions{})
	if err != nil {
		groupSnapshotClass := &groupVolumeSnasphotV1alpha1.VolumeGroupSnapshotClass{
			ObjectMeta: metav1.ObjectMeta{
				Name: DefaultVolumeGroupSnapshotClass,
			},
			Driver:         driverName,
			DeletionPolicy: "Delete",
			Parameters:     map[string]string{},
		}
		_, err = snapClient.GroupsnapshotV1alpha1().VolumeGroupSnapshotClasses().Create(context.TODO(), groupSnapshotClass, metav1.CreateOptions{})
		if err != nil {
			log.Errorf("checkInstallDefaultVolumeSnapshotClass:: failed to create volume group snapshot class: %v", err)
		}
	}
}
