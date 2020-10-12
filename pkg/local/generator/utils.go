package generator

import (
	"context"
	"errors"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

// UpdatePvcWithLabel update pvc
func UpdatePvcWithLabel(ctx context.Context, namespace, pvcName string, annotations map[string]string) error {
	pvc, err := types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumeClaims(namespace).Get(ctx, pvcName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	for key, value := range annotations {
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

// UpdatePvWithLabel udpate pv
func UpdatePvWithLabel(ctx context.Context, pvName string, annotations map[string]string) error {
	pv, err := types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumes().Get(ctx, pvName, metav1.GetOptions{})
	if err != nil {
		return err
	}
	for key, value := range annotations {
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

// CreateVolumeWithLabel create volume
func CreateVolumeWithLabel(namespace, pvcName string, annotations map[string]string) error {
	ctx := context.Background()
	if err := UpdatePvcWithLabel(ctx, namespace, pvcName, annotations); err != nil {
		//log.Errorf("CreateVolumeWithLabel: Failed to update PVC %s/%s: %v", namespace, pvcName, err)
		return err
	}

	for i := 0; i < 8; i++ {
		pvc, err := types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumeClaims(namespace).Get(ctx, pvcName, metav1.GetOptions{})
		if err != nil {
			//log.Errorf("CreateVolumeWithLabel: Failed to read PVC %s/%s: %v", namespace, pvcName, err)
			return err
		}
		for pvc.Annotations[types.VolumeLifecycleLabel] == types.VolumeLifecycleCreated {
			return nil
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
	//log.Errorf("CreateVolumeWithLabel: Fail to create volume with label timeout, %s, %s", namespace, pvcName)
	return errors.New("CreateVolumeWithLabel: Create Volume with label timeout error: " + namespace + pvcName)
}

// DeleteVolumeWithLabel delete volume
func DeleteVolumeWithLabel(pvName string, annotations map[string]string) error {
	ctx := context.Background()
	if err := UpdatePvWithLabel(ctx, pvName, annotations); err != nil {
		//log.Errorf("DeleteVolumeWithLabel: Failed to update PV %s: %v", pvName, err)
		return err
	}

	for i := 0; i < 8; i++ {
		pv, err := types.GlobalConfigVar.KubeClient.CoreV1().PersistentVolumes().Get(ctx, pvName, metav1.GetOptions{})
		if err != nil {
			//log.Errorf("DeleteVolumeWithLabel: Failed to read PV %s: %v", pvName, err)
			return err
		}
		for pv.Annotations[types.VolumeLifecycleLabel] == types.VolumeLifecycleDeleted {
			return nil
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
	//log.Errorf("DeleteVolumeWithLabel: delete volume with label time, %s", pvName)
	return errors.New("DeleteVolumeWithLabel: Delete Volume with label timeout error: " + pvName)
}
