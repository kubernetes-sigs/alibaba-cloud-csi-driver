package ossfs

import (
	"path/filepath"
	"testing"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/stretchr/testify/assert"
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
				filepath.Join(OssfsPasswdFile, mounter.KeyAccessKeyId):     "testAKID",
				filepath.Join(OssfsPasswdFile, mounter.KeyAccessKeySecret): "testAKSecret",
				filepath.Join(OssfsPasswdFile, mounter.KeyExpiration):      "testExpiration",
				filepath.Join(OssfsPasswdFile, mounter.KeySecurityToken):   "testSecurityToken",
			},
			wantFile: false,
			wantDir:  true,
			wantOpts: true,
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			file, dir, options, err := prepareCredentialFiles(tt.secrets)
			assert.Equal(t, tt.wantErr, err != nil)
			assert.Equal(t, tt.wantFile, file != "")
			assert.Equal(t, tt.wantDir, dir != "")
			assert.Equal(t, tt.wantOpts, len(options) != 0)
		})
	}

}
