package metric

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"strings"
	"sync"
)

//Handler is a package of promHttp,metric entry
type Handler struct {
}

var kubeletRootDir = os.Getenv("KUBELET_ROOT_DIR")

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
		logrus.Errorf("Couldn't create filtered metrics handler, err:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Couldn't create filtered metrics handler: %s", err)))
		return
	}
	handler.ServeHTTP(w, r)
}

//NewMetricHandler method returns a promHttp object
func NewMetricHandler(serviceType string) *Handler {
	if kubeletRootDir == "" {
		kubeletRootDir = "/var/lib/kubelet"
	}

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
