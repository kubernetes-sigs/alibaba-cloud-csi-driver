package credentials

import (
	"errors"
	"os"

	alicred_v1 "github.com/aliyun/alibaba-cloud-sdk-go/sdk/auth/credentials"
	alicred_old "github.com/aliyun/credentials-go/credentials"
	alicred "github.com/aliyun/credentials-go/credentials/providers"
	"k8s.io/klog/v2"
	"k8s.io/utils/clock"
)

const (
	EnvVarAccessKeyId     = "ALIBABA_CLOUD_ACCESS_KEY_ID"
	EnvVarAccessKeySecret = "ALIBABA_CLOUD_ACCESS_KEY_SECRET"
)

func MigrateEnv() {
	accessKeyID := os.Getenv("ACCESS_KEY_ID")
	accessSecret := os.Getenv("ACCESS_KEY_SECRET")
	if accessKeyID != "" || accessSecret != "" {
		klog.Warningf("environment variable ACCESS_KEY_ID and ACCESS_KEY_SECRET is deprecated, please use %s and %s instead.", EnvVarAccessKeyId, EnvVarAccessKeySecret)

		// if ALIBABA_CLOUD_* not set, use ACCESS_KEY_ID and ACCESS_KEY_SECRET
		if os.Getenv(EnvVarAccessKeyId) == "" && accessKeyID != "" {
			os.Setenv(EnvVarAccessKeyId, accessKeyID)
		}
		if os.Getenv(EnvVarAccessKeySecret) == "" && accessSecret != "" {
			os.Setenv(EnvVarAccessKeySecret, accessSecret)
		}
	}
}

func NewProvider() (alicred.CredentialsProvider, error) {
	if os.Getenv("USE_OIDC_AUTH_INNER") == "true" {
		return nil, errors.New("USE_OIDC_AUTH_INNER is no longer supported")
	}

	MigrateEnv()

	// try managed token credential
	provider, err := NewManagedTokenProvider(clock.RealClock{}, GetManagedTokenPath())
	if err == nil {
		klog.V(2).Info("using managed token credential")
		return provider, nil
	}
	if !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	return alicred.NewDefaultCredentialsProvider(), nil
}

func NewCredential() (alicred_old.Credential, error) {
	provider, err := NewProvider()
	if err != nil {
		return nil, err
	}
	return alicred_old.FromCredentialsProvider("default", provider), nil
}

type v1Provider struct {
	alicred.CredentialsProvider
}

func (p v1Provider) GetCredentials() (*alicred_v1.Credentials, error) {
	cred, err := p.CredentialsProvider.GetCredentials()
	if err != nil {
		return nil, err
	}
	return &alicred_v1.Credentials{
		AccessKeyId:     cred.AccessKeyId,
		AccessKeySecret: cred.AccessKeySecret,
		SecurityToken:   cred.SecurityToken,
		ProviderName:    cred.ProviderName,
	}, nil
}

func V1ProviderAdaptor(provider alicred.CredentialsProvider) alicred_v1.CredentialsProvider {
	return v1Provider{provider}
}
