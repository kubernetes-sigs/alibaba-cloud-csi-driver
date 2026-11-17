//go:build !windows

package metric

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetPVName(t *testing.T) {
	assert.Equal(t, "pv1", getPVName(&Mount{Mount: "/var/lib/kubelet/pods/uid/volumes/kubernetes.io~csi/pv1/mount"}))
	assert.Equal(t, "", getPVName(nil))
	assert.Equal(t, "", getPVName(&Mount{Mount: "mount"}))
}
