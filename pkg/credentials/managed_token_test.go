package credentials

import (
	"errors"
	"os"
	"testing"
	"time"

	cryptotest "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/crypto/testing"
	"github.com/stretchr/testify/assert"
	"k8s.io/utils/clock"
	clocktest "k8s.io/utils/clock/testing"
)

func TestManagedTokenCredential(t *testing.T) {
	tokenFile := t.TempDir() + "/token-config"
	err := cryptotest.WriteFakeToken(tokenFile)
	assert.NoError(t, err)

	credential, err := NewManagedTokenProvider(clock.RealClock{}, tokenFile)
	assert.NoError(t, err)

	cred, err := credential.GetCredentials()
	assert.NoError(t, err)

	assert.Equal(t, cryptotest.FakeAccessKeyId, cred.AccessKeyId)
	assert.Equal(t, cryptotest.FakeAccessKeySecret, cred.AccessKeySecret)
	assert.Equal(t, cryptotest.FakeSecurityToken, cred.SecurityToken)
}

func TestTokenNotExist(t *testing.T) {
	_, err := NewManagedTokenProvider(clock.RealClock{}, "/not/exist")
	if !errors.Is(err, os.ErrNotExist) {
		t.Errorf("expected file not exist error, got %v", err)
	}
}

var tokenLoadTime = cryptotest.Expiration.Add(-30 * time.Minute)

func TestTokenAutoReload(t *testing.T) {
	cases := []struct {
		name         string
		loadTime     time.Time
		stepTime     time.Duration
		expectReload bool
	}{
		{
			name:         "cache in 5 minutes",
			loadTime:     tokenLoadTime,
			stepTime:     4*time.Minute + 59*time.Second,
			expectReload: false,
		},
		{
			name:         "cache expired",
			loadTime:     tokenLoadTime,
			stepTime:     5*time.Minute + 1*time.Second,
			expectReload: true,
		},
		{
			name:         "token near expire",
			loadTime:     cryptotest.Expiration.Add(-11 * time.Minute),
			stepTime:     2 * time.Minute,
			expectReload: true,
		},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			clk := clocktest.NewFakeClock(c.loadTime)
			tokenFile := t.TempDir() + "/token-config"
			err := cryptotest.WriteFakeToken(tokenFile)
			assert.NoError(t, err)

			cred, err := NewManagedTokenProvider(clk, tokenFile)
			assert.NoError(t, err)

			clk.Step(c.stepTime)
			os.Remove(tokenFile)
			_, err = cred.GetCredentials()
			if c.expectReload {
				if !errors.Is(err, os.ErrNotExist) {
					t.Errorf("expected file not exist error, got %v", err)
				}
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
