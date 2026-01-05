package interceptors

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
)

var _ mounter.MountInterceptor = OssfsSecretInterceptor

// rawMounter is a shared mount.Interface instance for checking mount points.
// It's safe to reuse because mount.Interface implementations are stateless.
// We use a separate instance from ossfs_monitor.go to avoid import cycles,
// but both are created the same way and can be used interchangeably.
var rawMounter = mountutils.NewWithoutSystemd("")

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
		filePath := filepath.Join(dir, key)
		removeIgnoreNotExist(filePath)
	}
	// Try to remove the directory if it's empty
	removeIgnoreNotExist(dir)
}

func OssfsSecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	return ossfsSecretInterceptor(ctx, op, handler, mounterutils.OssFsType)
}

func Ossfs2SecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	return ossfsSecretInterceptor(ctx, op, handler, mounterutils.OssFs2Type)
}

func ossfsSecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler, fuseType string) error {
	return ossfsSecretInterceptorWithMounter(ctx, op, handler, fuseType, rawMounter)
}

// ossfsSecretInterceptorWithMounter is the internal implementation that accepts a mounter parameter.
// This allows tests to inject a fake mounter to simulate different mount point states.
func ossfsSecretInterceptorWithMounter(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler, fuseType string, mountInterface mountutils.Interface) error {
	if op == nil || op.Secrets == nil {
		return handler(ctx, op)
	}

	passwdFile, tokenDir, err := prepareCredentialFiles(fuseType, op.Target, op.Secrets)
	if err != nil {
		return fmt.Errorf("prepare credential files failed: %w", err)
	}

	// Check if mount point already exists for token rotation scenario.
	// In token rotation (republish), we only need to update the token files,
	// and the running ossfs client will automatically reload the new token.
	// We skip the mount operation to avoid creating duplicate mount points.
	//
	// Note: IsNotMountPoint handles path not existing by creating the directory
	// and returning (true, nil). If it returns an error, it's a real error
	// (e.g., permission denied, mkdir failed, unmount failed) that should be returned.
	// Only check mount point if mountInterface is available (not nil).
	if mountInterface != nil {
		notMnt, err := mounterutils.IsNotMountPoint(mountInterface, op.Target)
		if err != nil {
			return fmt.Errorf("failed to check if target %s is a mountpoint: %w", op.Target, err)
		}
		if !notMnt {
			// Mount point already exists, this is a token rotation scenario.
			// Token files have already been updated by rotateTokenFiles above.
			// Skip mount operation and let the existing ossfs client reload the new token.
			klog.V(4).InfoS("mount point already exists, skipping mount for token rotation", "target", op.Target)
			return mounter.ErrSkipMount
		}
	}

	if passwdFile != "" {
		klog.V(4).InfoS("created ossfs passwd file", "path", passwdFile)
		if fuseType == mounterutils.OssFsType {
			op.Options = append(op.Options, "passwd_file="+passwdFile)
		} else {
			// ossfs2
			op.Args = append(op.Args, []string{"-c", passwdFile}...)
		}
	}
	if tokenDir != "" {
		klog.V(4).InfoS("created ossfs token directory", "dir", tokenDir)
		if fuseType == mounterutils.OssFsType {
			op.Options = append(op.Options, "passwd_file="+tokenDir)
		} else {
			// ossfs2
			// For ossfs2, file-path is a common option configuration after -o, so append to op.Options
			op.Options = append(op.Options,
				fmt.Sprintf("oss_sts_multi_conf_ak_file=%s", filepath.Join(tokenDir, mounterutils.KeyAccessKeyId)),
				fmt.Sprintf("oss_sts_multi_conf_sk_file=%s", filepath.Join(tokenDir, mounterutils.KeyAccessKeySecret)),
				fmt.Sprintf("oss_sts_multi_conf_token_file=%s", filepath.Join(tokenDir, mounterutils.KeySecurityToken)),
			)
		}
	}

	if err = handler(ctx, op); err != nil {
		return err
	}

	if (passwdFile == "" && tokenDir == "") || op.MountResult == nil {
		return nil
	}
	result, ok := op.MountResult.(server.OssfsMountResult)
	if !ok {
		klog.ErrorS(
			errors.New("failed to assert ossfs mount result"),
			"skipping cleanup of passwd file", "mountpoint", op.Target, "path", passwdFile,
		)
		return nil
	}

	go func() {
		<-result.ExitChan
		removeIgnoreNotExist(passwdFile)
	}()
	return nil
}

