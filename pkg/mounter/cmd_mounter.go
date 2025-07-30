package mounter

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/klog/v2"
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

func (m *OssCmdMounter) Name() string {
	return "cmd-mounter"
}

func (m *OssCmdMounter) RotateToken(target, fstype string, secrets map[string]string) error {
	return ErrNotImplemented(m.Name(), fstype, "rotateToken")
}

func (m *OssCmdMounter) MountWithSecrets(source, target, fstype string, options []string, secrets map[string]string) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	passwd, err := saveOssSecretsToFile(secrets)
	if err != nil {
		return err
	}
	options = append(options, "passwd_file="+passwd)

	args := mount.MakeMountArgs(source, target, "", options)
	cmd := exec.CommandContext(ctx, "ossfs", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ossfs: %w", err)
	}
	return nil
}
func saveOssSecretsToFileIfNeeded(authCfg *utils.AuthConfig) (string, error) {
	if authCfg == nil || authCfg.Secrets == nil {
		return "", nil
	}
	return saveOssSecretsToFile(authCfg.Secrets)
}

func saveOssSecretsToFile(secrets map[string]string) (filePath string, err error) {
	passwd := secrets["passwd-ossfs"]
	if passwd == "" {
		return
	}

	tmpDir, err := os.MkdirTemp("", "ossfs-")
	if err != nil {
		return "", err
	}
	filePath = filepath.Join(tmpDir, "passwd")
	if err = os.WriteFile(filePath, []byte(passwd), 0o600); err != nil {
		return "", err
	}
	klog.V(4).InfoS("created ossfs passwd file", "path", filePath)
	return
}
