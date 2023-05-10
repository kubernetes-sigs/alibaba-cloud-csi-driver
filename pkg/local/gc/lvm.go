package gc

import (
	"context"
	"sync"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/lib"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	"github.com/sirupsen/logrus"
	storagev1 "k8s.io/api/storage/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

const lvmTodoGCTag = "todogc"

type lvmGC struct {
	client *kubernetes.Clientset
	log    *logrus.Entry
}

func newLvmGC(cfg types.GlobalConfig) *lvmGC {
	return &lvmGC{
		client: cfg.KubeClient,
		log:    logrus.WithField("local_gc", "lvm"),
	}
}

func (l *lvmGC) GC(ctx context.Context) {
	// list StorageClasses
	scs, err := l.client.StorageV1().StorageClasses().List(ctx, metav1.ListOptions{})
	if err != nil {
		l.log.Errorf("List StorageClasses: %v", err)
		return
	}
	var wg sync.WaitGroup
	for _, sc := range scs.Items {
		if !isLvmStorageClass(sc) {
			continue
		}
		vg := sc.Parameters["vgName"]
		if vg == "" {
			l.log.Warningf("Skip gc for StorageClass %s because parameters.vgName is empty", sc.Name)
			continue
		}
		wg.Add(1)
		go func(vg string) {
			defer wg.Done()
			l.gcForVG(ctx, vg)
		}(vg)
	}
	wg.Wait()
}

func (l *lvmGC) gcForVG(ctx context.Context, vg string) {
	l.log.Infof("Start gc for vg %s", vg)
	defer l.log.Infof("End gc for vg %s", vg)

	lvs, err := server.ListLV(vg)
	if err != nil {
		l.log.Errorf("List lvs of vg %s: %v", vg, err)
		return
	}

	for _, lv := range lvs {
		// ignore LVs with "protected" tag
		if lvHasTag(lv, server.ProtectedTagName) {
			continue
		}
		exists, err := checkPVExists(ctx, l.client, lv.Name)
		if err != nil {
			l.log.Errorf("Get pv %s: %v", lv.Name, err)
			if err == context.DeadlineExceeded {
				return
			}
			continue
		}
		tagged := lvHasTag(lv, lvmTodoGCTag)
		if exists {
			if tagged {
				out, err := server.RemoveTagLV(ctx, vg, lv.Name, []string{lvmTodoGCTag})
				if err != nil {
					l.log.Errorf("Failed to remove todogc tag on lv %s/%s: %v", vg, lv.Name, err)
				} else {
					l.log.Infof("Removed todogc tag on lv %s/%s: %s", vg, lv.Name, out)
				}
			}
		} else {
			if tagged {
				out, err := server.SafeRemoveLV(ctx, vg, lv.Name)
				if err != nil {
					l.log.Errorf("Failed to remove lv %s/%s: %v", vg, lv.Name, err)
				} else {
					l.log.Infof("Removed lv %s/%s: %s", vg, lv.Name, out)
				}
			} else {
				out, err := server.AddTagLV(ctx, vg, lv.Name, []string{lvmTodoGCTag})
				if err != nil {
					l.log.Errorf("Failed to add todogc tag on lv %s/%s: %v", vg, lv.Name, err)
				} else {
					l.log.Infof("Added todogc tag on lv %s/%s: %s", vg, lv.Name, out)
				}
			}
		}
	}
}

func lvHasTag(lv *lib.LV, tag string) bool {
	for _, t := range lv.Tags {
		if t == tag {
			return true
		}
	}
	return false
}

func isLvmStorageClass(sc storagev1.StorageClass) bool {
	return sc.Provisioner == types.LocalDriverName &&
		sc.Parameters != nil && sc.Parameters["volumeType"] == "LVM"
}
