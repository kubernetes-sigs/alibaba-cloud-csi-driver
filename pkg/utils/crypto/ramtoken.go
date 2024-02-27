package crypto

import (
	"encoding/base64"
	"errors"
)

type RamToken struct {
	AccessKeyId     string `json:"access.key.id"`
	AccessKeySecret string `json:"access.key.secret"`
	SecurityToken   string `json:"security.token"`
	Expiration      string `json:"expiration"`
	Keyring         string `json:"keyring"`
	ClusterType     string `json:"cluster.type"`
}

func RamTokenDecrypt(token *RamToken) (*RamToken, error) {
	if token == nil || token.AccessKeyId == "" || token.AccessKeySecret == "" || token.Keyring == "" {
		return nil, errors.New("invalid token")
	}
	rk := token.Keyring
	var err error
	newToken := &RamToken{
		Keyring:     token.Keyring,
		ClusterType: token.ClusterType,
		Expiration:  token.Expiration,
	}

	// accessKeyId
	encryptedAccessKeyId, err := base64.StdEncoding.DecodeString(token.AccessKeyId)
	if err != nil {
		return nil, err
	}
	accessKeyIdData, err := Decrypt(encryptedAccessKeyId, []byte(rk))
	if err != nil {
		return nil, err
	}
	newToken.AccessKeyId = string(accessKeyIdData)

	// accessKeySecret
	encryptedAccessKeySecret, err := base64.StdEncoding.DecodeString(token.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	accessKeySecretData, err := Decrypt(encryptedAccessKeySecret, []byte(rk))
	if err != nil {
		return nil, err
	}
	newToken.AccessKeySecret = string(accessKeySecretData)

	// securityToken
	encryptedSecurityToken, err := base64.StdEncoding.DecodeString(token.SecurityToken)
	if err != nil {
		return nil, err
	}
	securityTokenData, err := Decrypt(encryptedSecurityToken, []byte(rk))
	if err != nil {
		return nil, err
	}
	newToken.SecurityToken = string(securityTokenData)

	return newToken, nil
}
