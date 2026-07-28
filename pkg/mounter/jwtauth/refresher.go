package jwtauth

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"

	"k8s.io/klog/v2"
)

const (
	defaultRefreshMargin = 5 * time.Minute
	minSleepDuration     = 30 * time.Second
	maxRetryBackoff      = 30 * time.Second
	initialRetryBackoff  = 1 * time.Second
	maxFetchRetries      = 5
	stopWaitTimeout      = 15 * time.Second
)

// Refresher fetches scoped STS credentials for a jwtauth mount and keeps them
// fresh for the lifetime of the mount, delivering each credential through the
// configured CredentialSink.
type Refresher struct {
	opts          Opts
	sink          CredentialSink
	refreshMargin time.Duration

	mu     sync.Mutex
	stopCh chan struct{}
	done   chan struct{}
	client *http.Client
}

func NewRefresher(opts Opts, sink CredentialSink) *Refresher {
	return &Refresher{
		opts:          opts,
		sink:          sink,
		refreshMargin: defaultRefreshMargin,
		stopCh:        make(chan struct{}),
		done:          make(chan struct{}),
	}
}

// Start performs the initial credential fetch synchronously, applies it via
// the sink and, on success, launches the background refresh loop. The
// provided context bounds only the initial fetch; the loop is stopped via
// Stop.
func (r *Refresher) Start(ctx context.Context) error {
	client, err := buildHTTPClient(r.opts.CAFile)
	if err != nil {
		return fmt.Errorf("build http client: %w", err)
	}
	r.client = client

	cred, err := r.fetchCredentials(ctx)
	if err != nil {
		return fmt.Errorf("initial credential fetch: %w", err)
	}
	if err := r.sink.Apply(cred); err != nil {
		return fmt.Errorf("apply initial credentials: %w", err)
	}

	go r.refreshLoop(cred)
	return nil
}

// StartWith launches the background refresh loop using an already fetched
// credential, without applying it through the sink first. Intended for
// consumers that obtained the initial credential via FetchSTSToken and
// delivered it themselves (e.g. injected into mount options); the sink is
// only invoked on subsequent rotations.
func (r *Refresher) StartWith(cred *STSToken) error {
	client, err := buildHTTPClient(r.opts.CAFile)
	if err != nil {
		return fmt.Errorf("build http client: %w", err)
	}
	r.client = client

	go r.refreshLoop(cred)
	return nil
}

// Stop signals the refresh loop to stop and waits for it to exit, bounded by
// stopWaitTimeout so a hung HTTP call cannot stall driver Terminate.
func (r *Refresher) Stop() {
	r.mu.Lock()
	select {
	case <-r.stopCh:
	default:
		close(r.stopCh)
	}
	r.mu.Unlock()

	select {
	case <-r.done:
	case <-time.After(stopWaitTimeout):
		klog.Warningf("jwtauth refresher: Stop timed out after %v, goroutine may still be running", stopWaitTimeout)
	}
}

// Cleanup delegates to the sink. It should be called after Stop once the
// consuming mount has gone away.
func (r *Refresher) Cleanup() {
	r.sink.Cleanup()
}

func (r *Refresher) fetchCredentials(ctx context.Context) (*STSToken, error) {
	return exchangeSTSToken(ctx, r.client, r.opts)
}

func (r *Refresher) refreshLoop(lastCred *STSToken) {
	defer close(r.done)
	for {
		sleepDuration := r.calcSleepDuration(lastCred.Expiration)
		klog.V(4).Infof("jwtauth refresher: next refresh in %v", sleepDuration)

		select {
		case <-r.stopCh:
			return
		case <-time.After(sleepDuration):
		}

		cred, err := r.fetchWithRetry()
		if err != nil {
			klog.Errorf("jwtauth refresher: fetch failed after retries: %v", err)
			continue
		}
		if err := r.sink.Apply(cred); err != nil {
			klog.Errorf("jwtauth refresher: apply credentials failed: %v", err)
			continue
		}
		lastCred = cred
		klog.V(2).Infof("jwtauth refresher: credentials refreshed, expires %s", cred.Expiration)
	}
}

func (r *Refresher) fetchWithRetry() (*STSToken, error) {
	backoff := initialRetryBackoff
	var lastErr error
	for i := 0; i < maxFetchRetries; i++ {
		select {
		case <-r.stopCh:
			return nil, fmt.Errorf("stopped")
		default:
		}

		ctx, cancel := context.WithTimeout(context.Background(), httpTimeout)
		cred, err := r.fetchCredentials(ctx)
		cancel()
		if err == nil {
			return cred, nil
		}
		lastErr = err
		klog.Warningf("jwtauth refresher: fetch attempt %d failed: %v", i+1, err)

		select {
		case <-r.stopCh:
			return nil, fmt.Errorf("stopped")
		case <-time.After(backoff):
		}
		backoff *= 2
		if backoff > maxRetryBackoff {
			backoff = maxRetryBackoff
		}
	}
	return nil, fmt.Errorf("all retries exhausted: %w", lastErr)
}

func (r *Refresher) calcSleepDuration(expiration string) time.Duration {
	expTime, err := time.Parse(time.RFC3339, expiration)
	if err != nil {
		klog.Warningf("jwtauth refresher: parse expiration %q failed, using margin as sleep: %v", expiration, err)
		return r.refreshMargin
	}
	until := time.Until(expTime) - r.refreshMargin
	if until < minSleepDuration {
		return minSleepDuration
	}
	return until
}
