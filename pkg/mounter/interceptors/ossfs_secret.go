package interceptors

import (
	"fmt"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
)

type OssfsSecretInterceptor struct {
	passwdFile string
}

var _ mounter.MountInterceptor = OssfsSecretInterceptor{}

func NewOssfsSecretInterceptor() mounter.MountInterceptor {
	return OssfsSecretInterceptor{}
}

func (i OssfsSecretInterceptor) BeforeMount(req *mounter.MountOperation, _ error) (*mounter.MountOperation, error) {
	var err error
	i.passwdFile, err = utils.SaveOssSecretsToFile(req.Secrets, req.FsType)
	if err != nil {
		return req, err
	}
	if i.passwdFile != "" {
		req.Options = append(req.Options, "passwd_file="+i.passwdFile)
	}
	return req, nil
}

func (i OssfsSecretInterceptor) AfterMount(op *mounter.MountOperation, err error) error {
	if i.passwdFile == "" {
		return nil
	}
	if err := os.Remove(i.passwdFile); err != nil {
		return fmt.Errorf("error removing passwd file: %w, mountpoint=%s, path=%s", err, op.Target, i.passwdFile)
	}
	return nil
}
