package mounter

import (
	"context"
	"errors"
	"fmt"
	"hash/fnv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	watchtools "k8s.io/client-go/tools/watch"
	mountutils "k8s.io/mount-utils"
)

const fuseMountTimeout = time.Minute

const (
	FuseTypeLabelKey          = "csi.alibabacloud.com/fuse-type"
	FuseVolumeIdLabelKey      = "csi.alibabacloud.com/volume-id"
	FuseMountPathHashLabelKey = "csi.alibabacloud.com/mount-path-hash"
)

type FuseMounterType interface {
	name() string
	buildPodSpec(
		source, target, fstype string, options, mountFlags []string, nodeName string, config FuseContainerConfig,
	) (corev1.PodSpec, error)
}

type FuseContainerConfig struct {
	Resources corev1.ResourceRequirements
	Image     string
	Extra     map[string]string
}

func extractFuseContainerConfig(configmap *corev1.ConfigMap, fuseType FuseMounterType) (config FuseContainerConfig) {
	config.Resources.Requests = make(corev1.ResourceList)
	config.Resources.Limits = make(corev1.ResourceList)
	content := configmap.Data["fuse-"+fuseType.name()]
	for _, line := range strings.Split(content, "\n") {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		key, value, _ := strings.Cut(line, "=")
		invalid := false
		switch key {
		case "":
			invalid = true
		case "image":
			config.Image = value
		case "cpu-request", "cpu-limit", "memory-request", "memory-limit":
			quantity, err := resource.ParseQuantity(value)
			if err != nil {
				invalid = true
				break
			}
			resourceName, requireType, _ := strings.Cut(key, "-")
			switch requireType {
			case "request":
				config.Resources.Requests[corev1.ResourceName(resourceName)] = quantity
			case "limit":
				config.Resources.Limits[corev1.ResourceName(resourceName)] = quantity
			}
		default:
			if config.Extra == nil {
				config.Extra = make(map[string]string)
			}
			config.Extra[key] = value
		}
		if invalid {
			logrus.Warnf("ignore invalid configuration line: %q", line)
		}
	}
	return config
}

type ContainerizedFuseMounterFactory struct {
	fuseType            FuseMounterType
	config              FuseContainerConfig
	nodeName, namespace string
	client              kubernetes.Interface
}

func NewContainerizedFuseMounterFactory(
	fuseType FuseMounterType,
	configmap *corev1.ConfigMap,
	client kubernetes.Interface,
	nodeName string,
) *ContainerizedFuseMounterFactory {
	return &ContainerizedFuseMounterFactory{
		fuseType:  fuseType,
		config:    extractFuseContainerConfig(configmap, fuseType),
		nodeName:  nodeName,
		namespace: "kube-system",
		client:    client,
	}
}

func (fac *ContainerizedFuseMounterFactory) NewFuseMounter(volumeId string) *ContainerizedFuseMounter {
	return &ContainerizedFuseMounter{
		volumeId:  volumeId,
		nodeName:  fac.nodeName,
		namespace: fac.namespace,
		client:    fac.client,
		log: logrus.WithFields(logrus.Fields{
			"fuse-type": fac.fuseType.name(),
			"volume-id": volumeId,
		}),
		FuseMounterType: fac.fuseType,
		fuseConfig:      fac.config,
		Interface:       mountutils.New(""),
	}
}

type ContainerizedFuseMounter struct {
	volumeId   string
	nodeName   string
	namespace  string
	client     kubernetes.Interface
	log        *logrus.Entry
	fuseConfig FuseContainerConfig
	FuseMounterType
	mountutils.Interface
}

func (mounter *ContainerizedFuseMounter) Mount(source string, target string, fstype string, options []string) error {
	ctx, cancel := context.WithTimeout(context.Background(), fuseMountTimeout)
	defer cancel()
	err := mounter.launchFusePod(ctx, source, target, fstype, options, nil)
	if err != nil {
		return err
	}

	// check mountpoint
	notMounted, err := mounter.IsLikelyNotMountPoint(target)
	if err != nil {
		return err
	}
	if notMounted {
		return errors.New("fuse pod launched but target still not mounted")
	}
	return nil
}

