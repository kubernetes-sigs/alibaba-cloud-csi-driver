package generator

import (
	log "github.com/sirupsen/logrus"
)

// VolumeHandler watch pv/pvc
func VolumeHandler() {
	log.Infof("Local Plugin worker in Worker Provision mode")
	go pvcInformer()
	pvInformer()
}
