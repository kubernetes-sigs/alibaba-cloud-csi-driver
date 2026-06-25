package ossfs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestApplyOptionDefaults(t *testing.T) {
	t.Run("no CA file present, options unchanged", func(t *testing.T) {
		d := &Driver{}
		options := []string{"url=oss.aliyuncs.com", "allow_other"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"url=oss.aliyuncs.com", "allow_other"}, result)
	})

	t.Run("nil options, no CA file", func(t *testing.T) {
		d := &Driver{}
		result := d.ApplyOptionDefaults(nil)
		assert.Nil(t, result)
	})

	t.Run("user already specified agent_identity_ca_file", func(t *testing.T) {
		d := &Driver{}
		options := []string{"agent_identity_ca_file=/custom/path/ca.crt", "url=oss.aliyuncs.com"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"agent_identity_ca_file=/custom/path/ca.crt", "url=oss.aliyuncs.com"}, result)
	})

	t.Run("CA file present and readable, appends option", func(t *testing.T) {
		tmpDir := t.TempDir()
		caFile := filepath.Join(tmpDir, "ca.crt")
		require.NoError(t, os.WriteFile(caFile, []byte("fake-ca"), 0644))

		t.Setenv("SSL_CERT_FILE", caFile)
		d := &Driver{}
		options := []string{"url=oss.aliyuncs.com"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"url=oss.aliyuncs.com", "agent_identity_ca_file=" + caFile}, result)
	})

	t.Run("CA file not readable, options unchanged", func(t *testing.T) {
		if os.Getuid() == 0 {
			t.Skip("root bypasses file permission checks")
		}
		tmpDir := t.TempDir()
		caFile := filepath.Join(tmpDir, "ca.crt")
		require.NoError(t, os.WriteFile(caFile, []byte("fake-ca"), 0000))

		t.Setenv("SSL_CERT_FILE", caFile)
		d := &Driver{}
		options := []string{"url=oss.aliyuncs.com"}
		result := d.ApplyOptionDefaults(options)
		assert.Equal(t, []string{"url=oss.aliyuncs.com"}, result)
	})
}
