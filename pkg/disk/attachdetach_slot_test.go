package disk

import (
	"context"
	"sync"
	"testing"
	"time"
)

func TestDetachPriority(t *testing.T) {
	slots := NewSerialAttachDetachSlots()
	s := slots.GetSlotFor("node1")
	seq := 0

	wg := sync.WaitGroup{}
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func() {
			if err := s.AquireAttach(context.Background()); err != nil {
				t.Error(err)
				return
			}
			t.Log("Attach seq", seq)
			seq += 1
			time.Sleep(200 * time.Millisecond)
			s.Release()
			wg.Done()
		}()
	}
	time.Sleep(100 * time.Millisecond)
	if err := s.AquireDetach(context.Background()); err != nil {
		t.Fatal(err)
	}
	if seq != 1 {
		// detach should go after the first attach, but before the second
		t.Fatalf("Expected seq to be 1, got %d", seq)
	}
	s.Release()
	wg.Wait()
}

func BenchmarkGetSlot(b *testing.B) {
	slots := NewSerialAttachDetachSlots()
	nodeName := "node1"
	for i := 0; i < b.N; i++ {
		slots.GetSlotFor(nodeName)
	}
}
