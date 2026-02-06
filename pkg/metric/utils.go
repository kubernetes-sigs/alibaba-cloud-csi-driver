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

	promdto "github.com/prometheus/client_model/go"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	apicorev1 "k8s.io/api/core/v1"
	apismetav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

var vfOnce = new(sync.Once)
var isVF = false

const containerNetworkFileSystem = "containerNetworkFileSystem"

func newK8sClient() (kubernetes.Interface, error) {
	config, err := options.GetRestConfig()
	if err != nil {
		klog.ErrorS(err, "Failed to get rest config")
		return nil, nil
	}
	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

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

// readAllContent reads all content from a file and replaces newlines with spaces
func readAllContent(path string) (string, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	result := strings.ReplaceAll(string(content), "\n", " ")

	return result, nil
}

func getDiskPvcByPvName(client kubernetes.Interface, pvName, volDataPath string) (*apicorev1.ObjectReference, error) {
	if client == nil {
		return getPvcByVolData(volDataPath)
	}
	pv, err := client.CoreV1().PersistentVolumes().Get(context.Background(), pvName, apismetav1.GetOptions{})
	if err != nil {
		return nil, err
	}
	if pv.Status.Phase == apicorev1.VolumeBound {
		return pv.Spec.ClaimRef, nil
	}
	return nil, errors.New("pvName:" + pv.Name + " status is not bound.")
}

func getPvcByVolData(volDataPath string) (*apicorev1.ObjectReference, error) {
	volDataMap, err := utils.ReadJSONFile(volDataPath)
	klog.InfoS("Volume data map", "map", volDataMap)
	if err != nil {
		return nil, err
	}
	return &apicorev1.ObjectReference{
		Namespace: volDataMap["csi.alibabacloud.com/pvc-namespace"],
		Name:      volDataMap["csi.alibabacloud.com/pvc-name"],
	}, nil
}

func getNasPvcByPvName(client kubernetes.Interface, pvName, volDataPath string) (string, string, error) {
	if client == nil {
		pvc, err := getPvcByVolData(volDataPath)
		klog.InfoS("getNasPvcByPvName", "pvc", pvc)
		if err != nil {
			return "", "", err
		}
		return pvc.Namespace, pvc.Name, nil
	}

	pv, err := client.CoreV1().PersistentVolumes().Get(context.Background(), pvName, apismetav1.GetOptions{})
	if err != nil {
		return "", "", err
	}
	if pv.Spec.CSI != nil {
		return pv.Spec.ClaimRef.Namespace, pv.Spec.ClaimRef.Name, nil
	}
	return "", "", errors.New("pvName:" + pv.Name + " status is not bound.")
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

func parseLatencyThreshold(s string, defaults float64) (float64, error) {
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

func extractMetricsFromMetricVec(fqName, help string, metricVec prometheus.Collector, valueType prometheus.ValueType) (metrics []*Metric) {
	ch := make(chan prometheus.Metric)
	go func() {
		metricVec.Collect(ch)
		close(ch)
	}()

	for metric := range ch {
		desc := metric.Desc()

		gauge := &promdto.Metric{}
		if err := metric.Write(gauge); err != nil {
			klog.ErrorS(err, "Failed to write metric", "desc", desc)
			continue
		}

		value, err := getMetricValue(gauge, valueType)
		if err != nil {
			klog.ErrorS(err, "Failed to get metric value", "desc", desc)
			continue
		}

		metrics = append(metrics, &Metric{
			MetaDesc: &MetaDesc{
				Desc:   desc,
				FQName: fqName,
				Help:   help,
			},
			VariableLabelPairs: gauge.Label,
			Value:              value,
			ValueType:          valueType,
		})
	}
	return metrics
}

func getMetricValue(metric *promdto.Metric, valueType prometheus.ValueType) (float64, error) {
	if metric == nil {
		return 0, nil
	}
	switch valueType {
	case prometheus.CounterValue:
		if metric.Counter == nil || metric.Counter.Value == nil {
			return 0, errors.New("nil metric counter")
		}
		return *metric.Counter.Value, nil
	case prometheus.GaugeValue:
		if metric.Gauge == nil {
			return 0, errors.New("nil metric gauge")
		}
		return *metric.Gauge.Value, nil
	case prometheus.UntypedValue:
		if metric.Untyped == nil {
			return 0, errors.New("nil metric untyped")
		}
		return *metric.Untyped.Value, nil
	default:
		return 0, fmt.Errorf("unsupported value type: %s", valueType.ToDTO().String())
	}
}
