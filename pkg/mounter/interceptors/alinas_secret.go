package interceptors

import (
	"context"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	sts20150401 "github.com/alibabacloud-go/sts-20150401/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	alicred "github.com/aliyun/credentials-go/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/cloud"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/credentials"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"k8s.io/klog/v2"
	mountutils "k8s.io/mount-utils"
	"k8s.io/utils/clock"
)

const (
	rrsaAuthType = "rrsa"

	stsTokenDurationSeconds = 3600
)

var (
	credDir      = os.TempDir()
	rrsaMountDir = "/var/run/secrets/ack.alibabacloud.com/rrsa-tokens"
)

type AlinasSecretInterceptor struct {
	stsClient     cloud.STSInterface
	ramRoleTokens sync.Map
	clock         clock.Clock
	mounter       mountutils.Interface
}

type ramRoleToken struct {
	accessKey     string
	accessSecret  string
	securityToken string
	expiresAt     time.Time
	refreshAt     time.Time
}

func NewAlinasSecretInterceptor(regionID string) (*AlinasSecretInterceptor, error) {
	provider, err := credentials.NewProvider()
	if err != nil {
		return nil, err
	}

	cred := alicred.FromCredentialsProvider(provider.GetProviderName(), provider)
	stsConfig := utils.GetStsConfig(regionID).SetCredential(cred)
	stsClient, err := sts20150401.NewClient(stsConfig)
	if err != nil {
		return nil, err
	}

	return &AlinasSecretInterceptor{
		stsClient:     stsClient,
		clock:         clock.RealClock{},
		ramRoleTokens: sync.Map{},
		mounter:       mountutils.New(""),
	}, nil
}

func (i *AlinasSecretInterceptor) Intercept(ctx context.Context, op *mounter.MountOperation, handler mounter.MountHandler) (err error) {
	if op == nil || op.AuthConfig == nil {
		return handler(ctx, op)
	}
	if op.AuthConfig.AuthType != rrsaAuthType && (op.AuthConfig.AccessKey == "" || op.AuthConfig.AccessSecret == "") {
		return handler(ctx, op)
	}

	var credFilePath string
	if op.AuthConfig != nil {
		switch op.AuthConfig.AuthType {
		case rrsaAuthType:
			if credFilePath, err = i.refreshAndSaveRRSAToken(op); err != nil {
				return
			}
		default:
			if credFilePath, err = saveCredentials(op); err != nil {
				return
			}
		}
	}

	klog.V(4).InfoS("Created alinas credential file", "path", credFilePath)
	op.Options = append(op.Options, "ram_config_file="+credFilePath)
	if err = handler(ctx, op); err != nil {
		return
	}

	if op.AuthConfig != nil && op.AuthConfig.AuthType == rrsaAuthType {
		go i.startTokenRefreshLoop(op)
	}
	return
}

func (i *AlinasSecretInterceptor) startTokenRefreshLoop(op *mounter.MountOperation) {
	if op == nil || op.Target == "" || op.AuthConfig == nil {
		return
	}

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	klog.V(4).InfoS("Started RRSA token refresh loop", "target", op.Target, "volumeID", op.VolumeID)

	for {
		<-ticker.C
		if !i.isMountPoint(op.Target) {
			klog.InfoS("Target is no longer a mount point, stopping RRSA token refresh loop", "target", op.Target, "volumeID", op.VolumeID)
			return
		}

		if _, err := i.refreshAndSaveRRSAToken(op); err != nil {
			klog.ErrorS(err, "Failed to refresh RRSA token", "target", op.Target, "volumeID", op.VolumeID)
		}
	}
}

func (i *AlinasSecretInterceptor) isMountPoint(target string) bool {
	if target == "" {
		return false
	}

	notMnt, err := i.mounter.IsLikelyNotMountPoint(target)
	if err != nil {
		if os.IsNotExist(err) {
			klog.V(4).InfoS("Target path does not exist", "target", target)
			return false
		}
		if mountutils.IsCorruptedMnt(err) {
			klog.V(4).InfoS("Target path is corrupted", "target", target)
			return false
		}
		klog.V(4).InfoS("Error checking mount point", "target", target, "error", err)
		return false
	}

	return !notMnt
}

