package interceptors

import (
	"fmt"
	"os"
	"path"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"k8s.io/klog/v2"
)

type AlinasSecretInterceptor struct{}

var _ mounter.MountInterceptor = AlinasSecretInterceptor{}

func NewAlinasSecretInterceptor() mounter.MountInterceptor {
	return AlinasSecretInterceptor{}
}

func (AlinasSecretInterceptor) BeforeMount(op *mounter.MountOperation, _ error) (*mounter.MountOperation, error) {
	if op == nil || op.Secrets == nil {
		return op, nil
	}

	tmpDir, err := os.MkdirTemp("", "alinas-")
	if err != nil {
		return op, err
	}

	credFileContent := makeCredFileContent(op.Secrets)
	credFilePath := path.Join(tmpDir, op.VolumeID+".credentials")
	if err = os.WriteFile(credFilePath, credFileContent, 0o600); err != nil {
		return op, err
	}

	klog.V(4).InfoS("Created alinas credential file", "path", credFilePath)
	op.Options = append(op.Options, "ram_config_file="+credFilePath)
	return op, nil
}

func makeCredFileContent(secrets map[string]string) []byte {
	return fmt.Appendf(
		nil,
		"[NASCredentials]\naccessKeyID=%s\naccessKeySecret=%s",
		secrets["akId"],
		secrets["akSecret"],
	)
}

func (AlinasSecretInterceptor) AfterMount(op *mounter.MountOperation, err error) error {
	return nil
}
