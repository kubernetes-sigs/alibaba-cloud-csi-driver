//go:build !windows

package nas

import (
	"bytes"
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	mounterutils "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils"
	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/ttlcache"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/klog/v2"
)

// AuthTypeRRSA selects per-Pod STS credentials obtained from sts AssumeRoleWithOIDC
const AuthTypeRRSA = "rrsa"

const (
	// defaultTokenAudience is what RRSA uses when the token goes to STS itself. A
	// deployment that sends it to an oidc-proxy instead sets rrsaAudience to
	// whatever that proxy verifies, since the audience names the endpoint allowed
	// to consume the token. Every audience used has to be listed in the nasplugin
	// CSIDriver's tokenRequests, or kubelet never mints that token.
	defaultTokenAudience = "sts.aliyuncs.com"

	// serviceAccountTokensKey holds the tokens kubelet requested via
	// CSIDriver.tokenRequests, in the publish secrets when the CSIDriver sets
	// serviceAccountTokenInSecrets (Kubernetes 1.35+) and in the volume context
	// otherwise. Both are read, so the delivery stays the cluster's choice. A
	// missing key is a configuration error: without it nothing ties the credential
	// to a Pod.
	serviceAccountTokensKey = "csi.storage.k8s.io/serviceAccount.tokens"

	exchangeTimeout = 30 * time.Second

	// clusterIPTTL bounds how long a resolved ClusterIP is reused. A Service IP
	// is stable for the life of the Service, so this only matters after someone
	// recreates it.
	clusterIPTTL = 10 * time.Minute

	// refreshMargin is the latest a credential is renewed before it expires,
	// leaving room for the exchange to fail and be retried on later republishes.
	refreshMargin = 20 * time.Minute
)

type stsCredential struct {
	AkID          string    `json:"AccessKeyId"`
	AkSecret      string    `json:"AccessKeySecret"`
	SecurityToken string    `json:"SecurityToken"`
	Expiration    time.Time `json:"Expiration"`
}

type rrsaCredential struct {
	stsCredential
	// refreshAt is when the credential should be replaced, fixed at exchange
	// time so a republish only has to compare the clock to it. Whether the mount
	// is already running on this credential is not tracked here; see
	// installedCredentials.
	refreshAt time.Time
}

// refreshDeadline returns when a credential obtained now and expiring at
// expiration should be replaced:
//
// long lived: refresh 20m before expiration (ref. jwtauth, ossfs),
// NAS server close connection 15m before expiration, then client reconnects.
// we go before it for the normal case.
//
// short lived: every time kubelet call us (~1 min), with 30s cooldown just in case.
//
// NAS return EACCES when credential expires on established connection,
// or drop connection immediately when reconnect with expired credential, leave mountpoint hang.
// Invariant: keep on-disk credential present to alinas-utils at least valid for 10m at any moment.
// NAS only refresh credential once when reconnect, and only auto reconnect after at least 10m after establish
func refreshDeadline(obtained, expiration time.Time) time.Time {
	refreshAt := expiration.Add(-refreshMargin)
	if cooldown := obtained.Add(30 * time.Second); cooldown.After(refreshAt) {
		refreshAt = cooldown
	}
	return refreshAt
}

// rrsaExchanger turns Pod ServiceAccount tokens into scoped STS credentials,
// and remembers the credential of each mount so a republish
// can tell whether it is time to rotate.
type rrsaExchanger struct {
	kubeClient kubernetes.Interface
	// duration is the lifetime to ask the endpoint for; zero lets it decide.
	duration time.Duration
	// region builds the default STS endpoint when a volume names none.
	region string
	// accountID and clusterID complete a volume that names only a roleName.
	accountID string
	clusterID string

	clusterIPs *ttlcache.TTLCache[string, string]
	dial       func(ctx context.Context, network, addr string) (net.Conn, error)
	client     *http.Client

	creds sync.Map // map[mount target]*rrsaCredential
}

func newRRSAExchanger(kubeClient kubernetes.Interface, region, accountID, clusterID, caFile string, duration time.Duration) *rrsaExchanger {
	e := &rrsaExchanger{
		kubeClient: kubeClient,
		region:     region,
		accountID:  accountID,
		clusterID:  clusterID,
		duration:   duration,
		clusterIPs: ttlcache.NewTTLCache[string, string](clusterIPTTL),
		dial:       (&net.Dialer{}).DialContext,
	}
	e.client = &http.Client{Transport: &rrsaTransport{
		caFile: caFile,
		dial:   e.dialContext,
	}}
	return e
}

