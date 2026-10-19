package jwtauth

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/mounter/utils/agentidentity"
	"k8s.io/klog/v2"
)

const (
	// minSleep floors the interval between two refreshes so an already-expired
	// or near-expiry credential cannot turn the loop into a busy loop.
	minSleep            = 30 * time.Second
	maxRetryBackoff     = 30 * time.Second
	initialRetryBackoff = 1 * time.Second
	maxFetchRetries     = 5
	// Recheck the absolute refresh target after a suspend/resume or clock change.
	clockCheckInterval = 5 * time.Second
)

// Refresher fetches scoped STS credentials for a jwtauth mount and keeps them
// fresh for the lifetime of the mount, delivering each credential through the
// configured CredentialSink.
type Refresher struct {
	opts Opts
	sink CredentialSink
	// refreshMargin is how long before expiry the credential is renewed. It is
	// resolved from the environment per refresher, so an operator can tune it
	// without a rebuild.
	refreshMargin time.Duration
	// wallNow allows tests to simulate clock changes without changing system time.
	wallNow func() time.Time

	client *http.Client
}

func NewRefresher(opts Opts, sink CredentialSink) (*Refresher, error) {
	client, err := buildHTTPClient(opts.CAFile)
	if err != nil {
		return nil, fmt.Errorf("build http client: %w", err)
	}
	return &Refresher{
		opts:          opts,
		sink:          sink,
		refreshMargin: agentidentity.GetTokenRefreshMargin(),
		wallNow:       time.Now,
		client:        client,
	}, nil
}

// FetchCredentials performs a single exchange of the jwtauth token for an STS
// credential, without delivering it to the sink.
func (r *Refresher) FetchCredentials(ctx context.Context) (*STSToken, error) {
	return exchangeSTSToken(ctx, r.client, r.opts)
}

// refreshLoop keeps the credential fresh until ctx is cancelled. lastCred has
// already been delivered by the caller (e.g. injected into mount options), so
// the sink is only invoked on subsequent rotations.
func (r *Refresher) refreshLoop(ctx context.Context, lastCred *STSToken) {
	notifyBackoff := initialRetryBackoff
	for {
		target := r.calcRefreshTarget(lastCred.Expiration)
		klog.V(4).Infof("jwtauth refresher: next refresh at %v", target)

		if !r.waitUntilDue(ctx, target) {
			return
		}

		cred, err := r.fetchWithRetry(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return
			}
			klog.Errorf("jwtauth refresher: fetch failed after retries: %v", err)
			continue
		}
		if err := r.sink.Deliver(cred); err != nil {
			klog.Errorf("jwtauth refresher: deliver credentials failed: %v", err)
			continue
		}
		lastCred = cred
		for err := r.sink.Notify(); err != nil; err = r.sink.Notify() {
			klog.Errorf("jwtauth refresher: notify failed (credential is on disk): %v", err)
			select {
			case <-ctx.Done():
				return
			case <-time.After(notifyBackoff):
			}
			notifyBackoff = min(notifyBackoff*2, maxRetryBackoff)
		}
		notifyBackoff = initialRetryBackoff
		klog.V(2).Infof("jwtauth refresher: credentials refreshed, expires %s", cred.Expiration)
	}
}

func (r *Refresher) fetchWithRetry(ctx context.Context) (*STSToken, error) {
	backoff := initialRetryBackoff
	var lastErr error
	for i := 0; i < maxFetchRetries; i++ {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		reqCtx, cancel := context.WithTimeout(ctx, httpTimeout)
		cred, err := r.FetchCredentials(reqCtx)
		cancel()
		if err == nil {
			return cred, nil
		}
		lastErr = err
		klog.Warningf("jwtauth refresher: fetch attempt %d failed: %v", i+1, err)

		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(backoff):
		}
		backoff *= 2
		if backoff > maxRetryBackoff {
			backoff = maxRetryBackoff
		}
	}
	return nil, fmt.Errorf("all retries exhausted: %w", lastErr)
}

// waitUntilDue polls the wall-clock target prepared by refreshLoop so clock
// changes and sandbox resumes are reflected on the next check.
func (r *Refresher) waitUntilDue(ctx context.Context, target time.Time) bool {
	for {
		dur := target.Sub(r.wallNow())
		if dur <= 0 {
			return true
		}
		select {
		case <-ctx.Done():
			return false
		case <-time.After(min(dur, clockCheckInterval)):
		}
	}
}

func (r *Refresher) calcRefreshTarget(expiration string) time.Time {
	expTime, err := time.Parse(time.RFC3339, expiration)
	if err != nil {
		klog.Warningf("jwtauth refresher: parse expiration %q failed, refresh after %v: %v", expiration, r.refreshMargin, err)
		return r.wallNow().Round(0).Add(r.refreshMargin)
	}
	target := expTime.Add(-r.refreshMargin)
	floor := r.wallNow().Round(0).Add(minSleep)
	if target.Before(floor) {
		return floor
	}
	return target
}
