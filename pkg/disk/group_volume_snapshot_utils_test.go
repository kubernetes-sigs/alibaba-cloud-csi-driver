package disk

import (
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/stretchr/testify/assert"
)

func Test_parseGroupSnapshotAnnotations(t *testing.T) {
	annotations := make(map[string]string)
	ecsParams := &createGroupSnapshotParams{}

	// Test case 1: snapshotTTL is empty
	annotations["storage.alibabacloud.com/snapshot-ttl"] = ""
	err := parseGroupSnapshotAnnotations(annotations, ecsParams)
	assert.Nil(t, err)

	// Test case 2: snapshotTTL is not empty
	annotations["storage.alibabacloud.com/snapshot-ttl"] = "10"
	err = parseGroupSnapshotAnnotations(annotations, ecsParams)
	assert.Error(t, err)
}

func Test_formatGroupSnapshot(t *testing.T) {
	// Create a mock groupSnapshot object
	mockGroupSnapshot := &ecs.SnapshotGroup{
		CreationTime:    "2006-01-02T15:04:05Z",
		SnapshotGroupId: "snapshotGroup1",
		Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
			Snapshot: []ecs.Snapshot{
				{
					SnapshotId:     "snapshot1",
					Status:         "accomplished",
					InstantAccess:  false,
					Available:      true,
					SourceDiskSize: "10",
				},
				{
					SnapshotId:     "snapshot2",
					Status:         "progressing",
					InstantAccess:  true,
					Available:      false,
					SourceDiskSize: "10",
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
				SnapshotId:      "snapshot1",
				ReadyToUse:      true,
				CreationTime:    &timestamp.Timestamp{Seconds: stamp.Unix()},
				GroupSnapshotId: "snapshotGroup1",
				SizeBytes:       10 * 1024 * 1024 * 1024,
			},
			{
				SnapshotId:      "snapshot2",
				ReadyToUse:      false,
				CreationTime:    &timestamp.Timestamp{Seconds: stamp.Unix()},
				GroupSnapshotId: "snapshotGroup1",
				SizeBytes:       10 * 1024 * 1024 * 1024,
			},
		},
		CreationTime: &timestamp.Timestamp{Seconds: stamp.Unix()},
		ReadyToUse:   false,
	}

	assert.Equal(t, expectedResult, result)
}

func Test_ifExistsGroupSnapshotMatchSourceVolume(t *testing.T) {
	tests := []struct {
		name              string
		groupNameSnapshot *ecs.SnapshotGroup
		sourceVolumeIds   []string
		expectedResult    bool
	}{

		{
			name: "Matched snapshots and volume IDs",
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
		{
			name: "Different number of snapshots and volume IDs",
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
		{
			name: "Missing volume ID in snapshots",
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
		{
			name: "Extra volume ID not present in snapshots",
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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ifExistsGroupSnapshotMatchSourceVolume(tt.groupNameSnapshot, tt.sourceVolumeIds)
			if result != tt.expectedResult {
				t.Errorf("Expected %v but got %v for input: %v, %v",
					tt.expectedResult, result, tt.groupNameSnapshot, tt.sourceVolumeIds)
			}
		})
	}
}

func Test_ifExistsGroupSnapshotMatch(t *testing.T) {
	tests := []struct {
		name                string
		existsGroupSnapshot *ecs.SnapshotGroup
		snapshotIds         []string
		requireExactMatch   bool
		match               bool
	}{
		{
			name: "Exact match, requireExactMatch = true (default)",
			existsGroupSnapshot: &ecs.SnapshotGroup{
				Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
					Snapshot: []ecs.Snapshot{
						{SnapshotId: "id1"},
						{SnapshotId: "id2"},
					},
				},
			},
			snapshotIds:       []string{"id1", "id2"},
			requireExactMatch: true,
			match:             true,
		},
		{
			name: "Exact match, requireExactMatch = false",
			existsGroupSnapshot: &ecs.SnapshotGroup{
				Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
					Snapshot: []ecs.Snapshot{
						{SnapshotId: "id1"},
						{SnapshotId: "id2"},
					},
				},
			},
			snapshotIds:       []string{"id1", "id2"},
			requireExactMatch: false,
			match:             true,
		},
		{
			name: "Missing snapshot ID, requireExactMatch = true",
			existsGroupSnapshot: &ecs.SnapshotGroup{
				Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
					Snapshot: []ecs.Snapshot{
						{SnapshotId: "id1"},
					},
				},
			},
			snapshotIds:       []string{"id1", "id2"},
			requireExactMatch: true,
			match:             false,
		},
		{
			name: "Extra snapshot ID, requireExactMatch = false",
			existsGroupSnapshot: &ecs.SnapshotGroup{
				Snapshots: ecs.SnapshotsInDescribeSnapshotGroups{
					Snapshot: []ecs.Snapshot{
						{SnapshotId: "id1"},
					},
				},
			},
			snapshotIds:       []string{"id1", "id2"},
			requireExactMatch: false,
			match:             true,
		},
	}
	// Run tests
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ifExistsGroupSnapshotMatch(tt.existsGroupSnapshot, tt.snapshotIds, tt.requireExactMatch)
			assert.Equal(t, tt.match, result)
		})
	}
}