func (mounter *ContainerizedFuseMounter) Unmount(target string) error {
	unmountErr := mounter.Interface.Unmount(target)
	if unmountErr != nil {
		mounter.log.Errorf("umount %s: %v", target, unmountErr)
	} else {
		mounter.log.Infof("successfully unmount %s", target)
	}

	_, listOptions := mounter.labelsAndListOptionsFor(target)
	deletePodErr := mounter.client.CoreV1().Pods(mounter.namespace).DeleteCollection(context.TODO(), metav1.DeleteOptions{}, listOptions)
	if deletePodErr != nil {
		mounter.log.Errorf("failed to delete fuse pod: %v", unmountErr)
	} else {
		mounter.log.Infof("deleted fuse pod for %s", target)
	}

	// Only return nil when unmounted successfully and fuse pods are deleted
	if unmountErr != nil {
		return unmountErr
	}
	if deletePodErr != nil {
		return deletePodErr
	}
	return nil
}

func (mounter *ContainerizedFuseMounter) labelsAndListOptionsFor(target string) (map[string]string, metav1.ListOptions) {
	labels := map[string]string{
		FuseTypeLabelKey:          mounter.name(),
		FuseVolumeIdLabelKey:      mounter.volumeId,
		FuseMountPathHashLabelKey: computeMountPathHash(target),
	}
	listOptions := metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("spec.nodeName", mounter.nodeName).String(),
		LabelSelector: metav1.FormatLabelSelector(&metav1.LabelSelector{MatchLabels: labels}),
	}
	return labels, listOptions
}

func (mounter *ContainerizedFuseMounter) launchFusePod(ctx context.Context, source, target, fstype string, options, mountFlags []string) error {
	podClient := mounter.client.CoreV1().Pods(mounter.namespace)
	labels, listOptions := mounter.labelsAndListOptionsFor(target)
	podList, err := podClient.List(ctx, listOptions)
	if err != nil {
		return err
	}
	var pods []corev1.Pod
	for _, pod := range podList.Items {
		if pod.DeletionTimestamp.IsZero() {
			if isFusePodReady(&pod) {
				mounter.log.Infof("%s is already mounted by pod %s", target, pod.Name)
				return nil
			}
			pods = append(pods, pod)
		}
	}

	var fusePod *corev1.Pod
	if len(pods) == 0 {
		// create fuse pod for target
		var rawPod corev1.Pod
		rawPod.GenerateName = fmt.Sprintf("csi-%s-%s-", mounter.name(), mounter.volumeId)
		rawPod.Labels = labels
		rawPod.Spec, err = mounter.buildPodSpec(source, target, fstype, options, mountFlags, mounter.nodeName, mounter.fuseConfig)
		if err != nil {
			return err
		}
		mounter.log.Infof("creating fuse pod for %s", target)
		createdPod, err := podClient.Create(ctx, &rawPod, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		mounter.log.Infof("created fuse pod %s for %s", createdPod.Name, target)
		fusePod = createdPod
	} else {
		if len(pods) > 1 {
			mounter.log.Warnf("%d duplicated fuse pods for %s", len(pods), target)
			// TODO: clean duplicated fuse pods
		}
		mounter.log.Infof("found existed fuse pod %s for %s", pods[0].Name, target)
		fusePod = &pods[0]
	}

	mounter.log.Infof("wait util pod %s is ready", fusePod.Name)
	fieldSelector := fields.OneTermEqualSelector("metadata.name", fusePod.Name).String()
	lw := &cache.ListWatch{
		ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
			options.FieldSelector = fieldSelector
			return podClient.List(ctx, options)
		},
		WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
			options.FieldSelector = fieldSelector
			return podClient.Watch(ctx, options)
		},
	}
	event, err := watchtools.Until(ctx, fusePod.ResourceVersion, lw, func(event watch.Event) (bool, error) {
		if event.Type == watch.Deleted {
			return true, fmt.Errorf("fuse pod %s was deleted", fusePod.Name)
		}
		pod, ok := event.Object.(*corev1.Pod)
		return ok && isFusePodReady(pod), nil
	})
	if err != nil {
		return err
	}
	if pod, ok := event.Object.(*corev1.Pod); !(ok && isFusePodReady(pod)) {
		return fmt.Errorf("failed to wait for pod %s to be ready", fusePod.Name)
	}
	mounter.log.Infof("fuse pod %s is ready", fusePod.Name)

	return nil
}

func isFusePodReady(pod *corev1.Pod) bool {
	if pod.Status.Phase != corev1.PodRunning {
		return false
	}
	for _, cond := range pod.Status.Conditions {
		if cond.Type == corev1.PodReady && cond.Status == corev1.ConditionTrue {
			return true
		}
	}
	return false
}

func computeMountPathHash(target string) string {
	hasher := fnv.New32a()
	hasher.Write([]byte(target))
	return rand.SafeEncodeString(fmt.Sprint(hasher.Sum32()))
}
