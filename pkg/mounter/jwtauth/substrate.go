//go:build !windows

package jwtauth

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net/http"
	"os"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
)

func buildSubstrateHTTPClient() (*http.Client, error) {
	tlsConfig, err := substrateTLSConfig(agentidentity.PodCertBundleFile, agentidentity.PodCertTrustBundleFile)
	if err != nil {
		return nil, err
	}
	return &http.Client{
		Timeout:       httpTimeout,
		Transport:     &http.Transport{TLSClientConfig: tlsConfig},
		CheckRedirect: rejectRedirect,
	}, nil
}

func substrateTLSConfig(certFile, trustFile string) (*tls.Config, error) {
	_, err := tls.LoadX509KeyPair(certFile, certFile)
	if err != nil {
		return nil, fmt.Errorf("load PodCertificate %s: %w", certFile, err)
	}
	config := &tls.Config{
		GetClientCertificate: func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			cert, err := tls.LoadX509KeyPair(certFile, certFile)
			if err != nil {
				return nil, fmt.Errorf("reload PodCertificate %s: %w", certFile, err)
			}
			return &cert, nil
		},
	}
	caPEM, err := os.ReadFile(trustFile)
	if err != nil {
		return nil, fmt.Errorf("read trust bundle %s: %w", trustFile, err)
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(caPEM) {
		return nil, fmt.Errorf("parse trust bundle %s: no valid certificate found", trustFile)
	}
	config.RootCAs = pool
	return config, nil
}

func rejectRedirect(req *http.Request, via []*http.Request) error {
	return fmt.Errorf("credential endpoint attempted redirect to %s, rejected to prevent token leakage", req.URL)
}
