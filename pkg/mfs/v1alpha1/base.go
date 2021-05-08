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

func getManagedFileSystem(client dynamic.Interface, name string) (*ManagedFileSystem, error) {
	utd, err := client.Resource(GVR).Get(context.TODO(), name, metav1.GetOptions{})
	if err != nil {
		logrus.Errorf("Dynamic get ManagedFileSystem %s is failed, err:%s",name,err)
		return nil, err
	}
	data, err := utd.MarshalJSON()
	if err != nil {
		logrus.Errorf("MarshalJSON ManagedFileSystem %s is failed, err:%s",name,err)
		return nil, err
	}
	var mfs ManagedFileSystem
	if err := json.Unmarshal(data, &mfs); err != nil {
		logrus.Errorf("Unmarshal ManagedFileSystem %s is failed, err:%s",name,err)
		return nil, err
	}
	return &mfs, nil
}


func GetManagedFileSystemServer(client dynamic.Interface, name string) (string, error) {
	mfsObj, err := getManagedFileSystem(client, name)
	if err != nil {
		return "", err
	}
	if mfsObj.Status.Status != "Available" {
		msg := fmt.Sprintf("ManagedFileSystem %s is not available, status is %s", name, mfsObj.Status.Status)
		return "", errors.New(msg)
	}
	return mfsObj.Spec.FsAttributes.Server, nil
}