// rotateTokenFiles rotates (or initializes) token files using directory-level symlink for atomic updates.
// This function uses a simple symlink-based approach similar to Kubernetes secret volume plugin:
// 1. Create a temporary data directory (e.g., .dir.tmp_<timestamp>)
// 2. Write all token files to the temporary directory
// 3. Atomically switch the dir symlink to point to the new directory
// This ensures all files are updated atomically - readers either see all old files or all new files.
// Clients that have already opened file handles will continue to read from the old directory
// until they close and reopen, ensuring no interruption during rotation.
func rotateTokenFiles(dir string, secrets map[string]string) (rotated bool, err error) {
	// Currently, for ossfs2, expiration is not required.
	// But we still manage it (if offered) for the feature.
	tokenKey := []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration}
	// Check nil value in advanced.
	for _, key := range tokenKey {
		val := secrets[key]
		if val == "" {
			if key == mounterutils.KeyExpiration {
				continue
			}
			err = fmt.Errorf("invalid authorization. %s is empty", key)
			klog.Error(err)
			// Return false for rotated when error occurs
			return false, err
		}
	}

	// Check if any file needs update before writing
	// Read current data directory if symlink exists
	currentDataDir := ""
	if linkTarget, readErr := os.Readlink(dir); readErr == nil {
		// dir is a symlink, resolve to actual directory
		if filepath.IsAbs(linkTarget) {
			currentDataDir = linkTarget
		} else {
			currentDataDir = filepath.Join(filepath.Dir(dir), linkTarget)
		}
	}
	// Note: dir should always be a symlink path that doesn't exist on first call,
	// so we don't need to handle the case where dir is a regular directory

	anyNeedsUpdate := false
	if currentDataDir == "" {
		// No existing data directory, need to create
		anyNeedsUpdate = true
	} else {
		// Check if any file content changed
		for _, key := range tokenKey {
			val := secrets[key]
			if val == "" {
				continue
			}
			filePath := filepath.Join(currentDataDir, key)
			exists, sameContent, checkErr := checkFileContent(filePath, []byte(val))
			if checkErr != nil || !exists || !sameContent {
				// File doesn't exist, error reading, or content is different, needs update
				anyNeedsUpdate = true
				break
			}
		}
	}

	// If no files need update, return early
	if !anyNeedsUpdate {
		return false, nil
	}

	// Create temporary data directory with timestamp suffix in the parent directory
	// This ensures the symlink target is in the same filesystem for atomic rename
	parentDir := filepath.Dir(dir)
	baseName := filepath.Base(dir)
	tmpDataDir := filepath.Join(parentDir, fmt.Sprintf(".%s.tmp_%d", baseName, time.Now().UnixNano()))
	if err = os.MkdirAll(tmpDataDir, 0o755); err != nil {
		return false, fmt.Errorf("failed to create temporary data directory: %w", err)
	}

	// Use defer to ensure temporary directory is cleaned up on any error or panic
	defer func() {
		if !rotated {
			// Clean up temporary directory on error
			cleanupTokenDirectory(tmpDataDir, tokenKey)
		}
	}()

	// Write all token files to temporary directory
	for _, key := range tokenKey {
		val := secrets[key]
		if val == "" {
			continue
		}
		filePath := filepath.Join(tmpDataDir, key)
		if err = os.WriteFile(filePath, []byte(val), 0o600); err != nil {
			return false, fmt.Errorf("failed to write token file %s: %w", key, err)
		}
	}

	// Atomically switch the dir symlink
	// First, create a temporary symlink pointing to the new directory
	tmpLinkPath := filepath.Join(parentDir, fmt.Sprintf(".%s.tmp", baseName))
	relativeDataDir := filepath.Base(tmpDataDir)
	if err = os.Symlink(relativeDataDir, tmpLinkPath); err != nil {
		return false, fmt.Errorf("failed to create temporary symlink: %w", err)
	}

	// Use defer to ensure temporary symlink is cleaned up on any error or panic
	// After successful rename, the symlink will be renamed to dir, so removeIgnoreNotExist will safely ignore it
	defer removeIgnoreNotExist(tmpLinkPath)

	// Atomically rename the temporary symlink to dir
	// Note: dir should not exist on first call, so this will create the symlink
	// On subsequent calls, this will atomically replace the existing symlink
	// This is atomic on most filesystems and ensures all readers see either all old files or all new files
	// IMPORTANT: Once this rename completes, all new opens of dir/ will resolve to the new directory.
	// However, clients that have already opened file handles will continue to read from the old directory
	// (pointed to by their open file descriptors), ensuring no interruption during rotation.
	if err = os.Rename(tmpLinkPath, dir); err != nil {
		return false, fmt.Errorf("failed to atomically switch symlink: %w", err)
	}

	// Mark as rotated so defer won't clean up the new directory
	rotated = true

	// Clean up old data directory if it exists and is different from the new one
	// Only remove the token files we know about, not the entire directory
	if currentDataDir != "" && currentDataDir != tmpDataDir {
		cleanupTokenDirectory(currentDataDir, tokenKey)
	}

	return true, nil
}

