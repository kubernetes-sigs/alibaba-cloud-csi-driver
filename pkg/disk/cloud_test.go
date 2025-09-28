package disk

import (
	"context"
	"fmt"
	"testing"
	"time"

	alicloudErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	gomock "github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/batcher"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/waitstatus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/klog/v2/ktesting"
	"k8s.io/utils/clock"
)

var (
	resizeDiskRequest = ecs.CreateResizeDiskRequest()

	deleteDiskResponse = ecs.CreateDeleteDiskResponse()
	resizeDiskResponse = ecs.CreateResizeDiskResponse()
)

func init() {
	cloud.UnmarshalAcsResponse([]byte(`{
	"RequestId": "B6B6D6B6-6B6B-6B6B-6B6B-6B6B6B6B6B6"
}`), deleteDiskResponse)

	cloud.UnmarshalAcsResponse([]byte(`{
	"RequestId": "B6B6D6B6-6B6B-6B6B-6B6B-6B6B6B6B6B7"
}`), resizeDiskResponse)
}

func testCreateDelete(t *testing.T) (*cloud.MockECSInterface, *DiskCreateDelete) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	return c, &DiskCreateDelete{
		ecs:             c,
		createThrottler: defaultThrottler(),
		deleteThrottler: defaultThrottler(),
		batcher:         batcher.NewLowLatency(desc.Disk(c), clock.RealClock{}, 1*time.Second, 8),
	}
}

func TestDeleteDisk(t *testing.T) {
	c, cd := testCreateDelete(t)
	_, ctx := ktesting.NewTestContext(t)

	c.EXPECT().DeleteDisk(gomock.Any()).Return(deleteDiskResponse, nil)

	_, err := cd.deleteDisk(ctx, c, "test-disk")
	assert.Nil(t, err)
}

func TestDeleteDiskRetryOnInitError(t *testing.T) {
	t.Parallel()
	c, cd := testCreateDelete(t)
	_, ctx := ktesting.NewTestContext(t)

	initErr := alicloudErr.NewServerError(400, `{"Code": "IncorrectDiskStatus.Initializing"}`, "")
	c.EXPECT().DeleteDisk(gomock.Any()).Return(nil, initErr)
	c.EXPECT().DeleteDisk(gomock.Any()).Return(deleteDiskResponse, nil)

	_, err := cd.deleteDisk(ctx, c, "test-disk")
	assert.Nil(t, err)
}

func TestDeleteDiskPassthroughError(t *testing.T) {
	c, cd := testCreateDelete(t)
	_, ctx := ktesting.NewTestContext(t)

	serverErr := alicloudErr.NewServerError(400, `{"Code": "AnyOtherErrors"}`, "")
	c.EXPECT().DeleteDisk(gomock.Any()).Return(nil, serverErr)

	_, err := cd.deleteDisk(ctx, c, "test-disk")
	assert.Equal(t, serverErr, err)
}

func TestResizeDisk(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	c.EXPECT().ResizeDisk(gomock.Any()).Return(resizeDiskResponse, nil)

	_, err := resizeDisk(context.Background(), c, resizeDiskRequest)
	assert.Nil(t, err)
}

func TestResizeDiskRetryOnProcessingError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	initErr := alicloudErr.NewServerError(400, `{"Code": "LastOrderProcessing"}`, "")
	c.EXPECT().ResizeDisk(gomock.Any()).Return(nil, initErr)
	c.EXPECT().ResizeDisk(gomock.Any()).Return(resizeDiskResponse, nil)

	_, err := resizeDisk(context.Background(), c, resizeDiskRequest)
	assert.Nil(t, err)
}

func TestResizeDiskPassthroughError(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	serverErr := alicloudErr.NewServerError(400, `{"Code": "AnyOtherErrors"}`, "")
	c.EXPECT().ResizeDisk(gomock.Any()).Return(nil, serverErr)

	_, err := resizeDisk(context.Background(), c, resizeDiskRequest)
	assert.Equal(t, serverErr, err)
}

