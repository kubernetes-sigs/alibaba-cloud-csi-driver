package metric

import (
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	rawBlockStatLabelNames = []string{"namespace", "pvc", "device", "type"}
)

const (
	volumeDevice       = "/var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/"
	defaultTokenFormat = "/var/lib/kubelet/pods/%s/volumes/kubernetes.io~secret/default-token"
	//RawBlockUnit is 4MB
	pfsRawBlockUnit = 4 * 1024 * 1024
)

var (
	rawBlockTotalCapacityDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "pfs_block_capacity_total"),
		"The total capacity of pfs raw block.",
		diskStatLabelNames, nil,
	)
	rawBlockAvailableCapacityDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "pfs_block_capacity_available"),
		"The available capacity of pfs raw block.",
		diskStatLabelNames, nil,
	)
	rawBlockUsedCapacityDesc = prometheus.NewDesc(
		prometheus.BuildFQName(nodeNamespace, volumeSubSystem, "pfs_block_capacity_used"),
		"The used capacity of pfs raw block.",
		diskStatLabelNames, nil,
	)
)

type pfsRawBlockStatCollector struct {
	enable               bool
	descs                []typedFactorDesc
	lastPvStorageInfoMap map[string]storageInfo
	kubeClient           *kubernetes.Clientset
	dockerClient         *client.Client
}

func init() {
	registerCollector("pfsblockstat", NewPfsRawBlockStatCollector)
}

// NewPfsRawBlockStatCollector returns a new Collector exposing stats.
func NewPfsRawBlockStatCollector() (Collector, error) {
	enable := false
	enablePfsRawBlockMetric := os.Getenv("ENABLE_PFS_BLOCK_MEREIC")
	if enablePfsRawBlockMetric == "true" {
		enable = true
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		logrus.Errorf("Get kubernetes client config is failed, err:%s", err)
		return nil, err
	}
	client, err := client.NewClientWithOpts()
	if err != nil {
		logrus.Errorf("New docker client is failed, err:%s", err)
		return nil, err
	}
	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logrus.Errorf("New kubernetes client is failed, err:%s", err)
		return nil, err
	}

	return &pfsRawBlockStatCollector{
		enable: enable,
		descs: []typedFactorDesc{
			{desc: rawBlockTotalCapacityDesc, valueType: prometheus.CounterValue},
			{desc: rawBlockAvailableCapacityDesc, valueType: prometheus.CounterValue},
			{desc: rawBlockUsedCapacityDesc, valueType: prometheus.CounterValue},
		},
		lastPvStorageInfoMap: make(map[string]storageInfo, 0),
		kubeClient:           clientset,
		dockerClient:         client,
	}, nil
}

func (p *pfsRawBlockStatCollector) Update(ch chan<- prometheus.Metric) error {
	//1.getPfsRawBlockStats
	startTime := time.Now()
	if !p.enable {
		return nil
	}
	rawBlockStats, err := getPfsRawBlockStats(&p.dockerClient)
	if err != nil {
		msg := fmt.Sprintf("Couldn't get pfs raw block: %s", err)
		logrus.Errorf(msg)
		return errors.New(msg)
	}
	elapsedTime := time.Since(startTime)
	logrus.Info("getPfsRawBlockStats spent time:", elapsedTime)

	//2.findVolJSONByPfsRawBlock
	startTime = time.Now()
	volJSONPaths, err := findVolJSONByPfsRawBlock(rawBlockRootPath)
	if err != nil {
		logrus.Errorf("Find disk vol_data json is failed, err:%s", err)
		return err
	}
	elapsedTime = time.Since(startTime)
	logrus.Info("findVolJSONByPfsRawBlock spent time:", elapsedTime)

	//3.updateMap
	startTime = time.Now()
	updateMap(p.kubeClient, &p.lastPvStorageInfoMap, volJSONPaths, diskDriverName)
	elapsedTime = time.Since(startTime)
	logrus.Info("updateMap spent time:", elapsedTime)

	//4.setPfsRawBlockMetric
	startTime = time.Now()
	wg := sync.WaitGroup{}
	for pvName, stats := range rawBlockStats {
		info, ok := p.lastPvStorageInfoMap[pvName]
		if ok {
			wg.Add(1)
			go func(deviceNameArgs string, pvcNamespaceArgs string, pvcNameArgs string, statsArgs []string) {
				defer wg.Done()
				p.setPfsRawBlockMetric(deviceNameArgs, pvcNamespaceArgs, pvcNameArgs, statsArgs, ch)
			}(info.DeviceName, info.PvcNamespace, info.PvcName, stats)
		}
	}
	wg.Wait()
	elapsedTime = time.Since(startTime)
	logrus.Info("setPfsRawBlockMetric spent time:", elapsedTime)

	return nil
}

