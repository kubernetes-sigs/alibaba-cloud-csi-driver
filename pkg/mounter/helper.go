package mounter

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/options"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// https://github.com/kubernetes/kubernetes/blob/b5ba7bc4f5f49760c821cae2f152a8000922e72e/staging/src/k8s.io/apimachinery/pkg/api/validation/objectmeta.go#L36
// TotalAnnotationSizeLimitB only takes 128 kB here, and the rest is reserved for the default annotations.
const TotalAnnotationSizeLimitB int = 128 * (1 << 10) // 128 kB

const (
	// UidResource is ali-uid url subpath
	UidResource = "owner-account-id"
)

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
	valid, err := utils.CheckRequestArgs(annotations)
	if !valid {
		return err
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

// GetAliUid get aliUid from env or metadata server for RRSA
func GetAliUid() (aliUid string) {
	aliUid = os.Getenv("AlI_UID")
	if aliUid != "" {
		return
	}
	aliUid = utils.RetryGetMetaData(UidResource)
	return
}

// GetClusterId get clusterId from env or profile for RRSA
func GetClusterId() (clusterId string, err error) {
	clusterId = os.Getenv("CLUSTER_ID")
	if clusterId != "" {
		return
	}
	cfg, err := clientcmd.BuildConfigFromFlags(options.MasterURL, options.Kubeconfig)
	if err != nil {
		return
	}
	clientset, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return
	}
	profile, err := clientset.CoreV1().ConfigMaps("kube-system").Get(context.Background(), "ack-cluster-profile", metav1.GetOptions{})
	if err != nil {
		return
	}
	clusterId = profile.Data["clusterid"]
	return
}

// GetOIDCProvider get OIDC provider from env or ACK clusterId for RRSA
func GetOIDCProvider() (provider string, err error) {
	provider = os.Getenv("OIDC_PROVIDER")
	if provider != "" {
		return
	}
	clusterId, err := GetClusterId()
	if err != nil {
		return
	}
	provider = fmt.Sprintf("ack-rrsa-%s", clusterId)
	return
}

// GetArn get rrsa config for fuse contianer's env setting
func GetArn(provider, aliUid, roleName string) (roleArn, providerArn string) {
	if provider == "" || aliUid == "" || roleName == "" {
		return
	}
	roleArn = fmt.Sprintf("acs:ram::%s:role/%s", aliUid, roleName)
	providerArn = fmt.Sprintf("acs:ram::%s:oidc-provider/%s", aliUid, provider)
	return
}
