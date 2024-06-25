package metric

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const (
	storageMonitorSvcName       = "storage-monitor-service"
	storageMonitorPort          = "11280"
	storageMonitorClientTimeout = time.Millisecond * 500
	monitorIPUpdateInterval     = time.Minute * 30
)

type StorageMonitorClient struct {
	sync.RWMutex
	clusterIP  string
	kubeClient kubernetes.Interface
}

func NewStorageMonitorClient(kubeClient kubernetes.Interface) *StorageMonitorClient {
	c := &StorageMonitorClient{
		kubeClient: kubeClient,
	}
	if err := c.updateClusterIP(); err != nil {
		log.Warnf("failed to get clusterIP of storage-monitor: %v", err)
	}

	var (
		lastError   error
		lastLogTime time.Time
	)

	go func() {
		ticker := time.NewTicker(monitorIPUpdateInterval)
		defer ticker.Stop()
		for range ticker.C {
			err := c.updateClusterIP()
			if err != nil {
				now := time.Now()
				// ignore recent and duplicated error
				if !(lastError != nil && lastError.Error() == err.Error() && now.Sub(lastLogTime) < monitorIPUpdateInterval*24) {
					log.Warnf("failed to update clusterIP of storage-monitor: %v", err)
					lastLogTime = now
				}
			}
			lastError = err
		}
	}()
	return c
}

func (c *StorageMonitorClient) GetNasCapacityInfo(pv string) (*nfsCapacityInfo, error) {
	c.RLock()
	clusterIP := c.clusterIP
	c.RUnlock()
	if clusterIP == "" {
		return nil, errors.New("clusterIP not initialized")
	}
	params := url.Values{"multi-cnfs-nas": []string{pv}}
	uri := url.URL{
		Scheme:   "http",
		Host:     net.JoinHostPort(clusterIP, storageMonitorPort),
		Path:     "/metrics",
		RawQuery: params.Encode(),
	}
	resp, err := (&http.Client{Timeout: storageMonitorClientTimeout}).Get(uri.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http response status: %s", resp.Status)
	}
	var capacityInfos map[string]nfsCapacityInfo
	err = json.NewDecoder(resp.Body).Decode(&capacityInfos)
	if err != nil {
		return nil, err
	}
	capacityInfo, ok := capacityInfos[pv]
	if ok {
		return &capacityInfo, nil
	}
	return nil, nil
}

func (c *StorageMonitorClient) updateClusterIP() error {
	svc, err := c.kubeClient.CoreV1().Services("kube-system").Get(context.Background(), storageMonitorSvcName, v1.GetOptions{})
	if err != nil {
		return err
	}
	currentClusterIP := svc.Spec.ClusterIP
	if currentClusterIP == "" {
		return errors.New("spec.ClusterIP of service is empty")
	}
	c.Lock()
	defer c.Unlock()
	if c.clusterIP != currentClusterIP {
		log.Infof("update storage-monitor IP: %s", currentClusterIP)
		c.clusterIP = currentClusterIP
	}
	return nil
}
