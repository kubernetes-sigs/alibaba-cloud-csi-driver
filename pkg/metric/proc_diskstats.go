package metric

import (
	"sync"
	"time"

	"github.com/prometheus/procfs"
	"github.com/prometheus/procfs/blockdevice"
	"k8s.io/utils/clock"
)

type statHistoryEntry struct {
	// The time at which the stats were collected.
	time time.Time
	// The stats collected.
	stats []blockdevice.Diskstats
}

type ProcDiskStats struct {
	bdev blockdevice.FS

	history   []statHistoryEntry
	historyMu sync.RWMutex
	// a disk is considered to be hung if there is I/O in flight but no progress made in at least this duration.
	HungDuration time.Duration

	clk clock.PassiveClock
}

func NewDefaultProcDiskStats() (*ProcDiskStats, error) {
	return NewProcDiskStats(procfs.DefaultMountPoint, clock.RealClock{})
}

func NewProcDiskStats(procMountPoint string, clk clock.PassiveClock) (*ProcDiskStats, error) {
	fs, err := blockdevice.NewFS(procMountPoint, "/") // we do not use sys mount
	if err != nil {
		return nil, err
	}
	return &ProcDiskStats{
		bdev: fs,
		clk:  clk,
	}, nil
}

func (s *ProcDiskStats) GetStats() ([]blockdevice.Diskstats, error) {
	stat, err := s.bdev.ProcDiskstats()
	if err != nil {
		return nil, err
	}
	if s.HungDuration <= 0 {
		return stat, nil
	}

	now := s.clk.Now()
	nextEntry := statHistoryEntry{
		time:  now,
		stats: stat,
	}

	s.historyMu.Lock()
	defer s.historyMu.Unlock()
	his := s.history
	if len(his) == 0 {
		s.history = []statHistoryEntry{nextEntry}
		return stat, nil
	}
	if now.Sub(his[len(his)-1].time) < 1*time.Second {
		// avoid store too many history
		return stat, nil
	}

	for len(his) > 1 && now.Sub(his[1].time) > s.HungDuration {
		his = his[1:]
	}
	his = append(his, nextEntry)
	s.history = his
	return stat, nil
}

func (s *ProcDiskStats) GetHungDisks() []blockdevice.Info {
	s.historyMu.RLock()
	his := s.history
	s.historyMu.RUnlock()

	if len(his) < 2 {
		return nil
	}
	last := his[len(his)-1]
	first := his[0]
	if last.time.Sub(first.time) < s.HungDuration {
		return nil // not enough history yet
	}
	firstMap := make(map[blockdevice.Info]*blockdevice.IOStats, len(first.stats))
	for i, d := range first.stats {
		firstMap[d.Info] = &first.stats[i].IOStats
	}

	var hung []blockdevice.Info
	for _, d1 := range last.stats {
		if d1.IOsInProgress == 0 {
			continue // currently no IO request
		}
		d0, ok := firstMap[d1.Info]
		if !ok {
			continue
		}

		// bio based drivers are increasing ReadIOs at the beginning of IO in old kernel,
		// not at the completion. Fortunately, nvme or virtIO are blk_mq based.
		// see https://lore.kernel.org/all/20230223091226.1135678-1-yukuai1@huaweicloud.com/
		if d1.ReadIOs != d0.ReadIOs || d1.WriteIOs != d0.WriteIOs || d1.DiscardIOs != d0.DiscardIOs ||
			d1.IOsInProgress < d0.IOsInProgress {
			continue // IO made progress, not hung
		}

		// At least hungDuration has passed during IO or there are IOs across the two samples.
		// Note: IOsTotalTicks may only increase at IO completion in old kernel,
		// So it may not change during IO hang. Keep it here for faster hung detection in new kernel.
		// see https://lore.kernel.org/all/20220217064247.4041435-1-zhangwensheng5@huawei.com/
		if time.Duration(d1.IOsTotalTicks-d0.IOsTotalTicks)*time.Millisecond > s.HungDuration ||
			d0.IOsInProgress > 0 {
			hung = append(hung, d1.Info)
		}
	}
	return hung
}
