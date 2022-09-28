package auth

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
	"time"
)

var _ Provider = &oidcProvider{}

type Provider interface {
	AssumeRoleWithOIDC() (*sts.AssumeRoleWithOIDCResponse, error)
	NewClientWithStsToken(response *sts.AssumeRoleWithOIDCResponse) (*Client, error)
	GetStsTokenWithCache() (*sts.AssumeRoleWithOIDCResponse, error)
	GetClient() (*Client, error)
}

type oidcProvider struct{}

// NewOIDCProvider create a new OIDC Provider
func NewOIDCProvider(region, appName, providerName, roleName, UUID, network string, durationSeconds time.Duration) *oidcProvider {
	return nil
}

// NewOIDCProviderVPC create a new OIDC Provider with VPC network
func NewOIDCProviderVPC(region, appName, providerName, roleName, UUID string, durationSeconds time.Duration) *oidcProvider {
	return nil
}

// AssumeRoleWithOIDC assume role with OIDC provider, return sts token
func (p *oidcProvider) AssumeRoleWithOIDC() (*sts.AssumeRoleWithOIDCResponse, error) {
	return nil, nil
}

// GetStsTokenWithCache get sts token with OIDC provider, return assume role response
func (p *oidcProvider) GetStsTokenWithCache() (*sts.AssumeRoleWithOIDCResponse, error) {
	return nil, nil
}

// NewClientWithStsToken create a new alibaba cloud client with sts token
func (p *oidcProvider) NewClientWithStsToken(response *sts.AssumeRoleWithOIDCResponse) (*Client, error) {
	return nil, nil
}

// GetClient get the alibaba cloud client, Note: it automatically refreshes the client
func (p *oidcProvider) GetClient() (*Client, error) {
	return nil, nil
}

func (p *oidcProvider) refresh() {}

func (p *oidcProvider) refreshClient() {}
