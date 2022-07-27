package v1beta1

import (
	"context"
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/dynamic"
)

// getCnfsObject get cnfs crd
func getCnfsObject(client dynamic.Interface, name string) (*ContainerNetworkFileSystem, error) {
	utd, err := client.Resource(GVR).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		log.Errorf("Dynamic get cnfs %s is failed, err:%s", name, err)
		return nil, err
	}
	data, err := utd.MarshalJSON()
	if err != nil {
		log.Errorf("MarshalJSON cnfs %s is failed, err:%s", name, err)
		return nil, err
	}
	var cnfs ContainerNetworkFileSystem
	if err := json.Unmarshal(data, &cnfs); err != nil {
		log.Errorf("Unmarshal cnfs %s is failed, err:%s", name, err)
		return nil, err
	}
	return &cnfs, nil
}

// GetCnfsObject get cnfs's object
func GetCnfsObject(client dynamic.Interface, name string) (*ContainerNetworkFileSystem, error) {
	cnfsObj, err := getCnfsObject(client, name)
	if err != nil {
		log.Errorf("Get cnfs %s is failed, err:%s", name, err)
		return nil, err
	}
	if cnfsObj.Status.Status != "Available" {
		msg := fmt.Sprintf("ContainerNetworkFileSystem %s is not available, status is %s", name, cnfsObj.Status.Status)
		return nil, errors.New(msg)
	}
	return cnfsObj, nil
}
