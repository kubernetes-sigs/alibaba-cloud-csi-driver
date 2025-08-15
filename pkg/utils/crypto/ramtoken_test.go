package crypto

import (
	"bytes"
	"testing"
	"time"

	cryptotest "github.com/kubernetes-sigs/alibaba-cloud-csi-driver/pkg/utils/crypto/testing"
	"github.com/stretchr/testify/assert"
)

func assertFakeToken(t *testing.T, tokens *RamToken) {
	assert.Equal(t, cryptotest.FakeAccessKeyId, tokens.AccessKeyId)
	assert.Equal(t, cryptotest.FakeAccessKeySecret, tokens.AccessKeySecret)
	assert.Equal(t, cryptotest.FakeSecurityToken, tokens.SecurityToken)
	exp, err := time.Parse(time.RFC3339, tokens.Expiration)
	assert.NoError(t, err)
	assert.Equal(t, cryptotest.Expiration, exp)
}

func TestParseToken(t *testing.T) {
	f := bytes.NewReader([]byte(cryptotest.FakeEncryptedToken))

	tokens, err := RamTokenParse(f)
	assert.NoError(t, err)

	assertFakeToken(t, tokens)
}
