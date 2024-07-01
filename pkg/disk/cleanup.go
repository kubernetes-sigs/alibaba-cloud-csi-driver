package disk

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

func cleanupOneGlobalMount(path string) error {
	pGlobalMount := filepath.Join(path, "globalmount")
	if err := unix.Unmount(pGlobalMount, 0); err != nil {
		if !errors.Is(err, unix.EINVAL) && !errors.Is(err, unix.ENOENT) { // EINVAL: not a mount point
			return fmt.Errorf("unmount %s failed: %w", pGlobalMount, err)
		}
	}

	if err := unix.Rmdir(pGlobalMount); err != nil {
		if !errors.Is(err, unix.ENOENT) {
			return fmt.Errorf("remove %s failed: %w", pGlobalMount, err)
		}
	}

	pVolData := filepath.Join(path, "vol_data.json")
	if err := unix.Unlink(pVolData); err != nil {
		if !errors.Is(err, unix.ENOENT) {
			return fmt.Errorf("remove %s failed: %w", pVolData, err)
		}
	}

	if err := unix.Rmdir(path); err != nil {
		return fmt.Errorf("remove %s failed: %w", path, err)
	}

	log.Infof("cleanupOldGlobalMount: %s removed", path)
	return nil
}

var cleanupOldGlobalMountTriggered atomic.Bool

func TriggerCleanupOldGlobalMount() {
	triggered := cleanupOldGlobalMountTriggered.Swap(true)
	if !triggered {
		go doCleanupOldGlobalMount()
	}
}

func doCleanupOldGlobalMount() {
	if len(os.Getenv("SKIP_OLD_GLOBALMOUNT_CLEANUP")) > 0 {
		return
	}
	oldGlobalMountPath := filepath.Join(utils.KubeletRootDir, "plugins/kubernetes.io/csi/pv")
	if _, err := os.Stat(oldGlobalMountPath); err != nil {
		return
	}

	dirs, err := os.ReadDir(oldGlobalMountPath)
	if err != nil {
		log.Errorf("cleanupOldGlobalMount: read %s failed: %v", oldGlobalMountPath, err)
		return
	}
	hasError := false
	for _, dir := range dirs {
		if !dir.IsDir() {
			hasError = true
			log.Errorf("cleanupOldGlobalMount: %s is not a directory", dir.Name())
			continue
		}
		err := cleanupOneGlobalMount(filepath.Join(oldGlobalMountPath, dir.Name()))
		if err != nil {
			log.Errorf("cleanupOldGlobalMount: %v", err)
			hasError = true
		}
	}
	if hasError {
		return
	}

	err = unix.Rmdir(oldGlobalMountPath)
	if err != nil {
		log.Errorf("cleanupOldGlobalMount: remove %s failed: %v", oldGlobalMountPath, err)
	}
	log.Infof("cleanupOldGlobalMount: %s removed, finished", oldGlobalMountPath)
}
