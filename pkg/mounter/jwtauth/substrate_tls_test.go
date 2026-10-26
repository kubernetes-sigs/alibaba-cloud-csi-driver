package jwtauth

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"io"
	"math/big"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func testPodCertificate(t *testing.T, serial int64) (bundle, certPEM []byte) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	template := &x509.Certificate{
		SerialNumber: big.NewInt(serial),
		NotBefore:    time.Now().Add(-time.Minute),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	require.NoError(t, err)
	keyDER, err := x509.MarshalPKCS8PrivateKey(key)
	require.NoError(t, err)
	certPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
	bundle = append(pem.EncodeToMemory(&pem.Block{Type: "PRIVATE KEY", Bytes: keyDER}), certPEM...)
	return bundle, certPEM
}

func TestSubstrateTLSReloadsProjectedCertificate(t *testing.T) {
	first, firstCert := testPodCertificate(t, 1)
	second, secondCert := testPodCertificate(t, 2)
	clientRoots := x509.NewCertPool()
	require.True(t, clientRoots.AppendCertsFromPEM(firstCert))
	require.True(t, clientRoots.AppendCertsFromPEM(secondCert))
	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := io.WriteString(w, r.TLS.PeerCertificates[0].SerialNumber.String())
		assert.NoError(t, err)
	}))
	server.TLS = &tls.Config{ClientAuth: tls.RequireAndVerifyClientCert, ClientCAs: clientRoots}
	server.Config.SetKeepAlivesEnabled(false)
	server.StartTLS()
	defer server.Close()
	dir := t.TempDir()
	certFile, trustFile := filepath.Join(dir, "credential-bundle.pem"), filepath.Join(dir, "trust-bundle.pem")
	require.NoError(t, os.WriteFile(certFile, first, 0600))
	require.NoError(t, os.WriteFile(trustFile, pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: server.Certificate().Raw}), 0600))
	config, err := substrateTLSConfig(certFile, trustFile)
	require.NoError(t, err)
	// InsecureSkipVerify is true because VerifyConnection handles CA
	// verification with a freshly loaded trust bundle on every handshake.
	// Hostname verification is included via DNSName in VerifyConnection.
	require.True(t, config.InsecureSkipVerify)
	require.NotNil(t, config.VerifyConnection)
	client := &http.Client{Transport: &http.Transport{TLSClientConfig: config}, Timeout: time.Second}
	defer client.CloseIdleConnections()
	for i, bundle := range [][]byte{first, second} {
		require.NoError(t, os.WriteFile(certFile+".new", bundle, 0600))
		require.NoError(t, os.Rename(certFile+".new", certFile))
		response, err := client.Get(server.URL)
		require.NoError(t, err)
		body, err := io.ReadAll(response.Body)
		require.NoError(t, response.Body.Close())
		require.NoError(t, err)
		assert.Equal(t, big.NewInt(int64(i+1)).String(), string(body))
	}
	require.NoError(t, os.Remove(certFile))
	_, err = client.Get(server.URL)
	require.Error(t, err, "must not reuse a stale client certificate after the projected file disappears")
}

func TestSubstrateTLSTrustBundle(t *testing.T) {
	bundle, _ := testPodCertificate(t, 1)
	dir := t.TempDir()
	certFile, trustFile := filepath.Join(dir, "credential-bundle.pem"), filepath.Join(dir, "trust-bundle.pem")
	require.NoError(t, os.WriteFile(certFile, bundle, 0600))
	t.Run("missing trust bundle fails closed", func(t *testing.T) {
		config, err := substrateTLSConfig(certFile, trustFile)
		require.ErrorIs(t, err, os.ErrNotExist)
		assert.Nil(t, config)
	})
	t.Run("unreadable trust bundle fails closed", func(t *testing.T) {
		config, err := substrateTLSConfig(certFile, dir)
		require.Error(t, err)
		assert.Nil(t, config)
	})
	t.Run("malformed present trust bundle fails closed", func(t *testing.T) {
		require.NoError(t, os.WriteFile(trustFile, []byte("invalid PEM"), 0600))
		_, err := substrateTLSConfig(certFile, trustFile)
		require.Error(t, err)
	})
}

