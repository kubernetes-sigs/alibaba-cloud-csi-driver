package metric

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
)

var (
	rawBlockStatLabelNames = []string{"namespace", "pvc", "device", "type"}
	volumeDevice           = filepath.Join(utils.KubeletRootDir, "/plugins/kubernetes.io/csi/volumeDevices/")
	defaultTokenFormat     = filepath.Join(utils.KubeletRootDir, "/pods/%s/volumes/kubernetes.io~secret/default-token")
)

const (
	//RawBlockUnit is 4MB
	pfsRawBlockUnit       = 4 * 1024 * 1024
	notFoundvolumeDevices = "Not found volumeDevices"
)

type pfsRawBlockStatCollector struct {
	enable           bool
	descs            []typedFactorDesc
	pvPfsInfoMap     map[string]pfsInfo
	rawBlockStatsMap map[string][]string
	kubeClient       *kubernetes.Clientset
	dockerClient     *client.Client
	capacityMutex    sync.Mutex
}

type pfsInfo struct {
	PvcNamespace    string
	PvcName         string
	DiskID          string
	DeviceName      string
	VolDataPath     string
	GlobalMountPath string
}

func init() {
	registerCollector("pfs_block_stat", NewPfsRawBlockStatCollector)
}

func (p *pfsRawBlockStatCollector) isEnable() bool {
	if !p.enable {
		return false
	}
	return true
}

func (p *pfsRawBlockStatCollector) updateStatByPolling() {
	if !p.isEnable() {
		return
	}
	var err error
	for {
		p.capacityMutex.Lock()
		doUpdate := true
		var volJSONPaths []string
		p.rawBlockStatsMap, err = getPfsRawBlockStats(&p.dockerClient)
		if err != nil {
			if err.Error() != notFoundvolumeDevices {
				msg := fmt.Sprintf("Couldn't get pfs raw block: %s", err)
				logrus.Errorf(msg)
				doUpdate = false
			}
		} else {
			volJSONPaths, err = findVolJSONByPfsRawBlock(rawBlockRootPath)
			if err != nil {
				logrus.Errorf("Find disk vol_data json is failed, err:%s", err)
				doUpdate = false
			}
		}

		if doUpdate {
			p.updateMap(p.kubeClient, &p.pvPfsInfoMap, volJSONPaths, diskDriverName, "volumeDevices")
		}

		p.capacityMutex.Unlock()
		time.Sleep(60 * time.Second)
	}
}

// NewPfsRawBlockStatCollector returns a new Collector exposing stats.
func NewPfsRawBlockStatCollector() (Collector, error) {
	enable := false
	enablePfsRawBlockMetric := os.Getenv("ENABLE_PFS_BLOCK_MEREIC")
	if enablePfsRawBlockMetric == "true" {
		enable = true
	}
	config, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
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

	pfsBlockCollector := pfsRawBlockStatCollector{
		enable: enable,
		descs: []typedFactorDesc{
			{desc: diskCapacityTotalDesc, valueType: prometheus.CounterValue},
			{desc: diskCapacityAvailableDesc, valueType: prometheus.CounterValue},
			{desc: diskCapacityUsedDesc, valueType: prometheus.CounterValue},
		},
		pvPfsInfoMap: make(map[string]pfsInfo, 0),

		kubeClient:    clientset,
		dockerClient:  client,
		capacityMutex: sync.Mutex{},
	}
	go pfsBlockCollector.updateStatByPolling()
	return &pfsBlockCollector, nil
}

