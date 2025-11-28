package interceptors

import (
	"context"
	"errors"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/proxy/server"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
)

var _ mounter.MountInterceptor = OssfsSecretInterceptor

func OssfsSecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	return ossfsSecretInterceptor(ctx, op, handler, "ossfs")
}

func Ossfs2SecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) error {
	return ossfsSecretInterceptor(ctx, op, handler, "ossfs2")
}

func ossfsSecretInterceptor(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler, fuseType string) error {
	if op == nil || op.Secrets == nil {
		return handler(ctx, op)
	}

	passwdFile, err := utils.SaveOssSecretsToFile(op.Secrets, op.FsType)
	if err != nil {
		return err
	}
	if passwdFile != "" {
		if fuseType == "ossfs" {
			op.Args = append(op.Args, "passwd_file="+passwdFile)
		} else {
			op.Args = append(op.Args, []string{"-c", passwdFile}...)
		}
	}

	if err = handler(ctx, op); err != nil {
		return err
	}

	if passwdFile == "" || op.MountResult == nil {
		return nil
	}
	result, ok := op.MountResult.(server.OssfsMountResult)
	if !ok {
		klog.ErrorS(
			errors.New("failed to assert ossfs mount result"),
			"skipping cleanup of passwd file", "mountpoint", op.Target, "path", passwdFile,
		)
		return nil
	}

	go func() {
		<-result.ExitChan
		if err := os.Remove(passwdFile); err != nil {
			klog.ErrorS(err, "failed to cleanup ossfs passwd file", "mountpoint", op.Target, "path", passwdFile)
		}
	}()
	return nil
}
