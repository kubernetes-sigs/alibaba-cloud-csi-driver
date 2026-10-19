//go:build !windows

package nas

import (
	"testing"

	"github.com/container-storage-interface/spec/lib/go/csi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"
)

// Check successful parsing against invariants independently of error wording.
// Seed cases run in the regular suite; fuzzing exercises rounding and cap edges.
func FuzzComputeAgenticSpaceSizeLimit(f *testing.F) {
	f.Add(int64(10*GiB), int64(0), "")
	f.Add(int64(10*GiB+1), int64(0), "10.5Gi")
	f.Add(int64(0), int64(20*GiB), "10Gi")
	f.Add(int64(-1), int64(0), "invalid")
	f.Add(int64(maxAgenticSpaceSizeLimit), int64(0), "1024000Gi")
	f.Fuzz(func(t *testing.T, required, limit int64, capParam string) {
		if len(capParam) > 128 {
			t.Skip()
		}
		got, err := computeAgenticSpaceSizeLimit(&csi.CapacityRange{RequiredBytes: required, LimitBytes: limit}, capParam)
		if err != nil {
			return
		}
		require.Positive(t, got)
		assert.Zero(t, got%GiB)
		assert.LessOrEqual(t, got, maxAgenticSpaceSizeLimit)
		if required > 0 {
			assert.GreaterOrEqual(t, got, required)
		}
		if limit > 0 {
			assert.LessOrEqual(t, got, limit)
		}
		if capParam != "" {
			q, err := resource.ParseQuantity(capParam)
			require.NoError(t, err)
			assert.LessOrEqual(t, got, q.Value())
		}
	})
}
