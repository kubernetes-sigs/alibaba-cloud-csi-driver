package disk

import (
	"context"
	"fmt"
	"testing"

	alicloudErr "github.com/aliyun/alibaba-cloud-sdk-go/sdk/errors"
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
