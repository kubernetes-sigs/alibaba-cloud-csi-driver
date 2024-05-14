package disk

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes/timestamp"
	groupVolumeSnasphotV1alpha1 "github.com/kubernetes-csi/external-snapshotter/client/v7/apis/volumegroupsnapshot/v1alpha1"
	snapClientset "github.com/kubernetes-csi/external-snapshotter/client/v7/clientset/versioned"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/crds"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	crdv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	crd "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/yaml"
)

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
		case RETENTIONDAYS:
			//ecsParams.RetentionDays, err = strconv.Atoi(v)
			//if err != nil {
			//	return fmt.Errorf("failed to parse retentionDays: %w", err)
			//}
			return fmt.Errorf("groupSnapshot do not support retentionDays: %w", err)
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
	creationTime := timestamp.Timestamp{Seconds: t.Unix()}

	readyToUse := true
	snapshots := []*csi.Snapshot{}
	for _, ecsSnapshot := range groupSnapshot.Snapshots.Snapshot {
		snapshot, err := formatCSISnapshot(&ecsSnapshot)
		if err != nil {
			return nil, err
		}
		snapshot.CreationTime = &creationTime
		snapshot.GroupSnapshotId = groupSnapshot.SnapshotGroupId
		if !snapshot.ReadyToUse {
			// set readyToUse for groupsnapshots according to each snapshot status
			readyToUse = false
		}
		snapshots = append(snapshots, snapshot)
	}

	return &csi.VolumeGroupSnapshot{
		GroupSnapshotId: groupSnapshot.SnapshotGroupId,
		Snapshots:       snapshots,
		CreationTime:    &creationTime,
		ReadyToUse:      readyToUse,
	}, nil
}

func ifExistsGroupSnapshotMatchSourceVolume(existsGroupSnapshot *ecs.SnapshotGroup, sourceVolumeIds []string) bool {
	if len(existsGroupSnapshot.Snapshots.Snapshot) != len(sourceVolumeIds) {
		return false
	}
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

// ifExistsGroupSnapshotMatch checks whether there snapshotIDs in `snapshotIds`
// that correspond to all related snapshots in `existsGroupSnapshot`.
// If `requireExactMatch` is true, the function requires an exact match between
// the related snapshotIDs and `snapshotIds`, including the same number of elements.
// If `requireExactMatch` is false, it only checks if each related snapshotID
// can be found in the `snapshotIds`, that means actual snapshots in group can be a subset of `snapshotIds`
func ifExistsGroupSnapshotMatch(existsGroupSnapshot *ecs.SnapshotGroup, snapshotIds []string, requireExactMatch bool) bool {
	if requireExactMatch && len(snapshotIds) != len(existsGroupSnapshot.Snapshots.Snapshot) {
		return false
	}
	sourceSnapshotIdMap := make(map[string]struct{})
	for _, id := range snapshotIds {
		sourceSnapshotIdMap[id] = struct{}{}
	}
	matach := true
	for _, snapshot := range existsGroupSnapshot.Snapshots.Snapshot {
		if _, ok := sourceSnapshotIdMap[snapshot.SnapshotId]; !ok {
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
		switch len(disks) {
		case 0:
			return fmt.Errorf("no disk found: %s", sourceVolumeId)
		case 1:
		default:
			return fmt.Errorf("multi disk found: %s", sourceVolumeId)
		}

		switch disks[0].Category {
		case DiskESSD, DiskESSDAuto, DiskESSDEntry:
		default:
			return fmt.Errorf("groupSnapshot only support ESSD disks, but disk %s is %v", sourceVolumeId, disks[0].Category)
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

func requestAndDeleteGroupSnapshot(groupSnapshotID string) (*ecs.DeleteSnapshotGroupResponse, error) {
	// Delete Snapshotgroup
	deleteSnapshotGroupRequset := ecs.CreateDeleteSnapshotGroupRequest()
	deleteSnapshotGroupRequset.SnapshotGroupId = groupSnapshotID
	response, err := GlobalConfigVar.EcsClient.DeleteSnapshotGroup(deleteSnapshotGroupRequset)
	if err != nil {
		return response, err
	}
	return response, nil
}
