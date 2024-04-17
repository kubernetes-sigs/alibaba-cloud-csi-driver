package disk

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestDetachPriority(t *testing.T) {
	slots := NewSlots(true, true)
	s := slots.GetSlotFor("node1")
	seq := 0

	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			as := s.Attach()
			if err := as.Aquire(context.Background()); err != nil {
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
	if err := ds.Aquire(context.Background()); err != nil {
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
	slots := NewSlots(true, true)
	nodeName := "node1"
	for i := 0; i < b.N; i++ {
		slots.GetSlotFor(nodeName)
	}
}

func TestParallelGetSlot(t *testing.T) {
	slots := NewSlots(false, false)
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

func TestCancel(t *testing.T) {
	slots := NewSlots(true, true)
	s := slots.GetSlotFor("node1")
	ctx, cancel := context.WithCancel(context.Background())
	err := s.Detach().Aquire(ctx) // occupy the slot
	if err != nil {
		t.Fatal(err)
	}

	cancel()
	if err := s.Detach().Aquire(ctx); err != context.Canceled {
		t.Fatalf("Expected context.Canceled, got %v", err)
	}
	if err := s.Attach().Aquire(ctx); err != context.Canceled {
		t.Fatalf("Expected context.Canceled, got %v", err)
	}
}

func TestSerialDetach(t *testing.T) {
	slots := NewSlots(true, false)
	s := slots.GetSlotFor("node1")

	ctx := context.Background()
	as := s.Attach()
	err := as.Aquire(ctx)
	if err != nil {
		t.Fatal(err)
	}

	// Attach should not block detach
	ds := s.Detach()
	err = ds.Aquire(ctx)
	if err != nil {
		t.Fatal(err)
	}

	ds.Release()
	as.Release()
}

// run this with go test -race
func TestSerialDetach_NoRace(t *testing.T) {
	s := NewSlots(true, false).GetSlotFor("node1").Detach()
	ctx := context.Background()
	wg := sync.WaitGroup{}
	wg.Add(2)
	state := -1
	for i := 0; i < 2; i++ {
		go func(i int) {
			s.Aquire(ctx)
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
