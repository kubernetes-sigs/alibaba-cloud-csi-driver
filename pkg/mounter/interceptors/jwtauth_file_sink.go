package interceptors

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/jwtauth"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
)

// credentialDataDir is the sub-directory (managed via atomic symlink swap)
// inside the per-mount output directory that actually holds the STS files.
// Using a nested "sts" dir mirrors OssfsSecretInterceptor and ensures the
// symlink path never collides with a pre-existing regular directory.
const credentialDataDir = "sts"

// jwtAuthFileSink writes the credential as a set of files rotated atomically
// via a directory symlink swap, for FUSE clients that read AK/SK/token files
// from disk. All four files (AccessKeyId, AccessKeySecret, SecurityToken,
// Expiration) become visible together, so readers never observe a torn
// credential set.
//
// It lives in this package (rather than pkg/mounter/jwtauth) so it can share
// the rotateTokenFiles/cleanupTokenFiles helpers with OssfsSecretInterceptor.
type jwtAuthFileSink struct {
	outputDir   string // per-mount base dir; entrypoint reads from here
	dataSymlink string // outputDir/sts symlink swapped atomically on rotation
}

var _ jwtauth.CredentialSink = &jwtAuthFileSink{}

func newJWTAuthFileSink(outputDir string) *jwtAuthFileSink {
	return &jwtAuthFileSink{
		outputDir:   outputDir,
		dataSymlink: filepath.Join(outputDir, credentialDataDir),
	}
}

// Dir returns the directory the FUSE entrypoint should read credential files
// from. It is the symlink that is swapped atomically on each rotation.
func (s *jwtAuthFileSink) Dir() string {
	return s.dataSymlink
}

func (s *jwtAuthFileSink) Apply(cred *jwtauth.STSToken) error {
	if err := os.MkdirAll(s.outputDir, 0o700); err != nil {
		return fmt.Errorf("create credential dir: %w", err)
	}
	secrets := map[string]string{
		mounterutils.KeyAccessKeyId:     cred.AccessKeyID,
		mounterutils.KeyAccessKeySecret: cred.AccessKeySecret,
		mounterutils.KeySecurityToken:   cred.SecurityToken,
		mounterutils.KeyExpiration:      cred.Expiration,
	}
	if _, err := rotateTokenFiles(s.dataSymlink, secrets); err != nil {
		return fmt.Errorf("rotate credential files: %w", err)
	}
	return nil
}

// Cleanup removes the credential files and directories. It should be called
// once the consuming mount process has exited.
func (s *jwtAuthFileSink) Cleanup() {
	cleanupTokenFiles(s.dataSymlink)
	removeIgnoreNotExist(s.outputDir)
}
