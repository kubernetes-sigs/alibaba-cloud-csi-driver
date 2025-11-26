package interceptors

import (
	"context"
	"fmt"
	"os"
	"path"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"k8s.io/klog/v2"
)

var _ mounter.MountInterceptor = AlinasSecretInterceptor

func AlinasSecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	if op == nil || op.Secrets == nil {
		return handler(ctx, op)
	}

	tmpDir, err := os.MkdirTemp("", "alinas-")
	if err != nil {
		return err
	}

	credFileContent := makeCredFileContent(op.Secrets)
	credFilePath := path.Join(tmpDir, op.VolumeID+".credentials")
	if err = os.WriteFile(credFilePath, credFileContent, 0o600); err != nil {
		os.RemoveAll(tmpDir) // cleanup in case of error
		return err
	}

	klog.V(4).InfoS("Created alinas credential file", "path", credFilePath)
	op.Options = append(op.Options, "ram_config_file="+credFilePath)

	return handler(ctx, op)
}

func makeCredFileContent(secrets map[string]string) []byte {
	return fmt.Appendf(
		nil,
		"[NASCredentials]\naccessKeyID=%s\naccessKeySecret=%s",
		secrets["akId"],
		secrets["akSecret"],
	)
}
