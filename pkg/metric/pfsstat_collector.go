package metric

import (
	"errors"
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"strconv"
	"strings"
	"sync"
)

var (
	rdsRawBlockStatLabelNames = []string{"namespace", "pvc", "device", "type"}
)

const (
	volumeDevice       = "/var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices"
	defaultTokenFormat = "/var/lib/kubelet/pods/%s/kubernetes.io~secret/default-token"
	//rdsRawBlockUnit is 4MB
	rdsRawBlockUnit = 4 * 1024 * 1024
)

var (
	rawBlockTotalCapacityDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "rds_raw_block_capacity_total"),
		"The total number of reads completed successfully.",
		diskStatLabelNames, nil,
	)
	rawBlockAvailableCapacityDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "rds_raw_block_capacity_available"),
		"The total number of reads completed successfully.",
		diskStatLabelNames, nil,
	)
)

type rdsRawBlockStatCollector struct {
	enable            bool
	descs             []typedFactorDesc
	lastPvPathMapping map[string]string   //key:pvName, value:mountPath
	lastPvPvcMapping  map[string][]string //key:pvName, value:pvcNamespace, pvcName
	clientSet         *kubernetes.Clientset
}

func init() {
	registerCollector("rdsrawblockstat", NewRdsRawBlockStatCollector)
}

