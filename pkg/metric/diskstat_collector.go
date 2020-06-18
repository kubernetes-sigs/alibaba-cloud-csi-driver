package metric

import (
	"bufio"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"io"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"strconv"
	"strings"
	"sync"
)

type diskStatCollector struct {
	descs             []typedFactorDesc
	lastPvPathMapping map[string]string   //key:pvName,value:mountPath
	lastPvPvcMapping  map[string][]string //key:pvName, value:pvcNamespace,pvcName
	clientSet         *kubernetes.Clientset
}

func init() {
	registerCollector("diskstat", NewDiskStatCollector)
}

// NewPVUsageCollector returns a new Collector exposing pv stats.
func NewDiskStatCollector() (Collector, error) {
	config, err := rest.InClusterConfig()
	if err != nil {
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}

	return &diskStatCollector{
		descs: []typedFactorDesc{
			//4 - reads completed successfully
			{desc: diskReadsCompletedDesc, valueType: prometheus.CounterValue,},
			//5 - reads merged
			{desc: diskReadsMergeDesc, valueType: prometheus.CounterValue,},
			//6 - sectors read
			{desc: diskReadBytesDesc, valueType: prometheus.CounterValue, factor: diskSectorSize,},
			//7 - time spent reading (ms)
			{desc: diskReadTimeSecondsDesc, valueType: prometheus.CounterValue, factor: .001,},
			//8 - writes completed
			{desc: diskWritesCompletedDesc, valueType: prometheus.CounterValue,},
			//9 - writes merged
			{desc: diskWriteMergeDesc, valueType: prometheus.CounterValue,},
			//10 - sectors written
			{desc: diskWrittenBytesDesc, valueType: prometheus.CounterValue, factor: diskSectorSize,},
			//11 - time spent writing (ms)
			{desc: diskWriteTimeSecondsDesc, valueType: prometheus.CounterValue, factor: .001,},
			//12 - I/Os currently in progress
			{desc: diskIONowDesc, valueType: prometheus.GaugeValue,},
		},
		lastPvPathMapping: make(map[string]string, 0),
		lastPvPvcMapping:  make(map[string][]string, 0),
		clientSet:         clientset,
	}, nil
}

func (p *diskStatCollector) Update(ch chan<- prometheus.Metric) error {
	//startTime := time.Now()
	diskStats, err := getDiskStats()
	if err != nil {
		return fmt.Errorf("couldn't get diskstats: %s", err)
	}
	volDataJsonPath, err := findVolDataJsonFileByPattern(podsRootPath)
	if err != nil {
		return err
	}

	pvDeviceNameMapping := make(map[string]string, 0)
	thisPvPathMapping := make(map[string]string, 0)
	for _, path := range volDataJsonPath {
		//Get disk pvName
		pvName, err := getVolumeIDByJson(path, diskDriverName)
		if err != nil {
			continue
		}
		pvDevice, err := GetDeviceByVolumeID(pvName)
		if err != nil {
			continue
		}
		thisPvPathMapping[pvName] = path
		pvDeviceNameMapping[pvDevice] = pvName
	}

	//If there is a change:add, modify, delete
	updateLastPvcMapping(thisPvPathMapping, &p.lastPvPathMapping, p.clientSet, &p.lastPvPvcMapping)
	wg := sync.WaitGroup{}
	for dev, stats := range diskStats {
		pvName, ok := pvDeviceNameMapping[dev]
		pvcArr, ok := p.lastPvPvcMapping[pvName]
		if ok {
			pvcNamespace := pvcArr[0]
			pvcName := pvcArr[1]
			wg.Add(1)
			go func(devArgs string, pvcNamespaceArgs string, pvcNameArgs string, statsArgs []string) {
				defer wg.Done()
				p.setDiskMetric(devArgs, pvcNamespaceArgs, pvcNameArgs, statsArgs, ch)
			}(dev, pvcNamespace, pvcName, stats)
		}
	}
	wg.Wait()
	//elapsedTime := time.Since(startTime)
	//logrus.Info("DiskStat spent time:", elapsedTime)
	return nil
}

func (p *diskStatCollector) setDiskMetric(dev string, pvcNamespace string, pvcName string, stats []string, ch chan<- prometheus.Metric) {
	for i, value := range stats {
		if i >= len(p.descs) {
			return
		}
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		ch <- p.descs[i].mustNewConstMetric(v, pvcNamespace, pvcName, dev, DiskStorageName)
	}
}

func getDiskStats() (map[string][]string, error) {
	diskStatsPath := procFilePath(diskStatsFileName)

	file, err := os.Open(diskStatsPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return parseDiskStats(file)
}

func parseDiskStats(r io.Reader) (map[string][]string, error) {
	var (
		diskStats = map[string][]string{}
		scanner   = bufio.NewScanner(r)
	)

	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		if len(parts) < 4 { // we strip major, minor and dev
			return nil, fmt.Errorf("Invalid line in %s: %s", procFilePath(diskStatsFileName), scanner.Text())
		}
		dev := "/dev/" + parts[2]
		diskStats[dev] = parts[3:12]
	}

	return diskStats, scanner.Err()
}
