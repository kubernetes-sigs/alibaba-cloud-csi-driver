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
func rotateTokenFiles(dir string, secrets map[string]string) (rotated bool, err error) {
	if secrets == nil {
		return false, nil
	}
	// token
	tokenKey := []string{oss.KeyAccessKeyId, oss.KeyAccessKeySecret, oss.KeySecurityToken, oss.KeyExpiration}
	for _, key := range tokenKey {
		val := secrets[filepath.Join(utils.GetPasswdFileName("ossfs"), key)]
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
		rotated = true
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
	err = os.MkdirAll(hashDir, 0o644)
	if err != nil {
		klog.Errorf("mkdirall hashdir failed %v", err)
		return
	}

	if passwd := secrets[utils.GetPasswdFileName("ossfs")]; passwd != "" {
		err = os.WriteFile(filepath.Join(hashDir, utils.GetPasswdFileName("ossfs")), []byte(passwd), 0o600)
		if err != nil {
			return
		}
		file = filepath.Join(hashDir, utils.GetPasswdFileName("ossfs"))
		options = append(options, "passwd_file="+file)
		return
	}

	// token
	var token bool
	token, err = rotateTokenFiles(hashDir, secrets)
	if err != nil {
		return
	}
	if token {
		dir = hashDir
		options = append(options, "passwd_file="+dir)
		return
	}

	// other authorizarion methods
	err = os.Remove(hashDir)
	if err != nil {
		klog.Errorf("removeall hashdir failed %v", err)
	}
	return
}
