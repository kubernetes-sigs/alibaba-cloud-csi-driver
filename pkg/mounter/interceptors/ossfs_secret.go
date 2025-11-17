package interceptors

import (
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
)

type OssfsSecretInterceptor struct {
	credFileKey string
}

var _ mounter.MountInterceptor = OssfsSecretInterceptor{}

func NewOssfsSecretInterceptor() mounter.MountInterceptor {
	return OssfsSecretInterceptor{
		credFileKey: "passwd_file",
	}
}

func NewOssfs2SecretInterceptor() mounter.MountInterceptor {
	return OssfsSecretInterceptor{
		credFileKey: "-c",
	}
}

func (i OssfsSecretInterceptor) BeforeMount(req *mounter.MountOperation) (*mounter.MountOperation, error) {
	filePath, err := utils.SaveOssSecretsToFile(req.Secrets, req.FsType)
	if err != nil {
		return req, err
	}
	if filePath != "" {
		req.Options = append(req.Options, i.credFileKey+"="+filePath)
	}
	return req, nil
}
