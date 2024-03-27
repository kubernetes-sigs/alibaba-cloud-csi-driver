package waitstatus

import (
	"context"
	"slices"
	"time"

	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

type waitResponse[T any] struct {
	res T
	err error
}

type waitRequest[T any] struct {
	ctx        context.Context
	id         string
	pred       StatusPredicate[T]
	resultChan chan waitResponse[T]
}

type Batched[T any] struct {
	ecsClient ECSDescribeResources[T]

	requestChan chan *waitRequest[*T]
	requests    map[string][]*waitRequest[*T]
	idQueue     []string

	clk clock.WithTicker
}

func NewBatched[T any](ecsClient ECSDescribeResources[T], clk clock.WithTicker) *Batched[T] {
	return &Batched[T]{
		ecsClient:   ecsClient,
		requestChan: make(chan *waitRequest[*T]),
		requests:    make(map[string][]*waitRequest[*T]),
		clk:         clk,
	}
}

func (w *Batched[T]) WaitFor(ctx context.Context, id string, pred StatusPredicate[*T]) (*T, error) {
	resultChan := make(chan waitResponse[*T], 1) // in case this wait is canceled, don't block poll()
	w.requestChan <- &waitRequest[*T]{
		ctx:        ctx,
		id:         id,
		pred:       pred,
		resultChan: resultChan,
	}
	select {
	case resp := <-resultChan:
		return resp.res, resp.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func (w *Batched[T]) enqueueRequest(r *waitRequest[*T]) {
	reqs, ok := w.requests[r.id]
	w.requests[r.id] = append(reqs, r)
	if !ok {
		w.idQueue = append(w.idQueue, r.id)
	}
}

func (w *Batched[T]) Run(ctx context.Context) {
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
			w.idQueue = w.poll(w.idQueue)
			lastPollTime = t
			if len(w.requests) == 0 {
				// no more requests, stop polling
				pollChan = nil
			} else {
				pollChan = w.clk.After(pollInterval)
			}
		}
	}
}

func (w *Batched[T]) got(id string, resource *T, requestID string) (done bool) {
	pendingReqs := slices.DeleteFunc(w.requests[id], func(r *waitRequest[*T]) bool {
		if r.ctx.Err() != nil {
			return true
		}
		logger := klog.FromContext(r.ctx)
		if resource == nil || r.pred(resource) {
			logger.V(4).Info("reached desired status", "type", w.ecsClient.Type(), "requestID", requestID)
			r.resultChan <- waitResponse[*T]{res: resource}
			return true
		}
		logger.V(4).Info("status not reached, re-queue", "type", w.ecsClient.Type(), "requestID", requestID)
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
func (w *Batched[T]) poll(waitIDs []string) []string {
	var next []string
	thisBatch := make([]string, 0, batchSize)
	for i, id := range waitIDs {
		reqs := w.requests[id]
		if slices.ContainsFunc(reqs, func(r *waitRequest[*T]) bool {
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
		return nil
	}

	resp, err := w.ecsClient.Describe(thisBatch)
	if err != nil {
		for _, id := range thisBatch {
			for _, r := range w.requests[id] {
				r.resultChan <- waitResponse[*T]{err: err}
			}
		}
		return next
	}
	resources := make(map[string]*T, len(resp.Resources))
	for i := range resp.Resources {
		res := &resp.Resources[i]
		resources[w.ecsClient.GetID(res)] = res
	}
	for _, id := range thisBatch {
		if !w.got(id, resources[id], resp.RequestID) {
			next = append(next, id)
		}
	}
	klog.V(3).InfoS("polling batch", "type", w.ecsClient.Type(), "n", len(thisBatch),
		"interval", pollInterval, "remaining", len(next), "requestID", resp.RequestID)
	return next
}
