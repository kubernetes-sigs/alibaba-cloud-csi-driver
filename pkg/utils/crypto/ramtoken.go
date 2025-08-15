package crypto

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type RamToken struct {
	AccessKeyId     string `json:"access.key.id"`
	AccessKeySecret string `json:"access.key.secret"`
	SecurityToken   string `json:"security.token"`
	Expiration      string `json:"expiration"`
	Keyring         string `json:"keyring"`
	ClusterType     string `json:"cluster.type"`
}

func RamTokenParse(encodedToken io.Reader) (*RamToken, error) {
	var encryptedToken RamToken
	err := json.NewDecoder(encodedToken).Decode(&encryptedToken)
	if err != nil {
		return nil, err
	}
	return RamTokenDecrypt(&encryptedToken)
}

type tokenDecodeContext struct {
	*base64.Encoding
	Keyring []byte
}

func (ctx *tokenDecodeContext) Decode(s string) (string, error) {
	decoded, err := ctx.Encoding.DecodeString(s)
	if err != nil {
		return "", err
	}
	decrypted, err := Decrypt(decoded, ctx.Keyring)
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func RamTokenDecrypt(token *RamToken) (*RamToken, error) {
	if token == nil || token.AccessKeyId == "" || token.AccessKeySecret == "" || token.Keyring == "" {
		return nil, errors.New("invalid token")
	}
	var err error
	newToken := &RamToken{
		Keyring:     token.Keyring,
		ClusterType: token.ClusterType,
		Expiration:  token.Expiration,
	}

	decodeCtx := tokenDecodeContext{
		Encoding: base64.StdEncoding,
		Keyring:  []byte(token.Keyring),
	}

	newToken.AccessKeyId, err = decodeCtx.Decode(token.AccessKeyId)
	if err != nil {
		return nil, fmt.Errorf("failed to decode access key id: %w", err)
	}

	newToken.AccessKeySecret, err = decodeCtx.Decode(token.AccessKeySecret)
	if err != nil {
		return nil, fmt.Errorf("failed to decode access key secret: %w", err)
	}

	newToken.SecurityToken, err = decodeCtx.Decode(token.SecurityToken)
	if err != nil {
		return nil, fmt.Errorf("failed to decode security token: %w", err)
	}

	return newToken, nil
}
