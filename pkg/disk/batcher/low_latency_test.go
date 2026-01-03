package batcher

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc"
	testdesc "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc/testing"
	"github.com/stretchr/testify/assert"
	"k8s.io/klog/v2"
	"k8s.io/klog/v2/ktesting"
	"k8s.io/utils/clock"
	testclock "k8s.io/utils/clock/testing"
)

func TestZeroBucketIsFilled(t *testing.T) {
	bucket := tokenBucket{
		limit: 16 * time.Second, perToken: 2 * time.Second,
	}
	assert.Equal(t, 8., bucket.tokenAt(time.Now()))
}

func TestBucketTake(t *testing.T) {
	bucket := tokenBucket{
		limit: 16 * time.Second, perToken: 2 * time.Second,
	}
	now := time.Now()
	bucket.takeAt(now)
	assert.Equal(t, 7., bucket.tokenAt(now))

	bucket.takeAt(now)
	assert.Equal(t, 6., bucket.tokenAt(now))

	assert.Equal(t, 6.5, bucket.tokenAt(now.Add(time.Second)))
	assert.Equal(t, 8., bucket.tokenAt(now.Add(100*time.Second)))

	for range 10 {
		bucket.takeAt(now)
	}
	// bucket is empty, cannot be negative
	assert.Equal(t, 0., bucket.tokenAt(now))
}

func setup(t *testing.T) (client *testdesc.FakeClient, clk *testclock.FakeClock, batcher *LowLatency[ecs.Disk]) {
	client = &testdesc.FakeClient{}
	clk = testclock.NewFakeClock(time.Now())
	batcher = NewLowLatency(client, clk, 2*time.Second, 8)

	logger := ktesting.NewLogger(t, ktesting.DefaultConfig)
	ctx := klog.NewContext(t.Context(), logger)
	go batcher.Run(ctx)

	return client, clk, batcher
}

func TestCancel(t *testing.T) {
	t.Parallel()
	client, clk, batcher := setup(t)

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := batcher.Describe(ctx, "d-test")
	assert.ErrorIs(t, err, context.Canceled)

	time.Sleep(100 * time.Millisecond)
	clk.Step(10 * time.Second)
	time.Sleep(100 * time.Millisecond)
	assert.Len(t, client.Requests, 0) // no requests were made
}

func TestBasic(t *testing.T) {
	t.Parallel()
	client, clk, batcher := setup(t)
	client.Disks.Store("d-test", &ecs.Disk{
		DiskId:   "d-test",
		DiskName: "mydisk",
	})

	var disk *ecs.Disk
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		var err error
		disk, err = batcher.Describe(context.Background(), "d-test")
		assert.NoError(t, err)
		assert.Equal(t, "mydisk", disk.DiskName)
		wg.Done()
	}()
	for range 2 {
		go func() {
			disk, err := batcher.Describe(context.Background(), "d-test-not-found")
			assert.NoError(t, err)
			assert.Nil(t, disk)
			wg.Done()
		}()
	}
	time.Sleep(100 * time.Millisecond)
	assert.Nil(t, disk)             // not returned yet
	clk.Step(10 * time.Millisecond) // expected wait time for first request is 2/(2^8) seconds

	wg.Wait()
	assert.Len(t, client.Requests, 1)
	assert.Len(t, client.Requests[0], 2) // deduplicate requests
}

func TestImmediateFullBatch(t *testing.T) {
	t.Parallel()
	client, _, batcher := setup(t)

	res := make(chan *ecs.Disk)
	for i := range client.BatchSize() {
		diskID := fmt.Sprintf("d-disk-%d", i)
		go func() {
			disk, err := batcher.Describe(context.Background(), diskID)
			assert.NoError(t, err)
			res <- disk
		}()
	}

	for range client.BatchSize() {
		disk := <-res
		assert.Nil(t, disk)
	}

	// only one request was made (batch), even clk is not stepped
	assert.Len(t, client.Requests, 1)
	assert.Len(t, client.Requests[0], client.BatchSize())
}

var ErrFake = errors.New("fake error")

type failingClient struct {
	testdesc.FakeClient
}

func (c *failingClient) Describe(ids []string) (desc.Response[ecs.Disk], error) {
	return desc.Response[ecs.Disk]{}, ErrFake
}

func TestFailures(t *testing.T) {
	t.Parallel()
	client := &failingClient{}
	batcher := NewLowLatency(client, clock.RealClock{}, 2*time.Second, 10)

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	go batcher.Run(ctx)

	disk, err := batcher.Describe(context.Background(), "d-test")
	assert.ErrorIs(t, err, ErrFake)
	assert.Nil(t, disk)
}

func TestReportInefficient(t *testing.T) {
	client := &testdesc.FakeClient{}
	batcher := NewLowLatency(client, clock.RealClock{}, 20*time.Millisecond, 8)

	logger := ktesting.NewLogger(t, ktesting.NewConfig(ktesting.BufferLogs(true)))
	ctx := klog.NewContext(t.Context(), logger)
	go batcher.Run(ctx)

	client.Disks.Store("d-test", &ecs.Disk{
		DiskId:   "d-test",
		DiskName: "mydisk",
	})

	for range 20 {
		_, err := batcher.Describe(ctx, "d-test")
		assert.NoError(t, err)
	}
	assert.Contains(t, logger.GetSink().(ktesting.Underlier).GetBuffer().String(), "Inefficient batching")
}
