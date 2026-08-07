// Package jwtauth implements the jwtauth credential flow shared by the
// mount-proxy drivers: it exchanges a sandbox jwtauth token for a scoped STS
// credential and keeps it fresh for the lifetime of a mount. Credential
// delivery is pluggable via CredentialSink so each driver can consume the
// credential in its native form (files on disk, exec-based refresh, ...).
package jwtauth

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const (
	httpTimeout            = 10 * time.Second
	apiActionGetCredential = "GetResourceCredential"
	credentialTypeSTSToken = "stsToken"
)

// Opts is the resolved configuration for a jwtauth mount.
type Opts struct {
	TokenFile    string
	Endpoint     string
	CredProvider string
	CAFile       string
	SandboxId    string
}

// Validate checks that all settings required for the credential exchange are
// present.
func (o Opts) Validate() error {
	if o.SandboxId == "" {
		return fmt.Errorf("sandboxId is required")
	}
	if o.CredProvider == "" {
		return fmt.Errorf("credential provider name is required")
	}
	if o.TokenFile == "" {
		return fmt.Errorf("token file path could not be resolved")
	}
	if o.Endpoint == "" {
		return fmt.Errorf("endpoint could not be resolved")
	}
	return nil
}

type tokenFileContent struct {
	RequestID             string `json:"requestId"`
	AccessToken           string `json:"accessToken"`
	SandboxClientID       string `json:"sandboxClientId"`
	AccessTokenExpiration string `json:"accessTokenExpiration"`
}

type credentialRequest struct {
	CredentialType         string `json:"credentialType"`
	ResourceID             string `json:"resourceId"`
	CredentialProviderName string `json:"credentialProviderName"`
}

// STSToken is the scoped credential returned by the jwtauth credential
// provider.
type STSToken struct {
	AccessKeyID     string `json:"accessKeyId"`
	AccessKeySecret string `json:"accessKeySecret"`
	SecurityToken   string `json:"securityToken"`
	Expiration      string `json:"expiration"`
}

type credentialResponse struct {
	RequestID string    `json:"requestId"`
	STSToken  *STSToken `json:"stsToken"`
}

// buildHTTPClient builds the HTTP client used to exchange the jwtauth token
// for STS credentials. This is a security-sensitive channel (it carries
// AK/SK), so TLS verification is never disabled: when a CA file is configured
// it must be readable and parsable, otherwise it fails; when no CA file is
// configured the system root pool is used (tls.Config.RootCAs == nil).
func buildHTTPClient(caFile string) (*http.Client, error) {
	tlsConfig := &tls.Config{}
	if caFile != "" {
		caCert, err := os.ReadFile(caFile)
		if err != nil {
			return nil, fmt.Errorf("read CA file %s: %w", caFile, err)
		}
		pool := x509.NewCertPool()
		if !pool.AppendCertsFromPEM(caCert) {
			return nil, fmt.Errorf("parse CA file %s: no valid certificate found", caFile)
		}
		tlsConfig.RootCAs = pool
	}
	return &http.Client{
		Timeout:   httpTimeout,
		Transport: &http.Transport{TLSClientConfig: tlsConfig},
	}, nil
}

func readTokenFile(tokenFile string) (*tokenFileContent, error) {
	data, err := os.ReadFile(tokenFile)
	if err != nil {
		return nil, fmt.Errorf("read token file %s: %w", tokenFile, err)
	}
	var token tokenFileContent
	if err := json.Unmarshal(data, &token); err != nil {
		return nil, fmt.Errorf("parse token file: %w", err)
	}
	if token.AccessToken == "" {
		return nil, fmt.Errorf("token file has empty accessToken")
	}
	if token.SandboxClientID == "" {
		return nil, fmt.Errorf("token file has empty sandboxClientId")
	}
	return &token, nil
}

// FetchSTSToken performs a one-shot, stateless exchange of the jwtauth token
// for an STS credential. It builds its own HTTP client (honoring the CA
// settings) and never touches the filesystem beyond reading the configured
// token file. Intended for consumers that need the initial credential up
// front (e.g. to inject it into mount options) before starting a Refresher.
func FetchSTSToken(ctx context.Context, opts Opts) (*STSToken, error) {
	client, err := buildHTTPClient(opts.CAFile)
	if err != nil {
		return nil, fmt.Errorf("build http client: %w", err)
	}
	return exchangeSTSToken(ctx, client, opts)
}

func exchangeSTSToken(ctx context.Context, client *http.Client, opts Opts) (*STSToken, error) {
	token, err := readTokenFile(opts.TokenFile)
	if err != nil {
		return nil, err
	}

	reqBody := credentialRequest{
		CredentialType:         credentialTypeSTSToken,
		ResourceID:             token.SandboxClientID,
		CredentialProviderName: opts.CredProvider,
	}
	bodyBytes, err := json.Marshal(reqBody)
	if err != nil {
		return nil, fmt.Errorf("marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, opts.Endpoint, bytes.NewReader(bodyBytes))
	if err != nil {
		return nil, fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token.AccessToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Api-Action-Name", apiActionGetCredential)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("credential request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("credential endpoint returned %d: %s", resp.StatusCode, string(respBody))
	}

	var credResp credentialResponse
	if err := json.Unmarshal(respBody, &credResp); err != nil {
		return nil, fmt.Errorf("parse credential response: %w", err)
	}
	if credResp.STSToken == nil {
		return nil, fmt.Errorf("credential response has nil stsToken")
	}
	if credResp.STSToken.AccessKeyID == "" || credResp.STSToken.AccessKeySecret == "" {
		return nil, fmt.Errorf("credential response has empty credentials")
	}
	return credResp.STSToken, nil
}
