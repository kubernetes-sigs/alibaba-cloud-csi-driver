package nas

import (
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/features"
	"github.com/stretchr/testify/assert"
	featuregatetesting "k8s.io/component-base/featuregate/testing"
)

func TestGetMetricsPathPrefix(t *testing.T) {
	t.Run("agentMode", func(t *testing.T) {
		assert.Equal(t, "/var/run/efc/", getMetricsPathPrefix(true))
	})

	t.Run("hostPath", func(t *testing.T) {
		assert.Equal(t, "/host/var/run/efc/", getMetricsPathPrefix(false))
	})

	t.Run("mountProxy", func(t *testing.T) {
		featuregatetesting.SetFeatureGateDuringTest(t, features.FunctionalMutableFeatureGate, features.AlinasMountProxy, true)
		assert.Equal(t, "/run/cnfs/efc/", getMetricsPathPrefix(false))
	})
}
