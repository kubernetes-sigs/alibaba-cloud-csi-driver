package interceptors

import (
	"os"
	"path/filepath"
	"testing"

	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetParentDirAndBaseName(t *testing.T) {
	tests := []struct {
		name       string
		path       string
		wantParent string
		wantBase   string
	}{
		{
			name:       "simple path",
			path:       "/tmp/test/file",
			wantParent: "/tmp/test",
			wantBase:   "file",
		},
		{
			name:       "root path",
			path:       "/file",
			wantParent: "/",
			wantBase:   "file",
		},
		{
			name:       "relative path",
			path:       "test/file",
			wantParent: "test",
			wantBase:   "file",
		},
		{
			name:       "single filename",
			path:       "file",
			wantParent: ".",
			wantBase:   "file",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parent, base := getParentDirAndBaseName(tt.path)
			assert.Equal(t, tt.wantParent, parent)
			assert.Equal(t, tt.wantBase, base)
		})
	}
}

func TestResolveSymlinkTarget(t *testing.T) {
	tests := []struct {
		name         string
		symlinkPath  string
		linkTarget   string
		wantResolved string
	}{
		{
			name:         "absolute target",
			symlinkPath:  "/tmp/symlink",
			linkTarget:   "/absolute/target",
			wantResolved: "/absolute/target",
		},
		{
			name:         "relative target",
			symlinkPath:  "/tmp/symlink",
			linkTarget:   "relative/target",
			wantResolved: "/tmp/relative/target",
		},
		{
			name:         "relative target with parent",
			symlinkPath:  "/tmp/dir/symlink",
			linkTarget:   "../target",
			wantResolved: filepath.Join("/tmp/dir", "../target"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resolved := resolveSymlinkTarget(tt.symlinkPath, tt.linkTarget)
			assert.Equal(t, tt.wantResolved, resolved)
		})
	}
}

