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

	tmpCredFile, err := os.CreateTemp("", op.VolumeID+"-*.credentials")
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

	credFilePath := path.Join(os.TempDir(), op.VolumeID+".credentials")
	if err = os.Rename(tmpCredFile.Name(), credFilePath); err != nil {
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
