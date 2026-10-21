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
	if _, err := loadTrustBundle(trustFile); err != nil {
		return nil, err
	}
	return &tls.Config{
		GetClientCertificate: func(*tls.CertificateRequestInfo) (*tls.Certificate, error) {
			cert, err := tls.LoadX509KeyPair(certFile, certFile)
			if err != nil {
				return nil, fmt.Errorf("reload PodCertificate %s: %w", certFile, err)
			}
			return &cert, nil
		},
		VerifyConnection: func(cs tls.ConnectionState) error {
			pool, err := loadTrustBundle(trustFile)
			if err != nil {
				return err
			}
			opts := x509.VerifyOptions{
				Roots:         pool,
				DNSName:       cs.ServerName,
				Intermediates: x509.NewCertPool(),
			}
			for _, cert := range cs.PeerCertificates[1:] {
				opts.Intermediates.AddCert(cert)
			}
			_, err = cs.PeerCertificates[0].Verify(opts)
			return err
		},
		InsecureSkipVerify: true,
	}, nil
}

func loadTrustBundle(trustFile string) (*x509.CertPool, error) {
	caPEM, err := os.ReadFile(trustFile)
	if err != nil {
		return nil, fmt.Errorf("read trust bundle %s: %w", trustFile, err)
	}
	pool := x509.NewCertPool()
	if !pool.AppendCertsFromPEM(caPEM) {
		return nil, fmt.Errorf("parse trust bundle %s: no valid certificate found", trustFile)
	}
	return pool, nil
}

func rejectRedirect(req *http.Request, via []*http.Request) error {
	return fmt.Errorf("credential endpoint attempted redirect to %s, rejected to prevent token leakage", req.URL)
}
