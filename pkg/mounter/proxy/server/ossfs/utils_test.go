package ossfs

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
	OssfsPasswdFile = utils.GetPasswdFileName("ossfs")
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
			wantOpts: true,
			wantErr:  false,
		},
		{
			name: "TokenSecretsExists",
			secrets: map[string]string{
				oss.KeyAccessKeyId:     "testAKID",
				oss.KeyAccessKeySecret: "testAKSecret",
				oss.KeyExpiration:      "testExpiration",
				oss.KeySecurityToken:   "testSecurityToken",
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
	assert.Equal(t, false, rotated)
	assert.NoError(t, err)
	// case 2: initialize token
	secrets = map[string]string{
		oss.KeyAccessKeyId:     "testAKID",
		oss.KeyAccessKeySecret: "testAKSecret",
		oss.KeyExpiration:      "testExpiration",
		oss.KeySecurityToken:   "testSecurityToken",
	}
	rotated, err = rotateTokenFiles("/tmp/token-files", secrets)
	assert.Equal(t, true, rotated)
	assert.NoError(t, err)
	ak, _ := os.ReadFile(filepath.Join("/tmp/token-files", oss.KeyAccessKeyId))
	assert.Equal(t, "testAKID", string(ak))
	sk, _ := os.ReadFile(filepath.Join("/tmp/token-files", oss.KeyAccessKeySecret))
	assert.Equal(t, "testAKSecret", string(sk))
	exp, _ := os.ReadFile(filepath.Join("/tmp/token-files", oss.KeyExpiration))
	assert.Equal(t, "testExpiration", string(exp))
	st, _ := os.ReadFile(filepath.Join("/tmp/token-files", oss.KeySecurityToken))
	assert.Equal(t, "testSecurityToken", string(st))
	// case 3: rotate token
	secrets = map[string]string{
		oss.KeyAccessKeyId:     "newAKID",
		oss.KeyAccessKeySecret: "newAKSecret",
		oss.KeyExpiration:      "newExpiration",
		oss.KeySecurityToken:   "newSecurityToken",
	}
	rotated, err = rotateTokenFiles("/tmp/token-files", secrets)
	assert.Equal(t, true, rotated)
	assert.NoError(t, err)
	ak, _ = os.ReadFile(filepath.Join("/tmp/token-files", oss.KeyAccessKeyId))
	assert.Equal(t, "newAKID", string(ak))
	sk, _ = os.ReadFile(filepath.Join("/tmp/token-files", oss.KeyAccessKeySecret))
	assert.Equal(t, "newAKSecret", string(sk))
	exp, _ = os.ReadFile(filepath.Join("/tmp/token-files", oss.KeyExpiration))
	assert.Equal(t, "newExpiration", string(exp))
	st, _ = os.ReadFile(filepath.Join("/tmp/token-files", oss.KeySecurityToken))
	assert.Equal(t, "newSecurityToken", string(st))
	os.RemoveAll("/tmp/token-files")
}
