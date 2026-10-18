package interceptors

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"k8s.io/klog/v2"
)

var credDir = os.TempDir()
var _ mounter.MountInterceptor = AlinasSecretInterceptor

// Keys of the credential carried in MountRequest/RefreshRequest
// Secrets. mount.alinas reads all three from the ram_config_file; the token alone
// selects its ID token-only mode, so a partial set must never be written.
const (
	SecretKeyAccessKeyID     = "akId"
	SecretKeyAccessKeySecret = "akSecret"
	SecretKeySecurityToken   = "securityToken"
)

func AlinasSecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	if op == nil || op.Secrets == nil {
		return handler(ctx, op)
	}

	tmpCredFile, err := os.CreateTemp(credDir, op.VolumeID+"-*.credentials")
	if err != nil {
		return err
	}
	defer func() {
		if err = os.Remove(tmpCredFile.Name()); err != nil && !os.IsNotExist(err) {
			klog.ErrorS(err, "Failed to remove temporary alinas credential file", "path", tmpCredFile.Name())
		}
	}()

	credFileContent := makeCredFileContent(op.Secrets)
	if _, err = tmpCredFile.Write(credFileContent); err != nil {
		return err
	}
	if err = tmpCredFile.Close(); err != nil {
		return err
	}

	credFilePath := path.Join(credDir, op.VolumeID+".credentials")
	if err = os.Rename(tmpCredFile.Name(), credFilePath); err != nil {
		return err
	}

	klog.V(4).InfoS("Created alinas credential file", "path", credFilePath)
	op.Options = append(op.Options, "ram_config_file="+credFilePath)

	return handler(ctx, op)
}

func makeCredFileContent(secrets map[string]string) []byte {
	content := fmt.Appendf(
		nil,
		"[NASCredentials]\naccessKeyID=%s\naccessKeySecret=%s",
		secrets[SecretKeyAccessKeyID],
		secrets[SecretKeyAccessKeySecret],
	)
	// Without the token an STS credential is sent as bare AK/SK and the server
	// rejects it with "Unknown error 521".
	if token := secrets[SecretKeySecurityToken]; token != "" {
		content = fmt.Appendf(content, "\nsecurityToken=%s", token)
	}
	return content
}
