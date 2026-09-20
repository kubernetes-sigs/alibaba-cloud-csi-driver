package nas

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/nas/internal"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/fake"
)

func TestRefreshDeadline(t *testing.T) {
	now := time.Now()
	tests := []struct {
		name string
		exp  time.Duration
		want time.Duration
	}{
		// Long-lived: refresh 20m before expiry, keeping the on-disk credential
		// valid for >=20m at any moment (>10m survives a NAS reconnect interval).
		{"long lived renews 20m before expiry", time.Hour, 40 * time.Minute},
		{"medium renews 20m before expiry", 30 * time.Minute, 10 * time.Minute},
		// At exactly twice the margin, expiry-20m lands 20m in.
		{"at twice the margin", 40 * time.Minute, 20 * time.Minute},
		// Short-lived: expiry-20m is already past at issuance, so the 30s cooldown
		// floors it; kubelet's ~1m republishes then rotate it every call.
		{"short lived floored by cooldown", 15 * time.Minute, 30 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, now.Add(tt.want), refreshDeadline(now, now.Add(tt.exp)))
		})
	}
}

func TestPodServiceAccountToken(t *testing.T) {
	tests := []struct {
		name          string
		volumeContext map[string]string
		secrets       map[string]string
		want          string
		errLike       string
	}{
		{
			// kubelet's default delivery: the token rides in the volume context.
			name: "token for our audience in the volume context",
			volumeContext: map[string]string{
				serviceAccountTokensKey: `{"sts.aliyuncs.com":{"token":"tok","expirationTimestamp":"2026-01-01T00:00:00Z"}}`,
			},
			want: "tok",
		},
		{
			// Without tokenRequests the key is simply absent; that is a
			// misconfigured CSIDriver, not a volume that may mount unauthenticated.
			// A CSIDriver with serviceAccountTokenInSecrets set delivers it here
			// instead, so both sources are accepted.
			name: "token in the publish secrets",
			secrets: map[string]string{
				serviceAccountTokensKey: `{"sts.aliyuncs.com":{"token":"tok"}}`,
			},
			want: "tok",
		},
		{
			name:    "no tokens at all",
			secrets: nil,
			errLike: "tokenRequests",
		},
		{
			name: "tokens for other audiences only",
			volumeContext: map[string]string{
				serviceAccountTokensKey: `{"oidc-proxy.alibabacloud.com":{"token":"tok"}}`,
			},
			errLike: "audience",
		},
		{
			name:          "unparsable tokens",
			volumeContext: map[string]string{serviceAccountTokensKey: `not-json`},
			errLike:       "parse ServiceAccount tokens",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := podServiceAccountToken(defaultTokenAudience, tt.volumeContext, tt.secrets)
			if tt.errLike != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errLike)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResolveEndpoint(t *testing.T) {
	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "oidc-proxy", Namespace: "oidc-proxy-system"},
		Spec:       corev1.ServiceSpec{ClusterIP: "10.0.0.7"},
	}
	headless := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "headless", Namespace: "oidc-proxy-system"},
		Spec:       corev1.ServiceSpec{ClusterIP: "None"},
	}
	e := newRRSAExchanger(fake.NewSimpleClientset(svc, headless), "cn-beijing", "1857989822569166", "c123", "", 0)

	tests := []struct {
		name           string
		endpoint       string
		wantAddress    string
		wantServerName string
		errLike        string
	}{
		{
			// The Service name stays the verified identity while we dial the IP,
			// because csi-plugin cannot resolve cluster DNS.
			name:           "service host resolves to ClusterIP, name still verified",
			endpoint:       "https://oidc-proxy.oidc-proxy-system.svc",
			wantAddress:    "https://10.0.0.7",
			wantServerName: "oidc-proxy.oidc-proxy-system.svc",
		},
		{
			name:           "port is preserved",
			endpoint:       "https://oidc-proxy.oidc-proxy-system.svc:8443/",
			wantAddress:    "https://10.0.0.7:8443/",
			wantServerName: "oidc-proxy.oidc-proxy-system.svc",
		},
		{
			name:           "cluster.local suffix",
			endpoint:       "https://oidc-proxy.oidc-proxy-system.svc.cluster.local",
			wantAddress:    "https://10.0.0.7",
			wantServerName: "oidc-proxy.oidc-proxy-system.svc.cluster.local",
		},
		{
			name:           "literal IP is passed through",
			endpoint:       "https://10.1.2.3:443",
			wantAddress:    "https://10.1.2.3:443",
			wantServerName: "10.1.2.3",
		},
		{
			name:           "non-service host is left to the resolver",
			endpoint:       "https://sts.aliyuncs.com",
			wantAddress:    "https://sts.aliyuncs.com",
			wantServerName: "sts.aliyuncs.com",
		},
		{
			// The Pod token is a bearer credential, so plaintext is refused
			// rather than downgraded.
			name:     "http is rejected",
			endpoint: "http://oidc-proxy.oidc-proxy-system.svc",
			errLike:  "must use https",
		},
		{
			name:     "missing service",
			endpoint: "https://absent.oidc-proxy-system.svc",
			errLike:  "resolve Service",
		},
		{
			name:     "headless service has nothing to dial",
			endpoint: "https://headless.oidc-proxy-system.svc",
			errLike:  "no ClusterIP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			address, serverName, err := e.resolveEndpoint(context.Background(), tt.endpoint)
			if tt.errLike != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errLike)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, tt.wantAddress, address)
			assert.Equal(t, tt.wantServerName, serverName)
		})
	}
}

