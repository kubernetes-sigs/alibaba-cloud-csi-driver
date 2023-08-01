package metric

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	vgFreeBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "vg_free_bytes"),
		"Total amount of free bytes in LVM volume group",
		[]string{"vg_name"}, nil,
	)
	vgSizeBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "vg_size_bytes"),
		"Total size of LVM volume group in bytes",
		[]string{"vg_name"}, nil,
	)
)

func init() {
	registerCollector("local_stat", NewLocalStatCollector)
}

type localStatCollector struct {
}

func (*localStatCollector) Update(ch chan<- prometheus.Metric) error {
	vgs, err := server.ListVG()
	if err != nil {
		return err
	}

	for _, vg := range vgs {
		name := vg.Name
		ch <- prometheus.MustNewConstMetric(vgSizeBytesDesc, prometheus.GaugeValue, float64(vg.Size), name)
		ch <- prometheus.MustNewConstMetric(vgFreeBytesDesc, prometheus.GaugeValue, float64(vg.FreeSize), name)
	}

	return nil
}

func NewLocalStatCollector() (Collector, error) {
	return &localStatCollector{}, nil
}
