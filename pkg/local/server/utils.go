package server

import (
	"context"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"path/filepath"
	"regexp"
	"strings"
)

func checkStringAlpha(inputs []string) bool {
	log.Infof("Debug: checkStringAlpha %v", inputs)
	for _, input := range inputs {
		if matched, err := regexp.MatchString("^[A-Za-z0-9=._~/-]*$", input); err != nil || !matched {
			return false
		}
	}
	return true
}

func checkQuotaPath(path string) error {
	log.Infof("Debug: checkQuotaPath %v", path)
	pvName := filepath.Base(path)
	_, err := KubeClient.CoreV1().PersistentVolumes().Get(context.Background(), pvName, metav1.GetOptions{})
	if err != nil {
		log.Errorf("checkQuotaPath %s cannot find volume, error: %s", path, err.Error())
		return err
	}
	return nil
}

func isHostFileExist(path string) bool {
	args := []string{NsenterCmd, "stat", path}
	cmd := strings.Join(args, " ")
	out, err := utils.Run(cmd)
	if err != nil && strings.Contains(out, "No such file or directory") {
		return false
	}

	return true
}