// TestExchangeVerifiesProxyIdentity is the point of the SNI pinning: the driver
// connects to an IP but must still require a certificate issued for the Service
// name, signed by the configured CA. Asserting only the happy path would pass
// just as well with verification turned off, so the mismatch cases are here too.
func TestExchangeVerifiesProxyIdentity(t *testing.T) {
	const serviceHost = "oidc-proxy.oidc-proxy-system.svc"
	caPEM, serverCert := mustIssueServingCert(t, serviceHost)

	var gotForm map[string][]string
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		require.NoError(t, r.ParseForm())
		gotForm = r.Form
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{"Credentials": map[string]string{
			"AccessKeyId":     "ak",
			"AccessKeySecret": "sk",
			"SecurityToken":   "token",
			"Expiration":      time.Now().Add(time.Hour).UTC().Format(time.RFC3339),
		}})
	}))
	srv.TLS = &tls.Config{Certificates: []tls.Certificate{serverCert}}
	srv.StartTLS()
	defer srv.Close()

	_, port, err := net.SplitHostPort(srv.Listener.Addr().String())
	require.NoError(t, err)
	// The Service resolves to the loopback address the test server listens on,
	// standing in for a ClusterIP.
	clientset := fake.NewSimpleClientset(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "oidc-proxy", Namespace: "oidc-proxy-system"},
		Spec:       corev1.ServiceSpec{ClusterIP: "127.0.0.1"},
	})
	caFile := filepath.Join(t.TempDir(), "ca.crt")
	require.NoError(t, os.WriteFile(caFile, caPEM, 0o600))

	opt := &Options{
		RoleArn:         "acs:ram::1:role/test",
		OIDCProviderArn: "acs:ram::1:oidc-provider/ack-rrsa-c123",
		RRSAEndpoint:    fmt.Sprintf("https://%s:%s", serviceHost, port),
	}
	// Delivered the way kubelet does it by default: in the volume context.
	secrets := map[string]string{
		serviceAccountTokensKey: `{"sts.aliyuncs.com":{"token":"pod-token"}}`,
	}

	t.Run("credential exchanged over a verified channel", func(t *testing.T) {
		e := newRRSAExchanger(clientset, "cn-beijing", "1857989822569166", "c123", caFile, 0)
		cred, err := e.exchange(context.Background(), "vol", "/target", opt, secrets, nil)
		require.NoError(t, err)
		assert.Equal(t, "ak", cred.AkID)
		assert.Equal(t, "sk", cred.AkSecret)
		assert.Equal(t, "token", cred.SecurityToken)
		assert.WithinDuration(t, time.Now().Add(time.Hour), cred.Expiration, time.Minute)

		assert.Equal(t, []string{"AssumeRoleWithOIDC"}, gotForm["Action"])
		assert.Equal(t, []string{"pod-token"}, gotForm["OIDCToken"])
		assert.Equal(t, []string{opt.RoleArn}, gotForm["RoleArn"])
		assert.Equal(t, []string{"JSON"}, gotForm["Format"])
	})

	t.Run("unknown CA is refused", func(t *testing.T) {
		otherCA, _ := mustIssueServingCert(t, serviceHost)
		otherFile := filepath.Join(t.TempDir(), "ca.crt")
		require.NoError(t, os.WriteFile(otherFile, otherCA, 0o600))
		_, err := newRRSAExchanger(clientset, "cn-beijing", "1857989822569166", "c123", otherFile, 0).exchange(context.Background(), "vol", "/target", opt, secrets, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "certificate")
	})

	t.Run("certificate for another name is refused", func(t *testing.T) {
		e := newRRSAExchanger(clientset, "cn-beijing", "1857989822569166", "c123", caFile, 0)
		wrongName := *opt
		// Same address, different expected identity: the IP alone must not be
		// enough to satisfy verification.
		wrongName.RRSAEndpoint = fmt.Sprintf("https://impostor.oidc-proxy-system.svc:%s", port)
		_, err := clientset.CoreV1().Services("oidc-proxy-system").Create(context.Background(), &corev1.Service{
			ObjectMeta: metav1.ObjectMeta{Name: "impostor", Namespace: "oidc-proxy-system"},
			Spec:       corev1.ServiceSpec{ClusterIP: "127.0.0.1"},
		}, metav1.CreateOptions{})
		require.NoError(t, err)
		_, err = e.exchange(context.Background(), "vol", "/target", &wrongName, secrets, nil)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "impostor.oidc-proxy-system.svc")
	})
}

