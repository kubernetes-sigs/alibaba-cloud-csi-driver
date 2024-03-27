package disk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/golang/mock/gomock"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	testclock "k8s.io/utils/clock/testing"
)

func disk(id, status string, instances ...string) *ecs.Disk {
	disk := &ecs.Disk{
		DiskId: id,
		Status: status,
	}
	as := make([]ecs.Attachment, len(instances))
	for i, instance := range instances {
		as[i] = ecs.Attachment{
			InstanceId: instance,
		}
	}
	disk.Attachments.Attachment = as
	return disk
}

func diskShared(id, status string, instances ...string) *ecs.Disk {
	disk := &ecs.Disk{
		DiskId: id,
		Status: status,
	}
	mi := make([]ecs.MountInstance, len(instances))
	for i, instance := range instances {
		mi[i] = ecs.MountInstance{
			InstanceId: instance,
		}
	}
	disk.MountInstances.MountInstance = mi
	return disk
}

func setupDescribe(ecsClient *cloud.MockECSInterface, disks *sync.Map) *gomock.Call {
	return ecsClient.EXPECT().DescribeDisks(gomock.Any()).DoAndReturn(func(req *ecs.DescribeDisksRequest) (*ecs.DescribeDisksResponse, error) {
		var diskIds []string
		err := json.Unmarshal([]byte(req.DiskIds), &diskIds)
		if err != nil {
			panic(fmt.Sprintf("failed to unmarshal disk ids: %v", err))
		}
		resp := ecs.CreateDescribeDisksResponse()
		for _, id := range diskIds {
			if disk, ok := disks.Load(id); ok {
				resp.Disks.Disk = append(resp.Disks.Disk, *disk.(*ecs.Disk))
			}
		}
		return resp, nil
	})
}

func isNextStatus(previousStatus, instanceID string) DiskStatusPredicate {
	return func(disk *ecs.Disk) bool {
		found := isInstanceAttached(disk, instanceID)
		switch previousStatus {
		case DiskStatusAttaching:
			if found {
				return true
			}
		case DiskStatusDetaching:
			if !found {
				return true
			}
		}
		return disk.Status != previousStatus
	}
}

func TestWaitForDiskNextStatus(t *testing.T) {
	log.SetLevel(log.DebugLevel)
	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSInterface(ctrl)

	var disks sync.Map
	disks.Store("d1", disk("d1", "Attaching"))

	setupDescribe(ecsClient, &disks).Times(2)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	clk := testclock.NewFakeClock(time.Now())
	waiter := NewBatchedDiskStatusWaiter(ecsClient, clk)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		waiter.Run(ctx)
		wg.Done()
	}()

	finalDisk := disk("d1", "In_use", "i1")
	go func() {
		time.Sleep(100 * time.Millisecond) // for request to be picked up
		t.Logf("Step for first poll")
		clk.Step(100 * time.Millisecond)
		time.Sleep(100 * time.Millisecond) // for the fist poll to finish

		disks.Store("d1", finalDisk)
		t.Logf("Step for second poll")
		clk.Step(pollInterval)
	}()

	{
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		disk, err := waiter.WaitForDiskNextStatus(ctx, "d1", isNextStatus("Attaching", "i1"))
		assert.NoError(t, err)
		assert.Equal(t, *finalDisk, *disk)
	}
	cancel()
	wg.Wait()
}

