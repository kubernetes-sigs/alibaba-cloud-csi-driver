package credentials

import (
	"errors"
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	alicred "github.com/aliyun/credentials-go/credentials"
	"github.com/stretchr/testify/assert"
)

type fakeCredential struct{}

func (f fakeCredential) GetCredential() (*alicred.CredentialModel, error) {
	return &alicred.CredentialModel{
		Type:            tea.String("fake"),
		AccessKeyId:     tea.String("fake-accessKeyId"),
		AccessKeySecret: tea.String("fake-accessKeySecret"),
		SecurityToken:   tea.String("fake-securityToken"),
		BearerToken:     tea.String("fake-bearerToken"),
	}, nil
}

func (f fakeCredential) GetType() *string {
	return tea.String("fake")
}

func TestAdaptor(t *testing.T) {
	cred := fakeCredential{}
	adaptor := oldCredentialAdaptor{cred}

	accessKeyId, err := adaptor.GetAccessKeyId()
	assert.NoError(t, err)
	assert.Equal(t, "fake-accessKeyId", *accessKeyId)

	accessKeySecret, err := adaptor.GetAccessKeySecret()
	assert.NoError(t, err)
	assert.Equal(t, "fake-accessKeySecret", *accessKeySecret)

	securityToken, err := adaptor.GetSecurityToken()
	assert.NoError(t, err)
	assert.Equal(t, "fake-securityToken", *securityToken)

	bearerToken := adaptor.GetBearerToken()
	assert.Equal(t, "fake-bearerToken", *bearerToken)
}

type errorCredential struct{}

var errCred = errors.New("error credential")

func (e errorCredential) GetCredential() (*alicred.CredentialModel, error) {
	return nil, errCred
}

func (e errorCredential) GetType() *string {
	return tea.String("fake")
}

func TestAdaptorError(t *testing.T) {
	cred := errorCredential{}
	adaptor := oldCredentialAdaptor{cred}

	_, err := adaptor.GetAccessKeyId()
	assert.Equal(t, errCred, err)

	_, err = adaptor.GetAccessKeySecret()
	assert.Equal(t, errCred, err)

	_, err = adaptor.GetSecurityToken()
	assert.Equal(t, errCred, err)

	assert.Panics(t, func() {
		_ = adaptor.GetBearerToken()
	})
}