func TestGetTempSymlinkPath(t *testing.T) {
	tests := []struct {
		name      string
		parentDir string
		baseName  string
		want      string
	}{
		{
			name:      "simple case",
			parentDir: "/tmp",
			baseName:  "test",
			want:      "/tmp/.test.tmp",
		},
		{
			name:      "with subdirectory",
			parentDir: "/tmp/dir",
			baseName:  "file",
			want:      "/tmp/dir/.file.tmp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTempSymlinkPath(tt.parentDir, tt.baseName)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTempDataDirPath(t *testing.T) {
	tests := []struct {
		name      string
		parentDir string
		baseName  string
		timestamp int64
		want      string
	}{
		{
			name:      "simple case",
			parentDir: "/tmp",
			baseName:  "test",
			timestamp: 1234567890,
			want:      "/tmp/.test.tmp_1234567890",
		},
		{
			name:      "with subdirectory",
			parentDir: "/tmp/dir",
			baseName:  "file",
			timestamp: 9876543210,
			want:      "/tmp/dir/.file.tmp_9876543210",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTempDataDirPath(tt.parentDir, tt.baseName, tt.timestamp)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTempDataDirPathWithTimestamp(t *testing.T) {
	parentDir := "/tmp"
	baseName := "test"
	got := getTempDataDirPathWithTimestamp(parentDir, baseName)

	// Should contain the pattern
	assert.Contains(t, got, parentDir)
	assert.Contains(t, got, baseName)
	assert.Contains(t, got, ".tmp_")

	// Should have a timestamp (numeric suffix)
	require.Greater(t, len(got), len(parentDir)+len(baseName)+5, "should have timestamp suffix")
}

func TestGetTempDirPattern(t *testing.T) {
	tests := []struct {
		name     string
		baseName string
		want     string
	}{
		{
			name:     "simple case",
			baseName: "test",
			want:     ".test.tmp_",
		},
		{
			name:     "with special chars",
			baseName: "file-name",
			want:     ".file-name.tmp_",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTempDirPattern(tt.baseName)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTempPasswdFilePath(t *testing.T) {
	tests := []struct {
		name      string
		parentDir string
		baseName  string
		timestamp int64
		want      string
	}{
		{
			name:      "simple case",
			parentDir: "/tmp",
			baseName:  "passwd",
			timestamp: 1234567890,
			want:      "/tmp/.passwd.tmp.1234567890",
		},
		{
			name:      "with subdirectory",
			parentDir: "/tmp/dir",
			baseName:  "file",
			timestamp: 9876543210,
			want:      "/tmp/dir/.file.tmp.9876543210",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTempPasswdFilePath(tt.parentDir, tt.baseName, tt.timestamp)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTempPasswdFilePathWithTimestamp(t *testing.T) {
	parentDir := "/tmp"
	baseName := "passwd"
	got := getTempPasswdFilePathWithTimestamp(parentDir, baseName)

	// Should contain the pattern
	assert.Contains(t, got, parentDir)
	assert.Contains(t, got, baseName)
	assert.Contains(t, got, ".tmp.")

	// Should have a timestamp (numeric suffix)
	require.Greater(t, len(got), len(parentDir)+len(baseName)+5, "should have timestamp suffix")
}

func TestGetTempPasswdFilePattern(t *testing.T) {
	tests := []struct {
		name     string
		baseName string
		want     string
	}{
		{
			name:     "simple case",
			baseName: "passwd",
			want:     ".passwd.tmp.",
		},
		{
			name:     "with special chars",
			baseName: "file-name",
			want:     ".file-name.tmp.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTempPasswdFilePattern(tt.baseName)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTokenFilePath(t *testing.T) {
	tests := []struct {
		name    string
		dataDir string
		key     string
		want    string
	}{
		{
			name:    "simple case",
			dataDir: "/tmp/data",
			key:     "AccessKeyId",
			want:    "/tmp/data/AccessKeyId",
		},
		{
			name:    "with subdirectory",
			dataDir: "/tmp/dir/data",
			key:     "SecurityToken",
			want:    "/tmp/dir/data/SecurityToken",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getTokenFilePath(tt.dataDir, tt.key)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetRelativeDirName(t *testing.T) {
	tests := []struct {
		name    string
		dirPath string
		want    string
	}{
		{
			name:    "simple path",
			dirPath: "/tmp/test/dir",
			want:    "dir",
		},
		{
			name:    "root path",
			dirPath: "/dir",
			want:    "dir",
		},
		{
			name:    "relative path",
			dirPath: "test/dir",
			want:    "dir",
		},
		{
			name:    "single directory",
			dirPath: "dir",
			want:    "dir",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getRelativeDirName(tt.dirPath)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetTokenKeys(t *testing.T) {
	keys := getTokenKeys()

	// Should have 4 keys
	assert.Len(t, keys, 4)

	// Should contain all expected keys
	assert.Contains(t, keys, mounterutils.KeyAccessKeySecret)
	assert.Contains(t, keys, mounterutils.KeySecurityToken)
	assert.Contains(t, keys, mounterutils.KeyExpiration)
	assert.Contains(t, keys, mounterutils.KeyAccessKeyId)

	// AccessKeyId should be last
	assert.Equal(t, mounterutils.KeyAccessKeyId, keys[len(keys)-1], "AccessKeyId should be last")
}

func TestRemoveIgnoreNotExist(t *testing.T) {
	t.Run("remove existing file", func(t *testing.T) {
		tmpDir := t.TempDir()
		filePath := filepath.Join(tmpDir, "testfile")

		err := os.WriteFile(filePath, []byte("test"), 0o600)
		require.NoError(t, err)
		assert.FileExists(t, filePath)

		removeIgnoreNotExist(filePath)
		assert.NoFileExists(t, filePath)
	})

	t.Run("remove non-existent file - should not error", func(t *testing.T) {
		tmpDir := t.TempDir()
		filePath := filepath.Join(tmpDir, "nonexistent")

		// Should not panic or error
		removeIgnoreNotExist(filePath)
	})

	t.Run("remove existing directory", func(t *testing.T) {
		tmpDir := t.TempDir()
		dirPath := filepath.Join(tmpDir, "testdir")

		err := os.Mkdir(dirPath, 0o755)
		require.NoError(t, err)
		assert.DirExists(t, dirPath)

		removeIgnoreNotExist(dirPath)
		assert.NoDirExists(t, dirPath)
	})
}

func TestCheckFileContent(t *testing.T) {
	tests := []struct {
		name         string
		setupFile    bool
		fileContent  string
		checkContent string
		expectExists bool
		expectSame   bool
		expectErr    bool
		expectErrMsg string
	}{
		{
			name:         "file does not exist",
			setupFile:    false,
			checkContent: "test content",
			expectExists: false,
			expectSame:   false,
			expectErr:    false,
		},
		{
			name:         "file exists with same content",
			setupFile:    true,
			fileContent:  "test content",
			checkContent: "test content",
			expectExists: true,
			expectSame:   true,
			expectErr:    false,
		},
		{
			name:         "file exists with different content",
			setupFile:    true,
			fileContent:  "old content",
			checkContent: "new content",
			expectExists: true,
			expectSame:   false,
			expectErr:    false,
		},
		{
			name:         "file exists but empty, checking empty",
			setupFile:    true,
			fileContent:  "",
			checkContent: "",
			expectExists: true,
			expectSame:   true,
			expectErr:    false,
		},
		{
			name:         "file exists but empty, checking non-empty",
			setupFile:    true,
			fileContent:  "",
			checkContent: "non-empty",
			expectExists: true,
			expectSame:   false,
			expectErr:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			baseDir := t.TempDir()
			filePath := filepath.Join(baseDir, "testfile")

			if tt.setupFile {
				err := os.WriteFile(filePath, []byte(tt.fileContent), 0o600)
				require.NoError(t, err)
			}

			exists, sameContent, err := checkFileContent(filePath, []byte(tt.checkContent))

			if tt.expectErr {
				assert.Error(t, err)
				if tt.expectErrMsg != "" {
					assert.Contains(t, err.Error(), tt.expectErrMsg)
				}
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectExists, exists, "exists should match")
				assert.Equal(t, tt.expectSame, sameContent, "sameContent should match")
			}
		})
	}
}

func TestCleanupTokenDirectory(t *testing.T) {
	t.Run("cleanup directory with token files", func(t *testing.T) {
		tmpDir := t.TempDir()
		tokenDir := filepath.Join(tmpDir, "tokens")
		err := os.Mkdir(tokenDir, 0o755)
		require.NoError(t, err)

		// Create token files
		tokenKeys := getTokenKeys()
		for _, key := range tokenKeys {
			filePath := getTokenFilePath(tokenDir, key)
			err := os.WriteFile(filePath, []byte("test"), 0o600)
			require.NoError(t, err)
		}

		// Verify files exist
		for _, key := range tokenKeys {
			filePath := getTokenFilePath(tokenDir, key)
			assert.FileExists(t, filePath)
		}

		// Cleanup
		cleanupTokenDirectory(tokenDir, tokenKeys)

		// Verify files are removed
		for _, key := range tokenKeys {
			filePath := getTokenFilePath(tokenDir, key)
			assert.NoFileExists(t, filePath)
		}

		// Directory should be removed if empty
		assert.NoDirExists(t, tokenDir)
	})

	t.Run("cleanup empty directory", func(t *testing.T) {
		tmpDir := t.TempDir()
		tokenDir := filepath.Join(tmpDir, "tokens")
		err := os.Mkdir(tokenDir, 0o755)
		require.NoError(t, err)

		tokenKeys := getTokenKeys()
		cleanupTokenDirectory(tokenDir, tokenKeys)

		// Directory should be removed
		assert.NoDirExists(t, tokenDir)
	})

	t.Run("cleanup non-existent directory - should not error", func(t *testing.T) {
		tmpDir := t.TempDir()
		tokenDir := filepath.Join(tmpDir, "nonexistent")

		tokenKeys := getTokenKeys()
		cleanupTokenDirectory(tokenDir, tokenKeys)
		// Should not panic
	})
}

func TestCleanupTokenFiles(t *testing.T) {
	t.Run("cleanup token directory with symlink", func(t *testing.T) {
		tmpDir := t.TempDir()
		parentDir := filepath.Join(tmpDir, "parent")
		err := os.MkdirAll(parentDir, 0o755)
		require.NoError(t, err)

		// Create actual data directory
		dataDir := filepath.Join(parentDir, ".sts.tmp_123456")
		err = os.MkdirAll(dataDir, 0o755)
		require.NoError(t, err)

		// Create token files in data directory
		tokenKeys := getTokenKeys()
		for _, key := range tokenKeys {
			filePath := getTokenFilePath(dataDir, key)
			err := os.WriteFile(filePath, []byte("test"), 0o600)
			require.NoError(t, err)
		}

		// Create symlink
		tokenDir := filepath.Join(parentDir, "sts")
		relativeDataDir := getRelativeDirName(dataDir)
		err = os.Symlink(relativeDataDir, tokenDir)
		require.NoError(t, err)

		// Create temporary symlink
		tmpLinkPath := getTempSymlinkPath(parentDir, "sts")
		err = os.Symlink(relativeDataDir, tmpLinkPath)
		require.NoError(t, err)

		// Create old temporary directory
		oldDir := filepath.Join(parentDir, ".sts.tmp_999999")
		err = os.MkdirAll(oldDir, 0o755)
		require.NoError(t, err)

		// Cleanup
		cleanupTokenFiles(tokenDir)

		// Verify symlink is removed
		assert.NoFileExists(t, tokenDir)

		// Verify temporary symlink is removed
		assert.NoFileExists(t, tmpLinkPath)

		// Verify data directory and files are removed
		assert.NoDirExists(t, dataDir)

		// Verify old temporary directory is removed
		assert.NoDirExists(t, oldDir)
	})

	t.Run("cleanup empty token directory - should not error", func(t *testing.T) {
		cleanupTokenFiles("")
		// Should not panic
	})

	t.Run("cleanup non-existent token directory - should not error", func(t *testing.T) {
		tmpDir := t.TempDir()
		tokenDir := filepath.Join(tmpDir, "nonexistent")

		cleanupTokenFiles(tokenDir)
		// Should not panic
	})
}

func TestCleanupExpiredTempDirectories(t *testing.T) {
	t.Run("cleanup matching directories", func(t *testing.T) {
		tmpDir := t.TempDir()
		parentDir := filepath.Join(tmpDir, "parent")
		err := os.MkdirAll(parentDir, 0o755)
		require.NoError(t, err)

		// Create matching directories
		dir1 := filepath.Join(parentDir, ".test.tmp_123456")
		dir2 := filepath.Join(parentDir, ".test.tmp_789012")
		err = os.MkdirAll(dir1, 0o755)
		require.NoError(t, err)
		err = os.MkdirAll(dir2, 0o755)
		require.NoError(t, err)

		// Create token files in directories
		tokenKeys := getTokenKeys()
		for _, key := range tokenKeys {
			filePath1 := getTokenFilePath(dir1, key)
			filePath2 := getTokenFilePath(dir2, key)
			err := os.WriteFile(filePath1, []byte("test"), 0o600)
			require.NoError(t, err)
			err = os.WriteFile(filePath2, []byte("test"), 0o600)
			require.NoError(t, err)
		}

		// Create non-matching directory
		otherDir := filepath.Join(parentDir, ".other.tmp_123456")
		err = os.MkdirAll(otherDir, 0o755)
		require.NoError(t, err)

		// Cleanup
		pattern := getTempDirPattern("test")
		cleanupExpiredTempDirectories(parentDir, pattern, tokenKeys)

		// Verify matching directories are removed
		assert.NoDirExists(t, dir1)
		assert.NoDirExists(t, dir2)

		// Verify non-matching directory still exists
		assert.DirExists(t, otherDir)
	})

	t.Run("cleanup with non-existent parent - should not error", func(t *testing.T) {
		tmpDir := t.TempDir()
		parentDir := filepath.Join(tmpDir, "nonexistent")

		pattern := getTempDirPattern("test")
		cleanupExpiredTempDirectories(parentDir, pattern, getTokenKeys())
		// Should not panic
	})
}

func TestCleanupExpiredTempFiles(t *testing.T) {
	t.Run("cleanup matching files", func(t *testing.T) {
		tmpDir := t.TempDir()
		parentDir := filepath.Join(tmpDir, "parent")
		err := os.MkdirAll(parentDir, 0o755)
		require.NoError(t, err)

		// Create matching files
		file1 := filepath.Join(parentDir, ".passwd.tmp.123456")
		file2 := filepath.Join(parentDir, ".passwd.tmp.789012")
		err = os.WriteFile(file1, []byte("test1"), 0o600)
		require.NoError(t, err)
		err = os.WriteFile(file2, []byte("test2"), 0o600)
		require.NoError(t, err)

		// Create non-matching file
		otherFile := filepath.Join(parentDir, ".other.tmp.123456")
		err = os.WriteFile(otherFile, []byte("test"), 0o600)
		require.NoError(t, err)

		// Cleanup
		pattern := getTempPasswdFilePattern("passwd")
		cleanupExpiredTempFiles(parentDir, pattern)

		// Verify matching files are removed
		assert.NoFileExists(t, file1)
		assert.NoFileExists(t, file2)

		// Verify non-matching file still exists
		assert.FileExists(t, otherFile)
	})

	t.Run("cleanup with non-existent parent - should not error", func(t *testing.T) {
		tmpDir := t.TempDir()
		parentDir := filepath.Join(tmpDir, "nonexistent")

		pattern := getTempPasswdFilePattern("passwd")
		cleanupExpiredTempFiles(parentDir, pattern)
		// Should not panic
	})
}

func TestCleanupPasswdFile(t *testing.T) {
	t.Run("cleanup passwd file and temporary files", func(t *testing.T) {
		tmpDir := t.TempDir()
		parentDir := filepath.Join(tmpDir, "parent")
		err := os.MkdirAll(parentDir, 0o755)
		require.NoError(t, err)

		// Create passwd file
		passwdFile := filepath.Join(parentDir, "passwd")
		err = os.WriteFile(passwdFile, []byte("test"), 0o600)
		require.NoError(t, err)

		// Create temporary files
		tmpFile1 := filepath.Join(parentDir, ".passwd.tmp.123456")
		tmpFile2 := filepath.Join(parentDir, ".passwd.tmp.789012")
		err = os.WriteFile(tmpFile1, []byte("test1"), 0o600)
		require.NoError(t, err)
		err = os.WriteFile(tmpFile2, []byte("test2"), 0o600)
		require.NoError(t, err)

		// Create non-matching file
		otherFile := filepath.Join(parentDir, ".other.tmp.123456")
		err = os.WriteFile(otherFile, []byte("test"), 0o600)
		require.NoError(t, err)

		// Cleanup
		cleanupPasswdFile(passwdFile)

		// Verify passwd file is removed
		assert.NoFileExists(t, passwdFile)

		// Verify temporary files are removed
		assert.NoFileExists(t, tmpFile1)
		assert.NoFileExists(t, tmpFile2)

		// Verify non-matching file still exists
		assert.FileExists(t, otherFile)
	})

	t.Run("cleanup empty passwd file - should not error", func(t *testing.T) {
		cleanupPasswdFile("")
		// Should not panic
	})

	t.Run("cleanup non-existent passwd file - should not error", func(t *testing.T) {
		tmpDir := t.TempDir()
		passwdFile := filepath.Join(tmpDir, "nonexistent")

		cleanupPasswdFile(passwdFile)
		// Should not panic
	})
}
