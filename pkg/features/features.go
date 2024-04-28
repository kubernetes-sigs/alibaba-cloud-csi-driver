package features

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"
)

const (
	DiskADController    featuregate.Feature = "DiskADController"
	DBFSMetricByPlugin  featuregate.Feature = "DBFSMetricByPlugin"
	DBFSADController    featuregate.Feature = "DBFSADController"
	UpdatedOssfsVersion featuregate.Feature = "UpdatedOssfsVersion"
	RundCSIProtocol3    featuregate.Feature = "RundCSIProtocol3"
)

var (
	FunctionalMutableFeatureGate = featuregate.NewFeatureGate()
	defaultDiskFeatureGate       = map[featuregate.Feature]featuregate.FeatureSpec{
		DiskADController: {Default: false, PreRelease: featuregate.Alpha},
	}
	defaultDBFSFeatureGate = map[featuregate.Feature]featuregate.FeatureSpec{
		DBFSMetricByPlugin: {Default: false, PreRelease: featuregate.Alpha},
		DBFSADController:   {Default: false, PreRelease: featuregate.Alpha},
	}
	defaultOSSFeatureGate = map[featuregate.Feature]featuregate.FeatureSpec{
		UpdatedOssfsVersion: {Default: false, PreRelease: featuregate.Alpha},
	}
	defaultNASFeatureGate = map[featuregate.Feature]featuregate.FeatureSpec{
		RundCSIProtocol3: {Default: false, PreRelease: featuregate.Alpha},
	}
)

func init() {
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultDiskFeatureGate))
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultDBFSFeatureGate))
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultOSSFeatureGate))
	runtime.Must(FunctionalMutableFeatureGate.Add(defaultNASFeatureGate))
}
