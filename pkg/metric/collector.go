package metric

import (
	"errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

func registerCollector(collector string, factory func() (Collector, error)) {
	factories[collector] = factory
}

// Collector is the interface a collector has to implement.
type Collector interface {
	// Get new metrics and expose them via prometheus registry.
	Update(ch chan<- prometheus.Metric) error
}

// CSICollector implements the prometheus.Collector interface.
type CSICollector struct {
	Collectors map[string]Collector
}

func NewCSICollector() (*CSICollector, error) {
	if CSICollectorInstance == nil {
		collectors := make(map[string]Collector)
		for key := range factories {
			var collector Collector
			var err error
			switch metricType {
			case nodeServer:
				if !nodeMetricSet.Contains(key) {
					continue
				}
				collector, err = factories[key]()
			case clusterServer:
				if !clusterMetricSet.Contains(key) {
					continue
				}
				collector, err = factories[key]()
			default:
				return nil, errors.New("Unknown metricType:" + metricType)
			}
			if err != nil {
				return nil, err
			}
			collectors[key] = collector
		}
		CSICollectorInstance = &CSICollector{Collectors: collectors}
	}

	return CSICollectorInstance, nil
}

// Describe implements the prometheus.Collector interface.
func (csi CSICollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- scrapeDurationDesc
	ch <- scrapeSuccessDesc
}

// Collect implements the prometheus.Collector interface.
func (csi CSICollector) Collect(ch chan<- prometheus.Metric) {
	wg := sync.WaitGroup{}
	wg.Add(len(csi.Collectors))
	for name, c := range csi.Collectors {
		go func(name string, c Collector) {
			execute(name, c, ch)
			wg.Done()
		}(name, c)
	}
	wg.Wait()
}

func execute(name string, c Collector, ch chan<- prometheus.Metric) {
	begin := time.Now()
	err := c.Update(ch)
	duration := time.Since(begin)
	var success float64

	if err != nil {
		if IsNoDataError(err) {
			logrus.Infof("Collector returned no data,name:%, duration_seconds:%v, err:%s", name, duration.Seconds(), err.Error())
		} else {
			logrus.Errorf("Collector failed, name%s, duration_seconds:%v, err:%s", name, duration.Seconds(), err.Error())
		}
		success = 0
	} else {
		success = 1
	}
	ch <- prometheus.MustNewConstMetric(scrapeDurationDesc, prometheus.GaugeValue, duration.Seconds(), name)
	ch <- prometheus.MustNewConstMetric(scrapeSuccessDesc, prometheus.GaugeValue, success, name)
}

// ErrNoData indicates the collector found no data to collect, but had no other error.
var ErrNoData = errors.New("Collector returned no data")

func IsNoDataError(err error) bool {
	return err == ErrNoData
}
