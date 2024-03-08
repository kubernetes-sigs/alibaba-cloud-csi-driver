package common

import (
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/component-base/featuregate"
)

var (
	// DefaultMutableFeatureGate is a mutable version of DefaultFeatureGate.
	// Only top-level commands/options setup and the k8s.io/component-base/featuregate/testing package should make use of this.
	// Tests that need to modify feature gates for the duration of their test should use:
	//   defer featuregatetesting.SetFeatureGateDuringTest(t, utilfeature.DefaultFeatureGate, features.<FeatureName>, <value>)()
	DefaultMutableFeatureGate featuregate.MutableFeatureGate = featuregate.NewFeatureGate()

	// DefaultFeatureGate is a shared global FeatureGate.
	// Top-level commands/options setup that needs to modify this feature gate should use DefaultMutableFeatureGate.
	DefaultFeatureGate featuregate.FeatureGate = DefaultMutableFeatureGate
)

const (
	// VolumeGroupSnapshot enables CSI create volume group snapshot
	VolumeGroupSnapshot featuregate.Feature = "VolumeGroupSnapshot"
)

var defaultFeatureGates = map[featuregate.Feature]featuregate.FeatureSpec{
	VolumeGroupSnapshot: {Default: false, PreRelease: featuregate.Alpha},
}

func init() {
	runtime.Must(DefaultMutableFeatureGate.Add(defaultFeatureGates))
}
