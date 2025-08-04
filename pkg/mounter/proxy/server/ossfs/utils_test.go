package ossfs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/oss"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
			hash := utils.ComputeMountPathHash("/mnt/target1")
			file, dir, options, err := prepareCredentialFiles("/mnt/target1", tt.secrets)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantFile, file != "")
			assert.Equal(t, tt.wantDir, dir != "")
			assert.Equal(t, tt.wantOpts, len(options) != 0)
			err = os.RemoveAll("/tmp/" + hash)
			if err != nil {
				klog.ErrorS(err, "Remove token directory", "dir", "/tmp/"+hash)
			}
		})
	}

}

func TestRotateTokenFiles(t *testing.T) {
	mountPath := "/mnt/target2"
	hash := utils.ComputeMountPathHash(mountPath)
	hashDir := filepath.Join("/tmp", hash)
	err := os.MkdirAll(hashDir, 0o644)
	require.NoError(t, err)

	// case 1: initialize fiexd AKSK
	secrets := map[string]string{OssfsPasswdFile: "testPasswd"}
	rotated, err := rotateTokenFiles(hashDir, secrets)
	assert.Error(t, err)
	assert.Equal(t, false, rotated)

	// case 2: initialize token
	secrets = map[string]string{
		oss.KeyAccessKeyId:     "testAKID",
		oss.KeyAccessKeySecret: "testAKSecret",
		oss.KeyExpiration:      "testExpiration",
		oss.KeySecurityToken:   "testSecurityToken",
	}
	rotated, err = rotateTokenFiles(hashDir, secrets)
	assert.NoError(t, err)
	assert.Equal(t, true, rotated)
	ak, _ := os.ReadFile(filepath.Join(hashDir, oss.KeyAccessKeyId))
	assert.Equal(t, "testAKID", string(ak))
	sk, _ := os.ReadFile(filepath.Join(hashDir, oss.KeyAccessKeySecret))
	assert.Equal(t, "testAKSecret", string(sk))
	exp, _ := os.ReadFile(filepath.Join(hashDir, oss.KeyExpiration))
	assert.Equal(t, "testExpiration", string(exp))
	st, _ := os.ReadFile(filepath.Join(hashDir, oss.KeySecurityToken))
	assert.Equal(t, "testSecurityToken", string(st))

	// case 3: rotate token
	secrets = map[string]string{
		oss.KeyAccessKeyId:     "newAKID",
		oss.KeyAccessKeySecret: "newAKSecret",
		oss.KeyExpiration:      "newExpiration",
		oss.KeySecurityToken:   "newSecurityToken",
	}
	rotated, err = rotateTokenFiles(hashDir, secrets)
	assert.NoError(t, err)
	assert.Equal(t, true, rotated)
	ak, _ = os.ReadFile(filepath.Join(hashDir, oss.KeyAccessKeyId))
	assert.Equal(t, "newAKID", string(ak))
	sk, _ = os.ReadFile(filepath.Join(hashDir, oss.KeyAccessKeySecret))
	assert.Equal(t, "newAKSecret", string(sk))
	exp, _ = os.ReadFile(filepath.Join(hashDir, oss.KeyExpiration))
	assert.Equal(t, "newExpiration", string(exp))
	st, _ = os.ReadFile(filepath.Join(hashDir, oss.KeySecurityToken))
	assert.Equal(t, "newSecurityToken", string(st))
	err = os.RemoveAll(hashDir)
	if err != nil {
		t.Errorf("removeall hashdir failed %v", err)
	}
}
