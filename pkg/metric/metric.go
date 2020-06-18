package metric

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/version"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

type handler struct {
	handler http.Handler
}

// ServeHTTP implements http.Handler.
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// To serve filtered metrics, we create a filtering handler on the fly.
	filteredHandler, err := h.innerHandler()
	if err != nil {
		logrus.Errorf("Couldn't create filtered metrics handler, err:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("Couldn't create filtered metrics handler: %s", err)))
		return
	}
	filteredHandler.ServeHTTP(w, r)
}

func Run(metricsPort string, metricsPath string, serverType string) {
	setServerType(serverType)
	//csi collector singleton
	_, err := NewCSICollector()
	if err != nil {
		logrus.Errorf("Couldn't create collector: %s", err)
		os.Exit(1)
	}
	http.Handle(metricsPath, newHandler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>CSI-Metric</title></head>
			<body>
			<h1>CSI Metric</h1>
			<p><a href="` + metricsPath + `">Metrics</a></p>
			</body>
			</html>`))
	})
	logrus.Infof("Listening on address:%s", metricsPort)

	server := &http.Server{Addr: metricsPort}
	if err := server.ListenAndServe(); err != nil {
		logrus.Infof("Server listen and serve err:%s", err.Error())
		os.Exit(1)
	}
}

func (h *handler) innerHandler() (http.Handler, error) {
	r := prometheus.NewRegistry()
	r.MustRegister(version.NewCollector("alibaba_cloud_csi_driver"))
	if err := r.Register(CSICollectorInstance); err != nil {
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

func newHandler() *handler {
	h := &handler{}
	if _, err := h.innerHandler(); err != nil {
		panic(fmt.Sprintf("Couldn't create metrics handler: %s", err))
	}
	return h
}

func setServerType(serverType string) {
	metricType = serverType
}
