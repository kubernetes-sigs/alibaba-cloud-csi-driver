package interceptors

import (
	"fmt"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
)

type Ossfs2SecretInterceptor struct {
	passwdFile string
}

var _ mounter.MountInterceptor = Ossfs2SecretInterceptor{}

func NewOssfs2SecretInterceptor() mounter.MountInterceptor {
	return Ossfs2SecretInterceptor{}
}

func (i Ossfs2SecretInterceptor) BeforeMount(req *mounter.MountOperation, _ error) (*mounter.MountOperation, error) {
	var err error
	i.passwdFile, err = utils.SaveOssSecretsToFile(req.Secrets, req.FsType)
	if err != nil {
		return req, err
	}
	if i.passwdFile != "" {
		req.Args = append(req.Args, []string{"-c", i.passwdFile}...)
	}
	return req, nil
}

func (i Ossfs2SecretInterceptor) AfterMount(op *mounter.MountOperation, err error) error {
	if i.passwdFile == "" {
		return nil
	}
	if err := os.Remove(i.passwdFile); err != nil {
		return fmt.Errorf("error removing configuration file: %w, mountpoint=%s, path=%s", err, op.Target, i.passwdFile)
	}
	return nil
}
