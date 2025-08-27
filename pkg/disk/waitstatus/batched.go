package waitstatus

import (
	"context"
	"slices"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc"
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
	time       time.Time
	pred       StatusPredicate[T]
	resultChan chan waitResponse[T]
}

type responseFeedback[T any] struct {
	ids []string
	res desc.Response[T]
	err error
}

type Batched[T any] struct {
	ecsClient desc.Client[T]

	feedback    chan responseFeedback[T]
	requestChan chan *waitRequest[*T]
	requests    map[string][]*waitRequest[*T]
	idQueue     []string

	defaultInterval time.Duration
	maxInterval     time.Duration

	clk clock.WithTicker
}

// NewBatched creates a new batched status waiter.
//
// It aggregates multiple resources into a single OpenAPI request.
// The interval between the adjacent OpenAPI requests is defaultInterval.
// However, if a resource has not been polled for longer than maxInterval since it was enqueued or last polled,
// the interval is automatically shortened.
//
// With these rules, the poll interval for a specific resource is defaultInterval when the load is low,
// increases as the load increases, and caps at maxInterval.
func NewBatched[T any](ecsClient desc.Client[T], clk clock.WithTicker, defaultInterval, maxInterval time.Duration) *Batched[T] {
	return &Batched[T]{
		ecsClient:   ecsClient,
		requestChan: make(chan *waitRequest[*T]),
		requests:    make(map[string][]*waitRequest[*T]),
		feedback:    make(chan responseFeedback[T]),

		defaultInterval: defaultInterval,
		maxInterval:     maxInterval,
		clk:             clk,
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
	r.time = w.clk.Now()
	reqs, ok := w.requests[r.id]
	w.requests[r.id] = append(reqs, r)
	if !ok {
		w.idQueue = append(w.idQueue, r.id)
	}
}

func (w *Batched[T]) Run(ctx context.Context) {
	var pollChan <-chan time.Time
	var lastPollTime time.Time
	logger := klog.FromContext(ctx).WithValues("type", w.ecsClient.Type())

	for {
		if pollChan == nil && len(w.idQueue) > 0 {
			// If this is the first request, poll immediately.
			nextPoll := lastPollTime.Add(w.defaultInterval)
			upperBound := w.requests[w.idQueue[0]][0].time.Add(w.maxInterval)
			if upperBound.Before(nextPoll) {
				nextPoll = upperBound
			}
			d := -w.clk.Since(nextPoll)
			pollChan = w.clk.After(d)
			logger.V(3).Info("poll scheduled", "after", d, "queueDepth", len(w.idQueue))
		}

		select {
		case <-ctx.Done():
			return
		case r := <-w.requestChan:
			klog.FromContext(r.ctx).V(5).Info("enqueued waitstatus", "type", w.ecsClient.Type())
			w.enqueueRequest(r)
		case r := <-w.feedback:
			next := w.processFeedback(r)
			w.idQueue = append(w.idQueue, next...)
			logger.V(4).Info("poll response processed", "queueDepth", len(w.idQueue), "requeue", len(next))
		case t := <-pollChan:
			logger.V(4).Info("starting poll", "queueDepth", len(w.idQueue))
			w.idQueue = w.poll(t, w.idQueue)
			lastPollTime = t
			pollChan = nil
		}
	}
}

func (w *Batched[T]) got(t time.Time, id string, resource *T, requestID string) (done bool) {
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
		r.time = t
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

func (w *Batched[T]) processFeedback(fb responseFeedback[T]) (next []string) {
	if fb.err != nil {
		for _, id := range fb.ids {
			for _, r := range w.requests[id] {
				r.resultChan <- waitResponse[*T]{err: fb.err}
			}
			delete(w.requests, id)
		}
		return nil
	}
	resp := fb.res
	resources := make(map[string]*T, len(resp.Resources))
	for i := range resp.Resources {
		res := &resp.Resources[i]
		resources[w.ecsClient.GetID(res)] = res
	}
	t := w.clk.Now()
	for _, id := range fb.ids {
		if !w.got(t, id, resources[id], resp.RequestID) {
			next = append(next, id)
		}
	}
	return
}

// poll picks first batchSize waitIDs that are not canceled,
// polls ECS for their status and notify caller if done,
// returns the rest IDs for next poll.
// It removes canceled requests from w.requests.
// poll results are sent to w.feedback.
func (w *Batched[T]) poll(t time.Time, waitIDs []string) []string {
	var next []string
	batchSize := w.ecsClient.BatchSize()
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
	t0 := w.requests[thisBatch[0]][0].time
	interval := t.Sub(t0)
	go func() {
		client := w.ecsClient
		resp, err := client.Describe(thisBatch)
		w.feedback <- responseFeedback[T]{
			ids: thisBatch,
			res: resp,
			err: err,
		}
		klog.V(2).InfoS("polled batch", "type", client.Type(), "n", len(thisBatch),
			"interval", interval, "duration", w.clk.Since(t), "requestID", resp.RequestID)
	}()
	return next
}
