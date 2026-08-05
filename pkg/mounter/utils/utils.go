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

	"golang.org/x/sys/unix"
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

// OverlayBaseDir is the base directory for overlay lower/upper/work dirs.
// Defaults to /run/csi-overlay; override via CSI_OVERLAY_BASE_DIR env var
// (e.g. point to a temp dir in unit tests to avoid writing to /run).
var OverlayBaseDir = func() string {
	if dir := os.Getenv("CSI_OVERLAY_BASE_DIR"); dir != "" {
		return dir
	}
	return "/run/csi-overlay"
}()

// OverlayDirs computes the lower, upper, and work dir paths for a given overlay target.
//
// The key is the merged path (target) exposed to the business container, NOT the volumeID.
//
//   - volumeID is unstable: in Sandbox, the volumeID contains a random suffix that changes
//     across checkpoint clone and deep hibernate resume. Using it would orphan the upper dir
//     data preserved in the VM filesystem.
//
//   - target (merged path) is stable: it is deterministically derived from the container's
//     mountPath, which is preserved across checkpoint/clone/resume.
//
//   - Isolation is guaranteed by target uniqueness: each volume mount has a distinct mountPath,
//     so each gets a distinct overlay dir. Two mounts of the same PV to different paths are
//     fully isolated.
//
// NOTE: This design relies on VM-level isolation (each Sandbox is a separate VM). In RunC
// (future), multiple pods on the same host share the filesystem, so target alone cannot
// provide inter-pod isolation — a per-pod key (e.g. podUID) would be needed. RunC also does
// not need hostPath for overlay because the upper dir lives directly on the host filesystem,
// avoiding the overlay-on-overlay problem and the hibernate persistence requirement.
func OverlayDirs(target string) (lower, upper, work string) {
	h := sha256.Sum256([]byte(target))
	base := filepath.Join(OverlayBaseDir, hex.EncodeToString(h[:]))
	return filepath.Join(base, "lower"), filepath.Join(base, "upper"), filepath.Join(base, "work")
}

// OverlayLowerDir returns the lower dir path for a given overlay target (merged path).
func OverlayLowerDir(target string) string {
	lower, _, _ := OverlayDirs(target)
	return lower
}

// CleanupOverlayLowerDir performs best-effort cleanup of the overlay lower dir.
// After the overlay (merged) is unmounted by NodeUnpublish, the lower dir mount may
// still be alive. This function attempts to unmount it.
// targetPath is the merged path (same key used to compute overlay dirs).
// Safe to call even if the lower dir is not mounted (no-op in that case).
func CleanupOverlayLowerDir(targetPath string) {
	lower, _, _ := OverlayDirs(targetPath)
	m := mountutils.NewWithoutSystemd("")
	notMnt, err := m.IsLikelyNotMountPoint(lower)
	if err != nil || notMnt {
		return
	}
	klog.InfoS("Cleaning up overlay lower dir", "path", lower)
	if err := mountutils.CleanupMountPoint(lower, m, false); err != nil {
		klog.ErrorS(err, "Best-effort cleanup of overlay lower dir failed", "path", lower)
	}
}

// statfs is a package variable so tests can simulate a mount whose daemon is gone.
var statfs = unix.Statfs

// IsNotLiveMountPoint reports whether target needs a (re)mount: it is not a mount
// point at all, or it is one whose filesystem no longer answers.
//
// IsNotMountPoint decides on stat alone, which is not enough for FUSE. When the
// daemon dies the connection is aborted, yet the kernel can still answer stat for
// the mount root from the cached inode, so the mount looks healthy and the caller
// skips the remount that would repair it. statfs is not cached and always reaches
// the daemon, so it is what actually settles liveness.
//
// A mount that is present but unserviced is unmounted before returning true, so the
// caller can mount again without stacking a second filesystem on the same path.
//
// This is strictly additive: it never reports a failure that IsNotMountPoint would not
// have reported. An unexpected statfs error is logged and the mount is left alone,
// because this runs on the normal publish path and a probe must not be able to break a
// mount that works.
//
// Note the probe blocks if the daemon is gone but the connection is still open,
// which is what fd-passing recovery deliberately arranges. Callers that support that
// mode must skip this check.
func IsNotLiveMountPoint(mounter mountutils.Interface, target string) (bool, error) {
	notMnt, err := IsNotMountPoint(mounter, target)
	if err != nil || notMnt {
		return notMnt, err
	}

	var st unix.Statfs_t
	if serr := statfs(target, &st); serr != nil {
		if !mountutils.IsCorruptedMnt(serr) {
			klog.Warningf("Cannot probe mountpoint %s, assuming it is alive: %v", target, serr)
			return false, nil
		}
		klog.Warningf("Umount unserviced mountpoint %s", target)
		if uerr := mounter.Unmount(target); uerr != nil {
			return false, status.Errorf(codes.Internal, "umount unserviced mountpoint %s: %v", target, uerr)
		}
		return true, nil
	}
	return false, nil
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
