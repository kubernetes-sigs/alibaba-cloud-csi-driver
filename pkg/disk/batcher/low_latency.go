package batcher

import (
	"context"
	"fmt"
	"math"
	"slices"
	"time"

	"github.com/go-logr/logr"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/disk/desc"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

type getResponse[T any] struct {
	res T
	err error
}

type getRequest[T any] struct {
	ctx        context.Context
	id         string
	time       time.Time
	resultChan chan getResponse[T]
}

type LowLatency[T any] struct {
	ecsClient desc.Client[T]
	// remove this once we have a ecsClient that can refresh its credentials
	PollHook func() desc.Client[T]

	requestChan chan *getRequest[*T]

	tokens tokenBucket
	clk    clock.WithTicker
}

// NewLowLatency creates a batcher that has low latency when traffic is low, while maintaining high efficiency when traffic is high.
//
// The more frequent incoming requests are, the longer each request will wait to allow batching.
// The waiting time is capped at perRequest.
// The average interval of outgoing batched request traffic is also perRequest, allowing some burst.
//
// If the batch is fully filled, the request is issued immediately, skipping any remaining wait period, and is not limited by perRequest.
// Concurrent outgoing requests are possible if there are too many incomming requests.
func NewLowLatency[T any](ecsClient desc.Client[T], clk clock.WithTicker, perRequest time.Duration, burst int) *LowLatency[T] {
	return &LowLatency[T]{
		ecsClient:   ecsClient,
		requestChan: make(chan *getRequest[*T]),
		tokens: tokenBucket{
			limit: burst, perToken: perRequest,
		},
		clk: clk,
	}
}

func (w *LowLatency[T]) Describe(ctx context.Context, id string) (*T, error) {
	resultChan := make(chan getResponse[*T], 1) // in case this wait is canceled, don't block descBatch()
	w.requestChan <- &getRequest[*T]{
		ctx:        ctx,
		id:         id,
		resultChan: resultChan,
	}
	select {
	case resp := <-resultChan:
		return resp.res, resp.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

type tokenBucket struct {
	zeroAt time.Time

	limit    int
	perToken time.Duration
}

func (tb *tokenBucket) tokenAt(t time.Time) float64 {
	elapsed := t.Sub(tb.zeroAt)
	token := float64(elapsed) / float64(tb.perToken)
	if token > float64(tb.limit) {
		token = float64(tb.limit)
	}
	return token
}

func (tb *tokenBucket) takeAt(t time.Time) {
	elapsed := t.Sub(tb.zeroAt)
	if elapsed >= time.Duration(tb.limit)*tb.perToken {
		tb.zeroAt = t.Add(-time.Duration(tb.limit-1) * tb.perToken)
	} else {
		tb.zeroAt = tb.zeroAt.Add(tb.perToken)
		if tb.zeroAt.After(t) {
			tb.zeroAt = t
		}
	}
}

func (w *LowLatency[T]) Run(ctx context.Context) {
	var timeout <-chan time.Time
	batchSize := w.ecsClient.BatchSize()
	requests := make(map[string][]*getRequest[*T], batchSize)
	logger := klog.FromContext(ctx).WithValues("type", w.ecsClient.Type())

	descBatch := func(t time.Time) {
		w.tokens.takeAt(t)
		go w.descBatch(logger, t, requests)
		requests = make(map[string][]*getRequest[*T], batchSize)
		timeout = nil
	}

	for {
		select {
		case <-ctx.Done():
			sendToAll(requests, getResponse[*T]{err: fmt.Errorf("batcher terminating: %w", ctx.Err())})
			return
		case r := <-w.requestChan:
			klog.FromContext(r.ctx).V(5).Info("enqueued batcher", "type", w.ecsClient.Type())
			t := w.clk.Now()
			r.time = t
			requests[r.id] = append(requests[r.id], r)
			if len(requests) == batchSize {
				logger.V(4).Info("batch full", "n", batchSize)
				descBatch(t)
			} else if timeout == nil {
				// add some artificial delay for better batching
				// the less tokens left, the more we wait
				tokens := w.tokens.tokenAt(t)
				d := time.Duration(math.Pow(0.5, tokens) * float64(w.tokens.perToken))
				timeout = w.clk.After(d)
				logger.V(4).Info("batch waiting", "timeout", d)
			}
		case t := <-timeout:
			logger.V(4).Info("batch timeout", "n", len(requests))
			descBatch(t)
		}
	}
}

func (w *LowLatency[T]) descBatch(logger logr.Logger, t time.Time, requests map[string][]*getRequest[*T]) {
	thisBatch := make([]string, 0, len(requests))
	for id, reqs := range requests {
		if slices.ContainsFunc(reqs, func(r *getRequest[*T]) bool {
			return r.ctx.Err() == nil
		}) {
			thisBatch = append(thisBatch, id)
		}
	}
	if len(thisBatch) == 0 {
		return
	}

	if w.PollHook != nil {
		w.ecsClient = w.PollHook()
	}

	resp, err := w.ecsClient.Describe(thisBatch)
	if err != nil {
		sendToAll(requests, getResponse[*T]{err: err})
		return
	}
	firstTime := t
	for _, r := range resp.Resources {
		id := w.ecsClient.GetID(&r)
		for _, req := range requests[id] {
			if req.time.Before(firstTime) {
				firstTime = req.time
			}
			klog.FromContext(req.ctx).V(4).Info("got resource", "type", w.ecsClient.Type(),
				"wait", t.Sub(req.time), "requestID", resp.RequestID)
			req.resultChan <- getResponse[*T]{res: &r}
		}
		delete(requests, id)
	}
	// Not found
	sendToAll(requests, getResponse[*T]{})
	logger.V(3).Info("got batch", "n", len(thisBatch),
		"requestID", resp.RequestID, "duration", w.clk.Since(t), "wait", t.Sub(firstTime))
}

func sendToAll[T any](requests map[string][]*getRequest[*T], resp getResponse[*T]) {
	for _, reqs := range requests {
		for _, req := range reqs {
			req.resultChan <- resp
		}
	}
}