func TestExchangeRejectsBadResponses(t *testing.T) {
	const serviceHost = "oidc-proxy.oidc-proxy-system.svc"
	caPEM, serverCert := mustIssueServingCert(t, serviceHost)
	clientset := fake.NewSimpleClientset(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "oidc-proxy", Namespace: "oidc-proxy-system"},
		Spec:       corev1.ServiceSpec{ClusterIP: "127.0.0.1"},
	})
	caFile := filepath.Join(t.TempDir(), "ca.crt")
	require.NoError(t, os.WriteFile(caFile, caPEM, 0o600))
	// Delivered the way kubelet does it by default: in the volume context.
	secrets := map[string]string{
		serviceAccountTokensKey: `{"sts.aliyuncs.com":{"token":"pod-token"}}`,
	}

	tests := []struct {
		name    string
		status  int
		body    string
		errLike string
	}{
		{
			name:    "STS error is reported with its code",
			status:  http.StatusBadRequest,
			body:    `{"RequestId":"1","Code":"NoMountsFound","Message":"no NAS mounts"}`,
			errLike: "NoMountsFound",
		},
		{
			name:    "opaque failure still names the status",
			status:  http.StatusBadGateway,
			body:    `<html>gateway</html>`,
			errLike: "502",
		},
		{
			// A credential missing the token would mount in ID token-only mode and
			// hang, so it is rejected before it can reach mount.alinas.
			name:    "credential without a security token",
			status:  http.StatusOK,
			body:    `{"Credentials":{"AccessKeyId":"ak","AccessKeySecret":"sk"}}`,
			errLike: "incomplete credential",
		},
		{
			// Expiration is the only time in the response, so the decoder's own
			// message identifies it well enough.
			name:    "unparsable expiration",
			status:  http.StatusOK,
			body:    `{"Credentials":{"AccessKeyId":"ak","AccessKeySecret":"sk","SecurityToken":"t","Expiration":"soon"}}`,
			errLike: `parsing time "soon"`,
		},
		{
			// Without an expiration there is no rotation deadline, so the mount
			// would keep a credential until it silently stopped working.
			name:    "missing expiration",
			status:  http.StatusOK,
			body:    `{"Credentials":{"AccessKeyId":"ak","AccessKeySecret":"sk","SecurityToken":"t"}}`,
			errLike: "incomplete credential",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.status)
				_, _ = w.Write([]byte(tt.body))
			}))
			srv.TLS = &tls.Config{Certificates: []tls.Certificate{serverCert}}
			srv.StartTLS()
			defer srv.Close()
			_, port, err := net.SplitHostPort(srv.Listener.Addr().String())
			require.NoError(t, err)

			opt := &Options{
				RoleArn:         "acs:ram::1:role/test",
				OIDCProviderArn: "acs:ram::1:oidc-provider/ack-rrsa-c123",
				RRSAEndpoint:    fmt.Sprintf("https://%s:%s", serviceHost, port),
			}
			_, err = newRRSAExchanger(clientset, "cn-beijing", "1857989822569166", "c123", caFile, 0).exchange(context.Background(), "vol", "/target", opt, secrets, nil)
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errLike)
		})
	}
}

