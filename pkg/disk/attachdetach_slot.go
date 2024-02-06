package disk

import (
	"context"
	"sync"
)

type AttachDetachSlots struct {
	mu          sync.RWMutex
	nodes       map[string]slot
	slotFactory func() slot
}

func (a *AttachDetachSlots) GetSlotFor(node string) slot {
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
	s := a.slotFactory()
	a.nodes[node] = s
	return s
}

type slot interface {
	AquireDetach(ctx context.Context) error
	AquireAttach(ctx context.Context) error
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

func (s *serialSlot) AquireDetach(ctx context.Context) error {
	select {
	case s.highPriorityChannel <- struct{}{}:
		return nil
	case s.slot <- struct{}{}:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *serialSlot) AquireAttach(ctx context.Context) error {
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

func NewSerialAttachDetachSlots() AttachDetachSlots {
	return AttachDetachSlots{
		slotFactory: func() slot {
			return &serialSlot{
				highPriorityChannel: make(chan struct{}),
				slot:                make(chan struct{}, 1),
			}
		},
		nodes: map[string]slot{},
	}
}

type parallelSlot struct{}

func (s parallelSlot) AquireDetach(ctx context.Context) error { return nil }
func (s parallelSlot) AquireAttach(ctx context.Context) error { return nil }
func (s parallelSlot) Release()                               {}

func NewParallelAttachDetachSlots() AttachDetachSlots {
	return AttachDetachSlots{
		slotFactory: func() slot {
			return parallelSlot{}
		},
	}
}
