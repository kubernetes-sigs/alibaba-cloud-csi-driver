package customfuse

import "fmt"

// stderrTailLimit is how much of the entrypoint's startup stderr is kept to travel
// back inside the mount error.
//
// The whole stream already reaches this container's log unconditionally, so this is
// a second copy and its only reader is the error the caller gets back. The bound
// therefore comes from what that error can carry, not from any guess about how much
// a client might print — which is unknowable here, the entrypoint being the
// customer's own script. It does have to be a bound: the copy accumulates in the
// mount-proxy's heap, the fuse pod sets no memory limit by default, and the window
// it covers runs until the mount appears or the request deadline fires.
const stderrTailLimit = 64 << 10

// stderrTail retains the last limit bytes written to it.
//
// Only the os/exec goroutine copying the entrypoint's stderr writes here, and the
// one read happens after cmd.Wait, which does not return until that copy has
// finished. The two never overlap, so this needs no lock — but that ordering is a
// property of the caller, not of this type, and reading it from anywhere else loses
// it.
type stderrTail struct {
	buf []byte
	// pos is where the next byte lands, wrapping at len(buf). Once more than len(buf)
	// bytes have been written it is also where the oldest retained byte sits.
	pos     int
	written int64
}

func newStderrTail(limit int) *stderrTail {
	return &stderrTail{buf: make([]byte, limit)}
}

func (t *stderrTail) Write(p []byte) (int, error) {
	n := len(p)
	t.written += int64(n)
	limit := len(t.buf)
	switch {
	case n >= limit:
		// A single write longer than the whole ring: only its tail can survive.
		copy(t.buf, p[n-limit:])
		t.pos = 0
	case t.pos+n <= limit:
		copy(t.buf[t.pos:], p)
		t.pos += n
	default:
		head := limit - t.pos
		copy(t.buf[t.pos:], p[:head])
		copy(t.buf, p[head:])
		t.pos = n - head
	}
	return n, nil
}

// String returns what was retained, oldest byte first, and says so when anything was
// dropped. The note goes first because the message it rides in gets truncated again
// downstream, and a reader who never sees the note would take a partial stream for
// the whole of what the client said.
func (t *stderrTail) String() string {
	if t.written == 0 {
		return ""
	}
	limit := int64(len(t.buf))
	if t.written <= limit {
		// Nothing was dropped, so nothing wrapped: what was written sits contiguously
		// from the start. Reading to pos instead would come up empty for the single
		// write that exactly fills the ring, which leaves pos back at zero.
		return string(t.buf[:t.written])
	}
	return fmt.Sprintf("...[%d earlier byte(s) dropped; the entrypoint's full stderr is in this container's log]\n%s%s",
		t.written-limit, t.buf[t.pos:], t.buf[:t.pos])
}
