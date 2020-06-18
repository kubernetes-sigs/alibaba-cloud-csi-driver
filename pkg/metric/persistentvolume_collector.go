package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
)

type persistentVolumeCollector struct {
	counter                        *prometheus.Desc
	volumeTypeStatusCounterMapping map[string]map[string]int
	clientSet                      *kubernetes.Clientset
}

func init() {
	registerCollector("pv", NewPersistentVolumeCollector)
}

// NewPVUsageCollector returns a new Collector exposing pv stats.
func NewPersistentVolumeCollector() (Collector, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return &persistentVolumeCollector{
		counter:                        persistentVolumeCounterDesc,
		volumeTypeStatusCounterMapping: map[string]map[string]int{},
		clientSet:                      clientset,
	}, nil
}

func (p *persistentVolumeCollector) Update(ch chan<- prometheus.Metric) error {
	persistentVolumeList, err := getPersistentVolumeList(p.clientSet)
	if err != nil {
		return err
	}
	p.volumeTypeStatusCounterMapping = map[string]map[string]int{}
	for _, persistentVolume := range persistentVolumeList.Items {
		if persistentVolume.Spec.CSI == nil {
			continue
		}
		driverName := strings.ToLower(persistentVolume.Spec.CSI.Driver)
		currentStatus := strings.ToLower(string(persistentVolume.Status.Phase))
		switch driverName {
		case diskDriverName:
			p.setMetricVolumeTypeStatusMapping(DiskStorageName, currentStatus)
		case nasDriverName:
			p.setMetricVolumeTypeStatusMapping(NasStorageName, currentStatus)
		case ossDriverName:
			p.setMetricVolumeTypeStatusMapping(OssStorageName, currentStatus)
		default:
			p.setMetricVolumeTypeStatusMapping(unknownStorageName, currentStatus)
			logrus.Infof("Unknown persistent volume type:%+v", driverName)
		}
	}

	for volumeType, volumeStatusCounterMapping := range p.volumeTypeStatusCounterMapping {
		for volumeStatus, counter := range volumeStatusCounterMapping {
			ch <- prometheus.MustNewConstMetric(
				p.counter, prometheus.GaugeValue, float64(counter), volumeType, volumeStatus)
		}
	}
	return nil
}

func (p *persistentVolumeCollector) setMetricVolumeTypeStatusMapping(volumeType string, volumeStatus string) {
	if volumeStatusCounterMapping, ok := p.volumeTypeStatusCounterMapping[volumeType]; ok {
		if _, ok := volumeStatusCounterMapping[volumeStatus]; ok {
			p.volumeTypeStatusCounterMapping[volumeType][volumeStatus]++
		} else {
			p.volumeTypeStatusCounterMapping[volumeType][volumeStatus] = counterInitSize
		}
	} else {
		mapping := map[string]int{volumeStatus: counterInitSize}
		p.volumeTypeStatusCounterMapping[volumeType] = mapping
	}
}
