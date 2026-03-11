package om

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseReferenceMount(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected []string
	}{
		{
			name:     "valid single file",
			line:     `some message is still referenced from other Pods [/var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/pvc-abc/dev/vdb]`,
			expected: []string{"/var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/pvc-abc/dev/vdb"},
		},
		{
			name:     "no reference tag",
			line:     "some unrelated log message",
			expected: []string{},
		},
		{
			name:     "missing bracket",
			line:     "is still referenced from other Pods no brackets here",
			expected: []string{},
		},
		{
			name:     "empty bracket content",
			line:     "is still referenced from other Pods []",
			expected: []string{""},
		},
		{
			name:     "missing closing bracket",
			line:     "is still referenced from other Pods [/some/path",
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseReferenceMount(tt.line)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestIsFileBlockLink(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected bool
	}{
		{
			name:     "valid block link path",
			file:     "/var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/pvc-54574338/dev/vdb",
			expected: true,
		},
		{
			name:     "missing /dev/ segment",
			file:     "/var/lib/kubelet/plugins/kubernetes.io/csi/volumeDevices/pvc-54574338/vdb",
			expected: false,
		},
		{
			name:     "wrong prefix",
			file:     "/tmp/plugins/kubernetes.io/csi/volumeDevices/pvc-abc/dev/vdb",
			expected: false,
		},
		{
			name:     "empty string",
			file:     "",
			expected: false,
		},
		{
			name:     "random path",
			file:     "/var/log/messages",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isFileBlockLink(tt.file)
			assert.Equal(t, tt.expected, got, "isFileBlockLink(%q)", tt.file)
		})
	}
}

func TestFixReferenceMountIssue_InvalidLine(t *testing.T) {
	// A line without the reference pattern should return false
	result := FixReferenceMountIssue("some random log line without reference")
	assert.False(t, result, "expected false for invalid line")
}

func TestFixReferenceMountIssue_ValidFormatNonExistentFiles(t *testing.T) {
	// Files listed in the message don't exist - function should still return true
	line := "is still referenced from other Pods [/nonexistent/path/dev/vdb]"
	result := FixReferenceMountIssue(line)
	assert.True(t, result, "expected true for valid format line")
}
