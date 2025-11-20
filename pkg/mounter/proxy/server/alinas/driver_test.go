package alinas

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAutoFallbackNFSMountOptions(t *testing.T) {
	tests := []struct {
		name         string
		inputOptions []string
		expected     []string
	}{
		{
			name:         "Empty options",
			inputOptions: []string{},
			expected:     []string{},
		},
		{
			name:         "EFC without VSC should add both no_add_cgroup and auto_fallback_nfs",
			inputOptions: []string{"efc"},
			expected:     []string{"efc", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "EFC with VSC should only add no_add_cgroup",
			inputOptions: []string{"efc", "net=vsc"},
			expected:     []string{"efc", "net=vsc", "no_add_cgroup"},
		},
		{
			name:         "EFC with non-VSC net option should add both options",
			inputOptions: []string{"efc", "net=other"},
			expected:     []string{"efc", "net=other", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "Multiple options with EFC but no VSC",
			inputOptions: []string{"rw", "efc", "hard"},
			expected:     []string{"rw", "efc", "hard", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "Multiple options with EFC and VSC",
			inputOptions: []string{"rw", "efc", "net=vsc", "hard"},
			expected:     []string{"rw", "efc", "net=vsc", "hard", "no_add_cgroup"},
		},
		{
			name:         "No EFC should not add any options",
			inputOptions: []string{"rw", "hard"},
			expected:     []string{"rw", "hard"},
		},
		{
			name:         "Comma-separated options with EFC but no VSC",
			inputOptions: []string{"rw,efc,hard"},
			expected:     []string{"rw,efc,hard", "no_add_cgroup", "auto_fallback_nfs"},
		},
		{
			name:         "Comma-separated options with EFC and VSC",
			inputOptions: []string{"rw,efc,net=vsc,hard"},
			expected:     []string{"rw,efc,net=vsc,hard", "no_add_cgroup"},
		},
		{
			name:         "Empty option string in slice",
			inputOptions: []string{"", "efc"},
			expected:     []string{"", "efc", "no_add_cgroup", "auto_fallback_nfs"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := addAutoFallbackNFSMountOptions(tt.inputOptions)
			assert.Equal(t, tt.expected, result)
		})
	}
}
