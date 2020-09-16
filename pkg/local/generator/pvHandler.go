package generator

import (
	"context"
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kubeinformers "k8s.io/client-go/informers"
	"os"
	"strings"
	"sync"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"os/signal"
	"syscall"
	"time"
)

type DesiredStateOfWordPv struct {
	DesiredPvMap map[string]*corev1.PersistentVolume
	sync.RWMutex
}

// Add add pv
func (dsw *DesiredStateOfWordPv) Add(pv *corev1.PersistentVolume) {
	dsw.Lock()
	defer dsw.Unlock()
	dsw.DesiredPvMap[pv.Name] = pv
}

// Remove remove pv
func (dsw *DesiredStateOfWordPv) Remove(pv *corev1.PersistentVolume) {
	dsw.Lock()
	defer dsw.Unlock()
	delete(dsw.DesiredPvMap, pv.Name)
}

// ActualStateOfWordPv save the actual state of volume in node.
type ActualStateOfWordPv struct {
	ActualPvMap map[string]*corev1.PersistentVolume
	sync.RWMutex
}

func (dsw *ActualStateOfWordPv) Add(pv *corev1.PersistentVolume) {
	dsw.Lock()
	defer dsw.Unlock()
	dsw.ActualPvMap[pv.Name] = pv
}

func (dsw *ActualStateOfWordPv) Remove(pv *corev1.PersistentVolume) {
	dsw.Lock()
	defer dsw.Unlock()
	delete(dsw.ActualPvMap, pv.Name)
}

var DesiredStateOfPv = DesiredStateOfWordPv{}
var ActualStateOfPv = ActualStateOfWordPv{}

func isPvExpected(pv *corev1.PersistentVolume) bool {
	if pv.Spec.CSI == nil {
		return false
	}
	if pv.Spec.CSI.Driver != types.LocalDriverName {
		return false
	}
	if pv.Spec.NodeAffinity == nil {
		return false
	}
	if pv.Spec.NodeAffinity.Required == nil {
		return false
	}
	if len(pv.Spec.NodeAffinity.Required.NodeSelectorTerms) != 1 {
		return false
	}
	if len(pv.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions) != 1 {
		return false
	}
	nodes := pv.Spec.NodeAffinity.Required.NodeSelectorTerms[0].MatchExpressions[0].Values
	if len(nodes) == 0 {
		return false
	}
	if types.GlobalConfigVar.NodeID != nodes[0] {
		return false
	}
	if value, ok := pv.Spec.CSI.VolumeAttributes["volumeType"]; ok && value != "LVM" {
		return false
	}
	return true
}

func pvInformer() {
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := SetupSignalHandlerPV()
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(types.GlobalConfigVar.KubeClient, time.Second*30)
	pvInformer := kubeInformerFactory.Core().V1().PersistentVolumes()

	pvInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: handlePvAdd,
		UpdateFunc: func(old, new interface{}) {
			newDepl := new.(*corev1.PersistentVolume)
			oldDepl := old.(*corev1.PersistentVolume)
			if newDepl.ResourceVersion == oldDepl.ResourceVersion {
				// Periodic resync will send update events for all known Deployments.
				// Two different versions of the same Deployment will always have different RVs.
				return
			}
			handlePvUpdate(new)
		},
		DeleteFunc: handlePvDelete,
	})
	DesiredStateOfPv.DesiredPvMap = map[string]*corev1.PersistentVolume{}
	ActualStateOfPv.ActualPvMap = map[string]*corev1.PersistentVolume{}

	go pvReconcile()
	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	kubeInformerFactory.Start(stopCh)

}

func pvReconcile() {
	for {
		for _, pv := range DesiredStateOfPv.DesiredPvMap {
			if pv.Annotations[types.VolumeLifecycleLabel] == types.VolumeLifecycleDeleting {
				//if _, ok := ActualStateOfPv.ActualPvMap[pv.Name]; !ok {
				//	ActualStateOfPv.Add(pv)
				//}
				if err := processPv(pv); err != nil {
					log.Errorf("pvReconcile: Process PV in reconcile error %s", err.Error())
				}
			}
		}
		time.Sleep(time.Duration(2) * time.Second)
	}
}

func handlePvAdd(obj interface{}) {
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

	pvObj := obj.(*corev1.PersistentVolume)
	if !isPvExpected(pvObj) {
		return
	}
	log.Infof("Adding Pv: %s", object.GetName())

	DesiredStateOfPv.Add(pvObj)
	//ActualStateOfPv.Add(pvObj)
}

func handlePvDelete(obj interface{}) {
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
	pvObj := obj.(*corev1.PersistentVolume)
	if !isPvExpected(pvObj) {
		return
	}

	log.Infof("Delete Pv: %s", object.GetName())
	DesiredStateOfPv.Remove(pvObj)
	//ActualStateOfPv.Remove(pvObj)
}

func handlePvUpdate(obj interface{}) {
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

	pvObj := obj.(*corev1.PersistentVolume)
	if !isPvExpected(pvObj) {
		return
	}
	log.Infof("Update Pv: %s", object.GetName())

	if value, ok := pvObj.Annotations[types.VolumeLifecycleLabel]; ok && value == types.VolumeLifecycleDeleting {
		DesiredStateOfPv.Add(pvObj)
		if err := processPv(pvObj); err != nil {
			return
		}
		//ActualStateOfPv.Add(pvObj)
		log.Infof("Debug: delete pv: %s", pvObj)
	}
}

func processPv(pvObj *corev1.PersistentVolume) error {
	volumeSpec := pvObj.Annotations[types.VolumeSpecLabel]
	volumeSpecParts := strings.Split(volumeSpec, "/")
	if len(volumeSpecParts) != 2 {
		utilruntime.HandleError(fmt.Errorf("delete lvm with error spec %s", volumeSpec))
		return fmt.Errorf("delete lvm with error spec")
	}

	vgName := volumeSpecParts[0]
	volumeID := volumeSpecParts[1]
	_, err := server.RemoveLV(context.Background(), vgName, volumeID)
	if err != nil {
		utilruntime.HandleError(fmt.Errorf("error delete lvm object, %s", volumeSpec))
		return err
	}

	labels := map[string]string{}
	labels[types.VolumeLifecycleLabel] = types.VolumeLifecycleDeleted
	if err := UpdatePvWithLabel(context.Background(), pvObj.Name, labels); err != nil {
		return err
	}
	return nil
}

var onlyOneSignalHandlerPV = make(chan struct{})
var shutdownSignalsPV = []os.Signal{os.Interrupt, syscall.SIGTERM}

func SetupSignalHandlerPV() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandlerPV) // panics when called twice

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignalsPV...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop
}
