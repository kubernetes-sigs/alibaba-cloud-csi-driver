package creator

import (
	corev1 "k8s.io/api/core/v1"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
)


var PVCGlobalMap = map[string]*corev1.PersistentVolumeClaim{}
var PVGlobalMap = map[string]*corev1.PersistentVolume{}

func PvcHandler() {
	if types.GlobalConfigVar.ControllerProvision {
		return
	}
	go pvcInformer()
	pvInformer()
}