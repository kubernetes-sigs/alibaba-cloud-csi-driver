package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"strings"
)

type persistentVolumeClaimCollector struct {
	counter                        *prometheus.Desc
	volumeTypeStatusCounterMapping map[string]map[string]int
	clientSet                      *kubernetes.Clientset
}

func init() {
	registerCollector("pvc", NewPersistentVolumeClaimCollector)
}

// NewPVUsageCollector returns a new Collector exposing pv stats.
func NewPersistentVolumeClaimCollector() (Collector, error) {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	return &persistentVolumeClaimCollector{
		counter:                        persistentVolumeClaimCounterDesc,
		volumeTypeStatusCounterMapping: map[string]map[string]int{},
		clientSet:                      clientset,
	}, nil
}

func (p *persistentVolumeClaimCollector) Update(ch chan<- prometheus.Metric) error {
	persistentVolumeClaimList, err := getPersistentVolumeClaimList(p.clientSet, defaultNamespace)
	if err != nil {
		return err
	}
	p.volumeTypeStatusCounterMapping = map[string]map[string]int{}
	for _, persistentVolumeClaim := range persistentVolumeClaimList.Items {
		driverName, err := getPersistentVolumeDriver(p.clientSet, persistentVolumeClaim.Spec.VolumeName)
		if driverName == nil || err != nil {
			continue
		}

		currentStatus := strings.ToLower(string(persistentVolumeClaim.Status.Phase))
		switch *driverName {
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

func (p *persistentVolumeClaimCollector) setMetricVolumeTypeStatusMapping(volumeType string, volumeStatus string) {
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
