package metric

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetGlobalMountPathByDiskID(t *testing.T) {
	assert.Equal(t,
		"/var/lib/kubelet/plugins/kubernetes.io/csi/diskplugin.csi.alibabacloud.com/b3f919f99444cdb7b9a4e3906e14c76a1f12111725e2a8044de53a312c61ba2c/globalmount",
		getGlobalMountPathByDiskID("d-testdiskid"))
}

func writeVolData(t *testing.T, content string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), volDataFile)
	require.NoError(t, os.WriteFile(path, []byte(content), 0644))
	return path
}

func TestGetPvcByVolData(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		path := writeVolData(t, `{"csi.alibabacloud.com/pvc-namespace":"ns1","csi.alibabacloud.com/pvc-name":"pvc1"}`)
		ref, err := getPvcByVolData(path)
		require.NoError(t, err)
		assert.Equal(t, "ns1", ref.Namespace)
		assert.Equal(t, "pvc1", ref.Name)
	})

	t.Run("missingNamespace", func(t *testing.T) {
		path := writeVolData(t, `{"csi.alibabacloud.com/pvc-name":"pvc1"}`)
		_, err := getPvcByVolData(path)
		require.Error(t, err)
	})

	t.Run("missingPvcName", func(t *testing.T) {
		path := writeVolData(t, `{"driverName":"diskplugin.csi.alibabacloud.com"}`)
		_, err := getPvcByVolData(path)
		assert.Error(t, err)
	})

	t.Run("missingFile", func(t *testing.T) {
		_, err := getPvcByVolData(filepath.Join(t.TempDir(), "absent.json"))
		assert.Error(t, err)
	})
}

func TestGetPvcByPvNameWithoutClient(t *testing.T) {
	path := writeVolData(t, `{"csi.alibabacloud.com/pvc-namespace":"ns1","csi.alibabacloud.com/pvc-name":"pvc1"}`)

	ref, err := getDiskPvcByPvName(context.Background(), nil, "pv1", path)
	require.NoError(t, err)
	assert.Equal(t, "pvc1", ref.Name)

	ns, name, server, err := getNasPvcByPvName(context.Background(), nil, nil, "pv1", path)
	require.NoError(t, err)
	assert.Empty(t, server)
	assert.Equal(t, "ns1", ns)
	assert.Equal(t, "pvc1", name)
}
