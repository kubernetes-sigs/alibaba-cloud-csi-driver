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

func (s serialAD_DetachSlot) Aquire(ctx context.Context) error {
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

func (s serialAD_AttachSlot) Aquire(ctx context.Context) error {
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

type parallelSlot struct{}

func (s parallelSlot) GetSlotFor(node string) adSlot { return s }

func (parallelSlot) Detach() slot { return noOpSlot{} }
func (parallelSlot) Attach() slot { return noOpSlot{} }

type noOpSlot struct{}

func (noOpSlot) Aquire(ctx context.Context) error { return ctx.Err() }
func (noOpSlot) Release()                         {}

type serialOneDirSlot struct {
	parallelSlot
	serial serialSlot
}

type serialDetachSlot serialOneDirSlot
type serialAttachSlot serialOneDirSlot

func (s serialDetachSlot) Detach() slot { return s.serial }
func (s serialAttachSlot) Attach() slot { return s.serial }

type serialSlot struct {
	// slot is a buffered channel with size 1.
	// The buffer is filled if and only if a detach is in progress.
	slot chan struct{}
}

func (s serialSlot) Aquire(ctx context.Context) error {
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

func (s serialSlot) Release() {
	<-s.slot
}

func makeOneSideSlot() serialOneDirSlot {
	return serialOneDirSlot{
		serial: serialSlot{
			slot: make(chan struct{}, 1),
		},
	}
}

func NewSlots(serialDetach, serialAttach bool) AttachDetachSlots {
	if !serialAttach && !serialDetach {
		return parallelSlot{}
	}
	var makeSlot func() adSlot
	if serialDetach && serialAttach {
		makeSlot = func() adSlot {
			return &serialADSlot{
				highPriorityChannel: make(chan struct{}),
				slot:                make(chan struct{}, 1),
			}
		}
	} else if serialDetach {
		makeSlot = func() adSlot {
			return serialDetachSlot(makeOneSideSlot())
		}
	} else if serialAttach {
		makeSlot = func() adSlot {
			return serialAttachSlot(makeOneSideSlot())
		}
	} else {
		panic("unreachable")
	}
	return NewPerNodeSlots(makeSlot)
}
