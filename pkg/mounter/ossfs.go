package mounter

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

var defaultOssfsImageTag = "v1.88.4-80d165c-aliyun"
var defaultOssfsUpdatedImageTag = "v1.91.5.ack.1-ed398f6-aliyun"

const (
	hostPrefix                = "/host"
	OssfsCredentialSecretName = "csi-ossfs-credentials"
	OssfsDefMimeTypesFilePath = "/etc/mime.types"
	// The ossfs image includes a default MIME configuration located at /csi/mime.types
	OssfsCsiMimeTypesFilePath = "/csi/mime.types"
	OssfsPasswdFile           = "passwd-ossfs"

	defaultRegistry = "registry-cn-hangzhou.ack.aliyuncs.com"

	CsiSecretStoreDriver   = "secrets-store.csi.k8s.io"
	SecretProviderClassKey = "secretProviderClass"

	OssfsAttachDir = "/run/fuse.ossfs"

	PasswdMountDir = "/etc/ossfs"
	PasswdFilename = "passwd-ossfs"

	MaxRoleSessionNameLimit = 64
)

type fuseOssfs struct {
	config FuseContainerConfig
}

var ossfsDbglevels = map[string]string{
	DebugLevelDebug: "debug",
	DebugLevelInfo:  "info",
	DebugLevelWarn:  "warn",
	DebugLevelError: "err",
	DebugLevelFatal: "crit",
}

const defaultDbglevel = DebugLevelError

func NewFuseOssfs(configmap *corev1.ConfigMap, m metadata.MetadataProvider) FuseMounterType {
	config := extractFuseContainerConfig(configmap, "ossfs")
	// set default image
	if config.Image == "" {
		registry := os.Getenv("DEFAULT_REGISTRY")
		if registry == "" {
			region, err := m.Get(metadata.RegionID)
			if err == nil {
				registry = fmt.Sprintf("registry-%s-vpc.ack.aliyuncs.com", region)
			} else {
				klog.Warningf("DEFAULT_REGISTRY env not set, failed to get current region: %v, fallback to default registry: %s", err, defaultRegistry)
				registry = defaultRegistry
			}
		}
		if config.ImageTag == "" {
			if features.FunctionalMutableFeatureGate.Enabled(features.UpdatedOssfsVersion) {
				config.ImageTag = defaultOssfsUpdatedImageTag
			} else {
				config.ImageTag = defaultOssfsImageTag
			}
		}
		config.Image = fmt.Sprintf("%s/acs/csi-ossfs:%s", registry, config.ImageTag)
		klog.Infof("Use ossfs image: %s", config.Image)
	}
	// set default memory request
	if _, ok := config.Resources.Requests[corev1.ResourceMemory]; !ok {
		config.Resources.Requests[corev1.ResourceMemory] = resource.MustParse("50Mi")
	}

	return &fuseOssfs{config: config}
}

func (f *fuseOssfs) Name() string {
	return "ossfs"
}

func (f *fuseOssfs) PodTemplateSpec(c *FusePodContext, target string) (*corev1.PodTemplateSpec, error) {
	spec, err := f.buildPodSpec(c, target)
	if err != nil {
		return nil, err
	}

	pod := new(corev1.PodTemplateSpec)
	pod.Spec = spec

	pod.Annotations = make(map[string]string)
	for k, v := range f.config.Annotations {
		pod.Annotations[k] = v
	}
	pod.Labels = make(map[string]string)
	for k, v := range f.config.Labels {
		pod.Labels[k] = v
	}
	return pod, nil
}

