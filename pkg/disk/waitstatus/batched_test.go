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
			resp := <-req.resultChan
			expectedID := fmt.Sprintf("d%d", i)
			assert.Equal(t, expectedID, resp.DiskId)
		} else {
			resp := <-req.resultChan
			assert.Nil(t, resp)
		}
	}
}

var ErrFake = errors.New("fake error")

type failingClient struct{}

func (c *failingClient) DescribeDisks(*ecs.DescribeDisksRequest) (*ecs.DescribeDisksResponse, error) {
	return nil, ErrFake
}

func TestPollError(t *testing.T) {

	waiter := NewBatched(&failingClient{}, nil)
	req := &waitRequest{
		ctx:        context.Background(),
		diskID:     "d-test",
		pred:       isNextStatus("Detaching", "i1"),
		resultChan: make(chan *ecs.Disk, 1),
	}
	waiter.enqueueRequest(req)
	ids, err := waiter.poll(waiter.idQueue)
	assert.Equal(t, ErrFake, err)
	assert.Empty(t, req.resultChan)
	assert.Equal(t, []string{"d-test"}, ids)
}

func TestPollEmpty(t *testing.T) {
	waiter := NewBatched(&failingClient{}, nil)
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
