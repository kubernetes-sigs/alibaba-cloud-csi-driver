package waitstatus

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/stretchr/testify/assert"
)

func TestPoll(t *testing.T) {
	client := &fakeClient{}
	for i := 0; i < 5; i++ {
		diskID := fmt.Sprintf("d%d", i)
		client.disks.Store(diskID, disk(diskID, "Available"))
	}

	waiter := NewBatched(client, nil)
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
	ids := waiter.poll(waiter.idQueue)
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

var ErrFake = errors.New("fake error")

type failingClient struct {
	fakeClient
}

func (c *failingClient) Describe(ids []string) (DescribeResourceResponse[ecs.Disk], error) {
	return DescribeResourceResponse[ecs.Disk]{}, ErrFake
}

func TestPollError(t *testing.T) {
	waiter := NewBatched(&failingClient{}, nil)
	req := &waitRequest[*ecs.Disk]{
		ctx:        context.Background(),
		id:         "d-test",
		pred:       isNextStatus("Detaching", "i1"),
		resultChan: make(chan waitResponse[*ecs.Disk], 1),
	}
	waiter.enqueueRequest(req)
	ids := waiter.poll(waiter.idQueue)
	assert.Empty(t, ids)
	res := <-req.resultChan
	assert.ErrorIs(t, res.err, ErrFake)
	assert.Nil(t, res.res)
}

func TestPollEmpty(t *testing.T) {
	waiter := NewBatched(&failingClient{}, nil)
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	req := &waitRequest[*ecs.Disk]{
		ctx:        ctx,
		id:         "d-test",
		pred:       isNextStatus("Detaching", "i1"),
		resultChan: make(chan waitResponse[*ecs.Disk], 1),
	}
	waiter.enqueueRequest(req)
	ids := waiter.poll(waiter.idQueue)
	assert.Empty(t, ids)
	assert.Empty(t, waiter.requests)
}
