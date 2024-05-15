package mounter

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"strings"
	"time"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	watchtools "k8s.io/client-go/tools/watch"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

const fuseMountTimeout = time.Second * 30
const fuseMountNamespace = "kube-system"
const fuseServieAccountName = "csi-fuse-ossfs"

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

type FuseMounterType interface {
	name() string
	addPodMeta(pod *corev1.Pod)
	buildPodSpec(source, target, fstype string, authCfg *AuthConfig, options, mountFlags []string, nodeName, volumeId string) (corev1.PodSpec, error)
}

type FuseContainerConfig struct {
	Resources   corev1.ResourceRequirements
	Image       string
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
		namespace: fuseMountNamespace,
		client:    client,
	}
}

// NewFuseMounter creates a mounter for the volume id.
// When atomic is true, mount operations are responsible for cleaning up inflight fuse pods in case a timeout error occurs.
// This implies that mount operations will either succeed when the fuse pod is ready,
// or fail and ensure that no fuse pods are left behind.
func (fac *ContainerizedFuseMounterFactory) NewFuseMounter(
	ctx context.Context, volumeId string, authCfg *AuthConfig, atomic bool) *ContainerizedFuseMounter {
	return &ContainerizedFuseMounter{
		ctx:       ctx,
		atomic:    atomic,
		volumeId:  volumeId,
		nodeName:  fac.nodeName,
		namespace: fac.namespace,
		authCfg:   authCfg,
		client:    fac.client,
		log: klog.FromContext(ctx).WithValues(
			"fuseType", fac.fuseType.name(),
			"volumeID", volumeId,
		),
		FuseMounterType: fac.fuseType,
		Interface:       mountutils.New(""),
	}
}

type ContainerizedFuseMounter struct {
	ctx       context.Context
	atomic    bool
	volumeId  string
	nodeName  string
	namespace string
	authCfg   *AuthConfig
	client    kubernetes.Interface
	log       logr.Logger
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
	if mounter.authCfg != nil && mounter.authCfg.AuthType == AuthTypeRRSA {
		if mounter.authCfg.RrsaConfig.ServiceAccountName == fuseServieAccountName {
			err := mounter.checkServiceAccount(ctx)
			if err != nil {
				return err
			}
		} else {
			mounter.log.V(2).Info("serviceAccountName set, skip service account check", "serviceAccount", mounter.authCfg.RrsaConfig.ServiceAccountName)
		}
	}
	err := mounter.launchFusePod(ctx, source, target, fstype, mounter.authCfg, options, nil)
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
	err := mounter.cleanupFusePods(mounter.ctx, target)
	if err != nil {
		return fmt.Errorf("cleanup fuse pods: %w", err)
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

func (mounter *ContainerizedFuseMounter) checkServiceAccount(ctx context.Context) error {
	_, err := mounter.client.CoreV1().ServiceAccounts(mounter.namespace).Get(ctx, mounter.authCfg.RrsaConfig.ServiceAccountName, metav1.GetOptions{})
	if err != nil {
		return fmt.Errorf("check service account %s for RRSA: %w", mounter.authCfg.RrsaConfig.ServiceAccountName, err)
	}
	return nil
}

func (mounter *ContainerizedFuseMounter) launchFusePod(ctx context.Context, source, target, fstype string, authCfg *AuthConfig, options, mountFlags []string) error {
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
			mounter.log.V(1).Info("exited fuse pod, will be cleaned when unmount", "pod", pod.Name)
		case corev1.PodRunning:
			if isFusePodReady(&pod) {
				ready = true
				mounter.log.V(2).Info("already mounted", "target", target, "pod", pod.Name)
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
		rawPod.Annotations = map[string]string{FuseMountPathAnnoKey: target, FuseSafeToEvictAnnoKey: "true"}
		mounter.addPodMeta(&rawPod)
		rawPod.Spec, err = mounter.buildPodSpec(source, target, fstype, authCfg, options, mountFlags, mounter.nodeName, mounter.volumeId)
		if err != nil {
			return err
		}
		mounter.log.V(3).Info("creating fuse pod", "target", target)
		createdPod, err := podClient.Create(ctx, &rawPod, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		mounter.log.V(2).Info("created fuse pod", "pod", createdPod.Name, "target", target)
		fusePod = createdPod
	} else {
		if len(startingPods) > 1 {
			mounter.log.V(1).Info("duplicated fuse pods", "numPods", len(startingPods), "target", target)
		}
		mounter.log.V(2).Info("found existed fuse pod", "pod", startingPods[0].Name, "target", target)
		fusePod = &startingPods[0]
	}

	mounter.log.V(2).Info("waiting until pod ready", "pod", fusePod.Name)
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
		if mounter.atomic {
			mounter.log.V(1).Info("failed to wait for pod ready, deleting it", "pod", fusePod.Name)
			deleteErr := podClient.Delete(context.Background(), fusePod.Name, metav1.DeleteOptions{})
			if deleteErr != nil {
				mounter.log.Error(deleteErr, "delete fuse pod", "pod", fusePod.Name)
			}
		}
		return err
	}
	mounter.log.V(2).Info("fuse pod ready", "pod", fusePod.Name)
	return nil
}

func (mounter *ContainerizedFuseMounter) cleanupFusePods(ctx context.Context, target string) error {
	podClient := mounter.client.CoreV1().Pods(mounter.namespace)
	_, listOptions := mounter.labelsAndListOptionsFor(target)
	podList, err := podClient.List(ctx, listOptions)
	if err != nil {
		return err
	}
	var errs []error
	for _, pod := range podList.Items {
		if pod.Annotations[FuseMountPathAnnoKey] == target {
			mounter.log.V(2).Info("deleting fuse pod", "pod", pod.Name, "target", target)
			err := podClient.Delete(ctx, pod.Name, metav1.DeleteOptions{})
			errs = append(errs, err)
		}
	}
	return utilerrors.NewAggregate(errs)
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

func getRoleSessionName(volumeId, target string) string {
	return fmt.Sprintf("ossfs.%s.%s", volumeId, computeMountPathHash(target))
}