func (f *fuseOssfs) buildPodSpec(c *FusePodContext, target string) (spec corev1.PodSpec, _ error) {
	targetDir := filepath.Dir(target)
	targetDirVolume := corev1.Volume{
		Name: "target-dir",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: targetDir,
				Type: new(corev1.HostPathType),
			},
		},
	}
	*targetDirVolume.HostPath.Type = corev1.HostPathDirectoryOrCreate
	metricsDirVolume := corev1.Volume{
		Name: "metrics-dir",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/run/ossfs",
				Type: new(corev1.HostPathType),
			},
		},
	}
	*metricsDirVolume.HostPath.Type = corev1.HostPathDirectoryOrCreate
	// mime.types/csi-mime.types in /etc is used
	etcDirVolume := corev1.Volume{
		Name: "host-etc",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/etc",
				Type: new(corev1.HostPathType),
			},
		},
	}
	*etcDirVolume.HostPath.Type = corev1.HostPathDirectory

	spec.Volumes = []corev1.Volume{targetDirVolume, metricsDirVolume, etcDirVolume}

	bidirectional := corev1.MountPropagationBidirectional
	socketPath := GetOssfsMountProxySocketPath(c.VolumeId)
	container := corev1.Container{
		Name:      "ossfs",
		Image:     f.config.Image,
		Resources: f.config.Resources,
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:             targetDirVolume.Name,
				MountPath:        targetDir,
				MountPropagation: &bidirectional,
			}, {
				Name:      metricsDirVolume.Name,
				MountPath: metricsDirVolume.HostPath.Path,
			}, {
				Name:      etcDirVolume.Name,
				MountPath: path.Join(hostPrefix, etcDirVolume.HostPath.Path),
			},
		},
		SecurityContext: &corev1.SecurityContext{
			Privileged: tea.Bool(true),
		},
		ReadinessProbe: &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				Exec: &corev1.ExecAction{
					Command: []string{
						"test", "-S", socketPath,
					},
				},
			},
			PeriodSeconds:    2,
			FailureThreshold: 5,
		},
	}

	buildAuthSpec(c, target, &spec, &container)

	container.Args = []string{"-socket=" + socketPath, "-v=4"}

	spec.Containers = []corev1.Container{container}
	spec.NodeName = c.NodeName
	spec.HostNetwork = true
	spec.PriorityClassName = "system-node-critical"
	spec.Tolerations = []corev1.Toleration{{Operator: corev1.TolerationOpExists}}
	return
}

func (f *fuseOssfs) AddDefaultMountOptions(options []string) []string {
	alreadySet := false
	for _, option := range options {
		if strings.Contains(option, "dbglevel") {
			alreadySet = true
			break
		}
	}
	if !alreadySet {
		level, ok := ossfsDbglevels[f.config.Dbglevel]
		if ok {
			options = append(options, fmt.Sprintf("dbglevel=%s", level))
		} else {
			if f.config.Dbglevel != "" {
				klog.Warningf("invalid dbglevel for ossfs: %q, use default dbglevel %s", f.config.Dbglevel, defaultDbglevel)
			}
			options = append(options, fmt.Sprintf("dbglevel=%s", ossfsDbglevels[defaultDbglevel]))
		}
	}

	if !utils.IsFileExisting(filepath.Join(hostPrefix, OssfsDefMimeTypesFilePath)) && strings.ToLower(f.config.Extra["mime-support"]) == "true" {
		// mime.types not exists, use csi-mime.types
		options = append(options, fmt.Sprintf("mime=%s", OssfsCsiMimeTypesFilePath))
	}

	return options
}