func (p *pfsRawBlockStatCollector) setPfsRawBlockMetric(dev string, pvcNamespace string, pvcName string, stats []string, ch chan<- prometheus.Metric) {
	for i, value := range stats {
		if i >= len(p.descs) {
			return
		}
		v, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return
		}
		ch <- p.descs[i].mustNewConstMetric(v*pfsRawBlockUnit, pvcNamespace, pvcName, dev, pfsBlockName)
	}
}

func findVolJSONByPfsRawBlock(rootDir string) ([]string, error) {
	resDir := make([]string, 0)
	rootFiles, err := ioutil.ReadDir(rootDir)
	if err != nil {
		return nil, err
	}
	for _, rootFile := range rootFiles {
		csiDir := rootDir + rootFile.Name() + "/" + "data/"
		csiFiles, err := ioutil.ReadDir(csiDir)
		if err != nil {
			continue
		}
		for _, csiFile := range csiFiles {
			if csiFile.Name() == volDataFile {
				resDir = append(resDir, csiDir+volDataFile)
			}
		}
	}
	return resDir, err
}

//1.find pv
//2.find pod id by pv
//3.pod is running?
//4.find docker id by pod id
func getPfsRawBlockStats(dockerClient **client.Client) (map[string][]string, error) {
	pvNameArray, err := getPfsRawBlockPvName()
	if err != nil {
		logrus.Errorf("Get  raw block pv is failed, err:%s", err)
		return nil, err
	}
	//1.listContainerByDocker
	startTime := time.Now()
	pvStatMapping := make(map[string][]string)
	containers, err := listContainerByDocker(dockerClient)
	if err != nil {
		return nil, err
	}
	elapsedTime := time.Since(startTime)
	logrus.Info("listContainerByDocker spent time:", elapsedTime)

	for _, pvName := range pvNameArray {
		//2.getDockerIDByPvName
		startTime := time.Now()
		dockerIDArray, err := getDockerIDByPvName(pvName, containers)
		if err != nil {
			logrus.Errorf("Get pod name by pv is failed, err:%s", err)
			continue
		}
		elapsedTime := time.Since(startTime)
		logrus.Info("getDockerIDByPvName spent time:", elapsedTime)

		for _, dockerID := range dockerIDArray {
			//rawblock is used by multi pod
			//3.getStatByDockerID
			startTime := time.Now()
			if _, ok := pvStatMapping[pvName]; ok {
				continue
			}
			statArray, err := getStatByDockerID(dockerID, dockerClient)
			if err != nil {
				logrus.Errorf("Get stat by docker id %s is failed, err:%s", dockerID, err)
				continue
			}
			pvStatMapping[pvName] = statArray
			elapsedTime := time.Since(startTime)
			logrus.Info("getStatByDockerID spent time:", elapsedTime)
		}
	}
	return pvStatMapping, nil
}

