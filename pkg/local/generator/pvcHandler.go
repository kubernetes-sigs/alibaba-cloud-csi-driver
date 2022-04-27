package generator

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/client"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kubeinformers "k8s.io/client-go/informers"
	"os"
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"os/signal"
	"syscall"
	"time"
)

// DesiredStateOfWordPvc the desired pvc state
type DesiredStateOfWordPvc struct {
	DesiredPvcMap map[string]*corev1.PersistentVolumeClaim
	sync.RWMutex
}

// Add object
func (dsw *DesiredStateOfWordPvc) Add(pvc *corev1.PersistentVolumeClaim) {
	dsw.Lock()
	defer dsw.Unlock()
	dsw.DesiredPvcMap[pvc.Name] = pvc
}

// Remove object
func (dsw *DesiredStateOfWordPvc) Remove(pvc *corev1.PersistentVolumeClaim) {
	dsw.Lock()
	defer dsw.Unlock()
	delete(dsw.DesiredPvcMap, pvc.Name)
}

// ActualStateOfWordPvc the actual pvc state
type ActualStateOfWordPvc struct {
	ActualPvcMap map[string]*corev1.PersistentVolumeClaim
	sync.RWMutex
}

// Add object
func (dsw *ActualStateOfWordPvc) Add(pvc *corev1.PersistentVolumeClaim) {
	dsw.Lock()
	defer dsw.Unlock()
	dsw.ActualPvcMap[pvc.Name] = pvc
}

// Remove object
func (dsw *ActualStateOfWordPvc) Remove(pvc *corev1.PersistentVolumeClaim) {
	dsw.Lock()
	defer dsw.Unlock()
	delete(dsw.ActualPvcMap, pvc.Name)
}

// DesiredStateOfPvc record desired pvc
var DesiredStateOfPvc = DesiredStateOfWordPvc{}

// ActualStateOfPvc record actual pvc
var ActualStateOfPvc = ActualStateOfWordPvc{}

func pvcInformer() {
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := SetupSignalHandlerPVC()
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(types.GlobalConfigVar.KubeClient, time.Second*30)
	pvcInformer := kubeInformerFactory.Core().V1().PersistentVolumeClaims()

	pvcInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: handlePvcAdd,
		UpdateFunc: func(old, new interface{}) {
			newDepl := new.(*corev1.PersistentVolumeClaim)
			oldDepl := old.(*corev1.PersistentVolumeClaim)
			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
				// Periodic resync will send update events for all known Deployments.
				// Two different versions of the same Deployment will always have different RVs.
				return
			}
			handlePvcUpdate(new)
		},
		DeleteFunc: handlePvcDelete,
	})

	// init pvc maps
	DesiredStateOfPvc.DesiredPvcMap = map[string]*corev1.PersistentVolumeClaim{}
	ActualStateOfPvc.ActualPvcMap = map[string]*corev1.PersistentVolumeClaim{}

	// reconcile pvc state while process failed
	go pvcReconcile()

	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	kubeInformerFactory.Start(stopCh)
}

