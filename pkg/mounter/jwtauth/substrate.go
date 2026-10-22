//go:build !windows

package jwtauth

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var substrateKubeClient kubernetes.Interface

func SetSubstrateKubeClient(c kubernetes.Interface) {
	substrateKubeClient = c
}

func buildSubstrateHTTPClient() (*http.Client, error) {
	tlsConfig, err := substrateTLSConfig(agentidentity.PodCertBundleFile, agentidentity.PodCertTrustBundleFile)
	if err != nil {
		return nil, err
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	if substrateKubeClient != nil {
		transport.DialContext = serviceDialContext(substrateKubeClient)
	}
	return &http.Client{
		Timeout:       httpTimeout,
		Transport:     transport,
		CheckRedirect: rejectRedirect,
	}, nil
}

func serviceDialContext(kubeClient kubernetes.Interface) func(ctx context.Context, network, addr string) (net.Conn, error) {
	return func(ctx context.Context, network, addr string) (net.Conn, error) {
		host, port, err := net.SplitHostPort(addr)
		if err != nil {
			return (&net.Dialer{}).DialContext(ctx, network, addr)
		}
		if name, namespace, ok := splitServiceHost(host); ok {
			ip, err := resolveServiceIP(ctx, kubeClient, namespace, name, port)
			if err != nil {
				return nil, err
			}
			addr = net.JoinHostPort(ip, port)
		}
		return (&net.Dialer{}).DialContext(ctx, network, addr)
	}
}

func splitServiceHost(host string) (name, namespace string, ok bool) {
	parts := strings.Split(host, ".")
	if len(parts) < 3 || parts[2] != "svc" {
		return "", "", false
	}
	return parts[0], parts[1], true
}

func resolveServiceIP(ctx context.Context, kubeClient kubernetes.Interface, namespace, name, port string) (string, error) {
	svc, err := kubeClient.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		return "", fmt.Errorf("resolve Service %s/%s: %w", namespace, name, err)
	}
	if svc.Spec.ClusterIP != "" && svc.Spec.ClusterIP != "None" {
		return svc.Spec.ClusterIP, nil
	}
	slices, err := kubeClient.DiscoveryV1().EndpointSlices(namespace).List(ctx, metav1.ListOptions{
		LabelSelector: "kubernetes.io/service-name=" + name,
	})
	if err != nil {
		return "", fmt.Errorf("list EndpointSlices for headless Service %s/%s: %w", namespace, name, err)
	}
	for _, slice := range slices.Items {
		for _, ep := range slice.Endpoints {
			if ep.Conditions.Ready != nil && !*ep.Conditions.Ready {
				continue
			}
			if len(ep.Addresses) > 0 {
				return ep.Addresses[0], nil
			}
		}
	}
	return "", fmt.Errorf("no ready endpoints for headless Service %s/%s", namespace, name)
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