// TestCredentialForReusesUntilDue pins the property that makes republishing cheap:
// kubelet calls NodePublishVolume about once a minute per volume, and only the
// calls past the refresh point may reach the proxy.
func TestCredentialForReusesUntilDue(t *testing.T) {
	e := newRRSAExchanger(fake.NewSimpleClientset(), "cn-beijing", "1857989822569166", "c123", "", 0)
	const target = "/var/lib/kubelet/pods/uid/volumes/x"
	fresh := &rrsaCredential{
		stsCredential: stsCredential{
			AkID: "ak", AkSecret: "sk", SecurityToken: "token",
			Expiration: time.Now().Add(time.Hour),
		},
		refreshAt: time.Now().Add(30 * time.Minute),
	}
	e.creds.Store(target, fresh)

	cred, err := e.credentialFor(t.Context(), "vol", target, &Options{}, nil, nil)
	require.NoError(t, err)
	assert.Same(t, fresh, cred, "a credential at 0%% of its life must not be re-exchanged")

	// Past the refresh point the exchange is attempted; with no reachable proxy
	// it fails, which is what a caller must surface rather than skip.
	e.creds.Store(target, &rrsaCredential{
		stsCredential: stsCredential{Expiration: time.Now().Add(time.Minute)},
		refreshAt:     time.Now().Add(-time.Minute),
	})
	_, err = e.credentialFor(t.Context(), "vol", target, &Options{RRSAEndpoint: "https://absent.ns.svc"}, nil, nil)
	assert.Error(t, err)

	// forget drops the state so the next mount of the same path starts clean.
	e.forget(target)
	assert.Nil(t, e.load(target))
}

// mustIssueServingCert returns a CA certificate in PEM form and a serving
// certificate for host signed by it.
func mustIssueServingCert(t *testing.T, host string) ([]byte, tls.Certificate) {
	t.Helper()
	caKey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	caTmpl := &x509.Certificate{
		SerialNumber:          big.NewInt(1),
		Subject:               pkix.Name{CommonName: "test-ca"},
		NotBefore:             time.Now().Add(-time.Hour),
		NotAfter:              time.Now().Add(time.Hour),
		IsCA:                  true,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
		BasicConstraintsValid: true,
	}
	caDER, err := x509.CreateCertificate(rand.Reader, caTmpl, caTmpl, &caKey.PublicKey, caKey)
	require.NoError(t, err)
	caCert, err := x509.ParseCertificate(caDER)
	require.NoError(t, err)

	key, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)
	tmpl := &x509.Certificate{
		SerialNumber: big.NewInt(2),
		Subject:      pkix.Name{CommonName: host},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
		DNSNames:     []string{host},
	}
	der, err := x509.CreateCertificate(rand.Reader, tmpl, caCert, &key.PublicKey, caKey)
	require.NoError(t, err)

	caPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: caDER})
	return caPEM, tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key}
}

