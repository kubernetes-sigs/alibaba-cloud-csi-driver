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
	commonutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
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

	if passwdFile == "" || op.MountResult == nil {
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

// rotateTokenFiles rotates (or initializes) token files using symlink for atomic updates.
// This function uses a symlink-based approach similar to Kubernetes configmap volume plugin:
// 1. Create a temporary data directory (e.g., ..data_tmp_<timestamp>)
// 2. Write all token files to the temporary directory
// 3. Atomically switch the ..data symlink to point to the new directory
// 4. Create symlinks for each token file pointing to ..data/<filename>
// This ensures all files are updated atomically - readers either see all old files or all new files.
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
	dataLinkPath := filepath.Join(dir, "..data")
	currentDataDir := ""
	if linkTarget, readErr := os.Readlink(dataLinkPath); readErr == nil {
		currentDataDir = filepath.Join(dir, linkTarget)
	}

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
			existingData, readErr := os.ReadFile(filePath)
			if readErr != nil {
				// File doesn't exist, needs to be created
				anyNeedsUpdate = true
				break
			}
			if string(existingData) != val {
				// File exists but content is different, needs update
				anyNeedsUpdate = true
				break
			}
		}
	}

	// If no files need update, return early
	if !anyNeedsUpdate {
		return false, nil
	}

	// Create temporary data directory with timestamp suffix
	tmpDataDir := filepath.Join(dir, fmt.Sprintf("..data_tmp_%d", time.Now().UnixNano()))
	if err = os.MkdirAll(tmpDataDir, 0o755); err != nil {
		return false, fmt.Errorf("failed to create temporary data directory: %w", err)
	}

	// Helper function to clean up temporary directory by removing only the files we created
	cleanupTmpDir := func() {
		for _, key := range tokenKey {
			filePath := filepath.Join(tmpDataDir, key)
			removeIgnoreNotExist(filePath)
		}
		// Try to remove the directory if it's empty
		removeIgnoreNotExist(tmpDataDir)
	}

	// Write all token files to temporary directory
	for _, key := range tokenKey {
		val := secrets[key]
		if val == "" {
			continue
		}
		filePath := filepath.Join(tmpDataDir, key)
		if err = commonutils.WriteAndSyncFile(filePath, []byte(val), 0o600); err != nil {
			// Clean up temporary directory on error
			cleanupTmpDir()
			return false, fmt.Errorf("failed to write token file %s: %w", key, err)
		}
	}

	// Atomically switch the ..data symlink
	// First, create a temporary symlink
	tmpLinkPath := filepath.Join(dir, "..data_tmp")
	relativeDataDir := filepath.Base(tmpDataDir)
	if err = os.Symlink(relativeDataDir, tmpLinkPath); err != nil {
		cleanupTmpDir()
		return false, fmt.Errorf("failed to create temporary symlink: %w", err)
	}

	// Atomically rename the temporary symlink to ..data
	// This is atomic on most filesystems
	if err = os.Rename(tmpLinkPath, dataLinkPath); err != nil {
		removeIgnoreNotExist(tmpLinkPath)
		cleanupTmpDir()
		return false, fmt.Errorf("failed to atomically switch symlink: %w", err)
	}

	// Create symlinks for each token file pointing to ..data/<filename>
	// This ensures external references to tokenDir/<filename> work correctly
	// Use atomic replacement to avoid race condition: create temp symlink first, then rename
	for _, key := range tokenKey {
		val := secrets[key]
		if val == "" {
			continue
		}
		fileLinkPath := filepath.Join(dir, key)
		tmpFileLinkPath := filepath.Join(dir, fmt.Sprintf(".%s.tmp", key))

		// Create symlink pointing to ..data/<filename> with temporary name
		// Use forward slash for symlink target (standard for symlinks)
		linkTarget := fmt.Sprintf("..data/%s", key)
		if err = os.Symlink(linkTarget, tmpFileLinkPath); err != nil {
			klog.Errorf("failed to create temporary symlink for %s: %v", key, err)
			// Continue with other files
			continue
		}

		// Atomically replace the old symlink with the new one
		// This ensures the file is always accessible during the update
		if err = os.Rename(tmpFileLinkPath, fileLinkPath); err != nil {
			removeIgnoreNotExist(tmpFileLinkPath)
			klog.Errorf("failed to atomically replace symlink for %s: %v", key, err)
			// Continue with other files
		}
	}

	// Clean up old data directory if it exists
	// Only remove the token files we know about, not the entire directory
	if currentDataDir != "" && currentDataDir != tmpDataDir {
		// Remove old data directory asynchronously to avoid blocking
		go func() {
			for _, key := range tokenKey {
				filePath := filepath.Join(currentDataDir, key)
				removeIgnoreNotExist(filePath)
			}
			// Try to remove the directory if it's empty
			removeIgnoreNotExist(currentDataDir)
		}()
	}

	rotated = true
	return
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
		_, err = mounterutils.WriteFileWithLock(filepath.Join(hashDir, mounterutils.GetPasswdFileName(fuseType)), []byte(passwd), 0o600)
		if err != nil {
			return
		}
		file = filepath.Join(hashDir, mounterutils.GetPasswdFileName(fuseType))
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
		_, err = rotateTokenFiles(tokenDir, secrets)
		if err != nil {
			return
		}
		dir = tokenDir
		return
	}
	return
}