func (i *AlinasSecretInterceptor) refreshAndSaveRRSAToken(op *mounter.MountOperation) (credFilePath string, err error) {
	if op == nil || op.AuthConfig == nil || op.AuthConfig.AuthType != rrsaAuthType {
		return
	}

	credFilePath = path.Join(credDir, op.VolumeID+".credentials")
	save := func(token ramRoleToken) (string, error) {
		op.AuthConfig.AccessKey = token.accessKey
		op.AuthConfig.AccessSecret = token.accessSecret
		op.AuthConfig.SecurityToken = token.securityToken
		return saveCredentials(op)
	}

	if !i.shouldRefreshRRSAToken(op.AuthConfig.RoleName) {
		klog.V(4).InfoS("RAM role token is still fresh, no need to refresh", "roleName", op.AuthConfig.RoleName)
		return
	}

	var token ramRoleToken
	if token, err = i.refreshRRSAToken(op); err != nil {
		return
	}
	return save(token)
}

func (i *AlinasSecretInterceptor) refreshRRSAToken(op *mounter.MountOperation) (token ramRoleToken, err error) {
	if op == nil {
		return
	}

	oidcTokenFilePath := path.Join(rrsaMountDir, "token")
	oidcToken, err := os.ReadFile(oidcTokenFilePath)
	if err != nil {
		return
	}

	session := mounterutils.GetRoleSessionName(op.VolumeID, op.Target, "alinas")
	provider := mounterutils.GetOIDCProvider(op.AuthConfig.ClusterID)
	providerArn, roleArn := mounterutils.GetArn(provider, op.AuthConfig.AccountID, op.AuthConfig.RoleName)

	req := &sts20150401.AssumeRoleWithOIDCRequest{
		DurationSeconds: tea.Int64(stsTokenDurationSeconds),
		OIDCToken:       tea.String(string(oidcToken)),
		OIDCProviderArn: tea.String(providerArn),
		RoleArn:         tea.String(roleArn),
		RoleSessionName: tea.String(session),
	}
	resp, err := i.stsClient.AssumeRoleWithOIDC(req)
	if err != nil || resp.Body == nil || resp.Body.Credentials == nil {
		return
	}

	klog.V(4).InfoS("RAM role token refreshed", "roleName", op.AuthConfig.RoleName)
	cred := resp.Body.Credentials
	token = ramRoleToken{
		accessKey:     tea.StringValue(cred.AccessKeyId),
		accessSecret:  tea.StringValue(cred.AccessKeySecret),
		securityToken: tea.StringValue(cred.SecurityToken),
		expiresAt:     i.clock.Now().Add(time.Second * stsTokenDurationSeconds),
		refreshAt:     i.clock.Now(),
	}
	i.ramRoleTokens.Store(op.AuthConfig.RoleName, token)
	return
}

func (i *AlinasSecretInterceptor) shouldRefreshRRSAToken(roleName string) bool {
	value, exists := i.ramRoleTokens.Load(roleName)
	if !exists {
		return true
	}
	token, ok := value.(ramRoleToken)
	if !ok {
		return true
	}

	now := i.clock.Now()
	return now.After(token.refreshAt.Add(5*time.Minute)) || token.expiresAt.Before(now.Add(10*time.Minute))
}

func saveCredentials(op *mounter.MountOperation) (credFilePath string, err error) {
	if op == nil || op.AuthConfig == nil {
		return
	}

	tmpCredFile, err := os.CreateTemp(credDir, op.VolumeID+"-*.credentials")
	if err != nil {
		return
	}
	tmpFilePath := tmpCredFile.Name()

	defer func() {
		if err != nil {
			// Only remove temp file if there was an error (rename didn't succeed)
			if removeErr := os.Remove(tmpFilePath); removeErr != nil && !os.IsNotExist(removeErr) {
				klog.ErrorS(removeErr, "Failed to remove temporary alinas credential file", "path", tmpFilePath)
			}
		}
	}()

	credFileContent := makeCredFileContent(op.AuthConfig)
	if _, err = tmpCredFile.Write(credFileContent); err != nil {
		tmpCredFile.Close()
		return
	}

	if err = tmpCredFile.Close(); err != nil {
		return
	}

	credFilePath = path.Join(credDir, op.VolumeID+".credentials")
	err = os.Rename(tmpFilePath, credFilePath)
	return
}

func makeCredFileContent(authConfig *mounterutils.AuthConfig) []byte {
	if authConfig == nil {
		return nil
	}
	return fmt.Appendf(
		nil,
		"[NASCredentials]\naccessKeyID=%s\naccessKeySecret=%s\nsecurityToken=%s",
		authConfig.AccessKey,
		authConfig.AccessSecret,
		authConfig.SecurityToken,
	)
}
