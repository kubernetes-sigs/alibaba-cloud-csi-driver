package metric

import (
	"context"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/apimachinery/pkg/util/sets"
)

type pfsRawBlockStatCollector struct {
}

func init() {
	registerCollector("pfs_block_stat", NewPfsRawBlockStatCollector, diskDriverName)
}

// NewPfsRawBlockStatCollector returns a new Collector exposing stats.
func NewPfsRawBlockStatCollector() (Collector, error) {
	if os.Getenv("ENABLE_PFS_BLOCK_MEREIC") == "true" {
		panic("ENABLE_PFS_BLOCK_MEREIC is true, but no longer supported in this version")
	}
	return &pfsRawBlockStatCollector{}, nil
}

func (p *pfsRawBlockStatCollector) Update(ctx context.Context, pvcs sets.Set[string], ch chan<- prometheus.Metric) error {
	return nil
}
