package metric

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/stretchr/testify/assert"
)

func TestHandler_ServeHTTP(t *testing.T) {
	// Use empty temp directory to avoid reading real kubelet data
	podsRootPath = t.TempDir()

	handler := NewMetricHandler([]string{diskDriverName}, utils.Node)

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	body := rec.Body.String()
	// Should contain version collector metrics
	assert.Contains(t, body, "alibaba_cloud_csi_driver_build_info")
	// Should contain scrape duration metrics (from CSICollector)
	assert.Contains(t, body, "cluster_scrape_collector_duration_seconds")
}

func TestHandler_ServeHTTP_InvalidQuery(t *testing.T) {
	handler := NewMetricHandler(nil, utils.Node)

	// %ZZ is invalid percent-encoding, will cause url.ParseQuery to fail
	req := httptest.NewRequest(http.MethodGet, "/metrics?multipvc=%ZZ", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
	assert.Contains(t, rec.Body.String(), "Failed to parse query")
}

func TestHandler_ServeHTTP_ConcurrencyLimit(t *testing.T) {
	handler := NewMetricHandler(nil, utils.Node)

	// Fill the semaphore to simulate max concurrent requests
	for range maxRequestsInFlight {
		handler.inFlightSem <- struct{}{}
	}

	req := httptest.NewRequest(http.MethodGet, "/metrics", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusServiceUnavailable, rec.Code)
	assert.Contains(t, rec.Body.String(), "Limit of concurrent requests reached")
}
