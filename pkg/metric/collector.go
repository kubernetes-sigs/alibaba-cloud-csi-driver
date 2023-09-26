package metric

import (
	"errors"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

var (
	scrapeDurationDesc = prometheus.NewDesc(
		prometheus.BuildFQName(clusterNamespace, scrapeSubSystem, "collector_duration_seconds"),
		"csi_metric: Duration of a collector scrape.",
		[]string{"collector"},
		nil,
	)
	scrapeSuccessDesc = prometheus.NewDesc(
		prometheus.BuildFQName(clusterNamespace, scrapeSubSystem, "collector_success"),
		"csi_metric: Whether a collector succeeded.",
		[]string{"collector"},
		nil,
	)
)

func registerCollector(collector string, factory collectorFactoryFunc, relatedDrivers ...string) {
	registry = append(registry, collectorRegistryItem{
		Name:           collector,
		Factory:        factory,
		RelatedDrivers: relatedDrivers,
	})
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

// newCSICollector method returns the CSICollector object
func newCSICollector(metricType string, driverNames []string) error {
	if csiCollectorInstance != nil {
		return nil
	}
	collectors := make(map[string]Collector)
	if metricType == pluginService {
		enabledDrivers := map[string]struct{}{}
		for _, d := range driverNames {
			enabledDrivers[d] = struct{}{}
		}
		for _, reg := range registry {
			for _, d := range reg.RelatedDrivers {
				if _, ok := enabledDrivers[d]; ok {
					collector, err := reg.Factory()
					if err != nil {
						return err
					}
					collectors[reg.Name] = collector
				}
			}
		}

	}
	csiCollectorInstance = &CSICollector{Collectors: collectors}

	return nil
}

// Describe implements the prometheus.Collector interface.
func (csi CSICollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- scrapeDurationDesc
	ch <- scrapeSuccessDesc
}

// Collect implements the prometheus.Collector interface.
func (csi CSICollector) Collect(ch chan<- prometheus.Metric) {
	defer func() { scalerPvcMap = nil }()
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
			logrus.Infof("Collector returned no data,name: %s, duration_seconds: %f, err: %s", name, duration.Seconds(), err.Error())
		} else {
			logrus.Errorf("Collector failed, name: %s, duration_seconds: %f, err: %s", name, duration.Seconds(), err.Error())
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

// IsNoDataError method is to determine whether the Collector has no data to return
func IsNoDataError(err error) bool {
	return err == ErrNoData
}
