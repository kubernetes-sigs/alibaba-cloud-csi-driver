package ossfs2

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
	tokenKey := []string{oss.KeyAccessKeyId, oss.KeyAccessKeySecret, oss.KeySecurityToken}
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
		rotated = rotated || fileUpdate
	}
	return
}

// prepareCredentialFiles returns:
//  1. file:    path of ossfs2 credential file for fixed AKSK
//  2. dir:     dorectory of ossfs2 credential files for token
//  3. options: extra options
//  4. error
func prepareCredentialFiles(target string, secrets map[string]string) (file, dir string, options []string, err error) {
	// fixed AKSK
	hashDir := utils.GetPasswdHashDir(target)

	if passwd := secrets[utils.GetPasswdFileName("ossfs2")]; passwd != "" {
		err = os.MkdirAll(hashDir, 0o644)
		if err != nil {
			klog.Errorf("mkdirall hashdir failed %v", err)
			return
		}
		_, err = utils.WriteFileWithLock(filepath.Join(hashDir, utils.GetPasswdFileName("ossfs2")), []byte(passwd), 0o600)
		if err != nil {
			return
		}
		file = filepath.Join(hashDir, utils.GetPasswdFileName("ossfs2"))
		return
	}

	// token
	if token := secrets[oss.KeySecurityToken]; token != "" {
		tokenDir := filepath.Join(hashDir, utils.GetPasswdFileName("ossfs2"))
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
		options = append(options,
			fmt.Sprintf("oss_sts_multi_conf_ak_file=%s", filepath.Join(dir, oss.KeyAccessKeyId)),
			fmt.Sprintf("oss_sts_multi_conf_sk_file=%s", filepath.Join(dir, oss.KeyAccessKeySecret)),
			fmt.Sprintf("oss_sts_multi_conf_token_file=%s", filepath.Join(dir, oss.KeySecurityToken)),
		)
		return
	}
	return
}
