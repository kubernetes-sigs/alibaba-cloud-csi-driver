package disk

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/common"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func parseGroupSnapshotParameters(params map[string]string) (*createGroupSnapshotParams, error) {
	ecsParams := &createGroupSnapshotParams{}
	tags := make(map[string]string)
	for k, v := range params {
		switch k {
		case RETENTIONDAYS:
			return nil, fmt.Errorf("groupSnapshot do not support retentionDays")
		case SNAPSHOTRESOURCEGROUPID:
			ecsParams.ResourceGroupID = v
		case common.VolumeGroupSnapshotNameKey:
			tags[common.VolumeGroupSnapshotNameTag] = v
		case common.VolumeGroupSnapshotNamespaceKey:
			tags[common.VolumeGroupSnapshotNamespaceTag] = v
		default:
			if strings.HasPrefix(k, SNAPSHOT_TAG_PREFIX) {
				k = k[len(SNAPSHOT_TAG_PREFIX):]
				// don't override built-in tags
				if _, ok := tags[k]; !ok {
					tags[k] = v
				}
			}
		}
	}
	if len(tags) > 0 {
		keys := make([]string, 0, len(tags))
		for k := range tags {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		for _, k := range keys {
			ecsParams.SnapshotTags = append(ecsParams.SnapshotTags, ecs.CreateSnapshotGroupTag{
				Key:   k,
				Value: tags[k],
			})
		}
	}
	return ecsParams, nil
}

func formatGroupSnapshot(groupSnapshot *ecs.SnapshotGroup) (*csi.VolumeGroupSnapshot, error) {
	t, err := time.Parse(time.RFC3339, groupSnapshot.CreationTime)
	if err != nil {
		return nil, fmt.Errorf("failed to parse groupSnapshot creation time: %s", groupSnapshot.CreationTime)
	}
	creationTime := timestamppb.New(t)

	readyToUse := groupSnapshot.Status == SnapshotStatusAccomplished
	snapshots := []*csi.Snapshot{}
	for _, ecsSnapshot := range groupSnapshot.Snapshots.Snapshot {
		snapshot, err := formatCSISnapshot(&ecsSnapshot)
		if err != nil {
			return nil, err
		}
		snapshot.CreationTime = creationTime
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
		CreationTime:    creationTime,
		ReadyToUse:      readyToUse,
	}, nil
}

// ifExistsGroupSnapshotMatchSourceVolume checks whether there SourceDiskId in `existsGroupSnapshot`
// that correspond to all related snapshots in `sourceVolumeIds`.
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

// checkSourceVolumes checks whether a single source volume is valid.
// ref: https://www.alibabacloud.com/help/en/ecs/user-guide/create-a-snapshot-consistent-group?spm=a2c63.p38356.0.0.172c75477J9Sxz#a383a34054gtz
func checkSourceVolume(disk *ecs.Disk, zoneId string, capacityInGiB int) (
	nZoneId string, nCapacityInGiB int, err error) {
	nZoneId = zoneId
	nCapacityInGiB = capacityInGiB

	if !AllCategories[Category(disk.Category)].SnapshotConsistentGroup {
		err = fmt.Errorf("groupSnapshot only support ESSD disks, but disk %s is %v", disk.DiskId, disk.Category)
		return
	}

	if disk.Status != DiskStatusInuse {
		err = fmt.Errorf("groupSnapshot only support in-use disks, but disk %s is %s", disk.DiskId, disk.Status)
		return
	}

	if nZoneId == "" {
		nZoneId = disk.ZoneId
	}
	if nZoneId != disk.ZoneId {
		err = fmt.Errorf("groupSnapshot only support disks from one zone, but get disks from %s and %s", zoneId, disk.ZoneId)
		return
	}

	if disk.MultiAttach != DiskMultiAttachDisabled {
		err = fmt.Errorf("groupSnapshot do not support multiAttached disk, but disk %s enables multiAttached", disk.DiskId)
		return
	}

	nCapacityInGiB += disk.Size
	if nCapacityInGiB > MaxCapacityInGiB {
		err = fmt.Errorf("total capacity is larger than %d", MaxCapacityInGiB)
		return
	}
	return
}

func requestAndDeleteGroupSnapshot(groupSnapshotID string) (*ecs.DeleteSnapshotGroupResponse, error) {
	// Delete Snapshotgroup
	deleteSnapshotGroupRequest := ecs.CreateDeleteSnapshotGroupRequest()
	deleteSnapshotGroupRequest.SnapshotGroupId = groupSnapshotID
	response, err := GlobalConfigVar.EcsClient.DeleteSnapshotGroup(deleteSnapshotGroupRequest)
	if err != nil {
		return response, err
	}
	return response, nil
}

// TODO: groupSnapshotId is missing in ListSnapshotsResponse so we
//
//	1.could not get groupsnapshotId from DescribeSnpshots API response
//	2.could not get volumesnapshot name from ListSnapshotsRequest
//
// this will cause:
// volumesnapshot leased in cluster when volumegroupsnaoshot deleted,
//
//	as status.volumeGroupSnapshotHandle is missing in volumesnapshotcontent,
//	and "snapshot.storage.kubernetes.io/volumesnapshot-in-group-protection" finalizer
//	will not be removed in checkandRemoveSnapshotFinalizersAndCheckandDeleteContent
//	of external-csi-snapshotter
//
// In current version, we try to get snapshotGroupId by description or name of snapshot
func tryGetGroupSnapshotId(str string) string {
	if !strings.HasPrefix(str, "Created_from_ssg-") {
		return ""
	}
	return strings.TrimSpace(strings.TrimPrefix(str, "Created_from_"))
}
