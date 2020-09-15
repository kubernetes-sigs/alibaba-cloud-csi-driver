package creator

import (
	"context"
	"fmt"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"os/signal"
	"syscall"
	"time"
)

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

func pvInformer() {
	// set up signals so we handle the first shutdown signal gracefully
	stopCh := SetupSignalHandlerPV()
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	pvInformer := kubeInformerFactory.Core().V1().PersistentVolumes()

	//pvLister := pvInformer.Lister()
	//pvSynced := pvInformer.Informer().HasSynced

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

	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	kubeInformerFactory.Start(stopCh)

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
	log.Infof("Adding Pv: %s", object.GetName())

	pvObj := obj.(*corev1.PersistentVolume)
	pvIdx := GetPvIdx(pvObj)
	PVGlobalMap[pvIdx] = pvObj
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
	log.Infof("Delete Pod: %s", object.GetName())
	pvObj := obj.(*corev1.PersistentVolume)
	pvIdx := GetPvIdx(pvObj)
	delete(PVGlobalMap, pvIdx)
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
	log.Infof("Update Pod: %s", object.GetName())

	pvObj := obj.(*corev1.PersistentVolume)
	pvIdx := GetPvIdx(pvObj)
	if pvObj.Spec.CSI != nil && pvObj.Spec.CSI.Driver == types.LocalDriverName {
		if value, ok := pvObj.Spec.CSI.VolumeAttributes["volumeType"]; ok && value != "LVM" {
			delete(PVCGlobalMap, pvIdx)
			return
		}
		if value, ok := pvObj.Annotations[types.VolumeLifecycleLabel]; ok && value == types.VolumeLifecycleDeleting {
			volumeSpec := pvObj.Annotations[types.VolumeSpecLabel]
			volumeSpecParts := strings.Split(volumeSpec, "/")
			if len(volumeSpecParts) != 2 {
				utilruntime.HandleError(fmt.Errorf("delete lvm with error spec %s", volumeSpec))
				return
			}
			vgName := volumeSpecParts[0]
			volumeID := volumeSpecParts[1]
			_, err := server.RemoveLV(context.Background(), vgName, volumeID)
			if err != nil {
				utilruntime.HandleError(fmt.Errorf("error delete lvm object, %s", volumeSpec))
				return
			}
			labels := map[string]string{}
			labels[types.VolumeLifecycleLabel] = types.VolumeLifecycleDeleted
			UpdatePvWithLabel(pvObj.Name, labels)

			log.Infof("Debug: delete pv: %s", pvObj)
		}
	}

	PVGlobalMap[pvIdx] = pvObj
}

func GetPvIdx(pvObj *corev1.PersistentVolume) string {
	pvIdx := pvObj.Name
	return pvIdx
}
