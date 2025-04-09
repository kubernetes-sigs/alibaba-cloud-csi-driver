package utils

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	watchtools "k8s.io/client-go/tools/watch"
	"k8s.io/klog/v2"
)

const (
	fusePodManagerTimeout = time.Second * 30
	fuseServieAccountName = "csi-fuse-ossfs"
	// deprecated
	LegacyFusePodNamespace = "kube-system"
)

const (
	FuseTypeLabelKey          = "csi.alibabacloud.com/fuse-type"
	FuseVolumeIdLabelKey      = "csi.alibabacloud.com/volume-id"
	FuseMountPathHashLabelKey = "csi.alibabacloud.com/mount-path-hash"
	FuseMountPathAnnoKey      = "csi.alibabacloud.com/mount-path"
	FuseSafeToEvictAnnoKey    = "cluster-autoscaler.kubernetes.io/safe-to-evict"
	ACKDrainLabelKey          = "alibabacloud.com/drain-pod"
)

type AuthConfig struct {
	AuthType string
	// for RRSA
	RrsaConfig *RrsaConfig
	// for csi-secret-store
	SecretProviderClassName string
	// for AK/SK
	Secrets map[string]string
	// for Token from Secret
	SecretRef string
	// for STS(ECS worker role)
	RoleName string
}

type RrsaConfig struct {
	OidcProviderArn    string
	RoleArn            string
	ServiceAccountName string
	AssumeRoleArn      string
}

const (
	AuthTypeSTS    = "sts"
	AuthTypeRRSA   = "rrsa"
	AuthTypeCSS    = "csi-secret-store"
	AuthTypePublic = "public"
)

const (
	DebugLevelFatal = "fatal"
	DebugLevelError = "error"
	DebugLevelWarn  = "warn"
	DebugLevelInfo  = "info"
	DebugLevelDebug = "debug"
)

type FusePodContext struct {
	context.Context
	Namespace  string
	NodeName   string
	VolumeId   string
	FuseType   string
	AuthConfig *AuthConfig
}

type FuseMounterType interface {
	Name() string
	PodTemplateSpec(c *FusePodContext, target string) (*corev1.PodTemplateSpec, error)
	AddDefaultMountOptions(options []string) []string
	CheckAuthConfig(c *AuthConfig) error
}

type FuseContainerConfig struct {
	Resources   corev1.ResourceRequirements
	Image       string
	ImageTag    string
	Dbglevel    string
	Annotations map[string]string
	Labels      map[string]string
	Extra       map[string]string
}

func ExtractFuseContainerConfig(configmap *corev1.ConfigMap, name string) (config FuseContainerConfig) {
	config.Resources.Requests = make(corev1.ResourceList)
	config.Resources.Limits = make(corev1.ResourceList)

	if configmap == nil {
		return
	}
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
		case "dbglevel", "log_level":
			switch value {
			case DebugLevelFatal, DebugLevelError, DebugLevelWarn, DebugLevelInfo, DebugLevelDebug:
				config.Dbglevel = value
			default:
				invalid = true
			}
		case "image":
			klog.Warning("'image' config in configmap no longer supported")
		case "image-tag":
			config.ImageTag = value
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
		case "annotations":
			annotations := make(map[string]string)
			err := json.Unmarshal([]byte(value), &annotations)
			if err != nil {
				invalid = true
				break
			}
			err = ValidateAnnotations(annotations)
			if err != nil {
				invalid = true
				break
			}
			config.Annotations = annotations
		case "labels":
			labels := make(map[string]string)
			err := json.Unmarshal([]byte(value), &labels)
			if err != nil {
				invalid = true
				break
			}
			err = ValidateLabels(labels)
			if err != nil {
				invalid = true
				break
			}
			config.Labels = labels
		default:
			if config.Extra == nil {
				config.Extra = make(map[string]string)
			}
			config.Extra[key] = value
		}
		if invalid {
			klog.Warningf("ignore invalid configuration line: %q", line)
		}
	}
	return
}

type FusePodManager struct {
	client kubernetes.Interface
	FuseMounterType
}

func NewFusePodManager(fuseType FuseMounterType, client kubernetes.Interface) *FusePodManager {
	return &FusePodManager{
		client:          client,
		FuseMounterType: fuseType,
	}
}

func (fpm *FusePodManager) labelsAndListOptionsFor(c *FusePodContext, target string) (map[string]string, metav1.ListOptions) {
	labels := map[string]string{
		FuseVolumeIdLabelKey: c.VolumeId,
	}
	// ControllerUnPublish cannot get fuseType info,
	// so FuseTypeLabelKey cannot used as a label for Delete
	if c.FuseType != "" {
		labels[FuseTypeLabelKey] = c.FuseType
	}
	if target != "" {
		labels[FuseMountPathHashLabelKey] = computeMountPathHash(target)
	}
	listOptions := metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("spec.nodeName", c.NodeName).String(),
		LabelSelector: metav1.FormatLabelSelector(&metav1.LabelSelector{MatchLabels: labels}),
	}
	return labels, listOptions
}

