package mounter

import (
	"context"
	"encoding/json"
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
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	watchtools "k8s.io/client-go/tools/watch"
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
)

type AuthConfig struct {
	AuthType string
	// for RRSA
	RrsaConfig *RrsaConfig
	// for csi-secret-store
	SecretProviderClassName string
	// for AK/SK
	Secrets map[string]string
	// for STS(ECS worker role)
	RoleName string
}

type RrsaConfig struct {
	OidcProviderArn    string
	RoleArn            string
	ServiceAccountName string
}

const (
	AuthTypeSTS  = "sts"
	AuthTypeRRSA = "rrsa"
	AuthTypeCSS  = "csi-secret-store"
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
	AuthConfig *AuthConfig
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
	Annotations map[string]string
	Labels      map[string]string
	Extra       map[string]string
}

func extractFuseContainerConfig(configmap *corev1.ConfigMap, name string) (config FuseContainerConfig) {
	if configmap == nil {
		return
	}
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
		case "dbglevel":
			switch value {
			case DebugLevelFatal, DebugLevelError, DebugLevelWarn, DebugLevelInfo, DebugLevelDebug:
				config.Dbglevel = value
			default:
				invalid = true
				break
			}
		case "image":
			logrus.Warn("'image' config in configmap no longer supported")
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
			logrus.Warnf("ignore invalid configuration line: %q", line)
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
		FuseTypeLabelKey:     fpm.Name(),
		FuseVolumeIdLabelKey: c.VolumeId,
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

func (fpm *FusePodManager) Create(c *FusePodContext, target string, atomic bool) error {
	ctx, cancel := context.WithTimeout(c, fusePodManagerTimeout)
	defer cancel()

	log := logrus.WithFields(logrus.Fields{
		"volumeId": c.VolumeId,
		"nodeName": c.NodeName,
	})

	podClient := fpm.client.CoreV1().Pods(c.Namespace)
	labels, listOptions := fpm.labelsAndListOptionsFor(c, target)
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
			log.Warnf("exited fuse pod %s, will be cleaned when unmount", pod.Name)
		case corev1.PodRunning:
			if isFusePodReady(&pod) {
				ready = true
				log.Infof("already mounted by pod %s for %s", pod.Name, target)
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
		template, err := fpm.PodTemplateSpec(c, target)
		if err != nil {
			return err
		}
		rawPod := corev1.Pod{
			ObjectMeta: template.ObjectMeta,
			Spec:       template.Spec,
		}
		rawPod.GenerateName = fmt.Sprintf("csi-fuse-%s-", fpm.Name())
		rawPod.Labels = labels
		rawPod.Annotations = map[string]string{FuseMountPathAnnoKey: target, FuseSafeToEvictAnnoKey: "true"}

		// check service account
		if rawPod.Spec.ServiceAccountName != "" {
			_, err := fpm.client.CoreV1().ServiceAccounts(c.Namespace).Get(c, rawPod.Spec.ServiceAccountName, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("check service account %s for RRSA: %w", rawPod.Spec.ServiceAccountName, err)
			}
		}

		log.Infof("creating fuse pod for %s", target)
		createdPod, err := podClient.Create(ctx, &rawPod, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		log.Infof("created fuse pod %s for %s", createdPod.Name, target)
		fusePod = createdPod
	} else {
		if len(startingPods) > 1 {
			log.Warnf("%d duplicated fuse pods for %s", len(startingPods), target)
		}
		log.Infof("found existed fuse pod %s for %s", startingPods[0].Name, target)
		fusePod = &startingPods[0]
	}

	log.Infof("wait util pod %s is ready", fusePod.Name)
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
		if atomic {
			log.Warnf("failed to wait for pod %s to be ready, deleting it", fusePod.Name)
			deleteErr := podClient.Delete(context.Background(), fusePod.Name, metav1.DeleteOptions{})
			if deleteErr != nil {
				log.WithError(deleteErr).Errorf("delete fuse pod %s", fusePod.Name)
			}
		}
		return err
	}
	log.Infof("fuse pod %s is ready", fusePod.Name)
	return nil
}

func (fpm *FusePodManager) Delete(c *FusePodContext, target string) error {
	ctx, cancel := context.WithTimeout(c, fusePodManagerTimeout)
	defer cancel()

	podClient := fpm.client.CoreV1().Pods(c.Namespace)
	_, listOptions := fpm.labelsAndListOptionsFor(c, target)

	log := logrus.WithFields(logrus.Fields{
		"target":    target,
		"node":      c.NodeName,
		"volumeId":  c.VolumeId,
		"namespace": c.Namespace,
	})
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
				log.Infof("deleting fuse pod %s", pod.Name)
				err := podClient.Delete(ctx, pod.Name, metav1.DeleteOptions{})
				errs = append(errs, err)
			}
		}
		if len(errs) > 0 {
			return utilerrors.NewAggregate(errs)
		}
	} else {
		log.Info("deleting fuse pods")
		err := podClient.DeleteCollection(ctx, metav1.DeleteOptions{}, listOptions)
		if err != nil {
			return err
		}
	}

	return wait.PollUntilWithContext(ctx, time.Second*2, func(ctx context.Context) (done bool, err error) {
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

func computeMountPathHash(target string) string {
	hasher := fnv.New32a()
	hasher.Write([]byte(target))
	return rand.SafeEncodeString(fmt.Sprint(hasher.Sum32()))
}