func getStatByDockerID(dockerID string, dockerClient **client.Client) ([]string, error) {
	//1.docker exec
	startTime := time.Now()
	var stat []string
	cmd := []string{"sh", "-c", "cd /dev && pfsadm info vda"}
	execResult, err := execByDocker(context.TODO(), *dockerClient, dockerID, cmd)
	if err != nil {
		logrus.Errorf("Execute cmd %s is failed, err:%s", cmd, err)
		return nil, err
	}
	if len(execResult.errBuffer.String()) != 0 {
		logrus.Errorf("Execute cmd %s is failed, errBuffer:%s", cmd, execResult.errBuffer.String())
		return nil, errors.New(execResult.errBuffer.String())
	}
	elapsedTime := time.Since(startTime)
	logrus.Info("execByDocker spent time:", elapsedTime)

	//1.deal with result
	startTime = time.Now()
	pfsadmInfo := execResult.outBuffer.String()
	pfsadmInfoArray := strings.Split(pfsadmInfo, ",")
	nallFlag := false
	nfreeFlag := false
	for _, info := range pfsadmInfoArray {
		if strings.Contains(info, "nall ") && !nallFlag {
			nall := strings.Split(info, "nall ")
			stat = append(stat, nall[1])
			nallFlag = true
		}
		if strings.Contains(info, "nfree ") && !nfreeFlag {
			nfree := strings.Split(info, "nfree ")
			stat = append(stat, nfree[1])
			nfreeFlag = true
		}
		if nallFlag && nfreeFlag {
			break
		}
	}
	if len(stat) >= 2 {
		totalCapacity, err1 := strconv.Atoi(stat[0])
		if err1 != nil {
			logrus.Errorf("Convert totalCapacity %s to int is failed, err:%s", stat[0], err)
		}
		availableCapacity, err2 := strconv.Atoi(stat[1])
		if err2 != nil {
			logrus.Errorf("Convert availableCapacity %s to int is failed, err:%s", stat[1], err)
		}
		usedCapacity := totalCapacity - availableCapacity
		if err1 != nil || err2 != nil || usedCapacity < 0 {
			usedCapacity = 0
		}
		stat = append(stat, strconv.Itoa(usedCapacity))
	}
	elapsedTime = time.Since(startTime)
	logrus.Info("DealWithResult spent time:", elapsedTime)
	return stat, nil
}

func getPfsRawBlockPvName() ([]string, error) {
	var pvNameArray []string
	cmd := fmt.Sprintf("mount | grep /var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices | grep staging")
	mount, err := utils.Run(cmd)
	if err != nil {
		logrus.Errorf("Execute cmd %s is failed, err:%s", mount, err)
		return nil, err
	}
	mountArray := strings.Fields(mount)
	for _, s := range mountArray {
		if strings.Contains(s, volumeDevice+"staging/") {
			pvName := strings.Split(s, volumeDevice+"staging/")
			if len(pvName) >= 2 {
				keyWord := "/"
				end := strings.Index(pvName[1], keyWord)
				pvNameArray = append(pvNameArray, pvName[1][:end])
			}
		}
	}
	if len(pvNameArray) == 0 {
		return nil, errors.New("Not found pv")
	}
	return pvNameArray, nil

}

func isRunningByPodName(podID string) bool {
	runningPodMountPoint := fmt.Sprintf(defaultTokenFormat, podID)
	cmd := "mount | grep " + runningPodMountPoint
	out, err := utils.Run(cmd)
	if err != nil {
		return false
	}
	if len(strings.Trim(out, " ")) == 0 {
		//event pod is down
		return false
	}
	return true
}

func getDockerIDByPodID(podID string, containers []types.Container) (string, error) {
	for _, container := range containers {
		if strings.Contains(container.Command, "/pause") {
			continue
		}
		for _, name := range container.Names {
			if strings.Contains(name, podID) {
				return container.ID, nil
			}
		}
	}

	return "", fmt.Errorf(fmt.Sprintf("Not found pod name, podID:%s", podID))
}

func getDockerIDByPvName(pvName string, containers []types.Container) ([]string, error) {
	var dockerIDArray []string
	pvDevPath := volumeDevice + pvName + "/dev"
	podIDList, err := listDirectory(pvDevPath)
	if err != nil {
		logrus.Errorf("List Directory %s is failed, err:%s", pvDevPath, err)
		return nil, err
	}
	for _, podID := range podIDList {
		if !isRunningByPodName(podID) {
			continue
		}
		dockerID, err := getDockerIDByPodID(podID, containers)
		if err != nil {
			logrus.Errorf("Get pod name by docker ID is failed, err:%s", err)
			continue
		}
		dockerIDArray = append(dockerIDArray, dockerID)
	}
	return dockerIDArray, nil
}