func (fpm *FusePodManager) Create(c *FusePodContext, target string, atomic bool) (*corev1.Pod, error) {
	ctx, cancel := context.WithTimeout(c, fusePodManagerTimeout)
	defer cancel()

	logger := klog.Background().WithValues(
		"volumeId", c.VolumeId,
		"nodeName", c.NodeName,
	)

	podClient := fpm.client.CoreV1().Pods(c.Namespace)
	labels, listOptions := fpm.labelsAndListOptionsFor(c, target)
	podList, err := podClient.List(ctx, listOptions)
	if err != nil {
		return nil, err
	}

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
			logger.V(1).Info("fuse pod exited, will be cleaned when unmount", "pod", pod.Name)
		case corev1.PodRunning:
			if isFusePodReady(&pod) {
				logger.V(2).Info("already mounted by pod", "pod", pod.Name, "target", target)
				return &pod, nil
			} else {
				startingPods = append(startingPods, pod)
			}
		default:
			startingPods = append(startingPods, pod)
		}
	}

	var fusePod *corev1.Pod
	if len(startingPods) == 0 {
		// create fuse pod for target
		template, err := fpm.PodTemplateSpec(c, target)
		if err != nil {
			return nil, err
		}
		rawPod := corev1.Pod{
			ObjectMeta: template.ObjectMeta,
			Spec:       template.Spec,
		}
		rawPod.GenerateName = fmt.Sprintf("csi-fuse-%s-", fpm.Name())
		if rawPod.Labels == nil {
			rawPod.Labels = labels
		} else {
			for key, value := range labels {
				rawPod.Labels[key] = value
			}
		}
		// make ack drain skip fuse pods
		rawPod.Labels[ACKDrainLabelKey] = "skip"

		if rawPod.Annotations == nil {
			rawPod.Annotations = make(map[string]string)
		}
		rawPod.Annotations[FuseMountPathAnnoKey] = target
		rawPod.Annotations[FuseSafeToEvictAnnoKey] = "true"

		logger.V(2).Info("creating fuse pod", "target", target)
		createdPod, err := podClient.Create(ctx, &rawPod, metav1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		logger.V(2).Info("created fuse pod", "pod", createdPod.Name, "target", target)
		fusePod = createdPod
	} else {
		if len(startingPods) > 1 {
			logger.V(1).Info("duplicated fuse pods", "num", len(startingPods), "target", target)
		}
		logger.V(2).Info("found existing fuse pod", "pod", startingPods[0].Name, "target", target)
		fusePod = &startingPods[0]
	}

	logger.V(2).Info("wait until pod is ready", "pod", fusePod.Name)
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
		fusePod = pod
		if pod.Status.Phase == corev1.PodSucceeded || pod.Status.Phase == corev1.PodFailed {
			return false, fmt.Errorf("fuse pod %s exited", pod.Name)
		}
		return pod.Status.Phase == corev1.PodRunning && isFusePodReady(pod), nil
	})
	if err != nil {
		if atomic {
			logger.V(1).Info("failed to wait for pod to be ready, deleting it", "pod", fusePod.Name)
			deleteErr := podClient.Delete(context.Background(), fusePod.Name, metav1.DeleteOptions{})
			if deleteErr != nil {
				logger.Error(deleteErr, "delete fuse pod", "pod", fusePod.Name)
			}
		}
		return nil, err
	}
	logger.V(2).Info("fuse pod is ready", "pod", fusePod.Name)
	return fusePod, nil
}

func (fpm *FusePodManager) Delete(c *FusePodContext, target string) error {
	ctx, cancel := context.WithTimeout(c, fusePodManagerTimeout)
	defer cancel()

	podClient := fpm.client.CoreV1().Pods(c.Namespace)
	_, listOptions := fpm.labelsAndListOptionsFor(c, target)

	logger := klog.Background().WithValues(
		"target", target,
		"node", c.NodeName,
		"volumeId", c.VolumeId,
		"namespace", c.Namespace,
	)
	if target != "" {
		podList, err := podClient.List(ctx, listOptions)
		if err != nil {
			return err
		}
		if len(podList.Items) == 0 {
			return nil
		}
		var errs []error
		for _, pod := range podList.Items {
			if pod.Annotations[FuseMountPathAnnoKey] == target {
				logger.V(2).Info("deleting fuse pod", "pod", pod.Name)
				err := podClient.Delete(ctx, pod.Name, metav1.DeleteOptions{})
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			return utilerrors.NewAggregate(errs)
		}
	} else {
		logger.V(2).Info("deleting fuse pods")
		err := podClient.DeleteCollection(ctx, metav1.DeleteOptions{}, listOptions)
		if err != nil {
			return err
		}
	}

	return wait.PollUntilContextCancel(ctx, time.Second*2, true, func(ctx context.Context) (done bool, err error) {
		podList, err := podClient.List(ctx, listOptions)
		if err != nil {
			return false, err
		}
		if target == "" {
			return len(podList.Items) == 0, nil
		}
		for _, pod := range podList.Items {
			if pod.Annotations[FuseMountPathAnnoKey] == target {
				return false, nil
			}
		}
		return true, nil
	})
}

func isFusePodReady(pod *corev1.Pod) bool {
	for _, cond := range pod.Status.Conditions {
		if cond.Type == corev1.PodReady {
			return cond.Status == corev1.ConditionTrue
		}
	}
	return false
}