func TestListSnapshots(t *testing.T) {
	cases := []struct {
		name          string
		numRemaining  int
		maxEntries    int
		nextToken     string
		expectedNum   int
		firstID       string
		expectedToken string
	}{
		{
			name:         "empty",
			numRemaining: 0, maxEntries: 0, nextToken: "", expectedNum: 0, firstID: "",
		}, {
			name:         "one",
			numRemaining: 1, maxEntries: 0, nextToken: "", expectedNum: 1, firstID: "snap-0",
		}, {
			name:         "skip one",
			numRemaining: 2, maxEntries: 0, nextToken: "1@", expectedNum: 1, firstID: "snap-1",
		}, {
			name:         "paged",
			numRemaining: 13, maxEntries: 5, nextToken: "6@", expectedNum: 5, firstID: "snap-6",
			expectedToken: "0@next-page",
		}, {
			name:         "middle of page",
			numRemaining: 3, maxEntries: 1, nextToken: "1@next-page", expectedNum: 1, firstID: "snap-1",
			expectedToken: "2@next-page",
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			client := cloud.NewMockECSInterface(ctrl)

			client.EXPECT().DescribeSnapshots(gomock.Any()).DoAndReturn(func(req *ecs.DescribeSnapshotsRequest) (*ecs.DescribeSnapshotsResponse, error) {
				snapshots := make([]ecs.Snapshot, c.numRemaining)
				for i := 0; i < c.numRemaining; i++ {
					snapshots[i] = ecs.Snapshot{SnapshotId: fmt.Sprintf("snap-%d", i)}
				}
				if req.NextToken != "" {
					assert.Equal(t, "next-page", req.NextToken, "n@ should not be passed to the API")
				}
				max := 10
				if req.MaxResults.HasValue() {
					var err error
					max, err = req.MaxResults.GetValue()
					assert.NoError(t, err)
				}
				if max < 10 {
					max = 10 // mimic the API behavior
				}
				nextToken := ""
				if c.numRemaining > max {
					assert.Empty(t, req.NextToken, "not supporting page 2 for now")
					nextToken = "next-page"
					snapshots = snapshots[:max]
				}
				return &ecs.DescribeSnapshotsResponse{
					Snapshots: ecs.SnapshotsInDescribeSnapshots{
						Snapshot: snapshots,
					},
					NextToken: nextToken,
				}, nil
			})

			s, nextToken, err := listSnapshots(client, "test-disk", "my-cluster", c.nextToken, c.maxEntries)
			assert.NoError(t, err)
			assert.Len(t, s, c.expectedNum)
			if c.expectedNum > 0 {
				assert.Equal(t, c.firstID, s[0].SnapshotId)
			}
			assert.Equal(t, c.expectedToken, nextToken)
		})
	}
}

func TestListSnapshotsInvalidToken(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := cloud.NewMockECSInterface(ctrl)

	_, _, err := listSnapshots(client, "test-disk", "my-cluster", "invalid-token", 0)
	assert.Error(t, err)
}

func TestClientToken(t *testing.T) {
	// we should keep the token stable across versions
	assert.Equal(t, "n:disk-dcd6fdde-8c1e-45eb-8ec7-786a8b2e0b61", clientToken("disk-dcd6fdde-8c1e-45eb-8ec7-786a8b2e0b61"))
	// only ASCII characters are allowed
	assert.Equal(t, "h:LGH7vCPQbR31I47I1eCW7g", clientToken("disk-磁盘名称-1"))

	// the length should be <= 64
	assert.Equal(t, "n:01234567890123456789012345678901234567890123456789012345678901",
		clientToken("01234567890123456789012345678901234567890123456789012345678901"))
	assert.Equal(t, "h:NDeYXVDChDCom5xYgHLVQA",
		clientToken("012345678901234567890123456789012345678901234567890123456789012"))
}

func BenchmarkClientToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		clientToken("disk-dcd6fdde-8c1e-45eb-8ec7-786a8b2e0b61")
	}
}

func TestBuildCreateDiskRequest(t *testing.T) {
	args := &diskVolumeArgs{
		ZoneID: "cn-hangzhou",
	}
	req := buildCreateDiskRequest(args)
	assert.Equal(t, "cn-hangzhou", req.ZoneId)

	req2 := finalizeCreateDiskRequest(req, createAttempt{
		Category:         DiskESSD,
		PerformanceLevel: PERFORMANCE_LEVEL0,
	})
	assert.Equal(t, "cloud_essd", req2.DiskCategory)
	assert.Equal(t, "PL0", req2.PerformanceLevel)
	// fields is copied
	assert.Equal(t, "cn-hangzhou", req2.ZoneId)

	// send req2 should not affect req
	requests.InitParams(req2)
	assert.Greater(t, len(req2.QueryParams), len(req.QueryParams))
}

