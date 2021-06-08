package controller

import (
	corev1 "k8s.io/api/core/v1"
	coreinformers "k8s.io/client-go/informers/core/v1"
	corelisters "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"sync"
)

//CsiController use it to watch pv
type CsiController struct {
	pvLister corelisters.PersistentVolumeLister
	pvSynced cache.InformerSynced
	pvList   *sync.Map
}

//CsiControllerInstance is csi controller singleton
var CsiControllerInstance *CsiController

//NewController function is new CsiController object
func NewController(pvInformer coreinformers.PersistentVolumeInformer) {
	if CsiControllerInstance == nil {
		controller := &CsiController{
			pvLister: pvInformer.Lister(),
			pvSynced: pvInformer.Informer().HasSynced,
			pvList:   &sync.Map{},
		}
		pvInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
			AddFunc: controller.handlePvAdd,
			UpdateFunc: func(old, new interface{}) {
				newDepl := new.(*corev1.PersistentVolume)
				oldDepl := old.(*corev1.PersistentVolume)
				if newDepl.ResourceVersion == oldDepl.ResourceVersion {
					return
				}
				controller.handlePvUpdate(new)
			},
			DeleteFunc: controller.handlePvDelete,
		})
		CsiControllerInstance = controller
	}
}

//GetPvByDiskID is get pv object from diskid
func (c *CsiController) GetPvByDiskID(diskID string) *corev1.PersistentVolume {
	return c.loadPv(diskID)
}

func validPv(pv *corev1.PersistentVolume) bool {
	if pv.Spec.CSI == nil {
		return false
	}
	if len(pv.Spec.CSI.VolumeHandle) == 0 {
		return false
	}
	if pv.Spec.CSI.Driver != diskDriverName {
		return false
	}
	return true
}

func (c *CsiController) storePv(pv *corev1.PersistentVolume) {
	if validPv(pv) {
		diskID := pv.Spec.CSI.VolumeHandle
		c.pvList.Store(diskID, pv)
	}
}

func (c *CsiController) loadPv(diskID string) *corev1.PersistentVolume {
	value, exist := c.pvList.Load(diskID)
	if exist {
		pv := value.(*corev1.PersistentVolume)
		return pv
	}
	return nil
}

func (c *CsiController) deletePv(diskID string) {
	pv := c.loadPv(diskID)
	if pv != nil {
		c.pvList.Delete(diskID)
	}
}
