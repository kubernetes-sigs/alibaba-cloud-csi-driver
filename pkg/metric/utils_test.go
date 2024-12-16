package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGlobalMountPathByDiskID(t *testing.T) {
	assert.Equal(t,
		"/var/lib/kubelet/plugins/kubernetes.io/csi/diskplugin.csi.alibabacloud.com/b3f919f99444cdb7b9a4e3906e14c76a1f12111725e2a8044de53a312c61ba2c/globalmount",
		getGlobalMountPathByDiskID("d-testdiskid"))
}
