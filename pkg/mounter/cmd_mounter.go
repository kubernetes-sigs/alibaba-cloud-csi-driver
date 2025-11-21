package mounter

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"k8s.io/mount-utils"
)

const timeout = time.Second * 10

type OssCmdMounter struct {
	execPath string
	volumeID string
	mount.Interface
}

var _ Mounter = &OssCmdMounter{}

func NewOssCmdMounter(execPath, volumeID string, inner mount.Interface) Mounter {
	return &OssCmdMounter{
		execPath:  execPath,
		volumeID:  volumeID,
		Interface: inner,
	}
}

func (m *OssCmdMounter) ExtendedMount(_ context.Context, op *MountOperation) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	cmd := exec.CommandContext(ctx, m.execPath, getArgs(op)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ossfs: %w", err)
	}
	return nil
}

func getArgs(op *MountOperation) []string {
	if op == nil {
		return nil
	}
	switch op.FsType {
	case "ossfs":
		return mount.MakeMountArgs(op.Source, op.Target, "", op.Options)
	case "ossfs2":
		args := []string{"mount", op.Target}
		args = append(args, op.Args...)
		for _, o := range op.Options {
			args = append(args, fmt.Sprintf("--%s", o))
		}
		return args
	default:
		return nil
	}
}
