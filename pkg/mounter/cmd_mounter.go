package mounter

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud/metadata"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"k8s.io/mount-utils"
)

const timeout = time.Second * 10

type OssCmdMounter struct {
	mount.Interface
	volumeId string
	execPath string
	metadata metadata.MetadataProvider
}

func NewOssCmdMounter(execPath, volumeId string, metadata metadata.MetadataProvider, inner mount.Interface) Mounter {
	return &OssCmdMounter{
		execPath:  execPath,
		volumeId:  volumeId,
		Interface: inner,
		metadata:  metadata,
	}
}

func (m *OssCmdMounter) MountWithSecrets(source, target, fstype string, options []string, authCfg *utils.AuthConfig) error {
	if authCfg == nil {
		return errors.New("empty auth config")
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	options, err := m.appendAuthOptions(options, target, authCfg)
	if err != nil {
		return err
	}

	args := mount.MakeMountArgs(source, target, "", options)
	cmd := exec.CommandContext(ctx, "ossfs", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Run(); err != nil {
		return fmt.Errorf("failed to execute ossfs: %w", err)
	}
	return nil
}

func (m *OssCmdMounter) appendAuthOptions(options []string, target string, authCfg *utils.AuthConfig) ([]string, error) {
	if authCfg == nil {
		return options, nil
	}

	passwdFile, err := saveOssSecretsToFileIfNeeded(authCfg)
	if err != nil {
		return nil, err
	}

	switch authCfg.AuthType {
	case "rrsa":
		tokenFile, err := m.metadata.Get(metadata.RRSATokenFile)
		if err != nil {
			return nil, err
		}
		sessionName := utils.GetRoleSessionName(m.volumeId, target, "ossfs")
		options = append(options, fmt.Sprintf("rrsa_oidc_provider_arn=%s", authCfg.RrsaConfig.OidcProviderArn))
		options = append(options, fmt.Sprintf("rrsa_role_arn=%s", authCfg.RrsaConfig.RoleArn))
		options = append(options, fmt.Sprintf("rrsa_role_session_name=%s", sessionName))
		options = append(options, fmt.Sprintf("rrsa_token_file=%s", tokenFile))
	default:
		options = append(options, "passwd_file="+passwdFile)
	}
	return options, nil
}

func saveOssSecretsToFileIfNeeded(authCfg *utils.AuthConfig) (string, error) {
	if authCfg == nil || authCfg.Secrets == nil {
		return "", nil
	}
	return utils.SaveOssSecretsToFile(authCfg.Secrets)
}
