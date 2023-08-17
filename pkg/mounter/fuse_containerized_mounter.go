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

const fuseMountTimeout = time.Second * 30

const (
	FuseTypeLabelKey          = "csi.alibabacloud.com/fuse-type"
	FuseVolumeIdLabelKey      = "csi.alibabacloud.com/volume-id"
	FuseMountPathHashLabelKey = "csi.alibabacloud.com/mount-path-hash"
	FuseMountPathAnnoKey      = "csi.alibabacloud.com/mount-path"
)

type FuseMounterType interface {
	name() string
	buildPodSpec(source, target, fstype string, options, mountFlags []string, nodeName string) (corev1.PodSpec, error)
}

type FuseContainerConfig struct {
	Resources corev1.ResourceRequirements
	Image     string
	Extra     map[string]string
}

func extractFuseContainerConfig(configmap *corev1.ConfigMap, name string) (config FuseContainerConfig) {
	config.Resources.Requests = make(corev1.ResourceList)
	config.Resources.Limits = make(corev1.ResourceList)
	content := configmap.Data["fuse-"+name]
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
	nodeName, namespace string
	client              kubernetes.Interface
}

func NewContainerizedFuseMounterFactory(
	fuseType FuseMounterType,
	client kubernetes.Interface,
	nodeName string,
) *ContainerizedFuseMounterFactory {
	return &ContainerizedFuseMounterFactory{
		fuseType:  fuseType,
		nodeName:  nodeName,
		namespace: "kube-system",
		client:    client,
	}
}

func (fac *ContainerizedFuseMounterFactory) NewFuseMounter(ctx context.Context, volumeId string) *ContainerizedFuseMounter {
	return &ContainerizedFuseMounter{
		ctx:       ctx,
		volumeId:  volumeId,
		nodeName:  fac.nodeName,
		namespace: fac.namespace,
		client:    fac.client,
		log: logrus.WithFields(logrus.Fields{
			"fuse-type": fac.fuseType.name(),
			"volume-id": volumeId,
		}),
		FuseMounterType: fac.fuseType,
		Interface:       mountutils.New(""),
	}
}

type ContainerizedFuseMounter struct {
	ctx       context.Context
	volumeId  string
	nodeName  string
	namespace string
	client    kubernetes.Interface
	log       *logrus.Entry
	FuseMounterType
	mountutils.Interface
}

func (mounter *ContainerizedFuseMounter) Mount(source string, target string, fstype string, options []string) error {
	for _, option := range options {
		if option == "bind" {
			return mounter.Interface.Mount(source, target, fstype, options)
		}
	}

	ctx, cancel := context.WithTimeout(mounter.ctx, fuseMountTimeout)
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
	_, listOptions := mounter.labelsAndListOptionsFor(target)
	err := mounter.client.CoreV1().Pods(mounter.namespace).DeleteCollection(mounter.ctx, metav1.DeleteOptions{}, listOptions)
	if err != nil {
		return fmt.Errorf("delete fuse pod: %w", err)
	}
	mounter.log.Infof("deleted fuse pod for %s", target)
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

	ready := false
	var startingPods []corev1.Pod
	for _, pod := range podList.Items {
		// compare target path to avoid hash conflict
		if !pod.DeletionTimestamp.IsZero() || pod.Annotations[FuseMountPathAnnoKey] != target {
			continue
		}
		switch pod.Status.Phase {
		case corev1.PodSucceeded, corev1.PodFailed:
			// do not immediately delete the pod that has exited,
			// as we may need to view its status or logs to troubleshoot.
			mounter.log.Warnf("exited fuse pod %s, will be cleaned when unmount", pod.Name)
		case corev1.PodRunning:
			if isFusePodReady(&pod) {
				ready = true
				mounter.log.Infof("%s already mounted by pod %s", target, pod.Name)
			} else {
				startingPods = append(startingPods, pod)
			}
		default:
			startingPods = append(startingPods, pod)
		}
	}
	if ready {
		return nil
	}

	var fusePod *corev1.Pod
	if len(startingPods) == 0 {
		// create fuse pod for target
		var rawPod corev1.Pod
		rawPod.GenerateName = fmt.Sprintf("csi-fuse-%s-", mounter.name())
		rawPod.Labels = labels
		rawPod.Annotations = map[string]string{FuseMountPathAnnoKey: target}
		rawPod.Spec, err = mounter.buildPodSpec(source, target, fstype, options, mountFlags, mounter.nodeName)
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
		if len(startingPods) > 1 {
			mounter.log.Warnf("%d duplicated fuse pods for %s", len(startingPods), target)
		}
		mounter.log.Infof("found existed fuse pod %s for %s", startingPods[0].Name, target)
		fusePod = &startingPods[0]
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
	_, err = watchtools.Until(ctx, fusePod.ResourceVersion, lw, func(event watch.Event) (bool, error) {
		if event.Type == watch.Deleted {
			return false, fmt.Errorf("fuse pod %s was deleted", fusePod.Name)
		}
		pod, ok := event.Object.(*corev1.Pod)
		if !ok {
			return false, errors.New("failed to cast event Object to Pod")
		}
		if pod.Status.Phase == corev1.PodSucceeded || pod.Status.Phase == corev1.PodFailed {
			return false, fmt.Errorf("fuse pod %s exited", pod.Name)
		}
		return pod.Status.Phase == corev1.PodRunning && isFusePodReady(pod), nil
	})
	if err != nil {
		return err
	}
	mounter.log.Infof("fuse pod %s is ready", fusePod.Name)
	return nil
}

func isFusePodReady(pod *corev1.Pod) bool {
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
