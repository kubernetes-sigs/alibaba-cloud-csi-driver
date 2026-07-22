package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOverlayDirs(t *testing.T) {
	volumeID := "pvc-12345"

	lower, upper, work := overlayDirs(volumeID)

	// All should be under the same parent directory
	assert.Contains(t, lower, overlayBaseDir)
	assert.Contains(t, upper, overlayBaseDir)
	assert.Contains(t, work, overlayBaseDir)

	// Should end with the correct subdirectory
	assert.Contains(t, lower, "/lower")
	assert.Contains(t, upper, "/upper")
	assert.Contains(t, work, "/work")

	// Same volumeID should produce the same paths
	lower2, _, _ := overlayDirs(volumeID)
	assert.Equal(t, lower, lower2)

	// Different volumeID should produce different paths
	otherLower, _, _ := overlayDirs("pvc-other")
	assert.NotEqual(t, lower, otherLower)
}

func TestOverlayLowerDir(t *testing.T) {
	volumeID := "pvc-12345"

	lower := OverlayLowerDir(volumeID)
	expectedLower, _, _ := overlayDirs(volumeID)
	assert.Equal(t, expectedLower, lower)
}