func TestIsInstanceAttached(t *testing.T) {
	tests := []struct {
		name       string
		disk       *ecs.Disk
		instanceID string
		attached   bool
	}{
		{
			name:       "attaching",
			disk:       disk("d1", "Attaching"),
			instanceID: "i1",
			attached:   false,
		}, {
			name:       "detaching",
			disk:       disk("d1", "Detaching", "i1"),
			instanceID: "i1",
			attached:   true,
		}, {
			name:       "attaching - multi",
			disk:       disk("d1", "Attaching", "i2"),
			instanceID: "i1",
			attached:   false,
		}, {
			name:       "detaching - multi",
			disk:       disk("d1", "Detaching", "i1", "i2"),
			instanceID: "i1",
			attached:   true,
		}, {
			name:       "attaching - done",
			disk:       disk("d1", "In_use", "i1"),
			instanceID: "i1",
			attached:   true,
		}, {
			name:       "detaching - done",
			disk:       disk("d1", "Available"),
			instanceID: "i1",
			attached:   false,
		}, {
			name:       "attaching - done - shared",
			disk:       diskShared("d1", "In_use", "i1", "i2"),
			instanceID: "i1",
			attached:   true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.attached, isInstanceAttached(test.disk, test.instanceID))
		})
	}
}

func TestPoll(t *testing.T) {
	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSInterface(ctrl)

	disks := sync.Map{}
	for i := 0; i < 5; i++ {
		diskID := fmt.Sprintf("d%d", i)
		disks.Store(diskID, disk(diskID, "Available"))
	}
	setupDescribe(ecsClient, &disks)

	waiter := NewBatchedDiskStatusWaiter(ecsClient, nil)
	sentReqs := make([]*waitRequest, 0, 110) // large enough to trigger multiple batches
	for i := 0; i < cap(sentReqs); i++ {
		r := &waitRequest{
			ctx:        context.Background(),
			diskID:     fmt.Sprintf("d%d", i),
			pred:       isNextStatus("Detaching", "i1"),
			resultChan: make(chan *ecs.Disk, 1),
		}
		if i == 2 {
			// make this request already canceled
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			r.ctx = ctx
		}
		if i == 20 {
			// make a repeated request
			r.diskID = "d30"
		}
		waiter.enqueueRequest(r)
		sentReqs = append(sentReqs, r)
	}
	expectedIDs := waiter.idQueue[101:]     // 100 polled, 1 canceled
	assert.Equal(t, "d102", expectedIDs[0]) // 1 repeated request
	ids, err := waiter.poll(waiter.idQueue)
	assert.NoError(t, err)
	assert.Equal(t, expectedIDs, ids)

	for i, req := range sentReqs[:102] {
		if i == 2 { // this request was canceled
			continue
		}
		if i < 5 {
			disk := <-req.resultChan
			expectedID := fmt.Sprintf("d%d", i)
			assert.Equal(t, expectedID, disk.DiskId)
		} else {
			assert.Nil(t, <-req.resultChan)
		}
	}
}

func TestPollError(t *testing.T) {
	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSInterface(ctrl)

	expectedErr := errors.New("fake error")
	ecsClient.EXPECT().DescribeDisks(gomock.Any()).Return(nil, expectedErr)

	waiter := NewBatchedDiskStatusWaiter(ecsClient, nil)
	req := &waitRequest{
		ctx:        context.Background(),
		diskID:     "d-test",
		pred:       isNextStatus("Detaching", "i1"),
		resultChan: make(chan *ecs.Disk, 1),
	}
	waiter.enqueueRequest(req)
	ids, err := waiter.poll(waiter.idQueue)
	assert.Equal(t, expectedErr, err)
	assert.Empty(t, req.resultChan)
	assert.Equal(t, []string{"d-test"}, ids)
}

func TestPollEmpty(t *testing.T) {
	ctrl := gomock.NewController(t)
	ecsClient := cloud.NewMockECSInterface(ctrl)

	waiter := NewBatchedDiskStatusWaiter(ecsClient, nil)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	req := &waitRequest{
		ctx:        ctx,
		diskID:     "d-test",
		pred:       isNextStatus("Detaching", "i1"),
		resultChan: make(chan *ecs.Disk, 1),
	}
	waiter.enqueueRequest(req)
	ids, err := waiter.poll(waiter.idQueue)
	assert.Empty(t, ids)
	assert.Empty(t, waiter.requests)
	assert.NoError(t, err)
}
