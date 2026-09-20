package nas

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
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

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter"
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
	// caFile holds the trust anchor for the proxy's serving certificate. It is
	// read per exchange rather than cached, so a rotated CA takes effect without
	// restarting the driver; exchanges are rare enough for this to be free.
	caFile string

	clusterIPs *ttlcache.TTLCache[string, string]

	creds sync.Map // map[mount target]*rrsaCredential
}

func newRRSAExchanger(kubeClient kubernetes.Interface, region, accountID, clusterID, caFile string, duration time.Duration) *rrsaExchanger {
	return &rrsaExchanger{
		kubeClient: kubeClient,
		region:     region,
		accountID:  accountID,
		clusterID:  clusterID,
		caFile:     caFile,
		duration:   duration,
		clusterIPs: ttlcache.NewTTLCache[string, string](clusterIPTTL),
	}
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
	address, serverName, err := e.resolveEndpoint(ctx, endpoint)
	if err != nil {
		return nil, err
	}
	client, err := e.httpClient(serverName)
	if err != nil {
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
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, address, strings.NewReader(form.Encode()))
	if err != nil {
		return nil, fmt.Errorf("build credential request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("call %s: %w", address, err)
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

// resolveEndpoint turns the configured proxy URL into an address this process can
// actually reach, plus the name its certificate must present.
//
// csi-plugin runs with hostNetwork and the host's resolv.conf, so it cannot
// resolve a Service DNS name. The ClusterIP is therefore looked up through the
// API server and dialed directly, while TLS still verifies the DNS name the
// certificate was issued for (see httpClient). A literal IP or an
// externally-resolvable host is passed through untouched.
func (e *rrsaExchanger) resolveEndpoint(ctx context.Context, endpoint string) (address, serverName string, err error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return "", "", fmt.Errorf("parse rrsaEndpoint %q: %w", endpoint, err)
	}
	if u.Scheme != "https" {
		return "", "", fmt.Errorf("rrsaEndpoint %q must use https: the Pod token is a bearer credential", endpoint)
	}
	host := u.Hostname()
	if net.ParseIP(host) != nil {
		return endpoint, host, nil
	}
	name, namespace, ok := splitServiceHost(host)
	if !ok {
		return endpoint, host, nil
	}

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
		return "", "", err
	}

	resolved := *u
	if port := u.Port(); port != "" {
		resolved.Host = net.JoinHostPort(ip, port)
	} else {
		resolved.Host = ip
	}
	return resolved.String(), host, nil
}

// splitServiceHost parses <name>.<namespace>.svc[.cluster.local].
func splitServiceHost(host string) (name, namespace string, ok bool) {
	parts := strings.Split(host, ".")
	if len(parts) < 3 || parts[2] != "svc" {
		return "", "", false
	}
	return parts[0], parts[1], true
}

// httpClient builds a client that verifies the endpoint against serverName, which
// is the DNS name from the endpoint rather than the address dialed. Pinning the
// name this way lets a proxy certificate stay tied to its Service name while we
// connect to its ClusterIP.
//
// The configured CA is added to the system roots rather than replacing them,
// because the same setting serves both endpoints: a proxy presenting a private
// certificate, and STS presenting a publicly trusted one. Keeping the public roots
// does not weaken the proxy case, since no public CA can issue for a .svc name.
func (e *rrsaExchanger) httpClient(serverName string) (*http.Client, error) {
	tlsConfig := &tls.Config{ServerName: serverName}
	if e.caFile != "" {
		caPEM, err := os.ReadFile(e.caFile)
		if err != nil {
			return nil, fmt.Errorf("read CA %s: %w", e.caFile, err)
		}
		pool, err := x509.SystemCertPool()
		if err != nil {
			pool = x509.NewCertPool()
		}
		if !pool.AppendCertsFromPEM(caPEM) {
			return nil, fmt.Errorf("CA %s contains no certificate", e.caFile)
		}
		tlsConfig.RootCAs = pool
	}
	return &http.Client{Transport: &http.Transport{TLSClientConfig: tlsConfig}}, nil
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
// and then hangs against a server that rejects it, and a volume whose credential
// can never be rotated is refused before it mounts rather than after it expires.
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
	// An rrsa credential expires, so a node that could never install the next one is
	// refused here rather than one credential lifetime later, from inside the Pod.
	if err := ns.checkCredentialInstallSupport(ctx); err != nil {
		return status.Errorf(codes.FailedPrecondition,
			"authType %s needs a node that can rotate credentials: %v", AuthTypeRRSA, err)
	}

	cred, err := ns.rrsa.credentialFor(ctx, volumeID, target, opt, volumeContext, secrets)
	if err != nil {
		return status.Errorf(codes.Internal, "get NAS credentials: %v", err)
	}
	opt.stsCredential = cred.stsCredential
	return nil
}

// checkCredentialInstallSupport reports why this node cannot install a credential
// on a live mount, or nil when it can.
//
// The mount broker is the only mounter that can do it, and it ships in its own
// image, so its support is asked over the socket rather than assumed from this
// process being built with the client half.
func (ns *nodeServer) checkCredentialInstallSupport(ctx context.Context) error {
	// NasMounter satisfies ProxyRefresher whichever mounter it wraps, so which mode
	// NAS is in has to be read from the configuration.
	if ns.config.MountProxySocket == "" {
		return errors.New("NAS is not using the mount broker: enable the AlinasMountProxy feature gate on csi-plugin")
	}
	refresher, ok := ns.mounter.(mounter.ProxyRefresher)
	if !ok {
		return errors.New("this mounter cannot install a credential on a live mount")
	}
	can, err := refresher.CanRefresh(ctx)
	if err != nil {
		return fmt.Errorf("ask the mount broker whether it can refresh credentials: %w", err)
	}
	if !can {
		return errors.New("the mount broker does not support credential refresh: upgrade mount-proxy-server")
	}
	return nil
}
