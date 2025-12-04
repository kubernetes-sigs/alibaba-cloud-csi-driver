package disk

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"
)

type AttachDetachSlots interface {
	GetSlotFor(node string) adSlot
}

type PerNodeSlots struct {
	mu       sync.RWMutex
	nodes    map[string]adSlot
	makeSlot func() adSlot
}

func NewPerNodeSlots(makeSlot func() adSlot) AttachDetachSlots {
	return &PerNodeSlots{
		nodes:    map[string]adSlot{},
		makeSlot: makeSlot,
	}
}

func (a *PerNodeSlots) GetSlotFor(node string) adSlot {
	a.mu.RLock()
	if s, ok := a.nodes[node]; ok {
		a.mu.RUnlock()
		return s
	}
	a.mu.RUnlock()

	a.mu.Lock()
	defer a.mu.Unlock()
	if s, ok := a.nodes[node]; ok {
		return s
	}
	s := a.makeSlot()
	a.nodes[node] = s
	return s
}

type adSlot interface {
	Attach() blockableSlot
	Detach() blockableSlot
}

type slot interface {
	Acquire(ctx context.Context) error
	Release()
}

type blockableSlot interface {
	slot

	// Block blocks all the future Acquire() calls until the given time.
	// Use this if the node is in an unusual state, and is not expected to recover if retried immediately.
	Block(until time.Time)
}

type blockable struct {
	until atomic.Pointer[time.Time]
}

func (s *blockable) Block(until time.Time) {
	for {
		old := s.until.Load()
		if old != nil && old.After(until) {
			return
		}
		swapped := s.until.CompareAndSwap(old, &until)
		if swapped {
			return
		}
	}
}

func (s *blockable) Acquire(ctx context.Context) error {
	for {
		until := s.until.Load()
		if until == nil {
			return nil
		}
		select {
		case <-time.After(-time.Since(*until)):
			swapped := s.until.CompareAndSwap(until, nil)
			if swapped {
				return nil
			}
			// if not swapped, the block time is extended while we are waiting or already cleared by another concurrent Acquire()
			// we will see in the next iteration
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

type independentBlockableSlot[TSlot slot] struct {
	slot TSlot
	blockable
}

func (s *independentBlockableSlot[TSlot]) Acquire(ctx context.Context) error {
	err := s.slot.Acquire(ctx)
	if err != nil {
		return err
	}
	return s.blockable.Acquire(ctx)
}

func (s *independentBlockableSlot[TSlot]) Release() {
	s.slot.Release()
}

func newBlockable[TSlot slot](slot TSlot) *independentBlockableSlot[TSlot] {
	return &independentBlockableSlot[TSlot]{slot: slot}
}

type serialADSlot struct {
	// slot is a buffered channel with size 1.
	// The buffer is filled if and only if an attach or detach is in progress.
	slot chan struct{}

	// highPriorityChannel is a channel without buffer.
	// detach requests are fulfilled from this channel first.
	highPriorityChannel chan struct{}

	blockable
}

type serialAD_DetachSlot struct{ *serialADSlot }
type serialAD_AttachSlot struct{ *serialADSlot }

func (s serialAD_DetachSlot) Acquire(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	select {
	case s.highPriorityChannel <- struct{}{}:
	case s.slot <- struct{}{}:
	case <-ctx.Done():
		return maybeWaitingAD(ctx.Err())
	}
	return s.blockable.Acquire(ctx)
}

func (s serialAD_AttachSlot) Acquire(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	select {
	case s.slot <- struct{}{}:
	case <-ctx.Done():
		return maybeWaitingAD(ctx.Err())
	}
	return s.blockable.Acquire(ctx)
}

func (s *serialADSlot) Release() {
	select {
	// wake up a high priority request first
	case <-s.highPriorityChannel:
	default:
		// if no high priority request waiting, free the slot
		<-s.slot
	}
}

func (s *serialADSlot) Detach() blockableSlot { return serialAD_DetachSlot{s} }
func (s *serialADSlot) Attach() blockableSlot { return serialAD_AttachSlot{s} }

type maxConcurrentSlot struct {
	slots chan struct{}
}

func newMaxConcurrentSlot(maxConcurrency int) maxConcurrentSlot {
	return maxConcurrentSlot{
		slots: make(chan struct{}, maxConcurrency),
	}
}

func (s maxConcurrentSlot) Acquire(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	select {
	case s.slots <- struct{}{}:
		return nil
	case <-ctx.Done():
		return maybeWaitingAD(ctx.Err())
	}
}

func (s maxConcurrentSlot) Release() {
	<-s.slots
}

type noOpSlot struct{}

func (noOpSlot) Acquire(ctx context.Context) error { return ctx.Err() }
func (noOpSlot) Release()                          {}

type independentSlot struct {
	attach blockableSlot
	detach blockableSlot
}

func (s *independentSlot) Detach() blockableSlot { return s.detach }
func (s *independentSlot) Attach() blockableSlot { return s.attach }

func makeOneSide(concurrency int) blockableSlot {
	if concurrency == 0 {
		return newBlockable(noOpSlot{})
	}
	return newBlockable(newMaxConcurrentSlot(concurrency))
}

// NewSlots returns a new AttachDetachSlots that allows up to
// detachConcurrency detach operations and attachConcurrency attach operations in parallel.
// 0 means no limitation on concurrency.
// As a special case, if both values are 1, only one of them can be in progress at a time.
func NewSlots(detachConcurrency, attachConcurrency int) AttachDetachSlots {
	var makeSlot func() adSlot
	if detachConcurrency == 1 && attachConcurrency == 1 {
		makeSlot = func() adSlot {
			return &serialADSlot{
				highPriorityChannel: make(chan struct{}),
				slot:                make(chan struct{}, 1),
			}
		}
	} else {
		makeSlot = func() adSlot {
			return &independentSlot{
				attach: makeOneSide(attachConcurrency),
				detach: makeOneSide(detachConcurrency),
			}
		}
	}
	return NewPerNodeSlots(makeSlot)
}

type waitingAD struct{}

func (waitingAD) Error() string {
	return "still waiting for other disk(s) to finish attach/detach"
}

func (waitingAD) Is(target error) bool {
	return target == context.DeadlineExceeded
}

func maybeWaitingAD(err error) error {
	if errors.Is(err, context.DeadlineExceeded) {
		return waitingAD{}
	}
	return err
}
