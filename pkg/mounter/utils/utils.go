package utils

import (
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

const (
	OssFsType  = "ossfs"
	OssFs2Type = "ossfs2"
)

// keys for STS token
const (
	KeyAccessKeyId     = "AccessKeyId"
	KeyAccessKeySecret = "AccessKeySecret"
	KeyExpiration      = "Expiration"
	KeySecurityToken   = "SecurityToken"
)

const LegacyFusePodNamespace = "kube-system" // deprecated

// fuseAttachBaseDir is the base directory for fuse attach paths.
// Default is "/run", but can be overridden for testing.
var (
	fuseAttachBaseDir     = "/run"
	fuseAttachBaseDirLock sync.RWMutex
)

// SetFuseAttachBaseDir sets the base directory for fuse attach paths.
// This is primarily used for testing to avoid /run permission issues.
// It should be called before any calls to GetFuseAttachDir() or GetAttachPath().
func SetFuseAttachBaseDir(dir string) {
	fuseAttachBaseDirLock.Lock()
	defer fuseAttachBaseDirLock.Unlock()
	fuseAttachBaseDir = dir
}

// GetFuseAttachBaseDir returns the current base directory for fuse attach paths.
func GetFuseAttachBaseDir() string {
	fuseAttachBaseDirLock.RLock()
	defer fuseAttachBaseDirLock.RUnlock()
	return fuseAttachBaseDir
}

func ComputeMountPathHash(target string) string {
	hasher := fnv.New32a()
	hasher.Write([]byte(target))
	return rand.SafeEncodeString(fmt.Sprint(hasher.Sum32()))
}

func ComputeVolumeIdLabelVal(volumeId string) string {
	errs := validation.IsValidLabelValue(volumeId)
	if errs == nil {
		return volumeId
	}
	volSha := sha1.Sum([]byte(volumeId))
	return "h1." + hex.EncodeToString(volSha[:])
}

// https://github.com/kubernetes/kubernetes/blob/b5ba7bc4f5f49760c821cae2f152a8000922e72e/staging/src/k8s.io/apimachinery/pkg/api/validation/objectmeta.go#L36
// TotalAnnotationSizeLimitB only takes 128 kB here, and the rest is reserved for the default annotations.
const TotalAnnotationSizeLimitB int = 128 * (1 << 10) // 128 kB

// Copy from https://github.com/kubernetes/mount-utils/blob/41e8de37ef8a3782f9cd6c915699b98b2b24b2c4/mount_helper_unix.go#L164
func SplitMountOptions(s string) []string {
	inQuotes := false
	list := strings.FieldsFunc(s, func(r rune) bool {
		if r == '"' {
			inQuotes = !inQuotes
		}
		// Report a new field only when outside of double quotes.
		return r == ',' && !inQuotes
	})
	return list
}

// IndexMountOptions parses mount options into a key-value map for lookup.
// Options in "key=value" form are split at the first "=".
// Flag-only options like "ro" have an empty string value.
func IndexMountOptions(options []string) map[string]string {
	m := make(map[string]string, len(options))
	for _, o := range options {
		if o == "" {
			continue
		}
		k, v, _ := strings.Cut(o, "=")
		m[strings.TrimSpace(k)] = strings.TrimSpace(v)
	}
	return m
}

// MergeMountOptions merges additional options into base options.
// If the same key exists in both, the base value takes precedence.
// When values conflict, a warning is logged.
// Options can be "key=value" or "key" (flag-only).
func MergeMountOptions(base, additional []string) []string {
	existing := IndexMountOptions(base)
	for _, op := range additional {
		k, v, _ := strings.Cut(op, "=")
		k = strings.TrimSpace(k)
		v = strings.TrimSpace(v)
		if ev, ok := existing[k]; ok {
			if ev != v {
				klog.Warningf("MergeMountOptions: option %q conflicts with existing key %q, keeping existing value", op, k)
			}
			continue
		}
		existing[k] = v
		base = append(base, op)
	}
	return base
}

// Copy from https://github.com/kubernetes/kubernetes/blob/b5ba7bc4f5f49760c821cae2f152a8000922e72e/staging/src/k8s.io/apimachinery/pkg/api/validation/objectmeta.go#L43
// ValidateAnnotations validates that a set of annotations are correctly defined.
func ValidateAnnotations(annotations map[string]string) error {
	for k := range annotations {
		// The rule is QualifiedName except that case doesn't matter, so convert to lowercase before checking.
		err := ValidateKey(strings.ToLower(k))
		if err != nil {
			return err
		}
	}
	if err := ValidateAnnotationsSize(annotations); err != nil {
		return err
	}
	return nil
}

// Copy from https://github.com/kubernetes/kubernetes/blob/b5ba7bc4f5f49760c821cae2f152a8000922e72e/staging/src/k8s.io/apimachinery/pkg/apis/meta/v1/validation/validation.go#L105
// ValidateLabels validates that a set of labels are correctly defined.
func ValidateLabels(labels map[string]string) error {
	for k, v := range labels {
		err := ValidateKey(k)
		if err != nil {
			return err
		}
		err = ValidateLabelValue(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func ValidateKey(k string) error {
	if errs := validation.IsQualifiedName(k); len(errs) != 0 {
		return fmt.Errorf("invalid key %q: %s", k, strings.Join(errs, "; "))
	}
	return nil
}

func ValidateLabelValue(k, v string) error {
	if errs := validation.IsValidLabelValue(v); len(errs) != 0 {
		return fmt.Errorf("invalid label value: %q: at key: %q: %s", v, k, strings.Join(errs, "; "))
	}
	return nil
}

// Copy from https://github.com/kubernetes/kubernetes/blob/b5ba7bc4f5f49760c821cae2f152a8000922e72e/staging/src/k8s.io/apimachinery/pkg/api/validation/objectmeta.go#L58
func ValidateAnnotationsSize(annotations map[string]string) error {
	var totalSize int64
	for k, v := range annotations {
		totalSize += (int64)(len(k)) + (int64)(len(v))
	}
	if totalSize > (int64)(TotalAnnotationSizeLimitB) {
		return fmt.Errorf("annotations size %d is larger than limit %d", totalSize, TotalAnnotationSizeLimitB)
	}
	return nil
}

// GetOIDCProvider get OIDC provider from env or ACK clusterId for RRSA
func GetOIDCProvider(clusterId string) string {
	return fmt.Sprintf("ack-rrsa-%s", clusterId)
}

// GetArn get rrsa config for fuse container's env setting
func GetArn(provider, accountId, roleName string) (oidcProviderArn, roleArn string) {
	if provider == "" || accountId == "" || roleName == "" {
		return
	}
	roleArn = fmt.Sprintf("acs:ram::%s:role/%s", accountId, roleName)
	oidcProviderArn = fmt.Sprintf("acs:ram::%s:oidc-provider/%s", accountId, provider)
	return
}

func GetMountProxySocketPath(volumeId string) string {
	volSha := sha256.Sum256([]byte(volumeId))
	return filepath.Join(GetFuseAttachDir(), hex.EncodeToString(volSha[:]), "mounter.sock")
}

func GetFuseAttachDir() string {
	// Notes: as OSS driver used /run/fuse.ossfs/* mount dir before,
	// and NodeUnstageVolume request do not contain fuseType info for unmount,
	// so all kinds of fuseTypes share this unified mount dir.
	// A volumeId should only belong to one kind of fuseType, and mounted ONCE.
	fuseType := OssFsType
	baseDir := GetFuseAttachBaseDir()
	return filepath.Join(baseDir, fmt.Sprintf("fuse.%s", fuseType))
}

func GetAttachPath(volumeId string) string {
	volSha := sha256.Sum256([]byte(volumeId))
	return filepath.Join(GetFuseAttachDir(), hex.EncodeToString(volSha[:]), "globalmount")
}

func GetCredentialsSecretName(fuseType string) string {
	return fmt.Sprintf("csi-%s-credentials", fuseType)
}

func CleanupCredentialSecret(ctx context.Context, clientset kubernetes.Interface, node, volumeId, fuseType string) error {
	if clientset == nil {
		klog.V(2).InfoS("Skip cleaning up credential secret due to nil kube client")
		return nil
	}
	key := fmt.Sprintf("%s.%s", node, volumeId)
	secretName := GetCredentialsSecretName(fuseType)
	secretClient := clientset.CoreV1().Secrets(LegacyFusePodNamespace)
	secret, err := secretClient.Get(ctx, secretName, metav1.GetOptions{})
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
	_, err = secretClient.Patch(ctx, secretName, types.StrategicMergePatchType, patchData, metav1.PatchOptions{})
	if err == nil {
		klog.V(2).InfoS("patched secret to remove credentials", "secret", secretName, "volumeId", volumeId)
	}
	return err
}

const MaxRoleSessionNameLimit = 64

func GetRoleSessionName(volumeId, target, fuseType string) string {
	name := fmt.Sprintf("%s.%s.%s", fuseType, volumeId, ComputeMountPathHash(target))
	if len(name) > MaxRoleSessionNameLimit {
		name = name[:MaxRoleSessionNameLimit]
	}
	return name
}

// ConfigDir stores the config and passwd files for fuse to load
func GetConfigDir(fuseType string) string {
	return fmt.Sprintf("/etc/%s", fuseType)
}

func GetPasswdFileName(fuseType string) string {
	return fmt.Sprintf("passwd-%s", fuseType)
}

func GetPasswdHashDir(target string) string {
	return filepath.Join("/tmp", ComputeMountPathHash(target))
}

func IsNotMountPoint(mounter mountutils.Interface, target string) (notMnt bool, err error) {
	notMnt, err = mounter.IsLikelyNotMountPoint(target)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(target, os.ModePerm); err != nil {
				return false, status.Errorf(codes.Internal, "mkdir: %v", err)
			}
			return true, nil
		} else if mountutils.IsCorruptedMnt(err) {
			klog.Warningf("Umount corrupted mountpoint %s", target)
			err := mounter.Unmount(target)
			if err != nil {
				return false, status.Errorf(codes.Internal, "umount corrupted mountpoint %s: %v", target, err)
			}
			return true, nil
		}
		return false, status.Errorf(codes.Internal, "check mountpoint: %v", err)
	}
	return notMnt, nil
}
