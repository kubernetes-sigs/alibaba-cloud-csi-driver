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
func rotateTokenFiles(dir string, secrets map[string]string) (rotatedFiles map[string]string, err error) {
	if secrets == nil {
		return nil, nil
	}
	// token
	rotatedFiles = make(map[string]string)
	tokenKey := []string{oss.KeyAccessKeyId, oss.KeyAccessKeySecret, oss.KeySecurityToken}
	for _, key := range tokenKey {
		val := secrets[filepath.Join(utils.GetPasswdFileName("ossfs2"), key)]
		if val == "" {
			err = fmt.Errorf("invalid authorization. %s is empty", key)
			klog.Error(err)
			return
		}
		err = os.WriteFile(filepath.Join(dir, key), []byte(val), 0o600)
		if err != nil {
			klog.Errorf("writeFile %s failed %v", key, err)
			return
		}
		rotatedFiles[key] = filepath.Join(dir, key)
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
	err = os.MkdirAll(hashDir, 0o644)
	if err != nil {
		klog.Errorf("mkdirall hashdir failed %v", err)
		return
	}

	if passwd := secrets[utils.GetPasswdFileName("ossfs2")]; passwd != "" {
		err = os.WriteFile(filepath.Join(hashDir, utils.GetPasswdFileName("ossfs2")), []byte(passwd), 0o600)
		if err != nil {
			return
		}
		file = filepath.Join(hashDir, utils.GetPasswdFileName("ossfs2"))
		return
	}

	// token
	tokenFiles, err := rotateTokenFiles(hashDir, secrets)
	if len(tokenFiles) != 0 {
		dir = hashDir
		options = append(options,
			fmt.Sprintf("oss_sts_multi_conf_ak_file=%s", tokenFiles[oss.KeyAccessKeyId]),
			fmt.Sprintf("oss_sts_multi_conf_sk_file=%s", tokenFiles[oss.KeyAccessKeySecret]),
			fmt.Sprintf("oss_sts_multi_conf_token_file=%s", tokenFiles[oss.KeySecurityToken]),
		)
		return
	}

	// other authorizarion methods
	err = os.Remove(hashDir)
	if err != nil {
		klog.Errorf("removeall hashdir failed %v", err)
		return
	}
	return
}
