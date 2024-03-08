package disk

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func Test_parseGroupSnapshotAnnotations(t *testing.T) {
	annotations := make(map[string]string)
	ecsParams := &createGroupSnapshotParams{}

	// Test case 1: snapshotTTL is empty
	annotations["storage.alibabacloud.com/snapshot-ttl"] = ""
	err := parseGroupSnapshotAnnotations(annotations, ecsParams)
	if err != nil {
		t.Errorf("Expected nil error, but got: %v", err)
	}

	// Test case 2: snapshotTTL is not empty
	annotations["storage.alibabacloud.com/snapshot-ttl"] = "10"
	err = parseGroupSnapshotAnnotations(annotations, ecsParams)
	expectedErr := fmt.Errorf("groupSnapshot do not support retentionDays")
	if err == nil || err.Error() != expectedErr.Error() {
		t.Errorf("Expected error: %v, but got: %v", expectedErr, err)
	}
}

func Test_formatGroupSnapshot(t *testing.T) {
	// Create a mock groupSnapshot object
	mockGroupSnapshot := &ecs.SnapshotGroup{
		CreationTime:    "2006-01-02T15:04:05Z07:00",
		SnapshotGroupId: "snapshotGroup1",
		Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
			Snapshot: []ecs.Snapshot{
				{
					SnapshotId:    "snapshot1",
					Status:        "accomplished",
					InstantAccess: true,
				},
				{
					SnapshotId:    "snapshot2",
					Status:        "progressing",
					InstantAccess: false,
				},
			},
		},
	}

	// Call the formatGroupSnapshot function
	result, err := formatGroupSnapshot(mockGroupSnapshot)

	// Verify the result
	if err != nil {
		t.Errorf("formatGroupSnapshot returned an error: %v", err)
	}

	stamp, err := time.Parse(time.RFC3339, mockGroupSnapshot.CreationTime)
	expectedResult := &csi.VolumeGroupSnapshot{
		GroupSnapshotId: "snapshotGroup1",
		Snapshots: []*csi.Snapshot{
			{
				SnapshotId: "snapshot1",
				ReadyToUse: true,
			},
			{
				SnapshotId: "snapshot2",
				ReadyToUse: false,
			},
		},
		CreationTime: &timestamp.Timestamp{Seconds: stamp.Unix()},
		ReadyToUse:   false,
	}

	if !reflect.DeepEqual(result, expectedResult) {
		t.Errorf("formatGroupSnapshot returned an incorrect result.\nExpected: %+v\n Got: %+v", expectedResult, result)
	}
}

func Test_ifExistsGroupSnapshotMatch(t *testing.T) {
	// Create test SnapshotGroup and sourceVolumeIds
	existsGroupSnapshot := &ecs.SnapshotGroup{
		Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
			Snapshot: []ecs.Snapshot{
				{SourceDiskId: "id1"},
				{SourceDiskId: "id2"},
			},
		},
	}
	sourceVolumeIds := []string{"id1", "id2"}

	// Call the function being tested
	result := ifExistsGroupSnapshotMatch(existsGroupSnapshot, sourceVolumeIds)

	// Verify that the result matches the expected outcome
	if result != true {
		t.Errorf("Expected true, but got false")
	}
}