func (p *pfsRawBlockStatCollector) Update(ch chan<- prometheus.Metric) error {
	if !p.isEnable() {
		return nil
	}

	p.capacityMutex.Lock()
	defer p.capacityMutex.Unlock()
	wg := sync.WaitGroup{}
	for pvName, stats := range p.rawBlockStatsMap {
		info, ok := p.pvPfsInfoMap[pvName]
		if ok {
			wg.Add(1)
			go func(deviceNameArgs string, pvcNamespaceArgs string, pvcNameArgs string, statsArgs []string) {
				defer wg.Done()
				p.setPfsRawBlockMetric(deviceNameArgs, pvcNamespaceArgs, pvcNameArgs, statsArgs, ch)
			}(info.DeviceName, info.PvcNamespace, info.PvcName, stats)
		}
	}
	wg.Wait()

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

func (p *pfsRawBlockStatCollector) updateMap(clientSet *kubernetes.Clientset, lastPvPfsInfoMap *map[string]pfsInfo, jsonPaths []string, deriverName string, keyword string) {
	thisPvStorageInfoMap := make(map[string]pfsInfo, 0)
	lineArr, err := utils.RunWithFilter("mount", "csi", keyword)
	if err != nil {
		p.updatePfsInfoMap(clientSet, thisPvStorageInfoMap, lastPvPfsInfoMap)
		return
	}
	if err != nil {
		logrus.Errorf("Execute command pfs is failed, err: %s", err)
		return
	}
	for _, path := range jsonPaths {
		//Get disk pvName
		pvName, diskID, err := getVolumeInfoByJSON(path, deriverName)
		if err != nil {
			if err.Error() != "VolumeType is not the expected type" {
				logrus.Errorf("Get volume info by path %s is failed, err:%s", path, err)
			}
			continue
		}
		for _, line := range lineArr {
			if !strings.Contains(line, "/"+pvName+"/") {
				continue
			}
		}

		deviceName, err := getDeviceByVolumeID(pvName, diskID)
		if err != nil {
			logrus.Errorf("Get dev name by diskID %s is failed, err:%s", diskID, err)
			continue
		}
		strorageInfo := pfsInfo{
			DiskID:      diskID,
			DeviceName:  deviceName,
			VolDataPath: path,
		}
		thisPvStorageInfoMap[pvName] = strorageInfo
	}

	//If there is a change: add, modify, delete
	p.updatePfsInfoMap(clientSet, thisPvStorageInfoMap, lastPvPfsInfoMap)
}

func (p *pfsRawBlockStatCollector) updatePfsInfoMap(clientSet *kubernetes.Clientset, thisPvStorageInfoMap map[string]pfsInfo, lastPvStorageInfoMap *map[string]pfsInfo) {
	for pv, thisInfo := range thisPvStorageInfoMap {
		lastInfo, ok := (*lastPvStorageInfoMap)[pv]
		// add and modify
		if !ok || thisInfo.VolDataPath != lastInfo.VolDataPath {
			pvcNamespace, pvcName, err := getPvcByPvNameByDisk(clientSet, pv)
			if err != nil {
				continue
			}
			updateInfo := pfsInfo{
				DiskID:       thisInfo.DiskID,
				VolDataPath:  thisInfo.VolDataPath,
				DeviceName:   thisInfo.DeviceName,
				PvcName:      pvcName,
				PvcNamespace: pvcNamespace,
			}
			(*lastPvStorageInfoMap)[pv] = updateInfo
		}
	}
	//if pv exist thisPvStorageInfoMap and not exist lastPvStorageInfoMap, pv should be deleted
	for lastPv := range *lastPvStorageInfoMap {
		_, ok := thisPvStorageInfoMap[lastPv]
		if !ok {
			delete(*lastPvStorageInfoMap, lastPv)
		}
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

// 1.find pv
// 2.find pod id by pv
// 3.pod is running?
// 4.find docker id by pod id
func getPfsRawBlockStats(dockerClient **client.Client) (map[string][]string, error) {
	pvNameArray, err := getPfsRawBlockPvName()
	if pvNameArray == nil && err.Error() == notFoundvolumeDevices {
		return nil, err
	}
	if err != nil {
		logrus.Errorf("Get raw block pv is failed, err:%s", err)
		return nil, err
	}

	pvStatMapping := make(map[string][]string)
	filterArgs := filters.NewArgs()
	filterArgs.Add("status", "running")
	containerListOptions := types.ContainerListOptions{Filters: filterArgs}
	containers, err := listContainerByDocker(dockerClient, containerListOptions)
	if err != nil {
		return nil, err
	}

	for _, pvName := range pvNameArray {
		dockerIDArray, err := getDockerIDByPvName(pvName, containers)
		if err != nil {
			logrus.Errorf("Get pod name by pv is failed, err:%s", err)
			continue
		}

		for _, dockerID := range dockerIDArray {
			//rawblock is used by multi pod
			if _, ok := pvStatMapping[pvName]; ok {
				continue
			}
			statArray, err := getStatByDockerID(dockerID, dockerClient)
			if err != nil {
				logrus.Errorf("Get stat by docker id %s is failed, err:%s", dockerID, err)
				continue
			}
			pvStatMapping[pvName] = statArray
		}
	}
	return pvStatMapping, nil
}

func getStatByDockerID(dockerID string, dockerClient **client.Client) ([]string, error) {
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

	return stat, nil
}

func getPfsRawBlockPvName() ([]string, error) {
	var pvNameArray []string
	path := fmt.Sprintf("%s/plugins/kubernetes.io/csi/volumeDevices", utils.KubeletRootDir)
	mountArr, err := utils.RunWithFilter("mount", path, "staging")
	if err != nil {
		return nil, errors.New(notFoundvolumeDevices)
	}

	for _, s := range mountArr {
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
	stdout, err := utils.RunWithFilter("mount", runningPodMountPoint)
	if err != nil {
		return false
	}
	if len(stdout) == 0 {
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