// TestExchangeAsksForConfiguredDuration covers the knob a cluster uses to trade
// credential lifetime against AssumeRoleWithOIDC volume: a request is sent only
// when one is configured, so an endpoint's own default still applies otherwise.
func TestExchangeAsksForConfiguredDuration(t *testing.T) {
	const serviceHost = "oidc-proxy.oidc-proxy-system.svc"
	caPEM, serverCert := mustIssueServingCert(t, serviceHost)
	clientset := fake.NewSimpleClientset(&corev1.Service{
		ObjectMeta: metav1.ObjectMeta{Name: "oidc-proxy", Namespace: "oidc-proxy-system"},
		Spec:       corev1.ServiceSpec{ClusterIP: "127.0.0.1"},
	})
	caFile := filepath.Join(t.TempDir(), "ca.crt")
	require.NoError(t, os.WriteFile(caFile, caPEM, 0o600))
	secrets := map[string]string{
		serviceAccountTokensKey: `{"sts.aliyuncs.com":{"token":"pod-token"}}`,
	}

	tests := []struct {
		name     string
		duration time.Duration
		want     []string
	}{
		{name: "unset sends nothing", duration: 0, want: nil},
		{name: "12h for a large cluster", duration: 12 * time.Hour, want: []string{"43200"}},
		{name: "15m for a test", duration: 15 * time.Minute, want: []string{"900"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got []string
			srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				require.NoError(t, r.ParseForm())
				got = r.Form["DurationSeconds"]
				_ = json.NewEncoder(w).Encode(map[string]any{"Credentials": map[string]string{
					"AccessKeyId": "ak", "AccessKeySecret": "sk", "SecurityToken": "token",
					"Expiration": time.Now().Add(time.Hour).UTC().Format(time.RFC3339),
				}})
			}))
			srv.TLS = &tls.Config{Certificates: []tls.Certificate{serverCert}}
			srv.StartTLS()
			defer srv.Close()
			_, port, err := net.SplitHostPort(srv.Listener.Addr().String())
			require.NoError(t, err)

			opt := &Options{
				RoleArn:         "acs:ram::1:role/test",
				OIDCProviderArn: "acs:ram::1:oidc-provider/ack-rrsa-c123",
				RRSAEndpoint:    fmt.Sprintf("https://%s:%s", serviceHost, port),
			}
			_, err = newRRSAExchanger(clientset, "cn-beijing", "1857989822569166", "c123", caFile, tt.duration).exchange(context.Background(), "vol", "/target", opt, secrets, nil)
			require.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

// fakeRefreshMounter is a NAS mounter whose mount broker support is configurable,
// standing in for the three cases a node can be in: no broker, a broker too old
// to refresh, and one that can.
type fakeRefreshMounter struct {
	mounter.Mounter
	can        bool
	canErr     error
	refreshErr error
	installed  map[string]*mounter.RefreshOperation
}

func (m *fakeRefreshMounter) CanRefresh(context.Context) (bool, error) {
	return m.can, m.canErr
}

func (m *fakeRefreshMounter) Refresh(_ context.Context, op *mounter.RefreshOperation) error {
	if m.refreshErr != nil {
		return m.refreshErr
	}
	if m.installed == nil {
		m.installed = map[string]*mounter.RefreshOperation{}
	}
	m.installed[op.Target] = op
	return nil
}

// TestCheckCredentialInstallSupport pins where a node that cannot rotate
// credentials is caught: an rrsa volume is refused before a mount exists and
// before a Pod token is spent, rather than one credential lifetime later as EACCES
// inside the Pod.
func TestCheckCredentialInstallSupport(t *testing.T) {
	tests := []struct {
		name    string
		mounter mounter.Mounter
		socket  string
		errLike string
	}{
		{
			// Connector or agent mode. NasMounter implements the refresher interface
			// either way, so only the configuration tells them apart.
			name:    "NAS is not using the mount broker",
			mounter: &fakeRefreshMounter{},
			errLike: "AlinasMountProxy",
		},
		{
			name:    "mounter cannot install at all",
			mounter: &recordingMounter{},
			socket:  "/run/cnfs/alinas-mounter.sock",
			errLike: "cannot install a credential",
		},
		{
			name:    "broker too old to know the RPC",
			mounter: &fakeRefreshMounter{can: false},
			socket:  "/run/cnfs/alinas-mounter.sock",
			errLike: "upgrade mount-proxy-server",
		},
		{
			name:    "broker unreachable",
			mounter: &fakeRefreshMounter{canErr: errors.New("dial unix: no such file")},
			socket:  "/run/cnfs/alinas-mounter.sock",
			errLike: "no such file",
		},
		{
			name:    "broker supports refresh",
			mounter: &fakeRefreshMounter{can: true},
			socket:  "/run/cnfs/alinas-mounter.sock",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := &nodeServer{mounter: tt.mounter, config: &internal.NodeConfig{MountProxySocket: tt.socket}}
			err := ns.checkCredentialInstallSupport(t.Context())
			if tt.errLike == "" {
				assert.NoError(t, err)
				return
			}
			require.Error(t, err)
			assert.Contains(t, err.Error(), tt.errLike)

			// An rrsa volume turns it into a publish failure that names the fix.
			ns.rrsa = newRRSAExchanger(fake.NewSimpleClientset(), "cn-beijing", "1857989822569166", "c123", "", 0)
			gateErr := ns.prepareRRSACredentials(t.Context(), "vol", "/target",
				&Options{RoleName: "test-role"}, nil, nil)
			require.Error(t, gateErr)
			assert.Equal(t, codes.FailedPrecondition, status.Code(gateErr))
			assert.Contains(t, gateErr.Error(), tt.errLike)
		})
	}
}

