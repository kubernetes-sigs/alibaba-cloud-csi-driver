package server

import (
	"os"
	"path/filepath"
	"testing"

	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// realistic mountinfo lines: the container rootfs overlay (upperdir under containerd)
// must be ignored; only overlays whose upperdir is under OverlayBaseDir are ours.
const mountInfoFixture = `436 331 0:91 / / rw,relatime - overlay overlay rw,lowerdir=/var/lib/containerd/snapshots/219/fs,upperdir=/var/lib/containerd/snapshots/222/fs,workdir=/var/lib/containerd/snapshots/222/work
298 598 0:108 / /run/csi/mount-root/oss/aaaa rw,relatime shared:155 - overlay overlay rw,lowerdir=/run/csi-overlay/h1/lower,upperdir=/run/csi-overlay/h1/upper,workdir=/run/csi-overlay/h1/work
299 598 0:109 / /run/csi/mount-root/oss/bbbb rw,relatime shared:156 - overlay overlay rw,lowerdir=/run/csi-overlay/h2/lower,upperdir=/run/csi-overlay/h2/upper,workdir=/run/csi-overlay/h2/work
310 331 0:200 / /run/cnfs/x rw,relatime - fuse.ossfs ossfs rw,user_id=0
320 331 0:201 / /some/other rw,relatime - overlay overlay rw,lowerdir=/a,upperdir=/run/csi-overlay-decoy/upper,workdir=/run/csi-overlay-decoy/work
`

func writeFixture(t *testing.T, content string) string {
	t.Helper()
	dir := t.TempDir()
	p := filepath.Join(dir, "mountinfo")
	require.NoError(t, os.WriteFile(p, []byte(content), 0o644))
	return p
}

func TestListOwnedOverlays(t *testing.T) {
	origPath, origBase := mountInfoPath, mounterutils.OverlayBaseDir
	defer func() { mountInfoPath, mounterutils.OverlayBaseDir = origPath, origBase }()

	mountInfoPath = writeFixture(t, mountInfoFixture)
	mounterutils.OverlayBaseDir = "/run/csi-overlay"

	got, err := listOwnedOverlays()
	require.NoError(t, err)

	// Only the two overlays whose upperdir is under /run/csi-overlay. The container
	// rootfs (containerd upperdir), the fuse mount, and the /run/csi-overlay-decoy
	// mount (prefix without a directory boundary) must all be excluded.
	assert.ElementsMatch(t, []string{
		"/run/csi/mount-root/oss/aaaa",
		"/run/csi/mount-root/oss/bbbb",
	}, got)
}

func TestListOwnedOverlays_BaseDirBoundary(t *testing.T) {
	origPath, origBase := mountInfoPath, mounterutils.OverlayBaseDir
	defer func() { mountInfoPath, mounterutils.OverlayBaseDir = origPath, origBase }()

	// A configured base dir with a trailing slash must still match cleanly, and must
	// not accidentally match the sibling "-decoy" path.
	mountInfoPath = writeFixture(t, mountInfoFixture)
	mounterutils.OverlayBaseDir = "/run/csi-overlay/"

	got, err := listOwnedOverlays()
	require.NoError(t, err)
	assert.ElementsMatch(t, []string{
		"/run/csi/mount-root/oss/aaaa",
		"/run/csi/mount-root/oss/bbbb",
	}, got)
}

func TestListOwnedOverlays_NoOwned(t *testing.T) {
	origPath, origBase := mountInfoPath, mounterutils.OverlayBaseDir
	defer func() { mountInfoPath, mounterutils.OverlayBaseDir = origPath, origBase }()

	mountInfoPath = writeFixture(t, mountInfoFixture)
	// Point the base somewhere nothing in the fixture lives under.
	mounterutils.OverlayBaseDir = "/mnt/nothing-here"

	got, err := listOwnedOverlays()
	require.NoError(t, err)
	assert.Empty(t, got)
}

func TestUpperDirUnder(t *testing.T) {
	base := "/run/csi-overlay"
	tests := []struct {
		name    string
		options []string
		want    bool
	}{
		{"owned", []string{"lowerdir=/x", "upperdir=/run/csi-overlay/h/upper", "workdir=/x"}, true},
		{"container rootfs", []string{"upperdir=/var/lib/containerd/snapshots/222/fs"}, false},
		{"sibling prefix decoy", []string{"upperdir=/run/csi-overlay-decoy/upper"}, false},
		{"no upperdir", []string{"lowerdir=/x", "workdir=/y"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, upperDirUnder(tt.options, base))
		})
	}
}
