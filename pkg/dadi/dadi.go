package dadi

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"strconv"
	"time"
)

const (
	dadiCacheEndpointsName = "alinas-dadi-ds-service"
	dadiHostPath           = "/host/etc/"
	dadiEndPointFile       = "alinas-dadi-endpoint.json"
)

var errMaxCount = 0

type DadiEndPoint struct {
	ContainerEndPoints []string `json:"containerendpoints,omitempty"`
}

func getEndPoints(kubeClient *kubernetes.Clientset, namespace string) *corev1.Endpoints {
	endpoints, err := kubeClient.CoreV1().Endpoints(namespace).Get(context.Background(), dadiCacheEndpointsName, metav1.GetOptions{})
	if err != nil && errMaxCount < 5 {
		errMaxCount++
		log.Errorf("Get %s endpoints is failed, error:%s", dadiCacheEndpointsName, err.Error())
		return nil
	}
	return endpoints
}

func checkEndPoints(endpoints *corev1.Endpoints) bool {
	if len(endpoints.Subsets) == 0 {
		msg := fmt.Sprintf("%s endpoints.Subsets is empty.endpoints:%+v", dadiCacheEndpointsName, endpoints)
		log.Errorf(msg)
		return false
	}
	for _, endpointSubset := range endpoints.Subsets {
		if len(endpointSubset.Addresses) == 0 {
			msg := fmt.Sprintf("%s endpoints.Subsets.Addresses is empty.endpoints:%+v", dadiCacheEndpointsName, endpoints)
			log.Errorf(msg)
			return false
		}
	}
	return true
}

func getEndPoint(endpoints *corev1.Endpoints) []string {
	endPoint := []string{}
	for _, subset := range endpoints.Subsets {
		for _, address := range subset.Addresses {
			for _, port := range subset.Ports {
				endPoint = append(endPoint, address.IP+":"+strconv.Itoa(int(port.Port)))
			}
		}
	}
	return endPoint
}

func Run(kubeClient *kubernetes.Clientset) {
	for {
		endpoint := getEndPoints(kubeClient, "kube-system")
		if endpoint == nil {
			continue
		}
		if !checkEndPoints(endpoint) {
			continue
		}
		dadiEndPoint := DadiEndPoint{}
		dadiEndPoint.ContainerEndPoints = getEndPoint(endpoint)
		b, _ := json.Marshal(dadiEndPoint)
		err := utils.WriteAndSyncFile(dadiHostPath+dadiEndPointFile, b, 644)
		if err != nil {
			log.Errorf("Write %s json file is failed, err:%s", dadiHostPath+dadiEndPointFile, err.Error())
		}
		time.Sleep(time.Second * 10)
	}
}
