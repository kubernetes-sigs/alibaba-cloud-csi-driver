package metadata

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// LingjunTestCases contains multiple test cases for Lingjun metadata
const LingjunConfigJson = `{
    "RegionId":"cn-wulanchabu",
    "ZoneId":"cn-wulanchabu-a",
    "NodeId":"e01-cn-xxxxxx",
    "InstanceType":"efg2.C48eNH3ebn",
    "HyperNodeInstanceType":"",
    "EniMac":"00:16:3e:0f:00:00",
    "AckNicName":"eth0",
    "NimitzNetworkMode":"Virtio",
    "HyperNodeId":"",
    "NimitzInterface":["LENI"]
    }`

func TestNewLingJunMetadata(t *testing.T) {
	t.Run("successful parsing", func(t *testing.T) {
		// Setup test file
		tempDir := t.TempDir()
		lingjunConfigFile := filepath.Join(tempDir, "lingjun_config")

		require.NoError(t, os.WriteFile(lingjunConfigFile, []byte(LingjunConfigJson), 0644))

		// Test the function
		result, err := NewLingJunMetadata(lingjunConfigFile)
		require.NoError(t, err)
		require.NotNil(t, result)

		m := testMetadata(t, result)
		assert.Equal(t, "cn-wulanchabu", MustGet(m, RegionID))
		assert.Equal(t, "cn-wulanchabu-a", MustGet(m, ZoneID))
		assert.Equal(t, "e01-cn-xxxxxx", MustGet(m, InstanceID))
		assert.Equal(t, "efg2.C48eNH3ebn", MustGet(m, InstanceType))
		kind, err := m.MachineKind()
		assert.NoError(t, err)
		assert.Equal(t, MachineKindLingjun, kind)

		_, err = m.Get(AccountID)
		assert.ErrorIs(t, err, ErrUnknownMetadataKey)
	})

	t.Run("file not exist", func(t *testing.T) {
		// Setup non-existent file
		tempDir := t.TempDir()
		lingjunConfigFile := filepath.Join(tempDir, "non_existent_file")

		result, err := NewLingJunMetadata(lingjunConfigFile)
		assert.Error(t, err)
		assert.True(t, os.IsNotExist(err))
		assert.Nil(t, result)
	})

	t.Run("other file read error", func(t *testing.T) {
		// Test with a directory path instead of a file
		tempDir := t.TempDir()

		result, err := NewLingJunMetadata(tempDir) // Pass directory instead of file
		assert.Error(t, err)
		assert.False(t, os.IsNotExist(err)) // Not a "not exist" error
		assert.Nil(t, result)
	})

	t.Run("invalid json data", func(t *testing.T) {
		// Setup test file with invalid JSON
		tempDir := t.TempDir()
		lingjunConfigFile := filepath.Join(tempDir, "lingjun_config_invalid")

		err := os.WriteFile(lingjunConfigFile, []byte("{ invalid json }"), 0644)
		require.NoError(t, err)

		result, err := NewLingJunMetadata(lingjunConfigFile)
		assert.Error(t, err)
		assert.Nil(t, result)
	})
}
