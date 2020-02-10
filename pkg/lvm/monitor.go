package lvm

import (
	log "github.com/sirupsen/logrus"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/clientcmd"
	"time"

	"encoding/json"
	"fmt"
	"github.com/mesosphere/csilvm/pkg/lvm"
	"github.com/pkg/errors"
	"os"
	"strconv"
	"strings"
)

const (
	// DefaultSleepSecond is the time period to check update node status
	DefaultSleepSecond = 30
	// MEGABYTE is used to convert bytes string to  megabytes string
	MEGABYTE = 1024 * 1024
	// OperatorTypeRemove is the remove type to update node status
	OperatorTypeRemove = "remove"
	// OperatorTypeAdd is the add type to update node status
	OperatorTypeAdd = "add"
	// OperatorPath is the path to update node status
	OperatorPath = "/status/capacity/lvm.csi.alibabacloud.com~1"
	// LvmCapacityKey is the key to update node status
	LvmCapacityKey = "lvm.csi.alibabacloud.com"
)

//PatchStringValue specifies a patch operation for a string.
type PatchStringValue struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value"`
}

// PatchNodeParam specifies a patch operation
type PatchNodeParam struct {
	coreV1       v1.CoreV1Interface
	NodeName     string                 `json:"nodeName"`
	OperatorType string                 `json:"operator_type"`
	OperatorPath string                 `json:"operator_path"`
	OperatorData map[string]interface{} `json:"operator_data"`
}

// RemoveStringValue specifies a remove operation for a string
type RemoveStringValue struct {
	Op   string `json:"op"`
	Path string `json:"path"`
}

// RemoveNodeInfoParam specifies a remove operation
type RemoveNodeInfoParam struct {
	coreV1       v1.CoreV1Interface
	NodeName     string   `json:"nodeName"`
	OperatorType string   `json:"operator_type"`
	OperatorPath string   `json:"operator_path"`
	OperatorKey  []string `json:"operator_key"`
}

// Monitor start a loop to update node's capacity status
func Monitor() {

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	nodeName, err := getNodeName()
	if err != nil {
		log.Fatalf("Error get node name: %s", err.Error())
	}

	// monitoring in loop
	for {
		paramData, err := getPatchNodeParamData()
		if err != nil {
			log.Fatalf("Error monitoring vg resource: %s", err.Error())
		}

		removed := checkUpdate(nodeName, paramData, kubeClient)
		if len(removed) > 0 {
			param := &RemoveNodeInfoParam{}
			param.coreV1 = kubeClient.CoreV1()
			param.NodeName = nodeName
			param.OperatorType = OperatorTypeRemove
			param.OperatorPath = OperatorPath
			param.OperatorKey = removed
			_, err := removeNodeCapacityStatus(param)
			if err != nil {
				log.Fatalf("Error monitoring vg resource: %s", err.Error())
			}
		}

		param := &PatchNodeParam{}
		param.coreV1 = kubeClient.CoreV1()
		param.NodeName = nodeName
		param.OperatorType = OperatorTypeAdd
		param.OperatorPath = OperatorPath
		param.OperatorData = paramData

		_, err = updateNodeCapacityStatus(param)
		if err != nil {
			log.Fatalf("Error monitoring vg resource: %s", err.Error())
		}

		time.Sleep(DefaultSleepSecond * time.Second)
	}
}

//patch node info, example label, annotation
func updateNodeCapacityStatus(param *PatchNodeParam) (*apiv1.Node, error) {
	coreV1 := param.coreV1
	nodeSn := param.NodeName

	_, err := coreV1.Nodes().Get(nodeSn, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	operatorData := param.OperatorData
	operatorType := param.OperatorType
	operatorPath := param.OperatorPath

	var payloads []interface{}

	for key, value := range operatorData {
		payload := PatchStringValue{
			Op:    operatorType,
			Path:  operatorPath + key,
			Value: value,
		}

		payloads = append(payloads, payload)

	}

	payloadBytes, _ := json.Marshal(payloads)

	newNode, err := coreV1.Nodes().Patch(nodeSn, types.JSONPatchType, payloadBytes, "status")

	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}

	return newNode, err
}

func getPatchNodeParamData() (map[string]interface{}, error) {
	vgnames, err := lvm.ListVolumeGroupNames()
	if err != nil {
		return nil, fmt.Errorf("error ListVolumeGroupNames: %s", err.Error())
	}

	data := make(map[string]interface{})
	for _, vgname := range vgnames {
		vg, err := lvm.LookupVolumeGroup(vgname)
		if err != nil {
			return nil, fmt.Errorf("error LookupVolumeGroup: %s", err.Error())
		}
		bytesFree, err := vg.BytesFree()
		if err != nil {
			return nil, fmt.Errorf("error get BytesFree: %s", err.Error())
		}
		data[vgname] = strconv.FormatUint(bytesFree/MEGABYTE, 10) + "Mi"
	}

	return data, nil
}

func getNodeName() (string, error) {
	nodeName := os.Getenv("KUBE_NODE_NAME")
	if len(nodeName) == 0 {
		return "", errors.Errorf("KUBE_NODE_NAME environment variable is not set")
	}

	return nodeName, nil
}

func checkUpdate(nodeName string, data map[string]interface{}, kubeClient kubernetes.Interface) []string {
	added := []string{}
	removed := []string{}

	node, err := kubeClient.CoreV1().Nodes().Get(nodeName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error get node: %s", err.Error())
	}

	for k := range node.Status.Capacity {
		if strings.Contains(k.String(), LvmCapacityKey) {
			added = append(added, strings.Split(k.String(), "/")[1])
		}
	}

	for _, k := range added {
		if _, ok := data[k]; !ok {
			removed = append(removed, k)
		}
	}

	return removed
}

func removeNodeCapacityStatus(param *RemoveNodeInfoParam) (*apiv1.Node, error) {
	coreV1 := param.coreV1
	nodeSn := param.NodeName

	_, err := coreV1.Nodes().Get(nodeSn, metav1.GetOptions{})

	if err != nil {
		return nil, err
	}

	operatorKey := param.OperatorKey
	operatorType := param.OperatorType
	operatorPath := param.OperatorPath

	var payloads []interface{}

	for _, key := range operatorKey {
		payload := RemoveStringValue{
			Op:   operatorType,
			Path: operatorPath + key,
		}

		payloads = append(payloads, payload)
	}

	payloadBytes, _ := json.Marshal(payloads)

	newNode, err := coreV1.Nodes().Patch(nodeSn, types.JSONPatchType, payloadBytes, "status")

	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}

	return newNode, err
}
