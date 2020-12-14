package metric

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/sirupsen/logrus"
	"net/http"
)

//Handler is a package of promHttp,metric entry
type Handler struct {
}

// ServeHTTP implements http.Handler.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handler, err := h.innerHandler()
	if err != nil {
		logrus.Errorf("Couldn't create filtered metrics handler, err:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Couldn't create filtered metrics handler: %s", err)))
		return
	}
	handler.ServeHTTP(w, r)
}

//NewMetricHandler method returns a promHttp object
func NewMetricHandler(serviceType string) *Handler {
	setServiceType(serviceType)
	//csi collector singleton
	err := newCSICollector()
	if err != nil {
		logrus.Errorf("Couldn't create collector: %s", err)
	}
	return newHandler()
}

func (h *Handler) innerHandler() (http.Handler, error) {
	r := prometheus.NewRegistry()
	r.MustRegister(version.NewCollector("alibaba_cloud_csi_driver"))
	if err := r.Register(csiCollectorInstance); err != nil {
		return nil, fmt.Errorf("Couldn't register node collector: %s", err)
	}
	handler := promhttp.HandlerFor(
		prometheus.Gatherers{r},
		promhttp.HandlerOpts{
			ErrorHandling: promhttp.ContinueOnError,
		},
	)
	return handler, nil
}

func newHandler() *Handler {
	h := &Handler{}
	if _, err := h.innerHandler(); err != nil {
		logrus.Errorf("Couldn't create metrics handler: %s", err)
	}
	return h
}

func setServiceType(serviceType string) {
	metricType = serviceType
}
