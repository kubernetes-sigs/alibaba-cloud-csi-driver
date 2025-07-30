package ossfs2

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"k8s.io/klog/v2"
)

var (
	OssfsPasswdFile = utils.GetPasswdFileName("ossfs2")
)

func TestPrepareCredentialFiles(t *testing.T) {
	tests := []struct {
		name     string
		secrets  map[string]string
		wantFile bool
		wantDir  bool
		wantOpts bool
		wantErr  bool
	}{
		{
			name:    "EmptySecrets",
			secrets: map[string]string{},
		},
		{
			name:     "FixedAKSKExists",
			secrets:  map[string]string{OssfsPasswdFile: "testPasswd"},
			wantFile: true,
			wantDir:  false,
			wantOpts: false,
			wantErr:  false,
		},
		{
			name: "TokenSecretsExists",
			secrets: map[string]string{
				filepath.Join(OssfsPasswdFile, oss.KeyAccessKeyId):     "testAKID",
				filepath.Join(OssfsPasswdFile, oss.KeyAccessKeySecret): "testAKSecret",
				filepath.Join(OssfsPasswdFile, oss.KeySecurityToken):   "testSecurityToken",
			},
			wantFile: false,
			wantDir:  true,
			wantOpts: true,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			file, dir, options, err := prepareCredentialFiles("/tmp/token-files", tt.secrets)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantFile, file != "")
			assert.Equal(t, tt.wantDir, dir != "")
			assert.Equal(t, tt.wantOpts, len(options) != 0)
			err = os.RemoveAll("/tmp/token-files")
			klog.ErrorS(err, "Remove token directory", "dir", "/tmp/token-files")
		})
	}

}

func TestRotateTokenFiles(t *testing.T) {
	// case 1: initialize fiexd AKSK
	secrets := map[string]string{OssfsPasswdFile: "testPasswd"}
	rotated, err := rotateTokenFiles("/tmp/token-files", secrets)
	assert.Len(t, rotated, 0)
	assert.NoError(t, err)
	err = os.RemoveAll("/tmp/token-files")
	klog.ErrorS(err, "Remove token directory", "dir", "/tmp/token-files")

	// case 2: initialize token
	secrets = map[string]string{
		filepath.Join(OssfsPasswdFile, oss.KeyAccessKeyId):     "testAKID",
		filepath.Join(OssfsPasswdFile, oss.KeyAccessKeySecret): "testAKSecret",
		filepath.Join(OssfsPasswdFile, oss.KeyExpiration):      "testExpiration",
		filepath.Join(OssfsPasswdFile, oss.KeySecurityToken):   "testSecurityToken",
	}
	rotated, err = rotateTokenFiles("/tmp/token-files", secrets)
	assert.Len(t, rotated, 3)
	assert.NoError(t, err)
	ak, _ := os.ReadFile(filepath.Join("/tmp/token-files", oss.KeyAccessKeyId))
	assert.Equal(t, "testAKID", string(ak))
	sk, _ := os.ReadFile(filepath.Join("/tmp/token-files", oss.KeyAccessKeySecret))
	assert.Equal(t, "testAKSecret", string(sk))
	st, _ := os.ReadFile(filepath.Join("/tmp/token-files", oss.KeySecurityToken))
	assert.Equal(t, "testSecurityToken", string(st))
	err = os.RemoveAll("/tmp/token-files")
	klog.ErrorS(err, "Remove token directory", "dir", "/tmp/token-files")

	// case 3: rotate token
	secrets = map[string]string{
		filepath.Join(OssfsPasswdFile, oss.KeyAccessKeyId):     "newAKID",
		filepath.Join(OssfsPasswdFile, oss.KeyAccessKeySecret): "newAKSecret",
		filepath.Join(OssfsPasswdFile, oss.KeyExpiration):      "newExpiration",
		filepath.Join(OssfsPasswdFile, oss.KeySecurityToken):   "newSecurityToken",
	}
	rotated, err = rotateTokenFiles("/tmp/token-files", secrets)
	assert.Len(t, rotated, 3)
	assert.NoError(t, err)
	ak, _ = os.ReadFile(filepath.Join("/tmp/token-files", oss.KeyAccessKeyId))
	assert.Equal(t, "newAKID", string(ak))
	sk, _ = os.ReadFile(filepath.Join("/tmp/token-files", oss.KeyAccessKeySecret))
	assert.Equal(t, "newAKSecret", string(sk))
	st, _ = os.ReadFile(filepath.Join("/tmp/token-files", oss.KeySecurityToken))
	assert.Equal(t, "newSecurityToken", string(st))
	os.RemoveAll("/tmp/token-files")
	klog.ErrorS(err, "Remove token directory", "dir", "/tmp/token-files")
}
