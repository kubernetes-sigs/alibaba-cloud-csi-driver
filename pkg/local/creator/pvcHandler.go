package creator

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
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/tools/cache"
	"os/signal"
	"syscall"
	"time"
)
var (
	masterURL  string
	kubeconfig string
)

func pvcInformer() {
	// set up signals so we handle the first shutdown signal gracefully

	stopCh := SetupSignalHandlerPVC()
	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		log.Fatalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		log.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}
	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	pvcInformer := kubeInformerFactory.Core().V1().PersistentVolumeClaims()

	//pvLister := pvInformer.Lister()
	//pvSynced := pvInformer.Informer().HasSynced

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

	// notice that there is no need to run Start methods in a separate goroutine. (i.e. go kubeInformerFactory.Start(stopCh)
	// Start method is non-blocking and runs all registered informers in a dedicated goroutine.
	kubeInformerFactory.Start(stopCh)

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
	log.Infof("Adding Pv: %s", object.GetName())

	pvcObj := obj.(*corev1.PersistentVolumeClaim)
	pvcIdx := GetPvcIdx(pvcObj)
	PVCGlobalMap[pvcIdx] = pvcObj
	
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
	log.Infof("Delete Pod: %s", object.GetName())
	pvcObj := obj.(*corev1.PersistentVolumeClaim)
	pvcIdx := GetPvcIdx(pvcObj)
	delete(PVCGlobalMap, pvcIdx)
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
	log.Infof("Update PVC: %s", object.GetName())

	pvcObj := obj.(*corev1.PersistentVolumeClaim)
	podIdx := GetPvcIdx(pvcObj)
	value, ok := pvcObj.Labels[types.NodeSchedueTag];
	if !ok || types.GlobalConfigVar.NodeID != value {
		PVCGlobalMap[podIdx] = pvcObj
		return
	}

	// pv.csi.alibabacloud.com/volume.lifecycle
	if value, ok := pvcObj.Annotations[types.VolumeLifecycleLabel]; ok && value == types.VolumeLifecycleCreating {
		volumeSpec := pvcObj.Annotations[types.VolumeSpecLabel]
		volumeSpecMap := client.LVMOptions{}
		err := json.Unmarshal([]byte(volumeSpec), &volumeSpecMap)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("error Unmarshal object, invalid spec %s", volumeSpec))
			return
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
		tags := []string{}
		_, err = server.CreateLV(context.Background(), vgName, volumeID, pvSize, 0, tags, striping)
		if err != nil {
			utilruntime.HandleError(fmt.Errorf("error create lvm object, %s", volumeSpec))
			return
		}

		labels := map[string]string{}
		labels[types.VolumeLifecycleLabel] = types.VolumeLifecycleCreated
		UpdatePvcWithLabel(pvcObj.Namespace, pvcObj.Name, labels)
		log.Infof("Debug: create pv: %s", pvcObj)
	} else {
		PVCGlobalMap[podIdx] = pvcObj
	}
}

func GetPvcIdx(pvObj *corev1.PersistentVolumeClaim) string {
	pvIdx := pvObj.Name
	return pvIdx
}

var onlyOneSignalHandlerPVC = make(chan struct{})
var shutdownSignalsPVC = []os.Signal{os.Interrupt, syscall.SIGTERM}

// SetupSignalHandler registered for SIGTERM and SIGINT. A stop channel is returned
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
