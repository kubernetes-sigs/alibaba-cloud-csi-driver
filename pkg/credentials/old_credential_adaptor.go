package credentials

import (
	alicred "github.com/aliyun/credentials-go/credentials"
)

type newCredential interface {
	GetType() *string
	GetCredential() (*alicred.CredentialModel, error)
}

type oldCredentialAdaptor struct {
	newCredential
}

func (o *oldCredentialAdaptor) GetAccessKeyId() (*string, error) {
	c, err := o.GetCredential()
	if err != nil {
		return nil, err
	}
	return c.AccessKeyId, nil
}

func (o *oldCredentialAdaptor) GetAccessKeySecret() (*string, error) {
	c, err := o.GetCredential()
	if err != nil {
		return nil, err
	}
	return c.AccessKeySecret, nil
}

func (o *oldCredentialAdaptor) GetSecurityToken() (*string, error) {
	c, err := o.GetCredential()
	if err != nil {
		return nil, err
	}
	return c.SecurityToken, nil
}

func (o *oldCredentialAdaptor) GetBearerToken() *string {
	c, err := o.GetCredential()
	if err != nil {
		panic(err)
	}
	return c.BearerToken
}
