package ossfs

import (
	"fmt"
	"os"
	"path/filepath"

	"k8s.io/klog/v2"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
)

// rotateTokenFiles rotates (or initializes) token files
// This function assumes that token rotation is required when executed
func rotateTokenFiles(dir string, secrets map[string]string) (rotated bool, err error) {
	var fileUpdate bool
	tokenKey := []string{oss.KeyAccessKeyId, oss.KeyAccessKeySecret, oss.KeySecurityToken, oss.KeyExpiration}
	for _, key := range tokenKey {
		val := secrets[key]
		if val == "" {
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
//  3. options: extra options
//  4. error
func prepareCredentialFiles(target string, secrets map[string]string) (file, dir string, options []string, err error) {
	// fixed AKSK
	hashDir := utils.GetPasswdHashDir(target)

	if passwd := secrets[utils.GetPasswdFileName("ossfs")]; passwd != "" {
		err = os.MkdirAll(hashDir, 0o644)
		if err != nil {
			klog.Errorf("mkdirall hashdir failed %v", err)
			return
		}
		_, err = utils.WriteFileWithLock(filepath.Join(hashDir, utils.GetPasswdFileName("ossfs")), []byte(passwd), 0o600)
		if err != nil {
			return
		}
		file = filepath.Join(hashDir, utils.GetPasswdFileName("ossfs"))
		options = append(options, "passwd_file="+file)
		return
	}

	// token
	if token := secrets[oss.KeySecurityToken]; token != "" {
		tokenDir := filepath.Join(hashDir, utils.GetPasswdFileName("ossfs"))
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
		options = append(options, "passwd_file="+dir)
		return
	}
	return
}