// credentialFor returns the credential target should be running on, exchanging a
// new one only when there is none or the current one is due. Reusing the current
// one is what keeps a once-a-minute republish free of AssumeRoleWithOIDC quota;
// whether it has reached the mount is the caller's business, see installed.
func (e *rrsaExchanger) credentialFor(ctx context.Context, volumeID, target string, opt *Options, volumeContext, secrets map[string]string) (*rrsaCredential, error) {
	if cred := e.load(target); cred != nil && time.Now().Before(cred.refreshAt) {
		return cred, nil
	}

	cred, err := e.exchange(ctx, volumeID, target, opt, volumeContext, secrets)
	if err != nil {
		return nil, err
	}
	e.creds.Store(target, cred)
	return cred, nil
}

// load returns the credential remembered for target, or nil.
func (e *rrsaExchanger) load(target string) *rrsaCredential {
	v, _ := e.creds.Load(target)
	cred, _ := v.(*rrsaCredential)
	return cred
}

// forget drops the credential remembered for target, so an unmounted volume does
// not pin state and a later mount of the same path exchanges its own.
func (e *rrsaExchanger) forget(target string) {
	e.creds.Delete(target)
}

func (e *rrsaExchanger) exchange(ctx context.Context, volumeID, target string, opt *Options, tokenSources ...map[string]string) (*rrsaCredential, error) {
	// The same attributes OSS accepts, resolved by the same rule: a bare roleName
	// is completed from this cluster, or both ARNs are given.
	arns, err := mounterutils.ResolveRRSAArns(opt.RoleName, opt.RoleArn, opt.OIDCProviderArn, e.accountID, e.clusterID)
	if err != nil {
		return nil, err
	}
	audience := opt.RRSAAudience
	if audience == "" {
		audience = defaultTokenAudience
	}
	saToken, err := podServiceAccountToken(audience, tokenSources[0], tokenSources[1])
	if err != nil {
		return nil, err
	}
	// No endpoint means plain RRSA: the token goes to STS, which grants the role
	// as its own policy allows, with no per-Pod narrowing.
	endpoint := opt.RRSAEndpoint
	if endpoint == "" {
		endpoint, err = utils.GetSTSEndpoint(e.region)
		if err != nil {
			return nil, fmt.Errorf("sts endpoint: %w", err)
		}
	}
	if err := checkEndpoint(endpoint); err != nil {
		return nil, err
	}

	form := url.Values{
		"Action":          {"AssumeRoleWithOIDC"},
		"Version":         {"2015-04-01"},
		"Timestamp":       {time.Now().UTC().Format(time.RFC3339)},
		"Format":          {"JSON"},
		"OIDCProviderArn": {arns.OidcProviderArn},
		"RoleArn":         {arns.RoleArn},
		"OIDCToken":       {saToken},
		// Identifies the mount in STS audit logs, the way the OSS driver does.
		"RoleSessionName": {mounterutils.GetRoleSessionName(volumeID, target, "nas")},
	}
	if e.duration > 0 {
		form.Set("DurationSeconds", strconv.FormatInt(int64(e.duration.Seconds()), 10))
	}
	ctx, cancel := context.WithTimeout(ctx, exchangeTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("build credential request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := e.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	// The body carries no credential on the error path, but it does on the happy
	// path, so it is never logged as a whole.
	body, err := io.ReadAll(io.LimitReader(resp.Body, 1<<20))
	if err != nil {
		return nil, fmt.Errorf("read AssumeRoleWithOIDC response: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		var e struct{ Code, Message string }
		_ = json.Unmarshal(body, &e)
		if e.Code != "" {
			return nil, fmt.Errorf("AssumeRoleWithOIDC rejected the request: %s: %s", e.Code, e.Message)
		}
		return nil, fmt.Errorf("AssumeRoleWithOIDC returned %d", resp.StatusCode)
	}

	var parsed struct {
		Credentials stsCredential `json:"Credentials"`
	}
	if err := json.Unmarshal(body, &parsed); err != nil {
		return nil, fmt.Errorf("parse AssumeRoleWithOIDC response: %w", err)
	}
	c := parsed.Credentials
	if c.AkID == "" || c.AkSecret == "" || c.SecurityToken == "" || c.Expiration.IsZero() {
		return nil, fmt.Errorf("AssumeRoleWithOIDC returned an incomplete credential")
	}
	klog.V(2).InfoS("exchanged Pod token for NAS credentials", "roleArn", arns.RoleArn,
		"accessKeyId", c.AkID, "expiration", c.Expiration)
	obtained := time.Now()
	return &rrsaCredential{
		stsCredential: parsed.Credentials,
		refreshAt:     refreshDeadline(obtained, c.Expiration),
	}, nil
}

// rrsaVolumeAttributeKeys are what an authType=rrsa volume needs on the node.
// They are the same for every volume of a StorageClass, so dynamic provisioning
// can carry them over, unlike the agent-identity ones, which name a sandbox.
var rrsaVolumeAttributeKeys = []string{"authType", "roleName", "roleArn", "oidcProviderArn", "rrsaEndpoint", "rrsaAudience"}

// rrsaVolumeAttributes returns the attributes a StorageClass with authType=rrsa
// puts on its volumes, or nil for any other authType. It validates them the way
// the node will, so a misconfigured class fails at provisioning, before any cloud
// resource is created, rather than on every mount.
//
// A roleName is carried over as such rather than resolved here: it means the
// provider of whichever cluster mounts the volume, which keeps a restored PV
// working in another cluster.
func rrsaVolumeAttributes(parameters map[string]string) (map[string]string, error) {
	if !strings.EqualFold(parameters["authType"], AuthTypeRRSA) {
		return nil, nil
	}
	if err := mounterutils.CheckRRSAArns(parameters["roleName"], parameters["roleArn"], parameters["oidcProviderArn"]); err != nil {
		return nil, err
	}
	if endpoint := parameters["rrsaEndpoint"]; endpoint != "" {
		if err := checkEndpoint(endpoint); err != nil {
			return nil, err
		}
	}

	attrs := map[string]string{}
	for _, k := range rrsaVolumeAttributeKeys {
		if v := parameters[k]; v != "" {
			attrs[k] = v
		}
	}
	return attrs, nil
}

// checkEndpoint refuses an endpoint the Pod token must not be sent to.
func checkEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("parse rrsaEndpoint %q: %w", endpoint, err)
	}
	if u.Scheme != "https" {
		return fmt.Errorf("rrsaEndpoint %q must use https: the Pod token is a bearer credential", endpoint)
	}
	return nil
}

// dialContext dials a Service host at its ClusterIP, read from the API server:
// csi-plugin runs with hostNetwork and the host's resolv.conf, so it cannot
// resolve cluster DNS. Only the dial is redirected; the request URL keeps the
// Service name, so TLS still verifies the certificate against it. Any other host
// is dialed as is.
func (e *rrsaExchanger) dialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	if name, namespace, ok := splitServiceHost(host); ok {
		ip, err := e.clusterIPs.Get(ctx, host, func() (string, error) {
			svc, err := e.kubeClient.CoreV1().Services(namespace).Get(ctx, name, metav1.GetOptions{})
			if err != nil {
				return "", fmt.Errorf("resolve Service %s/%s for rrsaEndpoint: %w", namespace, name, err)
			}
			if svc.Spec.ClusterIP == "" || svc.Spec.ClusterIP == "None" {
				return "", fmt.Errorf("service %s/%s has no ClusterIP", namespace, name)
			}
			return svc.Spec.ClusterIP, nil
		})
		if err != nil {
			return nil, err
		}
		addr = net.JoinHostPort(ip, port)
	}
	return e.dial(ctx, network, addr)
}