// process pvc object
func pvcReconcile() {
	for {
		for _, pvc := range DesiredStateOfPvc.DesiredPvcMap {
			if pvc.Annotations[types.VolumeLifecycleLabel] == types.VolumeLifecycleCreating {
				if err := processPvc(pvc); err != nil {
					log.Errorf("pvcReconcile: Process PVC error %s", err.Error())
				}
			}
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func handlePvcAdd(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		log.Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	pvcObj := obj.(*corev1.PersistentVolumeClaim)
	// if pvc not desired pvcObj, just return
	if !isPvcExpected(pvcObj) {
		return
	}
	log.Infof("Adding Pvc: %s", object.GetName())

	DesiredStateOfPvc.Add(pvcObj)
	//if pvcObj.Spec.VolumeName != "" {
	//	ActualStateOfPvc.Add(pvcObj)
	//}
}

func handlePvcDelete(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		log.Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}
	pvcObj := obj.(*corev1.PersistentVolumeClaim)
	if !isPvcExpected(pvcObj) {
		return
	}
	log.Infof("Delete Pvc: %s", object.GetName())

	DesiredStateOfPvc.Remove(pvcObj)
	//	ActualStateOfPvc.Remove(pvcObj)
}

func handlePvcUpdate(obj interface{}) {
	var object metav1.Object
	var ok bool
	if object, ok = obj.(metav1.Object); !ok {
		tombstone, ok := obj.(cache.DeletedFinalStateUnknown)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object, invalid type"))
			return
		}
		object, ok = tombstone.Obj.(metav1.Object)
		if !ok {
			utilruntime.HandleError(fmt.Errorf("error decoding object tombstone, invalid type"))
			return
		}
		log.Infof("Recovered deleted object '%s' from tombstone", object.GetName())
	}

	pvcObj := obj.(*corev1.PersistentVolumeClaim)
	if !isPvcExpected(pvcObj) {
		return
	}
	log.Infof("Update PVC: %s", object.GetName())

	// pv.csi.alibabacloud.com/volume.lifecycle
	if value, ok := pvcObj.Annotations[types.VolumeLifecycleLabel]; ok && value == types.VolumeLifecycleCreating {
		DesiredStateOfPvc.Add(pvcObj)
		if err := processPvc(pvcObj); err != nil {
			log.Errorf("handlePvcUpdate: processPvc with error: %v", err)
			utilruntime.HandleError(err)
			return
		}
		log.Infof("handlePvUpdate: lvm volume created %s/%s", pvcObj.Namespace, pvcObj.Name)
	}
}

func isPvcExpected(pvc *corev1.PersistentVolumeClaim) bool {
	value, ok := pvc.Annotations[types.NodeSchedueTag]
	if !ok || types.GlobalConfigVar.NodeID != value {
		return false
	}
	if _, ok := pvc.Annotations[types.VolumeLifecycleLabel]; !ok {
		return false
	}
	return true
}

func processPvc(pvcObj *corev1.PersistentVolumeClaim) error {
	volumeSpec := pvcObj.Annotations[types.VolumeSpecLabel]
	volumeSpecMap := client.LVMOptions{}
	err := json.Unmarshal([]byte(volumeSpec), &volumeSpecMap)
	if err != nil {
		return fmt.Errorf("error Unmarshal object, invalid spec %s", volumeSpec)
	}
	vgName := ""
	volumeID := ""
	striping := false
	var pvSize uint64
	if volumeSpecMap.VolumeGroup != "" {
		vgName = volumeSpecMap.VolumeGroup
	}
	if volumeSpecMap.Name != "" {
		volumeID = volumeSpecMap.Name
	}
	if volumeSpecMap.Striping {
		striping = true
	}
	if volumeSpecMap.Size != 0 {
		pvSize = volumeSpecMap.Size
	}
	lvmName := vgName + "/" + volumeID
	rsp, err := server.ListLV(lvmName)
	if err != nil {
		log.Errorf("processPvc: Get Lvm with error: %s", err.Error())
		return err
	}
	if len(rsp) == 0 {
		tags := []string{}
		_, err = server.CreateLV(context.Background(), vgName, volumeID, pvSize, 0, tags, striping)
		if err != nil {
			return fmt.Errorf("processPvc: error create lvm object, %s, %v", volumeSpec, err)
		}
	}

	labels := map[string]string{}
	labels[types.VolumeLifecycleLabel] = types.VolumeLifecycleCreated
	if err := UpdatePvcWithAnnotations(context.Background(), pvcObj.Namespace, pvcObj.Name, labels); err != nil {
		return fmt.Errorf("update pvc object error, %s", err.Error())
	}
	return nil
}

var onlyOneSignalHandlerPVC = make(chan struct{})
var shutdownSignalsPVC = []os.Signal{os.Interrupt, syscall.SIGTERM}

// SetupSignalHandlerPVC registered for SIGTERM and SIGINT. A stop channel is returned
// which is closed on one of these signals. If a second signal is caught, the program
// is terminated with exit code 1.
func SetupSignalHandlerPVC() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandlerPVC) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignalsPVC...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
