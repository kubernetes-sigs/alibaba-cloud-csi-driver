package customfuse

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// tailBody returns what was retained, without the line saying something was
// dropped. That line is a property of the message the tail rides in rather than of
// the bytes themselves, so the tests comparing content read it apart from the ones
// asserting it is there.
func tailBody(t *testing.T, s string) string {
	t.Helper()
	if !strings.HasPrefix(s, "...[") {
		return s
	}
	_, body, found := strings.Cut(s, "\n")
	require.True(t, found, "the marker line must be followed by the retained bytes")
	return body
}

func TestStderrTailKeepsEverythingThatFits(t *testing.T) {
	for _, limit := range []int{1, 7, 64} {
		t.Run(fmt.Sprintf("limit=%d", limit), func(t *testing.T) {
			in := strings.Repeat("x", limit-1)
			tail := newStderrTail(limit)
			n, err := tail.Write([]byte(in))
			require.NoError(t, err)
			assert.Equal(t, len(in), n)
			assert.Equal(t, in, tail.String())
			assert.NotContains(t, tail.String(), "dropped", "nothing was dropped, so nothing may claim it was")
		})
	}
}

// A single write that exactly fills the ring leaves pos back at zero, which is also
// where a write that has not started wrapping yet would leave it. Reading up to pos
// would report an empty stream for a completely full one.
func TestStderrTailKeepsAWriteThatExactlyFillsIt(t *testing.T) {
	in := "0123456789"
	tail := newStderrTail(len(in))
	n, err := tail.Write([]byte(in))
	require.NoError(t, err)
	assert.Equal(t, len(in), n)
	assert.Equal(t, in, tail.String())
	assert.NotContains(t, tail.String(), "dropped")
}

func TestStderrTailKeepsTheEndOfOneOversizedWrite(t *testing.T) {
	in := "0123456789"
	tail := newStderrTail(4)
	n, err := tail.Write([]byte(in))
	require.NoError(t, err)
	assert.Equal(t, len(in), n, "bytes thrown away are still reported as accepted")
	assert.Equal(t, "6789", tailBody(t, tail.String()))
	assert.Contains(t, tail.String(), "6 earlier byte(s) dropped")
}

// Writes land on both sides of the ring's wrap: some entirely inside the free
// space, one split across the end. What survives has to be the end of the stream,
// not the end of whichever write happened to be last.
func TestStderrTailRetainsTheMostRecentBytes(t *testing.T) {
	const limit = 10
	in := "abcdefghijkl"
	tail := newStderrTail(limit)
	for i := 0; i < len(in); i += 5 {
		_, err := tail.Write([]byte(in[i:min(i+5, len(in))]))
		require.NoError(t, err)
	}
	assert.Equal(t, in[len(in)-limit:], tailBody(t, tail.String()))
	assert.Contains(t, tail.String(), fmt.Sprintf("%d earlier byte(s) dropped", len(in)-limit))
}

// The stream is chopped up by pipe buffers and by however the client flushes, and
// neither is knowable from here, so the retained bytes must not depend on it.
func TestStderrTailDoesNotDependOnHowTheStreamIsChunked(t *testing.T) {
	const limit = 17
	var lines []string
	for i := range 200 {
		lines = append(lines, fmt.Sprintf("line %03d\n", i))
	}
	in := strings.Join(lines, "")

	oneShot := newStderrTail(limit)
	_, err := oneShot.Write([]byte(in))
	require.NoError(t, err)

	chunked := newStderrTail(limit)
	for i := 0; i < len(in); i += 7 {
		_, err := chunked.Write([]byte(in[i:min(i+7, len(in))]))
		require.NoError(t, err)
	}

	assert.Equal(t, oneShot.String(), chunked.String())
	assert.Equal(t, in[len(in)-limit:], tailBody(t, chunked.String()))
	assert.False(t, strings.HasPrefix(tailBody(t, chunked.String()), "line "),
		"the cut is byte-exact, not snapped to a line")
}

func TestStderrTailReadsEmptyWhenNothingWasWritten(t *testing.T) {
	tail := newStderrTail(stderrTailLimit)
	assert.Equal(t, "", tail.String())

	n, err := tail.Write(nil)
	require.NoError(t, err)
	assert.Equal(t, 0, n)
	assert.Equal(t, "", tail.String())
}

// Holding the retained bytes to the limit is the whole reason this type exists
// instead of a buffer that grows. It shares its writer with the container's own
// log, and io.MultiWriter abandons the rest of its targets at the first short write
// or error, so every byte has to be reported as accepted however much is then
// thrown away.
func TestStderrTailNeverGrowsPastItsLimit(t *testing.T) {
	tail := newStderrTail(stderrTailLimit)
	chunk := make([]byte, 4096)
	for range chunk {
		n, err := tail.Write(chunk)
		require.NoError(t, err)
		assert.Equal(t, len(chunk), n)
	}
	assert.Contains(t, tail.String(), "dropped")
	assert.LessOrEqual(t, len(tail.String()), stderrTailLimit+128,
		"only the marker sits on top of the limit")
}