// splitServiceHost parses <name>.<namespace>.svc[.cluster.local].
func splitServiceHost(host string) (name, namespace string, ok bool) {
	parts := strings.Split(host, ".")
	if len(parts) < 3 || parts[2] != "svc" {
		return "", "", false
	}
	return parts[0], parts[1], true
}

// caRecheckInterval is how long a loaded CA is trusted before its file is read
// again, the same as client-go uses for its own CA file.
const caRecheckInterval = 5 * time.Minute

// rrsaTransport is the one http.Transport all exchanges share, so connections
// are pooled and bounded instead of each exchange leaving one behind.
//
// Its RootCAs follow caFile the way client-go follows its CA file: the file is
// read again at most every caRecheckInterval, and only changed content swaps in a
// new transport, after which the old one drops its idle connections. A file that
// turns unreadable or invalid keeps the last good CA.
//
// The CA is added to the system roots rather than replacing them, so a configured
// CA does not break an endpoint with a publicly trusted certificate, such as STS.
// That does not weaken the proxy case, since no public CA issues for a .svc name.
type rrsaTransport struct {
	caFile string
	dial   func(ctx context.Context, network, addr string) (net.Conn, error)

	mu        sync.Mutex
	current   *http.Transport
	caPEM     []byte
	checkedAt time.Time
}

