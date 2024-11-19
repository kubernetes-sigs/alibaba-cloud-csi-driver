package disk

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDetachPriority(t *testing.T) {
	slots := NewSlots(1, 1)
	s := slots.GetSlotFor("node1")
	seq := 0

	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			as := s.Attach()
			if err := as.Acquire(context.Background()); err != nil {
				t.Error(err)
				return
			}
			t.Log("Attach seq", seq)
			seq += 1
			time.Sleep(200 * time.Millisecond)
			as.Release()
			wg.Done()
		}()
	}
	time.Sleep(100 * time.Millisecond)
	ds := s.Detach()
	if err := ds.Acquire(context.Background()); err != nil {
		t.Fatal(err)
	}
	if seq != 1 {
		// detach should go after the first attach, but before the second
		t.Fatalf("Expected seq to be 1, got %d", seq)
	}
	ds.Release()
	wg.Wait()
}

func BenchmarkGetSlot(b *testing.B) {
	slots := NewSlots(1, 1)
	nodeName := "node1"
	for i := 0; i < b.N; i++ {
		slots.GetSlotFor(nodeName)
	}
}

func TestParallelGetSlot(t *testing.T) {
	slots := NewSlots(0, 0)
	nodeName := "node1"
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			slots.GetSlotFor(nodeName)
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestCancelWaiting(t *testing.T) {
	testSlot := func(t *testing.T, s adSlot) {
		ctx, cancel := context.WithCancel(context.Background())
		errs := make(chan error)
		go func() {
			errs <- s.Detach().Acquire(ctx)
		}()
		go func() {
			errs <- s.Attach().Acquire(ctx)
		}()
		time.Sleep(100 * time.Millisecond) // ensure we enter waiting state
		cancel()
		for i := 0; i < 2; i++ {
			select {
			case err := <-errs:
				if err != context.Canceled {
					t.Fatalf("Expected context.Canceled, got %v", err)
				}
			case <-time.After(5 * time.Second):
				t.Fatal("Timeout waiting for error")
			}
		}
	}
	t.Run("serial", func(t *testing.T) {
		s := NewSlots(1, 1).GetSlotFor("node1")
		err := s.Detach().Acquire(context.Background()) // occupy the slot
		assert.NoError(t, err)
		testSlot(t, s)
	})
	t.Run("independent", func(t *testing.T) {
		s := &independentSlot{
			attach: newBlockable(newMaxConcurrentSlot(1)),
			detach: newBlockable(newMaxConcurrentSlot(1)),
		}
		assert.NoError(t, s.attach.Acquire(context.Background()))
		assert.NoError(t, s.detach.Acquire(context.Background()))
		testSlot(t, s)
	})
}

func TestCancelNoOccupy(t *testing.T) {
	testSlot := func(t *testing.T, slots AttachDetachSlots) {
		s := slots.GetSlotFor("node1")
		ctx, cancel := context.WithCancel(context.Background())
		cancel()

		if err := s.Detach().Acquire(ctx); err != context.Canceled {
			t.Fatalf("Expected context.Canceled, got %v", err)
		}
		if err := s.Attach().Acquire(ctx); err != context.Canceled {
			t.Fatalf("Expected context.Canceled, got %v", err)
		}
	}
	t.Run("serial", func(t *testing.T) {
		testSlot(t, NewSlots(1, 1))
	})
	t.Run("parallel", func(t *testing.T) {
		testSlot(t, NewSlots(0, 0))
	})
	t.Run("maxConcurrent", func(t *testing.T) {
		testSlot(t, NewSlots(2, 1))
	})
}

func TestSerialDetach(t *testing.T) {
	slots := NewSlots(1, 0)
	s := slots.GetSlotFor("node1")

	ctx := context.Background()
	as := s.Attach()
	err := as.Acquire(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Attach should not block detach
	ds := s.Detach()
	err = ds.Acquire(ctx)
	if err != nil {
		t.Fatal(err)
	}

	ds.Release()
	as.Release()
}

// run this with go test -race
func TestSerialDetach_NoRace(t *testing.T) {
	s := NewSlots(1, 0).GetSlotFor("node1").Detach()
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(2)
	state := -1
	for i := 0; i < 2; i++ {
		go func(i int) {
			s.Acquire(ctx)
			state = i
			s.Release()
			wg.Done()
		}(i)
	}
	wg.Wait()
	if state == -1 {
		t.Fatal("state not updated")
	}
}

func TestWaitingADError(t *testing.T) {
	s := NewSlots(1, 0).GetSlotFor("node1").Detach()
	ctx := context.Background()
	assert.NoError(t, s.Acquire(ctx))

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()
	err := s.Acquire(ctx)
	assert.ErrorIs(t, err, waitingAD{})
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}

func TestBlock(t *testing.T) {
	s := NewSlots(1, 1).GetSlotFor("node1")

	now := time.Now()
	until := now.Add(300 * time.Millisecond)
	// regardless of the order, use the latest timeout
	s.Attach().Block(until.Add(-100 * time.Millisecond))
	s.Attach().Block(until)
	s.Attach().Block(until.Add(-200 * time.Millisecond))

	times := make(chan time.Time)
	// multiple concurrent Acquire() should be fine. All of them are controlled by the same timeout
	go func() {
		assert.NoError(t, s.Attach().Acquire(context.Background()))
		s.Attach().Release()
		times <- time.Now()
	}()
	go func() {
		assert.NoError(t, s.Detach().Acquire(context.Background()))
		s.Detach().Release()
		times <- time.Now()
	}()
	for range 2 {
		select {
		case acquired := <-times:
			assert.WithinDuration(t, until, acquired, 50*time.Millisecond)
		case <-time.After(500 * time.Millisecond):
			t.Fatal("Timeout waiting for Acquire")
		}
	}
}

func TestBlockCancel(t *testing.T) {
	s := NewSlots(1, 1).GetSlotFor("node1")
	s.Attach().Block(time.Now().Add(10 * time.Second))

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(100 * time.Millisecond)
		cancel()
	}()
	assert.ErrorIs(t, s.Attach().Acquire(ctx), context.Canceled)
}