func TestGenerateAttempts(t *testing.T) {
	cases := []struct {
		name     string
		args     *diskVolumeArgs
		attempts []createAttempt
	}{
		{
			name: "no PL",
			args: &diskVolumeArgs{
				Type: []Category{DiskESSD, DiskESSDAuto, DiskESSDXc0, DiskESSDXc1},
			},
			attempts: []createAttempt{
				{Category: DiskESSD},
				{Category: DiskESSDAuto},
				{Category: DiskESSDXc0},
				{Category: DiskESSDXc1},
			},
		}, {
			name: "with PL",
			args: &diskVolumeArgs{
				Type:             []Category{DiskESSDEntry, DiskESSD, DiskESSDAuto},
				PerformanceLevel: []PerformanceLevel{PERFORMANCE_LEVEL0, PERFORMANCE_LEVEL1},
			},
			attempts: []createAttempt{
				{Category: DiskESSDEntry},
				{Category: DiskESSD, PerformanceLevel: PERFORMANCE_LEVEL0},
				{Category: DiskESSD, PerformanceLevel: PERFORMANCE_LEVEL1},
				{Category: DiskESSDAuto},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			attempts := generateCreateAttempts(c.args)
			assert.Equal(t, c.attempts, attempts)
		})
	}
}

func TestCheckExistingDisk(t *testing.T) {
	disk := &ecs.Disk{
		Size:             20,
		Category:         "cloud_essd",
		PerformanceLevel: "PL0",
		Tags: ecs.TagsInDescribeDisks{
			Tag: []ecs.Tag{
				{Key: "k1", Value: "v1"},
			},
		},
	}
	cases := []struct {
		name  string
		args  *diskVolumeArgs
		match bool
	}{
		{
			name:  "match",
			args:  &diskVolumeArgs{RequestGB: 20, Type: []Category{DiskESSD, DiskESSDAuto}, PerformanceLevel: []PerformanceLevel{PERFORMANCE_LEVEL0}},
			match: true,
		}, {
			name: "mismatch category",
			args: &diskVolumeArgs{RequestGB: 20, Type: []Category{DiskESSDAuto}, PerformanceLevel: []PerformanceLevel{PERFORMANCE_LEVEL0}},
		}, {
			name: "mismatch PL",
			args: &diskVolumeArgs{RequestGB: 20, Type: []Category{DiskESSD}, PerformanceLevel: []PerformanceLevel{PERFORMANCE_LEVEL1}},
		}, {
			name: "mismatch MultiAttach",
			args: &diskVolumeArgs{
				RequestGB: 20, Type: []Category{DiskESSD}, PerformanceLevel: []PerformanceLevel{PERFORMANCE_LEVEL0},
				MultiAttach: true,
			},
		}, {
			name: "mismatch tag key",
			args: &diskVolumeArgs{
				RequestGB: 20, Type: []Category{DiskESSD}, PerformanceLevel: []PerformanceLevel{PERFORMANCE_LEVEL0},
				DiskTags: map[string]string{"k2": "v1"},
			},
		}, {
			name: "mismatch tag value",
			args: &diskVolumeArgs{
				RequestGB: 20, Type: []Category{DiskESSD}, PerformanceLevel: []PerformanceLevel{PERFORMANCE_LEVEL0},
				DiskTags: map[string]string{"k1": "v2"},
			},
		}, {
			name:  "match no PL requested",
			args:  &diskVolumeArgs{RequestGB: 20, Type: []Category{DiskESSD}},
			match: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			createAttempt, err := checkExistingDisk(disk, c.args)
			assert.Equal(t, c.match, err == nil)
			if c.match {
				assert.Equal(t, disk.Category, string(createAttempt.Category))
				assert.Equal(t, disk.PerformanceLevel, string(createAttempt.PerformanceLevel))
			}
		})
	}
}

