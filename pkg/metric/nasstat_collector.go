package metric

import (
	"github.com/prometheus/client_golang/prometheus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

type nasStatCollector struct {
	descs     []typedFactorDesc
	clientSet *kubernetes.Clientset
}

func init() {
	registerCollector("nasstat", NewNasStatCollector)
}

// NewPVUsageCollector returns a new Collector exposing pv stats.
func NewNasStatCollector() (Collector, error) {
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

	return &nasStatCollector{
		descs: []typedFactorDesc{
			//4 - reads completed successfully
			{desc: nasReadsCompletedDesc, valueType: prometheus.CounterValue,},
			//6 - sectors read
			{desc: nasReadBytesDesc, valueType: prometheus.CounterValue,},
			//7 - time spent reading (ms)
			{desc: nasReadTimeSecondsDesc, valueType: prometheus.CounterValue, factor: .001,},
			//8 - writes completed
			{desc: nasWritesCompletedDesc, valueType: prometheus.CounterValue,},
			//10 - sectors written
			{desc: nasWrittenBytesDesc, valueType: prometheus.CounterValue,},
			//11 - time spent writing (ms)
			{desc: nasWriteTimeSecondsDesc, valueType: prometheus.CounterValue, factor: .001,},
		},
		clientSet: clientset,
	}, nil
}

func (n *nasStatCollector) Update(ch chan<- prometheus.Metric) error {
	/*persistentVolumeList, err := getPersistentVolumeList(n.clientSet)
	if err != nil {
		return err
	}
	for _, persistentVolume := range persistentVolumeList.Items {
		if persistentVolume.Spec.CSI == nil {
			continue
		}
		driverName := strings.ToLower(persistentVolume.Spec.CSI.Driver)
		if driverName != nasDriverName {
			continue
		}
		pvName := persistentVolume.Name

		mountPath, err := getMountPathByVolumeName(pvName)
		if err != nil {
			continue
		}

		nasStat, err := getNasStats(mountPath)
		if err != nil {
			continue
		}

		for i, stat := range nasStat {
			num, err := strconv.ParseFloat(stat, 64)
			if err != nil {
				continue
			}
			ch <- n.descs[i].mustNewConstMetric(num, pvName, NasStorageName)
		}
	}*/
	return nil
}
/*func getNasStats(mountPath string) ([]string, error) {
	cmd := "nfsiostat 1 2 " + mountPath
	err, stdout, stderr := execByShell(cmd)

	nasStatArray := make([]string, 0)
	if err != nil {

		return nasStatArray, err
	}
	if len(stderr) != 0 {

		return nasStatArray, errors.New(stderr)
	}
	scanner := bufio.NewScanner(strings.NewReader(stdout))
	split := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			return advance, token, err
		}
		if atEOF {
			return 0, nil, io.EOF
		}
		return 0, nil, nil
	}

	scanner.Split(split)
	for scanner.Scan() {
		if utils.IsNumeric(scanner.Text()) {
			nasStatArray = append(nasStatArray, scanner.Text())
		}
	}
	return parseNasStats(nasStatArray)
}

func parseNasStats(nasStatArray []string) ([]string, error) {
	nasStats := make([]string, 0)
	length := len(nasStatArray) - 1
	nasStats = append(nasStats, nasStatArray[length-11]) //read IOPS
	nasStats = append(nasStats, nasStatArray[length-10]) //read ThroughPut
	nasStats = append(nasStats, nasStatArray[length-7])  //read latency
	nasStats = append(nasStats, nasStatArray[length-5])  //write IOPS
	nasStats = append(nasStats, nasStatArray[length-4])  //write ThroughPut
	nasStats = append(nasStats, nasStatArray[length-1])  //write latency

	return nasStats, nil
}*/
