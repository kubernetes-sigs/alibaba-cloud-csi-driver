//go:build linux && integration

package jwtauth

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type integrationSink struct {
	apply func(*STSToken) error
}

func (s integrationSink) Apply(credential *STSToken) error { return s.apply(credential) }
func (integrationSink) Cleanup()                           {}

func TestSubstrateExchangeAndRefreshIntegration(t *testing.T) {
	if os.Getenv("CSI_JWTAUTH_CONTAINER_TEST") != "1" {
		t.Skip("requires an isolated disposable container; run hack/check-jwtauth-integration.sh")
	}
	paths := []string{agentidentity.SATokenFile, agentidentity.PodCertBundleFile, agentidentity.PodCertTrustBundleFile}
	for _, path := range paths {
		_, err := os.Lstat(path)
		require.True(t, os.IsNotExist(err), "refusing to replace existing credential file %s", path)
		require.NoError(t, os.MkdirAll(filepath.Dir(path), 0700))
	}
	t.Cleanup(func() {
		for _, path := range paths {
			if err := os.Remove(path); err != nil && !os.IsNotExist(err) {
				t.Errorf("remove integration credential file: %v", err)
			}
		}
	})
	first, firstCert := testPodCertificate(t, 1)
	second, secondCert := testPodCertificate(t, 2)
	clientRoots := x509.NewCertPool()
	require.True(t, clientRoots.AppendCertsFromPEM(firstCert))
	require.True(t, clientRoots.AppendCertsFromPEM(secondCert))

	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "GetResourceCredential", r.Header.Get("X-Api-Action-Name"))
		var body map[string]string
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("decode request: %v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		assert.Len(t, body, 3)
		assert.Equal(t, "stsToken", body["credentialType"])
		generation := "legacy"
		if body["credentialProviderName"] == "legacy-provider" {
			assert.Equal(t, "legacy-sandbox-client", body["resourceId"])
			assert.Equal(t, "Bearer legacy-access-token", r.Header.Get("Authorization"))
			assert.Empty(t, r.TLS.PeerCertificates)
		} else {
			assert.Equal(t, "nas-integration", body["credentialProviderName"])
			assert.Equal(t, "actor-uid", body["resourceId"])
			if !assert.Len(t, r.TLS.PeerCertificates, 1) {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			generation = r.TLS.PeerCertificates[0].SerialNumber.String()
			assert.Contains(t, []string{"1", "2"}, generation)
			assert.Equal(t, "Bearer service-account-"+generation, r.Header.Get("Authorization"))
		}
		response := credentialResponse{STSToken: &STSToken{
			AccessKeyID:     "STS.integration-" + generation,
			AccessKeySecret: "test-secret-" + generation,
			SecurityToken:   "test-security-token-" + generation,
			Expiration:      time.Now().Add(time.Hour).UTC().Format(time.RFC3339),
		}}
		w.Header().Set("Content-Type", "application/json")
		assert.NoError(t, json.NewEncoder(w).Encode(response))
	}))
	server.TLS = &tls.Config{ClientAuth: tls.VerifyClientCertIfGiven, ClientCAs: clientRoots}
	server.Config.SetKeepAlivesEnabled(false)
	server.StartTLS()
	defer server.Close()
	require.NoError(t, os.WriteFile(agentidentity.PodCertBundleFile, first, 0600))
	require.NoError(t, os.WriteFile(agentidentity.SATokenFile, []byte("service-account-1\n"), 0600))
	require.NoError(t, os.WriteFile(agentidentity.PodCertTrustBundleFile,
		pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600))

	t.Run("ACS remains legacy with PodCertificate present", func(t *testing.T) {
		opts := ResolveOpts(map[string]string{
			OptSandboxId:    "legacy-sandbox-id",
			OptTokenFile:    writeTokenFile(t, t.TempDir(), "legacy-access-token", "legacy-sandbox-client"),
			OptEndpoint:     server.URL,
			OptCredProvider: "legacy-provider",
			OptCAFile:       agentidentity.PodCertTrustBundleFile,
		})
		require.False(t, opts.SubstrateMode)
		credential, err := FetchSTSToken(context.Background(), opts)
		require.NoError(t, err)
		assert.Equal(t, "STS.integration-legacy", credential.AccessKeyID)
	})

	t.Run("Substrate uses projected credentials and reloads both", func(t *testing.T) {
		opts := ResolveOpts(map[string]string{
			OptSubstrateMode: "true",
			OptSandboxId:     "actor-uid",
			OptEndpoint:      server.URL,
			OptCredProvider:  "nas-integration",
			OptTokenFile:     "/legacy-token-must-not-be-read",
		})
		require.True(t, opts.SubstrateMode)
		require.NoError(t, opts.Validate())
		rotated := make(chan *STSToken, 8)
		sink := integrationSink{apply: func(credential *STSToken) error {
			if credential.AccessKeyID == "STS.integration-1" {
				for path, data := range map[string][]byte{
					agentidentity.PodCertBundleFile: second,
					agentidentity.SATokenFile:       []byte("service-account-2\n"),
				} {
					if err := os.WriteFile(path+".next", data, 0600); err != nil {
						return err
					}
					if err := os.Rename(path+".next", path); err != nil {
						return err
					}
				}
				return nil
			}
			if credential.AccessKeyID != "STS.integration-2" {
				return fmt.Errorf("unexpected rotated credential")
			}
			select {
			case rotated <- credential:
			default:
			}
			return nil
		}}
		refresher := NewRefresher(opts, sink)
		refresher.refreshMargin = 2 * time.Hour
		refresher.minSleep = time.Millisecond
		require.NoError(t, refresher.Start(context.Background()))
		defer refresher.Stop()
		assert.False(t, refresher.client.Transport.(*http.Transport).TLSClientConfig.InsecureSkipVerify)
		select {
		case credential := <-rotated:
			assert.Equal(t, "STS.integration-2", credential.AccessKeyID)
			assert.Equal(t, "test-secret-2", credential.AccessKeySecret)
			assert.Equal(t, "test-security-token-2", credential.SecurityToken)
		case <-time.After(5 * time.Second):
			t.Fatal("no rotated credential delivered to the sink")
		}
	})
}