// Cases that only hit the server at most once
func TestCreateDisk_Basic(t *testing.T) {
	cases := []struct {
		name          string
		supports      sets.Set[Category]
		instance      string
		isVirtualNode bool
		args          *diskVolumeArgs
		serverFail    bool
		expected      createAttempt
		err           bool
	}{
		{
			name:     "success",
			supports: sets.New(DiskESSD),
			args:     &diskVolumeArgs{Type: []Category{DiskESSD}, RequestGB: 20},
			expected: createAttempt{Category: DiskESSD},
		}, {
			name:     "success - fallback",
			supports: sets.New(DiskESSD),
			args:     &diskVolumeArgs{Type: []Category{DiskSSD, DiskESSD}, RequestGB: 20},
			expected: createAttempt{Category: DiskESSD},
		}, {
			name:     "success - empty supports",
			args:     &diskVolumeArgs{Type: []Category{DiskESSD}, RequestGB: 20},
			expected: createAttempt{Category: DiskESSD},
		}, {
			name:     "success - EED",
			args:     &diskVolumeArgs{Type: []Category{DiskEEDStandard}, RequestGB: 100},
			instance: "i-someinstance",
			expected: createAttempt{Category: DiskEEDStandard, Instance: "i-someinstance"},
		}, {
			name: "EED no instance",
			args: &diskVolumeArgs{Type: []Category{DiskEEDStandard}, RequestGB: 100},
			err:  true,
		}, {
			name:          "EED virtual node",
			args:          &diskVolumeArgs{Type: []Category{DiskEEDStandard}, RequestGB: 100},
			isVirtualNode: true,
			expected:      createAttempt{Category: DiskEEDStandard},
		}, {
			name:     "unsupported",
			supports: sets.New(DiskSSD),
			args:     &diskVolumeArgs{Type: []Category{DiskESSD}, RequestGB: 20},
			err:      true,
		}, {
			name: "too small",
			args: &diskVolumeArgs{Type: []Category{DiskSSD}, RequestGB: 1},
			err:  true,
		}, {
			name:       "server fail",
			args:       &diskVolumeArgs{Type: []Category{DiskESSD}, RequestGB: 20},
			err:        true,
			serverFail: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			client, cd := testCreateDelete(t)
			_, ctx := ktesting.NewTestContext(t)

			if !c.err {
				client.EXPECT().CreateDisk(gomock.Any()).Return(&ecs.CreateDiskResponse{
					DiskId: "d-123",
				}, nil)
			}
			if c.serverFail {
				client.EXPECT().CreateDisk(gomock.Any()).Return(nil, alicloudErr.NewServerError(400, `{"Code": "AnyOtherErrors"}`, ""))
			}

			diskID, attempt, err := cd.createDisk(ctx, "disk-name", "", c.args, c.supports, c.instance, c.isVirtualNode)
			if c.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "d-123", diskID)
				assert.Equal(t, c.expected, attempt)
			}
		})
	}
}

func TestCreateDisk_ServerFailFallback(t *testing.T) {
	client, cd := testCreateDelete(t)
	_, ctx := ktesting.NewTestContext(t)

	client.EXPECT().CreateDisk(gomock.Any()).Return(nil, alicloudErr.NewServerError(400, `{"Code": "InvalidDataDiskSize.ValueNotSupported"}`, ""))
	client.EXPECT().CreateDisk(gomock.Any()).Return(&ecs.CreateDiskResponse{
		DiskId: "d-123",
	}, nil)

	args := &diskVolumeArgs{Type: []Category{DiskESSD, DiskESSDAuto}, RequestGB: 20}
	diskID, attempt, err := cd.createDisk(ctx, "disk-name", "", args, nil, "", false)
	assert.NoError(t, err)
	assert.Equal(t, "d-123", diskID)
	assert.Equal(t, DiskESSDAuto, attempt.Category)
	assert.Empty(t, attempt.PerformanceLevel)
}

