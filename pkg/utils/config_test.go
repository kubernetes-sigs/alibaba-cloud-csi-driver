package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBool(t *testing.T) {
	tests := []struct {
		arg      string
		expected bool
		wantErr  bool
	}{
		{
			arg:      "enable",
			expected: true,
			wantErr:  false,
		},
		{
			arg:      "no",
			expected: false,
			wantErr:  false,
		},
		{
			arg:      "t",
			expected: true,
			wantErr:  false,
		},
		{
			arg:      "test",
			expected: false,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.arg, func(t *testing.T) {
			dir := t.TempDir()
			assert.NoError(t, os.WriteFile(dir+"/test-config", []byte(tt.arg), 0644))

			cfg := Config{Path: dir}
			if tt.wantErr {
				assert.Panics(t, func() {
					cfg.GetBool("test-config", "", !tt.expected)
				})
			} else {
				actual := cfg.GetBool("test-config", "", !tt.expected)
				assert.Equal(t, tt.expected, actual)
			}
		})
	}
}

func TestDefault(t *testing.T) {
	dir := t.TempDir()
	cfg := Config{Path: dir}

	assert.Equal(t, "default-config", cfg.Get("test-config", "TEST_CONFIG_IN_UT", "default-config"))
}
