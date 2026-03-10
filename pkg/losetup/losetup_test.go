package losetup

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithAssociatedFile(t *testing.T) {
	file := "/dev/loop0"
	opt := WithAssociatedFile(file)
	assert.Len(t, opt, 2, "expected 2 elements")
	assert.Equal(t, "--associated", opt[0])
	assert.Equal(t, file, opt[1])
}

func TestList_Empty(t *testing.T) {
	// List without any loop devices should return empty or nil
	devices, err := List()
	if err != nil {
		// losetup command might not be available on some systems
		t.Skipf("losetup command not available: %v", err)
	}
	// On systems without loop devices, this should return nil or empty slice
	_ = devices
}

func TestList_WithAssociatedFile_NotFound(t *testing.T) {
	// Query for a non-existent file should return empty
	devices, err := List(WithAssociatedFile("/nonexistent/file.img"))
	if err != nil {
		// losetup command might not be available on some systems
		t.Skipf("losetup command not available: %v", err)
	}
	assert.Empty(t, devices, "expected 0 devices for non-existent file")
}

func TestOption_AppendArgs(t *testing.T) {
	base := []string{"--list", "--json"}
	opt := WithAssociatedFile("/tmp/file")
	combined := append(base, opt...)
	expected := []string{"--list", "--json", "--associated", "/tmp/file"}
	assert.Equal(t, expected, combined)
}