func TestCreateDisk_ParameterMismatch(t *testing.T) {
	cases := []struct {
		name     string
		existing []ecs.Disk
		err      bool
	}{
		{
			name: "retry",
		}, {
			name: "reuse",
			existing: []ecs.Disk{{
				DiskId:   "d-124",
				Category: "cloud_auto",
				Size:     20,
			}},
		}, {
			name: "mismatch",
			existing: []ecs.Disk{{
				DiskId:   "d-124",
				Category: "cloud_essd_entry",
				Size:     20,
			}},
			err: true,
		}, {
			name:     "multiple existing",
			existing: []ecs.Disk{{}, {}},
			err:      true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			client, cd := testCreateDelete(t)
			_, ctx := ktesting.NewTestContext(t)

			r1 := client.EXPECT().CreateDisk(gomock.Any()).Return(nil, alicloudErr.NewServerError(400, `{"Code": "IdempotentParameterMismatch"}`, ""))
			r2 := client.EXPECT().DescribeDisks(gomock.Any()).Return(&ecs.DescribeDisksResponse{
				Disks: ecs.DisksInDescribeDisks{
					Disk: c.existing,
				},
			}, nil).After(r1)
			if c.existing == nil {
				client.EXPECT().CreateDisk(gomock.Any()).Return(&ecs.CreateDiskResponse{
					DiskId: "d-123",
				}, nil).After(r2)
			}

			args := &diskVolumeArgs{Type: []Category{DiskESSD, DiskESSDAuto}, RequestGB: 20}
			diskID, attempt, err := cd.createDisk(ctx, "disk-name", "", args, nil, "", false)
			if c.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				if c.existing == nil {
					assert.Equal(t, "d-123", diskID)
					assert.Equal(t, DiskESSD, attempt.Category)
					assert.Empty(t, attempt.PerformanceLevel)
				} else {
					d := c.existing[0]
					assert.Equal(t, d.DiskId, diskID)
					assert.Equal(t, Category(d.Category), attempt.Category)
				}
			}
		})
	}
}

func TestCreateDisk_NoInfiniteLoop(t *testing.T) {
	client, cd := testCreateDelete(t)
	_, ctx := ktesting.NewTestContext(t)

	client.EXPECT().CreateDisk(gomock.Any()).Return(nil, alicloudErr.NewServerError(400, `{"Code": "IdempotentParameterMismatch"}`, "")).Times(2)
	client.EXPECT().DescribeDisks(gomock.Any()).Return(&ecs.DescribeDisksResponse{
		Disks: ecs.DisksInDescribeDisks{
			Disk: []ecs.Disk{},
		},
	}, nil)

	args := &diskVolumeArgs{Type: []Category{DiskESSD}, RequestGB: 20}
	_, _, err := cd.createDisk(ctx, "disk-name", "", args, nil, "", false)
	assert.Error(t, err)
}

func TestValidSnapshotName(t *testing.T) {
	test := func(name string) {
		t.Run(name, func(t *testing.T) {
			assert.True(t, isValidSnapshotName(name))
		})
	}
	test("快照")
	test("snapshot_name-1")
}

func TestInvalidSnapshotName(t *testing.T) {
	test := func(name string) {
		t.Run(name, func(t *testing.T) {
			assert.False(t, isValidSnapshotName(name))
		})
	}
	test("autosnap") // not to be confused with real auto snapshots
	test("a")        // too short
}

func TestValidDiskName(t *testing.T) {
	test := func(name string) {
		t.Run(name, func(t *testing.T) {
			assert.True(t, isValidDiskName(name))
		})
	}
	test("块存储")
	test("disk_name-1")
}

func TestInvalidDiskName(t *testing.T) {
	test := func(name string) {
		t.Run(name, func(t *testing.T) {
			assert.False(t, isValidDiskName(name))
		})
	}
	test("a")     // too short
	test("块")     // too short
	test("???")   // ? not supported
	test("0asdf") // must start with letter
	test("😊😊😊")   // not letter
}

func Test_getDiskDescribeRequest(t *testing.T) {
	tests := []struct {
		name     string
		diskIDs  []string
		expected string
	}{
		{
			name:     "single disk ID",
			diskIDs:  []string{"disk-1"},
			expected: "[\"disk-1\"]",
		},
		{
			name:     "multiple disk IDs",
			diskIDs:  []string{"disk-1", "disk-2"},
			expected: "[\"disk-1\",\"disk-2\"]",
		},
		{
			name:     "no disk IDs",
			diskIDs:  []string{},
			expected: "[]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			request := getDiskDescribeRequest(tt.diskIDs)
			assert.NotNil(t, request)
			assert.Equal(t, tt.expected, request.DiskIds)
		})
	}
}

func testAttachDetach(t *testing.T) (context.Context, *cloud.MockECSInterface, *DiskAttachDetach) {
	ctrl := gomock.NewController(t)
	ecs := cloud.NewMockECSInterface(ctrl)
	_, ctx := ktesting.NewTestContext(t)

	client := desc.Disk(ecs)
	b := batcher.NewPassthrough(client)
	return ctx, ecs, &DiskAttachDetach{
		slots:           NewSlots(0, 0),
		ecs:             ecs,
		waiter:          waitstatus.NewSimple(client, clock.RealClock{}),
		batcher:         b,
		attachThrottler: defaultThrottler(),
		detachThrottler: defaultThrottler(),
		dev:             DefaultDeviceManager,
	}
}

