package metric

import (
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	httpProtocol = "http"
	dnsName      = "storage-monitor-service.kube-system.svc.cluster.local"
	port         = "11280"
	metricPath   = "metrics"
)

// HTTPClient is http client with server client
type HTTPClient struct {
	url    string
	client *http.Client
}

// getURL to get the corresponding url
func getURL(httpProtocol string, dnsName string, port string, metricPath string) string {
	url := httpProtocol + "://" + dnsName + ":" + port + "/" + metricPath
	return url
}

// NewHTTPClient to create http client
func NewHTTPClient() *HTTPClient {
	url := getURL(httpProtocol, dnsName, port, metricPath)
	client := &http.Client{}
	client.Timeout = time.Millisecond * 500
	return &HTTPClient{url, client}
}

// Get to get response from server with http client
func (h *HTTPClient) Get(key string, val string) (*http.Response, error) {
	url := h.url + "?" + key + "=" + val
	resp, err := h.client.Get(url)
	if err != nil {
		//log.Errorf("Get url %s is failed, err: %s. ", url, err)
		return nil, err
	}
	return resp, nil
}

// ReadBody to read body from response
func (h *HTTPClient) ReadBody(response *http.Response) string {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Get response is failed, response:%+v, err:%s", response, err.Error())
		return string("")
	}
	return string(body)
}

// Close to close response
func (h *HTTPClient) Close(response *http.Response) {
	if response != nil {
		_ = response.Body.Close()
	}
}
