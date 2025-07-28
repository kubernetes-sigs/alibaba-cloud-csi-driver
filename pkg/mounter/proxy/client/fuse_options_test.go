package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var defaultFuseOptions = []string{
	"nodev",
	"nosuid",
	"rootmode=40000",
	"allow_other",
	"default_permissions",
	fmt.Sprintf("user_id=%d", os.Getuid()),
	fmt.Sprintf("group_id=%d", os.Getegid()),
}

func TestSplitFuseOptions(t *testing.T) {
	var modifyFuseOptions, debugFuseoptions []string

	modifyFuseOptions = append(modifyFuseOptions, defaultFuseOptions...)
	modifyFuseOptions[2] = "rootmode=0600"

	debugFuseoptions = append(debugFuseoptions, defaultFuseOptions...)
	debugFuseoptions = append(debugFuseoptions, "debug")

	tests := []struct {
		name            string
		options         []string
		defaultFuseOpts map[string]string
		fuseOptionKeys  map[string]int
		nullVal         string
		expectedFuse    []string
		expectedDaemon  []string
		expectError     bool
	}{
		{
			name:           "Empty options",
			options:        []string{},
			expectedFuse:   defaultFuseOptions,
			expectedDaemon: []string{},
			expectError:    false,
		},
		{
			name:           "Only default options with same values",
			options:        []string{"allow_other", "rootmode=40000"},
			expectedFuse:   defaultFuseOptions,
			expectedDaemon: []string{"allow_other", "rootmode=40000"},
			expectError:    false,
		},
		{
			name:           "Default options with different values",
			options:        []string{"rootmode=0600", "default_permissions"},
			expectedFuse:   modifyFuseOptions,
			expectedDaemon: []string{"rootmode=0600", "default_permissions"},
			expectError:    false,
		},
		{
			name:           "Supported non-default options",
			options:        []string{"debug"},
			expectedFuse:   debugFuseoptions,
			expectedDaemon: []string{"debug"},
			expectError:    false,
		},
		{
			name:           "Ignored options",
			options:        []string{"rw", "allow_other"},
			expectedFuse:   defaultFuseOptions,
			expectedDaemon: []string{"allow_other"},
			expectError:    false,
		},
		{
			name:           "Unsupported options should return error",
			options:        []string{"dev"},
			expectedFuse:   nil,
			expectedDaemon: nil,
			expectError:    true,
		},
		{
			name:           "Daemon options",
			options:        []string{"another_daemon_opt"},
			expectedFuse:   defaultFuseOptions,
			expectedDaemon: []string{"another_daemon_opt"},
			expectError:    false,
		},
		{
			name:           "Mixed options",
			options:        []string{"allow_other", "daemon_opt=value"},
			expectedFuse:   defaultFuseOptions,
			expectedDaemon: []string{"allow_other", "daemon_opt=value"},
			expectError:    false,
		},
		{
			name:           "Invalid empty kv option",
			options:        []string{""},
			expectedFuse:   defaultFuseOptions,
			expectedDaemon: []string{},
			expectError:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fuseOptions, daemonOptions, err := splitFuseOptions(tt.options)
			assert.Equal(t, tt.expectError, err != nil)
			if tt.expectError {
				return
			}
			assert.ElementsMatch(t, tt.expectedFuse, fuseOptions)
			assert.ElementsMatch(t, tt.expectedDaemon, daemonOptions)
		})
	}
}
