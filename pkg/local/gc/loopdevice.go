package gc

import (
	"context"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/manager"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	"github.com/sirupsen/logrus"
	"k8s.io/client-go/kubernetes"
)

// ".todogc" is a file for recording the list of files that need to be garbage collected.
const ldTodoGCFile = ".todogc"

type loopDeviceGC struct {
	ld     manager.LoopDevice
	dir    string
	client *kubernetes.Clientset
	log    *logrus.Entry
}

func newLoopDeviceGC(cfg types.GlobalConfig) *loopDeviceGC {
	dir := cfg.LocalSparseFileDir
	ld := manager.NewLoopDevice(dir, cfg.LocalSparseFileTempSize)
	return &loopDeviceGC{
		ld:     ld,
		dir:    dir,
		client: cfg.KubeClient,
		log:    logrus.WithField("local_gc", "loopdevice"),
	}
}

func (l *loopDeviceGC) GC(ctx context.Context) {
	l.log.Info("Start loopdevice gc")
	defer l.log.Info("End loopdevice gc")

	todolist, err := l.loadTodoGCFile()
	if err != nil {
		l.log.Errorf("Load .todogc: %v", err)
		return
	}
	var nextTodolist []string

	err = filepath.WalkDir(l.dir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		filename := d.Name()
		// ignore files not end with .img
		ext := filepath.Ext(filename)
		if ext != ".img" {
			return nil
		}
		// ignore template file
		if filename == manager.LOCAL_SPARSE_TEMPLATE_NAME {
			return nil
		}

		pvName := filename[:len(filename)-len(ext)]
		exists, err := checkPVExists(ctx, l.client, pvName)
		if err != nil {
			l.log.Errorf("Get pv %s: %v", pvName, err)
			return nil
		}

		_, todogc := todolist[pvName]
		if !exists {
			if todogc {
				out, err := l.ld.SafeDeleteLoopDevice(path)
				if err != nil {
					nextTodolist = append(nextTodolist, pvName)
					l.log.Errorf("Failed to remove loopdevice %s: %v", pvName, err)
				} else {
					l.log.Infof("Removed loopdevice %s: %s", pvName, out)
				}
			} else {
				nextTodolist = append(nextTodolist, pvName)
			}
		}

		return nil
	})
	if err != nil {
		l.log.Errorf("Walk sparsefile dir: %v", err)
	}
	if err := l.saveTodoGCFile(nextTodolist); err != nil {
		l.log.Errorf("Update .todogc: %v", err)
	}
	l.log.Infof("Updated .todogc: %s", strings.Join(nextTodolist, ","))
}

func (l *loopDeviceGC) loadTodoGCFile() (map[string]struct{}, error) {
	data, err := os.ReadFile(filepath.Join(l.dir, ldTodoGCFile))
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	todolist := make(map[string]struct{}, len(lines))
	for _, line := range lines {
		name := strings.TrimSpace(line)
		if name != "" {
			todolist[name] = struct{}{}
		}
	}
	return todolist, nil
}

func (l *loopDeviceGC) saveTodoGCFile(todolist []string) error {
	return os.WriteFile(filepath.Join(l.dir, ldTodoGCFile), []byte(strings.Join(todolist, "\n")), 0644)
}
