package metric

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-logr/logr"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	versioncollector "github.com/prometheus/client_golang/prometheus/collectors/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/apimachinery/pkg/util/sets"
	"k8s.io/component-base/metrics/legacyregistry"
	"k8s.io/klog/v2"
)

const maxRequestsInFlight = 8

// Handler is a package of promHttp,metric entry
type Handler struct {
	collectors  map[string]Collector
	inFlightSem chan struct{}
}

// ServeHTTP implements http.Handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Limit concurrent requests to prevent OOM
	select {
	case h.inFlightSem <- struct{}{}:
		defer func() { <-h.inFlightSem }()
	default:
		http.Error(w, fmt.Sprintf(
			"Limit of concurrent requests reached (%d), try again later.", maxRequestsInFlight,
		), http.StatusServiceUnavailable)
		return
	}

	var pvcs sets.Set[string]
	if r.URL != nil && len(r.URL.RawQuery) != 0 {
		values, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			http.Error(w, "Failed to parse query", http.StatusBadRequest)
			return
		}

		multipvc := values["multipvc"]
		pvcs = make(sets.Set[string], len(multipvc))
		for _, v := range multipvc {
			for v := range strings.SplitSeq(v, ",") {
				pvcs.Insert(v)
			}
		}
	}

	csiCollector := CSICollector{
		Collectors: h.collectors,
		ctx:        r.Context(),
		pvcs:       pvcs,
	}

	reg := prometheus.NewRegistry()
	reg.MustRegister(
		versioncollector.NewCollector("alibaba_cloud_csi_driver"),
		&csiCollector, &CsiGrpcExecTimeCollector)
	handler := promhttp.InstrumentMetricHandler(reg, promhttp.HandlerFor(
		prometheus.Gatherers{reg, legacyregistry.DefaultGatherer},
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
			ErrorLog:      logger{klog.Background()},
			Registry:      reg,
		},
	))
	handler.ServeHTTP(w, r)
}

// NewMetricHandler method returns a promHttp object
func NewMetricHandler(driverNames []string, serviceType utils.ServiceType) *Handler {
	return &Handler{
		collectors:  registeredCollectors(driverNames, serviceType),
		inFlightSem: make(chan struct{}, maxRequestsInFlight),
	}
}

type logger struct {
	logr.Logger
}

func (l logger) Println(v ...any) {
	if len(v) == 2 {
		msg, ok1 := v[0].(string)
		err, ok2 := v[1].(error)
		if ok1 && ok2 {
			l.Logger.Error(err, msg)
			return
		}
	}
	l.Logger.Error(nil, "failed to gather metrics", "error", v)
}
