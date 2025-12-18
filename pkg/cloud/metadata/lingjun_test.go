package metadata

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// LingjunTestCases contains multiple test cases for Lingjun metadata
var LingjunTestCases = map[string]string{
	"standard": `{
	"RegionId": "cn-hangzhou",
	"ZoneId": "cn-hangzhou-a",
	"NodeId": "i-1234567890abcdef0",
	"InstanceType": "ecs.g6.large",
	"EniMac": "00:11:22:33:44:55",
	"AckNicName": "eth1",
	"NimitzNetworkMode": "VPC",
	"NimitzInterface": ["eth0", "eth1"]
}`,
	"minimal": `{
	"RegionId": "cn-beijing",
	"ZoneId": "cn-beijing-c",
	"NodeId": "i-0987654321fedcba0",
	"InstanceType": "ecs.c5.large"
}`,
	"extended": `{
	"RegionId": "cn-shanghai",
	"ZoneId": "cn-shanghai-b",
	"NodeId": "i-1122334455aabbccdd",
	"InstanceType": "ecs.r5.xlarge",
	"EniMac": "aa:bb:cc:dd:ee:ff",
	"AckNicName": "eth0",
	"NimitzNetworkMode": "ENI",
	"NimitzInterface": ["eth0", "eth1", "eth2"]
}`,
}

func TestNewLingJunMetadata(t *testing.T) {
	t.Run("successful parsing - standard", func(t *testing.T) {
		// Setup test file
		tempDir := t.TempDir()
		lingjunConfigFile := filepath.Join(tempDir, "lingjun_config_standard")

		err := os.WriteFile(lingjunConfigFile, []byte(LingjunTestCases["standard"]), 0644)
		require.NoError(t, err)

		// Test the function
		result, err := NewLingJunMetadata(lingjunConfigFile)
		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, "cn-hangzhou", result.RegionId)
		assert.Equal(t, "cn-hangzhou-a", result.ZoneId)
		assert.Equal(t, "i-1234567890abcdef0", result.NodeId)
		assert.Equal(t, "ecs.g6.large", result.InstanceType)
	})

	t.Run("successful parsing - minimal", func(t *testing.T) {
		// Setup test file
		tempDir := t.TempDir()
		lingjunConfigFile := filepath.Join(tempDir, "lingjun_config_minimal")

		err := os.WriteFile(lingjunConfigFile, []byte(LingjunTestCases["minimal"]), 0644)
		require.NoError(t, err)

		// Test the function
		result, err := NewLingJunMetadata(lingjunConfigFile)
		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, "cn-beijing", result.RegionId)
		assert.Equal(t, "cn-beijing-c", result.ZoneId)
		assert.Equal(t, "i-0987654321fedcba0", result.NodeId)
		assert.Equal(t, "ecs.c5.large", result.InstanceType)
	})

	t.Run("successful parsing - extended", func(t *testing.T) {
		// Setup test file
		tempDir := t.TempDir()
		lingjunConfigFile := filepath.Join(tempDir, "lingjun_config_extended")

		err := os.WriteFile(lingjunConfigFile, []byte(LingjunTestCases["extended"]), 0644)
		require.NoError(t, err)

		// Test the function
		result, err := NewLingJunMetadata(lingjunConfigFile)
		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, "cn-shanghai", result.RegionId)
		assert.Equal(t, "cn-shanghai-b", result.ZoneId)
		assert.Equal(t, "i-1122334455aabbccdd", result.NodeId)
		assert.Equal(t, "ecs.r5.xlarge", result.InstanceType)
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

func TestLingjunMetaData_Get(t *testing.T) {
	// Create a temporary directory for test files
	tempDir := t.TempDir()
	lingjunConfigFile := filepath.Join(tempDir, "lingjun_config_get")

	// Setup test file
	err := os.WriteFile(lingjunConfigFile, []byte(LingjunTestCases["standard"]), 0644)
	require.NoError(t, err)

	// Get metadata instance
	metaData, err := NewLingJunMetadata(lingjunConfigFile)
	require.NoError(t, err)
	require.NotNil(t, metaData)

	tests := []struct {
		name        string
		key         MetadataKey
		expectedVal string
		expectErr   bool
	}{
		{
			name:        "get region id",
			key:         RegionID,
			expectedVal: "cn-hangzhou",
			expectErr:   false,
		},
		{
			name:        "get zone id",
			key:         ZoneID,
			expectedVal: "cn-hangzhou-a",
			expectErr:   false,
		},
		{
			name:        "get instance id",
			key:         InstanceID,
			expectedVal: "i-1234567890abcdef0",
			expectErr:   false,
		},
		{
			name:        "get instance type",
			key:         InstanceType,
			expectedVal: "ecs.g6.large",
			expectErr:   false,
		},
		{
			name:        "get unknown key",
			key:         AccountID,
			expectedVal: "",
			expectErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			value, err := metaData.Get(tt.key)
			if tt.expectErr {
				assert.Error(t, err)
				assert.Equal(t, "", value)
				assert.Equal(t, ErrUnknownMetadataKey, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedVal, value)
			}
		})
	}
}
