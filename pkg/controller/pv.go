package controller

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/client-go/tools/cache"
)

const (
	diskDriverName string = "diskplugin.csi.alibabacloud.com"
)

func parseObject(obj interface{}) bool {
	if object, ok := obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return false
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return false
		}
		log.Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	return true
}

func (c *CsiController) handlePvAdd(obj interface{}) {
	if !parseObject(obj) {
		return
	}
	pvObj := obj.(*corev1.PersistentVolume)
	c.storePv(pvObj)
}

func (c *CsiController) handlePvDelete(obj interface{}) {
	if !parseObject(obj) {
		return
	}
	pvObj := obj.(*corev1.PersistentVolume)
	if pvObj.Spec.CSI != nil && len(pvObj.Spec.CSI.VolumeHandle) != 0 {
		c.deletePv(pvObj.Spec.CSI.VolumeHandle)
	}
}

func (c *CsiController) handlePvUpdate(obj interface{}) {
	if !parseObject(obj) {
		return
	}
	pvObj := obj.(*corev1.PersistentVolume)
	c.storePv(pvObj)
}
