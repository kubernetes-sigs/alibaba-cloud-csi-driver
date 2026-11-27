//go:build !windows

package metric

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func TestFuseUpdateCollectsGuestWithEmptyHostDir(t *testing.T) {
	host, guest := t.TempDir(), t.TempDir()
	vol := filepath.Join(guest, "pod-a", "pv-a")
	require.NoError(t, os.MkdirAll(vol, 0755))
	require.NoError(t, os.WriteFile(filepath.Join(guest, "pod-a", utils.PodInfoFile), []byte("ns pod pod-a 10"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(vol, utils.MountPointInfoFile), []byte("ossfs oss bucket pv-a /target"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(vol, utils.MetricsMountPointStatus), []byte("1"), 0644))
	alias := filepath.Join(t.TempDir(), "alias")
	require.NoError(t, os.Symlink(guest, alias))
	oldDirs := fuseMetricDirs
	t.Cleanup(func() { fuseMetricDirs = oldDirs })
	stale := t.TempDir()
	require.NoError(t, os.MkdirAll(filepath.Join(stale, "pod-a/pv-a"), 0755))
	require.NoError(t, os.WriteFile(filepath.Join(stale, "pod-a", utils.PodInfoFile), []byte("ns pod pod-a 10"), 0644))
	require.NoError(t, os.WriteFile(filepath.Join(stale, "pod-a/pv-a", utils.MountPointInfoFile), []byte("ossfs oss bucket pv-a /target"), 0644))
	fuseMetricDirs = map[string][]string{"ossfs": {host, stale, guest, alias}}
	collector, err := NewFuseStatCollector()
	require.NoError(t, err)
	ch := make(chan prometheus.Metric, 4)
	require.NoError(t, collector.Update(t.Context(), nil, ch))
	require.Len(t, ch, 1)
	require.Contains(t, (<-ch).Desc().String(), "mount_point_status")
}
