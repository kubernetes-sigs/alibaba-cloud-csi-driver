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
	mount.Interface
	volumeId string
	execPath string
}

func NewOssCmdMounter(execPath, volumeId string, inner mount.Interface) Mounter {
	return &OssCmdMounter{
		execPath:  execPath,
		volumeId:  volumeId,
		Interface: inner,
	}
}

func (m *OssCmdMounter) ExtendedMount(source, target, fstype string, options []string, parms ExtendedMountParams) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	passwd, err := utils.SaveOssSecretsToFile(parms.Secrets, fstype)
	if err != nil {
		return err
	}

	args := getArgs(source, target, fstype, passwd, options)
	cmd := exec.CommandContext(ctx, m.execPath, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ossfs: %w", err)
	}
	return nil
}

func getArgs(source, target, fstype, passwdFile string, options []string) []string {
	if fstype == "ossfs" {
		if passwdFile != "" {
			options = append(options, "passwd_file="+passwdFile)
		}
		return mount.MakeMountArgs(source, target, "", options)
	}
	// ossfs2
	args := []string{"mount", target}
	if passwdFile != "" {
		args = append(args, []string{"-c", passwdFile}...)
	}
	for _, o := range options {
		args = append(args, fmt.Sprintf("--%s", o))
	}
	return args
}
