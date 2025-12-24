package interceptors

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
)

var _ mounter.MountInterceptor = OssfsSecretInterceptor

func OssfsSecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	return ossfsSecretInterceptor(ctx, op, handler, mounterutils.OssFsType)
}

func Ossfs2SecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	return ossfsSecretInterceptor(ctx, op, handler, mounterutils.OssFs2Type)
}

func ossfsSecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler, fuseType string) error {
	if op == nil || op.Secrets == nil {
		return handler(ctx, op)
	}

	passwdFile, tokenDir, err := prepareCredentialFiles(fuseType, op.Target, op.Secrets)
	if err != nil {
		return fmt.Errorf("prepare credential files failed: %w", err)
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
		if err := os.Remove(passwdFile); err != nil {
			klog.ErrorS(err, "failed to cleanup ossfs passwd file", "mountpoint", op.Target, "path", passwdFile)
		}
	}()
	return nil
}

// rotateTokenFiles rotates (or initializes) token files
// This function assumes that token rotation is required when executed
func rotateTokenFiles(dir string, secrets map[string]string) (rotated bool, err error) {
	var fileUpdate bool
	// Currently, for ossfs2, expiration is not required.
	// But we still manage it (if offered) for the feature.
	tokenKey := []string{mounterutils.KeyAccessKeyId, mounterutils.KeyAccessKeySecret, mounterutils.KeySecurityToken, mounterutils.KeyExpiration}
	for _, key := range tokenKey {
		val := secrets[key]
		if val == "" {
			if key == mounterutils.KeyExpiration {
				continue
			}
			err = fmt.Errorf("invalid authorization. %s is empty", key)
			klog.Error(err)
			return
		}
		fileUpdate, err = utils.WriteFileWithLock(filepath.Join(dir, key), []byte(val), 0o600)
		if err != nil {
			klog.Errorf("WriteFileWithLock %s failed %v", key, err)
			return
		}
		rotated = fileUpdate || rotated
	}
	return
}

// prepareCredentialFiles returns:
//  1. file:    path of ossfs credential file for fixed AKSK
//  2. dir:     dorectory of ossfs credential files for token
//  4. error
func prepareCredentialFiles(fuseType, target string, secrets map[string]string) (file, dir string, err error) {
	// fixed credientials
	hashDir := utils.GetPasswdHashDir(target)

	if passwd := secrets[utils.GetPasswdFileName(fuseType)]; passwd != "" {
		err = os.MkdirAll(hashDir, 0o644)
		if err != nil {
			klog.Errorf("mkdirall hashdir failed %v", err)
			return
		}
		_, err = utils.WriteFileWithLock(filepath.Join(hashDir, utils.GetPasswdFileName(fuseType)), []byte(passwd), 0o600)
		if err != nil {
			return
		}
		file = filepath.Join(hashDir, utils.GetPasswdFileName(fuseType))
		return
	}

	// token
	if token := secrets[mounterutils.KeySecurityToken]; token != "" {
		tokenDir := filepath.Join(hashDir, utils.GetPasswdFileName(fuseType))
		err = os.MkdirAll(tokenDir, 0o644)
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
