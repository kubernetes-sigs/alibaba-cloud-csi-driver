package disk

import (
	"context"
	"sync"
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
	Attach() slot
	Detach() slot
}

type slot interface {
	Acquire(ctx context.Context) error
	Release()
}

type serialADSlot struct {
	// slot is a buffered channel with size 1.
	// The buffer is filled if and only if an attach or detach is in progress.
	slot chan struct{}

	// highPriorityChannel is a channel without buffer.
	// detach requests are fulfilled from this channel first.
	highPriorityChannel chan struct{}
}

type serialAD_DetachSlot struct{ *serialADSlot }
type serialAD_AttachSlot struct{ *serialADSlot }

func (s serialAD_DetachSlot) Acquire(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	select {
	case s.highPriorityChannel <- struct{}{}:
		return nil
	case s.slot <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s serialAD_AttachSlot) Acquire(ctx context.Context) error {
	if ctx.Err() != nil {
		return ctx.Err()
	}
	select {
	case s.slot <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
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

func (s *serialADSlot) Detach() slot { return serialAD_DetachSlot{s} }
func (s *serialADSlot) Attach() slot { return serialAD_AttachSlot{s} }

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
		return ctx.Err()
	}
}

func (s maxConcurrentSlot) Release() {
	<-s.slots
}

type parallelSlot struct{}

func (s parallelSlot) GetSlotFor(node string) adSlot { return s }

func (parallelSlot) Detach() slot { return noOpSlot{} }
func (parallelSlot) Attach() slot { return noOpSlot{} }

type noOpSlot struct{}

func (noOpSlot) Acquire(ctx context.Context) error { return ctx.Err() }
func (noOpSlot) Release()                          {}

type independentSlot struct {
	attach slot
	detach slot
}

func (s *independentSlot) Detach() slot { return s.detach }
func (s *independentSlot) Attach() slot { return s.attach }

func makeOneSide(concurrency int) slot {
	if concurrency == 0 {
		return noOpSlot{}
	}
	return newMaxConcurrentSlot(concurrency)
}

// NewSlots returns a new AttachDetachSlots that allows up to
// detachConcurrency detach operations and attachConcurrency attach operations in parallel.
// 0 means no limitation on concurrency.
// As a special case, if both values are 1, only one of them can be in progress at a time.
func NewSlots(detachConcurrency, attachConcurrency int) AttachDetachSlots {
	if detachConcurrency == 0 && attachConcurrency == 0 {
		return parallelSlot{}
	}
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
