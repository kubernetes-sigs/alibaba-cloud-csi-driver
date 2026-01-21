package metric

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/go-logr/logr"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	versioncollector "github.com/prometheus/client_golang/prometheus/collectors/version"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"k8s.io/component-base/metrics/legacyregistry"
	"k8s.io/klog/v2"
)

// Handler is a package of promHttp,metric entry
type Handler struct {
}

// ServeHTTP implements http.Handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL != nil && len(r.URL.RawQuery) != 0 {
		queryField := strings.Split(r.URL.RawQuery, "=")
		if len(queryField) >= 2 && queryField[0] == "multipvc" {
			pvcNameArray := strings.Split(queryField[1], ",")
			scalerPvcMap = new(sync.Map)
			for _, pvcName := range pvcNameArray {
				scalerPvcMap.Store(pvcName, true)
			}
		}
	}
	handler, err := h.innerHandler()
	if err != nil {
		klog.Errorf("Couldn't create filtered metrics handler, err:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(fmt.Sprintf("Couldn't create filtered metrics handler: %s", err)))
		return
	}
	handler.ServeHTTP(w, r)
}

// NewMetricHandler method returns a promHttp object
func NewMetricHandler(driverNames []string, serviceType utils.ServiceType) *Handler {
	//csi collector singleton
	initCSICollector(driverNames, serviceType)
	return newHandler()
}

type logger struct {
	logr.Logger
}

func (l logger) Println(v ...interface{}) {
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

func (h *Handler) innerHandler() (http.Handler, error) {
	r := prometheus.NewRegistry()
	r.MustRegister(
		versioncollector.NewCollector("alibaba_cloud_csi_driver"),
		csiCollectorInstance)
	handler := promhttp.InstrumentMetricHandler(r, promhttp.HandlerFor(
		prometheus.Gatherers{r, legacyregistry.DefaultGatherer},
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
			ErrorLog:      logger{klog.Background()},
			Registry:      r,
		},
	))
	return handler, nil
}

func newHandler() *Handler {
	h := &Handler{}
	if _, err := h.innerHandler(); err != nil {
		klog.Errorf("Couldn't create metrics handler: %s", err)
	}
	return h
}
