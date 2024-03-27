package disk

import (
	"context"
	"encoding/json"
	"slices"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	log "github.com/sirupsen/logrus"
	"k8s.io/utils/clock"
)

const (
	pollInterval = 2 * time.Second
	batchSize    = 100 // DescribeDisks supports up to 100 disk IDs at once
)

type DiskStatusWaiter interface {
	WaitForDiskNextStatus(ctx context.Context, diskID string, pred DiskStatusPredicate) (*ecs.Disk, error)
}

// a function that asserts the disk should quit waiting
type DiskStatusPredicate func(disk *ecs.Disk) bool

type waitRequest struct {
	ctx        context.Context
	diskID     string
	pred       DiskStatusPredicate
	resultChan chan *ecs.Disk
}

type BatchedDiskStatusWaiter struct {
	ecsClient cloud.ECSInterface

	requestChan chan *waitRequest
	requests    map[string][]*waitRequest
	idQueue     []string

	clk clock.WithTicker
}

func NewBatchedDiskStatusWaiter(ecsClient cloud.ECSInterface, clk clock.WithTicker) *BatchedDiskStatusWaiter {
	return &BatchedDiskStatusWaiter{
		ecsClient:   ecsClient,
		requestChan: make(chan *waitRequest),
		requests:    make(map[string][]*waitRequest),
		clk:         clk,
	}
}

func (w *BatchedDiskStatusWaiter) WaitForDiskNextStatus(ctx context.Context, diskID string, pred DiskStatusPredicate) (*ecs.Disk, error) {
	resultChan := make(chan *ecs.Disk, 1) // in case this wait is canceled, don't block poll()
	w.requestChan <- &waitRequest{
		ctx:        ctx,
		diskID:     diskID,
		pred:       pred,
		resultChan: resultChan,
	}
	select {
	case disk := <-resultChan:
		return disk, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (w *BatchedDiskStatusWaiter) enqueueRequest(r *waitRequest) {
	reqs, ok := w.requests[r.diskID]
	w.requests[r.diskID] = append(reqs, r)
	if !ok {
		w.idQueue = append(w.idQueue, r.diskID)
	}
}

func (w *BatchedDiskStatusWaiter) Run(ctx context.Context) {
	var pollChan <-chan time.Time
	var lastPollTime time.Time

	for {
		select {
		case <-ctx.Done():
			return
		case r := <-w.requestChan:
			w.enqueueRequest(r)
			if pollChan == nil {
				// If this is the first request, poll immediately.
				// But don't poll again until pollInterval has passed.
				pollChan = w.clk.After(pollInterval - w.clk.Since(lastPollTime))
			}
		case t := <-pollChan:
			var err error
			w.idQueue, err = w.poll(w.idQueue)
			lastPollTime = t
			if err != nil {
				log.Errorf("failed to poll disk status: %v", err)
			}
			if len(w.requests) == 0 {
				// no more requests, stop polling
				pollChan = nil
			} else {
				pollChan = w.clk.After(pollInterval)
				log.Debugf("scheduled next poll in %v", pollInterval)
			}
		}
	}
}

func (w *BatchedDiskStatusWaiter) gotDisk(id string, disk *ecs.Disk) (done bool) {
	pendingReqs := slices.DeleteFunc(w.requests[id], func(r *waitRequest) bool {
		if r.ctx.Err() != nil {
			return true
		}
		if disk == nil || r.pred(disk) {
			r.resultChan <- disk
			return true
		}
		return false
	})
	if len(pendingReqs) == 0 {
		delete(w.requests, id)
		return true
	} else {
		w.requests[id] = pendingReqs
		return false
	}
}

// poll picks first batchSize waitIDs that are not canceled,
// polls ECS for their status and notify caller if done,
// returns the rest IDs for next poll.
// It removes done requests from w.requests
func (w *BatchedDiskStatusWaiter) poll(waitIDs []string) (next []string, err error) {
	thisBatch := make([]string, 0, batchSize)
	for i, id := range waitIDs {
		reqs := w.requests[id]
		if slices.ContainsFunc(reqs, func(r *waitRequest) bool {
			return r.ctx.Err() == nil
		}) {
			thisBatch = append(thisBatch, id)
		} else {
			delete(w.requests, id)
		}

		if len(thisBatch) == batchSize {
			next = waitIDs[i+1:]
			break
		}
	}
	if len(thisBatch) == 0 {
		return nil, nil
	}

	diskIDBytes, err := json.Marshal(thisBatch)
	if err != nil {
		panic(err)
	}
	req := ecs.CreateDescribeDisksRequest()
	req.DiskIds = string(diskIDBytes)
	req.PageSize = requests.NewInteger(batchSize)
	resp, err := w.ecsClient.DescribeDisks(req)
	if err != nil {
		next = append(next, thisBatch...)
		return
	}
	disks := make(map[string]*ecs.Disk, len(resp.Disks.Disk))
	for i, disk := range resp.Disks.Disk {
		disks[disk.DiskId] = &resp.Disks.Disk[i]
	}
	for _, id := range thisBatch {
		if !w.gotDisk(id, disks[id]) {
			next = append(next, id)
		}
	}
	return
}

type SimpleDiskStatusWaiter struct {
}

func (w *SimpleDiskStatusWaiter) WaitForDiskNextStatus(ctx context.Context, diskID string, pred DiskStatusPredicate) (*ecs.Disk, error) {
	tenantUserUID := GetTenantUserUID(ctx)
	ecsClient, err := getEcsClientByID("", tenantUserUID)
	if err != nil {
		return nil, err
	}
	for {
		disk, err := findDiskByID(diskID, ecsClient)
		if err != nil || disk == nil {
			return disk, err
		}
		if pred(disk) {
			return disk, nil
		}
		select {
		case <-time.After(pollInterval):
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}

func isInstanceAttached(disk *ecs.Disk, instanceID string) bool {
	for _, a := range disk.Attachments.Attachment {
		if a.InstanceId == instanceID {
			return true
		}
	}
	for _, a := range disk.MountInstances.MountInstance {
		if a.InstanceId == instanceID {
			return true
		}
	}
	return false
}
