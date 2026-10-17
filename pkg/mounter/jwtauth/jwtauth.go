// Package jwtauth implements the jwtauth credential flow shared by the
// mount-proxy drivers: it exchanges pod or ACS sandbox identity for a scoped STS
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

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
	"k8s.io/klog/v2"
)

const (
	httpTimeout            = 10 * time.Second
	apiActionGetCredential = "GetResourceCredential"
	credentialTypeSTSToken = "stsToken"
)

// Opts is the resolved configuration for a jwtauth mount.
type Opts struct {
	TokenFile    string // ACS sandbox JSON identity file.
	Endpoint     string
	CredProvider string
	CAFile       string // Optional server CA for ACS sandbox mode.
	// SandboxId is the actor UID in Substrate mode, or the sandbox ID in ACS mode.
	SandboxId string

	// SubstrateMode comes from the explicit request marker, not credential-file presence.
	SubstrateMode bool
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
	if o.Endpoint == "" {
		return fmt.Errorf("endpoint could not be resolved")
	}
	if !o.SubstrateMode && o.TokenFile == "" {
		return fmt.Errorf("token file path could not be resolved")
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
// for STS credentials. Substrate mode uses PodCertificate client authentication
// and its separately mounted server trust bundle. In ACS sandbox mode, a
// configured CA file must be readable and parsable, otherwise it fails.
//
// In ACS sandbox mode without a CA file, verification is skipped rather than
// deferred to the system root pool. The AgentIdentity endpoint is an in-cluster service holding a
// private certificate that no public root signs, so the system pool can only
// accept it while something injects that private CA into the process-wide trust
// store — and such an injection also hides every public root from the same
// process, breaking unrelated OSS connections. Skipping verification keeps the
// two concerns apart, matching how ossfs treats an empty agent_identity_ca_file.
func buildHTTPClient(opts Opts) (*http.Client, error) {
	tlsConfig := &tls.Config{}

	if opts.SubstrateMode {
		// Substrate mode: authenticate this Pod with its projected client certificate.
		var err error
		tlsConfig, err = substrateTLSConfig(agentidentity.PodCertBundleFile, agentidentity.PodCertTrustBundleFile)
		if err != nil {
			return nil, err
		}
	} else {
		// ACS sandbox mode: retain its configured-CA and missing-CA behavior.
		if opts.CAFile != "" {
			caCert, err := os.ReadFile(opts.CAFile)
			if err != nil {
				return nil, fmt.Errorf("read CA file %s: %w", opts.CAFile, err)
			}
			pool := x509.NewCertPool()
			if !pool.AppendCertsFromPEM(caCert) {
				return nil, fmt.Errorf("parse CA file %s: no valid certificate found", opts.CAFile)
			}
			tlsConfig.RootCAs = pool
		} else {
			klog.Warningf("no agent identity CA file configured, skipping TLS verification for the AgentIdentity endpoint")
			tlsConfig.InsecureSkipVerify = true
		}
	}
	return &http.Client{
		Timeout:   httpTimeout,
		Transport: &http.Transport{TLSClientConfig: tlsConfig},
	}, nil
}

// substrateTLSConfig configures both PodCertificate client authentication and
// the independent trust bundle used to verify the credential-provider server.
func substrateTLSConfig(certFile, trustFile string) (*tls.Config, error) {
	// Validate the projected client identity before accepting the mount setup.
	_, err := tls.LoadX509KeyPair(certFile, certFile)
	if err != nil {
		return nil, fmt.Errorf("load PodCertificate %s: %w", certFile, err)
	}
	config := &tls.Config{
		// Reload the PodCertificate on each new TLS handshake to pick up rotation.
		GetClientCertificate: func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			cert, err := tls.LoadX509KeyPair(certFile, certFile)
			if err != nil {
				return nil, fmt.Errorf("reload PodCertificate %s: %w", certFile, err)
			}
			return &cert, nil
		},
	}
	// The server trust root need not be the issuer of the client PodCertificate.
	caPEM, err := os.ReadFile(trustFile)
	if err != nil {
		return nil, fmt.Errorf("read PodCertificate trust bundle %s: %w", trustFile, err)
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(caPEM) {
		return nil, fmt.Errorf("parse PodCertificate trust bundle %s: no valid certificate found", trustFile)
	}
	config.RootCAs = pool
	return config, nil
}

// readTokenFile reads the ACS sandbox identity document, not the Pod SA token.
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
// settings) and only reads the selected mode's token and TLS credential files.
// Intended for consumers that need the initial credential up
// front (e.g. to inject it into mount options) before starting a Refresher.
func FetchSTSToken(ctx context.Context, opts Opts) (*STSToken, error) {
	client, err := buildHTTPClient(opts)
	if err != nil {
		return nil, fmt.Errorf("build http client: %w", err)
	}
	return exchangeSTSToken(ctx, client, opts)
}

func exchangeSTSToken(ctx context.Context, client *http.Client, opts Opts) (*STSToken, error) {
	var bearerToken, resourceID string

	if opts.SubstrateMode {
		// Substrate mode: the Pod SA token is the bearer and the actor UID is the resource.
		saToken, err := os.ReadFile(agentidentity.SATokenFile)
		if err != nil {
			return nil, fmt.Errorf("read SA token file %s: %w", agentidentity.SATokenFile, err)
		}
		bearerToken = string(bytes.TrimSpace(saToken))
		resourceID = opts.SandboxId
	} else {
		// ACS sandbox mode: both values come from the sandbox identity token file.
		token, err := readTokenFile(opts.TokenFile)
		if err != nil {
			return nil, err
		}
		bearerToken = token.AccessToken
		resourceID = token.SandboxClientID
	}

	reqBody := credentialRequest{
		CredentialType:         credentialTypeSTSToken,
		ResourceID:             resourceID,
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
	req.Header.Set("Authorization", "Bearer "+bearerToken)
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