func buildAuthSpec(c *FusePodContext, target string, spec *corev1.PodSpec, container *corev1.Container) {
	if spec == nil || container == nil {
		return
	}
	authCfg := c.AuthConfig
	if authCfg == nil {
		return
	}

	switch authCfg.AuthType {
	case AuthTypeSTS:
	case AuthTypeRRSA:
		if authCfg.RrsaConfig == nil {
			return
		}
		spec.ServiceAccountName = authCfg.RrsaConfig.ServiceAccountName
		rrsaMountDir := "/var/run/secrets/ack.alibabacloud.com/rrsa-tokens"
		rrsaVolume := corev1.Volume{
			Name: "rrsa-oidc-token",
			VolumeSource: corev1.VolumeSource{
				Projected: &corev1.ProjectedVolumeSource{
					DefaultMode: tea.Int32(0600),
					Sources: []corev1.VolumeProjection{
						{
							ServiceAccountToken: &corev1.ServiceAccountTokenProjection{
								Audience:          "sts.aliyuncs.com",
								ExpirationSeconds: tea.Int64(3600),
								Path:              "token",
							},
						},
					},
				},
			},
		}
		spec.Volumes = append(spec.Volumes, rrsaVolume)
		rrsaVolumeMount := corev1.VolumeMount{
			Name:      rrsaVolume.Name,
			MountPath: rrsaMountDir,
		}
		container.VolumeMounts = append(container.VolumeMounts, rrsaVolumeMount)
		envs := []corev1.EnvVar{
			{
				Name:  "ALIBABA_CLOUD_ROLE_ARN",
				Value: authCfg.RrsaConfig.RoleArn,
			},
			{
				Name:  "ALIBABA_CLOUD_OIDC_PROVIDER_ARN",
				Value: authCfg.RrsaConfig.OidcProviderArn,
			},
			{
				Name:  "ALIBABA_CLOUD_OIDC_TOKEN_FILE",
				Value: rrsaMountDir + "/token",
			},
			{
				Name:  "ROLE_SESSION_NAME",
				Value: getRoleSessionName(c.VolumeId, target),
			},
		}
		container.Env = append(container.Env, envs...)
	case AuthTypeCSS:
		secretStoreMountDir := "/etc/ossfs/secrets-store"
		secretStoreVolume := corev1.Volume{
			Name: "secrets-store",
			VolumeSource: corev1.VolumeSource{
				CSI: &corev1.CSIVolumeSource{
					Driver:           CsiSecretStoreDriver,
					ReadOnly:         tea.Bool(true),
					VolumeAttributes: map[string]string{SecretProviderClassKey: authCfg.SecretProviderClassName},
				},
			},
		}
		spec.Volumes = append(spec.Volumes, secretStoreVolume)
		secretStoreVolumeMount := corev1.VolumeMount{
			Name:      secretStoreVolume.Name,
			MountPath: secretStoreMountDir,
			ReadOnly:  true,
		}
		container.VolumeMounts = append(container.VolumeMounts, secretStoreVolumeMount)
	default:
		secretVolumeSource := getPasswdSecretVolume(authCfg.SecretRef)
		if secretVolumeSource != nil {
			passwdSecretVolume := corev1.Volume{
				Name: "passwd-ossfs",
				VolumeSource: corev1.VolumeSource{
					Secret: secretVolumeSource,
				},
			}
			spec.Volumes = append(spec.Volumes, passwdSecretVolume)
			passwdVolumeMont := corev1.VolumeMount{
				Name:      passwdSecretVolume.Name,
				MountPath: PasswdMountDir,
			}
			container.VolumeMounts = append(container.VolumeMounts, passwdVolumeMont)
		}
	}
}

func CleanupOssfsCredentialSecret(ctx context.Context, clientset kubernetes.Interface, node, volumeId string) error {
	key := fmt.Sprintf("%s.%s", node, volumeId)
	secretClient := clientset.CoreV1().Secrets(LegacyFusePodNamespace)
	secret, err := secretClient.Get(ctx, OssfsCredentialSecretName, metav1.GetOptions{})
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil
		}
		return err
	}
	_, exists := secret.Data[key]
	if !exists {
		return nil
	}
	// patch secret
	patch := corev1.Secret{
		Data: map[string][]byte{
			key: nil,
		},
	}
	patchData, err := json.Marshal(patch)
	if err != nil {
		return err
	}
	_, err = secretClient.Patch(ctx, OssfsCredentialSecretName, types.StrategicMergePatchType, patchData, metav1.PatchOptions{})
	if err == nil {
		klog.V(2).InfoS("patched secret to remove credentials", "secret", OssfsCredentialSecretName, "volumeId", volumeId)
	}
	return err
}

func getRoleSessionName(volumeId, target string) string {
	name := fmt.Sprintf("ossfs.%s.%s", volumeId, computeMountPathHash(target))
	if len(name) > MaxRoleSessionNameLimit {
		name = name[:MaxRoleSessionNameLimit]
	}
	return name
}

func GetOssfsMountProxySocketPath(volumeId string) string {
	volSha := sha256.Sum256([]byte(volumeId))
	return filepath.Join(OssfsAttachDir, hex.EncodeToString(volSha[:]), "mounter.sock")
}

func GetOssfsAttachPath(volumeId string) string {
	volSha := sha256.Sum256([]byte(volumeId))
	return filepath.Join(OssfsAttachDir, hex.EncodeToString(volSha[:]), "globalmount")
}

// keep consistent with RAM response
var secretRefKeysToParse []string = []string{
	"AccessKeyId",
	"AccessKeySecret",
	"Expiration",
	"SecurityToken",
}

func getPasswdSecretVolume(secretRef string) (secret *corev1.SecretVolumeSource) {
	passwdFilename := "passwd-ossfs"
	if secretRef == "" {
		return nil
	}
	items := []corev1.KeyToPath{}
	for _, key := range secretRefKeysToParse {
		item := corev1.KeyToPath{
			Key:  key,
			Path: fmt.Sprintf("%s/%s", passwdFilename, key),
			Mode: tea.Int32(0600),
		}
		items = append(items, item)
	}
	secret = &corev1.SecretVolumeSource{
		SecretName: secretRef,
		Items:      items,
	}
	return
}
