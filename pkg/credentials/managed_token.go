package credentials

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/alibabacloud-go/tea/tea"
	alicred "github.com/aliyun/credentials-go/credentials"
	"github.com/sirupsen/logrus"
	"k8s.io/utils/clock"
)

const (
	// the secret mount file
	DefaultTokenPath = "/var/addon/token-config"
	TokenPathEnvName = "ALIBABA_CLOUD_ACK_ADDON_TOKEN_PATH"
)

// data struct stored in ConfigPath
type AKInfo struct {
	AccessKeyID         string `json:"access.key.id"`
	AccessKeySecret     string `json:"access.key.secret"`
	SecurityToken       string `json:"security.token"`
	Expiration          string `json:"expiration"`
	Keyring             string `json:"keyring"`
	RoleAccessKeyID     string `json:"role.access.key.id"`
	RoleAccessKeySecret string `json:"role.access.key.secret"`
	RoleArn             string `json:"role.arn"`
}

// ManageTokens 定义资源账号 和 角色扮演账号
type ManageTokens struct {
	AccessKeyID     string
	AccessKeySecret string
	SecurityToken   string
	Expiration      time.Time

	RoleAccessKeyID     string
	RoleAccessKeySecret string
	RoleArn             string
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func Decrypt(s string, keyring []byte) ([]byte, error) {
	cdata, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64: %w", err)
	}
	block, err := aes.NewCipher(keyring)
	if err != nil {
		return nil, fmt.Errorf("failed to new cipher: %w", err)
	}
	blockSize := block.BlockSize()

	iv := cdata[:blockSize]
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(cdata)-blockSize)

	blockMode.CryptBlocks(origData, cdata[blockSize:])

	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func ParseManagedTokens(content []byte) (ManageTokens, error) {
	var akInfo AKInfo
	err := json.Unmarshal(content, &akInfo)
	if err != nil {
		return ManageTokens{}, fmt.Errorf("error unmarshal token config: %w", err)
	}
	keyring := akInfo.Keyring
	ak, err := Decrypt(akInfo.AccessKeyID, []byte(keyring))
	if err != nil {
		return ManageTokens{}, fmt.Errorf("failed to decode AccessKeyID: %w", err)
	}

	sk, err := Decrypt(akInfo.AccessKeySecret, []byte(keyring))
	if err != nil {
		return ManageTokens{}, fmt.Errorf("failed to decode AccessKeySecret: %w", err)
	}

	token, err := Decrypt(akInfo.SecurityToken, []byte(keyring))
	if err != nil {
		return ManageTokens{}, fmt.Errorf("failed to decode SecurityToken: %w", err)
	}

	t, err := time.Parse(time.RFC3339, akInfo.Expiration)
	if err != nil {
		return ManageTokens{}, fmt.Errorf("parse expiration failed: %w", err)
	}

	tokens := ManageTokens{
		AccessKeyID:     string(ak),
		AccessKeySecret: string(sk),
		SecurityToken:   string(token),
		Expiration:      t,
	}

	if akInfo.RoleAccessKeyID != "" && akInfo.RoleAccessKeySecret != "" {
		roleAK, err := Decrypt(akInfo.RoleAccessKeyID, []byte(keyring))
		if err != nil {
			return tokens, fmt.Errorf("failed to decode RoleAccessKeyID: %w", err)
		}
		roleSK, err := Decrypt(akInfo.RoleAccessKeySecret, []byte(keyring))
		if err != nil {
			return tokens, fmt.Errorf("failed to decode RoleAccessKeySecret: %w", err)
		}
		tokens.RoleAccessKeyID = string(roleAK)
		tokens.RoleAccessKeySecret = string(roleSK)
	}
	tokens.RoleArn = akInfo.RoleArn
	return tokens, nil
}

type ManagedTokenCredential struct {
	tokens     ManageTokens
	tokenPath  string
	reloadTime time.Time
	clock      clock.PassiveClock
}

func (m *ManagedTokenCredential) reloadToken() error {
	encodedTokens, err := os.ReadFile(m.tokenPath)
	if err != nil {
		return fmt.Errorf("failed to read token config: %w", err)
	}

	tokens, err := ParseManagedTokens(encodedTokens)
	if err != nil {
		return err
	}
	m.tokens = tokens
	m.reloadTime = m.clock.Now()

	if m.tokens.Expiration.Before(m.reloadTime) {
		logrus.Warnf("token already expired at %s on load", m.tokens.Expiration)
	}

	return nil
}

func (m *ManagedTokenCredential) reloadTokenIfNeeded() error {
	now := m.clock.Now()
	// reload token every 5 minutes, or 10 minutes before expiration
	if now.After(m.reloadTime.Add(5*time.Minute)) || m.tokens.Expiration.Before(now.Add(10*time.Minute)) {
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

func NewManagedTokenCredential(clk clock.PassiveClock, tokenPath string) (alicred.Credential, error) {
	c := &ManagedTokenCredential{
		clock:     clk,
		tokenPath: tokenPath,
	}
	err := c.reloadToken()
	if err != nil {
		return nil, err
	}
	return &oldCredentialAdaptor{c}, nil
}

func (m *ManagedTokenCredential) GetType() *string {
	return tea.String("sts")
}

func (m *ManagedTokenCredential) GetCredential() (*alicred.CredentialModel, error) {
	err := m.reloadTokenIfNeeded()
	if err != nil {
		return nil, err
	}
	return &alicred.CredentialModel{
		Type:            m.GetType(),
		AccessKeyId:     &m.tokens.AccessKeyID,
		AccessKeySecret: &m.tokens.AccessKeySecret,
		SecurityToken:   &m.tokens.SecurityToken,
	}, nil
}
