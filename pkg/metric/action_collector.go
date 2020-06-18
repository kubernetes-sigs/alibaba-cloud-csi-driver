package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"sync"
	"time"
)

type actionDesc struct {
	desc      *prometheus.Desc
	volumeID  string
	pvType    string
	action    string
	spendTime time.Duration
}

type actionCollector struct {
	descs  []actionDesc
	descsSyncLock sync.RWMutex
}

func init() {
	registerCollector("action", NewActionCollector)
}

func CollectDesc(volumeID string, action string, pvType string, timer *prometheus.Timer, collector *actionCollector) {
	if !nodeMetricSet.Contains("action") {
		return
	}
	collector.descsSyncLock.Lock()
	defer  collector.descsSyncLock.Unlock()
	duration := timer.ObserveDuration()
	desc := actionDesc{desc: actionTimeSecondsDesc, volumeID: volumeID, pvType: pvType, action: action, spendTime: duration}
	collector.descs = append(collector.descs, desc)
}

// NewPVUsageCollector returns a new Collector exposing pv stats.
func NewActionCollector() (Collector, error) {
	if ActionCollectorInstance == nil {
		collector := &actionCollector{
			descs:  make([]actionDesc, 0),
		}
		ActionCollectorInstance = collector
	}
	return ActionCollectorInstance, nil
}

func (a *actionCollector) Update(ch chan<- prometheus.Metric) error {
	a.descsSyncLock.Lock()
	defer  a.descsSyncLock.Unlock()
	for _, desc := range a.descs {
		ch <- prometheus.MustNewConstMetric(
			desc.desc, prometheus.CounterValue, float64(desc.spendTime), desc.volumeID, desc.pvType, desc.action)

	}
	a.descs = make([]actionDesc, 0)
	return nil
}
