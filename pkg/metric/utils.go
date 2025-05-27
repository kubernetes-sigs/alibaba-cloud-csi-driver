package metric

import (
	"bufio"
	"context"
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cnfs/v1beta1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	apicorev1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

var vfOnce = new(sync.Once)
var isVF = false

const containerNetworkFileSystem = "containerNetworkFileSystem"

func readFirstLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if len(lines) == 0 {
		return lines, scanner.Err()
	}
	lineStrArray := strings.Split(lines[0], " ")
	return lineStrArray, scanner.Err()
}

func getDiskPvcByPvName(clientSet *kubernetes.Clientset, pvName string) (*apicorev1.ObjectReference, error) {
	pv, err := clientSet.CoreV1().PersistentVolumes().Get(context.Background(), pvName, apismetav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if pv.Status.Phase == apicorev1.VolumeBound {
		return pv.Spec.ClaimRef, nil
	}
	return nil, errors.New("pvName:" + pv.Name + " status is not bound.")
}

func getNasPvcByPvName(clientSet *kubernetes.Clientset, cnfsClient dynamic.Interface, pvName string) (string, string, string, error) {
	pv, err := clientSet.CoreV1().PersistentVolumes().Get(context.Background(), pvName, apismetav1.GetOptions{})
	if err != nil {
		return "", "", "", err
	}
	if pv.Spec.CSI != nil {
		if val, ok := pv.Spec.CSI.VolumeAttributes["server"]; ok {
			if pv.Status.Phase == apicorev1.VolumeBound {
				return pv.Spec.ClaimRef.Namespace, pv.Spec.ClaimRef.Name, val, nil
			}
		} else if value, ok := pv.Spec.CSI.VolumeAttributes[containerNetworkFileSystem]; ok {
			cnfs, err := v1beta1.GetCnfsObject(cnfsClient, value)
			if err != nil {
				klog.Errorf("Get cnfs %s server is failed, err:%s", value, err)
				return "", "", "", err
			}
			return pv.Spec.ClaimRef.Namespace, pv.Spec.ClaimRef.Name, cnfs.Status.FsAttributes.Server, nil
		}
	}
	return "", "", "", errors.New("pvName:" + pv.Name + " status is not bound.")
}

var ErrUnexpectedVolumeType = errors.New("VolumeType is not the expected type")

func getVolumeInfoByJSON(volDataJSONPath string, volType string) (string, string, error) {
	volDataMap, err := utils.ReadJSONFile(volDataJSONPath)
	if err != nil {
		klog.Errorf("Read json path %s is failed, err:%s", volDataJSONPath, err)
		return "", "", err
	}
	if volDataMap["driverName"] == volType {
		return volDataMap["specVolID"], volDataMap["volumeHandle"], nil
	}
	return "", "", ErrUnexpectedVolumeType
}

func findVolJSON(rootDir string) ([]string, error) {
	return filepath.Glob(filepath.Join(rootDir, "*", csiMountKeyWords, "*", volDataFile))
}

func listDirectory(rootPath string) ([]string, error) {
	var fileLists []string
	files, err := os.ReadDir(rootPath)
	if err != nil {
		klog.Errorf("List Directory %s is failed, err:%s", rootPath, err.Error())
		return nil, err
	}
	for _, f := range files {
		fileLists = append(fileLists, f.Name())
	}
	return fileLists, nil
}

func parseLantencyThreshold(s string, defaults float64) (float64, error) {
	var thresholdNum int
	var thresholdUnit string
	_, err := fmt.Sscanf(s, "%d%s", &thresholdNum, &thresholdUnit)
	if err != nil {
		klog.Errorf("Parse latency threshold %s is failed, err:%s", s, err)
		return defaults, err
	}
	switch thresholdUnit {
	case "s", "second", "seconds":
		return float64(thresholdNum * 1000), nil
	case "ms", "millisecond", "milliseconds":
		return float64(thresholdNum), nil
	case "us", "microsecond", "microseconds":
		return float64(thresholdNum / 1000), nil
	default:
		return defaults, nil
	}
}

func parseCapacityThreshold(s string, defaults float64) (float64, error) {
	var thresholdNum float64
	_, err := fmt.Sscanf(s, "%f", &thresholdNum)
	if err != nil {
		klog.Errorf("Parse  threshold %s is failed, err:%s", s, err)
		return defaults, err
	}
	return thresholdNum, nil
}

func getGlobalMountPathByDiskID(diskID string) string {
	hash := sha256.Sum256([]byte(diskID))
	return filepath.Join(utils.KubeletRootDir, fmt.Sprintf("/plugins/kubernetes.io/csi/%s/%x/globalmount", diskDriverName, hash))
}
