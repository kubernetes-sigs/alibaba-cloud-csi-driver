package credentials

import (
	"errors"
	"os"

	alicred_old "github.com/aliyun/credentials-go/credentials"
	alicred "github.com/aliyun/credentials-go/credentials/providers"
	"k8s.io/utils/clock"
)

func NewProvider() (alicred.CredentialsProvider, error) {
	// try managed token credential
	provider, err := NewManagedTokenProvider(clock.RealClock{}, GetManagedTokenPath())
	if err == nil {
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