func TestACSTLSWithoutCARetainsFallback(t *testing.T) {
	server := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	defer server.Close()
	client, err := buildHTTPClient(Opts{})
	require.NoError(t, err)
	defer client.CloseIdleConnections()
	response, err := client.Get(server.URL)
	require.NoError(t, err)
	defer func() { assert.NoError(t, response.Body.Close()) }()
	assert.Equal(t, http.StatusNoContent, response.StatusCode)
}

func TestSubstrateModeRejectsPlainHTTP(t *testing.T) {
	opts := Opts{
		SubstrateMode: true,
		Endpoint:      "http://credential-provider.svc:8443/credential",
		SandboxId:     "actor-uid",
	}
	_, err := exchangeSTSToken(t.Context(), &http.Client{}, opts)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "requires HTTPS endpoint")
}

func TestSubstrateRedirectRejected(t *testing.T) {
	err := rejectRedirect(&http.Request{URL: mustParseURL("http://evil.example.com/steal")}, nil)
	require.Error(t, err)
	assert.Contains(t, err.Error(), "redirect")
	assert.Contains(t, err.Error(), "evil.example.com")
}

func mustParseURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

func TestSubstrateTLSReloadsTrustBundle(t *testing.T) {
	// Generate two self-signed CA certs to simulate CA rotation.
	makeCA := func(serial int64) (certPEM []byte, tlsCert tls.Certificate) {
		t.Helper()
		key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		require.NoError(t, err)
		template := &x509.Certificate{
			SerialNumber:          big.NewInt(serial),
			NotBefore:             time.Now().Add(-time.Minute),
			NotAfter:              time.Now().Add(time.Hour),
			KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
			BasicConstraintsValid: true,
			IsCA:                  true,
		}
		der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
		require.NoError(t, err)
		certPEM = pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: der})
		tlsCert = tls.Certificate{
			Certificate: [][]byte{der},
			PrivateKey:  key,
		}
		return certPEM, tlsCert
	}

	ca1PEM, ca1TLS := makeCA(100)
	ca2PEM, ca2TLS := makeCA(200)

	dir := t.TempDir()
	trustFile := filepath.Join(dir, "trust-bundle.pem")

	// Client cert (PodCertificate) - just needs to exist for substrateTLSConfig.
	clientBundle, _ := testPodCertificate(t, 1)
	certFile := filepath.Join(dir, "credential-bundle.pem")
	require.NoError(t, os.WriteFile(certFile, clientBundle, 0600))

	// Start with CA1 in the trust bundle.
	require.NoError(t, os.WriteFile(trustFile, ca1PEM, 0600))

	tlsConfig, err := substrateTLSConfig(certFile, trustFile)
	require.NoError(t, err)

	// Server using CA1 certificate — should succeed.
	server1 := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	server1.TLS = &tls.Config{Certificates: []tls.Certificate{ca1TLS}}
	server1.StartTLS()
	defer server1.Close()

	client := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig.Clone()}}
	resp, err := client.Get(server1.URL)
	require.NoError(t, err)
	_ = resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	client.CloseIdleConnections()

	// Rotate trust bundle to CA2.
	require.NoError(t, os.WriteFile(trustFile, ca2PEM, 0600))

	// Server using CA2 certificate — should succeed because trust bundle is re-read.
	server2 := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	}))
	server2.TLS = &tls.Config{Certificates: []tls.Certificate{ca2TLS}}
	server2.StartTLS()
	defer server2.Close()

	client2 := &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig.Clone()}}
	resp2, err := client2.Get(server2.URL)
	require.NoError(t, err)
	_ = resp2.Body.Close()
	assert.Equal(t, http.StatusNoContent, resp2.StatusCode)
}

func TestServiceDialContext_SubstrateModeOff(t *testing.T) {
	// When SubstrateMode is false, buildHTTPClient does not call
	// buildSubstrateHTTPClient, so no DialContext is set.
	client, err := buildHTTPClient(Opts{SubstrateMode: false})
	require.NoError(t, err)
	transport := client.Transport.(*http.Transport)
	assert.Nil(t, transport.DialContext)
}

func boolPtr(v bool) *bool { return &v }
