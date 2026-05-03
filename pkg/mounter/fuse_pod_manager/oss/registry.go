package oss

import (
	"maps"
	"os"
	"strconv"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	k8sver "k8s.io/apimachinery/pkg/util/version"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

var (
	fstypeToFactory      = map[string]func(utils.Config, metadata.MetadataProvider) OSSFuseMounterType{}
	fstypeToPath         = map[string]string{}
	fstypeToInterceptors = map[string][]mounter.MountInterceptor{}
)

// RegisterFuseMounter registers a fuse mounter factory for a given fstype
func RegisterFuseMounter(fstype string, factory func(utils.Config, metadata.MetadataProvider) OSSFuseMounterType) {
	fstypeToFactory[fstype] = factory
}

// RegisterFuseMounterPath registers the executable path for a given fstype
func RegisterFuseMounterPath(fstype, path string) {
	fstypeToPath[fstype] = path
}

// RegisterFuseInterceptors registers the mount interceptors for a given fstype
func RegisterFuseInterceptors(fstype string, interceptors []mounter.MountInterceptor) {
	fstypeToInterceptors[fstype] = interceptors
}

// GetFuseMounterPath returns the executable path for a given fstype
func GetFuseMounterPath(fstype string) (string, bool) {
	path, ok := fstypeToPath[fstype]
	return path, ok
}

// GetAllFuseMounterPaths returns a map of all registered fuse types to their executable paths
func GetAllFuseMounterPaths() map[string]string {
	return maps.Clone(fstypeToPath)
}

// GetFuseMountInterceptors returns the mount interceptors for a given fstype
func GetFuseMountInterceptors(fstype string) ([]mounter.MountInterceptor, bool) {
	path, ok := fstypeToInterceptors[fstype]
	return path, ok
}

// GetFuseMounter returns a fuse mounter instance for the given fstype
func GetFuseMounter(fstype string, csiCfg utils.Config, m metadata.MetadataProvider) (OSSFuseMounterType, error) {
	factory, ok := fstypeToFactory[fstype]
	if !ok {
		return nil, &UnsupportedFstypeError{Fstype: fstype}
	}
	return factory(csiCfg, m), nil
}

// GetAllRegisteredFuseTypes returns all registered fuse types
func GetAllRegisteredFuseTypes() []string {
	types := make([]string, 0, len(fstypeToFactory))
	for fstype := range fstypeToFactory {
		types = append(types, fstype)
	}
	return types
}

// ShouldConstrainResourceVersion determines whether to constrain ResourceVersion to "" for consistent reads
// Priority: ENV override > K8s version detection
// - ENV OSS_FUSE_CONSTRAIN_RV=true/false: If set, use this value
// - K8s >= 1.31: false (Consistent Reads from Cache supported, no need to constrain)
// - K8s < 1.31: true (Constrain RV="0" to use watch cache and avoid etcd pressure)
// - K8s unknown: true (Conservative approach, constrain RV="0")
func ShouldConstrainResourceVersion(k8sVersion *k8sver.Version) bool {
	// Check ENV override first
	if envVal := os.Getenv("OSS_FUSE_CONSTRAIN_RV"); envVal != "" {
		if constrained, err := strconv.ParseBool(envVal); err == nil {
			if constrained {
				klog.Warningf("ENV OSS_FUSE_CONSTRAIN_RV=true, constraining ResourceVersion to '0' for fuse pod delete operations")
			} else {
				klog.Infof("ENV OSS_FUSE_CONSTRAIN_RV=false, using ResourceVersion='' for fuse pod delete operations")
			}
			return constrained
		}
		klog.Warningf("ENV OSS_FUSE_CONSTRAIN_RV has invalid value %q, ignoring", envVal)
	}

	// Fall back to K8s version detection
	if k8sVersion == nil {
		klog.Errorf("K8s version unknown, constraining ResourceVersion to '0' for fuse pod delete operations (conservative approach)")
		return true
	}

	if k8sVersion.AtLeast(k8sver.MajorMinor(1, 31)) {
		klog.Infof("K8s version %s >= 1.31, not constraining ResourceVersion (Consistent Reads from Cache supported)", k8sVersion.String())
		return false
	}

	klog.Warningf("K8s version %s < 1.31, constraining ResourceVersion to '0' for fuse pod delete operations (avoid etcd pressure)", k8sVersion.String())
	return true
}

// GetAllOSSFusePodManagers creates a map of all registered OSS fuse pod managers
// configmap can be nil if not available (e.g., in CSI agent mode)
// client can be nil if not needed (e.g., in CSI agent mode)
func GetAllOSSFusePodManagers(csiCfg utils.Config, m metadata.MetadataProvider, client kubernetes.Interface, k8sVersion *k8sver.Version) map[string]*OSSFusePodManager {
	constrainRV := ShouldConstrainResourceVersion(k8sVersion)

	fusePodManagers := make(map[string]*OSSFusePodManager, len(fstypeToFactory))
	for fstype, factory := range fstypeToFactory {
		fusePodManagers[fstype] = NewOSSFusePodManager(factory(csiCfg, m), client, constrainRV)
	}
	return fusePodManagers
}

// UnsupportedFstypeError represents an error when fstype is not supported
type UnsupportedFstypeError struct {
	Fstype string
}

func (e *UnsupportedFstypeError) Error() string {
	return "unsupported fstype: " + e.Fstype
}
