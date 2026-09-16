package cloud

import (
	"context"
	"net/http"
	"sync/atomic"
	"testing"
	"time"

	sdk "github.com/alibabacloud-go/nas-20170626/v4/client"
	oldcred "github.com/aliyun/credentials-go/credentials"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The SDK's credential interface has no context. Simulate a refresh that only
// completes after the cleanup deadline, without relying on a long sleep.
type blockedDeleteCredential struct {
	oldcred.Credential
	started chan struct{}
	release chan struct{}
}

func (c *blockedDeleteCredential) GetCredential() (*oldcred.CredentialModel, error) {
	close(c.started)
	<-c.release
	return c.Credential.GetCredential()
}

func TestDeleteAccesspointDoesNotSendAfterCredentialRefreshExceedsDeadline(t *testing.T) {
	var hits atomic.Int32
	client := newTimeoutWireNasClient(t, func(w http.ResponseWriter, r *http.Request) {
		hits.Add(1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"RequestId":"unexpected-late-delete"}`))
	})
	real := client.client.(*sdk.Client)
	credential := &blockedDeleteCredential{
		Credential: real.Credential,
		started:    make(chan struct{}),
		release:    make(chan struct{}),
	}
	real.Credential = credential
	defer close(credential.release)

	ctx, cancel := context.WithTimeout(t.Context(), 100*time.Millisecond)
	defer cancel()
	done := make(chan error, 1)
	go func() {
		done <- client.DeleteAccesspoint(ctx, wireTimeoutFileSystemID, wireTimeoutAccessPointID)
	}()
	select {
	case <-credential.started:
	case <-time.After(5 * time.Second):
		t.Fatal("SDK did not enter credential resolution")
	}
	<-ctx.Done()
	select {
	case err := <-done:
		t.Fatalf("SDK call returned before the credential provider finished: %v", err)
	default:
	}
	assert.Zero(t, hits.Load())

	credential.release <- struct{}{}
	select {
	case err := <-done:
		require.ErrorIs(t, err, context.DeadlineExceeded)
		assert.NotErrorIs(t, err, ErrRateLimiterWait, "the request passed the limiter before resolving credentials")
	case <-time.After(5 * time.Second):
		t.Fatal("SDK did not return after credential resolution completed")
	}
	assert.Zero(t, hits.Load(), "expired cleanup must not issue a late DeleteAccessPoint request")
}