// NewRdsRawBlockStatCollector returns a new Collector exposing  stats.
func NewRdsRawBlockStatCollector() (Collector, error) {
	enable := false
	enableRdsRawBlockMetric := os.Getenv("ENABLE_RDS_RAW_BLOCK_MEREIC")
	if enableRdsRawBlockMetric == "true" {
		enable = true
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

	return &rdsRawBlockStatCollector{
		enable: enable,
		descs: []typedFactorDesc{
			{desc: rawBlockTotalCapacityDesc, valueType: prometheus.CounterValue},
			{desc: rawBlockAvailableCapacityDesc, valueType: prometheus.CounterValue},
		},
		lastPvPathMapping: make(map[string]string, 0),
		lastPvPvcMapping:  make(map[string][]string, 0),
		clientSet:         clientset,
	}, nil
}

func (p *rdsRawBlockStatCollector) Update(ch chan<- prometheus.Metric) error {
	//startTime := time.Now()
	if !p.enable {
		return nil
	}
	rawBlockStats, err := getRdsRawBlockStats()
	if err != nil {
		return fmt.Errorf("couldn't get rds raw block: %s", err)
	}
	pvDeviceNameMapping, err := updateMapping(p.clientSet, &p.lastPvPathMapping, &p.lastPvPvcMapping, diskDriverName)
	if err != nil {
		logrus.Errorf("Update rds %s Mapping is failed, err:%s", diskStorageName, err)
		return err
	}
	wg := sync.WaitGroup{}
	for pvName, stats := range rawBlockStats {
		devName, getPv := pvDeviceNameMapping[pvName]
		pvcArray, getPvc := p.lastPvPvcMapping[pvName]
		if getPv && getPvc && len(pvcArray) == 2 {
			pvcNamespace := pvcArray[0]
			pvcName := pvcArray[1]
			wg.Add(1)
			go func(devArgs string, pvcNamespaceArgs string, pvcNameArgs string, statsArgs []string) {
				defer wg.Done()
				p.setRdsBlockMetric(devArgs, pvcNamespaceArgs, pvcNameArgs, statsArgs, ch)
			}(devName, pvcNamespace, pvcName, stats)
		}
	}
	wg.Wait()
	return nil
}

func (p *rdsRawBlockStatCollector) setRdsBlockMetric(dev string, pvcNamespace string, pvcName string, stats []string, ch chan<- prometheus.Metric) {
	for i, value := range stats {
		if i >= len(p.descs) {
			return
		}
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		ch <- p.descs[i].mustNewConstMetric(v*rdsRawBlockUnit, pvcNamespace, pvcName, dev, diskStorageName)
	}
}

//1.find pv
//2.find pod id by pv
//3.pod is running?
//4.find docker id by pod id
func getRdsRawBlockStats() (map[string][]string, error) {
	pvNameArray, err := getRdsRawBlockPvName()
	if err != nil {
		logrus.Errorf("Get rds raw block pv is failed, err:%s", err)
		return nil, err
	}
	pvStatMapping := make(map[string][]string)
	for _, pvName := range pvNameArray {
		dockerIDArray, err := getRdsDockerIDByPvName(pvName)
		if err != nil {
			logrus.Errorf("Get rds pod name by pv is failed, err:%s", err)
			continue
		}
		for _, dockerID := range dockerIDArray {
			if _, ok := pvStatMapping[pvName]; ok {
				continue
			}
			statArray, err := getStatByDockerID(dockerID)
			if err != nil {
				logrus.Errorf("Get rds stat by docker id %s is failed, err:%s", err)
				continue
			}
			pvStatMapping[pvName] = statArray
		}
	}

	return pvStatMapping, nil
}

func getStatByDockerID(dockerID string) ([]string, error) {
	var stat []string
	cmd := fmt.Sprintf("docker exec %s sh -c \"cd /dev && pfsadm info vda\"", dockerID)
	devInfo, err := utils.Run(cmd)
	if err != nil {
		logrus.Errorf("Execute cmd %s is failed, err:%s", cmd, err)
		return nil, err
	}
	devInfoArray := strings.Split(devInfo, ",")
	nallFlag := false
	nfreeFlag := false
	for _, info := range devInfoArray {
		if strings.Contains(info, "nall") && !nallFlag {
			nall := strings.Split(info, "nall")
			stat = append(stat, nall[1])
			nallFlag = true
		}
		if strings.Contains(info, "nfree") && !nfreeFlag {
			nfree := strings.Split(info, "nfree")
			stat = append(stat, nfree[1])
			nfreeFlag = true
		}
		if nallFlag && nfreeFlag {
			break
		}
	}
	return stat, nil
}

func getRdsRawBlockPvName() ([]string, error) {
	var pvNameArray []string
	cmd := fmt.Sprintf("mount | grep /var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices | grep publish")
	mount, err := utils.Run(cmd)
	if err != nil {
		logrus.Errorf("Execute cmd %s is failed, err:%s", mount, err)
		return nil, err
	}
	mountArray := strings.Fields(mount)
	for _, s := range mountArray {
		if strings.Contains(s, volumeDevice) {
			pvName := strings.Split(s, volumeDevice)
			if len(pvName) == 2 {
				pvNameArray = append(pvNameArray, pvName[1])
			}
		}
	}
	if len(pvNameArray) == 0 {
		return nil, errors.New("Not found pv")
	}
	return pvNameArray, nil

}

func isRunningByRdsPodName(podID string) bool {
	runningPodMountPoint := fmt.Sprintf(defaultTokenFormat, podID)
	cmd := "mount | grep " + runningPodMountPoint
	out, err := utils.Run(cmd)
	if err != nil {
		logrus.Errorf("Execute cmd %s is failed, err:%s", cmd, err)
		return false
	}
	if len(strings.Trim(out, " ")) == 0 {
		//event pod is down
		return false
	}
	return true
}

func getRdsDockerIDByPodID(podID string) (string, error) {
	cmd := "docker ps | grep " + podID + " | grep -v /pause"
	dockerInfo, err := utils.Run(cmd)
	if err != nil {
		return "", nil
	}
	dockerInfoArray := strings.Fields(dockerInfo)
	if len(dockerInfoArray) > 0 {
		return dockerInfoArray[0], nil
	}
	return "", errors.New(fmt.Sprintf("Not found pod name, podID:%s", podID))
}

func getRdsDockerIDByPvName(pvName string) ([]string, error) {
	var dockerIDArray []string
	pvDevPath := volumeDevice + pvName + "/dev"
	podList, err := listDirectory(pvDevPath)
	if err != nil {
		return nil, err
	}
	for _, podID := range podList {
		if !isRunningByRdsPodName(podID) {
			continue
		}
		dockerID, err := getRdsDockerIDByPodID(podID)
		if err != nil {
			logrus.Errorf("Get pod name by docker ID is failed, err:%s", err)
			continue
		}
		dockerIDArray = append(dockerIDArray, dockerID)
	}
	return dockerIDArray, nil
}
