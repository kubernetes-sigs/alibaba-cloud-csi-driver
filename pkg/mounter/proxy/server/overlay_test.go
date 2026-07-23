package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverlayDirs(t *testing.T) {
	target := "/run/csi/mount-root/oss/abc123"

	lower, upper, work := overlayDirs(target)

	// All should be under the same parent directory
	assert.Contains(t, lower, overlayBaseDir)
	assert.Contains(t, upper, overlayBaseDir)
	assert.Contains(t, work, overlayBaseDir)

	// Should end with the correct subdirectory
	assert.Contains(t, lower, "/lower")
	assert.Contains(t, upper, "/upper")
	assert.Contains(t, work, "/work")

	// Same target should produce the same paths (deterministic)
	lower2, _, _ := overlayDirs(target)
	assert.Equal(t, lower, lower2)

	// Different target should produce different paths (isolation)
	otherLower, _, _ := overlayDirs("/run/csi/mount-root/oss/other")
	assert.NotEqual(t, lower, otherLower)
}

func TestOverlayLowerDir(t *testing.T) {
	target := "/run/csi/mount-root/oss/abc123"

	lower := OverlayLowerDir(target)
	expectedLower, _, _ := overlayDirs(target)
	assert.Equal(t, expectedLower, lower)
}