func (t *rrsaTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	tr, err := t.get(req.Context())
	if err != nil {
		return nil, err
	}
	return tr.RoundTrip(req)
}

func (t *rrsaTransport) get(ctx context.Context) (*http.Transport, error) {
	t.mu.Lock()
	defer t.mu.Unlock()
	if t.current != nil && time.Since(t.checkedAt) < caRecheckInterval {
		return t.current, nil
	}

	t.checkedAt = time.Now()
	caPEM, pool, err := t.loadCA()
	if err != nil {
		if t.current == nil {
			return nil, err
		}
		klog.FromContext(ctx).Error(err, "Keeping the previous CA for rrsa endpoints", "caFile", t.caFile)
		return t.current, nil
	}
	if t.current != nil && bytes.Equal(caPEM, t.caPEM) {
		return t.current, nil
	}

	old := t.current
	t.current = &http.Transport{
		DialContext:         t.dial,
		TLSClientConfig:     &tls.Config{RootCAs: pool},
		TLSHandshakeTimeout: 10 * time.Second,
		// Also closes what a swapped-out transport pools after CloseIdleConnections:
		// a request that picked it up before the swap re-enables its pool.
		IdleConnTimeout: 90 * time.Second,
	}
	t.caPEM = caPEM
	if old != nil {
		old.CloseIdleConnections()
	}
	return t.current, nil
}

// loadCA reads caFile, returning a nil pool (the system roots) when none is
// configured.
func (t *rrsaTransport) loadCA() ([]byte, *x509.CertPool, error) {
	if t.caFile == "" {
		return nil, nil, nil
	}
	caPEM, err := os.ReadFile(t.caFile)
	if err != nil {
		return nil, nil, fmt.Errorf("read CA %s: %w", t.caFile, err)
	}
	pool, err := x509.SystemCertPool()
	if err != nil {
		pool = x509.NewCertPool()
	}
	if !pool.AppendCertsFromPEM(caPEM) {
		return nil, nil, fmt.Errorf("CA %s contains no certificate", t.caFile)
	}
	return caPEM, pool, nil
}

// podServiceAccountToken returns the token kubelet requested for our audience.
// Requiring that audience keeps a misconfigured CSIDriver from falling back to
// some other token.
func podServiceAccountToken(audience string, volumeContext, secrets map[string]string) (string, error) {
	raw := volumeContext[serviceAccountTokensKey]
	if raw == "" {
		raw = secrets[serviceAccountTokensKey]
	}
	if raw == "" {
		return "", fmt.Errorf("no ServiceAccount token in the publish request: "+
			"CSIDriver %s must request audience %s in tokenRequests", driverName, audience)
	}
	var tokens map[string]struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal([]byte(raw), &tokens); err != nil {
		return "", fmt.Errorf("parse ServiceAccount tokens: %w", err)
	}
	token := tokens[audience].Token
	if token == "" {
		return "", fmt.Errorf("no ServiceAccount token for audience %s", audience)
	}
	return token, nil
}

// prepareRRSACredentials obtains the credential target should be running on and
// puts it where doMount passes credentials to the mount broker.
//
// Every failure here fails the publish: mount.alinas accepts a partial credential
// and then hangs against a server that rejects it.
func (ns *nodeServer) prepareRRSACredentials(ctx context.Context, volumeID, target string, opt *Options, volumeContext, secrets map[string]string) error {
	if opt.RoleName == "" && opt.RoleArn == "" {
		return status.Errorf(codes.InvalidArgument,
			"authType %s requires roleName, or roleArn with oidcProviderArn, in volumeAttributes", AuthTypeRRSA)
	}
	// A losetup volume mounts the NAS share somewhere else and exposes a loop
	// device at target, so the broker owns neither and no credential could ever be
	// installed on it.
	if opt.MountType == LosetupType {
		return status.Errorf(codes.InvalidArgument, "authType %s does not support losetup volumes", AuthTypeRRSA)
	}
	if ns.rrsa == nil {
		return status.Error(codes.FailedPrecondition, "no kube client: cannot exchange the Pod token for credentials")
	}

	cred, err := ns.rrsa.credentialFor(ctx, volumeID, target, opt, volumeContext, secrets)
	if err != nil {
		return status.Errorf(codes.Internal, "get NAS credentials: %v", err)
	}
	opt.stsCredential = cred.stsCredential
	return nil
}
