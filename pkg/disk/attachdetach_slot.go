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
	Aquire(ctx context.Context) error
	Release()
}

type serialSlot struct {
	// slot is a buffered channel with size 1.
	// The buffer is filled if and only if an attach or detach is in progress.
	slot chan struct{}

	// highPriorityChannel is a channel without buffer.
	// detach requests are fulfilled from this channel first.
	highPriorityChannel chan struct{}
}

type serialAD_DetachSlot struct{ *serialSlot }
type serialAD_AttachSlot struct{ *serialSlot }

func (s serialAD_DetachSlot) Aquire(ctx context.Context) error {
	select {
	case s.highPriorityChannel <- struct{}{}:
		return nil
	case s.slot <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s serialAD_AttachSlot) Aquire(ctx context.Context) error {
	select {
	case s.slot <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *serialSlot) Release() {
	select {
	// wake up a high priority request first
	case <-s.highPriorityChannel:
	default:
		// if no high priority request waiting, free the slot
		<-s.slot
	}
}

func (s *serialSlot) Detach() slot { return serialAD_DetachSlot{s} }
func (s *serialSlot) Attach() slot { return serialAD_AttachSlot{s} }

func NewSerialAttachDetachSlots() AttachDetachSlots {
	return NewPerNodeSlots(func() adSlot {
		return &serialSlot{
			highPriorityChannel: make(chan struct{}),
			slot:                make(chan struct{}, 1),
		}
	})
}

type parallelSlot struct{}

func (s parallelSlot) GetSlotFor(node string) adSlot { return s }

func (parallelSlot) Detach() slot { return noOpSlot{} }
func (parallelSlot) Attach() slot { return noOpSlot{} }

type noOpSlot struct{}

func (noOpSlot) Aquire(ctx context.Context) error { return nil }
func (noOpSlot) Release()                         {}

func NewParallelAttachDetachSlots() AttachDetachSlots {
	return parallelSlot{}
}

type serialDetachSlot struct {
	parallelSlot
	detach serialDetach_DetachSlot
}

func (s *serialDetachSlot) Detach() slot {
	return s.detach
}

type serialDetach_DetachSlot struct {
	// slot is a buffered channel with size 1.
	// The buffer is filled if and only if a detach is in progress.
	slot chan struct{}
}

func (s serialDetach_DetachSlot) Aquire(ctx context.Context) error {
	select {
	case s.slot <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s serialDetach_DetachSlot) Release() {
	<-s.slot
}

func NewSerialDetachSlots() AttachDetachSlots {
	return NewPerNodeSlots(func() adSlot {
		return &serialDetachSlot{
			detach: serialDetach_DetachSlot{
				slot: make(chan struct{}, 1),
			},
		}
	})
}
