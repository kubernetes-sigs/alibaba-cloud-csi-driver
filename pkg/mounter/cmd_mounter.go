package mounter

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/mount-utils"
)

const timeout = time.Second * 10

type OssCmdMounter struct {
	execPath string
	mount.Interface
}

func NewOssCmdMounter(execPath string, inner mount.Interface) Mounter {
	return &OssCmdMounter{
		execPath:  execPath,
		Interface: inner,
	}
}

func (m *OssCmdMounter) MountWithSecrets(source, target, fstype string, options []string, secrets map[string]string) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	passwdFile, err := utils.SaveOssSecretsToFile(secrets)
	if err != nil {
		return err
	}
	options = append(options, "passwd_file="+passwdFile)

	args := mount.MakeMountArgs(source, target, "", options)
	cmd := exec.CommandContext(ctx, "ossfs", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ossfs: %w", err)
	}
	return nil
}
