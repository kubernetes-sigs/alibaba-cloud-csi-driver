package v1alpha1

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/dynamic"
)

func getContainerNetworkFileSystem(client dynamic.Interface, name string) (*ContainerNetworkFileSystem, error) {
	utd, err := client.Resource(GVR).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		logrus.Errorf("Dynamic get ContainerNetworkFileSystem %s is failed, err:%s",name,err)
		return nil, err
	}
	data, err := utd.MarshalJSON()
	if err != nil {
		logrus.Errorf("MarshalJSON ContainerNetworkFileSystem %s is failed, err:%s",name,err)
		return nil, err
	}
	var cnfs ContainerNetworkFileSystem
	if err := json.Unmarshal(data, &cnfs); err != nil {
		logrus.Errorf("Unmarshal ContainerNetworkFileSystem %s is failed, err:%s",name,err)
		return nil, err
	}
	return &cnfs, nil
}


func GetContainerNetworkFileSystemServer(client dynamic.Interface, name string) (string, error) {
	cnfsObj, err := getContainerNetworkFileSystem(client, name)
	if err != nil {
		return "", err
	}
	if cnfsObj.Status.Status != "Available" {
		msg := fmt.Sprintf("ContainerNetworkFileSystem %s is not available, status is %s", name, cnfsObj.Status.Status)
		return "", errors.New(msg)
	}
	return cnfsObj.Status.FsAttributes.Server, nil
}
