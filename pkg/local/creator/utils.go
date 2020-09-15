package creator

import (
	"context"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
	log "github.com/sirupsen/logrus"
	"errors"
)
func UpdatePvcWithLabel(namespace, pvcName string, labels map[string]string) error {
	ctx := context.Background()
	pvc, err := types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumeClaims(namespace).Get(ctx, pvcName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	for key, value := range labels {
		if pvc.Annotations == nil {
			pvc.Annotations = map[string]string{key: value}
		} else {
			pvc.Annotations[key] = value
		}
	}
	_, err = types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumeClaims(namespace).Update(ctx, pvc, metav1.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}


func UpdatePvWithLabel(pvName string, labels map[string]string) error {
	ctx := context.Background()
	pv, err := types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumes().Get(ctx, pvName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	for key, value := range labels {
		if pv.Annotations == nil {
			pv.Annotations = map[string]string{key: value}
		} else {
			pv.Annotations[key] = value
		}
	}
	_, err = types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumes().Update(ctx, pv, metav1.UpdateOptions{})
	if err != nil {
		return err
	}
	return nil
}


func CreateVolumeWithLabel(namespace, pvcName string, labels map[string]string) error {
	if err := UpdatePvcWithLabel(namespace, pvcName, labels); err != nil {
		log.Errorf("Failed to update PVC %s/%s: %v", namespace, pvcName, err)
		return err
	}

	ctx := context.Background()
	for i := 0; i < 8; i++ {
		pvc, err := types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumeClaims(namespace).Get(ctx, pvcName, metav1.GetOptions{})
		if err != nil {
			log.Errorf("Failed to read PVC %s/%s: %v", namespace, pvcName, err)
			return err
		}
		for pvc.Annotations[types.VolumeLifecycleLabel] == types.VolumeLifecycleCreated {
			return nil
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
	return errors.New("Create Volume with label timeout error: " + namespace + pvcName)
}

func DeleteVolumeWithLabel(pvName string, labels map[string]string) error {
	if err := UpdatePvWithLabel(pvName, labels); err != nil {
		log.Errorf("Failed to update PV %s: %v", pvName, err)
		return err
	}

	ctx := context.Background()
	for i := 0; i < 8; i++ {
		pv, err := types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumes().Get(ctx, pvName, metav1.GetOptions{})
		if err != nil {
			log.Errorf("Failed to read PV %s: %v", pvName, err)
			return err
		}
		for pv.Annotations[types.VolumeLifecycleLabel] == types.VolumeLifecycleDeleted {
			return nil
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
	return errors.New("Delete Volume with label timeout error: " + pvName)
}
