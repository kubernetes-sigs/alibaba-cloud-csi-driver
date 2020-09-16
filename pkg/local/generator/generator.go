package generator

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	log "github.com/sirupsen/logrus"
)

func VolumeHandler() {
	if types.GlobalConfigVar.ControllerProvision {
		return
	}
	log.Infof("Local Plugin worker in Worker Provision mode")
	go pvcInformer()
	pvInformer()
}
