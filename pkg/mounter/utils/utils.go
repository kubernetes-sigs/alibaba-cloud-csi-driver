package utils

import (
	"context"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"hash/fnv"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sys/unix"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/rand"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

const (
	OssFsType  = "ossfs"
	OssFs2Type = "ossfs2"
)

const LegacyFusePodNamespace = "kube-system" // deprecated

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
	return fmt.Sprintf("/run/fuse.%s", fuseType)
}

func GetAttachPath(volumeId string) string {
	volSha := sha256.Sum256([]byte(volumeId))
	return filepath.Join(GetFuseAttachDir(), hex.EncodeToString(volSha[:]), "globalmount")
}

func GetCredientialsSecretName(fuseType string) string {
	return fmt.Sprintf("csi-%s-credentials", fuseType)
}

func CleanupCredentialSecret(ctx context.Context, clientset kubernetes.Interface, node, volumeId, fuseType string) error {
	key := fmt.Sprintf("%s.%s", node, volumeId)
	secretName := GetCredientialsSecretName(fuseType)
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

func WaitFdReadable(fd int, timeout time.Duration) error {
	tv := unix.Timeval{
		Sec: int64(timeout.Seconds()),
	}
	readSet := new(unix.FdSet)
	readSet.Zero()
	readSet.Set(fd)
	n, err := unix.Select(fd+1, readSet, nil, nil, &tv)
	if err != nil {
		return err
	}
	if n == 0 {
		return errors.New("timeout waiting fd")
	} else if n < 0 {
		return fmt.Errorf("unexpected select result: %d", n)
	}
	return nil
}

func SaveOssSecretsToFile(secrets map[string]string, fuseType string) (filePath string, err error) {
	passwd := secrets[GetPasswdFileName(fuseType)]
	if passwd == "" {
		return
	}

	tmpDir, err := os.MkdirTemp("", fuseType+"-")
	if err != nil {
		return "", err
	}
	filePath = filepath.Join(tmpDir, "passwd")
	if err = os.WriteFile(filePath, []byte(passwd), 0o600); err != nil {
		return "", err
	}
	klog.V(4).InfoS(fmt.Sprintf("created %s passwd file", fuseType), "path", filePath)
	return
}
