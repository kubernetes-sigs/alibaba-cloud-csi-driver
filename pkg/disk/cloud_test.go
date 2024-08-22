package disk

import (
	"context"
	"fmt"
	"testing"

	alicloudErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	gomock "github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/stretchr/testify/assert"
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

func TestDeleteDisk(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	c.EXPECT().DeleteDisk(gomock.Any()).Return(deleteDiskResponse, nil)

	_, err := deleteDisk(context.Background(), c, "test-disk")
	assert.Nil(t, err)
}

func TestDeleteDiskRetryOnInitError(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	initErr := alicloudErr.NewServerError(400, `{"Code": "IncorrectDiskStatus.Initializing"}`, "")
	c.EXPECT().DeleteDisk(gomock.Any()).Return(nil, initErr)
	c.EXPECT().DeleteDisk(gomock.Any()).Return(deleteDiskResponse, nil)

	_, err := deleteDisk(context.Background(), c, "test-disk")
	assert.Nil(t, err)
}

func TestDeleteDiskPassthroughError(t *testing.T) {
	ctrl := gomock.NewController(t)
	c := cloud.NewMockECSInterface(ctrl)

	serverErr := alicloudErr.NewServerError(400, `{"Code": "AnyOtherErrors"}`, "")
	c.EXPECT().DeleteDisk(gomock.Any()).Return(nil, serverErr)

	_, err := deleteDisk(context.Background(), c, "test-disk")
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
				Type: []Category{DiskESSD, DiskESSDAuto},
			},
			attempts: []createAttempt{
				{Category: DiskESSD},
				{Category: DiskESSDAuto},
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
		name       string
		args       *diskVolumeArgs
		err        bool
		serverFail bool
	}{
		{
			name: "success",
			args: &diskVolumeArgs{Type: []Category{DiskESSD}, RequestGB: 20},
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
			ctrl := gomock.NewController(t)
			client := cloud.NewMockECSInterface(ctrl)

			if !c.err {
				client.EXPECT().CreateDisk(gomock.Any()).Return(&ecs.CreateDiskResponse{
					DiskId: "d-123",
				}, nil)
			}
			if c.serverFail {
				client.EXPECT().CreateDisk(gomock.Any()).Return(nil, alicloudErr.NewServerError(400, `{"Code": "AnyOtherErrors"}`, ""))
			}

			diskID, attempt, err := createDisk(client, "disk-name", "", c.args)
			if c.err {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, "d-123", diskID)
				assert.Equal(t, DiskESSD, attempt.Category)
				assert.Empty(t, attempt.PerformanceLevel)
			}
		})
	}
}

func TestCreateDisk_ServerFailFallback(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := cloud.NewMockECSInterface(ctrl)

	client.EXPECT().CreateDisk(gomock.Any()).Return(nil, alicloudErr.NewServerError(400, `{"Code": "InvalidDataDiskSize.ValueNotSupported"}`, ""))
	client.EXPECT().CreateDisk(gomock.Any()).Return(&ecs.CreateDiskResponse{
		DiskId: "d-123",
	}, nil)

	args := &diskVolumeArgs{Type: []Category{DiskESSD, DiskESSDAuto}, RequestGB: 20}
	diskID, attempt, err := createDisk(client, "disk-name", "", args)
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
			ctrl := gomock.NewController(t)
			client := cloud.NewMockECSInterface(ctrl)

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
			diskID, attempt, err := createDisk(client, "disk-name", "", args)
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

func TestCreateDisk_NoInfinitLoop(t *testing.T) {
	ctrl := gomock.NewController(t)
	client := cloud.NewMockECSInterface(ctrl)

	client.EXPECT().CreateDisk(gomock.Any()).Return(nil, alicloudErr.NewServerError(400, `{"Code": "IdempotentParameterMismatch"}`, "")).Times(2)
	client.EXPECT().DescribeDisks(gomock.Any()).Return(&ecs.DescribeDisksResponse{
		Disks: ecs.DisksInDescribeDisks{
			Disk: []ecs.Disk{},
		},
	}, nil)

	args := &diskVolumeArgs{Type: []Category{DiskESSD}, RequestGB: 20}
	_, _, err := createDisk(client, "disk-name", "", args)
	assert.Error(t, err)
}
