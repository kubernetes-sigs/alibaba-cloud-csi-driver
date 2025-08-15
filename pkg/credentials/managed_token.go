package credentials

import (
	"fmt"
	"os"
	"time"

	alicred "github.com/aliyun/credentials-go/credentials/providers"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/crypto"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

const (
	// the secret mount file
	DefaultTokenPath = "/var/addon/token-config"
	TokenPathEnvName = "ALIBABA_CLOUD_ACK_ADDON_TOKEN_PATH"
)

type ManagedTokenCredential struct {
	tokens     *crypto.RamToken
	tokenPath  string
	reloadTime time.Time
	expiration time.Time
	clock      clock.PassiveClock
}

func (m *ManagedTokenCredential) reloadToken() error {
	f, err := os.Open(m.tokenPath)
	if err != nil {
		return fmt.Errorf("failed to read token config: %w", err)
	}
	defer f.Close()

	tokens, err := crypto.RamTokenParse(f)
	if err != nil {
		return err
	}
	m.tokens = tokens
	klog.V(4).InfoS("reloaded managed token", "expireAt", m.tokens.Expiration)
	m.expiration, err = time.Parse(time.RFC3339, m.tokens.Expiration)
	if err != nil {
		klog.Warningf("failed to parse token expiration: %s", m.tokens.Expiration)
	}

	m.reloadTime = m.clock.Now()
	if m.expiration.Before(m.reloadTime) {
		klog.Warningf("token already expired at %s on load", m.tokens.Expiration)
	}

	return nil
}

func (m *ManagedTokenCredential) reloadTokenIfNeeded() error {
	now := m.clock.Now()
	// reload token every 5 minutes, or 10 minutes before expiration
	if now.After(m.reloadTime.Add(5*time.Minute)) || m.expiration.Before(now.Add(10*time.Minute)) {
		return m.reloadToken()
	}
	return nil
}

func GetManagedTokenPath() string {
	tokenPath := os.Getenv(TokenPathEnvName)
	if tokenPath == "" {
		tokenPath = DefaultTokenPath
	}
	return tokenPath
}

func NewManagedTokenProvider(clk clock.PassiveClock, tokenPath string) (alicred.CredentialsProvider, error) {
	c := &ManagedTokenCredential{
		clock:     clk,
		tokenPath: tokenPath,
	}
	err := c.reloadToken()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (m *ManagedTokenCredential) GetProviderName() string {
	return "managed_token"
}

func (m *ManagedTokenCredential) GetCredentials() (*alicred.Credentials, error) {
	err := m.reloadTokenIfNeeded()
	if err != nil {
		return nil, err
	}
	return &alicred.Credentials{
		ProviderName:    m.GetProviderName(),
		AccessKeyId:     m.tokens.AccessKeyId,
		AccessKeySecret: m.tokens.AccessKeySecret,
		SecurityToken:   m.tokens.SecurityToken,
	}, nil
}
