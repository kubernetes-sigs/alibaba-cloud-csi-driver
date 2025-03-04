package ossfs

import (
	"os"
	"path/filepath"

	"k8s.io/klog/v2"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
)

func writeFile(dir, fileName, contents string, perm os.FileMode) error {
	file := filepath.Join(dir, fileName)
	if err := os.Truncate(file, 0); err != nil && !os.IsNotExist(err) {
		return err
	}
	return os.WriteFile(file, []byte(contents), perm)
}

// rotateTokenFiles rotates (or initializes) token files
func rotateTokenFiles(dir string, secrets map[string]string) (rotated bool, err error) {
	if secrets == nil {
		return false, nil
	}
	// token
	tokenKey := []string{mounter.KeyAccessKeyId, mounter.KeyAccessKeySecret, mounter.KeySecurityToken, mounter.KeyExpiration}
	for _, key := range tokenKey {
		val := secrets[filepath.Join(OssfsPasswdFile, key)]
		if val == "" {
			continue
		}
		err = os.MkdirAll(dir, 0o644)
		if err != nil {
			klog.Errorf("mkdirall tokendir failed %v", err)
			return
		}
		err = writeFile(dir, key, val, 0o600)
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
	hashDir := mounter.ComputeMountPathHash(target)
	if passwd := secrets[OssfsPasswdFile]; passwd != "" {
		err = writeFile(hashDir, OssfsPasswdFileName, passwd, 0o600)
		if err != nil {
			return
		}
		file = filepath.Join(hashDir, OssfsPasswdFileName)
		options = append(options, "passwd_file="+file)
		return
	}

	// token
	var token bool
	tokenDir := filepath.Join(hashDir, OssfsTokenFilesDir)
	token, err = rotateTokenFiles(tokenDir, secrets)
	if token {
		dir = tokenDir
		options = append(options, "passwd_file="+dir)
	}
	return
}