// TestPrepareRRSACredentialsRejectsLosetup: a losetup volume mounts the NAS share
// elsewhere and exposes a loop device at the target, so the broker owns neither
// and no credential could ever be installed. Saying so beats failing every
// republish for the life of the mount.
func TestPrepareRRSACredentialsRejectsLosetup(t *testing.T) {
	ns := &nodeServer{mounter: &fakeRefreshMounter{can: true}, config: &internal.NodeConfig{}}
	opt := &Options{RoleName: "test-role", MountType: LosetupType}
	err := ns.prepareRRSACredentials(t.Context(), "vol", "/target", opt, nil, nil)
	require.Error(t, err)
	assert.Equal(t, codes.InvalidArgument, status.Code(err))
	assert.Contains(t, err.Error(), "losetup")
}

// TestOptionsNeverCarryCredentialsThroughJSON pins the reason the embedded
// credential is tagged "-": Options is JSON-shaped for volume attributes, and the
// STS field names are the ones a PV author could otherwise set (or a log could
// otherwise print).
func TestOptionsNeverCarryCredentialsThroughJSON(t *testing.T) {
	var opt Options
	require.NoError(t, json.Unmarshal([]byte(
		`{"authType":"rrsa","AccessKeyId":"injected","AccessKeySecret":"injected","SecurityToken":"injected"}`), &opt))
	assert.Equal(t, AuthTypeRRSA, opt.AuthType)
	assert.Empty(t, opt.AkID)
	assert.Empty(t, opt.AkSecret)
	assert.Empty(t, opt.SecurityToken)

	opt.stsCredential = stsCredential{AkID: "ak", AkSecret: "sk", SecurityToken: "token"}
	out, err := json.Marshal(opt)
	require.NoError(t, err)
	for _, secret := range []string{"ak", "sk", "token"} {
		assert.NotContains(t, string(out), `"`+secret+`"`)
	}
}