// rotatePasswdFile rotates (or initializes) a passwd file atomically.
// It first checks if the file exists and has the same content to avoid redundant writes.
// If content is different or file doesn't exist, it writes to a temporary file
// and atomically renames it to the target file. This avoids the need for locks
// and ensures atomic updates even with concurrent writes.
func rotatePasswdFile(path string, data []byte, perm os.FileMode) (done bool, err error) {
	// First check if file exists and has the same content
	exists, sameContent, err := checkFileContent(path, data)
	if err != nil {
		return false, err
	}
	if exists && sameContent {
		return false, nil
	}

	// Content is different or file doesn't exist, need to write
	// Create temporary file in the same directory
	dir := filepath.Dir(path)
	tmpFile := filepath.Join(dir, fmt.Sprintf(".%s.tmp.%d", filepath.Base(path), time.Now().UnixNano()))

	// Use defer to ensure temporary file is cleaned up on any error or panic
	defer removeIgnoreNotExist(tmpFile)

	// Write to temporary file
	if err = os.WriteFile(tmpFile, data, perm); err != nil {
		return false, fmt.Errorf("failed to write temporary file: %w", err)
	}

	// Atomically rename temporary file to target file
	// This is atomic on most filesystems and handles concurrent writes gracefully
	if err = os.Rename(tmpFile, path); err != nil {
		return false, fmt.Errorf("failed to atomically replace file: %w", err)
	}

	return true, nil
}

// prepareCredentialFiles returns:
//  1. file:    path of ossfs credential file for fixed AKSK
//  2. dir:     directory of ossfs credential files for token
//  4. error
func prepareCredentialFiles(fuseType, target string, secrets map[string]string) (file, dir string, err error) {
	// fixed credientials
	hashDir := mounterutils.GetPasswdHashDir(target)

	if passwd := secrets[mounterutils.GetPasswdFileName(fuseType)]; passwd != "" {
		err = os.MkdirAll(hashDir, 0o755)
		if err != nil {
			klog.Errorf("mkdirall hashdir failed %v", err)
			return
		}
		filePath := filepath.Join(hashDir, mounterutils.GetPasswdFileName(fuseType))
		_, err = rotatePasswdFile(filePath, []byte(passwd), 0o600)
		if err != nil {
			return
		}
		file = filePath
		return
	}

	// token
	if token := secrets[mounterutils.KeySecurityToken]; token != "" {
		tokenDir := filepath.Join(hashDir, mounterutils.GetPasswdFileName(fuseType))
		err = os.MkdirAll(tokenDir, 0o755)
		if err != nil {
			klog.Errorf("mkdirall tokenDir failed %v", err)
			return
		}
		// Use tokenDir/sts as the symlink path to ensure it doesn't exist on first call
		// This avoids the need to handle directory-to-symlink conversion
		stsDir := filepath.Join(tokenDir, "sts")
		_, err = rotateTokenFiles(stsDir, secrets)
		if err != nil {
			return
		}
		dir = stsDir
		return
	}
	return
}
