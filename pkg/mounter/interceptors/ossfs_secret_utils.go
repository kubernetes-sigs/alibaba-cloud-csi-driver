package interceptors

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
)

// getParentDirAndBaseName splits a path into its parent directory and base name.
func getParentDirAndBaseName(path string) (parentDir, baseName string) {
	return filepath.Dir(path), filepath.Base(path)
}

// resolveSymlinkTarget resolves a symlink target to an absolute path.
// If linkTarget is already absolute, it returns it as is.
// Otherwise, it joins it with the parent directory of symlinkPath.
func resolveSymlinkTarget(symlinkPath, linkTarget string) string {
	if filepath.IsAbs(linkTarget) {
		return linkTarget
	}
	return filepath.Join(filepath.Dir(symlinkPath), linkTarget)
}

// getTempSymlinkPath returns the path for a temporary symlink used during atomic rotation.
// Format: <parentDir>/.<baseName>.tmp
func getTempSymlinkPath(parentDir, baseName string) string {
	return filepath.Join(parentDir, fmt.Sprintf(".%s.tmp", baseName))
}

// getTempDataDirPath returns the path for a temporary data directory used during token rotation.
// Format: <parentDir>/.<baseName>.tmp_<timestamp>
func getTempDataDirPath(parentDir, baseName string, timestamp int64) string {
	return filepath.Join(parentDir, fmt.Sprintf(".%s.tmp_%d", baseName, timestamp))
}

// getTempDataDirPathWithTimestamp returns the path for a temporary data directory with current timestamp.
func getTempDataDirPathWithTimestamp(parentDir, baseName string) string {
	return getTempDataDirPath(parentDir, baseName, time.Now().UnixNano())
}

// getTempDirPattern returns the pattern prefix for matching temporary data directories.
// Format: .<baseName>.tmp_
func getTempDirPattern(baseName string) string {
	return fmt.Sprintf(".%s.tmp_", baseName)
}

// getTempPasswdFilePath returns the path for a temporary passwd file used during atomic rotation.
// Format: <parentDir>/.<baseName>.tmp.<timestamp>
func getTempPasswdFilePath(parentDir, baseName string, timestamp int64) string {
	return filepath.Join(parentDir, fmt.Sprintf(".%s.tmp.%d", baseName, timestamp))
}

// getTempPasswdFilePathWithTimestamp returns the path for a temporary passwd file with current timestamp.
func getTempPasswdFilePathWithTimestamp(parentDir, baseName string) string {
	return getTempPasswdFilePath(parentDir, baseName, time.Now().UnixNano())
}

// getTempPasswdFilePattern returns the pattern prefix for matching temporary passwd files.
// Format: .<baseName>.tmp.
func getTempPasswdFilePattern(baseName string) string {
	return fmt.Sprintf(".%s.tmp.", baseName)
}

// getTokenFilePath returns the full path to a token file within a data directory.
func getTokenFilePath(dataDir, key string) string {
	return filepath.Join(dataDir, key)
}

// getRelativeDirName returns the base name of a directory path, used for creating relative symlinks.
func getRelativeDirName(dirPath string) string {
	return filepath.Base(dirPath)
}

// getTokenKeys returns the list of token keys in the order they should be processed.
//
// TODO: In ossfs and ossfs2, the client checks if AccessKeyId has changed to determine
// whether to update credential information. Therefore, AccessKeyId must be rotated last
// to ensure atomic updates - all other files are updated before AccessKeyId, so clients
// will see either all old files or all new files, never a mixed state.
func getTokenKeys() []string {
	return []string{mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration, mounterutils.KeyAccessKeyId}
}

// removeIgnoreNotExist removes a file or directory, ignoring "not exist" errors.
// Other errors are logged but not returned. This is useful for cleanup operations
// where the file may or may not exist.
func removeIgnoreNotExist(path string) {
	if err := os.Remove(path); err != nil {
		if !os.IsNotExist(err) {
			klog.V(4).Infof("failed to remove %s: %v", path, err)
		}
	}
}

// checkFileContent checks if a file exists and has the same content as the given data.
// Returns (exists, sameContent, error).
func checkFileContent(path string, data []byte) (exists bool, sameContent bool, err error) {
	existingData, readErr := os.ReadFile(path)
	if readErr != nil {
		if os.IsNotExist(readErr) {
			return false, false, nil
		}
		return false, false, readErr
	}
	return true, string(existingData) == string(data), nil
}

// cleanupTokenDirectory removes all token files from the given directory and then tries to remove the directory itself.
// This is used for cleaning up temporary directories on error or old data directories after rotation.
func cleanupTokenDirectory(dir string, tokenKeys []string) {
	for _, key := range tokenKeys {
		filePath := getTokenFilePath(dir, key)
		removeIgnoreNotExist(filePath)
	}
	// Try to remove the directory if it's empty
	removeIgnoreNotExist(dir)
}

// cleanupTokenFiles cleans up the token directory symlink and all associated files/directories.
// It resolves the symlink to get the actual data directory, cleans up token files in that directory,
// removes the symlink itself, and also cleans up any temporary symlinks or old directories.
// Since this is called when the client has exited, all files should be cleaned up.
func cleanupTokenFiles(tokenDir string) {
	if tokenDir == "" {
		return
	}

	parentDir, baseName := getParentDirAndBaseName(tokenDir)
	tokenKeys := getTokenKeys()

	// Clean up the actual data directory if symlink exists
	if linkTarget, err := os.Readlink(tokenDir); err == nil {
		// tokenDir is a symlink, resolve to actual directory
		actualDataDir := resolveSymlinkTarget(tokenDir, linkTarget)
		// Clean up token files in the actual directory
		cleanupTokenDirectory(actualDataDir, tokenKeys)
	}

	// Remove the symlink itself
	removeIgnoreNotExist(tokenDir)

	// Clean up temporary symlink if it exists
	tmpLinkPath := getTempSymlinkPath(parentDir, baseName)
	removeIgnoreNotExist(tmpLinkPath)

	// Clean up any old temporary directories matching the pattern returned by getTempDirPattern
	// This handles cases where old directories weren't cleaned up properly
	cleanupExpiredTempDirectories(parentDir, getTempDirPattern(baseName), tokenKeys)
}

// cleanupExpiredTempDirectories cleans up all directories matching the given pattern in the parent directory.
func cleanupExpiredTempDirectories(parentDir, pattern string, tokenKeys []string) {
	entries, err := os.ReadDir(parentDir)
	if err != nil {
		return
	}
	for _, entry := range entries {
		if entry.IsDir() && strings.HasPrefix(entry.Name(), pattern) {
			oldDir := filepath.Join(parentDir, entry.Name())
			cleanupTokenDirectory(oldDir, tokenKeys)
		}
	}
}

// cleanupExpiredTempFiles cleans up all files matching the given pattern in the parent directory.
func cleanupExpiredTempFiles(parentDir, pattern string) {
	entries, err := os.ReadDir(parentDir)
	if err != nil {
		return
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), pattern) {
			tmpFile := filepath.Join(parentDir, entry.Name())
			removeIgnoreNotExist(tmpFile)
		}
	}
}

// cleanupPasswdFile cleans up the passwd file and all temporary files matching the pattern.
func cleanupPasswdFile(passwdFile string) {
	if passwdFile == "" {
		return
	}

	// Remove the passwd file itself
	removeIgnoreNotExist(passwdFile)

	// Clean up any temporary files matching the pattern returned by getTempPasswdFilePattern
	parentDir, baseName := getParentDirAndBaseName(passwdFile)
	cleanupExpiredTempFiles(parentDir, getTempPasswdFilePattern(baseName))
}
