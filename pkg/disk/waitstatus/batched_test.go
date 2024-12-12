package waitstatus

import (
	"context"
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	testdesc "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc/testing"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/clock"
	testclock "k8s.io/utils/clock/testing"
)

func TestPoll(t *testing.T) {
	client := &testdesc.FakeClient{}
	for i := 0; i < 5; i++ {
		diskID := fmt.Sprintf("d%d", i)
		client.Disks.Store(diskID, disk(diskID, "Available"))
	}

	waiter := NewBatched(client, clock.RealClock{}, pollInterval, 2*pollInterval)
	go func() {
		waiter.processFeedback(<-waiter.feedback)
	}()

	sentReqs := make([]*waitRequest[*ecs.Disk], 0, 110) // large enough to trigger multiple batches
	for i := 0; i < cap(sentReqs); i++ {
		r := &waitRequest[*ecs.Disk]{
			ctx:        context.Background(),
			id:         fmt.Sprintf("d%d", i),
			pred:       isNextStatus("Detaching", "i1"),
			resultChan: make(chan waitResponse[*ecs.Disk], 1),
		}
		if i == 2 {
			// make this request already canceled
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			r.ctx = ctx
		}
		if i == 20 {
			// make a repeated request
			r.id = "d30"
		}
		waiter.enqueueRequest(r)
		sentReqs = append(sentReqs, r)
	}
	expectedIDs := waiter.idQueue[101:]     // 100 polled, 1 canceled
	assert.Equal(t, "d102", expectedIDs[0]) // 1 repeated request
	ids := waiter.poll(time.Now(), waiter.idQueue)
	assert.Equal(t, expectedIDs, ids)

	for i, req := range sentReqs[:102] {
		if i == 2 { // this request was canceled
			continue
		}
		resp := <-req.resultChan
		assert.NoError(t, resp.err)
		if i < 5 {
			expectedID := fmt.Sprintf("d%d", i)
			assert.Equal(t, expectedID, resp.res.DiskId)
		} else {
			assert.Nil(t, resp.res)
		}
	}
}

func TestPollError(t *testing.T) {
	waiter := NewBatched(&failingClient{}, clock.RealClock{}, pollInterval, 2*pollInterval)
	go func() {
		waiter.processFeedback(<-waiter.feedback)
	}()

	req := &waitRequest[*ecs.Disk]{
		ctx:        context.Background(),
		id:         "d-test",
		pred:       isNextStatus("Detaching", "i1"),
		resultChan: make(chan waitResponse[*ecs.Disk], 1),
	}
	waiter.enqueueRequest(req)
	ids := waiter.poll(time.Now(), waiter.idQueue)
	assert.Empty(t, ids)
	res := <-req.resultChan
	assert.ErrorIs(t, res.err, ErrFake)
	assert.Nil(t, res.res)
}

func TestPollEmpty(t *testing.T) {
	waiter := NewBatched(&testdesc.FakeClient{}, clock.RealClock{}, pollInterval, 2*pollInterval)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	req := &waitRequest[*ecs.Disk]{
		ctx:        ctx,
		id:         "d-test",
		pred:       isNextStatus("Detaching", "i1"),
		resultChan: make(chan waitResponse[*ecs.Disk], 1),
	}
	waiter.enqueueRequest(req)
	ids := waiter.poll(time.Now(), waiter.idQueue)
	assert.Empty(t, ids)
	assert.Empty(t, waiter.requests)
}

type noBatchClient struct {
	testdesc.FakeClient
}

func (*noBatchClient) BatchSize() int {
	return 1
}

func TestWaitTimeUpperBound(t *testing.T) {
	clk := testclock.NewFakeClock(time.Now())
	waiter := NewBatched(&noBatchClient{}, clk, 5*time.Second, 7*time.Second)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go waiter.Run(ctx)

	var iter atomic.Int32
	for i := range 3 {
		go func() {
			waiter.WaitFor(ctx, fmt.Sprintf("d-test-%d", i), isNextStatus("Detaching", "i1"))
			iter.Add(1)
		}()
	}

	time.Sleep(10 * time.Millisecond)
	clk.Step(1 * time.Microsecond)
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, int32(1), iter.Load())

	clk.Step(5 * time.Second)
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, int32(2), iter.Load())

	clk.Step(2 * time.Second)
	time.Sleep(10 * time.Millisecond)
	assert.Equal(t, int32(3), iter.Load())
}
