package fuse_pod_manager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	corev1 "k8s.io/api/core/v1"
	apiserrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/watch"
	informercorev1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	watchtools "k8s.io/client-go/tools/watch"
	"k8s.io/klog/v2"
)

const (
	fusePodManagerTimeout  = time.Second * 30
	FuseServiceAccountName = "csi-fuse-ossfs"
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
	// for AK/SK with or without token
	Secrets map[string]string
	// for Token from Secret
	SecretRef string
	// for STS(ECS worker role)/RRSA
	RoleName string
}

type RrsaConfig struct {
	OidcProviderArn    string
	RoleArn            string
	ServiceAccountName string
	AssumeRoleArn      string
}

type PodTemplateConfig struct {
	DnsPolicy corev1.DNSPolicy
}

const (
	DebugLevelFatal = "fatal"
	DebugLevelError = "error"
	DebugLevelWarn  = "warn"
	DebugLevelInfo  = "info"
	DebugLevelDebug = "debug"
)

const (
	MetricsModeDisabled = "disabled"
	MetricsModeEnabled  = "enabled"
)

type FusePodContext struct {
	context.Context
	Namespace         string
	NodeName          string
	VolumeId          string
	FuseType          string
	AuthConfig        *AuthConfig
	PodTemplateConfig *PodTemplateConfig
}

type FuseMounterType interface {
	Name() string
	PodTemplateSpec(c *FusePodContext, target string) (*corev1.PodTemplateSpec, error)
	AddDefaultMountOptions(options []string) []string
}

type FuseContainerConfig struct {
	Resources   corev1.ResourceRequirements
	Image       string
	ImageTag    string
	Dbglevel    string
	MetricsMode string
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
			err = utils.ValidateAnnotations(annotations)
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
			err = utils.ValidateLabels(labels)
			if err != nil {
				invalid = true
				break
			}
			config.Labels = labels
		case "metrics-mode":
			switch value {
			case MetricsModeDisabled, MetricsModeEnabled:
				config.MetricsMode = value
			default:
				invalid = true
			}
		default:
			// Initialize Extra map if nil. Currently supported keys in Extra:
			// - "set-dumpable": enables dumpable flag for process of ossfs 1.0 / ossfs 2.0
			// - "mime-support": enables MIME type support for ossfs 1.0
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
		FuseVolumeIdLabelKey: utils.ComputeVolumeIdLabelVal(c.VolumeId),
	}
	// ControllerUnPublish cannot get fuseType info,
	// so FuseTypeLabelKey cannot used as a label for Delete
	if c.FuseType != "" {
		labels[FuseTypeLabelKey] = c.FuseType
	}
	if target != "" {
		labels[FuseMountPathHashLabelKey] = utils.ComputeMountPathHash(target)
	}
	listOptions := metav1.ListOptions{
		FieldSelector: fields.OneTermEqualSelector("spec.nodeName", c.NodeName).String(),
		LabelSelector: metav1.FormatLabelSelector(&metav1.LabelSelector{MatchLabels: labels}),
	}
	return labels, listOptions
}

func (fpm *FusePodManager) Create(c *FusePodContext, target string) (*corev1.Pod, error) {
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
		return nil, err
	}
	logger.V(2).Info("fuse pod is ready", "pod", fusePod.Name)
	return fusePod, nil
}

func (fpm *FusePodManager) Delete(c *FusePodContext) error {
	ctx, cancel := context.WithTimeout(c, fusePodManagerTimeout)
	defer cancel()

	logger := klog.FromContext(ctx).WithValues("namespace", c.Namespace)

	_, listOptions := fpm.labelsAndListOptionsFor(c, "")
	informer := informercorev1.NewFilteredPodInformer(fpm.client, c.Namespace, 0, nil, func(options *metav1.ListOptions) {
		options.FieldSelector = listOptions.FieldSelector
		options.LabelSelector = listOptions.LabelSelector
	})
	deleteNotify := make(chan struct{}, 1)
	deleteNotify <- struct{}{}
	_, err := informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		DeleteFunc: func(obj interface{}) {
			select {
			case deleteNotify <- struct{}{}:
				logger.V(4).Info("pod deleted")
			default:
				logger.V(4).Info("pod deleted, already notified")
			}
		},
	})
	if err != nil {
		return fmt.Errorf("failed to add event handler to pod informer: %w", err)
	}
	go informer.Run(ctx.Done())
	synced := cache.WaitForCacheSync(ctx.Done(), informer.HasSynced)
	if !synced {
		return fmt.Errorf("failed to wait for caches to sync: %w", ctx.Err())
	}

	podClient := fpm.client.CoreV1().Pods(c.Namespace)
	for _, obj := range informer.GetStore().List() {
		pod := obj.(*corev1.Pod)
		if pod.DeletionTimestamp == nil {
			logger.V(2).Info("deleting fuse pod", "pod", klog.KObj(pod))
			err := podClient.Delete(ctx, pod.Name, metav1.DeleteOptions{})
			if err != nil {
				if apiserrors.IsNotFound(err) {
					continue
				}
				return fmt.Errorf("failed to delete pod %s: %w", pod.Name, err)
			}
		} else {
			logger.V(2).Info("fuse pod already terminating", "pod", klog.KObj(pod))
		}
	}

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("wait for pods deleted: %w", ctx.Err())
		case <-deleteNotify:
			n := len(informer.GetStore().ListKeys())
			if n == 0 {
				logger.V(2).Info("all pods deleted")
				return nil
			}
			logger.V(2).Info("wait for pods deleted", "count", n)
		}
	}
}

func isFusePodReady(pod *corev1.Pod) bool {
	for _, cond := range pod.Status.Conditions {
		if cond.Type == corev1.PodReady {
			return cond.Status == corev1.ConditionTrue
		}
	}
	return false
}