// Test_checkSourceVolume tests the checkSourceVolume function
func Test_checkSourceVolume(t *testing.T) {
	// Define your test cases
	tests := []struct {
		name             string
		disk             *ecs.Disk
		zoneId           string
		capacityInGiB    int
		expectedZoneId   string
		expectedCapacity int
		expectedError    bool
	}{
		{
			name: "valid disk (first disk)",
			disk: &ecs.Disk{
				DiskId:      "disk-1",
				Category:    string(DiskESSDAuto),
				ZoneId:      "zone-a",
				MultiAttach: DiskMultiAttachDisabled,
				Size:        10,
			},
			zoneId:           "",
			capacityInGiB:    0,
			expectedZoneId:   "zone-a",
			expectedCapacity: 10,
			expectedError:    false,
		},
		{
			name: "valid disk",
			disk: &ecs.Disk{
				DiskId:      "disk-1",
				Category:    string(DiskESSD),
				ZoneId:      "zone-a",
				MultiAttach: DiskMultiAttachDisabled,
				Size:        10,
			},
			zoneId:           "zone-a",
			capacityInGiB:    100,
			expectedZoneId:   "zone-a",
			expectedCapacity: 110,
			expectedError:    false,
		},
		{
			name: "invalid category",
			disk: &ecs.Disk{
				DiskId:      "disk-1",
				Category:    string(DiskSSD),
				ZoneId:      "zone-a",
				MultiAttach: DiskMultiAttachDisabled,
				Size:        10,
			},
			zoneId:        "zone-a",
			capacityInGiB: 100,
			expectedError: true,
		},
		{
			name: "from different zones",
			disk: &ecs.Disk{
				DiskId:      "disk-1",
				Category:    string(DiskSSD),
				ZoneId:      "zone-a",
				MultiAttach: DiskMultiAttachDisabled,
				Size:        10,
			},
			zoneId:        "zone-b",
			capacityInGiB: 100,
			expectedError: true,
		},
		{
			name: "multi-attach enabled",
			disk: &ecs.Disk{
				DiskId:      "disk-1",
				Category:    string(DiskESSD),
				ZoneId:      "zone-a",
				MultiAttach: DiskMultiAttachEnabled,
				Size:        10,
			},
			zoneId:        "zone-a",
			capacityInGiB: 100,
			expectedError: true,
		},
		{
			name: "exceed max capacity",
			disk: &ecs.Disk{
				DiskId:      "disk-1",
				Category:    string(DiskESSD),
				ZoneId:      "zone-a",
				MultiAttach: DiskMultiAttachEnabled,
				Size:        10,
			},
			zoneId:        "zone-a",
			capacityInGiB: MaxCapacityInGiB,
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function under test
			nZoneId, nCapacityInGiB, err := checkSourceVolume(tt.disk, tt.zoneId, tt.capacityInGiB)

			// Assert the results
			assert.Equal(t, tt.expectedError, err != nil)
			if !tt.expectedError {
				assert.Equal(t, tt.expectedZoneId, nZoneId)
				assert.Equal(t, tt.expectedCapacity, nCapacityInGiB)
			}
		})
	}
}

func Test_tryGetGroupSnapshotId(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "Created_from_ssg-snapshot123",
			expected: "ssg-snapshot123",
		},
		{
			input:    "Created_from_ssg",
			expected: "",
		},
		{
			input:    "Created_from_ssg-123-456",
			expected: "ssg-123-456",
		},
		{
			input:    "snapshot123",
			expected: "",
		},
	}

	for _, test := range tests {
		actual := tryGetGroupSnapshotId(test.input)
		if actual != test.expected {
			t.Errorf("tryGetGroupSnapshotId(%q) = %q; expected %q", test.input, actual, test.expected)
		}
	}
}
