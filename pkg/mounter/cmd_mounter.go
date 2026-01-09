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

func (m *OssCmdMounter) ExtendedMount(ctx context.Context, req *utils.MountRequest) error {
	if req == nil {
		return nil
	}

	cmd := exec.CommandContext(ctx, m.execPath, getArgs(req)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ossfs: %w", err)
	}
	return nil
}

func getArgs(req *utils.MountRequest) []string {
	if req == nil {
		return nil
	}
	switch req.Fstype {
	case utils.OssFsType:
		return mount.MakeMountArgs(req.Source, req.Target, "", req.Options)
	case utils.OssFs2Type:
		args := []string{"mount", req.Target}
		args = append(args, req.Args...)
		for _, o := range req.Options {
			args = append(args, fmt.Sprintf("--%s", o))
		}
		return args
	default:
		return nil
	}
}
