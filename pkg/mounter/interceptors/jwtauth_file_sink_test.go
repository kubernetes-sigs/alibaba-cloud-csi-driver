package interceptors

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func readSinkFile(t *testing.T, s *jwtAuthFileSink, key string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(s.Dir(), key))
	require.NoError(t, err)
	return string(data)
}

func TestJWTAuthFileSink_ApplyAndRotate(t *testing.T) {
	outputDir := filepath.Join(t.TempDir(), "creds")
	sink := newJWTAuthFileSink(outputDir)
	assert.Equal(t, filepath.Join(outputDir, credentialDataDir), sink.Dir())

	require.NoError(t, sink.Apply(&jwtauth.STSToken{
		AccessKeyID: "ak1", AccessKeySecret: "sk1", SecurityToken: "st1", Expiration: "exp1",
	}))
	assert.Equal(t, "ak1", readSinkFile(t, sink, mounterutils.KeyAccessKeyId))
	assert.Equal(t, "sk1", readSinkFile(t, sink, mounterutils.KeyAccessKeySecret))
	assert.Equal(t, "st1", readSinkFile(t, sink, mounterutils.KeySecurityToken))
	assert.Equal(t, "exp1", readSinkFile(t, sink, mounterutils.KeyExpiration))

	// Rotation replaces the whole set atomically via the symlink swap.
	require.NoError(t, sink.Apply(&jwtauth.STSToken{
		AccessKeyID: "ak2", AccessKeySecret: "sk2", SecurityToken: "st2", Expiration: "exp2",
	}))
	assert.Equal(t, "ak2", readSinkFile(t, sink, mounterutils.KeyAccessKeyId))
	assert.Equal(t, "st2", readSinkFile(t, sink, mounterutils.KeySecurityToken))
}

func TestJWTAuthFileSink_ApplyRejectsIncompleteCredential(t *testing.T) {
	sink := newJWTAuthFileSink(filepath.Join(t.TempDir(), "creds"))
	err := sink.Apply(&jwtauth.STSToken{AccessKeyID: "ak"})
	require.Error(t, err)
	assert.Contains(t, err.Error(), "rotate credential files")
}

func TestJWTAuthFileSink_CleanupRemovesEverything(t *testing.T) {
	outputDir := filepath.Join(t.TempDir(), "creds")
	sink := newJWTAuthFileSink(outputDir)
	require.NoError(t, sink.Apply(&jwtauth.STSToken{
		AccessKeyID: "ak", AccessKeySecret: "sk", SecurityToken: "st", Expiration: "exp",
	}))

	sink.Cleanup()

	_, err := os.Stat(outputDir)
	assert.True(t, os.IsNotExist(err), "output dir should be removed by Cleanup")

	// Cleanup is safe to call again on an already-removed directory.
	sink.Cleanup()
}
