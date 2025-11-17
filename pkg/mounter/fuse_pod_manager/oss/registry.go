package oss

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
)

var (
	fstypeToFactory      = map[string]func(*corev1.ConfigMap, metadata.MetadataProvider) OSSFuseMounterType{}
	fstypeToPath         = map[string]string{}
	fstypeToInterceptors = map[string][]mounter.MountInterceptor{}
)

// RegisterFuseMounter registers a fuse mounter factory for a given fstype
func RegisterFuseMounter(fstype string, factory func(*corev1.ConfigMap, metadata.MetadataProvider) OSSFuseMounterType) {
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
	result := make(map[string]string, len(fstypeToPath))
	for fstype, path := range fstypeToPath {
		result[fstype] = path
	}
	return result
}

// GetFuseMountInterceptors returns the mount interceptors for a given fstype
func GetFuseMountInterceptors(fstype string) ([]mounter.MountInterceptor, bool) {
	path, ok := fstypeToInterceptors[fstype]
	return path, ok
}

// GetFuseMounter returns a fuse mounter instance for the given fstype
func GetFuseMounter(fstype string, configmap *corev1.ConfigMap, m metadata.MetadataProvider) (OSSFuseMounterType, error) {
	factory, ok := fstypeToFactory[fstype]
	if !ok {
		return nil, &UnsupportedFstypeError{Fstype: fstype}
	}
	return factory(configmap, m), nil
}

// GetAllRegisteredFuseTypes returns all registered fuse types
func GetAllRegisteredFuseTypes() []string {
	types := make([]string, 0, len(fstypeToFactory))
	for fstype := range fstypeToFactory {
		types = append(types, fstype)
	}
	return types
}

// GetAllOSSFusePodManagers creates a map of all registered OSS fuse pod managers
// configmap can be nil if not available (e.g., in CSI agent mode)
// client can be nil if not needed (e.g., in CSI agent mode)
func GetAllOSSFusePodManagers(configmap *corev1.ConfigMap, m metadata.MetadataProvider, client kubernetes.Interface) map[string]*OSSFusePodManager {
	fusePodManagers := make(map[string]*OSSFusePodManager, len(fstypeToFactory))
	for fstype, factory := range fstypeToFactory {
		fusePodManagers[fstype] = NewOSSFusePodManager(factory(configmap, m), client)
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