func disk(status string, node string) ecs.Disk {
	disk := ecs.Disk{
		Status:       status,
		Category:     "cloud_regional_disk_auto", // only one support forceAttach
		MultiAttach:  "Disabled",
		DiskId:       "d-testdiskid",
		InstanceId:   node,
		SerialNumber: "fake-serial-number",
	}
	if node != "" {
		disk.Attachments.Attachment = []ecs.Attachment{
			{InstanceId: node},
		}
	}
	return disk
}

func diskResp(disk ecs.Disk) *ecs.DescribeDisksResponse {
	return &ecs.DescribeDisksResponse{
		Disks: ecs.DisksInDescribeDisks{
			Disk: []ecs.Disk{disk},
		},
	}
}

func TestAttachDisk(t *testing.T) {
	GlobalConfigVar.DetachBeforeAttach = true // This is the default
	cases := []struct {
		name          string
		before, after ecs.Disk
		detaching     bool
		detach        bool
		detached      ecs.Disk
		noAttach      bool
		forceAttach   bool
		expectErr     bool
	}{
		{
			name:     "already attached",
			before:   disk("In_use", "i-testinstanceid"),
			noAttach: true,
		},
		{
			name:        "attached to other",
			before:      disk("In_use", "i-anotherinstance"),
			detaching:   true,
			forceAttach: true,
			after:       disk("In_use", "i-testinstanceid"),
		},
		{
			name:     "attached to other (no force attach)",
			before:   disk("In_use", "i-anotherinstance"),
			detach:   true,
			detached: disk("Available", ""),
			after:    disk("In_use", "i-testinstanceid"),
		},
		{
			name:   "normal",
			before: disk("Available", ""),
			after:  disk("In_use", "i-testinstanceid"),
		},
		{
			name:      "attaching",
			before:    disk("Attaching", ""),
			noAttach:  true,
			expectErr: true,
		},
		{
			name:   "detaching from self",
			before: disk("Detaching", "i-testinstanceid"),
			after:  disk("In_use", "i-testinstanceid"), // FIXME
		},
		{
			// This not supported by ECS, but desired by us to speed up failover. Hopes they will support it someday.
			name:        "detaching from other",
			before:      disk("Detaching", "i-anotherinstance"),
			detaching:   true,
			forceAttach: true,
			after:       disk("In_use", "i-testinstanceid"),
		},
		{
			// This is likely to fail in real env. But we try it anyway, in case the detach just finished after we checked
			name:        "detaching from other (no force attach)",
			before:      disk("Detaching", "i-anotherinstance"),
			forceAttach: false,
			after:       disk("In_use", "i-testinstanceid"),
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, c, ad := testAttachDetach(t)

			if tc.detaching {
				ad.detaching.Store("d-testdiskid", "i-anotherinstance")
			}

			c.EXPECT().DescribeDisks(gomock.Any()).Return(diskResp(tc.before), nil)
			if tc.detach {
				detachCall := c.EXPECT().DetachDisk(gomock.Any()).Return(&ecs.DetachDiskResponse{}, nil)
				c.EXPECT().DescribeDisks(gomock.Any()).Return(diskResp(tc.detached), nil).After(detachCall)
			}
			force := false
			if !tc.noAttach {
				attachCall := c.EXPECT().AttachDisk(gomock.Any()).Return(&ecs.AttachDiskResponse{}, nil).
					Do(func(request *ecs.AttachDiskRequest) {
						if request.Force.HasValue() {
							var err error
							force, err = request.Force.GetValue()
							if err != nil {
								t.Fatalf("unexpected error: %v", err)
							}
						}
					})
				c.EXPECT().DescribeDisks(gomock.Any()).Return(diskResp(tc.after), nil).After(attachCall)
			}
			serial, err := ad.attachDisk(ctx, "d-testdiskid", "i-testinstanceid", false)

			assert.Equal(t, tc.forceAttach, force)
			if tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, "fake-serial-number", serial)
			}
		})
	}
}
