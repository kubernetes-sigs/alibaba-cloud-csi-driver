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
func rotateTokenFiles(secrets map[string]string) (rotated bool, err error) {
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
		err = os.MkdirAll(OssfsTokenFilesDir, 0o644)
		if err != nil {
			klog.Errorf("mkdirall tokendir failed %v", err)
			return
		}
		err = writeFile(OssfsTokenFilesDir, key, val, 0o600)
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
func prepareCredentialFiles(secrets map[string]string) (file, dir string, options []string, err error) {
	// fixed AKSK
	if passwd := secrets[OssfsPasswdFile]; passwd != "" {
		var tmpDir string
		tmpDir, err = os.MkdirTemp("", "ossfs-")
		if err != nil {
			return
		}
		err = writeFile(tmpDir, OssfsPasswdFileName, passwd, 0o600)
		if err != nil {
			return
		}
		file = filepath.Join(tmpDir, OssfsPasswdFileName)
		options = append(options, "passwd_file="+file)
		return
	}

	// token
	var token bool
	token, err = rotateTokenFiles(secrets)
	if token {
		dir = OssfsTokenFilesDir
		options = append(options, "passwd_file="+dir)
	}
	return
}
