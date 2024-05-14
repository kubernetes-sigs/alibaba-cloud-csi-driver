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

func Test_ifExistsGroupSnapshotMatchSourceVolume(t *testing.T) {
	// Define test cases
	testCases := []struct {
		groupNameSnapshot *ecs.SnapshotGroup
		sourceVolumeIds   []string
		expectedResult    bool
	}{
		// Test case 1: Matched snapshots and volume IDs
		{
			groupNameSnapshot: &ecs.SnapshotGroup{
				Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
					Snapshot: []ecs.Snapshot{
						{SourceDiskId: "vol-1"},
						{SourceDiskId: "vol-2"},
					},
				},
			},
			sourceVolumeIds: []string{"vol-1", "vol-2"},
			expectedResult:  true,
		},
		// Test case 2: Different number of snapshots and volume IDs
		{
			groupNameSnapshot: &ecs.SnapshotGroup{
				Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
					Snapshot: []ecs.Snapshot{
						{SourceDiskId: "vol-1"},
					},
				},
			},
			sourceVolumeIds: []string{"vol-1", "vol-2"},
			expectedResult:  false,
		},
		// Test case 3: Missing volume ID in snapshots
		{
			groupNameSnapshot: &ecs.SnapshotGroup{
				Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
					Snapshot: []ecs.Snapshot{
						{SourceDiskId: "vol-1"},
					},
				},
			},
			sourceVolumeIds: []string{"vol-1", "vol-2"},
			expectedResult:  false,
		},
		// Test case 4: Extra volume ID not present in snapshots
		{
			groupNameSnapshot: &ecs.SnapshotGroup{
				Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
					Snapshot: []ecs.Snapshot{
						{SourceDiskId: "vol-1"},
					},
				},
			},
			sourceVolumeIds: []string{"vol-1", "vol-2", "vol-3"},
			expectedResult:  false,
		},
	}

	// Run tests
	for _, tc := range testCases {
		result := ifExistsGroupSnapshotMatchSourceVolume(tc.groupNameSnapshot, tc.sourceVolumeIds)
		if result != tc.expectedResult {
			t.Errorf("Expected %v but got %v for input: %v, %v",
				tc.expectedResult, result, tc.groupNameSnapshot, tc.sourceVolumeIds)
		}
	}
}

func Test_ifExistsGroupSnapshotMatch(t *testing.T) {
	// Test case 1: Exact match, requireExactMatch = true (default)
	existsGroupSnapshot1 := &ecs.SnapshotGroup{
		Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
			Snapshot: []ecs.Snapshot{
				{SnapshotId: "id1"},
				{SnapshotId: "id2"},
			},
		},
	}
	snapshotIds1 := []string{"id1", "id2"}
	if result1 := ifExistsGroupSnapshotMatch(existsGroupSnapshot1, snapshotIds1, true); result1 != true {
		t.Errorf("Test 1: Expected true, but got false")
	}

	// Test case 2: Exact match, requireExactMatch = false
	existsGroupSnapshot2 := &ecs.SnapshotGroup{
		Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
			Snapshot: []ecs.Snapshot{
				{SnapshotId: "id1"},
				{SnapshotId: "id2"},
			},
		},
	}
	snapshotIds2 := []string{"id1", "id2"}
	if result2 := ifExistsGroupSnapshotMatch(existsGroupSnapshot2, snapshotIds2, false); result2 != true {
		t.Errorf("Test 2: Expected true, but got false")
	}

	// Test case 3: Missing snapshot ID, requireExactMatch = true
	existsGroupSnapshot3 := &ecs.SnapshotGroup{
		Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
			Snapshot: []ecs.Snapshot{
				{SnapshotId: "id1"},
			},
		},
	}
	snapshotIds3 := []string{"id1", "id2"}
	if result3 := ifExistsGroupSnapshotMatch(existsGroupSnapshot3, snapshotIds3, true); result3 != false {
		t.Errorf("Test 3: Expected false, but got true")
	}

	// Test case 4: Extra snapshot ID, requireExactMatch = false
	existsGroupSnapshot4 := &ecs.SnapshotGroup{
		Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
			Snapshot: []ecs.Snapshot{
				{SnapshotId: "id1"},
			},
		},
	}
	snapshotIds4 := []string{"id1", "id2"}
	if result4 := ifExistsGroupSnapshotMatch(existsGroupSnapshot4, snapshotIds4, false); result4 != true {
		t.Errorf("Test 4: Expected true, but got false")
	}
}
