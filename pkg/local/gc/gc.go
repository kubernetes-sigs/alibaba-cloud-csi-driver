package gc

import (
	"context"
	"math/rand"
	"os"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/local/types"
	log "github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/kubernetes"
)

const (
	LOCAL_GC_LVM_DURATION        = "LOCAL_GC_LVM_DURATION"
	LOCAL_GC_LOOPDEVICE_DURATION = "LOCAL_GC_LOOPDEVICE_DURATION"
	MinGCDuration                = time.Hour
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func Start(cfg types.GlobalConfig) {
	// start lvm gc
	start(newLvmGC(cfg), LOCAL_GC_LVM_DURATION)

	// start loopdevice gc
	start(newLoopDeviceGC(cfg), LOCAL_GC_LOOPDEVICE_DURATION)
}

type garbageCollector interface {
	GC(ctx context.Context)
}

// The start function initializes a garbage collector based on the environment variable settings.
// If the durationEnvName variable is not empty and has a non-zero value, the garbageCollector.GC()
// function will be called at regular intervals.
// Before the first call, the function will sleep for a random duration.
func start(c garbageCollector, durationEnvName string) {
	envVal := os.Getenv(durationEnvName)
	if envVal == "" {
		return
	}
	duration, err := time.ParseDuration(envVal)
	if err != nil {
		log.Errorf("Skip initializing gc: invalid env %s: %v", durationEnvName, err)
		return
	}
	if duration <= 0 {
		log.Errorf("Skip initializing gc: env %s is non-positive", durationEnvName)
		return
	}
	if duration < MinGCDuration {
		log.Warningf("%s[%v] is too small, will use %v", durationEnvName, duration, MinGCDuration)
		duration = MinGCDuration
	}
	log.Infof("Initialized gc for env %s setting: every %v", durationEnvName, duration)

	go func() {
		// sleep for a random duration
		sleepDuration := time.Duration(rand.Int63n(int64(duration)))
		log.Infof("[%s] Sleep %v before the first time gc", durationEnvName, sleepDuration)
		time.Sleep(sleepDuration)

		wait.Forever(func() {
			ctx, cancel := context.WithTimeout(context.Background(), duration)
			defer cancel()
			c.GC(ctx)
		}, duration)
	}()

}

func checkPVExists(ctx context.Context, client *kubernetes.Clientset, pvName string) (bool, error) {
	_, err := client.CoreV1().PersistentVolumes().Get(ctx, pvName, metav1.GetOptions{})
	if err == nil {
		return true, nil
	}
	if errors.IsNotFound(err) {
		return false, nil
	}
	return false, err
}
