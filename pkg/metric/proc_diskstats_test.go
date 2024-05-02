package metric

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	testingclock "k8s.io/utils/clock/testing"
)

func newTestingProcDiskStats(t *testing.T) (*ProcDiskStats, *testingclock.FakeClock, string) {
	clk := testingclock.NewFakeClock(time.Now())
	proc := filepath.Join(t.TempDir(), "proc")
	os.Mkdir(proc, 0755)
	s, err := NewProcDiskStats(proc, clk)
	assert.NoError(t, err)
	return s, clk, proc
}

const (
	statLineA          = " 253       0 vda 18909 5492 4137980 55982 1262487 495033 17866563 3957744 0 207976 4016266 0 0 0 0 42182 2539 52991 2932647 0\n"
	statLineA_NormalIO = " 253       0 vda 18909 5492 4137980 55982 1262487 495033 17866563 3957744 1 207977 4016266 0 0 0 0 42182 2539 52991 2932647 0\n"
	statLineA_Complete = " 253       0 vda 18910 5492 4137980 55982 1262487 495033 17866563 3957744 1 207979 4016266 0 0 0 0 42182 2539 52991 2932647 0\n"
	statLineB          = " 253      16 vdb 433 0 101464 1471 0 0 0 0 0 444 1471 0 0 0 0 0 0 1345 0 0\n"
	statLineB_hung     = " 253      16 vdb 433 0 101464 1471 0 0 0 0 1 16444 1471 0 0 0 0 0 0 1345 0 0\n"
)

func writeStatLines(t *testing.T, proc string, lines ...string) {
	content := strings.Join(lines, "")
	assert.NoError(t, os.WriteFile(filepath.Join(proc, "diskstats"), []byte(content), 0644))
}

// no history saved, just pass through
func TestProcDiskStats_Disabled(t *testing.T) {
	s, _, proc := newTestingProcDiskStats(t)
	writeStatLines(t, proc, statLineB)

	_, err := s.GetStats()
	assert.NoError(t, err)
	assert.Empty(t, s.history)
	assert.Empty(t, s.GetHungDisks())
}

func TestProcDiskStats_GetStats(t *testing.T) {
	s, clk, proc := newTestingProcDiskStats(t)
	s.HungDuration = 15 * time.Second
	writeStatLines(t, proc, statLineB)

	_, err := s.GetStats()
	assert.NoError(t, err)

	// DropTooClose
	clk.Step(1 * time.Millisecond)
	_, err = s.GetStats()
	assert.NoError(t, err)
	assert.Len(t, s.history, 1)

	// Save stats
	// index: 0 1 2 3 4
	//  time: 0 2 4 6 8
	for i := 2; i < 5; i++ {
		clk.Step(2 * time.Second)
		_, err = s.GetStats()
		assert.NoError(t, err)
		assert.Len(t, s.history, i)
		assert.Empty(t, s.GetHungDisks()) // not enough history
	}

	// DropTooOld
	clk.Step(14 * time.Second)
	_, err = s.GetStats()
	assert.NoError(t, err)
	assert.Len(t, s.history, 3)       // save stat at 6, 8, 22
	assert.Empty(t, s.GetHungDisks()) // no IO, not hung

	clk.Step(20 * time.Second)
	_, err = s.GetStats()
	assert.NoError(t, err)
	assert.Len(t, s.history, 2)       // save stat at 22, 42
	assert.Empty(t, s.GetHungDisks()) // no IO, not hung
}

func TestProcDiskStats_GetHungDisks(t *testing.T) {
	cases := []struct {
		name   string
		before []string
		after  []string
		hung   []string
	}{
		{
			name:   "normal IO",
			before: []string{statLineA},
			after:  []string{statLineA_NormalIO},
		}, {
			name:   "complete",
			before: []string{statLineA_NormalIO},
			after:  []string{statLineA_Complete},
		}, {
			name:   "new device",
			before: []string{},
			after:  []string{statLineA_NormalIO},
		}, {
			name:   "disappear device",
			before: []string{statLineA_NormalIO},
			after:  []string{},
		}, {
			name:   "hung",
			before: []string{statLineB},
			after:  []string{statLineB_hung},
			hung:   []string{"vdb"},
		}, {
			name:   "hung old kernel",
			before: []string{statLineB_hung},
			after:  []string{statLineB_hung},
			hung:   []string{"vdb"},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s, clk, proc := newTestingProcDiskStats(t)
			s.HungDuration = 15 * time.Second

			writeStatLines(t, proc, c.before...)
			_, err := s.GetStats()
			assert.NoError(t, err)

			clk.Step(20 * time.Second)
			writeStatLines(t, proc, c.after...)
			_, err = s.GetStats()
			assert.NoError(t, err)

			// Check hung disks
			hung := s.GetHungDisks()
			assert.Len(t, hung, len(c.hung))
			for _, d := range hung {
				assert.Contains(t, c.hung, d.DeviceName)
			}
		})
	}
}
