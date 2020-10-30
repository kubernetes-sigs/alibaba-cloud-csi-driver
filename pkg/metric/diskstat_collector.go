package metric

import (
	"bufio"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"io"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	diskStatLabelNames = []string{"namespace", "pvc", "device", "type"}
)

var (
	//4 - reads completed successfully
	diskReadsCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_read_completed_total"),
		"The total number of reads completed successfully.",
		diskStatLabelNames, nil,
	)
	//5 - reads merged
	diskReadsMergeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_read_merged_total"),
		"The total number of reads merged.",
		diskStatLabelNames,
		nil,
	)
	//6 - sectors read
	diskReadBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_read_bytes_total"),
		"The total number of bytes read successfully.",
		diskStatLabelNames, nil,
	)
	//7 - time spent reading (ms)
	diskReadTimeSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_read_time_seconds_total"),
		"The total number of seconds spent by all reads.",
		diskStatLabelNames,
		nil,
	)
	//8 - writes completed
	diskWritesCompletedDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_write_completed_total"),
		"The total number of writes completed successfully.",
		diskStatLabelNames, nil,
	)
	//9 - writes merged
	diskWriteMergeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_write_merged_total"),
		"The number of writes merged.",
		diskStatLabelNames,
		nil,
	)
	//10 - sectors written
	diskWrittenBytesDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_write_bytes_total"),
		"The total number of bytes written successfully.",
		diskStatLabelNames, nil,
	)
	//11 - time spent writing (ms)
	diskWriteTimeSecondsDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_write_time_seconds_total"),
		"This is the total number of seconds spent by all writes.",
		diskStatLabelNames,
		nil,
	)
	//12 - I/Os currently in progress
	diskIONowDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "disk_io_now"),
		"The number of I/Os currently in progress.",
		diskStatLabelNames,
		nil,
	)
)

type diskStatCollector struct {
	lantencyThreshold int //Unit: milliseconds
	descs             []typedFactorDesc
	lastPvPathMapping map[string]string   //key:pvName, value:mountPath
	lastPvPvcMapping  map[string][]string //key:pvName, value:pvcNamespace, pvcName
	clientSet         *kubernetes.Clientset
}

func init() {
	registerCollector("diskstat", NewDiskStatCollector)
}

// NewDiskStatCollector returns a new Collector exposing disk stats.
func NewDiskStatCollector() (Collector, error) {
	lantencyThreshold := 10
	lantencyThreshold = strings.Trim(os.Getenv("DISK_LATENCY_THRESHOLD"), " ")
	if lantencyThreshold == "true" {
	}
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
			{desc: diskReadsCompletedDesc, valueType: prometheus.CounterValue},
			//5 - reads merged
			{desc: diskReadsMergeDesc, valueType: prometheus.CounterValue},
			//6 - sectors read
			{desc: diskReadBytesDesc, valueType: prometheus.CounterValue, factor: diskSectorSize},
			//7 - time spent reading (ms)
			{desc: diskReadTimeSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			//8 - writes completed
			{desc: diskWritesCompletedDesc, valueType: prometheus.CounterValue},
			//9 - writes merged
			{desc: diskWriteMergeDesc, valueType: prometheus.CounterValue},
			//10 - sectors written
			{desc: diskWrittenBytesDesc, valueType: prometheus.CounterValue, factor: diskSectorSize},
			//11 - time spent writing (ms)
			{desc: diskWriteTimeSecondsDesc, valueType: prometheus.CounterValue, factor: .001},
			//12 - I/Os currently in progress
			{desc: diskIONowDesc, valueType: prometheus.GaugeValue},
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

	pvDeviceNameMapping,err := updateMapping(p.clientSet, &p.lastPvPathMapping, &p.lastPvPvcMapping, diskDriverName)
	if err != nil {
		logrus.Errorf("Update %s Mapping is failed, err:%s", diskStorageName, err)
		return err
	}
	wg := sync.WaitGroup{}
	for dev, stats := range diskStats {
		pvName, getPv := pvDeviceNameMapping[dev]
		pvcArr, getPvc := p.lastPvPvcMapping[pvName]
		if getPv && getPvc && len(pvcArr) == 2 {
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
		ch <- p.descs[i].mustNewConstMetric(v, pvcNamespace, pvcName, dev, diskStorageName)
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
